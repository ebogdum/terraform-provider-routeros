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
	_ resource.Resource                = &IPIpsecProposalResource{}
	_ resource.ResourceWithImportState = &IPIpsecProposalResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPIpsecProposalResource struct {
	reg *client.Registry
}

type IPIpsecProposalModel struct {
	ID             types.String `tfsdk:"id"`
	AuthAlgorithms types.String `tfsdk:"auth_algorithms"`
	Comment        types.String `tfsdk:"comment"`
	Default        types.Bool   `tfsdk:"default"`
	Disabled       types.Bool   `tfsdk:"disabled"`
	EncAlgorithms  types.Set    `tfsdk:"enc_algorithms"`
	EncrAlgorithms types.String `tfsdk:"encr_algorithms"`
	Lifetime       types.String `tfsdk:"lifetime"`
	Name           types.String `tfsdk:"name"`
	PfsGroup       types.String `tfsdk:"pfs_group"`
	Router         types.String `tfsdk:"router"`
}

func NewIPIpsecProposalResource() resource.Resource { return &IPIpsecProposalResource{} }

func (r *IPIpsecProposalResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_ipsec_proposal"
}

func (r *IPIpsecProposalResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPIpsecProposalResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/ipsec/proposal`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"auth_algorithms": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"md5", "sha1", "null", "sha256", "sha512"}...)},
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
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"enc_algorithms": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "",
			},
			"encr_algorithms": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"lifetime": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pfs_group": schema.StringAttribute{
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

func (r *IPIpsecProposalResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPIpsecProposalModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AuthAlgorithms.IsNull() || plan.AuthAlgorithms.IsUnknown()) {
		body["auth-algorithms"] = plan.AuthAlgorithms.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.EncAlgorithms.IsNull() || plan.EncAlgorithms.IsUnknown()) {
		body["enc-algorithms"] = encodeStringSet(ctx, plan.EncAlgorithms, &resp.Diagnostics)
	}
	if !(plan.Lifetime.IsNull() || plan.Lifetime.IsUnknown()) {
		body["lifetime"] = plan.Lifetime.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.PfsGroup.IsNull() || plan.PfsGroup.IsUnknown()) {
		body["pfs-group"] = plan.PfsGroup.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/ipsec/proposal", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/ipsec/proposal failed", err.Error())
		return
	}
	iPIpsecProposalApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIpsecProposalResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPIpsecProposalModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/ipsec/proposal", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/ipsec/proposal failed", err.Error())
		return
	}
	iPIpsecProposalApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPIpsecProposalResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPIpsecProposalModel
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
	if !plan.AuthAlgorithms.Equal(state.AuthAlgorithms) && !plan.AuthAlgorithms.IsUnknown() {
		body["auth-algorithms"] = plan.AuthAlgorithms.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.EncAlgorithms.Equal(state.EncAlgorithms) && !plan.EncAlgorithms.IsUnknown() {
		body["enc-algorithms"] = encodeStringSet(ctx, plan.EncAlgorithms, &resp.Diagnostics)
	}
	if !plan.Lifetime.Equal(state.Lifetime) && !plan.Lifetime.IsUnknown() {
		body["lifetime"] = plan.Lifetime.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.PfsGroup.Equal(state.PfsGroup) && !plan.PfsGroup.IsUnknown() {
		body["pfs-group"] = plan.PfsGroup.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/ipsec/proposal", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/ipsec/proposal failed", err.Error())
			return
		}
		iPIpsecProposalApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIpsecProposalResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPIpsecProposalModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/ipsec/proposal", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/ipsec/proposal failed", err.Error())
	}
}

func (r *IPIpsecProposalResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPIpsecProposalLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/ipsec/proposal matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPIpsecProposalLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPIpsecProposalLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/ipsec/proposal", id)
}

func iPIpsecProposalApply(ctx context.Context, obj client.Object, m *IPIpsecProposalModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["auth-algorithms"]; ok {
		_ = v
		if v != "" {
			m.AuthAlgorithms = types.StringValue(v)
		} else {
			m.AuthAlgorithms = types.StringNull()
		}
	} else {
		m.AuthAlgorithms = types.StringNull()
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
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else {
			m.Default = types.BoolNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["enc-algorithms"]; ok {
		_ = v
		m.EncAlgorithms = decodeStringSet(ctx, v)
	} else {
		m.EncAlgorithms = types.SetNull(types.StringType)
	}
	if v, ok := obj["encr-algorithms"]; ok {
		_ = v
		if v != "" {
			m.EncrAlgorithms = types.StringValue(v)
		} else {
			m.EncrAlgorithms = types.StringNull()
		}
	} else {
		m.EncrAlgorithms = types.StringNull()
	}
	if v, ok := obj["lifetime"]; ok {
		_ = v
		if v != "" {
			m.Lifetime = types.StringValue(v)
		} else {
			m.Lifetime = types.StringNull()
		}
	} else {
		m.Lifetime = types.StringNull()
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
	if v, ok := obj["pfs-group"]; ok {
		_ = v
		if v != "" {
			m.PfsGroup = types.StringValue(v)
		} else {
			m.PfsGroup = types.StringNull()
		}
	} else {
		m.PfsGroup = types.StringNull()
	}
}
