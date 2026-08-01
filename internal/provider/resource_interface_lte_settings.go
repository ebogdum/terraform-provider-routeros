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
	_ resource.Resource                = &InterfaceLTESettingsResource{}
	_ resource.ResourceWithImportState = &InterfaceLTESettingsResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceLTESettingsResource struct {
	reg *client.Registry
}

type InterfaceLTESettingsModel struct {
	ID                types.String `tfsdk:"id"`
	EsimChannel       types.String `tfsdk:"esim_channel"`
	FirmwarePath      types.String `tfsdk:"firmware_path"`
	LinkRecoveryTimer types.Int64  `tfsdk:"link_recovery_timer"`
	Mode              types.String `tfsdk:"mode"`
	Router            types.String `tfsdk:"router"`
}

func NewInterfaceLTESettingsResource() resource.Resource { return &InterfaceLTESettingsResource{} }

func (r *InterfaceLTESettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_lte_settings"
}

func (r *InterfaceLTESettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceLTESettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/lte/settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"esim_channel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `esim-channel`.",
			},
			"firmware_path": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `firmware-path`.",
			},
			"link_recovery_timer": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `link-recovery-timer`.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mode`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceLTESettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceLTESettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	interfaceLTESettingsUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceLTESettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InterfaceLTESettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state InterfaceLTESettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	interfaceLTESettingsUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceLTESettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceLTESettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/interface/lte/settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /interface/lte/settings failed", err.Error())
		return
	}
	interfaceLTESettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/interface/lte/settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceLTESettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *InterfaceLTESettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/interface/lte/settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/interface/lte/settings", types.StringValue(routerName))))...)
}

func interfaceLTESettingsUpsert(ctx context.Context, reg *client.Registry, plan, state *InterfaceLTESettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.EsimChannel.IsNull() || plan.EsimChannel.IsUnknown()) && (state == nil || !plan.EsimChannel.Equal(state.EsimChannel)) {
		body["esim-channel"] = plan.EsimChannel.ValueString()
	}
	if !(plan.FirmwarePath.IsNull() || plan.FirmwarePath.IsUnknown()) && (state == nil || !plan.FirmwarePath.Equal(state.FirmwarePath)) {
		body["firmware-path"] = plan.FirmwarePath.ValueString()
	}
	if !(plan.LinkRecoveryTimer.IsNull() || plan.LinkRecoveryTimer.IsUnknown()) && (state == nil || !plan.LinkRecoveryTimer.Equal(state.LinkRecoveryTimer)) {
		body["link-recovery-timer"] = client.FormatInt64(plan.LinkRecoveryTimer.ValueInt64())
	}
	if !(plan.Mode.IsNull() || plan.Mode.IsUnknown()) && (state == nil || !plan.Mode.Equal(state.Mode)) {
		body["mode"] = plan.Mode.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/interface/lte/settings", body)
	if err != nil {
		diags.AddError("Upsert /interface/lte/settings failed", err.Error())
		return
	}
	interfaceLTESettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/interface/lte/settings", plan.Router))
}

func interfaceLTESettingsApply(ctx context.Context, obj client.Object, m *InterfaceLTESettingsModel) {
	_ = ctx
	if v, ok := obj["esim-channel"]; ok && v != "" {
		m.EsimChannel = types.StringValue(v)
	} else {
		m.EsimChannel = types.StringNull()
	}
	if v, ok := obj["firmware-path"]; ok && v != "" {
		m.FirmwarePath = types.StringValue(v)
	} else {
		m.FirmwarePath = types.StringNull()
	}
	if v, ok := obj["link-recovery-timer"]; ok {
		if n, err := client.ParseInt64(v); err == nil {
			m.LinkRecoveryTimer = types.Int64Value(n)
		} else {
			m.LinkRecoveryTimer = types.Int64Null()
		}
	} else {
		m.LinkRecoveryTimer = types.Int64Null()
	}
	if v, ok := obj["mode"]; ok && v != "" {
		m.Mode = types.StringValue(v)
	} else {
		m.Mode = types.StringNull()
	}
}
