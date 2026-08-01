package provider

import (
	"context"
	"fmt"

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
	_ resource.Resource                = &IPTrafficFlowResource{}
	_ resource.ResourceWithImportState = &IPTrafficFlowResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type IPTrafficFlowResource struct {
	reg *client.Registry
}

type IPTrafficFlowModel struct {
	ID                  types.String `tfsdk:"id"`
	ActiveFlowTimeout   types.String `tfsdk:"active_flow_timeout"`
	CacheEntries        types.String `tfsdk:"cache_entries"`
	Enabled             types.Bool   `tfsdk:"enabled"`
	InactiveFlowTimeout types.String `tfsdk:"inactive_flow_timeout"`
	Interfaces          types.String `tfsdk:"interfaces"`
	PacketSampling      types.Bool   `tfsdk:"packet_sampling"`
	SamplingInterval    types.Int64  `tfsdk:"sampling_interval"`
	SamplingSpace       types.Int64  `tfsdk:"sampling_space"`
	Router              types.String `tfsdk:"router"`
}

func NewIPTrafficFlowResource() resource.Resource { return &IPTrafficFlowResource{} }

func (r *IPTrafficFlowResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_traffic_flow"
}

func (r *IPTrafficFlowResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPTrafficFlowResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/traffic-flow`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"active_flow_timeout": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "Maximum life-time of a flow.",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"cache_entries": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Number of flows which can be in router's memory simultaneously.",
			},
			"enabled": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"inactive_flow_timeout": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "How long to keep the flow active, if it is idle. If a connection does not see any packet within this timeout, then traffic-flow will send a packet out as a new flow. If this timeout is too small it can create a significant amount of flows and overflow the buffer.",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"interfaces": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Names of those interfaces will be used to gather statistics for traffic-flow. To specify more than one interface, separate them with a comma.",
			},
			"packet_sampling": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "Enable or disable packet sampling feature.",
			},
			"sampling_interval": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "The number of packets that are consecutively sampled.",
			},
			"sampling_space": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "The number of packets that are consecutively omitted.",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPTrafficFlowResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPTrafficFlowModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPTrafficFlowUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPTrafficFlowResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPTrafficFlowModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state IPTrafficFlowModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPTrafficFlowUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPTrafficFlowResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPTrafficFlowModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/traffic-flow")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/traffic-flow failed", err.Error())
		return
	}
	iPTrafficFlowApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/traffic-flow", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPTrafficFlowResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPTrafficFlowResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/traffic-flow" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ip/traffic-flow", types.StringValue(routerName))))...)
}

func iPTrafficFlowUpsert(ctx context.Context, reg *client.Registry, plan, state *IPTrafficFlowModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.ActiveFlowTimeout.IsNull() || plan.ActiveFlowTimeout.IsUnknown()) && (state == nil || !plan.ActiveFlowTimeout.Equal(state.ActiveFlowTimeout)) {
		body["active-flow-timeout"] = plan.ActiveFlowTimeout.ValueString()
	}
	if !(plan.CacheEntries.IsNull() || plan.CacheEntries.IsUnknown()) && (state == nil || !plan.CacheEntries.Equal(state.CacheEntries)) {
		body["cache-entries"] = plan.CacheEntries.ValueString()
	}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	if !(plan.InactiveFlowTimeout.IsNull() || plan.InactiveFlowTimeout.IsUnknown()) && (state == nil || !plan.InactiveFlowTimeout.Equal(state.InactiveFlowTimeout)) {
		body["inactive-flow-timeout"] = plan.InactiveFlowTimeout.ValueString()
	}
	if !(plan.Interfaces.IsNull() || plan.Interfaces.IsUnknown()) && (state == nil || !plan.Interfaces.Equal(state.Interfaces)) {
		body["interfaces"] = plan.Interfaces.ValueString()
	}
	if !(plan.PacketSampling.IsNull() || plan.PacketSampling.IsUnknown()) && (state == nil || !plan.PacketSampling.Equal(state.PacketSampling)) {
		body["packet-sampling"] = client.FormatBool(plan.PacketSampling.ValueBool())
	}
	if !(plan.SamplingInterval.IsNull() || plan.SamplingInterval.IsUnknown()) && (state == nil || !plan.SamplingInterval.Equal(state.SamplingInterval)) {
		body["sampling-interval"] = client.FormatInt64(plan.SamplingInterval.ValueInt64())
	}
	if !(plan.SamplingSpace.IsNull() || plan.SamplingSpace.IsUnknown()) && (state == nil || !plan.SamplingSpace.Equal(state.SamplingSpace)) {
		body["sampling-space"] = client.FormatInt64(plan.SamplingSpace.ValueInt64())
	}
	obj, err := c.SetSingleton(ctx, "/ip/traffic-flow", body)
	if err != nil {
		diags.AddError("Upsert /ip/traffic-flow failed", err.Error())
		return
	}
	iPTrafficFlowApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/traffic-flow", plan.Router))
}

func iPTrafficFlowApply(ctx context.Context, obj client.Object, m *IPTrafficFlowModel) {
	_ = ctx
	if v, ok := obj["active-flow-timeout"]; ok {
		_ = v
		if v != "" {
			m.ActiveFlowTimeout = types.StringValue(v)
		} else {
			m.ActiveFlowTimeout = types.StringNull()
		}
	}
	if v, ok := obj["cache-entries"]; ok {
		_ = v
		if v != "" {
			m.CacheEntries = types.StringValue(v)
		} else {
			m.CacheEntries = types.StringNull()
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
	if v, ok := obj["inactive-flow-timeout"]; ok {
		_ = v
		if v != "" {
			m.InactiveFlowTimeout = types.StringValue(v)
		} else {
			m.InactiveFlowTimeout = types.StringNull()
		}
	}
	if v, ok := obj["interfaces"]; ok {
		_ = v
		if v != "" {
			m.Interfaces = types.StringValue(v)
		} else {
			m.Interfaces = types.StringNull()
		}
	}
	if v, ok := obj["packet-sampling"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.PacketSampling = types.BoolValue(b)
		} else {
			m.PacketSampling = types.BoolNull()
		}
	}
	if v, ok := obj["sampling-interval"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.SamplingInterval = types.Int64Value(n)
		} else {
			m.SamplingInterval = types.Int64Null()
		}
	}
	if v, ok := obj["sampling-space"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.SamplingSpace = types.Int64Value(n)
		} else {
			m.SamplingSpace = types.Int64Null()
		}
	}
}
