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
	_ resource.Resource                = &RoutingBGPConnectionResource{}
	_ resource.ResourceWithImportState = &RoutingBGPConnectionResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingBGPConnectionResource struct {
	reg *client.Registry
}

type RoutingBGPConnectionModel struct {
	ID                          types.String `tfsdk:"id"`
	Afi                         types.String `tfsdk:"afi"`
	AllowAsIn                   types.String `tfsdk:"allow_as_in"`
	As                          types.String `tfsdk:"as"`
	AsOverride                  types.String `tfsdk:"as_override"`
	CiscoVplsNlriLengthFormat   types.String `tfsdk:"cisco_vpls_nlri_length_format"`
	Comment                     types.String `tfsdk:"comment"`
	Connect                     types.String `tfsdk:"connect"`
	DefaultOriginate            types.String `tfsdk:"default_originate"`
	DefaultPrepend              types.String `tfsdk:"default_prepend"`
	Disabled                    types.Bool   `tfsdk:"disabled"`
	Dynamic                     types.Bool   `tfsdk:"dynamic"`
	HoldTime                    types.String `tfsdk:"hold_time"`
	IgnoreAsPathLength          types.String `tfsdk:"ignore_as_path_length"`
	InputAcceptCommunities      types.String `tfsdk:"input_accept_communities"`
	InputAcceptExtCommunities   types.String `tfsdk:"input_accept_ext_communities"`
	InputAcceptLargeCommunities types.String `tfsdk:"input_accept_large_communities"`
	InputAcceptNlri             types.String `tfsdk:"input_accept_nlri"`
	InputAffinity               types.String `tfsdk:"input_affinity"`
	InputFilter                 types.String `tfsdk:"input_filter"`
	InputFilterCommunities      types.String `tfsdk:"input_filter_communities"`
	InputFilterExtCommunities   types.String `tfsdk:"input_filter_ext_communities"`
	InputFilterLargeCommunities types.String `tfsdk:"input_filter_large_communities"`
	InputFilterUnknown          types.String `tfsdk:"input_filter_unknown"`
	Instance                    types.String `tfsdk:"instance"`
	Invalid                     types.Bool   `tfsdk:"invalid"`
	KeepSentAttributes          types.String `tfsdk:"keep_sent_attributes"`
	KeepaliveTime               types.String `tfsdk:"keepalive_time"`
	Listen                      types.String `tfsdk:"listen"`
	LocalAddress                types.String `tfsdk:"local_address"`
	LocalPort                   types.String `tfsdk:"local_port"`
	LocalRole                   types.String `tfsdk:"local_role"`
	Multihop                    types.String `tfsdk:"multihop"`
	Name                        types.String `tfsdk:"name"`
	NetworkBlackhole            types.String `tfsdk:"network_blackhole"`
	NexthopChoice               types.String `tfsdk:"nexthop_choice"`
	NoClientToClientReflection  types.String `tfsdk:"no_client_to_client_reflection"`
	NoEarlyCut                  types.String `tfsdk:"no_early_cut"`
	OutputAffinity              types.String `tfsdk:"output_affinity"`
	OutputFilter                types.String `tfsdk:"output_filter"`
	OutputNetwork               types.String `tfsdk:"output_network"`
	OutputRedistribute          types.String `tfsdk:"output_redistribute"`
	OutputSelectionPolicy       types.String `tfsdk:"output_selection_policy"`
	RemoteAddress               types.String `tfsdk:"remote_address"`
	RemoteAllowAs               types.String `tfsdk:"remote_allow_as"`
	RemoteAs                    types.String `tfsdk:"remote_as"`
	RemotePort                  types.String `tfsdk:"remote_port"`
	RemovePrivateAs             types.String `tfsdk:"remove_private_as"`
	RouterID                    types.String `tfsdk:"router_id"`
	RoutingTable                types.String `tfsdk:"routing_table"`
	RxMinTtl                    types.String `tfsdk:"rx_min_ttl"`
	TCPMd5Key                   types.String `tfsdk:"tcp_md5_key"`
	Template                    types.String `tfsdk:"template"`
	TxTtl                       types.String `tfsdk:"tx_ttl"`
	UseBfd                      types.String `tfsdk:"use_bfd"`
	Vrf                         types.String `tfsdk:"vrf"`
	Router                      types.String `tfsdk:"router"`
}

func NewRoutingBGPConnectionResource() resource.Resource { return &RoutingBGPConnectionResource{} }

func (r *RoutingBGPConnectionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_bgp_connection"
}

func (r *RoutingBGPConnectionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *RoutingBGPConnectionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Auto-test requires a typed-reference precondition (e.g. an existing peer,\ninstance, bridge of the specific kind). The current acc-test generator's\ngeneric data.routeros_interface.all lookup can't satisfy these. Use this\nresource manually with explicit references to a precondition resource\nin your config.\n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"afi": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"allow_as_in": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"as": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"as_override": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"cisco_vpls_nlri_length_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"connect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"default_originate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"default_prepend": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"dynamic": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hold_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ignore_as_path_length": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"input_accept_communities": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"input_accept_ext_communities": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"input_accept_large_communities": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"input_accept_nlri": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"input_affinity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"input_filter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"input_filter_communities": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"input_filter_ext_communities": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"input_filter_large_communities": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"input_filter_unknown": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"instance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"keep_sent_attributes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"keepalive_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"listen": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"local_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"local_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"local_role": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"ibgp", "ibgp-rr", "ebgp", "ebgp-provider", "ebgp-rs", "ebgp-rs-client", "ebgp-customer", "ebgp-peer"}...)},
			},
			"multihop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"network_blackhole": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nexthop_choice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"no_client_to_client_reflection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"no_early_cut": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"output_affinity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"output_filter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"output_network": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"output_redistribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"output_selection_policy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_allow_as": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_as": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remove_private_as": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"router_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"routing_table": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_min_ttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tcp_md5_key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"template": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_ttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"use_bfd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vrf": schema.StringAttribute{
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

func (r *RoutingBGPConnectionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingBGPConnectionModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Afi.IsNull() || plan.Afi.IsUnknown()) {
		body["afi"] = plan.Afi.ValueString()
	}
	if !(plan.AllowAsIn.IsNull() || plan.AllowAsIn.IsUnknown()) {
		body["allow-as-in"] = plan.AllowAsIn.ValueString()
	}
	if !(plan.As.IsNull() || plan.As.IsUnknown()) {
		body["as"] = plan.As.ValueString()
	}
	if !(plan.AsOverride.IsNull() || plan.AsOverride.IsUnknown()) {
		body["as-override"] = plan.AsOverride.ValueString()
	}
	if !(plan.CiscoVplsNlriLengthFormat.IsNull() || plan.CiscoVplsNlriLengthFormat.IsUnknown()) {
		body["cisco-vpls-nlri-length-format"] = plan.CiscoVplsNlriLengthFormat.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Connect.IsNull() || plan.Connect.IsUnknown()) {
		body["connect"] = plan.Connect.ValueString()
	}
	if !(plan.DefaultOriginate.IsNull() || plan.DefaultOriginate.IsUnknown()) {
		body["default-originate"] = plan.DefaultOriginate.ValueString()
	}
	if !(plan.DefaultPrepend.IsNull() || plan.DefaultPrepend.IsUnknown()) {
		body["default-prepend"] = plan.DefaultPrepend.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.HoldTime.IsNull() || plan.HoldTime.IsUnknown()) {
		body["hold-time"] = plan.HoldTime.ValueString()
	}
	if !(plan.IgnoreAsPathLength.IsNull() || plan.IgnoreAsPathLength.IsUnknown()) {
		body["ignore-as-path-length"] = plan.IgnoreAsPathLength.ValueString()
	}
	if !(plan.InputAcceptCommunities.IsNull() || plan.InputAcceptCommunities.IsUnknown()) {
		body["input-accept-communities"] = plan.InputAcceptCommunities.ValueString()
	}
	if !(plan.InputAcceptExtCommunities.IsNull() || plan.InputAcceptExtCommunities.IsUnknown()) {
		body["input-accept-ext-communities"] = plan.InputAcceptExtCommunities.ValueString()
	}
	if !(plan.InputAcceptLargeCommunities.IsNull() || plan.InputAcceptLargeCommunities.IsUnknown()) {
		body["input-accept-large-communities"] = plan.InputAcceptLargeCommunities.ValueString()
	}
	if !(plan.InputAcceptNlri.IsNull() || plan.InputAcceptNlri.IsUnknown()) {
		body["input-accept-nlri"] = plan.InputAcceptNlri.ValueString()
	}
	if !(plan.InputAffinity.IsNull() || plan.InputAffinity.IsUnknown()) {
		body["input-affinity"] = plan.InputAffinity.ValueString()
	}
	if !(plan.InputFilter.IsNull() || plan.InputFilter.IsUnknown()) {
		body["input-filter"] = plan.InputFilter.ValueString()
	}
	if !(plan.InputFilterCommunities.IsNull() || plan.InputFilterCommunities.IsUnknown()) {
		body["input-filter-communities"] = plan.InputFilterCommunities.ValueString()
	}
	if !(plan.InputFilterExtCommunities.IsNull() || plan.InputFilterExtCommunities.IsUnknown()) {
		body["input-filter-ext-communities"] = plan.InputFilterExtCommunities.ValueString()
	}
	if !(plan.InputFilterLargeCommunities.IsNull() || plan.InputFilterLargeCommunities.IsUnknown()) {
		body["input-filter-large-communities"] = plan.InputFilterLargeCommunities.ValueString()
	}
	if !(plan.InputFilterUnknown.IsNull() || plan.InputFilterUnknown.IsUnknown()) {
		body["input-filter-unknown"] = plan.InputFilterUnknown.ValueString()
	}
	if !(plan.Instance.IsNull() || plan.Instance.IsUnknown()) {
		body["instance"] = plan.Instance.ValueString()
	}
	if !(plan.KeepSentAttributes.IsNull() || plan.KeepSentAttributes.IsUnknown()) {
		body["keep-sent-attributes"] = plan.KeepSentAttributes.ValueString()
	}
	if !(plan.KeepaliveTime.IsNull() || plan.KeepaliveTime.IsUnknown()) {
		body["keepalive-time"] = plan.KeepaliveTime.ValueString()
	}
	if !(plan.Listen.IsNull() || plan.Listen.IsUnknown()) {
		body["listen"] = plan.Listen.ValueString()
	}
	if !(plan.LocalAddress.IsNull() || plan.LocalAddress.IsUnknown()) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !(plan.LocalPort.IsNull() || plan.LocalPort.IsUnknown()) {
		body["local-port"] = plan.LocalPort.ValueString()
	}
	if !(plan.LocalRole.IsNull() || plan.LocalRole.IsUnknown()) {
		body["local-role"] = plan.LocalRole.ValueString()
	}
	if !(plan.Multihop.IsNull() || plan.Multihop.IsUnknown()) {
		body["multihop"] = plan.Multihop.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.NetworkBlackhole.IsNull() || plan.NetworkBlackhole.IsUnknown()) {
		body["network-blackhole"] = plan.NetworkBlackhole.ValueString()
	}
	if !(plan.NexthopChoice.IsNull() || plan.NexthopChoice.IsUnknown()) {
		body["nexthop-choice"] = plan.NexthopChoice.ValueString()
	}
	if !(plan.NoClientToClientReflection.IsNull() || plan.NoClientToClientReflection.IsUnknown()) {
		body["no-client-to-client-reflection"] = plan.NoClientToClientReflection.ValueString()
	}
	if !(plan.NoEarlyCut.IsNull() || plan.NoEarlyCut.IsUnknown()) {
		body["no-early-cut"] = plan.NoEarlyCut.ValueString()
	}
	if !(plan.OutputAffinity.IsNull() || plan.OutputAffinity.IsUnknown()) {
		body["output-affinity"] = plan.OutputAffinity.ValueString()
	}
	if !(plan.OutputFilter.IsNull() || plan.OutputFilter.IsUnknown()) {
		body["output-filter"] = plan.OutputFilter.ValueString()
	}
	if !(plan.OutputNetwork.IsNull() || plan.OutputNetwork.IsUnknown()) {
		body["output-network"] = plan.OutputNetwork.ValueString()
	}
	if !(plan.OutputRedistribute.IsNull() || plan.OutputRedistribute.IsUnknown()) {
		body["output-redistribute"] = plan.OutputRedistribute.ValueString()
	}
	if !(plan.OutputSelectionPolicy.IsNull() || plan.OutputSelectionPolicy.IsUnknown()) {
		body["output-selection-policy"] = plan.OutputSelectionPolicy.ValueString()
	}
	if !(plan.RemoteAddress.IsNull() || plan.RemoteAddress.IsUnknown()) {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !(plan.RemoteAllowAs.IsNull() || plan.RemoteAllowAs.IsUnknown()) {
		body["remote-allow-as"] = plan.RemoteAllowAs.ValueString()
	}
	if !(plan.RemoteAs.IsNull() || plan.RemoteAs.IsUnknown()) {
		body["remote-as"] = plan.RemoteAs.ValueString()
	}
	if !(plan.RemotePort.IsNull() || plan.RemotePort.IsUnknown()) {
		body["remote-port"] = plan.RemotePort.ValueString()
	}
	if !(plan.RemovePrivateAs.IsNull() || plan.RemovePrivateAs.IsUnknown()) {
		body["remove-private-as"] = plan.RemovePrivateAs.ValueString()
	}
	if !(plan.RouterID.IsNull() || plan.RouterID.IsUnknown()) {
		body["router-id"] = plan.RouterID.ValueString()
	}
	if !(plan.RoutingTable.IsNull() || plan.RoutingTable.IsUnknown()) {
		body["routing-table"] = plan.RoutingTable.ValueString()
	}
	if !(plan.RxMinTtl.IsNull() || plan.RxMinTtl.IsUnknown()) {
		body["rx-min-ttl"] = plan.RxMinTtl.ValueString()
	}
	if !(plan.TCPMd5Key.IsNull() || plan.TCPMd5Key.IsUnknown()) {
		body["tcp-md5-key"] = plan.TCPMd5Key.ValueString()
	}
	if !(plan.Template.IsNull() || plan.Template.IsUnknown()) {
		body["template"] = plan.Template.ValueString()
	}
	if !(plan.TxTtl.IsNull() || plan.TxTtl.IsUnknown()) {
		body["tx-ttl"] = plan.TxTtl.ValueString()
	}
	if !(plan.UseBfd.IsNull() || plan.UseBfd.IsUnknown()) {
		body["use-bfd"] = plan.UseBfd.ValueString()
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/bgp/connection", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/bgp/connection failed", err.Error())
		return
	}
	routingBGPConnectionApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingBGPConnectionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingBGPConnectionModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/bgp/connection", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/bgp/connection failed", err.Error())
		return
	}
	routingBGPConnectionApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingBGPConnectionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingBGPConnectionModel
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
	if !plan.Afi.Equal(state.Afi) {
		body["afi"] = plan.Afi.ValueString()
	}
	if !plan.AllowAsIn.Equal(state.AllowAsIn) {
		body["allow-as-in"] = plan.AllowAsIn.ValueString()
	}
	if !plan.As.Equal(state.As) {
		body["as"] = plan.As.ValueString()
	}
	if !plan.AsOverride.Equal(state.AsOverride) {
		body["as-override"] = plan.AsOverride.ValueString()
	}
	if !plan.CiscoVplsNlriLengthFormat.Equal(state.CiscoVplsNlriLengthFormat) {
		body["cisco-vpls-nlri-length-format"] = plan.CiscoVplsNlriLengthFormat.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Connect.Equal(state.Connect) {
		body["connect"] = plan.Connect.ValueString()
	}
	if !plan.DefaultOriginate.Equal(state.DefaultOriginate) {
		body["default-originate"] = plan.DefaultOriginate.ValueString()
	}
	if !plan.DefaultPrepend.Equal(state.DefaultPrepend) {
		body["default-prepend"] = plan.DefaultPrepend.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.HoldTime.Equal(state.HoldTime) {
		body["hold-time"] = plan.HoldTime.ValueString()
	}
	if !plan.IgnoreAsPathLength.Equal(state.IgnoreAsPathLength) {
		body["ignore-as-path-length"] = plan.IgnoreAsPathLength.ValueString()
	}
	if !plan.InputAcceptCommunities.Equal(state.InputAcceptCommunities) {
		body["input-accept-communities"] = plan.InputAcceptCommunities.ValueString()
	}
	if !plan.InputAcceptExtCommunities.Equal(state.InputAcceptExtCommunities) {
		body["input-accept-ext-communities"] = plan.InputAcceptExtCommunities.ValueString()
	}
	if !plan.InputAcceptLargeCommunities.Equal(state.InputAcceptLargeCommunities) {
		body["input-accept-large-communities"] = plan.InputAcceptLargeCommunities.ValueString()
	}
	if !plan.InputAcceptNlri.Equal(state.InputAcceptNlri) {
		body["input-accept-nlri"] = plan.InputAcceptNlri.ValueString()
	}
	if !plan.InputAffinity.Equal(state.InputAffinity) {
		body["input-affinity"] = plan.InputAffinity.ValueString()
	}
	if !plan.InputFilter.Equal(state.InputFilter) {
		body["input-filter"] = plan.InputFilter.ValueString()
	}
	if !plan.InputFilterCommunities.Equal(state.InputFilterCommunities) {
		body["input-filter-communities"] = plan.InputFilterCommunities.ValueString()
	}
	if !plan.InputFilterExtCommunities.Equal(state.InputFilterExtCommunities) {
		body["input-filter-ext-communities"] = plan.InputFilterExtCommunities.ValueString()
	}
	if !plan.InputFilterLargeCommunities.Equal(state.InputFilterLargeCommunities) {
		body["input-filter-large-communities"] = plan.InputFilterLargeCommunities.ValueString()
	}
	if !plan.InputFilterUnknown.Equal(state.InputFilterUnknown) {
		body["input-filter-unknown"] = plan.InputFilterUnknown.ValueString()
	}
	if !plan.Instance.Equal(state.Instance) {
		body["instance"] = plan.Instance.ValueString()
	}
	if !plan.KeepSentAttributes.Equal(state.KeepSentAttributes) {
		body["keep-sent-attributes"] = plan.KeepSentAttributes.ValueString()
	}
	if !plan.KeepaliveTime.Equal(state.KeepaliveTime) {
		body["keepalive-time"] = plan.KeepaliveTime.ValueString()
	}
	if !plan.Listen.Equal(state.Listen) {
		body["listen"] = plan.Listen.ValueString()
	}
	if !plan.LocalAddress.Equal(state.LocalAddress) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !plan.LocalPort.Equal(state.LocalPort) {
		body["local-port"] = plan.LocalPort.ValueString()
	}
	if !plan.LocalRole.Equal(state.LocalRole) {
		body["local-role"] = plan.LocalRole.ValueString()
	}
	if !plan.Multihop.Equal(state.Multihop) {
		body["multihop"] = plan.Multihop.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.NetworkBlackhole.Equal(state.NetworkBlackhole) {
		body["network-blackhole"] = plan.NetworkBlackhole.ValueString()
	}
	if !plan.NexthopChoice.Equal(state.NexthopChoice) {
		body["nexthop-choice"] = plan.NexthopChoice.ValueString()
	}
	if !plan.NoClientToClientReflection.Equal(state.NoClientToClientReflection) {
		body["no-client-to-client-reflection"] = plan.NoClientToClientReflection.ValueString()
	}
	if !plan.NoEarlyCut.Equal(state.NoEarlyCut) {
		body["no-early-cut"] = plan.NoEarlyCut.ValueString()
	}
	if !plan.OutputAffinity.Equal(state.OutputAffinity) {
		body["output-affinity"] = plan.OutputAffinity.ValueString()
	}
	if !plan.OutputFilter.Equal(state.OutputFilter) {
		body["output-filter"] = plan.OutputFilter.ValueString()
	}
	if !plan.OutputNetwork.Equal(state.OutputNetwork) {
		body["output-network"] = plan.OutputNetwork.ValueString()
	}
	if !plan.OutputRedistribute.Equal(state.OutputRedistribute) {
		body["output-redistribute"] = plan.OutputRedistribute.ValueString()
	}
	if !plan.OutputSelectionPolicy.Equal(state.OutputSelectionPolicy) {
		body["output-selection-policy"] = plan.OutputSelectionPolicy.ValueString()
	}
	if !plan.RemoteAddress.Equal(state.RemoteAddress) {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !plan.RemoteAllowAs.Equal(state.RemoteAllowAs) {
		body["remote-allow-as"] = plan.RemoteAllowAs.ValueString()
	}
	if !plan.RemoteAs.Equal(state.RemoteAs) {
		body["remote-as"] = plan.RemoteAs.ValueString()
	}
	if !plan.RemotePort.Equal(state.RemotePort) {
		body["remote-port"] = plan.RemotePort.ValueString()
	}
	if !plan.RemovePrivateAs.Equal(state.RemovePrivateAs) {
		body["remove-private-as"] = plan.RemovePrivateAs.ValueString()
	}
	if !plan.RouterID.Equal(state.RouterID) {
		body["router-id"] = plan.RouterID.ValueString()
	}
	if !plan.RoutingTable.Equal(state.RoutingTable) {
		body["routing-table"] = plan.RoutingTable.ValueString()
	}
	if !plan.RxMinTtl.Equal(state.RxMinTtl) {
		body["rx-min-ttl"] = plan.RxMinTtl.ValueString()
	}
	if !plan.TCPMd5Key.Equal(state.TCPMd5Key) {
		body["tcp-md5-key"] = plan.TCPMd5Key.ValueString()
	}
	if !plan.Template.Equal(state.Template) {
		body["template"] = plan.Template.ValueString()
	}
	if !plan.TxTtl.Equal(state.TxTtl) {
		body["tx-ttl"] = plan.TxTtl.ValueString()
	}
	if !plan.UseBfd.Equal(state.UseBfd) {
		body["use-bfd"] = plan.UseBfd.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/bgp/connection", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/bgp/connection failed", err.Error())
			return
		}
		routingBGPConnectionApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingBGPConnectionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingBGPConnectionModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/bgp/connection", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/bgp/connection failed", err.Error())
	}
}

func (r *RoutingBGPConnectionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingBGPConnectionLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/bgp/connection matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingBGPConnectionLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingBGPConnectionLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/bgp/connection", id)
}

func routingBGPConnectionApply(ctx context.Context, obj client.Object, m *RoutingBGPConnectionModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["afi"]; ok {
		_ = v
		if v != "" {
			m.Afi = types.StringValue(v)
		} else {
			m.Afi = types.StringNull()
		}
	} else {
		m.Afi = types.StringNull()
	}
	if v, ok := obj["allow-as-in"]; ok {
		_ = v
		if v != "" {
			m.AllowAsIn = types.StringValue(v)
		} else {
			m.AllowAsIn = types.StringNull()
		}
	} else {
		m.AllowAsIn = types.StringNull()
	}
	if v, ok := obj["as"]; ok {
		_ = v
		if v != "" {
			m.As = types.StringValue(v)
		} else {
			m.As = types.StringNull()
		}
	} else {
		m.As = types.StringNull()
	}
	if v, ok := obj["as-override"]; ok {
		_ = v
		if v != "" {
			m.AsOverride = types.StringValue(v)
		} else {
			m.AsOverride = types.StringNull()
		}
	} else {
		m.AsOverride = types.StringNull()
	}
	if v, ok := obj["cisco-vpls-nlri-length-format"]; ok {
		_ = v
		if v != "" {
			m.CiscoVplsNlriLengthFormat = types.StringValue(v)
		} else {
			m.CiscoVplsNlriLengthFormat = types.StringNull()
		}
	} else {
		m.CiscoVplsNlriLengthFormat = types.StringNull()
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
	if v, ok := obj["connect"]; ok {
		_ = v
		if v != "" {
			m.Connect = types.StringValue(v)
		} else {
			m.Connect = types.StringNull()
		}
	} else {
		m.Connect = types.StringNull()
	}
	if v, ok := obj["default-originate"]; ok {
		_ = v
		if v != "" {
			m.DefaultOriginate = types.StringValue(v)
		} else {
			m.DefaultOriginate = types.StringNull()
		}
	} else {
		m.DefaultOriginate = types.StringNull()
	}
	if v, ok := obj["default-prepend"]; ok {
		_ = v
		if v != "" {
			m.DefaultPrepend = types.StringValue(v)
		} else {
			m.DefaultPrepend = types.StringNull()
		}
	} else {
		m.DefaultPrepend = types.StringNull()
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
	if v, ok := obj["dynamic"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else {
			m.Dynamic = types.BoolNull()
		}
	} else {
		m.Dynamic = types.BoolNull()
	}
	if v, ok := obj["hold-time"]; ok {
		_ = v
		if v != "" {
			m.HoldTime = types.StringValue(v)
		} else {
			m.HoldTime = types.StringNull()
		}
	} else {
		m.HoldTime = types.StringNull()
	}
	if v, ok := obj["ignore-as-path-length"]; ok {
		_ = v
		if v != "" {
			m.IgnoreAsPathLength = types.StringValue(v)
		} else {
			m.IgnoreAsPathLength = types.StringNull()
		}
	} else {
		m.IgnoreAsPathLength = types.StringNull()
	}
	if v, ok := obj["input-accept-communities"]; ok {
		_ = v
		if v != "" {
			m.InputAcceptCommunities = types.StringValue(v)
		} else {
			m.InputAcceptCommunities = types.StringNull()
		}
	} else {
		m.InputAcceptCommunities = types.StringNull()
	}
	if v, ok := obj["input-accept-ext-communities"]; ok {
		_ = v
		if v != "" {
			m.InputAcceptExtCommunities = types.StringValue(v)
		} else {
			m.InputAcceptExtCommunities = types.StringNull()
		}
	} else {
		m.InputAcceptExtCommunities = types.StringNull()
	}
	if v, ok := obj["input-accept-large-communities"]; ok {
		_ = v
		if v != "" {
			m.InputAcceptLargeCommunities = types.StringValue(v)
		} else {
			m.InputAcceptLargeCommunities = types.StringNull()
		}
	} else {
		m.InputAcceptLargeCommunities = types.StringNull()
	}
	if v, ok := obj["input-accept-nlri"]; ok {
		_ = v
		if v != "" {
			m.InputAcceptNlri = types.StringValue(v)
		} else {
			m.InputAcceptNlri = types.StringNull()
		}
	} else {
		m.InputAcceptNlri = types.StringNull()
	}
	if v, ok := obj["input-affinity"]; ok {
		_ = v
		if v != "" {
			m.InputAffinity = types.StringValue(v)
		} else {
			m.InputAffinity = types.StringNull()
		}
	} else {
		m.InputAffinity = types.StringNull()
	}
	if v, ok := obj["input-filter"]; ok {
		_ = v
		if v != "" {
			m.InputFilter = types.StringValue(v)
		} else {
			m.InputFilter = types.StringNull()
		}
	} else {
		m.InputFilter = types.StringNull()
	}
	if v, ok := obj["input-filter-communities"]; ok {
		_ = v
		if v != "" {
			m.InputFilterCommunities = types.StringValue(v)
		} else {
			m.InputFilterCommunities = types.StringNull()
		}
	} else {
		m.InputFilterCommunities = types.StringNull()
	}
	if v, ok := obj["input-filter-ext-communities"]; ok {
		_ = v
		if v != "" {
			m.InputFilterExtCommunities = types.StringValue(v)
		} else {
			m.InputFilterExtCommunities = types.StringNull()
		}
	} else {
		m.InputFilterExtCommunities = types.StringNull()
	}
	if v, ok := obj["input-filter-large-communities"]; ok {
		_ = v
		if v != "" {
			m.InputFilterLargeCommunities = types.StringValue(v)
		} else {
			m.InputFilterLargeCommunities = types.StringNull()
		}
	} else {
		m.InputFilterLargeCommunities = types.StringNull()
	}
	if v, ok := obj["input-filter-unknown"]; ok {
		_ = v
		if v != "" {
			m.InputFilterUnknown = types.StringValue(v)
		} else {
			m.InputFilterUnknown = types.StringNull()
		}
	} else {
		m.InputFilterUnknown = types.StringNull()
	}
	if v, ok := obj["instance"]; ok {
		_ = v
		if v != "" {
			m.Instance = types.StringValue(v)
		} else {
			m.Instance = types.StringNull()
		}
	} else {
		m.Instance = types.StringNull()
	}
	if v, ok := obj["invalid"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Invalid = types.BoolValue(b)
		} else {
			m.Invalid = types.BoolNull()
		}
	} else {
		m.Invalid = types.BoolNull()
	}
	if v, ok := obj["keep-sent-attributes"]; ok {
		_ = v
		if v != "" {
			m.KeepSentAttributes = types.StringValue(v)
		} else {
			m.KeepSentAttributes = types.StringNull()
		}
	} else {
		m.KeepSentAttributes = types.StringNull()
	}
	if v, ok := obj["keepalive-time"]; ok {
		_ = v
		if v != "" {
			m.KeepaliveTime = types.StringValue(v)
		} else {
			m.KeepaliveTime = types.StringNull()
		}
	} else {
		m.KeepaliveTime = types.StringNull()
	}
	if v, ok := obj["listen"]; ok {
		_ = v
		if v != "" {
			m.Listen = types.StringValue(v)
		} else {
			m.Listen = types.StringNull()
		}
	} else {
		m.Listen = types.StringNull()
	}
	if v, ok := obj["local-address"]; ok {
		_ = v
		if v != "" {
			m.LocalAddress = types.StringValue(v)
		} else {
			m.LocalAddress = types.StringNull()
		}
	} else {
		m.LocalAddress = types.StringNull()
	}
	if v, ok := obj["local-port"]; ok {
		_ = v
		if v != "" {
			m.LocalPort = types.StringValue(v)
		} else {
			m.LocalPort = types.StringNull()
		}
	} else {
		m.LocalPort = types.StringNull()
	}
	if v, ok := obj["local-role"]; ok {
		_ = v
		if v != "" {
			m.LocalRole = types.StringValue(v)
		} else {
			m.LocalRole = types.StringNull()
		}
	} else {
		m.LocalRole = types.StringNull()
	}
	if v, ok := obj["multihop"]; ok {
		_ = v
		if v != "" {
			m.Multihop = types.StringValue(v)
		} else {
			m.Multihop = types.StringNull()
		}
	} else {
		m.Multihop = types.StringNull()
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
	if v, ok := obj["network-blackhole"]; ok {
		_ = v
		if v != "" {
			m.NetworkBlackhole = types.StringValue(v)
		} else {
			m.NetworkBlackhole = types.StringNull()
		}
	} else {
		m.NetworkBlackhole = types.StringNull()
	}
	if v, ok := obj["nexthop-choice"]; ok {
		_ = v
		if v != "" {
			m.NexthopChoice = types.StringValue(v)
		} else {
			m.NexthopChoice = types.StringNull()
		}
	} else {
		m.NexthopChoice = types.StringNull()
	}
	if v, ok := obj["no-client-to-client-reflection"]; ok {
		_ = v
		if v != "" {
			m.NoClientToClientReflection = types.StringValue(v)
		} else {
			m.NoClientToClientReflection = types.StringNull()
		}
	} else {
		m.NoClientToClientReflection = types.StringNull()
	}
	if v, ok := obj["no-early-cut"]; ok {
		_ = v
		if v != "" {
			m.NoEarlyCut = types.StringValue(v)
		} else {
			m.NoEarlyCut = types.StringNull()
		}
	} else {
		m.NoEarlyCut = types.StringNull()
	}
	if v, ok := obj["output-affinity"]; ok {
		_ = v
		if v != "" {
			m.OutputAffinity = types.StringValue(v)
		} else {
			m.OutputAffinity = types.StringNull()
		}
	} else {
		m.OutputAffinity = types.StringNull()
	}
	if v, ok := obj["output-filter"]; ok {
		_ = v
		if v != "" {
			m.OutputFilter = types.StringValue(v)
		} else {
			m.OutputFilter = types.StringNull()
		}
	} else {
		m.OutputFilter = types.StringNull()
	}
	if v, ok := obj["output-network"]; ok {
		_ = v
		if v != "" {
			m.OutputNetwork = types.StringValue(v)
		} else {
			m.OutputNetwork = types.StringNull()
		}
	} else {
		m.OutputNetwork = types.StringNull()
	}
	if v, ok := obj["output-redistribute"]; ok {
		_ = v
		if v != "" {
			m.OutputRedistribute = types.StringValue(v)
		} else {
			m.OutputRedistribute = types.StringNull()
		}
	} else {
		m.OutputRedistribute = types.StringNull()
	}
	if v, ok := obj["output-selection-policy"]; ok {
		_ = v
		if v != "" {
			m.OutputSelectionPolicy = types.StringValue(v)
		} else {
			m.OutputSelectionPolicy = types.StringNull()
		}
	} else {
		m.OutputSelectionPolicy = types.StringNull()
	}
	if v, ok := obj["remote-address"]; ok {
		_ = v
		if v != "" {
			m.RemoteAddress = types.StringValue(v)
		} else {
			m.RemoteAddress = types.StringNull()
		}
	} else {
		m.RemoteAddress = types.StringNull()
	}
	if v, ok := obj["remote-allow-as"]; ok {
		_ = v
		if v != "" {
			m.RemoteAllowAs = types.StringValue(v)
		} else {
			m.RemoteAllowAs = types.StringNull()
		}
	} else {
		m.RemoteAllowAs = types.StringNull()
	}
	if v, ok := obj["remote-as"]; ok {
		_ = v
		if v != "" {
			m.RemoteAs = types.StringValue(v)
		} else {
			m.RemoteAs = types.StringNull()
		}
	} else {
		m.RemoteAs = types.StringNull()
	}
	if v, ok := obj["remote-port"]; ok {
		_ = v
		if v != "" {
			m.RemotePort = types.StringValue(v)
		} else {
			m.RemotePort = types.StringNull()
		}
	} else {
		m.RemotePort = types.StringNull()
	}
	if v, ok := obj["remove-private-as"]; ok {
		_ = v
		if v != "" {
			m.RemovePrivateAs = types.StringValue(v)
		} else {
			m.RemovePrivateAs = types.StringNull()
		}
	} else {
		m.RemovePrivateAs = types.StringNull()
	}
	if v, ok := obj["router-id"]; ok {
		_ = v
		if v != "" {
			m.RouterID = types.StringValue(v)
		} else {
			m.RouterID = types.StringNull()
		}
	} else {
		m.RouterID = types.StringNull()
	}
	if v, ok := obj["routing-table"]; ok {
		_ = v
		if v != "" {
			m.RoutingTable = types.StringValue(v)
		} else {
			m.RoutingTable = types.StringNull()
		}
	} else {
		m.RoutingTable = types.StringNull()
	}
	if v, ok := obj["rx-min-ttl"]; ok {
		_ = v
		if v != "" {
			m.RxMinTtl = types.StringValue(v)
		} else {
			m.RxMinTtl = types.StringNull()
		}
	} else {
		m.RxMinTtl = types.StringNull()
	}
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.TCPMd5Key already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["tcp-md5-key"]; ok && v != "" {
		_ = v
		if v != "" {
			m.TCPMd5Key = types.StringValue(v)
		} else {
			m.TCPMd5Key = types.StringNull()
		}
	} else if m.TCPMd5Key.IsUnknown() {
		m.TCPMd5Key = types.StringNull()
	}
	if v, ok := obj["template"]; ok {
		_ = v
		if v != "" {
			m.Template = types.StringValue(v)
		} else {
			m.Template = types.StringNull()
		}
	} else {
		m.Template = types.StringNull()
	}
	if v, ok := obj["tx-ttl"]; ok {
		_ = v
		if v != "" {
			m.TxTtl = types.StringValue(v)
		} else {
			m.TxTtl = types.StringNull()
		}
	} else {
		m.TxTtl = types.StringNull()
	}
	if v, ok := obj["use-bfd"]; ok {
		_ = v
		if v != "" {
			m.UseBfd = types.StringValue(v)
		} else {
			m.UseBfd = types.StringNull()
		}
	} else {
		m.UseBfd = types.StringNull()
	}
	if v, ok := obj["vrf"]; ok {
		_ = v
		if v != "" {
			m.Vrf = types.StringValue(v)
		} else {
			m.Vrf = types.StringNull()
		}
	} else {
		m.Vrf = types.StringNull()
	}
}
