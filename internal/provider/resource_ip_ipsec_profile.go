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
	_ resource.Resource                = &IPIpsecProfileResource{}
	_ resource.ResourceWithImportState = &IPIpsecProfileResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPIpsecProfileResource struct {
	reg *client.Registry
}

type IPIpsecProfileModel struct {
	ID                  types.String `tfsdk:"id"`
	PrfAlgorithm        types.String `tfsdk:"prf_algorithm"`
	Default             types.Bool   `tfsdk:"default"`
	DhGroup             types.Set    `tfsdk:"dh_group"`
	DpdInterval         types.String `tfsdk:"dpd_interval"`
	DpdMaximumFailures  types.Int64  `tfsdk:"dpd_maximum_failures"`
	EncAlgorithm        types.Set    `tfsdk:"enc_algorithm"`
	EncryptionAlgorithm types.String `tfsdk:"encryption_algorithm"`
	HashAlgorithm       types.String `tfsdk:"hash_algorithm"`
	HashAlgorithms      types.String `tfsdk:"hash_algorithms"`
	Lifebytes           types.Int64  `tfsdk:"lifebytes"`
	Lifetime            types.String `tfsdk:"lifetime"`
	Name                types.String `tfsdk:"name"`
	NATTraversal        types.Bool   `tfsdk:"nat_traversal"`
	Ppk                 types.String `tfsdk:"ppk"`
	PrfAlgorithms       types.String `tfsdk:"prf_algorithms"`
	ProposalCheck       types.String `tfsdk:"proposal_check"`
	Router              types.String `tfsdk:"router"`
}

func NewIPIpsecProfileResource() resource.Resource { return &IPIpsecProfileResource{} }

func (r *IPIpsecProfileResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_ipsec_profile"
}

func (r *IPIpsecProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPIpsecProfileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/ipsec/profile`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"prf_algorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `prf-algorithm`.",
			},
			"default": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"dh_group": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "",
			},
			"dpd_interval": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationOrKeyword("disable-dpd")},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDurationExcept("disable-dpd")},
			},
			"dpd_maximum_failures": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"enc_algorithm": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "",
			},
			"encryption_algorithm": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"hash_algorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hash_algorithms": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"md5", "sha1", "sha256", "sha384", "sha512"}...)},
			},
			"lifebytes": schema.Int64Attribute{
				Optional:    true,
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
				Required:    true,
				Description: "",
			},
			"nat_traversal": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ppk": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "psk", "qkd", "psk-ike-initial"}...)},
			},
			"prf_algorithms": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"auto", "sha1", "sha256", "sha384", "sha512"}...)},
			},
			"proposal_check": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "obey", "strict", "claim", "exact"}...)},
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPIpsecProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPIpsecProfileModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.DhGroup.IsNull() || plan.DhGroup.IsUnknown()) {
		body["dh-group"] = encodeStringSet(ctx, plan.DhGroup, &resp.Diagnostics)
	}
	if !(plan.DpdInterval.IsNull() || plan.DpdInterval.IsUnknown()) {
		body["dpd-interval"] = plan.DpdInterval.ValueString()
	}
	if !(plan.DpdMaximumFailures.IsNull() || plan.DpdMaximumFailures.IsUnknown()) {
		body["dpd-maximum-failures"] = client.FormatInt64(plan.DpdMaximumFailures.ValueInt64())
	}
	if !(plan.EncAlgorithm.IsNull() || plan.EncAlgorithm.IsUnknown()) {
		body["enc-algorithm"] = encodeStringSet(ctx, plan.EncAlgorithm, &resp.Diagnostics)
	}
	if !(plan.HashAlgorithm.IsNull() || plan.HashAlgorithm.IsUnknown()) {
		body["hash-algorithm"] = plan.HashAlgorithm.ValueString()
	}
	if !(plan.Lifebytes.IsNull() || plan.Lifebytes.IsUnknown()) {
		body["lifebytes"] = client.FormatInt64(plan.Lifebytes.ValueInt64())
	}
	if !(plan.Lifetime.IsNull() || plan.Lifetime.IsUnknown()) {
		body["lifetime"] = plan.Lifetime.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.NATTraversal.IsNull() || plan.NATTraversal.IsUnknown()) {
		body["nat-traversal"] = client.FormatBool(plan.NATTraversal.ValueBool())
	}
	if !(plan.Ppk.IsNull() || plan.Ppk.IsUnknown()) {
		body["ppk"] = plan.Ppk.ValueString()
	}
	if !(plan.ProposalCheck.IsNull() || plan.ProposalCheck.IsUnknown()) {
		body["proposal-check"] = plan.ProposalCheck.ValueString()
	}
	if !(plan.PrfAlgorithm.IsNull() || plan.PrfAlgorithm.IsUnknown()) {
		body["prf-algorithm"] = plan.PrfAlgorithm.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/ipsec/profile", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/ipsec/profile failed", err.Error())
		return
	}
	iPIpsecProfileApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIpsecProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPIpsecProfileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/ipsec/profile", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/ipsec/profile failed", err.Error())
		return
	}
	iPIpsecProfileApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPIpsecProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPIpsecProfileModel
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
	if !plan.DhGroup.Equal(state.DhGroup) && !plan.DhGroup.IsUnknown() {
		body["dh-group"] = encodeStringSet(ctx, plan.DhGroup, &resp.Diagnostics)
	}
	if !plan.DpdInterval.Equal(state.DpdInterval) && !plan.DpdInterval.IsUnknown() {
		body["dpd-interval"] = plan.DpdInterval.ValueString()
	}
	if !plan.DpdMaximumFailures.Equal(state.DpdMaximumFailures) && !plan.DpdMaximumFailures.IsUnknown() {
		body["dpd-maximum-failures"] = client.FormatInt64(plan.DpdMaximumFailures.ValueInt64())
	}
	if !plan.EncAlgorithm.Equal(state.EncAlgorithm) && !plan.EncAlgorithm.IsUnknown() {
		body["enc-algorithm"] = encodeStringSet(ctx, plan.EncAlgorithm, &resp.Diagnostics)
	}
	if !plan.HashAlgorithm.Equal(state.HashAlgorithm) && !plan.HashAlgorithm.IsUnknown() {
		body["hash-algorithm"] = plan.HashAlgorithm.ValueString()
	}
	if !plan.Lifebytes.Equal(state.Lifebytes) && !plan.Lifebytes.IsUnknown() {
		body["lifebytes"] = client.FormatInt64(plan.Lifebytes.ValueInt64())
	}
	if !plan.Lifetime.Equal(state.Lifetime) && !plan.Lifetime.IsUnknown() {
		body["lifetime"] = plan.Lifetime.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.NATTraversal.Equal(state.NATTraversal) && !plan.NATTraversal.IsUnknown() {
		body["nat-traversal"] = client.FormatBool(plan.NATTraversal.ValueBool())
	}
	if !plan.Ppk.Equal(state.Ppk) && !plan.Ppk.IsUnknown() {
		body["ppk"] = plan.Ppk.ValueString()
	}
	if !plan.ProposalCheck.Equal(state.ProposalCheck) && !plan.ProposalCheck.IsUnknown() {
		body["proposal-check"] = plan.ProposalCheck.ValueString()
	}
	if !plan.PrfAlgorithm.Equal(state.PrfAlgorithm) && !plan.PrfAlgorithm.IsUnknown() {
		body["prf-algorithm"] = plan.PrfAlgorithm.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/ipsec/profile", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/ipsec/profile failed", err.Error())
			return
		}
		iPIpsecProfileApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIpsecProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPIpsecProfileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/ipsec/profile", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/ipsec/profile failed", err.Error())
	}
}

func (r *IPIpsecProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPIpsecProfileLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/ipsec/profile matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPIpsecProfileLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPIpsecProfileLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/ipsec/profile", id)
}

func iPIpsecProfileApply(ctx context.Context, obj client.Object, m *IPIpsecProfileModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["prf-algorithm"]; ok && v != "" {
		m.PrfAlgorithm = types.StringValue(v)
	} else {
		m.PrfAlgorithm = types.StringNull()
	}
	if v, ok := obj["default"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Default = types.BoolValue(true)
		} else {
			m.Default = types.BoolNull()
		}
	}
	if v, ok := obj["dh-group"]; ok {
		_ = v
		m.DhGroup = decodeStringSet(ctx, v)
	} else {
		m.DhGroup = types.SetNull(types.StringType)
	}
	if v, ok := obj["dpd-interval"]; ok {
		if v != "" {
			m.DpdInterval = types.StringValue(v)
		} else {
			m.DpdInterval = types.StringNull()
		}
	}
	if v, ok := obj["dpd-maximum-failures"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.DpdMaximumFailures = types.Int64Value(n)
		} else {
			m.DpdMaximumFailures = types.Int64Null()
		}
	} else {
		m.DpdMaximumFailures = types.Int64Null()
	}
	if v, ok := obj["enc-algorithm"]; ok {
		_ = v
		m.EncAlgorithm = decodeStringSet(ctx, v)
	} else {
		m.EncAlgorithm = types.SetNull(types.StringType)
	}
	if v, ok := obj["encryption-algorithm"]; ok {
		if v != "" {
			m.EncryptionAlgorithm = types.StringValue(v)
		} else {
			m.EncryptionAlgorithm = types.StringNull()
		}
	}
	if v, ok := obj["hash-algorithm"]; ok {
		if v != "" {
			m.HashAlgorithm = types.StringValue(v)
		} else {
			m.HashAlgorithm = types.StringNull()
		}
	}
	if v, ok := obj["hash-algorithms"]; ok {
		if v != "" {
			m.HashAlgorithms = types.StringValue(v)
		} else {
			m.HashAlgorithms = types.StringNull()
		}
	}
	if v, ok := obj["lifebytes"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Lifebytes = types.Int64Value(n)
		} else {
			m.Lifebytes = types.Int64Null()
		}
	} else {
		m.Lifebytes = types.Int64Null()
	}
	if v, ok := obj["lifetime"]; ok {
		if v != "" {
			m.Lifetime = types.StringValue(v)
		} else {
			m.Lifetime = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["nat-traversal"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.NATTraversal = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.NATTraversal = types.BoolValue(true)
		} else {
			m.NATTraversal = types.BoolNull()
		}
	}
	if v, ok := obj["ppk"]; ok {
		if v != "" {
			m.Ppk = types.StringValue(v)
		} else {
			m.Ppk = types.StringNull()
		}
	}
	if v, ok := obj["prf-algorithms"]; ok {
		if v != "" {
			m.PrfAlgorithms = types.StringValue(v)
		} else {
			m.PrfAlgorithms = types.StringNull()
		}
	}
	if v, ok := obj["proposal-check"]; ok {
		if v != "" {
			m.ProposalCheck = types.StringValue(v)
		} else {
			m.ProposalCheck = types.StringNull()
		}
	}
}
