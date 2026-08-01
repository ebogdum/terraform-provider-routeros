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
	_ resource.Resource                = &InterfaceL2TPClientResource{}
	_ resource.ResourceWithImportState = &InterfaceL2TPClientResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceL2TPClientResource struct {
	reg *client.Registry
}

type InterfaceL2TPClientModel struct {
	ID                   types.String `tfsdk:"id"`
	UsePeerDns           types.String `tfsdk:"use_peer_dns"`
	UseIpsec             types.String `tfsdk:"use_ipsec"`
	SrcAddress           types.String `tfsdk:"src_address"`
	RandomSourcePort     types.String `tfsdk:"random_source_port"`
	L2tpv3DigestHash     types.String `tfsdk:"l2tpv3_digest_hash"`
	L2tpv3CookieLength   types.String `tfsdk:"l2tpv3_cookie_length"`
	L2tpv3CircuitId      types.String `tfsdk:"l2tpv3_circuit_id"`
	L2tpProtoVersion     types.String `tfsdk:"l2tp_proto_version"`
	AllowFastPath        types.String `tfsdk:"allow_fast_path"`
	AddDefaultRoute      types.String `tfsdk:"add_default_route"`
	Allow                types.String `tfsdk:"allow"`
	Comment              types.String `tfsdk:"comment"`
	ConnectTo            types.String `tfsdk:"connect_to"`
	DefaultRouteDistance types.String `tfsdk:"default_route_distance"`
	DialOnDemand         types.String `tfsdk:"dial_on_demand"`
	Disabled             types.Bool   `tfsdk:"disabled"`
	IpsecSecret          types.String `tfsdk:"ipsec_secret"`
	KeepaliveTimeout     types.String `tfsdk:"keepalive_timeout"`
	MaxMru               types.String `tfsdk:"max_mru"`
	MaxMTU               types.String `tfsdk:"max_mtu"`
	Mrru                 types.String `tfsdk:"mrru"`
	Name                 types.String `tfsdk:"name"`
	Password             types.String `tfsdk:"password"`
	Profile              types.String `tfsdk:"profile"`
	User                 types.String `tfsdk:"user"`
	Router               types.String `tfsdk:"router"`
}

func NewInterfaceL2TPClientResource() resource.Resource { return &InterfaceL2TPClientResource{} }

func (r *InterfaceL2TPClientResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_l2tp_client"
}

func (r *InterfaceL2TPClientResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceL2TPClientResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/l2tp-client`.",
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
			"use_ipsec": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `use-ipsec`.",
			},
			"src_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `src-address`.",
			},
			"random_source_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `random-source-port`.",
			},
			"l2tpv3_digest_hash": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2tpv3-digest-hash`.",
			},
			"l2tpv3_cookie_length": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2tpv3-cookie-length`.",
			},
			"l2tpv3_circuit_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2tpv3-circuit-id`.",
			},
			"l2tp_proto_version": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2tp-proto-version`.",
			},
			"allow_fast_path": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `allow-fast-path`.",
			},
			"add_default_route": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `add-default-route`.",
			},
			"allow": schema.StringAttribute{
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
			"ipsec_secret": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "",
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
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceL2TPClientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceL2TPClientModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Allow.IsNull() || plan.Allow.IsUnknown()) {
		body["allow"] = plan.Allow.ValueString()
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
	if !(plan.IpsecSecret.IsNull() || plan.IpsecSecret.IsUnknown()) {
		body["ipsec-secret"] = plan.IpsecSecret.ValueString()
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
	if !(plan.AddDefaultRoute.IsNull() || plan.AddDefaultRoute.IsUnknown()) {
		body["add-default-route"] = plan.AddDefaultRoute.ValueString()
	}
	if !(plan.AllowFastPath.IsNull() || plan.AllowFastPath.IsUnknown()) {
		body["allow-fast-path"] = plan.AllowFastPath.ValueString()
	}
	if !(plan.L2tpProtoVersion.IsNull() || plan.L2tpProtoVersion.IsUnknown()) {
		body["l2tp-proto-version"] = plan.L2tpProtoVersion.ValueString()
	}
	if !(plan.L2tpv3CircuitId.IsNull() || plan.L2tpv3CircuitId.IsUnknown()) {
		body["l2tpv3-circuit-id"] = plan.L2tpv3CircuitId.ValueString()
	}
	if !(plan.L2tpv3CookieLength.IsNull() || plan.L2tpv3CookieLength.IsUnknown()) {
		body["l2tpv3-cookie-length"] = plan.L2tpv3CookieLength.ValueString()
	}
	if !(plan.L2tpv3DigestHash.IsNull() || plan.L2tpv3DigestHash.IsUnknown()) {
		body["l2tpv3-digest-hash"] = plan.L2tpv3DigestHash.ValueString()
	}
	if !(plan.RandomSourcePort.IsNull() || plan.RandomSourcePort.IsUnknown()) {
		body["random-source-port"] = plan.RandomSourcePort.ValueString()
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !(plan.UseIpsec.IsNull() || plan.UseIpsec.IsUnknown()) {
		body["use-ipsec"] = plan.UseIpsec.ValueString()
	}
	if !(plan.UsePeerDns.IsNull() || plan.UsePeerDns.IsUnknown()) {
		body["use-peer-dns"] = plan.UsePeerDns.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/l2tp-client", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/l2tp-client failed", err.Error())
		return
	}
	interfaceL2TPClientApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceL2TPClientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceL2TPClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/l2tp-client", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/l2tp-client failed", err.Error())
		return
	}
	interfaceL2TPClientApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceL2TPClientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceL2TPClientModel
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
	if !plan.Allow.Equal(state.Allow) && !plan.Allow.IsUnknown() {
		body["allow"] = plan.Allow.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.ConnectTo.Equal(state.ConnectTo) && !plan.ConnectTo.IsUnknown() {
		body["connect-to"] = plan.ConnectTo.ValueString()
	}
	if !plan.DefaultRouteDistance.Equal(state.DefaultRouteDistance) && !plan.DefaultRouteDistance.IsUnknown() {
		body["default-route-distance"] = plan.DefaultRouteDistance.ValueString()
	}
	if !plan.DialOnDemand.Equal(state.DialOnDemand) && !plan.DialOnDemand.IsUnknown() {
		body["dial-on-demand"] = plan.DialOnDemand.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.IpsecSecret.Equal(state.IpsecSecret) && !plan.IpsecSecret.IsUnknown() {
		body["ipsec-secret"] = plan.IpsecSecret.ValueString()
	}
	if !plan.KeepaliveTimeout.Equal(state.KeepaliveTimeout) && !plan.KeepaliveTimeout.IsUnknown() {
		body["keepalive-timeout"] = plan.KeepaliveTimeout.ValueString()
	}
	if !plan.MaxMru.Equal(state.MaxMru) && !plan.MaxMru.IsUnknown() {
		body["max-mru"] = plan.MaxMru.ValueString()
	}
	if !plan.MaxMTU.Equal(state.MaxMTU) && !plan.MaxMTU.IsUnknown() {
		body["max-mtu"] = plan.MaxMTU.ValueString()
	}
	if !plan.Mrru.Equal(state.Mrru) && !plan.Mrru.IsUnknown() {
		body["mrru"] = plan.Mrru.ValueString()
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
	if !plan.AddDefaultRoute.Equal(state.AddDefaultRoute) && !plan.AddDefaultRoute.IsUnknown() {
		body["add-default-route"] = plan.AddDefaultRoute.ValueString()
	}
	if !plan.AllowFastPath.Equal(state.AllowFastPath) && !plan.AllowFastPath.IsUnknown() {
		body["allow-fast-path"] = plan.AllowFastPath.ValueString()
	}
	if !plan.L2tpProtoVersion.Equal(state.L2tpProtoVersion) && !plan.L2tpProtoVersion.IsUnknown() {
		body["l2tp-proto-version"] = plan.L2tpProtoVersion.ValueString()
	}
	if !plan.L2tpv3CircuitId.Equal(state.L2tpv3CircuitId) && !plan.L2tpv3CircuitId.IsUnknown() {
		body["l2tpv3-circuit-id"] = plan.L2tpv3CircuitId.ValueString()
	}
	if !plan.L2tpv3CookieLength.Equal(state.L2tpv3CookieLength) && !plan.L2tpv3CookieLength.IsUnknown() {
		body["l2tpv3-cookie-length"] = plan.L2tpv3CookieLength.ValueString()
	}
	if !plan.L2tpv3DigestHash.Equal(state.L2tpv3DigestHash) && !plan.L2tpv3DigestHash.IsUnknown() {
		body["l2tpv3-digest-hash"] = plan.L2tpv3DigestHash.ValueString()
	}
	if !plan.RandomSourcePort.Equal(state.RandomSourcePort) && !plan.RandomSourcePort.IsUnknown() {
		body["random-source-port"] = plan.RandomSourcePort.ValueString()
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) && !plan.SrcAddress.IsUnknown() {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !plan.UseIpsec.Equal(state.UseIpsec) && !plan.UseIpsec.IsUnknown() {
		body["use-ipsec"] = plan.UseIpsec.ValueString()
	}
	if !plan.UsePeerDns.Equal(state.UsePeerDns) && !plan.UsePeerDns.IsUnknown() {
		body["use-peer-dns"] = plan.UsePeerDns.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/l2tp-client", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/l2tp-client failed", err.Error())
			return
		}
		interfaceL2TPClientApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceL2TPClientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceL2TPClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/l2tp-client", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/l2tp-client failed", err.Error())
	}
}

func (r *InterfaceL2TPClientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceL2TPClientLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/l2tp-client matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceL2TPClientLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceL2TPClientLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/l2tp-client", id)
}

func interfaceL2TPClientApply(ctx context.Context, obj client.Object, m *InterfaceL2TPClientModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["use-peer-dns"]; ok && v != "" {
		m.UsePeerDns = types.StringValue(v)
	} else {
		m.UsePeerDns = types.StringNull()
	}
	if v, ok := obj["use-ipsec"]; ok && v != "" {
		m.UseIpsec = types.StringValue(v)
	} else {
		m.UseIpsec = types.StringNull()
	}
	if v, ok := obj["src-address"]; ok && v != "" {
		m.SrcAddress = types.StringValue(v)
	} else {
		m.SrcAddress = types.StringNull()
	}
	if v, ok := obj["random-source-port"]; ok && v != "" {
		m.RandomSourcePort = types.StringValue(v)
	} else {
		m.RandomSourcePort = types.StringNull()
	}
	if v, ok := obj["l2tpv3-digest-hash"]; ok && v != "" {
		m.L2tpv3DigestHash = types.StringValue(v)
	} else {
		m.L2tpv3DigestHash = types.StringNull()
	}
	if v, ok := obj["l2tpv3-cookie-length"]; ok && v != "" {
		m.L2tpv3CookieLength = types.StringValue(v)
	} else {
		m.L2tpv3CookieLength = types.StringNull()
	}
	if v, ok := obj["l2tpv3-circuit-id"]; ok && v != "" {
		m.L2tpv3CircuitId = types.StringValue(v)
	} else {
		m.L2tpv3CircuitId = types.StringNull()
	}
	if v, ok := obj["l2tp-proto-version"]; ok && v != "" {
		m.L2tpProtoVersion = types.StringValue(v)
	} else {
		m.L2tpProtoVersion = types.StringNull()
	}
	if v, ok := obj["allow-fast-path"]; ok && v != "" {
		m.AllowFastPath = types.StringValue(v)
	} else {
		m.AllowFastPath = types.StringNull()
	}
	if v, ok := obj["add-default-route"]; ok && v != "" {
		m.AddDefaultRoute = types.StringValue(v)
	} else {
		m.AddDefaultRoute = types.StringNull()
	}
	if v, ok := obj["allow"]; ok {
		if v != "" {
			m.Allow = types.StringValue(v)
		} else {
			m.Allow = types.StringNull()
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
	if v, ok := obj["default-route-distance"]; ok {
		if v != "" {
			m.DefaultRouteDistance = types.StringValue(v)
		} else {
			m.DefaultRouteDistance = types.StringNull()
		}
	}
	if v, ok := obj["dial-on-demand"]; ok {
		if v != "" {
			m.DialOnDemand = types.StringValue(v)
		} else {
			m.DialOnDemand = types.StringNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["ipsec-secret"]; ok {
		if v != "" {
			m.IpsecSecret = types.StringValue(v)
		} else {
			m.IpsecSecret = types.StringNull()
		}
	}
	if v, ok := obj["keepalive-timeout"]; ok {
		if v != "" {
			m.KeepaliveTimeout = types.StringValue(v)
		} else {
			m.KeepaliveTimeout = types.StringNull()
		}
	}
	if v, ok := obj["max-mru"]; ok {
		if v != "" {
			m.MaxMru = types.StringValue(v)
		} else {
			m.MaxMru = types.StringNull()
		}
	}
	if v, ok := obj["max-mtu"]; ok {
		if v != "" {
			m.MaxMTU = types.StringValue(v)
		} else {
			m.MaxMTU = types.StringNull()
		}
	}
	if v, ok := obj["mrru"]; ok {
		if v != "" {
			m.Mrru = types.StringValue(v)
		} else {
			m.Mrru = types.StringNull()
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
}
