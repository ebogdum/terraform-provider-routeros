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
	_ resource.Resource                = &InterfaceWireguardPeersResource{}
	_ resource.ResourceWithImportState = &InterfaceWireguardPeersResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceWireguardPeersResource struct {
	reg *client.Registry
}

type InterfaceWireguardPeersModel struct {
	ID                     types.String  `tfsdk:"id"`
	AllowedAddress         types.String  `tfsdk:"allowed_address"`
	ClientAddress          types.String  `tfsdk:"client_address"`
	ClientAllowedAddress   types.String  `tfsdk:"client_allowed_address"`
	ClientConfig           types.String  `tfsdk:"client_config"`
	ClientDNS              types.String  `tfsdk:"client_dns"`
	ClientEndpoint         types.String  `tfsdk:"client_endpoint"`
	ClientKeepalive        durationValue `tfsdk:"client_keepalive"`
	ClientListenPort       types.Int64   `tfsdk:"client_listen_port"`
	ClientQr               types.String  `tfsdk:"client_qr"`
	Comment                types.String  `tfsdk:"comment"`
	CurrentEndpointAddress types.String  `tfsdk:"current_endpoint_address"`
	CurrentEndpointPort    types.Int64   `tfsdk:"current_endpoint_port"`
	Disabled               types.Bool    `tfsdk:"disabled"`
	Dynamic                types.Bool    `tfsdk:"dynamic"`
	Endpoint               types.String  `tfsdk:"endpoint"`
	EndpointAddress        types.String  `tfsdk:"endpoint_address"`
	EndpointPort           types.Int64   `tfsdk:"endpoint_port"`
	Interface              types.String  `tfsdk:"interface"`
	LastHandshake          durationValue `tfsdk:"last_handshake"`
	Name                   types.String  `tfsdk:"name"`
	PersistentKeepalive    durationValue `tfsdk:"persistent_keepalive"`
	PresharedKey           types.String  `tfsdk:"preshared_key"`
	PrivateKey             types.String  `tfsdk:"private_key"`
	PublicKey              types.String  `tfsdk:"public_key"`
	Responder              types.Bool    `tfsdk:"responder"`
	Rx                     types.String  `tfsdk:"rx"`
	Tx                     types.String  `tfsdk:"tx"`
	Router                 types.String  `tfsdk:"router"`
}

func NewInterfaceWireguardPeersResource() resource.Resource {
	return &InterfaceWireguardPeersResource{}
}

func (r *InterfaceWireguardPeersResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_wireguard_peers"
}

func (r *InterfaceWireguardPeersResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceWireguardPeersResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Peer attached to a /interface/wireguard interface. Set the `interface`\nattribute to an existing WireGuard interface name.\n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"allowed_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"client_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"client_allowed_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"client_config": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"client_dns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"client_endpoint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"client_keepalive": schema.StringAttribute{
				CustomType:  durationType{},
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"client_listen_port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"client_qr": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"current_endpoint_address": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"current_endpoint_port": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"endpoint": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"endpoint_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"endpoint_port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"last_handshake": schema.StringAttribute{
				CustomType:  durationType{},
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"persistent_keepalive": schema.StringAttribute{
				CustomType:  durationType{},
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"preshared_key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"private_key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"public_key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"responder": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *InterfaceWireguardPeersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceWireguardPeersModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AllowedAddress.IsNull() || plan.AllowedAddress.IsUnknown()) {
		body["allowed-address"] = plan.AllowedAddress.ValueString()
	}
	if !(plan.ClientAddress.IsNull() || plan.ClientAddress.IsUnknown()) {
		body["client-address"] = plan.ClientAddress.ValueString()
	}
	if !(plan.ClientAllowedAddress.IsNull() || plan.ClientAllowedAddress.IsUnknown()) {
		body["client-allowed-address"] = plan.ClientAllowedAddress.ValueString()
	}
	if !(plan.ClientDNS.IsNull() || plan.ClientDNS.IsUnknown()) {
		body["client-dns"] = plan.ClientDNS.ValueString()
	}
	if !(plan.ClientEndpoint.IsNull() || plan.ClientEndpoint.IsUnknown()) {
		body["client-endpoint"] = plan.ClientEndpoint.ValueString()
	}
	if !(plan.ClientKeepalive.IsNull() || plan.ClientKeepalive.IsUnknown()) {
		body["client-keepalive"] = plan.ClientKeepalive.ValueString()
	}
	if !(plan.ClientListenPort.IsNull() || plan.ClientListenPort.IsUnknown()) {
		body["client-listen-port"] = client.FormatInt64(plan.ClientListenPort.ValueInt64())
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.EndpointAddress.IsNull() || plan.EndpointAddress.IsUnknown()) {
		body["endpoint-address"] = plan.EndpointAddress.ValueString()
	}
	if !(plan.EndpointPort.IsNull() || plan.EndpointPort.IsUnknown()) {
		body["endpoint-port"] = client.FormatInt64(plan.EndpointPort.ValueInt64())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.PersistentKeepalive.IsNull() || plan.PersistentKeepalive.IsUnknown()) {
		body["persistent-keepalive"] = plan.PersistentKeepalive.ValueString()
	}
	if !(plan.PresharedKey.IsNull() || plan.PresharedKey.IsUnknown()) {
		body["preshared-key"] = plan.PresharedKey.ValueString()
	}
	if !(plan.PrivateKey.IsNull() || plan.PrivateKey.IsUnknown()) {
		body["private-key"] = plan.PrivateKey.ValueString()
	}
	if !(plan.PublicKey.IsNull() || plan.PublicKey.IsUnknown()) {
		body["public-key"] = plan.PublicKey.ValueString()
	}
	if !(plan.Responder.IsNull() || plan.Responder.IsUnknown()) {
		body["responder"] = client.FormatBool(plan.Responder.ValueBool())
	}
	obj, err := c.Add(ctx, "/interface/wireguard/peers", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/wireguard/peers failed", err.Error())
		return
	}
	interfaceWireguardPeersApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWireguardPeersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceWireguardPeersModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/wireguard/peers", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/wireguard/peers failed", err.Error())
		return
	}
	interfaceWireguardPeersApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceWireguardPeersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceWireguardPeersModel
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
	if !plan.AllowedAddress.Equal(state.AllowedAddress) && !plan.AllowedAddress.IsUnknown() {
		body["allowed-address"] = plan.AllowedAddress.ValueString()
	}
	if !plan.ClientAddress.Equal(state.ClientAddress) && !plan.ClientAddress.IsUnknown() {
		body["client-address"] = plan.ClientAddress.ValueString()
	}
	if !plan.ClientAllowedAddress.Equal(state.ClientAllowedAddress) && !plan.ClientAllowedAddress.IsUnknown() {
		body["client-allowed-address"] = plan.ClientAllowedAddress.ValueString()
	}
	if !plan.ClientDNS.Equal(state.ClientDNS) && !plan.ClientDNS.IsUnknown() {
		body["client-dns"] = plan.ClientDNS.ValueString()
	}
	if !plan.ClientEndpoint.Equal(state.ClientEndpoint) && !plan.ClientEndpoint.IsUnknown() {
		body["client-endpoint"] = plan.ClientEndpoint.ValueString()
	}
	if !plan.ClientKeepalive.Equal(state.ClientKeepalive) && !plan.ClientKeepalive.IsUnknown() {
		body["client-keepalive"] = plan.ClientKeepalive.ValueString()
	}
	if !plan.ClientListenPort.Equal(state.ClientListenPort) && !plan.ClientListenPort.IsUnknown() {
		body["client-listen-port"] = client.FormatInt64(plan.ClientListenPort.ValueInt64())
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.EndpointAddress.Equal(state.EndpointAddress) && !plan.EndpointAddress.IsUnknown() {
		body["endpoint-address"] = plan.EndpointAddress.ValueString()
	}
	if !plan.EndpointPort.Equal(state.EndpointPort) && !plan.EndpointPort.IsUnknown() {
		body["endpoint-port"] = client.FormatInt64(plan.EndpointPort.ValueInt64())
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.PersistentKeepalive.Equal(state.PersistentKeepalive) && !plan.PersistentKeepalive.IsUnknown() {
		body["persistent-keepalive"] = plan.PersistentKeepalive.ValueString()
	}
	if !plan.PresharedKey.Equal(state.PresharedKey) && !plan.PresharedKey.IsUnknown() {
		body["preshared-key"] = plan.PresharedKey.ValueString()
	}
	if !plan.PrivateKey.Equal(state.PrivateKey) && !plan.PrivateKey.IsUnknown() {
		body["private-key"] = plan.PrivateKey.ValueString()
	}
	if !plan.PublicKey.Equal(state.PublicKey) && !plan.PublicKey.IsUnknown() {
		body["public-key"] = plan.PublicKey.ValueString()
	}
	if !plan.Responder.Equal(state.Responder) && !plan.Responder.IsUnknown() {
		body["responder"] = client.FormatBool(plan.Responder.ValueBool())
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/wireguard/peers", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/wireguard/peers failed", err.Error())
			return
		}
		interfaceWireguardPeersApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWireguardPeersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceWireguardPeersModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/wireguard/peers", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/wireguard/peers failed", err.Error())
	}
}

func (r *InterfaceWireguardPeersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceWireguardPeersLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/wireguard/peers matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceWireguardPeersLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceWireguardPeersLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/wireguard/peers", id)
}

func interfaceWireguardPeersApply(ctx context.Context, obj client.Object, m *InterfaceWireguardPeersModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["allowed-address"]; ok {
		if v != "" {
			m.AllowedAddress = types.StringValue(v)
		} else {
			m.AllowedAddress = types.StringNull()
		}
	}
	if v, ok := obj["client-address"]; ok {
		if v != "" {
			m.ClientAddress = types.StringValue(v)
		} else {
			m.ClientAddress = types.StringNull()
		}
	}
	if v, ok := obj["client-allowed-address"]; ok {
		if v != "" {
			m.ClientAllowedAddress = types.StringValue(v)
		} else {
			m.ClientAllowedAddress = types.StringNull()
		}
	}
	if v, ok := obj["client-config"]; ok {
		if v != "" {
			m.ClientConfig = types.StringValue(v)
		} else {
			m.ClientConfig = types.StringNull()
		}
	}
	if v, ok := obj["client-dns"]; ok {
		if v != "" {
			m.ClientDNS = types.StringValue(v)
		} else {
			m.ClientDNS = types.StringNull()
		}
	}
	if v, ok := obj["client-endpoint"]; ok {
		if v != "" {
			m.ClientEndpoint = types.StringValue(v)
		} else {
			m.ClientEndpoint = types.StringNull()
		}
	}
	if v, ok := obj["client-keepalive"]; ok {
		if v != "" {
			m.ClientKeepalive = newDurationValue(v)
		} else {
			m.ClientKeepalive = newDurationNull()
		}
	}
	if v, ok := obj["client-listen-port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.ClientListenPort = types.Int64Value(n)
		} else {
			m.ClientListenPort = types.Int64Null()
		}
	} else {
		m.ClientListenPort = types.Int64Null()
	}
	if v, ok := obj["client-qr"]; ok {
		if v != "" {
			m.ClientQr = types.StringValue(v)
		} else {
			m.ClientQr = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["current-endpoint-address"]; ok {
		if v != "" {
			m.CurrentEndpointAddress = types.StringValue(v)
		} else {
			m.CurrentEndpointAddress = types.StringNull()
		}
	}
	if v, ok := obj["current-endpoint-port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.CurrentEndpointPort = types.Int64Value(n)
		} else {
			m.CurrentEndpointPort = types.Int64Null()
		}
	} else {
		m.CurrentEndpointPort = types.Int64Null()
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Disabled = types.BoolValue(true)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["dynamic"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Dynamic = types.BoolValue(true)
		} else {
			m.Dynamic = types.BoolNull()
		}
	}
	if v, ok := obj["endpoint"]; ok {
		if v != "" {
			m.Endpoint = types.StringValue(v)
		} else {
			m.Endpoint = types.StringNull()
		}
	}
	if v, ok := obj["endpoint-address"]; ok {
		if v != "" {
			m.EndpointAddress = types.StringValue(v)
		} else {
			m.EndpointAddress = types.StringNull()
		}
	}
	if v, ok := obj["endpoint-port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.EndpointPort = types.Int64Value(n)
		} else {
			m.EndpointPort = types.Int64Null()
		}
	} else {
		m.EndpointPort = types.Int64Null()
	}
	if v, ok := obj["interface"]; ok {
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	}
	if v, ok := obj["last-handshake"]; ok {
		if v != "" {
			m.LastHandshake = newDurationValue(v)
		} else {
			m.LastHandshake = newDurationNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["persistent-keepalive"]; ok {
		if v != "" {
			m.PersistentKeepalive = newDurationValue(v)
		} else {
			m.PersistentKeepalive = newDurationNull()
		}
	}
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.PresharedKey already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["preshared-key"]; ok && v != "" {
		_ = v
		if v != "" {
			m.PresharedKey = types.StringValue(v)
		} else {
			m.PresharedKey = types.StringNull()
		}
	} else if m.PresharedKey.IsUnknown() {
		m.PresharedKey = types.StringNull()
	}
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.PrivateKey already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["private-key"]; ok && v != "" {
		_ = v
		if v != "" {
			m.PrivateKey = types.StringValue(v)
		} else {
			m.PrivateKey = types.StringNull()
		}
	} else if m.PrivateKey.IsUnknown() {
		m.PrivateKey = types.StringNull()
	}
	if v, ok := obj["public-key"]; ok {
		if v != "" {
			m.PublicKey = types.StringValue(v)
		} else {
			m.PublicKey = types.StringNull()
		}
	}
	if v, ok := obj["responder"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Responder = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Responder = types.BoolValue(true)
		} else {
			m.Responder = types.BoolNull()
		}
	}
	if v, ok := obj["rx"]; ok {
		if v != "" {
			m.Rx = types.StringValue(v)
		} else {
			m.Rx = types.StringNull()
		}
	}
	if v, ok := obj["tx"]; ok {
		if v != "" {
			m.Tx = types.StringValue(v)
		} else {
			m.Tx = types.StringNull()
		}
	}
}
