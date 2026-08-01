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
	_ resource.Resource                = &ToolTrafficGeneratorRawPacketTemplateResource{}
	_ resource.ResourceWithImportState = &ToolTrafficGeneratorRawPacketTemplateResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type ToolTrafficGeneratorRawPacketTemplateResource struct {
	reg *client.Registry
}

type ToolTrafficGeneratorRawPacketTemplateModel struct {
	ID                        types.String `tfsdk:"id"`
	UdpComputeChecksum        types.String `tfsdk:"udp_compute_checksum"`
	TcpHeaderOffset           types.String `tfsdk:"tcp_header_offset"`
	ComputeChecksumFromOffset types.String `tfsdk:"compute_checksum_from_offset"`
	Comment                   types.String `tfsdk:"comment"`
	Data                      types.String `tfsdk:"data"`
	DataByte                  types.Int64  `tfsdk:"data_byte"`
	Dynamic                   types.Bool   `tfsdk:"dynamic"`
	Header                    types.String `tfsdk:"header"`
	HeaderLength              types.Int64  `tfsdk:"header_length"`
	IPHeaderOffset            types.String `tfsdk:"ip_header_offset"`
	IPV6HeaderOffset          types.String `tfsdk:"ipv6_header_offset"`
	Name                      types.String `tfsdk:"name"`
	Port                      types.String `tfsdk:"port"`
	Random                    types.String `tfsdk:"random"`
	RandomByteOffsetsAndMasks types.String `tfsdk:"random_byte_offsets_and_masks"`
	RandomRanges              types.String `tfsdk:"random_ranges"`
	Specbyte                  types.String `tfsdk:"specbyte"`
	SpecialFooter             types.Bool   `tfsdk:"special_footer"`
	UDPHeaderOffset           types.String `tfsdk:"udp_header_offset"`
	Router                    types.String `tfsdk:"router"`
}

func NewToolTrafficGeneratorRawPacketTemplateResource() resource.Resource {
	return &ToolTrafficGeneratorRawPacketTemplateResource{}
}

func (r *ToolTrafficGeneratorRawPacketTemplateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_traffic_generator_raw_packet_template"
}

func (r *ToolTrafficGeneratorRawPacketTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolTrafficGeneratorRawPacketTemplateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/tool/traffic-generator/raw-packet-template`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"udp_compute_checksum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `udp-compute-checksum`.",
			},
			"tcp_header_offset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tcp-header-offset`.",
			},
			"compute_checksum_from_offset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `compute-checksum-from-offset`.",
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
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"header": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"header_length": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"ip_header_offset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ipv6_header_offset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"random": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"random_byte_offsets_and_masks": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"random_ranges": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"specbyte": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"special_footer": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"udp_header_offset": schema.StringAttribute{
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

func (r *ToolTrafficGeneratorRawPacketTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolTrafficGeneratorRawPacketTemplateModel
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
	if !(plan.Header.IsNull() || plan.Header.IsUnknown()) {
		body["header"] = plan.Header.ValueString()
	}
	if !(plan.IPHeaderOffset.IsNull() || plan.IPHeaderOffset.IsUnknown()) {
		body["ip-header-offset"] = plan.IPHeaderOffset.ValueString()
	}
	if !(plan.IPV6HeaderOffset.IsNull() || plan.IPV6HeaderOffset.IsUnknown()) {
		body["ipv6-header-offset"] = plan.IPV6HeaderOffset.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.RandomByteOffsetsAndMasks.IsNull() || plan.RandomByteOffsetsAndMasks.IsUnknown()) {
		body["random-byte-offsets-and-masks"] = plan.RandomByteOffsetsAndMasks.ValueString()
	}
	if !(plan.RandomRanges.IsNull() || plan.RandomRanges.IsUnknown()) {
		body["random-ranges"] = plan.RandomRanges.ValueString()
	}
	if !(plan.SpecialFooter.IsNull() || plan.SpecialFooter.IsUnknown()) {
		body["special-footer"] = client.FormatBool(plan.SpecialFooter.ValueBool())
	}
	if !(plan.UDPHeaderOffset.IsNull() || plan.UDPHeaderOffset.IsUnknown()) {
		body["udp-header-offset"] = plan.UDPHeaderOffset.ValueString()
	}
	if !(plan.ComputeChecksumFromOffset.IsNull() || plan.ComputeChecksumFromOffset.IsUnknown()) {
		body["compute-checksum-from-offset"] = plan.ComputeChecksumFromOffset.ValueString()
	}
	if !(plan.TcpHeaderOffset.IsNull() || plan.TcpHeaderOffset.IsUnknown()) {
		body["tcp-header-offset"] = plan.TcpHeaderOffset.ValueString()
	}
	if !(plan.UdpComputeChecksum.IsNull() || plan.UdpComputeChecksum.IsUnknown()) {
		body["udp-compute-checksum"] = plan.UdpComputeChecksum.ValueString()
	}
	obj, err := c.Add(ctx, "/tool/traffic-generator/raw-packet-template", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /tool/traffic-generator/raw-packet-template failed", err.Error())
		return
	}
	toolTrafficGeneratorRawPacketTemplateApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolTrafficGeneratorRawPacketTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolTrafficGeneratorRawPacketTemplateModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/tool/traffic-generator/raw-packet-template", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /tool/traffic-generator/raw-packet-template failed", err.Error())
		return
	}
	toolTrafficGeneratorRawPacketTemplateApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolTrafficGeneratorRawPacketTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ToolTrafficGeneratorRawPacketTemplateModel
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
	if !plan.Header.Equal(state.Header) && !plan.Header.IsUnknown() {
		body["header"] = plan.Header.ValueString()
	}
	if !plan.IPHeaderOffset.Equal(state.IPHeaderOffset) && !plan.IPHeaderOffset.IsUnknown() {
		body["ip-header-offset"] = plan.IPHeaderOffset.ValueString()
	}
	if !plan.IPV6HeaderOffset.Equal(state.IPV6HeaderOffset) && !plan.IPV6HeaderOffset.IsUnknown() {
		body["ipv6-header-offset"] = plan.IPV6HeaderOffset.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.RandomByteOffsetsAndMasks.Equal(state.RandomByteOffsetsAndMasks) && !plan.RandomByteOffsetsAndMasks.IsUnknown() {
		body["random-byte-offsets-and-masks"] = plan.RandomByteOffsetsAndMasks.ValueString()
	}
	if !plan.RandomRanges.Equal(state.RandomRanges) && !plan.RandomRanges.IsUnknown() {
		body["random-ranges"] = plan.RandomRanges.ValueString()
	}
	if !plan.SpecialFooter.Equal(state.SpecialFooter) && !plan.SpecialFooter.IsUnknown() {
		body["special-footer"] = client.FormatBool(plan.SpecialFooter.ValueBool())
	}
	if !plan.UDPHeaderOffset.Equal(state.UDPHeaderOffset) && !plan.UDPHeaderOffset.IsUnknown() {
		body["udp-header-offset"] = plan.UDPHeaderOffset.ValueString()
	}
	if !plan.ComputeChecksumFromOffset.Equal(state.ComputeChecksumFromOffset) && !plan.ComputeChecksumFromOffset.IsUnknown() {
		body["compute-checksum-from-offset"] = plan.ComputeChecksumFromOffset.ValueString()
	}
	if !plan.TcpHeaderOffset.Equal(state.TcpHeaderOffset) && !plan.TcpHeaderOffset.IsUnknown() {
		body["tcp-header-offset"] = plan.TcpHeaderOffset.ValueString()
	}
	if !plan.UdpComputeChecksum.Equal(state.UdpComputeChecksum) && !plan.UdpComputeChecksum.IsUnknown() {
		body["udp-compute-checksum"] = plan.UdpComputeChecksum.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/tool/traffic-generator/raw-packet-template", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /tool/traffic-generator/raw-packet-template failed", err.Error())
			return
		}
		toolTrafficGeneratorRawPacketTemplateApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolTrafficGeneratorRawPacketTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ToolTrafficGeneratorRawPacketTemplateModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/tool/traffic-generator/raw-packet-template", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /tool/traffic-generator/raw-packet-template failed", err.Error())
	}
}

func (r *ToolTrafficGeneratorRawPacketTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := toolTrafficGeneratorRawPacketTemplateLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /tool/traffic-generator/raw-packet-template matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// toolTrafficGeneratorRawPacketTemplateLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func toolTrafficGeneratorRawPacketTemplateLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/tool/traffic-generator/raw-packet-template", id)
}

func toolTrafficGeneratorRawPacketTemplateApply(ctx context.Context, obj client.Object, m *ToolTrafficGeneratorRawPacketTemplateModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["udp-compute-checksum"]; ok && v != "" {
		m.UdpComputeChecksum = types.StringValue(v)
	} else {
		m.UdpComputeChecksum = types.StringNull()
	}
	if v, ok := obj["tcp-header-offset"]; ok && v != "" {
		m.TcpHeaderOffset = types.StringValue(v)
	} else {
		m.TcpHeaderOffset = types.StringNull()
	}
	if v, ok := obj["compute-checksum-from-offset"]; ok && v != "" {
		m.ComputeChecksumFromOffset = types.StringValue(v)
	} else {
		m.ComputeChecksumFromOffset = types.StringNull()
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
	if v, ok := obj["dynamic"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else {
			m.Dynamic = types.BoolNull()
		}
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
	if v, ok := obj["header-length"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.HeaderLength = types.Int64Value(n)
		} else {
			m.HeaderLength = types.Int64Null()
		}
	} else {
		m.HeaderLength = types.Int64Null()
	}
	if v, ok := obj["ip-header-offset"]; ok {
		_ = v
		if v != "" {
			m.IPHeaderOffset = types.StringValue(v)
		} else {
			m.IPHeaderOffset = types.StringNull()
		}
	} else {
		m.IPHeaderOffset = types.StringNull()
	}
	if v, ok := obj["ipv6-header-offset"]; ok {
		_ = v
		if v != "" {
			m.IPV6HeaderOffset = types.StringValue(v)
		} else {
			m.IPV6HeaderOffset = types.StringNull()
		}
	} else {
		m.IPV6HeaderOffset = types.StringNull()
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
	if v, ok := obj["random-byte-offsets-and-masks"]; ok {
		_ = v
		if v != "" {
			m.RandomByteOffsetsAndMasks = types.StringValue(v)
		} else {
			m.RandomByteOffsetsAndMasks = types.StringNull()
		}
	} else {
		m.RandomByteOffsetsAndMasks = types.StringNull()
	}
	if v, ok := obj["random-ranges"]; ok {
		_ = v
		if v != "" {
			m.RandomRanges = types.StringValue(v)
		} else {
			m.RandomRanges = types.StringNull()
		}
	} else {
		m.RandomRanges = types.StringNull()
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
	if v, ok := obj["special-footer"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SpecialFooter = types.BoolValue(b)
		} else {
			m.SpecialFooter = types.BoolNull()
		}
	}
	if v, ok := obj["udp-header-offset"]; ok {
		_ = v
		if v != "" {
			m.UDPHeaderOffset = types.StringValue(v)
		} else {
			m.UDPHeaderOffset = types.StringNull()
		}
	} else {
		m.UDPHeaderOffset = types.StringNull()
	}
}
