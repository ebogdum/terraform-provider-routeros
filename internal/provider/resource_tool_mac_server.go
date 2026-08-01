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
	"github.com/ebogdum/terraform-provider-routeros/internal/schemautil"
)

var (
	_ resource.Resource                = &ToolMACServerResource{}
	_ resource.ResourceWithImportState = &ToolMACServerResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type ToolMACServerResource struct {
	reg *client.Registry
}

type ToolMACServerModel struct {
	ID                   types.String `tfsdk:"id"`
	AllowedInterfaceList types.String `tfsdk:"allowed_interface_list"`
	Router               types.String `tfsdk:"router"`
	LockoutAck           types.Bool   `tfsdk:"lockout_ack"`
}

func NewToolMACServerResource() resource.Resource { return &ToolMACServerResource{} }

func (r *ToolMACServerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_mac_server"
}

func (r *ToolMACServerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolMACServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/tool/mac-server`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"allowed_interface_list": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
			"lockout_ack": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "Acknowledge that this rule may sever management traffic (required for unconditional input/forward drop/reject/tarpit rules with no match).",
			},
		},
	}
}

func (r *ToolMACServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolMACServerModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolMACServerUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolMACServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ToolMACServerModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state ToolMACServerModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolMACServerUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolMACServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolMACServerModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/tool/mac-server")
	if err != nil {
		resp.Diagnostics.AddError("Read /tool/mac-server failed", err.Error())
		return
	}
	toolMACServerApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/tool/mac-server", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolMACServerResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *ToolMACServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/tool/mac-server" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/tool/mac-server", types.StringValue(routerName))))...)
}

func toolMACServerUpsert(ctx context.Context, reg *client.Registry, plan, state *ToolMACServerModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AllowedInterfaceList.IsNull() || plan.AllowedInterfaceList.IsUnknown()) && (state == nil || !plan.AllowedInterfaceList.Equal(state.AllowedInterfaceList)) {
		body["allowed-interface-list"] = plan.AllowedInterfaceList.ValueString()
	}
	if err := schemautil.CheckMACServerLockout("/tool/mac-server", body, !plan.LockoutAck.IsNull() && plan.LockoutAck.ValueBool()); err != nil {
		diags.AddError("Refusing mac-server change", err.Error())
		return
	}
	obj, err := c.SetSingleton(ctx, "/tool/mac-server", body)
	if err != nil {
		diags.AddError("Upsert /tool/mac-server failed", err.Error())
		return
	}
	toolMACServerApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/tool/mac-server", plan.Router))
}

func toolMACServerApply(ctx context.Context, obj client.Object, m *ToolMACServerModel) {
	_ = ctx
	if v, ok := obj["allowed-interface-list"]; ok {
		_ = v
		if v != "" {
			m.AllowedInterfaceList = types.StringValue(v)
		} else {
			m.AllowedInterfaceList = types.StringNull()
		}
	}
}
