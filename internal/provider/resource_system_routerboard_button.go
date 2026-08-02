package provider

import (
	"context"
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
	_ resource.Resource                = &SystemRouterboardButtonResource{}
	_ resource.ResourceWithImportState = &SystemRouterboardButtonResource{}
)

// SystemRouterboardButtonResource backs the three hardware buttons under
// /system/routerboard: mode-button, wps-button and reset-button.
//
// The menus are separate singletons but carry an identical property set
// (enabled, hold-time, on-event), so one implementation is parameterised by
// menu path rather than copied three times.
//
// A button binding is easy to overlook because the script it fires is usually
// managed separately as routeros_system_script: the script survives a rebuild,
// the binding that invokes it does not.
type SystemRouterboardButtonResource struct {
	reg      *client.Registry
	menuPath string
	typeName string
	button   string
}

type SystemRouterboardButtonModel struct {
	ID       types.String `tfsdk:"id"`
	Enabled  types.Bool   `tfsdk:"enabled"`
	HoldTime types.String `tfsdk:"hold_time"`
	OnEvent  types.String `tfsdk:"on_event"`
	Router   types.String `tfsdk:"router"`
}

func NewSystemRouterboardModeButtonResource() resource.Resource {
	return &SystemRouterboardButtonResource{
		menuPath: "/system/routerboard/mode-button",
		typeName: "_system_routerboard_mode_button",
		button:   "mode",
	}
}

func NewSystemRouterboardWPSButtonResource() resource.Resource {
	return &SystemRouterboardButtonResource{
		menuPath: "/system/routerboard/wps-button",
		typeName: "_system_routerboard_wps_button",
		button:   "WPS",
	}
}

func NewSystemRouterboardResetButtonResource() resource.Resource {
	return &SystemRouterboardButtonResource{
		menuPath: "/system/routerboard/reset-button",
		typeName: "_system_routerboard_reset_button",
		button:   "reset",
	}
}

func (r *SystemRouterboardButtonResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + r.typeName
}

func (r *SystemRouterboardButtonResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemRouterboardButtonResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `" + r.menuPath + "` -- what the " + r.button +
			" button runs when held.\n\nThe script named by `on_event` is a separate object " +
			"(`routeros_system_script`); this resource only manages the binding to it.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "Whether the button fires `on_event` at all.",
			},
			"hold_time": schema.StringAttribute{Optional: true, Computed: true,
				Description: "How long the button must be held for the event to fire, as a RouterOS " +
					"range -- e.g. `0s..1m`. Not a plain duration.",
			},
			"on_event": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Name of the script to run, or a built-in event such as `dark-mode` or " +
					"`wps-accept`. Empty means nothing is bound.",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemRouterboardButtonResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemRouterboardButtonModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	r.upsert(ctx, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemRouterboardButtonResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SystemRouterboardButtonModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state SystemRouterboardButtonModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	r.upsert(ctx, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemRouterboardButtonResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemRouterboardButtonModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, r.menuPath)
	if err != nil {
		resp.Diagnostics.AddError("Read "+r.menuPath+" failed", err.Error())
		return
	}
	systemRouterboardButtonApply(obj, &state)
	state.ID = types.StringValue(stateIDFor(r.menuPath, state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemRouterboardButtonResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state. The binding is
	// left as-is rather than being cleared, matching the other singletons.
}

func (r *SystemRouterboardButtonResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == r.menuPath {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"),
		types.StringValue(stateIDFor(r.menuPath, types.StringValue(routerName))))...)
}

func (r *SystemRouterboardButtonResource) upsert(ctx context.Context, plan, state *SystemRouterboardButtonModel, diags *diagBuf) {
	c := pickClient(r.reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	if !(plan.HoldTime.IsNull() || plan.HoldTime.IsUnknown()) && (state == nil || !plan.HoldTime.Equal(state.HoldTime)) {
		body["hold-time"] = plan.HoldTime.ValueString()
	}
	if !(plan.OnEvent.IsNull() || plan.OnEvent.IsUnknown()) && (state == nil || !plan.OnEvent.Equal(state.OnEvent)) {
		body["on-event"] = plan.OnEvent.ValueString()
	}
	obj, err := c.SetSingleton(ctx, r.menuPath, body)
	if err != nil {
		diags.AddError("Upsert "+r.menuPath+" failed", err.Error())
		return
	}
	systemRouterboardButtonApply(obj, plan)
	plan.ID = types.StringValue(stateIDFor(r.menuPath, plan.Router))
}

func systemRouterboardButtonApply(obj client.Object, m *SystemRouterboardButtonModel) {
	if v, ok := obj["enabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Enabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Enabled = types.BoolValue(true)
		} else {
			m.Enabled = types.BoolNull()
		}
	} else {
		m.Enabled = types.BoolNull()
	}
	if v, ok := obj["hold-time"]; ok && v != "" {
		m.HoldTime = types.StringValue(v)
	} else {
		m.HoldTime = types.StringNull()
	}
	// on-event is routinely the empty string (reset-button ships unbound), which
	// is a meaningful value, not an absent one -- keep it distinct from null.
	if v, ok := obj["on-event"]; ok {
		m.OnEvent = types.StringValue(v)
	} else {
		m.OnEvent = types.StringNull()
	}
}
