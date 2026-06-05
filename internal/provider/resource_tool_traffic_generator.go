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
	_ resource.Resource                = &ToolTrafficGeneratorResource{}
	_ resource.ResourceWithImportState = &ToolTrafficGeneratorResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type ToolTrafficGeneratorResource struct {
	reg *client.Registry
}

type ToolTrafficGeneratorModel struct {
	ID                                 types.String `tfsdk:"id"`
	LatencyDistributionMax             types.String `tfsdk:"latency_distribution_max"`
	LatencyDistributionMeasureInterval types.String `tfsdk:"latency_distribution_measure_interval"`
	LatencyDistributionSamples         types.Int64  `tfsdk:"latency_distribution_samples"`
	MeasureOutOfOrder                  types.Bool   `tfsdk:"measure_out_of_order"`
	Running                            types.Bool   `tfsdk:"running"`
	StatsSamplesToKeep                 types.Int64  `tfsdk:"stats_samples_to_keep"`
	TestID                             types.Int64  `tfsdk:"test_id"`
	Router                             types.String `tfsdk:"router"`
}

func NewToolTrafficGeneratorResource() resource.Resource { return &ToolTrafficGeneratorResource{} }

func (r *ToolTrafficGeneratorResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_traffic_generator"
}

func (r *ToolTrafficGeneratorResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolTrafficGeneratorResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/tool/traffic-generator`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"latency_distribution_max": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"latency_distribution_measure_interval": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"latency_distribution_samples": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"measure_out_of_order": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"running": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"stats_samples_to_keep": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"test_id": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *ToolTrafficGeneratorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolTrafficGeneratorModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolTrafficGeneratorUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolTrafficGeneratorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ToolTrafficGeneratorModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolTrafficGeneratorUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolTrafficGeneratorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolTrafficGeneratorModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/tool/traffic-generator")
	if err != nil {
		resp.Diagnostics.AddError("Read /tool/traffic-generator failed", err.Error())
		return
	}
	toolTrafficGeneratorApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/tool/traffic-generator", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolTrafficGeneratorResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *ToolTrafficGeneratorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/tool/traffic-generator" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/tool/traffic-generator", types.StringValue(routerName))))...)
}

func toolTrafficGeneratorUpsert(ctx context.Context, reg *client.Registry, plan *ToolTrafficGeneratorModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.LatencyDistributionMax.IsNull() || plan.LatencyDistributionMax.IsUnknown()) {
		body["latency-distribution-max"] = plan.LatencyDistributionMax.ValueString()
	}
	if !(plan.MeasureOutOfOrder.IsNull() || plan.MeasureOutOfOrder.IsUnknown()) {
		body["measure-out-of-order"] = client.FormatBool(plan.MeasureOutOfOrder.ValueBool())
	}
	if !(plan.StatsSamplesToKeep.IsNull() || plan.StatsSamplesToKeep.IsUnknown()) {
		body["stats-samples-to-keep"] = client.FormatInt64(plan.StatsSamplesToKeep.ValueInt64())
	}
	if !(plan.TestID.IsNull() || plan.TestID.IsUnknown()) {
		body["test-id"] = client.FormatInt64(plan.TestID.ValueInt64())
	}
	obj, err := c.SetSingleton(ctx, "/tool/traffic-generator", body)
	if err != nil {
		diags.AddError("Upsert /tool/traffic-generator failed", err.Error())
		return
	}
	toolTrafficGeneratorApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/tool/traffic-generator", plan.Router))
}

func toolTrafficGeneratorApply(ctx context.Context, obj client.Object, m *ToolTrafficGeneratorModel) {
	_ = ctx
	if v, ok := obj["latency-distribution-max"]; ok {
		_ = v
		if v != "" {
			m.LatencyDistributionMax = types.StringValue(v)
		} else {
			m.LatencyDistributionMax = types.StringNull()
		}
	}
	if v, ok := obj["latency-distribution-measure-interval"]; ok {
		_ = v
		if v != "" {
			m.LatencyDistributionMeasureInterval = types.StringValue(v)
		} else {
			m.LatencyDistributionMeasureInterval = types.StringNull()
		}
	}
	if v, ok := obj["latency-distribution-samples"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.LatencyDistributionSamples = types.Int64Value(n)
		} else {
			m.LatencyDistributionSamples = types.Int64Null()
		}
	}
	if v, ok := obj["measure-out-of-order"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.MeasureOutOfOrder = types.BoolValue(b)
		} else {
			m.MeasureOutOfOrder = types.BoolNull()
		}
	}
	if v, ok := obj["running"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Running = types.BoolValue(b)
		} else {
			m.Running = types.BoolNull()
		}
	}
	if v, ok := obj["stats-samples-to-keep"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.StatsSamplesToKeep = types.Int64Value(n)
		} else {
			m.StatsSamplesToKeep = types.Int64Null()
		}
	}
	if v, ok := obj["test-id"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TestID = types.Int64Value(n)
		} else {
			m.TestID = types.Int64Null()
		}
	}
}
