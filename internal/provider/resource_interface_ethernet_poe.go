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
	_ resource.Resource                = &InterfaceEthernetPoeResource{}
	_ resource.ResourceWithImportState = &InterfaceEthernetPoeResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceEthernetPoeResource struct {
	reg *client.Registry
}

type InterfaceEthernetPoeModel struct {
	ID                    types.String    `tfsdk:"id"`
	Psu2MaxPower          types.String    `tfsdk:"psu2_max_power"`
	Psu1MaxPower          types.String    `tfsdk:"psu1_max_power"`
	PsuMaxPower           types.String    `tfsdk:"psu_max_power"`
	PowerCyclePingTimeout types.String    `tfsdk:"power_cycle_ping_timeout"`
	PowerCyclePingEnabled boolStringValue `tfsdk:"power_cycle_ping_enabled"`
	PowerCyclePingAddress types.String    `tfsdk:"power_cycle_ping_address"`
	PowerCycleInterval    types.String    `tfsdk:"power_cycle_interval"`
	PoeVoltage            types.String    `tfsdk:"poe_voltage"`
	PoePriority           types.String    `tfsdk:"poe_priority"`
	PoeOut                types.String    `tfsdk:"poe_out"`
	PoeInMaxPower         types.String    `tfsdk:"poe_in_max_power"`
	Name                  types.String    `tfsdk:"name"`
	Jack2MaxPower         types.String    `tfsdk:"jack2_max_power"`
	Jack1MaxPower         types.String    `tfsdk:"jack1_max_power"`
	JackMaxPower          types.String    `tfsdk:"jack_max_power"`
	Ether1PoeInLongCable  types.String    `tfsdk:"ether1_poe_in_long_cable"`
	PowerCycle            types.String    `tfsdk:"power_cycle"`
	Print                 types.String    `tfsdk:"print"`
	Router                types.String    `tfsdk:"router"`
}

func NewInterfaceEthernetPoeResource() resource.Resource { return &InterfaceEthernetPoeResource{} }

func (r *InterfaceEthernetPoeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ethernet_poe"
}

func (r *InterfaceEthernetPoeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceEthernetPoeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Requires PoE-capable ethernet port",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"psu2_max_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `psu2-max-power`.",
			},
			"psu1_max_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `psu1-max-power`.",
			},
			"psu_max_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `psu-max-power`.",
			},
			"power_cycle_ping_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `power-cycle-ping-timeout`.",
			},
			"power_cycle_ping_enabled": schema.StringAttribute{
				CustomType:  boolStringType{},
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `power-cycle-ping-enabled`.",
			},
			"power_cycle_ping_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `power-cycle-ping-address`.",
			},
			"power_cycle_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `power-cycle-interval`.",
			},
			"poe_voltage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `poe-voltage`.",
			},
			"poe_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `poe-priority`.",
			},
			"poe_out": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `poe-out`.",
			},
			"poe_in_max_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `poe-in-max-power`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `name`.",
			},
			"jack2_max_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `jack2-max-power`.",
			},
			"jack1_max_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `jack1-max-power`.",
			},
			"jack_max_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `jack-max-power`.",
			},
			"ether1_poe_in_long_cable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ether1-poe-in-long-cable`.",
			},
			"power_cycle": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Disables PoE-Out power for a specified period of time.",
			},
			"print": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Prints PoE-Out related settings.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceEthernetPoeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceEthernetPoeModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.PowerCycle.IsNull() || plan.PowerCycle.IsUnknown()) {
		body["power-cycle"] = plan.PowerCycle.ValueString()
	}
	if !(plan.Print.IsNull() || plan.Print.IsUnknown()) {
		body["print"] = plan.Print.ValueString()
	}
	if !(plan.Ether1PoeInLongCable.IsNull() || plan.Ether1PoeInLongCable.IsUnknown()) {
		body["ether1-poe-in-long-cable"] = plan.Ether1PoeInLongCable.ValueString()
	}
	if !(plan.JackMaxPower.IsNull() || plan.JackMaxPower.IsUnknown()) {
		body["jack-max-power"] = plan.JackMaxPower.ValueString()
	}
	if !(plan.Jack1MaxPower.IsNull() || plan.Jack1MaxPower.IsUnknown()) {
		body["jack1-max-power"] = plan.Jack1MaxPower.ValueString()
	}
	if !(plan.Jack2MaxPower.IsNull() || plan.Jack2MaxPower.IsUnknown()) {
		body["jack2-max-power"] = plan.Jack2MaxPower.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.PoeInMaxPower.IsNull() || plan.PoeInMaxPower.IsUnknown()) {
		body["poe-in-max-power"] = plan.PoeInMaxPower.ValueString()
	}
	if !(plan.PoeOut.IsNull() || plan.PoeOut.IsUnknown()) {
		body["poe-out"] = plan.PoeOut.ValueString()
	}
	if !(plan.PoePriority.IsNull() || plan.PoePriority.IsUnknown()) {
		body["poe-priority"] = plan.PoePriority.ValueString()
	}
	if !(plan.PoeVoltage.IsNull() || plan.PoeVoltage.IsUnknown()) {
		body["poe-voltage"] = plan.PoeVoltage.ValueString()
	}
	if !(plan.PowerCycleInterval.IsNull() || plan.PowerCycleInterval.IsUnknown()) {
		body["power-cycle-interval"] = plan.PowerCycleInterval.ValueString()
	}
	if !(plan.PowerCyclePingAddress.IsNull() || plan.PowerCyclePingAddress.IsUnknown()) {
		body["power-cycle-ping-address"] = plan.PowerCyclePingAddress.ValueString()
	}
	if !(plan.PowerCyclePingEnabled.IsNull() || plan.PowerCyclePingEnabled.IsUnknown()) {
		body["power-cycle-ping-enabled"] = plan.PowerCyclePingEnabled.ValueString()
	}
	if !(plan.PowerCyclePingTimeout.IsNull() || plan.PowerCyclePingTimeout.IsUnknown()) {
		body["power-cycle-ping-timeout"] = plan.PowerCyclePingTimeout.ValueString()
	}
	if !(plan.PsuMaxPower.IsNull() || plan.PsuMaxPower.IsUnknown()) {
		body["psu-max-power"] = plan.PsuMaxPower.ValueString()
	}
	if !(plan.Psu1MaxPower.IsNull() || plan.Psu1MaxPower.IsUnknown()) {
		body["psu1-max-power"] = plan.Psu1MaxPower.ValueString()
	}
	if !(plan.Psu2MaxPower.IsNull() || plan.Psu2MaxPower.IsUnknown()) {
		body["psu2-max-power"] = plan.Psu2MaxPower.ValueString()
	}
	rows, err := c.List(ctx, "/interface/ethernet/poe")
	if err != nil {
		resp.Diagnostics.AddError("Read /interface/ethernet/poe failed", err.Error())
		return
	}
	want := plan.Name.ValueString()
	var id string
	for _, row := range rows {
		if row["name"] == want || row["default-name"] == want {
			id = row[".id"]
			break
		}
	}
	if id == "" {
		resp.Diagnostics.AddError("Unknown /interface/ethernet/poe "+want, fmt.Sprintf("/interface/ethernet/poe is a fixed hardware row set; no row matches name %q. Import the interface instead of creating it.", want))
		return
	}
	obj, err := c.Set(ctx, "/interface/ethernet/poe", id, body)
	if err != nil {
		resp.Diagnostics.AddError("Adopt /interface/ethernet/poe failed", err.Error())
		return
	}
	interfaceEthernetPoeApply(ctx, obj, &plan)
	plan.ID = types.StringValue(id)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEthernetPoeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceEthernetPoeModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/ethernet/poe", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/ethernet/poe failed", err.Error())
		return
	}
	interfaceEthernetPoeApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceEthernetPoeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceEthernetPoeModel
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
	if !plan.PowerCycle.Equal(state.PowerCycle) && !plan.PowerCycle.IsUnknown() {
		body["power-cycle"] = plan.PowerCycle.ValueString()
	}
	if !plan.Print.Equal(state.Print) && !plan.Print.IsUnknown() {
		body["print"] = plan.Print.ValueString()
	}
	if !plan.Ether1PoeInLongCable.Equal(state.Ether1PoeInLongCable) && !plan.Ether1PoeInLongCable.IsUnknown() {
		body["ether1-poe-in-long-cable"] = plan.Ether1PoeInLongCable.ValueString()
	}
	if !plan.JackMaxPower.Equal(state.JackMaxPower) && !plan.JackMaxPower.IsUnknown() {
		body["jack-max-power"] = plan.JackMaxPower.ValueString()
	}
	if !plan.Jack1MaxPower.Equal(state.Jack1MaxPower) && !plan.Jack1MaxPower.IsUnknown() {
		body["jack1-max-power"] = plan.Jack1MaxPower.ValueString()
	}
	if !plan.Jack2MaxPower.Equal(state.Jack2MaxPower) && !plan.Jack2MaxPower.IsUnknown() {
		body["jack2-max-power"] = plan.Jack2MaxPower.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.PoeInMaxPower.Equal(state.PoeInMaxPower) && !plan.PoeInMaxPower.IsUnknown() {
		body["poe-in-max-power"] = plan.PoeInMaxPower.ValueString()
	}
	if !plan.PoeOut.Equal(state.PoeOut) && !plan.PoeOut.IsUnknown() {
		body["poe-out"] = plan.PoeOut.ValueString()
	}
	if !plan.PoePriority.Equal(state.PoePriority) && !plan.PoePriority.IsUnknown() {
		body["poe-priority"] = plan.PoePriority.ValueString()
	}
	if !plan.PoeVoltage.Equal(state.PoeVoltage) && !plan.PoeVoltage.IsUnknown() {
		body["poe-voltage"] = plan.PoeVoltage.ValueString()
	}
	if !plan.PowerCycleInterval.Equal(state.PowerCycleInterval) && !plan.PowerCycleInterval.IsUnknown() {
		body["power-cycle-interval"] = plan.PowerCycleInterval.ValueString()
	}
	if !plan.PowerCyclePingAddress.Equal(state.PowerCyclePingAddress) && !plan.PowerCyclePingAddress.IsUnknown() {
		body["power-cycle-ping-address"] = plan.PowerCyclePingAddress.ValueString()
	}
	if !plan.PowerCyclePingEnabled.Equal(state.PowerCyclePingEnabled) && !plan.PowerCyclePingEnabled.IsUnknown() {
		body["power-cycle-ping-enabled"] = plan.PowerCyclePingEnabled.ValueString()
	}
	if !plan.PowerCyclePingTimeout.Equal(state.PowerCyclePingTimeout) && !plan.PowerCyclePingTimeout.IsUnknown() {
		body["power-cycle-ping-timeout"] = plan.PowerCyclePingTimeout.ValueString()
	}
	if !plan.PsuMaxPower.Equal(state.PsuMaxPower) && !plan.PsuMaxPower.IsUnknown() {
		body["psu-max-power"] = plan.PsuMaxPower.ValueString()
	}
	if !plan.Psu1MaxPower.Equal(state.Psu1MaxPower) && !plan.Psu1MaxPower.IsUnknown() {
		body["psu1-max-power"] = plan.Psu1MaxPower.ValueString()
	}
	if !plan.Psu2MaxPower.Equal(state.Psu2MaxPower) && !plan.Psu2MaxPower.IsUnknown() {
		body["psu2-max-power"] = plan.Psu2MaxPower.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/ethernet/poe", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/ethernet/poe failed", err.Error())
			return
		}
		interfaceEthernetPoeApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEthernetPoeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Fixed hardware row: cannot be removed. Drop from state; the row keeps
	// its last-applied settings (adopt-only, like /ip/service).
	_ = ctx
	_ = req
	_ = resp
}

func (r *InterfaceEthernetPoeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceEthernetPoeLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/ethernet/poe matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceEthernetPoeLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceEthernetPoeLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/ethernet/poe", id)
}

func interfaceEthernetPoeApply(ctx context.Context, obj client.Object, m *InterfaceEthernetPoeModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["psu2-max-power"]; ok && v != "" {
		m.Psu2MaxPower = types.StringValue(v)
	} else {
		m.Psu2MaxPower = types.StringNull()
	}
	if v, ok := obj["psu1-max-power"]; ok && v != "" {
		m.Psu1MaxPower = types.StringValue(v)
	} else {
		m.Psu1MaxPower = types.StringNull()
	}
	if v, ok := obj["psu-max-power"]; ok && v != "" {
		m.PsuMaxPower = types.StringValue(v)
	} else {
		m.PsuMaxPower = types.StringNull()
	}
	if v, ok := obj["power-cycle-ping-timeout"]; ok && v != "" {
		m.PowerCyclePingTimeout = types.StringValue(v)
	} else {
		m.PowerCyclePingTimeout = types.StringNull()
	}
	if v, ok := obj["power-cycle-ping-enabled"]; ok && v != "" {
		m.PowerCyclePingEnabled = newBoolStringValue(v)
	}
	if v, ok := obj["power-cycle-ping-address"]; ok && v != "" {
		m.PowerCyclePingAddress = types.StringValue(v)
	} else {
		m.PowerCyclePingAddress = types.StringNull()
	}
	if v, ok := obj["power-cycle-interval"]; ok && v != "" {
		m.PowerCycleInterval = types.StringValue(v)
	} else {
		m.PowerCycleInterval = types.StringNull()
	}
	if v, ok := obj["poe-voltage"]; ok && v != "" {
		m.PoeVoltage = types.StringValue(v)
	} else {
		m.PoeVoltage = types.StringNull()
	}
	if v, ok := obj["poe-priority"]; ok && v != "" {
		m.PoePriority = types.StringValue(v)
	} else {
		m.PoePriority = types.StringNull()
	}
	if v, ok := obj["poe-out"]; ok && v != "" {
		m.PoeOut = types.StringValue(v)
	} else {
		m.PoeOut = types.StringNull()
	}
	if v, ok := obj["poe-in-max-power"]; ok && v != "" {
		m.PoeInMaxPower = types.StringValue(v)
	} else {
		m.PoeInMaxPower = types.StringNull()
	}
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["jack2-max-power"]; ok && v != "" {
		m.Jack2MaxPower = types.StringValue(v)
	} else {
		m.Jack2MaxPower = types.StringNull()
	}
	if v, ok := obj["jack1-max-power"]; ok && v != "" {
		m.Jack1MaxPower = types.StringValue(v)
	} else {
		m.Jack1MaxPower = types.StringNull()
	}
	if v, ok := obj["jack-max-power"]; ok && v != "" {
		m.JackMaxPower = types.StringValue(v)
	} else {
		m.JackMaxPower = types.StringNull()
	}
	if v, ok := obj["ether1-poe-in-long-cable"]; ok && v != "" {
		m.Ether1PoeInLongCable = types.StringValue(v)
	} else {
		m.Ether1PoeInLongCable = types.StringNull()
	}
	if v, ok := obj["power-cycle"]; ok {
		_ = v
		if v != "" {
			m.PowerCycle = types.StringValue(v)
		} else {
			m.PowerCycle = types.StringNull()
		}
	} else {
		m.PowerCycle = types.StringNull()
	}
	if v, ok := obj["print"]; ok {
		_ = v
		if v != "" {
			m.Print = types.StringValue(v)
		} else {
			m.Print = types.StringNull()
		}
	} else {
		m.Print = types.StringNull()
	}
}
