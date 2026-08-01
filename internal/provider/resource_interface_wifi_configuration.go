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
	_ resource.Resource                = &InterfaceWifiConfigurationResource{}
	_ resource.ResourceWithImportState = &InterfaceWifiConfigurationResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceWifiConfigurationResource struct {
	reg *client.Registry
}

type InterfaceWifiConfigurationModel struct {
	ID                        types.String `tfsdk:"id"`
	QosClassifier             types.String `tfsdk:"qos_classifier"`
	X2gProbeDelay             types.String `tfsdk:"x2g_probe_delay"`
	X3gppInfo                 types.String `tfsdk:"x3gpp_info"`
	X3gppInfoRaw              types.String `tfsdk:"x3gpp_info_raw"`
	Aaa                       types.String `tfsdk:"aaa"`
	AntennaGain               types.String `tfsdk:"antenna_gain"`
	AuthenticationTypes       types.String `tfsdk:"authentication_types"`
	Band                      types.String `tfsdk:"band"`
	BeaconInterval            types.String `tfsdk:"beacon_interval"`
	BeaconProtection          types.String `tfsdk:"beacon_protection"`
	Bridge                    types.String `tfsdk:"bridge"`
	BridgeCost                types.String `tfsdk:"bridge_cost"`
	BridgeHorizon             types.String `tfsdk:"bridge_horizon"`
	CalledFormat              types.String `tfsdk:"called_format"`
	CallingFormat             types.String `tfsdk:"calling_format"`
	Chains                    types.String `tfsdk:"chains"`
	Channel                   types.String `tfsdk:"channel"`
	ChannelWidth              types.String `tfsdk:"channel_width"`
	Ciphers                   types.String `tfsdk:"ciphers"`
	ClientIsolation           types.String `tfsdk:"client_isolation"`
	Comment                   types.String `tfsdk:"comment"`
	ConnectGroup              types.String `tfsdk:"connect_group"`
	ConnectPriority           types.String `tfsdk:"connect_priority"`
	ConnectionCapabilities    types.String `tfsdk:"connection_capabilities"`
	Country                   types.String `tfsdk:"country"`
	Datapath                  types.String `tfsdk:"datapath"`
	DeprioritizeUnii34        types.String `tfsdk:"deprioritize_unii_3_4"`
	Dgaf                      types.String `tfsdk:"dgaf"`
	DhGroups                  types.String `tfsdk:"dh_groups"`
	DisablePmkid              types.String `tfsdk:"disable_pmkid"`
	Disabled                  types.Bool   `tfsdk:"disabled"`
	Distance                  types.String `tfsdk:"distance"`
	DomainNames               types.String `tfsdk:"domain_names"`
	DtimPeriod                types.String `tfsdk:"dtim_period"`
	EAPAccounting             types.String `tfsdk:"eap_accounting"`
	EAPAnonymousIdentity      types.String `tfsdk:"eap_anonymous_identity"`
	EAPCertificateMode        types.String `tfsdk:"eap_certificate_mode"`
	EAPMethods                types.String `tfsdk:"eap_methods"`
	EAPPassword               types.String `tfsdk:"eap_password"`
	EAPTLSCertificate         types.String `tfsdk:"eap_tls_certificate"`
	EAPUsername               types.String `tfsdk:"eap_username"`
	Esr                       types.String `tfsdk:"esr"`
	Frequency                 types.String `tfsdk:"frequency"`
	FtEnabled                 types.String `tfsdk:"ft_enabled"`
	FtMobilityDomain          types.String `tfsdk:"ft_mobility_domain"`
	FtNasIdentifier           types.String `tfsdk:"ft_nas_identifier"`
	FtOverDs                  types.String `tfsdk:"ft_over_ds"`
	FtPreserveVLANID          types.String `tfsdk:"ft_preserve_vlan_id"`
	FtR0KeyLifetime           types.String `tfsdk:"ft_r0_key_lifetime"`
	FtReassocDeadline         types.String `tfsdk:"ft_reassoc_deadline"`
	GroupEncryption           types.String `tfsdk:"group_encryption"`
	GroupKeyUpdate            types.String `tfsdk:"group_key_update"`
	Hessid                    types.String `tfsdk:"hessid"`
	HideSsid                  types.String `tfsdk:"hide_ssid"`
	Hotspot20                 types.String `tfsdk:"hotspot_2_0"`
	HwProtectionMode          types.String `tfsdk:"hw_protection_mode"`
	Installation              types.String `tfsdk:"installation"`
	InterfaceList             types.String `tfsdk:"interface_list"`
	InterimUpdate             types.String `tfsdk:"interim_update"`
	Internet                  types.String `tfsdk:"internet"`
	Interworking              types.String `tfsdk:"interworking"`
	Ipv4Availability          types.String `tfsdk:"ipv4_availability"`
	IPV6Availability          types.String `tfsdk:"ipv6_availability"`
	MACCaching                types.String `tfsdk:"mac_caching"`
	ManagementEncryption      types.String `tfsdk:"management_encryption"`
	ManagementProtection      types.String `tfsdk:"management_protection"`
	Manager                   types.String `tfsdk:"manager"`
	MaxClients                types.String `tfsdk:"max_clients"`
	MaxTxPower                types.String `tfsdk:"max_tx_power"`
	Mode                      types.String `tfsdk:"mode"`
	MultiPassphraseGroup      types.String `tfsdk:"multi_passphrase_group"`
	MulticastEnhance          types.String `tfsdk:"multicast_enhance"`
	Name                      types.String `tfsdk:"name"`
	NasIdentifier             types.String `tfsdk:"nas_identifier"`
	NeighborGroup             types.String `tfsdk:"neighbor_group"`
	NetworkType               types.String `tfsdk:"network_type"`
	OpenFlowSwitch            types.String `tfsdk:"open_flow_switch"`
	Openflow                  types.String `tfsdk:"openflow"`
	OperationalClasses        types.String `tfsdk:"operational_classes"`
	OperatorNames             types.String `tfsdk:"operator_names"`
	OweTransitionInterface    types.String `tfsdk:"owe_transition_interface"`
	Passphrase                types.String `tfsdk:"passphrase"`
	PasswordFormat            types.String `tfsdk:"password_format"`
	QoSClassifier             types.String `tfsdk:"qo_s_classifier"`
	Realms                    types.String `tfsdk:"realms"`
	RealmsRaw                 types.String `tfsdk:"realms_raw"`
	ReselectInterval          types.String `tfsdk:"reselect_interval"`
	ReselectTime              types.String `tfsdk:"reselect_time"`
	RoamingOis                types.String `tfsdk:"roaming_ois"`
	Rrm                       types.String `tfsdk:"rrm"`
	SaeAntiCloggingThreshold  types.String `tfsdk:"sae_anti_clogging_threshold"`
	SaeMaxFailureRate         types.String `tfsdk:"sae_max_failure_rate"`
	SaePwe                    types.String `tfsdk:"sae_pwe"`
	SecondaryFrequency        types.String `tfsdk:"secondary_frequency"`
	Security                  types.String `tfsdk:"security"`
	SkipDfsChannels           types.String `tfsdk:"skip_dfs_channels"`
	Ssid                      types.String `tfsdk:"ssid"`
	StationRoaming            types.String `tfsdk:"station_roaming"`
	Steering                  types.String `tfsdk:"steering"`
	TrafficProcessing         types.String `tfsdk:"traffic_processing"`
	TransitionRequestCount    types.String `tfsdk:"transition_request_count"`
	TransitionThreshold       types.String `tfsdk:"transition_threshold"`
	TransitionThresholdPeriod types.String `tfsdk:"transition_threshold_period"`
	TransitionThresholdTime   types.String `tfsdk:"transition_threshold_time"`
	TransitionTime            types.String `tfsdk:"transition_time"`
	TxChains                  types.String `tfsdk:"tx_chains"`
	TxPower                   types.String `tfsdk:"tx_power"`
	Types                     types.String `tfsdk:"types"`
	Uesa                      types.String `tfsdk:"uesa"`
	UsernameFormat            types.String `tfsdk:"username_format"`
	Venue                     types.String `tfsdk:"venue"`
	VenueNames                types.String `tfsdk:"venue_names"`
	VLANID                    types.String `tfsdk:"vlan_id"`
	WanAtCapacity             types.String `tfsdk:"wan_at_capacity"`
	WanDownlink               types.String `tfsdk:"wan_downlink"`
	WanDownlinkLoad           types.String `tfsdk:"wan_downlink_load"`
	WanMeasurementDuration    types.String `tfsdk:"wan_measurement_duration"`
	WanStatus                 types.String `tfsdk:"wan_status"`
	WanSymmetric              types.String `tfsdk:"wan_symmetric"`
	WanUplink                 types.String `tfsdk:"wan_uplink"`
	WanUplinkLoad             types.String `tfsdk:"wan_uplink_load"`
	Wnm                       types.String `tfsdk:"wnm"`
	Wps                       types.String `tfsdk:"wps"`
	Router                    types.String `tfsdk:"router"`
}

func NewInterfaceWifiConfigurationResource() resource.Resource {
	return &InterfaceWifiConfigurationResource{}
}

func (r *InterfaceWifiConfigurationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_wifi_configuration"
}

func (r *InterfaceWifiConfigurationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceWifiConfigurationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/wifi/configuration`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"qos_classifier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `qos-classifier`.",
			},
			"x2g_probe_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"x3gpp_info": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"x3gpp_info_raw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"aaa": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"antenna_gain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"authentication_types": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"band": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"beacon_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"beacon_protection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge_cost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge_horizon": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"called_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"calling_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"chains": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"channel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"channel_width": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ciphers": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"client_isolation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"connect_group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"connect_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"connection_capabilities": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"country": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"datapath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"deprioritize_unii_3_4": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dgaf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dh_groups": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disable_pmkid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"distance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"domain_names": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dtim_period": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"eap_accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"eap_anonymous_identity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"eap_certificate_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"eap_methods": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"eap_password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "",
			},
			"eap_tls_certificate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"eap_username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"esr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"frequency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ft_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ft_mobility_domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ft_nas_identifier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ft_over_ds": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ft_preserve_vlan_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ft_r0_key_lifetime": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ft_reassoc_deadline": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"group_encryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"group_key_update": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hessid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hide_ssid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hotspot_2_0": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hw_protection_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"installation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interface_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interim_update": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"internet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interworking": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ipv4_availability": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ipv6_availability": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mac_caching": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"management_encryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"management_protection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"manager": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"max_clients": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"max_tx_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"multi_passphrase_group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"multicast_enhance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"nas_identifier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"neighbor_group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"network_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"open_flow_switch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"openflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"operational_classes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"operator_names": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"owe_transition_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"passphrase": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"password_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"qo_s_classifier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"realms": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"realms_raw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"reselect_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"reselect_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"roaming_ois": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rrm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sae_anti_clogging_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sae_max_failure_rate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sae_pwe": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"secondary_frequency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"security": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"skip_dfs_channels": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ssid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"station_roaming": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"steering": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"traffic_processing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"transition_request_count": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"transition_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"transition_threshold_period": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"transition_threshold_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"transition_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_chains": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"types": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"uesa": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"username_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"venue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"venue_names": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vlan_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_at_capacity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_downlink": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_downlink_load": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_measurement_duration": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_symmetric": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_uplink": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_uplink_load": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wnm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wps": schema.StringAttribute{
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

func (r *InterfaceWifiConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceWifiConfigurationModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.X2gProbeDelay.IsNull() || plan.X2gProbeDelay.IsUnknown()) {
		body["2g-probe-delay"] = plan.X2gProbeDelay.ValueString()
	}
	if !(plan.X3gppInfo.IsNull() || plan.X3gppInfo.IsUnknown()) {
		body["3gpp-info"] = plan.X3gppInfo.ValueString()
	}
	if !(plan.X3gppInfoRaw.IsNull() || plan.X3gppInfoRaw.IsUnknown()) {
		body["3gpp-info-raw"] = plan.X3gppInfoRaw.ValueString()
	}
	if !(plan.Aaa.IsNull() || plan.Aaa.IsUnknown()) {
		body["aaa"] = plan.Aaa.ValueString()
	}
	if !(plan.AntennaGain.IsNull() || plan.AntennaGain.IsUnknown()) {
		body["antenna-gain"] = plan.AntennaGain.ValueString()
	}
	if !(plan.AuthenticationTypes.IsNull() || plan.AuthenticationTypes.IsUnknown()) {
		body["security.authentication-types"] = plan.AuthenticationTypes.ValueString()
	}
	if !(plan.Band.IsNull() || plan.Band.IsUnknown()) {
		body["channel.band"] = plan.Band.ValueString()
	}
	if !(plan.BeaconInterval.IsNull() || plan.BeaconInterval.IsUnknown()) {
		body["beacon-interval"] = plan.BeaconInterval.ValueString()
	}
	if !(plan.BeaconProtection.IsNull() || plan.BeaconProtection.IsUnknown()) {
		body["beacon-protection"] = plan.BeaconProtection.ValueString()
	}
	if !(plan.Bridge.IsNull() || plan.Bridge.IsUnknown()) {
		body["datapath.bridge"] = plan.Bridge.ValueString()
	}
	if !(plan.BridgeCost.IsNull() || plan.BridgeCost.IsUnknown()) {
		body["datapath.bridge-cost"] = plan.BridgeCost.ValueString()
	}
	if !(plan.BridgeHorizon.IsNull() || plan.BridgeHorizon.IsUnknown()) {
		body["datapath.bridge-horizon"] = plan.BridgeHorizon.ValueString()
	}
	if !(plan.CalledFormat.IsNull() || plan.CalledFormat.IsUnknown()) {
		body["aaa.called-format"] = plan.CalledFormat.ValueString()
	}
	if !(plan.CallingFormat.IsNull() || plan.CallingFormat.IsUnknown()) {
		body["aaa.calling-format"] = plan.CallingFormat.ValueString()
	}
	if !(plan.Chains.IsNull() || plan.Chains.IsUnknown()) {
		body["chains"] = plan.Chains.ValueString()
	}
	if !(plan.Channel.IsNull() || plan.Channel.IsUnknown()) {
		body["channel"] = plan.Channel.ValueString()
	}
	if !(plan.ChannelWidth.IsNull() || plan.ChannelWidth.IsUnknown()) {
		body["channel-width"] = plan.ChannelWidth.ValueString()
	}
	if !(plan.Ciphers.IsNull() || plan.Ciphers.IsUnknown()) {
		body["ciphers"] = plan.Ciphers.ValueString()
	}
	if !(plan.ClientIsolation.IsNull() || plan.ClientIsolation.IsUnknown()) {
		body["datapath.client-isolation"] = plan.ClientIsolation.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.ConnectGroup.IsNull() || plan.ConnectGroup.IsUnknown()) {
		body["connect-group"] = plan.ConnectGroup.ValueString()
	}
	if !(plan.ConnectPriority.IsNull() || plan.ConnectPriority.IsUnknown()) {
		body["security.connect-priority"] = plan.ConnectPriority.ValueString()
	}
	if !(plan.ConnectionCapabilities.IsNull() || plan.ConnectionCapabilities.IsUnknown()) {
		body["connection-capabilities"] = plan.ConnectionCapabilities.ValueString()
	}
	if !(plan.Country.IsNull() || plan.Country.IsUnknown()) {
		body["country"] = plan.Country.ValueString()
	}
	if !(plan.Datapath.IsNull() || plan.Datapath.IsUnknown()) {
		body["datapath"] = plan.Datapath.ValueString()
	}
	if !(plan.DeprioritizeUnii34.IsNull() || plan.DeprioritizeUnii34.IsUnknown()) {
		body["deprioritize-unii-3-4"] = plan.DeprioritizeUnii34.ValueString()
	}
	if !(plan.Dgaf.IsNull() || plan.Dgaf.IsUnknown()) {
		body["dgaf"] = plan.Dgaf.ValueString()
	}
	if !(plan.DhGroups.IsNull() || plan.DhGroups.IsUnknown()) {
		body["security.dh-groups"] = plan.DhGroups.ValueString()
	}
	if !(plan.DisablePmkid.IsNull() || plan.DisablePmkid.IsUnknown()) {
		body["security.disable-pmkid"] = plan.DisablePmkid.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Distance.IsNull() || plan.Distance.IsUnknown()) {
		body["distance"] = plan.Distance.ValueString()
	}
	if !(plan.DomainNames.IsNull() || plan.DomainNames.IsUnknown()) {
		body["domain-names"] = plan.DomainNames.ValueString()
	}
	if !(plan.DtimPeriod.IsNull() || plan.DtimPeriod.IsUnknown()) {
		body["dtim-period"] = plan.DtimPeriod.ValueString()
	}
	if !(plan.EAPAccounting.IsNull() || plan.EAPAccounting.IsUnknown()) {
		body["security.eap-accounting"] = plan.EAPAccounting.ValueString()
	}
	if !(plan.EAPAnonymousIdentity.IsNull() || plan.EAPAnonymousIdentity.IsUnknown()) {
		body["security.eap-anonymous-identity"] = plan.EAPAnonymousIdentity.ValueString()
	}
	if !(plan.EAPCertificateMode.IsNull() || plan.EAPCertificateMode.IsUnknown()) {
		body["security.eap-certificate-mode"] = plan.EAPCertificateMode.ValueString()
	}
	if !(plan.EAPMethods.IsNull() || plan.EAPMethods.IsUnknown()) {
		body["security.eap-methods"] = plan.EAPMethods.ValueString()
	}
	if !(plan.EAPPassword.IsNull() || plan.EAPPassword.IsUnknown()) {
		body["security.eap-password"] = plan.EAPPassword.ValueString()
	}
	if !(plan.EAPTLSCertificate.IsNull() || plan.EAPTLSCertificate.IsUnknown()) {
		body["eap-tls-certificate"] = plan.EAPTLSCertificate.ValueString()
	}
	if !(plan.EAPUsername.IsNull() || plan.EAPUsername.IsUnknown()) {
		body["security.eap-username"] = plan.EAPUsername.ValueString()
	}
	if !(plan.Esr.IsNull() || plan.Esr.IsUnknown()) {
		body["esr"] = plan.Esr.ValueString()
	}
	if !(plan.Frequency.IsNull() || plan.Frequency.IsUnknown()) {
		body["channel.frequency"] = plan.Frequency.ValueString()
	}
	if !(plan.FtEnabled.IsNull() || plan.FtEnabled.IsUnknown()) {
		body["ft-enabled"] = plan.FtEnabled.ValueString()
	}
	if !(plan.FtMobilityDomain.IsNull() || plan.FtMobilityDomain.IsUnknown()) {
		body["ft-mobility-domain"] = plan.FtMobilityDomain.ValueString()
	}
	if !(plan.FtNasIdentifier.IsNull() || plan.FtNasIdentifier.IsUnknown()) {
		body["ft-nas-identifier"] = plan.FtNasIdentifier.ValueString()
	}
	if !(plan.FtOverDs.IsNull() || plan.FtOverDs.IsUnknown()) {
		body["security.ft-over-ds"] = plan.FtOverDs.ValueString()
	}
	if !(plan.FtPreserveVLANID.IsNull() || plan.FtPreserveVLANID.IsUnknown()) {
		body["ft-preserve-vlan-id"] = plan.FtPreserveVLANID.ValueString()
	}
	if !(plan.FtR0KeyLifetime.IsNull() || plan.FtR0KeyLifetime.IsUnknown()) {
		body["ft-r0-key-lifetime"] = plan.FtR0KeyLifetime.ValueString()
	}
	if !(plan.FtReassocDeadline.IsNull() || plan.FtReassocDeadline.IsUnknown()) {
		body["ft-reassoc-deadline"] = plan.FtReassocDeadline.ValueString()
	}
	if !(plan.GroupEncryption.IsNull() || plan.GroupEncryption.IsUnknown()) {
		body["security.group-encryption"] = plan.GroupEncryption.ValueString()
	}
	if !(plan.GroupKeyUpdate.IsNull() || plan.GroupKeyUpdate.IsUnknown()) {
		body["security.group-key-update"] = plan.GroupKeyUpdate.ValueString()
	}
	if !(plan.Hessid.IsNull() || plan.Hessid.IsUnknown()) {
		body["hessid"] = plan.Hessid.ValueString()
	}
	if !(plan.HideSsid.IsNull() || plan.HideSsid.IsUnknown()) {
		body["hide-ssid"] = plan.HideSsid.ValueString()
	}
	if !(plan.Hotspot20.IsNull() || plan.Hotspot20.IsUnknown()) {
		body["hotspot-2-0"] = plan.Hotspot20.ValueString()
	}
	if !(plan.HwProtectionMode.IsNull() || plan.HwProtectionMode.IsUnknown()) {
		body["hw-protection-mode"] = plan.HwProtectionMode.ValueString()
	}
	if !(plan.Installation.IsNull() || plan.Installation.IsUnknown()) {
		body["installation"] = plan.Installation.ValueString()
	}
	if !(plan.InterfaceList.IsNull() || plan.InterfaceList.IsUnknown()) {
		body["datapath.interface-list"] = plan.InterfaceList.ValueString()
	}
	if !(plan.InterimUpdate.IsNull() || plan.InterimUpdate.IsUnknown()) {
		body["aaa.interim-update"] = plan.InterimUpdate.ValueString()
	}
	if !(plan.Internet.IsNull() || plan.Internet.IsUnknown()) {
		body["internet"] = plan.Internet.ValueString()
	}
	if !(plan.Interworking.IsNull() || plan.Interworking.IsUnknown()) {
		body["interworking"] = plan.Interworking.ValueString()
	}
	if !(plan.Ipv4Availability.IsNull() || plan.Ipv4Availability.IsUnknown()) {
		body["ipv4-availability"] = plan.Ipv4Availability.ValueString()
	}
	if !(plan.IPV6Availability.IsNull() || plan.IPV6Availability.IsUnknown()) {
		body["ipv6-availability"] = plan.IPV6Availability.ValueString()
	}
	if !(plan.MACCaching.IsNull() || plan.MACCaching.IsUnknown()) {
		body["aaa.mac-caching"] = plan.MACCaching.ValueString()
	}
	if !(plan.ManagementEncryption.IsNull() || plan.ManagementEncryption.IsUnknown()) {
		body["management-encryption"] = plan.ManagementEncryption.ValueString()
	}
	if !(plan.ManagementProtection.IsNull() || plan.ManagementProtection.IsUnknown()) {
		body["security.management-protection"] = plan.ManagementProtection.ValueString()
	}
	if !(plan.Manager.IsNull() || plan.Manager.IsUnknown()) {
		body["manager"] = plan.Manager.ValueString()
	}
	if !(plan.MaxClients.IsNull() || plan.MaxClients.IsUnknown()) {
		body["max-clients"] = plan.MaxClients.ValueString()
	}
	if !(plan.MaxTxPower.IsNull() || plan.MaxTxPower.IsUnknown()) {
		body["max-tx-power"] = plan.MaxTxPower.ValueString()
	}
	if !(plan.Mode.IsNull() || plan.Mode.IsUnknown()) {
		body["mode"] = plan.Mode.ValueString()
	}
	if !(plan.MultiPassphraseGroup.IsNull() || plan.MultiPassphraseGroup.IsUnknown()) {
		body["multi-passphrase-group"] = plan.MultiPassphraseGroup.ValueString()
	}
	if !(plan.MulticastEnhance.IsNull() || plan.MulticastEnhance.IsUnknown()) {
		body["multicast-enhance"] = plan.MulticastEnhance.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.NasIdentifier.IsNull() || plan.NasIdentifier.IsUnknown()) {
		body["aaa.nas-identifier"] = plan.NasIdentifier.ValueString()
	}
	if !(plan.NeighborGroup.IsNull() || plan.NeighborGroup.IsUnknown()) {
		body["steering.neighbor-group"] = plan.NeighborGroup.ValueString()
	}
	if !(plan.NetworkType.IsNull() || plan.NetworkType.IsUnknown()) {
		body["network-type"] = plan.NetworkType.ValueString()
	}
	if !(plan.OpenFlowSwitch.IsNull() || plan.OpenFlowSwitch.IsUnknown()) {
		body["open-flow-switch"] = plan.OpenFlowSwitch.ValueString()
	}
	if !(plan.Openflow.IsNull() || plan.Openflow.IsUnknown()) {
		body["openflow"] = plan.Openflow.ValueString()
	}
	if !(plan.OperationalClasses.IsNull() || plan.OperationalClasses.IsUnknown()) {
		body["operational-classes"] = plan.OperationalClasses.ValueString()
	}
	if !(plan.OperatorNames.IsNull() || plan.OperatorNames.IsUnknown()) {
		body["operator-names"] = plan.OperatorNames.ValueString()
	}
	if !(plan.OweTransitionInterface.IsNull() || plan.OweTransitionInterface.IsUnknown()) {
		body["security.owe-transition-interface"] = plan.OweTransitionInterface.ValueString()
	}
	if !(plan.Passphrase.IsNull() || plan.Passphrase.IsUnknown()) {
		body["security.passphrase"] = plan.Passphrase.ValueString()
	}
	if !(plan.PasswordFormat.IsNull() || plan.PasswordFormat.IsUnknown()) {
		body["aaa.password-format"] = plan.PasswordFormat.ValueString()
	}
	if !(plan.QoSClassifier.IsNull() || plan.QoSClassifier.IsUnknown()) {
		body["qo-s-classifier"] = plan.QoSClassifier.ValueString()
	}
	if !(plan.Realms.IsNull() || plan.Realms.IsUnknown()) {
		body["realms"] = plan.Realms.ValueString()
	}
	if !(plan.RealmsRaw.IsNull() || plan.RealmsRaw.IsUnknown()) {
		body["realms-raw"] = plan.RealmsRaw.ValueString()
	}
	if !(plan.ReselectInterval.IsNull() || plan.ReselectInterval.IsUnknown()) {
		body["channel.reselect-interval"] = plan.ReselectInterval.ValueString()
	}
	if !(plan.ReselectTime.IsNull() || plan.ReselectTime.IsUnknown()) {
		body["channel.reselect-time"] = plan.ReselectTime.ValueString()
	}
	if !(plan.RoamingOis.IsNull() || plan.RoamingOis.IsUnknown()) {
		body["roaming-ois"] = plan.RoamingOis.ValueString()
	}
	if !(plan.Rrm.IsNull() || plan.Rrm.IsUnknown()) {
		body["steering.rrm"] = plan.Rrm.ValueString()
	}
	if !(plan.SaeAntiCloggingThreshold.IsNull() || plan.SaeAntiCloggingThreshold.IsUnknown()) {
		body["sae-anti-clogging-threshold"] = plan.SaeAntiCloggingThreshold.ValueString()
	}
	if !(plan.SaeMaxFailureRate.IsNull() || plan.SaeMaxFailureRate.IsUnknown()) {
		body["security.sae-max-failure-rate"] = plan.SaeMaxFailureRate.ValueString()
	}
	if !(plan.SaePwe.IsNull() || plan.SaePwe.IsUnknown()) {
		body["security.sae-pwe"] = plan.SaePwe.ValueString()
	}
	if !(plan.SecondaryFrequency.IsNull() || plan.SecondaryFrequency.IsUnknown()) {
		body["channel.secondary-frequency"] = plan.SecondaryFrequency.ValueString()
	}
	if !(plan.Security.IsNull() || plan.Security.IsUnknown()) {
		body["security"] = plan.Security.ValueString()
	}
	if !(plan.SkipDfsChannels.IsNull() || plan.SkipDfsChannels.IsUnknown()) {
		body["channel.skip-dfs-channels"] = plan.SkipDfsChannels.ValueString()
	}
	if !(plan.Ssid.IsNull() || plan.Ssid.IsUnknown()) {
		body["ssid"] = plan.Ssid.ValueString()
	}
	if !(plan.StationRoaming.IsNull() || plan.StationRoaming.IsUnknown()) {
		body["station-roaming"] = plan.StationRoaming.ValueString()
	}
	if !(plan.Steering.IsNull() || plan.Steering.IsUnknown()) {
		body["steering"] = plan.Steering.ValueString()
	}
	if !(plan.TrafficProcessing.IsNull() || plan.TrafficProcessing.IsUnknown()) {
		body["datapath.traffic-processing"] = plan.TrafficProcessing.ValueString()
	}
	if !(plan.TransitionRequestCount.IsNull() || plan.TransitionRequestCount.IsUnknown()) {
		body["transition-request-count"] = plan.TransitionRequestCount.ValueString()
	}
	if !(plan.TransitionThreshold.IsNull() || plan.TransitionThreshold.IsUnknown()) {
		body["transition-threshold"] = plan.TransitionThreshold.ValueString()
	}
	if !(plan.TransitionThresholdPeriod.IsNull() || plan.TransitionThresholdPeriod.IsUnknown()) {
		body["transition-threshold-period"] = plan.TransitionThresholdPeriod.ValueString()
	}
	if !(plan.TransitionThresholdTime.IsNull() || plan.TransitionThresholdTime.IsUnknown()) {
		body["transition-threshold-time"] = plan.TransitionThresholdTime.ValueString()
	}
	if !(plan.TransitionTime.IsNull() || plan.TransitionTime.IsUnknown()) {
		body["transition-time"] = plan.TransitionTime.ValueString()
	}
	if !(plan.TxChains.IsNull() || plan.TxChains.IsUnknown()) {
		body["tx-chains"] = plan.TxChains.ValueString()
	}
	if !(plan.TxPower.IsNull() || plan.TxPower.IsUnknown()) {
		body["tx-power"] = plan.TxPower.ValueString()
	}
	if !(plan.Types.IsNull() || plan.Types.IsUnknown()) {
		body["types"] = plan.Types.ValueString()
	}
	if !(plan.Uesa.IsNull() || plan.Uesa.IsUnknown()) {
		body["uesa"] = plan.Uesa.ValueString()
	}
	if !(plan.UsernameFormat.IsNull() || plan.UsernameFormat.IsUnknown()) {
		body["aaa.username-format"] = plan.UsernameFormat.ValueString()
	}
	if !(plan.Venue.IsNull() || plan.Venue.IsUnknown()) {
		body["venue"] = plan.Venue.ValueString()
	}
	if !(plan.VenueNames.IsNull() || plan.VenueNames.IsUnknown()) {
		body["venue-names"] = plan.VenueNames.ValueString()
	}
	if !(plan.VLANID.IsNull() || plan.VLANID.IsUnknown()) {
		body["datapath.vlan-id"] = plan.VLANID.ValueString()
	}
	if !(plan.WanAtCapacity.IsNull() || plan.WanAtCapacity.IsUnknown()) {
		body["wan-at-capacity"] = plan.WanAtCapacity.ValueString()
	}
	if !(plan.WanDownlink.IsNull() || plan.WanDownlink.IsUnknown()) {
		body["wan-downlink"] = plan.WanDownlink.ValueString()
	}
	if !(plan.WanDownlinkLoad.IsNull() || plan.WanDownlinkLoad.IsUnknown()) {
		body["wan-downlink-load"] = plan.WanDownlinkLoad.ValueString()
	}
	if !(plan.WanMeasurementDuration.IsNull() || plan.WanMeasurementDuration.IsUnknown()) {
		body["wan-measurement-duration"] = plan.WanMeasurementDuration.ValueString()
	}
	if !(plan.WanStatus.IsNull() || plan.WanStatus.IsUnknown()) {
		body["wan-status"] = plan.WanStatus.ValueString()
	}
	if !(plan.WanSymmetric.IsNull() || plan.WanSymmetric.IsUnknown()) {
		body["wan-symmetric"] = plan.WanSymmetric.ValueString()
	}
	if !(plan.WanUplink.IsNull() || plan.WanUplink.IsUnknown()) {
		body["wan-uplink"] = plan.WanUplink.ValueString()
	}
	if !(plan.WanUplinkLoad.IsNull() || plan.WanUplinkLoad.IsUnknown()) {
		body["wan-uplink-load"] = plan.WanUplinkLoad.ValueString()
	}
	if !(plan.Wnm.IsNull() || plan.Wnm.IsUnknown()) {
		body["steering.wnm"] = plan.Wnm.ValueString()
	}
	if !(plan.Wps.IsNull() || plan.Wps.IsUnknown()) {
		body["security.wps"] = plan.Wps.ValueString()
	}
	if !(plan.QosClassifier.IsNull() || plan.QosClassifier.IsUnknown()) {
		body["qos-classifier"] = plan.QosClassifier.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/wifi/configuration", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/wifi/configuration failed", err.Error())
		return
	}
	interfaceWifiConfigurationApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceWifiConfigurationModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/wifi/configuration", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/wifi/configuration failed", err.Error())
		return
	}
	interfaceWifiConfigurationApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceWifiConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceWifiConfigurationModel
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
	if !plan.X2gProbeDelay.Equal(state.X2gProbeDelay) && !plan.X2gProbeDelay.IsUnknown() {
		body["2g-probe-delay"] = plan.X2gProbeDelay.ValueString()
	}
	if !plan.X3gppInfo.Equal(state.X3gppInfo) && !plan.X3gppInfo.IsUnknown() {
		body["3gpp-info"] = plan.X3gppInfo.ValueString()
	}
	if !plan.X3gppInfoRaw.Equal(state.X3gppInfoRaw) && !plan.X3gppInfoRaw.IsUnknown() {
		body["3gpp-info-raw"] = plan.X3gppInfoRaw.ValueString()
	}
	if !plan.Aaa.Equal(state.Aaa) && !plan.Aaa.IsUnknown() {
		body["aaa"] = plan.Aaa.ValueString()
	}
	if !plan.AntennaGain.Equal(state.AntennaGain) && !plan.AntennaGain.IsUnknown() {
		body["antenna-gain"] = plan.AntennaGain.ValueString()
	}
	if !plan.AuthenticationTypes.Equal(state.AuthenticationTypes) && !plan.AuthenticationTypes.IsUnknown() {
		body["security.authentication-types"] = plan.AuthenticationTypes.ValueString()
	}
	if !plan.Band.Equal(state.Band) && !plan.Band.IsUnknown() {
		body["channel.band"] = plan.Band.ValueString()
	}
	if !plan.BeaconInterval.Equal(state.BeaconInterval) && !plan.BeaconInterval.IsUnknown() {
		body["beacon-interval"] = plan.BeaconInterval.ValueString()
	}
	if !plan.BeaconProtection.Equal(state.BeaconProtection) && !plan.BeaconProtection.IsUnknown() {
		body["beacon-protection"] = plan.BeaconProtection.ValueString()
	}
	if !plan.Bridge.Equal(state.Bridge) && !plan.Bridge.IsUnknown() {
		body["datapath.bridge"] = plan.Bridge.ValueString()
	}
	if !plan.BridgeCost.Equal(state.BridgeCost) && !plan.BridgeCost.IsUnknown() {
		body["datapath.bridge-cost"] = plan.BridgeCost.ValueString()
	}
	if !plan.BridgeHorizon.Equal(state.BridgeHorizon) && !plan.BridgeHorizon.IsUnknown() {
		body["datapath.bridge-horizon"] = plan.BridgeHorizon.ValueString()
	}
	if !plan.CalledFormat.Equal(state.CalledFormat) && !plan.CalledFormat.IsUnknown() {
		body["aaa.called-format"] = plan.CalledFormat.ValueString()
	}
	if !plan.CallingFormat.Equal(state.CallingFormat) && !plan.CallingFormat.IsUnknown() {
		body["aaa.calling-format"] = plan.CallingFormat.ValueString()
	}
	if !plan.Chains.Equal(state.Chains) && !plan.Chains.IsUnknown() {
		body["chains"] = plan.Chains.ValueString()
	}
	if !plan.Channel.Equal(state.Channel) && !plan.Channel.IsUnknown() {
		body["channel"] = plan.Channel.ValueString()
	}
	if !plan.ChannelWidth.Equal(state.ChannelWidth) && !plan.ChannelWidth.IsUnknown() {
		body["channel-width"] = plan.ChannelWidth.ValueString()
	}
	if !plan.Ciphers.Equal(state.Ciphers) && !plan.Ciphers.IsUnknown() {
		body["ciphers"] = plan.Ciphers.ValueString()
	}
	if !plan.ClientIsolation.Equal(state.ClientIsolation) && !plan.ClientIsolation.IsUnknown() {
		body["datapath.client-isolation"] = plan.ClientIsolation.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.ConnectGroup.Equal(state.ConnectGroup) && !plan.ConnectGroup.IsUnknown() {
		body["connect-group"] = plan.ConnectGroup.ValueString()
	}
	if !plan.ConnectPriority.Equal(state.ConnectPriority) && !plan.ConnectPriority.IsUnknown() {
		body["security.connect-priority"] = plan.ConnectPriority.ValueString()
	}
	if !plan.ConnectionCapabilities.Equal(state.ConnectionCapabilities) && !plan.ConnectionCapabilities.IsUnknown() {
		body["connection-capabilities"] = plan.ConnectionCapabilities.ValueString()
	}
	if !plan.Country.Equal(state.Country) && !plan.Country.IsUnknown() {
		body["country"] = plan.Country.ValueString()
	}
	if !plan.Datapath.Equal(state.Datapath) && !plan.Datapath.IsUnknown() {
		body["datapath"] = plan.Datapath.ValueString()
	}
	if !plan.DeprioritizeUnii34.Equal(state.DeprioritizeUnii34) && !plan.DeprioritizeUnii34.IsUnknown() {
		body["deprioritize-unii-3-4"] = plan.DeprioritizeUnii34.ValueString()
	}
	if !plan.Dgaf.Equal(state.Dgaf) && !plan.Dgaf.IsUnknown() {
		body["dgaf"] = plan.Dgaf.ValueString()
	}
	if !plan.DhGroups.Equal(state.DhGroups) && !plan.DhGroups.IsUnknown() {
		body["security.dh-groups"] = plan.DhGroups.ValueString()
	}
	if !plan.DisablePmkid.Equal(state.DisablePmkid) && !plan.DisablePmkid.IsUnknown() {
		body["security.disable-pmkid"] = plan.DisablePmkid.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Distance.Equal(state.Distance) && !plan.Distance.IsUnknown() {
		body["distance"] = plan.Distance.ValueString()
	}
	if !plan.DomainNames.Equal(state.DomainNames) && !plan.DomainNames.IsUnknown() {
		body["domain-names"] = plan.DomainNames.ValueString()
	}
	if !plan.DtimPeriod.Equal(state.DtimPeriod) && !plan.DtimPeriod.IsUnknown() {
		body["dtim-period"] = plan.DtimPeriod.ValueString()
	}
	if !plan.EAPAccounting.Equal(state.EAPAccounting) && !plan.EAPAccounting.IsUnknown() {
		body["security.eap-accounting"] = plan.EAPAccounting.ValueString()
	}
	if !plan.EAPAnonymousIdentity.Equal(state.EAPAnonymousIdentity) && !plan.EAPAnonymousIdentity.IsUnknown() {
		body["security.eap-anonymous-identity"] = plan.EAPAnonymousIdentity.ValueString()
	}
	if !plan.EAPCertificateMode.Equal(state.EAPCertificateMode) && !plan.EAPCertificateMode.IsUnknown() {
		body["security.eap-certificate-mode"] = plan.EAPCertificateMode.ValueString()
	}
	if !plan.EAPMethods.Equal(state.EAPMethods) && !plan.EAPMethods.IsUnknown() {
		body["security.eap-methods"] = plan.EAPMethods.ValueString()
	}
	if !plan.EAPPassword.Equal(state.EAPPassword) && !plan.EAPPassword.IsUnknown() {
		body["security.eap-password"] = plan.EAPPassword.ValueString()
	}
	if !plan.EAPTLSCertificate.Equal(state.EAPTLSCertificate) && !plan.EAPTLSCertificate.IsUnknown() {
		body["eap-tls-certificate"] = plan.EAPTLSCertificate.ValueString()
	}
	if !plan.EAPUsername.Equal(state.EAPUsername) && !plan.EAPUsername.IsUnknown() {
		body["security.eap-username"] = plan.EAPUsername.ValueString()
	}
	if !plan.Esr.Equal(state.Esr) && !plan.Esr.IsUnknown() {
		body["esr"] = plan.Esr.ValueString()
	}
	if !plan.Frequency.Equal(state.Frequency) && !plan.Frequency.IsUnknown() {
		body["channel.frequency"] = plan.Frequency.ValueString()
	}
	if !plan.FtEnabled.Equal(state.FtEnabled) && !plan.FtEnabled.IsUnknown() {
		body["ft-enabled"] = plan.FtEnabled.ValueString()
	}
	if !plan.FtMobilityDomain.Equal(state.FtMobilityDomain) && !plan.FtMobilityDomain.IsUnknown() {
		body["ft-mobility-domain"] = plan.FtMobilityDomain.ValueString()
	}
	if !plan.FtNasIdentifier.Equal(state.FtNasIdentifier) && !plan.FtNasIdentifier.IsUnknown() {
		body["ft-nas-identifier"] = plan.FtNasIdentifier.ValueString()
	}
	if !plan.FtOverDs.Equal(state.FtOverDs) && !plan.FtOverDs.IsUnknown() {
		body["security.ft-over-ds"] = plan.FtOverDs.ValueString()
	}
	if !plan.FtPreserveVLANID.Equal(state.FtPreserveVLANID) && !plan.FtPreserveVLANID.IsUnknown() {
		body["ft-preserve-vlan-id"] = plan.FtPreserveVLANID.ValueString()
	}
	if !plan.FtR0KeyLifetime.Equal(state.FtR0KeyLifetime) && !plan.FtR0KeyLifetime.IsUnknown() {
		body["ft-r0-key-lifetime"] = plan.FtR0KeyLifetime.ValueString()
	}
	if !plan.FtReassocDeadline.Equal(state.FtReassocDeadline) && !plan.FtReassocDeadline.IsUnknown() {
		body["ft-reassoc-deadline"] = plan.FtReassocDeadline.ValueString()
	}
	if !plan.GroupEncryption.Equal(state.GroupEncryption) && !plan.GroupEncryption.IsUnknown() {
		body["security.group-encryption"] = plan.GroupEncryption.ValueString()
	}
	if !plan.GroupKeyUpdate.Equal(state.GroupKeyUpdate) && !plan.GroupKeyUpdate.IsUnknown() {
		body["security.group-key-update"] = plan.GroupKeyUpdate.ValueString()
	}
	if !plan.Hessid.Equal(state.Hessid) && !plan.Hessid.IsUnknown() {
		body["hessid"] = plan.Hessid.ValueString()
	}
	if !plan.HideSsid.Equal(state.HideSsid) && !plan.HideSsid.IsUnknown() {
		body["hide-ssid"] = plan.HideSsid.ValueString()
	}
	if !plan.Hotspot20.Equal(state.Hotspot20) && !plan.Hotspot20.IsUnknown() {
		body["hotspot-2-0"] = plan.Hotspot20.ValueString()
	}
	if !plan.HwProtectionMode.Equal(state.HwProtectionMode) && !plan.HwProtectionMode.IsUnknown() {
		body["hw-protection-mode"] = plan.HwProtectionMode.ValueString()
	}
	if !plan.Installation.Equal(state.Installation) && !plan.Installation.IsUnknown() {
		body["installation"] = plan.Installation.ValueString()
	}
	if !plan.InterfaceList.Equal(state.InterfaceList) && !plan.InterfaceList.IsUnknown() {
		body["datapath.interface-list"] = plan.InterfaceList.ValueString()
	}
	if !plan.InterimUpdate.Equal(state.InterimUpdate) && !plan.InterimUpdate.IsUnknown() {
		body["aaa.interim-update"] = plan.InterimUpdate.ValueString()
	}
	if !plan.Internet.Equal(state.Internet) && !plan.Internet.IsUnknown() {
		body["internet"] = plan.Internet.ValueString()
	}
	if !plan.Interworking.Equal(state.Interworking) && !plan.Interworking.IsUnknown() {
		body["interworking"] = plan.Interworking.ValueString()
	}
	if !plan.Ipv4Availability.Equal(state.Ipv4Availability) && !plan.Ipv4Availability.IsUnknown() {
		body["ipv4-availability"] = plan.Ipv4Availability.ValueString()
	}
	if !plan.IPV6Availability.Equal(state.IPV6Availability) && !plan.IPV6Availability.IsUnknown() {
		body["ipv6-availability"] = plan.IPV6Availability.ValueString()
	}
	if !plan.MACCaching.Equal(state.MACCaching) && !plan.MACCaching.IsUnknown() {
		body["aaa.mac-caching"] = plan.MACCaching.ValueString()
	}
	if !plan.ManagementEncryption.Equal(state.ManagementEncryption) && !plan.ManagementEncryption.IsUnknown() {
		body["management-encryption"] = plan.ManagementEncryption.ValueString()
	}
	if !plan.ManagementProtection.Equal(state.ManagementProtection) && !plan.ManagementProtection.IsUnknown() {
		body["security.management-protection"] = plan.ManagementProtection.ValueString()
	}
	if !plan.Manager.Equal(state.Manager) && !plan.Manager.IsUnknown() {
		body["manager"] = plan.Manager.ValueString()
	}
	if !plan.MaxClients.Equal(state.MaxClients) && !plan.MaxClients.IsUnknown() {
		body["max-clients"] = plan.MaxClients.ValueString()
	}
	if !plan.MaxTxPower.Equal(state.MaxTxPower) && !plan.MaxTxPower.IsUnknown() {
		body["max-tx-power"] = plan.MaxTxPower.ValueString()
	}
	if !plan.Mode.Equal(state.Mode) && !plan.Mode.IsUnknown() {
		body["mode"] = plan.Mode.ValueString()
	}
	if !plan.MultiPassphraseGroup.Equal(state.MultiPassphraseGroup) && !plan.MultiPassphraseGroup.IsUnknown() {
		body["multi-passphrase-group"] = plan.MultiPassphraseGroup.ValueString()
	}
	if !plan.MulticastEnhance.Equal(state.MulticastEnhance) && !plan.MulticastEnhance.IsUnknown() {
		body["multicast-enhance"] = plan.MulticastEnhance.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.NasIdentifier.Equal(state.NasIdentifier) && !plan.NasIdentifier.IsUnknown() {
		body["aaa.nas-identifier"] = plan.NasIdentifier.ValueString()
	}
	if !plan.NeighborGroup.Equal(state.NeighborGroup) && !plan.NeighborGroup.IsUnknown() {
		body["steering.neighbor-group"] = plan.NeighborGroup.ValueString()
	}
	if !plan.NetworkType.Equal(state.NetworkType) && !plan.NetworkType.IsUnknown() {
		body["network-type"] = plan.NetworkType.ValueString()
	}
	if !plan.OpenFlowSwitch.Equal(state.OpenFlowSwitch) && !plan.OpenFlowSwitch.IsUnknown() {
		body["open-flow-switch"] = plan.OpenFlowSwitch.ValueString()
	}
	if !plan.Openflow.Equal(state.Openflow) && !plan.Openflow.IsUnknown() {
		body["openflow"] = plan.Openflow.ValueString()
	}
	if !plan.OperationalClasses.Equal(state.OperationalClasses) && !plan.OperationalClasses.IsUnknown() {
		body["operational-classes"] = plan.OperationalClasses.ValueString()
	}
	if !plan.OperatorNames.Equal(state.OperatorNames) && !plan.OperatorNames.IsUnknown() {
		body["operator-names"] = plan.OperatorNames.ValueString()
	}
	if !plan.OweTransitionInterface.Equal(state.OweTransitionInterface) && !plan.OweTransitionInterface.IsUnknown() {
		body["security.owe-transition-interface"] = plan.OweTransitionInterface.ValueString()
	}
	if !plan.Passphrase.Equal(state.Passphrase) && !plan.Passphrase.IsUnknown() {
		body["security.passphrase"] = plan.Passphrase.ValueString()
	}
	if !plan.PasswordFormat.Equal(state.PasswordFormat) && !plan.PasswordFormat.IsUnknown() {
		body["aaa.password-format"] = plan.PasswordFormat.ValueString()
	}
	if !plan.QoSClassifier.Equal(state.QoSClassifier) && !plan.QoSClassifier.IsUnknown() {
		body["qo-s-classifier"] = plan.QoSClassifier.ValueString()
	}
	if !plan.Realms.Equal(state.Realms) && !plan.Realms.IsUnknown() {
		body["realms"] = plan.Realms.ValueString()
	}
	if !plan.RealmsRaw.Equal(state.RealmsRaw) && !plan.RealmsRaw.IsUnknown() {
		body["realms-raw"] = plan.RealmsRaw.ValueString()
	}
	if !plan.ReselectInterval.Equal(state.ReselectInterval) && !plan.ReselectInterval.IsUnknown() {
		body["channel.reselect-interval"] = plan.ReselectInterval.ValueString()
	}
	if !plan.ReselectTime.Equal(state.ReselectTime) && !plan.ReselectTime.IsUnknown() {
		body["channel.reselect-time"] = plan.ReselectTime.ValueString()
	}
	if !plan.RoamingOis.Equal(state.RoamingOis) && !plan.RoamingOis.IsUnknown() {
		body["roaming-ois"] = plan.RoamingOis.ValueString()
	}
	if !plan.Rrm.Equal(state.Rrm) && !plan.Rrm.IsUnknown() {
		body["steering.rrm"] = plan.Rrm.ValueString()
	}
	if !plan.SaeAntiCloggingThreshold.Equal(state.SaeAntiCloggingThreshold) && !plan.SaeAntiCloggingThreshold.IsUnknown() {
		body["sae-anti-clogging-threshold"] = plan.SaeAntiCloggingThreshold.ValueString()
	}
	if !plan.SaeMaxFailureRate.Equal(state.SaeMaxFailureRate) && !plan.SaeMaxFailureRate.IsUnknown() {
		body["security.sae-max-failure-rate"] = plan.SaeMaxFailureRate.ValueString()
	}
	if !plan.SaePwe.Equal(state.SaePwe) && !plan.SaePwe.IsUnknown() {
		body["security.sae-pwe"] = plan.SaePwe.ValueString()
	}
	if !plan.SecondaryFrequency.Equal(state.SecondaryFrequency) && !plan.SecondaryFrequency.IsUnknown() {
		body["channel.secondary-frequency"] = plan.SecondaryFrequency.ValueString()
	}
	if !plan.Security.Equal(state.Security) && !plan.Security.IsUnknown() {
		body["security"] = plan.Security.ValueString()
	}
	if !plan.SkipDfsChannels.Equal(state.SkipDfsChannels) && !plan.SkipDfsChannels.IsUnknown() {
		body["channel.skip-dfs-channels"] = plan.SkipDfsChannels.ValueString()
	}
	if !plan.Ssid.Equal(state.Ssid) && !plan.Ssid.IsUnknown() {
		body["ssid"] = plan.Ssid.ValueString()
	}
	if !plan.StationRoaming.Equal(state.StationRoaming) && !plan.StationRoaming.IsUnknown() {
		body["station-roaming"] = plan.StationRoaming.ValueString()
	}
	if !plan.Steering.Equal(state.Steering) && !plan.Steering.IsUnknown() {
		body["steering"] = plan.Steering.ValueString()
	}
	if !plan.TrafficProcessing.Equal(state.TrafficProcessing) && !plan.TrafficProcessing.IsUnknown() {
		body["datapath.traffic-processing"] = plan.TrafficProcessing.ValueString()
	}
	if !plan.TransitionRequestCount.Equal(state.TransitionRequestCount) && !plan.TransitionRequestCount.IsUnknown() {
		body["transition-request-count"] = plan.TransitionRequestCount.ValueString()
	}
	if !plan.TransitionThreshold.Equal(state.TransitionThreshold) && !plan.TransitionThreshold.IsUnknown() {
		body["transition-threshold"] = plan.TransitionThreshold.ValueString()
	}
	if !plan.TransitionThresholdPeriod.Equal(state.TransitionThresholdPeriod) && !plan.TransitionThresholdPeriod.IsUnknown() {
		body["transition-threshold-period"] = plan.TransitionThresholdPeriod.ValueString()
	}
	if !plan.TransitionThresholdTime.Equal(state.TransitionThresholdTime) && !plan.TransitionThresholdTime.IsUnknown() {
		body["transition-threshold-time"] = plan.TransitionThresholdTime.ValueString()
	}
	if !plan.TransitionTime.Equal(state.TransitionTime) && !plan.TransitionTime.IsUnknown() {
		body["transition-time"] = plan.TransitionTime.ValueString()
	}
	if !plan.TxChains.Equal(state.TxChains) && !plan.TxChains.IsUnknown() {
		body["tx-chains"] = plan.TxChains.ValueString()
	}
	if !plan.TxPower.Equal(state.TxPower) && !plan.TxPower.IsUnknown() {
		body["tx-power"] = plan.TxPower.ValueString()
	}
	if !plan.Types.Equal(state.Types) && !plan.Types.IsUnknown() {
		body["types"] = plan.Types.ValueString()
	}
	if !plan.Uesa.Equal(state.Uesa) && !plan.Uesa.IsUnknown() {
		body["uesa"] = plan.Uesa.ValueString()
	}
	if !plan.UsernameFormat.Equal(state.UsernameFormat) && !plan.UsernameFormat.IsUnknown() {
		body["aaa.username-format"] = plan.UsernameFormat.ValueString()
	}
	if !plan.Venue.Equal(state.Venue) && !plan.Venue.IsUnknown() {
		body["venue"] = plan.Venue.ValueString()
	}
	if !plan.VenueNames.Equal(state.VenueNames) && !plan.VenueNames.IsUnknown() {
		body["venue-names"] = plan.VenueNames.ValueString()
	}
	if !plan.VLANID.Equal(state.VLANID) && !plan.VLANID.IsUnknown() {
		body["datapath.vlan-id"] = plan.VLANID.ValueString()
	}
	if !plan.WanAtCapacity.Equal(state.WanAtCapacity) && !plan.WanAtCapacity.IsUnknown() {
		body["wan-at-capacity"] = plan.WanAtCapacity.ValueString()
	}
	if !plan.WanDownlink.Equal(state.WanDownlink) && !plan.WanDownlink.IsUnknown() {
		body["wan-downlink"] = plan.WanDownlink.ValueString()
	}
	if !plan.WanDownlinkLoad.Equal(state.WanDownlinkLoad) && !plan.WanDownlinkLoad.IsUnknown() {
		body["wan-downlink-load"] = plan.WanDownlinkLoad.ValueString()
	}
	if !plan.WanMeasurementDuration.Equal(state.WanMeasurementDuration) && !plan.WanMeasurementDuration.IsUnknown() {
		body["wan-measurement-duration"] = plan.WanMeasurementDuration.ValueString()
	}
	if !plan.WanStatus.Equal(state.WanStatus) && !plan.WanStatus.IsUnknown() {
		body["wan-status"] = plan.WanStatus.ValueString()
	}
	if !plan.WanSymmetric.Equal(state.WanSymmetric) && !plan.WanSymmetric.IsUnknown() {
		body["wan-symmetric"] = plan.WanSymmetric.ValueString()
	}
	if !plan.WanUplink.Equal(state.WanUplink) && !plan.WanUplink.IsUnknown() {
		body["wan-uplink"] = plan.WanUplink.ValueString()
	}
	if !plan.WanUplinkLoad.Equal(state.WanUplinkLoad) && !plan.WanUplinkLoad.IsUnknown() {
		body["wan-uplink-load"] = plan.WanUplinkLoad.ValueString()
	}
	if !plan.Wnm.Equal(state.Wnm) && !plan.Wnm.IsUnknown() {
		body["steering.wnm"] = plan.Wnm.ValueString()
	}
	if !plan.Wps.Equal(state.Wps) && !plan.Wps.IsUnknown() {
		body["security.wps"] = plan.Wps.ValueString()
	}
	if !plan.QosClassifier.Equal(state.QosClassifier) && !plan.QosClassifier.IsUnknown() {
		body["qos-classifier"] = plan.QosClassifier.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/wifi/configuration", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/wifi/configuration failed", err.Error())
			return
		}
		interfaceWifiConfigurationApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceWifiConfigurationModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/wifi/configuration", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/wifi/configuration failed", err.Error())
	}
}

func (r *InterfaceWifiConfigurationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceWifiConfigurationLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/wifi/configuration matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceWifiConfigurationLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceWifiConfigurationLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/wifi/configuration", id)
}

func interfaceWifiConfigurationApply(ctx context.Context, obj client.Object, m *InterfaceWifiConfigurationModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["qos-classifier"]; ok && v != "" {
		m.QosClassifier = types.StringValue(v)
	} else {
		m.QosClassifier = types.StringNull()
	}
	if v, ok := obj["2g-probe-delay"]; ok {
		_ = v
		if v != "" {
			m.X2gProbeDelay = types.StringValue(v)
		} else {
			m.X2gProbeDelay = types.StringNull()
		}
	} else {
		m.X2gProbeDelay = types.StringNull()
	}
	if v, ok := obj["3gpp-info"]; ok {
		_ = v
		if v != "" {
			m.X3gppInfo = types.StringValue(v)
		} else {
			m.X3gppInfo = types.StringNull()
		}
	} else {
		m.X3gppInfo = types.StringNull()
	}
	if v, ok := obj["3gpp-info-raw"]; ok {
		_ = v
		if v != "" {
			m.X3gppInfoRaw = types.StringValue(v)
		} else {
			m.X3gppInfoRaw = types.StringNull()
		}
	} else {
		m.X3gppInfoRaw = types.StringNull()
	}
	if v, ok := obj["aaa"]; ok {
		_ = v
		if v != "" {
			m.Aaa = types.StringValue(v)
		} else {
			m.Aaa = types.StringNull()
		}
	} else {
		m.Aaa = types.StringNull()
	}
	if v, ok := obj["antenna-gain"]; ok {
		_ = v
		if v != "" {
			m.AntennaGain = types.StringValue(v)
		} else {
			m.AntennaGain = types.StringNull()
		}
	} else {
		m.AntennaGain = types.StringNull()
	}
	if v, ok := obj["security.authentication-types"]; ok {
		_ = v
		if v != "" {
			m.AuthenticationTypes = types.StringValue(v)
		} else {
			m.AuthenticationTypes = types.StringNull()
		}
	} else {
		m.AuthenticationTypes = types.StringNull()
	}
	if v, ok := obj["channel.band"]; ok {
		_ = v
		if v != "" {
			m.Band = types.StringValue(v)
		} else {
			m.Band = types.StringNull()
		}
	} else {
		m.Band = types.StringNull()
	}
	if v, ok := obj["beacon-interval"]; ok {
		_ = v
		if v != "" {
			m.BeaconInterval = types.StringValue(v)
		} else {
			m.BeaconInterval = types.StringNull()
		}
	} else {
		m.BeaconInterval = types.StringNull()
	}
	if v, ok := obj["beacon-protection"]; ok {
		_ = v
		if v != "" {
			m.BeaconProtection = types.StringValue(v)
		} else {
			m.BeaconProtection = types.StringNull()
		}
	} else {
		m.BeaconProtection = types.StringNull()
	}
	if v, ok := obj["datapath.bridge"]; ok {
		_ = v
		if v != "" {
			m.Bridge = types.StringValue(v)
		} else {
			m.Bridge = types.StringNull()
		}
	} else {
		m.Bridge = types.StringNull()
	}
	if v, ok := obj["datapath.bridge-cost"]; ok {
		_ = v
		if v != "" {
			m.BridgeCost = types.StringValue(v)
		} else {
			m.BridgeCost = types.StringNull()
		}
	} else {
		m.BridgeCost = types.StringNull()
	}
	if v, ok := obj["datapath.bridge-horizon"]; ok {
		_ = v
		if v != "" {
			m.BridgeHorizon = types.StringValue(v)
		} else {
			m.BridgeHorizon = types.StringNull()
		}
	} else {
		m.BridgeHorizon = types.StringNull()
	}
	if v, ok := obj["aaa.called-format"]; ok {
		_ = v
		if v != "" {
			m.CalledFormat = types.StringValue(v)
		} else {
			m.CalledFormat = types.StringNull()
		}
	} else {
		m.CalledFormat = types.StringNull()
	}
	if v, ok := obj["aaa.calling-format"]; ok {
		_ = v
		if v != "" {
			m.CallingFormat = types.StringValue(v)
		} else {
			m.CallingFormat = types.StringNull()
		}
	} else {
		m.CallingFormat = types.StringNull()
	}
	if v, ok := obj["chains"]; ok {
		_ = v
		if v != "" {
			m.Chains = types.StringValue(v)
		} else {
			m.Chains = types.StringNull()
		}
	} else {
		m.Chains = types.StringNull()
	}
	if v, ok := obj["channel"]; ok {
		_ = v
		if v != "" {
			m.Channel = types.StringValue(v)
		} else {
			m.Channel = types.StringNull()
		}
	} else {
		m.Channel = types.StringNull()
	}
	if v, ok := obj["channel-width"]; ok {
		_ = v
		if v != "" {
			m.ChannelWidth = types.StringValue(v)
		} else {
			m.ChannelWidth = types.StringNull()
		}
	} else {
		m.ChannelWidth = types.StringNull()
	}
	if v, ok := obj["ciphers"]; ok {
		_ = v
		if v != "" {
			m.Ciphers = types.StringValue(v)
		} else {
			m.Ciphers = types.StringNull()
		}
	} else {
		m.Ciphers = types.StringNull()
	}
	if v, ok := obj["datapath.client-isolation"]; ok {
		_ = v
		if v != "" {
			m.ClientIsolation = types.StringValue(v)
		} else {
			m.ClientIsolation = types.StringNull()
		}
	} else {
		m.ClientIsolation = types.StringNull()
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
	if v, ok := obj["connect-group"]; ok {
		_ = v
		if v != "" {
			m.ConnectGroup = types.StringValue(v)
		} else {
			m.ConnectGroup = types.StringNull()
		}
	} else {
		m.ConnectGroup = types.StringNull()
	}
	if v, ok := obj["security.connect-priority"]; ok {
		_ = v
		if v != "" {
			m.ConnectPriority = types.StringValue(v)
		} else {
			m.ConnectPriority = types.StringNull()
		}
	} else {
		m.ConnectPriority = types.StringNull()
	}
	if v, ok := obj["connection-capabilities"]; ok {
		_ = v
		if v != "" {
			m.ConnectionCapabilities = types.StringValue(v)
		} else {
			m.ConnectionCapabilities = types.StringNull()
		}
	} else {
		m.ConnectionCapabilities = types.StringNull()
	}
	if v, ok := obj["country"]; ok {
		_ = v
		if v != "" {
			m.Country = types.StringValue(v)
		} else {
			m.Country = types.StringNull()
		}
	} else {
		m.Country = types.StringNull()
	}
	if v, ok := obj["datapath"]; ok {
		_ = v
		if v != "" {
			m.Datapath = types.StringValue(v)
		} else {
			m.Datapath = types.StringNull()
		}
	} else {
		m.Datapath = types.StringNull()
	}
	if v, ok := obj["deprioritize-unii-3-4"]; ok {
		_ = v
		if v != "" {
			m.DeprioritizeUnii34 = types.StringValue(v)
		} else {
			m.DeprioritizeUnii34 = types.StringNull()
		}
	} else {
		m.DeprioritizeUnii34 = types.StringNull()
	}
	if v, ok := obj["dgaf"]; ok {
		_ = v
		if v != "" {
			m.Dgaf = types.StringValue(v)
		} else {
			m.Dgaf = types.StringNull()
		}
	} else {
		m.Dgaf = types.StringNull()
	}
	if v, ok := obj["security.dh-groups"]; ok {
		_ = v
		if v != "" {
			m.DhGroups = types.StringValue(v)
		} else {
			m.DhGroups = types.StringNull()
		}
	} else {
		m.DhGroups = types.StringNull()
	}
	if v, ok := obj["security.disable-pmkid"]; ok {
		_ = v
		if v != "" {
			m.DisablePmkid = types.StringValue(v)
		} else {
			m.DisablePmkid = types.StringNull()
		}
	} else {
		m.DisablePmkid = types.StringNull()
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
	if v, ok := obj["distance"]; ok {
		_ = v
		if v != "" {
			m.Distance = types.StringValue(v)
		} else {
			m.Distance = types.StringNull()
		}
	} else {
		m.Distance = types.StringNull()
	}
	if v, ok := obj["domain-names"]; ok {
		_ = v
		if v != "" {
			m.DomainNames = types.StringValue(v)
		} else {
			m.DomainNames = types.StringNull()
		}
	} else {
		m.DomainNames = types.StringNull()
	}
	if v, ok := obj["dtim-period"]; ok {
		_ = v
		if v != "" {
			m.DtimPeriod = types.StringValue(v)
		} else {
			m.DtimPeriod = types.StringNull()
		}
	} else {
		m.DtimPeriod = types.StringNull()
	}
	if v, ok := obj["security.eap-accounting"]; ok {
		_ = v
		if v != "" {
			m.EAPAccounting = types.StringValue(v)
		} else {
			m.EAPAccounting = types.StringNull()
		}
	} else {
		m.EAPAccounting = types.StringNull()
	}
	if v, ok := obj["security.eap-anonymous-identity"]; ok {
		_ = v
		if v != "" {
			m.EAPAnonymousIdentity = types.StringValue(v)
		} else {
			m.EAPAnonymousIdentity = types.StringNull()
		}
	} else {
		m.EAPAnonymousIdentity = types.StringNull()
	}
	if v, ok := obj["security.eap-certificate-mode"]; ok {
		_ = v
		if v != "" {
			m.EAPCertificateMode = types.StringValue(v)
		} else {
			m.EAPCertificateMode = types.StringNull()
		}
	} else {
		m.EAPCertificateMode = types.StringNull()
	}
	if v, ok := obj["security.eap-methods"]; ok {
		_ = v
		if v != "" {
			m.EAPMethods = types.StringValue(v)
		} else {
			m.EAPMethods = types.StringNull()
		}
	} else {
		m.EAPMethods = types.StringNull()
	}
	if v, ok := obj["security.eap-password"]; ok {
		_ = v
		if v != "" {
			m.EAPPassword = types.StringValue(v)
		} else {
			m.EAPPassword = types.StringNull()
		}
	} else {
		m.EAPPassword = types.StringNull()
	}
	if v, ok := obj["eap-tls-certificate"]; ok {
		_ = v
		if v != "" {
			m.EAPTLSCertificate = types.StringValue(v)
		} else {
			m.EAPTLSCertificate = types.StringNull()
		}
	} else {
		m.EAPTLSCertificate = types.StringNull()
	}
	if v, ok := obj["security.eap-username"]; ok {
		_ = v
		if v != "" {
			m.EAPUsername = types.StringValue(v)
		} else {
			m.EAPUsername = types.StringNull()
		}
	} else {
		m.EAPUsername = types.StringNull()
	}
	if v, ok := obj["esr"]; ok {
		_ = v
		if v != "" {
			m.Esr = types.StringValue(v)
		} else {
			m.Esr = types.StringNull()
		}
	} else {
		m.Esr = types.StringNull()
	}
	if v, ok := obj["channel.frequency"]; ok {
		_ = v
		if v != "" {
			m.Frequency = types.StringValue(v)
		} else {
			m.Frequency = types.StringNull()
		}
	} else {
		m.Frequency = types.StringNull()
	}
	if v, ok := obj["ft-enabled"]; ok {
		_ = v
		if v != "" {
			m.FtEnabled = types.StringValue(v)
		} else {
			m.FtEnabled = types.StringNull()
		}
	} else {
		m.FtEnabled = types.StringNull()
	}
	if v, ok := obj["ft-mobility-domain"]; ok {
		_ = v
		if v != "" {
			m.FtMobilityDomain = types.StringValue(v)
		} else {
			m.FtMobilityDomain = types.StringNull()
		}
	} else {
		m.FtMobilityDomain = types.StringNull()
	}
	if v, ok := obj["ft-nas-identifier"]; ok {
		_ = v
		if v != "" {
			m.FtNasIdentifier = types.StringValue(v)
		} else {
			m.FtNasIdentifier = types.StringNull()
		}
	} else {
		m.FtNasIdentifier = types.StringNull()
	}
	if v, ok := obj["security.ft-over-ds"]; ok {
		_ = v
		if v != "" {
			m.FtOverDs = types.StringValue(v)
		} else {
			m.FtOverDs = types.StringNull()
		}
	} else {
		m.FtOverDs = types.StringNull()
	}
	if v, ok := obj["ft-preserve-vlan-id"]; ok {
		_ = v
		if v != "" {
			m.FtPreserveVLANID = types.StringValue(v)
		} else {
			m.FtPreserveVLANID = types.StringNull()
		}
	} else {
		m.FtPreserveVLANID = types.StringNull()
	}
	if v, ok := obj["ft-r0-key-lifetime"]; ok {
		_ = v
		if v != "" {
			m.FtR0KeyLifetime = types.StringValue(v)
		} else {
			m.FtR0KeyLifetime = types.StringNull()
		}
	} else {
		m.FtR0KeyLifetime = types.StringNull()
	}
	if v, ok := obj["ft-reassoc-deadline"]; ok {
		_ = v
		if v != "" {
			m.FtReassocDeadline = types.StringValue(v)
		} else {
			m.FtReassocDeadline = types.StringNull()
		}
	} else {
		m.FtReassocDeadline = types.StringNull()
	}
	if v, ok := obj["security.group-encryption"]; ok {
		_ = v
		if v != "" {
			m.GroupEncryption = types.StringValue(v)
		} else {
			m.GroupEncryption = types.StringNull()
		}
	} else {
		m.GroupEncryption = types.StringNull()
	}
	if v, ok := obj["security.group-key-update"]; ok {
		_ = v
		if v != "" {
			m.GroupKeyUpdate = types.StringValue(v)
		} else {
			m.GroupKeyUpdate = types.StringNull()
		}
	} else {
		m.GroupKeyUpdate = types.StringNull()
	}
	if v, ok := obj["hessid"]; ok {
		_ = v
		if v != "" {
			m.Hessid = types.StringValue(v)
		} else {
			m.Hessid = types.StringNull()
		}
	} else {
		m.Hessid = types.StringNull()
	}
	if v, ok := obj["hide-ssid"]; ok {
		_ = v
		if v != "" {
			m.HideSsid = types.StringValue(v)
		} else {
			m.HideSsid = types.StringNull()
		}
	} else {
		m.HideSsid = types.StringNull()
	}
	if v, ok := obj["hotspot-2-0"]; ok {
		_ = v
		if v != "" {
			m.Hotspot20 = types.StringValue(v)
		} else {
			m.Hotspot20 = types.StringNull()
		}
	} else {
		m.Hotspot20 = types.StringNull()
	}
	if v, ok := obj["hw-protection-mode"]; ok {
		_ = v
		if v != "" {
			m.HwProtectionMode = types.StringValue(v)
		} else {
			m.HwProtectionMode = types.StringNull()
		}
	} else {
		m.HwProtectionMode = types.StringNull()
	}
	if v, ok := obj["installation"]; ok {
		_ = v
		if v != "" {
			m.Installation = types.StringValue(v)
		} else {
			m.Installation = types.StringNull()
		}
	} else {
		m.Installation = types.StringNull()
	}
	if v, ok := obj["datapath.interface-list"]; ok {
		_ = v
		if v != "" {
			m.InterfaceList = types.StringValue(v)
		} else {
			m.InterfaceList = types.StringNull()
		}
	} else {
		m.InterfaceList = types.StringNull()
	}
	if v, ok := obj["aaa.interim-update"]; ok {
		_ = v
		if v != "" {
			m.InterimUpdate = types.StringValue(v)
		} else {
			m.InterimUpdate = types.StringNull()
		}
	} else {
		m.InterimUpdate = types.StringNull()
	}
	if v, ok := obj["internet"]; ok {
		_ = v
		if v != "" {
			m.Internet = types.StringValue(v)
		} else {
			m.Internet = types.StringNull()
		}
	} else {
		m.Internet = types.StringNull()
	}
	if v, ok := obj["interworking"]; ok {
		_ = v
		if v != "" {
			m.Interworking = types.StringValue(v)
		} else {
			m.Interworking = types.StringNull()
		}
	} else {
		m.Interworking = types.StringNull()
	}
	if v, ok := obj["ipv4-availability"]; ok {
		_ = v
		if v != "" {
			m.Ipv4Availability = types.StringValue(v)
		} else {
			m.Ipv4Availability = types.StringNull()
		}
	} else {
		m.Ipv4Availability = types.StringNull()
	}
	if v, ok := obj["ipv6-availability"]; ok {
		_ = v
		if v != "" {
			m.IPV6Availability = types.StringValue(v)
		} else {
			m.IPV6Availability = types.StringNull()
		}
	} else {
		m.IPV6Availability = types.StringNull()
	}
	if v, ok := obj["aaa.mac-caching"]; ok {
		_ = v
		if v != "" {
			m.MACCaching = types.StringValue(v)
		} else {
			m.MACCaching = types.StringNull()
		}
	} else {
		m.MACCaching = types.StringNull()
	}
	if v, ok := obj["management-encryption"]; ok {
		_ = v
		if v != "" {
			m.ManagementEncryption = types.StringValue(v)
		} else {
			m.ManagementEncryption = types.StringNull()
		}
	} else {
		m.ManagementEncryption = types.StringNull()
	}
	if v, ok := obj["security.management-protection"]; ok {
		_ = v
		if v != "" {
			m.ManagementProtection = types.StringValue(v)
		} else {
			m.ManagementProtection = types.StringNull()
		}
	} else {
		m.ManagementProtection = types.StringNull()
	}
	if v, ok := obj["manager"]; ok {
		_ = v
		if v != "" {
			m.Manager = types.StringValue(v)
		} else {
			m.Manager = types.StringNull()
		}
	} else {
		m.Manager = types.StringNull()
	}
	if v, ok := obj["max-clients"]; ok {
		_ = v
		if v != "" {
			m.MaxClients = types.StringValue(v)
		} else {
			m.MaxClients = types.StringNull()
		}
	} else {
		m.MaxClients = types.StringNull()
	}
	if v, ok := obj["max-tx-power"]; ok {
		_ = v
		if v != "" {
			m.MaxTxPower = types.StringValue(v)
		} else {
			m.MaxTxPower = types.StringNull()
		}
	} else {
		m.MaxTxPower = types.StringNull()
	}
	if v, ok := obj["mode"]; ok {
		_ = v
		if v != "" {
			m.Mode = types.StringValue(v)
		} else {
			m.Mode = types.StringNull()
		}
	} else {
		m.Mode = types.StringNull()
	}
	if v, ok := obj["multi-passphrase-group"]; ok {
		_ = v
		if v != "" {
			m.MultiPassphraseGroup = types.StringValue(v)
		} else {
			m.MultiPassphraseGroup = types.StringNull()
		}
	} else {
		m.MultiPassphraseGroup = types.StringNull()
	}
	if v, ok := obj["multicast-enhance"]; ok {
		_ = v
		if v != "" {
			m.MulticastEnhance = types.StringValue(v)
		} else {
			m.MulticastEnhance = types.StringNull()
		}
	} else {
		m.MulticastEnhance = types.StringNull()
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
	if v, ok := obj["aaa.nas-identifier"]; ok {
		_ = v
		if v != "" {
			m.NasIdentifier = types.StringValue(v)
		} else {
			m.NasIdentifier = types.StringNull()
		}
	} else {
		m.NasIdentifier = types.StringNull()
	}
	if v, ok := obj["steering.neighbor-group"]; ok {
		_ = v
		if v != "" {
			m.NeighborGroup = types.StringValue(v)
		} else {
			m.NeighborGroup = types.StringNull()
		}
	} else {
		m.NeighborGroup = types.StringNull()
	}
	if v, ok := obj["network-type"]; ok {
		_ = v
		if v != "" {
			m.NetworkType = types.StringValue(v)
		} else {
			m.NetworkType = types.StringNull()
		}
	} else {
		m.NetworkType = types.StringNull()
	}
	if v, ok := obj["open-flow-switch"]; ok {
		_ = v
		if v != "" {
			m.OpenFlowSwitch = types.StringValue(v)
		} else {
			m.OpenFlowSwitch = types.StringNull()
		}
	} else {
		m.OpenFlowSwitch = types.StringNull()
	}
	if v, ok := obj["openflow"]; ok {
		_ = v
		if v != "" {
			m.Openflow = types.StringValue(v)
		} else {
			m.Openflow = types.StringNull()
		}
	} else {
		m.Openflow = types.StringNull()
	}
	if v, ok := obj["operational-classes"]; ok {
		_ = v
		if v != "" {
			m.OperationalClasses = types.StringValue(v)
		} else {
			m.OperationalClasses = types.StringNull()
		}
	} else {
		m.OperationalClasses = types.StringNull()
	}
	if v, ok := obj["operator-names"]; ok {
		_ = v
		if v != "" {
			m.OperatorNames = types.StringValue(v)
		} else {
			m.OperatorNames = types.StringNull()
		}
	} else {
		m.OperatorNames = types.StringNull()
	}
	if v, ok := obj["security.owe-transition-interface"]; ok {
		_ = v
		if v != "" {
			m.OweTransitionInterface = types.StringValue(v)
		} else {
			m.OweTransitionInterface = types.StringNull()
		}
	} else {
		m.OweTransitionInterface = types.StringNull()
	}
	if v, ok := obj["security.passphrase"]; ok {
		_ = v
		if v != "" {
			m.Passphrase = types.StringValue(v)
		} else {
			m.Passphrase = types.StringNull()
		}
	} else {
		m.Passphrase = types.StringNull()
	}
	if v, ok := obj["aaa.password-format"]; ok {
		_ = v
		if v != "" {
			m.PasswordFormat = types.StringValue(v)
		} else {
			m.PasswordFormat = types.StringNull()
		}
	} else {
		m.PasswordFormat = types.StringNull()
	}
	if v, ok := obj["qo-s-classifier"]; ok {
		_ = v
		if v != "" {
			m.QoSClassifier = types.StringValue(v)
		} else {
			m.QoSClassifier = types.StringNull()
		}
	} else {
		m.QoSClassifier = types.StringNull()
	}
	if v, ok := obj["realms"]; ok {
		_ = v
		if v != "" {
			m.Realms = types.StringValue(v)
		} else {
			m.Realms = types.StringNull()
		}
	} else {
		m.Realms = types.StringNull()
	}
	if v, ok := obj["realms-raw"]; ok {
		_ = v
		if v != "" {
			m.RealmsRaw = types.StringValue(v)
		} else {
			m.RealmsRaw = types.StringNull()
		}
	} else {
		m.RealmsRaw = types.StringNull()
	}
	if v, ok := obj["channel.reselect-interval"]; ok {
		_ = v
		if v != "" {
			m.ReselectInterval = types.StringValue(v)
		} else {
			m.ReselectInterval = types.StringNull()
		}
	} else {
		m.ReselectInterval = types.StringNull()
	}
	if v, ok := obj["channel.reselect-time"]; ok {
		_ = v
		if v != "" {
			m.ReselectTime = types.StringValue(v)
		} else {
			m.ReselectTime = types.StringNull()
		}
	} else {
		m.ReselectTime = types.StringNull()
	}
	if v, ok := obj["roaming-ois"]; ok {
		_ = v
		if v != "" {
			m.RoamingOis = types.StringValue(v)
		} else {
			m.RoamingOis = types.StringNull()
		}
	} else {
		m.RoamingOis = types.StringNull()
	}
	if v, ok := obj["steering.rrm"]; ok {
		_ = v
		if v != "" {
			m.Rrm = types.StringValue(v)
		} else {
			m.Rrm = types.StringNull()
		}
	} else {
		m.Rrm = types.StringNull()
	}
	if v, ok := obj["sae-anti-clogging-threshold"]; ok {
		_ = v
		if v != "" {
			m.SaeAntiCloggingThreshold = types.StringValue(v)
		} else {
			m.SaeAntiCloggingThreshold = types.StringNull()
		}
	} else {
		m.SaeAntiCloggingThreshold = types.StringNull()
	}
	if v, ok := obj["security.sae-max-failure-rate"]; ok {
		_ = v
		if v != "" {
			m.SaeMaxFailureRate = types.StringValue(v)
		} else {
			m.SaeMaxFailureRate = types.StringNull()
		}
	} else {
		m.SaeMaxFailureRate = types.StringNull()
	}
	if v, ok := obj["security.sae-pwe"]; ok {
		_ = v
		if v != "" {
			m.SaePwe = types.StringValue(v)
		} else {
			m.SaePwe = types.StringNull()
		}
	} else {
		m.SaePwe = types.StringNull()
	}
	if v, ok := obj["channel.secondary-frequency"]; ok {
		_ = v
		if v != "" {
			m.SecondaryFrequency = types.StringValue(v)
		} else {
			m.SecondaryFrequency = types.StringNull()
		}
	} else {
		m.SecondaryFrequency = types.StringNull()
	}
	if v, ok := obj["security"]; ok {
		_ = v
		if v != "" {
			m.Security = types.StringValue(v)
		} else {
			m.Security = types.StringNull()
		}
	} else {
		m.Security = types.StringNull()
	}
	if v, ok := obj["channel.skip-dfs-channels"]; ok {
		_ = v
		if v != "" {
			m.SkipDfsChannels = types.StringValue(v)
		} else {
			m.SkipDfsChannels = types.StringNull()
		}
	} else {
		m.SkipDfsChannels = types.StringNull()
	}
	if v, ok := obj["ssid"]; ok {
		_ = v
		if v != "" {
			m.Ssid = types.StringValue(v)
		} else {
			m.Ssid = types.StringNull()
		}
	} else {
		m.Ssid = types.StringNull()
	}
	if v, ok := obj["station-roaming"]; ok {
		_ = v
		if v != "" {
			m.StationRoaming = types.StringValue(v)
		} else {
			m.StationRoaming = types.StringNull()
		}
	} else {
		m.StationRoaming = types.StringNull()
	}
	if v, ok := obj["steering"]; ok {
		_ = v
		if v != "" {
			m.Steering = types.StringValue(v)
		} else {
			m.Steering = types.StringNull()
		}
	} else {
		m.Steering = types.StringNull()
	}
	if v, ok := obj["datapath.traffic-processing"]; ok {
		_ = v
		if v != "" {
			m.TrafficProcessing = types.StringValue(v)
		} else {
			m.TrafficProcessing = types.StringNull()
		}
	} else {
		m.TrafficProcessing = types.StringNull()
	}
	if v, ok := obj["transition-request-count"]; ok {
		_ = v
		if v != "" {
			m.TransitionRequestCount = types.StringValue(v)
		} else {
			m.TransitionRequestCount = types.StringNull()
		}
	} else {
		m.TransitionRequestCount = types.StringNull()
	}
	if v, ok := obj["transition-threshold"]; ok {
		_ = v
		if v != "" {
			m.TransitionThreshold = types.StringValue(v)
		} else {
			m.TransitionThreshold = types.StringNull()
		}
	} else {
		m.TransitionThreshold = types.StringNull()
	}
	if v, ok := obj["transition-threshold-period"]; ok {
		_ = v
		if v != "" {
			m.TransitionThresholdPeriod = types.StringValue(v)
		} else {
			m.TransitionThresholdPeriod = types.StringNull()
		}
	} else {
		m.TransitionThresholdPeriod = types.StringNull()
	}
	if v, ok := obj["transition-threshold-time"]; ok {
		_ = v
		if v != "" {
			m.TransitionThresholdTime = types.StringValue(v)
		} else {
			m.TransitionThresholdTime = types.StringNull()
		}
	} else {
		m.TransitionThresholdTime = types.StringNull()
	}
	if v, ok := obj["transition-time"]; ok {
		_ = v
		if v != "" {
			m.TransitionTime = types.StringValue(v)
		} else {
			m.TransitionTime = types.StringNull()
		}
	} else {
		m.TransitionTime = types.StringNull()
	}
	if v, ok := obj["tx-chains"]; ok {
		_ = v
		if v != "" {
			m.TxChains = types.StringValue(v)
		} else {
			m.TxChains = types.StringNull()
		}
	} else {
		m.TxChains = types.StringNull()
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
	if v, ok := obj["types"]; ok {
		_ = v
		if v != "" {
			m.Types = types.StringValue(v)
		} else {
			m.Types = types.StringNull()
		}
	} else {
		m.Types = types.StringNull()
	}
	if v, ok := obj["uesa"]; ok {
		_ = v
		if v != "" {
			m.Uesa = types.StringValue(v)
		} else {
			m.Uesa = types.StringNull()
		}
	} else {
		m.Uesa = types.StringNull()
	}
	if v, ok := obj["aaa.username-format"]; ok {
		_ = v
		if v != "" {
			m.UsernameFormat = types.StringValue(v)
		} else {
			m.UsernameFormat = types.StringNull()
		}
	} else {
		m.UsernameFormat = types.StringNull()
	}
	if v, ok := obj["venue"]; ok {
		_ = v
		if v != "" {
			m.Venue = types.StringValue(v)
		} else {
			m.Venue = types.StringNull()
		}
	} else {
		m.Venue = types.StringNull()
	}
	if v, ok := obj["venue-names"]; ok {
		_ = v
		if v != "" {
			m.VenueNames = types.StringValue(v)
		} else {
			m.VenueNames = types.StringNull()
		}
	} else {
		m.VenueNames = types.StringNull()
	}
	if v, ok := obj["datapath.vlan-id"]; ok {
		_ = v
		if v != "" {
			m.VLANID = types.StringValue(v)
		} else {
			m.VLANID = types.StringNull()
		}
	} else {
		m.VLANID = types.StringNull()
	}
	if v, ok := obj["wan-at-capacity"]; ok {
		_ = v
		if v != "" {
			m.WanAtCapacity = types.StringValue(v)
		} else {
			m.WanAtCapacity = types.StringNull()
		}
	} else {
		m.WanAtCapacity = types.StringNull()
	}
	if v, ok := obj["wan-downlink"]; ok {
		_ = v
		if v != "" {
			m.WanDownlink = types.StringValue(v)
		} else {
			m.WanDownlink = types.StringNull()
		}
	} else {
		m.WanDownlink = types.StringNull()
	}
	if v, ok := obj["wan-downlink-load"]; ok {
		_ = v
		if v != "" {
			m.WanDownlinkLoad = types.StringValue(v)
		} else {
			m.WanDownlinkLoad = types.StringNull()
		}
	} else {
		m.WanDownlinkLoad = types.StringNull()
	}
	if v, ok := obj["wan-measurement-duration"]; ok {
		_ = v
		if v != "" {
			m.WanMeasurementDuration = types.StringValue(v)
		} else {
			m.WanMeasurementDuration = types.StringNull()
		}
	} else {
		m.WanMeasurementDuration = types.StringNull()
	}
	if v, ok := obj["wan-status"]; ok {
		_ = v
		if v != "" {
			m.WanStatus = types.StringValue(v)
		} else {
			m.WanStatus = types.StringNull()
		}
	} else {
		m.WanStatus = types.StringNull()
	}
	if v, ok := obj["wan-symmetric"]; ok {
		_ = v
		if v != "" {
			m.WanSymmetric = types.StringValue(v)
		} else {
			m.WanSymmetric = types.StringNull()
		}
	} else {
		m.WanSymmetric = types.StringNull()
	}
	if v, ok := obj["wan-uplink"]; ok {
		_ = v
		if v != "" {
			m.WanUplink = types.StringValue(v)
		} else {
			m.WanUplink = types.StringNull()
		}
	} else {
		m.WanUplink = types.StringNull()
	}
	if v, ok := obj["wan-uplink-load"]; ok {
		_ = v
		if v != "" {
			m.WanUplinkLoad = types.StringValue(v)
		} else {
			m.WanUplinkLoad = types.StringNull()
		}
	} else {
		m.WanUplinkLoad = types.StringNull()
	}
	if v, ok := obj["steering.wnm"]; ok {
		_ = v
		if v != "" {
			m.Wnm = types.StringValue(v)
		} else {
			m.Wnm = types.StringNull()
		}
	} else {
		m.Wnm = types.StringNull()
	}
	if v, ok := obj["security.wps"]; ok {
		_ = v
		if v != "" {
			m.Wps = types.StringValue(v)
		} else {
			m.Wps = types.StringNull()
		}
	} else {
		m.Wps = types.StringNull()
	}
}
