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
	_ resource.Resource                = &IPHotspotProfileResource{}
	_ resource.ResourceWithImportState = &IPHotspotProfileResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPHotspotProfileResource struct {
	reg *client.Registry
}

type IPHotspotProfileModel struct {
	ID                    types.String `tfsdk:"id"`
	TrialUserProfile      types.String `tfsdk:"trial_user_profile"`
	TrialUptimeReset      types.String `tfsdk:"trial_uptime_reset"`
	TrialUptimeLimit      types.String `tfsdk:"trial_uptime_limit"`
	SslCertificate        types.String `tfsdk:"ssl_certificate"`
	RateLimit             types.String `tfsdk:"rate_limit"`
	RadiusMacFormat       types.String `tfsdk:"radius_mac_format"`
	RadiusLocationName    types.String `tfsdk:"radius_location_name"`
	RadiusLocationId      types.String `tfsdk:"radius_location_id"`
	RadiusInterimUpdate   types.String `tfsdk:"radius_interim_update"`
	RadiusDefaultDomain   types.String `tfsdk:"radius_default_domain"`
	RadiusAccounting      types.String `tfsdk:"radius_accounting"`
	NasPortType           types.String `tfsdk:"nas_port_type"`
	MacAuthPassword       types.String `tfsdk:"mac_auth_password"`
	MacAuthMode           types.String `tfsdk:"mac_auth_mode"`
	Default               types.Bool   `tfsdk:"default"`
	DNSName               types.String `tfsdk:"dns_name"`
	HotspotAddress        types.String `tfsdk:"hotspot_address"`
	HtmlDirectory         types.String `tfsdk:"html_directory"`
	HtmlDirectoryOverride types.String `tfsdk:"html_directory_override"`
	HTTPCookieLifetime    types.String `tfsdk:"http_cookie_lifetime"`
	HTTPProxy             types.String `tfsdk:"http_proxy"`
	InstallHotspotQueue   types.Bool   `tfsdk:"install_hotspot_queue"`
	LoginBy               types.Set    `tfsdk:"login_by"`
	Name                  types.String `tfsdk:"name"`
	SMTPServer            types.String `tfsdk:"smtp_server"`
	SplitUserDomain       types.Bool   `tfsdk:"split_user_domain"`
	UseRADIUS             types.Bool   `tfsdk:"use_radius"`
	Router                types.String `tfsdk:"router"`
}

func NewIPHotspotProfileResource() resource.Resource { return &IPHotspotProfileResource{} }

func (r *IPHotspotProfileResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_hotspot_profile"
}

func (r *IPHotspotProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPHotspotProfileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/hotspot/profile`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"trial_user_profile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `trial-user-profile`.",
			},
			"trial_uptime_reset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `trial-uptime-reset`.",
			},
			"trial_uptime_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `trial-uptime-limit`.",
			},
			"ssl_certificate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ssl-certificate`.",
			},
			"rate_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `rate-limit`.",
			},
			"radius_mac_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `radius-mac-format`.",
			},
			"radius_location_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `radius-location-name`.",
			},
			"radius_location_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `radius-location-id`.",
			},
			"radius_interim_update": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `radius-interim-update`.",
			},
			"radius_default_domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `radius-default-domain`.",
			},
			"radius_accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `radius-accounting`.",
			},
			"nas_port_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `nas-port-type`.",
			},
			"mac_auth_password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "RouterOS `mac-auth-password`.",
			},
			"mac_auth_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mac-auth-mode`.",
			},
			"default": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"dns_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hotspot_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"html_directory": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"html_directory_override": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"http_cookie_lifetime": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"http_proxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"install_hotspot_queue": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"login_by": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"smtp_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"split_user_domain": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"use_radius": schema.BoolAttribute{
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

func (r *IPHotspotProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPHotspotProfileModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.DNSName.IsNull() || plan.DNSName.IsUnknown()) {
		body["dns-name"] = plan.DNSName.ValueString()
	}
	if !(plan.HotspotAddress.IsNull() || plan.HotspotAddress.IsUnknown()) {
		body["hotspot-address"] = plan.HotspotAddress.ValueString()
	}
	if !(plan.HtmlDirectory.IsNull() || plan.HtmlDirectory.IsUnknown()) {
		body["html-directory"] = plan.HtmlDirectory.ValueString()
	}
	if !(plan.HtmlDirectoryOverride.IsNull() || plan.HtmlDirectoryOverride.IsUnknown()) {
		body["html-directory-override"] = plan.HtmlDirectoryOverride.ValueString()
	}
	if !(plan.HTTPCookieLifetime.IsNull() || plan.HTTPCookieLifetime.IsUnknown()) {
		body["http-cookie-lifetime"] = plan.HTTPCookieLifetime.ValueString()
	}
	if !(plan.HTTPProxy.IsNull() || plan.HTTPProxy.IsUnknown()) {
		body["http-proxy"] = plan.HTTPProxy.ValueString()
	}
	if !(plan.InstallHotspotQueue.IsNull() || plan.InstallHotspotQueue.IsUnknown()) {
		body["install-hotspot-queue"] = client.FormatBool(plan.InstallHotspotQueue.ValueBool())
	}
	if !(plan.LoginBy.IsNull() || plan.LoginBy.IsUnknown()) {
		body["login-by"] = encodeStringSet(ctx, plan.LoginBy, &resp.Diagnostics)
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.SMTPServer.IsNull() || plan.SMTPServer.IsUnknown()) {
		body["smtp-server"] = plan.SMTPServer.ValueString()
	}
	if !(plan.SplitUserDomain.IsNull() || plan.SplitUserDomain.IsUnknown()) {
		body["split-user-domain"] = client.FormatBool(plan.SplitUserDomain.ValueBool())
	}
	if !(plan.UseRADIUS.IsNull() || plan.UseRADIUS.IsUnknown()) {
		body["use-radius"] = client.FormatBool(plan.UseRADIUS.ValueBool())
	}
	if !(plan.MacAuthMode.IsNull() || plan.MacAuthMode.IsUnknown()) {
		body["mac-auth-mode"] = plan.MacAuthMode.ValueString()
	}
	if !(plan.MacAuthPassword.IsNull() || plan.MacAuthPassword.IsUnknown()) {
		body["mac-auth-password"] = plan.MacAuthPassword.ValueString()
	}
	if !(plan.NasPortType.IsNull() || plan.NasPortType.IsUnknown()) {
		body["nas-port-type"] = plan.NasPortType.ValueString()
	}
	if !(plan.RadiusAccounting.IsNull() || plan.RadiusAccounting.IsUnknown()) {
		body["radius-accounting"] = plan.RadiusAccounting.ValueString()
	}
	if !(plan.RadiusDefaultDomain.IsNull() || plan.RadiusDefaultDomain.IsUnknown()) {
		body["radius-default-domain"] = plan.RadiusDefaultDomain.ValueString()
	}
	if !(plan.RadiusInterimUpdate.IsNull() || plan.RadiusInterimUpdate.IsUnknown()) {
		body["radius-interim-update"] = plan.RadiusInterimUpdate.ValueString()
	}
	if !(plan.RadiusLocationId.IsNull() || plan.RadiusLocationId.IsUnknown()) {
		body["radius-location-id"] = plan.RadiusLocationId.ValueString()
	}
	if !(plan.RadiusLocationName.IsNull() || plan.RadiusLocationName.IsUnknown()) {
		body["radius-location-name"] = plan.RadiusLocationName.ValueString()
	}
	if !(plan.RadiusMacFormat.IsNull() || plan.RadiusMacFormat.IsUnknown()) {
		body["radius-mac-format"] = plan.RadiusMacFormat.ValueString()
	}
	if !(plan.RateLimit.IsNull() || plan.RateLimit.IsUnknown()) {
		body["rate-limit"] = plan.RateLimit.ValueString()
	}
	if !(plan.SslCertificate.IsNull() || plan.SslCertificate.IsUnknown()) {
		body["ssl-certificate"] = plan.SslCertificate.ValueString()
	}
	if !(plan.TrialUptimeLimit.IsNull() || plan.TrialUptimeLimit.IsUnknown()) {
		body["trial-uptime-limit"] = plan.TrialUptimeLimit.ValueString()
	}
	if !(plan.TrialUptimeReset.IsNull() || plan.TrialUptimeReset.IsUnknown()) {
		body["trial-uptime-reset"] = plan.TrialUptimeReset.ValueString()
	}
	if !(plan.TrialUserProfile.IsNull() || plan.TrialUserProfile.IsUnknown()) {
		body["trial-user-profile"] = plan.TrialUserProfile.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/hotspot/profile", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/hotspot/profile failed", err.Error())
		return
	}
	iPHotspotProfileApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPHotspotProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPHotspotProfileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/hotspot/profile", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/hotspot/profile failed", err.Error())
		return
	}
	iPHotspotProfileApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPHotspotProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPHotspotProfileModel
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
	if !plan.DNSName.Equal(state.DNSName) && !plan.DNSName.IsUnknown() {
		body["dns-name"] = plan.DNSName.ValueString()
	}
	if !plan.HotspotAddress.Equal(state.HotspotAddress) && !plan.HotspotAddress.IsUnknown() {
		body["hotspot-address"] = plan.HotspotAddress.ValueString()
	}
	if !plan.HtmlDirectory.Equal(state.HtmlDirectory) && !plan.HtmlDirectory.IsUnknown() {
		body["html-directory"] = plan.HtmlDirectory.ValueString()
	}
	if !plan.HtmlDirectoryOverride.Equal(state.HtmlDirectoryOverride) && !plan.HtmlDirectoryOverride.IsUnknown() {
		body["html-directory-override"] = plan.HtmlDirectoryOverride.ValueString()
	}
	if !plan.HTTPCookieLifetime.Equal(state.HTTPCookieLifetime) && !plan.HTTPCookieLifetime.IsUnknown() {
		body["http-cookie-lifetime"] = plan.HTTPCookieLifetime.ValueString()
	}
	if !plan.HTTPProxy.Equal(state.HTTPProxy) && !plan.HTTPProxy.IsUnknown() {
		body["http-proxy"] = plan.HTTPProxy.ValueString()
	}
	if !plan.InstallHotspotQueue.Equal(state.InstallHotspotQueue) && !plan.InstallHotspotQueue.IsUnknown() {
		body["install-hotspot-queue"] = client.FormatBool(plan.InstallHotspotQueue.ValueBool())
	}
	if !plan.LoginBy.Equal(state.LoginBy) && !plan.LoginBy.IsUnknown() {
		body["login-by"] = encodeStringSet(ctx, plan.LoginBy, &resp.Diagnostics)
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.SMTPServer.Equal(state.SMTPServer) && !plan.SMTPServer.IsUnknown() {
		body["smtp-server"] = plan.SMTPServer.ValueString()
	}
	if !plan.SplitUserDomain.Equal(state.SplitUserDomain) && !plan.SplitUserDomain.IsUnknown() {
		body["split-user-domain"] = client.FormatBool(plan.SplitUserDomain.ValueBool())
	}
	if !plan.UseRADIUS.Equal(state.UseRADIUS) && !plan.UseRADIUS.IsUnknown() {
		body["use-radius"] = client.FormatBool(plan.UseRADIUS.ValueBool())
	}
	if !plan.MacAuthMode.Equal(state.MacAuthMode) && !plan.MacAuthMode.IsUnknown() {
		body["mac-auth-mode"] = plan.MacAuthMode.ValueString()
	}
	if !plan.MacAuthPassword.Equal(state.MacAuthPassword) && !plan.MacAuthPassword.IsUnknown() {
		body["mac-auth-password"] = plan.MacAuthPassword.ValueString()
	}
	if !plan.NasPortType.Equal(state.NasPortType) && !plan.NasPortType.IsUnknown() {
		body["nas-port-type"] = plan.NasPortType.ValueString()
	}
	if !plan.RadiusAccounting.Equal(state.RadiusAccounting) && !plan.RadiusAccounting.IsUnknown() {
		body["radius-accounting"] = plan.RadiusAccounting.ValueString()
	}
	if !plan.RadiusDefaultDomain.Equal(state.RadiusDefaultDomain) && !plan.RadiusDefaultDomain.IsUnknown() {
		body["radius-default-domain"] = plan.RadiusDefaultDomain.ValueString()
	}
	if !plan.RadiusInterimUpdate.Equal(state.RadiusInterimUpdate) && !plan.RadiusInterimUpdate.IsUnknown() {
		body["radius-interim-update"] = plan.RadiusInterimUpdate.ValueString()
	}
	if !plan.RadiusLocationId.Equal(state.RadiusLocationId) && !plan.RadiusLocationId.IsUnknown() {
		body["radius-location-id"] = plan.RadiusLocationId.ValueString()
	}
	if !plan.RadiusLocationName.Equal(state.RadiusLocationName) && !plan.RadiusLocationName.IsUnknown() {
		body["radius-location-name"] = plan.RadiusLocationName.ValueString()
	}
	if !plan.RadiusMacFormat.Equal(state.RadiusMacFormat) && !plan.RadiusMacFormat.IsUnknown() {
		body["radius-mac-format"] = plan.RadiusMacFormat.ValueString()
	}
	if !plan.RateLimit.Equal(state.RateLimit) && !plan.RateLimit.IsUnknown() {
		body["rate-limit"] = plan.RateLimit.ValueString()
	}
	if !plan.SslCertificate.Equal(state.SslCertificate) && !plan.SslCertificate.IsUnknown() {
		body["ssl-certificate"] = plan.SslCertificate.ValueString()
	}
	if !plan.TrialUptimeLimit.Equal(state.TrialUptimeLimit) && !plan.TrialUptimeLimit.IsUnknown() {
		body["trial-uptime-limit"] = plan.TrialUptimeLimit.ValueString()
	}
	if !plan.TrialUptimeReset.Equal(state.TrialUptimeReset) && !plan.TrialUptimeReset.IsUnknown() {
		body["trial-uptime-reset"] = plan.TrialUptimeReset.ValueString()
	}
	if !plan.TrialUserProfile.Equal(state.TrialUserProfile) && !plan.TrialUserProfile.IsUnknown() {
		body["trial-user-profile"] = plan.TrialUserProfile.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/hotspot/profile", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/hotspot/profile failed", err.Error())
			return
		}
		iPHotspotProfileApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPHotspotProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPHotspotProfileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/hotspot/profile", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/hotspot/profile failed", err.Error())
	}
}

func (r *IPHotspotProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPHotspotProfileLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/hotspot/profile matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPHotspotProfileLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPHotspotProfileLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/hotspot/profile", id)
}

func iPHotspotProfileApply(ctx context.Context, obj client.Object, m *IPHotspotProfileModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["trial-user-profile"]; ok && v != "" {
		m.TrialUserProfile = types.StringValue(v)
	} else {
		m.TrialUserProfile = types.StringNull()
	}
	if v, ok := obj["trial-uptime-reset"]; ok && v != "" {
		m.TrialUptimeReset = types.StringValue(v)
	} else {
		m.TrialUptimeReset = types.StringNull()
	}
	if v, ok := obj["trial-uptime-limit"]; ok && v != "" {
		m.TrialUptimeLimit = types.StringValue(v)
	} else {
		m.TrialUptimeLimit = types.StringNull()
	}
	if v, ok := obj["ssl-certificate"]; ok && v != "" {
		m.SslCertificate = types.StringValue(v)
	} else {
		m.SslCertificate = types.StringNull()
	}
	if v, ok := obj["rate-limit"]; ok && v != "" {
		m.RateLimit = types.StringValue(v)
	}
	if v, ok := obj["radius-mac-format"]; ok && v != "" {
		m.RadiusMacFormat = types.StringValue(v)
	} else {
		m.RadiusMacFormat = types.StringNull()
	}
	if v, ok := obj["radius-location-name"]; ok && v != "" {
		m.RadiusLocationName = types.StringValue(v)
	} else {
		m.RadiusLocationName = types.StringNull()
	}
	if v, ok := obj["radius-location-id"]; ok && v != "" {
		m.RadiusLocationId = types.StringValue(v)
	} else {
		m.RadiusLocationId = types.StringNull()
	}
	if v, ok := obj["radius-interim-update"]; ok && v != "" {
		m.RadiusInterimUpdate = types.StringValue(v)
	} else {
		m.RadiusInterimUpdate = types.StringNull()
	}
	if v, ok := obj["radius-default-domain"]; ok && v != "" {
		m.RadiusDefaultDomain = types.StringValue(v)
	} else {
		m.RadiusDefaultDomain = types.StringNull()
	}
	if v, ok := obj["radius-accounting"]; ok && v != "" {
		m.RadiusAccounting = types.StringValue(v)
	} else {
		m.RadiusAccounting = types.StringNull()
	}
	if v, ok := obj["nas-port-type"]; ok && v != "" {
		m.NasPortType = types.StringValue(v)
	} else {
		m.NasPortType = types.StringNull()
	}
	if v, ok := obj["mac-auth-password"]; ok && v != "" {
		m.MacAuthPassword = types.StringValue(v)
	} else {
		m.MacAuthPassword = types.StringNull()
	}
	if v, ok := obj["mac-auth-mode"]; ok && v != "" {
		m.MacAuthMode = types.StringValue(v)
	} else {
		m.MacAuthMode = types.StringNull()
	}
	if v, ok := obj["default"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else {
			m.Default = types.BoolNull()
		}
	}
	if v, ok := obj["dns-name"]; ok {
		_ = v
		if v != "" {
			m.DNSName = types.StringValue(v)
		} else {
			m.DNSName = types.StringNull()
		}
	} else {
		m.DNSName = types.StringNull()
	}
	if v, ok := obj["hotspot-address"]; ok {
		_ = v
		if v != "" {
			m.HotspotAddress = types.StringValue(v)
		} else {
			m.HotspotAddress = types.StringNull()
		}
	} else {
		m.HotspotAddress = types.StringNull()
	}
	if v, ok := obj["html-directory"]; ok {
		_ = v
		if v != "" {
			m.HtmlDirectory = types.StringValue(v)
		} else {
			m.HtmlDirectory = types.StringNull()
		}
	} else {
		m.HtmlDirectory = types.StringNull()
	}
	if v, ok := obj["html-directory-override"]; ok {
		_ = v
		if v != "" {
			m.HtmlDirectoryOverride = types.StringValue(v)
		} else {
			m.HtmlDirectoryOverride = types.StringNull()
		}
	} else {
		m.HtmlDirectoryOverride = types.StringNull()
	}
	if v, ok := obj["http-cookie-lifetime"]; ok {
		_ = v
		if v != "" {
			m.HTTPCookieLifetime = types.StringValue(v)
		} else {
			m.HTTPCookieLifetime = types.StringNull()
		}
	} else {
		m.HTTPCookieLifetime = types.StringNull()
	}
	if v, ok := obj["http-proxy"]; ok {
		_ = v
		if v != "" {
			m.HTTPProxy = types.StringValue(v)
		} else {
			m.HTTPProxy = types.StringNull()
		}
	} else {
		m.HTTPProxy = types.StringNull()
	}
	if v, ok := obj["install-hotspot-queue"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.InstallHotspotQueue = types.BoolValue(b)
		} else {
			m.InstallHotspotQueue = types.BoolNull()
		}
	}
	if v, ok := obj["login-by"]; ok {
		_ = v
		m.LoginBy = decodeStringSet(ctx, v)
	} else {
		m.LoginBy = types.SetNull(types.StringType)
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
	if v, ok := obj["smtp-server"]; ok {
		_ = v
		if v != "" {
			m.SMTPServer = types.StringValue(v)
		} else {
			m.SMTPServer = types.StringNull()
		}
	} else {
		m.SMTPServer = types.StringNull()
	}
	if v, ok := obj["split-user-domain"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SplitUserDomain = types.BoolValue(b)
		} else {
			m.SplitUserDomain = types.BoolNull()
		}
	}
	if v, ok := obj["use-radius"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.UseRADIUS = types.BoolValue(b)
		} else {
			m.UseRADIUS = types.BoolNull()
		}
	}
}
