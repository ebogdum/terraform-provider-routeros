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
	_ resource.Resource                = &IPSocksResource{}
	_ resource.ResourceWithImportState = &IPSocksResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type IPSocksResource struct {
	reg *client.Registry
}

type IPSocksModel struct {
	ID                    types.String `tfsdk:"id"`
	AuthMethod            types.String `tfsdk:"auth_method"`
	ConnectionIdleTimeout types.String `tfsdk:"connection_idle_timeout"`
	Enabled               types.Bool   `tfsdk:"enabled"`
	MaxConnections        types.Int64  `tfsdk:"max_connections"`
	Port                  types.Int64  `tfsdk:"port"`
	Version               types.Int64  `tfsdk:"version"`
	Vrf                   types.String `tfsdk:"vrf"`
	Router                types.String `tfsdk:"router"`
}

func NewIPSocksResource() resource.Resource { return &IPSocksResource{} }

func (r *IPSocksResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_socks"
}

func (r *IPSocksResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPSocksResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/socks`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"auth_method": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"connection_idle_timeout": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"enabled": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"max_connections": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"port": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"version": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"vrf": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *IPSocksResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPSocksModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPSocksUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPSocksResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPSocksModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state IPSocksModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPSocksUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPSocksResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPSocksModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/socks")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/socks failed", err.Error())
		return
	}
	iPSocksApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/socks", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPSocksResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPSocksResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/socks" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ip/socks", types.StringValue(routerName))))...)
}

func iPSocksUpsert(ctx context.Context, reg *client.Registry, plan, state *IPSocksModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AuthMethod.IsNull() || plan.AuthMethod.IsUnknown()) && (state == nil || !plan.AuthMethod.Equal(state.AuthMethod)) {
		body["auth-method"] = plan.AuthMethod.ValueString()
	}
	if !(plan.ConnectionIdleTimeout.IsNull() || plan.ConnectionIdleTimeout.IsUnknown()) && (state == nil || !plan.ConnectionIdleTimeout.Equal(state.ConnectionIdleTimeout)) {
		body["connection-idle-timeout"] = plan.ConnectionIdleTimeout.ValueString()
	}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	if !(plan.MaxConnections.IsNull() || plan.MaxConnections.IsUnknown()) && (state == nil || !plan.MaxConnections.Equal(state.MaxConnections)) {
		body["max-connections"] = client.FormatInt64(plan.MaxConnections.ValueInt64())
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) && (state == nil || !plan.Port.Equal(state.Port)) {
		body["port"] = client.FormatInt64(plan.Port.ValueInt64())
	}
	if !(plan.Version.IsNull() || plan.Version.IsUnknown()) && (state == nil || !plan.Version.Equal(state.Version)) {
		body["version"] = client.FormatInt64(plan.Version.ValueInt64())
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) && (state == nil || !plan.Vrf.Equal(state.Vrf)) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/ip/socks", body)
	if err != nil {
		diags.AddError("Upsert /ip/socks failed", err.Error())
		return
	}
	iPSocksApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/socks", plan.Router))
}

func iPSocksApply(ctx context.Context, obj client.Object, m *IPSocksModel) {
	_ = ctx
	if v, ok := obj["auth-method"]; ok {
		_ = v
		if v != "" {
			m.AuthMethod = types.StringValue(v)
		} else {
			m.AuthMethod = types.StringNull()
		}
	}
	if v, ok := obj["connection-idle-timeout"]; ok {
		_ = v
		if v != "" {
			m.ConnectionIdleTimeout = types.StringValue(v)
		} else {
			m.ConnectionIdleTimeout = types.StringNull()
		}
	}
	if v, ok := obj["enabled"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Enabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Enabled = types.BoolValue(true)
		} else {
			m.Enabled = types.BoolNull()
		}
	}
	if v, ok := obj["max-connections"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MaxConnections = types.Int64Value(n)
		} else {
			m.MaxConnections = types.Int64Null()
		}
	}
	if v, ok := obj["port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Port = types.Int64Value(n)
		} else {
			m.Port = types.Int64Null()
		}
	}
	if v, ok := obj["version"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Version = types.Int64Value(n)
		} else {
			m.Version = types.Int64Null()
		}
	}
	if v, ok := obj["vrf"]; ok {
		_ = v
		if v != "" {
			m.Vrf = types.StringValue(v)
		} else {
			m.Vrf = types.StringNull()
		}
	}
}
