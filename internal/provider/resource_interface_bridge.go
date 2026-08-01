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
	_ resource.Resource                = &InterfaceBridgeResource{}
	_ resource.ResourceWithImportState = &InterfaceBridgeResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceBridgeResource struct {
	reg *client.Registry
}

type InterfaceBridgeModel struct {
	ID                       types.String    `tfsdk:"id"`
	ForwardReservedAddresses types.String    `tfsdk:"forward_reserved_addresses"`
	Dhcpv6Snooping           boolStringValue `tfsdk:"dhcpv6_snooping"`
	Dhcpv6AgentRemoteId      types.String    `tfsdk:"dhcpv6_agent_remote_id"`
	Dhcpv6AgentCircuitId     types.String    `tfsdk:"dhcpv6_agent_circuit_id"`
	DhcpAgentRemoteId        types.String    `tfsdk:"dhcp_agent_remote_id"`
	DhcpAgentCircuitId       types.String    `tfsdk:"dhcp_agent_circuit_id"`
	ActiveRole               types.String    `tfsdk:"active_role"`
	AddDHCPOption82          types.Bool      `tfsdk:"add_dhcp_option_82"`
	AdminMAC                 types.String    `tfsdk:"admin_mac"`
	AdminMACAddress          types.String    `tfsdk:"admin_mac_address"`
	AgeingTime               types.String    `tfsdk:"ageing_time"`
	ARP                      types.String    `tfsdk:"arp"`
	ARPTimeout               types.String    `tfsdk:"arp_timeout"`
	AutoMAC                  types.Bool      `tfsdk:"auto_mac"`
	Comment                  types.String    `tfsdk:"comment"`
	DHCPSnooping             types.Bool      `tfsdk:"dhcp_snooping"`
	Disabled                 types.Bool      `tfsdk:"disabled"`
	Dumb                     types.String    `tfsdk:"dumb"`
	EtherType                types.String    `tfsdk:"ether_type"`
	FastForward              types.Bool      `tfsdk:"fast_forward"`
	ForwardDelay             types.String    `tfsdk:"forward_delay"`
	ForwardReserved          types.Bool      `tfsdk:"forward_reserved"`
	FpTxRxPacketRate         types.String    `tfsdk:"fp_tx_rx_packet_rate"`
	FpTxRxRate               types.String    `tfsdk:"fp_tx_rx_rate"`
	FrameTypes               types.String    `tfsdk:"frame_types"`
	Heartbeat                types.String    `tfsdk:"heartbeat"`
	Igmp                     types.String    `tfsdk:"igmp"`
	IgmpSnooping             types.Bool      `tfsdk:"igmp_snooping"`
	IgmpVersion              types.String    `tfsdk:"igmp_version"`
	IngressFiltering         types.Bool      `tfsdk:"ingress_filtering"`
	LastMemberInterval       types.String    `tfsdk:"last_member_interval"`
	LastMemberQueryCount     types.Int64     `tfsdk:"last_member_query_count"`
	MACAddress               types.String    `tfsdk:"mac_address"`
	MaxHops                  types.Int64     `tfsdk:"max_hops"`
	MaxLearnedEntries        types.String    `tfsdk:"max_learned_entries"`
	MaxMessageAge            types.String    `tfsdk:"max_message_age"`
	MembershipInterval       types.String    `tfsdk:"membership_interval"`
	MlagHeartbeat            types.String    `tfsdk:"mlag_heartbeat"`
	MlagPeerPort             types.String    `tfsdk:"mlag_peer_port"`
	MlagPriority             types.Int64     `tfsdk:"mlag_priority"`
	MldVersion               types.String    `tfsdk:"mld_version"`
	Mstp                     types.String    `tfsdk:"mstp"`
	MTU                      types.String    `tfsdk:"mtu"`
	MulticastQuerier         types.Bool      `tfsdk:"multicast_querier"`
	MulticastRouter          types.String    `tfsdk:"multicast_router"`
	Mvrp                     types.Bool      `tfsdk:"mvrp"`
	Name                     types.String    `tfsdk:"name"`
	PeerPort                 types.String    `tfsdk:"peer_port"`
	PortCostMode             types.String    `tfsdk:"port_cost_mode"`
	Priority                 types.Int64     `tfsdk:"priority"`
	ProtocolMode             types.String    `tfsdk:"protocol_mode"`
	Pvid                     types.Int64     `tfsdk:"pvid"`
	QuerierInterval          types.String    `tfsdk:"querier_interval"`
	QueryInterval            types.String    `tfsdk:"query_interval"`
	QueryResponseInterval    types.String    `tfsdk:"query_response_interval"`
	RaGuard                  types.Bool      `tfsdk:"ra_guard"`
	RegionName               types.String    `tfsdk:"region_name"`
	RegionRevision           types.Int64     `tfsdk:"region_revision"`
	StartupQueryCount        types.Int64     `tfsdk:"startup_query_count"`
	StartupQueryInterval     types.String    `tfsdk:"startup_query_interval"`
	State                    types.String    `tfsdk:"state"`
	Status                   types.Int64     `tfsdk:"status"`
	TransmitHoldCount        types.Int64     `tfsdk:"transmit_hold_count"`
	TxRxPacketRate           types.String    `tfsdk:"tx_rx_packet_rate"`
	TxRxRate                 types.String    `tfsdk:"tx_rx_rate"`
	Type                     types.String    `tfsdk:"type"`
	VLAN                     types.String    `tfsdk:"vlan"`
	VLANFiltering            types.Bool      `tfsdk:"vlan_filtering"`
	Router                   types.String    `tfsdk:"router"`
}

func NewInterfaceBridgeResource() resource.Resource { return &InterfaceBridgeResource{} }

func (r *InterfaceBridgeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_bridge"
}

func (r *InterfaceBridgeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceBridgeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/bridge`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"forward_reserved_addresses": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `forward-reserved-addresses`.",
			},
			"dhcpv6_snooping": schema.StringAttribute{
				CustomType:  boolStringType{},
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dhcpv6-snooping`.",
			},
			"dhcpv6_agent_remote_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dhcpv6-agent-remote-id`.",
			},
			"dhcpv6_agent_circuit_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dhcpv6-agent-circuit-id`.",
			},
			"dhcp_agent_remote_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dhcp-agent-remote-id`.",
			},
			"dhcp_agent_circuit_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dhcp-agent-circuit-id`.",
			},
			"active_role": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"primary", "secondary"}...)},
			},
			"add_dhcp_option_82": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"admin_mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"admin_mac_address": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"ageing_time": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
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
			"auto_mac": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"dhcp_snooping": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"dumb": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"ether_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"0x8100", "0x88a8", "0x9100"}...)},
			},
			"fast_forward": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"forward_delay": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"forward_reserved": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"fp_tx_rx_packet_rate": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"fp_tx_rx_rate": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"frame_types": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"admit-all", "admit-only-vlan-tagged", "admit-only-untagged-and-priority-tagged"}...)},
			},
			"heartbeat": schema.StringAttribute{
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"igmp": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"igmp_snooping": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"igmp_version": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"2", "3"}...)},
			},
			"ingress_filtering": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"last_member_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"last_member_query_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mac_address": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"max_hops": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"max_learned_entries": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"unlimited", "auto"}...)},
			},
			"max_message_age": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"membership_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mlag_heartbeat": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"mlag_peer_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mlag_priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mld_version": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "1", "2"}...)},
			},
			"mstp": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Bridge MTU. A number, or `auto` (the default) to follow the smallest port MTU.",
			},
			"multicast_querier": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"multicast_router": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"disabled", "temporary-query", "permanent"}...)},
			},
			"mvrp": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"peer_port": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"port_cost_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"short", "long"}...)},
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"protocol_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"none", "stp", "rstp", "mstp"}...)},
			},
			"pvid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"querier_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"query_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"query_response_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ra_guard": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"region_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"region_revision": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"startup_query_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"startup_query_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"state": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"status": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"transmit_hold_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_packet_rate": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_rx_rate": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"vlan": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"vlan_filtering": schema.BoolAttribute{
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

func (r *InterfaceBridgeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceBridgeModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AdminMAC.IsNull() || plan.AdminMAC.IsUnknown()) {
		body["admin-mac"] = plan.AdminMAC.ValueString()
	}
	if !(plan.AgeingTime.IsNull() || plan.AgeingTime.IsUnknown()) {
		body["ageing-time"] = plan.AgeingTime.ValueString()
	}
	if !(plan.ARP.IsNull() || plan.ARP.IsUnknown()) {
		body["arp"] = plan.ARP.ValueString()
	}
	if !(plan.ARPTimeout.IsNull() || plan.ARPTimeout.IsUnknown()) {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !(plan.AutoMAC.IsNull() || plan.AutoMAC.IsUnknown()) {
		body["auto-mac"] = client.FormatBool(plan.AutoMAC.ValueBool())
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.DHCPSnooping.IsNull() || plan.DHCPSnooping.IsUnknown()) {
		body["dhcp-snooping"] = client.FormatBool(plan.DHCPSnooping.ValueBool())
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.EtherType.IsNull() || plan.EtherType.IsUnknown()) {
		body["ether-type"] = plan.EtherType.ValueString()
	}
	if !(plan.FastForward.IsNull() || plan.FastForward.IsUnknown()) {
		body["fast-forward"] = client.FormatBool(plan.FastForward.ValueBool())
	}
	if !(plan.ForwardDelay.IsNull() || plan.ForwardDelay.IsUnknown()) {
		body["forward-delay"] = plan.ForwardDelay.ValueString()
	}
	if !(plan.FrameTypes.IsNull() || plan.FrameTypes.IsUnknown()) {
		body["frame-types"] = plan.FrameTypes.ValueString()
	}
	if !(plan.IgmpSnooping.IsNull() || plan.IgmpSnooping.IsUnknown()) {
		body["igmp-snooping"] = client.FormatBool(plan.IgmpSnooping.ValueBool())
	}
	if !(plan.IgmpVersion.IsNull() || plan.IgmpVersion.IsUnknown()) {
		body["igmp-version"] = plan.IgmpVersion.ValueString()
	}
	if !(plan.IngressFiltering.IsNull() || plan.IngressFiltering.IsUnknown()) {
		body["ingress-filtering"] = client.FormatBool(plan.IngressFiltering.ValueBool())
	}
	if !(plan.LastMemberInterval.IsNull() || plan.LastMemberInterval.IsUnknown()) {
		body["last-member-interval"] = plan.LastMemberInterval.ValueString()
	}
	if !(plan.LastMemberQueryCount.IsNull() || plan.LastMemberQueryCount.IsUnknown()) {
		body["last-member-query-count"] = client.FormatInt64(plan.LastMemberQueryCount.ValueInt64())
	}
	if !(plan.MaxHops.IsNull() || plan.MaxHops.IsUnknown()) {
		body["max-hops"] = client.FormatInt64(plan.MaxHops.ValueInt64())
	}
	if !(plan.MaxLearnedEntries.IsNull() || plan.MaxLearnedEntries.IsUnknown()) {
		body["max-learned-entries"] = plan.MaxLearnedEntries.ValueString()
	}
	if !(plan.MaxMessageAge.IsNull() || plan.MaxMessageAge.IsUnknown()) {
		body["max-message-age"] = plan.MaxMessageAge.ValueString()
	}
	if !(plan.MembershipInterval.IsNull() || plan.MembershipInterval.IsUnknown()) {
		body["membership-interval"] = plan.MembershipInterval.ValueString()
	}
	if !(plan.MlagHeartbeat.IsNull() || plan.MlagHeartbeat.IsUnknown()) {
		body["mlag-heartbeat"] = plan.MlagHeartbeat.ValueString()
	}
	if !(plan.MlagPeerPort.IsNull() || plan.MlagPeerPort.IsUnknown()) {
		body["mlag-peer-port"] = plan.MlagPeerPort.ValueString()
	}
	if !(plan.MlagPriority.IsNull() || plan.MlagPriority.IsUnknown()) {
		body["mlag-priority"] = client.FormatInt64(plan.MlagPriority.ValueInt64())
	}
	if !(plan.MldVersion.IsNull() || plan.MldVersion.IsUnknown()) {
		body["mld-version"] = plan.MldVersion.ValueString()
	}
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !(plan.MulticastQuerier.IsNull() || plan.MulticastQuerier.IsUnknown()) {
		body["multicast-querier"] = client.FormatBool(plan.MulticastQuerier.ValueBool())
	}
	if !(plan.MulticastRouter.IsNull() || plan.MulticastRouter.IsUnknown()) {
		body["multicast-router"] = plan.MulticastRouter.ValueString()
	}
	if !(plan.Mvrp.IsNull() || plan.Mvrp.IsUnknown()) {
		body["mvrp"] = client.FormatBool(plan.Mvrp.ValueBool())
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.PortCostMode.IsNull() || plan.PortCostMode.IsUnknown()) {
		body["port-cost-mode"] = plan.PortCostMode.ValueString()
	}
	if !(plan.Priority.IsNull() || plan.Priority.IsUnknown()) {
		body["priority"] = client.FormatInt64(plan.Priority.ValueInt64())
	}
	if !(plan.ProtocolMode.IsNull() || plan.ProtocolMode.IsUnknown()) {
		body["protocol-mode"] = plan.ProtocolMode.ValueString()
	}
	if !(plan.Pvid.IsNull() || plan.Pvid.IsUnknown()) {
		body["pvid"] = client.FormatInt64(plan.Pvid.ValueInt64())
	}
	if !(plan.QuerierInterval.IsNull() || plan.QuerierInterval.IsUnknown()) {
		body["querier-interval"] = plan.QuerierInterval.ValueString()
	}
	if !(plan.QueryInterval.IsNull() || plan.QueryInterval.IsUnknown()) {
		body["query-interval"] = plan.QueryInterval.ValueString()
	}
	if !(plan.QueryResponseInterval.IsNull() || plan.QueryResponseInterval.IsUnknown()) {
		body["query-response-interval"] = plan.QueryResponseInterval.ValueString()
	}
	if !(plan.RaGuard.IsNull() || plan.RaGuard.IsUnknown()) {
		body["ra-guard"] = client.FormatBool(plan.RaGuard.ValueBool())
	}
	if !(plan.RegionName.IsNull() || plan.RegionName.IsUnknown()) {
		body["region-name"] = plan.RegionName.ValueString()
	}
	if !(plan.RegionRevision.IsNull() || plan.RegionRevision.IsUnknown()) {
		body["region-revision"] = client.FormatInt64(plan.RegionRevision.ValueInt64())
	}
	if !(plan.StartupQueryCount.IsNull() || plan.StartupQueryCount.IsUnknown()) {
		body["startup-query-count"] = client.FormatInt64(plan.StartupQueryCount.ValueInt64())
	}
	if !(plan.StartupQueryInterval.IsNull() || plan.StartupQueryInterval.IsUnknown()) {
		body["startup-query-interval"] = plan.StartupQueryInterval.ValueString()
	}
	if !(plan.TransmitHoldCount.IsNull() || plan.TransmitHoldCount.IsUnknown()) {
		body["transmit-hold-count"] = client.FormatInt64(plan.TransmitHoldCount.ValueInt64())
	}
	if !(plan.VLANFiltering.IsNull() || plan.VLANFiltering.IsUnknown()) {
		body["vlan-filtering"] = client.FormatBool(plan.VLANFiltering.ValueBool())
	}
	if !(plan.DhcpAgentCircuitId.IsNull() || plan.DhcpAgentCircuitId.IsUnknown()) {
		body["dhcp-agent-circuit-id"] = plan.DhcpAgentCircuitId.ValueString()
	}
	if !(plan.DhcpAgentRemoteId.IsNull() || plan.DhcpAgentRemoteId.IsUnknown()) {
		body["dhcp-agent-remote-id"] = plan.DhcpAgentRemoteId.ValueString()
	}
	if !(plan.Dhcpv6AgentCircuitId.IsNull() || plan.Dhcpv6AgentCircuitId.IsUnknown()) {
		body["dhcpv6-agent-circuit-id"] = plan.Dhcpv6AgentCircuitId.ValueString()
	}
	if !(plan.Dhcpv6AgentRemoteId.IsNull() || plan.Dhcpv6AgentRemoteId.IsUnknown()) {
		body["dhcpv6-agent-remote-id"] = plan.Dhcpv6AgentRemoteId.ValueString()
	}
	if !(plan.Dhcpv6Snooping.IsNull() || plan.Dhcpv6Snooping.IsUnknown()) {
		body["dhcpv6-snooping"] = plan.Dhcpv6Snooping.ValueString()
	}
	if !(plan.ForwardReservedAddresses.IsNull() || plan.ForwardReservedAddresses.IsUnknown()) {
		body["forward-reserved-addresses"] = plan.ForwardReservedAddresses.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/bridge", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/bridge failed", err.Error())
		return
	}
	interfaceBridgeApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceBridgeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceBridgeModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/bridge", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/bridge failed", err.Error())
		return
	}
	interfaceBridgeApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceBridgeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceBridgeModel
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
	if !plan.AdminMAC.Equal(state.AdminMAC) && !plan.AdminMAC.IsUnknown() {
		body["admin-mac"] = plan.AdminMAC.ValueString()
	}
	if !plan.AgeingTime.Equal(state.AgeingTime) && !plan.AgeingTime.IsUnknown() {
		body["ageing-time"] = plan.AgeingTime.ValueString()
	}
	if !plan.ARP.Equal(state.ARP) && !plan.ARP.IsUnknown() {
		body["arp"] = plan.ARP.ValueString()
	}
	if !plan.ARPTimeout.Equal(state.ARPTimeout) && !plan.ARPTimeout.IsUnknown() {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !plan.AutoMAC.Equal(state.AutoMAC) && !plan.AutoMAC.IsUnknown() {
		body["auto-mac"] = client.FormatBool(plan.AutoMAC.ValueBool())
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.DHCPSnooping.Equal(state.DHCPSnooping) && !plan.DHCPSnooping.IsUnknown() {
		body["dhcp-snooping"] = client.FormatBool(plan.DHCPSnooping.ValueBool())
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.EtherType.Equal(state.EtherType) && !plan.EtherType.IsUnknown() {
		body["ether-type"] = plan.EtherType.ValueString()
	}
	if !plan.FastForward.Equal(state.FastForward) && !plan.FastForward.IsUnknown() {
		body["fast-forward"] = client.FormatBool(plan.FastForward.ValueBool())
	}
	if !plan.ForwardDelay.Equal(state.ForwardDelay) && !plan.ForwardDelay.IsUnknown() {
		body["forward-delay"] = plan.ForwardDelay.ValueString()
	}
	if !plan.FrameTypes.Equal(state.FrameTypes) && !plan.FrameTypes.IsUnknown() {
		body["frame-types"] = plan.FrameTypes.ValueString()
	}
	if !plan.IgmpSnooping.Equal(state.IgmpSnooping) && !plan.IgmpSnooping.IsUnknown() {
		body["igmp-snooping"] = client.FormatBool(plan.IgmpSnooping.ValueBool())
	}
	if !plan.IgmpVersion.Equal(state.IgmpVersion) && !plan.IgmpVersion.IsUnknown() {
		body["igmp-version"] = plan.IgmpVersion.ValueString()
	}
	if !plan.IngressFiltering.Equal(state.IngressFiltering) && !plan.IngressFiltering.IsUnknown() {
		body["ingress-filtering"] = client.FormatBool(plan.IngressFiltering.ValueBool())
	}
	if !plan.LastMemberInterval.Equal(state.LastMemberInterval) && !plan.LastMemberInterval.IsUnknown() {
		body["last-member-interval"] = plan.LastMemberInterval.ValueString()
	}
	if !plan.LastMemberQueryCount.Equal(state.LastMemberQueryCount) && !plan.LastMemberQueryCount.IsUnknown() {
		body["last-member-query-count"] = client.FormatInt64(plan.LastMemberQueryCount.ValueInt64())
	}
	if !plan.MaxHops.Equal(state.MaxHops) && !plan.MaxHops.IsUnknown() {
		body["max-hops"] = client.FormatInt64(plan.MaxHops.ValueInt64())
	}
	if !plan.MaxLearnedEntries.Equal(state.MaxLearnedEntries) && !plan.MaxLearnedEntries.IsUnknown() {
		body["max-learned-entries"] = plan.MaxLearnedEntries.ValueString()
	}
	if !plan.MaxMessageAge.Equal(state.MaxMessageAge) && !plan.MaxMessageAge.IsUnknown() {
		body["max-message-age"] = plan.MaxMessageAge.ValueString()
	}
	if !plan.MembershipInterval.Equal(state.MembershipInterval) && !plan.MembershipInterval.IsUnknown() {
		body["membership-interval"] = plan.MembershipInterval.ValueString()
	}
	if !plan.MlagHeartbeat.Equal(state.MlagHeartbeat) && !plan.MlagHeartbeat.IsUnknown() {
		body["mlag-heartbeat"] = plan.MlagHeartbeat.ValueString()
	}
	if !plan.MlagPeerPort.Equal(state.MlagPeerPort) && !plan.MlagPeerPort.IsUnknown() {
		body["mlag-peer-port"] = plan.MlagPeerPort.ValueString()
	}
	if !plan.MlagPriority.Equal(state.MlagPriority) && !plan.MlagPriority.IsUnknown() {
		body["mlag-priority"] = client.FormatInt64(plan.MlagPriority.ValueInt64())
	}
	if !plan.MldVersion.Equal(state.MldVersion) && !plan.MldVersion.IsUnknown() {
		body["mld-version"] = plan.MldVersion.ValueString()
	}
	if !plan.MTU.Equal(state.MTU) && !plan.MTU.IsUnknown() {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !plan.MulticastQuerier.Equal(state.MulticastQuerier) && !plan.MulticastQuerier.IsUnknown() {
		body["multicast-querier"] = client.FormatBool(plan.MulticastQuerier.ValueBool())
	}
	if !plan.MulticastRouter.Equal(state.MulticastRouter) && !plan.MulticastRouter.IsUnknown() {
		body["multicast-router"] = plan.MulticastRouter.ValueString()
	}
	if !plan.Mvrp.Equal(state.Mvrp) && !plan.Mvrp.IsUnknown() {
		body["mvrp"] = client.FormatBool(plan.Mvrp.ValueBool())
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.PortCostMode.Equal(state.PortCostMode) && !plan.PortCostMode.IsUnknown() {
		body["port-cost-mode"] = plan.PortCostMode.ValueString()
	}
	if !plan.Priority.Equal(state.Priority) && !plan.Priority.IsUnknown() {
		body["priority"] = client.FormatInt64(plan.Priority.ValueInt64())
	}
	if !plan.ProtocolMode.Equal(state.ProtocolMode) && !plan.ProtocolMode.IsUnknown() {
		body["protocol-mode"] = plan.ProtocolMode.ValueString()
	}
	if !plan.Pvid.Equal(state.Pvid) && !plan.Pvid.IsUnknown() {
		body["pvid"] = client.FormatInt64(plan.Pvid.ValueInt64())
	}
	if !plan.QuerierInterval.Equal(state.QuerierInterval) && !plan.QuerierInterval.IsUnknown() {
		body["querier-interval"] = plan.QuerierInterval.ValueString()
	}
	if !plan.QueryInterval.Equal(state.QueryInterval) && !plan.QueryInterval.IsUnknown() {
		body["query-interval"] = plan.QueryInterval.ValueString()
	}
	if !plan.QueryResponseInterval.Equal(state.QueryResponseInterval) && !plan.QueryResponseInterval.IsUnknown() {
		body["query-response-interval"] = plan.QueryResponseInterval.ValueString()
	}
	if !plan.RaGuard.Equal(state.RaGuard) && !plan.RaGuard.IsUnknown() {
		body["ra-guard"] = client.FormatBool(plan.RaGuard.ValueBool())
	}
	if !plan.RegionName.Equal(state.RegionName) && !plan.RegionName.IsUnknown() {
		body["region-name"] = plan.RegionName.ValueString()
	}
	if !plan.RegionRevision.Equal(state.RegionRevision) && !plan.RegionRevision.IsUnknown() {
		body["region-revision"] = client.FormatInt64(plan.RegionRevision.ValueInt64())
	}
	if !plan.StartupQueryCount.Equal(state.StartupQueryCount) && !plan.StartupQueryCount.IsUnknown() {
		body["startup-query-count"] = client.FormatInt64(plan.StartupQueryCount.ValueInt64())
	}
	if !plan.StartupQueryInterval.Equal(state.StartupQueryInterval) && !plan.StartupQueryInterval.IsUnknown() {
		body["startup-query-interval"] = plan.StartupQueryInterval.ValueString()
	}
	if !plan.TransmitHoldCount.Equal(state.TransmitHoldCount) && !plan.TransmitHoldCount.IsUnknown() {
		body["transmit-hold-count"] = client.FormatInt64(plan.TransmitHoldCount.ValueInt64())
	}
	if !plan.VLANFiltering.Equal(state.VLANFiltering) && !plan.VLANFiltering.IsUnknown() {
		body["vlan-filtering"] = client.FormatBool(plan.VLANFiltering.ValueBool())
	}
	if !plan.DhcpAgentCircuitId.Equal(state.DhcpAgentCircuitId) && !plan.DhcpAgentCircuitId.IsUnknown() {
		body["dhcp-agent-circuit-id"] = plan.DhcpAgentCircuitId.ValueString()
	}
	if !plan.DhcpAgentRemoteId.Equal(state.DhcpAgentRemoteId) && !plan.DhcpAgentRemoteId.IsUnknown() {
		body["dhcp-agent-remote-id"] = plan.DhcpAgentRemoteId.ValueString()
	}
	if !plan.Dhcpv6AgentCircuitId.Equal(state.Dhcpv6AgentCircuitId) && !plan.Dhcpv6AgentCircuitId.IsUnknown() {
		body["dhcpv6-agent-circuit-id"] = plan.Dhcpv6AgentCircuitId.ValueString()
	}
	if !plan.Dhcpv6AgentRemoteId.Equal(state.Dhcpv6AgentRemoteId) && !plan.Dhcpv6AgentRemoteId.IsUnknown() {
		body["dhcpv6-agent-remote-id"] = plan.Dhcpv6AgentRemoteId.ValueString()
	}
	if !plan.Dhcpv6Snooping.Equal(state.Dhcpv6Snooping) && !plan.Dhcpv6Snooping.IsUnknown() {
		body["dhcpv6-snooping"] = plan.Dhcpv6Snooping.ValueString()
	}
	if !plan.ForwardReservedAddresses.Equal(state.ForwardReservedAddresses) && !plan.ForwardReservedAddresses.IsUnknown() {
		body["forward-reserved-addresses"] = plan.ForwardReservedAddresses.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/bridge", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/bridge failed", err.Error())
			return
		}
		interfaceBridgeApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceBridgeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceBridgeModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/bridge", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/bridge failed", err.Error())
	}
}

func (r *InterfaceBridgeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceBridgeLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/bridge matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceBridgeLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceBridgeLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/bridge", id)
}

func interfaceBridgeApply(ctx context.Context, obj client.Object, m *InterfaceBridgeModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["forward-reserved-addresses"]; ok && v != "" {
		m.ForwardReservedAddresses = types.StringValue(v)
	} else {
		m.ForwardReservedAddresses = types.StringNull()
	}
	if v, ok := obj["dhcpv6-snooping"]; ok && v != "" {
		m.Dhcpv6Snooping = newBoolStringValue(v)
	}
	if v, ok := obj["dhcpv6-agent-remote-id"]; ok && v != "" {
		m.Dhcpv6AgentRemoteId = types.StringValue(v)
	} else {
		m.Dhcpv6AgentRemoteId = types.StringNull()
	}
	if v, ok := obj["dhcpv6-agent-circuit-id"]; ok && v != "" {
		m.Dhcpv6AgentCircuitId = types.StringValue(v)
	} else {
		m.Dhcpv6AgentCircuitId = types.StringNull()
	}
	if v, ok := obj["dhcp-agent-remote-id"]; ok && v != "" {
		m.DhcpAgentRemoteId = types.StringValue(v)
	} else {
		m.DhcpAgentRemoteId = types.StringNull()
	}
	if v, ok := obj["dhcp-agent-circuit-id"]; ok && v != "" {
		m.DhcpAgentCircuitId = types.StringValue(v)
	} else {
		m.DhcpAgentCircuitId = types.StringNull()
	}
	if v, ok := obj["active-role"]; ok {
		_ = v
		if v != "" {
			m.ActiveRole = types.StringValue(v)
		} else {
			m.ActiveRole = types.StringNull()
		}
	} else {
		m.ActiveRole = types.StringNull()
	}
	if v, ok := obj["add-dhcp-option-82"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.AddDHCPOption82 = types.BoolValue(b)
		} else {
			m.AddDHCPOption82 = types.BoolNull()
		}
	}
	if v, ok := obj["admin-mac"]; ok {
		_ = v
		if v != "" {
			m.AdminMAC = types.StringValue(v)
		} else {
			m.AdminMAC = types.StringNull()
		}
	} else {
		m.AdminMAC = types.StringNull()
	}
	if v, ok := obj["admin-mac-address"]; ok {
		_ = v
		if v != "" {
			m.AdminMACAddress = types.StringValue(v)
		} else {
			m.AdminMACAddress = types.StringNull()
		}
	} else {
		m.AdminMACAddress = types.StringNull()
	}
	if v, ok := obj["ageing-time"]; ok {
		_ = v
		if v != "" {
			m.AgeingTime = types.StringValue(v)
		} else {
			m.AgeingTime = types.StringNull()
		}
	} else {
		m.AgeingTime = types.StringNull()
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
	if v, ok := obj["auto-mac"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.AutoMAC = types.BoolValue(b)
		} else {
			m.AutoMAC = types.BoolNull()
		}
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
	if v, ok := obj["dhcp-snooping"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.DHCPSnooping = types.BoolValue(b)
		} else {
			m.DHCPSnooping = types.BoolNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["dumb"]; ok {
		_ = v
		if v != "" {
			m.Dumb = types.StringValue(v)
		} else {
			m.Dumb = types.StringNull()
		}
	} else {
		m.Dumb = types.StringNull()
	}
	if v, ok := obj["ether-type"]; ok {
		_ = v
		if v != "" {
			m.EtherType = types.StringValue(v)
		} else {
			m.EtherType = types.StringNull()
		}
	} else {
		m.EtherType = types.StringNull()
	}
	if v, ok := obj["fast-forward"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.FastForward = types.BoolValue(b)
		} else {
			m.FastForward = types.BoolNull()
		}
	}
	if v, ok := obj["forward-delay"]; ok {
		_ = v
		if v != "" {
			m.ForwardDelay = types.StringValue(v)
		} else {
			m.ForwardDelay = types.StringNull()
		}
	} else {
		m.ForwardDelay = types.StringNull()
	}
	if v, ok := obj["forward-reserved"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.ForwardReserved = types.BoolValue(b)
		} else {
			m.ForwardReserved = types.BoolNull()
		}
	}
	if v, ok := obj["fp-tx-rx-packet-rate"]; ok {
		_ = v
		if v != "" {
			m.FpTxRxPacketRate = types.StringValue(v)
		} else {
			m.FpTxRxPacketRate = types.StringNull()
		}
	} else {
		m.FpTxRxPacketRate = types.StringNull()
	}
	if v, ok := obj["fp-tx-rx-rate"]; ok {
		_ = v
		if v != "" {
			m.FpTxRxRate = types.StringValue(v)
		} else {
			m.FpTxRxRate = types.StringNull()
		}
	} else {
		m.FpTxRxRate = types.StringNull()
	}
	if v, ok := obj["frame-types"]; ok {
		_ = v
		if v != "" {
			m.FrameTypes = types.StringValue(v)
		} else {
			m.FrameTypes = types.StringNull()
		}
	} else {
		m.FrameTypes = types.StringNull()
	}
	if v, ok := obj["heartbeat"]; ok {
		_ = v
		if v != "" {
			m.Heartbeat = types.StringValue(v)
		} else {
			m.Heartbeat = types.StringNull()
		}
	} else {
		m.Heartbeat = types.StringNull()
	}
	if v, ok := obj["igmp"]; ok {
		_ = v
		if v != "" {
			m.Igmp = types.StringValue(v)
		} else {
			m.Igmp = types.StringNull()
		}
	} else {
		m.Igmp = types.StringNull()
	}
	if v, ok := obj["igmp-snooping"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.IgmpSnooping = types.BoolValue(b)
		} else {
			m.IgmpSnooping = types.BoolNull()
		}
	}
	if v, ok := obj["igmp-version"]; ok {
		_ = v
		if v != "" {
			m.IgmpVersion = types.StringValue(v)
		} else {
			m.IgmpVersion = types.StringNull()
		}
	} else {
		m.IgmpVersion = types.StringNull()
	}
	if v, ok := obj["ingress-filtering"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.IngressFiltering = types.BoolValue(b)
		} else {
			m.IngressFiltering = types.BoolNull()
		}
	}
	if v, ok := obj["last-member-interval"]; ok {
		_ = v
		if v != "" {
			m.LastMemberInterval = types.StringValue(v)
		} else {
			m.LastMemberInterval = types.StringNull()
		}
	} else {
		m.LastMemberInterval = types.StringNull()
	}
	if v, ok := obj["last-member-query-count"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.LastMemberQueryCount = types.Int64Value(n)
		} else {
			m.LastMemberQueryCount = types.Int64Null()
		}
	} else {
		m.LastMemberQueryCount = types.Int64Null()
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
	if v, ok := obj["max-hops"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MaxHops = types.Int64Value(n)
		} else {
			m.MaxHops = types.Int64Null()
		}
	} else {
		m.MaxHops = types.Int64Null()
	}
	if v, ok := obj["max-learned-entries"]; ok {
		_ = v
		if v != "" {
			m.MaxLearnedEntries = types.StringValue(v)
		} else {
			m.MaxLearnedEntries = types.StringNull()
		}
	} else {
		m.MaxLearnedEntries = types.StringNull()
	}
	if v, ok := obj["max-message-age"]; ok {
		_ = v
		if v != "" {
			m.MaxMessageAge = types.StringValue(v)
		} else {
			m.MaxMessageAge = types.StringNull()
		}
	} else {
		m.MaxMessageAge = types.StringNull()
	}
	if v, ok := obj["membership-interval"]; ok {
		_ = v
		if v != "" {
			m.MembershipInterval = types.StringValue(v)
		} else {
			m.MembershipInterval = types.StringNull()
		}
	} else {
		m.MembershipInterval = types.StringNull()
	}
	if v, ok := obj["mlag-heartbeat"]; ok {
		_ = v
		if v != "" {
			m.MlagHeartbeat = types.StringValue(v)
		} else {
			m.MlagHeartbeat = types.StringNull()
		}
	} else {
		m.MlagHeartbeat = types.StringNull()
	}
	if v, ok := obj["mlag-peer-port"]; ok {
		_ = v
		if v != "" {
			m.MlagPeerPort = types.StringValue(v)
		} else {
			m.MlagPeerPort = types.StringNull()
		}
	} else {
		m.MlagPeerPort = types.StringNull()
	}
	if v, ok := obj["mlag-priority"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MlagPriority = types.Int64Value(n)
		} else {
			m.MlagPriority = types.Int64Null()
		}
	} else {
		m.MlagPriority = types.Int64Null()
	}
	if v, ok := obj["mld-version"]; ok {
		_ = v
		if v != "" {
			m.MldVersion = types.StringValue(v)
		} else {
			m.MldVersion = types.StringNull()
		}
	} else {
		m.MldVersion = types.StringNull()
	}
	if v, ok := obj["mstp"]; ok {
		_ = v
		if v != "" {
			m.Mstp = types.StringValue(v)
		} else {
			m.Mstp = types.StringNull()
		}
	} else {
		m.Mstp = types.StringNull()
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
	if v, ok := obj["multicast-querier"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.MulticastQuerier = types.BoolValue(b)
		} else {
			m.MulticastQuerier = types.BoolNull()
		}
	}
	if v, ok := obj["multicast-router"]; ok {
		_ = v
		if v != "" {
			m.MulticastRouter = types.StringValue(v)
		} else {
			m.MulticastRouter = types.StringNull()
		}
	} else {
		m.MulticastRouter = types.StringNull()
	}
	if v, ok := obj["mvrp"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Mvrp = types.BoolValue(b)
		} else {
			m.Mvrp = types.BoolNull()
		}
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
	if v, ok := obj["peer-port"]; ok {
		_ = v
		if v != "" {
			m.PeerPort = types.StringValue(v)
		} else {
			m.PeerPort = types.StringNull()
		}
	} else {
		m.PeerPort = types.StringNull()
	}
	if v, ok := obj["port-cost-mode"]; ok {
		_ = v
		if v != "" {
			m.PortCostMode = types.StringValue(v)
		} else {
			m.PortCostMode = types.StringNull()
		}
	} else {
		m.PortCostMode = types.StringNull()
	}
	if v, ok := obj["priority"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Priority = types.Int64Value(n)
		} else {
			m.Priority = types.Int64Null()
		}
	} else {
		m.Priority = types.Int64Null()
	}
	if v, ok := obj["protocol-mode"]; ok {
		_ = v
		if v != "" {
			m.ProtocolMode = types.StringValue(v)
		} else {
			m.ProtocolMode = types.StringNull()
		}
	} else {
		m.ProtocolMode = types.StringNull()
	}
	if v, ok := obj["pvid"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Pvid = types.Int64Value(n)
		} else {
			m.Pvid = types.Int64Null()
		}
	} else {
		m.Pvid = types.Int64Null()
	}
	if v, ok := obj["querier-interval"]; ok {
		_ = v
		if v != "" {
			m.QuerierInterval = types.StringValue(v)
		} else {
			m.QuerierInterval = types.StringNull()
		}
	} else {
		m.QuerierInterval = types.StringNull()
	}
	if v, ok := obj["query-interval"]; ok {
		_ = v
		if v != "" {
			m.QueryInterval = types.StringValue(v)
		} else {
			m.QueryInterval = types.StringNull()
		}
	} else {
		m.QueryInterval = types.StringNull()
	}
	if v, ok := obj["query-response-interval"]; ok {
		_ = v
		if v != "" {
			m.QueryResponseInterval = types.StringValue(v)
		} else {
			m.QueryResponseInterval = types.StringNull()
		}
	} else {
		m.QueryResponseInterval = types.StringNull()
	}
	if v, ok := obj["ra-guard"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.RaGuard = types.BoolValue(b)
		} else {
			m.RaGuard = types.BoolNull()
		}
	}
	if v, ok := obj["region-name"]; ok {
		_ = v
		if v != "" {
			m.RegionName = types.StringValue(v)
		} else {
			m.RegionName = types.StringNull()
		}
	} else {
		m.RegionName = types.StringNull()
	}
	if v, ok := obj["region-revision"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RegionRevision = types.Int64Value(n)
		} else {
			m.RegionRevision = types.Int64Null()
		}
	} else {
		m.RegionRevision = types.Int64Null()
	}
	if v, ok := obj["startup-query-count"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.StartupQueryCount = types.Int64Value(n)
		} else {
			m.StartupQueryCount = types.Int64Null()
		}
	} else {
		m.StartupQueryCount = types.Int64Null()
	}
	if v, ok := obj["startup-query-interval"]; ok {
		_ = v
		if v != "" {
			m.StartupQueryInterval = types.StringValue(v)
		} else {
			m.StartupQueryInterval = types.StringNull()
		}
	} else {
		m.StartupQueryInterval = types.StringNull()
	}
	if v, ok := obj["state"]; ok {
		_ = v
		if v != "" {
			m.State = types.StringValue(v)
		} else {
			m.State = types.StringNull()
		}
	} else {
		m.State = types.StringNull()
	}
	if v, ok := obj["status"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Status = types.Int64Value(n)
		} else {
			m.Status = types.Int64Null()
		}
	} else {
		m.Status = types.Int64Null()
	}
	if v, ok := obj["transmit-hold-count"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TransmitHoldCount = types.Int64Value(n)
		} else {
			m.TransmitHoldCount = types.Int64Null()
		}
	} else {
		m.TransmitHoldCount = types.Int64Null()
	}
	if v, ok := obj["tx-rx-packet-rate"]; ok {
		_ = v
		if v != "" {
			m.TxRxPacketRate = types.StringValue(v)
		} else {
			m.TxRxPacketRate = types.StringNull()
		}
	} else {
		m.TxRxPacketRate = types.StringNull()
	}
	if v, ok := obj["tx-rx-rate"]; ok {
		_ = v
		if v != "" {
			m.TxRxRate = types.StringValue(v)
		} else {
			m.TxRxRate = types.StringNull()
		}
	} else {
		m.TxRxRate = types.StringNull()
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
	if v, ok := obj["vlan"]; ok {
		_ = v
		if v != "" {
			m.VLAN = types.StringValue(v)
		} else {
			m.VLAN = types.StringNull()
		}
	} else {
		m.VLAN = types.StringNull()
	}
	if v, ok := obj["vlan-filtering"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.VLANFiltering = types.BoolValue(b)
		} else {
			m.VLANFiltering = types.BoolNull()
		}
	}
}
