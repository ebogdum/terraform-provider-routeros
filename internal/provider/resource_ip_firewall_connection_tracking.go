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
	_ resource.Resource                = &IPFirewallConnectionTrackingResource{}
	_ resource.ResourceWithImportState = &IPFirewallConnectionTrackingResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type IPFirewallConnectionTrackingResource struct {
	reg *client.Registry
}

type IPFirewallConnectionTrackingModel struct {
	ID                    types.String  `tfsdk:"id"`
	ActiveIpv4            types.Bool    `tfsdk:"active_ipv4"`
	ActiveIPV6            types.Bool    `tfsdk:"active_ipv6"`
	Enabled               types.String  `tfsdk:"enabled"`
	GenericTimeout        durationValue `tfsdk:"generic_timeout"`
	IcmpTimeout           durationValue `tfsdk:"icmp_timeout"`
	LiberalTCPTracking    types.Bool    `tfsdk:"liberal_tcp_tracking"`
	LooseTCPTracking      types.Bool    `tfsdk:"loose_tcp_tracking"`
	MaxEntries            types.Int64   `tfsdk:"max_entries"`
	TCPCloseTimeout       durationValue `tfsdk:"tcp_close_timeout"`
	TCPCloseWaitTimeout   durationValue `tfsdk:"tcp_close_wait_timeout"`
	TCPEstablishedTimeout durationValue `tfsdk:"tcp_established_timeout"`
	TCPFinWaitTimeout     durationValue `tfsdk:"tcp_fin_wait_timeout"`
	TCPLastAckTimeout     durationValue `tfsdk:"tcp_last_ack_timeout"`
	TCPMaxRetransTimeout  durationValue `tfsdk:"tcp_max_retrans_timeout"`
	TCPSynReceivedTimeout durationValue `tfsdk:"tcp_syn_received_timeout"`
	TCPSynSentTimeout     durationValue `tfsdk:"tcp_syn_sent_timeout"`
	TCPTimeWaitTimeout    durationValue `tfsdk:"tcp_time_wait_timeout"`
	TCPUnackedTimeout     durationValue `tfsdk:"tcp_unacked_timeout"`
	TotalEntries          types.Int64   `tfsdk:"total_entries"`
	TotalIp4Entries       types.Int64   `tfsdk:"total_ip4_entries"`
	TotalIp6Entries       types.Int64   `tfsdk:"total_ip6_entries"`
	UDPStreamTimeout      durationValue `tfsdk:"udp_stream_timeout"`
	UDPTimeout            durationValue `tfsdk:"udp_timeout"`
	Router                types.String  `tfsdk:"router"`
}

func NewIPFirewallConnectionTrackingResource() resource.Resource {
	return &IPFirewallConnectionTrackingResource{}
}

func (r *IPFirewallConnectionTrackingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_firewall_connection_tracking"
}

func (r *IPFirewallConnectionTrackingResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPFirewallConnectionTrackingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/firewall/connection/tracking`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"active_ipv4": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"active_ipv6": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"enabled": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"generic_timeout": schema.StringAttribute{
				CustomType: durationType{}, Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"icmp_timeout": schema.StringAttribute{
				CustomType: durationType{}, Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"liberal_tcp_tracking": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"loose_tcp_tracking": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"max_entries": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"tcp_close_timeout": schema.StringAttribute{
				CustomType: durationType{}, Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"tcp_close_wait_timeout": schema.StringAttribute{
				CustomType: durationType{}, Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"tcp_established_timeout": schema.StringAttribute{
				CustomType: durationType{}, Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"tcp_fin_wait_timeout": schema.StringAttribute{
				CustomType: durationType{}, Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"tcp_last_ack_timeout": schema.StringAttribute{
				CustomType: durationType{}, Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"tcp_max_retrans_timeout": schema.StringAttribute{
				CustomType: durationType{}, Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"tcp_syn_received_timeout": schema.StringAttribute{
				CustomType: durationType{}, Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"tcp_syn_sent_timeout": schema.StringAttribute{
				CustomType: durationType{}, Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"tcp_time_wait_timeout": schema.StringAttribute{
				CustomType: durationType{}, Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"tcp_unacked_timeout": schema.StringAttribute{
				CustomType: durationType{}, Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"total_entries": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"total_ip4_entries": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"total_ip6_entries": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"udp_stream_timeout": schema.StringAttribute{
				CustomType: durationType{}, Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"udp_timeout": schema.StringAttribute{
				CustomType: durationType{}, Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"router": schema.StringAttribute{Optional: true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *IPFirewallConnectionTrackingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPFirewallConnectionTrackingModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPFirewallConnectionTrackingUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPFirewallConnectionTrackingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPFirewallConnectionTrackingModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state IPFirewallConnectionTrackingModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPFirewallConnectionTrackingUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPFirewallConnectionTrackingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPFirewallConnectionTrackingModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/firewall/connection/tracking")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/firewall/connection/tracking failed", err.Error())
		return
	}
	iPFirewallConnectionTrackingApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/firewall/connection/tracking", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPFirewallConnectionTrackingResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPFirewallConnectionTrackingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/firewall/connection/tracking" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ip/firewall/connection/tracking", types.StringValue(routerName))))...)
}

func iPFirewallConnectionTrackingUpsert(ctx context.Context, reg *client.Registry, plan, state *IPFirewallConnectionTrackingModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = plan.Enabled.ValueString()
	}
	if !(plan.GenericTimeout.IsNull() || plan.GenericTimeout.IsUnknown()) && (state == nil || !plan.GenericTimeout.Equal(state.GenericTimeout)) {
		body["generic-timeout"] = plan.GenericTimeout.ValueString()
	}
	if !(plan.IcmpTimeout.IsNull() || plan.IcmpTimeout.IsUnknown()) && (state == nil || !plan.IcmpTimeout.Equal(state.IcmpTimeout)) {
		body["icmp-timeout"] = plan.IcmpTimeout.ValueString()
	}
	if !(plan.LiberalTCPTracking.IsNull() || plan.LiberalTCPTracking.IsUnknown()) && (state == nil || !plan.LiberalTCPTracking.Equal(state.LiberalTCPTracking)) {
		body["liberal-tcp-tracking"] = client.FormatBool(plan.LiberalTCPTracking.ValueBool())
	}
	if !(plan.LooseTCPTracking.IsNull() || plan.LooseTCPTracking.IsUnknown()) && (state == nil || !plan.LooseTCPTracking.Equal(state.LooseTCPTracking)) {
		body["loose-tcp-tracking"] = client.FormatBool(plan.LooseTCPTracking.ValueBool())
	}
	if !(plan.TCPCloseTimeout.IsNull() || plan.TCPCloseTimeout.IsUnknown()) && (state == nil || !plan.TCPCloseTimeout.Equal(state.TCPCloseTimeout)) {
		body["tcp-close-timeout"] = plan.TCPCloseTimeout.ValueString()
	}
	if !(plan.TCPCloseWaitTimeout.IsNull() || plan.TCPCloseWaitTimeout.IsUnknown()) && (state == nil || !plan.TCPCloseWaitTimeout.Equal(state.TCPCloseWaitTimeout)) {
		body["tcp-close-wait-timeout"] = plan.TCPCloseWaitTimeout.ValueString()
	}
	if !(plan.TCPEstablishedTimeout.IsNull() || plan.TCPEstablishedTimeout.IsUnknown()) && (state == nil || !plan.TCPEstablishedTimeout.Equal(state.TCPEstablishedTimeout)) {
		body["tcp-established-timeout"] = plan.TCPEstablishedTimeout.ValueString()
	}
	if !(plan.TCPFinWaitTimeout.IsNull() || plan.TCPFinWaitTimeout.IsUnknown()) && (state == nil || !plan.TCPFinWaitTimeout.Equal(state.TCPFinWaitTimeout)) {
		body["tcp-fin-wait-timeout"] = plan.TCPFinWaitTimeout.ValueString()
	}
	if !(plan.TCPLastAckTimeout.IsNull() || plan.TCPLastAckTimeout.IsUnknown()) && (state == nil || !plan.TCPLastAckTimeout.Equal(state.TCPLastAckTimeout)) {
		body["tcp-last-ack-timeout"] = plan.TCPLastAckTimeout.ValueString()
	}
	if !(plan.TCPMaxRetransTimeout.IsNull() || plan.TCPMaxRetransTimeout.IsUnknown()) && (state == nil || !plan.TCPMaxRetransTimeout.Equal(state.TCPMaxRetransTimeout)) {
		body["tcp-max-retrans-timeout"] = plan.TCPMaxRetransTimeout.ValueString()
	}
	if !(plan.TCPSynReceivedTimeout.IsNull() || plan.TCPSynReceivedTimeout.IsUnknown()) && (state == nil || !plan.TCPSynReceivedTimeout.Equal(state.TCPSynReceivedTimeout)) {
		body["tcp-syn-received-timeout"] = plan.TCPSynReceivedTimeout.ValueString()
	}
	if !(plan.TCPSynSentTimeout.IsNull() || plan.TCPSynSentTimeout.IsUnknown()) && (state == nil || !plan.TCPSynSentTimeout.Equal(state.TCPSynSentTimeout)) {
		body["tcp-syn-sent-timeout"] = plan.TCPSynSentTimeout.ValueString()
	}
	if !(plan.TCPTimeWaitTimeout.IsNull() || plan.TCPTimeWaitTimeout.IsUnknown()) && (state == nil || !plan.TCPTimeWaitTimeout.Equal(state.TCPTimeWaitTimeout)) {
		body["tcp-time-wait-timeout"] = plan.TCPTimeWaitTimeout.ValueString()
	}
	if !(plan.TCPUnackedTimeout.IsNull() || plan.TCPUnackedTimeout.IsUnknown()) && (state == nil || !plan.TCPUnackedTimeout.Equal(state.TCPUnackedTimeout)) {
		body["tcp-unacked-timeout"] = plan.TCPUnackedTimeout.ValueString()
	}
	if !(plan.UDPStreamTimeout.IsNull() || plan.UDPStreamTimeout.IsUnknown()) && (state == nil || !plan.UDPStreamTimeout.Equal(state.UDPStreamTimeout)) {
		body["udp-stream-timeout"] = plan.UDPStreamTimeout.ValueString()
	}
	if !(plan.UDPTimeout.IsNull() || plan.UDPTimeout.IsUnknown()) && (state == nil || !plan.UDPTimeout.Equal(state.UDPTimeout)) {
		body["udp-timeout"] = plan.UDPTimeout.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/ip/firewall/connection/tracking", body)
	if err != nil {
		diags.AddError("Upsert /ip/firewall/connection/tracking failed", err.Error())
		return
	}
	iPFirewallConnectionTrackingApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/firewall/connection/tracking", plan.Router))
}

func iPFirewallConnectionTrackingApply(ctx context.Context, obj client.Object, m *IPFirewallConnectionTrackingModel) {
	_ = ctx
	if v, ok := obj["active-ipv4"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.ActiveIpv4 = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.ActiveIpv4 = types.BoolValue(true)
		} else {
			m.ActiveIpv4 = types.BoolNull()
		}
	}
	if v, ok := obj["active-ipv6"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.ActiveIPV6 = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.ActiveIPV6 = types.BoolValue(true)
		} else {
			m.ActiveIPV6 = types.BoolNull()
		}
	}
	if v, ok := obj["enabled"]; ok {
		_ = v
		if v != "" {
			m.Enabled = types.StringValue(v)
		} else {
			m.Enabled = types.StringNull()
		}
	}
	if v, ok := obj["generic-timeout"]; ok {
		_ = v
		if v != "" {
			m.GenericTimeout = newDurationValue(v)
		} else {
			m.GenericTimeout = newDurationNull()
		}
	}
	if v, ok := obj["icmp-timeout"]; ok {
		_ = v
		if v != "" {
			m.IcmpTimeout = newDurationValue(v)
		} else {
			m.IcmpTimeout = newDurationNull()
		}
	}
	if v, ok := obj["liberal-tcp-tracking"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.LiberalTCPTracking = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.LiberalTCPTracking = types.BoolValue(true)
		} else {
			m.LiberalTCPTracking = types.BoolNull()
		}
	}
	if v, ok := obj["loose-tcp-tracking"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.LooseTCPTracking = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.LooseTCPTracking = types.BoolValue(true)
		} else {
			m.LooseTCPTracking = types.BoolNull()
		}
	}
	if v, ok := obj["max-entries"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MaxEntries = types.Int64Value(n)
		} else {
			m.MaxEntries = types.Int64Null()
		}
	}
	if v, ok := obj["tcp-close-timeout"]; ok {
		_ = v
		if v != "" {
			m.TCPCloseTimeout = newDurationValue(v)
		} else {
			m.TCPCloseTimeout = newDurationNull()
		}
	}
	if v, ok := obj["tcp-close-wait-timeout"]; ok {
		_ = v
		if v != "" {
			m.TCPCloseWaitTimeout = newDurationValue(v)
		} else {
			m.TCPCloseWaitTimeout = newDurationNull()
		}
	}
	if v, ok := obj["tcp-established-timeout"]; ok {
		_ = v
		if v != "" {
			m.TCPEstablishedTimeout = newDurationValue(v)
		} else {
			m.TCPEstablishedTimeout = newDurationNull()
		}
	}
	if v, ok := obj["tcp-fin-wait-timeout"]; ok {
		_ = v
		if v != "" {
			m.TCPFinWaitTimeout = newDurationValue(v)
		} else {
			m.TCPFinWaitTimeout = newDurationNull()
		}
	}
	if v, ok := obj["tcp-last-ack-timeout"]; ok {
		_ = v
		if v != "" {
			m.TCPLastAckTimeout = newDurationValue(v)
		} else {
			m.TCPLastAckTimeout = newDurationNull()
		}
	}
	if v, ok := obj["tcp-max-retrans-timeout"]; ok {
		_ = v
		if v != "" {
			m.TCPMaxRetransTimeout = newDurationValue(v)
		} else {
			m.TCPMaxRetransTimeout = newDurationNull()
		}
	}
	if v, ok := obj["tcp-syn-received-timeout"]; ok {
		_ = v
		if v != "" {
			m.TCPSynReceivedTimeout = newDurationValue(v)
		} else {
			m.TCPSynReceivedTimeout = newDurationNull()
		}
	}
	if v, ok := obj["tcp-syn-sent-timeout"]; ok {
		_ = v
		if v != "" {
			m.TCPSynSentTimeout = newDurationValue(v)
		} else {
			m.TCPSynSentTimeout = newDurationNull()
		}
	}
	if v, ok := obj["tcp-time-wait-timeout"]; ok {
		_ = v
		if v != "" {
			m.TCPTimeWaitTimeout = newDurationValue(v)
		} else {
			m.TCPTimeWaitTimeout = newDurationNull()
		}
	}
	if v, ok := obj["tcp-unacked-timeout"]; ok {
		_ = v
		if v != "" {
			m.TCPUnackedTimeout = newDurationValue(v)
		} else {
			m.TCPUnackedTimeout = newDurationNull()
		}
	}
	if v, ok := obj["total-entries"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TotalEntries = types.Int64Value(n)
		} else {
			m.TotalEntries = types.Int64Null()
		}
	}
	if v, ok := obj["total-ip4-entries"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TotalIp4Entries = types.Int64Value(n)
		} else {
			m.TotalIp4Entries = types.Int64Null()
		}
	}
	if v, ok := obj["total-ip6-entries"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TotalIp6Entries = types.Int64Value(n)
		} else {
			m.TotalIp6Entries = types.Int64Null()
		}
	}
	if v, ok := obj["udp-stream-timeout"]; ok {
		_ = v
		if v != "" {
			m.UDPStreamTimeout = newDurationValue(v)
		} else {
			m.UDPStreamTimeout = newDurationNull()
		}
	}
	if v, ok := obj["udp-timeout"]; ok {
		_ = v
		if v != "" {
			m.UDPTimeout = newDurationValue(v)
		} else {
			m.UDPTimeout = newDurationNull()
		}
	}
}
