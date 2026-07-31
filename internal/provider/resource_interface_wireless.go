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
	_ resource.Resource                = &InterfaceWirelessResource{}
	_ resource.ResourceWithImportState = &InterfaceWirelessResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceWirelessResource struct {
	reg *client.Registry
}

type InterfaceWirelessModel struct {
	ID                       types.String `tfsdk:"id"`
	VlanMode                 types.String `tfsdk:"vlan_mode"`
	VlanId                   types.String `tfsdk:"vlan_id"`
	SecondaryFrequency       types.String `tfsdk:"secondary_frequency"`
	Nv2SyncSecret            types.String `tfsdk:"nv2_sync_secret"`
	Nv2Mode                  types.String `tfsdk:"nv2_mode"`
	Nv2DownlinkRatio         types.String `tfsdk:"nv2_downlink_ratio"`
	DfsTestMode              types.String `tfsdk:"dfs_test_mode"`
	AdaptiveNoiseImmunity    types.String `tfsdk:"adaptive_noise_immunity"`
	AllowSharedkey           types.String `tfsdk:"allow_sharedkey"`
	AmpduPriorities          types.String `tfsdk:"ampdu_priorities"`
	AmsduLimit               types.String `tfsdk:"amsdu_limit"`
	AmsduThreshold           types.String `tfsdk:"amsdu_threshold"`
	AntennaGain              types.String `tfsdk:"antenna_gain"`
	AntennaMode              types.String `tfsdk:"antenna_mode"`
	Area                     types.String `tfsdk:"area"`
	ARP                      types.String `tfsdk:"arp"`
	ARPTimeout               types.String `tfsdk:"arp_timeout"`
	Band                     types.String `tfsdk:"band"`
	BasicRatesB              types.String `tfsdk:"basic_rates_b"`
	BridgeMode               types.String `tfsdk:"bridge_mode"`
	BurstTime                types.String `tfsdk:"burst_time"`
	ChannelWidth             types.String `tfsdk:"channel_width"`
	Comment                  types.String `tfsdk:"comment"`
	Compression              types.String `tfsdk:"compression"`
	Country                  types.String `tfsdk:"country"`
	DefaultApTxLimit         types.String `tfsdk:"default_ap_tx_limit"`
	DefaultAuthentication    types.String `tfsdk:"default_authentication"`
	DefaultClientTxLimit     types.String `tfsdk:"default_client_tx_limit"`
	DefaultForwarding        types.String `tfsdk:"default_forwarding"`
	DisableRunningCheck      types.String `tfsdk:"disable_running_check"`
	Disabled                 types.Bool   `tfsdk:"disabled"`
	DisconnectTimeout        types.String `tfsdk:"disconnect_timeout"`
	Distance                 types.String `tfsdk:"distance"`
	FrameLifetime            types.String `tfsdk:"frame_lifetime"`
	Frequency                types.String `tfsdk:"frequency"`
	FrequencyMode            types.String `tfsdk:"frequency_mode"`
	FrequencyOffset          types.String `tfsdk:"frequency_offset"`
	GuardInterval            types.String `tfsdk:"guard_interval"`
	HideSsid                 types.String `tfsdk:"hide_ssid"`
	HtBasicMcs               types.String `tfsdk:"ht_basic_mcs"`
	HtSupportedMcs           types.String `tfsdk:"ht_supported_mcs"`
	HwFragmentationThreshold types.String `tfsdk:"hw_fragmentation_threshold"`
	HwProtectionMode         types.String `tfsdk:"hw_protection_mode"`
	HwProtectionThreshold    types.String `tfsdk:"hw_protection_threshold"`
	HwRetries                types.String `tfsdk:"hw_retries"`
	Installation             types.String `tfsdk:"installation"`
	InterworkingProfile      types.String `tfsdk:"interworking_profile"`
	KeepaliveFrames          types.String `tfsdk:"keepalive_frames"`
	L2mtu                    types.String `tfsdk:"l2mtu"`
	MACAddress               types.String `tfsdk:"mac_address"`
	MasterInterface          types.String `tfsdk:"master_interface"`
	MaxStationCount          types.String `tfsdk:"max_station_count"`
	Mode                     types.String `tfsdk:"mode"`
	MTU                      types.String `tfsdk:"mtu"`
	MulticastBuffering       types.String `tfsdk:"multicast_buffering"`
	MulticastHelper          types.String `tfsdk:"multicast_helper"`
	Name                     types.String `tfsdk:"name"`
	NoiseFloorThreshold      types.String `tfsdk:"noise_floor_threshold"`
	Nv2CellRADIUS            types.String `tfsdk:"nv2_cell_radius"`
	Nv2NoiseFloorOffset      types.String `tfsdk:"nv2_noise_floor_offset"`
	Nv2PresharedKey          types.String `tfsdk:"nv2_preshared_key"`
	Nv2Qos                   types.String `tfsdk:"nv2_qos"`
	Nv2QueueCount            types.String `tfsdk:"nv2_queue_count"`
	Nv2Security              types.String `tfsdk:"nv2_security"`
	OnFailRetryTime          types.String `tfsdk:"on_fail_retry_time"`
	PreambleMode             types.String `tfsdk:"preamble_mode"`
	PrismCardtype            types.String `tfsdk:"prism_cardtype"`
	RadioName                types.String `tfsdk:"radio_name"`
	RateSelection            types.String `tfsdk:"rate_selection"`
	RateSet                  types.String `tfsdk:"rate_set"`
	RxChains                 types.String `tfsdk:"rx_chains"`
	RxHtChainNames           types.String `tfsdk:"rx_ht_chain_names"`
	RxHtChains               types.String `tfsdk:"rx_ht_chains"`
	ScanList                 types.String `tfsdk:"scan_list"`
	SecurityProfile          types.String `tfsdk:"security_profile"`
	SkipDfsChannels          types.String `tfsdk:"skip_dfs_channels"`
	Ssid                     types.String `tfsdk:"ssid"`
	StationBridgeCloneMAC    types.String `tfsdk:"station_bridge_clone_mac"`
	StationRoaming           types.String `tfsdk:"station_roaming"`
	SupportedRatesB          types.String `tfsdk:"supported_rates_b"`
	TdmaPeriodSize           types.String `tfsdk:"tdma_period_size"`
	TxChains                 types.String `tfsdk:"tx_chains"`
	TxHtChainNames           types.String `tfsdk:"tx_ht_chain_names"`
	TxHtChains               types.String `tfsdk:"tx_ht_chains"`
	TxPower                  types.String `tfsdk:"tx_power"`
	TxPowerMode              types.String `tfsdk:"tx_power_mode"`
	UpdateStatsInterval      types.String `tfsdk:"update_stats_interval"`
	VhtBasicMcs              types.String `tfsdk:"vht_basic_mcs"`
	VhtSupportedMcs          types.String `tfsdk:"vht_supported_mcs"`
	WdsCostRange             types.String `tfsdk:"wds_cost_range"`
	WdsDefaultBridge         types.String `tfsdk:"wds_default_bridge"`
	WdsDefaultCost           types.String `tfsdk:"wds_default_cost"`
	WdsIgnoreSsid            types.String `tfsdk:"wds_ignore_ssid"`
	WdsMode                  types.String `tfsdk:"wds_mode"`
	WirelessProtocol         types.String `tfsdk:"wireless_protocol"`
	WmmSupport               types.String `tfsdk:"wmm_support"`
	WpsMode                  types.String `tfsdk:"wps_mode"`
	Router                   types.String `tfsdk:"router"`
}

func NewInterfaceWirelessResource() resource.Resource { return &InterfaceWirelessResource{} }

func (r *InterfaceWirelessResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_wireless"
}

func (r *InterfaceWirelessResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceWirelessResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "802.11abgn interface needs a master-interface from the physical wireless hardware. Skipped on CHR.",
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
			"secondary_frequency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `secondary-frequency`.",
			},
			"nv2_sync_secret": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "RouterOS `nv2-sync-secret`.",
			},
			"nv2_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `nv2-mode`.",
			},
			"nv2_downlink_ratio": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `nv2-downlink-ratio`.",
			},
			"dfs_test_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dfs-test-mode`.",
			},
			"adaptive_noise_immunity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This property is only effective for cards based on Atheros chipset.",
			},
			"allow_sharedkey": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allow WEP Shared Key clients to connect. Note that no authentication is done for these clients (WEP Shared keys are not compared to anything) - they are just accepted at once (if access list allows that)",
			},
			"ampdu_priorities": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Frame priorities for which AMPDU sending (aggregating frames and sending using block acknowledgment) should get negotiated and used. Using AMPDUs will increase throughput, but may increase latency, therefore, may not be desirable for real-time traffic (voice, video). Due to this, by default AMPDUs are enabled only for best-effort traffic.",
			},
			"amsdu_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Max AMSDU that device is allowed to prepare when negotiated. AMSDU aggregation may significantly increase throughput especially for small frames, but may increase latency in case of packet loss due to retransmission of aggregated frame. Sending and receiving AMSDUs will also increase CPU usage.",
			},
			"amsdu_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Max frame size to allow including in AMSDU.",
			},
			"antenna_gain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Antenna gain in dBi, used to calculate maximum transmit power according to country regulations.",
			},
			"antenna_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Select antenna to use for transmitting and for receiving ant-a - use only 'a' antenna ant-b - use only 'b' antenna txa-rxb - use antenna 'a' for transmitting, antenna 'b' for receiving rxa-txb - use antenna 'b' for transmitting, antenna 'a' for receiving",
			},
			"area": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Identifies group of wireless networks. This value is announced by AP, and can be matched in connect-list by area-prefix . This is a proprietary extension.",
			},
			"arp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Read more >>",
			},
			"arp_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ARP timeout is time how long ARP record is kept in ARP table after no packets are received from IP. Value auto equals to the value of arp-timeout in /ip settings , default is 30s",
			},
			"band": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Defines set of used data rates, channel frequencies and widths.",
			},
			"basic_rates_b": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "List of basic rates, used for 2.4ghz-b, 2.4ghz-b/g and 2.4ghz-onlyg bands. Client will connect to AP only if it supports all basic rates announced by the AP. AP will establish WDS link only if it supports all basic rates of the other AP. This property has effect only in AP modes, and when value of rate-set is configured.",
			},
			"bridge_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allows to use station-bridge mode. Read more >>",
			},
			"burst_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time in microseconds which will be used to send data without stopping. Note that no other wireless cards in that network will be able to transmit data during burst-time microseconds. This setting is available only for AR5000, AR5001X, and AR5001X+ chipset based cards.",
			},
			"channel_width": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Use of extension channels (e.g. Ce, eC etc) allows additional 20MHz extension channels and if it should be located below or above the control (main) channel. Extension channel allows 802.11n devices to use up to 40MHz (802.11ac up to 160MHz) of spectrum in total thus increasing max throughput. Channel widths with XX and XXXX extensions automatically scan for a less crowded control channel frequency based on the number of concurrent devices running in every frequency and chooses the “C” - Control channel frequency automatically.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"compression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Setting this property to yes will allow the use of the hardware compression. Wireless interface must have support for hardware compression. Connections with devices that do not use compression will still work.",
			},
			"country": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Limits available bands, frequencies and maximum transmit power for each frequency. Also specifies default value of scan-list . Value no_country_set is an FCC compliant set of channels.",
			},
			"default_ap_tx_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the value of ap-tx-limit for clients that do not match any entry in the access-list . 0 means no limit.",
			},
			"default_authentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "For AP mode, this is the value of authentication for clients that do not match any entry in the access-list . For station mode, this is the value of connect for APs that do not match any entry in the connect-list",
			},
			"default_client_tx_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the value of client-tx-limit for clients that do not match any entry in the access-list . 0 means no limit",
			},
			"default_forwarding": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This is the value of forwarding for clients that do not match any entry in the access-list",
			},
			"disable_running_check": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When set to yes interface will always have running flag. If value is set to no' , the router determines whether the card is up and running - for AP one or more clients have to be registered to it, for station, it should be connected to an AP.",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"disconnect_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This interval is measured from third sending failure on the lowest data rate. At this point 3 * ( hw-retries + 1) frame transmits on the lowest data rate had failed. During disconnect-timeout packet transmission will be retried with on-fail-retry-time interval. If no frame can be transmitted successfully during disconnect-timeout , the connection is closed, and this event is logged as \"extensive data loss\". Successful frame transmission resets this timer.",
			},
			"distance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "How long to wait for confirmation of unicast frames ( ACKs ) before considering transmission unsuccessful, or in short ACK-Timeout . Distance value has these behaviors: Dynamic - causes AP to detect and use the smallest timeout that works with all connected clients. Indoor - uses the default ACK timeout value that the hardware chip manufacturer has set. Number - uses the input value in formula: ACK-timeout = (( distance * 1000) + 299) / 300 us; Acknowledgments are not used in Nstreme/NV2 protocols.",
			},
			"frame_lifetime": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Discard frames that have been queued for sending longer than frame-lifetime . By default, when value of this property is 0 , frames are discarded only after connection is closed.",
			},
			"frequency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Channel frequency value in MHz on which AP will operate. Allowed values depend on the selected band, and are restricted by country setting and wireless card capabilities. This setting has no effect if interface is in any of station modes, or in wds-slave mode, or if DFS is active. Note : If using mode \"superchannel\", any frequency supported by the card will be accepted, but on the RouterOS client, any non-standard frequency must be configured in the scan-list , otherwise it will not be scanning in non-standard range. In Winbox, scanlist frequencies are in bold , any other frequency means the clients will need scan-list configured.",
			},
			"frequency_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Three frequency modes are available: regulatory-domain - Limit available channels and maximum transmit power for each channel according to the value of country manual-txpower - Same as above, but do not limit maximum transmit power. superchannel - Conformance Testing Mode. Allow all channels supported by the card. List of available channels for each band can be seen in /interface wireless info allowed-channels . This mode allows you to test wireless channels outside the default scan-list and/or regulatory domain. This mode should only be used in controlled environments, or if you have special permission to use it in your region. Before v4.3 this was called Custom Frequency Upgrade, or Superchannel. Since RouterOS v4.3 this mode is available without special key upgrades to all installations.",
			},
			"frequency_offset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Allows to specify offset if the used wireless card operates at a different frequency than is shown in RouterOS, in case a frequency converter is used in the card. So if your card works at 4000MHz but RouterOS shows 5000MHz, set offset to 1000MHz and it will be displayed correctly. The value is in MHz and can be positive or negative.",
			},
			"guard_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether to allow use of short guard interval (refer to 802.11n MCS specification to see how this may affect throughput). \"any\" will use either short or long, depending on data rate, \"long\" will use long.",
			},
			"hide_ssid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "yes - AP does not include SSID in the beacon frames, and does not reply to probe requests that have broadcast SSID. no - AP includes SSID in the beacon frames, and replies to probe requests that have broadcast SSID. This property has an effect only in AP mode. Setting it to yes can remove this network from the list of wireless networks that are shown by some client software. Changing this setting does not improve the security of the wireless network, because SSID is included in other frames sent by the AP.",
			},
			"ht_basic_mcs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Modulation and Coding Schemes that every connecting client must support. Refer to 802.11n for MCS specification.",
			},
			"ht_supported_mcs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11n for MCS specification.",
			},
			"hw_fragmentation_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies maximum fragment size in bytes when transmitted over the wireless medium. 802.11 standard packet (MSDU in 802.11 terminologies) fragmentation allows packets to be fragmented before transmitting over a wireless medium to increase the probability of successful transmission (only fragments that did not transmit correctly are retransmitted). Note that transmission of a fragmented packet is less efficient than transmitting unfragmented packet because of protocol overhead and increased resource usage at both - transmitting and receiving party.",
			},
			"hw_protection_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Frame protection support property read more >>",
			},
			"hw_protection_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Frame protection support property read more >>",
			},
			"hw_retries": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of times sending frame is retried without considering it a transmission failure. Data-rate is decreased upon failure and the frame is sent again. Three sequential failures on the lowest supported rate suspend transmission to this destination for the duration of on-fail-retry-time . After that, the frame is sent again. The frame is being retransmitted until transmission success, or until the client is disconnected after disconnect-timeout . The frame can be discarded during this time if frame-lifetime is exceeded.",
			},
			"installation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Adjusts scan-list to use indoor, outdoor or all frequencies for the country that is set.",
			},
			"interworking_profile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"keepalive_frames": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Applies only if wireless interface is in mode= ap-bridge . If a client has not communicated for around 20 seconds, AP sends a \"keepalive-frame\". Note , disabling the feature can lead to \"ghost\" clients in registration-table.",
			},
			"l2mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"master_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of wireless interface that has virtual-ap capability. Virtual AP interface will only work if master interface is in ap-bridge , bridge , station or wds-slave mode. This property is only for virtual AP interfaces.",
			},
			"max_station_count": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of associated clients. WDS links also count toward this limit.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Selection between different station and access point (AP) modes. : station - Basic station mode. Find and connect to acceptable AP. station-wds - Same as station , but create WDS link with AP, using proprietary extension. AP configuration has to allow WDS links with this device. Note that this mode does not use entries in wds. station-pseudobridge - Same as station , but additionally perform MAC address translation of all traffic. Allows interface to be bridged. station-pseudobridge-clone - Same as station-pseudobridge , but use station-bridge-clone-mac address to connect to AP. station-bridge - Provides support for transparent protocol-independent L2 bridging on the station device. RouterOS AP accepts clients in station-bridge mode when enabled using bridge-mode parameter. In this mode, the AP maintains a forwarding table with information on which MAC addresses are reachable over which station device. Only works with RouterOS APs. With station-bridge mode, it is not possible to connect to CAPsMAN controlled CAP. The 'wireless' station-bridge mode,\u00a0 is incompatible with APs running the newer 'wifi' package and vice versa. AP modes: ap-bridge - Basic access point mode. bridge - Same as ap-bridge , but limited to one associated client. wds-slave - Same as ap-bridge , but scan for AP with the same ssid and establishes WDS link. If this link is lost or cannot be established, then continue scanning. If dfs-mode is radar-detect , then APs with enabled hide-ssid will not be found during scanning. Special modes: alignment-only - Put the interface in a continuous transmit mode that is used for aiming the remote antenna. nstreme-dual-slave - allow this interface to be used in nstreme-dual setup. MAC address translation in pseudobridge modes works by inspecting packets and building a table of corresponding IP and MAC addresses. All packets are sent to AP with the MAC address used by pseudobridge, and MAC addresses of received packets are restored from the address translation table. There is a single entry in the address translation table for all non-IP packets, hence more than one host in the bridged network cannot reliably use non-IP protocols. Note: Currently IPv6 doesn't work over Pseudobridge",
			},
			"mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"multicast_buffering": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "For a client that has power saving, buffer multicast packets until next beacon time. A client should wake up to receive a beacon, by receiving beacon it sees that there are multicast packets pending, and it should wait for multicast packets to be sent.",
			},
			"multicast_helper": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When set to full , multicast packets will be sent with a unicast destination MAC address, resolving multicast problem on the wireless link. This option should be enabled only on the access point, clients should be configured in station-bridge mode. Available starting from v5.15. disabled - disables the helper and sends multicast packets with multicast destination MAC addresses dhcp - dhcp packet mac addresses are changed to unicast mac addresses prior to sending them out full - all multicast packet mac address are changed to unicast mac addresses prior to sending them out default - default choice that currently is set to dhcp . Value can be changed in future releases.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "name of the interface",
			},
			"noise_floor_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "For advanced use only, as it can badly affect the performance of the interface. It is possible to manually set noise floor threshold value. By default, it is dynamically calculated. This property also affects received signal strength. This property is only effective on non-AC chips.",
			},
			"nv2_cell_radius": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Setting affects the size of contention time slot that AP allocates for clients to initiate connection and also size of time slots used for estimating distance to client. When setting is too small, clients that are farther away may have trouble connecting and/or disconnect with \"ranging timeout\" error. Although during normal operation the effect of this setting should be negligible, in order to maintain maximum performance, it is advised to not increase this setting if not necessary, so AP is not reserving time that is actually never used, but instead allocates it for actual data transfer. on AP: distance to farthest client in km on station: no effect",
			},
			"nv2_noise_floor_offset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nv2_preshared_key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"nv2_qos": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Sets the packet priority mechanism, firstly data from high priority queue is sent, then lower queue priority data until 0 queue priority is reached. When link is full with high priority queue data, lower priority data is not sent. Use it very carefully, setting works on AP frame-priority - manual setting that can be tuned with Mangle rules. default - default setting where small packets receive priority for best latency",
			},
			"nv2_queue_count": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nv2_security": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"on_fail_retry_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "After third sending failure on the lowest data rate, wait for specified time interval before retrying.",
			},
			"preamble_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Short preamble mode is an option of 802.11b standard that reduces per-frame overhead. On AP: long - Do not use short preamble. short - Announce short preamble capability. Do not accept connections from clients that do not have this capability. both - Announce short preamble capability. On station: long - do not use short preamble. short - do not connect to AP if it does not support short preamble. both - Use short preamble if AP supports it.",
			},
			"prism_cardtype": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify type of the installed Prism wireless card.",
			},
			"radio_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Descriptive name of the device, that is shown in registration table entries on the remote devices. This is a proprietary extension.",
			},
			"rate_selection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Starting from v5.9 default value is advanced since legacy mode was inefficient.",
			},
			"rate_set": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Two options are available: default - default basic and supported rate sets are used. Values from basic-rates and supported-rates parameters have no effect. configured - use values from basic-rates , supported-rates , basic-mcs , mcs . Read more >> .",
			},
			"rx_chains": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Which antennas to use for receive. In current MikroTik routers, both RX and TX chain must be enabled, for the chain to be enabled.",
			},
			"rx_ht_chain_names": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_ht_chains": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"scan_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The default value is all channels from selected band that are supported by card and allowed by the country and frequency-mode settings (this list can be seen in info ). For default scan list in 5ghz band channels are taken with 20MHz step, in 5ghz-turbo band - with 40MHz step, for all other bands - with 5MHz step. If scan-list is specified manually, then all matching channels are taken. (Example: scan-list = default,5200-5245,2412-2427 - This will use the default value of scan list for current band, and add to it supported frequencies from 5200-5245 or 2412-2427 range.) Since RouterOS v6.0 with Winbox or Webfig, for inputting of multiple frequencies, add each frequency or range of frequencies into separate multiple scan-lists. Using a comma to separate frequencies is no longer supported in Winbox/Webfig since v6.0. Since RouterOS v6.35 (wireless-rep) scan-list support step feature where it is possible to manually specify the scan step. Example: scan-list = 5500-5600:20 will generate such scan-list values 5500,5520,5540,5560,5580,5600",
			},
			"security_profile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of profile from security-profiles",
			},
			"skip_dfs_channels": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "These values are used to skip all DFS channels or specifically skip DFS CAC channels in range 5600-5650MHz which detection could go up to 10min.",
			},
			"ssid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSID (service set identifier) is a name that identifies wireless network.",
			},
			"station_bridge_clone_mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "This property has effect only in the station-pseudobridge-clone mode. Use this MAC address when connection to AP. If this value is 00:00:00:00:00:00 , station will initially use MAC address of the wireless interface. As soon as packet with MAC address of another device needs to be transmitted, station will reconnect to AP using that address.",
			},
			"station_roaming": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Station Roaming feature is available only for 802.11 wireless protocol and only for station modes. Read more >>",
			},
			"supported_rates_b": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "List of supported rates, used for 2ghz-b , 2ghz-b/g and 2ghz-b/g/n bands. Two devices will communicate only using rates that are supported by both devices. This property has effect only when value of rate-set is configured .",
			},
			"tdma_period_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies TDMA period in milliseconds. It could help on the longer distance links, it could slightly increase bandwidth, while latency is increased too.",
			},
			"tx_chains": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Which antennas to use for transmitting. In current MikroTik routers, both RX and TX chain must be enabled, for the chain to be enabled.",
			},
			"tx_ht_chain_names": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_ht_chains": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_power": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "For 802.11ac wireless interface it's total power but for 802.11a/b/g/n it's power per chain.",
			},
			"tx_power_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "sets up tx-power mode for wireless card default - use values stored in the card all-rates-fixed - use same transmit power for all data rates. Can damage the card if transmit power is set above rated value of the card for used rate. manual-table - define transmit power for each rate separately. Can damage the card if transmit power is set above rated value of the card for used rate. card-rates - use transmit power calculated for each rate based on value of tx-power parameter. Legacy mode only compatible with currently discontinued products.",
			},
			"update_stats_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "How often to request update of signals strength and ccq values from clients. Access to registration-table also triggers update of these values. This is proprietary extension.",
			},
			"vht_basic_mcs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Modulation and Coding Schemes that every connecting client must support. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - client must support MCS-0 to MCS-7 MCS 0-8 - client must support MCS-0 to MCS-8 MCS 0-9 - client must support MCS-0 to MCS-9",
			},
			"vht_supported_mcs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Modulation and Coding Schemes that this device advertises as supported. Refer to 802.11ac for MCS specification. You can set MCS interval for each of Spatial Stream none - will not use selected Spatial Stream MCS 0-7 - devices will advertise as supported MCS-0 to MCS-7 MCS 0-8 - devices will advertise as supported MCS-0 to MCS-8 MCS 0-9 - devices will advertise as supported MCS-0 to MCS-9",
			},
			"wds_cost_range": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bridge port cost of WDS links are automatically adjusted, depending on measured link throughput. Port cost is recalculated and adjusted every 5 seconds if it has changed by more than 10%, or if more than 20 seconds have passed since the last adjustment. Setting this property to 0 disables automatic cost adjustment. Automatic adjustment does not work for WDS links that are manually configured as a bridge port.",
			},
			"wds_default_bridge": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When WDS link is established and status of the wds interface becomes running , it will be added as a bridge port to the bridge interface specified by this property. When WDS link is lost, wds interface is removed from the bridge. If wds interface is already included in a bridge setup when WDS link becomes active, it will not be added to bridge specified by , and will (needs editing)",
			},
			"wds_default_cost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Initial bridge port cost of the WDS links.",
			},
			"wds_ignore_ssid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "By default, WDS link between two APs can be created only when they work on the same frequency and have the same SSID value. If this property is set to yes , then SSID of the remote AP will not be checked. This property has no effect on connections from clients in station-wds mode. It also does not work if wds-mode is static-mesh or dynamic-mesh .",
			},
			"wds_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Controls how WDS links with other devices (APs and clients in station-wds mode) are established. disabled does not allow WDS links. static only allows WDS links that are manually configured in WDS dynamic also allows WDS links with devices that are not configured in WDS, by creating required entries dynamically. Such dynamic WDS entries are removed automatically after the connection with the other AP is lost. -mesh modes use different (better) method for establishing link between AP, that is not compatible with APs in non-mesh mode. This method avoids one-sided WDS links that are created only by one of the two APs. Such links cannot pass any data.When AP or station is establishing WDS connection with another AP, it uses connect-list to check whether this connection is allowed. If station in station-wds mode is establishing connection with AP, AP uses access-list to check whether this connection is allowed.If mode is station-wds, then this property has no effect.",
			},
			"wireless_protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies protocol used on wireless interface; unspecified - protocol mode used on previous RouterOS versions (v3.x, v4.x). Nstreme is enabled by old enable-nstreme setting, Nv2 configuration is not possible. any \u00a0: on AP - regular 802.11 Access Point or Nstreme Access Point; on station - selects Access Point without specific sequence, it could be changed by connect-list rules. nstreme - enables Nstreme protocol (the same as old enable-nstreme setting). nv2 - enables Nv2 protocol. nv2 nstreme \u00a0: on AP - uses first wireless-protocol setting, always Nv2; on station - searches for Nv2 Access Point, then for Nstreme Access Point. nv2 nstreme 802.11 - on AP - uses first wireless-protocol setting, always Nv2; on station - searches for Nv2 Access Point, then for Nstreme Access Point, then for regular 802.11 Access Point. Warning! Nv2 doesn't have support for Virtual AP",
			},
			"wmm_support": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies whether to enable WMM . \u00a0Only applies to bands B and G. Other bands will have it enabled regardless of this setting",
			},
			"wps_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Read more >>",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceWirelessResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceWirelessModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AdaptiveNoiseImmunity.IsNull() || plan.AdaptiveNoiseImmunity.IsUnknown()) {
		body["adaptive-noise-immunity"] = plan.AdaptiveNoiseImmunity.ValueString()
	}
	if !(plan.AllowSharedkey.IsNull() || plan.AllowSharedkey.IsUnknown()) {
		body["allow-sharedkey"] = plan.AllowSharedkey.ValueString()
	}
	if !(plan.AmpduPriorities.IsNull() || plan.AmpduPriorities.IsUnknown()) {
		body["ampdu-priorities"] = plan.AmpduPriorities.ValueString()
	}
	if !(plan.AmsduLimit.IsNull() || plan.AmsduLimit.IsUnknown()) {
		body["amsdu-limit"] = plan.AmsduLimit.ValueString()
	}
	if !(plan.AmsduThreshold.IsNull() || plan.AmsduThreshold.IsUnknown()) {
		body["amsdu-threshold"] = plan.AmsduThreshold.ValueString()
	}
	if !(plan.AntennaGain.IsNull() || plan.AntennaGain.IsUnknown()) {
		body["antenna-gain"] = plan.AntennaGain.ValueString()
	}
	if !(plan.AntennaMode.IsNull() || plan.AntennaMode.IsUnknown()) {
		body["antenna-mode"] = plan.AntennaMode.ValueString()
	}
	if !(plan.Area.IsNull() || plan.Area.IsUnknown()) {
		body["area"] = plan.Area.ValueString()
	}
	if !(plan.ARP.IsNull() || plan.ARP.IsUnknown()) {
		body["arp"] = plan.ARP.ValueString()
	}
	if !(plan.ARPTimeout.IsNull() || plan.ARPTimeout.IsUnknown()) {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !(plan.Band.IsNull() || plan.Band.IsUnknown()) {
		body["band"] = plan.Band.ValueString()
	}
	if !(plan.BasicRatesB.IsNull() || plan.BasicRatesB.IsUnknown()) {
		body["basic-rates-b"] = plan.BasicRatesB.ValueString()
	}
	if !(plan.BridgeMode.IsNull() || plan.BridgeMode.IsUnknown()) {
		body["bridge-mode"] = plan.BridgeMode.ValueString()
	}
	if !(plan.BurstTime.IsNull() || plan.BurstTime.IsUnknown()) {
		body["burst-time"] = plan.BurstTime.ValueString()
	}
	if !(plan.ChannelWidth.IsNull() || plan.ChannelWidth.IsUnknown()) {
		body["channel-width"] = plan.ChannelWidth.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Compression.IsNull() || plan.Compression.IsUnknown()) {
		body["compression"] = plan.Compression.ValueString()
	}
	if !(plan.Country.IsNull() || plan.Country.IsUnknown()) {
		body["country"] = plan.Country.ValueString()
	}
	if !(plan.DefaultApTxLimit.IsNull() || plan.DefaultApTxLimit.IsUnknown()) {
		body["default-ap-tx-limit"] = plan.DefaultApTxLimit.ValueString()
	}
	if !(plan.DefaultAuthentication.IsNull() || plan.DefaultAuthentication.IsUnknown()) {
		body["default-authentication"] = plan.DefaultAuthentication.ValueString()
	}
	if !(plan.DefaultClientTxLimit.IsNull() || plan.DefaultClientTxLimit.IsUnknown()) {
		body["default-client-tx-limit"] = plan.DefaultClientTxLimit.ValueString()
	}
	if !(plan.DefaultForwarding.IsNull() || plan.DefaultForwarding.IsUnknown()) {
		body["default-forwarding"] = plan.DefaultForwarding.ValueString()
	}
	if !(plan.DisableRunningCheck.IsNull() || plan.DisableRunningCheck.IsUnknown()) {
		body["disable-running-check"] = plan.DisableRunningCheck.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.DisconnectTimeout.IsNull() || plan.DisconnectTimeout.IsUnknown()) {
		body["disconnect-timeout"] = plan.DisconnectTimeout.ValueString()
	}
	if !(plan.Distance.IsNull() || plan.Distance.IsUnknown()) {
		body["distance"] = plan.Distance.ValueString()
	}
	if !(plan.FrameLifetime.IsNull() || plan.FrameLifetime.IsUnknown()) {
		body["frame-lifetime"] = plan.FrameLifetime.ValueString()
	}
	if !(plan.Frequency.IsNull() || plan.Frequency.IsUnknown()) {
		body["frequency"] = plan.Frequency.ValueString()
	}
	if !(plan.FrequencyMode.IsNull() || plan.FrequencyMode.IsUnknown()) {
		body["frequency-mode"] = plan.FrequencyMode.ValueString()
	}
	if !(plan.FrequencyOffset.IsNull() || plan.FrequencyOffset.IsUnknown()) {
		body["frequency-offset"] = plan.FrequencyOffset.ValueString()
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
	if !(plan.HwFragmentationThreshold.IsNull() || plan.HwFragmentationThreshold.IsUnknown()) {
		body["hw-fragmentation-threshold"] = plan.HwFragmentationThreshold.ValueString()
	}
	if !(plan.HwProtectionMode.IsNull() || plan.HwProtectionMode.IsUnknown()) {
		body["hw-protection-mode"] = plan.HwProtectionMode.ValueString()
	}
	if !(plan.HwProtectionThreshold.IsNull() || plan.HwProtectionThreshold.IsUnknown()) {
		body["hw-protection-threshold"] = plan.HwProtectionThreshold.ValueString()
	}
	if !(plan.HwRetries.IsNull() || plan.HwRetries.IsUnknown()) {
		body["hw-retries"] = plan.HwRetries.ValueString()
	}
	if !(plan.Installation.IsNull() || plan.Installation.IsUnknown()) {
		body["installation"] = plan.Installation.ValueString()
	}
	if !(plan.InterworkingProfile.IsNull() || plan.InterworkingProfile.IsUnknown()) {
		body["interworking-profile"] = plan.InterworkingProfile.ValueString()
	}
	if !(plan.KeepaliveFrames.IsNull() || plan.KeepaliveFrames.IsUnknown()) {
		body["keepalive-frames"] = plan.KeepaliveFrames.ValueString()
	}
	if !(plan.L2mtu.IsNull() || plan.L2mtu.IsUnknown()) {
		body["l2mtu"] = plan.L2mtu.ValueString()
	}
	if !(plan.MACAddress.IsNull() || plan.MACAddress.IsUnknown()) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !(plan.MasterInterface.IsNull() || plan.MasterInterface.IsUnknown()) {
		body["master-interface"] = plan.MasterInterface.ValueString()
	}
	if !(plan.MaxStationCount.IsNull() || plan.MaxStationCount.IsUnknown()) {
		body["max-station-count"] = plan.MaxStationCount.ValueString()
	}
	if !(plan.Mode.IsNull() || plan.Mode.IsUnknown()) {
		body["mode"] = plan.Mode.ValueString()
	}
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !(plan.MulticastBuffering.IsNull() || plan.MulticastBuffering.IsUnknown()) {
		body["multicast-buffering"] = plan.MulticastBuffering.ValueString()
	}
	if !(plan.MulticastHelper.IsNull() || plan.MulticastHelper.IsUnknown()) {
		body["multicast-helper"] = plan.MulticastHelper.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.NoiseFloorThreshold.IsNull() || plan.NoiseFloorThreshold.IsUnknown()) {
		body["noise-floor-threshold"] = plan.NoiseFloorThreshold.ValueString()
	}
	if !(plan.Nv2CellRADIUS.IsNull() || plan.Nv2CellRADIUS.IsUnknown()) {
		body["nv2-cell-radius"] = plan.Nv2CellRADIUS.ValueString()
	}
	if !(plan.Nv2NoiseFloorOffset.IsNull() || plan.Nv2NoiseFloorOffset.IsUnknown()) {
		body["nv2-noise-floor-offset"] = plan.Nv2NoiseFloorOffset.ValueString()
	}
	if !(plan.Nv2PresharedKey.IsNull() || plan.Nv2PresharedKey.IsUnknown()) {
		body["nv2-preshared-key"] = plan.Nv2PresharedKey.ValueString()
	}
	if !(plan.Nv2Qos.IsNull() || plan.Nv2Qos.IsUnknown()) {
		body["nv2-qos"] = plan.Nv2Qos.ValueString()
	}
	if !(plan.Nv2QueueCount.IsNull() || plan.Nv2QueueCount.IsUnknown()) {
		body["nv2-queue-count"] = plan.Nv2QueueCount.ValueString()
	}
	if !(plan.Nv2Security.IsNull() || plan.Nv2Security.IsUnknown()) {
		body["nv2-security"] = plan.Nv2Security.ValueString()
	}
	if !(plan.OnFailRetryTime.IsNull() || plan.OnFailRetryTime.IsUnknown()) {
		body["on-fail-retry-time"] = plan.OnFailRetryTime.ValueString()
	}
	if !(plan.PreambleMode.IsNull() || plan.PreambleMode.IsUnknown()) {
		body["preamble-mode"] = plan.PreambleMode.ValueString()
	}
	if !(plan.PrismCardtype.IsNull() || plan.PrismCardtype.IsUnknown()) {
		body["prism-cardtype"] = plan.PrismCardtype.ValueString()
	}
	if !(plan.RadioName.IsNull() || plan.RadioName.IsUnknown()) {
		body["radio-name"] = plan.RadioName.ValueString()
	}
	if !(plan.RateSelection.IsNull() || plan.RateSelection.IsUnknown()) {
		body["rate-selection"] = plan.RateSelection.ValueString()
	}
	if !(plan.RateSet.IsNull() || plan.RateSet.IsUnknown()) {
		body["rate-set"] = plan.RateSet.ValueString()
	}
	if !(plan.RxChains.IsNull() || plan.RxChains.IsUnknown()) {
		body["rx-chains"] = plan.RxChains.ValueString()
	}
	if !(plan.RxHtChainNames.IsNull() || plan.RxHtChainNames.IsUnknown()) {
		body["rx-ht-chain-names"] = plan.RxHtChainNames.ValueString()
	}
	if !(plan.RxHtChains.IsNull() || plan.RxHtChains.IsUnknown()) {
		body["rx-ht-chains"] = plan.RxHtChains.ValueString()
	}
	if !(plan.ScanList.IsNull() || plan.ScanList.IsUnknown()) {
		body["scan-list"] = plan.ScanList.ValueString()
	}
	if !(plan.SecurityProfile.IsNull() || plan.SecurityProfile.IsUnknown()) {
		body["security-profile"] = plan.SecurityProfile.ValueString()
	}
	if !(plan.SkipDfsChannels.IsNull() || plan.SkipDfsChannels.IsUnknown()) {
		body["skip-dfs-channels"] = plan.SkipDfsChannels.ValueString()
	}
	if !(plan.Ssid.IsNull() || plan.Ssid.IsUnknown()) {
		body["ssid"] = plan.Ssid.ValueString()
	}
	if !(plan.StationBridgeCloneMAC.IsNull() || plan.StationBridgeCloneMAC.IsUnknown()) {
		body["station-bridge-clone-mac"] = plan.StationBridgeCloneMAC.ValueString()
	}
	if !(plan.StationRoaming.IsNull() || plan.StationRoaming.IsUnknown()) {
		body["station-roaming"] = plan.StationRoaming.ValueString()
	}
	if !(plan.SupportedRatesB.IsNull() || plan.SupportedRatesB.IsUnknown()) {
		body["supported-rates-b"] = plan.SupportedRatesB.ValueString()
	}
	if !(plan.TdmaPeriodSize.IsNull() || plan.TdmaPeriodSize.IsUnknown()) {
		body["tdma-period-size"] = plan.TdmaPeriodSize.ValueString()
	}
	if !(plan.TxChains.IsNull() || plan.TxChains.IsUnknown()) {
		body["tx-chains"] = plan.TxChains.ValueString()
	}
	if !(plan.TxHtChainNames.IsNull() || plan.TxHtChainNames.IsUnknown()) {
		body["tx-ht-chain-names"] = plan.TxHtChainNames.ValueString()
	}
	if !(plan.TxHtChains.IsNull() || plan.TxHtChains.IsUnknown()) {
		body["tx-ht-chains"] = plan.TxHtChains.ValueString()
	}
	if !(plan.TxPower.IsNull() || plan.TxPower.IsUnknown()) {
		body["tx-power"] = plan.TxPower.ValueString()
	}
	if !(plan.TxPowerMode.IsNull() || plan.TxPowerMode.IsUnknown()) {
		body["tx-power-mode"] = plan.TxPowerMode.ValueString()
	}
	if !(plan.UpdateStatsInterval.IsNull() || plan.UpdateStatsInterval.IsUnknown()) {
		body["update-stats-interval"] = plan.UpdateStatsInterval.ValueString()
	}
	if !(plan.VhtBasicMcs.IsNull() || plan.VhtBasicMcs.IsUnknown()) {
		body["vht-basic-mcs"] = plan.VhtBasicMcs.ValueString()
	}
	if !(plan.VhtSupportedMcs.IsNull() || plan.VhtSupportedMcs.IsUnknown()) {
		body["vht-supported-mcs"] = plan.VhtSupportedMcs.ValueString()
	}
	if !(plan.WdsCostRange.IsNull() || plan.WdsCostRange.IsUnknown()) {
		body["wds-cost-range"] = plan.WdsCostRange.ValueString()
	}
	if !(plan.WdsDefaultBridge.IsNull() || plan.WdsDefaultBridge.IsUnknown()) {
		body["wds-default-bridge"] = plan.WdsDefaultBridge.ValueString()
	}
	if !(plan.WdsDefaultCost.IsNull() || plan.WdsDefaultCost.IsUnknown()) {
		body["wds-default-cost"] = plan.WdsDefaultCost.ValueString()
	}
	if !(plan.WdsIgnoreSsid.IsNull() || plan.WdsIgnoreSsid.IsUnknown()) {
		body["wds-ignore-ssid"] = plan.WdsIgnoreSsid.ValueString()
	}
	if !(plan.WdsMode.IsNull() || plan.WdsMode.IsUnknown()) {
		body["wds-mode"] = plan.WdsMode.ValueString()
	}
	if !(plan.WirelessProtocol.IsNull() || plan.WirelessProtocol.IsUnknown()) {
		body["wireless-protocol"] = plan.WirelessProtocol.ValueString()
	}
	if !(plan.WmmSupport.IsNull() || plan.WmmSupport.IsUnknown()) {
		body["wmm-support"] = plan.WmmSupport.ValueString()
	}
	if !(plan.WpsMode.IsNull() || plan.WpsMode.IsUnknown()) {
		body["wps-mode"] = plan.WpsMode.ValueString()
	}
	if !(plan.DfsTestMode.IsNull() || plan.DfsTestMode.IsUnknown()) {
		body["dfs-test-mode"] = plan.DfsTestMode.ValueString()
	}
	if !(plan.Nv2DownlinkRatio.IsNull() || plan.Nv2DownlinkRatio.IsUnknown()) {
		body["nv2-downlink-ratio"] = plan.Nv2DownlinkRatio.ValueString()
	}
	if !(plan.Nv2Mode.IsNull() || plan.Nv2Mode.IsUnknown()) {
		body["nv2-mode"] = plan.Nv2Mode.ValueString()
	}
	if !(plan.Nv2SyncSecret.IsNull() || plan.Nv2SyncSecret.IsUnknown()) {
		body["nv2-sync-secret"] = plan.Nv2SyncSecret.ValueString()
	}
	if !(plan.SecondaryFrequency.IsNull() || plan.SecondaryFrequency.IsUnknown()) {
		body["secondary-frequency"] = plan.SecondaryFrequency.ValueString()
	}
	if !(plan.VlanId.IsNull() || plan.VlanId.IsUnknown()) {
		body["vlan-id"] = plan.VlanId.ValueString()
	}
	if !(plan.VlanMode.IsNull() || plan.VlanMode.IsUnknown()) {
		body["vlan-mode"] = plan.VlanMode.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/wireless", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/wireless failed", err.Error())
		return
	}
	interfaceWirelessApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWirelessResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceWirelessModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/wireless", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/wireless failed", err.Error())
		return
	}
	interfaceWirelessApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceWirelessResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceWirelessModel
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
	if !plan.AdaptiveNoiseImmunity.Equal(state.AdaptiveNoiseImmunity) {
		body["adaptive-noise-immunity"] = plan.AdaptiveNoiseImmunity.ValueString()
	}
	if !plan.AllowSharedkey.Equal(state.AllowSharedkey) {
		body["allow-sharedkey"] = plan.AllowSharedkey.ValueString()
	}
	if !plan.AmpduPriorities.Equal(state.AmpduPriorities) {
		body["ampdu-priorities"] = plan.AmpduPriorities.ValueString()
	}
	if !plan.AmsduLimit.Equal(state.AmsduLimit) {
		body["amsdu-limit"] = plan.AmsduLimit.ValueString()
	}
	if !plan.AmsduThreshold.Equal(state.AmsduThreshold) {
		body["amsdu-threshold"] = plan.AmsduThreshold.ValueString()
	}
	if !plan.AntennaGain.Equal(state.AntennaGain) {
		body["antenna-gain"] = plan.AntennaGain.ValueString()
	}
	if !plan.AntennaMode.Equal(state.AntennaMode) {
		body["antenna-mode"] = plan.AntennaMode.ValueString()
	}
	if !plan.Area.Equal(state.Area) {
		body["area"] = plan.Area.ValueString()
	}
	if !plan.ARP.Equal(state.ARP) {
		body["arp"] = plan.ARP.ValueString()
	}
	if !plan.ARPTimeout.Equal(state.ARPTimeout) {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !plan.Band.Equal(state.Band) {
		body["band"] = plan.Band.ValueString()
	}
	if !plan.BasicRatesB.Equal(state.BasicRatesB) {
		body["basic-rates-b"] = plan.BasicRatesB.ValueString()
	}
	if !plan.BridgeMode.Equal(state.BridgeMode) {
		body["bridge-mode"] = plan.BridgeMode.ValueString()
	}
	if !plan.BurstTime.Equal(state.BurstTime) {
		body["burst-time"] = plan.BurstTime.ValueString()
	}
	if !plan.ChannelWidth.Equal(state.ChannelWidth) {
		body["channel-width"] = plan.ChannelWidth.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Compression.Equal(state.Compression) {
		body["compression"] = plan.Compression.ValueString()
	}
	if !plan.Country.Equal(state.Country) {
		body["country"] = plan.Country.ValueString()
	}
	if !plan.DefaultApTxLimit.Equal(state.DefaultApTxLimit) {
		body["default-ap-tx-limit"] = plan.DefaultApTxLimit.ValueString()
	}
	if !plan.DefaultAuthentication.Equal(state.DefaultAuthentication) {
		body["default-authentication"] = plan.DefaultAuthentication.ValueString()
	}
	if !plan.DefaultClientTxLimit.Equal(state.DefaultClientTxLimit) {
		body["default-client-tx-limit"] = plan.DefaultClientTxLimit.ValueString()
	}
	if !plan.DefaultForwarding.Equal(state.DefaultForwarding) {
		body["default-forwarding"] = plan.DefaultForwarding.ValueString()
	}
	if !plan.DisableRunningCheck.Equal(state.DisableRunningCheck) {
		body["disable-running-check"] = plan.DisableRunningCheck.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DisconnectTimeout.Equal(state.DisconnectTimeout) {
		body["disconnect-timeout"] = plan.DisconnectTimeout.ValueString()
	}
	if !plan.Distance.Equal(state.Distance) {
		body["distance"] = plan.Distance.ValueString()
	}
	if !plan.FrameLifetime.Equal(state.FrameLifetime) {
		body["frame-lifetime"] = plan.FrameLifetime.ValueString()
	}
	if !plan.Frequency.Equal(state.Frequency) {
		body["frequency"] = plan.Frequency.ValueString()
	}
	if !plan.FrequencyMode.Equal(state.FrequencyMode) {
		body["frequency-mode"] = plan.FrequencyMode.ValueString()
	}
	if !plan.FrequencyOffset.Equal(state.FrequencyOffset) {
		body["frequency-offset"] = plan.FrequencyOffset.ValueString()
	}
	if !plan.GuardInterval.Equal(state.GuardInterval) {
		body["guard-interval"] = plan.GuardInterval.ValueString()
	}
	if !plan.HideSsid.Equal(state.HideSsid) {
		body["hide-ssid"] = plan.HideSsid.ValueString()
	}
	if !plan.HtBasicMcs.Equal(state.HtBasicMcs) {
		body["ht-basic-mcs"] = plan.HtBasicMcs.ValueString()
	}
	if !plan.HtSupportedMcs.Equal(state.HtSupportedMcs) {
		body["ht-supported-mcs"] = plan.HtSupportedMcs.ValueString()
	}
	if !plan.HwFragmentationThreshold.Equal(state.HwFragmentationThreshold) {
		body["hw-fragmentation-threshold"] = plan.HwFragmentationThreshold.ValueString()
	}
	if !plan.HwProtectionMode.Equal(state.HwProtectionMode) {
		body["hw-protection-mode"] = plan.HwProtectionMode.ValueString()
	}
	if !plan.HwProtectionThreshold.Equal(state.HwProtectionThreshold) {
		body["hw-protection-threshold"] = plan.HwProtectionThreshold.ValueString()
	}
	if !plan.HwRetries.Equal(state.HwRetries) {
		body["hw-retries"] = plan.HwRetries.ValueString()
	}
	if !plan.Installation.Equal(state.Installation) {
		body["installation"] = plan.Installation.ValueString()
	}
	if !plan.InterworkingProfile.Equal(state.InterworkingProfile) {
		body["interworking-profile"] = plan.InterworkingProfile.ValueString()
	}
	if !plan.KeepaliveFrames.Equal(state.KeepaliveFrames) {
		body["keepalive-frames"] = plan.KeepaliveFrames.ValueString()
	}
	if !plan.L2mtu.Equal(state.L2mtu) {
		body["l2mtu"] = plan.L2mtu.ValueString()
	}
	if !plan.MACAddress.Equal(state.MACAddress) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.MasterInterface.Equal(state.MasterInterface) {
		body["master-interface"] = plan.MasterInterface.ValueString()
	}
	if !plan.MaxStationCount.Equal(state.MaxStationCount) {
		body["max-station-count"] = plan.MaxStationCount.ValueString()
	}
	if !plan.Mode.Equal(state.Mode) {
		body["mode"] = plan.Mode.ValueString()
	}
	if !plan.MTU.Equal(state.MTU) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !plan.MulticastBuffering.Equal(state.MulticastBuffering) {
		body["multicast-buffering"] = plan.MulticastBuffering.ValueString()
	}
	if !plan.MulticastHelper.Equal(state.MulticastHelper) {
		body["multicast-helper"] = plan.MulticastHelper.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.NoiseFloorThreshold.Equal(state.NoiseFloorThreshold) {
		body["noise-floor-threshold"] = plan.NoiseFloorThreshold.ValueString()
	}
	if !plan.Nv2CellRADIUS.Equal(state.Nv2CellRADIUS) {
		body["nv2-cell-radius"] = plan.Nv2CellRADIUS.ValueString()
	}
	if !plan.Nv2NoiseFloorOffset.Equal(state.Nv2NoiseFloorOffset) {
		body["nv2-noise-floor-offset"] = plan.Nv2NoiseFloorOffset.ValueString()
	}
	if !plan.Nv2PresharedKey.Equal(state.Nv2PresharedKey) {
		body["nv2-preshared-key"] = plan.Nv2PresharedKey.ValueString()
	}
	if !plan.Nv2Qos.Equal(state.Nv2Qos) {
		body["nv2-qos"] = plan.Nv2Qos.ValueString()
	}
	if !plan.Nv2QueueCount.Equal(state.Nv2QueueCount) {
		body["nv2-queue-count"] = plan.Nv2QueueCount.ValueString()
	}
	if !plan.Nv2Security.Equal(state.Nv2Security) {
		body["nv2-security"] = plan.Nv2Security.ValueString()
	}
	if !plan.OnFailRetryTime.Equal(state.OnFailRetryTime) {
		body["on-fail-retry-time"] = plan.OnFailRetryTime.ValueString()
	}
	if !plan.PreambleMode.Equal(state.PreambleMode) {
		body["preamble-mode"] = plan.PreambleMode.ValueString()
	}
	if !plan.PrismCardtype.Equal(state.PrismCardtype) {
		body["prism-cardtype"] = plan.PrismCardtype.ValueString()
	}
	if !plan.RadioName.Equal(state.RadioName) {
		body["radio-name"] = plan.RadioName.ValueString()
	}
	if !plan.RateSelection.Equal(state.RateSelection) {
		body["rate-selection"] = plan.RateSelection.ValueString()
	}
	if !plan.RateSet.Equal(state.RateSet) {
		body["rate-set"] = plan.RateSet.ValueString()
	}
	if !plan.RxChains.Equal(state.RxChains) {
		body["rx-chains"] = plan.RxChains.ValueString()
	}
	if !plan.RxHtChainNames.Equal(state.RxHtChainNames) {
		body["rx-ht-chain-names"] = plan.RxHtChainNames.ValueString()
	}
	if !plan.RxHtChains.Equal(state.RxHtChains) {
		body["rx-ht-chains"] = plan.RxHtChains.ValueString()
	}
	if !plan.ScanList.Equal(state.ScanList) {
		body["scan-list"] = plan.ScanList.ValueString()
	}
	if !plan.SecurityProfile.Equal(state.SecurityProfile) {
		body["security-profile"] = plan.SecurityProfile.ValueString()
	}
	if !plan.SkipDfsChannels.Equal(state.SkipDfsChannels) {
		body["skip-dfs-channels"] = plan.SkipDfsChannels.ValueString()
	}
	if !plan.Ssid.Equal(state.Ssid) {
		body["ssid"] = plan.Ssid.ValueString()
	}
	if !plan.StationBridgeCloneMAC.Equal(state.StationBridgeCloneMAC) {
		body["station-bridge-clone-mac"] = plan.StationBridgeCloneMAC.ValueString()
	}
	if !plan.StationRoaming.Equal(state.StationRoaming) {
		body["station-roaming"] = plan.StationRoaming.ValueString()
	}
	if !plan.SupportedRatesB.Equal(state.SupportedRatesB) {
		body["supported-rates-b"] = plan.SupportedRatesB.ValueString()
	}
	if !plan.TdmaPeriodSize.Equal(state.TdmaPeriodSize) {
		body["tdma-period-size"] = plan.TdmaPeriodSize.ValueString()
	}
	if !plan.TxChains.Equal(state.TxChains) {
		body["tx-chains"] = plan.TxChains.ValueString()
	}
	if !plan.TxHtChainNames.Equal(state.TxHtChainNames) {
		body["tx-ht-chain-names"] = plan.TxHtChainNames.ValueString()
	}
	if !plan.TxHtChains.Equal(state.TxHtChains) {
		body["tx-ht-chains"] = plan.TxHtChains.ValueString()
	}
	if !plan.TxPower.Equal(state.TxPower) {
		body["tx-power"] = plan.TxPower.ValueString()
	}
	if !plan.TxPowerMode.Equal(state.TxPowerMode) {
		body["tx-power-mode"] = plan.TxPowerMode.ValueString()
	}
	if !plan.UpdateStatsInterval.Equal(state.UpdateStatsInterval) {
		body["update-stats-interval"] = plan.UpdateStatsInterval.ValueString()
	}
	if !plan.VhtBasicMcs.Equal(state.VhtBasicMcs) {
		body["vht-basic-mcs"] = plan.VhtBasicMcs.ValueString()
	}
	if !plan.VhtSupportedMcs.Equal(state.VhtSupportedMcs) {
		body["vht-supported-mcs"] = plan.VhtSupportedMcs.ValueString()
	}
	if !plan.WdsCostRange.Equal(state.WdsCostRange) {
		body["wds-cost-range"] = plan.WdsCostRange.ValueString()
	}
	if !plan.WdsDefaultBridge.Equal(state.WdsDefaultBridge) {
		body["wds-default-bridge"] = plan.WdsDefaultBridge.ValueString()
	}
	if !plan.WdsDefaultCost.Equal(state.WdsDefaultCost) {
		body["wds-default-cost"] = plan.WdsDefaultCost.ValueString()
	}
	if !plan.WdsIgnoreSsid.Equal(state.WdsIgnoreSsid) {
		body["wds-ignore-ssid"] = plan.WdsIgnoreSsid.ValueString()
	}
	if !plan.WdsMode.Equal(state.WdsMode) {
		body["wds-mode"] = plan.WdsMode.ValueString()
	}
	if !plan.WirelessProtocol.Equal(state.WirelessProtocol) {
		body["wireless-protocol"] = plan.WirelessProtocol.ValueString()
	}
	if !plan.WmmSupport.Equal(state.WmmSupport) {
		body["wmm-support"] = plan.WmmSupport.ValueString()
	}
	if !plan.WpsMode.Equal(state.WpsMode) {
		body["wps-mode"] = plan.WpsMode.ValueString()
	}
	if !plan.DfsTestMode.Equal(state.DfsTestMode) && !plan.DfsTestMode.IsUnknown() {
		body["dfs-test-mode"] = plan.DfsTestMode.ValueString()
	}
	if !plan.Nv2DownlinkRatio.Equal(state.Nv2DownlinkRatio) && !plan.Nv2DownlinkRatio.IsUnknown() {
		body["nv2-downlink-ratio"] = plan.Nv2DownlinkRatio.ValueString()
	}
	if !plan.Nv2Mode.Equal(state.Nv2Mode) && !plan.Nv2Mode.IsUnknown() {
		body["nv2-mode"] = plan.Nv2Mode.ValueString()
	}
	if !plan.Nv2SyncSecret.Equal(state.Nv2SyncSecret) && !plan.Nv2SyncSecret.IsUnknown() {
		body["nv2-sync-secret"] = plan.Nv2SyncSecret.ValueString()
	}
	if !plan.SecondaryFrequency.Equal(state.SecondaryFrequency) && !plan.SecondaryFrequency.IsUnknown() {
		body["secondary-frequency"] = plan.SecondaryFrequency.ValueString()
	}
	if !plan.VlanId.Equal(state.VlanId) && !plan.VlanId.IsUnknown() {
		body["vlan-id"] = plan.VlanId.ValueString()
	}
	if !plan.VlanMode.Equal(state.VlanMode) && !plan.VlanMode.IsUnknown() {
		body["vlan-mode"] = plan.VlanMode.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/wireless", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/wireless failed", err.Error())
			return
		}
		interfaceWirelessApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWirelessResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceWirelessModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/wireless", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/wireless failed", err.Error())
	}
}

func (r *InterfaceWirelessResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceWirelessLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/wireless matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceWirelessLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceWirelessLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/wireless", id)
}

func interfaceWirelessApply(ctx context.Context, obj client.Object, m *InterfaceWirelessModel) {
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
	if v, ok := obj["secondary-frequency"]; ok && v != "" {
		m.SecondaryFrequency = types.StringValue(v)
	} else {
		m.SecondaryFrequency = types.StringNull()
	}
	if v, ok := obj["nv2-sync-secret"]; ok && v != "" {
		m.Nv2SyncSecret = types.StringValue(v)
	} else {
		m.Nv2SyncSecret = types.StringNull()
	}
	if v, ok := obj["nv2-mode"]; ok && v != "" {
		m.Nv2Mode = types.StringValue(v)
	} else {
		m.Nv2Mode = types.StringNull()
	}
	if v, ok := obj["nv2-downlink-ratio"]; ok && v != "" {
		m.Nv2DownlinkRatio = types.StringValue(v)
	} else {
		m.Nv2DownlinkRatio = types.StringNull()
	}
	if v, ok := obj["dfs-test-mode"]; ok && v != "" {
		m.DfsTestMode = types.StringValue(v)
	} else {
		m.DfsTestMode = types.StringNull()
	}
	if v, ok := obj["adaptive-noise-immunity"]; ok {
		_ = v
		if v != "" {
			m.AdaptiveNoiseImmunity = types.StringValue(v)
		} else {
			m.AdaptiveNoiseImmunity = types.StringNull()
		}
	} else {
		m.AdaptiveNoiseImmunity = types.StringNull()
	}
	if v, ok := obj["allow-sharedkey"]; ok {
		_ = v
		if v != "" {
			m.AllowSharedkey = types.StringValue(v)
		} else {
			m.AllowSharedkey = types.StringNull()
		}
	} else {
		m.AllowSharedkey = types.StringNull()
	}
	if v, ok := obj["ampdu-priorities"]; ok {
		_ = v
		if v != "" {
			m.AmpduPriorities = types.StringValue(v)
		} else {
			m.AmpduPriorities = types.StringNull()
		}
	} else {
		m.AmpduPriorities = types.StringNull()
	}
	if v, ok := obj["amsdu-limit"]; ok {
		_ = v
		if v != "" {
			m.AmsduLimit = types.StringValue(v)
		} else {
			m.AmsduLimit = types.StringNull()
		}
	} else {
		m.AmsduLimit = types.StringNull()
	}
	if v, ok := obj["amsdu-threshold"]; ok {
		_ = v
		if v != "" {
			m.AmsduThreshold = types.StringValue(v)
		} else {
			m.AmsduThreshold = types.StringNull()
		}
	} else {
		m.AmsduThreshold = types.StringNull()
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
	if v, ok := obj["antenna-mode"]; ok {
		_ = v
		if v != "" {
			m.AntennaMode = types.StringValue(v)
		} else {
			m.AntennaMode = types.StringNull()
		}
	} else {
		m.AntennaMode = types.StringNull()
	}
	if v, ok := obj["area"]; ok {
		_ = v
		if v != "" {
			m.Area = types.StringValue(v)
		} else {
			m.Area = types.StringNull()
		}
	} else {
		m.Area = types.StringNull()
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
	if v, ok := obj["band"]; ok {
		_ = v
		if v != "" {
			m.Band = types.StringValue(v)
		} else {
			m.Band = types.StringNull()
		}
	} else {
		m.Band = types.StringNull()
	}
	if v, ok := obj["basic-rates-b"]; ok {
		_ = v
		if v != "" {
			m.BasicRatesB = types.StringValue(v)
		} else {
			m.BasicRatesB = types.StringNull()
		}
	} else {
		m.BasicRatesB = types.StringNull()
	}
	if v, ok := obj["bridge-mode"]; ok {
		_ = v
		if v != "" {
			m.BridgeMode = types.StringValue(v)
		} else {
			m.BridgeMode = types.StringNull()
		}
	} else {
		m.BridgeMode = types.StringNull()
	}
	if v, ok := obj["burst-time"]; ok {
		_ = v
		if v != "" {
			m.BurstTime = types.StringValue(v)
		} else {
			m.BurstTime = types.StringNull()
		}
	} else {
		m.BurstTime = types.StringNull()
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
	if v, ok := obj["compression"]; ok {
		_ = v
		if v != "" {
			m.Compression = types.StringValue(v)
		} else {
			m.Compression = types.StringNull()
		}
	} else {
		m.Compression = types.StringNull()
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
	if v, ok := obj["default-ap-tx-limit"]; ok {
		_ = v
		if v != "" {
			m.DefaultApTxLimit = types.StringValue(v)
		} else {
			m.DefaultApTxLimit = types.StringNull()
		}
	} else {
		m.DefaultApTxLimit = types.StringNull()
	}
	if v, ok := obj["default-authentication"]; ok {
		_ = v
		if v != "" {
			m.DefaultAuthentication = types.StringValue(v)
		} else {
			m.DefaultAuthentication = types.StringNull()
		}
	} else {
		m.DefaultAuthentication = types.StringNull()
	}
	if v, ok := obj["default-client-tx-limit"]; ok {
		_ = v
		if v != "" {
			m.DefaultClientTxLimit = types.StringValue(v)
		} else {
			m.DefaultClientTxLimit = types.StringNull()
		}
	} else {
		m.DefaultClientTxLimit = types.StringNull()
	}
	if v, ok := obj["default-forwarding"]; ok {
		_ = v
		if v != "" {
			m.DefaultForwarding = types.StringValue(v)
		} else {
			m.DefaultForwarding = types.StringNull()
		}
	} else {
		m.DefaultForwarding = types.StringNull()
	}
	if v, ok := obj["disable-running-check"]; ok {
		_ = v
		if v != "" {
			m.DisableRunningCheck = types.StringValue(v)
		} else {
			m.DisableRunningCheck = types.StringNull()
		}
	} else {
		m.DisableRunningCheck = types.StringNull()
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
	if v, ok := obj["disconnect-timeout"]; ok {
		_ = v
		if v != "" {
			m.DisconnectTimeout = types.StringValue(v)
		} else {
			m.DisconnectTimeout = types.StringNull()
		}
	} else {
		m.DisconnectTimeout = types.StringNull()
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
	if v, ok := obj["frame-lifetime"]; ok {
		_ = v
		if v != "" {
			m.FrameLifetime = types.StringValue(v)
		} else {
			m.FrameLifetime = types.StringNull()
		}
	} else {
		m.FrameLifetime = types.StringNull()
	}
	if v, ok := obj["frequency"]; ok {
		_ = v
		if v != "" {
			m.Frequency = types.StringValue(v)
		} else {
			m.Frequency = types.StringNull()
		}
	} else {
		m.Frequency = types.StringNull()
	}
	if v, ok := obj["frequency-mode"]; ok {
		_ = v
		if v != "" {
			m.FrequencyMode = types.StringValue(v)
		} else {
			m.FrequencyMode = types.StringNull()
		}
	} else {
		m.FrequencyMode = types.StringNull()
	}
	if v, ok := obj["frequency-offset"]; ok {
		_ = v
		if v != "" {
			m.FrequencyOffset = types.StringValue(v)
		} else {
			m.FrequencyOffset = types.StringNull()
		}
	} else {
		m.FrequencyOffset = types.StringNull()
	}
	if v, ok := obj["guard-interval"]; ok {
		_ = v
		if v != "" {
			m.GuardInterval = types.StringValue(v)
		} else {
			m.GuardInterval = types.StringNull()
		}
	} else {
		m.GuardInterval = types.StringNull()
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
	if v, ok := obj["ht-basic-mcs"]; ok {
		_ = v
		if v != "" {
			m.HtBasicMcs = types.StringValue(v)
		} else {
			m.HtBasicMcs = types.StringNull()
		}
	} else {
		m.HtBasicMcs = types.StringNull()
	}
	if v, ok := obj["ht-supported-mcs"]; ok {
		_ = v
		if v != "" {
			m.HtSupportedMcs = types.StringValue(v)
		} else {
			m.HtSupportedMcs = types.StringNull()
		}
	} else {
		m.HtSupportedMcs = types.StringNull()
	}
	if v, ok := obj["hw-fragmentation-threshold"]; ok {
		_ = v
		if v != "" {
			m.HwFragmentationThreshold = types.StringValue(v)
		} else {
			m.HwFragmentationThreshold = types.StringNull()
		}
	} else {
		m.HwFragmentationThreshold = types.StringNull()
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
	if v, ok := obj["hw-protection-threshold"]; ok {
		_ = v
		if v != "" {
			m.HwProtectionThreshold = types.StringValue(v)
		} else {
			m.HwProtectionThreshold = types.StringNull()
		}
	} else {
		m.HwProtectionThreshold = types.StringNull()
	}
	if v, ok := obj["hw-retries"]; ok {
		_ = v
		if v != "" {
			m.HwRetries = types.StringValue(v)
		} else {
			m.HwRetries = types.StringNull()
		}
	} else {
		m.HwRetries = types.StringNull()
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
	if v, ok := obj["interworking-profile"]; ok {
		_ = v
		if v != "" {
			m.InterworkingProfile = types.StringValue(v)
		} else {
			m.InterworkingProfile = types.StringNull()
		}
	} else {
		m.InterworkingProfile = types.StringNull()
	}
	if v, ok := obj["keepalive-frames"]; ok {
		_ = v
		if v != "" {
			m.KeepaliveFrames = types.StringValue(v)
		} else {
			m.KeepaliveFrames = types.StringNull()
		}
	} else {
		m.KeepaliveFrames = types.StringNull()
	}
	if v, ok := obj["l2mtu"]; ok {
		_ = v
		if v != "" {
			m.L2mtu = types.StringValue(v)
		} else {
			m.L2mtu = types.StringNull()
		}
	} else {
		m.L2mtu = types.StringNull()
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
	if v, ok := obj["master-interface"]; ok {
		_ = v
		if v != "" {
			m.MasterInterface = types.StringValue(v)
		} else {
			m.MasterInterface = types.StringNull()
		}
	} else {
		m.MasterInterface = types.StringNull()
	}
	if v, ok := obj["max-station-count"]; ok {
		_ = v
		if v != "" {
			m.MaxStationCount = types.StringValue(v)
		} else {
			m.MaxStationCount = types.StringNull()
		}
	} else {
		m.MaxStationCount = types.StringNull()
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
	if v, ok := obj["mtu"]; ok {
		_ = v
		if v != "" {
			m.MTU = types.StringValue(v)
		} else {
			m.MTU = types.StringNull()
		}
	} else {
		m.MTU = types.StringNull()
	}
	if v, ok := obj["multicast-buffering"]; ok {
		_ = v
		if v != "" {
			m.MulticastBuffering = types.StringValue(v)
		} else {
			m.MulticastBuffering = types.StringNull()
		}
	} else {
		m.MulticastBuffering = types.StringNull()
	}
	if v, ok := obj["multicast-helper"]; ok {
		_ = v
		if v != "" {
			m.MulticastHelper = types.StringValue(v)
		} else {
			m.MulticastHelper = types.StringNull()
		}
	} else {
		m.MulticastHelper = types.StringNull()
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
	if v, ok := obj["noise-floor-threshold"]; ok {
		_ = v
		if v != "" {
			m.NoiseFloorThreshold = types.StringValue(v)
		} else {
			m.NoiseFloorThreshold = types.StringNull()
		}
	} else {
		m.NoiseFloorThreshold = types.StringNull()
	}
	if v, ok := obj["nv2-cell-radius"]; ok {
		_ = v
		if v != "" {
			m.Nv2CellRADIUS = types.StringValue(v)
		} else {
			m.Nv2CellRADIUS = types.StringNull()
		}
	} else {
		m.Nv2CellRADIUS = types.StringNull()
	}
	if v, ok := obj["nv2-noise-floor-offset"]; ok {
		_ = v
		if v != "" {
			m.Nv2NoiseFloorOffset = types.StringValue(v)
		} else {
			m.Nv2NoiseFloorOffset = types.StringNull()
		}
	} else {
		m.Nv2NoiseFloorOffset = types.StringNull()
	}
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.Nv2PresharedKey already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["nv2-preshared-key"]; ok && v != "" {
		_ = v
		if v != "" {
			m.Nv2PresharedKey = types.StringValue(v)
		} else {
			m.Nv2PresharedKey = types.StringNull()
		}
	} else if m.Nv2PresharedKey.IsUnknown() {
		m.Nv2PresharedKey = types.StringNull()
	}
	if v, ok := obj["nv2-qos"]; ok {
		_ = v
		if v != "" {
			m.Nv2Qos = types.StringValue(v)
		} else {
			m.Nv2Qos = types.StringNull()
		}
	} else {
		m.Nv2Qos = types.StringNull()
	}
	if v, ok := obj["nv2-queue-count"]; ok {
		_ = v
		if v != "" {
			m.Nv2QueueCount = types.StringValue(v)
		} else {
			m.Nv2QueueCount = types.StringNull()
		}
	} else {
		m.Nv2QueueCount = types.StringNull()
	}
	if v, ok := obj["nv2-security"]; ok {
		_ = v
		if v != "" {
			m.Nv2Security = types.StringValue(v)
		} else {
			m.Nv2Security = types.StringNull()
		}
	} else {
		m.Nv2Security = types.StringNull()
	}
	if v, ok := obj["on-fail-retry-time"]; ok {
		_ = v
		if v != "" {
			m.OnFailRetryTime = types.StringValue(v)
		} else {
			m.OnFailRetryTime = types.StringNull()
		}
	} else {
		m.OnFailRetryTime = types.StringNull()
	}
	if v, ok := obj["preamble-mode"]; ok {
		_ = v
		if v != "" {
			m.PreambleMode = types.StringValue(v)
		} else {
			m.PreambleMode = types.StringNull()
		}
	} else {
		m.PreambleMode = types.StringNull()
	}
	if v, ok := obj["prism-cardtype"]; ok {
		_ = v
		if v != "" {
			m.PrismCardtype = types.StringValue(v)
		} else {
			m.PrismCardtype = types.StringNull()
		}
	} else {
		m.PrismCardtype = types.StringNull()
	}
	if v, ok := obj["radio-name"]; ok {
		_ = v
		if v != "" {
			m.RadioName = types.StringValue(v)
		} else {
			m.RadioName = types.StringNull()
		}
	} else {
		m.RadioName = types.StringNull()
	}
	if v, ok := obj["rate-selection"]; ok {
		_ = v
		if v != "" {
			m.RateSelection = types.StringValue(v)
		} else {
			m.RateSelection = types.StringNull()
		}
	} else {
		m.RateSelection = types.StringNull()
	}
	if v, ok := obj["rate-set"]; ok {
		_ = v
		if v != "" {
			m.RateSet = types.StringValue(v)
		} else {
			m.RateSet = types.StringNull()
		}
	} else {
		m.RateSet = types.StringNull()
	}
	if v, ok := obj["rx-chains"]; ok {
		_ = v
		if v != "" {
			m.RxChains = types.StringValue(v)
		} else {
			m.RxChains = types.StringNull()
		}
	} else {
		m.RxChains = types.StringNull()
	}
	if v, ok := obj["rx-ht-chain-names"]; ok {
		_ = v
		if v != "" {
			m.RxHtChainNames = types.StringValue(v)
		} else {
			m.RxHtChainNames = types.StringNull()
		}
	} else {
		m.RxHtChainNames = types.StringNull()
	}
	if v, ok := obj["rx-ht-chains"]; ok {
		_ = v
		if v != "" {
			m.RxHtChains = types.StringValue(v)
		} else {
			m.RxHtChains = types.StringNull()
		}
	} else {
		m.RxHtChains = types.StringNull()
	}
	if v, ok := obj["scan-list"]; ok {
		_ = v
		if v != "" {
			m.ScanList = types.StringValue(v)
		} else {
			m.ScanList = types.StringNull()
		}
	} else {
		m.ScanList = types.StringNull()
	}
	if v, ok := obj["security-profile"]; ok {
		_ = v
		if v != "" {
			m.SecurityProfile = types.StringValue(v)
		} else {
			m.SecurityProfile = types.StringNull()
		}
	} else {
		m.SecurityProfile = types.StringNull()
	}
	if v, ok := obj["skip-dfs-channels"]; ok {
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
	if v, ok := obj["station-bridge-clone-mac"]; ok {
		_ = v
		if v != "" {
			m.StationBridgeCloneMAC = types.StringValue(v)
		} else {
			m.StationBridgeCloneMAC = types.StringNull()
		}
	} else {
		m.StationBridgeCloneMAC = types.StringNull()
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
	if v, ok := obj["supported-rates-b"]; ok {
		_ = v
		if v != "" {
			m.SupportedRatesB = types.StringValue(v)
		} else {
			m.SupportedRatesB = types.StringNull()
		}
	} else {
		m.SupportedRatesB = types.StringNull()
	}
	if v, ok := obj["tdma-period-size"]; ok {
		_ = v
		if v != "" {
			m.TdmaPeriodSize = types.StringValue(v)
		} else {
			m.TdmaPeriodSize = types.StringNull()
		}
	} else {
		m.TdmaPeriodSize = types.StringNull()
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
	if v, ok := obj["tx-ht-chain-names"]; ok {
		_ = v
		if v != "" {
			m.TxHtChainNames = types.StringValue(v)
		} else {
			m.TxHtChainNames = types.StringNull()
		}
	} else {
		m.TxHtChainNames = types.StringNull()
	}
	if v, ok := obj["tx-ht-chains"]; ok {
		_ = v
		if v != "" {
			m.TxHtChains = types.StringValue(v)
		} else {
			m.TxHtChains = types.StringNull()
		}
	} else {
		m.TxHtChains = types.StringNull()
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
	if v, ok := obj["tx-power-mode"]; ok {
		_ = v
		if v != "" {
			m.TxPowerMode = types.StringValue(v)
		} else {
			m.TxPowerMode = types.StringNull()
		}
	} else {
		m.TxPowerMode = types.StringNull()
	}
	if v, ok := obj["update-stats-interval"]; ok {
		_ = v
		if v != "" {
			m.UpdateStatsInterval = types.StringValue(v)
		} else {
			m.UpdateStatsInterval = types.StringNull()
		}
	} else {
		m.UpdateStatsInterval = types.StringNull()
	}
	if v, ok := obj["vht-basic-mcs"]; ok {
		_ = v
		if v != "" {
			m.VhtBasicMcs = types.StringValue(v)
		} else {
			m.VhtBasicMcs = types.StringNull()
		}
	} else {
		m.VhtBasicMcs = types.StringNull()
	}
	if v, ok := obj["vht-supported-mcs"]; ok {
		_ = v
		if v != "" {
			m.VhtSupportedMcs = types.StringValue(v)
		} else {
			m.VhtSupportedMcs = types.StringNull()
		}
	} else {
		m.VhtSupportedMcs = types.StringNull()
	}
	if v, ok := obj["wds-cost-range"]; ok {
		_ = v
		if v != "" {
			m.WdsCostRange = types.StringValue(v)
		} else {
			m.WdsCostRange = types.StringNull()
		}
	} else {
		m.WdsCostRange = types.StringNull()
	}
	if v, ok := obj["wds-default-bridge"]; ok {
		_ = v
		if v != "" {
			m.WdsDefaultBridge = types.StringValue(v)
		} else {
			m.WdsDefaultBridge = types.StringNull()
		}
	} else {
		m.WdsDefaultBridge = types.StringNull()
	}
	if v, ok := obj["wds-default-cost"]; ok {
		_ = v
		if v != "" {
			m.WdsDefaultCost = types.StringValue(v)
		} else {
			m.WdsDefaultCost = types.StringNull()
		}
	} else {
		m.WdsDefaultCost = types.StringNull()
	}
	if v, ok := obj["wds-ignore-ssid"]; ok {
		_ = v
		if v != "" {
			m.WdsIgnoreSsid = types.StringValue(v)
		} else {
			m.WdsIgnoreSsid = types.StringNull()
		}
	} else {
		m.WdsIgnoreSsid = types.StringNull()
	}
	if v, ok := obj["wds-mode"]; ok {
		_ = v
		if v != "" {
			m.WdsMode = types.StringValue(v)
		} else {
			m.WdsMode = types.StringNull()
		}
	} else {
		m.WdsMode = types.StringNull()
	}
	if v, ok := obj["wireless-protocol"]; ok {
		_ = v
		if v != "" {
			m.WirelessProtocol = types.StringValue(v)
		} else {
			m.WirelessProtocol = types.StringNull()
		}
	} else {
		m.WirelessProtocol = types.StringNull()
	}
	if v, ok := obj["wmm-support"]; ok {
		_ = v
		if v != "" {
			m.WmmSupport = types.StringValue(v)
		} else {
			m.WmmSupport = types.StringNull()
		}
	} else {
		m.WmmSupport = types.StringNull()
	}
	if v, ok := obj["wps-mode"]; ok {
		_ = v
		if v != "" {
			m.WpsMode = types.StringValue(v)
		} else {
			m.WpsMode = types.StringNull()
		}
	} else {
		m.WpsMode = types.StringNull()
	}
}
