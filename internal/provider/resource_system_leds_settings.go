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
	_ resource.Resource                = &SystemLedsSettingsResource{}
	_ resource.ResourceWithImportState = &SystemLedsSettingsResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type SystemLedsSettingsResource struct {
	reg *client.Registry
}

type SystemLedsSettingsModel struct {
	ID         types.String `tfsdk:"id"`
	AllLedsOff types.String `tfsdk:"all_leds_off"`
	Router     types.String `tfsdk:"router"`
}

func NewSystemLedsSettingsResource() resource.Resource { return &SystemLedsSettingsResource{} }

func (r *SystemLedsSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_leds_settings"
}

func (r *SystemLedsSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemLedsSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/system/leds/settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"all_leds_off": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemLedsSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemLedsSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemLedsSettingsUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemLedsSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SystemLedsSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state SystemLedsSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemLedsSettingsUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemLedsSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemLedsSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/system/leds/settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /system/leds/settings failed", err.Error())
		return
	}
	systemLedsSettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/system/leds/settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemLedsSettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *SystemLedsSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/system/leds/settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/system/leds/settings", types.StringValue(routerName))))...)
}

func systemLedsSettingsUpsert(ctx context.Context, reg *client.Registry, plan, state *SystemLedsSettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AllLedsOff.IsNull() || plan.AllLedsOff.IsUnknown()) && (state == nil || !plan.AllLedsOff.Equal(state.AllLedsOff)) {
		body["all-leds-off"] = plan.AllLedsOff.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/system/leds/settings", body)
	if err != nil {
		diags.AddError("Upsert /system/leds/settings failed", err.Error())
		return
	}
	systemLedsSettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/system/leds/settings", plan.Router))
}

func systemLedsSettingsApply(ctx context.Context, obj client.Object, m *SystemLedsSettingsModel) {
	_ = ctx
	if v, ok := obj["all-leds-off"]; ok {
		_ = v
		if v != "" {
			m.AllLedsOff = types.StringValue(v)
		} else {
			m.AllLedsOff = types.StringNull()
		}
	}
}
