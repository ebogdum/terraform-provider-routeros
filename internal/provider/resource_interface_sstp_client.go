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
	_ resource.Resource                = &InterfaceSSTPClientResource{}
	_ resource.ResourceWithImportState = &InterfaceSSTPClientResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceSSTPClientResource struct {
	reg *client.Registry
}

type InterfaceSSTPClientModel struct {
	ID                                 types.String `tfsdk:"id"`
	VerifyServerAddressFromCertificate types.String `tfsdk:"verify_server_address_from_certificate"`
	TlsVersion                         types.String `tfsdk:"tls_version"`
	ProxyPort                          types.String `tfsdk:"proxy_port"`
	Port                               types.String `tfsdk:"port"`
	Pfs                                types.String `tfsdk:"pfs"`
	HttpProxy                          types.String `tfsdk:"http_proxy"`
	Ciphers                            types.String `tfsdk:"ciphers"`
	AddSni                             types.String `tfsdk:"add_sni"`
	AddDefaultRoute                    types.String `tfsdk:"add_default_route"`
	Authentication                     types.String `tfsdk:"authentication"`
	Certificate                        types.String `tfsdk:"certificate"`
	Comment                            types.String `tfsdk:"comment"`
	ConnectTo                          types.String `tfsdk:"connect_to"`
	DefaultRouteDistance               types.String `tfsdk:"default_route_distance"`
	DialOnDemand                       types.String `tfsdk:"dial_on_demand"`
	Disabled                           types.Bool   `tfsdk:"disabled"`
	KeepaliveTimeout                   types.String `tfsdk:"keepalive_timeout"`
	MaxMru                             types.String `tfsdk:"max_mru"`
	MaxMTU                             types.String `tfsdk:"max_mtu"`
	Mrru                               types.String `tfsdk:"mrru"`
	Name                               types.String `tfsdk:"name"`
	Password                           types.String `tfsdk:"password"`
	Profile                            types.String `tfsdk:"profile"`
	User                               types.String `tfsdk:"user"`
	VerifyServerCertificate            types.String `tfsdk:"verify_server_certificate"`
	Router                             types.String `tfsdk:"router"`
}

func NewInterfaceSSTPClientResource() resource.Resource { return &InterfaceSSTPClientResource{} }

func (r *InterfaceSSTPClientResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_sstp_client"
}

func (r *InterfaceSSTPClientResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceSSTPClientResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/sstp-client`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"verify_server_address_from_certificate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `verify-server-address-from-certificate`.",
			},
			"tls_version": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tls-version`.",
			},
			"proxy_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `proxy-port`.",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `port`.",
			},
			"pfs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `pfs`.",
			},
			"http_proxy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `http-proxy`.",
			},
			"ciphers": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ciphers`.",
			},
			"add_sni": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `add-sni`.",
			},
			"add_default_route": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `add-default-route`.",
			},
			"authentication": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"certificate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"connect_to": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"default_route_distance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dial_on_demand": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"keepalive_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"max_mru": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"max_mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mrru": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"profile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"user": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"verify_server_certificate": schema.StringAttribute{
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

func (r *InterfaceSSTPClientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceSSTPClientModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Authentication.IsNull() || plan.Authentication.IsUnknown()) {
		body["authentication"] = plan.Authentication.ValueString()
	}
	if !(plan.Certificate.IsNull() || plan.Certificate.IsUnknown()) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.ConnectTo.IsNull() || plan.ConnectTo.IsUnknown()) {
		body["connect-to"] = plan.ConnectTo.ValueString()
	}
	if !(plan.DefaultRouteDistance.IsNull() || plan.DefaultRouteDistance.IsUnknown()) {
		body["default-route-distance"] = plan.DefaultRouteDistance.ValueString()
	}
	if !(plan.DialOnDemand.IsNull() || plan.DialOnDemand.IsUnknown()) {
		body["dial-on-demand"] = plan.DialOnDemand.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.KeepaliveTimeout.IsNull() || plan.KeepaliveTimeout.IsUnknown()) {
		body["keepalive-timeout"] = plan.KeepaliveTimeout.ValueString()
	}
	if !(plan.MaxMru.IsNull() || plan.MaxMru.IsUnknown()) {
		body["max-mru"] = plan.MaxMru.ValueString()
	}
	if !(plan.MaxMTU.IsNull() || plan.MaxMTU.IsUnknown()) {
		body["max-mtu"] = plan.MaxMTU.ValueString()
	}
	if !(plan.Mrru.IsNull() || plan.Mrru.IsUnknown()) {
		body["mrru"] = plan.Mrru.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Password.IsNull() || plan.Password.IsUnknown()) {
		body["password"] = plan.Password.ValueString()
	}
	if !(plan.Profile.IsNull() || plan.Profile.IsUnknown()) {
		body["profile"] = plan.Profile.ValueString()
	}
	if !(plan.User.IsNull() || plan.User.IsUnknown()) {
		body["user"] = plan.User.ValueString()
	}
	if !(plan.VerifyServerCertificate.IsNull() || plan.VerifyServerCertificate.IsUnknown()) {
		body["verify-server-certificate"] = plan.VerifyServerCertificate.ValueString()
	}
	if !(plan.AddDefaultRoute.IsNull() || plan.AddDefaultRoute.IsUnknown()) {
		body["add-default-route"] = plan.AddDefaultRoute.ValueString()
	}
	if !(plan.AddSni.IsNull() || plan.AddSni.IsUnknown()) {
		body["add-sni"] = plan.AddSni.ValueString()
	}
	if !(plan.Ciphers.IsNull() || plan.Ciphers.IsUnknown()) {
		body["ciphers"] = plan.Ciphers.ValueString()
	}
	if !(plan.HttpProxy.IsNull() || plan.HttpProxy.IsUnknown()) {
		body["http-proxy"] = plan.HttpProxy.ValueString()
	}
	if !(plan.Pfs.IsNull() || plan.Pfs.IsUnknown()) {
		body["pfs"] = plan.Pfs.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.ProxyPort.IsNull() || plan.ProxyPort.IsUnknown()) {
		body["proxy-port"] = plan.ProxyPort.ValueString()
	}
	if !(plan.TlsVersion.IsNull() || plan.TlsVersion.IsUnknown()) {
		body["tls-version"] = plan.TlsVersion.ValueString()
	}
	if !(plan.VerifyServerAddressFromCertificate.IsNull() || plan.VerifyServerAddressFromCertificate.IsUnknown()) {
		body["verify-server-address-from-certificate"] = plan.VerifyServerAddressFromCertificate.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/sstp-client", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/sstp-client failed", err.Error())
		return
	}
	interfaceSSTPClientApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceSSTPClientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceSSTPClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/sstp-client", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/sstp-client failed", err.Error())
		return
	}
	interfaceSSTPClientApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceSSTPClientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceSSTPClientModel
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
	if !plan.Authentication.Equal(state.Authentication) {
		body["authentication"] = plan.Authentication.ValueString()
	}
	if !plan.Certificate.Equal(state.Certificate) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.ConnectTo.Equal(state.ConnectTo) {
		body["connect-to"] = plan.ConnectTo.ValueString()
	}
	if !plan.DefaultRouteDistance.Equal(state.DefaultRouteDistance) {
		body["default-route-distance"] = plan.DefaultRouteDistance.ValueString()
	}
	if !plan.DialOnDemand.Equal(state.DialOnDemand) {
		body["dial-on-demand"] = plan.DialOnDemand.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.KeepaliveTimeout.Equal(state.KeepaliveTimeout) {
		body["keepalive-timeout"] = plan.KeepaliveTimeout.ValueString()
	}
	if !plan.MaxMru.Equal(state.MaxMru) {
		body["max-mru"] = plan.MaxMru.ValueString()
	}
	if !plan.MaxMTU.Equal(state.MaxMTU) {
		body["max-mtu"] = plan.MaxMTU.ValueString()
	}
	if !plan.Mrru.Equal(state.Mrru) {
		body["mrru"] = plan.Mrru.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Password.Equal(state.Password) {
		body["password"] = plan.Password.ValueString()
	}
	if !plan.Profile.Equal(state.Profile) {
		body["profile"] = plan.Profile.ValueString()
	}
	if !plan.User.Equal(state.User) {
		body["user"] = plan.User.ValueString()
	}
	if !plan.VerifyServerCertificate.Equal(state.VerifyServerCertificate) {
		body["verify-server-certificate"] = plan.VerifyServerCertificate.ValueString()
	}
	if !plan.AddDefaultRoute.Equal(state.AddDefaultRoute) && !plan.AddDefaultRoute.IsUnknown() {
		body["add-default-route"] = plan.AddDefaultRoute.ValueString()
	}
	if !plan.AddSni.Equal(state.AddSni) && !plan.AddSni.IsUnknown() {
		body["add-sni"] = plan.AddSni.ValueString()
	}
	if !plan.Ciphers.Equal(state.Ciphers) && !plan.Ciphers.IsUnknown() {
		body["ciphers"] = plan.Ciphers.ValueString()
	}
	if !plan.HttpProxy.Equal(state.HttpProxy) && !plan.HttpProxy.IsUnknown() {
		body["http-proxy"] = plan.HttpProxy.ValueString()
	}
	if !plan.Pfs.Equal(state.Pfs) && !plan.Pfs.IsUnknown() {
		body["pfs"] = plan.Pfs.ValueString()
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.ProxyPort.Equal(state.ProxyPort) && !plan.ProxyPort.IsUnknown() {
		body["proxy-port"] = plan.ProxyPort.ValueString()
	}
	if !plan.TlsVersion.Equal(state.TlsVersion) && !plan.TlsVersion.IsUnknown() {
		body["tls-version"] = plan.TlsVersion.ValueString()
	}
	if !plan.VerifyServerAddressFromCertificate.Equal(state.VerifyServerAddressFromCertificate) && !plan.VerifyServerAddressFromCertificate.IsUnknown() {
		body["verify-server-address-from-certificate"] = plan.VerifyServerAddressFromCertificate.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/sstp-client", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/sstp-client failed", err.Error())
			return
		}
		interfaceSSTPClientApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceSSTPClientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceSSTPClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/sstp-client", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/sstp-client failed", err.Error())
	}
}

func (r *InterfaceSSTPClientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceSSTPClientLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/sstp-client matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceSSTPClientLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceSSTPClientLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/sstp-client", id)
}

func interfaceSSTPClientApply(ctx context.Context, obj client.Object, m *InterfaceSSTPClientModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["verify-server-address-from-certificate"]; ok && v != "" {
		m.VerifyServerAddressFromCertificate = types.StringValue(v)
	} else {
		m.VerifyServerAddressFromCertificate = types.StringNull()
	}
	if v, ok := obj["tls-version"]; ok && v != "" {
		m.TlsVersion = types.StringValue(v)
	} else {
		m.TlsVersion = types.StringNull()
	}
	if v, ok := obj["proxy-port"]; ok && v != "" {
		m.ProxyPort = types.StringValue(v)
	} else {
		m.ProxyPort = types.StringNull()
	}
	if v, ok := obj["port"]; ok && v != "" {
		m.Port = types.StringValue(v)
	} else {
		m.Port = types.StringNull()
	}
	if v, ok := obj["pfs"]; ok && v != "" {
		m.Pfs = types.StringValue(v)
	} else {
		m.Pfs = types.StringNull()
	}
	if v, ok := obj["http-proxy"]; ok && v != "" {
		m.HttpProxy = types.StringValue(v)
	} else {
		m.HttpProxy = types.StringNull()
	}
	if v, ok := obj["ciphers"]; ok && v != "" {
		m.Ciphers = types.StringValue(v)
	} else {
		m.Ciphers = types.StringNull()
	}
	if v, ok := obj["add-sni"]; ok && v != "" {
		m.AddSni = types.StringValue(v)
	} else {
		m.AddSni = types.StringNull()
	}
	if v, ok := obj["add-default-route"]; ok && v != "" {
		m.AddDefaultRoute = types.StringValue(v)
	} else {
		m.AddDefaultRoute = types.StringNull()
	}
	if v, ok := obj["authentication"]; ok {
		_ = v
		if v != "" {
			m.Authentication = types.StringValue(v)
		} else {
			m.Authentication = types.StringNull()
		}
	} else {
		m.Authentication = types.StringNull()
	}
	if v, ok := obj["certificate"]; ok {
		_ = v
		if v != "" {
			m.Certificate = types.StringValue(v)
		} else {
			m.Certificate = types.StringNull()
		}
	} else {
		m.Certificate = types.StringNull()
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
	if v, ok := obj["connect-to"]; ok {
		_ = v
		if v != "" {
			m.ConnectTo = types.StringValue(v)
		} else {
			m.ConnectTo = types.StringNull()
		}
	} else {
		m.ConnectTo = types.StringNull()
	}
	if v, ok := obj["default-route-distance"]; ok {
		_ = v
		if v != "" {
			m.DefaultRouteDistance = types.StringValue(v)
		} else {
			m.DefaultRouteDistance = types.StringNull()
		}
	} else {
		m.DefaultRouteDistance = types.StringNull()
	}
	if v, ok := obj["dial-on-demand"]; ok {
		_ = v
		if v != "" {
			m.DialOnDemand = types.StringValue(v)
		} else {
			m.DialOnDemand = types.StringNull()
		}
	} else {
		m.DialOnDemand = types.StringNull()
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
	if v, ok := obj["keepalive-timeout"]; ok {
		_ = v
		if v != "" {
			m.KeepaliveTimeout = types.StringValue(v)
		} else {
			m.KeepaliveTimeout = types.StringNull()
		}
	} else {
		m.KeepaliveTimeout = types.StringNull()
	}
	if v, ok := obj["max-mru"]; ok {
		_ = v
		if v != "" {
			m.MaxMru = types.StringValue(v)
		} else {
			m.MaxMru = types.StringNull()
		}
	} else {
		m.MaxMru = types.StringNull()
	}
	if v, ok := obj["max-mtu"]; ok {
		_ = v
		if v != "" {
			m.MaxMTU = types.StringValue(v)
		} else {
			m.MaxMTU = types.StringNull()
		}
	} else {
		m.MaxMTU = types.StringNull()
	}
	if v, ok := obj["mrru"]; ok {
		_ = v
		if v != "" {
			m.Mrru = types.StringValue(v)
		} else {
			m.Mrru = types.StringNull()
		}
	} else {
		m.Mrru = types.StringNull()
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
	if v, ok := obj["password"]; ok {
		_ = v
		if v != "" {
			m.Password = types.StringValue(v)
		} else {
			m.Password = types.StringNull()
		}
	} else {
		m.Password = types.StringNull()
	}
	if v, ok := obj["profile"]; ok {
		_ = v
		if v != "" {
			m.Profile = types.StringValue(v)
		} else {
			m.Profile = types.StringNull()
		}
	} else {
		m.Profile = types.StringNull()
	}
	if v, ok := obj["user"]; ok {
		_ = v
		if v != "" {
			m.User = types.StringValue(v)
		} else {
			m.User = types.StringNull()
		}
	} else {
		m.User = types.StringNull()
	}
	if v, ok := obj["verify-server-certificate"]; ok {
		_ = v
		if v != "" {
			m.VerifyServerCertificate = types.StringValue(v)
		} else {
			m.VerifyServerCertificate = types.StringNull()
		}
	} else {
		m.VerifyServerCertificate = types.StringNull()
	}
}
