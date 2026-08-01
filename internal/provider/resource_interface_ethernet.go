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
	SfpRateSelect             types.String `tfsdk:"sfp_rate_select"`
	SfpIgnoreRxLos            types.String `tfsdk:"sfp_ignore_rx_los"`
	MdixEnable                types.String `tfsdk:"mdix_enable"`
	L2mtu                     types.String `tfsdk:"l2mtu"`
	Advertise                 types.Set    `tfsdk:"advertise"`
	Advertising               types.String `tfsdk:"advertising"`
	ARP                       types.String `tfsdk:"arp"`
	ARPTimeout                types.String `tfsdk:"arp_timeout"`
	AutoNegotiation           types.Bool   `tfsdk:"auto_negotiation"`
	Autoneg                   types.Bool   `tfsdk:"autoneg"`
	Blink                     types.String `tfsdk:"blink"`
	CableAssemblyLinkLength   types.String `tfsdk:"cable_assembly_link_length"`
	Bandwidth                 types.String `tfsdk:"bandwidth"`
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
	PoEOut                    types.String `tfsdk:"poe_out"`
	PoEOutCurrent             types.Int64  `tfsdk:"poe_out_current"`
	PoEOutPower               types.String `tfsdk:"poe_out_power"`
	PoEOutStatus              types.String `tfsdk:"poe_out_status"`
	PoEOutVoltage             types.String `tfsdk:"poe_out_voltage"`
	PoEPriority               types.Int64  `tfsdk:"poe_priority"`
	PoEVoltage                types.String `tfsdk:"poe_voltage"`
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
			"sfp_rate_select": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `sfp-rate-select`.",
			},
			"sfp_ignore_rx_los": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `sfp-ignore-rx-los`.",
			},
			"mdix_enable": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mdix-enable`.",
			},
			"l2mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2mtu`.",
			},
			"advertise": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "",
			},
			"advertising": schema.StringAttribute{
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
				Validators:    []validator.String{schemautil.IsDurationOrKeyword("auto")},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"auto_negotiation": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Description: "Whether the port negotiates link speed and duplex with its peer. " +
					"Disable it only when forcing `speed`/`full_duplex` on both ends.",
			},
			"autoneg": schema.BoolAttribute{
				Optional:           true,
				Computed:           true,
				Description:        "",
				DeprecationMessage: "Not a RouterOS REST property (WebFig-only spelling); use auto_negotiation instead. This attribute is read-only and ignored on write.",
			},
			"blink": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"cable_assembly_link_length": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"bandwidth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Rx/tx rate limit in `<rx>/<tx>` form, e.g. `unlimited/unlimited` (the default) or `100M/100M`.",
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
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"low-power", "power-up", "ready", "power-down", "fault"}...)},
			},
			"cmis_revision": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"combo": schema.Int64Attribute{
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
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"unknown", "sc", "lc", "optical-pigtail", "multifiber-parallel-optic-1x12", "copper-pigtail", "rj45", "no-separable-connector", "multifiber-parallel-optic-1x16"}...)},
			},
			"copper_active_om4_link_length": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"default_name": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"disable_running_check": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disable_time": schema.StringAttribute{
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
				Computed:    true,
				Description: "",
			},
			"full_duplex": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"hastxqueuestats": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"ignore_rx_los": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"link_partner_advertising": schema.StringAttribute{
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
				Computed:    true,
				Description: "",
			},
			"max_l2_mtu": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"max_power": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"module_present": schema.BoolAttribute{
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
				Optional:           true,
				Computed:           true,
				Description:        "",
				DeprecationMessage: "Not a RouterOS REST property (WebFig-only spelling); use auto_negotiation instead. This attribute is read-only and ignored on write.",
			},
			"non_mgmt": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"om1_link_length": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"om2_link_length": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"om3_link_length": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"om4_link_length": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"om5_link_length": schema.Int64Attribute{
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
				Computed:    true,
				Description: "",
			},
			"pcie_passthrough": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"poe_out": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"off", "auto-on", "forced-on"}...)},
			},
			"poe_out_current": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"poe_out_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"poe_out_status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "disabled", "waiting-for-load", "powered-on", "overload", "short-circuit", "voltage-too-low", "current-too-low", "power-cycle", "voltage-too-high", "controller-error", "controller-upgrade", "voltage-on-poe-in", "no-valid-psu"}...)},
			},
			"poe_out_voltage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"poe_priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"poe_voltage": schema.StringAttribute{
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
				Computed:    true,
				Description: "",
			},
			"poecurr": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"poeping": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"poepower": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"poevolt": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"power_class": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"power_cycle": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"power_cycle_after": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"power_cycle_host_alive": schema.BoolAttribute{
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
				Computed:    true,
				Description: "",
			},
			"power_cycle_ping_enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"power_cycle_ping_timeout": schema.StringAttribute{
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
				Computed:    true,
				Description: "",
			},
			"reset_mac_address": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"running": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_align_error": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_broadcast": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"rx_bytes": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"rx_carrier_error": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_code_error": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_control": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_drop": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_error_events": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_fcs_error": schema.StringAttribute{
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
				Computed:    true,
				Description: "",
			},
			"rx_jabber": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_length_error": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_loss": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_multicast": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"rx_overflow": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_packet": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"rx_pause": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_power": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_too_long": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_too_short": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_unicast": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_unknown_op": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"send_interval": schema.StringAttribute{
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
				Computed:    true,
				Description: "",
			},
			"sfprate": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"sfpshutdown": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"sm_link_length": schema.StringAttribute{
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
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "off", "on", "disabled"}...)},
			},
			"supply_voltage": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"supported": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"temperature": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_bias_current": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"tx_broadcast": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"tx_bytes": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"tx_collision": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_control": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_deferred": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_drop": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_excessive_collision": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_excessive_deferred": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_fault": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_fcs_error": schema.StringAttribute{
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
				Computed:    true,
				Description: "",
			},
			"tx_jabber": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_late_collision": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_multicast": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"tx_multiple_collision": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_packet": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"tx_pause": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_pause_honorred": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_power": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_rx_1024_1518": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_rx_1024_max": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_rx_128_255": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_rx_1519_max": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_rx_256_511": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_rx_512_1023": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_rx_64": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_rx_65_127": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_rx_bytes": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_rx_packets": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_single_collision": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_too_short": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_total_collision": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_underrun": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_unicast": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"vendor_name": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"vendor_part_number": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"vendor_revision": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"vendor_serial": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"wavelength": schema.StringAttribute{
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
		body["advertise"] = encodeStringSet(ctx, plan.Advertise, &resp.Diagnostics)
	}
	if !(plan.ARP.IsNull() || plan.ARP.IsUnknown()) {
		body["arp"] = plan.ARP.ValueString()
	}
	if !(plan.ARPTimeout.IsNull() || plan.ARPTimeout.IsUnknown()) {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !(plan.AutoNegotiation.IsNull() || plan.AutoNegotiation.IsUnknown()) {
		body["auto-negotiation"] = client.FormatBool(plan.AutoNegotiation.ValueBool())
	}
	if !(plan.Blink.IsNull() || plan.Blink.IsUnknown()) {
		body["blink"] = plan.Blink.ValueString()
	}
	if !(plan.Bandwidth.IsNull() || plan.Bandwidth.IsUnknown()) {
		body["bandwidth"] = plan.Bandwidth.ValueString()
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

	if !(plan.OrigMACAddress.IsNull() || plan.OrigMACAddress.IsUnknown()) {
		body["orig-mac-address"] = plan.OrigMACAddress.ValueString()
	}
	if !(plan.PoEOut.IsNull() || plan.PoEOut.IsUnknown()) {
		body["poe-out"] = plan.PoEOut.ValueString()
	}
	if !(plan.PoEPriority.IsNull() || plan.PoEPriority.IsUnknown()) {
		body["poe-priority"] = client.FormatInt64(plan.PoEPriority.ValueInt64())
	}
	if !(plan.PoEVoltage.IsNull() || plan.PoEVoltage.IsUnknown()) {
		body["poe-voltage"] = plan.PoEVoltage.ValueString()
	}
	if !(plan.Poe.IsNull() || plan.Poe.IsUnknown()) {
		body["poe"] = plan.Poe.ValueString()
	}
	if !(plan.Poeping.IsNull() || plan.Poeping.IsUnknown()) {
		body["poeping"] = plan.Poeping.ValueString()
	}
	if !(plan.PowerCycleInterval.IsNull() || plan.PowerCycleInterval.IsUnknown()) {
		body["power-cycle-interval"] = plan.PowerCycleInterval.ValueString()
	}
	if !(plan.PowerCyclePingEnabled.IsNull() || plan.PowerCyclePingEnabled.IsUnknown()) {
		body["power-cycle-ping-enabled"] = client.FormatBool(plan.PowerCyclePingEnabled.ValueBool())
	}
	if !(plan.Qstats.IsNull() || plan.Qstats.IsUnknown()) {
		body["qstats"] = plan.Qstats.ValueString()
	}
	if !(plan.RateSelect.IsNull() || plan.RateSelect.IsUnknown()) {
		body["rate-select"] = plan.RateSelect.ValueString()
	}
	if !(plan.RxFlowControl.IsNull() || plan.RxFlowControl.IsUnknown()) {
		body["rx-flow-control"] = plan.RxFlowControl.ValueString()
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
	if !(plan.L2mtu.IsNull() || plan.L2mtu.IsUnknown()) {
		body["l2mtu"] = plan.L2mtu.ValueString()
	}
	if !(plan.MdixEnable.IsNull() || plan.MdixEnable.IsUnknown()) {
		body["mdix-enable"] = plan.MdixEnable.ValueString()
	}
	if !(plan.SfpIgnoreRxLos.IsNull() || plan.SfpIgnoreRxLos.IsUnknown()) {
		body["sfp-ignore-rx-los"] = plan.SfpIgnoreRxLos.ValueString()
	}
	if !(plan.SfpRateSelect.IsNull() || plan.SfpRateSelect.IsUnknown()) {
		body["sfp-rate-select"] = plan.SfpRateSelect.ValueString()
	}
	rows, err := c.List(ctx, "/interface/ethernet")
	if err != nil {
		resp.Diagnostics.AddError("Read /interface/ethernet failed", err.Error())
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
		resp.Diagnostics.AddError("Unknown /interface/ethernet "+want, fmt.Sprintf("/interface/ethernet is a fixed hardware row set; no row matches name %q. Import the interface instead of creating it.", want))
		return
	}
	obj, err := c.Set(ctx, "/interface/ethernet", id, body)
	if err != nil {
		resp.Diagnostics.AddError("Adopt /interface/ethernet failed", err.Error())
		return
	}
	interfaceEthernetApply(ctx, obj, &plan)
	plan.ID = types.StringValue(id)
	nullifyUnknownAttrs(&plan)
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
	if !plan.Advertise.Equal(state.Advertise) && !plan.Advertise.IsUnknown() {
		body["advertise"] = encodeStringSet(ctx, plan.Advertise, &resp.Diagnostics)
	}
	if !plan.ARP.Equal(state.ARP) && !plan.ARP.IsUnknown() {
		body["arp"] = plan.ARP.ValueString()
	}
	if !plan.ARPTimeout.Equal(state.ARPTimeout) && !plan.ARPTimeout.IsUnknown() {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !plan.AutoNegotiation.Equal(state.AutoNegotiation) && !plan.AutoNegotiation.IsUnknown() {
		body["auto-negotiation"] = client.FormatBool(plan.AutoNegotiation.ValueBool())
	}
	if !plan.Blink.Equal(state.Blink) && !plan.Blink.IsUnknown() {
		body["blink"] = plan.Blink.ValueString()
	}
	if !plan.Bandwidth.Equal(state.Bandwidth) && !plan.Bandwidth.IsUnknown() {
		body["bandwidth"] = plan.Bandwidth.ValueString()
	}
	if !plan.CableSettings.Equal(state.CableSettings) && !plan.CableSettings.IsUnknown() {
		body["cable-settings"] = plan.CableSettings.ValueString()
	}
	if !plan.CableTest.Equal(state.CableTest) && !plan.CableTest.IsUnknown() {
		body["cable-test"] = plan.CableTest.ValueString()
	}
	if !plan.ComboMode.Equal(state.ComboMode) && !plan.ComboMode.IsUnknown() {
		body["combo-mode"] = plan.ComboMode.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.DisableRunningCheck.Equal(state.DisableRunningCheck) && !plan.DisableRunningCheck.IsUnknown() {
		body["disable-running-check"] = client.FormatBool(plan.DisableRunningCheck.ValueBool())
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Extrastats.Equal(state.Extrastats) && !plan.Extrastats.IsUnknown() {
		body["extrastats"] = plan.Extrastats.ValueString()
	}
	if !plan.FecMode.Equal(state.FecMode) && !plan.FecMode.IsUnknown() {
		body["fec-mode"] = plan.FecMode.ValueString()
	}
	if !plan.Flowcntrl.Equal(state.Flowcntrl) && !plan.Flowcntrl.IsUnknown() {
		body["flowcntrl"] = plan.Flowcntrl.ValueString()
	}
	if !plan.IgnoreRxLos.Equal(state.IgnoreRxLos) && !plan.IgnoreRxLos.IsUnknown() {
		body["ignore-rx-los"] = client.FormatBool(plan.IgnoreRxLos.ValueBool())
	}
	if !plan.LoopProtect.Equal(state.LoopProtect) && !plan.LoopProtect.IsUnknown() {
		body["loop-protect"] = plan.LoopProtect.ValueString()
	}
	if !plan.LoopProtectDisableTime.Equal(state.LoopProtectDisableTime) && !plan.LoopProtectDisableTime.IsUnknown() {
		body["loop-protect-disable-time"] = plan.LoopProtectDisableTime.ValueString()
	}
	if !plan.LoopProtectSendInterval.Equal(state.LoopProtectSendInterval) && !plan.LoopProtectSendInterval.IsUnknown() {
		body["loop-protect-send-interval"] = plan.LoopProtectSendInterval.ValueString()
	}
	if !plan.MACAddress.Equal(state.MACAddress) && !plan.MACAddress.IsUnknown() {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.MTU.Equal(state.MTU) && !plan.MTU.IsUnknown() {
		body["mtu"] = client.FormatInt64(plan.MTU.ValueInt64())
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}

	if !plan.OrigMACAddress.Equal(state.OrigMACAddress) && !plan.OrigMACAddress.IsUnknown() {
		body["orig-mac-address"] = plan.OrigMACAddress.ValueString()
	}
	if !plan.PoEOut.Equal(state.PoEOut) && !plan.PoEOut.IsUnknown() {
		body["poe-out"] = plan.PoEOut.ValueString()
	}
	if !plan.PoEPriority.Equal(state.PoEPriority) && !plan.PoEPriority.IsUnknown() {
		body["poe-priority"] = client.FormatInt64(plan.PoEPriority.ValueInt64())
	}
	if !plan.PoEVoltage.Equal(state.PoEVoltage) && !plan.PoEVoltage.IsUnknown() {
		body["poe-voltage"] = plan.PoEVoltage.ValueString()
	}
	if !plan.Poe.Equal(state.Poe) && !plan.Poe.IsUnknown() {
		body["poe"] = plan.Poe.ValueString()
	}
	if !plan.Poeping.Equal(state.Poeping) && !plan.Poeping.IsUnknown() {
		body["poeping"] = plan.Poeping.ValueString()
	}
	if !plan.PowerCycleInterval.Equal(state.PowerCycleInterval) && !plan.PowerCycleInterval.IsUnknown() {
		body["power-cycle-interval"] = plan.PowerCycleInterval.ValueString()
	}
	if !plan.PowerCyclePingEnabled.Equal(state.PowerCyclePingEnabled) && !plan.PowerCyclePingEnabled.IsUnknown() {
		body["power-cycle-ping-enabled"] = client.FormatBool(plan.PowerCyclePingEnabled.ValueBool())
	}
	if !plan.Qstats.Equal(state.Qstats) && !plan.Qstats.IsUnknown() {
		body["qstats"] = plan.Qstats.ValueString()
	}
	if !plan.RateSelect.Equal(state.RateSelect) && !plan.RateSelect.IsUnknown() {
		body["rate-select"] = plan.RateSelect.ValueString()
	}
	if !plan.RxFlowControl.Equal(state.RxFlowControl) && !plan.RxFlowControl.IsUnknown() {
		body["rx-flow-control"] = plan.RxFlowControl.ValueString()
	}
	if !plan.Sfp.Equal(state.Sfp) && !plan.Sfp.IsUnknown() {
		body["sfp"] = client.FormatBool(plan.Sfp.ValueBool())
	}
	if !plan.SfpShutdownTemperature.Equal(state.SfpShutdownTemperature) && !plan.SfpShutdownTemperature.IsUnknown() {
		body["sfp-shutdown-temperature"] = client.FormatInt64(plan.SfpShutdownTemperature.ValueInt64())
	}
	if !plan.Speed.Equal(state.Speed) && !plan.Speed.IsUnknown() {
		body["speed"] = plan.Speed.ValueString()
	}
	if !plan.TxFlowControl.Equal(state.TxFlowControl) && !plan.TxFlowControl.IsUnknown() {
		body["tx-flow-control"] = plan.TxFlowControl.ValueString()
	}
	if !plan.L2mtu.Equal(state.L2mtu) && !plan.L2mtu.IsUnknown() {
		body["l2mtu"] = plan.L2mtu.ValueString()
	}
	if !plan.MdixEnable.Equal(state.MdixEnable) && !plan.MdixEnable.IsUnknown() {
		body["mdix-enable"] = plan.MdixEnable.ValueString()
	}
	if !plan.SfpIgnoreRxLos.Equal(state.SfpIgnoreRxLos) && !plan.SfpIgnoreRxLos.IsUnknown() {
		body["sfp-ignore-rx-los"] = plan.SfpIgnoreRxLos.ValueString()
	}
	if !plan.SfpRateSelect.Equal(state.SfpRateSelect) && !plan.SfpRateSelect.IsUnknown() {
		body["sfp-rate-select"] = plan.SfpRateSelect.ValueString()
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
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEthernetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Fixed hardware row: cannot be removed. Drop from state; the row keeps
	// its last-applied settings (adopt-only, like /ip/service).
	_ = ctx
	_ = req
	_ = resp
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
	if v, ok := obj["sfp-rate-select"]; ok && v != "" {
		m.SfpRateSelect = types.StringValue(v)
	} else {
		m.SfpRateSelect = types.StringNull()
	}
	if v, ok := obj["sfp-ignore-rx-los"]; ok && v != "" {
		m.SfpIgnoreRxLos = types.StringValue(v)
	} else {
		m.SfpIgnoreRxLos = types.StringNull()
	}
	if v, ok := obj["mdix-enable"]; ok && v != "" {
		m.MdixEnable = types.StringValue(v)
	} else {
		m.MdixEnable = types.StringNull()
	}
	if v, ok := obj["l2mtu"]; ok && v != "" {
		m.L2mtu = types.StringValue(v)
	} else {
		m.L2mtu = types.StringNull()
	}
	if v, ok := obj["advertise"]; ok {
		_ = v
		m.Advertise = decodeStringSet(ctx, v)
	} else {
		m.Advertise = types.SetNull(types.StringType)
	}
	if v, ok := obj["advertising"]; ok {
		if v != "" {
			m.Advertising = types.StringValue(v)
		} else {
			m.Advertising = types.StringNull()
		}
	}
	if v, ok := obj["arp"]; ok {
		if v != "" {
			m.ARP = types.StringValue(v)
		} else {
			m.ARP = types.StringNull()
		}
	}
	if v, ok := obj["arp-timeout"]; ok {
		if v != "" {
			m.ARPTimeout = types.StringValue(v)
		} else {
			m.ARPTimeout = types.StringNull()
		}
	}
	if v, ok := obj["auto-negotiation"]; ok {
		_ = v
		// Tolerate a device that answers with a link state rather than the
		// toggle; an unparseable value becomes null instead of a bogus false.
		if b, err := client.ParseBool(v); err == nil {
			m.AutoNegotiation = types.BoolValue(b)
		} else {
			m.AutoNegotiation = types.BoolNull()
		}
	} else {
		m.AutoNegotiation = types.BoolNull()
	}
	if v, ok := obj["autoneg"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Autoneg = types.BoolValue(b)
		} else {
			m.Autoneg = types.BoolNull()
		}
	}
	if v, ok := obj["blink"]; ok {
		if v != "" {
			m.Blink = types.StringValue(v)
		} else {
			m.Blink = types.StringNull()
		}
	}
	if v, ok := obj["cable-assembly-link-length"]; ok {
		if v != "" {
			m.CableAssemblyLinkLength = types.StringValue(v)
		} else {
			m.CableAssemblyLinkLength = types.StringNull()
		}
	}
	if v, ok := obj["bandwidth"]; ok {
		if v != "" {
			m.Bandwidth = types.StringValue(v)
		} else {
			m.Bandwidth = types.StringNull()
		}
	}
	if v, ok := obj["cable-settings"]; ok {
		if v != "" {
			m.CableSettings = types.StringValue(v)
		} else {
			m.CableSettings = types.StringNull()
		}
	}
	if v, ok := obj["cable-test"]; ok {
		if v != "" {
			m.CableTest = types.StringValue(v)
		} else {
			m.CableTest = types.StringNull()
		}
	}
	if v, ok := obj["cmis-module-state"]; ok {
		if v != "" {
			m.CmisModuleState = types.StringValue(v)
		} else {
			m.CmisModuleState = types.StringNull()
		}
	}
	if v, ok := obj["cmis-revision"]; ok {
		if v != "" {
			m.CmisRevision = types.StringValue(v)
		} else {
			m.CmisRevision = types.StringNull()
		}
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
		if v != "" {
			m.ComboMode = types.StringValue(v)
		} else {
			m.ComboMode = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["connector-type"]; ok {
		if v != "" {
			m.ConnectorType = types.StringValue(v)
		} else {
			m.ConnectorType = types.StringNull()
		}
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
		if v != "" {
			m.DefaultName = types.StringValue(v)
		} else {
			m.DefaultName = types.StringNull()
		}
	}
	if v, ok := obj["disable-running-check"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.DisableRunningCheck = types.BoolValue(b)
		} else {
			m.DisableRunningCheck = types.BoolNull()
		}
	}
	if v, ok := obj["disable-time"]; ok {
		if v != "" {
			m.DisableTime = types.StringValue(v)
		} else {
			m.DisableTime = types.StringNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["encoding"]; ok {
		if v != "" {
			m.Encoding = types.StringValue(v)
		} else {
			m.Encoding = types.StringNull()
		}
	}
	if v, ok := obj["extrastats"]; ok {
		if v != "" {
			m.Extrastats = types.StringValue(v)
		} else {
			m.Extrastats = types.StringNull()
		}
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
		if v != "" {
			m.FecMode = types.StringValue(v)
		} else {
			m.FecMode = types.StringNull()
		}
	}
	if v, ok := obj["flowcntrl"]; ok {
		if v != "" {
			m.Flowcntrl = types.StringValue(v)
		} else {
			m.Flowcntrl = types.StringNull()
		}
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
		if v != "" {
			m.FullDuplex = types.StringValue(v)
		} else {
			m.FullDuplex = types.StringNull()
		}
	}
	if v, ok := obj["hastxqueuestats"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Hastxqueuestats = types.BoolValue(b)
		} else {
			m.Hastxqueuestats = types.BoolNull()
		}
	}
	if v, ok := obj["ignore-rx-los"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.IgnoreRxLos = types.BoolValue(b)
		} else {
			m.IgnoreRxLos = types.BoolNull()
		}
	}
	if v, ok := obj["link-partner-advertising"]; ok {
		if v != "" {
			m.LinkPartnerAdvertising = types.StringValue(v)
		} else {
			m.LinkPartnerAdvertising = types.StringNull()
		}
	}
	if v, ok := obj["loop-protect"]; ok {
		if v != "" {
			m.LoopProtect = types.StringValue(v)
		} else {
			m.LoopProtect = types.StringNull()
		}
	}
	if v, ok := obj["loop-protect-disable-time"]; ok {
		if v != "" {
			m.LoopProtectDisableTime = types.StringValue(v)
		} else {
			m.LoopProtectDisableTime = types.StringNull()
		}
	}
	if v, ok := obj["loop-protect-send-interval"]; ok {
		if v != "" {
			m.LoopProtectSendInterval = types.StringValue(v)
		} else {
			m.LoopProtectSendInterval = types.StringNull()
		}
	}
	if v, ok := obj["loop-protect-status"]; ok {
		if v != "" {
			m.LoopProtectStatus = types.StringValue(v)
		} else {
			m.LoopProtectStatus = types.StringNull()
		}
	}
	if v, ok := obj["mac-address"]; ok {
		if v != "" {
			m.MACAddress = types.StringValue(v)
		} else {
			m.MACAddress = types.StringNull()
		}
	}
	if v, ok := obj["manufacturing-date"]; ok {
		if v != "" {
			m.ManufacturingDate = types.StringValue(v)
		} else {
			m.ManufacturingDate = types.StringNull()
		}
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
		if v != "" {
			m.MaxPower = types.StringValue(v)
		} else {
			m.MaxPower = types.StringNull()
		}
	}
	if v, ok := obj["module-present"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.ModulePresent = types.BoolValue(b)
		} else {
			m.ModulePresent = types.BoolNull()
		}
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
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["noautoneg"]; ok {
		if v != "" {
			m.Noautoneg = types.StringValue(v)
		} else {
			m.Noautoneg = types.StringNull()
		}
	}
	if v, ok := obj["non-mgmt"]; ok {
		if v != "" {
			m.NonMgmt = types.StringValue(v)
		} else {
			m.NonMgmt = types.StringNull()
		}
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
		if v != "" {
			m.OrigMACAddress = types.StringValue(v)
		} else {
			m.OrigMACAddress = types.StringNull()
		}
	}
	if v, ok := obj["passthrough-interface"]; ok {
		if v != "" {
			m.PassthroughInterface = types.StringValue(v)
		} else {
			m.PassthroughInterface = types.StringNull()
		}
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
	if v, ok := obj["poe-out"]; ok {
		if v != "" {
			m.PoEOut = types.StringValue(v)
		} else {
			m.PoEOut = types.StringNull()
		}
	}
	if v, ok := obj["poe-out-current"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PoEOutCurrent = types.Int64Value(n)
		} else {
			m.PoEOutCurrent = types.Int64Null()
		}
	} else {
		m.PoEOutCurrent = types.Int64Null()
	}
	if v, ok := obj["poe-out-power"]; ok {
		if v != "" {
			m.PoEOutPower = types.StringValue(v)
		} else {
			m.PoEOutPower = types.StringNull()
		}
	}
	if v, ok := obj["poe-out-status"]; ok {
		if v != "" {
			m.PoEOutStatus = types.StringValue(v)
		} else {
			m.PoEOutStatus = types.StringNull()
		}
	}
	if v, ok := obj["poe-out-voltage"]; ok {
		if v != "" {
			m.PoEOutVoltage = types.StringValue(v)
		} else {
			m.PoEOutVoltage = types.StringNull()
		}
	}
	if v, ok := obj["poe-priority"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PoEPriority = types.Int64Value(n)
		} else {
			m.PoEPriority = types.Int64Null()
		}
	} else {
		m.PoEPriority = types.Int64Null()
	}
	if v, ok := obj["poe-voltage"]; ok {
		if v != "" {
			m.PoEVoltage = types.StringValue(v)
		} else {
			m.PoEVoltage = types.StringNull()
		}
	}
	if v, ok := obj["poe"]; ok {
		if v != "" {
			m.Poe = types.StringValue(v)
		} else {
			m.Poe = types.StringNull()
		}
	}
	if v, ok := obj["poe-v"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.PoeV = types.BoolValue(b)
		} else {
			m.PoeV = types.BoolNull()
		}
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
		if v != "" {
			m.Poeping = types.StringValue(v)
		} else {
			m.Poeping = types.StringNull()
		}
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
		if v != "" {
			m.PowerCycle = types.StringValue(v)
		} else {
			m.PowerCycle = types.StringNull()
		}
	}
	if v, ok := obj["power-cycle-after"]; ok {
		if v != "" {
			m.PowerCycleAfter = types.StringValue(v)
		} else {
			m.PowerCycleAfter = types.StringNull()
		}
	}
	if v, ok := obj["power-cycle-host-alive"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.PowerCycleHostAlive = types.BoolValue(b)
		} else {
			m.PowerCycleHostAlive = types.BoolNull()
		}
	}
	if v, ok := obj["power-cycle-interval"]; ok {
		if v != "" {
			m.PowerCycleInterval = types.StringValue(v)
		} else {
			m.PowerCycleInterval = types.StringNull()
		}
	}
	if v, ok := obj["power-cycle-ping-address"]; ok {
		if v != "" {
			m.PowerCyclePingAddress = types.StringValue(v)
		} else {
			m.PowerCyclePingAddress = types.StringNull()
		}
	}
	if v, ok := obj["power-cycle-ping-enabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.PowerCyclePingEnabled = types.BoolValue(b)
		} else {
			m.PowerCyclePingEnabled = types.BoolNull()
		}
	}
	if v, ok := obj["power-cycle-ping-timeout"]; ok {
		if v != "" {
			m.PowerCyclePingTimeout = types.StringValue(v)
		} else {
			m.PowerCyclePingTimeout = types.StringNull()
		}
	}
	if v, ok := obj["qstats"]; ok {
		if v != "" {
			m.Qstats = types.StringValue(v)
		} else {
			m.Qstats = types.StringNull()
		}
	}
	if v, ok := obj["rate"]; ok {
		if v != "" {
			m.Rate = types.StringValue(v)
		} else {
			m.Rate = types.StringNull()
		}
	}
	if v, ok := obj["rate-select"]; ok {
		if v != "" {
			m.RateSelect = types.StringValue(v)
		} else {
			m.RateSelect = types.StringNull()
		}
	}
	if v, ok := obj["reset-counters"]; ok {
		if v != "" {
			m.ResetCounters = types.StringValue(v)
		} else {
			m.ResetCounters = types.StringNull()
		}
	}
	if v, ok := obj["reset-mac-address"]; ok {
		if v != "" {
			m.ResetMACAddress = types.StringValue(v)
		} else {
			m.ResetMACAddress = types.StringNull()
		}
	}
	if v, ok := obj["running"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Running = types.BoolValue(b)
		} else {
			m.Running = types.BoolNull()
		}
	}
	if v, ok := obj["rx-align-error"]; ok {
		if v != "" {
			m.RxAlignError = types.StringValue(v)
		} else {
			m.RxAlignError = types.StringNull()
		}
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
		if v != "" {
			m.RxCarrierError = types.StringValue(v)
		} else {
			m.RxCarrierError = types.StringNull()
		}
	}
	if v, ok := obj["rx-code-error"]; ok {
		if v != "" {
			m.RxCodeError = types.StringValue(v)
		} else {
			m.RxCodeError = types.StringNull()
		}
	}
	if v, ok := obj["rx-control"]; ok {
		if v != "" {
			m.RxControl = types.StringValue(v)
		} else {
			m.RxControl = types.StringNull()
		}
	}
	if v, ok := obj["rx-drop"]; ok {
		if v != "" {
			m.RxDrop = types.StringValue(v)
		} else {
			m.RxDrop = types.StringNull()
		}
	}
	if v, ok := obj["rx-error-events"]; ok {
		if v != "" {
			m.RxErrorEvents = types.StringValue(v)
		} else {
			m.RxErrorEvents = types.StringNull()
		}
	}
	if v, ok := obj["rx-fcs-error"]; ok {
		if v != "" {
			m.RxFcsError = types.StringValue(v)
		} else {
			m.RxFcsError = types.StringNull()
		}
	}
	if v, ok := obj["rx-flow-control"]; ok {
		if v != "" {
			m.RxFlowControl = types.StringValue(v)
		} else {
			m.RxFlowControl = types.StringNull()
		}
	}
	if v, ok := obj["rx-fragment"]; ok {
		if v != "" {
			m.RxFragment = types.StringValue(v)
		} else {
			m.RxFragment = types.StringNull()
		}
	}
	if v, ok := obj["rx-jabber"]; ok {
		if v != "" {
			m.RxJabber = types.StringValue(v)
		} else {
			m.RxJabber = types.StringNull()
		}
	}
	if v, ok := obj["rx-length-error"]; ok {
		if v != "" {
			m.RxLengthError = types.StringValue(v)
		} else {
			m.RxLengthError = types.StringNull()
		}
	}
	if v, ok := obj["rx-loss"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.RxLoss = types.BoolValue(b)
		} else {
			m.RxLoss = types.BoolNull()
		}
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
		if v != "" {
			m.RxOverflow = types.StringValue(v)
		} else {
			m.RxOverflow = types.StringNull()
		}
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
		if v != "" {
			m.RxPause = types.StringValue(v)
		} else {
			m.RxPause = types.StringNull()
		}
	}
	if v, ok := obj["rx-power"]; ok {
		if v != "" {
			m.RxPower = types.StringValue(v)
		} else {
			m.RxPower = types.StringNull()
		}
	}
	if v, ok := obj["rx-too-long"]; ok {
		if v != "" {
			m.RxTooLong = types.StringValue(v)
		} else {
			m.RxTooLong = types.StringNull()
		}
	}
	if v, ok := obj["rx-too-short"]; ok {
		if v != "" {
			m.RxTooShort = types.StringValue(v)
		} else {
			m.RxTooShort = types.StringNull()
		}
	}
	if v, ok := obj["rx-unicast"]; ok {
		if v != "" {
			m.RxUnicast = types.StringValue(v)
		} else {
			m.RxUnicast = types.StringNull()
		}
	}
	if v, ok := obj["rx-unknown-op"]; ok {
		if v != "" {
			m.RxUnknownOp = types.StringValue(v)
		} else {
			m.RxUnknownOp = types.StringNull()
		}
	}
	if v, ok := obj["send-interval"]; ok {
		if v != "" {
			m.SendInterval = types.StringValue(v)
		} else {
			m.SendInterval = types.StringNull()
		}
	}
	if v, ok := obj["sfp"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Sfp = types.BoolValue(b)
		} else {
			m.Sfp = types.BoolNull()
		}
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
		if v != "" {
			m.SfpSupported = types.StringValue(v)
		} else {
			m.SfpSupported = types.StringNull()
		}
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
		if b, err := client.ParseBool(v); err == nil {
			m.Sfpshutdown = types.BoolValue(b)
		} else {
			m.Sfpshutdown = types.BoolNull()
		}
	}
	if v, ok := obj["sm-link-length"]; ok {
		if v != "" {
			m.SmLinkLength = types.StringValue(v)
		} else {
			m.SmLinkLength = types.StringNull()
		}
	}
	if v, ok := obj["speed"]; ok {
		if v != "" {
			m.Speed = types.StringValue(v)
		} else {
			m.Speed = types.StringNull()
		}
	}
	if v, ok := obj["status"]; ok {
		if v != "" {
			m.Status = types.StringValue(v)
		} else {
			m.Status = types.StringNull()
		}
	}
	if v, ok := obj["supply-voltage"]; ok {
		if v != "" {
			m.SupplyVoltage = types.StringValue(v)
		} else {
			m.SupplyVoltage = types.StringNull()
		}
	}
	if v, ok := obj["supported"]; ok {
		if v != "" {
			m.Supported = types.StringValue(v)
		} else {
			m.Supported = types.StringNull()
		}
	}
	if v, ok := obj["temperature"]; ok {
		if v != "" {
			m.Temperature = types.StringValue(v)
		} else {
			m.Temperature = types.StringNull()
		}
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
		if v != "" {
			m.TxCollision = types.StringValue(v)
		} else {
			m.TxCollision = types.StringNull()
		}
	}
	if v, ok := obj["tx-control"]; ok {
		if v != "" {
			m.TxControl = types.StringValue(v)
		} else {
			m.TxControl = types.StringNull()
		}
	}
	if v, ok := obj["tx-deferred"]; ok {
		if v != "" {
			m.TxDeferred = types.StringValue(v)
		} else {
			m.TxDeferred = types.StringNull()
		}
	}
	if v, ok := obj["tx-drop"]; ok {
		if v != "" {
			m.TxDrop = types.StringValue(v)
		} else {
			m.TxDrop = types.StringNull()
		}
	}
	if v, ok := obj["tx-excessive-collision"]; ok {
		if v != "" {
			m.TxExcessiveCollision = types.StringValue(v)
		} else {
			m.TxExcessiveCollision = types.StringNull()
		}
	}
	if v, ok := obj["tx-excessive-deferred"]; ok {
		if v != "" {
			m.TxExcessiveDeferred = types.StringValue(v)
		} else {
			m.TxExcessiveDeferred = types.StringNull()
		}
	}
	if v, ok := obj["tx-fault"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.TxFault = types.BoolValue(b)
		} else {
			m.TxFault = types.BoolNull()
		}
	}
	if v, ok := obj["tx-fcs-error"]; ok {
		if v != "" {
			m.TxFcsError = types.StringValue(v)
		} else {
			m.TxFcsError = types.StringNull()
		}
	}
	if v, ok := obj["tx-flow-control"]; ok {
		if v != "" {
			m.TxFlowControl = types.StringValue(v)
		} else {
			m.TxFlowControl = types.StringNull()
		}
	}
	if v, ok := obj["tx-fragment"]; ok {
		if v != "" {
			m.TxFragment = types.StringValue(v)
		} else {
			m.TxFragment = types.StringNull()
		}
	}
	if v, ok := obj["tx-jabber"]; ok {
		if v != "" {
			m.TxJabber = types.StringValue(v)
		} else {
			m.TxJabber = types.StringNull()
		}
	}
	if v, ok := obj["tx-late-collision"]; ok {
		if v != "" {
			m.TxLateCollision = types.StringValue(v)
		} else {
			m.TxLateCollision = types.StringNull()
		}
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
		if v != "" {
			m.TxMultipleCollision = types.StringValue(v)
		} else {
			m.TxMultipleCollision = types.StringNull()
		}
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
		if v != "" {
			m.TxPause = types.StringValue(v)
		} else {
			m.TxPause = types.StringNull()
		}
	}
	if v, ok := obj["tx-pause-honorred"]; ok {
		if v != "" {
			m.TxPauseHonorred = types.StringValue(v)
		} else {
			m.TxPauseHonorred = types.StringNull()
		}
	}
	if v, ok := obj["tx-power"]; ok {
		if v != "" {
			m.TxPower = types.StringValue(v)
		} else {
			m.TxPower = types.StringNull()
		}
	}
	if v, ok := obj["tx-rx-1024-1518"]; ok {
		if v != "" {
			m.TxRx10241518 = types.StringValue(v)
		} else {
			m.TxRx10241518 = types.StringNull()
		}
	}
	if v, ok := obj["tx-rx-1024-max"]; ok {
		if v != "" {
			m.TxRx1024Max = types.StringValue(v)
		} else {
			m.TxRx1024Max = types.StringNull()
		}
	}
	if v, ok := obj["tx-rx-128-255"]; ok {
		if v != "" {
			m.TxRx128255 = types.StringValue(v)
		} else {
			m.TxRx128255 = types.StringNull()
		}
	}
	if v, ok := obj["tx-rx-1519-max"]; ok {
		if v != "" {
			m.TxRx1519Max = types.StringValue(v)
		} else {
			m.TxRx1519Max = types.StringNull()
		}
	}
	if v, ok := obj["tx-rx-256-511"]; ok {
		if v != "" {
			m.TxRx256511 = types.StringValue(v)
		} else {
			m.TxRx256511 = types.StringNull()
		}
	}
	if v, ok := obj["tx-rx-512-1023"]; ok {
		if v != "" {
			m.TxRx5121023 = types.StringValue(v)
		} else {
			m.TxRx5121023 = types.StringNull()
		}
	}
	if v, ok := obj["tx-rx-64"]; ok {
		if v != "" {
			m.TxRx64 = types.StringValue(v)
		} else {
			m.TxRx64 = types.StringNull()
		}
	}
	if v, ok := obj["tx-rx-65-127"]; ok {
		if v != "" {
			m.TxRx65127 = types.StringValue(v)
		} else {
			m.TxRx65127 = types.StringNull()
		}
	}
	if v, ok := obj["tx-rx-bytes"]; ok {
		if v != "" {
			m.TxRxBytes = types.StringValue(v)
		} else {
			m.TxRxBytes = types.StringNull()
		}
	}
	if v, ok := obj["tx-rx-packets"]; ok {
		if v != "" {
			m.TxRxPackets = types.StringValue(v)
		} else {
			m.TxRxPackets = types.StringNull()
		}
	}
	if v, ok := obj["tx-single-collision"]; ok {
		if v != "" {
			m.TxSingleCollision = types.StringValue(v)
		} else {
			m.TxSingleCollision = types.StringNull()
		}
	}
	if v, ok := obj["tx-too-short"]; ok {
		if v != "" {
			m.TxTooShort = types.StringValue(v)
		} else {
			m.TxTooShort = types.StringNull()
		}
	}
	if v, ok := obj["tx-total-collision"]; ok {
		if v != "" {
			m.TxTotalCollision = types.StringValue(v)
		} else {
			m.TxTotalCollision = types.StringNull()
		}
	}
	if v, ok := obj["tx-underrun"]; ok {
		if v != "" {
			m.TxUnderrun = types.StringValue(v)
		} else {
			m.TxUnderrun = types.StringNull()
		}
	}
	if v, ok := obj["tx-unicast"]; ok {
		if v != "" {
			m.TxUnicast = types.StringValue(v)
		} else {
			m.TxUnicast = types.StringNull()
		}
	}
	if v, ok := obj["vendor-name"]; ok {
		if v != "" {
			m.VendorName = types.StringValue(v)
		} else {
			m.VendorName = types.StringNull()
		}
	}
	if v, ok := obj["vendor-part-number"]; ok {
		if v != "" {
			m.VendorPartNumber = types.StringValue(v)
		} else {
			m.VendorPartNumber = types.StringNull()
		}
	}
	if v, ok := obj["vendor-revision"]; ok {
		if v != "" {
			m.VendorRevision = types.StringValue(v)
		} else {
			m.VendorRevision = types.StringNull()
		}
	}
	if v, ok := obj["vendor-serial"]; ok {
		if v != "" {
			m.VendorSerial = types.StringValue(v)
		} else {
			m.VendorSerial = types.StringNull()
		}
	}
	if v, ok := obj["wavelength"]; ok {
		if v != "" {
			m.Wavelength = types.StringValue(v)
		} else {
			m.Wavelength = types.StringNull()
		}
	}
}
