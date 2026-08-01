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
	_ resource.Resource                = &InterfaceWifiSecurityResource{}
	_ resource.ResourceWithImportState = &InterfaceWifiSecurityResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceWifiSecurityResource struct {
	reg *client.Registry
}

type InterfaceWifiSecurityModel struct {
	ID                       types.String `tfsdk:"id"`
	FtReassociationDeadline  types.String `tfsdk:"ft_reassociation_deadline"`
	FtPreserveVlanid         types.String `tfsdk:"ft_preserve_vlanid"`
	Ft                       types.String `tfsdk:"ft"`
	AuthenticationTypes      types.String `tfsdk:"authentication_types"`
	BeaconProtection         types.String `tfsdk:"beacon_protection"`
	Ciphers                  types.String `tfsdk:"ciphers"`
	Comment                  types.String `tfsdk:"comment"`
	ConnectGroup             types.String `tfsdk:"connect_group"`
	ConnectPriority          types.String `tfsdk:"connect_priority"`
	DhGroups                 types.String `tfsdk:"dh_groups"`
	DisablePmkid             types.String `tfsdk:"disable_pmkid"`
	Disabled                 types.Bool   `tfsdk:"disabled"`
	EAPAccounting            types.String `tfsdk:"eap_accounting"`
	EAPAnonymousIdentity     types.String `tfsdk:"eap_anonymous_identity"`
	EAPCertificateMode       types.String `tfsdk:"eap_certificate_mode"`
	EAPMethods               types.String `tfsdk:"eap_methods"`
	EAPPassword              types.String `tfsdk:"eap_password"`
	EAPTLSCertificate        types.String `tfsdk:"eap_tls_certificate"`
	EAPUsername              types.String `tfsdk:"eap_username"`
	Encryption               types.String `tfsdk:"encryption"`
	FtEnabled                types.String `tfsdk:"ft_enabled"`
	FtMobilityDomain         types.String `tfsdk:"ft_mobility_domain"`
	FtNasIdentifier          types.String `tfsdk:"ft_nas_identifier"`
	FtOverDs                 types.String `tfsdk:"ft_over_ds"`
	FtR0KeyLifetime          types.String `tfsdk:"ft_r0_key_lifetime"`
	FtReassocDeadline        types.String `tfsdk:"ft_reassoc_deadline"`
	GroupEncryption          types.String `tfsdk:"group_encryption"`
	GroupKeyUpdate           types.String `tfsdk:"group_key_update"`
	ManagementEncryption     types.String `tfsdk:"management_encryption"`
	ManagementProtection     types.String `tfsdk:"management_protection"`
	MultiPassphraseGroup     types.String `tfsdk:"multi_passphrase_group"`
	Name                     types.String `tfsdk:"name"`
	OweTransitionInterface   types.String `tfsdk:"owe_transition_interface"`
	Passphrase               types.String `tfsdk:"passphrase"`
	SaeAntiCloggingThreshold types.String `tfsdk:"sae_anti_clogging_threshold"`
	SaeMaxFailureRate        types.String `tfsdk:"sae_max_failure_rate"`
	SaePwe                   types.String `tfsdk:"sae_pwe"`
	Types                    types.String `tfsdk:"types"`
	Wps                      types.String `tfsdk:"wps"`
	Router                   types.String `tfsdk:"router"`
}

func NewInterfaceWifiSecurityResource() resource.Resource { return &InterfaceWifiSecurityResource{} }

func (r *InterfaceWifiSecurityResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_wifi_security"
}

func (r *InterfaceWifiSecurityResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceWifiSecurityResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/wifi/security`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ft_reassociation_deadline": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ft-reassociation-deadline`.",
			},
			"ft_preserve_vlanid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ft-preserve-vlanid`.",
			},
			"ft": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ft`.",
			},
			"authentication_types": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `authentication-types`.",
			},
			"beacon_protection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ciphers": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"connect_group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"connect_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dh_groups": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disable_pmkid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"eap_accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"eap_anonymous_identity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"eap_certificate_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"eap_methods": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"eap_password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "",
			},
			"eap_tls_certificate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"eap_username": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"encryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ft_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ft_mobility_domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ft_nas_identifier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ft_over_ds": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ft_r0_key_lifetime": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ft_reassoc_deadline": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"group_encryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"group_key_update": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"management_encryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"management_protection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"multi_passphrase_group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"owe_transition_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"passphrase": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"sae_anti_clogging_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sae_max_failure_rate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sae_pwe": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"types": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wps": schema.StringAttribute{
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

func (r *InterfaceWifiSecurityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceWifiSecurityModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.BeaconProtection.IsNull() || plan.BeaconProtection.IsUnknown()) {
		body["beacon-protection"] = plan.BeaconProtection.ValueString()
	}
	if !(plan.Ciphers.IsNull() || plan.Ciphers.IsUnknown()) {
		body["ciphers"] = plan.Ciphers.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.ConnectGroup.IsNull() || plan.ConnectGroup.IsUnknown()) {
		body["connect-group"] = plan.ConnectGroup.ValueString()
	}
	if !(plan.ConnectPriority.IsNull() || plan.ConnectPriority.IsUnknown()) {
		body["connect-priority"] = plan.ConnectPriority.ValueString()
	}
	if !(plan.DhGroups.IsNull() || plan.DhGroups.IsUnknown()) {
		body["dh-groups"] = plan.DhGroups.ValueString()
	}
	if !(plan.DisablePmkid.IsNull() || plan.DisablePmkid.IsUnknown()) {
		body["disable-pmkid"] = plan.DisablePmkid.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.EAPAccounting.IsNull() || plan.EAPAccounting.IsUnknown()) {
		body["eap-accounting"] = plan.EAPAccounting.ValueString()
	}
	if !(plan.EAPAnonymousIdentity.IsNull() || plan.EAPAnonymousIdentity.IsUnknown()) {
		body["eap-anonymous-identity"] = plan.EAPAnonymousIdentity.ValueString()
	}
	if !(plan.EAPCertificateMode.IsNull() || plan.EAPCertificateMode.IsUnknown()) {
		body["eap-certificate-mode"] = plan.EAPCertificateMode.ValueString()
	}
	if !(plan.EAPMethods.IsNull() || plan.EAPMethods.IsUnknown()) {
		body["eap-methods"] = plan.EAPMethods.ValueString()
	}
	if !(plan.EAPPassword.IsNull() || plan.EAPPassword.IsUnknown()) {
		body["eap-password"] = plan.EAPPassword.ValueString()
	}
	if !(plan.EAPTLSCertificate.IsNull() || plan.EAPTLSCertificate.IsUnknown()) {
		body["eap-tls-certificate"] = plan.EAPTLSCertificate.ValueString()
	}
	if !(plan.EAPUsername.IsNull() || plan.EAPUsername.IsUnknown()) {
		body["eap-username"] = plan.EAPUsername.ValueString()
	}
	if !(plan.Encryption.IsNull() || plan.Encryption.IsUnknown()) {
		body["encryption"] = plan.Encryption.ValueString()
	}
	if !(plan.FtEnabled.IsNull() || plan.FtEnabled.IsUnknown()) {
		body["ft-enabled"] = plan.FtEnabled.ValueString()
	}
	if !(plan.FtMobilityDomain.IsNull() || plan.FtMobilityDomain.IsUnknown()) {
		body["ft-mobility-domain"] = plan.FtMobilityDomain.ValueString()
	}
	if !(plan.FtNasIdentifier.IsNull() || plan.FtNasIdentifier.IsUnknown()) {
		body["ft-nas-identifier"] = plan.FtNasIdentifier.ValueString()
	}
	if !(plan.FtOverDs.IsNull() || plan.FtOverDs.IsUnknown()) {
		body["ft-over-ds"] = plan.FtOverDs.ValueString()
	}
	if !(plan.FtR0KeyLifetime.IsNull() || plan.FtR0KeyLifetime.IsUnknown()) {
		body["ft-r0-key-lifetime"] = plan.FtR0KeyLifetime.ValueString()
	}
	if !(plan.FtReassocDeadline.IsNull() || plan.FtReassocDeadline.IsUnknown()) {
		body["ft-reassoc-deadline"] = plan.FtReassocDeadline.ValueString()
	}
	if !(plan.GroupEncryption.IsNull() || plan.GroupEncryption.IsUnknown()) {
		body["group-encryption"] = plan.GroupEncryption.ValueString()
	}
	if !(plan.GroupKeyUpdate.IsNull() || plan.GroupKeyUpdate.IsUnknown()) {
		body["group-key-update"] = plan.GroupKeyUpdate.ValueString()
	}
	if !(plan.ManagementEncryption.IsNull() || plan.ManagementEncryption.IsUnknown()) {
		body["management-encryption"] = plan.ManagementEncryption.ValueString()
	}
	if !(plan.ManagementProtection.IsNull() || plan.ManagementProtection.IsUnknown()) {
		body["management-protection"] = plan.ManagementProtection.ValueString()
	}
	if !(plan.MultiPassphraseGroup.IsNull() || plan.MultiPassphraseGroup.IsUnknown()) {
		body["multi-passphrase-group"] = plan.MultiPassphraseGroup.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.OweTransitionInterface.IsNull() || plan.OweTransitionInterface.IsUnknown()) {
		body["owe-transition-interface"] = plan.OweTransitionInterface.ValueString()
	}
	if !(plan.Passphrase.IsNull() || plan.Passphrase.IsUnknown()) {
		body["passphrase"] = plan.Passphrase.ValueString()
	}
	if !(plan.SaeAntiCloggingThreshold.IsNull() || plan.SaeAntiCloggingThreshold.IsUnknown()) {
		body["sae-anti-clogging-threshold"] = plan.SaeAntiCloggingThreshold.ValueString()
	}
	if !(plan.SaeMaxFailureRate.IsNull() || plan.SaeMaxFailureRate.IsUnknown()) {
		body["sae-max-failure-rate"] = plan.SaeMaxFailureRate.ValueString()
	}
	if !(plan.SaePwe.IsNull() || plan.SaePwe.IsUnknown()) {
		body["sae-pwe"] = plan.SaePwe.ValueString()
	}
	if !(plan.Types.IsNull() || plan.Types.IsUnknown()) {
		body["types"] = plan.Types.ValueString()
	}
	if !(plan.Wps.IsNull() || plan.Wps.IsUnknown()) {
		body["wps"] = plan.Wps.ValueString()
	}
	if !(plan.AuthenticationTypes.IsNull() || plan.AuthenticationTypes.IsUnknown()) {
		body["authentication-types"] = plan.AuthenticationTypes.ValueString()
	}
	if !(plan.Ft.IsNull() || plan.Ft.IsUnknown()) {
		body["ft"] = plan.Ft.ValueString()
	}
	if !(plan.FtPreserveVlanid.IsNull() || plan.FtPreserveVlanid.IsUnknown()) {
		body["ft-preserve-vlanid"] = plan.FtPreserveVlanid.ValueString()
	}
	if !(plan.FtReassociationDeadline.IsNull() || plan.FtReassociationDeadline.IsUnknown()) {
		body["ft-reassociation-deadline"] = plan.FtReassociationDeadline.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/wifi/security", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/wifi/security failed", err.Error())
		return
	}
	interfaceWifiSecurityApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiSecurityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceWifiSecurityModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/wifi/security", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/wifi/security failed", err.Error())
		return
	}
	interfaceWifiSecurityApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceWifiSecurityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceWifiSecurityModel
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
	if !plan.BeaconProtection.Equal(state.BeaconProtection) && !plan.BeaconProtection.IsUnknown() {
		body["beacon-protection"] = plan.BeaconProtection.ValueString()
	}
	if !plan.Ciphers.Equal(state.Ciphers) && !plan.Ciphers.IsUnknown() {
		body["ciphers"] = plan.Ciphers.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.ConnectGroup.Equal(state.ConnectGroup) && !plan.ConnectGroup.IsUnknown() {
		body["connect-group"] = plan.ConnectGroup.ValueString()
	}
	if !plan.ConnectPriority.Equal(state.ConnectPriority) && !plan.ConnectPriority.IsUnknown() {
		body["connect-priority"] = plan.ConnectPriority.ValueString()
	}
	if !plan.DhGroups.Equal(state.DhGroups) && !plan.DhGroups.IsUnknown() {
		body["dh-groups"] = plan.DhGroups.ValueString()
	}
	if !plan.DisablePmkid.Equal(state.DisablePmkid) && !plan.DisablePmkid.IsUnknown() {
		body["disable-pmkid"] = plan.DisablePmkid.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.EAPAccounting.Equal(state.EAPAccounting) && !plan.EAPAccounting.IsUnknown() {
		body["eap-accounting"] = plan.EAPAccounting.ValueString()
	}
	if !plan.EAPAnonymousIdentity.Equal(state.EAPAnonymousIdentity) && !plan.EAPAnonymousIdentity.IsUnknown() {
		body["eap-anonymous-identity"] = plan.EAPAnonymousIdentity.ValueString()
	}
	if !plan.EAPCertificateMode.Equal(state.EAPCertificateMode) && !plan.EAPCertificateMode.IsUnknown() {
		body["eap-certificate-mode"] = plan.EAPCertificateMode.ValueString()
	}
	if !plan.EAPMethods.Equal(state.EAPMethods) && !plan.EAPMethods.IsUnknown() {
		body["eap-methods"] = plan.EAPMethods.ValueString()
	}
	if !plan.EAPPassword.Equal(state.EAPPassword) && !plan.EAPPassword.IsUnknown() {
		body["eap-password"] = plan.EAPPassword.ValueString()
	}
	if !plan.EAPTLSCertificate.Equal(state.EAPTLSCertificate) && !plan.EAPTLSCertificate.IsUnknown() {
		body["eap-tls-certificate"] = plan.EAPTLSCertificate.ValueString()
	}
	if !plan.EAPUsername.Equal(state.EAPUsername) && !plan.EAPUsername.IsUnknown() {
		body["eap-username"] = plan.EAPUsername.ValueString()
	}
	if !plan.Encryption.Equal(state.Encryption) && !plan.Encryption.IsUnknown() {
		body["encryption"] = plan.Encryption.ValueString()
	}
	if !plan.FtEnabled.Equal(state.FtEnabled) && !plan.FtEnabled.IsUnknown() {
		body["ft-enabled"] = plan.FtEnabled.ValueString()
	}
	if !plan.FtMobilityDomain.Equal(state.FtMobilityDomain) && !plan.FtMobilityDomain.IsUnknown() {
		body["ft-mobility-domain"] = plan.FtMobilityDomain.ValueString()
	}
	if !plan.FtNasIdentifier.Equal(state.FtNasIdentifier) && !plan.FtNasIdentifier.IsUnknown() {
		body["ft-nas-identifier"] = plan.FtNasIdentifier.ValueString()
	}
	if !plan.FtOverDs.Equal(state.FtOverDs) && !plan.FtOverDs.IsUnknown() {
		body["ft-over-ds"] = plan.FtOverDs.ValueString()
	}
	if !plan.FtR0KeyLifetime.Equal(state.FtR0KeyLifetime) && !plan.FtR0KeyLifetime.IsUnknown() {
		body["ft-r0-key-lifetime"] = plan.FtR0KeyLifetime.ValueString()
	}
	if !plan.FtReassocDeadline.Equal(state.FtReassocDeadline) && !plan.FtReassocDeadline.IsUnknown() {
		body["ft-reassoc-deadline"] = plan.FtReassocDeadline.ValueString()
	}
	if !plan.GroupEncryption.Equal(state.GroupEncryption) && !plan.GroupEncryption.IsUnknown() {
		body["group-encryption"] = plan.GroupEncryption.ValueString()
	}
	if !plan.GroupKeyUpdate.Equal(state.GroupKeyUpdate) && !plan.GroupKeyUpdate.IsUnknown() {
		body["group-key-update"] = plan.GroupKeyUpdate.ValueString()
	}
	if !plan.ManagementEncryption.Equal(state.ManagementEncryption) && !plan.ManagementEncryption.IsUnknown() {
		body["management-encryption"] = plan.ManagementEncryption.ValueString()
	}
	if !plan.ManagementProtection.Equal(state.ManagementProtection) && !plan.ManagementProtection.IsUnknown() {
		body["management-protection"] = plan.ManagementProtection.ValueString()
	}
	if !plan.MultiPassphraseGroup.Equal(state.MultiPassphraseGroup) && !plan.MultiPassphraseGroup.IsUnknown() {
		body["multi-passphrase-group"] = plan.MultiPassphraseGroup.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.OweTransitionInterface.Equal(state.OweTransitionInterface) && !plan.OweTransitionInterface.IsUnknown() {
		body["owe-transition-interface"] = plan.OweTransitionInterface.ValueString()
	}
	if !plan.Passphrase.Equal(state.Passphrase) && !plan.Passphrase.IsUnknown() {
		body["passphrase"] = plan.Passphrase.ValueString()
	}
	if !plan.SaeAntiCloggingThreshold.Equal(state.SaeAntiCloggingThreshold) && !plan.SaeAntiCloggingThreshold.IsUnknown() {
		body["sae-anti-clogging-threshold"] = plan.SaeAntiCloggingThreshold.ValueString()
	}
	if !plan.SaeMaxFailureRate.Equal(state.SaeMaxFailureRate) && !plan.SaeMaxFailureRate.IsUnknown() {
		body["sae-max-failure-rate"] = plan.SaeMaxFailureRate.ValueString()
	}
	if !plan.SaePwe.Equal(state.SaePwe) && !plan.SaePwe.IsUnknown() {
		body["sae-pwe"] = plan.SaePwe.ValueString()
	}
	if !plan.Types.Equal(state.Types) && !plan.Types.IsUnknown() {
		body["types"] = plan.Types.ValueString()
	}
	if !plan.Wps.Equal(state.Wps) && !plan.Wps.IsUnknown() {
		body["wps"] = plan.Wps.ValueString()
	}
	if !plan.AuthenticationTypes.Equal(state.AuthenticationTypes) && !plan.AuthenticationTypes.IsUnknown() {
		body["authentication-types"] = plan.AuthenticationTypes.ValueString()
	}
	if !plan.Ft.Equal(state.Ft) && !plan.Ft.IsUnknown() {
		body["ft"] = plan.Ft.ValueString()
	}
	if !plan.FtPreserveVlanid.Equal(state.FtPreserveVlanid) && !plan.FtPreserveVlanid.IsUnknown() {
		body["ft-preserve-vlanid"] = plan.FtPreserveVlanid.ValueString()
	}
	if !plan.FtReassociationDeadline.Equal(state.FtReassociationDeadline) && !plan.FtReassociationDeadline.IsUnknown() {
		body["ft-reassociation-deadline"] = plan.FtReassociationDeadline.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/wifi/security", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/wifi/security failed", err.Error())
			return
		}
		interfaceWifiSecurityApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiSecurityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceWifiSecurityModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/wifi/security", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/wifi/security failed", err.Error())
	}
}

func (r *InterfaceWifiSecurityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceWifiSecurityLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/wifi/security matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceWifiSecurityLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceWifiSecurityLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/wifi/security", id)
}

func interfaceWifiSecurityApply(ctx context.Context, obj client.Object, m *InterfaceWifiSecurityModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["ft-reassociation-deadline"]; ok && v != "" {
		m.FtReassociationDeadline = types.StringValue(v)
	} else {
		m.FtReassociationDeadline = types.StringNull()
	}
	if v, ok := obj["ft-preserve-vlanid"]; ok && v != "" {
		m.FtPreserveVlanid = types.StringValue(v)
	} else {
		m.FtPreserveVlanid = types.StringNull()
	}
	if v, ok := obj["ft"]; ok && v != "" {
		m.Ft = types.StringValue(v)
	} else {
		m.Ft = types.StringNull()
	}
	if v, ok := obj["authentication-types"]; ok && v != "" {
		m.AuthenticationTypes = types.StringValue(v)
	} else {
		m.AuthenticationTypes = types.StringNull()
	}
	if v, ok := obj["beacon-protection"]; ok {
		_ = v
		if v != "" {
			m.BeaconProtection = types.StringValue(v)
		} else {
			m.BeaconProtection = types.StringNull()
		}
	} else {
		m.BeaconProtection = types.StringNull()
	}
	if v, ok := obj["ciphers"]; ok {
		_ = v
		if v != "" {
			m.Ciphers = types.StringValue(v)
		} else {
			m.Ciphers = types.StringNull()
		}
	} else {
		m.Ciphers = types.StringNull()
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
	if v, ok := obj["connect-group"]; ok {
		_ = v
		if v != "" {
			m.ConnectGroup = types.StringValue(v)
		} else {
			m.ConnectGroup = types.StringNull()
		}
	} else {
		m.ConnectGroup = types.StringNull()
	}
	if v, ok := obj["connect-priority"]; ok {
		_ = v
		if v != "" {
			m.ConnectPriority = types.StringValue(v)
		} else {
			m.ConnectPriority = types.StringNull()
		}
	} else {
		m.ConnectPriority = types.StringNull()
	}
	if v, ok := obj["dh-groups"]; ok {
		_ = v
		if v != "" {
			m.DhGroups = types.StringValue(v)
		} else {
			m.DhGroups = types.StringNull()
		}
	} else {
		m.DhGroups = types.StringNull()
	}
	if v, ok := obj["disable-pmkid"]; ok {
		_ = v
		if v != "" {
			m.DisablePmkid = types.StringValue(v)
		} else {
			m.DisablePmkid = types.StringNull()
		}
	} else {
		m.DisablePmkid = types.StringNull()
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
	if v, ok := obj["eap-accounting"]; ok {
		_ = v
		if v != "" {
			m.EAPAccounting = types.StringValue(v)
		} else {
			m.EAPAccounting = types.StringNull()
		}
	} else {
		m.EAPAccounting = types.StringNull()
	}
	if v, ok := obj["eap-anonymous-identity"]; ok {
		_ = v
		if v != "" {
			m.EAPAnonymousIdentity = types.StringValue(v)
		} else {
			m.EAPAnonymousIdentity = types.StringNull()
		}
	} else {
		m.EAPAnonymousIdentity = types.StringNull()
	}
	if v, ok := obj["eap-certificate-mode"]; ok {
		_ = v
		if v != "" {
			m.EAPCertificateMode = types.StringValue(v)
		} else {
			m.EAPCertificateMode = types.StringNull()
		}
	} else {
		m.EAPCertificateMode = types.StringNull()
	}
	if v, ok := obj["eap-methods"]; ok {
		_ = v
		if v != "" {
			m.EAPMethods = types.StringValue(v)
		} else {
			m.EAPMethods = types.StringNull()
		}
	} else {
		m.EAPMethods = types.StringNull()
	}
	if v, ok := obj["eap-password"]; ok {
		_ = v
		if v != "" {
			m.EAPPassword = types.StringValue(v)
		} else {
			m.EAPPassword = types.StringNull()
		}
	} else {
		m.EAPPassword = types.StringNull()
	}
	if v, ok := obj["eap-tls-certificate"]; ok {
		_ = v
		if v != "" {
			m.EAPTLSCertificate = types.StringValue(v)
		} else {
			m.EAPTLSCertificate = types.StringNull()
		}
	} else {
		m.EAPTLSCertificate = types.StringNull()
	}
	if v, ok := obj["eap-username"]; ok {
		_ = v
		if v != "" {
			m.EAPUsername = types.StringValue(v)
		} else {
			m.EAPUsername = types.StringNull()
		}
	} else {
		m.EAPUsername = types.StringNull()
	}
	if v, ok := obj["encryption"]; ok {
		_ = v
		if v != "" {
			m.Encryption = types.StringValue(v)
		} else {
			m.Encryption = types.StringNull()
		}
	} else {
		m.Encryption = types.StringNull()
	}
	if v, ok := obj["ft-enabled"]; ok {
		_ = v
		if v != "" {
			m.FtEnabled = types.StringValue(v)
		} else {
			m.FtEnabled = types.StringNull()
		}
	} else {
		m.FtEnabled = types.StringNull()
	}
	if v, ok := obj["ft-mobility-domain"]; ok {
		_ = v
		if v != "" {
			m.FtMobilityDomain = types.StringValue(v)
		} else {
			m.FtMobilityDomain = types.StringNull()
		}
	} else {
		m.FtMobilityDomain = types.StringNull()
	}
	if v, ok := obj["ft-nas-identifier"]; ok {
		_ = v
		if v != "" {
			m.FtNasIdentifier = types.StringValue(v)
		} else {
			m.FtNasIdentifier = types.StringNull()
		}
	} else {
		m.FtNasIdentifier = types.StringNull()
	}
	if v, ok := obj["ft-over-ds"]; ok {
		_ = v
		if v != "" {
			m.FtOverDs = types.StringValue(v)
		} else {
			m.FtOverDs = types.StringNull()
		}
	} else {
		m.FtOverDs = types.StringNull()
	}
	if v, ok := obj["ft-r0-key-lifetime"]; ok {
		_ = v
		if v != "" {
			m.FtR0KeyLifetime = types.StringValue(v)
		} else {
			m.FtR0KeyLifetime = types.StringNull()
		}
	} else {
		m.FtR0KeyLifetime = types.StringNull()
	}
	if v, ok := obj["ft-reassoc-deadline"]; ok {
		_ = v
		if v != "" {
			m.FtReassocDeadline = types.StringValue(v)
		} else {
			m.FtReassocDeadline = types.StringNull()
		}
	} else {
		m.FtReassocDeadline = types.StringNull()
	}
	if v, ok := obj["group-encryption"]; ok {
		_ = v
		if v != "" {
			m.GroupEncryption = types.StringValue(v)
		} else {
			m.GroupEncryption = types.StringNull()
		}
	} else {
		m.GroupEncryption = types.StringNull()
	}
	if v, ok := obj["group-key-update"]; ok {
		_ = v
		if v != "" {
			m.GroupKeyUpdate = types.StringValue(v)
		} else {
			m.GroupKeyUpdate = types.StringNull()
		}
	} else {
		m.GroupKeyUpdate = types.StringNull()
	}
	if v, ok := obj["management-encryption"]; ok {
		_ = v
		if v != "" {
			m.ManagementEncryption = types.StringValue(v)
		} else {
			m.ManagementEncryption = types.StringNull()
		}
	} else {
		m.ManagementEncryption = types.StringNull()
	}
	if v, ok := obj["management-protection"]; ok {
		_ = v
		if v != "" {
			m.ManagementProtection = types.StringValue(v)
		} else {
			m.ManagementProtection = types.StringNull()
		}
	} else {
		m.ManagementProtection = types.StringNull()
	}
	if v, ok := obj["multi-passphrase-group"]; ok {
		_ = v
		if v != "" {
			m.MultiPassphraseGroup = types.StringValue(v)
		} else {
			m.MultiPassphraseGroup = types.StringNull()
		}
	} else {
		m.MultiPassphraseGroup = types.StringNull()
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
	if v, ok := obj["owe-transition-interface"]; ok {
		_ = v
		if v != "" {
			m.OweTransitionInterface = types.StringValue(v)
		} else {
			m.OweTransitionInterface = types.StringNull()
		}
	} else {
		m.OweTransitionInterface = types.StringNull()
	}
	if v, ok := obj["passphrase"]; ok {
		_ = v
		if v != "" {
			m.Passphrase = types.StringValue(v)
		} else {
			m.Passphrase = types.StringNull()
		}
	} else {
		m.Passphrase = types.StringNull()
	}
	if v, ok := obj["sae-anti-clogging-threshold"]; ok {
		_ = v
		if v != "" {
			m.SaeAntiCloggingThreshold = types.StringValue(v)
		} else {
			m.SaeAntiCloggingThreshold = types.StringNull()
		}
	} else {
		m.SaeAntiCloggingThreshold = types.StringNull()
	}
	if v, ok := obj["sae-max-failure-rate"]; ok {
		_ = v
		if v != "" {
			m.SaeMaxFailureRate = types.StringValue(v)
		} else {
			m.SaeMaxFailureRate = types.StringNull()
		}
	} else {
		m.SaeMaxFailureRate = types.StringNull()
	}
	if v, ok := obj["sae-pwe"]; ok {
		_ = v
		if v != "" {
			m.SaePwe = types.StringValue(v)
		} else {
			m.SaePwe = types.StringNull()
		}
	} else {
		m.SaePwe = types.StringNull()
	}
	if v, ok := obj["types"]; ok {
		_ = v
		if v != "" {
			m.Types = types.StringValue(v)
		} else {
			m.Types = types.StringNull()
		}
	} else {
		m.Types = types.StringNull()
	}
	if v, ok := obj["wps"]; ok {
		_ = v
		if v != "" {
			m.Wps = types.StringValue(v)
		} else {
			m.Wps = types.StringNull()
		}
	} else {
		m.Wps = types.StringNull()
	}
}
