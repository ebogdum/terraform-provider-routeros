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
	_ resource.Resource                = &InterfaceWifiSteeringResource{}
	_ resource.ResourceWithImportState = &InterfaceWifiSteeringResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceWifiSteeringResource struct {
	reg *client.Registry
}

type InterfaceWifiSteeringModel struct {
	ID                        types.String `tfsdk:"id"`
	TransitionRequestPeriod   types.String `tfsdk:"transition_request_period"`
	X2gProbeDelay             types.String `tfsdk:"x2g_probe_delay"`
	Comment                   types.String `tfsdk:"comment"`
	Disabled                  types.Bool   `tfsdk:"disabled"`
	Name                      types.String `tfsdk:"name"`
	NeighborGroup             types.String `tfsdk:"neighbor_group"`
	NeighborGroups            types.String `tfsdk:"neighbor_groups"`
	Rrm                       types.String `tfsdk:"rrm"`
	TransitionRequestCount    types.String `tfsdk:"transition_request_count"`
	TransitionThreshold       types.String `tfsdk:"transition_threshold"`
	TransitionThresholdPeriod types.String `tfsdk:"transition_threshold_period"`
	TransitionThresholdTime   types.String `tfsdk:"transition_threshold_time"`
	TransitionTime            types.String `tfsdk:"transition_time"`
	Wnm                       types.String `tfsdk:"wnm"`
	Router                    types.String `tfsdk:"router"`
}

func NewInterfaceWifiSteeringResource() resource.Resource { return &InterfaceWifiSteeringResource{} }

func (r *InterfaceWifiSteeringResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_wifi_steering"
}

func (r *InterfaceWifiSteeringResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceWifiSteeringResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/wifi/steering`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"transition_request_period": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `transition-request-period`.",
			},
			"x2g_probe_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"neighbor_group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"neighbor_groups": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rrm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"transition_request_count": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"transition_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"transition_threshold_period": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"transition_threshold_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"transition_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wnm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceWifiSteeringResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceWifiSteeringModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.X2gProbeDelay.IsNull() || plan.X2gProbeDelay.IsUnknown()) {
		body["2g-probe-delay"] = plan.X2gProbeDelay.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.NeighborGroup.IsNull() || plan.NeighborGroup.IsUnknown()) {
		body["neighbor-group"] = plan.NeighborGroup.ValueString()
	}
	if !(plan.NeighborGroups.IsNull() || plan.NeighborGroups.IsUnknown()) {
		body["neighbor-groups"] = plan.NeighborGroups.ValueString()
	}
	if !(plan.Rrm.IsNull() || plan.Rrm.IsUnknown()) {
		body["rrm"] = plan.Rrm.ValueString()
	}
	if !(plan.TransitionRequestCount.IsNull() || plan.TransitionRequestCount.IsUnknown()) {
		body["transition-request-count"] = plan.TransitionRequestCount.ValueString()
	}
	if !(plan.TransitionThreshold.IsNull() || plan.TransitionThreshold.IsUnknown()) {
		body["transition-threshold"] = plan.TransitionThreshold.ValueString()
	}
	if !(plan.TransitionThresholdPeriod.IsNull() || plan.TransitionThresholdPeriod.IsUnknown()) {
		body["transition-threshold-period"] = plan.TransitionThresholdPeriod.ValueString()
	}
	if !(plan.TransitionThresholdTime.IsNull() || plan.TransitionThresholdTime.IsUnknown()) {
		body["transition-threshold-time"] = plan.TransitionThresholdTime.ValueString()
	}
	if !(plan.TransitionTime.IsNull() || plan.TransitionTime.IsUnknown()) {
		body["transition-time"] = plan.TransitionTime.ValueString()
	}
	if !(plan.Wnm.IsNull() || plan.Wnm.IsUnknown()) {
		body["wnm"] = plan.Wnm.ValueString()
	}
	if !(plan.TransitionRequestPeriod.IsNull() || plan.TransitionRequestPeriod.IsUnknown()) {
		body["transition-request-period"] = plan.TransitionRequestPeriod.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/wifi/steering", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/wifi/steering failed", err.Error())
		return
	}
	interfaceWifiSteeringApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiSteeringResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceWifiSteeringModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/wifi/steering", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/wifi/steering failed", err.Error())
		return
	}
	interfaceWifiSteeringApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceWifiSteeringResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceWifiSteeringModel
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
	if !plan.X2gProbeDelay.Equal(state.X2gProbeDelay) && !plan.X2gProbeDelay.IsUnknown() {
		body["2g-probe-delay"] = plan.X2gProbeDelay.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.NeighborGroup.Equal(state.NeighborGroup) && !plan.NeighborGroup.IsUnknown() {
		body["neighbor-group"] = plan.NeighborGroup.ValueString()
	}
	if !plan.NeighborGroups.Equal(state.NeighborGroups) && !plan.NeighborGroups.IsUnknown() {
		body["neighbor-groups"] = plan.NeighborGroups.ValueString()
	}
	if !plan.Rrm.Equal(state.Rrm) && !plan.Rrm.IsUnknown() {
		body["rrm"] = plan.Rrm.ValueString()
	}
	if !plan.TransitionRequestCount.Equal(state.TransitionRequestCount) && !plan.TransitionRequestCount.IsUnknown() {
		body["transition-request-count"] = plan.TransitionRequestCount.ValueString()
	}
	if !plan.TransitionThreshold.Equal(state.TransitionThreshold) && !plan.TransitionThreshold.IsUnknown() {
		body["transition-threshold"] = plan.TransitionThreshold.ValueString()
	}
	if !plan.TransitionThresholdPeriod.Equal(state.TransitionThresholdPeriod) && !plan.TransitionThresholdPeriod.IsUnknown() {
		body["transition-threshold-period"] = plan.TransitionThresholdPeriod.ValueString()
	}
	if !plan.TransitionThresholdTime.Equal(state.TransitionThresholdTime) && !plan.TransitionThresholdTime.IsUnknown() {
		body["transition-threshold-time"] = plan.TransitionThresholdTime.ValueString()
	}
	if !plan.TransitionTime.Equal(state.TransitionTime) && !plan.TransitionTime.IsUnknown() {
		body["transition-time"] = plan.TransitionTime.ValueString()
	}
	if !plan.Wnm.Equal(state.Wnm) && !plan.Wnm.IsUnknown() {
		body["wnm"] = plan.Wnm.ValueString()
	}
	if !plan.TransitionRequestPeriod.Equal(state.TransitionRequestPeriod) && !plan.TransitionRequestPeriod.IsUnknown() {
		body["transition-request-period"] = plan.TransitionRequestPeriod.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/wifi/steering", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/wifi/steering failed", err.Error())
			return
		}
		interfaceWifiSteeringApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiSteeringResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceWifiSteeringModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/wifi/steering", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/wifi/steering failed", err.Error())
	}
}

func (r *InterfaceWifiSteeringResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceWifiSteeringLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/wifi/steering matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceWifiSteeringLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceWifiSteeringLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/wifi/steering", id)
}

func interfaceWifiSteeringApply(ctx context.Context, obj client.Object, m *InterfaceWifiSteeringModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["transition-request-period"]; ok && v != "" {
		m.TransitionRequestPeriod = types.StringValue(v)
	} else {
		m.TransitionRequestPeriod = types.StringNull()
	}
	if v, ok := obj["2g-probe-delay"]; ok {
		_ = v
		if v != "" {
			m.X2gProbeDelay = types.StringValue(v)
		} else {
			m.X2gProbeDelay = types.StringNull()
		}
	} else {
		m.X2gProbeDelay = types.StringNull()
	}
	if v, ok := obj["comment"]; ok {
		_ = v
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	} else {
		m.Comment = types.StringNull()
	}
	if v, ok := obj["disabled"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	} else {
		m.Disabled = types.BoolNull()
	}
	if v, ok := obj["name"]; ok {
		_ = v
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["neighbor-group"]; ok {
		_ = v
		if v != "" {
			m.NeighborGroup = types.StringValue(v)
		} else {
			m.NeighborGroup = types.StringNull()
		}
	} else {
		m.NeighborGroup = types.StringNull()
	}
	if v, ok := obj["neighbor-groups"]; ok {
		_ = v
		if v != "" {
			m.NeighborGroups = types.StringValue(v)
		} else {
			m.NeighborGroups = types.StringNull()
		}
	} else {
		m.NeighborGroups = types.StringNull()
	}
	if v, ok := obj["rrm"]; ok {
		_ = v
		if v != "" {
			m.Rrm = types.StringValue(v)
		} else {
			m.Rrm = types.StringNull()
		}
	} else {
		m.Rrm = types.StringNull()
	}
	if v, ok := obj["transition-request-count"]; ok {
		_ = v
		if v != "" {
			m.TransitionRequestCount = types.StringValue(v)
		} else {
			m.TransitionRequestCount = types.StringNull()
		}
	} else {
		m.TransitionRequestCount = types.StringNull()
	}
	if v, ok := obj["transition-threshold"]; ok {
		_ = v
		if v != "" {
			m.TransitionThreshold = types.StringValue(v)
		} else {
			m.TransitionThreshold = types.StringNull()
		}
	} else {
		m.TransitionThreshold = types.StringNull()
	}
	if v, ok := obj["transition-threshold-period"]; ok {
		_ = v
		if v != "" {
			m.TransitionThresholdPeriod = types.StringValue(v)
		} else {
			m.TransitionThresholdPeriod = types.StringNull()
		}
	} else {
		m.TransitionThresholdPeriod = types.StringNull()
	}
	if v, ok := obj["transition-threshold-time"]; ok {
		_ = v
		if v != "" {
			m.TransitionThresholdTime = types.StringValue(v)
		} else {
			m.TransitionThresholdTime = types.StringNull()
		}
	} else {
		m.TransitionThresholdTime = types.StringNull()
	}
	if v, ok := obj["transition-time"]; ok {
		_ = v
		if v != "" {
			m.TransitionTime = types.StringValue(v)
		} else {
			m.TransitionTime = types.StringNull()
		}
	} else {
		m.TransitionTime = types.StringNull()
	}
	if v, ok := obj["wnm"]; ok {
		_ = v
		if v != "" {
			m.Wnm = types.StringValue(v)
		} else {
			m.Wnm = types.StringNull()
		}
	} else {
		m.Wnm = types.StringNull()
	}
}
