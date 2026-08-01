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
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &SystemLedsResource{}
	_ resource.ResourceWithImportState = &SystemLedsResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemLedsResource struct {
	reg *client.Registry
}

type SystemLedsModel struct {
	ID                   types.String `tfsdk:"id"`
	ModemSignalThreshold types.String `tfsdk:"modem_signal_threshold"`
	Color                types.String `tfsdk:"color"`
	Disabled             types.Bool   `tfsdk:"disabled"`
	Interface            types.String `tfsdk:"interface"`
	Leds                 types.String `tfsdk:"leds"`
	Type                 types.String `tfsdk:"type"`
	Router               types.String `tfsdk:"router"`
}

func NewSystemLedsResource() resource.Resource { return &SystemLedsResource{} }

func (r *SystemLedsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_leds"
}

func (r *SystemLedsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemLedsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "LED bindings — type/leds values depend on the specific device's available LEDs; not portable in an auto-test.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"modem_signal_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `modem-signal-threshold`.",
			},
			"color": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `color`.",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the interface which will be used for led status. Applicable only if \u00a0 type \u00a0 is interface specific.",
			},
			"leds": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "List of led names used for a status report. For example, wireless signal strength will require more than one led.",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of the status: align-down \u00a0 - light the led if the w60g device needs to be aligned downwards for the best signal quality align-left \u00a0 - light the led if the w60g device needs to be aligned to the left align-right \u00a0 - light the led if the w60g device needs to be aligned to the right align-up \u00a0 - light the led if the w60g device needs to be aligned upwards ap-cap \u00a0 - blink on CAP initializing with CAPsMAN, steady on once connected fan-fault \u00a0 - light the led when any of the devices controlled fans stop working flash-access \u00a0 - blink the led on flash access interface-activity \u00a0 - blink the led on interface (traffic) activity interface-receive \u00a0 - blink the led on interface received a traffic interface-speed \u00a0 - light the led when interface works in 10Gbit rate interface-speed-1G \u00a0 - light the led when interface works in 1Gbit rate interface-speed-25G \u00a0 - light the led when interface works in 25Gbit rate interface-speed-100G - light the led when interface works in 100Gbit rate interface-status \u00a0 - light the led on interface status change interface-transmit \u00a0 - blink the led on interface transmitted traffic modem-signal \u00a0 - blink the led on 3G modem signal (either USB or miniPCIe) modem-technology \u00a0 - turns on LEDs in order of modem technology generation: GSM; 3G; LTE; single led turns on only when LTE is active. off \u00a0 - turn off the led on \u00a0 - turn on the led poe-fault \u00a0 - light the led when PoE out budget is close to the maximum supported limit poe-out \u00a0 - light the led when interface PoE out turns on wireless-signal-strength \u00a0 - light the leds displaying wireless signal (requires more than one led) wireless-status \u00a0 - light the led on wireless status change.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemLedsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemLedsModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.Leds.IsNull() || plan.Leds.IsUnknown()) {
		body["leds"] = plan.Leds.ValueString()
	}
	if !(plan.Type.IsNull() || plan.Type.IsUnknown()) {
		body["type"] = plan.Type.ValueString()
	}
	if !(plan.Color.IsNull() || plan.Color.IsUnknown()) {
		body["color"] = plan.Color.ValueString()
	}
	if !(plan.ModemSignalThreshold.IsNull() || plan.ModemSignalThreshold.IsUnknown()) {
		body["modem-signal-threshold"] = plan.ModemSignalThreshold.ValueString()
	}
	obj, err := c.Add(ctx, "/system/leds", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /system/leds failed", err.Error())
		return
	}
	systemLedsApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemLedsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemLedsModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/system/leds", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /system/leds failed", err.Error())
		return
	}
	systemLedsApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemLedsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemLedsModel
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
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.Leds.Equal(state.Leds) && !plan.Leds.IsUnknown() {
		body["leds"] = plan.Leds.ValueString()
	}
	if !plan.Type.Equal(state.Type) && !plan.Type.IsUnknown() {
		body["type"] = plan.Type.ValueString()
	}
	if !plan.Color.Equal(state.Color) && !plan.Color.IsUnknown() {
		body["color"] = plan.Color.ValueString()
	}
	if !plan.ModemSignalThreshold.Equal(state.ModemSignalThreshold) && !plan.ModemSignalThreshold.IsUnknown() {
		body["modem-signal-threshold"] = plan.ModemSignalThreshold.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/system/leds", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /system/leds failed", err.Error())
			return
		}
		systemLedsApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemLedsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemLedsModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/system/leds", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /system/leds failed", err.Error())
	}
}

func (r *SystemLedsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := systemLedsLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /system/leds matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// systemLedsLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func systemLedsLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/system/leds", id)
}

func systemLedsApply(ctx context.Context, obj client.Object, m *SystemLedsModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["modem-signal-threshold"]; ok && v != "" {
		m.ModemSignalThreshold = types.StringValue(v)
	} else {
		m.ModemSignalThreshold = types.StringNull()
	}
	if v, ok := obj["color"]; ok && v != "" {
		m.Color = types.StringValue(v)
	} else {
		m.Color = types.StringNull()
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["interface"]; ok {
		_ = v
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	} else {
		m.Interface = types.StringNull()
	}
	if v, ok := obj["leds"]; ok {
		_ = v
		if v != "" {
			m.Leds = types.StringValue(v)
		} else {
			m.Leds = types.StringNull()
		}
	} else {
		m.Leds = types.StringNull()
	}
	if v, ok := obj["type"]; ok {
		_ = v
		if v != "" {
			m.Type = types.StringValue(v)
		} else {
			m.Type = types.StringNull()
		}
	} else {
		m.Type = types.StringNull()
	}
}
