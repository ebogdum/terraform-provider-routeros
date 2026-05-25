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
	_ resource.Resource                = &InterfaceEthernetResource{}
	_ resource.ResourceWithImportState = &InterfaceEthernetResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceEthernetResource struct {
	reg *client.Registry
}

type InterfaceEthernetModel struct {
	ID                        types.String `tfsdk:"id"`
	Advertise                 types.List   `tfsdk:"advertise"`
	Advertising               types.String `tfsdk:"advertising"`
	ARP                       types.String `tfsdk:"arp"`
	ARPTimeout                types.String `tfsdk:"arp_timeout"`
	AutoNegotiation           types.String `tfsdk:"auto_negotiation"`
	Autoneg                   types.Bool   `tfsdk:"autoneg"`
	Blink                     types.String `tfsdk:"blink"`
	CableAssemblyLinkLength   types.String `tfsdk:"cable_assembly_link_length"`
	CableSettings             types.String `tfsdk:"cable_settings"`
	CableTest                 types.String `tfsdk:"cable_test"`
	CmisModuleState           types.String `tfsdk:"cmis_module_state"`
	CmisRevision              types.String `tfsdk:"cmis_revision"`
	Combo                     types.Int64  `tfsdk:"combo"`
	ComboMode                 types.String `tfsdk:"combo_mode"`
	Comment                   types.String `tfsdk:"comment"`
	ConnectorType             types.String `tfsdk:"connector_type"`
	CopperActiveOm4LinkLength types.Int64  `tfsdk:"copper_active_om4_link_length"`
	DefaultName               types.String `tfsdk:"default_name"`
	DisableRunningCheck       types.Bool   `tfsdk:"disable_running_check"`
	DisableTime               types.String `tfsdk:"disable_time"`
	Disabled                  types.Bool   `tfsdk:"disabled"`
	Encoding                  types.String `tfsdk:"encoding"`
	Extrastats                types.String `tfsdk:"extrastats"`
	Fec                       types.Int64  `tfsdk:"fec"`
	FecMode                   types.String `tfsdk:"fec_mode"`
	Flowcntrl                 types.String `tfsdk:"flowcntrl"`
	Flowcontrol               types.Int64  `tfsdk:"flowcontrol"`
	FullDuplex                types.String `tfsdk:"full_duplex"`
	Hastxqueuestats           types.Bool   `tfsdk:"hastxqueuestats"`
	IgnoreRxLos               types.Bool   `tfsdk:"ignore_rx_los"`
	L2MTU                     types.Int64  `tfsdk:"l2_mtu"`
	LinkPartnerAdvertising    types.String `tfsdk:"link_partner_advertising"`
	LoopProtect               types.String `tfsdk:"loop_protect"`
	LoopProtectDisableTime    types.String `tfsdk:"loop_protect_disable_time"`
	LoopProtectSendInterval   types.String `tfsdk:"loop_protect_send_interval"`
	LoopProtectStatus         types.String `tfsdk:"loop_protect_status"`
	MACAddress                types.String `tfsdk:"mac_address"`
	ManufacturingDate         types.String `tfsdk:"manufacturing_date"`
	MaxL2MTU                  types.Int64  `tfsdk:"max_l2_mtu"`
	MaxPower                  types.String `tfsdk:"max_power"`
	ModulePresent             types.Bool   `tfsdk:"module_present"`
	MTU                       types.Int64  `tfsdk:"mtu"`
	Name                      types.String `tfsdk:"name"`
	Noautoneg                 types.String `tfsdk:"noautoneg"`
	NonMgmt                   types.String `tfsdk:"non_mgmt"`
	Om1LinkLength             types.Int64  `tfsdk:"om1_link_length"`
	Om2LinkLength             types.Int64  `tfsdk:"om2_link_length"`
	Om3LinkLength             types.Int64  `tfsdk:"om3_link_length"`
	Om4LinkLength             types.Int64  `tfsdk:"om4_link_length"`
	Om5LinkLength             types.Int64  `tfsdk:"om5_link_length"`
	OrigMACAddress            types.String `tfsdk:"orig_mac_address"`
	PassthroughInterface      types.String `tfsdk:"passthrough_interface"`
	PciePassthrough           types.Int64  `tfsdk:"pcie_passthrough"`
	PoEOut                    types.String `tfsdk:"po_e_out"`
	PoEOutCurrent             types.Int64  `tfsdk:"po_e_out_current"`
	PoEOutPower               types.String `tfsdk:"po_e_out_power"`
	PoEOutStatus              types.String `tfsdk:"po_e_out_status"`
	PoEOutVoltage             types.String `tfsdk:"po_e_out_voltage"`
	PoEPriority               types.Int64  `tfsdk:"po_e_priority"`
	PoEVoltage                types.String `tfsdk:"po_e_voltage"`
	Poe                       types.String `tfsdk:"poe"`
	PoeV                      types.Bool   `tfsdk:"poe_v"`
	Poecurr                   types.Int64  `tfsdk:"poecurr"`
	Poeping                   types.String `tfsdk:"poeping"`
	Poepower                  types.Int64  `tfsdk:"poepower"`
	Poevolt                   types.Int64  `tfsdk:"poevolt"`
	PowerClass                types.Int64  `tfsdk:"power_class"`
	PowerCycle                types.String `tfsdk:"power_cycle"`
	PowerCycleAfter           types.String `tfsdk:"power_cycle_after"`
	PowerCycleHostAlive       types.Bool   `tfsdk:"power_cycle_host_alive"`
	PowerCycleInterval        types.String `tfsdk:"power_cycle_interval"`
	PowerCyclePingAddress     types.String `tfsdk:"power_cycle_ping_address"`
	PowerCyclePingEnabled     types.Bool   `tfsdk:"power_cycle_ping_enabled"`
	PowerCyclePingTimeout     types.String `tfsdk:"power_cycle_ping_timeout"`
	Qstats                    types.String `tfsdk:"qstats"`
	Rate                      types.String `tfsdk:"rate"`
	RateSelect                types.String `tfsdk:"rate_select"`
	ResetCounters             types.String `tfsdk:"reset_counters"`
	ResetMACAddress           types.String `tfsdk:"reset_mac_address"`
	Running                   types.Bool   `tfsdk:"running"`
	RxAlignError              types.String `tfsdk:"rx_align_error"`
	RxBroadcast               types.Int64  `tfsdk:"rx_broadcast"`
	RxBytes                   types.Int64  `tfsdk:"rx_bytes"`
	RxCarrierError            types.String `tfsdk:"rx_carrier_error"`
	RxCodeError               types.String `tfsdk:"rx_code_error"`
	RxControl                 types.String `tfsdk:"rx_control"`
	RxDrop                    types.String `tfsdk:"rx_drop"`
	RxErrorEvents             types.String `tfsdk:"rx_error_events"`
	RxFcsError                types.String `tfsdk:"rx_fcs_error"`
	RxFlowControl             types.String `tfsdk:"rx_flow_control"`
	RxFragment                types.String `tfsdk:"rx_fragment"`
	RxJabber                  types.String `tfsdk:"rx_jabber"`
	RxLengthError             types.String `tfsdk:"rx_length_error"`
	RxLoss                    types.Bool   `tfsdk:"rx_loss"`
	RxMulticast               types.Int64  `tfsdk:"rx_multicast"`
	RxOverflow                types.String `tfsdk:"rx_overflow"`
	RxPacket                  types.Int64  `tfsdk:"rx_packet"`
	RxPause                   types.String `tfsdk:"rx_pause"`
	RxPower                   types.String `tfsdk:"rx_power"`
	RxTooLong                 types.String `tfsdk:"rx_too_long"`
	RxTooShort                types.String `tfsdk:"rx_too_short"`
	RxUnicast                 types.String `tfsdk:"rx_unicast"`
	RxUnknownOp               types.String `tfsdk:"rx_unknown_op"`
	SendInterval              types.String `tfsdk:"send_interval"`
	Sfp                       types.Bool   `tfsdk:"sfp"`
	SfpShutdownTemperature    types.Int64  `tfsdk:"sfp_shutdown_temperature"`
	SfpSupported              types.String `tfsdk:"sfp_supported"`
	Sfprate                   types.Int64  `tfsdk:"sfprate"`
	Sfpshutdown               types.Bool   `tfsdk:"sfpshutdown"`
	SmLinkLength              types.String `tfsdk:"sm_link_length"`
	Speed                     types.String `tfsdk:"speed"`
	Status                    types.String `tfsdk:"status"`
	SupplyVoltage             types.String `tfsdk:"supply_voltage"`
	Supported                 types.String `tfsdk:"supported"`
	Temperature               types.String `tfsdk:"temperature"`
	TxBiasCurrent             types.Int64  `tfsdk:"tx_bias_current"`
	TxBroadcast               types.Int64  `tfsdk:"tx_broadcast"`
	TxBytes                   types.Int64  `tfsdk:"tx_bytes"`
	TxCollision               types.String `tfsdk:"tx_collision"`
	TxControl                 types.String `tfsdk:"tx_control"`
	TxDeferred                types.String `tfsdk:"tx_deferred"`
	TxDrop                    types.String `tfsdk:"tx_drop"`
	TxExcessiveCollision      types.String `tfsdk:"tx_excessive_collision"`
	TxExcessiveDeferred       types.String `tfsdk:"tx_excessive_deferred"`
	TxFault                   types.Bool   `tfsdk:"tx_fault"`
	TxFcsError                types.String `tfsdk:"tx_fcs_error"`
	TxFlowControl             types.String `tfsdk:"tx_flow_control"`
	TxFragment                types.String `tfsdk:"tx_fragment"`
	TxJabber                  types.String `tfsdk:"tx_jabber"`
	TxLateCollision           types.String `tfsdk:"tx_late_collision"`
	TxMulticast               types.Int64  `tfsdk:"tx_multicast"`
	TxMultipleCollision       types.String `tfsdk:"tx_multiple_collision"`
	TxPacket                  types.Int64  `tfsdk:"tx_packet"`
	TxPause                   types.String `tfsdk:"tx_pause"`
	TxPauseHonorred           types.String `tfsdk:"tx_pause_honorred"`
	TxPower                   types.String `tfsdk:"tx_power"`
	TxRx10241518              types.String `tfsdk:"tx_rx_1024_1518"`
	TxRx1024Max               types.String `tfsdk:"tx_rx_1024_max"`
	TxRx128255                types.String `tfsdk:"tx_rx_128_255"`
	TxRx1519Max               types.String `tfsdk:"tx_rx_1519_max"`
	TxRx256511                types.String `tfsdk:"tx_rx_256_511"`
	TxRx5121023               types.String `tfsdk:"tx_rx_512_1023"`
	TxRx64                    types.String `tfsdk:"tx_rx_64"`
	TxRx65127                 types.String `tfsdk:"tx_rx_65_127"`
	TxRxBytes                 types.String `tfsdk:"tx_rx_bytes"`
	TxRxPackets               types.String `tfsdk:"tx_rx_packets"`
	TxSingleCollision         types.String `tfsdk:"tx_single_collision"`
	TxTooShort                types.String `tfsdk:"tx_too_short"`
	TxTotalCollision          types.String `tfsdk:"tx_total_collision"`
	TxUnderrun                types.String `tfsdk:"tx_underrun"`
	TxUnicast                 types.String `tfsdk:"tx_unicast"`
	VendorName                types.String `tfsdk:"vendor_name"`
	VendorPartNumber          types.String `tfsdk:"vendor_part_number"`
	VendorRevision            types.String `tfsdk:"vendor_revision"`
	VendorSerial              types.String `tfsdk:"vendor_serial"`
	Wavelength                types.String `tfsdk:"wavelength"`
	Router                    types.String `tfsdk:"router"`
}

func NewInterfaceEthernetResource() resource.Resource { return &InterfaceEthernetResource{} }

func (r *InterfaceEthernetResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ethernet"
}

func (r *InterfaceEthernetResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *InterfaceEthernetResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "/interface/ethernet entries are auto-created from the physical NICs; can't be added via TF.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"advertise": schema.ListAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "",
			},
			"advertising": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"arp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"disabled", "enabled", "proxy-arp", "reply-only", "local-proxy-arp"}...)},
			},
			"arp_timeout": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"auto_negotiation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"incomplete", "done", "no-negotiation", "failed", "restarted", "disabled", "not-available"}...)},
			},
			"autoneg": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"blink": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"cable_assembly_link_length": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"cable_settings": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"cable_test": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"cmis_module_state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"low-power", "power-up", "ready", "power-down", "fault"}...)},
			},
			"cmis_revision": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"combo": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"combo_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"auto", "copper", "sfp"}...)},
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"connector_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"unknown", "sc", "lc", "optical-pigtail", "multifiber-parallel-optic-1x12", "copper-pigtail", "rj45", "no-separable-connector", "multifiber-parallel-optic-1x16"}...)},
			},
			"copper_active_om4_link_length": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"default_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disable_running_check": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disable_time": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"encoding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"unspecified", "8b/10b", "4b/5b", "nrz", "manchester", "sonet", "64b/66b", "256b/257b", "pam4"}...)},
			},
			"extrastats": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"fec": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"fec_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"off", "auto", "fec74", "fec91"}...)},
			},
			"flowcntrl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"flowcontrol": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"full_duplex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hastxqueuestats": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ignore_rx_los": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"l2_mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"link_partner_advertising": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"loop_protect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"default", "off", "on"}...)},
			},
			"loop_protect_disable_time": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"loop_protect_send_interval": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"loop_protect_status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mac_address": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "MAC address to be mapped to",
				Validators:    []validator.String{schemautil.IsMAC()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeMAC()},
			},
			"manufacturing_date": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"max_l2_mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"max_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"module_present": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"noautoneg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"non_mgmt": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"om1_link_length": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"om2_link_length": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"om3_link_length": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"om4_link_length": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"om5_link_length": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"orig_mac_address": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsMAC()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeMAC()},
			},
			"passthrough_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pcie_passthrough": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"po_e_out": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"off", "auto-on", "forced-on"}...)},
			},
			"po_e_out_current": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"po_e_out_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"po_e_out_status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "disabled", "waiting-for-load", "powered-on", "overload", "short-circuit", "voltage-too-low", "current-too-low", "power-cycle", "voltage-too-high", "controller-error", "controller-upgrade", "voltage-on-poe-in", "no-valid-psu"}...)},
			},
			"po_e_out_voltage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"po_e_priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"po_e_voltage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"auto", "low", "high"}...)},
			},
			"poe": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"poe_v": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"poecurr": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"poeping": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"poepower": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"poevolt": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"power_class": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"power_cycle": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"power_cycle_after": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"power_cycle_host_alive": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"power_cycle_interval": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"power_cycle_ping_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"power_cycle_ping_enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"power_cycle_ping_timeout": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"qstats": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rate_select": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"low", "high"}...)},
			},
			"reset_counters": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"reset_mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"running": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_align_error": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_broadcast": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_bytes": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_carrier_error": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_code_error": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_control": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_drop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_error_events": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_fcs_error": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_flow_control": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"off", "on", "auto"}...)},
			},
			"rx_fragment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_jabber": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_length_error": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_loss": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_multicast": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_overflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_packet": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_pause": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_too_long": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_too_short": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_unicast": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_unknown_op": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"send_interval": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"sfp": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sfp_shutdown_temperature": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sfp_supported": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sfprate": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sfpshutdown": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sm_link_length": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"speed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"10m-baset-half", "10m-baset-full", "100m-baset-half", "100m-baset-full", "1g-baset-half", "1g-baset-full", "10g-baset", "2.5g-basex", "40g-basecr4", "40g-basesr4-lr4", "25g-basecr", "25g-basesr-lr", "50g-basecr2", "100g-basesr4-lr4", "100g-basecr4", "50g-basesr2-lr2", "1g-basex", "10g-basecr", "10g-basesr-lr", "2.5g-baset", "5g-baset", "50g-basesr-lr", "50g-basecr", "100g-basesr2-lr2", "100g-basecr2", "200g-basesr4-lr4", "200g-basecr4", "400g-basesr8-lr8", "400g-basecr8", "100m-basefx-half", "100m-basefx-full"}...)},
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "off", "on", "disabled"}...)},
			},
			"supply_voltage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"supported": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"temperature": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_bias_current": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_broadcast": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_bytes": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_collision": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_control": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_deferred": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_drop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_excessive_collision": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_excessive_deferred": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_fault": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_fcs_error": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_flow_control": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"off", "on", "auto"}...)},
			},
			"tx_fragment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_jabber": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_late_collision": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_multicast": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_multiple_collision": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_packet": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_pause": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_pause_honorred": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_1024_1518": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_1024_max": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_128_255": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_1519_max": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_256_511": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_512_1023": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_64": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_65_127": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_bytes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_packets": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_single_collision": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_too_short": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_total_collision": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_underrun": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_unicast": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vendor_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vendor_part_number": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vendor_revision": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vendor_serial": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wavelength": schema.StringAttribute{
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

func (r *InterfaceEthernetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceEthernetModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Advertise.IsNull() || plan.Advertise.IsUnknown()) {
		body["advertise"] = encodeStringList(ctx, plan.Advertise, &resp.Diagnostics)
	}
	if !(plan.ARP.IsNull() || plan.ARP.IsUnknown()) {
		body["arp"] = plan.ARP.ValueString()
	}
	if !(plan.ARPTimeout.IsNull() || plan.ARPTimeout.IsUnknown()) {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !(plan.Autoneg.IsNull() || plan.Autoneg.IsUnknown()) {
		body["autoneg"] = client.FormatBool(plan.Autoneg.ValueBool())
	}
	if !(plan.Blink.IsNull() || plan.Blink.IsUnknown()) {
		body["blink"] = plan.Blink.ValueString()
	}
	if !(plan.CableSettings.IsNull() || plan.CableSettings.IsUnknown()) {
		body["cable-settings"] = plan.CableSettings.ValueString()
	}
	if !(plan.CableTest.IsNull() || plan.CableTest.IsUnknown()) {
		body["cable-test"] = plan.CableTest.ValueString()
	}
	if !(plan.ComboMode.IsNull() || plan.ComboMode.IsUnknown()) {
		body["combo-mode"] = plan.ComboMode.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.DisableRunningCheck.IsNull() || plan.DisableRunningCheck.IsUnknown()) {
		body["disable-running-check"] = client.FormatBool(plan.DisableRunningCheck.ValueBool())
	}
	if !(plan.DisableTime.IsNull() || plan.DisableTime.IsUnknown()) {
		body["disable-time"] = plan.DisableTime.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Extrastats.IsNull() || plan.Extrastats.IsUnknown()) {
		body["extrastats"] = plan.Extrastats.ValueString()
	}
	if !(plan.FecMode.IsNull() || plan.FecMode.IsUnknown()) {
		body["fec-mode"] = plan.FecMode.ValueString()
	}
	if !(plan.Flowcntrl.IsNull() || plan.Flowcntrl.IsUnknown()) {
		body["flowcntrl"] = plan.Flowcntrl.ValueString()
	}
	if !(plan.IgnoreRxLos.IsNull() || plan.IgnoreRxLos.IsUnknown()) {
		body["ignore-rx-los"] = client.FormatBool(plan.IgnoreRxLos.ValueBool())
	}
	if !(plan.L2MTU.IsNull() || plan.L2MTU.IsUnknown()) {
		body["l2-mtu"] = client.FormatInt64(plan.L2MTU.ValueInt64())
	}
	if !(plan.LoopProtect.IsNull() || plan.LoopProtect.IsUnknown()) {
		body["loop-protect"] = plan.LoopProtect.ValueString()
	}
	if !(plan.LoopProtectDisableTime.IsNull() || plan.LoopProtectDisableTime.IsUnknown()) {
		body["loop-protect-disable-time"] = plan.LoopProtectDisableTime.ValueString()
	}
	if !(plan.LoopProtectSendInterval.IsNull() || plan.LoopProtectSendInterval.IsUnknown()) {
		body["loop-protect-send-interval"] = plan.LoopProtectSendInterval.ValueString()
	}
	if !(plan.MACAddress.IsNull() || plan.MACAddress.IsUnknown()) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = client.FormatInt64(plan.MTU.ValueInt64())
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Noautoneg.IsNull() || plan.Noautoneg.IsUnknown()) {
		body["noautoneg"] = plan.Noautoneg.ValueString()
	}
	if !(plan.NonMgmt.IsNull() || plan.NonMgmt.IsUnknown()) {
		body["non-mgmt"] = plan.NonMgmt.ValueString()
	}
	if !(plan.OrigMACAddress.IsNull() || plan.OrigMACAddress.IsUnknown()) {
		body["orig-mac-address"] = plan.OrigMACAddress.ValueString()
	}
	if !(plan.PassthroughInterface.IsNull() || plan.PassthroughInterface.IsUnknown()) {
		body["passthrough-interface"] = plan.PassthroughInterface.ValueString()
	}
	if !(plan.PoEOut.IsNull() || plan.PoEOut.IsUnknown()) {
		body["po-e-out"] = plan.PoEOut.ValueString()
	}
	if !(plan.PoEPriority.IsNull() || plan.PoEPriority.IsUnknown()) {
		body["po-e-priority"] = client.FormatInt64(plan.PoEPriority.ValueInt64())
	}
	if !(plan.PoEVoltage.IsNull() || plan.PoEVoltage.IsUnknown()) {
		body["po-e-voltage"] = plan.PoEVoltage.ValueString()
	}
	if !(plan.Poe.IsNull() || plan.Poe.IsUnknown()) {
		body["poe"] = plan.Poe.ValueString()
	}
	if !(plan.Poeping.IsNull() || plan.Poeping.IsUnknown()) {
		body["poeping"] = plan.Poeping.ValueString()
	}
	if !(plan.PowerCycle.IsNull() || plan.PowerCycle.IsUnknown()) {
		body["power-cycle"] = plan.PowerCycle.ValueString()
	}
	if !(plan.PowerCycleInterval.IsNull() || plan.PowerCycleInterval.IsUnknown()) {
		body["power-cycle-interval"] = plan.PowerCycleInterval.ValueString()
	}
	if !(plan.PowerCyclePingAddress.IsNull() || plan.PowerCyclePingAddress.IsUnknown()) {
		body["power-cycle-ping-address"] = plan.PowerCyclePingAddress.ValueString()
	}
	if !(plan.PowerCyclePingEnabled.IsNull() || plan.PowerCyclePingEnabled.IsUnknown()) {
		body["power-cycle-ping-enabled"] = client.FormatBool(plan.PowerCyclePingEnabled.ValueBool())
	}
	if !(plan.PowerCyclePingTimeout.IsNull() || plan.PowerCyclePingTimeout.IsUnknown()) {
		body["power-cycle-ping-timeout"] = plan.PowerCyclePingTimeout.ValueString()
	}
	if !(plan.Qstats.IsNull() || plan.Qstats.IsUnknown()) {
		body["qstats"] = plan.Qstats.ValueString()
	}
	if !(plan.RateSelect.IsNull() || plan.RateSelect.IsUnknown()) {
		body["rate-select"] = plan.RateSelect.ValueString()
	}
	if !(plan.ResetCounters.IsNull() || plan.ResetCounters.IsUnknown()) {
		body["reset-counters"] = plan.ResetCounters.ValueString()
	}
	if !(plan.ResetMACAddress.IsNull() || plan.ResetMACAddress.IsUnknown()) {
		body["reset-mac-address"] = plan.ResetMACAddress.ValueString()
	}
	if !(plan.RxFlowControl.IsNull() || plan.RxFlowControl.IsUnknown()) {
		body["rx-flow-control"] = plan.RxFlowControl.ValueString()
	}
	if !(plan.SendInterval.IsNull() || plan.SendInterval.IsUnknown()) {
		body["send-interval"] = plan.SendInterval.ValueString()
	}
	if !(plan.Sfp.IsNull() || plan.Sfp.IsUnknown()) {
		body["sfp"] = client.FormatBool(plan.Sfp.ValueBool())
	}
	if !(plan.SfpShutdownTemperature.IsNull() || plan.SfpShutdownTemperature.IsUnknown()) {
		body["sfp-shutdown-temperature"] = client.FormatInt64(plan.SfpShutdownTemperature.ValueInt64())
	}
	if !(plan.Speed.IsNull() || plan.Speed.IsUnknown()) {
		body["speed"] = plan.Speed.ValueString()
	}
	if !(plan.TxFlowControl.IsNull() || plan.TxFlowControl.IsUnknown()) {
		body["tx-flow-control"] = plan.TxFlowControl.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/ethernet", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/ethernet failed", err.Error())
		return
	}
	interfaceEthernetApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEthernetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceEthernetModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/ethernet", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/ethernet failed", err.Error())
		return
	}
	interfaceEthernetApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceEthernetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceEthernetModel
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
	if !plan.Advertise.Equal(state.Advertise) {
		body["advertise"] = encodeStringList(ctx, plan.Advertise, &resp.Diagnostics)
	}
	if !plan.ARP.Equal(state.ARP) {
		body["arp"] = plan.ARP.ValueString()
	}
	if !plan.ARPTimeout.Equal(state.ARPTimeout) {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !plan.Autoneg.Equal(state.Autoneg) {
		body["autoneg"] = client.FormatBool(plan.Autoneg.ValueBool())
	}
	if !plan.Blink.Equal(state.Blink) {
		body["blink"] = plan.Blink.ValueString()
	}
	if !plan.CableSettings.Equal(state.CableSettings) {
		body["cable-settings"] = plan.CableSettings.ValueString()
	}
	if !plan.CableTest.Equal(state.CableTest) {
		body["cable-test"] = plan.CableTest.ValueString()
	}
	if !plan.ComboMode.Equal(state.ComboMode) {
		body["combo-mode"] = plan.ComboMode.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.DisableRunningCheck.Equal(state.DisableRunningCheck) {
		body["disable-running-check"] = client.FormatBool(plan.DisableRunningCheck.ValueBool())
	}
	if !plan.DisableTime.Equal(state.DisableTime) {
		body["disable-time"] = plan.DisableTime.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Extrastats.Equal(state.Extrastats) {
		body["extrastats"] = plan.Extrastats.ValueString()
	}
	if !plan.FecMode.Equal(state.FecMode) {
		body["fec-mode"] = plan.FecMode.ValueString()
	}
	if !plan.Flowcntrl.Equal(state.Flowcntrl) {
		body["flowcntrl"] = plan.Flowcntrl.ValueString()
	}
	if !plan.IgnoreRxLos.Equal(state.IgnoreRxLos) {
		body["ignore-rx-los"] = client.FormatBool(plan.IgnoreRxLos.ValueBool())
	}
	if !plan.L2MTU.Equal(state.L2MTU) {
		body["l2-mtu"] = client.FormatInt64(plan.L2MTU.ValueInt64())
	}
	if !plan.LoopProtect.Equal(state.LoopProtect) {
		body["loop-protect"] = plan.LoopProtect.ValueString()
	}
	if !plan.LoopProtectDisableTime.Equal(state.LoopProtectDisableTime) {
		body["loop-protect-disable-time"] = plan.LoopProtectDisableTime.ValueString()
	}
	if !plan.LoopProtectSendInterval.Equal(state.LoopProtectSendInterval) {
		body["loop-protect-send-interval"] = plan.LoopProtectSendInterval.ValueString()
	}
	if !plan.MACAddress.Equal(state.MACAddress) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.MTU.Equal(state.MTU) {
		body["mtu"] = client.FormatInt64(plan.MTU.ValueInt64())
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Noautoneg.Equal(state.Noautoneg) {
		body["noautoneg"] = plan.Noautoneg.ValueString()
	}
	if !plan.NonMgmt.Equal(state.NonMgmt) {
		body["non-mgmt"] = plan.NonMgmt.ValueString()
	}
	if !plan.OrigMACAddress.Equal(state.OrigMACAddress) {
		body["orig-mac-address"] = plan.OrigMACAddress.ValueString()
	}
	if !plan.PassthroughInterface.Equal(state.PassthroughInterface) {
		body["passthrough-interface"] = plan.PassthroughInterface.ValueString()
	}
	if !plan.PoEOut.Equal(state.PoEOut) {
		body["po-e-out"] = plan.PoEOut.ValueString()
	}
	if !plan.PoEPriority.Equal(state.PoEPriority) {
		body["po-e-priority"] = client.FormatInt64(plan.PoEPriority.ValueInt64())
	}
	if !plan.PoEVoltage.Equal(state.PoEVoltage) {
		body["po-e-voltage"] = plan.PoEVoltage.ValueString()
	}
	if !plan.Poe.Equal(state.Poe) {
		body["poe"] = plan.Poe.ValueString()
	}
	if !plan.Poeping.Equal(state.Poeping) {
		body["poeping"] = plan.Poeping.ValueString()
	}
	if !plan.PowerCycle.Equal(state.PowerCycle) {
		body["power-cycle"] = plan.PowerCycle.ValueString()
	}
	if !plan.PowerCycleInterval.Equal(state.PowerCycleInterval) {
		body["power-cycle-interval"] = plan.PowerCycleInterval.ValueString()
	}
	if !plan.PowerCyclePingAddress.Equal(state.PowerCyclePingAddress) {
		body["power-cycle-ping-address"] = plan.PowerCyclePingAddress.ValueString()
	}
	if !plan.PowerCyclePingEnabled.Equal(state.PowerCyclePingEnabled) {
		body["power-cycle-ping-enabled"] = client.FormatBool(plan.PowerCyclePingEnabled.ValueBool())
	}
	if !plan.PowerCyclePingTimeout.Equal(state.PowerCyclePingTimeout) {
		body["power-cycle-ping-timeout"] = plan.PowerCyclePingTimeout.ValueString()
	}
	if !plan.Qstats.Equal(state.Qstats) {
		body["qstats"] = plan.Qstats.ValueString()
	}
	if !plan.RateSelect.Equal(state.RateSelect) {
		body["rate-select"] = plan.RateSelect.ValueString()
	}
	if !plan.ResetCounters.Equal(state.ResetCounters) {
		body["reset-counters"] = plan.ResetCounters.ValueString()
	}
	if !plan.ResetMACAddress.Equal(state.ResetMACAddress) {
		body["reset-mac-address"] = plan.ResetMACAddress.ValueString()
	}
	if !plan.RxFlowControl.Equal(state.RxFlowControl) {
		body["rx-flow-control"] = plan.RxFlowControl.ValueString()
	}
	if !plan.SendInterval.Equal(state.SendInterval) {
		body["send-interval"] = plan.SendInterval.ValueString()
	}
	if !plan.Sfp.Equal(state.Sfp) {
		body["sfp"] = client.FormatBool(plan.Sfp.ValueBool())
	}
	if !plan.SfpShutdownTemperature.Equal(state.SfpShutdownTemperature) {
		body["sfp-shutdown-temperature"] = client.FormatInt64(plan.SfpShutdownTemperature.ValueInt64())
	}
	if !plan.Speed.Equal(state.Speed) {
		body["speed"] = plan.Speed.ValueString()
	}
	if !plan.TxFlowControl.Equal(state.TxFlowControl) {
		body["tx-flow-control"] = plan.TxFlowControl.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/ethernet", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/ethernet failed", err.Error())
			return
		}
		interfaceEthernetApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEthernetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceEthernetModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/ethernet", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/ethernet failed", err.Error())
	}
}

func (r *InterfaceEthernetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceEthernetLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/ethernet matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceEthernetLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceEthernetLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/ethernet", id)
}

func interfaceEthernetApply(ctx context.Context, obj client.Object, m *InterfaceEthernetModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["advertise"]; ok {
		_ = v
		m.Advertise = decodeStringList(ctx, v)
	} else {
		m.Advertise = types.ListNull(types.StringType)
	}
	if v, ok := obj["advertising"]; ok {
		_ = v
		if v != "" {
			m.Advertising = types.StringValue(v)
		} else {
			m.Advertising = types.StringNull()
		}
	} else {
		m.Advertising = types.StringNull()
	}
	if v, ok := obj["arp"]; ok {
		_ = v
		if v != "" {
			m.ARP = types.StringValue(v)
		} else {
			m.ARP = types.StringNull()
		}
	} else {
		m.ARP = types.StringNull()
	}
	if v, ok := obj["arp-timeout"]; ok {
		_ = v
		if v != "" {
			m.ARPTimeout = types.StringValue(v)
		} else {
			m.ARPTimeout = types.StringNull()
		}
	} else {
		m.ARPTimeout = types.StringNull()
	}
	if v, ok := obj["auto-negotiation"]; ok {
		_ = v
		if v != "" {
			m.AutoNegotiation = types.StringValue(v)
		} else {
			m.AutoNegotiation = types.StringNull()
		}
	} else {
		m.AutoNegotiation = types.StringNull()
	}
	if v, ok := obj["autoneg"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Autoneg = types.BoolValue(b)
		} else {
			m.Autoneg = types.BoolNull()
		}
	} else {
		m.Autoneg = types.BoolNull()
	}
	if v, ok := obj["blink"]; ok {
		_ = v
		if v != "" {
			m.Blink = types.StringValue(v)
		} else {
			m.Blink = types.StringNull()
		}
	} else {
		m.Blink = types.StringNull()
	}
	if v, ok := obj["cable-assembly-link-length"]; ok {
		_ = v
		if v != "" {
			m.CableAssemblyLinkLength = types.StringValue(v)
		} else {
			m.CableAssemblyLinkLength = types.StringNull()
		}
	} else {
		m.CableAssemblyLinkLength = types.StringNull()
	}
	if v, ok := obj["cable-settings"]; ok {
		_ = v
		if v != "" {
			m.CableSettings = types.StringValue(v)
		} else {
			m.CableSettings = types.StringNull()
		}
	} else {
		m.CableSettings = types.StringNull()
	}
	if v, ok := obj["cable-test"]; ok {
		_ = v
		if v != "" {
			m.CableTest = types.StringValue(v)
		} else {
			m.CableTest = types.StringNull()
		}
	} else {
		m.CableTest = types.StringNull()
	}
	if v, ok := obj["cmis-module-state"]; ok {
		_ = v
		if v != "" {
			m.CmisModuleState = types.StringValue(v)
		} else {
			m.CmisModuleState = types.StringNull()
		}
	} else {
		m.CmisModuleState = types.StringNull()
	}
	if v, ok := obj["cmis-revision"]; ok {
		_ = v
		if v != "" {
			m.CmisRevision = types.StringValue(v)
		} else {
			m.CmisRevision = types.StringNull()
		}
	} else {
		m.CmisRevision = types.StringNull()
	}
	if v, ok := obj["combo"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Combo = types.Int64Value(n)
		} else {
			m.Combo = types.Int64Null()
		}
	} else {
		m.Combo = types.Int64Null()
	}
	if v, ok := obj["combo-mode"]; ok {
		_ = v
		if v != "" {
			m.ComboMode = types.StringValue(v)
		} else {
			m.ComboMode = types.StringNull()
		}
	} else {
		m.ComboMode = types.StringNull()
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
	if v, ok := obj["connector-type"]; ok {
		_ = v
		if v != "" {
			m.ConnectorType = types.StringValue(v)
		} else {
			m.ConnectorType = types.StringNull()
		}
	} else {
		m.ConnectorType = types.StringNull()
	}
	if v, ok := obj["copper-active-om4-link-length"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.CopperActiveOm4LinkLength = types.Int64Value(n)
		} else {
			m.CopperActiveOm4LinkLength = types.Int64Null()
		}
	} else {
		m.CopperActiveOm4LinkLength = types.Int64Null()
	}
	if v, ok := obj["default-name"]; ok {
		_ = v
		if v != "" {
			m.DefaultName = types.StringValue(v)
		} else {
			m.DefaultName = types.StringNull()
		}
	} else {
		m.DefaultName = types.StringNull()
	}
	if v, ok := obj["disable-running-check"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.DisableRunningCheck = types.BoolValue(b)
		} else {
			m.DisableRunningCheck = types.BoolNull()
		}
	} else {
		m.DisableRunningCheck = types.BoolNull()
	}
	if v, ok := obj["disable-time"]; ok {
		_ = v
		if v != "" {
			m.DisableTime = types.StringValue(v)
		} else {
			m.DisableTime = types.StringNull()
		}
	} else {
		m.DisableTime = types.StringNull()
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
	if v, ok := obj["encoding"]; ok {
		_ = v
		if v != "" {
			m.Encoding = types.StringValue(v)
		} else {
			m.Encoding = types.StringNull()
		}
	} else {
		m.Encoding = types.StringNull()
	}
	if v, ok := obj["extrastats"]; ok {
		_ = v
		if v != "" {
			m.Extrastats = types.StringValue(v)
		} else {
			m.Extrastats = types.StringNull()
		}
	} else {
		m.Extrastats = types.StringNull()
	}
	if v, ok := obj["fec"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Fec = types.Int64Value(n)
		} else {
			m.Fec = types.Int64Null()
		}
	} else {
		m.Fec = types.Int64Null()
	}
	if v, ok := obj["fec-mode"]; ok {
		_ = v
		if v != "" {
			m.FecMode = types.StringValue(v)
		} else {
			m.FecMode = types.StringNull()
		}
	} else {
		m.FecMode = types.StringNull()
	}
	if v, ok := obj["flowcntrl"]; ok {
		_ = v
		if v != "" {
			m.Flowcntrl = types.StringValue(v)
		} else {
			m.Flowcntrl = types.StringNull()
		}
	} else {
		m.Flowcntrl = types.StringNull()
	}
	if v, ok := obj["flowcontrol"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Flowcontrol = types.Int64Value(n)
		} else {
			m.Flowcontrol = types.Int64Null()
		}
	} else {
		m.Flowcontrol = types.Int64Null()
	}
	if v, ok := obj["full-duplex"]; ok {
		_ = v
		if v != "" {
			m.FullDuplex = types.StringValue(v)
		} else {
			m.FullDuplex = types.StringNull()
		}
	} else {
		m.FullDuplex = types.StringNull()
	}
	if v, ok := obj["hastxqueuestats"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Hastxqueuestats = types.BoolValue(b)
		} else {
			m.Hastxqueuestats = types.BoolNull()
		}
	} else {
		m.Hastxqueuestats = types.BoolNull()
	}
	if v, ok := obj["ignore-rx-los"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.IgnoreRxLos = types.BoolValue(b)
		} else {
			m.IgnoreRxLos = types.BoolNull()
		}
	} else {
		m.IgnoreRxLos = types.BoolNull()
	}
	if v, ok := obj["l2-mtu"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.L2MTU = types.Int64Value(n)
		} else {
			m.L2MTU = types.Int64Null()
		}
	} else {
		m.L2MTU = types.Int64Null()
	}
	if v, ok := obj["link-partner-advertising"]; ok {
		_ = v
		if v != "" {
			m.LinkPartnerAdvertising = types.StringValue(v)
		} else {
			m.LinkPartnerAdvertising = types.StringNull()
		}
	} else {
		m.LinkPartnerAdvertising = types.StringNull()
	}
	if v, ok := obj["loop-protect"]; ok {
		_ = v
		if v != "" {
			m.LoopProtect = types.StringValue(v)
		} else {
			m.LoopProtect = types.StringNull()
		}
	} else {
		m.LoopProtect = types.StringNull()
	}
	if v, ok := obj["loop-protect-disable-time"]; ok {
		_ = v
		if v != "" {
			m.LoopProtectDisableTime = types.StringValue(v)
		} else {
			m.LoopProtectDisableTime = types.StringNull()
		}
	} else {
		m.LoopProtectDisableTime = types.StringNull()
	}
	if v, ok := obj["loop-protect-send-interval"]; ok {
		_ = v
		if v != "" {
			m.LoopProtectSendInterval = types.StringValue(v)
		} else {
			m.LoopProtectSendInterval = types.StringNull()
		}
	} else {
		m.LoopProtectSendInterval = types.StringNull()
	}
	if v, ok := obj["loop-protect-status"]; ok {
		_ = v
		if v != "" {
			m.LoopProtectStatus = types.StringValue(v)
		} else {
			m.LoopProtectStatus = types.StringNull()
		}
	} else {
		m.LoopProtectStatus = types.StringNull()
	}
	if v, ok := obj["mac-address"]; ok {
		_ = v
		if v != "" {
			m.MACAddress = types.StringValue(v)
		} else {
			m.MACAddress = types.StringNull()
		}
	} else {
		m.MACAddress = types.StringNull()
	}
	if v, ok := obj["manufacturing-date"]; ok {
		_ = v
		if v != "" {
			m.ManufacturingDate = types.StringValue(v)
		} else {
			m.ManufacturingDate = types.StringNull()
		}
	} else {
		m.ManufacturingDate = types.StringNull()
	}
	if v, ok := obj["max-l2-mtu"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MaxL2MTU = types.Int64Value(n)
		} else {
			m.MaxL2MTU = types.Int64Null()
		}
	} else {
		m.MaxL2MTU = types.Int64Null()
	}
	if v, ok := obj["max-power"]; ok {
		_ = v
		if v != "" {
			m.MaxPower = types.StringValue(v)
		} else {
			m.MaxPower = types.StringNull()
		}
	} else {
		m.MaxPower = types.StringNull()
	}
	if v, ok := obj["module-present"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.ModulePresent = types.BoolValue(b)
		} else {
			m.ModulePresent = types.BoolNull()
		}
	} else {
		m.ModulePresent = types.BoolNull()
	}
	if v, ok := obj["mtu"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MTU = types.Int64Value(n)
		} else {
			m.MTU = types.Int64Null()
		}
	} else {
		m.MTU = types.Int64Null()
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
	if v, ok := obj["noautoneg"]; ok {
		_ = v
		if v != "" {
			m.Noautoneg = types.StringValue(v)
		} else {
			m.Noautoneg = types.StringNull()
		}
	} else {
		m.Noautoneg = types.StringNull()
	}
	if v, ok := obj["non-mgmt"]; ok {
		_ = v
		if v != "" {
			m.NonMgmt = types.StringValue(v)
		} else {
			m.NonMgmt = types.StringNull()
		}
	} else {
		m.NonMgmt = types.StringNull()
	}
	if v, ok := obj["om1-link-length"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Om1LinkLength = types.Int64Value(n)
		} else {
			m.Om1LinkLength = types.Int64Null()
		}
	} else {
		m.Om1LinkLength = types.Int64Null()
	}
	if v, ok := obj["om2-link-length"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Om2LinkLength = types.Int64Value(n)
		} else {
			m.Om2LinkLength = types.Int64Null()
		}
	} else {
		m.Om2LinkLength = types.Int64Null()
	}
	if v, ok := obj["om3-link-length"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Om3LinkLength = types.Int64Value(n)
		} else {
			m.Om3LinkLength = types.Int64Null()
		}
	} else {
		m.Om3LinkLength = types.Int64Null()
	}
	if v, ok := obj["om4-link-length"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Om4LinkLength = types.Int64Value(n)
		} else {
			m.Om4LinkLength = types.Int64Null()
		}
	} else {
		m.Om4LinkLength = types.Int64Null()
	}
	if v, ok := obj["om5-link-length"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Om5LinkLength = types.Int64Value(n)
		} else {
			m.Om5LinkLength = types.Int64Null()
		}
	} else {
		m.Om5LinkLength = types.Int64Null()
	}
	if v, ok := obj["orig-mac-address"]; ok {
		_ = v
		if v != "" {
			m.OrigMACAddress = types.StringValue(v)
		} else {
			m.OrigMACAddress = types.StringNull()
		}
	} else {
		m.OrigMACAddress = types.StringNull()
	}
	if v, ok := obj["passthrough-interface"]; ok {
		_ = v
		if v != "" {
			m.PassthroughInterface = types.StringValue(v)
		} else {
			m.PassthroughInterface = types.StringNull()
		}
	} else {
		m.PassthroughInterface = types.StringNull()
	}
	if v, ok := obj["pcie-passthrough"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PciePassthrough = types.Int64Value(n)
		} else {
			m.PciePassthrough = types.Int64Null()
		}
	} else {
		m.PciePassthrough = types.Int64Null()
	}
	if v, ok := obj["po-e-out"]; ok {
		_ = v
		if v != "" {
			m.PoEOut = types.StringValue(v)
		} else {
			m.PoEOut = types.StringNull()
		}
	} else {
		m.PoEOut = types.StringNull()
	}
	if v, ok := obj["po-e-out-current"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PoEOutCurrent = types.Int64Value(n)
		} else {
			m.PoEOutCurrent = types.Int64Null()
		}
	} else {
		m.PoEOutCurrent = types.Int64Null()
	}
	if v, ok := obj["po-e-out-power"]; ok {
		_ = v
		if v != "" {
			m.PoEOutPower = types.StringValue(v)
		} else {
			m.PoEOutPower = types.StringNull()
		}
	} else {
		m.PoEOutPower = types.StringNull()
	}
	if v, ok := obj["po-e-out-status"]; ok {
		_ = v
		if v != "" {
			m.PoEOutStatus = types.StringValue(v)
		} else {
			m.PoEOutStatus = types.StringNull()
		}
	} else {
		m.PoEOutStatus = types.StringNull()
	}
	if v, ok := obj["po-e-out-voltage"]; ok {
		_ = v
		if v != "" {
			m.PoEOutVoltage = types.StringValue(v)
		} else {
			m.PoEOutVoltage = types.StringNull()
		}
	} else {
		m.PoEOutVoltage = types.StringNull()
	}
	if v, ok := obj["po-e-priority"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PoEPriority = types.Int64Value(n)
		} else {
			m.PoEPriority = types.Int64Null()
		}
	} else {
		m.PoEPriority = types.Int64Null()
	}
	if v, ok := obj["po-e-voltage"]; ok {
		_ = v
		if v != "" {
			m.PoEVoltage = types.StringValue(v)
		} else {
			m.PoEVoltage = types.StringNull()
		}
	} else {
		m.PoEVoltage = types.StringNull()
	}
	if v, ok := obj["poe"]; ok {
		_ = v
		if v != "" {
			m.Poe = types.StringValue(v)
		} else {
			m.Poe = types.StringNull()
		}
	} else {
		m.Poe = types.StringNull()
	}
	if v, ok := obj["poe-v"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.PoeV = types.BoolValue(b)
		} else {
			m.PoeV = types.BoolNull()
		}
	} else {
		m.PoeV = types.BoolNull()
	}
	if v, ok := obj["poecurr"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Poecurr = types.Int64Value(n)
		} else {
			m.Poecurr = types.Int64Null()
		}
	} else {
		m.Poecurr = types.Int64Null()
	}
	if v, ok := obj["poeping"]; ok {
		_ = v
		if v != "" {
			m.Poeping = types.StringValue(v)
		} else {
			m.Poeping = types.StringNull()
		}
	} else {
		m.Poeping = types.StringNull()
	}
	if v, ok := obj["poepower"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Poepower = types.Int64Value(n)
		} else {
			m.Poepower = types.Int64Null()
		}
	} else {
		m.Poepower = types.Int64Null()
	}
	if v, ok := obj["poevolt"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Poevolt = types.Int64Value(n)
		} else {
			m.Poevolt = types.Int64Null()
		}
	} else {
		m.Poevolt = types.Int64Null()
	}
	if v, ok := obj["power-class"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PowerClass = types.Int64Value(n)
		} else {
			m.PowerClass = types.Int64Null()
		}
	} else {
		m.PowerClass = types.Int64Null()
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
	if v, ok := obj["power-cycle-after"]; ok {
		_ = v
		if v != "" {
			m.PowerCycleAfter = types.StringValue(v)
		} else {
			m.PowerCycleAfter = types.StringNull()
		}
	} else {
		m.PowerCycleAfter = types.StringNull()
	}
	if v, ok := obj["power-cycle-host-alive"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.PowerCycleHostAlive = types.BoolValue(b)
		} else {
			m.PowerCycleHostAlive = types.BoolNull()
		}
	} else {
		m.PowerCycleHostAlive = types.BoolNull()
	}
	if v, ok := obj["power-cycle-interval"]; ok {
		_ = v
		if v != "" {
			m.PowerCycleInterval = types.StringValue(v)
		} else {
			m.PowerCycleInterval = types.StringNull()
		}
	} else {
		m.PowerCycleInterval = types.StringNull()
	}
	if v, ok := obj["power-cycle-ping-address"]; ok {
		_ = v
		if v != "" {
			m.PowerCyclePingAddress = types.StringValue(v)
		} else {
			m.PowerCyclePingAddress = types.StringNull()
		}
	} else {
		m.PowerCyclePingAddress = types.StringNull()
	}
	if v, ok := obj["power-cycle-ping-enabled"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.PowerCyclePingEnabled = types.BoolValue(b)
		} else {
			m.PowerCyclePingEnabled = types.BoolNull()
		}
	} else {
		m.PowerCyclePingEnabled = types.BoolNull()
	}
	if v, ok := obj["power-cycle-ping-timeout"]; ok {
		_ = v
		if v != "" {
			m.PowerCyclePingTimeout = types.StringValue(v)
		} else {
			m.PowerCyclePingTimeout = types.StringNull()
		}
	} else {
		m.PowerCyclePingTimeout = types.StringNull()
	}
	if v, ok := obj["qstats"]; ok {
		_ = v
		if v != "" {
			m.Qstats = types.StringValue(v)
		} else {
			m.Qstats = types.StringNull()
		}
	} else {
		m.Qstats = types.StringNull()
	}
	if v, ok := obj["rate"]; ok {
		_ = v
		if v != "" {
			m.Rate = types.StringValue(v)
		} else {
			m.Rate = types.StringNull()
		}
	} else {
		m.Rate = types.StringNull()
	}
	if v, ok := obj["rate-select"]; ok {
		_ = v
		if v != "" {
			m.RateSelect = types.StringValue(v)
		} else {
			m.RateSelect = types.StringNull()
		}
	} else {
		m.RateSelect = types.StringNull()
	}
	if v, ok := obj["reset-counters"]; ok {
		_ = v
		if v != "" {
			m.ResetCounters = types.StringValue(v)
		} else {
			m.ResetCounters = types.StringNull()
		}
	} else {
		m.ResetCounters = types.StringNull()
	}
	if v, ok := obj["reset-mac-address"]; ok {
		_ = v
		if v != "" {
			m.ResetMACAddress = types.StringValue(v)
		} else {
			m.ResetMACAddress = types.StringNull()
		}
	} else {
		m.ResetMACAddress = types.StringNull()
	}
	if v, ok := obj["running"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Running = types.BoolValue(b)
		} else {
			m.Running = types.BoolNull()
		}
	} else {
		m.Running = types.BoolNull()
	}
	if v, ok := obj["rx-align-error"]; ok {
		_ = v
		if v != "" {
			m.RxAlignError = types.StringValue(v)
		} else {
			m.RxAlignError = types.StringNull()
		}
	} else {
		m.RxAlignError = types.StringNull()
	}
	if v, ok := obj["rx-broadcast"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RxBroadcast = types.Int64Value(n)
		} else {
			m.RxBroadcast = types.Int64Null()
		}
	} else {
		m.RxBroadcast = types.Int64Null()
	}
	if v, ok := obj["rx-bytes"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RxBytes = types.Int64Value(n)
		} else {
			m.RxBytes = types.Int64Null()
		}
	} else {
		m.RxBytes = types.Int64Null()
	}
	if v, ok := obj["rx-carrier-error"]; ok {
		_ = v
		if v != "" {
			m.RxCarrierError = types.StringValue(v)
		} else {
			m.RxCarrierError = types.StringNull()
		}
	} else {
		m.RxCarrierError = types.StringNull()
	}
	if v, ok := obj["rx-code-error"]; ok {
		_ = v
		if v != "" {
			m.RxCodeError = types.StringValue(v)
		} else {
			m.RxCodeError = types.StringNull()
		}
	} else {
		m.RxCodeError = types.StringNull()
	}
	if v, ok := obj["rx-control"]; ok {
		_ = v
		if v != "" {
			m.RxControl = types.StringValue(v)
		} else {
			m.RxControl = types.StringNull()
		}
	} else {
		m.RxControl = types.StringNull()
	}
	if v, ok := obj["rx-drop"]; ok {
		_ = v
		if v != "" {
			m.RxDrop = types.StringValue(v)
		} else {
			m.RxDrop = types.StringNull()
		}
	} else {
		m.RxDrop = types.StringNull()
	}
	if v, ok := obj["rx-error-events"]; ok {
		_ = v
		if v != "" {
			m.RxErrorEvents = types.StringValue(v)
		} else {
			m.RxErrorEvents = types.StringNull()
		}
	} else {
		m.RxErrorEvents = types.StringNull()
	}
	if v, ok := obj["rx-fcs-error"]; ok {
		_ = v
		if v != "" {
			m.RxFcsError = types.StringValue(v)
		} else {
			m.RxFcsError = types.StringNull()
		}
	} else {
		m.RxFcsError = types.StringNull()
	}
	if v, ok := obj["rx-flow-control"]; ok {
		_ = v
		if v != "" {
			m.RxFlowControl = types.StringValue(v)
		} else {
			m.RxFlowControl = types.StringNull()
		}
	} else {
		m.RxFlowControl = types.StringNull()
	}
	if v, ok := obj["rx-fragment"]; ok {
		_ = v
		if v != "" {
			m.RxFragment = types.StringValue(v)
		} else {
			m.RxFragment = types.StringNull()
		}
	} else {
		m.RxFragment = types.StringNull()
	}
	if v, ok := obj["rx-jabber"]; ok {
		_ = v
		if v != "" {
			m.RxJabber = types.StringValue(v)
		} else {
			m.RxJabber = types.StringNull()
		}
	} else {
		m.RxJabber = types.StringNull()
	}
	if v, ok := obj["rx-length-error"]; ok {
		_ = v
		if v != "" {
			m.RxLengthError = types.StringValue(v)
		} else {
			m.RxLengthError = types.StringNull()
		}
	} else {
		m.RxLengthError = types.StringNull()
	}
	if v, ok := obj["rx-loss"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.RxLoss = types.BoolValue(b)
		} else {
			m.RxLoss = types.BoolNull()
		}
	} else {
		m.RxLoss = types.BoolNull()
	}
	if v, ok := obj["rx-multicast"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RxMulticast = types.Int64Value(n)
		} else {
			m.RxMulticast = types.Int64Null()
		}
	} else {
		m.RxMulticast = types.Int64Null()
	}
	if v, ok := obj["rx-overflow"]; ok {
		_ = v
		if v != "" {
			m.RxOverflow = types.StringValue(v)
		} else {
			m.RxOverflow = types.StringNull()
		}
	} else {
		m.RxOverflow = types.StringNull()
	}
	if v, ok := obj["rx-packet"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RxPacket = types.Int64Value(n)
		} else {
			m.RxPacket = types.Int64Null()
		}
	} else {
		m.RxPacket = types.Int64Null()
	}
	if v, ok := obj["rx-pause"]; ok {
		_ = v
		if v != "" {
			m.RxPause = types.StringValue(v)
		} else {
			m.RxPause = types.StringNull()
		}
	} else {
		m.RxPause = types.StringNull()
	}
	if v, ok := obj["rx-power"]; ok {
		_ = v
		if v != "" {
			m.RxPower = types.StringValue(v)
		} else {
			m.RxPower = types.StringNull()
		}
	} else {
		m.RxPower = types.StringNull()
	}
	if v, ok := obj["rx-too-long"]; ok {
		_ = v
		if v != "" {
			m.RxTooLong = types.StringValue(v)
		} else {
			m.RxTooLong = types.StringNull()
		}
	} else {
		m.RxTooLong = types.StringNull()
	}
	if v, ok := obj["rx-too-short"]; ok {
		_ = v
		if v != "" {
			m.RxTooShort = types.StringValue(v)
		} else {
			m.RxTooShort = types.StringNull()
		}
	} else {
		m.RxTooShort = types.StringNull()
	}
	if v, ok := obj["rx-unicast"]; ok {
		_ = v
		if v != "" {
			m.RxUnicast = types.StringValue(v)
		} else {
			m.RxUnicast = types.StringNull()
		}
	} else {
		m.RxUnicast = types.StringNull()
	}
	if v, ok := obj["rx-unknown-op"]; ok {
		_ = v
		if v != "" {
			m.RxUnknownOp = types.StringValue(v)
		} else {
			m.RxUnknownOp = types.StringNull()
		}
	} else {
		m.RxUnknownOp = types.StringNull()
	}
	if v, ok := obj["send-interval"]; ok {
		_ = v
		if v != "" {
			m.SendInterval = types.StringValue(v)
		} else {
			m.SendInterval = types.StringNull()
		}
	} else {
		m.SendInterval = types.StringNull()
	}
	if v, ok := obj["sfp"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Sfp = types.BoolValue(b)
		} else {
			m.Sfp = types.BoolNull()
		}
	} else {
		m.Sfp = types.BoolNull()
	}
	if v, ok := obj["sfp-shutdown-temperature"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.SfpShutdownTemperature = types.Int64Value(n)
		} else {
			m.SfpShutdownTemperature = types.Int64Null()
		}
	} else {
		m.SfpShutdownTemperature = types.Int64Null()
	}
	if v, ok := obj["sfp-supported"]; ok {
		_ = v
		if v != "" {
			m.SfpSupported = types.StringValue(v)
		} else {
			m.SfpSupported = types.StringNull()
		}
	} else {
		m.SfpSupported = types.StringNull()
	}
	if v, ok := obj["sfprate"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Sfprate = types.Int64Value(n)
		} else {
			m.Sfprate = types.Int64Null()
		}
	} else {
		m.Sfprate = types.Int64Null()
	}
	if v, ok := obj["sfpshutdown"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Sfpshutdown = types.BoolValue(b)
		} else {
			m.Sfpshutdown = types.BoolNull()
		}
	} else {
		m.Sfpshutdown = types.BoolNull()
	}
	if v, ok := obj["sm-link-length"]; ok {
		_ = v
		if v != "" {
			m.SmLinkLength = types.StringValue(v)
		} else {
			m.SmLinkLength = types.StringNull()
		}
	} else {
		m.SmLinkLength = types.StringNull()
	}
	if v, ok := obj["speed"]; ok {
		_ = v
		if v != "" {
			m.Speed = types.StringValue(v)
		} else {
			m.Speed = types.StringNull()
		}
	} else {
		m.Speed = types.StringNull()
	}
	if v, ok := obj["status"]; ok {
		_ = v
		if v != "" {
			m.Status = types.StringValue(v)
		} else {
			m.Status = types.StringNull()
		}
	} else {
		m.Status = types.StringNull()
	}
	if v, ok := obj["supply-voltage"]; ok {
		_ = v
		if v != "" {
			m.SupplyVoltage = types.StringValue(v)
		} else {
			m.SupplyVoltage = types.StringNull()
		}
	} else {
		m.SupplyVoltage = types.StringNull()
	}
	if v, ok := obj["supported"]; ok {
		_ = v
		if v != "" {
			m.Supported = types.StringValue(v)
		} else {
			m.Supported = types.StringNull()
		}
	} else {
		m.Supported = types.StringNull()
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
	if v, ok := obj["tx-bias-current"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TxBiasCurrent = types.Int64Value(n)
		} else {
			m.TxBiasCurrent = types.Int64Null()
		}
	} else {
		m.TxBiasCurrent = types.Int64Null()
	}
	if v, ok := obj["tx-broadcast"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TxBroadcast = types.Int64Value(n)
		} else {
			m.TxBroadcast = types.Int64Null()
		}
	} else {
		m.TxBroadcast = types.Int64Null()
	}
	if v, ok := obj["tx-bytes"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TxBytes = types.Int64Value(n)
		} else {
			m.TxBytes = types.Int64Null()
		}
	} else {
		m.TxBytes = types.Int64Null()
	}
	if v, ok := obj["tx-collision"]; ok {
		_ = v
		if v != "" {
			m.TxCollision = types.StringValue(v)
		} else {
			m.TxCollision = types.StringNull()
		}
	} else {
		m.TxCollision = types.StringNull()
	}
	if v, ok := obj["tx-control"]; ok {
		_ = v
		if v != "" {
			m.TxControl = types.StringValue(v)
		} else {
			m.TxControl = types.StringNull()
		}
	} else {
		m.TxControl = types.StringNull()
	}
	if v, ok := obj["tx-deferred"]; ok {
		_ = v
		if v != "" {
			m.TxDeferred = types.StringValue(v)
		} else {
			m.TxDeferred = types.StringNull()
		}
	} else {
		m.TxDeferred = types.StringNull()
	}
	if v, ok := obj["tx-drop"]; ok {
		_ = v
		if v != "" {
			m.TxDrop = types.StringValue(v)
		} else {
			m.TxDrop = types.StringNull()
		}
	} else {
		m.TxDrop = types.StringNull()
	}
	if v, ok := obj["tx-excessive-collision"]; ok {
		_ = v
		if v != "" {
			m.TxExcessiveCollision = types.StringValue(v)
		} else {
			m.TxExcessiveCollision = types.StringNull()
		}
	} else {
		m.TxExcessiveCollision = types.StringNull()
	}
	if v, ok := obj["tx-excessive-deferred"]; ok {
		_ = v
		if v != "" {
			m.TxExcessiveDeferred = types.StringValue(v)
		} else {
			m.TxExcessiveDeferred = types.StringNull()
		}
	} else {
		m.TxExcessiveDeferred = types.StringNull()
	}
	if v, ok := obj["tx-fault"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.TxFault = types.BoolValue(b)
		} else {
			m.TxFault = types.BoolNull()
		}
	} else {
		m.TxFault = types.BoolNull()
	}
	if v, ok := obj["tx-fcs-error"]; ok {
		_ = v
		if v != "" {
			m.TxFcsError = types.StringValue(v)
		} else {
			m.TxFcsError = types.StringNull()
		}
	} else {
		m.TxFcsError = types.StringNull()
	}
	if v, ok := obj["tx-flow-control"]; ok {
		_ = v
		if v != "" {
			m.TxFlowControl = types.StringValue(v)
		} else {
			m.TxFlowControl = types.StringNull()
		}
	} else {
		m.TxFlowControl = types.StringNull()
	}
	if v, ok := obj["tx-fragment"]; ok {
		_ = v
		if v != "" {
			m.TxFragment = types.StringValue(v)
		} else {
			m.TxFragment = types.StringNull()
		}
	} else {
		m.TxFragment = types.StringNull()
	}
	if v, ok := obj["tx-jabber"]; ok {
		_ = v
		if v != "" {
			m.TxJabber = types.StringValue(v)
		} else {
			m.TxJabber = types.StringNull()
		}
	} else {
		m.TxJabber = types.StringNull()
	}
	if v, ok := obj["tx-late-collision"]; ok {
		_ = v
		if v != "" {
			m.TxLateCollision = types.StringValue(v)
		} else {
			m.TxLateCollision = types.StringNull()
		}
	} else {
		m.TxLateCollision = types.StringNull()
	}
	if v, ok := obj["tx-multicast"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TxMulticast = types.Int64Value(n)
		} else {
			m.TxMulticast = types.Int64Null()
		}
	} else {
		m.TxMulticast = types.Int64Null()
	}
	if v, ok := obj["tx-multiple-collision"]; ok {
		_ = v
		if v != "" {
			m.TxMultipleCollision = types.StringValue(v)
		} else {
			m.TxMultipleCollision = types.StringNull()
		}
	} else {
		m.TxMultipleCollision = types.StringNull()
	}
	if v, ok := obj["tx-packet"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TxPacket = types.Int64Value(n)
		} else {
			m.TxPacket = types.Int64Null()
		}
	} else {
		m.TxPacket = types.Int64Null()
	}
	if v, ok := obj["tx-pause"]; ok {
		_ = v
		if v != "" {
			m.TxPause = types.StringValue(v)
		} else {
			m.TxPause = types.StringNull()
		}
	} else {
		m.TxPause = types.StringNull()
	}
	if v, ok := obj["tx-pause-honorred"]; ok {
		_ = v
		if v != "" {
			m.TxPauseHonorred = types.StringValue(v)
		} else {
			m.TxPauseHonorred = types.StringNull()
		}
	} else {
		m.TxPauseHonorred = types.StringNull()
	}
	if v, ok := obj["tx-power"]; ok {
		_ = v
		if v != "" {
			m.TxPower = types.StringValue(v)
		} else {
			m.TxPower = types.StringNull()
		}
	} else {
		m.TxPower = types.StringNull()
	}
	if v, ok := obj["tx-rx-1024-1518"]; ok {
		_ = v
		if v != "" {
			m.TxRx10241518 = types.StringValue(v)
		} else {
			m.TxRx10241518 = types.StringNull()
		}
	} else {
		m.TxRx10241518 = types.StringNull()
	}
	if v, ok := obj["tx-rx-1024-max"]; ok {
		_ = v
		if v != "" {
			m.TxRx1024Max = types.StringValue(v)
		} else {
			m.TxRx1024Max = types.StringNull()
		}
	} else {
		m.TxRx1024Max = types.StringNull()
	}
	if v, ok := obj["tx-rx-128-255"]; ok {
		_ = v
		if v != "" {
			m.TxRx128255 = types.StringValue(v)
		} else {
			m.TxRx128255 = types.StringNull()
		}
	} else {
		m.TxRx128255 = types.StringNull()
	}
	if v, ok := obj["tx-rx-1519-max"]; ok {
		_ = v
		if v != "" {
			m.TxRx1519Max = types.StringValue(v)
		} else {
			m.TxRx1519Max = types.StringNull()
		}
	} else {
		m.TxRx1519Max = types.StringNull()
	}
	if v, ok := obj["tx-rx-256-511"]; ok {
		_ = v
		if v != "" {
			m.TxRx256511 = types.StringValue(v)
		} else {
			m.TxRx256511 = types.StringNull()
		}
	} else {
		m.TxRx256511 = types.StringNull()
	}
	if v, ok := obj["tx-rx-512-1023"]; ok {
		_ = v
		if v != "" {
			m.TxRx5121023 = types.StringValue(v)
		} else {
			m.TxRx5121023 = types.StringNull()
		}
	} else {
		m.TxRx5121023 = types.StringNull()
	}
	if v, ok := obj["tx-rx-64"]; ok {
		_ = v
		if v != "" {
			m.TxRx64 = types.StringValue(v)
		} else {
			m.TxRx64 = types.StringNull()
		}
	} else {
		m.TxRx64 = types.StringNull()
	}
	if v, ok := obj["tx-rx-65-127"]; ok {
		_ = v
		if v != "" {
			m.TxRx65127 = types.StringValue(v)
		} else {
			m.TxRx65127 = types.StringNull()
		}
	} else {
		m.TxRx65127 = types.StringNull()
	}
	if v, ok := obj["tx-rx-bytes"]; ok {
		_ = v
		if v != "" {
			m.TxRxBytes = types.StringValue(v)
		} else {
			m.TxRxBytes = types.StringNull()
		}
	} else {
		m.TxRxBytes = types.StringNull()
	}
	if v, ok := obj["tx-rx-packets"]; ok {
		_ = v
		if v != "" {
			m.TxRxPackets = types.StringValue(v)
		} else {
			m.TxRxPackets = types.StringNull()
		}
	} else {
		m.TxRxPackets = types.StringNull()
	}
	if v, ok := obj["tx-single-collision"]; ok {
		_ = v
		if v != "" {
			m.TxSingleCollision = types.StringValue(v)
		} else {
			m.TxSingleCollision = types.StringNull()
		}
	} else {
		m.TxSingleCollision = types.StringNull()
	}
	if v, ok := obj["tx-too-short"]; ok {
		_ = v
		if v != "" {
			m.TxTooShort = types.StringValue(v)
		} else {
			m.TxTooShort = types.StringNull()
		}
	} else {
		m.TxTooShort = types.StringNull()
	}
	if v, ok := obj["tx-total-collision"]; ok {
		_ = v
		if v != "" {
			m.TxTotalCollision = types.StringValue(v)
		} else {
			m.TxTotalCollision = types.StringNull()
		}
	} else {
		m.TxTotalCollision = types.StringNull()
	}
	if v, ok := obj["tx-underrun"]; ok {
		_ = v
		if v != "" {
			m.TxUnderrun = types.StringValue(v)
		} else {
			m.TxUnderrun = types.StringNull()
		}
	} else {
		m.TxUnderrun = types.StringNull()
	}
	if v, ok := obj["tx-unicast"]; ok {
		_ = v
		if v != "" {
			m.TxUnicast = types.StringValue(v)
		} else {
			m.TxUnicast = types.StringNull()
		}
	} else {
		m.TxUnicast = types.StringNull()
	}
	if v, ok := obj["vendor-name"]; ok {
		_ = v
		if v != "" {
			m.VendorName = types.StringValue(v)
		} else {
			m.VendorName = types.StringNull()
		}
	} else {
		m.VendorName = types.StringNull()
	}
	if v, ok := obj["vendor-part-number"]; ok {
		_ = v
		if v != "" {
			m.VendorPartNumber = types.StringValue(v)
		} else {
			m.VendorPartNumber = types.StringNull()
		}
	} else {
		m.VendorPartNumber = types.StringNull()
	}
	if v, ok := obj["vendor-revision"]; ok {
		_ = v
		if v != "" {
			m.VendorRevision = types.StringValue(v)
		} else {
			m.VendorRevision = types.StringNull()
		}
	} else {
		m.VendorRevision = types.StringNull()
	}
	if v, ok := obj["vendor-serial"]; ok {
		_ = v
		if v != "" {
			m.VendorSerial = types.StringValue(v)
		} else {
			m.VendorSerial = types.StringNull()
		}
	} else {
		m.VendorSerial = types.StringNull()
	}
	if v, ok := obj["wavelength"]; ok {
		_ = v
		if v != "" {
			m.Wavelength = types.StringValue(v)
		} else {
			m.Wavelength = types.StringNull()
		}
	} else {
		m.Wavelength = types.StringNull()
	}
}
