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
	_ resource.Resource                = &InterfaceBridgeSettingsResource{}
	_ resource.ResourceWithImportState = &InterfaceBridgeSettingsResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type InterfaceBridgeSettingsResource struct {
	reg *client.Registry
}

type InterfaceBridgeSettingsModel struct {
	ID                       types.String `tfsdk:"id"`
	AllowFastPath            types.Bool   `tfsdk:"allow_fast_path"`
	BridgeFastForwardBytes   types.Int64  `tfsdk:"bridge_fast_forward_bytes"`
	BridgeFastForwardPackets types.Int64  `tfsdk:"bridge_fast_forward_packets"`
	BridgeFastPathActive     types.Bool   `tfsdk:"bridge_fast_path_active"`
	BridgeFastPathBytes      types.Int64  `tfsdk:"bridge_fast_path_bytes"`
	BridgeFastPathPackets    types.Int64  `tfsdk:"bridge_fast_path_packets"`
	UseIPFirewall            types.Bool   `tfsdk:"use_ip_firewall"`
	UseIPFirewallForPppoe    types.Bool   `tfsdk:"use_ip_firewall_for_pppoe"`
	UseIPFirewallForVLAN     types.Bool   `tfsdk:"use_ip_firewall_for_vlan"`
	Router                   types.String `tfsdk:"router"`
}

func NewInterfaceBridgeSettingsResource() resource.Resource {
	return &InterfaceBridgeSettingsResource{}
}

func (r *InterfaceBridgeSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_bridge_settings"
}

func (r *InterfaceBridgeSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceBridgeSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/bridge/settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"allow_fast_path": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"bridge_fast_forward_bytes": schema.Int64Attribute{Computed: true,
				Description: "",
			},
			"bridge_fast_forward_packets": schema.Int64Attribute{Computed: true,
				Description: "",
			},
			"bridge_fast_path_active": schema.BoolAttribute{Computed: true,
				Description: "",
			},
			"bridge_fast_path_bytes": schema.Int64Attribute{Computed: true,
				Description: "",
			},
			"bridge_fast_path_packets": schema.Int64Attribute{Computed: true,
				Description: "",
			},
			"use_ip_firewall": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"use_ip_firewall_for_pppoe": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"use_ip_firewall_for_vlan": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *InterfaceBridgeSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceBridgeSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	interfaceBridgeSettingsUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceBridgeSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InterfaceBridgeSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state InterfaceBridgeSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	interfaceBridgeSettingsUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceBridgeSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceBridgeSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/interface/bridge/settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /interface/bridge/settings failed", err.Error())
		return
	}
	interfaceBridgeSettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/interface/bridge/settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceBridgeSettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *InterfaceBridgeSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/interface/bridge/settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/interface/bridge/settings", types.StringValue(routerName))))...)
}

func interfaceBridgeSettingsUpsert(ctx context.Context, reg *client.Registry, plan, state *InterfaceBridgeSettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AllowFastPath.IsNull() || plan.AllowFastPath.IsUnknown()) && (state == nil || !plan.AllowFastPath.Equal(state.AllowFastPath)) {
		body["allow-fast-path"] = client.FormatBool(plan.AllowFastPath.ValueBool())
	}
	if !(plan.UseIPFirewall.IsNull() || plan.UseIPFirewall.IsUnknown()) && (state == nil || !plan.UseIPFirewall.Equal(state.UseIPFirewall)) {
		body["use-ip-firewall"] = client.FormatBool(plan.UseIPFirewall.ValueBool())
	}
	if !(plan.UseIPFirewallForPppoe.IsNull() || plan.UseIPFirewallForPppoe.IsUnknown()) && (state == nil || !plan.UseIPFirewallForPppoe.Equal(state.UseIPFirewallForPppoe)) {
		body["use-ip-firewall-for-pppoe"] = client.FormatBool(plan.UseIPFirewallForPppoe.ValueBool())
	}
	if !(plan.UseIPFirewallForVLAN.IsNull() || plan.UseIPFirewallForVLAN.IsUnknown()) && (state == nil || !plan.UseIPFirewallForVLAN.Equal(state.UseIPFirewallForVLAN)) {
		body["use-ip-firewall-for-vlan"] = client.FormatBool(plan.UseIPFirewallForVLAN.ValueBool())
	}
	obj, err := c.SetSingleton(ctx, "/interface/bridge/settings", body)
	if err != nil {
		diags.AddError("Upsert /interface/bridge/settings failed", err.Error())
		return
	}
	interfaceBridgeSettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/interface/bridge/settings", plan.Router))
}

func interfaceBridgeSettingsApply(ctx context.Context, obj client.Object, m *InterfaceBridgeSettingsModel) {
	_ = ctx
	if v, ok := obj["allow-fast-path"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AllowFastPath = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.AllowFastPath = types.BoolValue(true)
		} else {
			m.AllowFastPath = types.BoolNull()
		}
	}
	if v, ok := obj["bridge-fast-forward-bytes"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.BridgeFastForwardBytes = types.Int64Value(n)
		} else {
			m.BridgeFastForwardBytes = types.Int64Null()
		}
	}
	if v, ok := obj["bridge-fast-forward-packets"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.BridgeFastForwardPackets = types.Int64Value(n)
		} else {
			m.BridgeFastForwardPackets = types.Int64Null()
		}
	}
	if v, ok := obj["bridge-fast-path-active"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.BridgeFastPathActive = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.BridgeFastPathActive = types.BoolValue(true)
		} else {
			m.BridgeFastPathActive = types.BoolNull()
		}
	}
	if v, ok := obj["bridge-fast-path-bytes"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.BridgeFastPathBytes = types.Int64Value(n)
		} else {
			m.BridgeFastPathBytes = types.Int64Null()
		}
	}
	if v, ok := obj["bridge-fast-path-packets"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.BridgeFastPathPackets = types.Int64Value(n)
		} else {
			m.BridgeFastPathPackets = types.Int64Null()
		}
	}
	if v, ok := obj["use-ip-firewall"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.UseIPFirewall = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.UseIPFirewall = types.BoolValue(true)
		} else {
			m.UseIPFirewall = types.BoolNull()
		}
	}
	if v, ok := obj["use-ip-firewall-for-pppoe"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.UseIPFirewallForPppoe = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.UseIPFirewallForPppoe = types.BoolValue(true)
		} else {
			m.UseIPFirewallForPppoe = types.BoolNull()
		}
	}
	if v, ok := obj["use-ip-firewall-for-vlan"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.UseIPFirewallForVLAN = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.UseIPFirewallForVLAN = types.BoolValue(true)
		} else {
			m.UseIPFirewallForVLAN = types.BoolNull()
		}
	}
}
