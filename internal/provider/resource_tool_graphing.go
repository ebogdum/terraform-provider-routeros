package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &ToolGraphingResource{}
	_ resource.ResourceWithImportState = &ToolGraphingResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type ToolGraphingResource struct {
	reg *client.Registry
}

type ToolGraphingModel struct {
	ID          types.String `tfsdk:"id"`
	PageRefresh types.Int64  `tfsdk:"page_refresh"`
	StoreEvery  types.String `tfsdk:"store_every"`
	Router      types.String `tfsdk:"router"`
}

func NewToolGraphingResource() resource.Resource { return &ToolGraphingResource{} }

func (r *ToolGraphingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_graphing"
}

func (r *ToolGraphingResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolGraphingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/tool/graphing`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"page_refresh": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"store_every": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *ToolGraphingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolGraphingModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolGraphingUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolGraphingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ToolGraphingModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state ToolGraphingModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolGraphingUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolGraphingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolGraphingModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/tool/graphing")
	if err != nil {
		resp.Diagnostics.AddError("Read /tool/graphing failed", err.Error())
		return
	}
	toolGraphingApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/tool/graphing", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolGraphingResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *ToolGraphingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/tool/graphing" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/tool/graphing", types.StringValue(routerName))))...)
}

func toolGraphingUpsert(ctx context.Context, reg *client.Registry, plan, state *ToolGraphingModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.PageRefresh.IsNull() || plan.PageRefresh.IsUnknown()) && (state == nil || !plan.PageRefresh.Equal(state.PageRefresh)) {
		body["page-refresh"] = client.FormatInt64(plan.PageRefresh.ValueInt64())
	}
	if !(plan.StoreEvery.IsNull() || plan.StoreEvery.IsUnknown()) && (state == nil || !plan.StoreEvery.Equal(state.StoreEvery)) {
		body["store-every"] = plan.StoreEvery.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/tool/graphing", body)
	if err != nil {
		diags.AddError("Upsert /tool/graphing failed", err.Error())
		return
	}
	toolGraphingApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/tool/graphing", plan.Router))
}

func toolGraphingApply(ctx context.Context, obj client.Object, m *ToolGraphingModel) {
	_ = ctx
	if v, ok := obj["page-refresh"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PageRefresh = types.Int64Value(n)
		} else {
			m.PageRefresh = types.Int64Null()
		}
	}
	if v, ok := obj["store-every"]; ok {
		_ = v
		if v != "" {
			m.StoreEvery = types.StringValue(v)
		} else {
			m.StoreEvery = types.StringNull()
		}
	}
}
