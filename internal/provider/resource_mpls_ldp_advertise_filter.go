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
	_ resource.Resource                = &MPLSLdpAdvertiseFilterResource{}
	_ resource.ResourceWithImportState = &MPLSLdpAdvertiseFilterResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type MPLSLdpAdvertiseFilterResource struct {
	reg *client.Registry
}

type MPLSLdpAdvertiseFilterModel struct {
	ID        types.String `tfsdk:"id"`
	Neighbor  types.String `tfsdk:"neighbor"`
	Advertise types.String `tfsdk:"advertise"`
	Comment   types.String `tfsdk:"comment"`
	Disabled  types.Bool   `tfsdk:"disabled"`
	Prefix    types.String `tfsdk:"prefix"`
	Vrf       types.String `tfsdk:"vrf"`
	Router    types.String `tfsdk:"router"`
}

func NewMPLSLdpAdvertiseFilterResource() resource.Resource { return &MPLSLdpAdvertiseFilterResource{} }

func (r *MPLSLdpAdvertiseFilterResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mpls_ldp_advertise_filter"
}

func (r *MPLSLdpAdvertiseFilterResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *MPLSLdpAdvertiseFilterResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/mpls/ldp/advertise-filter`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"neighbor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `neighbor`.",
			},
			"advertise": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `advertise`.",
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
			"prefix": schema.StringAttribute{
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

func (r *MPLSLdpAdvertiseFilterResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan MPLSLdpAdvertiseFilterModel
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
	if !(plan.Prefix.IsNull() || plan.Prefix.IsUnknown()) {
		body["prefix"] = plan.Prefix.ValueString()
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if !(plan.Advertise.IsNull() || plan.Advertise.IsUnknown()) {
		body["advertise"] = plan.Advertise.ValueString()
	}
	if !(plan.Neighbor.IsNull() || plan.Neighbor.IsUnknown()) {
		body["neighbor"] = plan.Neighbor.ValueString()
	}
	obj, err := c.Add(ctx, "/mpls/ldp/advertise-filter", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /mpls/ldp/advertise-filter failed", err.Error())
		return
	}
	mPLSLdpAdvertiseFilterApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MPLSLdpAdvertiseFilterResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MPLSLdpAdvertiseFilterModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/mpls/ldp/advertise-filter", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /mpls/ldp/advertise-filter failed", err.Error())
		return
	}
	mPLSLdpAdvertiseFilterApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *MPLSLdpAdvertiseFilterResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state MPLSLdpAdvertiseFilterModel
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
	if !plan.Prefix.Equal(state.Prefix) && !plan.Prefix.IsUnknown() {
		body["prefix"] = plan.Prefix.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) && !plan.Vrf.IsUnknown() {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if !plan.Advertise.Equal(state.Advertise) && !plan.Advertise.IsUnknown() {
		body["advertise"] = plan.Advertise.ValueString()
	}
	if !plan.Neighbor.Equal(state.Neighbor) && !plan.Neighbor.IsUnknown() {
		body["neighbor"] = plan.Neighbor.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/mpls/ldp/advertise-filter", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /mpls/ldp/advertise-filter failed", err.Error())
			return
		}
		mPLSLdpAdvertiseFilterApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MPLSLdpAdvertiseFilterResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state MPLSLdpAdvertiseFilterModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/mpls/ldp/advertise-filter", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /mpls/ldp/advertise-filter failed", err.Error())
	}
}

func (r *MPLSLdpAdvertiseFilterResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := mPLSLdpAdvertiseFilterLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /mpls/ldp/advertise-filter matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// mPLSLdpAdvertiseFilterLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func mPLSLdpAdvertiseFilterLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/mpls/ldp/advertise-filter", id)
}

func mPLSLdpAdvertiseFilterApply(ctx context.Context, obj client.Object, m *MPLSLdpAdvertiseFilterModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["neighbor"]; ok && v != "" {
		m.Neighbor = types.StringValue(v)
	} else {
		m.Neighbor = types.StringNull()
	}
	if v, ok := obj["advertise"]; ok && v != "" {
		m.Advertise = types.StringValue(v)
	} else {
		m.Advertise = types.StringNull()
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["prefix"]; ok {
		if v != "" {
			m.Prefix = types.StringValue(v)
		} else {
			m.Prefix = types.StringNull()
		}
	}
	if v, ok := obj["vrf"]; ok {
		if v != "" {
			m.Vrf = types.StringValue(v)
		} else {
			m.Vrf = types.StringNull()
		}
	}
}
