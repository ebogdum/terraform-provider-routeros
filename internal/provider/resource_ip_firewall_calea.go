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
	_ resource.Resource                = &IPFirewallCaleaResource{}
	_ resource.ResourceWithImportState = &IPFirewallCaleaResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPFirewallCaleaResource struct {
	reg *client.Registry
}

type IPFirewallCaleaModel struct {
	ID                      types.String `tfsdk:"id"`
	Ttl                     types.String `tfsdk:"ttl"`
	Tos                     types.String `tfsdk:"tos"`
	TlsHost                 types.String `tfsdk:"tls_host"`
	Time                    types.String `tfsdk:"time"`
	TcpMss                  types.String `tfsdk:"tcp_mss"`
	SrcPort                 types.String `tfsdk:"src_port"`
	SrcMacAddress           types.String `tfsdk:"src_mac_address"`
	SrcAddressType          types.String `tfsdk:"src_address_type"`
	SrcAddressList          types.String `tfsdk:"src_address_list"`
	SrcAddress              types.String `tfsdk:"src_address"`
	SniffTargetPort         types.String `tfsdk:"sniff_target_port"`
	SniffTarget             types.String `tfsdk:"sniff_target"`
	SniffId                 types.String `tfsdk:"sniff_id"`
	RoutingMark             types.String `tfsdk:"routing_mark"`
	Realm                   types.String `tfsdk:"realm"`
	Random                  types.String `tfsdk:"random"`
	Psd                     types.String `tfsdk:"psd"`
	Protocol                types.String `tfsdk:"protocol"`
	Priority                types.String `tfsdk:"priority"`
	Port                    types.String `tfsdk:"port"`
	PerConnectionClassifier types.String `tfsdk:"per_connection_classifier"`
	PacketSize              types.String `tfsdk:"packet_size"`
	PacketMark              types.String `tfsdk:"packet_mark"`
	OutInterfaceList        types.String `tfsdk:"out_interface_list"`
	OutInterface            types.String `tfsdk:"out_interface"`
	OutBridgePortList       types.String `tfsdk:"out_bridge_port_list"`
	OutBridgePort           types.String `tfsdk:"out_bridge_port"`
	Nth                     types.String `tfsdk:"nth"`
	LogPrefix               types.String `tfsdk:"log_prefix"`
	Log                     types.String `tfsdk:"log"`
	Limit                   types.String `tfsdk:"limit"`
	Layer7Protocol          types.String `tfsdk:"layer7_protocol"`
	Ipv4Options             types.String `tfsdk:"ipv4_options"`
	IpsecPolicy             types.String `tfsdk:"ipsec_policy"`
	IngressPriority         types.String `tfsdk:"ingress_priority"`
	InInterfaceList         types.String `tfsdk:"in_interface_list"`
	InInterface             types.String `tfsdk:"in_interface"`
	InBridgePortList        types.String `tfsdk:"in_bridge_port_list"`
	InBridgePort            types.String `tfsdk:"in_bridge_port"`
	IcmpOptions             types.String `tfsdk:"icmp_options"`
	Hotspot                 types.String `tfsdk:"hotspot"`
	Fragment                types.String `tfsdk:"fragment"`
	DstPort                 types.String `tfsdk:"dst_port"`
	DstLimit                types.String `tfsdk:"dst_limit"`
	DstAddressType          types.String `tfsdk:"dst_address_type"`
	DstAddressList          types.String `tfsdk:"dst_address_list"`
	DstAddress              types.String `tfsdk:"dst_address"`
	Dscp                    types.String `tfsdk:"dscp"`
	Content                 types.String `tfsdk:"content"`
	ConnectionType          types.String `tfsdk:"connection_type"`
	ConnectionRate          types.String `tfsdk:"connection_rate"`
	ConnectionMark          types.String `tfsdk:"connection_mark"`
	ConnectionLimit         types.String `tfsdk:"connection_limit"`
	ConnectionBytes         types.String `tfsdk:"connection_bytes"`
	Chain                   types.String `tfsdk:"chain"`
	AddressListTimeout      types.String `tfsdk:"address_list_timeout"`
	AddressList             types.String `tfsdk:"address_list"`
	Action                  types.String `tfsdk:"action"`
	Comment                 types.String `tfsdk:"comment"`
	Disabled                types.Bool   `tfsdk:"disabled"`
	Router                  types.String `tfsdk:"router"`
}

func NewIPFirewallCaleaResource() resource.Resource { return &IPFirewallCaleaResource{} }

func (r *IPFirewallCaleaResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_firewall_calea"
}

func (r *IPFirewallCaleaResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPFirewallCaleaResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "IP firewall CALEA creates a session that drops the management connection on CHR. Skipped.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ttl`.",
			},
			"tos": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tos`.",
			},
			"tls_host": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tls-host`.",
			},
			"time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `time`.",
			},
			"tcp_mss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tcp-mss`.",
			},
			"src_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `src-port`.",
			},
			"src_mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `src-mac-address`.",
			},
			"src_address_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `src-address-type`.",
			},
			"src_address_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `src-address-list`.",
			},
			"src_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `src-address`.",
			},
			"sniff_target_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `sniff-target-port`.",
			},
			"sniff_target": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `sniff-target`.",
			},
			"sniff_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `sniff-id`.",
			},
			"routing_mark": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `routing-mark`.",
			},
			"realm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `realm`.",
			},
			"random": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `random`.",
			},
			"psd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `psd`.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `protocol`.",
			},
			"priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `priority`.",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `port`.",
			},
			"per_connection_classifier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `per-connection-classifier`.",
			},
			"packet_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `packet-size`.",
			},
			"packet_mark": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `packet-mark`.",
			},
			"out_interface_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `out-interface-list`.",
			},
			"out_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `out-interface`.",
			},
			"out_bridge_port_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `out-bridge-port-list`.",
			},
			"out_bridge_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `out-bridge-port`.",
			},
			"nth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `nth`.",
			},
			"log_prefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `log-prefix`.",
			},
			"log": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `log`.",
			},
			"limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `limit`.",
			},
			"layer7_protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `layer7-protocol`.",
			},
			"ipv4_options": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ipv4-options`.",
			},
			"ipsec_policy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ipsec-policy`.",
			},
			"ingress_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ingress-priority`.",
			},
			"in_interface_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `in-interface-list`.",
			},
			"in_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `in-interface`.",
			},
			"in_bridge_port_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `in-bridge-port-list`.",
			},
			"in_bridge_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `in-bridge-port`.",
			},
			"icmp_options": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `icmp-options`.",
			},
			"hotspot": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hotspot`.",
			},
			"fragment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fragment`.",
			},
			"dst_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dst-port`.",
			},
			"dst_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dst-limit`.",
			},
			"dst_address_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dst-address-type`.",
			},
			"dst_address_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dst-address-list`.",
			},
			"dst_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dst-address`.",
			},
			"dscp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dscp`.",
			},
			"content": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `content`.",
			},
			"connection_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `connection-type`.",
			},
			"connection_rate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `connection-rate`.",
			},
			"connection_mark": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `connection-mark`.",
			},
			"connection_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `connection-limit`.",
			},
			"connection_bytes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `connection-bytes`.",
			},
			"chain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `chain`.",
			},
			"address_list_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `address-list-timeout`.",
			},
			"address_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `address-list`.",
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `action`.",
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
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPFirewallCaleaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPFirewallCaleaModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
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
	if !(plan.ConnectionBytes.IsNull() || plan.ConnectionBytes.IsUnknown()) {
		body["connection-bytes"] = plan.ConnectionBytes.ValueString()
	}
	if !(plan.ConnectionLimit.IsNull() || plan.ConnectionLimit.IsUnknown()) {
		body["connection-limit"] = plan.ConnectionLimit.ValueString()
	}
	if !(plan.ConnectionMark.IsNull() || plan.ConnectionMark.IsUnknown()) {
		body["connection-mark"] = plan.ConnectionMark.ValueString()
	}
	if !(plan.ConnectionRate.IsNull() || plan.ConnectionRate.IsUnknown()) {
		body["connection-rate"] = plan.ConnectionRate.ValueString()
	}
	if !(plan.ConnectionType.IsNull() || plan.ConnectionType.IsUnknown()) {
		body["connection-type"] = plan.ConnectionType.ValueString()
	}
	if !(plan.Content.IsNull() || plan.Content.IsUnknown()) {
		body["content"] = plan.Content.ValueString()
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
	if !(plan.RoutingMark.IsNull() || plan.RoutingMark.IsUnknown()) {
		body["routing-mark"] = plan.RoutingMark.ValueString()
	}
	if !(plan.SniffId.IsNull() || plan.SniffId.IsUnknown()) {
		body["sniff-id"] = plan.SniffId.ValueString()
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
	if !(plan.SrcMacAddress.IsNull() || plan.SrcMacAddress.IsUnknown()) {
		body["src-mac-address"] = plan.SrcMacAddress.ValueString()
	}
	if !(plan.SrcPort.IsNull() || plan.SrcPort.IsUnknown()) {
		body["src-port"] = plan.SrcPort.ValueString()
	}
	if !(plan.TcpMss.IsNull() || plan.TcpMss.IsUnknown()) {
		body["tcp-mss"] = plan.TcpMss.ValueString()
	}
	if !(plan.Time.IsNull() || plan.Time.IsUnknown()) {
		body["time"] = plan.Time.ValueString()
	}
	if !(plan.TlsHost.IsNull() || plan.TlsHost.IsUnknown()) {
		body["tls-host"] = plan.TlsHost.ValueString()
	}
	if !(plan.Tos.IsNull() || plan.Tos.IsUnknown()) {
		body["tos"] = plan.Tos.ValueString()
	}
	if !(plan.Ttl.IsNull() || plan.Ttl.IsUnknown()) {
		body["ttl"] = plan.Ttl.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/firewall/calea", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/firewall/calea failed", err.Error())
		return
	}
	iPFirewallCaleaApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPFirewallCaleaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPFirewallCaleaModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/firewall/calea", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/firewall/calea failed", err.Error())
		return
	}
	iPFirewallCaleaApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPFirewallCaleaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPFirewallCaleaModel
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
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Action.Equal(state.Action) && !plan.Action.IsUnknown() {
		body["action"] = plan.Action.ValueString()
	}
	if !plan.AddressList.Equal(state.AddressList) && !plan.AddressList.IsUnknown() {
		body["address-list"] = plan.AddressList.ValueString()
	}
	if !plan.AddressListTimeout.Equal(state.AddressListTimeout) && !plan.AddressListTimeout.IsUnknown() {
		body["address-list-timeout"] = plan.AddressListTimeout.ValueString()
	}
	if !plan.Chain.Equal(state.Chain) && !plan.Chain.IsUnknown() {
		body["chain"] = plan.Chain.ValueString()
	}
	if !plan.ConnectionBytes.Equal(state.ConnectionBytes) && !plan.ConnectionBytes.IsUnknown() {
		body["connection-bytes"] = plan.ConnectionBytes.ValueString()
	}
	if !plan.ConnectionLimit.Equal(state.ConnectionLimit) && !plan.ConnectionLimit.IsUnknown() {
		body["connection-limit"] = plan.ConnectionLimit.ValueString()
	}
	if !plan.ConnectionMark.Equal(state.ConnectionMark) && !plan.ConnectionMark.IsUnknown() {
		body["connection-mark"] = plan.ConnectionMark.ValueString()
	}
	if !plan.ConnectionRate.Equal(state.ConnectionRate) && !plan.ConnectionRate.IsUnknown() {
		body["connection-rate"] = plan.ConnectionRate.ValueString()
	}
	if !plan.ConnectionType.Equal(state.ConnectionType) && !plan.ConnectionType.IsUnknown() {
		body["connection-type"] = plan.ConnectionType.ValueString()
	}
	if !plan.Content.Equal(state.Content) && !plan.Content.IsUnknown() {
		body["content"] = plan.Content.ValueString()
	}
	if !plan.Dscp.Equal(state.Dscp) && !plan.Dscp.IsUnknown() {
		body["dscp"] = plan.Dscp.ValueString()
	}
	if !plan.DstAddress.Equal(state.DstAddress) && !plan.DstAddress.IsUnknown() {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !plan.DstAddressList.Equal(state.DstAddressList) && !plan.DstAddressList.IsUnknown() {
		body["dst-address-list"] = plan.DstAddressList.ValueString()
	}
	if !plan.DstAddressType.Equal(state.DstAddressType) && !plan.DstAddressType.IsUnknown() {
		body["dst-address-type"] = plan.DstAddressType.ValueString()
	}
	if !plan.DstLimit.Equal(state.DstLimit) && !plan.DstLimit.IsUnknown() {
		body["dst-limit"] = plan.DstLimit.ValueString()
	}
	if !plan.DstPort.Equal(state.DstPort) && !plan.DstPort.IsUnknown() {
		body["dst-port"] = plan.DstPort.ValueString()
	}
	if !plan.Fragment.Equal(state.Fragment) && !plan.Fragment.IsUnknown() {
		body["fragment"] = plan.Fragment.ValueString()
	}
	if !plan.Hotspot.Equal(state.Hotspot) && !plan.Hotspot.IsUnknown() {
		body["hotspot"] = plan.Hotspot.ValueString()
	}
	if !plan.IcmpOptions.Equal(state.IcmpOptions) && !plan.IcmpOptions.IsUnknown() {
		body["icmp-options"] = plan.IcmpOptions.ValueString()
	}
	if !plan.InBridgePort.Equal(state.InBridgePort) && !plan.InBridgePort.IsUnknown() {
		body["in-bridge-port"] = plan.InBridgePort.ValueString()
	}
	if !plan.InBridgePortList.Equal(state.InBridgePortList) && !plan.InBridgePortList.IsUnknown() {
		body["in-bridge-port-list"] = plan.InBridgePortList.ValueString()
	}
	if !plan.InInterface.Equal(state.InInterface) && !plan.InInterface.IsUnknown() {
		body["in-interface"] = plan.InInterface.ValueString()
	}
	if !plan.InInterfaceList.Equal(state.InInterfaceList) && !plan.InInterfaceList.IsUnknown() {
		body["in-interface-list"] = plan.InInterfaceList.ValueString()
	}
	if !plan.IngressPriority.Equal(state.IngressPriority) && !plan.IngressPriority.IsUnknown() {
		body["ingress-priority"] = plan.IngressPriority.ValueString()
	}
	if !plan.IpsecPolicy.Equal(state.IpsecPolicy) && !plan.IpsecPolicy.IsUnknown() {
		body["ipsec-policy"] = plan.IpsecPolicy.ValueString()
	}
	if !plan.Ipv4Options.Equal(state.Ipv4Options) && !plan.Ipv4Options.IsUnknown() {
		body["ipv4-options"] = plan.Ipv4Options.ValueString()
	}
	if !plan.Layer7Protocol.Equal(state.Layer7Protocol) && !plan.Layer7Protocol.IsUnknown() {
		body["layer7-protocol"] = plan.Layer7Protocol.ValueString()
	}
	if !plan.Limit.Equal(state.Limit) && !plan.Limit.IsUnknown() {
		body["limit"] = plan.Limit.ValueString()
	}
	if !plan.Log.Equal(state.Log) && !plan.Log.IsUnknown() {
		body["log"] = plan.Log.ValueString()
	}
	if !plan.LogPrefix.Equal(state.LogPrefix) && !plan.LogPrefix.IsUnknown() {
		body["log-prefix"] = plan.LogPrefix.ValueString()
	}
	if !plan.Nth.Equal(state.Nth) && !plan.Nth.IsUnknown() {
		body["nth"] = plan.Nth.ValueString()
	}
	if !plan.OutBridgePort.Equal(state.OutBridgePort) && !plan.OutBridgePort.IsUnknown() {
		body["out-bridge-port"] = plan.OutBridgePort.ValueString()
	}
	if !plan.OutBridgePortList.Equal(state.OutBridgePortList) && !plan.OutBridgePortList.IsUnknown() {
		body["out-bridge-port-list"] = plan.OutBridgePortList.ValueString()
	}
	if !plan.OutInterface.Equal(state.OutInterface) && !plan.OutInterface.IsUnknown() {
		body["out-interface"] = plan.OutInterface.ValueString()
	}
	if !plan.OutInterfaceList.Equal(state.OutInterfaceList) && !plan.OutInterfaceList.IsUnknown() {
		body["out-interface-list"] = plan.OutInterfaceList.ValueString()
	}
	if !plan.PacketMark.Equal(state.PacketMark) && !plan.PacketMark.IsUnknown() {
		body["packet-mark"] = plan.PacketMark.ValueString()
	}
	if !plan.PacketSize.Equal(state.PacketSize) && !plan.PacketSize.IsUnknown() {
		body["packet-size"] = plan.PacketSize.ValueString()
	}
	if !plan.PerConnectionClassifier.Equal(state.PerConnectionClassifier) && !plan.PerConnectionClassifier.IsUnknown() {
		body["per-connection-classifier"] = plan.PerConnectionClassifier.ValueString()
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.Priority.Equal(state.Priority) && !plan.Priority.IsUnknown() {
		body["priority"] = plan.Priority.ValueString()
	}
	if !plan.Protocol.Equal(state.Protocol) && !plan.Protocol.IsUnknown() {
		body["protocol"] = plan.Protocol.ValueString()
	}
	if !plan.Psd.Equal(state.Psd) && !plan.Psd.IsUnknown() {
		body["psd"] = plan.Psd.ValueString()
	}
	if !plan.Random.Equal(state.Random) && !plan.Random.IsUnknown() {
		body["random"] = plan.Random.ValueString()
	}
	if !plan.Realm.Equal(state.Realm) && !plan.Realm.IsUnknown() {
		body["realm"] = plan.Realm.ValueString()
	}
	if !plan.RoutingMark.Equal(state.RoutingMark) && !plan.RoutingMark.IsUnknown() {
		body["routing-mark"] = plan.RoutingMark.ValueString()
	}
	if !plan.SniffId.Equal(state.SniffId) && !plan.SniffId.IsUnknown() {
		body["sniff-id"] = plan.SniffId.ValueString()
	}
	if !plan.SniffTarget.Equal(state.SniffTarget) && !plan.SniffTarget.IsUnknown() {
		body["sniff-target"] = plan.SniffTarget.ValueString()
	}
	if !plan.SniffTargetPort.Equal(state.SniffTargetPort) && !plan.SniffTargetPort.IsUnknown() {
		body["sniff-target-port"] = plan.SniffTargetPort.ValueString()
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) && !plan.SrcAddress.IsUnknown() {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !plan.SrcAddressList.Equal(state.SrcAddressList) && !plan.SrcAddressList.IsUnknown() {
		body["src-address-list"] = plan.SrcAddressList.ValueString()
	}
	if !plan.SrcAddressType.Equal(state.SrcAddressType) && !plan.SrcAddressType.IsUnknown() {
		body["src-address-type"] = plan.SrcAddressType.ValueString()
	}
	if !plan.SrcMacAddress.Equal(state.SrcMacAddress) && !plan.SrcMacAddress.IsUnknown() {
		body["src-mac-address"] = plan.SrcMacAddress.ValueString()
	}
	if !plan.SrcPort.Equal(state.SrcPort) && !plan.SrcPort.IsUnknown() {
		body["src-port"] = plan.SrcPort.ValueString()
	}
	if !plan.TcpMss.Equal(state.TcpMss) && !plan.TcpMss.IsUnknown() {
		body["tcp-mss"] = plan.TcpMss.ValueString()
	}
	if !plan.Time.Equal(state.Time) && !plan.Time.IsUnknown() {
		body["time"] = plan.Time.ValueString()
	}
	if !plan.TlsHost.Equal(state.TlsHost) && !plan.TlsHost.IsUnknown() {
		body["tls-host"] = plan.TlsHost.ValueString()
	}
	if !plan.Tos.Equal(state.Tos) && !plan.Tos.IsUnknown() {
		body["tos"] = plan.Tos.ValueString()
	}
	if !plan.Ttl.Equal(state.Ttl) && !plan.Ttl.IsUnknown() {
		body["ttl"] = plan.Ttl.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/firewall/calea", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/firewall/calea failed", err.Error())
			return
		}
		iPFirewallCaleaApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPFirewallCaleaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPFirewallCaleaModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/firewall/calea", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/firewall/calea failed", err.Error())
	}
}

func (r *IPFirewallCaleaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPFirewallCaleaLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/firewall/calea matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPFirewallCaleaLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPFirewallCaleaLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/firewall/calea", id)
}

func iPFirewallCaleaApply(ctx context.Context, obj client.Object, m *IPFirewallCaleaModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["ttl"]; ok && v != "" {
		m.Ttl = types.StringValue(v)
	} else {
		m.Ttl = types.StringNull()
	}
	if v, ok := obj["tos"]; ok && v != "" {
		m.Tos = types.StringValue(v)
	} else {
		m.Tos = types.StringNull()
	}
	if v, ok := obj["tls-host"]; ok && v != "" {
		m.TlsHost = types.StringValue(v)
	} else {
		m.TlsHost = types.StringNull()
	}
	if v, ok := obj["time"]; ok && v != "" {
		m.Time = types.StringValue(v)
	} else {
		m.Time = types.StringNull()
	}
	if v, ok := obj["tcp-mss"]; ok && v != "" {
		m.TcpMss = types.StringValue(v)
	} else {
		m.TcpMss = types.StringNull()
	}
	if v, ok := obj["src-port"]; ok && v != "" {
		m.SrcPort = types.StringValue(v)
	} else {
		m.SrcPort = types.StringNull()
	}
	if v, ok := obj["src-mac-address"]; ok && v != "" {
		m.SrcMacAddress = types.StringValue(v)
	} else {
		m.SrcMacAddress = types.StringNull()
	}
	if v, ok := obj["src-address-type"]; ok && v != "" {
		m.SrcAddressType = types.StringValue(v)
	} else {
		m.SrcAddressType = types.StringNull()
	}
	if v, ok := obj["src-address-list"]; ok && v != "" {
		m.SrcAddressList = types.StringValue(v)
	} else {
		m.SrcAddressList = types.StringNull()
	}
	if v, ok := obj["src-address"]; ok && v != "" {
		m.SrcAddress = types.StringValue(v)
	} else {
		m.SrcAddress = types.StringNull()
	}
	if v, ok := obj["sniff-target-port"]; ok && v != "" {
		m.SniffTargetPort = types.StringValue(v)
	} else {
		m.SniffTargetPort = types.StringNull()
	}
	if v, ok := obj["sniff-target"]; ok && v != "" {
		m.SniffTarget = types.StringValue(v)
	} else {
		m.SniffTarget = types.StringNull()
	}
	if v, ok := obj["sniff-id"]; ok && v != "" {
		m.SniffId = types.StringValue(v)
	} else {
		m.SniffId = types.StringNull()
	}
	if v, ok := obj["routing-mark"]; ok && v != "" {
		m.RoutingMark = types.StringValue(v)
	} else {
		m.RoutingMark = types.StringNull()
	}
	if v, ok := obj["realm"]; ok && v != "" {
		m.Realm = types.StringValue(v)
	} else {
		m.Realm = types.StringNull()
	}
	if v, ok := obj["random"]; ok && v != "" {
		m.Random = types.StringValue(v)
	} else {
		m.Random = types.StringNull()
	}
	if v, ok := obj["psd"]; ok && v != "" {
		m.Psd = types.StringValue(v)
	} else {
		m.Psd = types.StringNull()
	}
	if v, ok := obj["protocol"]; ok && v != "" {
		m.Protocol = types.StringValue(v)
	} else {
		m.Protocol = types.StringNull()
	}
	if v, ok := obj["priority"]; ok && v != "" {
		m.Priority = types.StringValue(v)
	} else {
		m.Priority = types.StringNull()
	}
	if v, ok := obj["port"]; ok && v != "" {
		m.Port = types.StringValue(v)
	} else {
		m.Port = types.StringNull()
	}
	if v, ok := obj["per-connection-classifier"]; ok && v != "" {
		m.PerConnectionClassifier = types.StringValue(v)
	} else {
		m.PerConnectionClassifier = types.StringNull()
	}
	if v, ok := obj["packet-size"]; ok && v != "" {
		m.PacketSize = types.StringValue(v)
	} else {
		m.PacketSize = types.StringNull()
	}
	if v, ok := obj["packet-mark"]; ok && v != "" {
		m.PacketMark = types.StringValue(v)
	} else {
		m.PacketMark = types.StringNull()
	}
	if v, ok := obj["out-interface-list"]; ok && v != "" {
		m.OutInterfaceList = types.StringValue(v)
	} else {
		m.OutInterfaceList = types.StringNull()
	}
	if v, ok := obj["out-interface"]; ok && v != "" {
		m.OutInterface = types.StringValue(v)
	} else {
		m.OutInterface = types.StringNull()
	}
	if v, ok := obj["out-bridge-port-list"]; ok && v != "" {
		m.OutBridgePortList = types.StringValue(v)
	} else {
		m.OutBridgePortList = types.StringNull()
	}
	if v, ok := obj["out-bridge-port"]; ok && v != "" {
		m.OutBridgePort = types.StringValue(v)
	} else {
		m.OutBridgePort = types.StringNull()
	}
	if v, ok := obj["nth"]; ok && v != "" {
		m.Nth = types.StringValue(v)
	} else {
		m.Nth = types.StringNull()
	}
	if v, ok := obj["log-prefix"]; ok && v != "" {
		m.LogPrefix = types.StringValue(v)
	} else {
		m.LogPrefix = types.StringNull()
	}
	if v, ok := obj["log"]; ok && v != "" {
		m.Log = types.StringValue(v)
	} else {
		m.Log = types.StringNull()
	}
	if v, ok := obj["limit"]; ok && v != "" {
		m.Limit = types.StringValue(v)
	} else {
		m.Limit = types.StringNull()
	}
	if v, ok := obj["layer7-protocol"]; ok && v != "" {
		m.Layer7Protocol = types.StringValue(v)
	} else {
		m.Layer7Protocol = types.StringNull()
	}
	if v, ok := obj["ipv4-options"]; ok && v != "" {
		m.Ipv4Options = types.StringValue(v)
	} else {
		m.Ipv4Options = types.StringNull()
	}
	if v, ok := obj["ipsec-policy"]; ok && v != "" {
		m.IpsecPolicy = types.StringValue(v)
	} else {
		m.IpsecPolicy = types.StringNull()
	}
	if v, ok := obj["ingress-priority"]; ok && v != "" {
		m.IngressPriority = types.StringValue(v)
	} else {
		m.IngressPriority = types.StringNull()
	}
	if v, ok := obj["in-interface-list"]; ok && v != "" {
		m.InInterfaceList = types.StringValue(v)
	} else {
		m.InInterfaceList = types.StringNull()
	}
	if v, ok := obj["in-interface"]; ok && v != "" {
		m.InInterface = types.StringValue(v)
	} else {
		m.InInterface = types.StringNull()
	}
	if v, ok := obj["in-bridge-port-list"]; ok && v != "" {
		m.InBridgePortList = types.StringValue(v)
	} else {
		m.InBridgePortList = types.StringNull()
	}
	if v, ok := obj["in-bridge-port"]; ok && v != "" {
		m.InBridgePort = types.StringValue(v)
	} else {
		m.InBridgePort = types.StringNull()
	}
	if v, ok := obj["icmp-options"]; ok && v != "" {
		m.IcmpOptions = types.StringValue(v)
	} else {
		m.IcmpOptions = types.StringNull()
	}
	if v, ok := obj["hotspot"]; ok && v != "" {
		m.Hotspot = types.StringValue(v)
	} else {
		m.Hotspot = types.StringNull()
	}
	if v, ok := obj["fragment"]; ok && v != "" {
		m.Fragment = types.StringValue(v)
	} else {
		m.Fragment = types.StringNull()
	}
	if v, ok := obj["dst-port"]; ok && v != "" {
		m.DstPort = types.StringValue(v)
	} else {
		m.DstPort = types.StringNull()
	}
	if v, ok := obj["dst-limit"]; ok && v != "" {
		m.DstLimit = types.StringValue(v)
	} else {
		m.DstLimit = types.StringNull()
	}
	if v, ok := obj["dst-address-type"]; ok && v != "" {
		m.DstAddressType = types.StringValue(v)
	} else {
		m.DstAddressType = types.StringNull()
	}
	if v, ok := obj["dst-address-list"]; ok && v != "" {
		m.DstAddressList = types.StringValue(v)
	} else {
		m.DstAddressList = types.StringNull()
	}
	if v, ok := obj["dst-address"]; ok && v != "" {
		m.DstAddress = types.StringValue(v)
	} else {
		m.DstAddress = types.StringNull()
	}
	if v, ok := obj["dscp"]; ok && v != "" {
		m.Dscp = types.StringValue(v)
	} else {
		m.Dscp = types.StringNull()
	}
	if v, ok := obj["content"]; ok && v != "" {
		m.Content = types.StringValue(v)
	} else {
		m.Content = types.StringNull()
	}
	if v, ok := obj["connection-type"]; ok && v != "" {
		m.ConnectionType = types.StringValue(v)
	} else {
		m.ConnectionType = types.StringNull()
	}
	if v, ok := obj["connection-rate"]; ok && v != "" {
		m.ConnectionRate = types.StringValue(v)
	} else {
		m.ConnectionRate = types.StringNull()
	}
	if v, ok := obj["connection-mark"]; ok && v != "" {
		m.ConnectionMark = types.StringValue(v)
	} else {
		m.ConnectionMark = types.StringNull()
	}
	if v, ok := obj["connection-limit"]; ok && v != "" {
		m.ConnectionLimit = types.StringValue(v)
	} else {
		m.ConnectionLimit = types.StringNull()
	}
	if v, ok := obj["connection-bytes"]; ok && v != "" {
		m.ConnectionBytes = types.StringValue(v)
	} else {
		m.ConnectionBytes = types.StringNull()
	}
	if v, ok := obj["chain"]; ok && v != "" {
		m.Chain = types.StringValue(v)
	} else {
		m.Chain = types.StringNull()
	}
	if v, ok := obj["address-list-timeout"]; ok && v != "" {
		m.AddressListTimeout = types.StringValue(v)
	} else {
		m.AddressListTimeout = types.StringNull()
	}
	if v, ok := obj["address-list"]; ok && v != "" {
		m.AddressList = types.StringValue(v)
	} else {
		m.AddressList = types.StringNull()
	}
	if v, ok := obj["action"]; ok && v != "" {
		m.Action = types.StringValue(v)
	} else {
		m.Action = types.StringNull()
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
}
