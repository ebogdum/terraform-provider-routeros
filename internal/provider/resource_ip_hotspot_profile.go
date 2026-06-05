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
	Default               types.Bool   `tfsdk:"default"`
	DNSName               types.String `tfsdk:"dns_name"`
	HotspotAddress        types.String `tfsdk:"hotspot_address"`
	HtmlDirectory         types.String `tfsdk:"html_directory"`
	HtmlDirectoryOverride types.String `tfsdk:"html_directory_override"`
	HTTPCookieLifetime    types.String `tfsdk:"http_cookie_lifetime"`
	HTTPProxy             types.String `tfsdk:"http_proxy"`
	InstallHotspotQueue   types.Bool   `tfsdk:"install_hotspot_queue"`
	LoginBy               types.List   `tfsdk:"login_by"`
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
	_ = fmt.Sprintf
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
			"default": schema.BoolAttribute{
				Optional:    true,
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
			"login_by": schema.ListAttribute{
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
		body["login-by"] = encodeStringList(ctx, plan.LoginBy)
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
	obj, err := c.Add(ctx, "/ip/hotspot/profile", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/hotspot/profile failed", err.Error())
		return
	}
	iPHotspotProfileApply(ctx, obj, &plan)
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
	if !plan.DNSName.Equal(state.DNSName) {
		body["dns-name"] = plan.DNSName.ValueString()
	}
	if !plan.HotspotAddress.Equal(state.HotspotAddress) {
		body["hotspot-address"] = plan.HotspotAddress.ValueString()
	}
	if !plan.HtmlDirectory.Equal(state.HtmlDirectory) {
		body["html-directory"] = plan.HtmlDirectory.ValueString()
	}
	if !plan.HtmlDirectoryOverride.Equal(state.HtmlDirectoryOverride) {
		body["html-directory-override"] = plan.HtmlDirectoryOverride.ValueString()
	}
	if !plan.HTTPCookieLifetime.Equal(state.HTTPCookieLifetime) {
		body["http-cookie-lifetime"] = plan.HTTPCookieLifetime.ValueString()
	}
	if !plan.HTTPProxy.Equal(state.HTTPProxy) {
		body["http-proxy"] = plan.HTTPProxy.ValueString()
	}
	if !plan.InstallHotspotQueue.Equal(state.InstallHotspotQueue) {
		body["install-hotspot-queue"] = client.FormatBool(plan.InstallHotspotQueue.ValueBool())
	}
	if !plan.LoginBy.Equal(state.LoginBy) {
		body["login-by"] = encodeStringList(ctx, plan.LoginBy)
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.SMTPServer.Equal(state.SMTPServer) {
		body["smtp-server"] = plan.SMTPServer.ValueString()
	}
	if !plan.SplitUserDomain.Equal(state.SplitUserDomain) {
		body["split-user-domain"] = client.FormatBool(plan.SplitUserDomain.ValueBool())
	}
	if !plan.UseRADIUS.Equal(state.UseRADIUS) {
		body["use-radius"] = client.FormatBool(plan.UseRADIUS.ValueBool())
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
	id := req.ID
	routerName := ""
	if i := strings.Index(id, "/"); i > 0 && !strings.HasPrefix(id, "*") {
		routerName, id = id[:i], id[i+1:]
	}
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
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/ip/hotspot/profile", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func iPHotspotProfileApply(ctx context.Context, obj client.Object, m *IPHotspotProfileModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["default"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else {
			m.Default = types.BoolNull()
		}
	} else {
		m.Default = types.BoolNull()
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
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.InstallHotspotQueue = types.BoolValue(b)
		} else {
			m.InstallHotspotQueue = types.BoolNull()
		}
	} else {
		m.InstallHotspotQueue = types.BoolNull()
	}
	if v, ok := obj["login-by"]; ok {
		_ = v
		m.LoginBy = decodeStringList(ctx, v)
	} else {
		m.LoginBy = types.ListNull(types.StringType)
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
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.SplitUserDomain = types.BoolValue(b)
		} else {
			m.SplitUserDomain = types.BoolNull()
		}
	} else {
		m.SplitUserDomain = types.BoolNull()
	}
	if v, ok := obj["use-radius"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.UseRADIUS = types.BoolValue(b)
		} else {
			m.UseRADIUS = types.BoolNull()
		}
	} else {
		m.UseRADIUS = types.BoolNull()
	}
}
