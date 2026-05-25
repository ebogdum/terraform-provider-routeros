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
	_ resource.Resource                = &SystemGpsResource{}
	_ resource.ResourceWithImportState = &SystemGpsResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemGpsResource struct {
	reg *client.Registry
}

type SystemGpsModel struct {
	ID               types.String `tfsdk:"id"`
	Channel          types.String `tfsdk:"channel"`
	CoordinateFormat types.String `tfsdk:"coordinate_format"`
	Enabled          types.String `tfsdk:"enabled"`
	GpsAntennaSelect types.String `tfsdk:"gps_antenna_select"`
	InitChannel      types.String `tfsdk:"init_channel"`
	InitString       types.String `tfsdk:"init_string"`
	Port             types.String `tfsdk:"port"`
	SetSystemTime    types.String `tfsdk:"set_system_time"`
	Router           types.String `tfsdk:"router"`
}

func NewSystemGpsResource() resource.Resource { return &SystemGpsResource{} }

func (r *SystemGpsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_gps"
}

func (r *SystemGpsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *SystemGpsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Requires GPS hardware/package",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"channel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port channel used by the device",
			},
			"coordinate_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Which coordinate format to use, \"Decimal Degrees\", \"Degrees Minutes Seconds\" or \"NMEA format DDDMM.MM[MM]\"",
			},
			"enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether GPS is enabled",
			},
			"gps_antenna_select": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Depending on the model. Internal antenna can be selected, if the device has one installed.",
			},
			"init_channel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Channel for init-string execution",
			},
			"init_string": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "AT init string for GPS initialization",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the USB/Serial port where the GPS receiver is connected",
			},
			"set_system_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether to set the router's date and time to one received from GPS.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemGpsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemGpsModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Channel.IsNull() || plan.Channel.IsUnknown()) {
		body["channel"] = plan.Channel.ValueString()
	}
	if !(plan.CoordinateFormat.IsNull() || plan.CoordinateFormat.IsUnknown()) {
		body["coordinate-format"] = plan.CoordinateFormat.ValueString()
	}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) {
		body["enabled"] = plan.Enabled.ValueString()
	}
	if !(plan.GpsAntennaSelect.IsNull() || plan.GpsAntennaSelect.IsUnknown()) {
		body["gps-antenna-select"] = plan.GpsAntennaSelect.ValueString()
	}
	if !(plan.InitChannel.IsNull() || plan.InitChannel.IsUnknown()) {
		body["init-channel"] = plan.InitChannel.ValueString()
	}
	if !(plan.InitString.IsNull() || plan.InitString.IsUnknown()) {
		body["init-string"] = plan.InitString.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.SetSystemTime.IsNull() || plan.SetSystemTime.IsUnknown()) {
		body["set-system-time"] = plan.SetSystemTime.ValueString()
	}
	obj, err := c.Add(ctx, "/system/gps", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /system/gps failed", err.Error())
		return
	}
	systemGpsApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemGpsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemGpsModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/system/gps", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /system/gps failed", err.Error())
		return
	}
	systemGpsApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemGpsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemGpsModel
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
	if !plan.Channel.Equal(state.Channel) {
		body["channel"] = plan.Channel.ValueString()
	}
	if !plan.CoordinateFormat.Equal(state.CoordinateFormat) {
		body["coordinate-format"] = plan.CoordinateFormat.ValueString()
	}
	if !plan.Enabled.Equal(state.Enabled) {
		body["enabled"] = plan.Enabled.ValueString()
	}
	if !plan.GpsAntennaSelect.Equal(state.GpsAntennaSelect) {
		body["gps-antenna-select"] = plan.GpsAntennaSelect.ValueString()
	}
	if !plan.InitChannel.Equal(state.InitChannel) {
		body["init-channel"] = plan.InitChannel.ValueString()
	}
	if !plan.InitString.Equal(state.InitString) {
		body["init-string"] = plan.InitString.ValueString()
	}
	if !plan.Port.Equal(state.Port) {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.SetSystemTime.Equal(state.SetSystemTime) {
		body["set-system-time"] = plan.SetSystemTime.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/system/gps", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /system/gps failed", err.Error())
			return
		}
		systemGpsApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemGpsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemGpsModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/system/gps", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /system/gps failed", err.Error())
	}
}

func (r *SystemGpsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := systemGpsLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /system/gps matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// systemGpsLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func systemGpsLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/system/gps", id)
}

func systemGpsApply(ctx context.Context, obj client.Object, m *SystemGpsModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["channel"]; ok {
		_ = v
		if v != "" {
			m.Channel = types.StringValue(v)
		} else {
			m.Channel = types.StringNull()
		}
	} else {
		m.Channel = types.StringNull()
	}
	if v, ok := obj["coordinate-format"]; ok {
		_ = v
		if v != "" {
			m.CoordinateFormat = types.StringValue(v)
		} else {
			m.CoordinateFormat = types.StringNull()
		}
	} else {
		m.CoordinateFormat = types.StringNull()
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
	if v, ok := obj["gps-antenna-select"]; ok {
		_ = v
		if v != "" {
			m.GpsAntennaSelect = types.StringValue(v)
		} else {
			m.GpsAntennaSelect = types.StringNull()
		}
	} else {
		m.GpsAntennaSelect = types.StringNull()
	}
	if v, ok := obj["init-channel"]; ok {
		_ = v
		if v != "" {
			m.InitChannel = types.StringValue(v)
		} else {
			m.InitChannel = types.StringNull()
		}
	} else {
		m.InitChannel = types.StringNull()
	}
	if v, ok := obj["init-string"]; ok {
		_ = v
		if v != "" {
			m.InitString = types.StringValue(v)
		} else {
			m.InitString = types.StringNull()
		}
	} else {
		m.InitString = types.StringNull()
	}
	if v, ok := obj["port"]; ok {
		_ = v
		if v != "" {
			m.Port = types.StringValue(v)
		} else {
			m.Port = types.StringNull()
		}
	} else {
		m.Port = types.StringNull()
	}
	if v, ok := obj["set-system-time"]; ok {
		_ = v
		if v != "" {
			m.SetSystemTime = types.StringValue(v)
		} else {
			m.SetSystemTime = types.StringNull()
		}
	} else {
		m.SetSystemTime = types.StringNull()
	}
}
