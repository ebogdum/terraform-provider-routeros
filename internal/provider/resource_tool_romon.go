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
	_ resource.Resource                = &ToolRomonResource{}
	_ resource.ResourceWithImportState = &ToolRomonResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type ToolRomonResource struct {
	reg *client.Registry
}

type ToolRomonModel struct {
	ID      types.String `tfsdk:"id"`
	Enabled types.Bool   `tfsdk:"enabled"`
	Secrets types.String `tfsdk:"secrets"`
	Router  types.String `tfsdk:"router"`
}

func NewToolRomonResource() resource.Resource { return &ToolRomonResource{} }

func (r *ToolRomonResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_romon"
}

func (r *ToolRomonResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolRomonResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/tool/romon`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"secrets": schema.StringAttribute{Optional: true, Computed: true, Sensitive: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *ToolRomonResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolRomonModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolRomonUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolRomonResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ToolRomonModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state ToolRomonModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolRomonUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolRomonResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolRomonModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/tool/romon")
	if err != nil {
		resp.Diagnostics.AddError("Read /tool/romon failed", err.Error())
		return
	}
	toolRomonApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/tool/romon", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolRomonResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *ToolRomonResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/tool/romon" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/tool/romon", types.StringValue(routerName))))...)
}

func toolRomonUpsert(ctx context.Context, reg *client.Registry, plan, state *ToolRomonModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	if !(plan.Secrets.IsNull() || plan.Secrets.IsUnknown()) && (state == nil || !plan.Secrets.Equal(state.Secrets)) {
		body["secrets"] = plan.Secrets.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/tool/romon", body)
	if err != nil {
		diags.AddError("Upsert /tool/romon failed", err.Error())
		return
	}
	toolRomonApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/tool/romon", plan.Router))
}

func toolRomonApply(ctx context.Context, obj client.Object, m *ToolRomonModel) {
	_ = ctx
	if v, ok := obj["enabled"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Enabled = types.BoolValue(b)
		} else {
			m.Enabled = types.BoolNull()
		}
	}
	if v, ok := obj["secrets"]; ok && v != "" {
		_ = v
		if v != "" {
			m.Secrets = types.StringValue(v)
		} else {
			m.Secrets = types.StringNull()
		}
	} else if m.Secrets.IsUnknown() {
		m.Secrets = types.StringNull()
	}
}
