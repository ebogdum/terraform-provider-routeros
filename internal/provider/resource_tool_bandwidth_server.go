package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &ToolBandwidthServerResource{}
	_ resource.ResourceWithImportState = &ToolBandwidthServerResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type ToolBandwidthServerResource struct {
	reg *client.Registry
}

type ToolBandwidthServerModel struct {
	ID                   types.String  `tfsdk:"id"`
	AllocateUDPPortsFrom types.Int64   `tfsdk:"allocate_udp_ports_from"`
	AllowedAddresses4    hostAddrValue `tfsdk:"allowed_addresses4"`
	AllowedAddresses6    hostAddrValue `tfsdk:"allowed_addresses6"`
	Authenticate         types.Bool    `tfsdk:"authenticate"`
	Enabled              types.Bool    `tfsdk:"enabled"`
	MaxSessions          types.Int64   `tfsdk:"max_sessions"`
	Router               types.String  `tfsdk:"router"`
}

func NewToolBandwidthServerResource() resource.Resource { return &ToolBandwidthServerResource{} }

func (r *ToolBandwidthServerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_bandwidth_server"
}

func (r *ToolBandwidthServerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolBandwidthServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/tool/bandwidth-server`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"allocate_udp_ports_from": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "Beginning of UDP port range",
			},
			"allowed_addresses4": schema.StringAttribute{Optional: true, Computed: true, CustomType: hostAddrType{},
				Description: "",
			},
			"allowed_addresses6": schema.StringAttribute{Optional: true, Computed: true, CustomType: hostAddrType{},
				Description: "",
			},
			"authenticate": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "Communicate only with authenticated clients",
			},
			"enabled": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "Defines whether bandwidth server is enabled or not",
			},
			"max_sessions": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "Maximal simultaneous test count",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *ToolBandwidthServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolBandwidthServerModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolBandwidthServerUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolBandwidthServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ToolBandwidthServerModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state ToolBandwidthServerModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolBandwidthServerUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolBandwidthServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolBandwidthServerModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/tool/bandwidth-server")
	if err != nil {
		resp.Diagnostics.AddError("Read /tool/bandwidth-server failed", err.Error())
		return
	}
	toolBandwidthServerApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/tool/bandwidth-server", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolBandwidthServerResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *ToolBandwidthServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/tool/bandwidth-server" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/tool/bandwidth-server", types.StringValue(routerName))))...)
}

func toolBandwidthServerUpsert(ctx context.Context, reg *client.Registry, plan, state *ToolBandwidthServerModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AllocateUDPPortsFrom.IsNull() || plan.AllocateUDPPortsFrom.IsUnknown()) && (state == nil || !plan.AllocateUDPPortsFrom.Equal(state.AllocateUDPPortsFrom)) {
		body["allocate-udp-ports-from"] = client.FormatInt64(plan.AllocateUDPPortsFrom.ValueInt64())
	}
	if !(plan.AllowedAddresses4.IsNull() || plan.AllowedAddresses4.IsUnknown()) && (state == nil || !plan.AllowedAddresses4.Equal(state.AllowedAddresses4)) {
		body["allowed-addresses4"] = plan.AllowedAddresses4.ValueString()
	}
	if !(plan.AllowedAddresses6.IsNull() || plan.AllowedAddresses6.IsUnknown()) && (state == nil || !plan.AllowedAddresses6.Equal(state.AllowedAddresses6)) {
		body["allowed-addresses6"] = plan.AllowedAddresses6.ValueString()
	}
	if !(plan.Authenticate.IsNull() || plan.Authenticate.IsUnknown()) && (state == nil || !plan.Authenticate.Equal(state.Authenticate)) {
		body["authenticate"] = client.FormatBool(plan.Authenticate.ValueBool())
	}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	if !(plan.MaxSessions.IsNull() || plan.MaxSessions.IsUnknown()) && (state == nil || !plan.MaxSessions.Equal(state.MaxSessions)) {
		body["max-sessions"] = client.FormatInt64(plan.MaxSessions.ValueInt64())
	}
	obj, err := c.SetSingleton(ctx, "/tool/bandwidth-server", body)
	if err != nil {
		diags.AddError("Upsert /tool/bandwidth-server failed", err.Error())
		return
	}
	toolBandwidthServerApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/tool/bandwidth-server", plan.Router))
}

func toolBandwidthServerApply(ctx context.Context, obj client.Object, m *ToolBandwidthServerModel) {
	_ = ctx
	if v, ok := obj["allocate-udp-ports-from"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.AllocateUDPPortsFrom = types.Int64Value(n)
		} else {
			m.AllocateUDPPortsFrom = types.Int64Null()
		}
	}
	if v, ok := obj["allowed-addresses4"]; ok {
		_ = v
		if v != "" {
			m.AllowedAddresses4 = newHostAddrValue(v)
		} else {
			m.AllowedAddresses4 = newHostAddrNull()
		}
	}
	if v, ok := obj["allowed-addresses6"]; ok {
		_ = v
		if v != "" {
			m.AllowedAddresses6 = newHostAddrValue(v)
		} else {
			m.AllowedAddresses6 = newHostAddrNull()
		}
	}
	if v, ok := obj["authenticate"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Authenticate = types.BoolValue(b)
		} else {
			m.Authenticate = types.BoolNull()
		}
	}
	if v, ok := obj["enabled"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Enabled = types.BoolValue(b)
		} else {
			m.Enabled = types.BoolNull()
		}
	}
	if v, ok := obj["max-sessions"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MaxSessions = types.Int64Value(n)
		} else {
			m.MaxSessions = types.Int64Null()
		}
	}
}
