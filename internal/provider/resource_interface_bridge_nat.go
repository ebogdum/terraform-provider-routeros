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
	_ resource.Resource                = &InterfaceBridgeNATResource{}
	_ resource.ResourceWithImportState = &InterfaceBridgeNATResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceBridgeNATResource struct {
	reg *client.Registry
}

type InterfaceBridgeNATModel struct {
	ID                   types.String `tfsdk:"id"`
	VlanPriority         types.String `tfsdk:"vlan_priority"`
	VlanId               types.String `tfsdk:"vlan_id"`
	VlanEncap            types.String `tfsdk:"vlan_encap"`
	ToSrcMacAddress      macValue     `tfsdk:"to_src_mac_address"`
	ToDstMacAddress      macValue     `tfsdk:"to_dst_mac_address"`
	ToArpReplyMacAddress macValue     `tfsdk:"to_arp_reply_mac_address"`
	StpType              types.String `tfsdk:"stp_type"`
	StpSenderPriority    types.String `tfsdk:"stp_sender_priority"`
	StpSenderAddress     types.String `tfsdk:"stp_sender_address"`
	StpRootPriority      types.String `tfsdk:"stp_root_priority"`
	StpRootCost          types.String `tfsdk:"stp_root_cost"`
	StpRootAddress       types.String `tfsdk:"stp_root_address"`
	StpPort              types.String `tfsdk:"stp_port"`
	StpMsgAge            types.String `tfsdk:"stp_msg_age"`
	StpMaxAge            types.String `tfsdk:"stp_max_age"`
	StpHelloTime         types.String `tfsdk:"stp_hello_time"`
	StpForwardDelay      types.String `tfsdk:"stp_forward_delay"`
	StpFlags             types.String `tfsdk:"stp_flags"`
	SrcAddress6          types.String `tfsdk:"src_address6"`
	Passthrough          types.String `tfsdk:"passthrough"`
	PacketType           types.String `tfsdk:"packet_type"`
	OutBridgeList        types.String `tfsdk:"out_bridge_list"`
	OutBridge            types.String `tfsdk:"out_bridge"`
	NewPriority          types.String `tfsdk:"new_priority"`
	NewPacketMark        types.String `tfsdk:"new_packet_mark"`
	MacProtocol          types.String `tfsdk:"mac_protocol"`
	IpProtocol           types.String `tfsdk:"ip_protocol"`
	InBridgeList         types.String `tfsdk:"in_bridge_list"`
	InBridge             types.String `tfsdk:"in_bridge"`
	DstMacAddress        types.String `tfsdk:"dst_mac_address"`
	DstAddress6          types.String `tfsdk:"dst_address6"`
	ArpSrcMacAddress     types.String `tfsdk:"arp_src_mac_address"`
	ArpSrcAddress        types.String `tfsdk:"arp_src_address"`
	ArpPacketType        types.String `tfsdk:"arp_packet_type"`
	ArpOpcode            types.String `tfsdk:"arp_opcode"`
	ArpHardwareType      types.String `tfsdk:"arp_hardware_type"`
	ArpGratuitous        types.String `tfsdk:"arp_gratuitous"`
	ArpDstMacAddress     types.String `tfsdk:"arp_dst_mac_address"`
	ArpDstAddress        types.String `tfsdk:"arp_dst_address"`
	Action               types.String `tfsdk:"action"`
	Chain                types.String `tfsdk:"chain"`
	Comment              types.String `tfsdk:"comment"`
	Disabled             types.Bool   `tfsdk:"disabled"`
	DstAddress           types.String `tfsdk:"dst_address"`
	DstPort              types.String `tfsdk:"dst_port"`
	InInterface          types.String `tfsdk:"in_interface"`
	InInterfaceList      types.String `tfsdk:"in_interface_list"`
	IngressPriority      types.String `tfsdk:"ingress_priority"`
	JumpTarget           types.String `tfsdk:"jump_target"`
	Limit                types.String `tfsdk:"limit"`
	Log                  types.String `tfsdk:"log"`
	LogPrefix            types.String `tfsdk:"log_prefix"`
	OutInterface         types.String `tfsdk:"out_interface"`
	OutInterfaceList     types.String `tfsdk:"out_interface_list"`
	PacketMark           types.String `tfsdk:"packet_mark"`
	SrcAddress           types.String `tfsdk:"src_address"`
	SrcMACAddress        types.String `tfsdk:"src_mac_address"`
	SrcPort              types.String `tfsdk:"src_port"`
	TLSHost              types.String `tfsdk:"tls_host"`
	Router               types.String `tfsdk:"router"`
}

func NewInterfaceBridgeNATResource() resource.Resource { return &InterfaceBridgeNATResource{} }

func (r *InterfaceBridgeNATResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_bridge_nat"
}

func (r *InterfaceBridgeNATResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceBridgeNATResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/bridge/nat`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"vlan_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `vlan-priority`.",
			},
			"vlan_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `vlan-id`.",
			},
			"vlan_encap": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `vlan-encap`.",
			},
			"to_src_mac_address": schema.StringAttribute{
				CustomType:  macType{},
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `to-src-mac-address`.",
				Validators:  []validator.String{schemautil.IsMAC()},
			},
			"to_dst_mac_address": schema.StringAttribute{
				CustomType:  macType{},
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `to-dst-mac-address`.",
				Validators:  []validator.String{schemautil.IsMAC()},
			},
			"to_arp_reply_mac_address": schema.StringAttribute{
				CustomType:  macType{},
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `to-arp-reply-mac-address`.",
				Validators:  []validator.String{schemautil.IsMAC()},
			},
			"stp_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `stp-type`.",
			},
			"stp_sender_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `stp-sender-priority`.",
			},
			"stp_sender_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `stp-sender-address`.",
			},
			"stp_root_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `stp-root-priority`.",
			},
			"stp_root_cost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `stp-root-cost`.",
			},
			"stp_root_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `stp-root-address`.",
			},
			"stp_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `stp-port`.",
			},
			"stp_msg_age": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `stp-msg-age`.",
			},
			"stp_max_age": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `stp-max-age`.",
			},
			"stp_hello_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `stp-hello-time`.",
			},
			"stp_forward_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `stp-forward-delay`.",
			},
			"stp_flags": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `stp-flags`.",
			},
			"src_address6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `src-address6`.",
			},
			"passthrough": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `passthrough`.",
			},
			"packet_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `packet-type`.",
			},
			"out_bridge_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `out-bridge-list`.",
			},
			"out_bridge": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `out-bridge`.",
			},
			"new_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `new-priority`.",
			},
			"new_packet_mark": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `new-packet-mark`.",
			},
			"mac_protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mac-protocol`.",
			},
			"ip_protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ip-protocol`.",
			},
			"in_bridge_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `in-bridge-list`.",
			},
			"in_bridge": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `in-bridge`.",
			},
			"dst_mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dst-mac-address`.",
			},
			"dst_address6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dst-address6`.",
			},
			"arp_src_mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `arp-src-mac-address`.",
			},
			"arp_src_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `arp-src-address`.",
			},
			"arp_packet_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `arp-packet-type`.",
			},
			"arp_opcode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `arp-opcode`.",
			},
			"arp_hardware_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `arp-hardware-type`.",
			},
			"arp_gratuitous": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `arp-gratuitous`.",
			},
			"arp_dst_mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `arp-dst-mac-address`.",
			},
			"arp_dst_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `arp-dst-address`.",
			},
			"action": schema.StringAttribute{
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
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"dst_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dst_port": schema.StringAttribute{
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
			"src_address": schema.StringAttribute{
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
			"tls_host": schema.StringAttribute{
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

func (r *InterfaceBridgeNATResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceBridgeNATModel
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
	if !(plan.Chain.IsNull() || plan.Chain.IsUnknown()) {
		body["chain"] = plan.Chain.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.DstAddress.IsNull() || plan.DstAddress.IsUnknown()) {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !(plan.DstPort.IsNull() || plan.DstPort.IsUnknown()) {
		body["dst-port"] = plan.DstPort.ValueString()
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
	if !(plan.OutInterface.IsNull() || plan.OutInterface.IsUnknown()) {
		body["out-interface"] = plan.OutInterface.ValueString()
	}
	if !(plan.OutInterfaceList.IsNull() || plan.OutInterfaceList.IsUnknown()) {
		body["out-interface-list"] = plan.OutInterfaceList.ValueString()
	}
	if !(plan.PacketMark.IsNull() || plan.PacketMark.IsUnknown()) {
		body["packet-mark"] = plan.PacketMark.ValueString()
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !(plan.SrcMACAddress.IsNull() || plan.SrcMACAddress.IsUnknown()) {
		body["src-mac-address"] = plan.SrcMACAddress.ValueString()
	}
	if !(plan.SrcPort.IsNull() || plan.SrcPort.IsUnknown()) {
		body["src-port"] = plan.SrcPort.ValueString()
	}
	if !(plan.TLSHost.IsNull() || plan.TLSHost.IsUnknown()) {
		body["tls-host"] = plan.TLSHost.ValueString()
	}
	if !(plan.ArpDstAddress.IsNull() || plan.ArpDstAddress.IsUnknown()) {
		body["arp-dst-address"] = plan.ArpDstAddress.ValueString()
	}
	if !(plan.ArpDstMacAddress.IsNull() || plan.ArpDstMacAddress.IsUnknown()) {
		body["arp-dst-mac-address"] = plan.ArpDstMacAddress.ValueString()
	}
	if !(plan.ArpGratuitous.IsNull() || plan.ArpGratuitous.IsUnknown()) {
		body["arp-gratuitous"] = plan.ArpGratuitous.ValueString()
	}
	if !(plan.ArpHardwareType.IsNull() || plan.ArpHardwareType.IsUnknown()) {
		body["arp-hardware-type"] = plan.ArpHardwareType.ValueString()
	}
	if !(plan.ArpOpcode.IsNull() || plan.ArpOpcode.IsUnknown()) {
		body["arp-opcode"] = plan.ArpOpcode.ValueString()
	}
	if !(plan.ArpPacketType.IsNull() || plan.ArpPacketType.IsUnknown()) {
		body["arp-packet-type"] = plan.ArpPacketType.ValueString()
	}
	if !(plan.ArpSrcAddress.IsNull() || plan.ArpSrcAddress.IsUnknown()) {
		body["arp-src-address"] = plan.ArpSrcAddress.ValueString()
	}
	if !(plan.ArpSrcMacAddress.IsNull() || plan.ArpSrcMacAddress.IsUnknown()) {
		body["arp-src-mac-address"] = plan.ArpSrcMacAddress.ValueString()
	}
	if !(plan.DstAddress6.IsNull() || plan.DstAddress6.IsUnknown()) {
		body["dst-address6"] = plan.DstAddress6.ValueString()
	}
	if !(plan.DstMacAddress.IsNull() || plan.DstMacAddress.IsUnknown()) {
		body["dst-mac-address"] = plan.DstMacAddress.ValueString()
	}
	if !(plan.InBridge.IsNull() || plan.InBridge.IsUnknown()) {
		body["in-bridge"] = plan.InBridge.ValueString()
	}
	if !(plan.InBridgeList.IsNull() || plan.InBridgeList.IsUnknown()) {
		body["in-bridge-list"] = plan.InBridgeList.ValueString()
	}
	if !(plan.IpProtocol.IsNull() || plan.IpProtocol.IsUnknown()) {
		body["ip-protocol"] = plan.IpProtocol.ValueString()
	}
	if !(plan.MacProtocol.IsNull() || plan.MacProtocol.IsUnknown()) {
		body["mac-protocol"] = plan.MacProtocol.ValueString()
	}
	if !(plan.NewPacketMark.IsNull() || plan.NewPacketMark.IsUnknown()) {
		body["new-packet-mark"] = plan.NewPacketMark.ValueString()
	}
	if !(plan.NewPriority.IsNull() || plan.NewPriority.IsUnknown()) {
		body["new-priority"] = plan.NewPriority.ValueString()
	}
	if !(plan.OutBridge.IsNull() || plan.OutBridge.IsUnknown()) {
		body["out-bridge"] = plan.OutBridge.ValueString()
	}
	if !(plan.OutBridgeList.IsNull() || plan.OutBridgeList.IsUnknown()) {
		body["out-bridge-list"] = plan.OutBridgeList.ValueString()
	}
	if !(plan.PacketType.IsNull() || plan.PacketType.IsUnknown()) {
		body["packet-type"] = plan.PacketType.ValueString()
	}
	if !(plan.Passthrough.IsNull() || plan.Passthrough.IsUnknown()) {
		body["passthrough"] = plan.Passthrough.ValueString()
	}
	if !(plan.SrcAddress6.IsNull() || plan.SrcAddress6.IsUnknown()) {
		body["src-address6"] = plan.SrcAddress6.ValueString()
	}
	if !(plan.StpFlags.IsNull() || plan.StpFlags.IsUnknown()) {
		body["stp-flags"] = plan.StpFlags.ValueString()
	}
	if !(plan.StpForwardDelay.IsNull() || plan.StpForwardDelay.IsUnknown()) {
		body["stp-forward-delay"] = plan.StpForwardDelay.ValueString()
	}
	if !(plan.StpHelloTime.IsNull() || plan.StpHelloTime.IsUnknown()) {
		body["stp-hello-time"] = plan.StpHelloTime.ValueString()
	}
	if !(plan.StpMaxAge.IsNull() || plan.StpMaxAge.IsUnknown()) {
		body["stp-max-age"] = plan.StpMaxAge.ValueString()
	}
	if !(plan.StpMsgAge.IsNull() || plan.StpMsgAge.IsUnknown()) {
		body["stp-msg-age"] = plan.StpMsgAge.ValueString()
	}
	if !(plan.StpPort.IsNull() || plan.StpPort.IsUnknown()) {
		body["stp-port"] = plan.StpPort.ValueString()
	}
	if !(plan.StpRootAddress.IsNull() || plan.StpRootAddress.IsUnknown()) {
		body["stp-root-address"] = plan.StpRootAddress.ValueString()
	}
	if !(plan.StpRootCost.IsNull() || plan.StpRootCost.IsUnknown()) {
		body["stp-root-cost"] = plan.StpRootCost.ValueString()
	}
	if !(plan.StpRootPriority.IsNull() || plan.StpRootPriority.IsUnknown()) {
		body["stp-root-priority"] = plan.StpRootPriority.ValueString()
	}
	if !(plan.StpSenderAddress.IsNull() || plan.StpSenderAddress.IsUnknown()) {
		body["stp-sender-address"] = plan.StpSenderAddress.ValueString()
	}
	if !(plan.StpSenderPriority.IsNull() || plan.StpSenderPriority.IsUnknown()) {
		body["stp-sender-priority"] = plan.StpSenderPriority.ValueString()
	}
	if !(plan.StpType.IsNull() || plan.StpType.IsUnknown()) {
		body["stp-type"] = plan.StpType.ValueString()
	}
	if !(plan.ToArpReplyMacAddress.IsNull() || plan.ToArpReplyMacAddress.IsUnknown()) {
		body["to-arp-reply-mac-address"] = plan.ToArpReplyMacAddress.ValueString()
	}
	if !(plan.ToDstMacAddress.IsNull() || plan.ToDstMacAddress.IsUnknown()) {
		body["to-dst-mac-address"] = plan.ToDstMacAddress.ValueString()
	}
	if !(plan.ToSrcMacAddress.IsNull() || plan.ToSrcMacAddress.IsUnknown()) {
		body["to-src-mac-address"] = plan.ToSrcMacAddress.ValueString()
	}
	if !(plan.VlanEncap.IsNull() || plan.VlanEncap.IsUnknown()) {
		body["vlan-encap"] = plan.VlanEncap.ValueString()
	}
	if !(plan.VlanId.IsNull() || plan.VlanId.IsUnknown()) {
		body["vlan-id"] = plan.VlanId.ValueString()
	}
	if !(plan.VlanPriority.IsNull() || plan.VlanPriority.IsUnknown()) {
		body["vlan-priority"] = plan.VlanPriority.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/bridge/nat", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/bridge/nat failed", err.Error())
		return
	}
	interfaceBridgeNATApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceBridgeNATResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceBridgeNATModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/bridge/nat", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/bridge/nat failed", err.Error())
		return
	}
	interfaceBridgeNATApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceBridgeNATResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceBridgeNATModel
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
	if !plan.Chain.Equal(state.Chain) && !plan.Chain.IsUnknown() {
		body["chain"] = plan.Chain.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DstAddress.Equal(state.DstAddress) && !plan.DstAddress.IsUnknown() {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !plan.DstPort.Equal(state.DstPort) && !plan.DstPort.IsUnknown() {
		body["dst-port"] = plan.DstPort.ValueString()
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
	if !plan.OutInterface.Equal(state.OutInterface) && !plan.OutInterface.IsUnknown() {
		body["out-interface"] = plan.OutInterface.ValueString()
	}
	if !plan.OutInterfaceList.Equal(state.OutInterfaceList) && !plan.OutInterfaceList.IsUnknown() {
		body["out-interface-list"] = plan.OutInterfaceList.ValueString()
	}
	if !plan.PacketMark.Equal(state.PacketMark) && !plan.PacketMark.IsUnknown() {
		body["packet-mark"] = plan.PacketMark.ValueString()
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) && !plan.SrcAddress.IsUnknown() {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !plan.SrcMACAddress.Equal(state.SrcMACAddress) && !plan.SrcMACAddress.IsUnknown() {
		body["src-mac-address"] = plan.SrcMACAddress.ValueString()
	}
	if !plan.SrcPort.Equal(state.SrcPort) && !plan.SrcPort.IsUnknown() {
		body["src-port"] = plan.SrcPort.ValueString()
	}
	if !plan.TLSHost.Equal(state.TLSHost) && !plan.TLSHost.IsUnknown() {
		body["tls-host"] = plan.TLSHost.ValueString()
	}
	if !plan.ArpDstAddress.Equal(state.ArpDstAddress) && !plan.ArpDstAddress.IsUnknown() {
		body["arp-dst-address"] = plan.ArpDstAddress.ValueString()
	}
	if !plan.ArpDstMacAddress.Equal(state.ArpDstMacAddress) && !plan.ArpDstMacAddress.IsUnknown() {
		body["arp-dst-mac-address"] = plan.ArpDstMacAddress.ValueString()
	}
	if !plan.ArpGratuitous.Equal(state.ArpGratuitous) && !plan.ArpGratuitous.IsUnknown() {
		body["arp-gratuitous"] = plan.ArpGratuitous.ValueString()
	}
	if !plan.ArpHardwareType.Equal(state.ArpHardwareType) && !plan.ArpHardwareType.IsUnknown() {
		body["arp-hardware-type"] = plan.ArpHardwareType.ValueString()
	}
	if !plan.ArpOpcode.Equal(state.ArpOpcode) && !plan.ArpOpcode.IsUnknown() {
		body["arp-opcode"] = plan.ArpOpcode.ValueString()
	}
	if !plan.ArpPacketType.Equal(state.ArpPacketType) && !plan.ArpPacketType.IsUnknown() {
		body["arp-packet-type"] = plan.ArpPacketType.ValueString()
	}
	if !plan.ArpSrcAddress.Equal(state.ArpSrcAddress) && !plan.ArpSrcAddress.IsUnknown() {
		body["arp-src-address"] = plan.ArpSrcAddress.ValueString()
	}
	if !plan.ArpSrcMacAddress.Equal(state.ArpSrcMacAddress) && !plan.ArpSrcMacAddress.IsUnknown() {
		body["arp-src-mac-address"] = plan.ArpSrcMacAddress.ValueString()
	}
	if !plan.DstAddress6.Equal(state.DstAddress6) && !plan.DstAddress6.IsUnknown() {
		body["dst-address6"] = plan.DstAddress6.ValueString()
	}
	if !plan.DstMacAddress.Equal(state.DstMacAddress) && !plan.DstMacAddress.IsUnknown() {
		body["dst-mac-address"] = plan.DstMacAddress.ValueString()
	}
	if !plan.InBridge.Equal(state.InBridge) && !plan.InBridge.IsUnknown() {
		body["in-bridge"] = plan.InBridge.ValueString()
	}
	if !plan.InBridgeList.Equal(state.InBridgeList) && !plan.InBridgeList.IsUnknown() {
		body["in-bridge-list"] = plan.InBridgeList.ValueString()
	}
	if !plan.IpProtocol.Equal(state.IpProtocol) && !plan.IpProtocol.IsUnknown() {
		body["ip-protocol"] = plan.IpProtocol.ValueString()
	}
	if !plan.MacProtocol.Equal(state.MacProtocol) && !plan.MacProtocol.IsUnknown() {
		body["mac-protocol"] = plan.MacProtocol.ValueString()
	}
	if !plan.NewPacketMark.Equal(state.NewPacketMark) && !plan.NewPacketMark.IsUnknown() {
		body["new-packet-mark"] = plan.NewPacketMark.ValueString()
	}
	if !plan.NewPriority.Equal(state.NewPriority) && !plan.NewPriority.IsUnknown() {
		body["new-priority"] = plan.NewPriority.ValueString()
	}
	if !plan.OutBridge.Equal(state.OutBridge) && !plan.OutBridge.IsUnknown() {
		body["out-bridge"] = plan.OutBridge.ValueString()
	}
	if !plan.OutBridgeList.Equal(state.OutBridgeList) && !plan.OutBridgeList.IsUnknown() {
		body["out-bridge-list"] = plan.OutBridgeList.ValueString()
	}
	if !plan.PacketType.Equal(state.PacketType) && !plan.PacketType.IsUnknown() {
		body["packet-type"] = plan.PacketType.ValueString()
	}
	if !plan.Passthrough.Equal(state.Passthrough) && !plan.Passthrough.IsUnknown() {
		body["passthrough"] = plan.Passthrough.ValueString()
	}
	if !plan.SrcAddress6.Equal(state.SrcAddress6) && !plan.SrcAddress6.IsUnknown() {
		body["src-address6"] = plan.SrcAddress6.ValueString()
	}
	if !plan.StpFlags.Equal(state.StpFlags) && !plan.StpFlags.IsUnknown() {
		body["stp-flags"] = plan.StpFlags.ValueString()
	}
	if !plan.StpForwardDelay.Equal(state.StpForwardDelay) && !plan.StpForwardDelay.IsUnknown() {
		body["stp-forward-delay"] = plan.StpForwardDelay.ValueString()
	}
	if !plan.StpHelloTime.Equal(state.StpHelloTime) && !plan.StpHelloTime.IsUnknown() {
		body["stp-hello-time"] = plan.StpHelloTime.ValueString()
	}
	if !plan.StpMaxAge.Equal(state.StpMaxAge) && !plan.StpMaxAge.IsUnknown() {
		body["stp-max-age"] = plan.StpMaxAge.ValueString()
	}
	if !plan.StpMsgAge.Equal(state.StpMsgAge) && !plan.StpMsgAge.IsUnknown() {
		body["stp-msg-age"] = plan.StpMsgAge.ValueString()
	}
	if !plan.StpPort.Equal(state.StpPort) && !plan.StpPort.IsUnknown() {
		body["stp-port"] = plan.StpPort.ValueString()
	}
	if !plan.StpRootAddress.Equal(state.StpRootAddress) && !plan.StpRootAddress.IsUnknown() {
		body["stp-root-address"] = plan.StpRootAddress.ValueString()
	}
	if !plan.StpRootCost.Equal(state.StpRootCost) && !plan.StpRootCost.IsUnknown() {
		body["stp-root-cost"] = plan.StpRootCost.ValueString()
	}
	if !plan.StpRootPriority.Equal(state.StpRootPriority) && !plan.StpRootPriority.IsUnknown() {
		body["stp-root-priority"] = plan.StpRootPriority.ValueString()
	}
	if !plan.StpSenderAddress.Equal(state.StpSenderAddress) && !plan.StpSenderAddress.IsUnknown() {
		body["stp-sender-address"] = plan.StpSenderAddress.ValueString()
	}
	if !plan.StpSenderPriority.Equal(state.StpSenderPriority) && !plan.StpSenderPriority.IsUnknown() {
		body["stp-sender-priority"] = plan.StpSenderPriority.ValueString()
	}
	if !plan.StpType.Equal(state.StpType) && !plan.StpType.IsUnknown() {
		body["stp-type"] = plan.StpType.ValueString()
	}
	if !plan.ToArpReplyMacAddress.Equal(state.ToArpReplyMacAddress) && !plan.ToArpReplyMacAddress.IsUnknown() {
		body["to-arp-reply-mac-address"] = plan.ToArpReplyMacAddress.ValueString()
	}
	if !plan.ToDstMacAddress.Equal(state.ToDstMacAddress) && !plan.ToDstMacAddress.IsUnknown() {
		body["to-dst-mac-address"] = plan.ToDstMacAddress.ValueString()
	}
	if !plan.ToSrcMacAddress.Equal(state.ToSrcMacAddress) && !plan.ToSrcMacAddress.IsUnknown() {
		body["to-src-mac-address"] = plan.ToSrcMacAddress.ValueString()
	}
	if !plan.VlanEncap.Equal(state.VlanEncap) && !plan.VlanEncap.IsUnknown() {
		body["vlan-encap"] = plan.VlanEncap.ValueString()
	}
	if !plan.VlanId.Equal(state.VlanId) && !plan.VlanId.IsUnknown() {
		body["vlan-id"] = plan.VlanId.ValueString()
	}
	if !plan.VlanPriority.Equal(state.VlanPriority) && !plan.VlanPriority.IsUnknown() {
		body["vlan-priority"] = plan.VlanPriority.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/bridge/nat", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/bridge/nat failed", err.Error())
			return
		}
		interfaceBridgeNATApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceBridgeNATResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceBridgeNATModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/bridge/nat", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/bridge/nat failed", err.Error())
	}
}

func (r *InterfaceBridgeNATResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceBridgeNATLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/bridge/nat matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceBridgeNATLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceBridgeNATLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/bridge/nat", id)
}

func interfaceBridgeNATApply(ctx context.Context, obj client.Object, m *InterfaceBridgeNATModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["vlan-priority"]; ok && v != "" {
		m.VlanPriority = types.StringValue(v)
	} else {
		m.VlanPriority = types.StringNull()
	}
	if v, ok := obj["vlan-id"]; ok && v != "" {
		m.VlanId = types.StringValue(v)
	} else {
		m.VlanId = types.StringNull()
	}
	if v, ok := obj["vlan-encap"]; ok && v != "" {
		m.VlanEncap = types.StringValue(v)
	} else {
		m.VlanEncap = types.StringNull()
	}
	if v, ok := obj["to-src-mac-address"]; ok && v != "" {
		m.ToSrcMacAddress = newMACValue(v)
	} else {
		m.ToSrcMacAddress = newMACNull()
	}
	if v, ok := obj["to-dst-mac-address"]; ok && v != "" {
		m.ToDstMacAddress = newMACValue(v)
	} else {
		m.ToDstMacAddress = newMACNull()
	}
	if v, ok := obj["to-arp-reply-mac-address"]; ok && v != "" {
		m.ToArpReplyMacAddress = newMACValue(v)
	} else {
		m.ToArpReplyMacAddress = newMACNull()
	}
	if v, ok := obj["stp-type"]; ok && v != "" {
		m.StpType = types.StringValue(v)
	} else {
		m.StpType = types.StringNull()
	}
	if v, ok := obj["stp-sender-priority"]; ok && v != "" {
		m.StpSenderPriority = types.StringValue(v)
	} else {
		m.StpSenderPriority = types.StringNull()
	}
	if v, ok := obj["stp-sender-address"]; ok && v != "" {
		m.StpSenderAddress = types.StringValue(v)
	} else {
		m.StpSenderAddress = types.StringNull()
	}
	if v, ok := obj["stp-root-priority"]; ok && v != "" {
		m.StpRootPriority = types.StringValue(v)
	} else {
		m.StpRootPriority = types.StringNull()
	}
	if v, ok := obj["stp-root-cost"]; ok && v != "" {
		m.StpRootCost = types.StringValue(v)
	} else {
		m.StpRootCost = types.StringNull()
	}
	if v, ok := obj["stp-root-address"]; ok && v != "" {
		m.StpRootAddress = types.StringValue(v)
	} else {
		m.StpRootAddress = types.StringNull()
	}
	if v, ok := obj["stp-port"]; ok && v != "" {
		m.StpPort = types.StringValue(v)
	} else {
		m.StpPort = types.StringNull()
	}
	if v, ok := obj["stp-msg-age"]; ok && v != "" {
		m.StpMsgAge = types.StringValue(v)
	} else {
		m.StpMsgAge = types.StringNull()
	}
	if v, ok := obj["stp-max-age"]; ok && v != "" {
		m.StpMaxAge = types.StringValue(v)
	} else {
		m.StpMaxAge = types.StringNull()
	}
	if v, ok := obj["stp-hello-time"]; ok && v != "" {
		m.StpHelloTime = types.StringValue(v)
	} else {
		m.StpHelloTime = types.StringNull()
	}
	if v, ok := obj["stp-forward-delay"]; ok && v != "" {
		m.StpForwardDelay = types.StringValue(v)
	} else {
		m.StpForwardDelay = types.StringNull()
	}
	if v, ok := obj["stp-flags"]; ok && v != "" {
		m.StpFlags = types.StringValue(v)
	} else {
		m.StpFlags = types.StringNull()
	}
	if v, ok := obj["src-address6"]; ok && v != "" {
		m.SrcAddress6 = types.StringValue(v)
	} else {
		m.SrcAddress6 = types.StringNull()
	}
	if v, ok := obj["passthrough"]; ok && v != "" {
		m.Passthrough = types.StringValue(v)
	} else {
		m.Passthrough = types.StringNull()
	}
	if v, ok := obj["packet-type"]; ok && v != "" {
		m.PacketType = types.StringValue(v)
	} else {
		m.PacketType = types.StringNull()
	}
	if v, ok := obj["out-bridge-list"]; ok && v != "" {
		m.OutBridgeList = types.StringValue(v)
	} else {
		m.OutBridgeList = types.StringNull()
	}
	if v, ok := obj["out-bridge"]; ok && v != "" {
		m.OutBridge = types.StringValue(v)
	} else {
		m.OutBridge = types.StringNull()
	}
	if v, ok := obj["new-priority"]; ok && v != "" {
		m.NewPriority = types.StringValue(v)
	} else {
		m.NewPriority = types.StringNull()
	}
	if v, ok := obj["new-packet-mark"]; ok && v != "" {
		m.NewPacketMark = types.StringValue(v)
	} else {
		m.NewPacketMark = types.StringNull()
	}
	if v, ok := obj["mac-protocol"]; ok && v != "" {
		m.MacProtocol = types.StringValue(v)
	} else {
		m.MacProtocol = types.StringNull()
	}
	if v, ok := obj["ip-protocol"]; ok && v != "" {
		m.IpProtocol = types.StringValue(v)
	} else {
		m.IpProtocol = types.StringNull()
	}
	if v, ok := obj["in-bridge-list"]; ok && v != "" {
		m.InBridgeList = types.StringValue(v)
	} else {
		m.InBridgeList = types.StringNull()
	}
	if v, ok := obj["in-bridge"]; ok && v != "" {
		m.InBridge = types.StringValue(v)
	} else {
		m.InBridge = types.StringNull()
	}
	if v, ok := obj["dst-mac-address"]; ok && v != "" {
		m.DstMacAddress = types.StringValue(v)
	} else {
		m.DstMacAddress = types.StringNull()
	}
	if v, ok := obj["dst-address6"]; ok && v != "" {
		m.DstAddress6 = types.StringValue(v)
	} else {
		m.DstAddress6 = types.StringNull()
	}
	if v, ok := obj["arp-src-mac-address"]; ok && v != "" {
		m.ArpSrcMacAddress = types.StringValue(v)
	} else {
		m.ArpSrcMacAddress = types.StringNull()
	}
	if v, ok := obj["arp-src-address"]; ok && v != "" {
		m.ArpSrcAddress = types.StringValue(v)
	} else {
		m.ArpSrcAddress = types.StringNull()
	}
	if v, ok := obj["arp-packet-type"]; ok && v != "" {
		m.ArpPacketType = types.StringValue(v)
	} else {
		m.ArpPacketType = types.StringNull()
	}
	if v, ok := obj["arp-opcode"]; ok && v != "" {
		m.ArpOpcode = types.StringValue(v)
	} else {
		m.ArpOpcode = types.StringNull()
	}
	if v, ok := obj["arp-hardware-type"]; ok && v != "" {
		m.ArpHardwareType = types.StringValue(v)
	} else {
		m.ArpHardwareType = types.StringNull()
	}
	if v, ok := obj["arp-gratuitous"]; ok && v != "" {
		m.ArpGratuitous = types.StringValue(v)
	} else {
		m.ArpGratuitous = types.StringNull()
	}
	if v, ok := obj["arp-dst-mac-address"]; ok && v != "" {
		m.ArpDstMacAddress = types.StringValue(v)
	} else {
		m.ArpDstMacAddress = types.StringNull()
	}
	if v, ok := obj["arp-dst-address"]; ok && v != "" {
		m.ArpDstAddress = types.StringValue(v)
	} else {
		m.ArpDstAddress = types.StringNull()
	}
	if v, ok := obj["action"]; ok {
		if v != "" {
			m.Action = types.StringValue(v)
		} else {
			m.Action = types.StringNull()
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
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Disabled = types.BoolValue(true)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["dst-address"]; ok {
		if v != "" {
			m.DstAddress = types.StringValue(v)
		} else {
			m.DstAddress = types.StringNull()
		}
	}
	if v, ok := obj["dst-port"]; ok {
		if v != "" {
			m.DstPort = types.StringValue(v)
		} else {
			m.DstPort = types.StringNull()
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
	if v, ok := obj["src-address"]; ok {
		if v != "" {
			m.SrcAddress = types.StringValue(v)
		} else {
			m.SrcAddress = types.StringNull()
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
	if v, ok := obj["tls-host"]; ok {
		if v != "" {
			m.TLSHost = types.StringValue(v)
		} else {
			m.TLSHost = types.StringNull()
		}
	}
}
