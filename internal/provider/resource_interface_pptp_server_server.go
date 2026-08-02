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
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &InterfacePPTPServerServerResource{}
	_ resource.ResourceWithImportState = &InterfacePPTPServerServerResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfacePPTPServerServerResource struct {
	reg *client.Registry
}

type InterfacePPTPServerServerModel struct {
	ID               types.String `tfsdk:"id"`
	Authentication   csvSetValue  `tfsdk:"authentication"`
	DefaultProfile   types.String `tfsdk:"default_profile"`
	Enabled          types.Bool   `tfsdk:"enabled"`
	KeepaliveTimeout types.Int64  `tfsdk:"keepalive_timeout"`
	MaxMru           types.Int64  `tfsdk:"max_mru"`
	MaxMtu           types.Int64  `tfsdk:"max_mtu"`
	Mrru             types.String `tfsdk:"mrru"`
	Router           types.String `tfsdk:"router"`
}

func NewInterfacePPTPServerServerResource() resource.Resource {
	return &InterfacePPTPServerServerResource{}
}

func (r *InterfacePPTPServerServerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_pptp_server_server"
}

func (r *InterfacePPTPServerServerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfacePPTPServerServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/pptp-server/server`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"authentication": schema.StringAttribute{
				CustomType:  csvSetType{},
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `authentication`.",
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
			"keepalive_timeout": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `keepalive-timeout`.",
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
			"mrru": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mrru`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfacePPTPServerServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfacePPTPServerServerModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	interfacePPTPServerServerUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfacePPTPServerServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InterfacePPTPServerServerModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state InterfacePPTPServerServerModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	interfacePPTPServerServerUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfacePPTPServerServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfacePPTPServerServerModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/interface/pptp-server/server")
	if err != nil {
		resp.Diagnostics.AddError("Read /interface/pptp-server/server failed", err.Error())
		return
	}
	interfacePPTPServerServerApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/interface/pptp-server/server", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfacePPTPServerServerResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *InterfacePPTPServerServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/interface/pptp-server/server" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/interface/pptp-server/server", types.StringValue(routerName))))...)
}

func interfacePPTPServerServerUpsert(ctx context.Context, reg *client.Registry, plan, state *InterfacePPTPServerServerModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Authentication.IsNull() || plan.Authentication.IsUnknown()) && (state == nil || !plan.Authentication.Equal(state.Authentication)) {
		body["authentication"] = plan.Authentication.ValueString()
	}
	if !(plan.DefaultProfile.IsNull() || plan.DefaultProfile.IsUnknown()) && (state == nil || !plan.DefaultProfile.Equal(state.DefaultProfile)) {
		body["default-profile"] = plan.DefaultProfile.ValueString()
	}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	if !(plan.KeepaliveTimeout.IsNull() || plan.KeepaliveTimeout.IsUnknown()) && (state == nil || !plan.KeepaliveTimeout.Equal(state.KeepaliveTimeout)) {
		body["keepalive-timeout"] = client.FormatInt64(plan.KeepaliveTimeout.ValueInt64())
	}
	if !(plan.MaxMru.IsNull() || plan.MaxMru.IsUnknown()) && (state == nil || !plan.MaxMru.Equal(state.MaxMru)) {
		body["max-mru"] = client.FormatInt64(plan.MaxMru.ValueInt64())
	}
	if !(plan.MaxMtu.IsNull() || plan.MaxMtu.IsUnknown()) && (state == nil || !plan.MaxMtu.Equal(state.MaxMtu)) {
		body["max-mtu"] = client.FormatInt64(plan.MaxMtu.ValueInt64())
	}
	if !(plan.Mrru.IsNull() || plan.Mrru.IsUnknown()) && (state == nil || !plan.Mrru.Equal(state.Mrru)) {
		body["mrru"] = plan.Mrru.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/interface/pptp-server/server", body)
	if err != nil {
		diags.AddError("Upsert /interface/pptp-server/server failed", err.Error())
		return
	}
	interfacePPTPServerServerApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/interface/pptp-server/server", plan.Router))
}

func interfacePPTPServerServerApply(ctx context.Context, obj client.Object, m *InterfacePPTPServerServerModel) {
	_ = ctx
	if v, ok := obj["authentication"]; ok && v != "" {
		m.Authentication = newCSVSetValue(v)
	} else {
		m.Authentication = newCSVSetNull()
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
	if v, ok := obj["keepalive-timeout"]; ok {
		if n, err := client.ParseInt64(v); err == nil {
			m.KeepaliveTimeout = types.Int64Value(n)
		} else {
			m.KeepaliveTimeout = types.Int64Null()
		}
	} else {
		m.KeepaliveTimeout = types.Int64Null()
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
	if v, ok := obj["mrru"]; ok && v != "" {
		m.Mrru = types.StringValue(v)
	} else {
		m.Mrru = types.StringNull()
	}
}
