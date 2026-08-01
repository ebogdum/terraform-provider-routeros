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
	_ resource.Resource                = &InterfaceL2TPEtherResource{}
	_ resource.ResourceWithImportState = &InterfaceL2TPEtherResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceL2TPEtherResource struct {
	reg *client.Registry
}

type InterfaceL2TPEtherModel struct {
	ID                    types.String `tfsdk:"id"`
	UseL2SpecificSublayer types.String `tfsdk:"use_l2_specific_sublayer"`
	UseIpsec              types.String `tfsdk:"use_ipsec"`
	UnmanagedMode         types.String `tfsdk:"unmanaged_mode"`
	SendCookie            types.String `tfsdk:"send_cookie"`
	RemoteTunnelId        types.String `tfsdk:"remote_tunnel_id"`
	RemoteSessionId       types.String `tfsdk:"remote_session_id"`
	PeerCookie            types.String `tfsdk:"peer_cookie"`
	LocalTunnelId         types.String `tfsdk:"local_tunnel_id"`
	LocalSessionId        types.String `tfsdk:"local_session_id"`
	L2tpProtoVersion      types.String `tfsdk:"l2tp_proto_version"`
	DigestHash            types.String `tfsdk:"digest_hash"`
	CookieLength          types.String `tfsdk:"cookie_length"`
	ConnectTo             types.String `tfsdk:"connect_to"`
	CircuitId             types.String `tfsdk:"circuit_id"`
	AllowFastPath         types.String `tfsdk:"allow_fast_path"`
	Comment               types.String `tfsdk:"comment"`
	Disabled              types.Bool   `tfsdk:"disabled"`
	IpsecSecret           types.String `tfsdk:"ipsec_secret"`
	LocalAddress          types.String `tfsdk:"local_address"`
	MACAddress            types.String `tfsdk:"mac_address"`
	MTU                   types.String `tfsdk:"mtu"`
	Name                  types.String `tfsdk:"name"`
	Router                types.String `tfsdk:"router"`
}

func NewInterfaceL2TPEtherResource() resource.Resource { return &InterfaceL2TPEtherResource{} }

func (r *InterfaceL2TPEtherResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_l2tp_ether"
}

func (r *InterfaceL2TPEtherResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceL2TPEtherResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/l2tp-ether`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"use_l2_specific_sublayer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `use-l2-specific-sublayer`.",
			},
			"use_ipsec": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `use-ipsec`.",
			},
			"unmanaged_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `unmanaged-mode`.",
			},
			"send_cookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `send-cookie`.",
			},
			"remote_tunnel_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `remote-tunnel-id`.",
			},
			"remote_session_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `remote-session-id`.",
			},
			"peer_cookie": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `peer-cookie`.",
			},
			"local_tunnel_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `local-tunnel-id`.",
			},
			"local_session_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `local-session-id`.",
			},
			"l2tp_proto_version": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2tp-proto-version`.",
			},
			"digest_hash": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `digest-hash`.",
			},
			"cookie_length": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cookie-length`.",
			},
			"connect_to": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `connect-to`.",
			},
			"circuit_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `circuit-id`.",
			},
			"allow_fast_path": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `allow-fast-path`.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
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
			"local_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
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

func (r *InterfaceL2TPEtherResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceL2TPEtherModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.IpsecSecret.IsNull() || plan.IpsecSecret.IsUnknown()) {
		body["ipsec-secret"] = plan.IpsecSecret.ValueString()
	}
	if !(plan.LocalAddress.IsNull() || plan.LocalAddress.IsUnknown()) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !(plan.MACAddress.IsNull() || plan.MACAddress.IsUnknown()) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.AllowFastPath.IsNull() || plan.AllowFastPath.IsUnknown()) {
		body["allow-fast-path"] = plan.AllowFastPath.ValueString()
	}
	if !(plan.CircuitId.IsNull() || plan.CircuitId.IsUnknown()) {
		body["circuit-id"] = plan.CircuitId.ValueString()
	}
	if !(plan.ConnectTo.IsNull() || plan.ConnectTo.IsUnknown()) {
		body["connect-to"] = plan.ConnectTo.ValueString()
	}
	if !(plan.CookieLength.IsNull() || plan.CookieLength.IsUnknown()) {
		body["cookie-length"] = plan.CookieLength.ValueString()
	}
	if !(plan.DigestHash.IsNull() || plan.DigestHash.IsUnknown()) {
		body["digest-hash"] = plan.DigestHash.ValueString()
	}
	if !(plan.L2tpProtoVersion.IsNull() || plan.L2tpProtoVersion.IsUnknown()) {
		body["l2tp-proto-version"] = plan.L2tpProtoVersion.ValueString()
	}
	if !(plan.LocalSessionId.IsNull() || plan.LocalSessionId.IsUnknown()) {
		body["local-session-id"] = plan.LocalSessionId.ValueString()
	}
	if !(plan.LocalTunnelId.IsNull() || plan.LocalTunnelId.IsUnknown()) {
		body["local-tunnel-id"] = plan.LocalTunnelId.ValueString()
	}
	if !(plan.PeerCookie.IsNull() || plan.PeerCookie.IsUnknown()) {
		body["peer-cookie"] = plan.PeerCookie.ValueString()
	}
	if !(plan.RemoteSessionId.IsNull() || plan.RemoteSessionId.IsUnknown()) {
		body["remote-session-id"] = plan.RemoteSessionId.ValueString()
	}
	if !(plan.RemoteTunnelId.IsNull() || plan.RemoteTunnelId.IsUnknown()) {
		body["remote-tunnel-id"] = plan.RemoteTunnelId.ValueString()
	}
	if !(plan.SendCookie.IsNull() || plan.SendCookie.IsUnknown()) {
		body["send-cookie"] = plan.SendCookie.ValueString()
	}
	if !(plan.UnmanagedMode.IsNull() || plan.UnmanagedMode.IsUnknown()) {
		body["unmanaged-mode"] = plan.UnmanagedMode.ValueString()
	}
	if !(plan.UseIpsec.IsNull() || plan.UseIpsec.IsUnknown()) {
		body["use-ipsec"] = plan.UseIpsec.ValueString()
	}
	if !(plan.UseL2SpecificSublayer.IsNull() || plan.UseL2SpecificSublayer.IsUnknown()) {
		body["use-l2-specific-sublayer"] = plan.UseL2SpecificSublayer.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/l2tp-ether", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/l2tp-ether failed", err.Error())
		return
	}
	interfaceL2TPEtherApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceL2TPEtherResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceL2TPEtherModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/l2tp-ether", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/l2tp-ether failed", err.Error())
		return
	}
	interfaceL2TPEtherApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceL2TPEtherResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceL2TPEtherModel
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
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.IpsecSecret.Equal(state.IpsecSecret) && !plan.IpsecSecret.IsUnknown() {
		body["ipsec-secret"] = plan.IpsecSecret.ValueString()
	}
	if !plan.LocalAddress.Equal(state.LocalAddress) && !plan.LocalAddress.IsUnknown() {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !plan.MACAddress.Equal(state.MACAddress) && !plan.MACAddress.IsUnknown() {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.MTU.Equal(state.MTU) && !plan.MTU.IsUnknown() {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.AllowFastPath.Equal(state.AllowFastPath) && !plan.AllowFastPath.IsUnknown() {
		body["allow-fast-path"] = plan.AllowFastPath.ValueString()
	}
	if !plan.CircuitId.Equal(state.CircuitId) && !plan.CircuitId.IsUnknown() {
		body["circuit-id"] = plan.CircuitId.ValueString()
	}
	if !plan.ConnectTo.Equal(state.ConnectTo) && !plan.ConnectTo.IsUnknown() {
		body["connect-to"] = plan.ConnectTo.ValueString()
	}
	if !plan.CookieLength.Equal(state.CookieLength) && !plan.CookieLength.IsUnknown() {
		body["cookie-length"] = plan.CookieLength.ValueString()
	}
	if !plan.DigestHash.Equal(state.DigestHash) && !plan.DigestHash.IsUnknown() {
		body["digest-hash"] = plan.DigestHash.ValueString()
	}
	if !plan.L2tpProtoVersion.Equal(state.L2tpProtoVersion) && !plan.L2tpProtoVersion.IsUnknown() {
		body["l2tp-proto-version"] = plan.L2tpProtoVersion.ValueString()
	}
	if !plan.LocalSessionId.Equal(state.LocalSessionId) && !plan.LocalSessionId.IsUnknown() {
		body["local-session-id"] = plan.LocalSessionId.ValueString()
	}
	if !plan.LocalTunnelId.Equal(state.LocalTunnelId) && !plan.LocalTunnelId.IsUnknown() {
		body["local-tunnel-id"] = plan.LocalTunnelId.ValueString()
	}
	if !plan.PeerCookie.Equal(state.PeerCookie) && !plan.PeerCookie.IsUnknown() {
		body["peer-cookie"] = plan.PeerCookie.ValueString()
	}
	if !plan.RemoteSessionId.Equal(state.RemoteSessionId) && !plan.RemoteSessionId.IsUnknown() {
		body["remote-session-id"] = plan.RemoteSessionId.ValueString()
	}
	if !plan.RemoteTunnelId.Equal(state.RemoteTunnelId) && !plan.RemoteTunnelId.IsUnknown() {
		body["remote-tunnel-id"] = plan.RemoteTunnelId.ValueString()
	}
	if !plan.SendCookie.Equal(state.SendCookie) && !plan.SendCookie.IsUnknown() {
		body["send-cookie"] = plan.SendCookie.ValueString()
	}
	if !plan.UnmanagedMode.Equal(state.UnmanagedMode) && !plan.UnmanagedMode.IsUnknown() {
		body["unmanaged-mode"] = plan.UnmanagedMode.ValueString()
	}
	if !plan.UseIpsec.Equal(state.UseIpsec) && !plan.UseIpsec.IsUnknown() {
		body["use-ipsec"] = plan.UseIpsec.ValueString()
	}
	if !plan.UseL2SpecificSublayer.Equal(state.UseL2SpecificSublayer) && !plan.UseL2SpecificSublayer.IsUnknown() {
		body["use-l2-specific-sublayer"] = plan.UseL2SpecificSublayer.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/l2tp-ether", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/l2tp-ether failed", err.Error())
			return
		}
		interfaceL2TPEtherApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceL2TPEtherResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceL2TPEtherModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/l2tp-ether", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/l2tp-ether failed", err.Error())
	}
}

func (r *InterfaceL2TPEtherResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceL2TPEtherLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/l2tp-ether matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceL2TPEtherLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceL2TPEtherLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/l2tp-ether", id)
}

func interfaceL2TPEtherApply(ctx context.Context, obj client.Object, m *InterfaceL2TPEtherModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["use-l2-specific-sublayer"]; ok && v != "" {
		m.UseL2SpecificSublayer = types.StringValue(v)
	} else {
		m.UseL2SpecificSublayer = types.StringNull()
	}
	if v, ok := obj["use-ipsec"]; ok && v != "" {
		m.UseIpsec = types.StringValue(v)
	} else {
		m.UseIpsec = types.StringNull()
	}
	if v, ok := obj["unmanaged-mode"]; ok && v != "" {
		m.UnmanagedMode = types.StringValue(v)
	} else {
		m.UnmanagedMode = types.StringNull()
	}
	if v, ok := obj["send-cookie"]; ok && v != "" {
		m.SendCookie = types.StringValue(v)
	} else {
		m.SendCookie = types.StringNull()
	}
	if v, ok := obj["remote-tunnel-id"]; ok && v != "" {
		m.RemoteTunnelId = types.StringValue(v)
	} else {
		m.RemoteTunnelId = types.StringNull()
	}
	if v, ok := obj["remote-session-id"]; ok && v != "" {
		m.RemoteSessionId = types.StringValue(v)
	} else {
		m.RemoteSessionId = types.StringNull()
	}
	if v, ok := obj["peer-cookie"]; ok && v != "" {
		m.PeerCookie = types.StringValue(v)
	} else {
		m.PeerCookie = types.StringNull()
	}
	if v, ok := obj["local-tunnel-id"]; ok && v != "" {
		m.LocalTunnelId = types.StringValue(v)
	} else {
		m.LocalTunnelId = types.StringNull()
	}
	if v, ok := obj["local-session-id"]; ok && v != "" {
		m.LocalSessionId = types.StringValue(v)
	} else {
		m.LocalSessionId = types.StringNull()
	}
	if v, ok := obj["l2tp-proto-version"]; ok && v != "" {
		m.L2tpProtoVersion = types.StringValue(v)
	} else {
		m.L2tpProtoVersion = types.StringNull()
	}
	if v, ok := obj["digest-hash"]; ok && v != "" {
		m.DigestHash = types.StringValue(v)
	} else {
		m.DigestHash = types.StringNull()
	}
	if v, ok := obj["cookie-length"]; ok && v != "" {
		m.CookieLength = types.StringValue(v)
	} else {
		m.CookieLength = types.StringNull()
	}
	if v, ok := obj["connect-to"]; ok && v != "" {
		m.ConnectTo = types.StringValue(v)
	} else {
		m.ConnectTo = types.StringNull()
	}
	if v, ok := obj["circuit-id"]; ok && v != "" {
		m.CircuitId = types.StringValue(v)
	} else {
		m.CircuitId = types.StringNull()
	}
	if v, ok := obj["allow-fast-path"]; ok && v != "" {
		m.AllowFastPath = types.StringValue(v)
	} else {
		m.AllowFastPath = types.StringNull()
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
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
	if v, ok := obj["local-address"]; ok {
		if v != "" {
			m.LocalAddress = types.StringValue(v)
		} else {
			m.LocalAddress = types.StringNull()
		}
	}
	if v, ok := obj["mac-address"]; ok {
		if v != "" {
			m.MACAddress = types.StringValue(v)
		} else {
			m.MACAddress = types.StringNull()
		}
	}
	if v, ok := obj["mtu"]; ok {
		if v != "" {
			m.MTU = types.StringValue(v)
		} else {
			m.MTU = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
}
