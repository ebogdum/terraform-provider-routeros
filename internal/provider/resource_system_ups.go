package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
	"github.com/ebogdum/terraform-provider-routeros/internal/schemautil"
)

var (
	_ resource.Resource                = &SystemUpsResource{}
	_ resource.ResourceWithImportState = &SystemUpsResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemUpsResource struct {
	reg *client.Registry
}

type SystemUpsModel struct {
	ID                    types.String `tfsdk:"id"`
	AlarmSetting          types.String `tfsdk:"alarm_setting"`
	BatteryCharge         types.Int64  `tfsdk:"battery_charge"`
	BatteryVoltage        types.String `tfsdk:"battery_voltage"`
	Beep                  types.String `tfsdk:"beep"`
	CheckCapabilities     types.String `tfsdk:"check_capabilities"`
	Comment               types.String `tfsdk:"comment"`
	Disabled              types.Bool   `tfsdk:"disabled"`
	Frequency             types.Int64  `tfsdk:"frequency"`
	Invalid               types.Bool   `tfsdk:"invalid"`
	LineVoltage           types.String `tfsdk:"line_voltage"`
	Load                  types.Int64  `tfsdk:"load"`
	LowBattery            types.Bool   `tfsdk:"low_battery"`
	ManufactureDate       types.String `tfsdk:"manufacture_date"`
	MinRuntime            types.String `tfsdk:"min_runtime"`
	Model                 types.String `tfsdk:"model"`
	Name                  types.String `tfsdk:"name"`
	NominalBatteryVoltage types.Int64  `tfsdk:"nominal_battery_voltage"`
	OfflineAfter          types.String `tfsdk:"offline_after"`
	OfflineTime           types.String `tfsdk:"offline_time"`
	OnBattery             types.Bool   `tfsdk:"on_battery"`
	OnLine                types.Bool   `tfsdk:"on_line"`
	OuputVoltage          types.String `tfsdk:"ouput_voltage"`
	Overload              types.Bool   `tfsdk:"overload"`
	Port                  types.String `tfsdk:"port"`
	ReplaceBattery        types.Bool   `tfsdk:"replace_battery"`
	RunTimeLeft           types.String `tfsdk:"run_time_left"`
	SerialNumber          types.String `tfsdk:"serial_number"`
	SmartBoost            types.Bool   `tfsdk:"smart_boost"`
	SmartTrim             types.Bool   `tfsdk:"smart_trim"`
	Temperature           types.String `tfsdk:"temperature"`
	TransferCause         types.String `tfsdk:"transfer_cause"`
	Version               types.String `tfsdk:"version"`
	Router                types.String `tfsdk:"router"`
}

func NewSystemUpsResource() resource.Resource { return &SystemUpsResource{} }

func (r *SystemUpsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_ups"
}

func (r *SystemUpsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *SystemUpsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/system/ups`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"alarm_setting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "UPS sound alarm setting: delayed - alarm is delayed to the on-battery event immediate - alarm immediately after the on-battery event low-battery - alarm only when the battery is low none - do not alarm",
				Validators:  []validator.String{schemautil.OneOf([]string{"immediate", "delayed", "low-battery", "none"}...)},
			},
			"battery_charge": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"battery_voltage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"beep": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"check_capabilities": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether to check UPS capabilities before reading information. Disabling it can fix compatibility issues with some UPS models. (Applies to RouterOS version 6, implemented since v6.17)",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"frequency": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"line_voltage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"load": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"low_battery": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"manufacture_date": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"min_runtime": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Minimal run time remaining. After a 'utility' failure, the router will monitor the runtime-left value. When the value reaches the min-runtime value, the router will go to hibernate mode. never - the router will go to hibernate mode when the \"battery low\" signal is sent indicating that the battery power is below 10% 0s - the router will continue to work as long as the battery is supplying sufficient voltage",
			},
			"model": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nominal_battery_voltage": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"offline_after": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"offline_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "How long to work on batteries. The router waits that amount of time and then goes into hibernate mode until the UPS reports that the 'utility' power is back 0s - the router will go into hibernate mode according to the min-runtime setting. In this case, the router will wait until the UPS reports that the battery power is below 10%",
			},
			"on_battery": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"on_line": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ouput_voltage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"overload": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Communication port of the router.",
			},
			"replace_battery": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"run_time_left": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"serial_number": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"smart_boost": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"smart_trim": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"temperature": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"transfer_cause": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"version": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemUpsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemUpsModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AlarmSetting.IsNull() || plan.AlarmSetting.IsUnknown()) {
		body["alarm-setting"] = plan.AlarmSetting.ValueString()
	}
	if !(plan.Beep.IsNull() || plan.Beep.IsUnknown()) {
		body["beep"] = plan.Beep.ValueString()
	}
	if !(plan.CheckCapabilities.IsNull() || plan.CheckCapabilities.IsUnknown()) {
		body["check-capabilities"] = plan.CheckCapabilities.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.MinRuntime.IsNull() || plan.MinRuntime.IsUnknown()) {
		body["min-runtime"] = plan.MinRuntime.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.OfflineTime.IsNull() || plan.OfflineTime.IsUnknown()) {
		body["offline-time"] = plan.OfflineTime.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	obj, err := c.Add(ctx, "/system/ups", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /system/ups failed", err.Error())
		return
	}
	systemUpsApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemUpsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemUpsModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/system/ups", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /system/ups failed", err.Error())
		return
	}
	systemUpsApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemUpsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemUpsModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !plan.AlarmSetting.Equal(state.AlarmSetting) {
		body["alarm-setting"] = plan.AlarmSetting.ValueString()
	}
	if !plan.Beep.Equal(state.Beep) {
		body["beep"] = plan.Beep.ValueString()
	}
	if !plan.CheckCapabilities.Equal(state.CheckCapabilities) {
		body["check-capabilities"] = plan.CheckCapabilities.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.MinRuntime.Equal(state.MinRuntime) {
		body["min-runtime"] = plan.MinRuntime.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.OfflineTime.Equal(state.OfflineTime) {
		body["offline-time"] = plan.OfflineTime.ValueString()
	}
	if !plan.Port.Equal(state.Port) {
		body["port"] = plan.Port.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/system/ups", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /system/ups failed", err.Error())
			return
		}
		systemUpsApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemUpsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemUpsModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/system/ups", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /system/ups failed", err.Error())
	}
}

func (r *SystemUpsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import formats accepted:
	//   *<id>                            -> bare RouterOS .id on the default router
	//   <router>/*<id>                   -> .id on the named router
	//   <router>/<naturalkey>            -> resolved via List + filter
	//   <naturalkey>                     -> resolved on the default router
	routerName, id := parseImportID(r.reg, req.ID)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	if strings.HasPrefix(id, "*") {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(id))...)
		return
	}
	c := pickClient(r.reg, types.StringValue(routerName), &resp.Diagnostics)
	if c == nil {
		return
	}
	rows, err := systemUpsLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /system/ups matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// systemUpsLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func systemUpsLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/system/ups", id)
}

func systemUpsApply(ctx context.Context, obj client.Object, m *SystemUpsModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["alarm-setting"]; ok {
		_ = v
		if v != "" {
			m.AlarmSetting = types.StringValue(v)
		} else {
			m.AlarmSetting = types.StringNull()
		}
	} else {
		m.AlarmSetting = types.StringNull()
	}
	if v, ok := obj["battery-charge"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.BatteryCharge = types.Int64Value(n)
		} else {
			m.BatteryCharge = types.Int64Null()
		}
	} else {
		m.BatteryCharge = types.Int64Null()
	}
	if v, ok := obj["battery-voltage"]; ok {
		_ = v
		if v != "" {
			m.BatteryVoltage = types.StringValue(v)
		} else {
			m.BatteryVoltage = types.StringNull()
		}
	} else {
		m.BatteryVoltage = types.StringNull()
	}
	if v, ok := obj["beep"]; ok {
		_ = v
		if v != "" {
			m.Beep = types.StringValue(v)
		} else {
			m.Beep = types.StringNull()
		}
	} else {
		m.Beep = types.StringNull()
	}
	if v, ok := obj["check-capabilities"]; ok {
		_ = v
		if v != "" {
			m.CheckCapabilities = types.StringValue(v)
		} else {
			m.CheckCapabilities = types.StringNull()
		}
	} else {
		m.CheckCapabilities = types.StringNull()
	}
	if v, ok := obj["comment"]; ok {
		_ = v
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	} else {
		m.Comment = types.StringNull()
	}
	if v, ok := obj["disabled"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	} else {
		m.Disabled = types.BoolNull()
	}
	if v, ok := obj["frequency"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Frequency = types.Int64Value(n)
		} else {
			m.Frequency = types.Int64Null()
		}
	} else {
		m.Frequency = types.Int64Null()
	}
	if v, ok := obj["invalid"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Invalid = types.BoolValue(b)
		} else {
			m.Invalid = types.BoolNull()
		}
	} else {
		m.Invalid = types.BoolNull()
	}
	if v, ok := obj["line-voltage"]; ok {
		_ = v
		if v != "" {
			m.LineVoltage = types.StringValue(v)
		} else {
			m.LineVoltage = types.StringNull()
		}
	} else {
		m.LineVoltage = types.StringNull()
	}
	if v, ok := obj["load"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Load = types.Int64Value(n)
		} else {
			m.Load = types.Int64Null()
		}
	} else {
		m.Load = types.Int64Null()
	}
	if v, ok := obj["low-battery"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.LowBattery = types.BoolValue(b)
		} else {
			m.LowBattery = types.BoolNull()
		}
	} else {
		m.LowBattery = types.BoolNull()
	}
	if v, ok := obj["manufacture-date"]; ok {
		_ = v
		if v != "" {
			m.ManufactureDate = types.StringValue(v)
		} else {
			m.ManufactureDate = types.StringNull()
		}
	} else {
		m.ManufactureDate = types.StringNull()
	}
	if v, ok := obj["min-runtime"]; ok {
		_ = v
		if v != "" {
			m.MinRuntime = types.StringValue(v)
		} else {
			m.MinRuntime = types.StringNull()
		}
	} else {
		m.MinRuntime = types.StringNull()
	}
	if v, ok := obj["model"]; ok {
		_ = v
		if v != "" {
			m.Model = types.StringValue(v)
		} else {
			m.Model = types.StringNull()
		}
	} else {
		m.Model = types.StringNull()
	}
	if v, ok := obj["name"]; ok {
		_ = v
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["nominal-battery-voltage"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.NominalBatteryVoltage = types.Int64Value(n)
		} else {
			m.NominalBatteryVoltage = types.Int64Null()
		}
	} else {
		m.NominalBatteryVoltage = types.Int64Null()
	}
	if v, ok := obj["offline-after"]; ok {
		_ = v
		if v != "" {
			m.OfflineAfter = types.StringValue(v)
		} else {
			m.OfflineAfter = types.StringNull()
		}
	} else {
		m.OfflineAfter = types.StringNull()
	}
	if v, ok := obj["offline-time"]; ok {
		_ = v
		if v != "" {
			m.OfflineTime = types.StringValue(v)
		} else {
			m.OfflineTime = types.StringNull()
		}
	} else {
		m.OfflineTime = types.StringNull()
	}
	if v, ok := obj["on-battery"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.OnBattery = types.BoolValue(b)
		} else {
			m.OnBattery = types.BoolNull()
		}
	} else {
		m.OnBattery = types.BoolNull()
	}
	if v, ok := obj["on-line"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.OnLine = types.BoolValue(b)
		} else {
			m.OnLine = types.BoolNull()
		}
	} else {
		m.OnLine = types.BoolNull()
	}
	if v, ok := obj["ouput-voltage"]; ok {
		_ = v
		if v != "" {
			m.OuputVoltage = types.StringValue(v)
		} else {
			m.OuputVoltage = types.StringNull()
		}
	} else {
		m.OuputVoltage = types.StringNull()
	}
	if v, ok := obj["overload"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Overload = types.BoolValue(b)
		} else {
			m.Overload = types.BoolNull()
		}
	} else {
		m.Overload = types.BoolNull()
	}
	if v, ok := obj["port"]; ok {
		_ = v
		if v != "" {
			m.Port = types.StringValue(v)
		} else {
			m.Port = types.StringNull()
		}
	} else {
		m.Port = types.StringNull()
	}
	if v, ok := obj["replace-battery"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.ReplaceBattery = types.BoolValue(b)
		} else {
			m.ReplaceBattery = types.BoolNull()
		}
	} else {
		m.ReplaceBattery = types.BoolNull()
	}
	if v, ok := obj["run-time-left"]; ok {
		_ = v
		if v != "" {
			m.RunTimeLeft = types.StringValue(v)
		} else {
			m.RunTimeLeft = types.StringNull()
		}
	} else {
		m.RunTimeLeft = types.StringNull()
	}
	if v, ok := obj["serial-number"]; ok {
		_ = v
		if v != "" {
			m.SerialNumber = types.StringValue(v)
		} else {
			m.SerialNumber = types.StringNull()
		}
	} else {
		m.SerialNumber = types.StringNull()
	}
	if v, ok := obj["smart-boost"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.SmartBoost = types.BoolValue(b)
		} else {
			m.SmartBoost = types.BoolNull()
		}
	} else {
		m.SmartBoost = types.BoolNull()
	}
	if v, ok := obj["smart-trim"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.SmartTrim = types.BoolValue(b)
		} else {
			m.SmartTrim = types.BoolNull()
		}
	} else {
		m.SmartTrim = types.BoolNull()
	}
	if v, ok := obj["temperature"]; ok {
		_ = v
		if v != "" {
			m.Temperature = types.StringValue(v)
		} else {
			m.Temperature = types.StringNull()
		}
	} else {
		m.Temperature = types.StringNull()
	}
	if v, ok := obj["transfer-cause"]; ok {
		_ = v
		if v != "" {
			m.TransferCause = types.StringValue(v)
		} else {
			m.TransferCause = types.StringNull()
		}
	} else {
		m.TransferCause = types.StringNull()
	}
	if v, ok := obj["version"]; ok {
		_ = v
		if v != "" {
			m.Version = types.StringValue(v)
		} else {
			m.Version = types.StringNull()
		}
	} else {
		m.Version = types.StringNull()
	}
}
