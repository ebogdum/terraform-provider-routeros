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
	_ resource.Resource                = &IPNeighborDiscoverySettingsResource{}
	_ resource.ResourceWithImportState = &IPNeighborDiscoverySettingsResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPNeighborDiscoverySettingsResource struct {
	reg *client.Registry
}

type IPNeighborDiscoverySettingsModel struct {
	ID                    types.String `tfsdk:"id"`
	LldpMed               types.String `tfsdk:"lldp_med"`
	AddDnsEntriesSuffix   types.String `tfsdk:"add_dns_entries_suffix"`
	AddDnsEntries         types.String `tfsdk:"add_dns_entries"`
	DiscoverInterfaceList types.String `tfsdk:"discover_interface_list"`
	DiscoverInterval      types.String `tfsdk:"discover_interval"`
	LldpMACPhyConfig      types.Bool   `tfsdk:"lldp_mac_phy_config"`
	LldpMaxFrameSize      types.Bool   `tfsdk:"lldp_max_frame_size"`
	LldpMedNetPolicyVlan  types.String `tfsdk:"lldp_med_net_policy_vlan"`
	LldpPoePower          types.Bool   `tfsdk:"lldp_poe_power"`
	LldpVlanInfo          types.Bool   `tfsdk:"lldp_vlan_info"`
	Mode                  types.String `tfsdk:"mode"`
	Protocol              types.String `tfsdk:"protocol"`
	Router                types.String `tfsdk:"router"`
}

func NewIPNeighborDiscoverySettingsResource() resource.Resource {
	return &IPNeighborDiscoverySettingsResource{}
}

func (r *IPNeighborDiscoverySettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_neighbor_discovery_settings"
}

func (r *IPNeighborDiscoverySettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPNeighborDiscoverySettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/neighbor/discovery-settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"lldp_med": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `lldp-med`.",
			},
			"add_dns_entries_suffix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `add-dns-entries-suffix`.",
			},
			"add_dns_entries": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `add-dns-entries`.",
			},
			"discover_interface_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `discover-interface-list`.",
			},
			"discover_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `discover-interval`.",
			},
			"lldp_mac_phy_config": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `lldp-mac-phy-config`.",
			},
			"lldp_max_frame_size": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `lldp-max-frame-size`.",
			},
			"lldp_med_net_policy_vlan": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `lldp-med-net-policy-vlan`.",
			},
			"lldp_poe_power": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `lldp-poe-power`.",
			},
			"lldp_vlan_info": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `lldp-vlan-info`.",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mode`.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `protocol`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPNeighborDiscoverySettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPNeighborDiscoverySettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPNeighborDiscoverySettingsUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPNeighborDiscoverySettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPNeighborDiscoverySettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPNeighborDiscoverySettingsUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPNeighborDiscoverySettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPNeighborDiscoverySettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/neighbor/discovery-settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/neighbor/discovery-settings failed", err.Error())
		return
	}
	iPNeighborDiscoverySettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/neighbor/discovery-settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPNeighborDiscoverySettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPNeighborDiscoverySettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/neighbor/discovery-settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ip/neighbor/discovery-settings", types.StringValue(routerName))))...)
}

func iPNeighborDiscoverySettingsUpsert(ctx context.Context, reg *client.Registry, plan *IPNeighborDiscoverySettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.DiscoverInterfaceList.IsNull() || plan.DiscoverInterfaceList.IsUnknown()) {
		body["discover-interface-list"] = plan.DiscoverInterfaceList.ValueString()
	}
	if !(plan.DiscoverInterval.IsNull() || plan.DiscoverInterval.IsUnknown()) {
		body["discover-interval"] = plan.DiscoverInterval.ValueString()
	}
	if !(plan.LldpMACPhyConfig.IsNull() || plan.LldpMACPhyConfig.IsUnknown()) {
		body["lldp-mac-phy-config"] = client.FormatBool(plan.LldpMACPhyConfig.ValueBool())
	}
	if !(plan.LldpMaxFrameSize.IsNull() || plan.LldpMaxFrameSize.IsUnknown()) {
		body["lldp-max-frame-size"] = client.FormatBool(plan.LldpMaxFrameSize.ValueBool())
	}
	if !(plan.LldpMedNetPolicyVlan.IsNull() || plan.LldpMedNetPolicyVlan.IsUnknown()) {
		body["lldp-med-net-policy-vlan"] = plan.LldpMedNetPolicyVlan.ValueString()
	}
	if !(plan.LldpPoePower.IsNull() || plan.LldpPoePower.IsUnknown()) {
		body["lldp-poe-power"] = client.FormatBool(plan.LldpPoePower.ValueBool())
	}
	if !(plan.LldpVlanInfo.IsNull() || plan.LldpVlanInfo.IsUnknown()) {
		body["lldp-vlan-info"] = client.FormatBool(plan.LldpVlanInfo.ValueBool())
	}
	if !(plan.Mode.IsNull() || plan.Mode.IsUnknown()) {
		body["mode"] = plan.Mode.ValueString()
	}
	if !(plan.Protocol.IsNull() || plan.Protocol.IsUnknown()) {
		body["protocol"] = plan.Protocol.ValueString()
	}
	if !(plan.AddDnsEntries.IsNull() || plan.AddDnsEntries.IsUnknown()) {
		body["add-dns-entries"] = plan.AddDnsEntries.ValueString()
	}
	if !(plan.AddDnsEntriesSuffix.IsNull() || plan.AddDnsEntriesSuffix.IsUnknown()) {
		body["add-dns-entries-suffix"] = plan.AddDnsEntriesSuffix.ValueString()
	}
	if !(plan.LldpMed.IsNull() || plan.LldpMed.IsUnknown()) {
		body["lldp-med"] = plan.LldpMed.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/ip/neighbor/discovery-settings", body)
	if err != nil {
		diags.AddError("Upsert /ip/neighbor/discovery-settings failed", err.Error())
		return
	}
	iPNeighborDiscoverySettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/neighbor/discovery-settings", plan.Router))
}

func iPNeighborDiscoverySettingsApply(ctx context.Context, obj client.Object, m *IPNeighborDiscoverySettingsModel) {
	_ = ctx
	if v, ok := obj["lldp-med"]; ok && v != "" {
		m.LldpMed = types.StringValue(v)
	} else {
		m.LldpMed = types.StringNull()
	}
	if v, ok := obj["add-dns-entries-suffix"]; ok && v != "" {
		m.AddDnsEntriesSuffix = types.StringValue(v)
	} else {
		m.AddDnsEntriesSuffix = types.StringNull()
	}
	if v, ok := obj["add-dns-entries"]; ok && v != "" {
		m.AddDnsEntries = types.StringValue(v)
	} else {
		m.AddDnsEntries = types.StringNull()
	}
	if v, ok := obj["discover-interface-list"]; ok && v != "" {
		m.DiscoverInterfaceList = types.StringValue(v)
	} else {
		m.DiscoverInterfaceList = types.StringNull()
	}
	if v, ok := obj["discover-interval"]; ok && v != "" {
		m.DiscoverInterval = types.StringValue(v)
	} else {
		m.DiscoverInterval = types.StringNull()
	}
	if v, ok := obj["lldp-mac-phy-config"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.LldpMACPhyConfig = types.BoolValue(b)
		} else {
			m.LldpMACPhyConfig = types.BoolNull()
		}
	} else {
		m.LldpMACPhyConfig = types.BoolNull()
	}
	if v, ok := obj["lldp-max-frame-size"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.LldpMaxFrameSize = types.BoolValue(b)
		} else {
			m.LldpMaxFrameSize = types.BoolNull()
		}
	} else {
		m.LldpMaxFrameSize = types.BoolNull()
	}
	if v, ok := obj["lldp-med-net-policy-vlan"]; ok && v != "" {
		m.LldpMedNetPolicyVlan = types.StringValue(v)
	} else {
		m.LldpMedNetPolicyVlan = types.StringNull()
	}
	if v, ok := obj["lldp-poe-power"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.LldpPoePower = types.BoolValue(b)
		} else {
			m.LldpPoePower = types.BoolNull()
		}
	} else {
		m.LldpPoePower = types.BoolNull()
	}
	if v, ok := obj["lldp-vlan-info"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.LldpVlanInfo = types.BoolValue(b)
		} else {
			m.LldpVlanInfo = types.BoolNull()
		}
	} else {
		m.LldpVlanInfo = types.BoolNull()
	}
	if v, ok := obj["mode"]; ok && v != "" {
		m.Mode = types.StringValue(v)
	} else {
		m.Mode = types.StringNull()
	}
	if v, ok := obj["protocol"]; ok && v != "" {
		m.Protocol = types.StringValue(v)
	} else {
		m.Protocol = types.StringNull()
	}
}
