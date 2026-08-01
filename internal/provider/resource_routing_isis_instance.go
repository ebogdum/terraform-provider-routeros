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
	_ resource.Resource                = &RoutingIsisInstanceResource{}
	_ resource.ResourceWithImportState = &RoutingIsisInstanceResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingIsisInstanceResource struct {
	reg *client.Registry
}

type RoutingIsisInstanceModel struct {
	ID                   types.String `tfsdk:"id"`
	Vrf                  types.String `tfsdk:"vrf"`
	SystemId             types.String `tfsdk:"system_id"`
	Name                 types.String `tfsdk:"name"`
	MetricType           types.String `tfsdk:"metric_type"`
	InFilterChain        types.String `tfsdk:"in_filter_chain"`
	AreasMax             types.String `tfsdk:"areas_max"`
	Areas                types.String `tfsdk:"areas"`
	Afi                  types.String `tfsdk:"afi"`
	L2Redistribute       types.String `tfsdk:"l2_redistribute"`
	L2OutFilterSelect    types.String `tfsdk:"l2_out_filter_select"`
	L2OutFilterChain     types.String `tfsdk:"l2_out_filter_chain"`
	L2OriginateDefault   types.String `tfsdk:"l2_originate_default"`
	L2LspUpdateInterval  types.String `tfsdk:"l2_lsp_update_interval"`
	L2LspMaxSize         types.String `tfsdk:"l2_lsp_max_size"`
	L2LspMaxAge          types.String `tfsdk:"l2_lsp_max_age"`
	L1Redistribute       types.String `tfsdk:"l1_redistribute"`
	L1OutFilterSelect    types.String `tfsdk:"l1_out_filter_select"`
	L1OutFilterChain     types.String `tfsdk:"l1_out_filter_chain"`
	L1OriginateDefault   types.String `tfsdk:"l1_originate_default"`
	L1LspUpdateInterval  types.String `tfsdk:"l1_lsp_update_interval"`
	L1LspRefreshInterval types.String `tfsdk:"l1_lsp_refresh_interval"`
	L1LspMaxSize         types.String `tfsdk:"l1_lsp_max_size"`
	L1LspMaxAge          types.String `tfsdk:"l1_lsp_max_age"`
	Comment              types.String `tfsdk:"comment"`
	Disabled             types.Bool   `tfsdk:"disabled"`
	Router               types.String `tfsdk:"router"`
}

func NewRoutingIsisInstanceResource() resource.Resource { return &RoutingIsisInstanceResource{} }

func (r *RoutingIsisInstanceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_isis_instance"
}

func (r *RoutingIsisInstanceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingIsisInstanceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "ISIS instance argument set differs across ROS releases. Skipped.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"vrf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `vrf`.",
			},
			"system_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `system-id`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `name`.",
			},
			"metric_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `metric-type`.",
			},
			"in_filter_chain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `in-filter-chain`.",
			},
			"areas_max": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `areas-max`.",
			},
			"areas": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `areas`.",
			},
			"afi": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `afi`.",
			},
			"l2_redistribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2.redistribute`.",
			},
			"l2_out_filter_select": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2.out-filter-select`.",
			},
			"l2_out_filter_chain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2.out-filter-chain`.",
			},
			"l2_originate_default": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2.originate-default`.",
			},
			"l2_lsp_update_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2.lsp-update-interval`.",
			},
			"l2_lsp_max_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2.lsp-max-size`.",
			},
			"l2_lsp_max_age": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2.lsp-max-age`.",
			},
			"l1_redistribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l1.redistribute`.",
			},
			"l1_out_filter_select": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l1.out-filter-select`.",
			},
			"l1_out_filter_chain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l1.out-filter-chain`.",
			},
			"l1_originate_default": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l1.originate-default`.",
			},
			"l1_lsp_update_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l1.lsp-update-interval`.",
			},
			"l1_lsp_refresh_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l1.lsp-refresh-interval`.",
			},
			"l1_lsp_max_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l1.lsp-max-size`.",
			},
			"l1_lsp_max_age": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l1.lsp-max-age`.",
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
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *RoutingIsisInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingIsisInstanceModel
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
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.L1LspMaxAge.IsNull() || plan.L1LspMaxAge.IsUnknown()) {
		body["l1.lsp-max-age"] = plan.L1LspMaxAge.ValueString()
	}
	if !(plan.L1LspMaxSize.IsNull() || plan.L1LspMaxSize.IsUnknown()) {
		body["l1.lsp-max-size"] = plan.L1LspMaxSize.ValueString()
	}
	if !(plan.L1LspRefreshInterval.IsNull() || plan.L1LspRefreshInterval.IsUnknown()) {
		body["l1.lsp-refresh-interval"] = plan.L1LspRefreshInterval.ValueString()
	}
	if !(plan.L1LspUpdateInterval.IsNull() || plan.L1LspUpdateInterval.IsUnknown()) {
		body["l1.lsp-update-interval"] = plan.L1LspUpdateInterval.ValueString()
	}
	if !(plan.L1OriginateDefault.IsNull() || plan.L1OriginateDefault.IsUnknown()) {
		body["l1.originate-default"] = plan.L1OriginateDefault.ValueString()
	}
	if !(plan.L1OutFilterChain.IsNull() || plan.L1OutFilterChain.IsUnknown()) {
		body["l1.out-filter-chain"] = plan.L1OutFilterChain.ValueString()
	}
	if !(plan.L1OutFilterSelect.IsNull() || plan.L1OutFilterSelect.IsUnknown()) {
		body["l1.out-filter-select"] = plan.L1OutFilterSelect.ValueString()
	}
	if !(plan.L1Redistribute.IsNull() || plan.L1Redistribute.IsUnknown()) {
		body["l1.redistribute"] = plan.L1Redistribute.ValueString()
	}
	if !(plan.L2LspMaxAge.IsNull() || plan.L2LspMaxAge.IsUnknown()) {
		body["l2.lsp-max-age"] = plan.L2LspMaxAge.ValueString()
	}
	if !(plan.L2LspMaxSize.IsNull() || plan.L2LspMaxSize.IsUnknown()) {
		body["l2.lsp-max-size"] = plan.L2LspMaxSize.ValueString()
	}
	if !(plan.L2LspUpdateInterval.IsNull() || plan.L2LspUpdateInterval.IsUnknown()) {
		body["l2.lsp-update-interval"] = plan.L2LspUpdateInterval.ValueString()
	}
	if !(plan.L2OriginateDefault.IsNull() || plan.L2OriginateDefault.IsUnknown()) {
		body["l2.originate-default"] = plan.L2OriginateDefault.ValueString()
	}
	if !(plan.L2OutFilterChain.IsNull() || plan.L2OutFilterChain.IsUnknown()) {
		body["l2.out-filter-chain"] = plan.L2OutFilterChain.ValueString()
	}
	if !(plan.L2OutFilterSelect.IsNull() || plan.L2OutFilterSelect.IsUnknown()) {
		body["l2.out-filter-select"] = plan.L2OutFilterSelect.ValueString()
	}
	if !(plan.L2Redistribute.IsNull() || plan.L2Redistribute.IsUnknown()) {
		body["l2.redistribute"] = plan.L2Redistribute.ValueString()
	}
	if !(plan.Afi.IsNull() || plan.Afi.IsUnknown()) {
		body["afi"] = plan.Afi.ValueString()
	}
	if !(plan.Areas.IsNull() || plan.Areas.IsUnknown()) {
		body["areas"] = plan.Areas.ValueString()
	}
	if !(plan.AreasMax.IsNull() || plan.AreasMax.IsUnknown()) {
		body["areas-max"] = plan.AreasMax.ValueString()
	}
	if !(plan.InFilterChain.IsNull() || plan.InFilterChain.IsUnknown()) {
		body["in-filter-chain"] = plan.InFilterChain.ValueString()
	}
	if !(plan.MetricType.IsNull() || plan.MetricType.IsUnknown()) {
		body["metric-type"] = plan.MetricType.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.SystemId.IsNull() || plan.SystemId.IsUnknown()) {
		body["system-id"] = plan.SystemId.ValueString()
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/isis/instance", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/isis/instance failed", err.Error())
		return
	}
	routingIsisInstanceApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingIsisInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingIsisInstanceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/isis/instance", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/isis/instance failed", err.Error())
		return
	}
	routingIsisInstanceApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingIsisInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingIsisInstanceModel
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
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.L1LspMaxAge.Equal(state.L1LspMaxAge) && !plan.L1LspMaxAge.IsUnknown() {
		body["l1.lsp-max-age"] = plan.L1LspMaxAge.ValueString()
	}
	if !plan.L1LspMaxSize.Equal(state.L1LspMaxSize) && !plan.L1LspMaxSize.IsUnknown() {
		body["l1.lsp-max-size"] = plan.L1LspMaxSize.ValueString()
	}
	if !plan.L1LspRefreshInterval.Equal(state.L1LspRefreshInterval) && !plan.L1LspRefreshInterval.IsUnknown() {
		body["l1.lsp-refresh-interval"] = plan.L1LspRefreshInterval.ValueString()
	}
	if !plan.L1LspUpdateInterval.Equal(state.L1LspUpdateInterval) && !plan.L1LspUpdateInterval.IsUnknown() {
		body["l1.lsp-update-interval"] = plan.L1LspUpdateInterval.ValueString()
	}
	if !plan.L1OriginateDefault.Equal(state.L1OriginateDefault) && !plan.L1OriginateDefault.IsUnknown() {
		body["l1.originate-default"] = plan.L1OriginateDefault.ValueString()
	}
	if !plan.L1OutFilterChain.Equal(state.L1OutFilterChain) && !plan.L1OutFilterChain.IsUnknown() {
		body["l1.out-filter-chain"] = plan.L1OutFilterChain.ValueString()
	}
	if !plan.L1OutFilterSelect.Equal(state.L1OutFilterSelect) && !plan.L1OutFilterSelect.IsUnknown() {
		body["l1.out-filter-select"] = plan.L1OutFilterSelect.ValueString()
	}
	if !plan.L1Redistribute.Equal(state.L1Redistribute) && !plan.L1Redistribute.IsUnknown() {
		body["l1.redistribute"] = plan.L1Redistribute.ValueString()
	}
	if !plan.L2LspMaxAge.Equal(state.L2LspMaxAge) && !plan.L2LspMaxAge.IsUnknown() {
		body["l2.lsp-max-age"] = plan.L2LspMaxAge.ValueString()
	}
	if !plan.L2LspMaxSize.Equal(state.L2LspMaxSize) && !plan.L2LspMaxSize.IsUnknown() {
		body["l2.lsp-max-size"] = plan.L2LspMaxSize.ValueString()
	}
	if !plan.L2LspUpdateInterval.Equal(state.L2LspUpdateInterval) && !plan.L2LspUpdateInterval.IsUnknown() {
		body["l2.lsp-update-interval"] = plan.L2LspUpdateInterval.ValueString()
	}
	if !plan.L2OriginateDefault.Equal(state.L2OriginateDefault) && !plan.L2OriginateDefault.IsUnknown() {
		body["l2.originate-default"] = plan.L2OriginateDefault.ValueString()
	}
	if !plan.L2OutFilterChain.Equal(state.L2OutFilterChain) && !plan.L2OutFilterChain.IsUnknown() {
		body["l2.out-filter-chain"] = plan.L2OutFilterChain.ValueString()
	}
	if !plan.L2OutFilterSelect.Equal(state.L2OutFilterSelect) && !plan.L2OutFilterSelect.IsUnknown() {
		body["l2.out-filter-select"] = plan.L2OutFilterSelect.ValueString()
	}
	if !plan.L2Redistribute.Equal(state.L2Redistribute) && !plan.L2Redistribute.IsUnknown() {
		body["l2.redistribute"] = plan.L2Redistribute.ValueString()
	}
	if !plan.Afi.Equal(state.Afi) && !plan.Afi.IsUnknown() {
		body["afi"] = plan.Afi.ValueString()
	}
	if !plan.Areas.Equal(state.Areas) && !plan.Areas.IsUnknown() {
		body["areas"] = plan.Areas.ValueString()
	}
	if !plan.AreasMax.Equal(state.AreasMax) && !plan.AreasMax.IsUnknown() {
		body["areas-max"] = plan.AreasMax.ValueString()
	}
	if !plan.InFilterChain.Equal(state.InFilterChain) && !plan.InFilterChain.IsUnknown() {
		body["in-filter-chain"] = plan.InFilterChain.ValueString()
	}
	if !plan.MetricType.Equal(state.MetricType) && !plan.MetricType.IsUnknown() {
		body["metric-type"] = plan.MetricType.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.SystemId.Equal(state.SystemId) && !plan.SystemId.IsUnknown() {
		body["system-id"] = plan.SystemId.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) && !plan.Vrf.IsUnknown() {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/isis/instance", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/isis/instance failed", err.Error())
			return
		}
		routingIsisInstanceApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingIsisInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingIsisInstanceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/isis/instance", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/isis/instance failed", err.Error())
	}
}

func (r *RoutingIsisInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingIsisInstanceLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/isis/instance matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingIsisInstanceLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingIsisInstanceLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/isis/instance", id)
}

func routingIsisInstanceApply(ctx context.Context, obj client.Object, m *RoutingIsisInstanceModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["vrf"]; ok && v != "" {
		m.Vrf = types.StringValue(v)
	} else {
		m.Vrf = types.StringNull()
	}
	if v, ok := obj["system-id"]; ok && v != "" {
		m.SystemId = types.StringValue(v)
	} else {
		m.SystemId = types.StringNull()
	}
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["metric-type"]; ok && v != "" {
		m.MetricType = types.StringValue(v)
	} else {
		m.MetricType = types.StringNull()
	}
	if v, ok := obj["in-filter-chain"]; ok && v != "" {
		m.InFilterChain = types.StringValue(v)
	} else {
		m.InFilterChain = types.StringNull()
	}
	if v, ok := obj["areas-max"]; ok && v != "" {
		m.AreasMax = types.StringValue(v)
	} else {
		m.AreasMax = types.StringNull()
	}
	if v, ok := obj["areas"]; ok && v != "" {
		m.Areas = types.StringValue(v)
	} else {
		m.Areas = types.StringNull()
	}
	if v, ok := obj["afi"]; ok && v != "" {
		m.Afi = types.StringValue(v)
	} else {
		m.Afi = types.StringNull()
	}
	if v, ok := obj["l2.redistribute"]; ok && v != "" {
		m.L2Redistribute = types.StringValue(v)
	} else {
		m.L2Redistribute = types.StringNull()
	}
	if v, ok := obj["l2.out-filter-select"]; ok && v != "" {
		m.L2OutFilterSelect = types.StringValue(v)
	} else {
		m.L2OutFilterSelect = types.StringNull()
	}
	if v, ok := obj["l2.out-filter-chain"]; ok && v != "" {
		m.L2OutFilterChain = types.StringValue(v)
	} else {
		m.L2OutFilterChain = types.StringNull()
	}
	if v, ok := obj["l2.originate-default"]; ok && v != "" {
		m.L2OriginateDefault = types.StringValue(v)
	} else {
		m.L2OriginateDefault = types.StringNull()
	}
	if v, ok := obj["l2.lsp-update-interval"]; ok && v != "" {
		m.L2LspUpdateInterval = types.StringValue(v)
	} else {
		m.L2LspUpdateInterval = types.StringNull()
	}
	if v, ok := obj["l2.lsp-max-size"]; ok && v != "" {
		m.L2LspMaxSize = types.StringValue(v)
	} else {
		m.L2LspMaxSize = types.StringNull()
	}
	if v, ok := obj["l2.lsp-max-age"]; ok && v != "" {
		m.L2LspMaxAge = types.StringValue(v)
	} else {
		m.L2LspMaxAge = types.StringNull()
	}
	if v, ok := obj["l1.redistribute"]; ok && v != "" {
		m.L1Redistribute = types.StringValue(v)
	} else {
		m.L1Redistribute = types.StringNull()
	}
	if v, ok := obj["l1.out-filter-select"]; ok && v != "" {
		m.L1OutFilterSelect = types.StringValue(v)
	} else {
		m.L1OutFilterSelect = types.StringNull()
	}
	if v, ok := obj["l1.out-filter-chain"]; ok && v != "" {
		m.L1OutFilterChain = types.StringValue(v)
	} else {
		m.L1OutFilterChain = types.StringNull()
	}
	if v, ok := obj["l1.originate-default"]; ok && v != "" {
		m.L1OriginateDefault = types.StringValue(v)
	} else {
		m.L1OriginateDefault = types.StringNull()
	}
	if v, ok := obj["l1.lsp-update-interval"]; ok && v != "" {
		m.L1LspUpdateInterval = types.StringValue(v)
	} else {
		m.L1LspUpdateInterval = types.StringNull()
	}
	if v, ok := obj["l1.lsp-refresh-interval"]; ok && v != "" {
		m.L1LspRefreshInterval = types.StringValue(v)
	} else {
		m.L1LspRefreshInterval = types.StringNull()
	}
	if v, ok := obj["l1.lsp-max-size"]; ok && v != "" {
		m.L1LspMaxSize = types.StringValue(v)
	} else {
		m.L1LspMaxSize = types.StringNull()
	}
	if v, ok := obj["l1.lsp-max-age"]; ok && v != "" {
		m.L1LspMaxAge = types.StringValue(v)
	} else {
		m.L1LspMaxAge = types.StringNull()
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
}
