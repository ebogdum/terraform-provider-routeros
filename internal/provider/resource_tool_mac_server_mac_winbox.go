package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &ToolMACServerMACWinboxResource{}
	_ resource.ResourceWithImportState = &ToolMACServerMACWinboxResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type ToolMACServerMACWinboxResource struct {
	reg *client.Registry
}

type ToolMACServerMACWinboxModel struct {
	ID                   types.String `tfsdk:"id"`
	AllowedInterfaceList types.String `tfsdk:"allowed_interface_list"`
	Router               types.String `tfsdk:"router"`
}

func NewToolMACServerMACWinboxResource() resource.Resource { return &ToolMACServerMACWinboxResource{} }

func (r *ToolMACServerMACWinboxResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_mac_server_mac_winbox"
}

func (r *ToolMACServerMACWinboxResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolMACServerMACWinboxResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/tool/mac-server/mac-winbox`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"allowed_interface_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `allowed-interface-list`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *ToolMACServerMACWinboxResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolMACServerMACWinboxModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolMACServerMACWinboxUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolMACServerMACWinboxResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ToolMACServerMACWinboxModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state ToolMACServerMACWinboxModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolMACServerMACWinboxUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolMACServerMACWinboxResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolMACServerMACWinboxModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/tool/mac-server/mac-winbox")
	if err != nil {
		resp.Diagnostics.AddError("Read /tool/mac-server/mac-winbox failed", err.Error())
		return
	}
	toolMACServerMACWinboxApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/tool/mac-server/mac-winbox", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolMACServerMACWinboxResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *ToolMACServerMACWinboxResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/tool/mac-server/mac-winbox" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/tool/mac-server/mac-winbox", types.StringValue(routerName))))...)
}

func toolMACServerMACWinboxUpsert(ctx context.Context, reg *client.Registry, plan, state *ToolMACServerMACWinboxModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AllowedInterfaceList.IsNull() || plan.AllowedInterfaceList.IsUnknown()) && (state == nil || !plan.AllowedInterfaceList.Equal(state.AllowedInterfaceList)) {
		body["allowed-interface-list"] = plan.AllowedInterfaceList.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/tool/mac-server/mac-winbox", body)
	if err != nil {
		diags.AddError("Upsert /tool/mac-server/mac-winbox failed", err.Error())
		return
	}
	toolMACServerMACWinboxApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/tool/mac-server/mac-winbox", plan.Router))
}

func toolMACServerMACWinboxApply(ctx context.Context, obj client.Object, m *ToolMACServerMACWinboxModel) {
	_ = ctx
	if v, ok := obj["allowed-interface-list"]; ok && v != "" {
		m.AllowedInterfaceList = types.StringValue(v)
	} else {
		m.AllowedInterfaceList = types.StringNull()
	}
}
