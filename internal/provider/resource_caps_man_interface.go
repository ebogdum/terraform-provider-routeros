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
	_ resource.Resource                = &CapsManInterfaceResource{}
	_ resource.ResourceWithImportState = &CapsManInterfaceResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type CapsManInterfaceResource struct {
	reg *client.Registry
}

type CapsManInterfaceModel struct {
	ID                       types.String `tfsdk:"id"`
	VlanMode                 types.String `tfsdk:"vlan_mode"`
	VlanId                   types.String `tfsdk:"vlan_id"`
	VhtSupportedMcs          types.String `tfsdk:"vht_supported_mcs"`
	VhtBasicMcs              types.String `tfsdk:"vht_basic_mcs"`
	TxPower                  types.String `tfsdk:"tx_power"`
	TxChains                 types.String `tfsdk:"tx_chains"`
	TlsMode                  types.String `tfsdk:"tls_mode"`
	TlsCertificate           types.String `tfsdk:"tls_certificate"`
	Supported                types.String `tfsdk:"supported"`
	Ssid                     types.String `tfsdk:"ssid"`
	SkipDfsChannels          types.String `tfsdk:"skip_dfs_channels"`
	Security                 types.String `tfsdk:"security"`
	SecondaryFrequency       types.String `tfsdk:"secondary_frequency"`
	SaveSelected             types.String `tfsdk:"save_selected"`
	RxChains                 types.String `tfsdk:"rx_chains"`
	ReselectInterval         types.String `tfsdk:"reselect_interval"`
	Rates                    types.String `tfsdk:"rates"`
	RadioName                types.String `tfsdk:"radio_name"`
	Passphrase               types.String `tfsdk:"passphrase"`
	OpenflowSwitch           types.String `tfsdk:"openflow_switch"`
	MulticastHelper          types.String `tfsdk:"multicast_helper"`
	Mtu                      types.String `tfsdk:"mtu"`
	Mode                     types.String `tfsdk:"mode"`
	MaxStaCount              types.String `tfsdk:"max_sta_count"`
	LocalForwarding          types.String `tfsdk:"local_forwarding"`
	LoadBalancingGroup       types.String `tfsdk:"load_balancing_group"`
	L2mtu                    types.String `tfsdk:"l2mtu"`
	KeepaliveFrames          types.String `tfsdk:"keepalive_frames"`
	InterfaceList            types.String `tfsdk:"interface_list"`
	Installation             types.String `tfsdk:"installation"`
	HwRetries                types.String `tfsdk:"hw_retries"`
	HwProtectionMode         types.String `tfsdk:"hw_protection_mode"`
	HtSupportedMcs           types.String `tfsdk:"ht_supported_mcs"`
	HtBasicMcs               types.String `tfsdk:"ht_basic_mcs"`
	HideSsid                 types.String `tfsdk:"hide_ssid"`
	GuardInterval            types.String `tfsdk:"guard_interval"`
	GroupKeyUpdate           types.String `tfsdk:"group_key_update"`
	GroupEncryption          types.String `tfsdk:"group_encryption"`
	Frequency                types.String `tfsdk:"frequency"`
	FrameLifetime            types.String `tfsdk:"frame_lifetime"`
	ExtensionChannel         types.String `tfsdk:"extension_channel"`
	Encryption               types.String `tfsdk:"encryption"`
	EapRadiusAccounting      types.String `tfsdk:"eap_radius_accounting"`
	EapMethods               types.String `tfsdk:"eap_methods"`
	Distance                 types.String `tfsdk:"distance"`
	DisconnectTimeout        types.String `tfsdk:"disconnect_timeout"`
	DisableRunningCheck      types.String `tfsdk:"disable_running_check"`
	DisablePmkid             types.String `tfsdk:"disable_pmkid"`
	Datapath                 types.String `tfsdk:"datapath"`
	Country                  types.String `tfsdk:"country"`
	ControlChannelWidth      types.String `tfsdk:"control_channel_width"`
	Configuration            types.String `tfsdk:"configuration"`
	ClientToClientForwarding types.String `tfsdk:"client_to_client_forwarding"`
	Channel                  types.String `tfsdk:"channel"`
	BridgeHorizon            types.String `tfsdk:"bridge_horizon"`
	BridgeCost               types.String `tfsdk:"bridge_cost"`
	Bridge                   types.String `tfsdk:"bridge"`
	Basic                    types.String `tfsdk:"basic"`
	Band                     types.String `tfsdk:"band"`
	AuthenticationTypes      types.String `tfsdk:"authentication_types"`
	Arp                      types.String `tfsdk:"arp"`
	ARPTimeout               types.String `tfsdk:"arp_timeout"`
	Comment                  types.String `tfsdk:"comment"`
	Disabled                 types.Bool   `tfsdk:"disabled"`
	MACAddress               macValue     `tfsdk:"mac_address"`
	MasterInterface          types.String `tfsdk:"master_interface"`
	Name                     types.String `tfsdk:"name"`
	RadioMAC                 macValue     `tfsdk:"radio_mac"`
	Router                   types.String `tfsdk:"router"`
}

func NewCapsManInterfaceResource() resource.Resource { return &CapsManInterfaceResource{} }

func (r *CapsManInterfaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_caps_man_interface"
}

func (r *CapsManInterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *CapsManInterfaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "CAPsMAN virtual interfaces are typically created automatically; manual creation collides with the master.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"vlan_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `vlan-mode`.",
			},
			"vlan_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `vlan-id`.",
			},
			"vht_supported_mcs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `vht-supported-mcs`.",
			},
			"vht_basic_mcs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `vht-basic-mcs`.",
			},
			"tx_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tx-power`.",
			},
			"tx_chains": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tx-chains`.",
			},
			"tls_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tls-mode`.",
			},
			"tls_certificate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tls-certificate`.",
			},
			"supported": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `supported`.",
			},
			"ssid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ssid`.",
			},
			"skip_dfs_channels": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `skip-dfs-channels`.",
			},
			"security": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `security`.",
			},
			"secondary_frequency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `secondary-frequency`.",
			},
			"save_selected": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `save-selected`.",
			},
			"rx_chains": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `rx-chains`.",
			},
			"reselect_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `reselect-interval`.",
			},
			"rates": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `rates`.",
			},
			"radio_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `radio-name`.",
			},
			"passphrase": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "RouterOS `passphrase`.",
			},
			"openflow_switch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `openflow-switch`.",
			},
			"multicast_helper": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `multicast-helper`.",
			},
			"mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mtu`.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mode`.",
			},
			"max_sta_count": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `max-sta-count`.",
			},
			"local_forwarding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `local-forwarding`.",
			},
			"load_balancing_group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `load-balancing-group`.",
			},
			"l2mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2mtu`.",
			},
			"keepalive_frames": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `keepalive-frames`.",
			},
			"interface_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `interface-list`.",
			},
			"installation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `installation`.",
			},
			"hw_retries": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hw-retries`.",
			},
			"hw_protection_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hw-protection-mode`.",
			},
			"ht_supported_mcs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ht-supported-mcs`.",
			},
			"ht_basic_mcs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ht-basic-mcs`.",
			},
			"hide_ssid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hide-ssid`.",
			},
			"guard_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `guard-interval`.",
			},
			"group_key_update": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `group-key-update`.",
			},
			"group_encryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `group-encryption`.",
			},
			"frequency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `frequency`.",
			},
			"frame_lifetime": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `frame-lifetime`.",
			},
			"extension_channel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `extension-channel`.",
			},
			"encryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `encryption`.",
			},
			"eap_radius_accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `eap-radius-accounting`.",
			},
			"eap_methods": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `eap-methods`.",
			},
			"distance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `distance`.",
			},
			"disconnect_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `disconnect-timeout`.",
			},
			"disable_running_check": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `disable-running-check`.",
			},
			"disable_pmkid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `disable-pmkid`.",
			},
			"datapath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `datapath`.",
			},
			"country": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `country`.",
			},
			"control_channel_width": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `control-channel-width`.",
			},
			"configuration": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `configuration`.",
			},
			"client_to_client_forwarding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `client-to-client-forwarding`.",
			},
			"channel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `channel`.",
			},
			"bridge_horizon": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bridge-horizon`.",
			},
			"bridge_cost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bridge-cost`.",
			},
			"bridge": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bridge`.",
			},
			"basic": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `basic`.",
			},
			"band": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `band`.",
			},
			"authentication_types": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `authentication-types`.",
			},
			"arp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `arp`.",
			},
			"arp_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
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
			"mac_address": schema.StringAttribute{
				CustomType:  macType{},
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsMAC()},
			},
			"master_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"radio_mac": schema.StringAttribute{
				CustomType:  macType{},
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsMAC()},
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *CapsManInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CapsManInterfaceModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.ARPTimeout.IsNull() || plan.ARPTimeout.IsUnknown()) {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.MACAddress.IsNull() || plan.MACAddress.IsUnknown()) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !(plan.MasterInterface.IsNull() || plan.MasterInterface.IsUnknown()) {
		body["master-interface"] = plan.MasterInterface.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.RadioMAC.IsNull() || plan.RadioMAC.IsUnknown()) {
		body["radio-mac"] = plan.RadioMAC.ValueString()
	}
	if !(plan.Arp.IsNull() || plan.Arp.IsUnknown()) {
		body["arp"] = plan.Arp.ValueString()
	}
	if !(plan.AuthenticationTypes.IsNull() || plan.AuthenticationTypes.IsUnknown()) {
		body["authentication-types"] = plan.AuthenticationTypes.ValueString()
	}
	if !(plan.Band.IsNull() || plan.Band.IsUnknown()) {
		body["band"] = plan.Band.ValueString()
	}
	if !(plan.Basic.IsNull() || plan.Basic.IsUnknown()) {
		body["basic"] = plan.Basic.ValueString()
	}
	if !(plan.Bridge.IsNull() || plan.Bridge.IsUnknown()) {
		body["bridge"] = plan.Bridge.ValueString()
	}
	if !(plan.BridgeCost.IsNull() || plan.BridgeCost.IsUnknown()) {
		body["bridge-cost"] = plan.BridgeCost.ValueString()
	}
	if !(plan.BridgeHorizon.IsNull() || plan.BridgeHorizon.IsUnknown()) {
		body["bridge-horizon"] = plan.BridgeHorizon.ValueString()
	}
	if !(plan.Channel.IsNull() || plan.Channel.IsUnknown()) {
		body["channel"] = plan.Channel.ValueString()
	}
	if !(plan.ClientToClientForwarding.IsNull() || plan.ClientToClientForwarding.IsUnknown()) {
		body["client-to-client-forwarding"] = plan.ClientToClientForwarding.ValueString()
	}
	if !(plan.Configuration.IsNull() || plan.Configuration.IsUnknown()) {
		body["configuration"] = plan.Configuration.ValueString()
	}
	if !(plan.ControlChannelWidth.IsNull() || plan.ControlChannelWidth.IsUnknown()) {
		body["control-channel-width"] = plan.ControlChannelWidth.ValueString()
	}
	if !(plan.Country.IsNull() || plan.Country.IsUnknown()) {
		body["country"] = plan.Country.ValueString()
	}
	if !(plan.Datapath.IsNull() || plan.Datapath.IsUnknown()) {
		body["datapath"] = plan.Datapath.ValueString()
	}
	if !(plan.DisablePmkid.IsNull() || plan.DisablePmkid.IsUnknown()) {
		body["disable-pmkid"] = plan.DisablePmkid.ValueString()
	}
	if !(plan.DisableRunningCheck.IsNull() || plan.DisableRunningCheck.IsUnknown()) {
		body["disable-running-check"] = plan.DisableRunningCheck.ValueString()
	}
	if !(plan.DisconnectTimeout.IsNull() || plan.DisconnectTimeout.IsUnknown()) {
		body["disconnect-timeout"] = plan.DisconnectTimeout.ValueString()
	}
	if !(plan.Distance.IsNull() || plan.Distance.IsUnknown()) {
		body["distance"] = plan.Distance.ValueString()
	}
	if !(plan.EapMethods.IsNull() || plan.EapMethods.IsUnknown()) {
		body["eap-methods"] = plan.EapMethods.ValueString()
	}
	if !(plan.EapRadiusAccounting.IsNull() || plan.EapRadiusAccounting.IsUnknown()) {
		body["eap-radius-accounting"] = plan.EapRadiusAccounting.ValueString()
	}
	if !(plan.Encryption.IsNull() || plan.Encryption.IsUnknown()) {
		body["encryption"] = plan.Encryption.ValueString()
	}
	if !(plan.ExtensionChannel.IsNull() || plan.ExtensionChannel.IsUnknown()) {
		body["extension-channel"] = plan.ExtensionChannel.ValueString()
	}
	if !(plan.FrameLifetime.IsNull() || plan.FrameLifetime.IsUnknown()) {
		body["frame-lifetime"] = plan.FrameLifetime.ValueString()
	}
	if !(plan.Frequency.IsNull() || plan.Frequency.IsUnknown()) {
		body["frequency"] = plan.Frequency.ValueString()
	}
	if !(plan.GroupEncryption.IsNull() || plan.GroupEncryption.IsUnknown()) {
		body["group-encryption"] = plan.GroupEncryption.ValueString()
	}
	if !(plan.GroupKeyUpdate.IsNull() || plan.GroupKeyUpdate.IsUnknown()) {
		body["group-key-update"] = plan.GroupKeyUpdate.ValueString()
	}
	if !(plan.GuardInterval.IsNull() || plan.GuardInterval.IsUnknown()) {
		body["guard-interval"] = plan.GuardInterval.ValueString()
	}
	if !(plan.HideSsid.IsNull() || plan.HideSsid.IsUnknown()) {
		body["hide-ssid"] = plan.HideSsid.ValueString()
	}
	if !(plan.HtBasicMcs.IsNull() || plan.HtBasicMcs.IsUnknown()) {
		body["ht-basic-mcs"] = plan.HtBasicMcs.ValueString()
	}
	if !(plan.HtSupportedMcs.IsNull() || plan.HtSupportedMcs.IsUnknown()) {
		body["ht-supported-mcs"] = plan.HtSupportedMcs.ValueString()
	}
	if !(plan.HwProtectionMode.IsNull() || plan.HwProtectionMode.IsUnknown()) {
		body["hw-protection-mode"] = plan.HwProtectionMode.ValueString()
	}
	if !(plan.HwRetries.IsNull() || plan.HwRetries.IsUnknown()) {
		body["hw-retries"] = plan.HwRetries.ValueString()
	}
	if !(plan.Installation.IsNull() || plan.Installation.IsUnknown()) {
		body["installation"] = plan.Installation.ValueString()
	}
	if !(plan.InterfaceList.IsNull() || plan.InterfaceList.IsUnknown()) {
		body["interface-list"] = plan.InterfaceList.ValueString()
	}
	if !(plan.KeepaliveFrames.IsNull() || plan.KeepaliveFrames.IsUnknown()) {
		body["keepalive-frames"] = plan.KeepaliveFrames.ValueString()
	}
	if !(plan.L2mtu.IsNull() || plan.L2mtu.IsUnknown()) {
		body["l2mtu"] = plan.L2mtu.ValueString()
	}
	if !(plan.LoadBalancingGroup.IsNull() || plan.LoadBalancingGroup.IsUnknown()) {
		body["load-balancing-group"] = plan.LoadBalancingGroup.ValueString()
	}
	if !(plan.LocalForwarding.IsNull() || plan.LocalForwarding.IsUnknown()) {
		body["local-forwarding"] = plan.LocalForwarding.ValueString()
	}
	if !(plan.MaxStaCount.IsNull() || plan.MaxStaCount.IsUnknown()) {
		body["max-sta-count"] = plan.MaxStaCount.ValueString()
	}
	if !(plan.Mode.IsNull() || plan.Mode.IsUnknown()) {
		body["mode"] = plan.Mode.ValueString()
	}
	if !(plan.Mtu.IsNull() || plan.Mtu.IsUnknown()) {
		body["mtu"] = plan.Mtu.ValueString()
	}
	if !(plan.MulticastHelper.IsNull() || plan.MulticastHelper.IsUnknown()) {
		body["multicast-helper"] = plan.MulticastHelper.ValueString()
	}
	if !(plan.OpenflowSwitch.IsNull() || plan.OpenflowSwitch.IsUnknown()) {
		body["openflow-switch"] = plan.OpenflowSwitch.ValueString()
	}
	if !(plan.Passphrase.IsNull() || plan.Passphrase.IsUnknown()) {
		body["passphrase"] = plan.Passphrase.ValueString()
	}
	if !(plan.RadioName.IsNull() || plan.RadioName.IsUnknown()) {
		body["radio-name"] = plan.RadioName.ValueString()
	}
	if !(plan.Rates.IsNull() || plan.Rates.IsUnknown()) {
		body["rates"] = plan.Rates.ValueString()
	}
	if !(plan.ReselectInterval.IsNull() || plan.ReselectInterval.IsUnknown()) {
		body["reselect-interval"] = plan.ReselectInterval.ValueString()
	}
	if !(plan.RxChains.IsNull() || plan.RxChains.IsUnknown()) {
		body["rx-chains"] = plan.RxChains.ValueString()
	}
	if !(plan.SaveSelected.IsNull() || plan.SaveSelected.IsUnknown()) {
		body["save-selected"] = plan.SaveSelected.ValueString()
	}
	if !(plan.SecondaryFrequency.IsNull() || plan.SecondaryFrequency.IsUnknown()) {
		body["secondary-frequency"] = plan.SecondaryFrequency.ValueString()
	}
	if !(plan.Security.IsNull() || plan.Security.IsUnknown()) {
		body["security"] = plan.Security.ValueString()
	}
	if !(plan.SkipDfsChannels.IsNull() || plan.SkipDfsChannels.IsUnknown()) {
		body["skip-dfs-channels"] = plan.SkipDfsChannels.ValueString()
	}
	if !(plan.Ssid.IsNull() || plan.Ssid.IsUnknown()) {
		body["ssid"] = plan.Ssid.ValueString()
	}
	if !(plan.Supported.IsNull() || plan.Supported.IsUnknown()) {
		body["supported"] = plan.Supported.ValueString()
	}
	if !(plan.TlsCertificate.IsNull() || plan.TlsCertificate.IsUnknown()) {
		body["tls-certificate"] = plan.TlsCertificate.ValueString()
	}
	if !(plan.TlsMode.IsNull() || plan.TlsMode.IsUnknown()) {
		body["tls-mode"] = plan.TlsMode.ValueString()
	}
	if !(plan.TxChains.IsNull() || plan.TxChains.IsUnknown()) {
		body["tx-chains"] = plan.TxChains.ValueString()
	}
	if !(plan.TxPower.IsNull() || plan.TxPower.IsUnknown()) {
		body["tx-power"] = plan.TxPower.ValueString()
	}
	if !(plan.VhtBasicMcs.IsNull() || plan.VhtBasicMcs.IsUnknown()) {
		body["vht-basic-mcs"] = plan.VhtBasicMcs.ValueString()
	}
	if !(plan.VhtSupportedMcs.IsNull() || plan.VhtSupportedMcs.IsUnknown()) {
		body["vht-supported-mcs"] = plan.VhtSupportedMcs.ValueString()
	}
	if !(plan.VlanId.IsNull() || plan.VlanId.IsUnknown()) {
		body["vlan-id"] = plan.VlanId.ValueString()
	}
	if !(plan.VlanMode.IsNull() || plan.VlanMode.IsUnknown()) {
		body["vlan-mode"] = plan.VlanMode.ValueString()
	}
	obj, err := c.Add(ctx, "/caps-man/interface", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /caps-man/interface failed", err.Error())
		return
	}
	capsManInterfaceApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CapsManInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CapsManInterfaceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/caps-man/interface", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /caps-man/interface failed", err.Error())
		return
	}
	capsManInterfaceApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CapsManInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CapsManInterfaceModel
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
	if !plan.ARPTimeout.Equal(state.ARPTimeout) && !plan.ARPTimeout.IsUnknown() {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.MACAddress.Equal(state.MACAddress) && !plan.MACAddress.IsUnknown() {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.MasterInterface.Equal(state.MasterInterface) && !plan.MasterInterface.IsUnknown() {
		body["master-interface"] = plan.MasterInterface.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.RadioMAC.Equal(state.RadioMAC) && !plan.RadioMAC.IsUnknown() {
		body["radio-mac"] = plan.RadioMAC.ValueString()
	}
	if !plan.Arp.Equal(state.Arp) && !plan.Arp.IsUnknown() {
		body["arp"] = plan.Arp.ValueString()
	}
	if !plan.AuthenticationTypes.Equal(state.AuthenticationTypes) && !plan.AuthenticationTypes.IsUnknown() {
		body["authentication-types"] = plan.AuthenticationTypes.ValueString()
	}
	if !plan.Band.Equal(state.Band) && !plan.Band.IsUnknown() {
		body["band"] = plan.Band.ValueString()
	}
	if !plan.Basic.Equal(state.Basic) && !plan.Basic.IsUnknown() {
		body["basic"] = plan.Basic.ValueString()
	}
	if !plan.Bridge.Equal(state.Bridge) && !plan.Bridge.IsUnknown() {
		body["bridge"] = plan.Bridge.ValueString()
	}
	if !plan.BridgeCost.Equal(state.BridgeCost) && !plan.BridgeCost.IsUnknown() {
		body["bridge-cost"] = plan.BridgeCost.ValueString()
	}
	if !plan.BridgeHorizon.Equal(state.BridgeHorizon) && !plan.BridgeHorizon.IsUnknown() {
		body["bridge-horizon"] = plan.BridgeHorizon.ValueString()
	}
	if !plan.Channel.Equal(state.Channel) && !plan.Channel.IsUnknown() {
		body["channel"] = plan.Channel.ValueString()
	}
	if !plan.ClientToClientForwarding.Equal(state.ClientToClientForwarding) && !plan.ClientToClientForwarding.IsUnknown() {
		body["client-to-client-forwarding"] = plan.ClientToClientForwarding.ValueString()
	}
	if !plan.Configuration.Equal(state.Configuration) && !plan.Configuration.IsUnknown() {
		body["configuration"] = plan.Configuration.ValueString()
	}
	if !plan.ControlChannelWidth.Equal(state.ControlChannelWidth) && !plan.ControlChannelWidth.IsUnknown() {
		body["control-channel-width"] = plan.ControlChannelWidth.ValueString()
	}
	if !plan.Country.Equal(state.Country) && !plan.Country.IsUnknown() {
		body["country"] = plan.Country.ValueString()
	}
	if !plan.Datapath.Equal(state.Datapath) && !plan.Datapath.IsUnknown() {
		body["datapath"] = plan.Datapath.ValueString()
	}
	if !plan.DisablePmkid.Equal(state.DisablePmkid) && !plan.DisablePmkid.IsUnknown() {
		body["disable-pmkid"] = plan.DisablePmkid.ValueString()
	}
	if !plan.DisableRunningCheck.Equal(state.DisableRunningCheck) && !plan.DisableRunningCheck.IsUnknown() {
		body["disable-running-check"] = plan.DisableRunningCheck.ValueString()
	}
	if !plan.DisconnectTimeout.Equal(state.DisconnectTimeout) && !plan.DisconnectTimeout.IsUnknown() {
		body["disconnect-timeout"] = plan.DisconnectTimeout.ValueString()
	}
	if !plan.Distance.Equal(state.Distance) && !plan.Distance.IsUnknown() {
		body["distance"] = plan.Distance.ValueString()
	}
	if !plan.EapMethods.Equal(state.EapMethods) && !plan.EapMethods.IsUnknown() {
		body["eap-methods"] = plan.EapMethods.ValueString()
	}
	if !plan.EapRadiusAccounting.Equal(state.EapRadiusAccounting) && !plan.EapRadiusAccounting.IsUnknown() {
		body["eap-radius-accounting"] = plan.EapRadiusAccounting.ValueString()
	}
	if !plan.Encryption.Equal(state.Encryption) && !plan.Encryption.IsUnknown() {
		body["encryption"] = plan.Encryption.ValueString()
	}
	if !plan.ExtensionChannel.Equal(state.ExtensionChannel) && !plan.ExtensionChannel.IsUnknown() {
		body["extension-channel"] = plan.ExtensionChannel.ValueString()
	}
	if !plan.FrameLifetime.Equal(state.FrameLifetime) && !plan.FrameLifetime.IsUnknown() {
		body["frame-lifetime"] = plan.FrameLifetime.ValueString()
	}
	if !plan.Frequency.Equal(state.Frequency) && !plan.Frequency.IsUnknown() {
		body["frequency"] = plan.Frequency.ValueString()
	}
	if !plan.GroupEncryption.Equal(state.GroupEncryption) && !plan.GroupEncryption.IsUnknown() {
		body["group-encryption"] = plan.GroupEncryption.ValueString()
	}
	if !plan.GroupKeyUpdate.Equal(state.GroupKeyUpdate) && !plan.GroupKeyUpdate.IsUnknown() {
		body["group-key-update"] = plan.GroupKeyUpdate.ValueString()
	}
	if !plan.GuardInterval.Equal(state.GuardInterval) && !plan.GuardInterval.IsUnknown() {
		body["guard-interval"] = plan.GuardInterval.ValueString()
	}
	if !plan.HideSsid.Equal(state.HideSsid) && !plan.HideSsid.IsUnknown() {
		body["hide-ssid"] = plan.HideSsid.ValueString()
	}
	if !plan.HtBasicMcs.Equal(state.HtBasicMcs) && !plan.HtBasicMcs.IsUnknown() {
		body["ht-basic-mcs"] = plan.HtBasicMcs.ValueString()
	}
	if !plan.HtSupportedMcs.Equal(state.HtSupportedMcs) && !plan.HtSupportedMcs.IsUnknown() {
		body["ht-supported-mcs"] = plan.HtSupportedMcs.ValueString()
	}
	if !plan.HwProtectionMode.Equal(state.HwProtectionMode) && !plan.HwProtectionMode.IsUnknown() {
		body["hw-protection-mode"] = plan.HwProtectionMode.ValueString()
	}
	if !plan.HwRetries.Equal(state.HwRetries) && !plan.HwRetries.IsUnknown() {
		body["hw-retries"] = plan.HwRetries.ValueString()
	}
	if !plan.Installation.Equal(state.Installation) && !plan.Installation.IsUnknown() {
		body["installation"] = plan.Installation.ValueString()
	}
	if !plan.InterfaceList.Equal(state.InterfaceList) && !plan.InterfaceList.IsUnknown() {
		body["interface-list"] = plan.InterfaceList.ValueString()
	}
	if !plan.KeepaliveFrames.Equal(state.KeepaliveFrames) && !plan.KeepaliveFrames.IsUnknown() {
		body["keepalive-frames"] = plan.KeepaliveFrames.ValueString()
	}
	if !plan.L2mtu.Equal(state.L2mtu) && !plan.L2mtu.IsUnknown() {
		body["l2mtu"] = plan.L2mtu.ValueString()
	}
	if !plan.LoadBalancingGroup.Equal(state.LoadBalancingGroup) && !plan.LoadBalancingGroup.IsUnknown() {
		body["load-balancing-group"] = plan.LoadBalancingGroup.ValueString()
	}
	if !plan.LocalForwarding.Equal(state.LocalForwarding) && !plan.LocalForwarding.IsUnknown() {
		body["local-forwarding"] = plan.LocalForwarding.ValueString()
	}
	if !plan.MaxStaCount.Equal(state.MaxStaCount) && !plan.MaxStaCount.IsUnknown() {
		body["max-sta-count"] = plan.MaxStaCount.ValueString()
	}
	if !plan.Mode.Equal(state.Mode) && !plan.Mode.IsUnknown() {
		body["mode"] = plan.Mode.ValueString()
	}
	if !plan.Mtu.Equal(state.Mtu) && !plan.Mtu.IsUnknown() {
		body["mtu"] = plan.Mtu.ValueString()
	}
	if !plan.MulticastHelper.Equal(state.MulticastHelper) && !plan.MulticastHelper.IsUnknown() {
		body["multicast-helper"] = plan.MulticastHelper.ValueString()
	}
	if !plan.OpenflowSwitch.Equal(state.OpenflowSwitch) && !plan.OpenflowSwitch.IsUnknown() {
		body["openflow-switch"] = plan.OpenflowSwitch.ValueString()
	}
	if !plan.Passphrase.Equal(state.Passphrase) && !plan.Passphrase.IsUnknown() {
		body["passphrase"] = plan.Passphrase.ValueString()
	}
	if !plan.RadioName.Equal(state.RadioName) && !plan.RadioName.IsUnknown() {
		body["radio-name"] = plan.RadioName.ValueString()
	}
	if !plan.Rates.Equal(state.Rates) && !plan.Rates.IsUnknown() {
		body["rates"] = plan.Rates.ValueString()
	}
	if !plan.ReselectInterval.Equal(state.ReselectInterval) && !plan.ReselectInterval.IsUnknown() {
		body["reselect-interval"] = plan.ReselectInterval.ValueString()
	}
	if !plan.RxChains.Equal(state.RxChains) && !plan.RxChains.IsUnknown() {
		body["rx-chains"] = plan.RxChains.ValueString()
	}
	if !plan.SaveSelected.Equal(state.SaveSelected) && !plan.SaveSelected.IsUnknown() {
		body["save-selected"] = plan.SaveSelected.ValueString()
	}
	if !plan.SecondaryFrequency.Equal(state.SecondaryFrequency) && !plan.SecondaryFrequency.IsUnknown() {
		body["secondary-frequency"] = plan.SecondaryFrequency.ValueString()
	}
	if !plan.Security.Equal(state.Security) && !plan.Security.IsUnknown() {
		body["security"] = plan.Security.ValueString()
	}
	if !plan.SkipDfsChannels.Equal(state.SkipDfsChannels) && !plan.SkipDfsChannels.IsUnknown() {
		body["skip-dfs-channels"] = plan.SkipDfsChannels.ValueString()
	}
	if !plan.Ssid.Equal(state.Ssid) && !plan.Ssid.IsUnknown() {
		body["ssid"] = plan.Ssid.ValueString()
	}
	if !plan.Supported.Equal(state.Supported) && !plan.Supported.IsUnknown() {
		body["supported"] = plan.Supported.ValueString()
	}
	if !plan.TlsCertificate.Equal(state.TlsCertificate) && !plan.TlsCertificate.IsUnknown() {
		body["tls-certificate"] = plan.TlsCertificate.ValueString()
	}
	if !plan.TlsMode.Equal(state.TlsMode) && !plan.TlsMode.IsUnknown() {
		body["tls-mode"] = plan.TlsMode.ValueString()
	}
	if !plan.TxChains.Equal(state.TxChains) && !plan.TxChains.IsUnknown() {
		body["tx-chains"] = plan.TxChains.ValueString()
	}
	if !plan.TxPower.Equal(state.TxPower) && !plan.TxPower.IsUnknown() {
		body["tx-power"] = plan.TxPower.ValueString()
	}
	if !plan.VhtBasicMcs.Equal(state.VhtBasicMcs) && !plan.VhtBasicMcs.IsUnknown() {
		body["vht-basic-mcs"] = plan.VhtBasicMcs.ValueString()
	}
	if !plan.VhtSupportedMcs.Equal(state.VhtSupportedMcs) && !plan.VhtSupportedMcs.IsUnknown() {
		body["vht-supported-mcs"] = plan.VhtSupportedMcs.ValueString()
	}
	if !plan.VlanId.Equal(state.VlanId) && !plan.VlanId.IsUnknown() {
		body["vlan-id"] = plan.VlanId.ValueString()
	}
	if !plan.VlanMode.Equal(state.VlanMode) && !plan.VlanMode.IsUnknown() {
		body["vlan-mode"] = plan.VlanMode.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/caps-man/interface", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /caps-man/interface failed", err.Error())
			return
		}
		capsManInterfaceApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CapsManInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CapsManInterfaceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/caps-man/interface", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /caps-man/interface failed", err.Error())
	}
}

func (r *CapsManInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := capsManInterfaceLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /caps-man/interface matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// capsManInterfaceLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func capsManInterfaceLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/caps-man/interface", id)
}

func capsManInterfaceApply(ctx context.Context, obj client.Object, m *CapsManInterfaceModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["vlan-mode"]; ok && v != "" {
		m.VlanMode = types.StringValue(v)
	} else {
		m.VlanMode = types.StringNull()
	}
	if v, ok := obj["vlan-id"]; ok && v != "" {
		m.VlanId = types.StringValue(v)
	} else {
		m.VlanId = types.StringNull()
	}
	if v, ok := obj["vht-supported-mcs"]; ok && v != "" {
		m.VhtSupportedMcs = types.StringValue(v)
	} else {
		m.VhtSupportedMcs = types.StringNull()
	}
	if v, ok := obj["vht-basic-mcs"]; ok && v != "" {
		m.VhtBasicMcs = types.StringValue(v)
	} else {
		m.VhtBasicMcs = types.StringNull()
	}
	if v, ok := obj["tx-power"]; ok && v != "" {
		m.TxPower = types.StringValue(v)
	} else {
		m.TxPower = types.StringNull()
	}
	if v, ok := obj["tx-chains"]; ok && v != "" {
		m.TxChains = types.StringValue(v)
	} else {
		m.TxChains = types.StringNull()
	}
	if v, ok := obj["tls-mode"]; ok && v != "" {
		m.TlsMode = types.StringValue(v)
	} else {
		m.TlsMode = types.StringNull()
	}
	if v, ok := obj["tls-certificate"]; ok && v != "" {
		m.TlsCertificate = types.StringValue(v)
	} else {
		m.TlsCertificate = types.StringNull()
	}
	if v, ok := obj["supported"]; ok && v != "" {
		m.Supported = types.StringValue(v)
	} else {
		m.Supported = types.StringNull()
	}
	if v, ok := obj["ssid"]; ok && v != "" {
		m.Ssid = types.StringValue(v)
	} else {
		m.Ssid = types.StringNull()
	}
	if v, ok := obj["skip-dfs-channels"]; ok && v != "" {
		m.SkipDfsChannels = types.StringValue(v)
	} else {
		m.SkipDfsChannels = types.StringNull()
	}
	if v, ok := obj["security"]; ok && v != "" {
		m.Security = types.StringValue(v)
	} else {
		m.Security = types.StringNull()
	}
	if v, ok := obj["secondary-frequency"]; ok && v != "" {
		m.SecondaryFrequency = types.StringValue(v)
	} else {
		m.SecondaryFrequency = types.StringNull()
	}
	if v, ok := obj["save-selected"]; ok && v != "" {
		m.SaveSelected = types.StringValue(v)
	} else {
		m.SaveSelected = types.StringNull()
	}
	if v, ok := obj["rx-chains"]; ok && v != "" {
		m.RxChains = types.StringValue(v)
	} else {
		m.RxChains = types.StringNull()
	}
	if v, ok := obj["reselect-interval"]; ok && v != "" {
		m.ReselectInterval = types.StringValue(v)
	} else {
		m.ReselectInterval = types.StringNull()
	}
	if v, ok := obj["rates"]; ok && v != "" {
		m.Rates = types.StringValue(v)
	} else {
		m.Rates = types.StringNull()
	}
	if v, ok := obj["radio-name"]; ok && v != "" {
		m.RadioName = types.StringValue(v)
	} else {
		m.RadioName = types.StringNull()
	}
	if v, ok := obj["passphrase"]; ok && v != "" {
		m.Passphrase = types.StringValue(v)
	} else {
		m.Passphrase = types.StringNull()
	}
	if v, ok := obj["openflow-switch"]; ok && v != "" {
		m.OpenflowSwitch = types.StringValue(v)
	} else {
		m.OpenflowSwitch = types.StringNull()
	}
	if v, ok := obj["multicast-helper"]; ok && v != "" {
		m.MulticastHelper = types.StringValue(v)
	} else {
		m.MulticastHelper = types.StringNull()
	}
	if v, ok := obj["mtu"]; ok && v != "" {
		m.Mtu = types.StringValue(v)
	} else {
		m.Mtu = types.StringNull()
	}
	if v, ok := obj["mode"]; ok && v != "" {
		m.Mode = types.StringValue(v)
	} else {
		m.Mode = types.StringNull()
	}
	if v, ok := obj["max-sta-count"]; ok && v != "" {
		m.MaxStaCount = types.StringValue(v)
	} else {
		m.MaxStaCount = types.StringNull()
	}
	if v, ok := obj["local-forwarding"]; ok && v != "" {
		m.LocalForwarding = types.StringValue(v)
	} else {
		m.LocalForwarding = types.StringNull()
	}
	if v, ok := obj["load-balancing-group"]; ok && v != "" {
		m.LoadBalancingGroup = types.StringValue(v)
	} else {
		m.LoadBalancingGroup = types.StringNull()
	}
	if v, ok := obj["l2mtu"]; ok && v != "" {
		m.L2mtu = types.StringValue(v)
	} else {
		m.L2mtu = types.StringNull()
	}
	if v, ok := obj["keepalive-frames"]; ok && v != "" {
		m.KeepaliveFrames = types.StringValue(v)
	} else {
		m.KeepaliveFrames = types.StringNull()
	}
	if v, ok := obj["interface-list"]; ok && v != "" {
		m.InterfaceList = types.StringValue(v)
	} else {
		m.InterfaceList = types.StringNull()
	}
	if v, ok := obj["installation"]; ok && v != "" {
		m.Installation = types.StringValue(v)
	} else {
		m.Installation = types.StringNull()
	}
	if v, ok := obj["hw-retries"]; ok && v != "" {
		m.HwRetries = types.StringValue(v)
	} else {
		m.HwRetries = types.StringNull()
	}
	if v, ok := obj["hw-protection-mode"]; ok && v != "" {
		m.HwProtectionMode = types.StringValue(v)
	} else {
		m.HwProtectionMode = types.StringNull()
	}
	if v, ok := obj["ht-supported-mcs"]; ok && v != "" {
		m.HtSupportedMcs = types.StringValue(v)
	} else {
		m.HtSupportedMcs = types.StringNull()
	}
	if v, ok := obj["ht-basic-mcs"]; ok && v != "" {
		m.HtBasicMcs = types.StringValue(v)
	} else {
		m.HtBasicMcs = types.StringNull()
	}
	if v, ok := obj["hide-ssid"]; ok && v != "" {
		m.HideSsid = types.StringValue(v)
	} else {
		m.HideSsid = types.StringNull()
	}
	if v, ok := obj["guard-interval"]; ok && v != "" {
		m.GuardInterval = types.StringValue(v)
	} else {
		m.GuardInterval = types.StringNull()
	}
	if v, ok := obj["group-key-update"]; ok && v != "" {
		m.GroupKeyUpdate = types.StringValue(v)
	} else {
		m.GroupKeyUpdate = types.StringNull()
	}
	if v, ok := obj["group-encryption"]; ok && v != "" {
		m.GroupEncryption = types.StringValue(v)
	} else {
		m.GroupEncryption = types.StringNull()
	}
	if v, ok := obj["frequency"]; ok && v != "" {
		m.Frequency = types.StringValue(v)
	} else {
		m.Frequency = types.StringNull()
	}
	if v, ok := obj["frame-lifetime"]; ok && v != "" {
		m.FrameLifetime = types.StringValue(v)
	} else {
		m.FrameLifetime = types.StringNull()
	}
	if v, ok := obj["extension-channel"]; ok && v != "" {
		m.ExtensionChannel = types.StringValue(v)
	} else {
		m.ExtensionChannel = types.StringNull()
	}
	if v, ok := obj["encryption"]; ok && v != "" {
		m.Encryption = types.StringValue(v)
	} else {
		m.Encryption = types.StringNull()
	}
	if v, ok := obj["eap-radius-accounting"]; ok && v != "" {
		m.EapRadiusAccounting = types.StringValue(v)
	} else {
		m.EapRadiusAccounting = types.StringNull()
	}
	if v, ok := obj["eap-methods"]; ok && v != "" {
		m.EapMethods = types.StringValue(v)
	} else {
		m.EapMethods = types.StringNull()
	}
	if v, ok := obj["distance"]; ok && v != "" {
		m.Distance = types.StringValue(v)
	} else {
		m.Distance = types.StringNull()
	}
	if v, ok := obj["disconnect-timeout"]; ok && v != "" {
		m.DisconnectTimeout = types.StringValue(v)
	} else {
		m.DisconnectTimeout = types.StringNull()
	}
	if v, ok := obj["disable-running-check"]; ok && v != "" {
		m.DisableRunningCheck = types.StringValue(v)
	} else {
		m.DisableRunningCheck = types.StringNull()
	}
	if v, ok := obj["disable-pmkid"]; ok && v != "" {
		m.DisablePmkid = types.StringValue(v)
	} else {
		m.DisablePmkid = types.StringNull()
	}
	if v, ok := obj["datapath"]; ok && v != "" {
		m.Datapath = types.StringValue(v)
	} else {
		m.Datapath = types.StringNull()
	}
	if v, ok := obj["country"]; ok && v != "" {
		m.Country = types.StringValue(v)
	} else {
		m.Country = types.StringNull()
	}
	if v, ok := obj["control-channel-width"]; ok && v != "" {
		m.ControlChannelWidth = types.StringValue(v)
	} else {
		m.ControlChannelWidth = types.StringNull()
	}
	if v, ok := obj["configuration"]; ok && v != "" {
		m.Configuration = types.StringValue(v)
	} else {
		m.Configuration = types.StringNull()
	}
	if v, ok := obj["client-to-client-forwarding"]; ok && v != "" {
		m.ClientToClientForwarding = types.StringValue(v)
	} else {
		m.ClientToClientForwarding = types.StringNull()
	}
	if v, ok := obj["channel"]; ok && v != "" {
		m.Channel = types.StringValue(v)
	} else {
		m.Channel = types.StringNull()
	}
	if v, ok := obj["bridge-horizon"]; ok && v != "" {
		m.BridgeHorizon = types.StringValue(v)
	} else {
		m.BridgeHorizon = types.StringNull()
	}
	if v, ok := obj["bridge-cost"]; ok && v != "" {
		m.BridgeCost = types.StringValue(v)
	} else {
		m.BridgeCost = types.StringNull()
	}
	if v, ok := obj["bridge"]; ok && v != "" {
		m.Bridge = types.StringValue(v)
	} else {
		m.Bridge = types.StringNull()
	}
	if v, ok := obj["basic"]; ok && v != "" {
		m.Basic = types.StringValue(v)
	} else {
		m.Basic = types.StringNull()
	}
	if v, ok := obj["band"]; ok && v != "" {
		m.Band = types.StringValue(v)
	} else {
		m.Band = types.StringNull()
	}
	if v, ok := obj["authentication-types"]; ok && v != "" {
		m.AuthenticationTypes = types.StringValue(v)
	} else {
		m.AuthenticationTypes = types.StringNull()
	}
	if v, ok := obj["arp"]; ok && v != "" {
		m.Arp = types.StringValue(v)
	} else {
		m.Arp = types.StringNull()
	}
	if v, ok := obj["arp-timeout"]; ok {
		if v != "" {
			m.ARPTimeout = types.StringValue(v)
		} else {
			m.ARPTimeout = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Disabled = types.BoolValue(true)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["mac-address"]; ok {
		if v != "" {
			m.MACAddress = newMACValue(v)
		} else {
			m.MACAddress = newMACNull()
		}
	}
	if v, ok := obj["master-interface"]; ok {
		if v != "" {
			m.MasterInterface = types.StringValue(v)
		} else {
			m.MasterInterface = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["radio-mac"]; ok {
		if v != "" {
			m.RadioMAC = newMACValue(v)
		} else {
			m.RadioMAC = newMACNull()
		}
	}
}
