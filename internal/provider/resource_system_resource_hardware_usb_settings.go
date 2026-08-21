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
	_ resource.Resource                = &SystemResourceHardwareUSBSettingsResource{}
	_ resource.ResourceWithImportState = &SystemResourceHardwareUSBSettingsResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemResourceHardwareUSBSettingsResource struct {
	reg *client.Registry
}

type SystemResourceHardwareUSBSettingsModel struct {
	ID            types.String `tfsdk:"id"`
	Authorization types.Bool   `tfsdk:"authorization"`
	Router        types.String `tfsdk:"router"`
}

func NewSystemResourceHardwareUSBSettingsResource() resource.Resource {
	return &SystemResourceHardwareUSBSettingsResource{}
}

func (r *SystemResourceHardwareUSBSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_resource_hardware_usb_settings"
}

func (r *SystemResourceHardwareUSBSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemResourceHardwareUSBSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/system/resource/hardware/usb-settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"authorization": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `authorization`.",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *SystemResourceHardwareUSBSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemResourceHardwareUSBSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemResourceHardwareUSBSettingsUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemResourceHardwareUSBSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SystemResourceHardwareUSBSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state SystemResourceHardwareUSBSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemResourceHardwareUSBSettingsUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemResourceHardwareUSBSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemResourceHardwareUSBSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/system/resource/hardware/usb-settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /system/resource/hardware/usb-settings failed", err.Error())
		return
	}
	systemResourceHardwareUSBSettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/system/resource/hardware/usb-settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemResourceHardwareUSBSettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *SystemResourceHardwareUSBSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/system/resource/hardware/usb-settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/system/resource/hardware/usb-settings", types.StringValue(routerName))))...)
}

func systemResourceHardwareUSBSettingsUpsert(ctx context.Context, reg *client.Registry, plan, state *SystemResourceHardwareUSBSettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Authorization.IsNull() || plan.Authorization.IsUnknown()) && (state == nil || !plan.Authorization.Equal(state.Authorization)) {
		body["authorization"] = client.FormatBool(plan.Authorization.ValueBool())
	}
	obj, err := c.SetSingleton(ctx, "/system/resource/hardware/usb-settings", body)
	if err != nil {
		diags.AddError("Upsert /system/resource/hardware/usb-settings failed", err.Error())
		return
	}
	systemResourceHardwareUSBSettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/system/resource/hardware/usb-settings", plan.Router))
}

func systemResourceHardwareUSBSettingsApply(ctx context.Context, obj client.Object, m *SystemResourceHardwareUSBSettingsModel) {
	_ = ctx
	if v, ok := obj["authorization"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Authorization = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Authorization = types.BoolValue(true)
		} else {
			m.Authorization = types.BoolNull()
		}
	} else {
		m.Authorization = types.BoolNull()
	}
}
