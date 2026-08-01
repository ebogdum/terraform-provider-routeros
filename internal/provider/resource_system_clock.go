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
	_ resource.Resource                = &SystemClockResource{}
	_ resource.ResourceWithImportState = &SystemClockResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type SystemClockResource struct {
	reg *client.Registry
}

type SystemClockModel struct {
	ID                 types.String `tfsdk:"id"`
	Date               types.String `tfsdk:"date"`
	DstActive          types.Bool   `tfsdk:"dst_active"`
	GmtOffset          types.String `tfsdk:"gmt_offset"`
	Time               types.String `tfsdk:"time"`
	TimeZoneAutodetect types.Bool   `tfsdk:"time_zone_autodetect"`
	TimeZoneName       types.String `tfsdk:"time_zone_name"`
	Router             types.String `tfsdk:"router"`
}

func NewSystemClockResource() resource.Resource { return &SystemClockResource{} }

func (r *SystemClockResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_clock"
}

func (r *SystemClockResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemClockResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Setting clock outside automated test scope — would skew router time.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"date": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"dst_active": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"gmt_offset": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"time": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"time_zone_autodetect": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"time_zone_name": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemClockResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemClockModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemClockUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemClockResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SystemClockModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state SystemClockModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemClockUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemClockResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemClockModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/system/clock")
	if err != nil {
		resp.Diagnostics.AddError("Read /system/clock failed", err.Error())
		return
	}
	systemClockApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/system/clock", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemClockResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *SystemClockResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/system/clock" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/system/clock", types.StringValue(routerName))))...)
}

func systemClockUpsert(ctx context.Context, reg *client.Registry, plan, state *SystemClockModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Date.IsNull() || plan.Date.IsUnknown()) && (state == nil || !plan.Date.Equal(state.Date)) {
		body["date"] = plan.Date.ValueString()
	}
	if !(plan.Time.IsNull() || plan.Time.IsUnknown()) && (state == nil || !plan.Time.Equal(state.Time)) {
		body["time"] = plan.Time.ValueString()
	}
	if !(plan.TimeZoneAutodetect.IsNull() || plan.TimeZoneAutodetect.IsUnknown()) && (state == nil || !plan.TimeZoneAutodetect.Equal(state.TimeZoneAutodetect)) {
		body["time-zone-autodetect"] = client.FormatBool(plan.TimeZoneAutodetect.ValueBool())
	}
	if !(plan.TimeZoneName.IsNull() || plan.TimeZoneName.IsUnknown()) && (state == nil || !plan.TimeZoneName.Equal(state.TimeZoneName)) {
		body["time-zone-name"] = plan.TimeZoneName.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/system/clock", body)
	if err != nil {
		diags.AddError("Upsert /system/clock failed", err.Error())
		return
	}
	systemClockApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/system/clock", plan.Router))
}

func systemClockApply(ctx context.Context, obj client.Object, m *SystemClockModel) {
	_ = ctx
	if v, ok := obj["date"]; ok {
		_ = v
		if v != "" {
			m.Date = types.StringValue(v)
		} else {
			m.Date = types.StringNull()
		}
	}
	if v, ok := obj["dst-active"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.DstActive = types.BoolValue(b)
		} else {
			m.DstActive = types.BoolNull()
		}
	}
	if v, ok := obj["gmt-offset"]; ok {
		_ = v
		if v != "" {
			m.GmtOffset = types.StringValue(v)
		} else {
			m.GmtOffset = types.StringNull()
		}
	}
	if v, ok := obj["time"]; ok {
		_ = v
		if v != "" {
			m.Time = types.StringValue(v)
		} else {
			m.Time = types.StringNull()
		}
	}
	if v, ok := obj["time-zone-autodetect"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.TimeZoneAutodetect = types.BoolValue(b)
		} else {
			m.TimeZoneAutodetect = types.BoolNull()
		}
	}
	if v, ok := obj["time-zone-name"]; ok {
		_ = v
		if v != "" {
			m.TimeZoneName = types.StringValue(v)
		} else {
			m.TimeZoneName = types.StringNull()
		}
	}
}
