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
	_ resource.Resource                = &RoutingOSPFInstanceResource{}
	_ resource.ResourceWithImportState = &RoutingOSPFInstanceResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingOSPFInstanceResource struct {
	reg *client.Registry
}

type RoutingOSPFInstanceModel struct {
	ID               types.String `tfsdk:"id"`
	UseDn            types.String `tfsdk:"use_dn"`
	OutFilterChain   types.String `tfsdk:"out_filter_chain"`
	InFilterChain    types.String `tfsdk:"in_filter_chain"`
	Comment          types.String `tfsdk:"comment"`
	Disabled         types.Bool   `tfsdk:"disabled"`
	DomainID         types.String `tfsdk:"domain_id"`
	DomainTag        types.String `tfsdk:"domain_tag"`
	InFilter         types.String `tfsdk:"in_filter"`
	Invalid          types.Bool   `tfsdk:"invalid"`
	MPLSTeAddress    types.String `tfsdk:"mpls_te_address"`
	MPLSTeArea       types.String `tfsdk:"mpls_te_area"`
	Name             types.String `tfsdk:"name"`
	OriginateDefault types.String `tfsdk:"originate_default"`
	OutFilter        types.String `tfsdk:"out_filter"`
	OutFilterSelect  types.String `tfsdk:"out_filter_select"`
	Redistribute     types.String `tfsdk:"redistribute"`
	RouterID         types.String `tfsdk:"router_id"`
	RoutingTable     types.String `tfsdk:"routing_table"`
	Version          types.String `tfsdk:"version"`
	Vrf              types.String `tfsdk:"vrf"`
	Router           types.String `tfsdk:"router"`
}

func NewRoutingOSPFInstanceResource() resource.Resource { return &RoutingOSPFInstanceResource{} }

func (r *RoutingOSPFInstanceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_ospf_instance"
}

func (r *RoutingOSPFInstanceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingOSPFInstanceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/routing/ospf/instance`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"use_dn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `use-dn`.",
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
			"domain_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"domain_tag": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"in_filter": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"mpls_te_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mpls_te_area": schema.StringAttribute{
				Optional:    true,
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
			"out_filter": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"out_filter_select": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"redistribute": schema.StringAttribute{
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
			"version": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"2", "3"}...)},
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

func (r *RoutingOSPFInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingOSPFInstanceModel
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
	if !(plan.DomainID.IsNull() || plan.DomainID.IsUnknown()) {
		body["domain-id"] = plan.DomainID.ValueString()
	}
	if !(plan.DomainTag.IsNull() || plan.DomainTag.IsUnknown()) {
		body["domain-tag"] = plan.DomainTag.ValueString()
	}
	if !(plan.MPLSTeAddress.IsNull() || plan.MPLSTeAddress.IsUnknown()) {
		body["mpls-te-address"] = plan.MPLSTeAddress.ValueString()
	}
	if !(plan.MPLSTeArea.IsNull() || plan.MPLSTeArea.IsUnknown()) {
		body["mpls-te-area"] = plan.MPLSTeArea.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.OriginateDefault.IsNull() || plan.OriginateDefault.IsUnknown()) {
		body["originate-default"] = plan.OriginateDefault.ValueString()
	}
	if !(plan.OutFilterSelect.IsNull() || plan.OutFilterSelect.IsUnknown()) {
		body["out-filter-select"] = plan.OutFilterSelect.ValueString()
	}
	if !(plan.Redistribute.IsNull() || plan.Redistribute.IsUnknown()) {
		body["redistribute"] = plan.Redistribute.ValueString()
	}
	if !(plan.RouterID.IsNull() || plan.RouterID.IsUnknown()) {
		body["router-id"] = plan.RouterID.ValueString()
	}
	if !(plan.RoutingTable.IsNull() || plan.RoutingTable.IsUnknown()) {
		body["routing-table"] = plan.RoutingTable.ValueString()
	}
	if !(plan.Version.IsNull() || plan.Version.IsUnknown()) {
		body["version"] = plan.Version.ValueString()
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
	if !(plan.UseDn.IsNull() || plan.UseDn.IsUnknown()) {
		body["use-dn"] = plan.UseDn.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/ospf/instance", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/ospf/instance failed", err.Error())
		return
	}
	routingOSPFInstanceApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingOSPFInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingOSPFInstanceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/ospf/instance", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/ospf/instance failed", err.Error())
		return
	}
	routingOSPFInstanceApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingOSPFInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingOSPFInstanceModel
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
	if !plan.DomainID.Equal(state.DomainID) && !plan.DomainID.IsUnknown() {
		body["domain-id"] = plan.DomainID.ValueString()
	}
	if !plan.DomainTag.Equal(state.DomainTag) && !plan.DomainTag.IsUnknown() {
		body["domain-tag"] = plan.DomainTag.ValueString()
	}
	if !plan.MPLSTeAddress.Equal(state.MPLSTeAddress) && !plan.MPLSTeAddress.IsUnknown() {
		body["mpls-te-address"] = plan.MPLSTeAddress.ValueString()
	}
	if !plan.MPLSTeArea.Equal(state.MPLSTeArea) && !plan.MPLSTeArea.IsUnknown() {
		body["mpls-te-area"] = plan.MPLSTeArea.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.OriginateDefault.Equal(state.OriginateDefault) && !plan.OriginateDefault.IsUnknown() {
		body["originate-default"] = plan.OriginateDefault.ValueString()
	}
	if !plan.OutFilterSelect.Equal(state.OutFilterSelect) && !plan.OutFilterSelect.IsUnknown() {
		body["out-filter-select"] = plan.OutFilterSelect.ValueString()
	}
	if !plan.Redistribute.Equal(state.Redistribute) && !plan.Redistribute.IsUnknown() {
		body["redistribute"] = plan.Redistribute.ValueString()
	}
	if !plan.RouterID.Equal(state.RouterID) && !plan.RouterID.IsUnknown() {
		body["router-id"] = plan.RouterID.ValueString()
	}
	if !plan.RoutingTable.Equal(state.RoutingTable) && !plan.RoutingTable.IsUnknown() {
		body["routing-table"] = plan.RoutingTable.ValueString()
	}
	if !plan.Version.Equal(state.Version) && !plan.Version.IsUnknown() {
		body["version"] = plan.Version.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) && !plan.Vrf.IsUnknown() {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if !plan.InFilterChain.Equal(state.InFilterChain) && !plan.InFilterChain.IsUnknown() {
		body["in-filter-chain"] = plan.InFilterChain.ValueString()
	}
	if !plan.OutFilterChain.Equal(state.OutFilterChain) && !plan.OutFilterChain.IsUnknown() {
		body["out-filter-chain"] = plan.OutFilterChain.ValueString()
	}
	if !plan.UseDn.Equal(state.UseDn) && !plan.UseDn.IsUnknown() {
		body["use-dn"] = plan.UseDn.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/ospf/instance", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/ospf/instance failed", err.Error())
			return
		}
		routingOSPFInstanceApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingOSPFInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingOSPFInstanceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/ospf/instance", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/ospf/instance failed", err.Error())
	}
}

func (r *RoutingOSPFInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingOSPFInstanceLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/ospf/instance matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingOSPFInstanceLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingOSPFInstanceLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/ospf/instance", id)
}

func routingOSPFInstanceApply(ctx context.Context, obj client.Object, m *RoutingOSPFInstanceModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["use-dn"]; ok && v != "" {
		m.UseDn = types.StringValue(v)
	} else {
		m.UseDn = types.StringNull()
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
	if v, ok := obj["domain-id"]; ok {
		_ = v
		if v != "" {
			m.DomainID = types.StringValue(v)
		} else {
			m.DomainID = types.StringNull()
		}
	} else {
		m.DomainID = types.StringNull()
	}
	if v, ok := obj["domain-tag"]; ok {
		_ = v
		if v != "" {
			m.DomainTag = types.StringValue(v)
		} else {
			m.DomainTag = types.StringNull()
		}
	} else {
		m.DomainTag = types.StringNull()
	}
	if v, ok := obj["in-filter"]; ok {
		_ = v
		if v != "" {
			m.InFilter = types.StringValue(v)
		} else {
			m.InFilter = types.StringNull()
		}
	} else {
		m.InFilter = types.StringNull()
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
	if v, ok := obj["mpls-te-address"]; ok {
		_ = v
		if v != "" {
			m.MPLSTeAddress = types.StringValue(v)
		} else {
			m.MPLSTeAddress = types.StringNull()
		}
	} else {
		m.MPLSTeAddress = types.StringNull()
	}
	if v, ok := obj["mpls-te-area"]; ok {
		_ = v
		if v != "" {
			m.MPLSTeArea = types.StringValue(v)
		} else {
			m.MPLSTeArea = types.StringNull()
		}
	} else {
		m.MPLSTeArea = types.StringNull()
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
	if v, ok := obj["out-filter"]; ok {
		_ = v
		if v != "" {
			m.OutFilter = types.StringValue(v)
		} else {
			m.OutFilter = types.StringNull()
		}
	} else {
		m.OutFilter = types.StringNull()
	}
	if v, ok := obj["out-filter-select"]; ok {
		_ = v
		if v != "" {
			m.OutFilterSelect = types.StringValue(v)
		} else {
			m.OutFilterSelect = types.StringNull()
		}
	} else {
		m.OutFilterSelect = types.StringNull()
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
	if v, ok := obj["version"]; ok {
		_ = v
		if v != "" {
			m.Version = types.StringValue(v)
		} else {
			m.Version = types.StringNull()
		}
	} else {
		m.Version = types.StringNull()
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
