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
	_ resource.Resource                = &MPLSLdpResource{}
	_ resource.ResourceWithImportState = &MPLSLdpResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type MPLSLdpResource struct {
	reg *client.Registry
}

type MPLSLdpModel struct {
	ID                   types.String `tfsdk:"id"`
	Vrf                  types.String `tfsdk:"vrf"`
	UseExplicitNull      types.String `tfsdk:"use_explicit_null"`
	TransportAddresses   types.String `tfsdk:"transport_addresses"`
	PreferredAfi         types.String `tfsdk:"preferred_afi"`
	PathVectorLimit      types.String `tfsdk:"path_vector_limit"`
	LsrId                types.String `tfsdk:"lsr_id"`
	LoopDetect           types.String `tfsdk:"loop_detect"`
	HopLimit             types.String `tfsdk:"hop_limit"`
	DistributeForDefault types.String `tfsdk:"distribute_for_default"`
	Afi                  types.String `tfsdk:"afi"`
	Comment              types.String `tfsdk:"comment"`
	Disabled             types.Bool   `tfsdk:"disabled"`
	Router               types.String `tfsdk:"router"`
}

func NewMPLSLdpResource() resource.Resource { return &MPLSLdpResource{} }

func (r *MPLSLdpResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mpls_ldp"
}

func (r *MPLSLdpResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *MPLSLdpResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "LDP allows one active instance per VRF; if one already exists this fails.",
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
			"use_explicit_null": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `use-explicit-null`.",
			},
			"transport_addresses": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `transport-addresses`.",
			},
			"preferred_afi": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `preferred-afi`.",
			},
			"path_vector_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `path-vector-limit`.",
			},
			"lsr_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `lsr-id`.",
			},
			"loop_detect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `loop-detect`.",
			},
			"hop_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hop-limit`.",
			},
			"distribute_for_default": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `distribute-for-default`.",
			},
			"afi": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `afi`.",
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

func (r *MPLSLdpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan MPLSLdpModel
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
	if !(plan.Afi.IsNull() || plan.Afi.IsUnknown()) {
		body["afi"] = plan.Afi.ValueString()
	}
	if !(plan.DistributeForDefault.IsNull() || plan.DistributeForDefault.IsUnknown()) {
		body["distribute-for-default"] = plan.DistributeForDefault.ValueString()
	}
	if !(plan.HopLimit.IsNull() || plan.HopLimit.IsUnknown()) {
		body["hop-limit"] = plan.HopLimit.ValueString()
	}
	if !(plan.LoopDetect.IsNull() || plan.LoopDetect.IsUnknown()) {
		body["loop-detect"] = plan.LoopDetect.ValueString()
	}
	if !(plan.LsrId.IsNull() || plan.LsrId.IsUnknown()) {
		body["lsr-id"] = plan.LsrId.ValueString()
	}
	if !(plan.PathVectorLimit.IsNull() || plan.PathVectorLimit.IsUnknown()) {
		body["path-vector-limit"] = plan.PathVectorLimit.ValueString()
	}
	if !(plan.PreferredAfi.IsNull() || plan.PreferredAfi.IsUnknown()) {
		body["preferred-afi"] = plan.PreferredAfi.ValueString()
	}
	if !(plan.TransportAddresses.IsNull() || plan.TransportAddresses.IsUnknown()) {
		body["transport-addresses"] = plan.TransportAddresses.ValueString()
	}
	if !(plan.UseExplicitNull.IsNull() || plan.UseExplicitNull.IsUnknown()) {
		body["use-explicit-null"] = plan.UseExplicitNull.ValueString()
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	obj, err := c.Add(ctx, "/mpls/ldp", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /mpls/ldp failed", err.Error())
		return
	}
	mPLSLdpApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MPLSLdpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MPLSLdpModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/mpls/ldp", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /mpls/ldp failed", err.Error())
		return
	}
	mPLSLdpApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *MPLSLdpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state MPLSLdpModel
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
	if !plan.Afi.Equal(state.Afi) && !plan.Afi.IsUnknown() {
		body["afi"] = plan.Afi.ValueString()
	}
	if !plan.DistributeForDefault.Equal(state.DistributeForDefault) && !plan.DistributeForDefault.IsUnknown() {
		body["distribute-for-default"] = plan.DistributeForDefault.ValueString()
	}
	if !plan.HopLimit.Equal(state.HopLimit) && !plan.HopLimit.IsUnknown() {
		body["hop-limit"] = plan.HopLimit.ValueString()
	}
	if !plan.LoopDetect.Equal(state.LoopDetect) && !plan.LoopDetect.IsUnknown() {
		body["loop-detect"] = plan.LoopDetect.ValueString()
	}
	if !plan.LsrId.Equal(state.LsrId) && !plan.LsrId.IsUnknown() {
		body["lsr-id"] = plan.LsrId.ValueString()
	}
	if !plan.PathVectorLimit.Equal(state.PathVectorLimit) && !plan.PathVectorLimit.IsUnknown() {
		body["path-vector-limit"] = plan.PathVectorLimit.ValueString()
	}
	if !plan.PreferredAfi.Equal(state.PreferredAfi) && !plan.PreferredAfi.IsUnknown() {
		body["preferred-afi"] = plan.PreferredAfi.ValueString()
	}
	if !plan.TransportAddresses.Equal(state.TransportAddresses) && !plan.TransportAddresses.IsUnknown() {
		body["transport-addresses"] = plan.TransportAddresses.ValueString()
	}
	if !plan.UseExplicitNull.Equal(state.UseExplicitNull) && !plan.UseExplicitNull.IsUnknown() {
		body["use-explicit-null"] = plan.UseExplicitNull.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) && !plan.Vrf.IsUnknown() {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/mpls/ldp", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /mpls/ldp failed", err.Error())
			return
		}
		mPLSLdpApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MPLSLdpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state MPLSLdpModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/mpls/ldp", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /mpls/ldp failed", err.Error())
	}
}

func (r *MPLSLdpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := mPLSLdpLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /mpls/ldp matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// mPLSLdpLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func mPLSLdpLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/mpls/ldp", id)
}

func mPLSLdpApply(ctx context.Context, obj client.Object, m *MPLSLdpModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["vrf"]; ok && v != "" {
		m.Vrf = types.StringValue(v)
	} else {
		m.Vrf = types.StringNull()
	}
	if v, ok := obj["use-explicit-null"]; ok && v != "" {
		m.UseExplicitNull = types.StringValue(v)
	} else {
		m.UseExplicitNull = types.StringNull()
	}
	if v, ok := obj["transport-addresses"]; ok && v != "" {
		m.TransportAddresses = types.StringValue(v)
	} else {
		m.TransportAddresses = types.StringNull()
	}
	if v, ok := obj["preferred-afi"]; ok && v != "" {
		m.PreferredAfi = types.StringValue(v)
	} else {
		m.PreferredAfi = types.StringNull()
	}
	if v, ok := obj["path-vector-limit"]; ok && v != "" {
		m.PathVectorLimit = types.StringValue(v)
	} else {
		m.PathVectorLimit = types.StringNull()
	}
	if v, ok := obj["lsr-id"]; ok && v != "" {
		m.LsrId = types.StringValue(v)
	} else {
		m.LsrId = types.StringNull()
	}
	if v, ok := obj["loop-detect"]; ok && v != "" {
		m.LoopDetect = types.StringValue(v)
	} else {
		m.LoopDetect = types.StringNull()
	}
	if v, ok := obj["hop-limit"]; ok && v != "" {
		m.HopLimit = types.StringValue(v)
	} else {
		m.HopLimit = types.StringNull()
	}
	if v, ok := obj["distribute-for-default"]; ok && v != "" {
		m.DistributeForDefault = types.StringValue(v)
	} else {
		m.DistributeForDefault = types.StringNull()
	}
	if v, ok := obj["afi"]; ok && v != "" {
		m.Afi = types.StringValue(v)
	} else {
		m.Afi = types.StringNull()
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
		} else if strings.TrimSpace(v) == "" {
			m.Disabled = types.BoolValue(true)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
}
