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
	_ resource.Resource                = &ToolMACServerPingResource{}
	_ resource.ResourceWithImportState = &ToolMACServerPingResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type ToolMACServerPingResource struct {
	reg *client.Registry
}

type ToolMACServerPingModel struct {
	ID      types.String    `tfsdk:"id"`
	Enabled boolStringValue `tfsdk:"enabled"`
	Router  types.String    `tfsdk:"router"`
}

func NewToolMACServerPingResource() resource.Resource { return &ToolMACServerPingResource{} }

func (r *ToolMACServerPingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_mac_server_ping"
}

func (r *ToolMACServerPingResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolMACServerPingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/tool/mac-server/ping`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.StringAttribute{
				CustomType: boolStringType{}, Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *ToolMACServerPingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolMACServerPingModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolMACServerPingUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolMACServerPingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ToolMACServerPingModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state ToolMACServerPingModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolMACServerPingUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolMACServerPingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolMACServerPingModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/tool/mac-server/ping")
	if err != nil {
		resp.Diagnostics.AddError("Read /tool/mac-server/ping failed", err.Error())
		return
	}
	toolMACServerPingApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/tool/mac-server/ping", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolMACServerPingResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *ToolMACServerPingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/tool/mac-server/ping" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/tool/mac-server/ping", types.StringValue(routerName))))...)
}

func toolMACServerPingUpsert(ctx context.Context, reg *client.Registry, plan, state *ToolMACServerPingModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = plan.Enabled.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/tool/mac-server/ping", body)
	if err != nil {
		diags.AddError("Upsert /tool/mac-server/ping failed", err.Error())
		return
	}
	toolMACServerPingApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/tool/mac-server/ping", plan.Router))
}

func toolMACServerPingApply(ctx context.Context, obj client.Object, m *ToolMACServerPingModel) {
	_ = ctx
	if v, ok := obj["enabled"]; ok {
		_ = v
		if v != "" {
			m.Enabled = newBoolStringValue(v)
		}
	}
}
