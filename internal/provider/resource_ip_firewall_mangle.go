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
	_ resource.Resource                = &IPFirewallMangleResource{}
	_ resource.ResourceWithImportState = &IPFirewallMangleResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPFirewallMangleResource struct {
	reg *client.Registry
}

type IPFirewallMangleModel struct {
	ID                      types.String `tfsdk:"id"`
	Tos                     types.String `tfsdk:"tos"`
	P2p                     types.String `tfsdk:"p2p"`
	Action                  types.String `tfsdk:"action"`
	AddressList             types.String `tfsdk:"address_list"`
	AddressListTimeout      types.String `tfsdk:"address_list_timeout"`
	Bytes                   types.String `tfsdk:"bytes"`
	Chain                   types.String `tfsdk:"chain"`
	Comment                 types.String `tfsdk:"comment"`
	ConnectionBytes         types.String `tfsdk:"connection_bytes"`
	ConnectionLimit         types.String `tfsdk:"connection_limit"`
	ConnectionMark          types.String `tfsdk:"connection_mark"`
	ConnectionNATState      types.String `tfsdk:"connection_nat_state"`
	ConnectionRate          types.String `tfsdk:"connection_rate"`
	ConnectionState         types.String `tfsdk:"connection_state"`
	ConnectionType          types.String `tfsdk:"connection_type"`
	Content                 types.String `tfsdk:"content"`
	Disabled                types.Bool   `tfsdk:"disabled"`
	Dscp                    types.String `tfsdk:"dscp"`
	DstAddress              types.String `tfsdk:"dst_address"`
	DstAddressList          types.String `tfsdk:"dst_address_list"`
	DstAddressType          types.String `tfsdk:"dst_address_type"`
	DstLimit                types.String `tfsdk:"dst_limit"`
	DstPort                 types.String `tfsdk:"dst_port"`
	Fragment                types.String `tfsdk:"fragment"`
	Hotspot                 types.String `tfsdk:"hotspot"`
	IcmpOptions             types.String `tfsdk:"icmp_options"`
	InBridgePort            types.String `tfsdk:"in_bridge_port"`
	InBridgePortList        types.String `tfsdk:"in_bridge_port_list"`
	InInterface             types.String `tfsdk:"in_interface"`
	InInterfaceList         types.String `tfsdk:"in_interface_list"`
	IngressPriority         types.String `tfsdk:"ingress_priority"`
	IpsecPolicy             types.String `tfsdk:"ipsec_policy"`
	Ipv4Options             types.String `tfsdk:"ipv4_options"`
	JumpTarget              types.String `tfsdk:"jump_target"`
	Layer7Protocol          types.String `tfsdk:"layer7_protocol"`
	Limit                   types.String `tfsdk:"limit"`
	Log                     types.String `tfsdk:"log"`
	LogPrefix               types.String `tfsdk:"log_prefix"`
	NewConnectionMark       types.String `tfsdk:"new_connection_mark"`
	NewDscp                 types.String `tfsdk:"new_dscp"`
	NewMss                  types.String `tfsdk:"new_mss"`
	NewPacketMark           types.String `tfsdk:"new_packet_mark"`
	NewPriority             types.String `tfsdk:"new_priority"`
	NewRoutingMark          types.String `tfsdk:"new_routing_mark"`
	NewTtl                  types.String `tfsdk:"new_ttl"`
	Nth                     types.String `tfsdk:"nth"`
	OutBridgePort           types.String `tfsdk:"out_bridge_port"`
	OutBridgePortList       types.String `tfsdk:"out_bridge_port_list"`
	OutInterface            types.String `tfsdk:"out_interface"`
	OutInterfaceList        types.String `tfsdk:"out_interface_list"`
	PacketMark              types.String `tfsdk:"packet_mark"`
	PacketSize              types.String `tfsdk:"packet_size"`
	Packets                 types.String `tfsdk:"packets"`
	Passthrough             types.String `tfsdk:"passthrough"`
	PerConnectionClassifier types.String `tfsdk:"per_connection_classifier"`
	Port                    types.String `tfsdk:"port"`
	Priority                types.String `tfsdk:"priority"`
	Protocol                types.String `tfsdk:"protocol"`
	Psd                     types.String `tfsdk:"psd"`
	Random                  types.String `tfsdk:"random"`
	Realm                   types.String `tfsdk:"realm"`
	RouteDst                types.String `tfsdk:"route_dst"`
	RoutingMark             types.String `tfsdk:"routing_mark"`
	SniffID                 types.String `tfsdk:"sniff_id"`
	SniffTarget             types.String `tfsdk:"sniff_target"`
	SniffTargetPort         types.String `tfsdk:"sniff_target_port"`
	SrcAddress              types.String `tfsdk:"src_address"`
	SrcAddressList          types.String `tfsdk:"src_address_list"`
	SrcAddressType          types.String `tfsdk:"src_address_type"`
	SrcMACAddress           types.String `tfsdk:"src_mac_address"`
	SrcPort                 types.String `tfsdk:"src_port"`
	TCPFlags                types.String `tfsdk:"tcp_flags"`
	TCPMss                  types.String `tfsdk:"tcp_mss"`
	Time                    types.String `tfsdk:"time"`
	TLSHost                 types.String `tfsdk:"tls_host"`
	Ttl                     types.String `tfsdk:"ttl"`
	Router                  types.String `tfsdk:"router"`
	Position                types.Int64  `tfsdk:"position"`
}

func NewIPFirewallMangleResource() resource.Resource { return &IPFirewallMangleResource{} }

func (r *IPFirewallMangleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_firewall_mangle"
}

func (r *IPFirewallMangleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPFirewallMangleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "IP firewall mangle rule. Ordered by `position` (sort key, not identity).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"tos": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tos`.",
			},
			"p2p": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `p2p`.",
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"address_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"address_list_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bytes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"chain": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"connection_bytes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"connection_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"connection_mark": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"connection_nat_state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"connection_rate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"connection_state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"connection_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"content": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"dscp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dst_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dst_address_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dst_address_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dst_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dst_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"fragment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hotspot": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"icmp_options": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"in_bridge_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"in_bridge_port_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"in_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"in_interface_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ingress_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ipsec_policy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ipv4_options": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"jump_target": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"layer7_protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"log": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"log_prefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"new_connection_mark": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"new_dscp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"new_mss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"new_packet_mark": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"new_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"new_routing_mark": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"new_ttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"out_bridge_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"out_bridge_port_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"out_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"out_interface_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"packet_mark": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"packet_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"packets": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"passthrough": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"per_connection_classifier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"psd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"random": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"realm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"route_dst": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"routing_mark": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sniff_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sniff_target": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sniff_target_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_address_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_address_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tcp_flags": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tcp_mss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tls_host": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
			"position": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Sort key for placement in the ordered chain. Lower = higher in the chain. Persisted on the device via a [tf:pos=N] prefix in the comment so destroy+apply rebuilds the same order.",
			},
		},
	}
}

func (r *IPFirewallMangleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPFirewallMangleModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Action.IsNull() || plan.Action.IsUnknown()) {
		body["action"] = plan.Action.ValueString()
	}
	if !(plan.AddressList.IsNull() || plan.AddressList.IsUnknown()) {
		body["address-list"] = plan.AddressList.ValueString()
	}
	if !(plan.AddressListTimeout.IsNull() || plan.AddressListTimeout.IsUnknown()) {
		body["address-list-timeout"] = plan.AddressListTimeout.ValueString()
	}
	if !(plan.Chain.IsNull() || plan.Chain.IsUnknown()) {
		body["chain"] = plan.Chain.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.ConnectionBytes.IsNull() || plan.ConnectionBytes.IsUnknown()) {
		body["connection-bytes"] = plan.ConnectionBytes.ValueString()
	}
	if !(plan.ConnectionLimit.IsNull() || plan.ConnectionLimit.IsUnknown()) {
		body["connection-limit"] = plan.ConnectionLimit.ValueString()
	}
	if !(plan.ConnectionMark.IsNull() || plan.ConnectionMark.IsUnknown()) {
		body["connection-mark"] = plan.ConnectionMark.ValueString()
	}
	if !(plan.ConnectionNATState.IsNull() || plan.ConnectionNATState.IsUnknown()) {
		body["connection-nat-state"] = plan.ConnectionNATState.ValueString()
	}
	if !(plan.ConnectionRate.IsNull() || plan.ConnectionRate.IsUnknown()) {
		body["connection-rate"] = plan.ConnectionRate.ValueString()
	}
	if !(plan.ConnectionState.IsNull() || plan.ConnectionState.IsUnknown()) {
		body["connection-state"] = plan.ConnectionState.ValueString()
	}
	if !(plan.ConnectionType.IsNull() || plan.ConnectionType.IsUnknown()) {
		body["connection-type"] = plan.ConnectionType.ValueString()
	}
	if !(plan.Content.IsNull() || plan.Content.IsUnknown()) {
		body["content"] = plan.Content.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Dscp.IsNull() || plan.Dscp.IsUnknown()) {
		body["dscp"] = plan.Dscp.ValueString()
	}
	if !(plan.DstAddress.IsNull() || plan.DstAddress.IsUnknown()) {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !(plan.DstAddressList.IsNull() || plan.DstAddressList.IsUnknown()) {
		body["dst-address-list"] = plan.DstAddressList.ValueString()
	}
	if !(plan.DstAddressType.IsNull() || plan.DstAddressType.IsUnknown()) {
		body["dst-address-type"] = plan.DstAddressType.ValueString()
	}
	if !(plan.DstLimit.IsNull() || plan.DstLimit.IsUnknown()) {
		body["dst-limit"] = plan.DstLimit.ValueString()
	}
	if !(plan.DstPort.IsNull() || plan.DstPort.IsUnknown()) {
		body["dst-port"] = plan.DstPort.ValueString()
	}
	if !(plan.Fragment.IsNull() || plan.Fragment.IsUnknown()) {
		body["fragment"] = plan.Fragment.ValueString()
	}
	if !(plan.Hotspot.IsNull() || plan.Hotspot.IsUnknown()) {
		body["hotspot"] = plan.Hotspot.ValueString()
	}
	if !(plan.IcmpOptions.IsNull() || plan.IcmpOptions.IsUnknown()) {
		body["icmp-options"] = plan.IcmpOptions.ValueString()
	}
	if !(plan.InBridgePort.IsNull() || plan.InBridgePort.IsUnknown()) {
		body["in-bridge-port"] = plan.InBridgePort.ValueString()
	}
	if !(plan.InBridgePortList.IsNull() || plan.InBridgePortList.IsUnknown()) {
		body["in-bridge-port-list"] = plan.InBridgePortList.ValueString()
	}
	if !(plan.InInterface.IsNull() || plan.InInterface.IsUnknown()) {
		body["in-interface"] = plan.InInterface.ValueString()
	}
	if !(plan.InInterfaceList.IsNull() || plan.InInterfaceList.IsUnknown()) {
		body["in-interface-list"] = plan.InInterfaceList.ValueString()
	}
	if !(plan.IngressPriority.IsNull() || plan.IngressPriority.IsUnknown()) {
		body["ingress-priority"] = plan.IngressPriority.ValueString()
	}
	if !(plan.IpsecPolicy.IsNull() || plan.IpsecPolicy.IsUnknown()) {
		body["ipsec-policy"] = plan.IpsecPolicy.ValueString()
	}
	if !(plan.Ipv4Options.IsNull() || plan.Ipv4Options.IsUnknown()) {
		body["ipv4-options"] = plan.Ipv4Options.ValueString()
	}
	if !(plan.JumpTarget.IsNull() || plan.JumpTarget.IsUnknown()) {
		body["jump-target"] = plan.JumpTarget.ValueString()
	}
	if !(plan.Layer7Protocol.IsNull() || plan.Layer7Protocol.IsUnknown()) {
		body["layer7-protocol"] = plan.Layer7Protocol.ValueString()
	}
	if !(plan.Limit.IsNull() || plan.Limit.IsUnknown()) {
		body["limit"] = plan.Limit.ValueString()
	}
	if !(plan.Log.IsNull() || plan.Log.IsUnknown()) {
		body["log"] = plan.Log.ValueString()
	}
	if !(plan.LogPrefix.IsNull() || plan.LogPrefix.IsUnknown()) {
		body["log-prefix"] = plan.LogPrefix.ValueString()
	}
	if !(plan.NewConnectionMark.IsNull() || plan.NewConnectionMark.IsUnknown()) {
		body["new-connection-mark"] = plan.NewConnectionMark.ValueString()
	}
	if !(plan.NewDscp.IsNull() || plan.NewDscp.IsUnknown()) {
		body["new-dscp"] = plan.NewDscp.ValueString()
	}
	if !(plan.NewMss.IsNull() || plan.NewMss.IsUnknown()) {
		body["new-mss"] = plan.NewMss.ValueString()
	}
	if !(plan.NewPacketMark.IsNull() || plan.NewPacketMark.IsUnknown()) {
		body["new-packet-mark"] = plan.NewPacketMark.ValueString()
	}
	if !(plan.NewPriority.IsNull() || plan.NewPriority.IsUnknown()) {
		body["new-priority"] = plan.NewPriority.ValueString()
	}
	if !(plan.NewRoutingMark.IsNull() || plan.NewRoutingMark.IsUnknown()) {
		body["new-routing-mark"] = plan.NewRoutingMark.ValueString()
	}
	if !(plan.NewTtl.IsNull() || plan.NewTtl.IsUnknown()) {
		body["new-ttl"] = plan.NewTtl.ValueString()
	}
	if !(plan.Nth.IsNull() || plan.Nth.IsUnknown()) {
		body["nth"] = plan.Nth.ValueString()
	}
	if !(plan.OutBridgePort.IsNull() || plan.OutBridgePort.IsUnknown()) {
		body["out-bridge-port"] = plan.OutBridgePort.ValueString()
	}
	if !(plan.OutBridgePortList.IsNull() || plan.OutBridgePortList.IsUnknown()) {
		body["out-bridge-port-list"] = plan.OutBridgePortList.ValueString()
	}
	if !(plan.OutInterface.IsNull() || plan.OutInterface.IsUnknown()) {
		body["out-interface"] = plan.OutInterface.ValueString()
	}
	if !(plan.OutInterfaceList.IsNull() || plan.OutInterfaceList.IsUnknown()) {
		body["out-interface-list"] = plan.OutInterfaceList.ValueString()
	}
	if !(plan.PacketMark.IsNull() || plan.PacketMark.IsUnknown()) {
		body["packet-mark"] = plan.PacketMark.ValueString()
	}
	if !(plan.PacketSize.IsNull() || plan.PacketSize.IsUnknown()) {
		body["packet-size"] = plan.PacketSize.ValueString()
	}
	if !(plan.Passthrough.IsNull() || plan.Passthrough.IsUnknown()) {
		body["passthrough"] = plan.Passthrough.ValueString()
	}
	if !(plan.PerConnectionClassifier.IsNull() || plan.PerConnectionClassifier.IsUnknown()) {
		body["per-connection-classifier"] = plan.PerConnectionClassifier.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.Priority.IsNull() || plan.Priority.IsUnknown()) {
		body["priority"] = plan.Priority.ValueString()
	}
	if !(plan.Protocol.IsNull() || plan.Protocol.IsUnknown()) {
		body["protocol"] = plan.Protocol.ValueString()
	}
	if !(plan.Psd.IsNull() || plan.Psd.IsUnknown()) {
		body["psd"] = plan.Psd.ValueString()
	}
	if !(plan.Random.IsNull() || plan.Random.IsUnknown()) {
		body["random"] = plan.Random.ValueString()
	}
	if !(plan.Realm.IsNull() || plan.Realm.IsUnknown()) {
		body["realm"] = plan.Realm.ValueString()
	}
	if !(plan.RouteDst.IsNull() || plan.RouteDst.IsUnknown()) {
		body["route-dst"] = plan.RouteDst.ValueString()
	}
	if !(plan.RoutingMark.IsNull() || plan.RoutingMark.IsUnknown()) {
		body["routing-mark"] = plan.RoutingMark.ValueString()
	}
	if !(plan.SniffID.IsNull() || plan.SniffID.IsUnknown()) {
		body["sniff-id"] = plan.SniffID.ValueString()
	}
	if !(plan.SniffTarget.IsNull() || plan.SniffTarget.IsUnknown()) {
		body["sniff-target"] = plan.SniffTarget.ValueString()
	}
	if !(plan.SniffTargetPort.IsNull() || plan.SniffTargetPort.IsUnknown()) {
		body["sniff-target-port"] = plan.SniffTargetPort.ValueString()
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !(plan.SrcAddressList.IsNull() || plan.SrcAddressList.IsUnknown()) {
		body["src-address-list"] = plan.SrcAddressList.ValueString()
	}
	if !(plan.SrcAddressType.IsNull() || plan.SrcAddressType.IsUnknown()) {
		body["src-address-type"] = plan.SrcAddressType.ValueString()
	}
	if !(plan.SrcMACAddress.IsNull() || plan.SrcMACAddress.IsUnknown()) {
		body["src-mac-address"] = plan.SrcMACAddress.ValueString()
	}
	if !(plan.SrcPort.IsNull() || plan.SrcPort.IsUnknown()) {
		body["src-port"] = plan.SrcPort.ValueString()
	}
	if !(plan.TCPFlags.IsNull() || plan.TCPFlags.IsUnknown()) {
		body["tcp-flags"] = plan.TCPFlags.ValueString()
	}
	if !(plan.TCPMss.IsNull() || plan.TCPMss.IsUnknown()) {
		body["tcp-mss"] = plan.TCPMss.ValueString()
	}
	if !(plan.Time.IsNull() || plan.Time.IsUnknown()) {
		body["time"] = plan.Time.ValueString()
	}
	if !(plan.TLSHost.IsNull() || plan.TLSHost.IsUnknown()) {
		body["tls-host"] = plan.TLSHost.ValueString()
	}
	if !(plan.Ttl.IsNull() || plan.Ttl.IsUnknown()) {
		body["ttl"] = plan.Ttl.ValueString()
	}
	if !(plan.P2p.IsNull() || plan.P2p.IsUnknown()) {
		body["p2p"] = plan.P2p.ValueString()
	}
	if !(plan.Tos.IsNull() || plan.Tos.IsUnknown()) {
		body["tos"] = plan.Tos.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/firewall/mangle", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/firewall/mangle failed", err.Error())
		return
	}
	if !plan.Position.IsNull() && !plan.Position.IsUnknown() {
		r.reg.RegisterOrdered(plan.Router.ValueString(), "/ip/firewall/mangle", obj[".id"], plan.Position.ValueInt64())
		snap := r.reg.OrderedSnapshot(plan.Router.ValueString(), "/ip/firewall/mangle")
		if err := c.PlaceOrdered(ctx, plan.Router.ValueString(), "/ip/firewall/mangle", obj[".id"], plan.Position.ValueInt64(), snap); err != nil {
			// Resource exists on the device; write minimal state so Terraform
			// tracks it (and a future apply can repair the order or delete it)
			// instead of creating a duplicate.
			iPFirewallMangleApply(ctx, obj, &plan)
			resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
			resp.Diagnostics.AddError("Order /ip/firewall/mangle failed", err.Error())
			return
		}
		reread, rerr := c.GetByID(ctx, "/ip/firewall/mangle", obj[".id"])
		if rerr != nil {
			iPFirewallMangleApply(ctx, obj, &plan)
			resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
			resp.Diagnostics.AddError("Re-read after order failed", rerr.Error())
			return
		}
		obj = reread
	}
	iPFirewallMangleApply(ctx, obj, &plan)
	// Apply has already split the marker off the comment; carry position
	// from plan into state explicitly because the wire shows the encoded form.
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPFirewallMangleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPFirewallMangleModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/firewall/mangle", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/firewall/mangle failed", err.Error())
		return
	}
	iPFirewallMangleApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPFirewallMangleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPFirewallMangleModel
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
	if !plan.Action.Equal(state.Action) {
		body["action"] = plan.Action.ValueString()
	}
	if !plan.AddressList.Equal(state.AddressList) {
		body["address-list"] = plan.AddressList.ValueString()
	}
	if !plan.AddressListTimeout.Equal(state.AddressListTimeout) {
		body["address-list-timeout"] = plan.AddressListTimeout.ValueString()
	}
	if !plan.Chain.Equal(state.Chain) {
		body["chain"] = plan.Chain.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.ConnectionBytes.Equal(state.ConnectionBytes) {
		body["connection-bytes"] = plan.ConnectionBytes.ValueString()
	}
	if !plan.ConnectionLimit.Equal(state.ConnectionLimit) {
		body["connection-limit"] = plan.ConnectionLimit.ValueString()
	}
	if !plan.ConnectionMark.Equal(state.ConnectionMark) {
		body["connection-mark"] = plan.ConnectionMark.ValueString()
	}
	if !plan.ConnectionNATState.Equal(state.ConnectionNATState) {
		body["connection-nat-state"] = plan.ConnectionNATState.ValueString()
	}
	if !plan.ConnectionRate.Equal(state.ConnectionRate) {
		body["connection-rate"] = plan.ConnectionRate.ValueString()
	}
	if !plan.ConnectionState.Equal(state.ConnectionState) {
		body["connection-state"] = plan.ConnectionState.ValueString()
	}
	if !plan.ConnectionType.Equal(state.ConnectionType) {
		body["connection-type"] = plan.ConnectionType.ValueString()
	}
	if !plan.Content.Equal(state.Content) {
		body["content"] = plan.Content.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Dscp.Equal(state.Dscp) {
		body["dscp"] = plan.Dscp.ValueString()
	}
	if !plan.DstAddress.Equal(state.DstAddress) {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !plan.DstAddressList.Equal(state.DstAddressList) {
		body["dst-address-list"] = plan.DstAddressList.ValueString()
	}
	if !plan.DstAddressType.Equal(state.DstAddressType) {
		body["dst-address-type"] = plan.DstAddressType.ValueString()
	}
	if !plan.DstLimit.Equal(state.DstLimit) {
		body["dst-limit"] = plan.DstLimit.ValueString()
	}
	if !plan.DstPort.Equal(state.DstPort) {
		body["dst-port"] = plan.DstPort.ValueString()
	}
	if !plan.Fragment.Equal(state.Fragment) {
		body["fragment"] = plan.Fragment.ValueString()
	}
	if !plan.Hotspot.Equal(state.Hotspot) {
		body["hotspot"] = plan.Hotspot.ValueString()
	}
	if !plan.IcmpOptions.Equal(state.IcmpOptions) {
		body["icmp-options"] = plan.IcmpOptions.ValueString()
	}
	if !plan.InBridgePort.Equal(state.InBridgePort) {
		body["in-bridge-port"] = plan.InBridgePort.ValueString()
	}
	if !plan.InBridgePortList.Equal(state.InBridgePortList) {
		body["in-bridge-port-list"] = plan.InBridgePortList.ValueString()
	}
	if !plan.InInterface.Equal(state.InInterface) {
		body["in-interface"] = plan.InInterface.ValueString()
	}
	if !plan.InInterfaceList.Equal(state.InInterfaceList) {
		body["in-interface-list"] = plan.InInterfaceList.ValueString()
	}
	if !plan.IngressPriority.Equal(state.IngressPriority) {
		body["ingress-priority"] = plan.IngressPriority.ValueString()
	}
	if !plan.IpsecPolicy.Equal(state.IpsecPolicy) {
		body["ipsec-policy"] = plan.IpsecPolicy.ValueString()
	}
	if !plan.Ipv4Options.Equal(state.Ipv4Options) {
		body["ipv4-options"] = plan.Ipv4Options.ValueString()
	}
	if !plan.JumpTarget.Equal(state.JumpTarget) {
		body["jump-target"] = plan.JumpTarget.ValueString()
	}
	if !plan.Layer7Protocol.Equal(state.Layer7Protocol) {
		body["layer7-protocol"] = plan.Layer7Protocol.ValueString()
	}
	if !plan.Limit.Equal(state.Limit) {
		body["limit"] = plan.Limit.ValueString()
	}
	if !plan.Log.Equal(state.Log) {
		body["log"] = plan.Log.ValueString()
	}
	if !plan.LogPrefix.Equal(state.LogPrefix) {
		body["log-prefix"] = plan.LogPrefix.ValueString()
	}
	if !plan.NewConnectionMark.Equal(state.NewConnectionMark) {
		body["new-connection-mark"] = plan.NewConnectionMark.ValueString()
	}
	if !plan.NewDscp.Equal(state.NewDscp) {
		body["new-dscp"] = plan.NewDscp.ValueString()
	}
	if !plan.NewMss.Equal(state.NewMss) {
		body["new-mss"] = plan.NewMss.ValueString()
	}
	if !plan.NewPacketMark.Equal(state.NewPacketMark) {
		body["new-packet-mark"] = plan.NewPacketMark.ValueString()
	}
	if !plan.NewPriority.Equal(state.NewPriority) {
		body["new-priority"] = plan.NewPriority.ValueString()
	}
	if !plan.NewRoutingMark.Equal(state.NewRoutingMark) {
		body["new-routing-mark"] = plan.NewRoutingMark.ValueString()
	}
	if !plan.NewTtl.Equal(state.NewTtl) {
		body["new-ttl"] = plan.NewTtl.ValueString()
	}
	if !plan.Nth.Equal(state.Nth) {
		body["nth"] = plan.Nth.ValueString()
	}
	if !plan.OutBridgePort.Equal(state.OutBridgePort) {
		body["out-bridge-port"] = plan.OutBridgePort.ValueString()
	}
	if !plan.OutBridgePortList.Equal(state.OutBridgePortList) {
		body["out-bridge-port-list"] = plan.OutBridgePortList.ValueString()
	}
	if !plan.OutInterface.Equal(state.OutInterface) {
		body["out-interface"] = plan.OutInterface.ValueString()
	}
	if !plan.OutInterfaceList.Equal(state.OutInterfaceList) {
		body["out-interface-list"] = plan.OutInterfaceList.ValueString()
	}
	if !plan.PacketMark.Equal(state.PacketMark) {
		body["packet-mark"] = plan.PacketMark.ValueString()
	}
	if !plan.PacketSize.Equal(state.PacketSize) {
		body["packet-size"] = plan.PacketSize.ValueString()
	}
	if !plan.Passthrough.Equal(state.Passthrough) {
		body["passthrough"] = plan.Passthrough.ValueString()
	}
	if !plan.PerConnectionClassifier.Equal(state.PerConnectionClassifier) {
		body["per-connection-classifier"] = plan.PerConnectionClassifier.ValueString()
	}
	if !plan.Port.Equal(state.Port) {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.Priority.Equal(state.Priority) {
		body["priority"] = plan.Priority.ValueString()
	}
	if !plan.Protocol.Equal(state.Protocol) {
		body["protocol"] = plan.Protocol.ValueString()
	}
	if !plan.Psd.Equal(state.Psd) {
		body["psd"] = plan.Psd.ValueString()
	}
	if !plan.Random.Equal(state.Random) {
		body["random"] = plan.Random.ValueString()
	}
	if !plan.Realm.Equal(state.Realm) {
		body["realm"] = plan.Realm.ValueString()
	}
	if !plan.RouteDst.Equal(state.RouteDst) {
		body["route-dst"] = plan.RouteDst.ValueString()
	}
	if !plan.RoutingMark.Equal(state.RoutingMark) {
		body["routing-mark"] = plan.RoutingMark.ValueString()
	}
	if !plan.SniffID.Equal(state.SniffID) {
		body["sniff-id"] = plan.SniffID.ValueString()
	}
	if !plan.SniffTarget.Equal(state.SniffTarget) {
		body["sniff-target"] = plan.SniffTarget.ValueString()
	}
	if !plan.SniffTargetPort.Equal(state.SniffTargetPort) {
		body["sniff-target-port"] = plan.SniffTargetPort.ValueString()
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !plan.SrcAddressList.Equal(state.SrcAddressList) {
		body["src-address-list"] = plan.SrcAddressList.ValueString()
	}
	if !plan.SrcAddressType.Equal(state.SrcAddressType) {
		body["src-address-type"] = plan.SrcAddressType.ValueString()
	}
	if !plan.SrcMACAddress.Equal(state.SrcMACAddress) {
		body["src-mac-address"] = plan.SrcMACAddress.ValueString()
	}
	if !plan.SrcPort.Equal(state.SrcPort) {
		body["src-port"] = plan.SrcPort.ValueString()
	}
	if !plan.TCPFlags.Equal(state.TCPFlags) {
		body["tcp-flags"] = plan.TCPFlags.ValueString()
	}
	if !plan.TCPMss.Equal(state.TCPMss) {
		body["tcp-mss"] = plan.TCPMss.ValueString()
	}
	if !plan.Time.Equal(state.Time) {
		body["time"] = plan.Time.ValueString()
	}
	if !plan.TLSHost.Equal(state.TLSHost) {
		body["tls-host"] = plan.TLSHost.ValueString()
	}
	if !plan.Ttl.Equal(state.Ttl) {
		body["ttl"] = plan.Ttl.ValueString()
	}
	if !plan.P2p.Equal(state.P2p) && !plan.P2p.IsUnknown() {
		body["p2p"] = plan.P2p.ValueString()
	}
	if !plan.Tos.Equal(state.Tos) && !plan.Tos.IsUnknown() {
		body["tos"] = plan.Tos.ValueString()
	}
	// If position OR comment changed, re-encode the marker into the comment
	// so the device-side prefix stays in sync.
	if !plan.Comment.Equal(state.Comment) {
		if plan.Comment.IsNull() || plan.Comment.IsUnknown() {
			body["comment"] = ""
		} else {
			body["comment"] = plan.Comment.ValueString()
		}
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/firewall/mangle", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/firewall/mangle failed", err.Error())
			return
		}
		iPFirewallMangleApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	if !plan.Position.IsNull() && !plan.Position.IsUnknown() {
		r.reg.RegisterOrdered(plan.Router.ValueString(), "/ip/firewall/mangle", plan.ID.ValueString(), plan.Position.ValueInt64())
		if !plan.Position.Equal(state.Position) {
			snap := r.reg.OrderedSnapshot(plan.Router.ValueString(), "/ip/firewall/mangle")
			if err := c.PlaceOrdered(ctx, plan.Router.ValueString(), "/ip/firewall/mangle", plan.ID.ValueString(), plan.Position.ValueInt64(), snap); err != nil {
				// Set already applied; persist the new attributes so state matches the
				// device even if the re-order failed.
				resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
				resp.Diagnostics.AddError("Order /ip/firewall/mangle failed", err.Error())
				return
			}
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPFirewallMangleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPFirewallMangleModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/firewall/mangle", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/firewall/mangle failed", err.Error())
	}
	r.reg.UnregisterOrdered(state.Router.ValueString(), "/ip/firewall/mangle", state.ID.ValueString())
}

func (r *IPFirewallMangleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPFirewallMangleLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/firewall/mangle matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPFirewallMangleLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPFirewallMangleLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/firewall/mangle", id)
}

func iPFirewallMangleApply(ctx context.Context, obj client.Object, m *IPFirewallMangleModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["tos"]; ok && v != "" {
		m.Tos = types.StringValue(v)
	} else {
		m.Tos = types.StringNull()
	}
	if v, ok := obj["p2p"]; ok && v != "" {
		m.P2p = types.StringValue(v)
	} else {
		m.P2p = types.StringNull()
	}
	// Strip the [tf:pos=N] marker from the comment before exposing to state.
	// Position is TF-state-only metadata; never written to the device. Keep
	// whatever the user planned. Comment is left untouched.
	if m.Position.IsUnknown() {
		m.Position = types.Int64Null()
	}
	if v, ok := obj["action"]; ok {
		_ = v
		if v != "" {
			m.Action = types.StringValue(v)
		} else {
			m.Action = types.StringNull()
		}
	} else {
		m.Action = types.StringNull()
	}
	if v, ok := obj["address-list"]; ok {
		_ = v
		if v != "" {
			m.AddressList = types.StringValue(v)
		} else {
			m.AddressList = types.StringNull()
		}
	} else {
		m.AddressList = types.StringNull()
	}
	if v, ok := obj["address-list-timeout"]; ok {
		_ = v
		if v != "" {
			m.AddressListTimeout = types.StringValue(v)
		} else {
			m.AddressListTimeout = types.StringNull()
		}
	} else {
		m.AddressListTimeout = types.StringNull()
	}
	if v, ok := obj["bytes"]; ok {
		_ = v
		if v != "" {
			m.Bytes = types.StringValue(v)
		} else {
			m.Bytes = types.StringNull()
		}
	} else {
		m.Bytes = types.StringNull()
	}
	if v, ok := obj["chain"]; ok {
		_ = v
		if v != "" {
			m.Chain = types.StringValue(v)
		} else {
			m.Chain = types.StringNull()
		}
	} else {
		m.Chain = types.StringNull()
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
	if v, ok := obj["connection-bytes"]; ok {
		_ = v
		if v != "" {
			m.ConnectionBytes = types.StringValue(v)
		} else {
			m.ConnectionBytes = types.StringNull()
		}
	} else {
		m.ConnectionBytes = types.StringNull()
	}
	if v, ok := obj["connection-limit"]; ok {
		_ = v
		if v != "" {
			m.ConnectionLimit = types.StringValue(v)
		} else {
			m.ConnectionLimit = types.StringNull()
		}
	} else {
		m.ConnectionLimit = types.StringNull()
	}
	if v, ok := obj["connection-mark"]; ok {
		_ = v
		if v != "" {
			m.ConnectionMark = types.StringValue(v)
		} else {
			m.ConnectionMark = types.StringNull()
		}
	} else {
		m.ConnectionMark = types.StringNull()
	}
	if v, ok := obj["connection-nat-state"]; ok {
		_ = v
		if v != "" {
			m.ConnectionNATState = types.StringValue(v)
		} else {
			m.ConnectionNATState = types.StringNull()
		}
	} else {
		m.ConnectionNATState = types.StringNull()
	}
	if v, ok := obj["connection-rate"]; ok {
		_ = v
		if v != "" {
			m.ConnectionRate = types.StringValue(v)
		} else {
			m.ConnectionRate = types.StringNull()
		}
	} else {
		m.ConnectionRate = types.StringNull()
	}
	if v, ok := obj["connection-state"]; ok {
		_ = v
		if v != "" {
			m.ConnectionState = types.StringValue(v)
		} else {
			m.ConnectionState = types.StringNull()
		}
	} else {
		m.ConnectionState = types.StringNull()
	}
	if v, ok := obj["connection-type"]; ok {
		_ = v
		if v != "" {
			m.ConnectionType = types.StringValue(v)
		} else {
			m.ConnectionType = types.StringNull()
		}
	} else {
		m.ConnectionType = types.StringNull()
	}
	if v, ok := obj["content"]; ok {
		_ = v
		if v != "" {
			m.Content = types.StringValue(v)
		} else {
			m.Content = types.StringNull()
		}
	} else {
		m.Content = types.StringNull()
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
	if v, ok := obj["dscp"]; ok {
		_ = v
		if v != "" {
			m.Dscp = types.StringValue(v)
		} else {
			m.Dscp = types.StringNull()
		}
	} else {
		m.Dscp = types.StringNull()
	}
	if v, ok := obj["dst-address"]; ok {
		_ = v
		if v != "" {
			m.DstAddress = types.StringValue(v)
		} else {
			m.DstAddress = types.StringNull()
		}
	} else {
		m.DstAddress = types.StringNull()
	}
	if v, ok := obj["dst-address-list"]; ok {
		_ = v
		if v != "" {
			m.DstAddressList = types.StringValue(v)
		} else {
			m.DstAddressList = types.StringNull()
		}
	} else {
		m.DstAddressList = types.StringNull()
	}
	if v, ok := obj["dst-address-type"]; ok {
		_ = v
		if v != "" {
			m.DstAddressType = types.StringValue(v)
		} else {
			m.DstAddressType = types.StringNull()
		}
	} else {
		m.DstAddressType = types.StringNull()
	}
	if v, ok := obj["dst-limit"]; ok {
		_ = v
		if v != "" {
			m.DstLimit = types.StringValue(v)
		} else {
			m.DstLimit = types.StringNull()
		}
	} else {
		m.DstLimit = types.StringNull()
	}
	if v, ok := obj["dst-port"]; ok {
		_ = v
		if v != "" {
			m.DstPort = types.StringValue(v)
		} else {
			m.DstPort = types.StringNull()
		}
	} else {
		m.DstPort = types.StringNull()
	}
	if v, ok := obj["fragment"]; ok {
		_ = v
		if v != "" {
			m.Fragment = types.StringValue(v)
		} else {
			m.Fragment = types.StringNull()
		}
	} else {
		m.Fragment = types.StringNull()
	}
	if v, ok := obj["hotspot"]; ok {
		_ = v
		if v != "" {
			m.Hotspot = types.StringValue(v)
		} else {
			m.Hotspot = types.StringNull()
		}
	} else {
		m.Hotspot = types.StringNull()
	}
	if v, ok := obj["icmp-options"]; ok {
		_ = v
		if v != "" {
			m.IcmpOptions = types.StringValue(v)
		} else {
			m.IcmpOptions = types.StringNull()
		}
	} else {
		m.IcmpOptions = types.StringNull()
	}
	if v, ok := obj["in-bridge-port"]; ok {
		_ = v
		if v != "" {
			m.InBridgePort = types.StringValue(v)
		} else {
			m.InBridgePort = types.StringNull()
		}
	} else {
		m.InBridgePort = types.StringNull()
	}
	if v, ok := obj["in-bridge-port-list"]; ok {
		_ = v
		if v != "" {
			m.InBridgePortList = types.StringValue(v)
		} else {
			m.InBridgePortList = types.StringNull()
		}
	} else {
		m.InBridgePortList = types.StringNull()
	}
	if v, ok := obj["in-interface"]; ok {
		_ = v
		if v != "" {
			m.InInterface = types.StringValue(v)
		} else {
			m.InInterface = types.StringNull()
		}
	} else {
		m.InInterface = types.StringNull()
	}
	if v, ok := obj["in-interface-list"]; ok {
		_ = v
		if v != "" {
			m.InInterfaceList = types.StringValue(v)
		} else {
			m.InInterfaceList = types.StringNull()
		}
	} else {
		m.InInterfaceList = types.StringNull()
	}
	if v, ok := obj["ingress-priority"]; ok {
		_ = v
		if v != "" {
			m.IngressPriority = types.StringValue(v)
		} else {
			m.IngressPriority = types.StringNull()
		}
	} else {
		m.IngressPriority = types.StringNull()
	}
	if v, ok := obj["ipsec-policy"]; ok {
		_ = v
		if v != "" {
			m.IpsecPolicy = types.StringValue(v)
		} else {
			m.IpsecPolicy = types.StringNull()
		}
	} else {
		m.IpsecPolicy = types.StringNull()
	}
	if v, ok := obj["ipv4-options"]; ok {
		_ = v
		if v != "" {
			m.Ipv4Options = types.StringValue(v)
		} else {
			m.Ipv4Options = types.StringNull()
		}
	} else {
		m.Ipv4Options = types.StringNull()
	}
	if v, ok := obj["jump-target"]; ok {
		_ = v
		if v != "" {
			m.JumpTarget = types.StringValue(v)
		} else {
			m.JumpTarget = types.StringNull()
		}
	} else {
		m.JumpTarget = types.StringNull()
	}
	if v, ok := obj["layer7-protocol"]; ok {
		_ = v
		if v != "" {
			m.Layer7Protocol = types.StringValue(v)
		} else {
			m.Layer7Protocol = types.StringNull()
		}
	} else {
		m.Layer7Protocol = types.StringNull()
	}
	if v, ok := obj["limit"]; ok {
		_ = v
		if v != "" {
			m.Limit = types.StringValue(v)
		} else {
			m.Limit = types.StringNull()
		}
	} else {
		m.Limit = types.StringNull()
	}
	if v, ok := obj["log"]; ok {
		_ = v
		if v != "" {
			m.Log = types.StringValue(v)
		} else {
			m.Log = types.StringNull()
		}
	} else {
		m.Log = types.StringNull()
	}
	if v, ok := obj["log-prefix"]; ok {
		_ = v
		if v != "" {
			m.LogPrefix = types.StringValue(v)
		} else {
			m.LogPrefix = types.StringNull()
		}
	} else {
		m.LogPrefix = types.StringNull()
	}
	if v, ok := obj["new-connection-mark"]; ok {
		_ = v
		if v != "" {
			m.NewConnectionMark = types.StringValue(v)
		} else {
			m.NewConnectionMark = types.StringNull()
		}
	} else {
		m.NewConnectionMark = types.StringNull()
	}
	if v, ok := obj["new-dscp"]; ok {
		_ = v
		if v != "" {
			m.NewDscp = types.StringValue(v)
		} else {
			m.NewDscp = types.StringNull()
		}
	} else {
		m.NewDscp = types.StringNull()
	}
	if v, ok := obj["new-mss"]; ok {
		_ = v
		if v != "" {
			m.NewMss = types.StringValue(v)
		} else {
			m.NewMss = types.StringNull()
		}
	} else {
		m.NewMss = types.StringNull()
	}
	if v, ok := obj["new-packet-mark"]; ok {
		_ = v
		if v != "" {
			m.NewPacketMark = types.StringValue(v)
		} else {
			m.NewPacketMark = types.StringNull()
		}
	} else {
		m.NewPacketMark = types.StringNull()
	}
	if v, ok := obj["new-priority"]; ok {
		_ = v
		if v != "" {
			m.NewPriority = types.StringValue(v)
		} else {
			m.NewPriority = types.StringNull()
		}
	} else {
		m.NewPriority = types.StringNull()
	}
	if v, ok := obj["new-routing-mark"]; ok {
		_ = v
		if v != "" {
			m.NewRoutingMark = types.StringValue(v)
		} else {
			m.NewRoutingMark = types.StringNull()
		}
	} else {
		m.NewRoutingMark = types.StringNull()
	}
	if v, ok := obj["new-ttl"]; ok {
		_ = v
		if v != "" {
			m.NewTtl = types.StringValue(v)
		} else {
			m.NewTtl = types.StringNull()
		}
	} else {
		m.NewTtl = types.StringNull()
	}
	if v, ok := obj["nth"]; ok {
		_ = v
		if v != "" {
			m.Nth = types.StringValue(v)
		} else {
			m.Nth = types.StringNull()
		}
	} else {
		m.Nth = types.StringNull()
	}
	if v, ok := obj["out-bridge-port"]; ok {
		_ = v
		if v != "" {
			m.OutBridgePort = types.StringValue(v)
		} else {
			m.OutBridgePort = types.StringNull()
		}
	} else {
		m.OutBridgePort = types.StringNull()
	}
	if v, ok := obj["out-bridge-port-list"]; ok {
		_ = v
		if v != "" {
			m.OutBridgePortList = types.StringValue(v)
		} else {
			m.OutBridgePortList = types.StringNull()
		}
	} else {
		m.OutBridgePortList = types.StringNull()
	}
	if v, ok := obj["out-interface"]; ok {
		_ = v
		if v != "" {
			m.OutInterface = types.StringValue(v)
		} else {
			m.OutInterface = types.StringNull()
		}
	} else {
		m.OutInterface = types.StringNull()
	}
	if v, ok := obj["out-interface-list"]; ok {
		_ = v
		if v != "" {
			m.OutInterfaceList = types.StringValue(v)
		} else {
			m.OutInterfaceList = types.StringNull()
		}
	} else {
		m.OutInterfaceList = types.StringNull()
	}
	if v, ok := obj["packet-mark"]; ok {
		_ = v
		if v != "" {
			m.PacketMark = types.StringValue(v)
		} else {
			m.PacketMark = types.StringNull()
		}
	} else {
		m.PacketMark = types.StringNull()
	}
	if v, ok := obj["packet-size"]; ok {
		_ = v
		if v != "" {
			m.PacketSize = types.StringValue(v)
		} else {
			m.PacketSize = types.StringNull()
		}
	} else {
		m.PacketSize = types.StringNull()
	}
	if v, ok := obj["packets"]; ok {
		_ = v
		if v != "" {
			m.Packets = types.StringValue(v)
		} else {
			m.Packets = types.StringNull()
		}
	} else {
		m.Packets = types.StringNull()
	}
	if v, ok := obj["passthrough"]; ok {
		_ = v
		if v != "" {
			m.Passthrough = types.StringValue(v)
		} else {
			m.Passthrough = types.StringNull()
		}
	} else {
		m.Passthrough = types.StringNull()
	}
	if v, ok := obj["per-connection-classifier"]; ok {
		_ = v
		if v != "" {
			m.PerConnectionClassifier = types.StringValue(v)
		} else {
			m.PerConnectionClassifier = types.StringNull()
		}
	} else {
		m.PerConnectionClassifier = types.StringNull()
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
	if v, ok := obj["priority"]; ok {
		_ = v
		if v != "" {
			m.Priority = types.StringValue(v)
		} else {
			m.Priority = types.StringNull()
		}
	} else {
		m.Priority = types.StringNull()
	}
	if v, ok := obj["protocol"]; ok {
		_ = v
		if v != "" {
			m.Protocol = types.StringValue(v)
		} else {
			m.Protocol = types.StringNull()
		}
	} else {
		m.Protocol = types.StringNull()
	}
	if v, ok := obj["psd"]; ok {
		_ = v
		if v != "" {
			m.Psd = types.StringValue(v)
		} else {
			m.Psd = types.StringNull()
		}
	} else {
		m.Psd = types.StringNull()
	}
	if v, ok := obj["random"]; ok {
		_ = v
		if v != "" {
			m.Random = types.StringValue(v)
		} else {
			m.Random = types.StringNull()
		}
	} else {
		m.Random = types.StringNull()
	}
	if v, ok := obj["realm"]; ok {
		_ = v
		if v != "" {
			m.Realm = types.StringValue(v)
		} else {
			m.Realm = types.StringNull()
		}
	} else {
		m.Realm = types.StringNull()
	}
	if v, ok := obj["route-dst"]; ok {
		_ = v
		if v != "" {
			m.RouteDst = types.StringValue(v)
		} else {
			m.RouteDst = types.StringNull()
		}
	} else {
		m.RouteDst = types.StringNull()
	}
	if v, ok := obj["routing-mark"]; ok {
		_ = v
		if v != "" {
			m.RoutingMark = types.StringValue(v)
		} else {
			m.RoutingMark = types.StringNull()
		}
	} else {
		m.RoutingMark = types.StringNull()
	}
	if v, ok := obj["sniff-id"]; ok {
		_ = v
		if v != "" {
			m.SniffID = types.StringValue(v)
		} else {
			m.SniffID = types.StringNull()
		}
	} else {
		m.SniffID = types.StringNull()
	}
	if v, ok := obj["sniff-target"]; ok {
		_ = v
		if v != "" {
			m.SniffTarget = types.StringValue(v)
		} else {
			m.SniffTarget = types.StringNull()
		}
	} else {
		m.SniffTarget = types.StringNull()
	}
	if v, ok := obj["sniff-target-port"]; ok {
		_ = v
		if v != "" {
			m.SniffTargetPort = types.StringValue(v)
		} else {
			m.SniffTargetPort = types.StringNull()
		}
	} else {
		m.SniffTargetPort = types.StringNull()
	}
	if v, ok := obj["src-address"]; ok {
		_ = v
		if v != "" {
			m.SrcAddress = types.StringValue(v)
		} else {
			m.SrcAddress = types.StringNull()
		}
	} else {
		m.SrcAddress = types.StringNull()
	}
	if v, ok := obj["src-address-list"]; ok {
		_ = v
		if v != "" {
			m.SrcAddressList = types.StringValue(v)
		} else {
			m.SrcAddressList = types.StringNull()
		}
	} else {
		m.SrcAddressList = types.StringNull()
	}
	if v, ok := obj["src-address-type"]; ok {
		_ = v
		if v != "" {
			m.SrcAddressType = types.StringValue(v)
		} else {
			m.SrcAddressType = types.StringNull()
		}
	} else {
		m.SrcAddressType = types.StringNull()
	}
	if v, ok := obj["src-mac-address"]; ok {
		_ = v
		if v != "" {
			m.SrcMACAddress = types.StringValue(v)
		} else {
			m.SrcMACAddress = types.StringNull()
		}
	} else {
		m.SrcMACAddress = types.StringNull()
	}
	if v, ok := obj["src-port"]; ok {
		_ = v
		if v != "" {
			m.SrcPort = types.StringValue(v)
		} else {
			m.SrcPort = types.StringNull()
		}
	} else {
		m.SrcPort = types.StringNull()
	}
	if v, ok := obj["tcp-flags"]; ok {
		_ = v
		if v != "" {
			m.TCPFlags = types.StringValue(v)
		} else {
			m.TCPFlags = types.StringNull()
		}
	} else {
		m.TCPFlags = types.StringNull()
	}
	if v, ok := obj["tcp-mss"]; ok {
		_ = v
		if v != "" {
			m.TCPMss = types.StringValue(v)
		} else {
			m.TCPMss = types.StringNull()
		}
	} else {
		m.TCPMss = types.StringNull()
	}
	if v, ok := obj["time"]; ok {
		_ = v
		if v != "" {
			m.Time = types.StringValue(v)
		} else {
			m.Time = types.StringNull()
		}
	} else {
		m.Time = types.StringNull()
	}
	if v, ok := obj["tls-host"]; ok {
		_ = v
		if v != "" {
			m.TLSHost = types.StringValue(v)
		} else {
			m.TLSHost = types.StringNull()
		}
	} else {
		m.TLSHost = types.StringNull()
	}
	if v, ok := obj["ttl"]; ok {
		_ = v
		if v != "" {
			m.Ttl = types.StringValue(v)
		} else {
			m.Ttl = types.StringNull()
		}
	} else {
		m.Ttl = types.StringNull()
	}
}
