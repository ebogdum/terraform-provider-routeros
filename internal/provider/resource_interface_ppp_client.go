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
	_ resource.Resource                = &InterfacePPPClientResource{}
	_ resource.ResourceWithImportState = &InterfacePPPClientResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfacePPPClientResource struct {
	reg *client.Registry
}

type InterfacePPPClientModel struct {
	ID                   types.String `tfsdk:"id"`
	UsePeerDns           types.String `tfsdk:"use_peer_dns"`
	Port                 types.String `tfsdk:"port"`
	Pin                  types.String `tfsdk:"pin"`
	Phone                types.String `tfsdk:"phone"`
	NullModem            types.String `tfsdk:"null_modem"`
	NetworkMode          types.String `tfsdk:"network_mode"`
	ModemInit            types.String `tfsdk:"modem_init"`
	InfoChannel          types.String `tfsdk:"info_channel"`
	DialCommand          types.String `tfsdk:"dial_command"`
	DataChannel          types.String `tfsdk:"data_channel"`
	Apn                  types.String `tfsdk:"apn"`
	AddDefaultRoute      types.String `tfsdk:"add_default_route"`
	Allow                types.String `tfsdk:"allow"`
	Comment              types.String `tfsdk:"comment"`
	DefaultRouteDistance types.String `tfsdk:"default_route_distance"`
	DialOnDemand         types.String `tfsdk:"dial_on_demand"`
	Disabled             types.Bool   `tfsdk:"disabled"`
	KeepaliveTimeout     types.String `tfsdk:"keepalive_timeout"`
	MaxMru               types.String `tfsdk:"max_mru"`
	MaxMTU               types.String `tfsdk:"max_mtu"`
	Mrru                 types.String `tfsdk:"mrru"`
	Name                 types.String `tfsdk:"name"`
	Password             types.String `tfsdk:"password"`
	Profile              types.String `tfsdk:"profile"`
	RemoteAddress        types.String `tfsdk:"remote_address"`
	User                 types.String `tfsdk:"user"`
	Router               types.String `tfsdk:"router"`
}

func NewInterfacePPPClientResource() resource.Resource { return &InterfacePPPClientResource{} }

func (r *InterfacePPPClientResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ppp_client"
}

func (r *InterfacePPPClientResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfacePPPClientResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/ppp-client`.",
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
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `port`.",
			},
			"pin": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `pin`.",
			},
			"phone": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `phone`.",
			},
			"null_modem": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `null-modem`.",
			},
			"network_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `network-mode`.",
			},
			"modem_init": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `modem-init`.",
			},
			"info_channel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `info-channel`.",
			},
			"dial_command": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dial-command`.",
			},
			"data_channel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `data-channel`.",
			},
			"apn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `apn`.",
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
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "",
			},
			"profile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"user": schema.StringAttribute{
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

func (r *InterfacePPPClientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfacePPPClientModel
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
	if !(plan.RemoteAddress.IsNull() || plan.RemoteAddress.IsUnknown()) {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !(plan.User.IsNull() || plan.User.IsUnknown()) {
		body["user"] = plan.User.ValueString()
	}
	if !(plan.AddDefaultRoute.IsNull() || plan.AddDefaultRoute.IsUnknown()) {
		body["add-default-route"] = plan.AddDefaultRoute.ValueString()
	}
	if !(plan.Apn.IsNull() || plan.Apn.IsUnknown()) {
		body["apn"] = plan.Apn.ValueString()
	}
	if !(plan.DataChannel.IsNull() || plan.DataChannel.IsUnknown()) {
		body["data-channel"] = plan.DataChannel.ValueString()
	}
	if !(plan.DialCommand.IsNull() || plan.DialCommand.IsUnknown()) {
		body["dial-command"] = plan.DialCommand.ValueString()
	}
	if !(plan.InfoChannel.IsNull() || plan.InfoChannel.IsUnknown()) {
		body["info-channel"] = plan.InfoChannel.ValueString()
	}
	if !(plan.ModemInit.IsNull() || plan.ModemInit.IsUnknown()) {
		body["modem-init"] = plan.ModemInit.ValueString()
	}
	if !(plan.NetworkMode.IsNull() || plan.NetworkMode.IsUnknown()) {
		body["network-mode"] = plan.NetworkMode.ValueString()
	}
	if !(plan.NullModem.IsNull() || plan.NullModem.IsUnknown()) {
		body["null-modem"] = plan.NullModem.ValueString()
	}
	if !(plan.Phone.IsNull() || plan.Phone.IsUnknown()) {
		body["phone"] = plan.Phone.ValueString()
	}
	if !(plan.Pin.IsNull() || plan.Pin.IsUnknown()) {
		body["pin"] = plan.Pin.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.UsePeerDns.IsNull() || plan.UsePeerDns.IsUnknown()) {
		body["use-peer-dns"] = plan.UsePeerDns.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/ppp-client", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/ppp-client failed", err.Error())
		return
	}
	interfacePPPClientApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfacePPPClientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfacePPPClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/ppp-client", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/ppp-client failed", err.Error())
		return
	}
	interfacePPPClientApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfacePPPClientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfacePPPClientModel
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
	if !plan.DefaultRouteDistance.Equal(state.DefaultRouteDistance) && !plan.DefaultRouteDistance.IsUnknown() {
		body["default-route-distance"] = plan.DefaultRouteDistance.ValueString()
	}
	if !plan.DialOnDemand.Equal(state.DialOnDemand) && !plan.DialOnDemand.IsUnknown() {
		body["dial-on-demand"] = plan.DialOnDemand.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
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
	if !plan.RemoteAddress.Equal(state.RemoteAddress) && !plan.RemoteAddress.IsUnknown() {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !plan.User.Equal(state.User) && !plan.User.IsUnknown() {
		body["user"] = plan.User.ValueString()
	}
	if !plan.AddDefaultRoute.Equal(state.AddDefaultRoute) && !plan.AddDefaultRoute.IsUnknown() {
		body["add-default-route"] = plan.AddDefaultRoute.ValueString()
	}
	if !plan.Apn.Equal(state.Apn) && !plan.Apn.IsUnknown() {
		body["apn"] = plan.Apn.ValueString()
	}
	if !plan.DataChannel.Equal(state.DataChannel) && !plan.DataChannel.IsUnknown() {
		body["data-channel"] = plan.DataChannel.ValueString()
	}
	if !plan.DialCommand.Equal(state.DialCommand) && !plan.DialCommand.IsUnknown() {
		body["dial-command"] = plan.DialCommand.ValueString()
	}
	if !plan.InfoChannel.Equal(state.InfoChannel) && !plan.InfoChannel.IsUnknown() {
		body["info-channel"] = plan.InfoChannel.ValueString()
	}
	if !plan.ModemInit.Equal(state.ModemInit) && !plan.ModemInit.IsUnknown() {
		body["modem-init"] = plan.ModemInit.ValueString()
	}
	if !plan.NetworkMode.Equal(state.NetworkMode) && !plan.NetworkMode.IsUnknown() {
		body["network-mode"] = plan.NetworkMode.ValueString()
	}
	if !plan.NullModem.Equal(state.NullModem) && !plan.NullModem.IsUnknown() {
		body["null-modem"] = plan.NullModem.ValueString()
	}
	if !plan.Phone.Equal(state.Phone) && !plan.Phone.IsUnknown() {
		body["phone"] = plan.Phone.ValueString()
	}
	if !plan.Pin.Equal(state.Pin) && !plan.Pin.IsUnknown() {
		body["pin"] = plan.Pin.ValueString()
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.UsePeerDns.Equal(state.UsePeerDns) && !plan.UsePeerDns.IsUnknown() {
		body["use-peer-dns"] = plan.UsePeerDns.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/ppp-client", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/ppp-client failed", err.Error())
			return
		}
		interfacePPPClientApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfacePPPClientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfacePPPClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/ppp-client", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/ppp-client failed", err.Error())
	}
}

func (r *InterfacePPPClientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfacePPPClientLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/ppp-client matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfacePPPClientLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfacePPPClientLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/ppp-client", id)
}

func interfacePPPClientApply(ctx context.Context, obj client.Object, m *InterfacePPPClientModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["use-peer-dns"]; ok && v != "" {
		m.UsePeerDns = types.StringValue(v)
	} else {
		m.UsePeerDns = types.StringNull()
	}
	if v, ok := obj["port"]; ok && v != "" {
		m.Port = types.StringValue(v)
	} else {
		m.Port = types.StringNull()
	}
	if v, ok := obj["pin"]; ok && v != "" {
		m.Pin = types.StringValue(v)
	} else {
		m.Pin = types.StringNull()
	}
	if v, ok := obj["phone"]; ok && v != "" {
		m.Phone = types.StringValue(v)
	} else {
		m.Phone = types.StringNull()
	}
	if v, ok := obj["null-modem"]; ok && v != "" {
		m.NullModem = types.StringValue(v)
	} else {
		m.NullModem = types.StringNull()
	}
	if v, ok := obj["network-mode"]; ok && v != "" {
		m.NetworkMode = types.StringValue(v)
	} else {
		m.NetworkMode = types.StringNull()
	}
	if v, ok := obj["modem-init"]; ok && v != "" {
		m.ModemInit = types.StringValue(v)
	} else {
		m.ModemInit = types.StringNull()
	}
	if v, ok := obj["info-channel"]; ok && v != "" {
		m.InfoChannel = types.StringValue(v)
	} else {
		m.InfoChannel = types.StringNull()
	}
	if v, ok := obj["dial-command"]; ok && v != "" {
		m.DialCommand = types.StringValue(v)
	} else {
		m.DialCommand = types.StringNull()
	}
	if v, ok := obj["data-channel"]; ok && v != "" {
		m.DataChannel = types.StringValue(v)
	} else {
		m.DataChannel = types.StringNull()
	}
	if v, ok := obj["apn"]; ok && v != "" {
		m.Apn = types.StringValue(v)
	} else {
		m.Apn = types.StringNull()
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
	if v, ok := obj["remote-address"]; ok {
		if v != "" {
			m.RemoteAddress = types.StringValue(v)
		} else {
			m.RemoteAddress = types.StringNull()
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
