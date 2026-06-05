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
	_ resource.Resource                = &RoutingBGPVPNResource{}
	_ resource.ResourceWithImportState = &RoutingBGPVPNResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingBGPVPNResource struct {
	reg *client.Registry
}

type RoutingBGPVPNModel struct {
	ID                    types.String `tfsdk:"id"`
	Comment               types.String `tfsdk:"comment"`
	Disabled              types.Bool   `tfsdk:"disabled"`
	ExportFilter          types.String `tfsdk:"export_filter"`
	ExportRouteTargets    types.String `tfsdk:"export_route_targets"`
	ExportSelect          types.String `tfsdk:"export_select"`
	ImportFilter          types.String `tfsdk:"import_filter"`
	ImportRouteTargets    types.String `tfsdk:"import_route_targets"`
	Instance              types.String `tfsdk:"instance"`
	Invalid               types.Bool   `tfsdk:"invalid"`
	LabelAllocationPolicy types.String `tfsdk:"label_allocation_policy"`
	Redistribute          types.String `tfsdk:"redistribute"`
	RouteDistinguisher    types.String `tfsdk:"route_distinguisher"`
	Vrf                   types.String `tfsdk:"vrf"`
	Router                types.String `tfsdk:"router"`
}

func NewRoutingBGPVPNResource() resource.Resource { return &RoutingBGPVPNResource{} }

func (r *RoutingBGPVPNResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_bgp_vpn"
}

func (r *RoutingBGPVPNResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *RoutingBGPVPNResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Auto-test requires a typed-reference precondition (e.g. an existing peer,\ninstance, bridge of the specific kind). The current acc-test generator's\ngeneric data.routeros_interface.all lookup can't satisfy these. Use this\nresource manually with explicit references to a precondition resource\nin your config.\n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"export_filter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"export_route_targets": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"export_select": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"import_filter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"import_route_targets": schema.StringAttribute{
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
			"label_allocation_policy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "per-vrf", "per-prefix"}...)},
			},
			"redistribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"route_distinguisher": schema.StringAttribute{
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

func (r *RoutingBGPVPNResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingBGPVPNModel
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
	if !(plan.ExportFilter.IsNull() || plan.ExportFilter.IsUnknown()) {
		body["export-filter"] = plan.ExportFilter.ValueString()
	}
	if !(plan.ExportRouteTargets.IsNull() || plan.ExportRouteTargets.IsUnknown()) {
		body["export-route-targets"] = plan.ExportRouteTargets.ValueString()
	}
	if !(plan.ExportSelect.IsNull() || plan.ExportSelect.IsUnknown()) {
		body["export-select"] = plan.ExportSelect.ValueString()
	}
	if !(plan.ImportFilter.IsNull() || plan.ImportFilter.IsUnknown()) {
		body["import-filter"] = plan.ImportFilter.ValueString()
	}
	if !(plan.ImportRouteTargets.IsNull() || plan.ImportRouteTargets.IsUnknown()) {
		body["import-route-targets"] = plan.ImportRouteTargets.ValueString()
	}
	if !(plan.Instance.IsNull() || plan.Instance.IsUnknown()) {
		body["instance"] = plan.Instance.ValueString()
	}
	if !(plan.LabelAllocationPolicy.IsNull() || plan.LabelAllocationPolicy.IsUnknown()) {
		body["label-allocation-policy"] = plan.LabelAllocationPolicy.ValueString()
	}
	if !(plan.Redistribute.IsNull() || plan.Redistribute.IsUnknown()) {
		body["redistribute"] = plan.Redistribute.ValueString()
	}
	if !(plan.RouteDistinguisher.IsNull() || plan.RouteDistinguisher.IsUnknown()) {
		body["route-distinguisher"] = plan.RouteDistinguisher.ValueString()
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/bgp/vpn", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/bgp/vpn failed", err.Error())
		return
	}
	routingBGPVPNApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingBGPVPNResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingBGPVPNModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/bgp/vpn", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/bgp/vpn failed", err.Error())
		return
	}
	routingBGPVPNApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingBGPVPNResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingBGPVPNModel
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
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.ExportFilter.Equal(state.ExportFilter) {
		body["export-filter"] = plan.ExportFilter.ValueString()
	}
	if !plan.ExportRouteTargets.Equal(state.ExportRouteTargets) {
		body["export-route-targets"] = plan.ExportRouteTargets.ValueString()
	}
	if !plan.ExportSelect.Equal(state.ExportSelect) {
		body["export-select"] = plan.ExportSelect.ValueString()
	}
	if !plan.ImportFilter.Equal(state.ImportFilter) {
		body["import-filter"] = plan.ImportFilter.ValueString()
	}
	if !plan.ImportRouteTargets.Equal(state.ImportRouteTargets) {
		body["import-route-targets"] = plan.ImportRouteTargets.ValueString()
	}
	if !plan.Instance.Equal(state.Instance) {
		body["instance"] = plan.Instance.ValueString()
	}
	if !plan.LabelAllocationPolicy.Equal(state.LabelAllocationPolicy) {
		body["label-allocation-policy"] = plan.LabelAllocationPolicy.ValueString()
	}
	if !plan.Redistribute.Equal(state.Redistribute) {
		body["redistribute"] = plan.Redistribute.ValueString()
	}
	if !plan.RouteDistinguisher.Equal(state.RouteDistinguisher) {
		body["route-distinguisher"] = plan.RouteDistinguisher.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/bgp/vpn", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/bgp/vpn failed", err.Error())
			return
		}
		routingBGPVPNApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingBGPVPNResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingBGPVPNModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/bgp/vpn", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/bgp/vpn failed", err.Error())
	}
}

func (r *RoutingBGPVPNResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingBGPVPNLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/bgp/vpn matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingBGPVPNLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingBGPVPNLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/routing/bgp/vpn", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func routingBGPVPNApply(ctx context.Context, obj client.Object, m *RoutingBGPVPNModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
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
	if v, ok := obj["export-filter"]; ok {
		_ = v
		if v != "" {
			m.ExportFilter = types.StringValue(v)
		} else {
			m.ExportFilter = types.StringNull()
		}
	} else {
		m.ExportFilter = types.StringNull()
	}
	if v, ok := obj["export-route-targets"]; ok {
		_ = v
		if v != "" {
			m.ExportRouteTargets = types.StringValue(v)
		} else {
			m.ExportRouteTargets = types.StringNull()
		}
	} else {
		m.ExportRouteTargets = types.StringNull()
	}
	if v, ok := obj["export-select"]; ok {
		_ = v
		if v != "" {
			m.ExportSelect = types.StringValue(v)
		} else {
			m.ExportSelect = types.StringNull()
		}
	} else {
		m.ExportSelect = types.StringNull()
	}
	if v, ok := obj["import-filter"]; ok {
		_ = v
		if v != "" {
			m.ImportFilter = types.StringValue(v)
		} else {
			m.ImportFilter = types.StringNull()
		}
	} else {
		m.ImportFilter = types.StringNull()
	}
	if v, ok := obj["import-route-targets"]; ok {
		_ = v
		if v != "" {
			m.ImportRouteTargets = types.StringValue(v)
		} else {
			m.ImportRouteTargets = types.StringNull()
		}
	} else {
		m.ImportRouteTargets = types.StringNull()
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
	if v, ok := obj["label-allocation-policy"]; ok {
		_ = v
		if v != "" {
			m.LabelAllocationPolicy = types.StringValue(v)
		} else {
			m.LabelAllocationPolicy = types.StringNull()
		}
	} else {
		m.LabelAllocationPolicy = types.StringNull()
	}
	if v, ok := obj["redistribute"]; ok {
		_ = v
		if v != "" {
			m.Redistribute = types.StringValue(v)
		} else {
			m.Redistribute = types.StringNull()
		}
	} else {
		m.Redistribute = types.StringNull()
	}
	if v, ok := obj["route-distinguisher"]; ok {
		_ = v
		if v != "" {
			m.RouteDistinguisher = types.StringValue(v)
		} else {
			m.RouteDistinguisher = types.StringNull()
		}
	} else {
		m.RouteDistinguisher = types.StringNull()
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
