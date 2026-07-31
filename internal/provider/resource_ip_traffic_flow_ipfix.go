package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &IPTrafficFlowIPFIXResource{}
	_ resource.ResourceWithImportState = &IPTrafficFlowIPFIXResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPTrafficFlowIPFIXResource struct {
	reg *client.Registry
}

type IPTrafficFlowIPFIXModel struct {
	ID             types.String `tfsdk:"id"`
	Bytes          types.Bool   `tfsdk:"bytes"`
	DstAddress     types.Bool   `tfsdk:"dst_address"`
	DstAddressMask types.Bool   `tfsdk:"dst_address_mask"`
	DstMACAddress  types.Bool   `tfsdk:"dst_mac_address"`
	DstPort        types.Bool   `tfsdk:"dst_port"`
	FirstForwarded types.Bool   `tfsdk:"first_forwarded"`
	Gateway        types.Bool   `tfsdk:"gateway"`
	IcmpCode       types.Bool   `tfsdk:"icmp_code"`
	IcmpType       types.Bool   `tfsdk:"icmp_type"`
	IgmpType       types.Bool   `tfsdk:"igmp_type"`
	InInterface    types.Bool   `tfsdk:"in_interface"`
	IPHeaderLength types.Bool   `tfsdk:"ip_header_length"`
	IPTotalLength  types.Bool   `tfsdk:"ip_total_length"`
	IPv6FlowLabel  types.Bool   `tfsdk:"ipv6_flow_label"`
	IsMulticast    types.Bool   `tfsdk:"is_multicast"`
	LastForwarded  types.Bool   `tfsdk:"last_forwarded"`
	NatDstAddress  types.Bool   `tfsdk:"nat_dst_address"`
	NatDstPort     types.Bool   `tfsdk:"nat_dst_port"`
	NatEvents      types.Bool   `tfsdk:"nat_events"`
	NatSrcAddress  types.Bool   `tfsdk:"nat_src_address"`
	NatSrcPort     types.Bool   `tfsdk:"nat_src_port"`
	OutInterface   types.Bool   `tfsdk:"out_interface"`
	Packets        types.Bool   `tfsdk:"packets"`
	Protocol       types.Bool   `tfsdk:"protocol"`
	SrcAddress     types.Bool   `tfsdk:"src_address"`
	SrcAddressMask types.Bool   `tfsdk:"src_address_mask"`
	SrcMACAddress  types.Bool   `tfsdk:"src_mac_address"`
	SrcPort        types.Bool   `tfsdk:"src_port"`
	SysInitTime    types.Bool   `tfsdk:"sys_init_time"`
	TCPAckNum      types.Bool   `tfsdk:"tcp_ack_num"`
	TCPFlags       types.Bool   `tfsdk:"tcp_flags"`
	TCPSeqNum      types.Bool   `tfsdk:"tcp_seq_num"`
	TCPWindowSize  types.Bool   `tfsdk:"tcp_window_size"`
	Tos            types.Bool   `tfsdk:"tos"`
	Ttl            types.Bool   `tfsdk:"ttl"`
	UDPLength      types.Bool   `tfsdk:"udp_length"`
	Router         types.String `tfsdk:"router"`
}

func NewIPTrafficFlowIPFIXResource() resource.Resource { return &IPTrafficFlowIPFIXResource{} }

func (r *IPTrafficFlowIPFIXResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_traffic_flow_ipfix"
}

func (r *IPTrafficFlowIPFIXResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPTrafficFlowIPFIXResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/traffic-flow/ipfix`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"bytes": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bytes`.",
			},
			"dst_address": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dst-address`.",
			},
			"dst_address_mask": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dst-address-mask`.",
			},
			"dst_mac_address": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dst-mac-address`.",
			},
			"dst_port": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dst-port`.",
			},
			"first_forwarded": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `first-forwarded`.",
			},
			"gateway": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `gateway`.",
			},
			"icmp_code": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `icmp-code`.",
			},
			"icmp_type": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `icmp-type`.",
			},
			"igmp_type": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `igmp-type`.",
			},
			"in_interface": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `in-interface`.",
			},
			"ip_header_length": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ip-header-length`.",
			},
			"ip_total_length": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ip-total-length`.",
			},
			"ipv6_flow_label": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ipv6-flow-label`.",
			},
			"is_multicast": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `is-multicast`.",
			},
			"last_forwarded": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `last-forwarded`.",
			},
			"nat_dst_address": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `nat-dst-address`.",
			},
			"nat_dst_port": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `nat-dst-port`.",
			},
			"nat_events": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `nat-events`.",
			},
			"nat_src_address": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `nat-src-address`.",
			},
			"nat_src_port": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `nat-src-port`.",
			},
			"out_interface": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `out-interface`.",
			},
			"packets": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `packets`.",
			},
			"protocol": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `protocol`.",
			},
			"src_address": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `src-address`.",
			},
			"src_address_mask": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `src-address-mask`.",
			},
			"src_mac_address": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `src-mac-address`.",
			},
			"src_port": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `src-port`.",
			},
			"sys_init_time": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `sys-init-time`.",
			},
			"tcp_ack_num": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tcp-ack-num`.",
			},
			"tcp_flags": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tcp-flags`.",
			},
			"tcp_seq_num": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tcp-seq-num`.",
			},
			"tcp_window_size": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tcp-window-size`.",
			},
			"tos": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tos`.",
			},
			"ttl": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ttl`.",
			},
			"udp_length": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `udp-length`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPTrafficFlowIPFIXResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPTrafficFlowIPFIXModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPTrafficFlowIPFIXUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPTrafficFlowIPFIXResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPTrafficFlowIPFIXModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPTrafficFlowIPFIXUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPTrafficFlowIPFIXResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPTrafficFlowIPFIXModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/traffic-flow/ipfix")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/traffic-flow/ipfix failed", err.Error())
		return
	}
	iPTrafficFlowIPFIXApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/traffic-flow/ipfix", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPTrafficFlowIPFIXResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPTrafficFlowIPFIXResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/traffic-flow/ipfix" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ip/traffic-flow/ipfix", types.StringValue(routerName))))...)
}

func iPTrafficFlowIPFIXUpsert(ctx context.Context, reg *client.Registry, plan *IPTrafficFlowIPFIXModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Bytes.IsNull() || plan.Bytes.IsUnknown()) {
		body["bytes"] = client.FormatBool(plan.Bytes.ValueBool())
	}
	if !(plan.DstAddress.IsNull() || plan.DstAddress.IsUnknown()) {
		body["dst-address"] = client.FormatBool(plan.DstAddress.ValueBool())
	}
	if !(plan.DstAddressMask.IsNull() || plan.DstAddressMask.IsUnknown()) {
		body["dst-address-mask"] = client.FormatBool(plan.DstAddressMask.ValueBool())
	}
	if !(plan.DstMACAddress.IsNull() || plan.DstMACAddress.IsUnknown()) {
		body["dst-mac-address"] = client.FormatBool(plan.DstMACAddress.ValueBool())
	}
	if !(plan.DstPort.IsNull() || plan.DstPort.IsUnknown()) {
		body["dst-port"] = client.FormatBool(plan.DstPort.ValueBool())
	}
	if !(plan.FirstForwarded.IsNull() || plan.FirstForwarded.IsUnknown()) {
		body["first-forwarded"] = client.FormatBool(plan.FirstForwarded.ValueBool())
	}
	if !(plan.Gateway.IsNull() || plan.Gateway.IsUnknown()) {
		body["gateway"] = client.FormatBool(plan.Gateway.ValueBool())
	}
	if !(plan.IcmpCode.IsNull() || plan.IcmpCode.IsUnknown()) {
		body["icmp-code"] = client.FormatBool(plan.IcmpCode.ValueBool())
	}
	if !(plan.IcmpType.IsNull() || plan.IcmpType.IsUnknown()) {
		body["icmp-type"] = client.FormatBool(plan.IcmpType.ValueBool())
	}
	if !(plan.IgmpType.IsNull() || plan.IgmpType.IsUnknown()) {
		body["igmp-type"] = client.FormatBool(plan.IgmpType.ValueBool())
	}
	if !(plan.InInterface.IsNull() || plan.InInterface.IsUnknown()) {
		body["in-interface"] = client.FormatBool(plan.InInterface.ValueBool())
	}
	if !(plan.IPHeaderLength.IsNull() || plan.IPHeaderLength.IsUnknown()) {
		body["ip-header-length"] = client.FormatBool(plan.IPHeaderLength.ValueBool())
	}
	if !(plan.IPTotalLength.IsNull() || plan.IPTotalLength.IsUnknown()) {
		body["ip-total-length"] = client.FormatBool(plan.IPTotalLength.ValueBool())
	}
	if !(plan.IPv6FlowLabel.IsNull() || plan.IPv6FlowLabel.IsUnknown()) {
		body["ipv6-flow-label"] = client.FormatBool(plan.IPv6FlowLabel.ValueBool())
	}
	if !(plan.IsMulticast.IsNull() || plan.IsMulticast.IsUnknown()) {
		body["is-multicast"] = client.FormatBool(plan.IsMulticast.ValueBool())
	}
	if !(plan.LastForwarded.IsNull() || plan.LastForwarded.IsUnknown()) {
		body["last-forwarded"] = client.FormatBool(plan.LastForwarded.ValueBool())
	}
	if !(plan.NatDstAddress.IsNull() || plan.NatDstAddress.IsUnknown()) {
		body["nat-dst-address"] = client.FormatBool(plan.NatDstAddress.ValueBool())
	}
	if !(plan.NatDstPort.IsNull() || plan.NatDstPort.IsUnknown()) {
		body["nat-dst-port"] = client.FormatBool(plan.NatDstPort.ValueBool())
	}
	if !(plan.NatEvents.IsNull() || plan.NatEvents.IsUnknown()) {
		body["nat-events"] = client.FormatBool(plan.NatEvents.ValueBool())
	}
	if !(plan.NatSrcAddress.IsNull() || plan.NatSrcAddress.IsUnknown()) {
		body["nat-src-address"] = client.FormatBool(plan.NatSrcAddress.ValueBool())
	}
	if !(plan.NatSrcPort.IsNull() || plan.NatSrcPort.IsUnknown()) {
		body["nat-src-port"] = client.FormatBool(plan.NatSrcPort.ValueBool())
	}
	if !(plan.OutInterface.IsNull() || plan.OutInterface.IsUnknown()) {
		body["out-interface"] = client.FormatBool(plan.OutInterface.ValueBool())
	}
	if !(plan.Packets.IsNull() || plan.Packets.IsUnknown()) {
		body["packets"] = client.FormatBool(plan.Packets.ValueBool())
	}
	if !(plan.Protocol.IsNull() || plan.Protocol.IsUnknown()) {
		body["protocol"] = client.FormatBool(plan.Protocol.ValueBool())
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = client.FormatBool(plan.SrcAddress.ValueBool())
	}
	if !(plan.SrcAddressMask.IsNull() || plan.SrcAddressMask.IsUnknown()) {
		body["src-address-mask"] = client.FormatBool(plan.SrcAddressMask.ValueBool())
	}
	if !(plan.SrcMACAddress.IsNull() || plan.SrcMACAddress.IsUnknown()) {
		body["src-mac-address"] = client.FormatBool(plan.SrcMACAddress.ValueBool())
	}
	if !(plan.SrcPort.IsNull() || plan.SrcPort.IsUnknown()) {
		body["src-port"] = client.FormatBool(plan.SrcPort.ValueBool())
	}
	if !(plan.SysInitTime.IsNull() || plan.SysInitTime.IsUnknown()) {
		body["sys-init-time"] = client.FormatBool(plan.SysInitTime.ValueBool())
	}
	if !(plan.TCPAckNum.IsNull() || plan.TCPAckNum.IsUnknown()) {
		body["tcp-ack-num"] = client.FormatBool(plan.TCPAckNum.ValueBool())
	}
	if !(plan.TCPFlags.IsNull() || plan.TCPFlags.IsUnknown()) {
		body["tcp-flags"] = client.FormatBool(plan.TCPFlags.ValueBool())
	}
	if !(plan.TCPSeqNum.IsNull() || plan.TCPSeqNum.IsUnknown()) {
		body["tcp-seq-num"] = client.FormatBool(plan.TCPSeqNum.ValueBool())
	}
	if !(plan.TCPWindowSize.IsNull() || plan.TCPWindowSize.IsUnknown()) {
		body["tcp-window-size"] = client.FormatBool(plan.TCPWindowSize.ValueBool())
	}
	if !(plan.Tos.IsNull() || plan.Tos.IsUnknown()) {
		body["tos"] = client.FormatBool(plan.Tos.ValueBool())
	}
	if !(plan.Ttl.IsNull() || plan.Ttl.IsUnknown()) {
		body["ttl"] = client.FormatBool(plan.Ttl.ValueBool())
	}
	if !(plan.UDPLength.IsNull() || plan.UDPLength.IsUnknown()) {
		body["udp-length"] = client.FormatBool(plan.UDPLength.ValueBool())
	}
	obj, err := c.SetSingleton(ctx, "/ip/traffic-flow/ipfix", body)
	if err != nil {
		diags.AddError("Upsert /ip/traffic-flow/ipfix failed", err.Error())
		return
	}
	iPTrafficFlowIPFIXApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/traffic-flow/ipfix", plan.Router))
}

func iPTrafficFlowIPFIXApply(ctx context.Context, obj client.Object, m *IPTrafficFlowIPFIXModel) {
	_ = ctx
	if v, ok := obj["bytes"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Bytes = types.BoolValue(b)
		} else {
			m.Bytes = types.BoolNull()
		}
	} else {
		m.Bytes = types.BoolNull()
	}
	if v, ok := obj["dst-address"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.DstAddress = types.BoolValue(b)
		} else {
			m.DstAddress = types.BoolNull()
		}
	} else {
		m.DstAddress = types.BoolNull()
	}
	if v, ok := obj["dst-address-mask"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.DstAddressMask = types.BoolValue(b)
		} else {
			m.DstAddressMask = types.BoolNull()
		}
	} else {
		m.DstAddressMask = types.BoolNull()
	}
	if v, ok := obj["dst-mac-address"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.DstMACAddress = types.BoolValue(b)
		} else {
			m.DstMACAddress = types.BoolNull()
		}
	} else {
		m.DstMACAddress = types.BoolNull()
	}
	if v, ok := obj["dst-port"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.DstPort = types.BoolValue(b)
		} else {
			m.DstPort = types.BoolNull()
		}
	} else {
		m.DstPort = types.BoolNull()
	}
	if v, ok := obj["first-forwarded"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.FirstForwarded = types.BoolValue(b)
		} else {
			m.FirstForwarded = types.BoolNull()
		}
	} else {
		m.FirstForwarded = types.BoolNull()
	}
	if v, ok := obj["gateway"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Gateway = types.BoolValue(b)
		} else {
			m.Gateway = types.BoolNull()
		}
	} else {
		m.Gateway = types.BoolNull()
	}
	if v, ok := obj["icmp-code"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.IcmpCode = types.BoolValue(b)
		} else {
			m.IcmpCode = types.BoolNull()
		}
	} else {
		m.IcmpCode = types.BoolNull()
	}
	if v, ok := obj["icmp-type"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.IcmpType = types.BoolValue(b)
		} else {
			m.IcmpType = types.BoolNull()
		}
	} else {
		m.IcmpType = types.BoolNull()
	}
	if v, ok := obj["igmp-type"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.IgmpType = types.BoolValue(b)
		} else {
			m.IgmpType = types.BoolNull()
		}
	} else {
		m.IgmpType = types.BoolNull()
	}
	if v, ok := obj["in-interface"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.InInterface = types.BoolValue(b)
		} else {
			m.InInterface = types.BoolNull()
		}
	} else {
		m.InInterface = types.BoolNull()
	}
	if v, ok := obj["ip-header-length"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.IPHeaderLength = types.BoolValue(b)
		} else {
			m.IPHeaderLength = types.BoolNull()
		}
	} else {
		m.IPHeaderLength = types.BoolNull()
	}
	if v, ok := obj["ip-total-length"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.IPTotalLength = types.BoolValue(b)
		} else {
			m.IPTotalLength = types.BoolNull()
		}
	} else {
		m.IPTotalLength = types.BoolNull()
	}
	if v, ok := obj["ipv6-flow-label"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.IPv6FlowLabel = types.BoolValue(b)
		} else {
			m.IPv6FlowLabel = types.BoolNull()
		}
	} else {
		m.IPv6FlowLabel = types.BoolNull()
	}
	if v, ok := obj["is-multicast"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.IsMulticast = types.BoolValue(b)
		} else {
			m.IsMulticast = types.BoolNull()
		}
	} else {
		m.IsMulticast = types.BoolNull()
	}
	if v, ok := obj["last-forwarded"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.LastForwarded = types.BoolValue(b)
		} else {
			m.LastForwarded = types.BoolNull()
		}
	} else {
		m.LastForwarded = types.BoolNull()
	}
	if v, ok := obj["nat-dst-address"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.NatDstAddress = types.BoolValue(b)
		} else {
			m.NatDstAddress = types.BoolNull()
		}
	} else {
		m.NatDstAddress = types.BoolNull()
	}
	if v, ok := obj["nat-dst-port"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.NatDstPort = types.BoolValue(b)
		} else {
			m.NatDstPort = types.BoolNull()
		}
	} else {
		m.NatDstPort = types.BoolNull()
	}
	if v, ok := obj["nat-events"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.NatEvents = types.BoolValue(b)
		} else {
			m.NatEvents = types.BoolNull()
		}
	} else {
		m.NatEvents = types.BoolNull()
	}
	if v, ok := obj["nat-src-address"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.NatSrcAddress = types.BoolValue(b)
		} else {
			m.NatSrcAddress = types.BoolNull()
		}
	} else {
		m.NatSrcAddress = types.BoolNull()
	}
	if v, ok := obj["nat-src-port"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.NatSrcPort = types.BoolValue(b)
		} else {
			m.NatSrcPort = types.BoolNull()
		}
	} else {
		m.NatSrcPort = types.BoolNull()
	}
	if v, ok := obj["out-interface"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.OutInterface = types.BoolValue(b)
		} else {
			m.OutInterface = types.BoolNull()
		}
	} else {
		m.OutInterface = types.BoolNull()
	}
	if v, ok := obj["packets"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Packets = types.BoolValue(b)
		} else {
			m.Packets = types.BoolNull()
		}
	} else {
		m.Packets = types.BoolNull()
	}
	if v, ok := obj["protocol"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Protocol = types.BoolValue(b)
		} else {
			m.Protocol = types.BoolNull()
		}
	} else {
		m.Protocol = types.BoolNull()
	}
	if v, ok := obj["src-address"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SrcAddress = types.BoolValue(b)
		} else {
			m.SrcAddress = types.BoolNull()
		}
	} else {
		m.SrcAddress = types.BoolNull()
	}
	if v, ok := obj["src-address-mask"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SrcAddressMask = types.BoolValue(b)
		} else {
			m.SrcAddressMask = types.BoolNull()
		}
	} else {
		m.SrcAddressMask = types.BoolNull()
	}
	if v, ok := obj["src-mac-address"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SrcMACAddress = types.BoolValue(b)
		} else {
			m.SrcMACAddress = types.BoolNull()
		}
	} else {
		m.SrcMACAddress = types.BoolNull()
	}
	if v, ok := obj["src-port"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SrcPort = types.BoolValue(b)
		} else {
			m.SrcPort = types.BoolNull()
		}
	} else {
		m.SrcPort = types.BoolNull()
	}
	if v, ok := obj["sys-init-time"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SysInitTime = types.BoolValue(b)
		} else {
			m.SysInitTime = types.BoolNull()
		}
	} else {
		m.SysInitTime = types.BoolNull()
	}
	if v, ok := obj["tcp-ack-num"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.TCPAckNum = types.BoolValue(b)
		} else {
			m.TCPAckNum = types.BoolNull()
		}
	} else {
		m.TCPAckNum = types.BoolNull()
	}
	if v, ok := obj["tcp-flags"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.TCPFlags = types.BoolValue(b)
		} else {
			m.TCPFlags = types.BoolNull()
		}
	} else {
		m.TCPFlags = types.BoolNull()
	}
	if v, ok := obj["tcp-seq-num"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.TCPSeqNum = types.BoolValue(b)
		} else {
			m.TCPSeqNum = types.BoolNull()
		}
	} else {
		m.TCPSeqNum = types.BoolNull()
	}
	if v, ok := obj["tcp-window-size"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.TCPWindowSize = types.BoolValue(b)
		} else {
			m.TCPWindowSize = types.BoolNull()
		}
	} else {
		m.TCPWindowSize = types.BoolNull()
	}
	if v, ok := obj["tos"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Tos = types.BoolValue(b)
		} else {
			m.Tos = types.BoolNull()
		}
	} else {
		m.Tos = types.BoolNull()
	}
	if v, ok := obj["ttl"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Ttl = types.BoolValue(b)
		} else {
			m.Ttl = types.BoolNull()
		}
	} else {
		m.Ttl = types.BoolNull()
	}
	if v, ok := obj["udp-length"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.UDPLength = types.BoolValue(b)
		} else {
			m.UDPLength = types.BoolNull()
		}
	} else {
		m.UDPLength = types.BoolNull()
	}
}
