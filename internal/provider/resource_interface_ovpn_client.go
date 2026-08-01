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
	_ resource.Resource                = &InterfaceOVPNClientResource{}
	_ resource.ResourceWithImportState = &InterfaceOVPNClientResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceOVPNClientResource struct {
	reg *client.Registry
}

type InterfaceOVPNClientModel struct {
	ID                      types.String `tfsdk:"id"`
	UsePeerDns              types.String `tfsdk:"use_peer_dns"`
	TlsVersion              types.String `tfsdk:"tls_version"`
	RouteNopull             types.String `tfsdk:"route_nopull"`
	Protocol                types.String `tfsdk:"protocol"`
	Port                    types.String `tfsdk:"port"`
	DisconnectNotify        types.String `tfsdk:"disconnect_notify"`
	Auth                    types.String `tfsdk:"auth"`
	AddDefaultRoute         types.String `tfsdk:"add_default_route"`
	Certificate             types.String `tfsdk:"certificate"`
	Cipher                  types.String `tfsdk:"cipher"`
	Comment                 types.String `tfsdk:"comment"`
	ConnectTo               types.String `tfsdk:"connect_to"`
	Disabled                types.Bool   `tfsdk:"disabled"`
	MACAddress              types.String `tfsdk:"mac_address"`
	MaxMTU                  types.String `tfsdk:"max_mtu"`
	Mode                    types.String `tfsdk:"mode"`
	Name                    types.String `tfsdk:"name"`
	Password                types.String `tfsdk:"password"`
	Profile                 types.String `tfsdk:"profile"`
	User                    types.String `tfsdk:"user"`
	VerifyServerCertificate types.String `tfsdk:"verify_server_certificate"`
	Router                  types.String `tfsdk:"router"`
}

func NewInterfaceOVPNClientResource() resource.Resource { return &InterfaceOVPNClientResource{} }

func (r *InterfaceOVPNClientResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ovpn_client"
}

func (r *InterfaceOVPNClientResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceOVPNClientResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/ovpn-client`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"use_peer_dns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `use-peer-dns`.",
			},
			"tls_version": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tls-version`.",
			},
			"route_nopull": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `route-nopull`.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `protocol`.",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `port`.",
			},
			"disconnect_notify": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `disconnect-notify`.",
			},
			"auth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `auth`.",
			},
			"add_default_route": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `add-default-route`.",
			},
			"certificate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"cipher": schema.StringAttribute{
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
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"max_mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mode": schema.StringAttribute{
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

func (r *InterfaceOVPNClientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceOVPNClientModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Certificate.IsNull() || plan.Certificate.IsUnknown()) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !(plan.Cipher.IsNull() || plan.Cipher.IsUnknown()) {
		body["cipher"] = plan.Cipher.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.ConnectTo.IsNull() || plan.ConnectTo.IsUnknown()) {
		body["connect-to"] = plan.ConnectTo.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.MACAddress.IsNull() || plan.MACAddress.IsUnknown()) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !(plan.MaxMTU.IsNull() || plan.MaxMTU.IsUnknown()) {
		body["max-mtu"] = plan.MaxMTU.ValueString()
	}
	if !(plan.Mode.IsNull() || plan.Mode.IsUnknown()) {
		body["mode"] = plan.Mode.ValueString()
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
	if !(plan.Auth.IsNull() || plan.Auth.IsUnknown()) {
		body["auth"] = plan.Auth.ValueString()
	}
	if !(plan.DisconnectNotify.IsNull() || plan.DisconnectNotify.IsUnknown()) {
		body["disconnect-notify"] = plan.DisconnectNotify.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.Protocol.IsNull() || plan.Protocol.IsUnknown()) {
		body["protocol"] = plan.Protocol.ValueString()
	}
	if !(plan.RouteNopull.IsNull() || plan.RouteNopull.IsUnknown()) {
		body["route-nopull"] = plan.RouteNopull.ValueString()
	}
	if !(plan.TlsVersion.IsNull() || plan.TlsVersion.IsUnknown()) {
		body["tls-version"] = plan.TlsVersion.ValueString()
	}
	if !(plan.UsePeerDns.IsNull() || plan.UsePeerDns.IsUnknown()) {
		body["use-peer-dns"] = plan.UsePeerDns.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/ovpn-client", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/ovpn-client failed", err.Error())
		return
	}
	interfaceOVPNClientApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceOVPNClientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceOVPNClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/ovpn-client", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/ovpn-client failed", err.Error())
		return
	}
	interfaceOVPNClientApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceOVPNClientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceOVPNClientModel
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
	if !plan.Certificate.Equal(state.Certificate) && !plan.Certificate.IsUnknown() {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !plan.Cipher.Equal(state.Cipher) && !plan.Cipher.IsUnknown() {
		body["cipher"] = plan.Cipher.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.ConnectTo.Equal(state.ConnectTo) && !plan.ConnectTo.IsUnknown() {
		body["connect-to"] = plan.ConnectTo.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.MACAddress.Equal(state.MACAddress) && !plan.MACAddress.IsUnknown() {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.MaxMTU.Equal(state.MaxMTU) && !plan.MaxMTU.IsUnknown() {
		body["max-mtu"] = plan.MaxMTU.ValueString()
	}
	if !plan.Mode.Equal(state.Mode) && !plan.Mode.IsUnknown() {
		body["mode"] = plan.Mode.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Password.Equal(state.Password) && !plan.Password.IsUnknown() {
		body["password"] = plan.Password.ValueString()
	}
	if !plan.Profile.Equal(state.Profile) && !plan.Profile.IsUnknown() {
		body["profile"] = plan.Profile.ValueString()
	}
	if !plan.User.Equal(state.User) && !plan.User.IsUnknown() {
		body["user"] = plan.User.ValueString()
	}
	if !plan.VerifyServerCertificate.Equal(state.VerifyServerCertificate) && !plan.VerifyServerCertificate.IsUnknown() {
		body["verify-server-certificate"] = plan.VerifyServerCertificate.ValueString()
	}
	if !plan.AddDefaultRoute.Equal(state.AddDefaultRoute) && !plan.AddDefaultRoute.IsUnknown() {
		body["add-default-route"] = plan.AddDefaultRoute.ValueString()
	}
	if !plan.Auth.Equal(state.Auth) && !plan.Auth.IsUnknown() {
		body["auth"] = plan.Auth.ValueString()
	}
	if !plan.DisconnectNotify.Equal(state.DisconnectNotify) && !plan.DisconnectNotify.IsUnknown() {
		body["disconnect-notify"] = plan.DisconnectNotify.ValueString()
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.Protocol.Equal(state.Protocol) && !plan.Protocol.IsUnknown() {
		body["protocol"] = plan.Protocol.ValueString()
	}
	if !plan.RouteNopull.Equal(state.RouteNopull) && !plan.RouteNopull.IsUnknown() {
		body["route-nopull"] = plan.RouteNopull.ValueString()
	}
	if !plan.TlsVersion.Equal(state.TlsVersion) && !plan.TlsVersion.IsUnknown() {
		body["tls-version"] = plan.TlsVersion.ValueString()
	}
	if !plan.UsePeerDns.Equal(state.UsePeerDns) && !plan.UsePeerDns.IsUnknown() {
		body["use-peer-dns"] = plan.UsePeerDns.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/ovpn-client", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/ovpn-client failed", err.Error())
			return
		}
		interfaceOVPNClientApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceOVPNClientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceOVPNClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/ovpn-client", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/ovpn-client failed", err.Error())
	}
}

func (r *InterfaceOVPNClientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceOVPNClientLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/ovpn-client matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceOVPNClientLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceOVPNClientLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/ovpn-client", id)
}

func interfaceOVPNClientApply(ctx context.Context, obj client.Object, m *InterfaceOVPNClientModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["use-peer-dns"]; ok && v != "" {
		m.UsePeerDns = types.StringValue(v)
	} else {
		m.UsePeerDns = types.StringNull()
	}
	if v, ok := obj["tls-version"]; ok && v != "" {
		m.TlsVersion = types.StringValue(v)
	} else {
		m.TlsVersion = types.StringNull()
	}
	if v, ok := obj["route-nopull"]; ok && v != "" {
		m.RouteNopull = types.StringValue(v)
	} else {
		m.RouteNopull = types.StringNull()
	}
	if v, ok := obj["protocol"]; ok && v != "" {
		m.Protocol = types.StringValue(v)
	} else {
		m.Protocol = types.StringNull()
	}
	if v, ok := obj["port"]; ok && v != "" {
		m.Port = types.StringValue(v)
	} else {
		m.Port = types.StringNull()
	}
	if v, ok := obj["disconnect-notify"]; ok && v != "" {
		m.DisconnectNotify = types.StringValue(v)
	} else {
		m.DisconnectNotify = types.StringNull()
	}
	if v, ok := obj["auth"]; ok && v != "" {
		m.Auth = types.StringValue(v)
	} else {
		m.Auth = types.StringNull()
	}
	if v, ok := obj["add-default-route"]; ok && v != "" {
		m.AddDefaultRoute = types.StringValue(v)
	} else {
		m.AddDefaultRoute = types.StringNull()
	}
	if v, ok := obj["certificate"]; ok {
		if v != "" {
			m.Certificate = types.StringValue(v)
		} else {
			m.Certificate = types.StringNull()
		}
	}
	if v, ok := obj["cipher"]; ok {
		if v != "" {
			m.Cipher = types.StringValue(v)
		} else {
			m.Cipher = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["connect-to"]; ok {
		if v != "" {
			m.ConnectTo = types.StringValue(v)
		} else {
			m.ConnectTo = types.StringNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["mac-address"]; ok {
		if v != "" {
			m.MACAddress = types.StringValue(v)
		} else {
			m.MACAddress = types.StringNull()
		}
	}
	if v, ok := obj["max-mtu"]; ok {
		if v != "" {
			m.MaxMTU = types.StringValue(v)
		} else {
			m.MaxMTU = types.StringNull()
		}
	}
	if v, ok := obj["mode"]; ok {
		if v != "" {
			m.Mode = types.StringValue(v)
		} else {
			m.Mode = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["password"]; ok {
		if v != "" {
			m.Password = types.StringValue(v)
		} else {
			m.Password = types.StringNull()
		}
	}
	if v, ok := obj["profile"]; ok {
		if v != "" {
			m.Profile = types.StringValue(v)
		} else {
			m.Profile = types.StringNull()
		}
	}
	if v, ok := obj["user"]; ok {
		if v != "" {
			m.User = types.StringValue(v)
		} else {
			m.User = types.StringNull()
		}
	}
	if v, ok := obj["verify-server-certificate"]; ok {
		if v != "" {
			m.VerifyServerCertificate = types.StringValue(v)
		} else {
			m.VerifyServerCertificate = types.StringNull()
		}
	}
}
