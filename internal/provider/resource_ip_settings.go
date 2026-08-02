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
	_ resource.Resource                = &IPSettingsResource{}
	_ resource.ResourceWithImportState = &IPSettingsResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type IPSettingsResource struct {
	reg *client.Registry
}

type IPSettingsModel struct {
	ID                                   types.String `tfsdk:"id"`
	Ipv4HighFragmentThresh               types.String `tfsdk:"ipv4_high_fragment_thresh"`
	Ipv4FragmentTime                     types.String `tfsdk:"ipv4_fragment_time"`
	AcceptRedirects                      types.Bool   `tfsdk:"accept_redirects"`
	AcceptSourceRoute                    types.Bool   `tfsdk:"accept_source_route"`
	AllowFastPath                        types.Bool   `tfsdk:"allow_fast_path"`
	ARPTimeout                           types.String `tfsdk:"arp_timeout"`
	IcmpErrorsUseInboundInterfaceAddress types.Bool   `tfsdk:"icmp_errors_use_inbound_interface_address"`
	IcmpRateLimit                        types.Int64  `tfsdk:"icmp_rate_limit"`
	IcmpRateMask                         types.Int64  `tfsdk:"icmp_rate_mask"`
	IPForward                            types.Bool   `tfsdk:"ip_forward"`
	Ipv4FastPathActive                   types.Bool   `tfsdk:"ipv4_fast_path_active"`
	Ipv4FastPathBytes                    types.Int64  `tfsdk:"ipv4_fast_path_bytes"`
	Ipv4FastPathPackets                  types.Int64  `tfsdk:"ipv4_fast_path_packets"`
	Ipv4FasttrackActive                  types.Bool   `tfsdk:"ipv4_fasttrack_active"`
	Ipv4FasttrackBytes                   types.Int64  `tfsdk:"ipv4_fasttrack_bytes"`
	Ipv4FasttrackPackets                 types.Int64  `tfsdk:"ipv4_fasttrack_packets"`
	Ipv4MultipathHashPolicy              types.String `tfsdk:"ipv4_multipath_hash_policy"`
	MaxNeighborEntries                   types.Int64  `tfsdk:"max_neighbor_entries"`
	RpFilter                             types.Bool   `tfsdk:"rp_filter"`
	SecureRedirects                      types.Bool   `tfsdk:"secure_redirects"`
	SendRedirects                        types.Bool   `tfsdk:"send_redirects"`
	TCPSyncookies                        types.Bool   `tfsdk:"tcp_syncookies"`
	TCPTimestamps                        types.String `tfsdk:"tcp_timestamps"`
	Router                               types.String `tfsdk:"router"`
}

func NewIPSettingsResource() resource.Resource { return &IPSettingsResource{} }

func (r *IPSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_settings"
}

func (r *IPSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ipv4_high_fragment_thresh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ipv4-high-fragment-thresh`.",
			},
			"ipv4_fragment_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ipv4-fragment-time`.",
			},
			"accept_redirects": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"accept_source_route": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"allow_fast_path": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"arp_timeout": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"icmp_errors_use_inbound_interface_address": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"icmp_rate_limit": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"icmp_rate_mask": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"ip_forward": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"ipv4_fast_path_active": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"ipv4_fast_path_bytes": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"ipv4_fast_path_packets": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"ipv4_fasttrack_active": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"ipv4_fasttrack_bytes": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"ipv4_fasttrack_packets": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"ipv4_multipath_hash_policy": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"max_neighbor_entries": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"rp_filter": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"secure_redirects": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"send_redirects": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"tcp_syncookies": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"tcp_timestamps": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPSettingsUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state IPSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPSettingsUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/settings failed", err.Error())
		return
	}
	iPSettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPSettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ip/settings", types.StringValue(routerName))))...)
}

func iPSettingsUpsert(ctx context.Context, reg *client.Registry, plan, state *IPSettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AcceptRedirects.IsNull() || plan.AcceptRedirects.IsUnknown()) && (state == nil || !plan.AcceptRedirects.Equal(state.AcceptRedirects)) {
		body["accept-redirects"] = client.FormatBool(plan.AcceptRedirects.ValueBool())
	}
	if !(plan.AcceptSourceRoute.IsNull() || plan.AcceptSourceRoute.IsUnknown()) && (state == nil || !plan.AcceptSourceRoute.Equal(state.AcceptSourceRoute)) {
		body["accept-source-route"] = client.FormatBool(plan.AcceptSourceRoute.ValueBool())
	}
	if !(plan.AllowFastPath.IsNull() || plan.AllowFastPath.IsUnknown()) && (state == nil || !plan.AllowFastPath.Equal(state.AllowFastPath)) {
		body["allow-fast-path"] = client.FormatBool(plan.AllowFastPath.ValueBool())
	}
	if !(plan.ARPTimeout.IsNull() || plan.ARPTimeout.IsUnknown()) && (state == nil || !plan.ARPTimeout.Equal(state.ARPTimeout)) {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !(plan.IcmpErrorsUseInboundInterfaceAddress.IsNull() || plan.IcmpErrorsUseInboundInterfaceAddress.IsUnknown()) && (state == nil || !plan.IcmpErrorsUseInboundInterfaceAddress.Equal(state.IcmpErrorsUseInboundInterfaceAddress)) {
		body["icmp-errors-use-inbound-interface-address"] = client.FormatBool(plan.IcmpErrorsUseInboundInterfaceAddress.ValueBool())
	}
	if !(plan.IcmpRateLimit.IsNull() || plan.IcmpRateLimit.IsUnknown()) && (state == nil || !plan.IcmpRateLimit.Equal(state.IcmpRateLimit)) {
		body["icmp-rate-limit"] = client.FormatInt64(plan.IcmpRateLimit.ValueInt64())
	}
	if !(plan.IcmpRateMask.IsNull() || plan.IcmpRateMask.IsUnknown()) && (state == nil || !plan.IcmpRateMask.Equal(state.IcmpRateMask)) {
		body["icmp-rate-mask"] = client.FormatInt64(plan.IcmpRateMask.ValueInt64())
	}
	if !(plan.IPForward.IsNull() || plan.IPForward.IsUnknown()) && (state == nil || !plan.IPForward.Equal(state.IPForward)) {
		body["ip-forward"] = client.FormatBool(plan.IPForward.ValueBool())
	}
	if !(plan.Ipv4MultipathHashPolicy.IsNull() || plan.Ipv4MultipathHashPolicy.IsUnknown()) && (state == nil || !plan.Ipv4MultipathHashPolicy.Equal(state.Ipv4MultipathHashPolicy)) {
		body["ipv4-multipath-hash-policy"] = plan.Ipv4MultipathHashPolicy.ValueString()
	}
	if !(plan.MaxNeighborEntries.IsNull() || plan.MaxNeighborEntries.IsUnknown()) && (state == nil || !plan.MaxNeighborEntries.Equal(state.MaxNeighborEntries)) {
		body["max-neighbor-entries"] = client.FormatInt64(plan.MaxNeighborEntries.ValueInt64())
	}
	if !(plan.RpFilter.IsNull() || plan.RpFilter.IsUnknown()) && (state == nil || !plan.RpFilter.Equal(state.RpFilter)) {
		body["rp-filter"] = client.FormatBool(plan.RpFilter.ValueBool())
	}
	if !(plan.SecureRedirects.IsNull() || plan.SecureRedirects.IsUnknown()) && (state == nil || !plan.SecureRedirects.Equal(state.SecureRedirects)) {
		body["secure-redirects"] = client.FormatBool(plan.SecureRedirects.ValueBool())
	}
	if !(plan.SendRedirects.IsNull() || plan.SendRedirects.IsUnknown()) && (state == nil || !plan.SendRedirects.Equal(state.SendRedirects)) {
		body["send-redirects"] = client.FormatBool(plan.SendRedirects.ValueBool())
	}
	if !(plan.TCPSyncookies.IsNull() || plan.TCPSyncookies.IsUnknown()) && (state == nil || !plan.TCPSyncookies.Equal(state.TCPSyncookies)) {
		body["tcp-syncookies"] = client.FormatBool(plan.TCPSyncookies.ValueBool())
	}
	if !(plan.TCPTimestamps.IsNull() || plan.TCPTimestamps.IsUnknown()) && (state == nil || !plan.TCPTimestamps.Equal(state.TCPTimestamps)) {
		body["tcp-timestamps"] = plan.TCPTimestamps.ValueString()
	}
	if !(plan.Ipv4FragmentTime.IsNull() || plan.Ipv4FragmentTime.IsUnknown()) && (state == nil || !plan.Ipv4FragmentTime.Equal(state.Ipv4FragmentTime)) {
		body["ipv4-fragment-time"] = plan.Ipv4FragmentTime.ValueString()
	}
	if !(plan.Ipv4HighFragmentThresh.IsNull() || plan.Ipv4HighFragmentThresh.IsUnknown()) && (state == nil || !plan.Ipv4HighFragmentThresh.Equal(state.Ipv4HighFragmentThresh)) {
		body["ipv4-high-fragment-thresh"] = plan.Ipv4HighFragmentThresh.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/ip/settings", body)
	if err != nil {
		diags.AddError("Upsert /ip/settings failed", err.Error())
		return
	}
	iPSettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/settings", plan.Router))
}

func iPSettingsApply(ctx context.Context, obj client.Object, m *IPSettingsModel) {
	_ = ctx
	if v, ok := obj["ipv4-high-fragment-thresh"]; ok && v != "" {
		m.Ipv4HighFragmentThresh = types.StringValue(v)
	} else {
		m.Ipv4HighFragmentThresh = types.StringNull()
	}
	if v, ok := obj["ipv4-fragment-time"]; ok && v != "" {
		m.Ipv4FragmentTime = types.StringValue(v)
	} else {
		m.Ipv4FragmentTime = types.StringNull()
	}
	if v, ok := obj["accept-redirects"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AcceptRedirects = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.AcceptRedirects = types.BoolValue(true)
		} else {
			m.AcceptRedirects = types.BoolNull()
		}
	}
	if v, ok := obj["accept-source-route"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AcceptSourceRoute = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.AcceptSourceRoute = types.BoolValue(true)
		} else {
			m.AcceptSourceRoute = types.BoolNull()
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
	if v, ok := obj["arp-timeout"]; ok {
		_ = v
		if v != "" {
			m.ARPTimeout = types.StringValue(v)
		} else {
			m.ARPTimeout = types.StringNull()
		}
	}
	if v, ok := obj["icmp-errors-use-inbound-interface-address"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.IcmpErrorsUseInboundInterfaceAddress = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.IcmpErrorsUseInboundInterfaceAddress = types.BoolValue(true)
		} else {
			m.IcmpErrorsUseInboundInterfaceAddress = types.BoolNull()
		}
	}
	if v, ok := obj["icmp-rate-limit"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.IcmpRateLimit = types.Int64Value(n)
		} else {
			m.IcmpRateLimit = types.Int64Null()
		}
	}
	if v, ok := obj["icmp-rate-mask"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.IcmpRateMask = types.Int64Value(n)
		} else {
			m.IcmpRateMask = types.Int64Null()
		}
	}
	if v, ok := obj["ip-forward"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.IPForward = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.IPForward = types.BoolValue(true)
		} else {
			m.IPForward = types.BoolNull()
		}
	}
	if v, ok := obj["ipv4-fast-path-active"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Ipv4FastPathActive = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Ipv4FastPathActive = types.BoolValue(true)
		} else {
			m.Ipv4FastPathActive = types.BoolNull()
		}
	}
	if v, ok := obj["ipv4-fast-path-bytes"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Ipv4FastPathBytes = types.Int64Value(n)
		} else {
			m.Ipv4FastPathBytes = types.Int64Null()
		}
	}
	if v, ok := obj["ipv4-fast-path-packets"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Ipv4FastPathPackets = types.Int64Value(n)
		} else {
			m.Ipv4FastPathPackets = types.Int64Null()
		}
	}
	if v, ok := obj["ipv4-fasttrack-active"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Ipv4FasttrackActive = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Ipv4FasttrackActive = types.BoolValue(true)
		} else {
			m.Ipv4FasttrackActive = types.BoolNull()
		}
	}
	if v, ok := obj["ipv4-fasttrack-bytes"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Ipv4FasttrackBytes = types.Int64Value(n)
		} else {
			m.Ipv4FasttrackBytes = types.Int64Null()
		}
	}
	if v, ok := obj["ipv4-fasttrack-packets"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Ipv4FasttrackPackets = types.Int64Value(n)
		} else {
			m.Ipv4FasttrackPackets = types.Int64Null()
		}
	}
	if v, ok := obj["ipv4-multipath-hash-policy"]; ok {
		_ = v
		if v != "" {
			m.Ipv4MultipathHashPolicy = types.StringValue(v)
		} else {
			m.Ipv4MultipathHashPolicy = types.StringNull()
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
	if v, ok := obj["rp-filter"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.RpFilter = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.RpFilter = types.BoolValue(true)
		} else {
			m.RpFilter = types.BoolNull()
		}
	}
	if v, ok := obj["secure-redirects"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.SecureRedirects = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.SecureRedirects = types.BoolValue(true)
		} else {
			m.SecureRedirects = types.BoolNull()
		}
	}
	if v, ok := obj["send-redirects"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.SendRedirects = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.SendRedirects = types.BoolValue(true)
		} else {
			m.SendRedirects = types.BoolNull()
		}
	}
	if v, ok := obj["tcp-syncookies"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.TCPSyncookies = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.TCPSyncookies = types.BoolValue(true)
		} else {
			m.TCPSyncookies = types.BoolNull()
		}
	}
	if v, ok := obj["tcp-timestamps"]; ok {
		_ = v
		if v != "" {
			m.TCPTimestamps = types.StringValue(v)
		} else {
			m.TCPTimestamps = types.StringNull()
		}
	}
}
