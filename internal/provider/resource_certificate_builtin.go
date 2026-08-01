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
	_ resource.Resource                = &CertificateBuiltinResource{}
	_ resource.ResourceWithImportState = &CertificateBuiltinResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type CertificateBuiltinResource struct {
	reg *client.Registry
}

type CertificateBuiltinModel struct {
	ID             types.String `tfsdk:"id"`
	KeySize        types.String `tfsdk:"key_size"`
	Akid           types.String `tfsdk:"akid"`
	Comment        types.String `tfsdk:"comment"`
	CommonName     types.String `tfsdk:"common_name"`
	Country        types.String `tfsdk:"country"`
	DaysValid      types.Int64  `tfsdk:"days_valid"`
	Disabled       types.Bool   `tfsdk:"disabled"`
	InvalidAfter   types.String `tfsdk:"invalid_after"`
	InvalidBefore  types.String `tfsdk:"invalid_before"`
	Issuer         types.String `tfsdk:"issuer"`
	KeyType        types.String `tfsdk:"key_type"`
	KeyUsage       types.Set    `tfsdk:"key_usage"`
	Locality       types.String `tfsdk:"locality"`
	Organization   types.String `tfsdk:"organization"`
	SerialNumber   types.String `tfsdk:"serial_number"`
	Skid           types.String `tfsdk:"skid"`
	State          types.String `tfsdk:"state"`
	SubjectAltName types.String `tfsdk:"subject_alt_name"`
	Unit           types.String `tfsdk:"unit"`
	Router         types.String `tfsdk:"router"`
}

func NewCertificateBuiltinResource() resource.Resource { return &CertificateBuiltinResource{} }

func (r *CertificateBuiltinResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate_builtin"
}

func (r *CertificateBuiltinResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *CertificateBuiltinResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "System-generated certificates — read-only.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"key_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `key-size`.",
			},
			"akid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"common_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"country": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"days_valid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"invalid_after": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"invalid_before": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"issuer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"key_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"key_usage": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "",
			},
			"locality": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"organization": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"serial_number": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"skid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"subject_alt_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"unit": schema.StringAttribute{
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

func (r *CertificateBuiltinResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CertificateBuiltinModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Akid.IsNull() || plan.Akid.IsUnknown()) {
		body["akid"] = plan.Akid.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.CommonName.IsNull() || plan.CommonName.IsUnknown()) {
		body["common-name"] = plan.CommonName.ValueString()
	}
	if !(plan.Country.IsNull() || plan.Country.IsUnknown()) {
		body["country"] = plan.Country.ValueString()
	}
	if !(plan.DaysValid.IsNull() || plan.DaysValid.IsUnknown()) {
		body["days-valid"] = client.FormatInt64(plan.DaysValid.ValueInt64())
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.InvalidAfter.IsNull() || plan.InvalidAfter.IsUnknown()) {
		body["invalid-after"] = plan.InvalidAfter.ValueString()
	}
	if !(plan.InvalidBefore.IsNull() || plan.InvalidBefore.IsUnknown()) {
		body["invalid-before"] = plan.InvalidBefore.ValueString()
	}
	if !(plan.Issuer.IsNull() || plan.Issuer.IsUnknown()) {
		body["issuer"] = plan.Issuer.ValueString()
	}
	if !(plan.KeyType.IsNull() || plan.KeyType.IsUnknown()) {
		body["key-type"] = plan.KeyType.ValueString()
	}
	if !(plan.KeyUsage.IsNull() || plan.KeyUsage.IsUnknown()) {
		body["key-usage"] = encodeStringSet(ctx, plan.KeyUsage, &resp.Diagnostics)
	}
	if !(plan.Locality.IsNull() || plan.Locality.IsUnknown()) {
		body["locality"] = plan.Locality.ValueString()
	}
	if !(plan.Organization.IsNull() || plan.Organization.IsUnknown()) {
		body["organization"] = plan.Organization.ValueString()
	}
	if !(plan.SerialNumber.IsNull() || plan.SerialNumber.IsUnknown()) {
		body["serial-number"] = plan.SerialNumber.ValueString()
	}
	if !(plan.Skid.IsNull() || plan.Skid.IsUnknown()) {
		body["skid"] = plan.Skid.ValueString()
	}
	if !(plan.State.IsNull() || plan.State.IsUnknown()) {
		body["state"] = plan.State.ValueString()
	}
	if !(plan.SubjectAltName.IsNull() || plan.SubjectAltName.IsUnknown()) {
		body["subject-alt-name"] = plan.SubjectAltName.ValueString()
	}
	if !(plan.Unit.IsNull() || plan.Unit.IsUnknown()) {
		body["unit"] = plan.Unit.ValueString()
	}
	if !(plan.KeySize.IsNull() || plan.KeySize.IsUnknown()) {
		body["key-size"] = plan.KeySize.ValueString()
	}
	obj, err := c.Add(ctx, "/certificate/builtin", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /certificate/builtin failed", err.Error())
		return
	}
	certificateBuiltinApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CertificateBuiltinResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CertificateBuiltinModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/certificate/builtin", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /certificate/builtin failed", err.Error())
		return
	}
	certificateBuiltinApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CertificateBuiltinResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CertificateBuiltinModel
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
	if !plan.Akid.Equal(state.Akid) && !plan.Akid.IsUnknown() {
		body["akid"] = plan.Akid.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.CommonName.Equal(state.CommonName) && !plan.CommonName.IsUnknown() {
		body["common-name"] = plan.CommonName.ValueString()
	}
	if !plan.Country.Equal(state.Country) && !plan.Country.IsUnknown() {
		body["country"] = plan.Country.ValueString()
	}
	if !plan.DaysValid.Equal(state.DaysValid) && !plan.DaysValid.IsUnknown() {
		body["days-valid"] = client.FormatInt64(plan.DaysValid.ValueInt64())
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.InvalidAfter.Equal(state.InvalidAfter) && !plan.InvalidAfter.IsUnknown() {
		body["invalid-after"] = plan.InvalidAfter.ValueString()
	}
	if !plan.InvalidBefore.Equal(state.InvalidBefore) && !plan.InvalidBefore.IsUnknown() {
		body["invalid-before"] = plan.InvalidBefore.ValueString()
	}
	if !plan.Issuer.Equal(state.Issuer) && !plan.Issuer.IsUnknown() {
		body["issuer"] = plan.Issuer.ValueString()
	}
	if !plan.KeyType.Equal(state.KeyType) && !plan.KeyType.IsUnknown() {
		body["key-type"] = plan.KeyType.ValueString()
	}
	if !plan.KeyUsage.Equal(state.KeyUsage) && !plan.KeyUsage.IsUnknown() {
		body["key-usage"] = encodeStringSet(ctx, plan.KeyUsage, &resp.Diagnostics)
	}
	if !plan.Locality.Equal(state.Locality) && !plan.Locality.IsUnknown() {
		body["locality"] = plan.Locality.ValueString()
	}
	if !plan.Organization.Equal(state.Organization) && !plan.Organization.IsUnknown() {
		body["organization"] = plan.Organization.ValueString()
	}
	if !plan.SerialNumber.Equal(state.SerialNumber) && !plan.SerialNumber.IsUnknown() {
		body["serial-number"] = plan.SerialNumber.ValueString()
	}
	if !plan.Skid.Equal(state.Skid) && !plan.Skid.IsUnknown() {
		body["skid"] = plan.Skid.ValueString()
	}
	if !plan.State.Equal(state.State) && !plan.State.IsUnknown() {
		body["state"] = plan.State.ValueString()
	}
	if !plan.SubjectAltName.Equal(state.SubjectAltName) && !plan.SubjectAltName.IsUnknown() {
		body["subject-alt-name"] = plan.SubjectAltName.ValueString()
	}
	if !plan.Unit.Equal(state.Unit) && !plan.Unit.IsUnknown() {
		body["unit"] = plan.Unit.ValueString()
	}
	if !plan.KeySize.Equal(state.KeySize) && !plan.KeySize.IsUnknown() {
		body["key-size"] = plan.KeySize.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/certificate/builtin", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /certificate/builtin failed", err.Error())
			return
		}
		certificateBuiltinApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CertificateBuiltinResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CertificateBuiltinModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/certificate/builtin", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /certificate/builtin failed", err.Error())
	}
}

func (r *CertificateBuiltinResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := certificateBuiltinLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /certificate/builtin matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// certificateBuiltinLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func certificateBuiltinLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/certificate/builtin", id)
}

func certificateBuiltinApply(ctx context.Context, obj client.Object, m *CertificateBuiltinModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["key-size"]; ok && v != "" {
		m.KeySize = types.StringValue(v)
	} else {
		m.KeySize = types.StringNull()
	}
	if v, ok := obj["akid"]; ok {
		if v != "" {
			m.Akid = types.StringValue(v)
		} else {
			m.Akid = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["common-name"]; ok {
		if v != "" {
			m.CommonName = types.StringValue(v)
		} else {
			m.CommonName = types.StringNull()
		}
	}
	if v, ok := obj["country"]; ok {
		if v != "" {
			m.Country = types.StringValue(v)
		} else {
			m.Country = types.StringNull()
		}
	}
	if v, ok := obj["days-valid"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.DaysValid = types.Int64Value(n)
		} else {
			m.DaysValid = types.Int64Null()
		}
	} else {
		m.DaysValid = types.Int64Null()
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["invalid-after"]; ok {
		if v != "" {
			m.InvalidAfter = types.StringValue(v)
		} else {
			m.InvalidAfter = types.StringNull()
		}
	}
	if v, ok := obj["invalid-before"]; ok {
		if v != "" {
			m.InvalidBefore = types.StringValue(v)
		} else {
			m.InvalidBefore = types.StringNull()
		}
	}
	if v, ok := obj["issuer"]; ok {
		if v != "" {
			m.Issuer = types.StringValue(v)
		} else {
			m.Issuer = types.StringNull()
		}
	}
	if v, ok := obj["key-type"]; ok {
		if v != "" {
			m.KeyType = types.StringValue(v)
		} else {
			m.KeyType = types.StringNull()
		}
	}
	if v, ok := obj["key-usage"]; ok {
		_ = v
		m.KeyUsage = decodeStringSet(ctx, v)
	} else {
		m.KeyUsage = types.SetNull(types.StringType)
	}
	if v, ok := obj["locality"]; ok {
		if v != "" {
			m.Locality = types.StringValue(v)
		} else {
			m.Locality = types.StringNull()
		}
	}
	if v, ok := obj["organization"]; ok {
		if v != "" {
			m.Organization = types.StringValue(v)
		} else {
			m.Organization = types.StringNull()
		}
	}
	if v, ok := obj["serial-number"]; ok {
		if v != "" {
			m.SerialNumber = types.StringValue(v)
		} else {
			m.SerialNumber = types.StringNull()
		}
	}
	if v, ok := obj["skid"]; ok {
		if v != "" {
			m.Skid = types.StringValue(v)
		} else {
			m.Skid = types.StringNull()
		}
	}
	if v, ok := obj["state"]; ok {
		if v != "" {
			m.State = types.StringValue(v)
		} else {
			m.State = types.StringNull()
		}
	}
	if v, ok := obj["subject-alt-name"]; ok {
		if v != "" {
			m.SubjectAltName = types.StringValue(v)
		} else {
			m.SubjectAltName = types.StringNull()
		}
	}
	if v, ok := obj["unit"]; ok {
		if v != "" {
			m.Unit = types.StringValue(v)
		} else {
			m.Unit = types.StringNull()
		}
	}
}
