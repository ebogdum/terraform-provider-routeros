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
	"github.com/ebogdum/terraform-provider-routeros/internal/schemautil"
)

var (
	_ resource.Resource                = &IPV6FirewallFilterResource{}
	_ resource.ResourceWithImportState = &IPV6FirewallFilterResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPV6FirewallFilterResource struct {
	reg *client.Registry
}

type IPV6FirewallFilterModel struct {
	ID                      types.String `tfsdk:"id"`
	Tos                     types.String `tfsdk:"tos"`
	Headers                 types.String `tfsdk:"headers"`
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
	HopLimit                types.String `tfsdk:"hop_limit"`
	IcmpOptions             types.String `tfsdk:"icmp_options"`
	InBridgePort            types.String `tfsdk:"in_bridge_port"`
	InBridgePortList        types.String `tfsdk:"in_bridge_port_list"`
	InInterface             types.String `tfsdk:"in_interface"`
	InInterfaceList         types.String `tfsdk:"in_interface_list"`
	IngressPriority         types.String `tfsdk:"ingress_priority"`
	IpsecPolicy             types.String `tfsdk:"ipsec_policy"`
	JumpTarget              types.String `tfsdk:"jump_target"`
	Limit                   types.String `tfsdk:"limit"`
	Log                     types.String `tfsdk:"log"`
	LogPrefix               types.String `tfsdk:"log_prefix"`
	Nth                     types.String `tfsdk:"nth"`
	OutBridgePort           types.String `tfsdk:"out_bridge_port"`
	OutBridgePortList       types.String `tfsdk:"out_bridge_port_list"`
	OutInterface            types.String `tfsdk:"out_interface"`
	OutInterfaceList        types.String `tfsdk:"out_interface_list"`
	PacketMark              types.String `tfsdk:"packet_mark"`
	PacketSize              types.String `tfsdk:"packet_size"`
	Packets                 types.String `tfsdk:"packets"`
	PerConnectionClassifier types.String `tfsdk:"per_connection_classifier"`
	Port                    types.String `tfsdk:"port"`
	Priority                types.String `tfsdk:"priority"`
	Protocol                types.String `tfsdk:"protocol"`
	Random                  types.String `tfsdk:"random"`
	RejectWith              types.String `tfsdk:"reject_with"`
	RoutingMark             types.String `tfsdk:"routing_mark"`
	SrcAddress              types.String `tfsdk:"src_address"`
	SrcAddressList          types.String `tfsdk:"src_address_list"`
	SrcAddressType          types.String `tfsdk:"src_address_type"`
	SrcMACAddress           types.String `tfsdk:"src_mac_address"`
	SrcPort                 types.String `tfsdk:"src_port"`
	TCPFlags                types.String `tfsdk:"tcp_flags"`
	TCPMss                  types.String `tfsdk:"tcp_mss"`
	Time                    csvSetValue  `tfsdk:"time"`
	TLSHost                 types.String `tfsdk:"tls_host"`
	Router                  types.String `tfsdk:"router"`
	Position                types.Int64  `tfsdk:"position"`
	LockoutAck              types.Bool   `tfsdk:"lockout_ack"`
}

func NewIPV6FirewallFilterResource() resource.Resource { return &IPV6FirewallFilterResource{} }

func (r *IPV6FirewallFilterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv6_firewall_filter"
}

func (r *IPV6FirewallFilterResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPV6FirewallFilterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ipv6/firewall/filter`.",
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
			"headers": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `headers`.",
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
			"hop_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Matches the IPv6 hop limit, as `<condition>:<value>` -- e.g. `equal:1`, `not-equal:1`, `less-than:5`, `greater-than:2`. Required to express the defconf rfc4890 rule; without it that rule matches all ICMPv6 instead of only hop-limit 1.",
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
			"jump_target": schema.StringAttribute{
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
			"random": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"reject_with": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"routing_mark": schema.StringAttribute{
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
				CustomType:  csvSetType{},
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tls_host": schema.StringAttribute{
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
			"lockout_ack": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Acknowledge that this rule may sever management traffic (required for unconditional input/forward drop/reject/tarpit rules with no match).",
			},
		},
	}
}

func (r *IPV6FirewallFilterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPV6FirewallFilterModel
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
	if !(plan.HopLimit.IsNull() || plan.HopLimit.IsUnknown()) {
		body["hop-limit"] = plan.HopLimit.ValueString()
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
	if !(plan.JumpTarget.IsNull() || plan.JumpTarget.IsUnknown()) {
		body["jump-target"] = plan.JumpTarget.ValueString()
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
	if !(plan.Random.IsNull() || plan.Random.IsUnknown()) {
		body["random"] = plan.Random.ValueString()
	}
	if !(plan.RejectWith.IsNull() || plan.RejectWith.IsUnknown()) {
		body["reject-with"] = plan.RejectWith.ValueString()
	}
	if !(plan.RoutingMark.IsNull() || plan.RoutingMark.IsUnknown()) {
		body["routing-mark"] = plan.RoutingMark.ValueString()
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
	if err := schemautil.CheckFirewallLockout("/ipv6/firewall/filter", body, !plan.LockoutAck.IsNull() && plan.LockoutAck.ValueBool()); err != nil {
		resp.Diagnostics.AddError("Refusing dangerous firewall rule", err.Error())
		return
	}
	if !(plan.Headers.IsNull() || plan.Headers.IsUnknown()) {
		body["headers"] = plan.Headers.ValueString()
	}
	if !(plan.Tos.IsNull() || plan.Tos.IsUnknown()) {
		body["tos"] = plan.Tos.ValueString()
	}
	obj, err := c.Add(ctx, "/ipv6/firewall/filter", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ipv6/firewall/filter failed", err.Error())
		return
	}
	if !plan.Position.IsNull() && !plan.Position.IsUnknown() {
		r.reg.RegisterOrdered(plan.Router.ValueString(), "/ipv6/firewall/filter", obj[".id"], plan.Position.ValueInt64())
		snap := r.reg.OrderedSnapshot(plan.Router.ValueString(), "/ipv6/firewall/filter")
		if err := c.PlaceOrdered(ctx, plan.Router.ValueString(), "/ipv6/firewall/filter", obj[".id"], plan.Position.ValueInt64(), snap); err != nil {
			// Resource exists on the device; write minimal state so Terraform
			// tracks it (and a future apply can repair the order or delete it)
			// instead of creating a duplicate.
			iPV6FirewallFilterApply(ctx, obj, &plan)
			nullifyUnknownAttrs(&plan)
			resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
			resp.Diagnostics.AddError("Order /ipv6/firewall/filter failed", err.Error())
			return
		}
		reread, rerr := c.GetByID(ctx, "/ipv6/firewall/filter", obj[".id"])
		if rerr != nil {
			iPV6FirewallFilterApply(ctx, obj, &plan)
			nullifyUnknownAttrs(&plan)
			resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
			resp.Diagnostics.AddError("Re-read after order failed", rerr.Error())
			return
		}
		obj = reread
	}
	iPV6FirewallFilterApply(ctx, obj, &plan)
	// Apply has already split the marker off the comment; carry position
	// from plan into state explicitly because the wire shows the encoded form.
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6FirewallFilterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPV6FirewallFilterModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ipv6/firewall/filter", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ipv6/firewall/filter failed", err.Error())
		return
	}
	iPV6FirewallFilterApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPV6FirewallFilterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPV6FirewallFilterModel
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
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
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
	if !plan.ConnectionNATState.Equal(state.ConnectionNATState) && !plan.ConnectionNATState.IsUnknown() {
		body["connection-nat-state"] = plan.ConnectionNATState.ValueString()
	}
	if !plan.ConnectionRate.Equal(state.ConnectionRate) && !plan.ConnectionRate.IsUnknown() {
		body["connection-rate"] = plan.ConnectionRate.ValueString()
	}
	if !plan.ConnectionState.Equal(state.ConnectionState) && !plan.ConnectionState.IsUnknown() {
		body["connection-state"] = plan.ConnectionState.ValueString()
	}
	if !plan.ConnectionType.Equal(state.ConnectionType) && !plan.ConnectionType.IsUnknown() {
		body["connection-type"] = plan.ConnectionType.ValueString()
	}
	if !plan.Content.Equal(state.Content) && !plan.Content.IsUnknown() {
		body["content"] = plan.Content.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
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
	if !plan.HopLimit.Equal(state.HopLimit) && !plan.HopLimit.IsUnknown() {
		body["hop-limit"] = plan.HopLimit.ValueString()
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
	if !plan.JumpTarget.Equal(state.JumpTarget) && !plan.JumpTarget.IsUnknown() {
		body["jump-target"] = plan.JumpTarget.ValueString()
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
	if !plan.Random.Equal(state.Random) && !plan.Random.IsUnknown() {
		body["random"] = plan.Random.ValueString()
	}
	if !plan.RejectWith.Equal(state.RejectWith) && !plan.RejectWith.IsUnknown() {
		body["reject-with"] = plan.RejectWith.ValueString()
	}
	if !plan.RoutingMark.Equal(state.RoutingMark) && !plan.RoutingMark.IsUnknown() {
		body["routing-mark"] = plan.RoutingMark.ValueString()
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
	if !plan.SrcMACAddress.Equal(state.SrcMACAddress) && !plan.SrcMACAddress.IsUnknown() {
		body["src-mac-address"] = plan.SrcMACAddress.ValueString()
	}
	if !plan.SrcPort.Equal(state.SrcPort) && !plan.SrcPort.IsUnknown() {
		body["src-port"] = plan.SrcPort.ValueString()
	}
	if !plan.TCPFlags.Equal(state.TCPFlags) && !plan.TCPFlags.IsUnknown() {
		body["tcp-flags"] = plan.TCPFlags.ValueString()
	}
	if !plan.TCPMss.Equal(state.TCPMss) && !plan.TCPMss.IsUnknown() {
		body["tcp-mss"] = plan.TCPMss.ValueString()
	}
	if !plan.Time.Equal(state.Time) && !plan.Time.IsUnknown() {
		body["time"] = plan.Time.ValueString()
	}
	if !plan.TLSHost.Equal(state.TLSHost) && !plan.TLSHost.IsUnknown() {
		body["tls-host"] = plan.TLSHost.ValueString()
	}
	if !plan.Headers.Equal(state.Headers) && !plan.Headers.IsUnknown() {
		body["headers"] = plan.Headers.ValueString()
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
	// Build the EFFECTIVE rule (state merged with diff) so the guard sees the
	// final shape -- not just the changed fields. A diff might only set
	// action=drop without re-stating chain or matches; without merging we'd
	// short-circuit on "chain is empty" and miss the lockout.
	effective := client.Object{}
	if v, ok := body["action"]; ok {
		effective["action"] = v
	} else {
		if !state.Action.IsNull() && !state.Action.IsUnknown() {
			effective["action"] = state.Action.ValueString()
		}
	}
	if v, ok := body["address-list"]; ok {
		effective["address-list"] = v
	} else {
		if !state.AddressList.IsNull() && !state.AddressList.IsUnknown() {
			effective["address-list"] = state.AddressList.ValueString()
		}
	}
	if v, ok := body["address-list-timeout"]; ok {
		effective["address-list-timeout"] = v
	} else {
		if !state.AddressListTimeout.IsNull() && !state.AddressListTimeout.IsUnknown() {
			effective["address-list-timeout"] = state.AddressListTimeout.ValueString()
		}
	}
	if v, ok := body["chain"]; ok {
		effective["chain"] = v
	} else {
		if !state.Chain.IsNull() && !state.Chain.IsUnknown() {
			effective["chain"] = state.Chain.ValueString()
		}
	}
	if v, ok := body["comment"]; ok {
		effective["comment"] = v
	} else {
		if !state.Comment.IsNull() && !state.Comment.IsUnknown() {
			effective["comment"] = state.Comment.ValueString()
		}
	}
	if v, ok := body["connection-bytes"]; ok {
		effective["connection-bytes"] = v
	} else {
		if !state.ConnectionBytes.IsNull() && !state.ConnectionBytes.IsUnknown() {
			effective["connection-bytes"] = state.ConnectionBytes.ValueString()
		}
	}
	if v, ok := body["connection-limit"]; ok {
		effective["connection-limit"] = v
	} else {
		if !state.ConnectionLimit.IsNull() && !state.ConnectionLimit.IsUnknown() {
			effective["connection-limit"] = state.ConnectionLimit.ValueString()
		}
	}
	if v, ok := body["connection-mark"]; ok {
		effective["connection-mark"] = v
	} else {
		if !state.ConnectionMark.IsNull() && !state.ConnectionMark.IsUnknown() {
			effective["connection-mark"] = state.ConnectionMark.ValueString()
		}
	}
	if v, ok := body["connection-nat-state"]; ok {
		effective["connection-nat-state"] = v
	} else {
		if !state.ConnectionNATState.IsNull() && !state.ConnectionNATState.IsUnknown() {
			effective["connection-nat-state"] = state.ConnectionNATState.ValueString()
		}
	}
	if v, ok := body["connection-rate"]; ok {
		effective["connection-rate"] = v
	} else {
		if !state.ConnectionRate.IsNull() && !state.ConnectionRate.IsUnknown() {
			effective["connection-rate"] = state.ConnectionRate.ValueString()
		}
	}
	if v, ok := body["connection-state"]; ok {
		effective["connection-state"] = v
	} else {
		if !state.ConnectionState.IsNull() && !state.ConnectionState.IsUnknown() {
			effective["connection-state"] = state.ConnectionState.ValueString()
		}
	}
	if v, ok := body["connection-type"]; ok {
		effective["connection-type"] = v
	} else {
		if !state.ConnectionType.IsNull() && !state.ConnectionType.IsUnknown() {
			effective["connection-type"] = state.ConnectionType.ValueString()
		}
	}
	if v, ok := body["content"]; ok {
		effective["content"] = v
	} else {
		if !state.Content.IsNull() && !state.Content.IsUnknown() {
			effective["content"] = state.Content.ValueString()
		}
	}
	if v, ok := body["disabled"]; ok {
		effective["disabled"] = v
	} else {
		if !state.Disabled.IsNull() && !state.Disabled.IsUnknown() {
			effective["disabled"] = client.FormatBool(state.Disabled.ValueBool())
		}
	}
	if v, ok := body["dscp"]; ok {
		effective["dscp"] = v
	} else {
		if !state.Dscp.IsNull() && !state.Dscp.IsUnknown() {
			effective["dscp"] = state.Dscp.ValueString()
		}
	}
	if v, ok := body["dst-address"]; ok {
		effective["dst-address"] = v
	} else {
		if !state.DstAddress.IsNull() && !state.DstAddress.IsUnknown() {
			effective["dst-address"] = state.DstAddress.ValueString()
		}
	}
	if v, ok := body["dst-address-list"]; ok {
		effective["dst-address-list"] = v
	} else {
		if !state.DstAddressList.IsNull() && !state.DstAddressList.IsUnknown() {
			effective["dst-address-list"] = state.DstAddressList.ValueString()
		}
	}
	if v, ok := body["dst-address-type"]; ok {
		effective["dst-address-type"] = v
	} else {
		if !state.DstAddressType.IsNull() && !state.DstAddressType.IsUnknown() {
			effective["dst-address-type"] = state.DstAddressType.ValueString()
		}
	}
	if v, ok := body["dst-limit"]; ok {
		effective["dst-limit"] = v
	} else {
		if !state.DstLimit.IsNull() && !state.DstLimit.IsUnknown() {
			effective["dst-limit"] = state.DstLimit.ValueString()
		}
	}
	if v, ok := body["dst-port"]; ok {
		effective["dst-port"] = v
	} else {
		if !state.DstPort.IsNull() && !state.DstPort.IsUnknown() {
			effective["dst-port"] = state.DstPort.ValueString()
		}
	}
	if v, ok := body["icmp-options"]; ok {
		effective["icmp-options"] = v
	} else {
		if !state.IcmpOptions.IsNull() && !state.IcmpOptions.IsUnknown() {
			effective["icmp-options"] = state.IcmpOptions.ValueString()
		}
	}
	if v, ok := body["in-bridge-port"]; ok {
		effective["in-bridge-port"] = v
	} else {
		if !state.InBridgePort.IsNull() && !state.InBridgePort.IsUnknown() {
			effective["in-bridge-port"] = state.InBridgePort.ValueString()
		}
	}
	if v, ok := body["in-bridge-port-list"]; ok {
		effective["in-bridge-port-list"] = v
	} else {
		if !state.InBridgePortList.IsNull() && !state.InBridgePortList.IsUnknown() {
			effective["in-bridge-port-list"] = state.InBridgePortList.ValueString()
		}
	}
	if v, ok := body["in-interface"]; ok {
		effective["in-interface"] = v
	} else {
		if !state.InInterface.IsNull() && !state.InInterface.IsUnknown() {
			effective["in-interface"] = state.InInterface.ValueString()
		}
	}
	if v, ok := body["in-interface-list"]; ok {
		effective["in-interface-list"] = v
	} else {
		if !state.InInterfaceList.IsNull() && !state.InInterfaceList.IsUnknown() {
			effective["in-interface-list"] = state.InInterfaceList.ValueString()
		}
	}
	if v, ok := body["ingress-priority"]; ok {
		effective["ingress-priority"] = v
	} else {
		if !state.IngressPriority.IsNull() && !state.IngressPriority.IsUnknown() {
			effective["ingress-priority"] = state.IngressPriority.ValueString()
		}
	}
	if v, ok := body["ipsec-policy"]; ok {
		effective["ipsec-policy"] = v
	} else {
		if !state.IpsecPolicy.IsNull() && !state.IpsecPolicy.IsUnknown() {
			effective["ipsec-policy"] = state.IpsecPolicy.ValueString()
		}
	}
	if v, ok := body["jump-target"]; ok {
		effective["jump-target"] = v
	} else {
		if !state.JumpTarget.IsNull() && !state.JumpTarget.IsUnknown() {
			effective["jump-target"] = state.JumpTarget.ValueString()
		}
	}
	if v, ok := body["limit"]; ok {
		effective["limit"] = v
	} else {
		if !state.Limit.IsNull() && !state.Limit.IsUnknown() {
			effective["limit"] = state.Limit.ValueString()
		}
	}
	if v, ok := body["log"]; ok {
		effective["log"] = v
	} else {
		if !state.Log.IsNull() && !state.Log.IsUnknown() {
			effective["log"] = state.Log.ValueString()
		}
	}
	if v, ok := body["log-prefix"]; ok {
		effective["log-prefix"] = v
	} else {
		if !state.LogPrefix.IsNull() && !state.LogPrefix.IsUnknown() {
			effective["log-prefix"] = state.LogPrefix.ValueString()
		}
	}
	if v, ok := body["nth"]; ok {
		effective["nth"] = v
	} else {
		if !state.Nth.IsNull() && !state.Nth.IsUnknown() {
			effective["nth"] = state.Nth.ValueString()
		}
	}
	if v, ok := body["out-bridge-port"]; ok {
		effective["out-bridge-port"] = v
	} else {
		if !state.OutBridgePort.IsNull() && !state.OutBridgePort.IsUnknown() {
			effective["out-bridge-port"] = state.OutBridgePort.ValueString()
		}
	}
	if v, ok := body["out-bridge-port-list"]; ok {
		effective["out-bridge-port-list"] = v
	} else {
		if !state.OutBridgePortList.IsNull() && !state.OutBridgePortList.IsUnknown() {
			effective["out-bridge-port-list"] = state.OutBridgePortList.ValueString()
		}
	}
	if v, ok := body["out-interface"]; ok {
		effective["out-interface"] = v
	} else {
		if !state.OutInterface.IsNull() && !state.OutInterface.IsUnknown() {
			effective["out-interface"] = state.OutInterface.ValueString()
		}
	}
	if v, ok := body["out-interface-list"]; ok {
		effective["out-interface-list"] = v
	} else {
		if !state.OutInterfaceList.IsNull() && !state.OutInterfaceList.IsUnknown() {
			effective["out-interface-list"] = state.OutInterfaceList.ValueString()
		}
	}
	if v, ok := body["packet-mark"]; ok {
		effective["packet-mark"] = v
	} else {
		if !state.PacketMark.IsNull() && !state.PacketMark.IsUnknown() {
			effective["packet-mark"] = state.PacketMark.ValueString()
		}
	}
	if v, ok := body["packet-size"]; ok {
		effective["packet-size"] = v
	} else {
		if !state.PacketSize.IsNull() && !state.PacketSize.IsUnknown() {
			effective["packet-size"] = state.PacketSize.ValueString()
		}
	}
	if v, ok := body["per-connection-classifier"]; ok {
		effective["per-connection-classifier"] = v
	} else {
		if !state.PerConnectionClassifier.IsNull() && !state.PerConnectionClassifier.IsUnknown() {
			effective["per-connection-classifier"] = state.PerConnectionClassifier.ValueString()
		}
	}
	if v, ok := body["port"]; ok {
		effective["port"] = v
	} else {
		if !state.Port.IsNull() && !state.Port.IsUnknown() {
			effective["port"] = state.Port.ValueString()
		}
	}
	if v, ok := body["priority"]; ok {
		effective["priority"] = v
	} else {
		if !state.Priority.IsNull() && !state.Priority.IsUnknown() {
			effective["priority"] = state.Priority.ValueString()
		}
	}
	if v, ok := body["protocol"]; ok {
		effective["protocol"] = v
	} else {
		if !state.Protocol.IsNull() && !state.Protocol.IsUnknown() {
			effective["protocol"] = state.Protocol.ValueString()
		}
	}
	if v, ok := body["random"]; ok {
		effective["random"] = v
	} else {
		if !state.Random.IsNull() && !state.Random.IsUnknown() {
			effective["random"] = state.Random.ValueString()
		}
	}
	if v, ok := body["reject-with"]; ok {
		effective["reject-with"] = v
	} else {
		if !state.RejectWith.IsNull() && !state.RejectWith.IsUnknown() {
			effective["reject-with"] = state.RejectWith.ValueString()
		}
	}
	if v, ok := body["routing-mark"]; ok {
		effective["routing-mark"] = v
	} else {
		if !state.RoutingMark.IsNull() && !state.RoutingMark.IsUnknown() {
			effective["routing-mark"] = state.RoutingMark.ValueString()
		}
	}
	if v, ok := body["src-address"]; ok {
		effective["src-address"] = v
	} else {
		if !state.SrcAddress.IsNull() && !state.SrcAddress.IsUnknown() {
			effective["src-address"] = state.SrcAddress.ValueString()
		}
	}
	if v, ok := body["src-address-list"]; ok {
		effective["src-address-list"] = v
	} else {
		if !state.SrcAddressList.IsNull() && !state.SrcAddressList.IsUnknown() {
			effective["src-address-list"] = state.SrcAddressList.ValueString()
		}
	}
	if v, ok := body["src-address-type"]; ok {
		effective["src-address-type"] = v
	} else {
		if !state.SrcAddressType.IsNull() && !state.SrcAddressType.IsUnknown() {
			effective["src-address-type"] = state.SrcAddressType.ValueString()
		}
	}
	if v, ok := body["src-mac-address"]; ok {
		effective["src-mac-address"] = v
	} else {
		if !state.SrcMACAddress.IsNull() && !state.SrcMACAddress.IsUnknown() {
			effective["src-mac-address"] = state.SrcMACAddress.ValueString()
		}
	}
	if v, ok := body["src-port"]; ok {
		effective["src-port"] = v
	} else {
		if !state.SrcPort.IsNull() && !state.SrcPort.IsUnknown() {
			effective["src-port"] = state.SrcPort.ValueString()
		}
	}
	if v, ok := body["tcp-flags"]; ok {
		effective["tcp-flags"] = v
	} else {
		if !state.TCPFlags.IsNull() && !state.TCPFlags.IsUnknown() {
			effective["tcp-flags"] = state.TCPFlags.ValueString()
		}
	}
	if v, ok := body["tcp-mss"]; ok {
		effective["tcp-mss"] = v
	} else {
		if !state.TCPMss.IsNull() && !state.TCPMss.IsUnknown() {
			effective["tcp-mss"] = state.TCPMss.ValueString()
		}
	}
	if v, ok := body["time"]; ok {
		effective["time"] = v
	} else {
		if !state.Time.IsNull() && !state.Time.IsUnknown() {
			effective["time"] = state.Time.ValueString()
		}
	}
	if v, ok := body["tls-host"]; ok {
		effective["tls-host"] = v
	} else {
		if !state.TLSHost.IsNull() && !state.TLSHost.IsUnknown() {
			effective["tls-host"] = state.TLSHost.ValueString()
		}
	}
	if err := schemautil.CheckFirewallLockout("/ipv6/firewall/filter", effective, !plan.LockoutAck.IsNull() && plan.LockoutAck.ValueBool()); err != nil {
		resp.Diagnostics.AddError("Refusing dangerous firewall rule", err.Error())
		return
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ipv6/firewall/filter", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ipv6/firewall/filter failed", err.Error())
			return
		}
		iPV6FirewallFilterApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	if !plan.Position.IsNull() && !plan.Position.IsUnknown() {
		r.reg.RegisterOrdered(plan.Router.ValueString(), "/ipv6/firewall/filter", plan.ID.ValueString(), plan.Position.ValueInt64())
		if !plan.Position.Equal(state.Position) {
			snap := r.reg.OrderedSnapshot(plan.Router.ValueString(), "/ipv6/firewall/filter")
			if err := c.PlaceOrdered(ctx, plan.Router.ValueString(), "/ipv6/firewall/filter", plan.ID.ValueString(), plan.Position.ValueInt64(), snap); err != nil {
				// Set already applied; persist the new attributes so state matches the
				// device even if the re-order failed.
				nullifyUnknownAttrs(&plan)
				resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
				resp.Diagnostics.AddError("Order /ipv6/firewall/filter failed", err.Error())
				return
			}
		}
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6FirewallFilterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPV6FirewallFilterModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ipv6/firewall/filter", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ipv6/firewall/filter failed", err.Error())
	}
	r.reg.UnregisterOrdered(state.Router.ValueString(), "/ipv6/firewall/filter", state.ID.ValueString())
}

func (r *IPV6FirewallFilterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPV6FirewallFilterLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ipv6/firewall/filter matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPV6FirewallFilterLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPV6FirewallFilterLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ipv6/firewall/filter", id)
}

func iPV6FirewallFilterApply(ctx context.Context, obj client.Object, m *IPV6FirewallFilterModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["tos"]; ok && v != "" {
		m.Tos = types.StringValue(v)
	} else {
		m.Tos = types.StringNull()
	}
	if v, ok := obj["headers"]; ok && v != "" {
		m.Headers = types.StringValue(v)
	} else {
		m.Headers = types.StringNull()
	}
	// Strip the [tf:pos=N] marker from the comment before exposing to state.
	// Position is TF-state-only metadata; never written to the device. Keep
	// whatever the user planned. Comment is left untouched.
	if m.Position.IsUnknown() {
		m.Position = types.Int64Null()
	}
	// LockoutAck is local-only and not persisted on the wire. Carry the
	// plan's value through to state (Null if the user didn't set it).
	if m.LockoutAck.IsUnknown() {
		m.LockoutAck = types.BoolNull()
	}
	if v, ok := obj["action"]; ok {
		if v != "" {
			m.Action = types.StringValue(v)
		} else {
			m.Action = types.StringNull()
		}
	}
	if v, ok := obj["address-list"]; ok {
		if v != "" {
			m.AddressList = types.StringValue(v)
		} else {
			m.AddressList = types.StringNull()
		}
	}
	if v, ok := obj["address-list-timeout"]; ok {
		if v != "" {
			m.AddressListTimeout = types.StringValue(v)
		} else {
			m.AddressListTimeout = types.StringNull()
		}
	}
	if v, ok := obj["bytes"]; ok {
		if v != "" {
			m.Bytes = types.StringValue(v)
		} else {
			m.Bytes = types.StringNull()
		}
	}
	if v, ok := obj["chain"]; ok {
		if v != "" {
			m.Chain = types.StringValue(v)
		} else {
			m.Chain = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["connection-bytes"]; ok {
		if v != "" {
			m.ConnectionBytes = types.StringValue(v)
		} else {
			m.ConnectionBytes = types.StringNull()
		}
	}
	if v, ok := obj["connection-limit"]; ok {
		if v != "" {
			m.ConnectionLimit = types.StringValue(v)
		} else {
			m.ConnectionLimit = types.StringNull()
		}
	}
	if v, ok := obj["connection-mark"]; ok {
		if v != "" {
			m.ConnectionMark = types.StringValue(v)
		} else {
			m.ConnectionMark = types.StringNull()
		}
	}
	if v, ok := obj["connection-nat-state"]; ok {
		if v != "" {
			m.ConnectionNATState = types.StringValue(v)
		} else {
			m.ConnectionNATState = types.StringNull()
		}
	}
	if v, ok := obj["connection-rate"]; ok {
		if v != "" {
			m.ConnectionRate = types.StringValue(v)
		} else {
			m.ConnectionRate = types.StringNull()
		}
	}
	if v, ok := obj["connection-state"]; ok {
		if v != "" {
			m.ConnectionState = types.StringValue(v)
		} else {
			m.ConnectionState = types.StringNull()
		}
	}
	if v, ok := obj["connection-type"]; ok {
		if v != "" {
			m.ConnectionType = types.StringValue(v)
		} else {
			m.ConnectionType = types.StringNull()
		}
	}
	if v, ok := obj["content"]; ok {
		if v != "" {
			m.Content = types.StringValue(v)
		} else {
			m.Content = types.StringNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["dscp"]; ok {
		if v != "" {
			m.Dscp = types.StringValue(v)
		} else {
			m.Dscp = types.StringNull()
		}
	}
	if v, ok := obj["dst-address"]; ok {
		if v != "" {
			m.DstAddress = types.StringValue(v)
		} else {
			m.DstAddress = types.StringNull()
		}
	}
	if v, ok := obj["dst-address-list"]; ok {
		if v != "" {
			m.DstAddressList = types.StringValue(v)
		} else {
			m.DstAddressList = types.StringNull()
		}
	}
	if v, ok := obj["dst-address-type"]; ok {
		if v != "" {
			m.DstAddressType = types.StringValue(v)
		} else {
			m.DstAddressType = types.StringNull()
		}
	}
	if v, ok := obj["dst-limit"]; ok {
		if v != "" {
			m.DstLimit = types.StringValue(v)
		} else {
			m.DstLimit = types.StringNull()
		}
	}
	if v, ok := obj["dst-port"]; ok {
		if v != "" {
			m.DstPort = types.StringValue(v)
		} else {
			m.DstPort = types.StringNull()
		}
	}
	if v, ok := obj["hop-limit"]; ok {
		if v != "" {
			m.HopLimit = types.StringValue(v)
		} else {
			m.HopLimit = types.StringNull()
		}
	}
	if v, ok := obj["icmp-options"]; ok {
		if v != "" {
			m.IcmpOptions = types.StringValue(v)
		} else {
			m.IcmpOptions = types.StringNull()
		}
	}
	if v, ok := obj["in-bridge-port"]; ok {
		if v != "" {
			m.InBridgePort = types.StringValue(v)
		} else {
			m.InBridgePort = types.StringNull()
		}
	}
	if v, ok := obj["in-bridge-port-list"]; ok {
		if v != "" {
			m.InBridgePortList = types.StringValue(v)
		} else {
			m.InBridgePortList = types.StringNull()
		}
	}
	if v, ok := obj["in-interface"]; ok {
		if v != "" {
			m.InInterface = types.StringValue(v)
		} else {
			m.InInterface = types.StringNull()
		}
	}
	if v, ok := obj["in-interface-list"]; ok {
		if v != "" {
			m.InInterfaceList = types.StringValue(v)
		} else {
			m.InInterfaceList = types.StringNull()
		}
	}
	if v, ok := obj["ingress-priority"]; ok {
		if v != "" {
			m.IngressPriority = types.StringValue(v)
		} else {
			m.IngressPriority = types.StringNull()
		}
	}
	if v, ok := obj["ipsec-policy"]; ok {
		if v != "" {
			m.IpsecPolicy = types.StringValue(v)
		} else {
			m.IpsecPolicy = types.StringNull()
		}
	}
	if v, ok := obj["jump-target"]; ok {
		if v != "" {
			m.JumpTarget = types.StringValue(v)
		} else {
			m.JumpTarget = types.StringNull()
		}
	}
	if v, ok := obj["limit"]; ok {
		if v != "" {
			m.Limit = types.StringValue(v)
		} else {
			m.Limit = types.StringNull()
		}
	}
	if v, ok := obj["log"]; ok {
		if v != "" {
			m.Log = types.StringValue(v)
		} else {
			m.Log = types.StringNull()
		}
	}
	if v, ok := obj["log-prefix"]; ok {
		if v != "" {
			m.LogPrefix = types.StringValue(v)
		} else {
			m.LogPrefix = types.StringNull()
		}
	}
	if v, ok := obj["nth"]; ok {
		if v != "" {
			m.Nth = types.StringValue(v)
		} else {
			m.Nth = types.StringNull()
		}
	}
	if v, ok := obj["out-bridge-port"]; ok {
		if v != "" {
			m.OutBridgePort = types.StringValue(v)
		} else {
			m.OutBridgePort = types.StringNull()
		}
	}
	if v, ok := obj["out-bridge-port-list"]; ok {
		if v != "" {
			m.OutBridgePortList = types.StringValue(v)
		} else {
			m.OutBridgePortList = types.StringNull()
		}
	}
	if v, ok := obj["out-interface"]; ok {
		if v != "" {
			m.OutInterface = types.StringValue(v)
		} else {
			m.OutInterface = types.StringNull()
		}
	}
	if v, ok := obj["out-interface-list"]; ok {
		if v != "" {
			m.OutInterfaceList = types.StringValue(v)
		} else {
			m.OutInterfaceList = types.StringNull()
		}
	}
	if v, ok := obj["packet-mark"]; ok {
		if v != "" {
			m.PacketMark = types.StringValue(v)
		} else {
			m.PacketMark = types.StringNull()
		}
	}
	if v, ok := obj["packet-size"]; ok {
		if v != "" {
			m.PacketSize = types.StringValue(v)
		} else {
			m.PacketSize = types.StringNull()
		}
	}
	if v, ok := obj["packets"]; ok {
		if v != "" {
			m.Packets = types.StringValue(v)
		} else {
			m.Packets = types.StringNull()
		}
	}
	if v, ok := obj["per-connection-classifier"]; ok {
		if v != "" {
			m.PerConnectionClassifier = types.StringValue(v)
		} else {
			m.PerConnectionClassifier = types.StringNull()
		}
	}
	if v, ok := obj["port"]; ok {
		if v != "" {
			m.Port = types.StringValue(v)
		} else {
			m.Port = types.StringNull()
		}
	}
	if v, ok := obj["priority"]; ok {
		if v != "" {
			m.Priority = types.StringValue(v)
		} else {
			m.Priority = types.StringNull()
		}
	}
	if v, ok := obj["protocol"]; ok {
		if v != "" {
			m.Protocol = types.StringValue(v)
		} else {
			m.Protocol = types.StringNull()
		}
	}
	if v, ok := obj["random"]; ok {
		if v != "" {
			m.Random = types.StringValue(v)
		} else {
			m.Random = types.StringNull()
		}
	}
	if v, ok := obj["reject-with"]; ok {
		if v != "" {
			m.RejectWith = types.StringValue(v)
		} else {
			m.RejectWith = types.StringNull()
		}
	}
	if v, ok := obj["routing-mark"]; ok {
		if v != "" {
			m.RoutingMark = types.StringValue(v)
		} else {
			m.RoutingMark = types.StringNull()
		}
	}
	if v, ok := obj["src-address"]; ok {
		if v != "" {
			m.SrcAddress = types.StringValue(v)
		} else {
			m.SrcAddress = types.StringNull()
		}
	}
	if v, ok := obj["src-address-list"]; ok {
		if v != "" {
			m.SrcAddressList = types.StringValue(v)
		} else {
			m.SrcAddressList = types.StringNull()
		}
	}
	if v, ok := obj["src-address-type"]; ok {
		if v != "" {
			m.SrcAddressType = types.StringValue(v)
		} else {
			m.SrcAddressType = types.StringNull()
		}
	}
	if v, ok := obj["src-mac-address"]; ok {
		if v != "" {
			m.SrcMACAddress = types.StringValue(v)
		} else {
			m.SrcMACAddress = types.StringNull()
		}
	}
	if v, ok := obj["src-port"]; ok {
		if v != "" {
			m.SrcPort = types.StringValue(v)
		} else {
			m.SrcPort = types.StringNull()
		}
	}
	if v, ok := obj["tcp-flags"]; ok {
		if v != "" {
			m.TCPFlags = types.StringValue(v)
		} else {
			m.TCPFlags = types.StringNull()
		}
	}
	if v, ok := obj["tcp-mss"]; ok {
		if v != "" {
			m.TCPMss = types.StringValue(v)
		} else {
			m.TCPMss = types.StringNull()
		}
	}
	if v, ok := obj["time"]; ok {
		_ = v
		if v != "" {
			m.Time = newCSVSetValue(v)
		} else {
			m.Time = newCSVSetNull()
		}
	} else {
		m.Time = newCSVSetNull()
	}
	if v, ok := obj["tls-host"]; ok {
		if v != "" {
			m.TLSHost = types.StringValue(v)
		} else {
			m.TLSHost = types.StringNull()
		}
	}
}
