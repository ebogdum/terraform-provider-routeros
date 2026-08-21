package provider

import (
	"context"
	"fmt"
	"strings"

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
	_ resource.Resource                = &InterfaceL2TPServerServerResource{}
	_ resource.ResourceWithImportState = &InterfaceL2TPServerServerResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceL2TPServerServerResource struct {
	reg *client.Registry
}

type InterfaceL2TPServerServerModel struct {
	ID                       types.String `tfsdk:"id"`
	L2tpv3EtherInterfaceList types.String `tfsdk:"l2tpv3_ether_interface_list"`
	AcceptProtoVersion       types.String `tfsdk:"accept_proto_version"`
	AcceptPseudowireType     types.String `tfsdk:"accept_pseudowire_type"`
	AllowFastPath            types.Bool   `tfsdk:"allow_fast_path"`
	Authentication           csvSetValue  `tfsdk:"authentication"`
	CallerIDType             types.String `tfsdk:"caller_id_type"`
	DefaultProfile           types.String `tfsdk:"default_profile"`
	Enabled                  types.Bool   `tfsdk:"enabled"`
	IPsecSecret              types.String `tfsdk:"ipsec_secret"`
	KeepaliveTimeout         types.Int64  `tfsdk:"keepalive_timeout"`
	L2tpv3CircuitID          types.String `tfsdk:"l2tpv3_circuit_id"`
	L2tpv3CookieLength       types.Int64  `tfsdk:"l2tpv3_cookie_length"`
	L2tpv3DigestHash         types.String `tfsdk:"l2tpv3_digest_hash"`
	MaxMru                   types.Int64  `tfsdk:"max_mru"`
	MaxMtu                   types.Int64  `tfsdk:"max_mtu"`
	MaxSessions              types.String `tfsdk:"max_sessions"`
	Mrru                     types.String `tfsdk:"mrru"`
	OneSessionPerHost        types.Bool   `tfsdk:"one_session_per_host"`
	UseIPsec                 types.String `tfsdk:"use_ipsec"`
	Router                   types.String `tfsdk:"router"`
}

func NewInterfaceL2TPServerServerResource() resource.Resource {
	return &InterfaceL2TPServerServerResource{}
}

func (r *InterfaceL2TPServerServerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_l2tp_server_server"
}

func (r *InterfaceL2TPServerServerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceL2TPServerServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/l2tp-server/server`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"l2tpv3_ether_interface_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2tpv3-ether-interface-list`.",
			},
			"accept_proto_version": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `accept-proto-version`.",
			},
			"accept_pseudowire_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `accept-pseudowire-type`.",
			},
			"allow_fast_path": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `allow-fast-path`.",
			},
			"authentication": schema.StringAttribute{
				CustomType:  csvSetType{},
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `authentication`.",
			},
			"caller_id_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `caller-id-type`.",
			},
			"default_profile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `default-profile`.",
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `enabled`.",
			},
			"ipsec_secret": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "RouterOS `ipsec-secret`.",
			},
			"keepalive_timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `keepalive-timeout`.",
			},
			"l2tpv3_circuit_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2tpv3-circuit-id`.",
			},
			"l2tpv3_cookie_length": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2tpv3-cookie-length`.",
			},
			"l2tpv3_digest_hash": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `l2tpv3-digest-hash`.",
			},
			"max_mru": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `max-mru`.",
			},
			"max_mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `max-mtu`.",
			},
			"max_sessions": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `max-sessions`.",
			},
			"mrru": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mrru`.",
			},
			"one_session_per_host": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `one-session-per-host`.",
			},
			"use_ipsec": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPsec usage for the L2TP server: `no`, `yes` (offer IPsec) or `required` (refuse plain L2TP).",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "required", "yes"}...)},
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *InterfaceL2TPServerServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceL2TPServerServerModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	interfaceL2TPServerServerUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceL2TPServerServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InterfaceL2TPServerServerModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state InterfaceL2TPServerServerModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	interfaceL2TPServerServerUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceL2TPServerServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceL2TPServerServerModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/interface/l2tp-server/server")
	if err != nil {
		resp.Diagnostics.AddError("Read /interface/l2tp-server/server failed", err.Error())
		return
	}
	interfaceL2TPServerServerApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/interface/l2tp-server/server", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceL2TPServerServerResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *InterfaceL2TPServerServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/interface/l2tp-server/server" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/interface/l2tp-server/server", types.StringValue(routerName))))...)
}

func interfaceL2TPServerServerUpsert(ctx context.Context, reg *client.Registry, plan, state *InterfaceL2TPServerServerModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AcceptProtoVersion.IsNull() || plan.AcceptProtoVersion.IsUnknown()) && (state == nil || !plan.AcceptProtoVersion.Equal(state.AcceptProtoVersion)) {
		body["accept-proto-version"] = plan.AcceptProtoVersion.ValueString()
	}
	if !(plan.AcceptPseudowireType.IsNull() || plan.AcceptPseudowireType.IsUnknown()) && (state == nil || !plan.AcceptPseudowireType.Equal(state.AcceptPseudowireType)) {
		body["accept-pseudowire-type"] = plan.AcceptPseudowireType.ValueString()
	}
	if !(plan.AllowFastPath.IsNull() || plan.AllowFastPath.IsUnknown()) && (state == nil || !plan.AllowFastPath.Equal(state.AllowFastPath)) {
		body["allow-fast-path"] = client.FormatBool(plan.AllowFastPath.ValueBool())
	}
	if !(plan.Authentication.IsNull() || plan.Authentication.IsUnknown()) && (state == nil || !plan.Authentication.Equal(state.Authentication)) {
		body["authentication"] = plan.Authentication.ValueString()
	}
	if !(plan.CallerIDType.IsNull() || plan.CallerIDType.IsUnknown()) && (state == nil || !plan.CallerIDType.Equal(state.CallerIDType)) {
		body["caller-id-type"] = plan.CallerIDType.ValueString()
	}
	if !(plan.DefaultProfile.IsNull() || plan.DefaultProfile.IsUnknown()) && (state == nil || !plan.DefaultProfile.Equal(state.DefaultProfile)) {
		body["default-profile"] = plan.DefaultProfile.ValueString()
	}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	if !(plan.IPsecSecret.IsNull() || plan.IPsecSecret.IsUnknown()) && (state == nil || !plan.IPsecSecret.Equal(state.IPsecSecret)) {
		body["ipsec-secret"] = plan.IPsecSecret.ValueString()
	}
	if !(plan.KeepaliveTimeout.IsNull() || plan.KeepaliveTimeout.IsUnknown()) && (state == nil || !plan.KeepaliveTimeout.Equal(state.KeepaliveTimeout)) {
		body["keepalive-timeout"] = client.FormatInt64(plan.KeepaliveTimeout.ValueInt64())
	}
	if !(plan.L2tpv3CircuitID.IsNull() || plan.L2tpv3CircuitID.IsUnknown()) && (state == nil || !plan.L2tpv3CircuitID.Equal(state.L2tpv3CircuitID)) {
		body["l2tpv3-circuit-id"] = plan.L2tpv3CircuitID.ValueString()
	}
	if !(plan.L2tpv3CookieLength.IsNull() || plan.L2tpv3CookieLength.IsUnknown()) && (state == nil || !plan.L2tpv3CookieLength.Equal(state.L2tpv3CookieLength)) {
		body["l2tpv3-cookie-length"] = client.FormatInt64(plan.L2tpv3CookieLength.ValueInt64())
	}
	if !(plan.L2tpv3DigestHash.IsNull() || plan.L2tpv3DigestHash.IsUnknown()) && (state == nil || !plan.L2tpv3DigestHash.Equal(state.L2tpv3DigestHash)) {
		body["l2tpv3-digest-hash"] = plan.L2tpv3DigestHash.ValueString()
	}
	if !(plan.MaxMru.IsNull() || plan.MaxMru.IsUnknown()) && (state == nil || !plan.MaxMru.Equal(state.MaxMru)) {
		body["max-mru"] = client.FormatInt64(plan.MaxMru.ValueInt64())
	}
	if !(plan.MaxMtu.IsNull() || plan.MaxMtu.IsUnknown()) && (state == nil || !plan.MaxMtu.Equal(state.MaxMtu)) {
		body["max-mtu"] = client.FormatInt64(plan.MaxMtu.ValueInt64())
	}
	if !(plan.MaxSessions.IsNull() || plan.MaxSessions.IsUnknown()) && (state == nil || !plan.MaxSessions.Equal(state.MaxSessions)) {
		body["max-sessions"] = plan.MaxSessions.ValueString()
	}
	if !(plan.Mrru.IsNull() || plan.Mrru.IsUnknown()) && (state == nil || !plan.Mrru.Equal(state.Mrru)) {
		body["mrru"] = plan.Mrru.ValueString()
	}
	if !(plan.OneSessionPerHost.IsNull() || plan.OneSessionPerHost.IsUnknown()) && (state == nil || !plan.OneSessionPerHost.Equal(state.OneSessionPerHost)) {
		body["one-session-per-host"] = client.FormatBool(plan.OneSessionPerHost.ValueBool())
	}
	if !(plan.UseIPsec.IsNull() || plan.UseIPsec.IsUnknown()) && (state == nil || !plan.UseIPsec.Equal(state.UseIPsec)) {
		body["use-ipsec"] = plan.UseIPsec.ValueString()
	}
	if !(plan.L2tpv3EtherInterfaceList.IsNull() || plan.L2tpv3EtherInterfaceList.IsUnknown()) && (state == nil || !plan.L2tpv3EtherInterfaceList.Equal(state.L2tpv3EtherInterfaceList)) {
		body["l2tpv3-ether-interface-list"] = plan.L2tpv3EtherInterfaceList.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/interface/l2tp-server/server", body)
	if err != nil {
		diags.AddError("Upsert /interface/l2tp-server/server failed", err.Error())
		return
	}
	interfaceL2TPServerServerApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/interface/l2tp-server/server", plan.Router))
}

func interfaceL2TPServerServerApply(ctx context.Context, obj client.Object, m *InterfaceL2TPServerServerModel) {
	_ = ctx
	if v, ok := obj["l2tpv3-ether-interface-list"]; ok && v != "" {
		m.L2tpv3EtherInterfaceList = types.StringValue(v)
	} else {
		m.L2tpv3EtherInterfaceList = types.StringNull()
	}
	if v, ok := obj["accept-proto-version"]; ok && v != "" {
		m.AcceptProtoVersion = types.StringValue(v)
	} else {
		m.AcceptProtoVersion = types.StringNull()
	}
	if v, ok := obj["accept-pseudowire-type"]; ok && v != "" {
		m.AcceptPseudowireType = types.StringValue(v)
	} else {
		m.AcceptPseudowireType = types.StringNull()
	}
	if v, ok := obj["allow-fast-path"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.AllowFastPath = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.AllowFastPath = types.BoolValue(true)
		} else {
			m.AllowFastPath = types.BoolNull()
		}
	} else {
		m.AllowFastPath = types.BoolNull()
	}
	if v, ok := obj["authentication"]; ok && v != "" {
		m.Authentication = newCSVSetValue(v)
	} else {
		m.Authentication = newCSVSetNull()
	}
	if v, ok := obj["caller-id-type"]; ok && v != "" {
		m.CallerIDType = types.StringValue(v)
	} else {
		m.CallerIDType = types.StringNull()
	}
	if v, ok := obj["default-profile"]; ok && v != "" {
		m.DefaultProfile = types.StringValue(v)
	} else {
		m.DefaultProfile = types.StringNull()
	}
	if v, ok := obj["enabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Enabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Enabled = types.BoolValue(true)
		} else {
			m.Enabled = types.BoolNull()
		}
	} else {
		m.Enabled = types.BoolNull()
	}
	if v, ok := obj["ipsec-secret"]; ok && v != "" {
		m.IPsecSecret = types.StringValue(v)
	} else {
		m.IPsecSecret = types.StringNull()
	}
	if v, ok := obj["keepalive-timeout"]; ok {
		if n, err := client.ParseInt64(v); err == nil {
			m.KeepaliveTimeout = types.Int64Value(n)
		} else {
			m.KeepaliveTimeout = types.Int64Null()
		}
	} else {
		m.KeepaliveTimeout = types.Int64Null()
	}
	if v, ok := obj["l2tpv3-circuit-id"]; ok && v != "" {
		m.L2tpv3CircuitID = types.StringValue(v)
	} else {
		m.L2tpv3CircuitID = types.StringNull()
	}
	if v, ok := obj["l2tpv3-cookie-length"]; ok {
		if n, err := client.ParseInt64(v); err == nil {
			m.L2tpv3CookieLength = types.Int64Value(n)
		} else {
			m.L2tpv3CookieLength = types.Int64Null()
		}
	} else {
		m.L2tpv3CookieLength = types.Int64Null()
	}
	if v, ok := obj["l2tpv3-digest-hash"]; ok && v != "" {
		m.L2tpv3DigestHash = types.StringValue(v)
	} else {
		m.L2tpv3DigestHash = types.StringNull()
	}
	if v, ok := obj["max-mru"]; ok {
		if n, err := client.ParseInt64(v); err == nil {
			m.MaxMru = types.Int64Value(n)
		} else {
			m.MaxMru = types.Int64Null()
		}
	} else {
		m.MaxMru = types.Int64Null()
	}
	if v, ok := obj["max-mtu"]; ok {
		if n, err := client.ParseInt64(v); err == nil {
			m.MaxMtu = types.Int64Value(n)
		} else {
			m.MaxMtu = types.Int64Null()
		}
	} else {
		m.MaxMtu = types.Int64Null()
	}
	if v, ok := obj["max-sessions"]; ok && v != "" {
		m.MaxSessions = types.StringValue(v)
	} else {
		m.MaxSessions = types.StringNull()
	}
	if v, ok := obj["mrru"]; ok && v != "" {
		m.Mrru = types.StringValue(v)
	} else {
		m.Mrru = types.StringNull()
	}
	if v, ok := obj["one-session-per-host"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.OneSessionPerHost = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.OneSessionPerHost = types.BoolValue(true)
		} else {
			m.OneSessionPerHost = types.BoolNull()
		}
	} else {
		m.OneSessionPerHost = types.BoolNull()
	}
	if v, ok := obj["use-ipsec"]; ok {
		if v != "" {
			m.UseIPsec = types.StringValue(v)
		} else {
			m.UseIPsec = types.StringNull()
		}
	} else {
		m.UseIPsec = types.StringNull()
	}
}
