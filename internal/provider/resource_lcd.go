package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &LcdResource{}
	_ resource.ResourceWithImportState = &LcdResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type LcdResource struct {
	reg *client.Registry
}

type LcdModel struct {
	ID               types.String `tfsdk:"id"`
	BacklightTimeout types.String `tfsdk:"backlight_timeout"`
	ColorScheme      types.String `tfsdk:"color_scheme"`
	DefaultScreen    types.String `tfsdk:"default_screen"`
	Enabled          types.String `tfsdk:"enabled"`
	ReadOnlyMode     types.String `tfsdk:"read_only_mode"`
	TimeInterval     types.String `tfsdk:"time_interval"`
	TouchScreen      types.String `tfsdk:"touch_screen"`
	Router           types.String `tfsdk:"router"`
}

func NewLcdResource() resource.Resource { return &LcdResource{} }

func (r *LcdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_lcd"
}

func (r *LcdResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *LcdResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Requires LCD-equipped board (e.g. RB1100AHx4)",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"backlight_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time after which LCD touchscreen is turned off",
			},
			"color_scheme": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Changes to color scheme with a dark or light background.",
			},
			"default_screen": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Default screen that is showed after startup.",
			},
			"enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Turns LCD touchscreen on/off. When off, it stops and resets statistics gathering and closes the LCD program.",
			},
			"read_only_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables or disables Read-Only mode. If Read-Only mode is enabled, then menus which can be used to change configuration are hidden.",
			},
			"time_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time interval of displayed interface statistics in Stats screen",
			},
			"touch_screen": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enable/disable touch screen input.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *LcdResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan LcdModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.BacklightTimeout.IsNull() || plan.BacklightTimeout.IsUnknown()) {
		body["backlight-timeout"] = plan.BacklightTimeout.ValueString()
	}
	if !(plan.ColorScheme.IsNull() || plan.ColorScheme.IsUnknown()) {
		body["color-scheme"] = plan.ColorScheme.ValueString()
	}
	if !(plan.DefaultScreen.IsNull() || plan.DefaultScreen.IsUnknown()) {
		body["default-screen"] = plan.DefaultScreen.ValueString()
	}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) {
		body["enabled"] = plan.Enabled.ValueString()
	}
	if !(plan.ReadOnlyMode.IsNull() || plan.ReadOnlyMode.IsUnknown()) {
		body["read-only-mode"] = plan.ReadOnlyMode.ValueString()
	}
	if !(plan.TimeInterval.IsNull() || plan.TimeInterval.IsUnknown()) {
		body["time-interval"] = plan.TimeInterval.ValueString()
	}
	if !(plan.TouchScreen.IsNull() || plan.TouchScreen.IsUnknown()) {
		body["touch-screen"] = plan.TouchScreen.ValueString()
	}
	obj, err := c.Add(ctx, "/lcd", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /lcd failed", err.Error())
		return
	}
	lcdApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *LcdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state LcdModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/lcd", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /lcd failed", err.Error())
		return
	}
	lcdApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *LcdResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state LcdModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !plan.BacklightTimeout.Equal(state.BacklightTimeout) {
		body["backlight-timeout"] = plan.BacklightTimeout.ValueString()
	}
	if !plan.ColorScheme.Equal(state.ColorScheme) {
		body["color-scheme"] = plan.ColorScheme.ValueString()
	}
	if !plan.DefaultScreen.Equal(state.DefaultScreen) {
		body["default-screen"] = plan.DefaultScreen.ValueString()
	}
	if !plan.Enabled.Equal(state.Enabled) {
		body["enabled"] = plan.Enabled.ValueString()
	}
	if !plan.ReadOnlyMode.Equal(state.ReadOnlyMode) {
		body["read-only-mode"] = plan.ReadOnlyMode.ValueString()
	}
	if !plan.TimeInterval.Equal(state.TimeInterval) {
		body["time-interval"] = plan.TimeInterval.ValueString()
	}
	if !plan.TouchScreen.Equal(state.TouchScreen) {
		body["touch-screen"] = plan.TouchScreen.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/lcd", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /lcd failed", err.Error())
			return
		}
		lcdApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *LcdResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state LcdModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/lcd", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /lcd failed", err.Error())
	}
}

func (r *LcdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import formats accepted:
	//   *<id>                            -> bare RouterOS .id on the default router
	//   <router>/*<id>                   -> .id on the named router
	//   <router>/<naturalkey>            -> resolved via List + filter
	//   <naturalkey>                     -> resolved on the default router
	routerName, id := parseImportID(r.reg, req.ID)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	if strings.HasPrefix(id, "*") {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(id))...)
		return
	}
	c := pickClient(r.reg, types.StringValue(routerName), &resp.Diagnostics)
	if c == nil {
		return
	}
	rows, err := lcdLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /lcd matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// lcdLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func lcdLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/lcd", id)
}

func lcdApply(ctx context.Context, obj client.Object, m *LcdModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["backlight-timeout"]; ok {
		_ = v
		if v != "" {
			m.BacklightTimeout = types.StringValue(v)
		} else {
			m.BacklightTimeout = types.StringNull()
		}
	} else {
		m.BacklightTimeout = types.StringNull()
	}
	if v, ok := obj["color-scheme"]; ok {
		_ = v
		if v != "" {
			m.ColorScheme = types.StringValue(v)
		} else {
			m.ColorScheme = types.StringNull()
		}
	} else {
		m.ColorScheme = types.StringNull()
	}
	if v, ok := obj["default-screen"]; ok {
		_ = v
		if v != "" {
			m.DefaultScreen = types.StringValue(v)
		} else {
			m.DefaultScreen = types.StringNull()
		}
	} else {
		m.DefaultScreen = types.StringNull()
	}
	if v, ok := obj["enabled"]; ok {
		_ = v
		if v != "" {
			m.Enabled = types.StringValue(v)
		} else {
			m.Enabled = types.StringNull()
		}
	} else {
		m.Enabled = types.StringNull()
	}
	if v, ok := obj["read-only-mode"]; ok {
		_ = v
		if v != "" {
			m.ReadOnlyMode = types.StringValue(v)
		} else {
			m.ReadOnlyMode = types.StringNull()
		}
	} else {
		m.ReadOnlyMode = types.StringNull()
	}
	if v, ok := obj["time-interval"]; ok {
		_ = v
		if v != "" {
			m.TimeInterval = types.StringValue(v)
		} else {
			m.TimeInterval = types.StringNull()
		}
	} else {
		m.TimeInterval = types.StringNull()
	}
	if v, ok := obj["touch-screen"]; ok {
		_ = v
		if v != "" {
			m.TouchScreen = types.StringValue(v)
		} else {
			m.TouchScreen = types.StringNull()
		}
	} else {
		m.TouchScreen = types.StringNull()
	}
}
