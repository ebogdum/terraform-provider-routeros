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
	_ resource.Resource                = &IPV6SettingsResource{}
	_ resource.ResourceWithImportState = &IPV6SettingsResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type IPV6SettingsResource struct {
	reg *client.Registry
}

type IPV6SettingsModel struct {
	ID                           types.String `tfsdk:"id"`
	AcceptRedirects              types.String `tfsdk:"accept_redirects"`
	AcceptRouterAdvertisements   types.String `tfsdk:"accept_router_advertisements"`
	AcceptRouterAdvertisementsOn types.String `tfsdk:"accept_router_advertisements_on"`
	AllowFastPath                types.Bool   `tfsdk:"allow_fast_path"`
	DisableIPV6                  types.Bool   `tfsdk:"disable_ipv6"`
	DisableLinkLocalAddress      types.Bool   `tfsdk:"disable_link_local_address"`
	Forward                      types.Bool   `tfsdk:"forward"`
	IPV6FastPathActive           types.Bool   `tfsdk:"ipv6_fast_path_active"`
	IPV6FastPathBytes            types.Int64  `tfsdk:"ipv6_fast_path_bytes"`
	IPV6FastPathPackets          types.Int64  `tfsdk:"ipv6_fast_path_packets"`
	IPV6FasttrackActive          types.Bool   `tfsdk:"ipv6_fasttrack_active"`
	IPV6FasttrackBytes           types.Int64  `tfsdk:"ipv6_fasttrack_bytes"`
	IPV6FasttrackPackets         types.Int64  `tfsdk:"ipv6_fasttrack_packets"`
	MaxNeighborEntries           types.Int64  `tfsdk:"max_neighbor_entries"`
	MinNeighborEntries           types.Int64  `tfsdk:"min_neighbor_entries"`
	MultipathHashPolicy          types.String `tfsdk:"multipath_hash_policy"`
	SoftMaxNeighborEntries       types.Int64  `tfsdk:"soft_max_neighbor_entries"`
	StaleNeighborDetectInterval  types.Int64  `tfsdk:"stale_neighbor_detect_interval"`
	StaleNeighborTimeout         types.Int64  `tfsdk:"stale_neighbor_timeout"`
	Router                       types.String `tfsdk:"router"`
}

func NewIPV6SettingsResource() resource.Resource { return &IPV6SettingsResource{} }

func (r *IPV6SettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv6_settings"
}

func (r *IPV6SettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPV6SettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ipv6/settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"accept_redirects": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"accept_router_advertisements": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"accept_router_advertisements_on": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"allow_fast_path": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"disable_ipv6": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"disable_link_local_address": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"forward": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"ipv6_fast_path_active": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"ipv6_fast_path_bytes": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"ipv6_fast_path_packets": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"ipv6_fasttrack_active": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"ipv6_fasttrack_bytes": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"ipv6_fasttrack_packets": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"max_neighbor_entries": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"min_neighbor_entries": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"multipath_hash_policy": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"soft_max_neighbor_entries": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"stale_neighbor_detect_interval": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"stale_neighbor_timeout": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *IPV6SettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPV6SettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPV6SettingsUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6SettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPV6SettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state IPV6SettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPV6SettingsUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6SettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPV6SettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ipv6/settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /ipv6/settings failed", err.Error())
		return
	}
	iPV6SettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ipv6/settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPV6SettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPV6SettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ipv6/settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ipv6/settings", types.StringValue(routerName))))...)
}

func iPV6SettingsUpsert(ctx context.Context, reg *client.Registry, plan, state *IPV6SettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AcceptRedirects.IsNull() || plan.AcceptRedirects.IsUnknown()) && (state == nil || !plan.AcceptRedirects.Equal(state.AcceptRedirects)) {
		body["accept-redirects"] = plan.AcceptRedirects.ValueString()
	}
	if !(plan.AcceptRouterAdvertisements.IsNull() || plan.AcceptRouterAdvertisements.IsUnknown()) && (state == nil || !plan.AcceptRouterAdvertisements.Equal(state.AcceptRouterAdvertisements)) {
		body["accept-router-advertisements"] = plan.AcceptRouterAdvertisements.ValueString()
	}
	if !(plan.AcceptRouterAdvertisementsOn.IsNull() || plan.AcceptRouterAdvertisementsOn.IsUnknown()) && (state == nil || !plan.AcceptRouterAdvertisementsOn.Equal(state.AcceptRouterAdvertisementsOn)) {
		body["accept-router-advertisements-on"] = plan.AcceptRouterAdvertisementsOn.ValueString()
	}
	if !(plan.AllowFastPath.IsNull() || plan.AllowFastPath.IsUnknown()) && (state == nil || !plan.AllowFastPath.Equal(state.AllowFastPath)) {
		body["allow-fast-path"] = client.FormatBool(plan.AllowFastPath.ValueBool())
	}
	if !(plan.DisableIPV6.IsNull() || plan.DisableIPV6.IsUnknown()) && (state == nil || !plan.DisableIPV6.Equal(state.DisableIPV6)) {
		body["disable-ipv6"] = client.FormatBool(plan.DisableIPV6.ValueBool())
	}
	if !(plan.DisableLinkLocalAddress.IsNull() || plan.DisableLinkLocalAddress.IsUnknown()) && (state == nil || !plan.DisableLinkLocalAddress.Equal(state.DisableLinkLocalAddress)) {
		body["disable-link-local-address"] = client.FormatBool(plan.DisableLinkLocalAddress.ValueBool())
	}
	if !(plan.Forward.IsNull() || plan.Forward.IsUnknown()) && (state == nil || !plan.Forward.Equal(state.Forward)) {
		body["forward"] = client.FormatBool(plan.Forward.ValueBool())
	}
	if !(plan.MaxNeighborEntries.IsNull() || plan.MaxNeighborEntries.IsUnknown()) && (state == nil || !plan.MaxNeighborEntries.Equal(state.MaxNeighborEntries)) {
		body["max-neighbor-entries"] = client.FormatInt64(plan.MaxNeighborEntries.ValueInt64())
	}
	if !(plan.MinNeighborEntries.IsNull() || plan.MinNeighborEntries.IsUnknown()) && (state == nil || !plan.MinNeighborEntries.Equal(state.MinNeighborEntries)) {
		body["min-neighbor-entries"] = client.FormatInt64(plan.MinNeighborEntries.ValueInt64())
	}
	if !(plan.MultipathHashPolicy.IsNull() || plan.MultipathHashPolicy.IsUnknown()) && (state == nil || !plan.MultipathHashPolicy.Equal(state.MultipathHashPolicy)) {
		body["multipath-hash-policy"] = plan.MultipathHashPolicy.ValueString()
	}
	if !(plan.SoftMaxNeighborEntries.IsNull() || plan.SoftMaxNeighborEntries.IsUnknown()) && (state == nil || !plan.SoftMaxNeighborEntries.Equal(state.SoftMaxNeighborEntries)) {
		body["soft-max-neighbor-entries"] = client.FormatInt64(plan.SoftMaxNeighborEntries.ValueInt64())
	}
	if !(plan.StaleNeighborDetectInterval.IsNull() || plan.StaleNeighborDetectInterval.IsUnknown()) && (state == nil || !plan.StaleNeighborDetectInterval.Equal(state.StaleNeighborDetectInterval)) {
		body["stale-neighbor-detect-interval"] = client.FormatInt64(plan.StaleNeighborDetectInterval.ValueInt64())
	}
	if !(plan.StaleNeighborTimeout.IsNull() || plan.StaleNeighborTimeout.IsUnknown()) && (state == nil || !plan.StaleNeighborTimeout.Equal(state.StaleNeighborTimeout)) {
		body["stale-neighbor-timeout"] = client.FormatInt64(plan.StaleNeighborTimeout.ValueInt64())
	}
	obj, err := c.SetSingleton(ctx, "/ipv6/settings", body)
	if err != nil {
		diags.AddError("Upsert /ipv6/settings failed", err.Error())
		return
	}
	iPV6SettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ipv6/settings", plan.Router))
}

func iPV6SettingsApply(ctx context.Context, obj client.Object, m *IPV6SettingsModel) {
	_ = ctx
	if v, ok := obj["accept-redirects"]; ok {
		_ = v
		if v != "" {
			m.AcceptRedirects = types.StringValue(v)
		} else {
			m.AcceptRedirects = types.StringNull()
		}
	}
	if v, ok := obj["accept-router-advertisements"]; ok {
		_ = v
		if v != "" {
			m.AcceptRouterAdvertisements = types.StringValue(v)
		} else {
			m.AcceptRouterAdvertisements = types.StringNull()
		}
	}
	if v, ok := obj["accept-router-advertisements-on"]; ok {
		_ = v
		if v != "" {
			m.AcceptRouterAdvertisementsOn = types.StringValue(v)
		} else {
			m.AcceptRouterAdvertisementsOn = types.StringNull()
		}
	}
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
	if v, ok := obj["disable-ipv6"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.DisableIPV6 = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.DisableIPV6 = types.BoolValue(true)
		} else {
			m.DisableIPV6 = types.BoolNull()
		}
	}
	if v, ok := obj["disable-link-local-address"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.DisableLinkLocalAddress = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.DisableLinkLocalAddress = types.BoolValue(true)
		} else {
			m.DisableLinkLocalAddress = types.BoolNull()
		}
	}
	if v, ok := obj["forward"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Forward = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Forward = types.BoolValue(true)
		} else {
			m.Forward = types.BoolNull()
		}
	}
	if v, ok := obj["ipv6-fast-path-active"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.IPV6FastPathActive = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.IPV6FastPathActive = types.BoolValue(true)
		} else {
			m.IPV6FastPathActive = types.BoolNull()
		}
	}
	if v, ok := obj["ipv6-fast-path-bytes"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.IPV6FastPathBytes = types.Int64Value(n)
		} else {
			m.IPV6FastPathBytes = types.Int64Null()
		}
	}
	if v, ok := obj["ipv6-fast-path-packets"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.IPV6FastPathPackets = types.Int64Value(n)
		} else {
			m.IPV6FastPathPackets = types.Int64Null()
		}
	}
	if v, ok := obj["ipv6-fasttrack-active"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.IPV6FasttrackActive = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.IPV6FasttrackActive = types.BoolValue(true)
		} else {
			m.IPV6FasttrackActive = types.BoolNull()
		}
	}
	if v, ok := obj["ipv6-fasttrack-bytes"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.IPV6FasttrackBytes = types.Int64Value(n)
		} else {
			m.IPV6FasttrackBytes = types.Int64Null()
		}
	}
	if v, ok := obj["ipv6-fasttrack-packets"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.IPV6FasttrackPackets = types.Int64Value(n)
		} else {
			m.IPV6FasttrackPackets = types.Int64Null()
		}
	}
	if v, ok := obj["max-neighbor-entries"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MaxNeighborEntries = types.Int64Value(n)
		} else {
			m.MaxNeighborEntries = types.Int64Null()
		}
	}
	if v, ok := obj["min-neighbor-entries"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MinNeighborEntries = types.Int64Value(n)
		} else {
			m.MinNeighborEntries = types.Int64Null()
		}
	}
	if v, ok := obj["multipath-hash-policy"]; ok {
		_ = v
		if v != "" {
			m.MultipathHashPolicy = types.StringValue(v)
		} else {
			m.MultipathHashPolicy = types.StringNull()
		}
	}
	if v, ok := obj["soft-max-neighbor-entries"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.SoftMaxNeighborEntries = types.Int64Value(n)
		} else {
			m.SoftMaxNeighborEntries = types.Int64Null()
		}
	}
	if v, ok := obj["stale-neighbor-detect-interval"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.StaleNeighborDetectInterval = types.Int64Value(n)
		} else {
			m.StaleNeighborDetectInterval = types.Int64Null()
		}
	}
	if v, ok := obj["stale-neighbor-timeout"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.StaleNeighborTimeout = types.Int64Value(n)
		} else {
			m.StaleNeighborTimeout = types.Int64Null()
		}
	}
}
