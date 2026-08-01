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
	_ resource.Resource                = &ToolTrafficGeneratorPacketTemplateResource{}
	_ resource.ResourceWithImportState = &ToolTrafficGeneratorPacketTemplateResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type ToolTrafficGeneratorPacketTemplateResource struct {
	reg *client.Registry
}

type ToolTrafficGeneratorPacketTemplateModel struct {
	ID                        types.String `tfsdk:"id"`
	VlanProtocol              types.String `tfsdk:"vlan_protocol"`
	VlanPriority              types.String `tfsdk:"vlan_priority"`
	UdpSrcPort                types.String `tfsdk:"udp_src_port"`
	UdpDstPort                types.String `tfsdk:"udp_dst_port"`
	UdpChecksum               types.String `tfsdk:"udp_checksum"`
	SpecialFooter             types.String `tfsdk:"special_footer"`
	RawHeader                 types.String `tfsdk:"raw_header"`
	RandomRanges              types.String `tfsdk:"random_ranges"`
	RandomByteOffsetsAndMasks types.String `tfsdk:"random_byte_offsets_and_masks"`
	MacSrc                    types.String `tfsdk:"mac_src"`
	MacProtocol               types.String `tfsdk:"mac_protocol"`
	MacDst                    types.String `tfsdk:"mac_dst"`
	Ipv6TrafficClass          types.String `tfsdk:"ipv6_traffic_class"`
	Ipv6Src                   types.String `tfsdk:"ipv6_src"`
	Ipv6NextHeader            types.String `tfsdk:"ipv6_next_header"`
	Ipv6HopLimit              types.String `tfsdk:"ipv6_hop_limit"`
	Ipv6Gateway               types.String `tfsdk:"ipv6_gateway"`
	Ipv6FlowLabel             types.String `tfsdk:"ipv6_flow_label"`
	Ipv6Dst                   types.String `tfsdk:"ipv6_dst"`
	IpTtl                     types.String `tfsdk:"ip_ttl"`
	IpSrc                     types.String `tfsdk:"ip_src"`
	IpProtocol                types.String `tfsdk:"ip_protocol"`
	IpGateway                 types.String `tfsdk:"ip_gateway"`
	IpFragOff                 types.String `tfsdk:"ip_frag_off"`
	IpDst                     types.String `tfsdk:"ip_dst"`
	IpDscp                    types.String `tfsdk:"ip_dscp"`
	ComputeChecksumFromOffset types.String `tfsdk:"compute_checksum_from_offset"`
	AssumedDscpEcn            types.String `tfsdk:"assumed_dscp_ecn"`
	AssumedDst                types.String `tfsdk:"assumed_dst"`
	AssumedDstPort            types.String `tfsdk:"assumed_dst_port"`
	AssumedFlowLabel          types.String `tfsdk:"assumed_flow_label"`
	AssumedFragOffset         types.String `tfsdk:"assumed_frag_offset"`
	AssumedHeader             types.String `tfsdk:"assumed_header"`
	AssumedInterface          types.String `tfsdk:"assumed_interface"`
	AssumedIPID               types.String `tfsdk:"assumed_ip_id"`
	AssumedNextHeader         types.String `tfsdk:"assumed_next_header"`
	AssumedPort               types.String `tfsdk:"assumed_port"`
	AssumedPriority           types.String `tfsdk:"assumed_priority"`
	AssumedProtocol           types.String `tfsdk:"assumed_protocol"`
	AssumedSrc                types.String `tfsdk:"assumed_src"`
	AssumedSrcPort            types.String `tfsdk:"assumed_src_port"`
	AssumedTCPAck             types.String `tfsdk:"assumed_tcp_ack"`
	AssumedTCPDataOffset      types.String `tfsdk:"assumed_tcp_data_offset"`
	AssumedTCPDstPort         types.String `tfsdk:"assumed_tcp_dst_port"`
	AssumedTCPFlags           types.String `tfsdk:"assumed_tcp_flags"`
	AssumedTCPSrcPort         types.String `tfsdk:"assumed_tcp_src_port"`
	AssumedTCPSyn             types.String `tfsdk:"assumed_tcp_syn"`
	AssumedTCPUrgentPointer   types.String `tfsdk:"assumed_tcp_urgent_pointer"`
	AssumedTCPWindowSize      types.String `tfsdk:"assumed_tcp_window_size"`
	AssumedTrafficClass       types.String `tfsdk:"assumed_traffic_class"`
	AssumedTtl                types.String `tfsdk:"assumed_ttl"`
	AssumedVLANID             types.String `tfsdk:"assumed_vlan_id"`
	Comment                   types.String `tfsdk:"comment"`
	Data                      types.String `tfsdk:"data"`
	DataByte                  types.Int64  `tfsdk:"data_byte"`
	DscpEcn                   types.String `tfsdk:"dscp_ecn"`
	Dst                       types.String `tfsdk:"dst"`
	DstPort                   types.String `tfsdk:"dst_port"`
	FlowLabel                 types.String `tfsdk:"flow_label"`
	FragOffset                types.String `tfsdk:"frag_offset"`
	Gateway                   types.String `tfsdk:"gateway"`
	Header                    types.String `tfsdk:"header"`
	HeaderStack               types.String `tfsdk:"header_stack"`
	HopLimit                  types.String `tfsdk:"hop_limit"`
	Interface                 types.String `tfsdk:"interface"`
	IP                        types.String `tfsdk:"ip"`
	IPID                      types.String `tfsdk:"ip_id"`
	IPV6                      types.String `tfsdk:"ipv6"`
	MAC                       types.String `tfsdk:"mac"`
	Name                      types.String `tfsdk:"name"`
	NextHeader                types.String `tfsdk:"next_header"`
	Port                      types.String `tfsdk:"port"`
	Priority                  types.String `tfsdk:"priority"`
	Protocol                  types.String `tfsdk:"protocol"`
	Raw                       types.String `tfsdk:"raw"`
	RawPacketTemplates        types.String `tfsdk:"raw_packet_templates"`
	Specbyte                  types.String `tfsdk:"specbyte"`
	Src                       types.String `tfsdk:"src"`
	SrcPort                   types.String `tfsdk:"src_port"`
	TCP                       types.String `tfsdk:"tcp"`
	TCPAck                    types.String `tfsdk:"tcp_ack"`
	TCPDataOffset             types.String `tfsdk:"tcp_data_offset"`
	TCPDstPort                types.String `tfsdk:"tcp_dst_port"`
	TCPFlags                  types.String `tfsdk:"tcp_flags"`
	TCPSrcPort                types.String `tfsdk:"tcp_src_port"`
	TCPSyn                    types.String `tfsdk:"tcp_syn"`
	TCPUrgentPointer          types.String `tfsdk:"tcp_urgent_pointer"`
	TCPWindowSize             types.String `tfsdk:"tcp_window_size"`
	TrafficClass              types.String `tfsdk:"traffic_class"`
	Ttl                       types.String `tfsdk:"ttl"`
	UDP                       types.String `tfsdk:"udp"`
	VLAN                      types.String `tfsdk:"vlan"`
	VLANID                    types.String `tfsdk:"vlan_id"`
	Router                    types.String `tfsdk:"router"`
}

func NewToolTrafficGeneratorPacketTemplateResource() resource.Resource {
	return &ToolTrafficGeneratorPacketTemplateResource{}
}

func (r *ToolTrafficGeneratorPacketTemplateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_traffic_generator_packet_template"
}

func (r *ToolTrafficGeneratorPacketTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolTrafficGeneratorPacketTemplateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/tool/traffic-generator/packet-template`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"vlan_protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `vlan-protocol`.",
			},
			"vlan_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `vlan-priority`.",
			},
			"udp_src_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `udp-src-port`.",
			},
			"udp_dst_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `udp-dst-port`.",
			},
			"udp_checksum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `udp-checksum`.",
			},
			"special_footer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `special-footer`.",
			},
			"raw_header": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `raw-header`.",
			},
			"random_ranges": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `random-ranges`.",
			},
			"random_byte_offsets_and_masks": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `random-byte-offsets-and-masks`.",
			},
			"mac_src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mac-src`.",
			},
			"mac_protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mac-protocol`.",
			},
			"mac_dst": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mac-dst`.",
			},
			"ipv6_traffic_class": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ipv6-traffic-class`.",
			},
			"ipv6_src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ipv6-src`.",
			},
			"ipv6_next_header": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ipv6-next-header`.",
			},
			"ipv6_hop_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ipv6-hop-limit`.",
			},
			"ipv6_gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ipv6-gateway`.",
			},
			"ipv6_flow_label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ipv6-flow-label`.",
			},
			"ipv6_dst": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ipv6-dst`.",
			},
			"ip_ttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ip-ttl`.",
			},
			"ip_src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ip-src`.",
			},
			"ip_protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ip-protocol`.",
			},
			"ip_gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ip-gateway`.",
			},
			"ip_frag_off": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ip-frag-off`.",
			},
			"ip_dst": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ip-dst`.",
			},
			"ip_dscp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ip-dscp`.",
			},
			"compute_checksum_from_offset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `compute-checksum-from-offset`.",
			},
			"assumed_dscp_ecn": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_dst": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_dst_port": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_flow_label": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_frag_offset": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_header": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_interface": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_ip_id": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_next_header": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_port": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_priority": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_protocol": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_src": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_src_port": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_ack": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_data_offset": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_dst_port": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_flags": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_src_port": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_syn": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_urgent_pointer": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_window_size": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_traffic_class": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_ttl": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"assumed_vlan_id": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"data": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"uninitialized", "random", "specific-byte", "incrementing"}...)},
			},
			"data_byte": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dscp_ecn": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"dst": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"dst_port": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"flow_label": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"frag_offset": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"gateway": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"header": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"header_stack": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hop_limit": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ip": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"ip_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ipv6": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"mac": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"next_header": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"priority": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"protocol": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"raw": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"raw_packet_templates": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"specbyte": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"src": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"src_port": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tcp": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tcp_ack": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tcp_data_offset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tcp_dst_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tcp_flags": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tcp_src_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tcp_syn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tcp_urgent_pointer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tcp_window_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"traffic_class": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"ttl": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"udp": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"vlan": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"vlan_id": schema.StringAttribute{
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

func (r *ToolTrafficGeneratorPacketTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolTrafficGeneratorPacketTemplateModel
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
	if !(plan.Data.IsNull() || plan.Data.IsUnknown()) {
		body["data"] = plan.Data.ValueString()
	}
	if !(plan.DataByte.IsNull() || plan.DataByte.IsUnknown()) {
		body["data-byte"] = client.FormatInt64(plan.DataByte.ValueInt64())
	}
	if !(plan.HeaderStack.IsNull() || plan.HeaderStack.IsUnknown()) {
		body["header-stack"] = plan.HeaderStack.ValueString()
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.IPID.IsNull() || plan.IPID.IsUnknown()) {
		body["ip-id"] = plan.IPID.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.TCPAck.IsNull() || plan.TCPAck.IsUnknown()) {
		body["tcp-ack"] = plan.TCPAck.ValueString()
	}
	if !(plan.TCPDataOffset.IsNull() || plan.TCPDataOffset.IsUnknown()) {
		body["tcp-data-offset"] = plan.TCPDataOffset.ValueString()
	}
	if !(plan.TCPDstPort.IsNull() || plan.TCPDstPort.IsUnknown()) {
		body["tcp-dst-port"] = plan.TCPDstPort.ValueString()
	}
	if !(plan.TCPFlags.IsNull() || plan.TCPFlags.IsUnknown()) {
		body["tcp-flags"] = plan.TCPFlags.ValueString()
	}
	if !(plan.TCPSrcPort.IsNull() || plan.TCPSrcPort.IsUnknown()) {
		body["tcp-src-port"] = plan.TCPSrcPort.ValueString()
	}
	if !(plan.TCPSyn.IsNull() || plan.TCPSyn.IsUnknown()) {
		body["tcp-syn"] = plan.TCPSyn.ValueString()
	}
	if !(plan.TCPUrgentPointer.IsNull() || plan.TCPUrgentPointer.IsUnknown()) {
		body["tcp-urgent-pointer"] = plan.TCPUrgentPointer.ValueString()
	}
	if !(plan.TCPWindowSize.IsNull() || plan.TCPWindowSize.IsUnknown()) {
		body["tcp-window-size"] = plan.TCPWindowSize.ValueString()
	}
	if !(plan.VLANID.IsNull() || plan.VLANID.IsUnknown()) {
		body["vlan-id"] = plan.VLANID.ValueString()
	}
	if !(plan.ComputeChecksumFromOffset.IsNull() || plan.ComputeChecksumFromOffset.IsUnknown()) {
		body["compute-checksum-from-offset"] = plan.ComputeChecksumFromOffset.ValueString()
	}
	if !(plan.IpDscp.IsNull() || plan.IpDscp.IsUnknown()) {
		body["ip-dscp"] = plan.IpDscp.ValueString()
	}
	if !(plan.IpDst.IsNull() || plan.IpDst.IsUnknown()) {
		body["ip-dst"] = plan.IpDst.ValueString()
	}
	if !(plan.IpFragOff.IsNull() || plan.IpFragOff.IsUnknown()) {
		body["ip-frag-off"] = plan.IpFragOff.ValueString()
	}
	if !(plan.IpGateway.IsNull() || plan.IpGateway.IsUnknown()) {
		body["ip-gateway"] = plan.IpGateway.ValueString()
	}
	if !(plan.IpProtocol.IsNull() || plan.IpProtocol.IsUnknown()) {
		body["ip-protocol"] = plan.IpProtocol.ValueString()
	}
	if !(plan.IpSrc.IsNull() || plan.IpSrc.IsUnknown()) {
		body["ip-src"] = plan.IpSrc.ValueString()
	}
	if !(plan.IpTtl.IsNull() || plan.IpTtl.IsUnknown()) {
		body["ip-ttl"] = plan.IpTtl.ValueString()
	}
	if !(plan.Ipv6Dst.IsNull() || plan.Ipv6Dst.IsUnknown()) {
		body["ipv6-dst"] = plan.Ipv6Dst.ValueString()
	}
	if !(plan.Ipv6FlowLabel.IsNull() || plan.Ipv6FlowLabel.IsUnknown()) {
		body["ipv6-flow-label"] = plan.Ipv6FlowLabel.ValueString()
	}
	if !(plan.Ipv6Gateway.IsNull() || plan.Ipv6Gateway.IsUnknown()) {
		body["ipv6-gateway"] = plan.Ipv6Gateway.ValueString()
	}
	if !(plan.Ipv6HopLimit.IsNull() || plan.Ipv6HopLimit.IsUnknown()) {
		body["ipv6-hop-limit"] = plan.Ipv6HopLimit.ValueString()
	}
	if !(plan.Ipv6NextHeader.IsNull() || plan.Ipv6NextHeader.IsUnknown()) {
		body["ipv6-next-header"] = plan.Ipv6NextHeader.ValueString()
	}
	if !(plan.Ipv6Src.IsNull() || plan.Ipv6Src.IsUnknown()) {
		body["ipv6-src"] = plan.Ipv6Src.ValueString()
	}
	if !(plan.Ipv6TrafficClass.IsNull() || plan.Ipv6TrafficClass.IsUnknown()) {
		body["ipv6-traffic-class"] = plan.Ipv6TrafficClass.ValueString()
	}
	if !(plan.MacDst.IsNull() || plan.MacDst.IsUnknown()) {
		body["mac-dst"] = plan.MacDst.ValueString()
	}
	if !(plan.MacProtocol.IsNull() || plan.MacProtocol.IsUnknown()) {
		body["mac-protocol"] = plan.MacProtocol.ValueString()
	}
	if !(plan.MacSrc.IsNull() || plan.MacSrc.IsUnknown()) {
		body["mac-src"] = plan.MacSrc.ValueString()
	}
	if !(plan.RandomByteOffsetsAndMasks.IsNull() || plan.RandomByteOffsetsAndMasks.IsUnknown()) {
		body["random-byte-offsets-and-masks"] = plan.RandomByteOffsetsAndMasks.ValueString()
	}
	if !(plan.RandomRanges.IsNull() || plan.RandomRanges.IsUnknown()) {
		body["random-ranges"] = plan.RandomRanges.ValueString()
	}
	if !(plan.RawHeader.IsNull() || plan.RawHeader.IsUnknown()) {
		body["raw-header"] = plan.RawHeader.ValueString()
	}
	if !(plan.SpecialFooter.IsNull() || plan.SpecialFooter.IsUnknown()) {
		body["special-footer"] = plan.SpecialFooter.ValueString()
	}
	if !(plan.UdpChecksum.IsNull() || plan.UdpChecksum.IsUnknown()) {
		body["udp-checksum"] = plan.UdpChecksum.ValueString()
	}
	if !(plan.UdpDstPort.IsNull() || plan.UdpDstPort.IsUnknown()) {
		body["udp-dst-port"] = plan.UdpDstPort.ValueString()
	}
	if !(plan.UdpSrcPort.IsNull() || plan.UdpSrcPort.IsUnknown()) {
		body["udp-src-port"] = plan.UdpSrcPort.ValueString()
	}
	if !(plan.VlanPriority.IsNull() || plan.VlanPriority.IsUnknown()) {
		body["vlan-priority"] = plan.VlanPriority.ValueString()
	}
	if !(plan.VlanProtocol.IsNull() || plan.VlanProtocol.IsUnknown()) {
		body["vlan-protocol"] = plan.VlanProtocol.ValueString()
	}
	obj, err := c.Add(ctx, "/tool/traffic-generator/packet-template", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /tool/traffic-generator/packet-template failed", err.Error())
		return
	}
	toolTrafficGeneratorPacketTemplateApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolTrafficGeneratorPacketTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolTrafficGeneratorPacketTemplateModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/tool/traffic-generator/packet-template", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /tool/traffic-generator/packet-template failed", err.Error())
		return
	}
	toolTrafficGeneratorPacketTemplateApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolTrafficGeneratorPacketTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ToolTrafficGeneratorPacketTemplateModel
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
	if !plan.Data.Equal(state.Data) && !plan.Data.IsUnknown() {
		body["data"] = plan.Data.ValueString()
	}
	if !plan.DataByte.Equal(state.DataByte) && !plan.DataByte.IsUnknown() {
		body["data-byte"] = client.FormatInt64(plan.DataByte.ValueInt64())
	}
	if !plan.HeaderStack.Equal(state.HeaderStack) && !plan.HeaderStack.IsUnknown() {
		body["header-stack"] = plan.HeaderStack.ValueString()
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.IPID.Equal(state.IPID) && !plan.IPID.IsUnknown() {
		body["ip-id"] = plan.IPID.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.TCPAck.Equal(state.TCPAck) && !plan.TCPAck.IsUnknown() {
		body["tcp-ack"] = plan.TCPAck.ValueString()
	}
	if !plan.TCPDataOffset.Equal(state.TCPDataOffset) && !plan.TCPDataOffset.IsUnknown() {
		body["tcp-data-offset"] = plan.TCPDataOffset.ValueString()
	}
	if !plan.TCPDstPort.Equal(state.TCPDstPort) && !plan.TCPDstPort.IsUnknown() {
		body["tcp-dst-port"] = plan.TCPDstPort.ValueString()
	}
	if !plan.TCPFlags.Equal(state.TCPFlags) && !plan.TCPFlags.IsUnknown() {
		body["tcp-flags"] = plan.TCPFlags.ValueString()
	}
	if !plan.TCPSrcPort.Equal(state.TCPSrcPort) && !plan.TCPSrcPort.IsUnknown() {
		body["tcp-src-port"] = plan.TCPSrcPort.ValueString()
	}
	if !plan.TCPSyn.Equal(state.TCPSyn) && !plan.TCPSyn.IsUnknown() {
		body["tcp-syn"] = plan.TCPSyn.ValueString()
	}
	if !plan.TCPUrgentPointer.Equal(state.TCPUrgentPointer) && !plan.TCPUrgentPointer.IsUnknown() {
		body["tcp-urgent-pointer"] = plan.TCPUrgentPointer.ValueString()
	}
	if !plan.TCPWindowSize.Equal(state.TCPWindowSize) && !plan.TCPWindowSize.IsUnknown() {
		body["tcp-window-size"] = plan.TCPWindowSize.ValueString()
	}
	if !plan.VLANID.Equal(state.VLANID) && !plan.VLANID.IsUnknown() {
		body["vlan-id"] = plan.VLANID.ValueString()
	}
	if !plan.ComputeChecksumFromOffset.Equal(state.ComputeChecksumFromOffset) && !plan.ComputeChecksumFromOffset.IsUnknown() {
		body["compute-checksum-from-offset"] = plan.ComputeChecksumFromOffset.ValueString()
	}
	if !plan.IpDscp.Equal(state.IpDscp) && !plan.IpDscp.IsUnknown() {
		body["ip-dscp"] = plan.IpDscp.ValueString()
	}
	if !plan.IpDst.Equal(state.IpDst) && !plan.IpDst.IsUnknown() {
		body["ip-dst"] = plan.IpDst.ValueString()
	}
	if !plan.IpFragOff.Equal(state.IpFragOff) && !plan.IpFragOff.IsUnknown() {
		body["ip-frag-off"] = plan.IpFragOff.ValueString()
	}
	if !plan.IpGateway.Equal(state.IpGateway) && !plan.IpGateway.IsUnknown() {
		body["ip-gateway"] = plan.IpGateway.ValueString()
	}
	if !plan.IpProtocol.Equal(state.IpProtocol) && !plan.IpProtocol.IsUnknown() {
		body["ip-protocol"] = plan.IpProtocol.ValueString()
	}
	if !plan.IpSrc.Equal(state.IpSrc) && !plan.IpSrc.IsUnknown() {
		body["ip-src"] = plan.IpSrc.ValueString()
	}
	if !plan.IpTtl.Equal(state.IpTtl) && !plan.IpTtl.IsUnknown() {
		body["ip-ttl"] = plan.IpTtl.ValueString()
	}
	if !plan.Ipv6Dst.Equal(state.Ipv6Dst) && !plan.Ipv6Dst.IsUnknown() {
		body["ipv6-dst"] = plan.Ipv6Dst.ValueString()
	}
	if !plan.Ipv6FlowLabel.Equal(state.Ipv6FlowLabel) && !plan.Ipv6FlowLabel.IsUnknown() {
		body["ipv6-flow-label"] = plan.Ipv6FlowLabel.ValueString()
	}
	if !plan.Ipv6Gateway.Equal(state.Ipv6Gateway) && !plan.Ipv6Gateway.IsUnknown() {
		body["ipv6-gateway"] = plan.Ipv6Gateway.ValueString()
	}
	if !plan.Ipv6HopLimit.Equal(state.Ipv6HopLimit) && !plan.Ipv6HopLimit.IsUnknown() {
		body["ipv6-hop-limit"] = plan.Ipv6HopLimit.ValueString()
	}
	if !plan.Ipv6NextHeader.Equal(state.Ipv6NextHeader) && !plan.Ipv6NextHeader.IsUnknown() {
		body["ipv6-next-header"] = plan.Ipv6NextHeader.ValueString()
	}
	if !plan.Ipv6Src.Equal(state.Ipv6Src) && !plan.Ipv6Src.IsUnknown() {
		body["ipv6-src"] = plan.Ipv6Src.ValueString()
	}
	if !plan.Ipv6TrafficClass.Equal(state.Ipv6TrafficClass) && !plan.Ipv6TrafficClass.IsUnknown() {
		body["ipv6-traffic-class"] = plan.Ipv6TrafficClass.ValueString()
	}
	if !plan.MacDst.Equal(state.MacDst) && !plan.MacDst.IsUnknown() {
		body["mac-dst"] = plan.MacDst.ValueString()
	}
	if !plan.MacProtocol.Equal(state.MacProtocol) && !plan.MacProtocol.IsUnknown() {
		body["mac-protocol"] = plan.MacProtocol.ValueString()
	}
	if !plan.MacSrc.Equal(state.MacSrc) && !plan.MacSrc.IsUnknown() {
		body["mac-src"] = plan.MacSrc.ValueString()
	}
	if !plan.RandomByteOffsetsAndMasks.Equal(state.RandomByteOffsetsAndMasks) && !plan.RandomByteOffsetsAndMasks.IsUnknown() {
		body["random-byte-offsets-and-masks"] = plan.RandomByteOffsetsAndMasks.ValueString()
	}
	if !plan.RandomRanges.Equal(state.RandomRanges) && !plan.RandomRanges.IsUnknown() {
		body["random-ranges"] = plan.RandomRanges.ValueString()
	}
	if !plan.RawHeader.Equal(state.RawHeader) && !plan.RawHeader.IsUnknown() {
		body["raw-header"] = plan.RawHeader.ValueString()
	}
	if !plan.SpecialFooter.Equal(state.SpecialFooter) && !plan.SpecialFooter.IsUnknown() {
		body["special-footer"] = plan.SpecialFooter.ValueString()
	}
	if !plan.UdpChecksum.Equal(state.UdpChecksum) && !plan.UdpChecksum.IsUnknown() {
		body["udp-checksum"] = plan.UdpChecksum.ValueString()
	}
	if !plan.UdpDstPort.Equal(state.UdpDstPort) && !plan.UdpDstPort.IsUnknown() {
		body["udp-dst-port"] = plan.UdpDstPort.ValueString()
	}
	if !plan.UdpSrcPort.Equal(state.UdpSrcPort) && !plan.UdpSrcPort.IsUnknown() {
		body["udp-src-port"] = plan.UdpSrcPort.ValueString()
	}
	if !plan.VlanPriority.Equal(state.VlanPriority) && !plan.VlanPriority.IsUnknown() {
		body["vlan-priority"] = plan.VlanPriority.ValueString()
	}
	if !plan.VlanProtocol.Equal(state.VlanProtocol) && !plan.VlanProtocol.IsUnknown() {
		body["vlan-protocol"] = plan.VlanProtocol.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/tool/traffic-generator/packet-template", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /tool/traffic-generator/packet-template failed", err.Error())
			return
		}
		toolTrafficGeneratorPacketTemplateApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolTrafficGeneratorPacketTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ToolTrafficGeneratorPacketTemplateModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/tool/traffic-generator/packet-template", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /tool/traffic-generator/packet-template failed", err.Error())
	}
}

func (r *ToolTrafficGeneratorPacketTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := toolTrafficGeneratorPacketTemplateLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /tool/traffic-generator/packet-template matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// toolTrafficGeneratorPacketTemplateLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func toolTrafficGeneratorPacketTemplateLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/tool/traffic-generator/packet-template", id)
}

func toolTrafficGeneratorPacketTemplateApply(ctx context.Context, obj client.Object, m *ToolTrafficGeneratorPacketTemplateModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["vlan-protocol"]; ok && v != "" {
		m.VlanProtocol = types.StringValue(v)
	} else {
		m.VlanProtocol = types.StringNull()
	}
	if v, ok := obj["vlan-priority"]; ok && v != "" {
		m.VlanPriority = types.StringValue(v)
	} else {
		m.VlanPriority = types.StringNull()
	}
	if v, ok := obj["udp-src-port"]; ok && v != "" {
		m.UdpSrcPort = types.StringValue(v)
	} else {
		m.UdpSrcPort = types.StringNull()
	}
	if v, ok := obj["udp-dst-port"]; ok && v != "" {
		m.UdpDstPort = types.StringValue(v)
	} else {
		m.UdpDstPort = types.StringNull()
	}
	if v, ok := obj["udp-checksum"]; ok && v != "" {
		m.UdpChecksum = types.StringValue(v)
	} else {
		m.UdpChecksum = types.StringNull()
	}
	if v, ok := obj["special-footer"]; ok && v != "" {
		m.SpecialFooter = types.StringValue(v)
	} else {
		m.SpecialFooter = types.StringNull()
	}
	if v, ok := obj["raw-header"]; ok && v != "" {
		m.RawHeader = types.StringValue(v)
	} else {
		m.RawHeader = types.StringNull()
	}
	if v, ok := obj["random-ranges"]; ok && v != "" {
		m.RandomRanges = types.StringValue(v)
	} else {
		m.RandomRanges = types.StringNull()
	}
	if v, ok := obj["random-byte-offsets-and-masks"]; ok && v != "" {
		m.RandomByteOffsetsAndMasks = types.StringValue(v)
	} else {
		m.RandomByteOffsetsAndMasks = types.StringNull()
	}
	if v, ok := obj["mac-src"]; ok && v != "" {
		m.MacSrc = types.StringValue(v)
	} else {
		m.MacSrc = types.StringNull()
	}
	if v, ok := obj["mac-protocol"]; ok && v != "" {
		m.MacProtocol = types.StringValue(v)
	} else {
		m.MacProtocol = types.StringNull()
	}
	if v, ok := obj["mac-dst"]; ok && v != "" {
		m.MacDst = types.StringValue(v)
	} else {
		m.MacDst = types.StringNull()
	}
	if v, ok := obj["ipv6-traffic-class"]; ok && v != "" {
		m.Ipv6TrafficClass = types.StringValue(v)
	} else {
		m.Ipv6TrafficClass = types.StringNull()
	}
	if v, ok := obj["ipv6-src"]; ok && v != "" {
		m.Ipv6Src = types.StringValue(v)
	} else {
		m.Ipv6Src = types.StringNull()
	}
	if v, ok := obj["ipv6-next-header"]; ok && v != "" {
		m.Ipv6NextHeader = types.StringValue(v)
	} else {
		m.Ipv6NextHeader = types.StringNull()
	}
	if v, ok := obj["ipv6-hop-limit"]; ok && v != "" {
		m.Ipv6HopLimit = types.StringValue(v)
	} else {
		m.Ipv6HopLimit = types.StringNull()
	}
	if v, ok := obj["ipv6-gateway"]; ok && v != "" {
		m.Ipv6Gateway = types.StringValue(v)
	} else {
		m.Ipv6Gateway = types.StringNull()
	}
	if v, ok := obj["ipv6-flow-label"]; ok && v != "" {
		m.Ipv6FlowLabel = types.StringValue(v)
	} else {
		m.Ipv6FlowLabel = types.StringNull()
	}
	if v, ok := obj["ipv6-dst"]; ok && v != "" {
		m.Ipv6Dst = types.StringValue(v)
	} else {
		m.Ipv6Dst = types.StringNull()
	}
	if v, ok := obj["ip-ttl"]; ok && v != "" {
		m.IpTtl = types.StringValue(v)
	} else {
		m.IpTtl = types.StringNull()
	}
	if v, ok := obj["ip-src"]; ok && v != "" {
		m.IpSrc = types.StringValue(v)
	} else {
		m.IpSrc = types.StringNull()
	}
	if v, ok := obj["ip-protocol"]; ok && v != "" {
		m.IpProtocol = types.StringValue(v)
	} else {
		m.IpProtocol = types.StringNull()
	}
	if v, ok := obj["ip-gateway"]; ok && v != "" {
		m.IpGateway = types.StringValue(v)
	} else {
		m.IpGateway = types.StringNull()
	}
	if v, ok := obj["ip-frag-off"]; ok && v != "" {
		m.IpFragOff = types.StringValue(v)
	} else {
		m.IpFragOff = types.StringNull()
	}
	if v, ok := obj["ip-dst"]; ok && v != "" {
		m.IpDst = types.StringValue(v)
	} else {
		m.IpDst = types.StringNull()
	}
	if v, ok := obj["ip-dscp"]; ok && v != "" {
		m.IpDscp = types.StringValue(v)
	} else {
		m.IpDscp = types.StringNull()
	}
	if v, ok := obj["compute-checksum-from-offset"]; ok && v != "" {
		m.ComputeChecksumFromOffset = types.StringValue(v)
	} else {
		m.ComputeChecksumFromOffset = types.StringNull()
	}
	if v, ok := obj["assumed-dscp-ecn"]; ok {
		_ = v
		if v != "" {
			m.AssumedDscpEcn = types.StringValue(v)
		} else {
			m.AssumedDscpEcn = types.StringNull()
		}
	} else {
		m.AssumedDscpEcn = types.StringNull()
	}
	if v, ok := obj["assumed-dst"]; ok {
		_ = v
		if v != "" {
			m.AssumedDst = types.StringValue(v)
		} else {
			m.AssumedDst = types.StringNull()
		}
	} else {
		m.AssumedDst = types.StringNull()
	}
	if v, ok := obj["assumed-dst-port"]; ok {
		_ = v
		if v != "" {
			m.AssumedDstPort = types.StringValue(v)
		} else {
			m.AssumedDstPort = types.StringNull()
		}
	} else {
		m.AssumedDstPort = types.StringNull()
	}
	if v, ok := obj["assumed-flow-label"]; ok {
		_ = v
		if v != "" {
			m.AssumedFlowLabel = types.StringValue(v)
		} else {
			m.AssumedFlowLabel = types.StringNull()
		}
	} else {
		m.AssumedFlowLabel = types.StringNull()
	}
	if v, ok := obj["assumed-frag-offset"]; ok {
		_ = v
		if v != "" {
			m.AssumedFragOffset = types.StringValue(v)
		} else {
			m.AssumedFragOffset = types.StringNull()
		}
	} else {
		m.AssumedFragOffset = types.StringNull()
	}
	if v, ok := obj["assumed-header"]; ok {
		_ = v
		if v != "" {
			m.AssumedHeader = types.StringValue(v)
		} else {
			m.AssumedHeader = types.StringNull()
		}
	} else {
		m.AssumedHeader = types.StringNull()
	}
	if v, ok := obj["assumed-interface"]; ok {
		_ = v
		if v != "" {
			m.AssumedInterface = types.StringValue(v)
		} else {
			m.AssumedInterface = types.StringNull()
		}
	} else {
		m.AssumedInterface = types.StringNull()
	}
	if v, ok := obj["assumed-ip-id"]; ok {
		_ = v
		if v != "" {
			m.AssumedIPID = types.StringValue(v)
		} else {
			m.AssumedIPID = types.StringNull()
		}
	} else {
		m.AssumedIPID = types.StringNull()
	}
	if v, ok := obj["assumed-next-header"]; ok {
		_ = v
		if v != "" {
			m.AssumedNextHeader = types.StringValue(v)
		} else {
			m.AssumedNextHeader = types.StringNull()
		}
	} else {
		m.AssumedNextHeader = types.StringNull()
	}
	if v, ok := obj["assumed-port"]; ok {
		_ = v
		if v != "" {
			m.AssumedPort = types.StringValue(v)
		} else {
			m.AssumedPort = types.StringNull()
		}
	} else {
		m.AssumedPort = types.StringNull()
	}
	if v, ok := obj["assumed-priority"]; ok {
		_ = v
		if v != "" {
			m.AssumedPriority = types.StringValue(v)
		} else {
			m.AssumedPriority = types.StringNull()
		}
	} else {
		m.AssumedPriority = types.StringNull()
	}
	if v, ok := obj["assumed-protocol"]; ok {
		_ = v
		if v != "" {
			m.AssumedProtocol = types.StringValue(v)
		} else {
			m.AssumedProtocol = types.StringNull()
		}
	} else {
		m.AssumedProtocol = types.StringNull()
	}
	if v, ok := obj["assumed-src"]; ok {
		_ = v
		if v != "" {
			m.AssumedSrc = types.StringValue(v)
		} else {
			m.AssumedSrc = types.StringNull()
		}
	} else {
		m.AssumedSrc = types.StringNull()
	}
	if v, ok := obj["assumed-src-port"]; ok {
		_ = v
		if v != "" {
			m.AssumedSrcPort = types.StringValue(v)
		} else {
			m.AssumedSrcPort = types.StringNull()
		}
	} else {
		m.AssumedSrcPort = types.StringNull()
	}
	if v, ok := obj["assumed-tcp-ack"]; ok {
		_ = v
		if v != "" {
			m.AssumedTCPAck = types.StringValue(v)
		} else {
			m.AssumedTCPAck = types.StringNull()
		}
	} else {
		m.AssumedTCPAck = types.StringNull()
	}
	if v, ok := obj["assumed-tcp-data-offset"]; ok {
		_ = v
		if v != "" {
			m.AssumedTCPDataOffset = types.StringValue(v)
		} else {
			m.AssumedTCPDataOffset = types.StringNull()
		}
	} else {
		m.AssumedTCPDataOffset = types.StringNull()
	}
	if v, ok := obj["assumed-tcp-dst-port"]; ok {
		_ = v
		if v != "" {
			m.AssumedTCPDstPort = types.StringValue(v)
		} else {
			m.AssumedTCPDstPort = types.StringNull()
		}
	} else {
		m.AssumedTCPDstPort = types.StringNull()
	}
	if v, ok := obj["assumed-tcp-flags"]; ok {
		_ = v
		if v != "" {
			m.AssumedTCPFlags = types.StringValue(v)
		} else {
			m.AssumedTCPFlags = types.StringNull()
		}
	} else {
		m.AssumedTCPFlags = types.StringNull()
	}
	if v, ok := obj["assumed-tcp-src-port"]; ok {
		_ = v
		if v != "" {
			m.AssumedTCPSrcPort = types.StringValue(v)
		} else {
			m.AssumedTCPSrcPort = types.StringNull()
		}
	} else {
		m.AssumedTCPSrcPort = types.StringNull()
	}
	if v, ok := obj["assumed-tcp-syn"]; ok {
		_ = v
		if v != "" {
			m.AssumedTCPSyn = types.StringValue(v)
		} else {
			m.AssumedTCPSyn = types.StringNull()
		}
	} else {
		m.AssumedTCPSyn = types.StringNull()
	}
	if v, ok := obj["assumed-tcp-urgent-pointer"]; ok {
		_ = v
		if v != "" {
			m.AssumedTCPUrgentPointer = types.StringValue(v)
		} else {
			m.AssumedTCPUrgentPointer = types.StringNull()
		}
	} else {
		m.AssumedTCPUrgentPointer = types.StringNull()
	}
	if v, ok := obj["assumed-tcp-window-size"]; ok {
		_ = v
		if v != "" {
			m.AssumedTCPWindowSize = types.StringValue(v)
		} else {
			m.AssumedTCPWindowSize = types.StringNull()
		}
	} else {
		m.AssumedTCPWindowSize = types.StringNull()
	}
	if v, ok := obj["assumed-traffic-class"]; ok {
		_ = v
		if v != "" {
			m.AssumedTrafficClass = types.StringValue(v)
		} else {
			m.AssumedTrafficClass = types.StringNull()
		}
	} else {
		m.AssumedTrafficClass = types.StringNull()
	}
	if v, ok := obj["assumed-ttl"]; ok {
		_ = v
		if v != "" {
			m.AssumedTtl = types.StringValue(v)
		} else {
			m.AssumedTtl = types.StringNull()
		}
	} else {
		m.AssumedTtl = types.StringNull()
	}
	if v, ok := obj["assumed-vlan-id"]; ok {
		_ = v
		if v != "" {
			m.AssumedVLANID = types.StringValue(v)
		} else {
			m.AssumedVLANID = types.StringNull()
		}
	} else {
		m.AssumedVLANID = types.StringNull()
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
	if v, ok := obj["data"]; ok {
		_ = v
		if v != "" {
			m.Data = types.StringValue(v)
		} else {
			m.Data = types.StringNull()
		}
	} else {
		m.Data = types.StringNull()
	}
	if v, ok := obj["data-byte"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.DataByte = types.Int64Value(n)
		} else {
			m.DataByte = types.Int64Null()
		}
	} else {
		m.DataByte = types.Int64Null()
	}
	if v, ok := obj["dscp-ecn"]; ok {
		_ = v
		if v != "" {
			m.DscpEcn = types.StringValue(v)
		} else {
			m.DscpEcn = types.StringNull()
		}
	} else {
		m.DscpEcn = types.StringNull()
	}
	if v, ok := obj["dst"]; ok {
		_ = v
		if v != "" {
			m.Dst = types.StringValue(v)
		} else {
			m.Dst = types.StringNull()
		}
	} else {
		m.Dst = types.StringNull()
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
	if v, ok := obj["flow-label"]; ok {
		_ = v
		if v != "" {
			m.FlowLabel = types.StringValue(v)
		} else {
			m.FlowLabel = types.StringNull()
		}
	} else {
		m.FlowLabel = types.StringNull()
	}
	if v, ok := obj["frag-offset"]; ok {
		_ = v
		if v != "" {
			m.FragOffset = types.StringValue(v)
		} else {
			m.FragOffset = types.StringNull()
		}
	} else {
		m.FragOffset = types.StringNull()
	}
	if v, ok := obj["gateway"]; ok {
		_ = v
		if v != "" {
			m.Gateway = types.StringValue(v)
		} else {
			m.Gateway = types.StringNull()
		}
	} else {
		m.Gateway = types.StringNull()
	}
	if v, ok := obj["header"]; ok {
		_ = v
		if v != "" {
			m.Header = types.StringValue(v)
		} else {
			m.Header = types.StringNull()
		}
	} else {
		m.Header = types.StringNull()
	}
	if v, ok := obj["header-stack"]; ok {
		_ = v
		if v != "" {
			m.HeaderStack = types.StringValue(v)
		} else {
			m.HeaderStack = types.StringNull()
		}
	} else {
		m.HeaderStack = types.StringNull()
	}
	if v, ok := obj["hop-limit"]; ok {
		_ = v
		if v != "" {
			m.HopLimit = types.StringValue(v)
		} else {
			m.HopLimit = types.StringNull()
		}
	} else {
		m.HopLimit = types.StringNull()
	}
	if v, ok := obj["interface"]; ok {
		_ = v
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	} else {
		m.Interface = types.StringNull()
	}
	if v, ok := obj["ip"]; ok {
		_ = v
		if v != "" {
			m.IP = types.StringValue(v)
		} else {
			m.IP = types.StringNull()
		}
	} else {
		m.IP = types.StringNull()
	}
	if v, ok := obj["ip-id"]; ok {
		_ = v
		if v != "" {
			m.IPID = types.StringValue(v)
		} else {
			m.IPID = types.StringNull()
		}
	} else {
		m.IPID = types.StringNull()
	}
	if v, ok := obj["ipv6"]; ok {
		_ = v
		if v != "" {
			m.IPV6 = types.StringValue(v)
		} else {
			m.IPV6 = types.StringNull()
		}
	} else {
		m.IPV6 = types.StringNull()
	}
	if v, ok := obj["mac"]; ok {
		_ = v
		if v != "" {
			m.MAC = types.StringValue(v)
		} else {
			m.MAC = types.StringNull()
		}
	} else {
		m.MAC = types.StringNull()
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
	if v, ok := obj["next-header"]; ok {
		_ = v
		if v != "" {
			m.NextHeader = types.StringValue(v)
		} else {
			m.NextHeader = types.StringNull()
		}
	} else {
		m.NextHeader = types.StringNull()
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
	if v, ok := obj["raw"]; ok {
		_ = v
		if v != "" {
			m.Raw = types.StringValue(v)
		} else {
			m.Raw = types.StringNull()
		}
	} else {
		m.Raw = types.StringNull()
	}
	if v, ok := obj["raw-packet-templates"]; ok {
		_ = v
		if v != "" {
			m.RawPacketTemplates = types.StringValue(v)
		} else {
			m.RawPacketTemplates = types.StringNull()
		}
	} else {
		m.RawPacketTemplates = types.StringNull()
	}
	if v, ok := obj["specbyte"]; ok {
		_ = v
		if v != "" {
			m.Specbyte = types.StringValue(v)
		} else {
			m.Specbyte = types.StringNull()
		}
	} else {
		m.Specbyte = types.StringNull()
	}
	if v, ok := obj["src"]; ok {
		_ = v
		if v != "" {
			m.Src = types.StringValue(v)
		} else {
			m.Src = types.StringNull()
		}
	} else {
		m.Src = types.StringNull()
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
	if v, ok := obj["tcp"]; ok {
		_ = v
		if v != "" {
			m.TCP = types.StringValue(v)
		} else {
			m.TCP = types.StringNull()
		}
	} else {
		m.TCP = types.StringNull()
	}
	if v, ok := obj["tcp-ack"]; ok {
		_ = v
		if v != "" {
			m.TCPAck = types.StringValue(v)
		} else {
			m.TCPAck = types.StringNull()
		}
	} else {
		m.TCPAck = types.StringNull()
	}
	if v, ok := obj["tcp-data-offset"]; ok {
		_ = v
		if v != "" {
			m.TCPDataOffset = types.StringValue(v)
		} else {
			m.TCPDataOffset = types.StringNull()
		}
	} else {
		m.TCPDataOffset = types.StringNull()
	}
	if v, ok := obj["tcp-dst-port"]; ok {
		_ = v
		if v != "" {
			m.TCPDstPort = types.StringValue(v)
		} else {
			m.TCPDstPort = types.StringNull()
		}
	} else {
		m.TCPDstPort = types.StringNull()
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
	if v, ok := obj["tcp-src-port"]; ok {
		_ = v
		if v != "" {
			m.TCPSrcPort = types.StringValue(v)
		} else {
			m.TCPSrcPort = types.StringNull()
		}
	} else {
		m.TCPSrcPort = types.StringNull()
	}
	if v, ok := obj["tcp-syn"]; ok {
		_ = v
		if v != "" {
			m.TCPSyn = types.StringValue(v)
		} else {
			m.TCPSyn = types.StringNull()
		}
	} else {
		m.TCPSyn = types.StringNull()
	}
	if v, ok := obj["tcp-urgent-pointer"]; ok {
		_ = v
		if v != "" {
			m.TCPUrgentPointer = types.StringValue(v)
		} else {
			m.TCPUrgentPointer = types.StringNull()
		}
	} else {
		m.TCPUrgentPointer = types.StringNull()
	}
	if v, ok := obj["tcp-window-size"]; ok {
		_ = v
		if v != "" {
			m.TCPWindowSize = types.StringValue(v)
		} else {
			m.TCPWindowSize = types.StringNull()
		}
	} else {
		m.TCPWindowSize = types.StringNull()
	}
	if v, ok := obj["traffic-class"]; ok {
		_ = v
		if v != "" {
			m.TrafficClass = types.StringValue(v)
		} else {
			m.TrafficClass = types.StringNull()
		}
	} else {
		m.TrafficClass = types.StringNull()
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
	if v, ok := obj["udp"]; ok {
		_ = v
		if v != "" {
			m.UDP = types.StringValue(v)
		} else {
			m.UDP = types.StringNull()
		}
	} else {
		m.UDP = types.StringNull()
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
	if v, ok := obj["vlan-id"]; ok {
		_ = v
		if v != "" {
			m.VLANID = types.StringValue(v)
		} else {
			m.VLANID = types.StringNull()
		}
	} else {
		m.VLANID = types.StringNull()
	}
}
