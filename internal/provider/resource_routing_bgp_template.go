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
	_ resource.Resource                = &RoutingBGPTemplateResource{}
	_ resource.ResourceWithImportState = &RoutingBGPTemplateResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingBGPTemplateResource struct {
	reg *client.Registry
}

type RoutingBGPTemplateModel struct {
	ID                               types.String `tfsdk:"id"`
	SaveTo                           types.String `tfsdk:"save_to"`
	CiscoVplsNlriLenFmt              types.String `tfsdk:"cisco_vpls_nlri_len_fmt"`
	OutputRemovePrivateAs            types.String `tfsdk:"output_remove_private_as"`
	OutputNoEarlyCut                 types.String `tfsdk:"output_no_early_cut"`
	OutputNoClientToClientReflection types.String `tfsdk:"output_no_client_to_client_reflection"`
	OutputNetworkBlackhole           types.String `tfsdk:"output_network_blackhole"`
	OutputKeepSentAttributes         types.String `tfsdk:"output_keep_sent_attributes"`
	OutputFilterSelect               types.String `tfsdk:"output_filter_select"`
	OutputFilterChain                types.String `tfsdk:"output_filter_chain"`
	OutputDefaultPrepend             types.String `tfsdk:"output_default_prepend"`
	OutputDefaultOriginate           types.String `tfsdk:"output_default_originate"`
	OutputAsOverride                 types.String `tfsdk:"output_as_override"`
	OutputAddPath                    types.String `tfsdk:"output_add_path"`
	InputLimitProcessRoutesIpv6      types.String `tfsdk:"input_limit_process_routes_ipv6"`
	InputLimitProcessRoutesIpv4      types.String `tfsdk:"input_limit_process_routes_ipv4"`
	InputFilterNlri                  types.String `tfsdk:"input_filter_nlri"`
	InputAttrErrorHandling           types.String `tfsdk:"input_attr_error_handling"`
	InputAllowAs                     types.String `tfsdk:"input_allow_as"`
	InputAddPath                     types.String `tfsdk:"input_add_path"`
	Afi                              types.String `tfsdk:"afi"`
	AllowAsIn                        types.String `tfsdk:"allow_as_in"`
	As                               types.String `tfsdk:"as"`
	AsOverride                       types.String `tfsdk:"as_override"`
	CiscoVplsNlriLengthFormat        types.String `tfsdk:"cisco_vpls_nlri_length_format"`
	ClusterID                        types.String `tfsdk:"cluster_id"`
	Comment                          types.String `tfsdk:"comment"`
	Default                          types.Bool   `tfsdk:"default"`
	DefaultOriginate                 types.String `tfsdk:"default_originate"`
	DefaultPrepend                   types.String `tfsdk:"default_prepend"`
	Disabled                         types.Bool   `tfsdk:"disabled"`
	HoldTime                         types.String `tfsdk:"hold_time"`
	IgnoreAsPathLength               types.String `tfsdk:"ignore_as_path_length"`
	InputAcceptCommunities           types.String `tfsdk:"input_accept_communities"`
	InputAcceptExtCommunities        types.String `tfsdk:"input_accept_ext_communities"`
	InputAcceptLargeCommunities      types.String `tfsdk:"input_accept_large_communities"`
	InputAcceptNlri                  types.String `tfsdk:"input_accept_nlri"`
	InputAffinity                    types.String `tfsdk:"input_affinity"`
	InputFilter                      types.String `tfsdk:"input_filter"`
	InputFilterCommunities           types.String `tfsdk:"input_filter_communities"`
	InputFilterExtCommunities        types.String `tfsdk:"input_filter_ext_communities"`
	InputFilterLargeCommunities      types.String `tfsdk:"input_filter_large_communities"`
	InputFilterUnknown               types.String `tfsdk:"input_filter_unknown"`
	Invalid                          types.Bool   `tfsdk:"invalid"`
	KeepSentAttributes               types.String `tfsdk:"keep_sent_attributes"`
	KeepaliveTime                    types.String `tfsdk:"keepalive_time"`
	Multihop                         types.String `tfsdk:"multihop"`
	Name                             types.String `tfsdk:"name"`
	NexthopChoice                    types.String `tfsdk:"nexthop_choice"`
	NoClientToClientReflection       types.String `tfsdk:"no_client_to_client_reflection"`
	NoEarlyCut                       types.String `tfsdk:"no_early_cut"`
	OutputAffinity                   types.String `tfsdk:"output_affinity"`
	OutputFilter                     types.String `tfsdk:"output_filter"`
	OutputNetwork                    types.String `tfsdk:"output_network"`
	OutputRedistribute               types.String `tfsdk:"output_redistribute"`
	OutputSelectionPolicy            types.String `tfsdk:"output_selection_policy"`
	RemovePrivateAs                  types.String `tfsdk:"remove_private_as"`
	RouterID                         types.String `tfsdk:"router_id"`
	RoutingTable                     types.String `tfsdk:"routing_table"`
	Templates                        types.String `tfsdk:"templates"`
	UseBfd                           types.String `tfsdk:"use_bfd"`
	Vrf                              types.String `tfsdk:"vrf"`
	Router                           types.String `tfsdk:"router"`
}

func NewRoutingBGPTemplateResource() resource.Resource { return &RoutingBGPTemplateResource{} }

func (r *RoutingBGPTemplateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_bgp_template"
}

func (r *RoutingBGPTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingBGPTemplateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/routing/bgp/template`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"save_to": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `save-to`.",
			},
			"cisco_vpls_nlri_len_fmt": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cisco-vpls-nlri-len-fmt`.",
			},
			"output_remove_private_as": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `output.remove-private-as`.",
			},
			"output_no_early_cut": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `output.no-early-cut`.",
			},
			"output_no_client_to_client_reflection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `output.no-client-to-client-reflection`.",
			},
			"output_network_blackhole": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `output.network-blackhole`.",
			},
			"output_keep_sent_attributes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `output.keep-sent-attributes`.",
			},
			"output_filter_select": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `output.filter-select`.",
			},
			"output_filter_chain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `output.filter-chain`.",
			},
			"output_default_prepend": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `output.default-prepend`.",
			},
			"output_default_originate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `output.default-originate`.",
			},
			"output_as_override": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `output.as-override`.",
			},
			"output_add_path": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `output.add-path`.",
			},
			"input_limit_process_routes_ipv6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `input.limit-process-routes-ipv6`.",
			},
			"input_limit_process_routes_ipv4": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `input.limit-process-routes-ipv4`.",
			},
			"input_filter_nlri": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `input.filter-nlri`.",
			},
			"input_attr_error_handling": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `input.attr-error-handling`.",
			},
			"input_allow_as": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `input.allow-as`.",
			},
			"input_add_path": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `input.add-path`.",
			},
			"afi": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"allow_as_in": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"as": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"as_override": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"cisco_vpls_nlri_length_format": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"cluster_id": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"default": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"default_originate": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"default_prepend": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"hold_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ignore_as_path_length": schema.StringAttribute{
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
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"keep_sent_attributes": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"keepalive_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"multihop": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"nexthop_choice": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"no_client_to_client_reflection": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"no_early_cut": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"output_affinity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"output_filter": schema.StringAttribute{
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
				Computed:    true,
				Description: "",
			},
			"remove_private_as": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"router_id": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"routing_table": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"templates": schema.StringAttribute{
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

func (r *RoutingBGPTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingBGPTemplateModel
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
	if !(plan.As.IsNull() || plan.As.IsUnknown()) {
		body["as"] = plan.As.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.HoldTime.IsNull() || plan.HoldTime.IsUnknown()) {
		body["hold-time"] = plan.HoldTime.ValueString()
	}
	if !(plan.InputAcceptCommunities.IsNull() || plan.InputAcceptCommunities.IsUnknown()) {
		body["input.accept-communities"] = plan.InputAcceptCommunities.ValueString()
	}
	if !(plan.InputAcceptExtCommunities.IsNull() || plan.InputAcceptExtCommunities.IsUnknown()) {
		body["input.accept-ext-communities"] = plan.InputAcceptExtCommunities.ValueString()
	}
	if !(plan.InputAcceptLargeCommunities.IsNull() || plan.InputAcceptLargeCommunities.IsUnknown()) {
		body["input.accept-large-communities"] = plan.InputAcceptLargeCommunities.ValueString()
	}
	if !(plan.InputAcceptNlri.IsNull() || plan.InputAcceptNlri.IsUnknown()) {
		body["input.accept-nlri"] = plan.InputAcceptNlri.ValueString()
	}
	if !(plan.InputAffinity.IsNull() || plan.InputAffinity.IsUnknown()) {
		body["input.affinity"] = plan.InputAffinity.ValueString()
	}
	if !(plan.InputFilter.IsNull() || plan.InputFilter.IsUnknown()) {
		body["input.filter"] = plan.InputFilter.ValueString()
	}
	if !(plan.InputFilterCommunities.IsNull() || plan.InputFilterCommunities.IsUnknown()) {
		body["input.filter-communities"] = plan.InputFilterCommunities.ValueString()
	}
	if !(plan.InputFilterExtCommunities.IsNull() || plan.InputFilterExtCommunities.IsUnknown()) {
		body["input.filter-ext-communities"] = plan.InputFilterExtCommunities.ValueString()
	}
	if !(plan.InputFilterLargeCommunities.IsNull() || plan.InputFilterLargeCommunities.IsUnknown()) {
		body["input.filter-large-communities"] = plan.InputFilterLargeCommunities.ValueString()
	}
	if !(plan.InputFilterUnknown.IsNull() || plan.InputFilterUnknown.IsUnknown()) {
		body["input.filter-unknown"] = plan.InputFilterUnknown.ValueString()
	}
	if !(plan.KeepaliveTime.IsNull() || plan.KeepaliveTime.IsUnknown()) {
		body["keepalive-time"] = plan.KeepaliveTime.ValueString()
	}
	if !(plan.Multihop.IsNull() || plan.Multihop.IsUnknown()) {
		body["multihop"] = plan.Multihop.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.NexthopChoice.IsNull() || plan.NexthopChoice.IsUnknown()) {
		body["nexthop-choice"] = plan.NexthopChoice.ValueString()
	}
	if !(plan.OutputAffinity.IsNull() || plan.OutputAffinity.IsUnknown()) {
		body["output.affinity"] = plan.OutputAffinity.ValueString()
	}
	if !(plan.OutputNetwork.IsNull() || plan.OutputNetwork.IsUnknown()) {
		body["output.network"] = plan.OutputNetwork.ValueString()
	}
	if !(plan.OutputRedistribute.IsNull() || plan.OutputRedistribute.IsUnknown()) {
		body["output.redistribute"] = plan.OutputRedistribute.ValueString()
	}
	if !(plan.RoutingTable.IsNull() || plan.RoutingTable.IsUnknown()) {
		body["routing-table"] = plan.RoutingTable.ValueString()
	}
	if !(plan.Templates.IsNull() || plan.Templates.IsUnknown()) {
		body["templates"] = plan.Templates.ValueString()
	}
	if !(plan.UseBfd.IsNull() || plan.UseBfd.IsUnknown()) {
		body["use-bfd"] = plan.UseBfd.ValueString()
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if !(plan.InputAddPath.IsNull() || plan.InputAddPath.IsUnknown()) {
		body["input.add-path"] = plan.InputAddPath.ValueString()
	}
	if !(plan.InputAllowAs.IsNull() || plan.InputAllowAs.IsUnknown()) {
		body["input.allow-as"] = plan.InputAllowAs.ValueString()
	}
	if !(plan.InputAttrErrorHandling.IsNull() || plan.InputAttrErrorHandling.IsUnknown()) {
		body["input.attr-error-handling"] = plan.InputAttrErrorHandling.ValueString()
	}
	if !(plan.InputFilterNlri.IsNull() || plan.InputFilterNlri.IsUnknown()) {
		body["input.filter-nlri"] = plan.InputFilterNlri.ValueString()
	}
	if !(plan.InputLimitProcessRoutesIpv4.IsNull() || plan.InputLimitProcessRoutesIpv4.IsUnknown()) {
		body["input.limit-process-routes-ipv4"] = plan.InputLimitProcessRoutesIpv4.ValueString()
	}
	if !(plan.InputLimitProcessRoutesIpv6.IsNull() || plan.InputLimitProcessRoutesIpv6.IsUnknown()) {
		body["input.limit-process-routes-ipv6"] = plan.InputLimitProcessRoutesIpv6.ValueString()
	}
	if !(plan.OutputAddPath.IsNull() || plan.OutputAddPath.IsUnknown()) {
		body["output.add-path"] = plan.OutputAddPath.ValueString()
	}
	if !(plan.OutputAsOverride.IsNull() || plan.OutputAsOverride.IsUnknown()) {
		body["output.as-override"] = plan.OutputAsOverride.ValueString()
	}
	if !(plan.OutputDefaultOriginate.IsNull() || plan.OutputDefaultOriginate.IsUnknown()) {
		body["output.default-originate"] = plan.OutputDefaultOriginate.ValueString()
	}
	if !(plan.OutputDefaultPrepend.IsNull() || plan.OutputDefaultPrepend.IsUnknown()) {
		body["output.default-prepend"] = plan.OutputDefaultPrepend.ValueString()
	}
	if !(plan.OutputFilterChain.IsNull() || plan.OutputFilterChain.IsUnknown()) {
		body["output.filter-chain"] = plan.OutputFilterChain.ValueString()
	}
	if !(plan.OutputFilterSelect.IsNull() || plan.OutputFilterSelect.IsUnknown()) {
		body["output.filter-select"] = plan.OutputFilterSelect.ValueString()
	}
	if !(plan.OutputKeepSentAttributes.IsNull() || plan.OutputKeepSentAttributes.IsUnknown()) {
		body["output.keep-sent-attributes"] = plan.OutputKeepSentAttributes.ValueString()
	}
	if !(plan.OutputNetworkBlackhole.IsNull() || plan.OutputNetworkBlackhole.IsUnknown()) {
		body["output.network-blackhole"] = plan.OutputNetworkBlackhole.ValueString()
	}
	if !(plan.OutputNoClientToClientReflection.IsNull() || plan.OutputNoClientToClientReflection.IsUnknown()) {
		body["output.no-client-to-client-reflection"] = plan.OutputNoClientToClientReflection.ValueString()
	}
	if !(plan.OutputNoEarlyCut.IsNull() || plan.OutputNoEarlyCut.IsUnknown()) {
		body["output.no-early-cut"] = plan.OutputNoEarlyCut.ValueString()
	}
	if !(plan.OutputRemovePrivateAs.IsNull() || plan.OutputRemovePrivateAs.IsUnknown()) {
		body["output.remove-private-as"] = plan.OutputRemovePrivateAs.ValueString()
	}
	if !(plan.CiscoVplsNlriLenFmt.IsNull() || plan.CiscoVplsNlriLenFmt.IsUnknown()) {
		body["cisco-vpls-nlri-len-fmt"] = plan.CiscoVplsNlriLenFmt.ValueString()
	}
	if !(plan.SaveTo.IsNull() || plan.SaveTo.IsUnknown()) {
		body["save-to"] = plan.SaveTo.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/bgp/template", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/bgp/template failed", err.Error())
		return
	}
	routingBGPTemplateApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingBGPTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingBGPTemplateModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/bgp/template", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/bgp/template failed", err.Error())
		return
	}
	routingBGPTemplateApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingBGPTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingBGPTemplateModel
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
	if !plan.Afi.Equal(state.Afi) && !plan.Afi.IsUnknown() {
		body["afi"] = plan.Afi.ValueString()
	}
	if !plan.As.Equal(state.As) && !plan.As.IsUnknown() {
		body["as"] = plan.As.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.HoldTime.Equal(state.HoldTime) && !plan.HoldTime.IsUnknown() {
		body["hold-time"] = plan.HoldTime.ValueString()
	}
	if !plan.InputAcceptCommunities.Equal(state.InputAcceptCommunities) && !plan.InputAcceptCommunities.IsUnknown() {
		body["input.accept-communities"] = plan.InputAcceptCommunities.ValueString()
	}
	if !plan.InputAcceptExtCommunities.Equal(state.InputAcceptExtCommunities) && !plan.InputAcceptExtCommunities.IsUnknown() {
		body["input.accept-ext-communities"] = plan.InputAcceptExtCommunities.ValueString()
	}
	if !plan.InputAcceptLargeCommunities.Equal(state.InputAcceptLargeCommunities) && !plan.InputAcceptLargeCommunities.IsUnknown() {
		body["input.accept-large-communities"] = plan.InputAcceptLargeCommunities.ValueString()
	}
	if !plan.InputAcceptNlri.Equal(state.InputAcceptNlri) && !plan.InputAcceptNlri.IsUnknown() {
		body["input.accept-nlri"] = plan.InputAcceptNlri.ValueString()
	}
	if !plan.InputAffinity.Equal(state.InputAffinity) && !plan.InputAffinity.IsUnknown() {
		body["input.affinity"] = plan.InputAffinity.ValueString()
	}
	if !plan.InputFilter.Equal(state.InputFilter) && !plan.InputFilter.IsUnknown() {
		body["input.filter"] = plan.InputFilter.ValueString()
	}
	if !plan.InputFilterCommunities.Equal(state.InputFilterCommunities) && !plan.InputFilterCommunities.IsUnknown() {
		body["input.filter-communities"] = plan.InputFilterCommunities.ValueString()
	}
	if !plan.InputFilterExtCommunities.Equal(state.InputFilterExtCommunities) && !plan.InputFilterExtCommunities.IsUnknown() {
		body["input.filter-ext-communities"] = plan.InputFilterExtCommunities.ValueString()
	}
	if !plan.InputFilterLargeCommunities.Equal(state.InputFilterLargeCommunities) && !plan.InputFilterLargeCommunities.IsUnknown() {
		body["input.filter-large-communities"] = plan.InputFilterLargeCommunities.ValueString()
	}
	if !plan.InputFilterUnknown.Equal(state.InputFilterUnknown) && !plan.InputFilterUnknown.IsUnknown() {
		body["input.filter-unknown"] = plan.InputFilterUnknown.ValueString()
	}
	if !plan.KeepaliveTime.Equal(state.KeepaliveTime) && !plan.KeepaliveTime.IsUnknown() {
		body["keepalive-time"] = plan.KeepaliveTime.ValueString()
	}
	if !plan.Multihop.Equal(state.Multihop) && !plan.Multihop.IsUnknown() {
		body["multihop"] = plan.Multihop.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.NexthopChoice.Equal(state.NexthopChoice) && !plan.NexthopChoice.IsUnknown() {
		body["nexthop-choice"] = plan.NexthopChoice.ValueString()
	}
	if !plan.OutputAffinity.Equal(state.OutputAffinity) && !plan.OutputAffinity.IsUnknown() {
		body["output.affinity"] = plan.OutputAffinity.ValueString()
	}
	if !plan.OutputNetwork.Equal(state.OutputNetwork) && !plan.OutputNetwork.IsUnknown() {
		body["output.network"] = plan.OutputNetwork.ValueString()
	}
	if !plan.OutputRedistribute.Equal(state.OutputRedistribute) && !plan.OutputRedistribute.IsUnknown() {
		body["output.redistribute"] = plan.OutputRedistribute.ValueString()
	}
	if !plan.RoutingTable.Equal(state.RoutingTable) && !plan.RoutingTable.IsUnknown() {
		body["routing-table"] = plan.RoutingTable.ValueString()
	}
	if !plan.Templates.Equal(state.Templates) && !plan.Templates.IsUnknown() {
		body["templates"] = plan.Templates.ValueString()
	}
	if !plan.UseBfd.Equal(state.UseBfd) && !plan.UseBfd.IsUnknown() {
		body["use-bfd"] = plan.UseBfd.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) && !plan.Vrf.IsUnknown() {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if !plan.InputAddPath.Equal(state.InputAddPath) && !plan.InputAddPath.IsUnknown() {
		body["input.add-path"] = plan.InputAddPath.ValueString()
	}
	if !plan.InputAllowAs.Equal(state.InputAllowAs) && !plan.InputAllowAs.IsUnknown() {
		body["input.allow-as"] = plan.InputAllowAs.ValueString()
	}
	if !plan.InputAttrErrorHandling.Equal(state.InputAttrErrorHandling) && !plan.InputAttrErrorHandling.IsUnknown() {
		body["input.attr-error-handling"] = plan.InputAttrErrorHandling.ValueString()
	}
	if !plan.InputFilterNlri.Equal(state.InputFilterNlri) && !plan.InputFilterNlri.IsUnknown() {
		body["input.filter-nlri"] = plan.InputFilterNlri.ValueString()
	}
	if !plan.InputLimitProcessRoutesIpv4.Equal(state.InputLimitProcessRoutesIpv4) && !plan.InputLimitProcessRoutesIpv4.IsUnknown() {
		body["input.limit-process-routes-ipv4"] = plan.InputLimitProcessRoutesIpv4.ValueString()
	}
	if !plan.InputLimitProcessRoutesIpv6.Equal(state.InputLimitProcessRoutesIpv6) && !plan.InputLimitProcessRoutesIpv6.IsUnknown() {
		body["input.limit-process-routes-ipv6"] = plan.InputLimitProcessRoutesIpv6.ValueString()
	}
	if !plan.OutputAddPath.Equal(state.OutputAddPath) && !plan.OutputAddPath.IsUnknown() {
		body["output.add-path"] = plan.OutputAddPath.ValueString()
	}
	if !plan.OutputAsOverride.Equal(state.OutputAsOverride) && !plan.OutputAsOverride.IsUnknown() {
		body["output.as-override"] = plan.OutputAsOverride.ValueString()
	}
	if !plan.OutputDefaultOriginate.Equal(state.OutputDefaultOriginate) && !plan.OutputDefaultOriginate.IsUnknown() {
		body["output.default-originate"] = plan.OutputDefaultOriginate.ValueString()
	}
	if !plan.OutputDefaultPrepend.Equal(state.OutputDefaultPrepend) && !plan.OutputDefaultPrepend.IsUnknown() {
		body["output.default-prepend"] = plan.OutputDefaultPrepend.ValueString()
	}
	if !plan.OutputFilterChain.Equal(state.OutputFilterChain) && !plan.OutputFilterChain.IsUnknown() {
		body["output.filter-chain"] = plan.OutputFilterChain.ValueString()
	}
	if !plan.OutputFilterSelect.Equal(state.OutputFilterSelect) && !plan.OutputFilterSelect.IsUnknown() {
		body["output.filter-select"] = plan.OutputFilterSelect.ValueString()
	}
	if !plan.OutputKeepSentAttributes.Equal(state.OutputKeepSentAttributes) && !plan.OutputKeepSentAttributes.IsUnknown() {
		body["output.keep-sent-attributes"] = plan.OutputKeepSentAttributes.ValueString()
	}
	if !plan.OutputNetworkBlackhole.Equal(state.OutputNetworkBlackhole) && !plan.OutputNetworkBlackhole.IsUnknown() {
		body["output.network-blackhole"] = plan.OutputNetworkBlackhole.ValueString()
	}
	if !plan.OutputNoClientToClientReflection.Equal(state.OutputNoClientToClientReflection) && !plan.OutputNoClientToClientReflection.IsUnknown() {
		body["output.no-client-to-client-reflection"] = plan.OutputNoClientToClientReflection.ValueString()
	}
	if !plan.OutputNoEarlyCut.Equal(state.OutputNoEarlyCut) && !plan.OutputNoEarlyCut.IsUnknown() {
		body["output.no-early-cut"] = plan.OutputNoEarlyCut.ValueString()
	}
	if !plan.OutputRemovePrivateAs.Equal(state.OutputRemovePrivateAs) && !plan.OutputRemovePrivateAs.IsUnknown() {
		body["output.remove-private-as"] = plan.OutputRemovePrivateAs.ValueString()
	}
	if !plan.CiscoVplsNlriLenFmt.Equal(state.CiscoVplsNlriLenFmt) && !plan.CiscoVplsNlriLenFmt.IsUnknown() {
		body["cisco-vpls-nlri-len-fmt"] = plan.CiscoVplsNlriLenFmt.ValueString()
	}
	if !plan.SaveTo.Equal(state.SaveTo) && !plan.SaveTo.IsUnknown() {
		body["save-to"] = plan.SaveTo.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/bgp/template", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/bgp/template failed", err.Error())
			return
		}
		routingBGPTemplateApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingBGPTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingBGPTemplateModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/bgp/template", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/bgp/template failed", err.Error())
	}
}

func (r *RoutingBGPTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingBGPTemplateLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/bgp/template matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingBGPTemplateLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingBGPTemplateLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/bgp/template", id)
}

func routingBGPTemplateApply(ctx context.Context, obj client.Object, m *RoutingBGPTemplateModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["save-to"]; ok && v != "" {
		m.SaveTo = types.StringValue(v)
	} else {
		m.SaveTo = types.StringNull()
	}
	if v, ok := obj["cisco-vpls-nlri-len-fmt"]; ok && v != "" {
		m.CiscoVplsNlriLenFmt = types.StringValue(v)
	} else {
		m.CiscoVplsNlriLenFmt = types.StringNull()
	}
	if v, ok := obj["output.remove-private-as"]; ok && v != "" {
		m.OutputRemovePrivateAs = types.StringValue(v)
	} else {
		m.OutputRemovePrivateAs = types.StringNull()
	}
	if v, ok := obj["output.no-early-cut"]; ok && v != "" {
		m.OutputNoEarlyCut = types.StringValue(v)
	} else {
		m.OutputNoEarlyCut = types.StringNull()
	}
	if v, ok := obj["output.no-client-to-client-reflection"]; ok && v != "" {
		m.OutputNoClientToClientReflection = types.StringValue(v)
	} else {
		m.OutputNoClientToClientReflection = types.StringNull()
	}
	if v, ok := obj["output.network-blackhole"]; ok && v != "" {
		m.OutputNetworkBlackhole = types.StringValue(v)
	} else {
		m.OutputNetworkBlackhole = types.StringNull()
	}
	if v, ok := obj["output.keep-sent-attributes"]; ok && v != "" {
		m.OutputKeepSentAttributes = types.StringValue(v)
	} else {
		m.OutputKeepSentAttributes = types.StringNull()
	}
	if v, ok := obj["output.filter-select"]; ok && v != "" {
		m.OutputFilterSelect = types.StringValue(v)
	} else {
		m.OutputFilterSelect = types.StringNull()
	}
	if v, ok := obj["output.filter-chain"]; ok && v != "" {
		m.OutputFilterChain = types.StringValue(v)
	} else {
		m.OutputFilterChain = types.StringNull()
	}
	if v, ok := obj["output.default-prepend"]; ok && v != "" {
		m.OutputDefaultPrepend = types.StringValue(v)
	} else {
		m.OutputDefaultPrepend = types.StringNull()
	}
	if v, ok := obj["output.default-originate"]; ok && v != "" {
		m.OutputDefaultOriginate = types.StringValue(v)
	} else {
		m.OutputDefaultOriginate = types.StringNull()
	}
	if v, ok := obj["output.as-override"]; ok && v != "" {
		m.OutputAsOverride = types.StringValue(v)
	} else {
		m.OutputAsOverride = types.StringNull()
	}
	if v, ok := obj["output.add-path"]; ok && v != "" {
		m.OutputAddPath = types.StringValue(v)
	} else {
		m.OutputAddPath = types.StringNull()
	}
	if v, ok := obj["input.limit-process-routes-ipv6"]; ok && v != "" {
		m.InputLimitProcessRoutesIpv6 = types.StringValue(v)
	} else {
		m.InputLimitProcessRoutesIpv6 = types.StringNull()
	}
	if v, ok := obj["input.limit-process-routes-ipv4"]; ok && v != "" {
		m.InputLimitProcessRoutesIpv4 = types.StringValue(v)
	} else {
		m.InputLimitProcessRoutesIpv4 = types.StringNull()
	}
	if v, ok := obj["input.filter-nlri"]; ok && v != "" {
		m.InputFilterNlri = types.StringValue(v)
	} else {
		m.InputFilterNlri = types.StringNull()
	}
	if v, ok := obj["input.attr-error-handling"]; ok && v != "" {
		m.InputAttrErrorHandling = types.StringValue(v)
	} else {
		m.InputAttrErrorHandling = types.StringNull()
	}
	if v, ok := obj["input.allow-as"]; ok && v != "" {
		m.InputAllowAs = types.StringValue(v)
	} else {
		m.InputAllowAs = types.StringNull()
	}
	if v, ok := obj["input.add-path"]; ok && v != "" {
		m.InputAddPath = types.StringValue(v)
	} else {
		m.InputAddPath = types.StringNull()
	}
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
	if v, ok := obj["cluster-id"]; ok {
		_ = v
		if v != "" {
			m.ClusterID = types.StringValue(v)
		} else {
			m.ClusterID = types.StringNull()
		}
	} else {
		m.ClusterID = types.StringNull()
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
	if v, ok := obj["default"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else {
			m.Default = types.BoolNull()
		}
	} else {
		m.Default = types.BoolNull()
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
	if v, ok := obj["input.accept-communities"]; ok {
		_ = v
		if v != "" {
			m.InputAcceptCommunities = types.StringValue(v)
		} else {
			m.InputAcceptCommunities = types.StringNull()
		}
	} else {
		m.InputAcceptCommunities = types.StringNull()
	}
	if v, ok := obj["input.accept-ext-communities"]; ok {
		_ = v
		if v != "" {
			m.InputAcceptExtCommunities = types.StringValue(v)
		} else {
			m.InputAcceptExtCommunities = types.StringNull()
		}
	} else {
		m.InputAcceptExtCommunities = types.StringNull()
	}
	if v, ok := obj["input.accept-large-communities"]; ok {
		_ = v
		if v != "" {
			m.InputAcceptLargeCommunities = types.StringValue(v)
		} else {
			m.InputAcceptLargeCommunities = types.StringNull()
		}
	} else {
		m.InputAcceptLargeCommunities = types.StringNull()
	}
	if v, ok := obj["input.accept-nlri"]; ok {
		_ = v
		if v != "" {
			m.InputAcceptNlri = types.StringValue(v)
		} else {
			m.InputAcceptNlri = types.StringNull()
		}
	} else {
		m.InputAcceptNlri = types.StringNull()
	}
	if v, ok := obj["input.affinity"]; ok {
		_ = v
		if v != "" {
			m.InputAffinity = types.StringValue(v)
		} else {
			m.InputAffinity = types.StringNull()
		}
	} else {
		m.InputAffinity = types.StringNull()
	}
	if v, ok := obj["input.filter"]; ok {
		_ = v
		if v != "" {
			m.InputFilter = types.StringValue(v)
		} else {
			m.InputFilter = types.StringNull()
		}
	} else {
		m.InputFilter = types.StringNull()
	}
	if v, ok := obj["input.filter-communities"]; ok {
		_ = v
		if v != "" {
			m.InputFilterCommunities = types.StringValue(v)
		} else {
			m.InputFilterCommunities = types.StringNull()
		}
	} else {
		m.InputFilterCommunities = types.StringNull()
	}
	if v, ok := obj["input.filter-ext-communities"]; ok {
		_ = v
		if v != "" {
			m.InputFilterExtCommunities = types.StringValue(v)
		} else {
			m.InputFilterExtCommunities = types.StringNull()
		}
	} else {
		m.InputFilterExtCommunities = types.StringNull()
	}
	if v, ok := obj["input.filter-large-communities"]; ok {
		_ = v
		if v != "" {
			m.InputFilterLargeCommunities = types.StringValue(v)
		} else {
			m.InputFilterLargeCommunities = types.StringNull()
		}
	} else {
		m.InputFilterLargeCommunities = types.StringNull()
	}
	if v, ok := obj["input.filter-unknown"]; ok {
		_ = v
		if v != "" {
			m.InputFilterUnknown = types.StringValue(v)
		} else {
			m.InputFilterUnknown = types.StringNull()
		}
	} else {
		m.InputFilterUnknown = types.StringNull()
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
	if v, ok := obj["output.affinity"]; ok {
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
	if v, ok := obj["output.network"]; ok {
		_ = v
		if v != "" {
			m.OutputNetwork = types.StringValue(v)
		} else {
			m.OutputNetwork = types.StringNull()
		}
	} else {
		m.OutputNetwork = types.StringNull()
	}
	if v, ok := obj["output.redistribute"]; ok {
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
	if v, ok := obj["templates"]; ok {
		_ = v
		if v != "" {
			m.Templates = types.StringValue(v)
		} else {
			m.Templates = types.StringNull()
		}
	} else {
		m.Templates = types.StringNull()
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
