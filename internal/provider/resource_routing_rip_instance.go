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
	_ resource.Resource                = &RoutingRipInstanceResource{}
	_ resource.ResourceWithImportState = &RoutingRipInstanceResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingRipInstanceResource struct {
	reg *client.Registry
}

type RoutingRipInstanceModel struct {
	ID                 types.String `tfsdk:"id"`
	OutFilterSelect    types.String `tfsdk:"out_filter_select"`
	OutFilterChain     types.String `tfsdk:"out_filter_chain"`
	InFilterChain      types.String `tfsdk:"in_filter_chain"`
	Afi                types.String `tfsdk:"afi"`
	Disabled           types.Bool   `tfsdk:"disabled"`
	InputFilter        types.String `tfsdk:"input_filter"`
	Name               types.String `tfsdk:"name"`
	OriginateDefault   types.String `tfsdk:"originate_default"`
	OutputFilter       types.String `tfsdk:"output_filter"`
	Redistribute       types.String `tfsdk:"redistribute"`
	RouteGcTimeout     types.String `tfsdk:"route_gc_timeout"`
	RouteTimeout       types.String `tfsdk:"route_timeout"`
	RoutingTable       types.String `tfsdk:"routing_table"`
	SelectOutputFilter types.String `tfsdk:"select_output_filter"`
	UpdateInterval     types.String `tfsdk:"update_interval"`
	Vrf                types.String `tfsdk:"vrf"`
	Router             types.String `tfsdk:"router"`
}

func NewRoutingRipInstanceResource() resource.Resource { return &RoutingRipInstanceResource{} }

func (r *RoutingRipInstanceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_rip_instance"
}

func (r *RoutingRipInstanceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingRipInstanceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/routing/rip/instance`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"out_filter_select": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `out-filter-select`.",
			},
			"out_filter_chain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `out-filter-chain`.",
			},
			"in_filter_chain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `in-filter-chain`.",
			},
			"afi": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"input_filter": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"originate_default": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"output_filter": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"redistribute": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"route_gc_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"route_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"routing_table": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"select_output_filter": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"update_interval": schema.StringAttribute{
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

func (r *RoutingRipInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingRipInstanceModel
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
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.OriginateDefault.IsNull() || plan.OriginateDefault.IsUnknown()) {
		body["originate-default"] = plan.OriginateDefault.ValueString()
	}
	if !(plan.Redistribute.IsNull() || plan.Redistribute.IsUnknown()) {
		body["redistribute"] = plan.Redistribute.ValueString()
	}
	if !(plan.RouteGcTimeout.IsNull() || plan.RouteGcTimeout.IsUnknown()) {
		body["route-gc-timeout"] = plan.RouteGcTimeout.ValueString()
	}
	if !(plan.RouteTimeout.IsNull() || plan.RouteTimeout.IsUnknown()) {
		body["route-timeout"] = plan.RouteTimeout.ValueString()
	}
	if !(plan.RoutingTable.IsNull() || plan.RoutingTable.IsUnknown()) {
		body["routing-table"] = plan.RoutingTable.ValueString()
	}
	if !(plan.UpdateInterval.IsNull() || plan.UpdateInterval.IsUnknown()) {
		body["update-interval"] = plan.UpdateInterval.ValueString()
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if !(plan.InFilterChain.IsNull() || plan.InFilterChain.IsUnknown()) {
		body["in-filter-chain"] = plan.InFilterChain.ValueString()
	}
	if !(plan.OutFilterChain.IsNull() || plan.OutFilterChain.IsUnknown()) {
		body["out-filter-chain"] = plan.OutFilterChain.ValueString()
	}
	if !(plan.OutFilterSelect.IsNull() || plan.OutFilterSelect.IsUnknown()) {
		body["out-filter-select"] = plan.OutFilterSelect.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/rip/instance", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/rip/instance failed", err.Error())
		return
	}
	routingRipInstanceApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingRipInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingRipInstanceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/rip/instance", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/rip/instance failed", err.Error())
		return
	}
	routingRipInstanceApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingRipInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingRipInstanceModel
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
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.OriginateDefault.Equal(state.OriginateDefault) {
		body["originate-default"] = plan.OriginateDefault.ValueString()
	}
	if !plan.Redistribute.Equal(state.Redistribute) {
		body["redistribute"] = plan.Redistribute.ValueString()
	}
	if !plan.RouteGcTimeout.Equal(state.RouteGcTimeout) {
		body["route-gc-timeout"] = plan.RouteGcTimeout.ValueString()
	}
	if !plan.RouteTimeout.Equal(state.RouteTimeout) {
		body["route-timeout"] = plan.RouteTimeout.ValueString()
	}
	if !plan.RoutingTable.Equal(state.RoutingTable) {
		body["routing-table"] = plan.RoutingTable.ValueString()
	}
	if !plan.UpdateInterval.Equal(state.UpdateInterval) {
		body["update-interval"] = plan.UpdateInterval.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if !plan.InFilterChain.Equal(state.InFilterChain) && !plan.InFilterChain.IsUnknown() {
		body["in-filter-chain"] = plan.InFilterChain.ValueString()
	}
	if !plan.OutFilterChain.Equal(state.OutFilterChain) && !plan.OutFilterChain.IsUnknown() {
		body["out-filter-chain"] = plan.OutFilterChain.ValueString()
	}
	if !plan.OutFilterSelect.Equal(state.OutFilterSelect) && !plan.OutFilterSelect.IsUnknown() {
		body["out-filter-select"] = plan.OutFilterSelect.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/rip/instance", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/rip/instance failed", err.Error())
			return
		}
		routingRipInstanceApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingRipInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingRipInstanceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/rip/instance", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/rip/instance failed", err.Error())
	}
}

func (r *RoutingRipInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingRipInstanceLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/rip/instance matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingRipInstanceLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingRipInstanceLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/rip/instance", id)
}

func routingRipInstanceApply(ctx context.Context, obj client.Object, m *RoutingRipInstanceModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["out-filter-select"]; ok && v != "" {
		m.OutFilterSelect = types.StringValue(v)
	} else {
		m.OutFilterSelect = types.StringNull()
	}
	if v, ok := obj["out-filter-chain"]; ok && v != "" {
		m.OutFilterChain = types.StringValue(v)
	} else {
		m.OutFilterChain = types.StringNull()
	}
	if v, ok := obj["in-filter-chain"]; ok && v != "" {
		m.InFilterChain = types.StringValue(v)
	} else {
		m.InFilterChain = types.StringNull()
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
	if v, ok := obj["originate-default"]; ok {
		_ = v
		if v != "" {
			m.OriginateDefault = types.StringValue(v)
		} else {
			m.OriginateDefault = types.StringNull()
		}
	} else {
		m.OriginateDefault = types.StringNull()
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
	if v, ok := obj["route-gc-timeout"]; ok {
		_ = v
		if v != "" {
			m.RouteGcTimeout = types.StringValue(v)
		} else {
			m.RouteGcTimeout = types.StringNull()
		}
	} else {
		m.RouteGcTimeout = types.StringNull()
	}
	if v, ok := obj["route-timeout"]; ok {
		_ = v
		if v != "" {
			m.RouteTimeout = types.StringValue(v)
		} else {
			m.RouteTimeout = types.StringNull()
		}
	} else {
		m.RouteTimeout = types.StringNull()
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
	if v, ok := obj["select-output-filter"]; ok {
		_ = v
		if v != "" {
			m.SelectOutputFilter = types.StringValue(v)
		} else {
			m.SelectOutputFilter = types.StringNull()
		}
	} else {
		m.SelectOutputFilter = types.StringNull()
	}
	if v, ok := obj["update-interval"]; ok {
		_ = v
		if v != "" {
			m.UpdateInterval = types.StringValue(v)
		} else {
			m.UpdateInterval = types.StringNull()
		}
	} else {
		m.UpdateInterval = types.StringNull()
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
