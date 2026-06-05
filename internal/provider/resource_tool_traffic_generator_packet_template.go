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
	ID                      types.String `tfsdk:"id"`
	AssumedDscpEcn          types.String `tfsdk:"assumed_dscp_ecn"`
	AssumedDst              types.String `tfsdk:"assumed_dst"`
	AssumedDstPort          types.String `tfsdk:"assumed_dst_port"`
	AssumedFlowLabel        types.String `tfsdk:"assumed_flow_label"`
	AssumedFragOffset       types.String `tfsdk:"assumed_frag_offset"`
	AssumedHeader           types.String `tfsdk:"assumed_header"`
	AssumedInterface        types.String `tfsdk:"assumed_interface"`
	AssumedIPID             types.String `tfsdk:"assumed_ip_id"`
	AssumedNextHeader       types.String `tfsdk:"assumed_next_header"`
	AssumedPort             types.String `tfsdk:"assumed_port"`
	AssumedPriority         types.String `tfsdk:"assumed_priority"`
	AssumedProtocol         types.String `tfsdk:"assumed_protocol"`
	AssumedSrc              types.String `tfsdk:"assumed_src"`
	AssumedSrcPort          types.String `tfsdk:"assumed_src_port"`
	AssumedTCPAck           types.String `tfsdk:"assumed_tcp_ack"`
	AssumedTCPDataOffset    types.String `tfsdk:"assumed_tcp_data_offset"`
	AssumedTCPDstPort       types.String `tfsdk:"assumed_tcp_dst_port"`
	AssumedTCPFlags         types.String `tfsdk:"assumed_tcp_flags"`
	AssumedTCPSrcPort       types.String `tfsdk:"assumed_tcp_src_port"`
	AssumedTCPSyn           types.String `tfsdk:"assumed_tcp_syn"`
	AssumedTCPUrgentPointer types.String `tfsdk:"assumed_tcp_urgent_pointer"`
	AssumedTCPWindowSize    types.String `tfsdk:"assumed_tcp_window_size"`
	AssumedTrafficClass     types.String `tfsdk:"assumed_traffic_class"`
	AssumedTtl              types.String `tfsdk:"assumed_ttl"`
	AssumedVLANID           types.String `tfsdk:"assumed_vlan_id"`
	Comment                 types.String `tfsdk:"comment"`
	Data                    types.String `tfsdk:"data"`
	DataByte                types.Int64  `tfsdk:"data_byte"`
	DscpEcn                 types.String `tfsdk:"dscp_ecn"`
	Dst                     types.String `tfsdk:"dst"`
	DstPort                 types.String `tfsdk:"dst_port"`
	FlowLabel               types.String `tfsdk:"flow_label"`
	FragOffset              types.String `tfsdk:"frag_offset"`
	Gateway                 types.String `tfsdk:"gateway"`
	Header                  types.String `tfsdk:"header"`
	HeaderStack             types.String `tfsdk:"header_stack"`
	HopLimit                types.String `tfsdk:"hop_limit"`
	Interface               types.String `tfsdk:"interface"`
	IP                      types.String `tfsdk:"ip"`
	IPID                    types.String `tfsdk:"ip_id"`
	IPV6                    types.String `tfsdk:"ipv6"`
	MAC                     types.String `tfsdk:"mac"`
	Name                    types.String `tfsdk:"name"`
	NextHeader              types.String `tfsdk:"next_header"`
	Port                    types.String `tfsdk:"port"`
	Priority                types.String `tfsdk:"priority"`
	Protocol                types.String `tfsdk:"protocol"`
	Raw                     types.String `tfsdk:"raw"`
	RawPacketTemplates      types.String `tfsdk:"raw_packet_templates"`
	Specbyte                types.String `tfsdk:"specbyte"`
	Src                     types.String `tfsdk:"src"`
	SrcPort                 types.String `tfsdk:"src_port"`
	TCP                     types.String `tfsdk:"tcp"`
	TCPAck                  types.String `tfsdk:"tcp_ack"`
	TCPDataOffset           types.String `tfsdk:"tcp_data_offset"`
	TCPDstPort              types.String `tfsdk:"tcp_dst_port"`
	TCPFlags                types.String `tfsdk:"tcp_flags"`
	TCPSrcPort              types.String `tfsdk:"tcp_src_port"`
	TCPSyn                  types.String `tfsdk:"tcp_syn"`
	TCPUrgentPointer        types.String `tfsdk:"tcp_urgent_pointer"`
	TCPWindowSize           types.String `tfsdk:"tcp_window_size"`
	TrafficClass            types.String `tfsdk:"traffic_class"`
	Ttl                     types.String `tfsdk:"ttl"`
	UDP                     types.String `tfsdk:"udp"`
	VLAN                    types.String `tfsdk:"vlan"`
	VLANID                  types.String `tfsdk:"vlan_id"`
	Router                  types.String `tfsdk:"router"`
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
	_ = fmt.Sprintf
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
			"assumed_dscp_ecn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_dst": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_dst_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_flow_label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_frag_offset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_header": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_ip_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_next_header": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_src_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_ack": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_data_offset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_dst_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_flags": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_src_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_syn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_urgent_pointer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_tcp_window_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_traffic_class": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_ttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"assumed_vlan_id": schema.StringAttribute{
				Optional:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dst": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dst_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"flow_label": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"frag_offset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"header": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"header_stack": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hop_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ip_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ipv6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"next_header": schema.StringAttribute{
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
			"raw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"raw_packet_templates": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"specbyte": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tcp": schema.StringAttribute{
				Optional:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"udp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vlan": schema.StringAttribute{
				Optional:    true,
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
	if !(plan.DscpEcn.IsNull() || plan.DscpEcn.IsUnknown()) {
		body["dscp-ecn"] = plan.DscpEcn.ValueString()
	}
	if !(plan.Dst.IsNull() || plan.Dst.IsUnknown()) {
		body["dst"] = plan.Dst.ValueString()
	}
	if !(plan.DstPort.IsNull() || plan.DstPort.IsUnknown()) {
		body["dst-port"] = plan.DstPort.ValueString()
	}
	if !(plan.FlowLabel.IsNull() || plan.FlowLabel.IsUnknown()) {
		body["flow-label"] = plan.FlowLabel.ValueString()
	}
	if !(plan.FragOffset.IsNull() || plan.FragOffset.IsUnknown()) {
		body["frag-offset"] = plan.FragOffset.ValueString()
	}
	if !(plan.Gateway.IsNull() || plan.Gateway.IsUnknown()) {
		body["gateway"] = plan.Gateway.ValueString()
	}
	if !(plan.Header.IsNull() || plan.Header.IsUnknown()) {
		body["header"] = plan.Header.ValueString()
	}
	if !(plan.HeaderStack.IsNull() || plan.HeaderStack.IsUnknown()) {
		body["header-stack"] = plan.HeaderStack.ValueString()
	}
	if !(plan.HopLimit.IsNull() || plan.HopLimit.IsUnknown()) {
		body["hop-limit"] = plan.HopLimit.ValueString()
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.IP.IsNull() || plan.IP.IsUnknown()) {
		body["ip"] = plan.IP.ValueString()
	}
	if !(plan.IPID.IsNull() || plan.IPID.IsUnknown()) {
		body["ip-id"] = plan.IPID.ValueString()
	}
	if !(plan.IPV6.IsNull() || plan.IPV6.IsUnknown()) {
		body["ipv6"] = plan.IPV6.ValueString()
	}
	if !(plan.MAC.IsNull() || plan.MAC.IsUnknown()) {
		body["mac"] = plan.MAC.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.NextHeader.IsNull() || plan.NextHeader.IsUnknown()) {
		body["next-header"] = plan.NextHeader.ValueString()
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
	if !(plan.Raw.IsNull() || plan.Raw.IsUnknown()) {
		body["raw"] = plan.Raw.ValueString()
	}
	if !(plan.RawPacketTemplates.IsNull() || plan.RawPacketTemplates.IsUnknown()) {
		body["raw-packet-templates"] = plan.RawPacketTemplates.ValueString()
	}
	if !(plan.Specbyte.IsNull() || plan.Specbyte.IsUnknown()) {
		body["specbyte"] = plan.Specbyte.ValueString()
	}
	if !(plan.Src.IsNull() || plan.Src.IsUnknown()) {
		body["src"] = plan.Src.ValueString()
	}
	if !(plan.SrcPort.IsNull() || plan.SrcPort.IsUnknown()) {
		body["src-port"] = plan.SrcPort.ValueString()
	}
	if !(plan.TCP.IsNull() || plan.TCP.IsUnknown()) {
		body["tcp"] = plan.TCP.ValueString()
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
	if !(plan.TrafficClass.IsNull() || plan.TrafficClass.IsUnknown()) {
		body["traffic-class"] = plan.TrafficClass.ValueString()
	}
	if !(plan.Ttl.IsNull() || plan.Ttl.IsUnknown()) {
		body["ttl"] = plan.Ttl.ValueString()
	}
	if !(plan.UDP.IsNull() || plan.UDP.IsUnknown()) {
		body["udp"] = plan.UDP.ValueString()
	}
	if !(plan.VLAN.IsNull() || plan.VLAN.IsUnknown()) {
		body["vlan"] = plan.VLAN.ValueString()
	}
	if !(plan.VLANID.IsNull() || plan.VLANID.IsUnknown()) {
		body["vlan-id"] = plan.VLANID.ValueString()
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
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Data.Equal(state.Data) {
		body["data"] = plan.Data.ValueString()
	}
	if !plan.DataByte.Equal(state.DataByte) {
		body["data-byte"] = client.FormatInt64(plan.DataByte.ValueInt64())
	}
	if !plan.DscpEcn.Equal(state.DscpEcn) {
		body["dscp-ecn"] = plan.DscpEcn.ValueString()
	}
	if !plan.Dst.Equal(state.Dst) {
		body["dst"] = plan.Dst.ValueString()
	}
	if !plan.DstPort.Equal(state.DstPort) {
		body["dst-port"] = plan.DstPort.ValueString()
	}
	if !plan.FlowLabel.Equal(state.FlowLabel) {
		body["flow-label"] = plan.FlowLabel.ValueString()
	}
	if !plan.FragOffset.Equal(state.FragOffset) {
		body["frag-offset"] = plan.FragOffset.ValueString()
	}
	if !plan.Gateway.Equal(state.Gateway) {
		body["gateway"] = plan.Gateway.ValueString()
	}
	if !plan.Header.Equal(state.Header) {
		body["header"] = plan.Header.ValueString()
	}
	if !plan.HeaderStack.Equal(state.HeaderStack) {
		body["header-stack"] = plan.HeaderStack.ValueString()
	}
	if !plan.HopLimit.Equal(state.HopLimit) {
		body["hop-limit"] = plan.HopLimit.ValueString()
	}
	if !plan.Interface.Equal(state.Interface) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.IP.Equal(state.IP) {
		body["ip"] = plan.IP.ValueString()
	}
	if !plan.IPID.Equal(state.IPID) {
		body["ip-id"] = plan.IPID.ValueString()
	}
	if !plan.IPV6.Equal(state.IPV6) {
		body["ipv6"] = plan.IPV6.ValueString()
	}
	if !plan.MAC.Equal(state.MAC) {
		body["mac"] = plan.MAC.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.NextHeader.Equal(state.NextHeader) {
		body["next-header"] = plan.NextHeader.ValueString()
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
	if !plan.Raw.Equal(state.Raw) {
		body["raw"] = plan.Raw.ValueString()
	}
	if !plan.RawPacketTemplates.Equal(state.RawPacketTemplates) {
		body["raw-packet-templates"] = plan.RawPacketTemplates.ValueString()
	}
	if !plan.Specbyte.Equal(state.Specbyte) {
		body["specbyte"] = plan.Specbyte.ValueString()
	}
	if !plan.Src.Equal(state.Src) {
		body["src"] = plan.Src.ValueString()
	}
	if !plan.SrcPort.Equal(state.SrcPort) {
		body["src-port"] = plan.SrcPort.ValueString()
	}
	if !plan.TCP.Equal(state.TCP) {
		body["tcp"] = plan.TCP.ValueString()
	}
	if !plan.TCPAck.Equal(state.TCPAck) {
		body["tcp-ack"] = plan.TCPAck.ValueString()
	}
	if !plan.TCPDataOffset.Equal(state.TCPDataOffset) {
		body["tcp-data-offset"] = plan.TCPDataOffset.ValueString()
	}
	if !plan.TCPDstPort.Equal(state.TCPDstPort) {
		body["tcp-dst-port"] = plan.TCPDstPort.ValueString()
	}
	if !plan.TCPFlags.Equal(state.TCPFlags) {
		body["tcp-flags"] = plan.TCPFlags.ValueString()
	}
	if !plan.TCPSrcPort.Equal(state.TCPSrcPort) {
		body["tcp-src-port"] = plan.TCPSrcPort.ValueString()
	}
	if !plan.TCPSyn.Equal(state.TCPSyn) {
		body["tcp-syn"] = plan.TCPSyn.ValueString()
	}
	if !plan.TCPUrgentPointer.Equal(state.TCPUrgentPointer) {
		body["tcp-urgent-pointer"] = plan.TCPUrgentPointer.ValueString()
	}
	if !plan.TCPWindowSize.Equal(state.TCPWindowSize) {
		body["tcp-window-size"] = plan.TCPWindowSize.ValueString()
	}
	if !plan.TrafficClass.Equal(state.TrafficClass) {
		body["traffic-class"] = plan.TrafficClass.ValueString()
	}
	if !plan.Ttl.Equal(state.Ttl) {
		body["ttl"] = plan.Ttl.ValueString()
	}
	if !plan.UDP.Equal(state.UDP) {
		body["udp"] = plan.UDP.ValueString()
	}
	if !plan.VLAN.Equal(state.VLAN) {
		body["vlan"] = plan.VLAN.ValueString()
	}
	if !plan.VLANID.Equal(state.VLANID) {
		body["vlan-id"] = plan.VLANID.ValueString()
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
	id := req.ID
	routerName := ""
	if i := strings.Index(id, "/"); i > 0 && !strings.HasPrefix(id, "*") {
		routerName, id = id[:i], id[i+1:]
	}
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
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/tool/traffic-generator/packet-template", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func toolTrafficGeneratorPacketTemplateApply(ctx context.Context, obj client.Object, m *ToolTrafficGeneratorPacketTemplateModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
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
