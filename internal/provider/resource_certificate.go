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
	_ resource.Resource                = &CertificateResource{}
	_ resource.ResourceWithImportState = &CertificateResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type CertificateResource struct {
	reg *client.Registry
}

type CertificateModel struct {
	ID                types.String `tfsdk:"id"`
	Acme              types.Bool   `tfsdk:"acme"`
	AcmeStatus        types.String `tfsdk:"acme_status"`
	AddAcme           types.String `tfsdk:"add_acme"`
	Akid              types.String `tfsdk:"akid"`
	Authority         types.Bool   `tfsdk:"authority"`
	CA                types.String `tfsdk:"ca"`
	CACrlHost         types.String `tfsdk:"ca_crl_host"`
	CAFingerprint     types.String `tfsdk:"ca_fingerprint"`
	CardReinstall     types.String `tfsdk:"card_reinstall"`
	CardVerify        types.String `tfsdk:"card_verify"`
	CommonName        types.String `tfsdk:"common_name"`
	Country           types.String `tfsdk:"country"`
	CreateCertRequest types.String `tfsdk:"create_cert_request"`
	Crl               types.Bool   `tfsdk:"crl"`
	DaysValid         types.Int64  `tfsdk:"days_valid"`
	DigestAlgorithm   types.String `tfsdk:"digest_algorithm"`
	Dynamic           types.Bool   `tfsdk:"dynamic"`
	Expired           types.Bool   `tfsdk:"expired"`
	ExpiresAfter      types.String `tfsdk:"expires_after"`
	Export            types.String `tfsdk:"export"`
	Fingerprint       types.String `tfsdk:"fingerprint"`
	HasAcmeStatus     types.String `tfsdk:"has_acme_status"`
	Import            types.String `tfsdk:"import"`
	InvalidAfter      types.String `tfsdk:"invalid_after"`
	InvalidBefore     types.String `tfsdk:"invalid_before"`
	Issued            types.Bool   `tfsdk:"issued"`
	Issuer            types.String `tfsdk:"issuer"`
	KeySize           types.String `tfsdk:"key_size"`
	KeyType           types.String `tfsdk:"key_type"`
	KeyUsage          types.String `tfsdk:"key_usage"`
	Locality          types.String `tfsdk:"locality"`
	Name              types.String `tfsdk:"name"`
	Notsealed         types.String `tfsdk:"notsealed"`
	Organization      types.String `tfsdk:"organization"`
	PrivateKey        types.Bool   `tfsdk:"private_key"`
	ReqFingerprint    types.String `tfsdk:"req_fingerprint"`
	Revoke            types.String `tfsdk:"revoke"`
	Revoked           types.Bool   `tfsdk:"revoked"`
	RevokedTime       types.String `tfsdk:"revoked_time"`
	ScepURL           types.String `tfsdk:"scep_url"`
	Sealed            types.String `tfsdk:"sealed"`
	SealedAndHide     types.String `tfsdk:"sealed_and_hide"`
	SerialNumber      types.String `tfsdk:"serial_number"`
	Sign              types.String `tfsdk:"sign"`
	SignViaScep       types.String `tfsdk:"sign_via_scep"`
	Skid              types.String `tfsdk:"skid"`
	SmartCardKey      types.Bool   `tfsdk:"smart_card_key"`
	State             types.String `tfsdk:"state"`
	SubjectAltName    types.String `tfsdk:"subject_alt_name"`
	TrustStore        types.String `tfsdk:"trust_store"`
	Trusted           types.Bool   `tfsdk:"trusted"`
	Type              types.Int64  `tfsdk:"type"`
	Unit              types.String `tfsdk:"unit"`
	Router            types.String `tfsdk:"router"`
}

func NewCertificateResource() resource.Resource { return &CertificateResource{} }

func (r *CertificateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate"
}

func (r *CertificateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *CertificateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Certificate template/store. RouterOS issues self-signed CAs and CA-signed\nleaf certs in two steps: declare the template via this resource, then\ntrigger signing via the /certificate/sign action.\n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"acme": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"acme_status": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"add_acme": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"akid": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"authority": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"ca": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"ca_crl_host": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"ca_fingerprint": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"card_reinstall": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"card_verify": schema.StringAttribute{
				Computed:    true,
				Description: "",
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
			"create_cert_request": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"crl": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"days_valid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"digest_algorithm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"md5", "sha1", "sha256", "sha384", "sha512"}...)},
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"expired": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"expires_after": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"export": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"fingerprint": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"has_acme_status": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"import": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"invalid_after": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"invalid_before": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"issued": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"issuer": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"key_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"prime256v1", "secp384r1", "secp521r1", "1024", "1536", "2048", "4096", "8192", "16384"}...)},
			},
			"key_type": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"rsa", "dsa", "ec"}...)},
			},
			"key_usage": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"locality": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"notsealed": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"organization": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"private_key": schema.BoolAttribute{
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"req_fingerprint": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"revoke": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"revoked": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"revoked_time": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"scep_url": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"sealed": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"sealed_and_hide": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"serial_number": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"sign": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"sign_via_scep": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"skid": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"smart_card_key": schema.BoolAttribute{
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
			"trust_store": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"trusted": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"type": schema.Int64Attribute{
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

func (r *CertificateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CertificateModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.CommonName.IsNull() || plan.CommonName.IsUnknown()) {
		body["common-name"] = plan.CommonName.ValueString()
	}
	if !(plan.Country.IsNull() || plan.Country.IsUnknown()) {
		body["country"] = plan.Country.ValueString()
	}
	if !(plan.DaysValid.IsNull() || plan.DaysValid.IsUnknown()) {
		body["days-valid"] = client.FormatInt64(plan.DaysValid.ValueInt64())
	}
	if !(plan.DigestAlgorithm.IsNull() || plan.DigestAlgorithm.IsUnknown()) {
		body["digest-algorithm"] = plan.DigestAlgorithm.ValueString()
	}
	if !(plan.KeySize.IsNull() || plan.KeySize.IsUnknown()) {
		body["key-size"] = plan.KeySize.ValueString()
	}
	if !(plan.KeyUsage.IsNull() || plan.KeyUsage.IsUnknown()) {
		body["key-usage"] = plan.KeyUsage.ValueString()
	}
	if !(plan.Locality.IsNull() || plan.Locality.IsUnknown()) {
		body["locality"] = plan.Locality.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Organization.IsNull() || plan.Organization.IsUnknown()) {
		body["organization"] = plan.Organization.ValueString()
	}
	if !(plan.State.IsNull() || plan.State.IsUnknown()) {
		body["state"] = plan.State.ValueString()
	}
	if !(plan.SubjectAltName.IsNull() || plan.SubjectAltName.IsUnknown()) {
		body["subject-alt-name"] = plan.SubjectAltName.ValueString()
	}
	if !(plan.TrustStore.IsNull() || plan.TrustStore.IsUnknown()) {
		body["trust-store"] = plan.TrustStore.ValueString()
	}
	if !(plan.Trusted.IsNull() || plan.Trusted.IsUnknown()) {
		body["trusted"] = client.FormatBool(plan.Trusted.ValueBool())
	}
	if !(plan.Unit.IsNull() || plan.Unit.IsUnknown()) {
		body["unit"] = plan.Unit.ValueString()
	}
	obj, err := c.Add(ctx, "/certificate", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /certificate failed", err.Error())
		return
	}
	certificateApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CertificateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CertificateModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/certificate", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /certificate failed", err.Error())
		return
	}
	certificateApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CertificateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CertificateModel
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
	if !plan.CommonName.Equal(state.CommonName) && !plan.CommonName.IsUnknown() {
		body["common-name"] = plan.CommonName.ValueString()
	}
	if !plan.Country.Equal(state.Country) && !plan.Country.IsUnknown() {
		body["country"] = plan.Country.ValueString()
	}
	if !plan.DaysValid.Equal(state.DaysValid) && !plan.DaysValid.IsUnknown() {
		body["days-valid"] = client.FormatInt64(plan.DaysValid.ValueInt64())
	}
	if !plan.DigestAlgorithm.Equal(state.DigestAlgorithm) && !plan.DigestAlgorithm.IsUnknown() {
		body["digest-algorithm"] = plan.DigestAlgorithm.ValueString()
	}
	if !plan.KeySize.Equal(state.KeySize) && !plan.KeySize.IsUnknown() {
		body["key-size"] = plan.KeySize.ValueString()
	}
	if !plan.KeyUsage.Equal(state.KeyUsage) && !plan.KeyUsage.IsUnknown() {
		body["key-usage"] = plan.KeyUsage.ValueString()
	}
	if !plan.Locality.Equal(state.Locality) && !plan.Locality.IsUnknown() {
		body["locality"] = plan.Locality.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Organization.Equal(state.Organization) && !plan.Organization.IsUnknown() {
		body["organization"] = plan.Organization.ValueString()
	}
	if !plan.State.Equal(state.State) && !plan.State.IsUnknown() {
		body["state"] = plan.State.ValueString()
	}
	if !plan.SubjectAltName.Equal(state.SubjectAltName) && !plan.SubjectAltName.IsUnknown() {
		body["subject-alt-name"] = plan.SubjectAltName.ValueString()
	}
	if !plan.TrustStore.Equal(state.TrustStore) && !plan.TrustStore.IsUnknown() {
		body["trust-store"] = plan.TrustStore.ValueString()
	}
	if !plan.Trusted.Equal(state.Trusted) && !plan.Trusted.IsUnknown() {
		body["trusted"] = client.FormatBool(plan.Trusted.ValueBool())
	}
	if !plan.Unit.Equal(state.Unit) && !plan.Unit.IsUnknown() {
		body["unit"] = plan.Unit.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/certificate", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /certificate failed", err.Error())
			return
		}
		certificateApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CertificateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CertificateModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/certificate", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /certificate failed", err.Error())
	}
}

func (r *CertificateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := certificateLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /certificate matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// certificateLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func certificateLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/certificate", id)
}

func certificateApply(ctx context.Context, obj client.Object, m *CertificateModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["acme"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Acme = types.BoolValue(b)
		} else {
			m.Acme = types.BoolNull()
		}
	} else {
		m.Acme = types.BoolNull()
	}
	if v, ok := obj["acme-status"]; ok {
		_ = v
		if v != "" {
			m.AcmeStatus = types.StringValue(v)
		} else {
			m.AcmeStatus = types.StringNull()
		}
	} else {
		m.AcmeStatus = types.StringNull()
	}
	if v, ok := obj["add-acme"]; ok {
		_ = v
		if v != "" {
			m.AddAcme = types.StringValue(v)
		} else {
			m.AddAcme = types.StringNull()
		}
	} else {
		m.AddAcme = types.StringNull()
	}
	if v, ok := obj["akid"]; ok {
		_ = v
		if v != "" {
			m.Akid = types.StringValue(v)
		} else {
			m.Akid = types.StringNull()
		}
	} else {
		m.Akid = types.StringNull()
	}
	if v, ok := obj["authority"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Authority = types.BoolValue(b)
		} else {
			m.Authority = types.BoolNull()
		}
	} else {
		m.Authority = types.BoolNull()
	}
	if v, ok := obj["ca"]; ok {
		_ = v
		if v != "" {
			m.CA = types.StringValue(v)
		} else {
			m.CA = types.StringNull()
		}
	} else {
		m.CA = types.StringNull()
	}
	if v, ok := obj["ca-crl-host"]; ok {
		_ = v
		if v != "" {
			m.CACrlHost = types.StringValue(v)
		} else {
			m.CACrlHost = types.StringNull()
		}
	} else {
		m.CACrlHost = types.StringNull()
	}
	if v, ok := obj["ca-fingerprint"]; ok {
		_ = v
		if v != "" {
			m.CAFingerprint = types.StringValue(v)
		} else {
			m.CAFingerprint = types.StringNull()
		}
	} else {
		m.CAFingerprint = types.StringNull()
	}
	if v, ok := obj["card-reinstall"]; ok {
		_ = v
		if v != "" {
			m.CardReinstall = types.StringValue(v)
		} else {
			m.CardReinstall = types.StringNull()
		}
	} else {
		m.CardReinstall = types.StringNull()
	}
	if v, ok := obj["card-verify"]; ok {
		_ = v
		if v != "" {
			m.CardVerify = types.StringValue(v)
		} else {
			m.CardVerify = types.StringNull()
		}
	} else {
		m.CardVerify = types.StringNull()
	}
	if v, ok := obj["common-name"]; ok {
		_ = v
		if v != "" {
			m.CommonName = types.StringValue(v)
		} else {
			m.CommonName = types.StringNull()
		}
	} else {
		m.CommonName = types.StringNull()
	}
	if v, ok := obj["country"]; ok {
		_ = v
		if v != "" {
			m.Country = types.StringValue(v)
		} else {
			m.Country = types.StringNull()
		}
	} else {
		m.Country = types.StringNull()
	}
	if v, ok := obj["create-cert-request"]; ok {
		_ = v
		if v != "" {
			m.CreateCertRequest = types.StringValue(v)
		} else {
			m.CreateCertRequest = types.StringNull()
		}
	} else {
		m.CreateCertRequest = types.StringNull()
	}
	if v, ok := obj["crl"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Crl = types.BoolValue(b)
		} else {
			m.Crl = types.BoolNull()
		}
	} else {
		m.Crl = types.BoolNull()
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
	if v, ok := obj["digest-algorithm"]; ok {
		_ = v
		if v != "" {
			m.DigestAlgorithm = types.StringValue(v)
		} else {
			m.DigestAlgorithm = types.StringNull()
		}
	} else {
		m.DigestAlgorithm = types.StringNull()
	}
	if v, ok := obj["dynamic"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else {
			m.Dynamic = types.BoolNull()
		}
	} else {
		m.Dynamic = types.BoolNull()
	}
	if v, ok := obj["expired"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Expired = types.BoolValue(b)
		} else {
			m.Expired = types.BoolNull()
		}
	} else {
		m.Expired = types.BoolNull()
	}
	if v, ok := obj["expires-after"]; ok {
		_ = v
		if v != "" {
			m.ExpiresAfter = types.StringValue(v)
		} else {
			m.ExpiresAfter = types.StringNull()
		}
	} else {
		m.ExpiresAfter = types.StringNull()
	}
	if v, ok := obj["export"]; ok {
		_ = v
		if v != "" {
			m.Export = types.StringValue(v)
		} else {
			m.Export = types.StringNull()
		}
	} else {
		m.Export = types.StringNull()
	}
	if v, ok := obj["fingerprint"]; ok {
		_ = v
		if v != "" {
			m.Fingerprint = types.StringValue(v)
		} else {
			m.Fingerprint = types.StringNull()
		}
	} else {
		m.Fingerprint = types.StringNull()
	}
	if v, ok := obj["has-acme-status"]; ok {
		_ = v
		if v != "" {
			m.HasAcmeStatus = types.StringValue(v)
		} else {
			m.HasAcmeStatus = types.StringNull()
		}
	} else {
		m.HasAcmeStatus = types.StringNull()
	}
	if v, ok := obj["import"]; ok {
		_ = v
		if v != "" {
			m.Import = types.StringValue(v)
		} else {
			m.Import = types.StringNull()
		}
	} else {
		m.Import = types.StringNull()
	}
	if v, ok := obj["invalid-after"]; ok {
		_ = v
		if v != "" {
			m.InvalidAfter = types.StringValue(v)
		} else {
			m.InvalidAfter = types.StringNull()
		}
	} else {
		m.InvalidAfter = types.StringNull()
	}
	if v, ok := obj["invalid-before"]; ok {
		_ = v
		if v != "" {
			m.InvalidBefore = types.StringValue(v)
		} else {
			m.InvalidBefore = types.StringNull()
		}
	} else {
		m.InvalidBefore = types.StringNull()
	}
	if v, ok := obj["issued"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Issued = types.BoolValue(b)
		} else {
			m.Issued = types.BoolNull()
		}
	} else {
		m.Issued = types.BoolNull()
	}
	if v, ok := obj["issuer"]; ok {
		_ = v
		if v != "" {
			m.Issuer = types.StringValue(v)
		} else {
			m.Issuer = types.StringNull()
		}
	} else {
		m.Issuer = types.StringNull()
	}
	if v, ok := obj["key-size"]; ok {
		_ = v
		if v != "" {
			m.KeySize = types.StringValue(v)
		} else {
			m.KeySize = types.StringNull()
		}
	} else {
		m.KeySize = types.StringNull()
	}
	if v, ok := obj["key-type"]; ok {
		_ = v
		if v != "" {
			m.KeyType = types.StringValue(v)
		} else {
			m.KeyType = types.StringNull()
		}
	} else {
		m.KeyType = types.StringNull()
	}
	if v, ok := obj["key-usage"]; ok {
		_ = v
		if v != "" {
			m.KeyUsage = types.StringValue(v)
		} else {
			m.KeyUsage = types.StringNull()
		}
	} else {
		m.KeyUsage = types.StringNull()
	}
	if v, ok := obj["locality"]; ok {
		_ = v
		if v != "" {
			m.Locality = types.StringValue(v)
		} else {
			m.Locality = types.StringNull()
		}
	} else {
		m.Locality = types.StringNull()
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
	if v, ok := obj["notsealed"]; ok {
		_ = v
		if v != "" {
			m.Notsealed = types.StringValue(v)
		} else {
			m.Notsealed = types.StringNull()
		}
	} else {
		m.Notsealed = types.StringNull()
	}
	if v, ok := obj["organization"]; ok {
		_ = v
		if v != "" {
			m.Organization = types.StringValue(v)
		} else {
			m.Organization = types.StringNull()
		}
	} else {
		m.Organization = types.StringNull()
	}
	if v, ok := obj["private-key"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.PrivateKey = types.BoolValue(b)
		} else {
			m.PrivateKey = types.BoolNull()
		}
	} else {
		m.PrivateKey = types.BoolNull()
	}
	if v, ok := obj["req-fingerprint"]; ok {
		_ = v
		if v != "" {
			m.ReqFingerprint = types.StringValue(v)
		} else {
			m.ReqFingerprint = types.StringNull()
		}
	} else {
		m.ReqFingerprint = types.StringNull()
	}
	if v, ok := obj["revoke"]; ok {
		_ = v
		if v != "" {
			m.Revoke = types.StringValue(v)
		} else {
			m.Revoke = types.StringNull()
		}
	} else {
		m.Revoke = types.StringNull()
	}
	if v, ok := obj["revoked"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Revoked = types.BoolValue(b)
		} else {
			m.Revoked = types.BoolNull()
		}
	} else {
		m.Revoked = types.BoolNull()
	}
	if v, ok := obj["revoked-time"]; ok {
		_ = v
		if v != "" {
			m.RevokedTime = types.StringValue(v)
		} else {
			m.RevokedTime = types.StringNull()
		}
	} else {
		m.RevokedTime = types.StringNull()
	}
	if v, ok := obj["scep-url"]; ok {
		_ = v
		if v != "" {
			m.ScepURL = types.StringValue(v)
		} else {
			m.ScepURL = types.StringNull()
		}
	} else {
		m.ScepURL = types.StringNull()
	}
	if v, ok := obj["sealed"]; ok {
		_ = v
		if v != "" {
			m.Sealed = types.StringValue(v)
		} else {
			m.Sealed = types.StringNull()
		}
	} else {
		m.Sealed = types.StringNull()
	}
	if v, ok := obj["sealed-and-hide"]; ok {
		_ = v
		if v != "" {
			m.SealedAndHide = types.StringValue(v)
		} else {
			m.SealedAndHide = types.StringNull()
		}
	} else {
		m.SealedAndHide = types.StringNull()
	}
	if v, ok := obj["serial-number"]; ok {
		_ = v
		if v != "" {
			m.SerialNumber = types.StringValue(v)
		} else {
			m.SerialNumber = types.StringNull()
		}
	} else {
		m.SerialNumber = types.StringNull()
	}
	if v, ok := obj["sign"]; ok {
		_ = v
		if v != "" {
			m.Sign = types.StringValue(v)
		} else {
			m.Sign = types.StringNull()
		}
	} else {
		m.Sign = types.StringNull()
	}
	if v, ok := obj["sign-via-scep"]; ok {
		_ = v
		if v != "" {
			m.SignViaScep = types.StringValue(v)
		} else {
			m.SignViaScep = types.StringNull()
		}
	} else {
		m.SignViaScep = types.StringNull()
	}
	if v, ok := obj["skid"]; ok {
		_ = v
		if v != "" {
			m.Skid = types.StringValue(v)
		} else {
			m.Skid = types.StringNull()
		}
	} else {
		m.Skid = types.StringNull()
	}
	if v, ok := obj["smart-card-key"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.SmartCardKey = types.BoolValue(b)
		} else {
			m.SmartCardKey = types.BoolNull()
		}
	} else {
		m.SmartCardKey = types.BoolNull()
	}
	if v, ok := obj["state"]; ok {
		_ = v
		if v != "" {
			m.State = types.StringValue(v)
		} else {
			m.State = types.StringNull()
		}
	} else {
		m.State = types.StringNull()
	}
	if v, ok := obj["subject-alt-name"]; ok {
		_ = v
		if v != "" {
			m.SubjectAltName = types.StringValue(v)
		} else {
			m.SubjectAltName = types.StringNull()
		}
	} else {
		m.SubjectAltName = types.StringNull()
	}
	if v, ok := obj["trust-store"]; ok {
		_ = v
		if v != "" {
			m.TrustStore = types.StringValue(v)
		} else {
			m.TrustStore = types.StringNull()
		}
	} else {
		m.TrustStore = types.StringNull()
	}
	if v, ok := obj["trusted"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Trusted = types.BoolValue(b)
		} else {
			m.Trusted = types.BoolNull()
		}
	} else {
		m.Trusted = types.BoolNull()
	}
	if v, ok := obj["type"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Type = types.Int64Value(n)
		} else {
			m.Type = types.Int64Null()
		}
	} else {
		m.Type = types.Int64Null()
	}
	if v, ok := obj["unit"]; ok {
		_ = v
		if v != "" {
			m.Unit = types.StringValue(v)
		} else {
			m.Unit = types.StringNull()
		}
	} else {
		m.Unit = types.StringNull()
	}
}
