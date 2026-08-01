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
	_ resource.Resource                = &ToolSnifferStopResource{}
	_ resource.ResourceWithImportState = &ToolSnifferStopResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type ToolSnifferStopResource struct {
	reg *client.Registry
}

type ToolSnifferStopModel struct {
	ID                           types.String `tfsdk:"id"`
	FileLimit                    types.String `tfsdk:"file_limit"`
	FileName                     types.String `tfsdk:"file_name"`
	FilterCpu                    types.String `tfsdk:"filter_cpu"`
	FilterDirection              types.String `tfsdk:"filter_direction"`
	FilterDstIPAddress           types.String `tfsdk:"filter_dst_ip_address"`
	FilterDstIPV6Address         types.String `tfsdk:"filter_dst_ipv6_address"`
	FilterDstMACAddress          types.String `tfsdk:"filter_dst_mac_address"`
	FilterDstPort                types.String `tfsdk:"filter_dst_port"`
	FilterInterface              types.String `tfsdk:"filter_interface"`
	FilterIPAddress              types.String `tfsdk:"filter_ip_address"`
	FilterIPProtocol             types.String `tfsdk:"filter_ip_protocol"`
	FilterIPV6Address            types.String `tfsdk:"filter_ipv6_address"`
	FilterMACAddress             types.String `tfsdk:"filter_mac_address"`
	FilterMACProtocol            types.String `tfsdk:"filter_mac_protocol"`
	FilterOperatorBetweenEntries types.String `tfsdk:"filter_operator_between_entries"`
	FilterPort                   types.String `tfsdk:"filter_port"`
	FilterSize                   types.String `tfsdk:"filter_size"`
	FilterSrcIPAddress           types.String `tfsdk:"filter_src_ip_address"`
	FilterSrcIPV6Address         types.String `tfsdk:"filter_src_ipv6_address"`
	FilterSrcMACAddress          types.String `tfsdk:"filter_src_mac_address"`
	FilterSrcPort                types.String `tfsdk:"filter_src_port"`
	FilterStream                 types.String `tfsdk:"filter_stream"`
	FilterVLAN                   types.String `tfsdk:"filter_vlan"`
	MemoryLimit                  types.String `tfsdk:"memory_limit"`
	MemoryScroll                 types.String `tfsdk:"memory_scroll"`
	OnlyHeaders                  types.String `tfsdk:"only_headers"`
	ShowFrame                    types.String `tfsdk:"show_frame"`
	StreamingEnabled             types.String `tfsdk:"streaming_enabled"`
	StreamingPort                types.String `tfsdk:"streaming_port"`
	StreamingServer              types.String `tfsdk:"streaming_server"`
	Router                       types.String `tfsdk:"router"`
}

func NewToolSnifferStopResource() resource.Resource { return &ToolSnifferStopResource{} }

func (r *ToolSnifferStopResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_sniffer_stop"
}

func (r *ToolSnifferStopResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolSnifferStopResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Action only; not CRUD",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"file_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "File size limit. Sniffer will stop when a limit is reached.",
			},
			"file_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the file where sniffed packets will be saved.",
			},
			"filter_cpu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "CPU core used as a filter.",
			},
			"filter_direction": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies which direction filtering will be applied.",
			},
			"filter_dst_ip_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 IP destination addresses used as a filter.",
			},
			"filter_dst_ipv6_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 IPv6 destination addresses used as a filter.",
			},
			"filter_dst_mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 MAC destination addresses and MAC address masks used as a filter.",
			},
			"filter_dst_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 comma-separated destination ports used as a filter. A list of predefined port names is also available, like ssh and telnet.",
			},
			"filter_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interface name on which sniffer will be running. \u00a0 all \u00a0 indicates that the sniffer will sniff packets on all interfaces.",
			},
			"filter_ip_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 IP addresses used as a filter.",
			},
			"filter_ip_protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 comma-separated IP/IPv6 protocols used as a filter. IP protocols (instead of protocol names, protocol numbers can be used): ipsec-ah \u00a0 - IPsec AH protocol ipsec-esp \u00a0 - IPsec ESP protocol ddp \u00a0 - datagram delivery protocol egp \u00a0 - exterior gateway protocol ggp \u00a0 - gateway-gateway protocol gre \u00a0 - general routing encapsulation hmp \u00a0 - host monitoring protocol idpr-cmtp \u00a0 - idpr control message transport icmp \u00a0 - internet control message protocol icmpv6 \u00a0 - internet control message protocol v6 igmp \u00a0 - internet group management protocol ipencap \u00a0 - ip encapsulated in ip ipip \u00a0 - ip encapsulation encap \u00a0 - ip encapsulation iso-tp4 \u00a0 - iso transport protocol class 4 ospf \u00a0 - open shortest path first pup \u00a0 - parc universal packet protocol pim \u00a0 - protocol independent multicast rspf \u00a0 - radio shortest path first rdp \u00a0 - reliable datagram protocol st \u00a0 - st datagram mode tcp \u00a0 - transmission control protocol udp \u00a0 - user datagram protocol vmtp \u00a0 - versatile message transport vrrp \u00a0 - virtual router redundancy protocol xns-idp \u00a0 - xerox xns idp xtp \u00a0 - xpress transfer protocol",
			},
			"filter_ipv6_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 IPv6 addresses used as a filter.",
			},
			"filter_mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 MAC addresses and MAC address masks used as a filter.",
			},
			"filter_mac_protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 comma separated entries used as a filter. Mac protocols (instead of protocol names, protocol number can be used): 802.2 \u00a0 - 802.2 Frames (0x0004) arp \u00a0 - Address Resolution Protocol (0x0806) capsman - CAPsMAN to CAP MAC layer connection ( 0x88BB ) dot1x - EAPoL IEEE 802.1X ( 0x888E ) homeplug-av \u00a0 - HomePlug AV MME (0x88E1) ip \u00a0 - Internet Protocol version 4 (0x0800) ipv6 \u00a0 - Internet Protocol Version 6 (0x86DD) ipx \u00a0 - Internetwork Packet Exchange (0x8137) lacp - Link Aggregation Control Protocol ( 0x8809 ) lldp \u00a0 - Link Layer Discovery Protocol (0x88CC) loop-protect \u00a0 - Loop Protect Protocol (0x9003) macsec - MAC security IEEE 802.1AE (0x88E5) mpls-multicast \u00a0 - MPLS multicast (0x8848) mpls-unicast \u00a0 - MPLS unicast (0x8847) mvrp - Multiple VLAN Registration protocol (0x88F5) packing-compr \u00a0 - Encapsulated packets with compressed \u00a0 IP packing \u00a0 (0x9001) packing-simple \u00a0 - Encapsulated packets with simple \u00a0 IP packing \u00a0 (0x9000) pppoe \u00a0 - PPPoE Session Stage (0x8864) pppoe-discovery \u00a0 - PPPoE Discovery Stage (0x8863) rarp \u00a0 - Reverse Address Resolution Protocol (0x8035) romon - Router Management Overlay Network RoMON ( 0x88BF ) service-vlan \u00a0 - Provider Bridging (IEEE 802.1ad) & Shortest Path Bridging IEEE 802.1aq (0x88A8) vlan \u00a0 - VLAN-tagged frame (IEEE 802.1Q) and Shortest Path Bridging IEEE 802.1aq with NNI compatibility (0x8100)",
			},
			"filter_operator_between_entries": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Changes the logic for filters with multiple entries.",
			},
			"filter_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 comma-separated ports used as a filter. A list of predefined port names is also available, like ssh and telnet.",
			},
			"filter_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Filters packets of specified size or size range in bytes.",
			},
			"filter_src_ip_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 IP source addresses used as a filter.",
			},
			"filter_src_ipv6_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 IPv6 source addresses used as a filter.",
			},
			"filter_src_mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 MAC source addresses and MAC address masks used as a filter.",
			},
			"filter_src_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 comma-separated source ports used as a filter. A list of predefined port names is also available, like ssh and telnet.",
			},
			"filter_stream": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Sniffed packets that are devised for the sniffer server are ignored.",
			},
			"filter_vlan": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Up to 16 VLAN IDs used as a filter.",
			},
			"memory_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Memory amount used to store sniffed data.",
			},
			"memory_scroll": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether to rewrite older sniffed data when the memory limit is reached.",
			},
			"only_headers": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Save in the memory only the packet's headers, not the whole packet.",
			},
			"show_frame": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether to see the content of the frame when running quick sniffer in command line.",
			},
			"streaming_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Defines whether to send sniffed packets to the streaming server.",
			},
			"streaming_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port to stream the TZSP packets to.",
			},
			"streaming_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Tazmen Sniffer Protocol (TZSP) stream receiver.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *ToolSnifferStopResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolSnifferStopModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.FileLimit.IsNull() || plan.FileLimit.IsUnknown()) {
		body["file-limit"] = plan.FileLimit.ValueString()
	}
	if !(plan.FileName.IsNull() || plan.FileName.IsUnknown()) {
		body["file-name"] = plan.FileName.ValueString()
	}
	if !(plan.FilterCpu.IsNull() || plan.FilterCpu.IsUnknown()) {
		body["filter-cpu"] = plan.FilterCpu.ValueString()
	}
	if !(plan.FilterDirection.IsNull() || plan.FilterDirection.IsUnknown()) {
		body["filter-direction"] = plan.FilterDirection.ValueString()
	}
	if !(plan.FilterDstIPAddress.IsNull() || plan.FilterDstIPAddress.IsUnknown()) {
		body["filter-dst-ip-address"] = plan.FilterDstIPAddress.ValueString()
	}
	if !(plan.FilterDstIPV6Address.IsNull() || plan.FilterDstIPV6Address.IsUnknown()) {
		body["filter-dst-ipv6-address"] = plan.FilterDstIPV6Address.ValueString()
	}
	if !(plan.FilterDstMACAddress.IsNull() || plan.FilterDstMACAddress.IsUnknown()) {
		body["filter-dst-mac-address"] = plan.FilterDstMACAddress.ValueString()
	}
	if !(plan.FilterDstPort.IsNull() || plan.FilterDstPort.IsUnknown()) {
		body["filter-dst-port"] = plan.FilterDstPort.ValueString()
	}
	if !(plan.FilterInterface.IsNull() || plan.FilterInterface.IsUnknown()) {
		body["filter-interface"] = plan.FilterInterface.ValueString()
	}
	if !(plan.FilterIPAddress.IsNull() || plan.FilterIPAddress.IsUnknown()) {
		body["filter-ip-address"] = plan.FilterIPAddress.ValueString()
	}
	if !(plan.FilterIPProtocol.IsNull() || plan.FilterIPProtocol.IsUnknown()) {
		body["filter-ip-protocol"] = plan.FilterIPProtocol.ValueString()
	}
	if !(plan.FilterIPV6Address.IsNull() || plan.FilterIPV6Address.IsUnknown()) {
		body["filter-ipv6-address"] = plan.FilterIPV6Address.ValueString()
	}
	if !(plan.FilterMACAddress.IsNull() || plan.FilterMACAddress.IsUnknown()) {
		body["filter-mac-address"] = plan.FilterMACAddress.ValueString()
	}
	if !(plan.FilterMACProtocol.IsNull() || plan.FilterMACProtocol.IsUnknown()) {
		body["filter-mac-protocol"] = plan.FilterMACProtocol.ValueString()
	}
	if !(plan.FilterOperatorBetweenEntries.IsNull() || plan.FilterOperatorBetweenEntries.IsUnknown()) {
		body["filter-operator-between-entries"] = plan.FilterOperatorBetweenEntries.ValueString()
	}
	if !(plan.FilterPort.IsNull() || plan.FilterPort.IsUnknown()) {
		body["filter-port"] = plan.FilterPort.ValueString()
	}
	if !(plan.FilterSize.IsNull() || plan.FilterSize.IsUnknown()) {
		body["filter-size"] = plan.FilterSize.ValueString()
	}
	if !(plan.FilterSrcIPAddress.IsNull() || plan.FilterSrcIPAddress.IsUnknown()) {
		body["filter-src-ip-address"] = plan.FilterSrcIPAddress.ValueString()
	}
	if !(plan.FilterSrcIPV6Address.IsNull() || plan.FilterSrcIPV6Address.IsUnknown()) {
		body["filter-src-ipv6-address"] = plan.FilterSrcIPV6Address.ValueString()
	}
	if !(plan.FilterSrcMACAddress.IsNull() || plan.FilterSrcMACAddress.IsUnknown()) {
		body["filter-src-mac-address"] = plan.FilterSrcMACAddress.ValueString()
	}
	if !(plan.FilterSrcPort.IsNull() || plan.FilterSrcPort.IsUnknown()) {
		body["filter-src-port"] = plan.FilterSrcPort.ValueString()
	}
	if !(plan.FilterStream.IsNull() || plan.FilterStream.IsUnknown()) {
		body["filter-stream"] = plan.FilterStream.ValueString()
	}
	if !(plan.FilterVLAN.IsNull() || plan.FilterVLAN.IsUnknown()) {
		body["filter-vlan"] = plan.FilterVLAN.ValueString()
	}
	if !(plan.MemoryLimit.IsNull() || plan.MemoryLimit.IsUnknown()) {
		body["memory-limit"] = plan.MemoryLimit.ValueString()
	}
	if !(plan.MemoryScroll.IsNull() || plan.MemoryScroll.IsUnknown()) {
		body["memory-scroll"] = plan.MemoryScroll.ValueString()
	}
	if !(plan.OnlyHeaders.IsNull() || plan.OnlyHeaders.IsUnknown()) {
		body["only-headers"] = plan.OnlyHeaders.ValueString()
	}
	if !(plan.ShowFrame.IsNull() || plan.ShowFrame.IsUnknown()) {
		body["show-frame"] = plan.ShowFrame.ValueString()
	}
	if !(plan.StreamingEnabled.IsNull() || plan.StreamingEnabled.IsUnknown()) {
		body["streaming-enabled"] = plan.StreamingEnabled.ValueString()
	}
	if !(plan.StreamingPort.IsNull() || plan.StreamingPort.IsUnknown()) {
		body["streaming-port"] = plan.StreamingPort.ValueString()
	}
	if !(plan.StreamingServer.IsNull() || plan.StreamingServer.IsUnknown()) {
		body["streaming-server"] = plan.StreamingServer.ValueString()
	}
	obj, err := c.Add(ctx, "/tool/sniffer/stop", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /tool/sniffer/stop failed", err.Error())
		return
	}
	toolSnifferStopApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolSnifferStopResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolSnifferStopModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/tool/sniffer/stop", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /tool/sniffer/stop failed", err.Error())
		return
	}
	toolSnifferStopApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolSnifferStopResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ToolSnifferStopModel
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
	if !plan.FileLimit.Equal(state.FileLimit) && !plan.FileLimit.IsUnknown() {
		body["file-limit"] = plan.FileLimit.ValueString()
	}
	if !plan.FileName.Equal(state.FileName) && !plan.FileName.IsUnknown() {
		body["file-name"] = plan.FileName.ValueString()
	}
	if !plan.FilterCpu.Equal(state.FilterCpu) && !plan.FilterCpu.IsUnknown() {
		body["filter-cpu"] = plan.FilterCpu.ValueString()
	}
	if !plan.FilterDirection.Equal(state.FilterDirection) && !plan.FilterDirection.IsUnknown() {
		body["filter-direction"] = plan.FilterDirection.ValueString()
	}
	if !plan.FilterDstIPAddress.Equal(state.FilterDstIPAddress) && !plan.FilterDstIPAddress.IsUnknown() {
		body["filter-dst-ip-address"] = plan.FilterDstIPAddress.ValueString()
	}
	if !plan.FilterDstIPV6Address.Equal(state.FilterDstIPV6Address) && !plan.FilterDstIPV6Address.IsUnknown() {
		body["filter-dst-ipv6-address"] = plan.FilterDstIPV6Address.ValueString()
	}
	if !plan.FilterDstMACAddress.Equal(state.FilterDstMACAddress) && !plan.FilterDstMACAddress.IsUnknown() {
		body["filter-dst-mac-address"] = plan.FilterDstMACAddress.ValueString()
	}
	if !plan.FilterDstPort.Equal(state.FilterDstPort) && !plan.FilterDstPort.IsUnknown() {
		body["filter-dst-port"] = plan.FilterDstPort.ValueString()
	}
	if !plan.FilterInterface.Equal(state.FilterInterface) && !plan.FilterInterface.IsUnknown() {
		body["filter-interface"] = plan.FilterInterface.ValueString()
	}
	if !plan.FilterIPAddress.Equal(state.FilterIPAddress) && !plan.FilterIPAddress.IsUnknown() {
		body["filter-ip-address"] = plan.FilterIPAddress.ValueString()
	}
	if !plan.FilterIPProtocol.Equal(state.FilterIPProtocol) && !plan.FilterIPProtocol.IsUnknown() {
		body["filter-ip-protocol"] = plan.FilterIPProtocol.ValueString()
	}
	if !plan.FilterIPV6Address.Equal(state.FilterIPV6Address) && !plan.FilterIPV6Address.IsUnknown() {
		body["filter-ipv6-address"] = plan.FilterIPV6Address.ValueString()
	}
	if !plan.FilterMACAddress.Equal(state.FilterMACAddress) && !plan.FilterMACAddress.IsUnknown() {
		body["filter-mac-address"] = plan.FilterMACAddress.ValueString()
	}
	if !plan.FilterMACProtocol.Equal(state.FilterMACProtocol) && !plan.FilterMACProtocol.IsUnknown() {
		body["filter-mac-protocol"] = plan.FilterMACProtocol.ValueString()
	}
	if !plan.FilterOperatorBetweenEntries.Equal(state.FilterOperatorBetweenEntries) && !plan.FilterOperatorBetweenEntries.IsUnknown() {
		body["filter-operator-between-entries"] = plan.FilterOperatorBetweenEntries.ValueString()
	}
	if !plan.FilterPort.Equal(state.FilterPort) && !plan.FilterPort.IsUnknown() {
		body["filter-port"] = plan.FilterPort.ValueString()
	}
	if !plan.FilterSize.Equal(state.FilterSize) && !plan.FilterSize.IsUnknown() {
		body["filter-size"] = plan.FilterSize.ValueString()
	}
	if !plan.FilterSrcIPAddress.Equal(state.FilterSrcIPAddress) && !plan.FilterSrcIPAddress.IsUnknown() {
		body["filter-src-ip-address"] = plan.FilterSrcIPAddress.ValueString()
	}
	if !plan.FilterSrcIPV6Address.Equal(state.FilterSrcIPV6Address) && !plan.FilterSrcIPV6Address.IsUnknown() {
		body["filter-src-ipv6-address"] = plan.FilterSrcIPV6Address.ValueString()
	}
	if !plan.FilterSrcMACAddress.Equal(state.FilterSrcMACAddress) && !plan.FilterSrcMACAddress.IsUnknown() {
		body["filter-src-mac-address"] = plan.FilterSrcMACAddress.ValueString()
	}
	if !plan.FilterSrcPort.Equal(state.FilterSrcPort) && !plan.FilterSrcPort.IsUnknown() {
		body["filter-src-port"] = plan.FilterSrcPort.ValueString()
	}
	if !plan.FilterStream.Equal(state.FilterStream) && !plan.FilterStream.IsUnknown() {
		body["filter-stream"] = plan.FilterStream.ValueString()
	}
	if !plan.FilterVLAN.Equal(state.FilterVLAN) && !plan.FilterVLAN.IsUnknown() {
		body["filter-vlan"] = plan.FilterVLAN.ValueString()
	}
	if !plan.MemoryLimit.Equal(state.MemoryLimit) && !plan.MemoryLimit.IsUnknown() {
		body["memory-limit"] = plan.MemoryLimit.ValueString()
	}
	if !plan.MemoryScroll.Equal(state.MemoryScroll) && !plan.MemoryScroll.IsUnknown() {
		body["memory-scroll"] = plan.MemoryScroll.ValueString()
	}
	if !plan.OnlyHeaders.Equal(state.OnlyHeaders) && !plan.OnlyHeaders.IsUnknown() {
		body["only-headers"] = plan.OnlyHeaders.ValueString()
	}
	if !plan.ShowFrame.Equal(state.ShowFrame) && !plan.ShowFrame.IsUnknown() {
		body["show-frame"] = plan.ShowFrame.ValueString()
	}
	if !plan.StreamingEnabled.Equal(state.StreamingEnabled) && !plan.StreamingEnabled.IsUnknown() {
		body["streaming-enabled"] = plan.StreamingEnabled.ValueString()
	}
	if !plan.StreamingPort.Equal(state.StreamingPort) && !plan.StreamingPort.IsUnknown() {
		body["streaming-port"] = plan.StreamingPort.ValueString()
	}
	if !plan.StreamingServer.Equal(state.StreamingServer) && !plan.StreamingServer.IsUnknown() {
		body["streaming-server"] = plan.StreamingServer.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/tool/sniffer/stop", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /tool/sniffer/stop failed", err.Error())
			return
		}
		toolSnifferStopApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolSnifferStopResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ToolSnifferStopModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/tool/sniffer/stop", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /tool/sniffer/stop failed", err.Error())
	}
}

func (r *ToolSnifferStopResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := toolSnifferStopLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /tool/sniffer/stop matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// toolSnifferStopLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func toolSnifferStopLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/tool/sniffer/stop", id)
}

func toolSnifferStopApply(ctx context.Context, obj client.Object, m *ToolSnifferStopModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["file-limit"]; ok {
		_ = v
		if v != "" {
			m.FileLimit = types.StringValue(v)
		} else {
			m.FileLimit = types.StringNull()
		}
	} else {
		m.FileLimit = types.StringNull()
	}
	if v, ok := obj["file-name"]; ok {
		_ = v
		if v != "" {
			m.FileName = types.StringValue(v)
		} else {
			m.FileName = types.StringNull()
		}
	} else {
		m.FileName = types.StringNull()
	}
	if v, ok := obj["filter-cpu"]; ok {
		_ = v
		if v != "" {
			m.FilterCpu = types.StringValue(v)
		} else {
			m.FilterCpu = types.StringNull()
		}
	} else {
		m.FilterCpu = types.StringNull()
	}
	if v, ok := obj["filter-direction"]; ok {
		_ = v
		if v != "" {
			m.FilterDirection = types.StringValue(v)
		} else {
			m.FilterDirection = types.StringNull()
		}
	} else {
		m.FilterDirection = types.StringNull()
	}
	if v, ok := obj["filter-dst-ip-address"]; ok {
		_ = v
		if v != "" {
			m.FilterDstIPAddress = types.StringValue(v)
		} else {
			m.FilterDstIPAddress = types.StringNull()
		}
	} else {
		m.FilterDstIPAddress = types.StringNull()
	}
	if v, ok := obj["filter-dst-ipv6-address"]; ok {
		_ = v
		if v != "" {
			m.FilterDstIPV6Address = types.StringValue(v)
		} else {
			m.FilterDstIPV6Address = types.StringNull()
		}
	} else {
		m.FilterDstIPV6Address = types.StringNull()
	}
	if v, ok := obj["filter-dst-mac-address"]; ok {
		_ = v
		if v != "" {
			m.FilterDstMACAddress = types.StringValue(v)
		} else {
			m.FilterDstMACAddress = types.StringNull()
		}
	} else {
		m.FilterDstMACAddress = types.StringNull()
	}
	if v, ok := obj["filter-dst-port"]; ok {
		_ = v
		if v != "" {
			m.FilterDstPort = types.StringValue(v)
		} else {
			m.FilterDstPort = types.StringNull()
		}
	} else {
		m.FilterDstPort = types.StringNull()
	}
	if v, ok := obj["filter-interface"]; ok {
		_ = v
		if v != "" {
			m.FilterInterface = types.StringValue(v)
		} else {
			m.FilterInterface = types.StringNull()
		}
	} else {
		m.FilterInterface = types.StringNull()
	}
	if v, ok := obj["filter-ip-address"]; ok {
		_ = v
		if v != "" {
			m.FilterIPAddress = types.StringValue(v)
		} else {
			m.FilterIPAddress = types.StringNull()
		}
	} else {
		m.FilterIPAddress = types.StringNull()
	}
	if v, ok := obj["filter-ip-protocol"]; ok {
		_ = v
		if v != "" {
			m.FilterIPProtocol = types.StringValue(v)
		} else {
			m.FilterIPProtocol = types.StringNull()
		}
	} else {
		m.FilterIPProtocol = types.StringNull()
	}
	if v, ok := obj["filter-ipv6-address"]; ok {
		_ = v
		if v != "" {
			m.FilterIPV6Address = types.StringValue(v)
		} else {
			m.FilterIPV6Address = types.StringNull()
		}
	} else {
		m.FilterIPV6Address = types.StringNull()
	}
	if v, ok := obj["filter-mac-address"]; ok {
		_ = v
		if v != "" {
			m.FilterMACAddress = types.StringValue(v)
		} else {
			m.FilterMACAddress = types.StringNull()
		}
	} else {
		m.FilterMACAddress = types.StringNull()
	}
	if v, ok := obj["filter-mac-protocol"]; ok {
		_ = v
		if v != "" {
			m.FilterMACProtocol = types.StringValue(v)
		} else {
			m.FilterMACProtocol = types.StringNull()
		}
	} else {
		m.FilterMACProtocol = types.StringNull()
	}
	if v, ok := obj["filter-operator-between-entries"]; ok {
		_ = v
		if v != "" {
			m.FilterOperatorBetweenEntries = types.StringValue(v)
		} else {
			m.FilterOperatorBetweenEntries = types.StringNull()
		}
	} else {
		m.FilterOperatorBetweenEntries = types.StringNull()
	}
	if v, ok := obj["filter-port"]; ok {
		_ = v
		if v != "" {
			m.FilterPort = types.StringValue(v)
		} else {
			m.FilterPort = types.StringNull()
		}
	} else {
		m.FilterPort = types.StringNull()
	}
	if v, ok := obj["filter-size"]; ok {
		_ = v
		if v != "" {
			m.FilterSize = types.StringValue(v)
		} else {
			m.FilterSize = types.StringNull()
		}
	} else {
		m.FilterSize = types.StringNull()
	}
	if v, ok := obj["filter-src-ip-address"]; ok {
		_ = v
		if v != "" {
			m.FilterSrcIPAddress = types.StringValue(v)
		} else {
			m.FilterSrcIPAddress = types.StringNull()
		}
	} else {
		m.FilterSrcIPAddress = types.StringNull()
	}
	if v, ok := obj["filter-src-ipv6-address"]; ok {
		_ = v
		if v != "" {
			m.FilterSrcIPV6Address = types.StringValue(v)
		} else {
			m.FilterSrcIPV6Address = types.StringNull()
		}
	} else {
		m.FilterSrcIPV6Address = types.StringNull()
	}
	if v, ok := obj["filter-src-mac-address"]; ok {
		_ = v
		if v != "" {
			m.FilterSrcMACAddress = types.StringValue(v)
		} else {
			m.FilterSrcMACAddress = types.StringNull()
		}
	} else {
		m.FilterSrcMACAddress = types.StringNull()
	}
	if v, ok := obj["filter-src-port"]; ok {
		_ = v
		if v != "" {
			m.FilterSrcPort = types.StringValue(v)
		} else {
			m.FilterSrcPort = types.StringNull()
		}
	} else {
		m.FilterSrcPort = types.StringNull()
	}
	if v, ok := obj["filter-stream"]; ok {
		_ = v
		if v != "" {
			m.FilterStream = types.StringValue(v)
		} else {
			m.FilterStream = types.StringNull()
		}
	} else {
		m.FilterStream = types.StringNull()
	}
	if v, ok := obj["filter-vlan"]; ok {
		_ = v
		if v != "" {
			m.FilterVLAN = types.StringValue(v)
		} else {
			m.FilterVLAN = types.StringNull()
		}
	} else {
		m.FilterVLAN = types.StringNull()
	}
	if v, ok := obj["memory-limit"]; ok {
		_ = v
		if v != "" {
			m.MemoryLimit = types.StringValue(v)
		} else {
			m.MemoryLimit = types.StringNull()
		}
	} else {
		m.MemoryLimit = types.StringNull()
	}
	if v, ok := obj["memory-scroll"]; ok {
		_ = v
		if v != "" {
			m.MemoryScroll = types.StringValue(v)
		} else {
			m.MemoryScroll = types.StringNull()
		}
	} else {
		m.MemoryScroll = types.StringNull()
	}
	if v, ok := obj["only-headers"]; ok {
		_ = v
		if v != "" {
			m.OnlyHeaders = types.StringValue(v)
		} else {
			m.OnlyHeaders = types.StringNull()
		}
	} else {
		m.OnlyHeaders = types.StringNull()
	}
	if v, ok := obj["show-frame"]; ok {
		_ = v
		if v != "" {
			m.ShowFrame = types.StringValue(v)
		} else {
			m.ShowFrame = types.StringNull()
		}
	} else {
		m.ShowFrame = types.StringNull()
	}
	if v, ok := obj["streaming-enabled"]; ok {
		_ = v
		if v != "" {
			m.StreamingEnabled = types.StringValue(v)
		} else {
			m.StreamingEnabled = types.StringNull()
		}
	} else {
		m.StreamingEnabled = types.StringNull()
	}
	if v, ok := obj["streaming-port"]; ok {
		_ = v
		if v != "" {
			m.StreamingPort = types.StringValue(v)
		} else {
			m.StreamingPort = types.StringNull()
		}
	} else {
		m.StreamingPort = types.StringNull()
	}
	if v, ok := obj["streaming-server"]; ok {
		_ = v
		if v != "" {
			m.StreamingServer = types.StringValue(v)
		} else {
			m.StreamingServer = types.StringNull()
		}
	} else {
		m.StreamingServer = types.StringNull()
	}
}
