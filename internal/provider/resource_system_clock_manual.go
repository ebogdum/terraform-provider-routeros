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
	_ resource.Resource                = &SystemClockManualResource{}
	_ resource.ResourceWithImportState = &SystemClockManualResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type SystemClockManualResource struct {
	reg *client.Registry
}

type SystemClockManualModel struct {
	ID       types.String `tfsdk:"id"`
	DstDelta types.String `tfsdk:"dst_delta"`
	DstEnd   types.String `tfsdk:"dst_end"`
	DstStart types.String `tfsdk:"dst_start"`
	TimeZone types.String `tfsdk:"time_zone"`
	Router   types.String `tfsdk:"router"`
}

func NewSystemClockManualResource() resource.Resource { return &SystemClockManualResource{} }

func (r *SystemClockManualResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_clock_manual"
}

func (r *SystemClockManualResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemClockManualResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Singleton for manual time configuration. Curl confirms POST /set with an\nempty body returns 200, but the acc test framework times out — likely\nbecause RouterOS adjusts the clock to a value that breaks TLS validation\non the next request. Skipped from automated acc tests.\n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"dst_delta": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"dst_end": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"dst_start": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"time_zone": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *SystemClockManualResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemClockManualModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemClockManualUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemClockManualResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SystemClockManualModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state SystemClockManualModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemClockManualUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemClockManualResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemClockManualModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/system/clock/manual")
	if err != nil {
		resp.Diagnostics.AddError("Read /system/clock/manual failed", err.Error())
		return
	}
	systemClockManualApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/system/clock/manual", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemClockManualResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *SystemClockManualResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/system/clock/manual" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/system/clock/manual", types.StringValue(routerName))))...)
}

func systemClockManualUpsert(ctx context.Context, reg *client.Registry, plan, state *SystemClockManualModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.DstDelta.IsNull() || plan.DstDelta.IsUnknown()) && (state == nil || !plan.DstDelta.Equal(state.DstDelta)) {
		body["dst-delta"] = plan.DstDelta.ValueString()
	}
	if !(plan.DstEnd.IsNull() || plan.DstEnd.IsUnknown()) && (state == nil || !plan.DstEnd.Equal(state.DstEnd)) {
		body["dst-end"] = plan.DstEnd.ValueString()
	}
	if !(plan.DstStart.IsNull() || plan.DstStart.IsUnknown()) && (state == nil || !plan.DstStart.Equal(state.DstStart)) {
		body["dst-start"] = plan.DstStart.ValueString()
	}
	if !(plan.TimeZone.IsNull() || plan.TimeZone.IsUnknown()) && (state == nil || !plan.TimeZone.Equal(state.TimeZone)) {
		body["time-zone"] = plan.TimeZone.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/system/clock/manual", body)
	if err != nil {
		diags.AddError("Upsert /system/clock/manual failed", err.Error())
		return
	}
	systemClockManualApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/system/clock/manual", plan.Router))
}

func systemClockManualApply(ctx context.Context, obj client.Object, m *SystemClockManualModel) {
	_ = ctx
	if v, ok := obj["dst-delta"]; ok {
		_ = v
		if v != "" {
			m.DstDelta = types.StringValue(v)
		} else {
			m.DstDelta = types.StringNull()
		}
	}
	if v, ok := obj["dst-end"]; ok {
		_ = v
		if v != "" {
			m.DstEnd = types.StringValue(v)
		} else {
			m.DstEnd = types.StringNull()
		}
	}
	if v, ok := obj["dst-start"]; ok {
		_ = v
		if v != "" {
			m.DstStart = types.StringValue(v)
		} else {
			m.DstStart = types.StringNull()
		}
	}
	if v, ok := obj["time-zone"]; ok {
		_ = v
		if v != "" {
			m.TimeZone = types.StringValue(v)
		} else {
			m.TimeZone = types.StringNull()
		}
	}
}
