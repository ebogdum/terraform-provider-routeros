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
	_ resource.Resource                = &SystemHealthSettingsResource{}
	_ resource.ResourceWithImportState = &SystemHealthSettingsResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemHealthSettingsResource struct {
	reg *client.Registry
}

type SystemHealthSettingsModel struct {
	ID                      types.String `tfsdk:"id"`
	UseFan                  types.String `tfsdk:"use_fan"`
	FanTargetTemp           types.String `tfsdk:"fan_target_temp"`
	FanSwitch               types.String `tfsdk:"fan_switch"`
	FanOnThreshold          types.String `tfsdk:"fan_on_threshold"`
	FanMode                 types.String `tfsdk:"fan_mode"`
	FanMinSpeedPercent      types.String `tfsdk:"fan_min_speed_percent"`
	FanFullSpeedTemp        types.String `tfsdk:"fan_full_speed_temp"`
	FanControlInterval      types.String `tfsdk:"fan_control_interval"`
	CPUOvertempCheck        types.Bool   `tfsdk:"cpu_overtemp_check"`
	CPUOvertempStartupDelay types.String `tfsdk:"cpu_overtemp_startup_delay"`
	CPUOvertempThreshold    types.String `tfsdk:"cpu_overtemp_threshold"`
	Router                  types.String `tfsdk:"router"`
}

func NewSystemHealthSettingsResource() resource.Resource { return &SystemHealthSettingsResource{} }

func (r *SystemHealthSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_health_settings"
}

func (r *SystemHealthSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemHealthSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/system/health/settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"use_fan": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `use-fan`.",
			},
			"fan_target_temp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fan-target-temp`.",
			},
			"fan_switch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fan-switch`.",
			},
			"fan_on_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fan-on-threshold`.",
			},
			"fan_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fan-mode`.",
			},
			"fan_min_speed_percent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fan-min-speed-percent`.",
			},
			"fan_full_speed_temp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fan-full-speed-temp`.",
			},
			"fan_control_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fan-control-interval`.",
			},
			"cpu_overtemp_check": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cpu-overtemp-check`.",
			},
			"cpu_overtemp_startup_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cpu-overtemp-startup-delay`.",
			},
			"cpu_overtemp_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cpu-overtemp-threshold`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemHealthSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemHealthSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemHealthSettingsUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemHealthSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SystemHealthSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemHealthSettingsUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemHealthSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemHealthSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/system/health/settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /system/health/settings failed", err.Error())
		return
	}
	systemHealthSettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/system/health/settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemHealthSettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *SystemHealthSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/system/health/settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/system/health/settings", types.StringValue(routerName))))...)
}

func systemHealthSettingsUpsert(ctx context.Context, reg *client.Registry, plan *SystemHealthSettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.CPUOvertempCheck.IsNull() || plan.CPUOvertempCheck.IsUnknown()) {
		body["cpu-overtemp-check"] = client.FormatBool(plan.CPUOvertempCheck.ValueBool())
	}
	if !(plan.CPUOvertempStartupDelay.IsNull() || plan.CPUOvertempStartupDelay.IsUnknown()) {
		body["cpu-overtemp-startup-delay"] = plan.CPUOvertempStartupDelay.ValueString()
	}
	if !(plan.CPUOvertempThreshold.IsNull() || plan.CPUOvertempThreshold.IsUnknown()) {
		body["cpu-overtemp-threshold"] = plan.CPUOvertempThreshold.ValueString()
	}
	if !(plan.FanControlInterval.IsNull() || plan.FanControlInterval.IsUnknown()) {
		body["fan-control-interval"] = plan.FanControlInterval.ValueString()
	}
	if !(plan.FanFullSpeedTemp.IsNull() || plan.FanFullSpeedTemp.IsUnknown()) {
		body["fan-full-speed-temp"] = plan.FanFullSpeedTemp.ValueString()
	}
	if !(plan.FanMinSpeedPercent.IsNull() || plan.FanMinSpeedPercent.IsUnknown()) {
		body["fan-min-speed-percent"] = plan.FanMinSpeedPercent.ValueString()
	}
	if !(plan.FanMode.IsNull() || plan.FanMode.IsUnknown()) {
		body["fan-mode"] = plan.FanMode.ValueString()
	}
	if !(plan.FanOnThreshold.IsNull() || plan.FanOnThreshold.IsUnknown()) {
		body["fan-on-threshold"] = plan.FanOnThreshold.ValueString()
	}
	if !(plan.FanSwitch.IsNull() || plan.FanSwitch.IsUnknown()) {
		body["fan-switch"] = plan.FanSwitch.ValueString()
	}
	if !(plan.FanTargetTemp.IsNull() || plan.FanTargetTemp.IsUnknown()) {
		body["fan-target-temp"] = plan.FanTargetTemp.ValueString()
	}
	if !(plan.UseFan.IsNull() || plan.UseFan.IsUnknown()) {
		body["use-fan"] = plan.UseFan.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/system/health/settings", body)
	if err != nil {
		diags.AddError("Upsert /system/health/settings failed", err.Error())
		return
	}
	systemHealthSettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/system/health/settings", plan.Router))
}

func systemHealthSettingsApply(ctx context.Context, obj client.Object, m *SystemHealthSettingsModel) {
	_ = ctx
	if v, ok := obj["use-fan"]; ok && v != "" {
		m.UseFan = types.StringValue(v)
	} else {
		m.UseFan = types.StringNull()
	}
	if v, ok := obj["fan-target-temp"]; ok && v != "" {
		m.FanTargetTemp = types.StringValue(v)
	} else {
		m.FanTargetTemp = types.StringNull()
	}
	if v, ok := obj["fan-switch"]; ok && v != "" {
		m.FanSwitch = types.StringValue(v)
	} else {
		m.FanSwitch = types.StringNull()
	}
	if v, ok := obj["fan-on-threshold"]; ok && v != "" {
		m.FanOnThreshold = types.StringValue(v)
	} else {
		m.FanOnThreshold = types.StringNull()
	}
	if v, ok := obj["fan-mode"]; ok && v != "" {
		m.FanMode = types.StringValue(v)
	} else {
		m.FanMode = types.StringNull()
	}
	if v, ok := obj["fan-min-speed-percent"]; ok && v != "" {
		m.FanMinSpeedPercent = types.StringValue(v)
	} else {
		m.FanMinSpeedPercent = types.StringNull()
	}
	if v, ok := obj["fan-full-speed-temp"]; ok && v != "" {
		m.FanFullSpeedTemp = types.StringValue(v)
	} else {
		m.FanFullSpeedTemp = types.StringNull()
	}
	if v, ok := obj["fan-control-interval"]; ok && v != "" {
		m.FanControlInterval = types.StringValue(v)
	} else {
		m.FanControlInterval = types.StringNull()
	}
	if v, ok := obj["cpu-overtemp-check"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.CPUOvertempCheck = types.BoolValue(b)
		} else {
			m.CPUOvertempCheck = types.BoolNull()
		}
	} else {
		m.CPUOvertempCheck = types.BoolNull()
	}
	if v, ok := obj["cpu-overtemp-startup-delay"]; ok && v != "" {
		m.CPUOvertempStartupDelay = types.StringValue(v)
	} else {
		m.CPUOvertempStartupDelay = types.StringNull()
	}
	if v, ok := obj["cpu-overtemp-threshold"]; ok && v != "" {
		m.CPUOvertempThreshold = types.StringValue(v)
	} else {
		m.CPUOvertempThreshold = types.StringNull()
	}
}
