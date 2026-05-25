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
	ID                          types.String `tfsdk:"id"`
	Afi                         types.String `tfsdk:"afi"`
	AllowAsIn                   types.String `tfsdk:"allow_as_in"`
	As                          types.String `tfsdk:"as"`
	AsOverride                  types.String `tfsdk:"as_override"`
	CiscoVplsNlriLengthFormat   types.String `tfsdk:"cisco_vpls_nlri_length_format"`
	ClusterID                   types.String `tfsdk:"cluster_id"`
	Comment                     types.String `tfsdk:"comment"`
	Default                     types.Bool   `tfsdk:"default"`
	DefaultOriginate            types.String `tfsdk:"default_originate"`
	DefaultPrepend              types.String `tfsdk:"default_prepend"`
	Disabled                    types.Bool   `tfsdk:"disabled"`
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
	Invalid                     types.Bool   `tfsdk:"invalid"`
	KeepSentAttributes          types.String `tfsdk:"keep_sent_attributes"`
	KeepaliveTime               types.String `tfsdk:"keepalive_time"`
	Multihop                    types.String `tfsdk:"multihop"`
	Name                        types.String `tfsdk:"name"`
	NexthopChoice               types.String `tfsdk:"nexthop_choice"`
	NoClientToClientReflection  types.String `tfsdk:"no_client_to_client_reflection"`
	NoEarlyCut                  types.String `tfsdk:"no_early_cut"`
	OutputAffinity              types.String `tfsdk:"output_affinity"`
	OutputFilter                types.String `tfsdk:"output_filter"`
	OutputNetwork               types.String `tfsdk:"output_network"`
	OutputRedistribute          types.String `tfsdk:"output_redistribute"`
	OutputSelectionPolicy       types.String `tfsdk:"output_selection_policy"`
	RemovePrivateAs             types.String `tfsdk:"remove_private_as"`
	RouterID                    types.String `tfsdk:"router_id"`
	RoutingTable                types.String `tfsdk:"routing_table"`
	Templates                   types.String `tfsdk:"templates"`
	UseBfd                      types.String `tfsdk:"use_bfd"`
	Vrf                         types.String `tfsdk:"vrf"`
	Router                      types.String `tfsdk:"router"`
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
	_ = fmt.Sprintf
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
			"cluster_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"default": schema.BoolAttribute{
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
	if !(plan.ClusterID.IsNull() || plan.ClusterID.IsUnknown()) {
		body["cluster-id"] = plan.ClusterID.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
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
	if !(plan.KeepSentAttributes.IsNull() || plan.KeepSentAttributes.IsUnknown()) {
		body["keep-sent-attributes"] = plan.KeepSentAttributes.ValueString()
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
	if !(plan.RemovePrivateAs.IsNull() || plan.RemovePrivateAs.IsUnknown()) {
		body["remove-private-as"] = plan.RemovePrivateAs.ValueString()
	}
	if !(plan.RouterID.IsNull() || plan.RouterID.IsUnknown()) {
		body["router-id"] = plan.RouterID.ValueString()
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
	obj, err := c.Add(ctx, "/routing/bgp/template", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/bgp/template failed", err.Error())
		return
	}
	routingBGPTemplateApply(ctx, obj, &plan)
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
	if !plan.ClusterID.Equal(state.ClusterID) {
		body["cluster-id"] = plan.ClusterID.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
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
	if !plan.KeepSentAttributes.Equal(state.KeepSentAttributes) {
		body["keep-sent-attributes"] = plan.KeepSentAttributes.ValueString()
	}
	if !plan.KeepaliveTime.Equal(state.KeepaliveTime) {
		body["keepalive-time"] = plan.KeepaliveTime.ValueString()
	}
	if !plan.Multihop.Equal(state.Multihop) {
		body["multihop"] = plan.Multihop.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
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
	if !plan.RemovePrivateAs.Equal(state.RemovePrivateAs) {
		body["remove-private-as"] = plan.RemovePrivateAs.ValueString()
	}
	if !plan.RouterID.Equal(state.RouterID) {
		body["router-id"] = plan.RouterID.ValueString()
	}
	if !plan.RoutingTable.Equal(state.RoutingTable) {
		body["routing-table"] = plan.RoutingTable.ValueString()
	}
	if !plan.Templates.Equal(state.Templates) {
		body["templates"] = plan.Templates.ValueString()
	}
	if !plan.UseBfd.Equal(state.UseBfd) {
		body["use-bfd"] = plan.UseBfd.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) {
		body["vrf"] = plan.Vrf.ValueString()
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
