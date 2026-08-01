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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
	"github.com/ebogdum/terraform-provider-routeros/internal/schemautil"
)

var (
	_ resource.Resource                = &ToolTrafficMonitorResource{}
	_ resource.ResourceWithImportState = &ToolTrafficMonitorResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type ToolTrafficMonitorResource struct {
	reg *client.Registry
}

type ToolTrafficMonitorModel struct {
	ID        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Disabled  types.Bool   `tfsdk:"disabled"`
	Interface types.String `tfsdk:"interface"`
	Name      types.String `tfsdk:"name"`
	OnEvent   types.String `tfsdk:"on_event"`
	Threshold types.String `tfsdk:"threshold"`
	Traffic   types.String `tfsdk:"traffic"`
	Trigger   types.String `tfsdk:"trigger"`
	Router    types.String `tfsdk:"router"`
}

func NewToolTrafficMonitorResource() resource.Resource { return &ToolTrafficMonitorResource{} }

func (r *ToolTrafficMonitorResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_traffic_monitor"
}

func (r *ToolTrafficMonitorResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolTrafficMonitorResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/tool/traffic-monitor`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"interface": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"on_event": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"traffic": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "transmitted", "received"}...)},
			},
			"trigger": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "above", "below", "always"}...)},
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *ToolTrafficMonitorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolTrafficMonitorModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.OnEvent.IsNull() || plan.OnEvent.IsUnknown()) {
		body["on-event"] = plan.OnEvent.ValueString()
	}
	if !(plan.Threshold.IsNull() || plan.Threshold.IsUnknown()) {
		body["threshold"] = plan.Threshold.ValueString()
	}
	if !(plan.Traffic.IsNull() || plan.Traffic.IsUnknown()) {
		body["traffic"] = plan.Traffic.ValueString()
	}
	if !(plan.Trigger.IsNull() || plan.Trigger.IsUnknown()) {
		body["trigger"] = plan.Trigger.ValueString()
	}
	obj, err := c.Add(ctx, "/tool/traffic-monitor", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /tool/traffic-monitor failed", err.Error())
		return
	}
	toolTrafficMonitorApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolTrafficMonitorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolTrafficMonitorModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/tool/traffic-monitor", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /tool/traffic-monitor failed", err.Error())
		return
	}
	toolTrafficMonitorApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolTrafficMonitorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ToolTrafficMonitorModel
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
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.OnEvent.Equal(state.OnEvent) && !plan.OnEvent.IsUnknown() {
		body["on-event"] = plan.OnEvent.ValueString()
	}
	if !plan.Threshold.Equal(state.Threshold) && !plan.Threshold.IsUnknown() {
		body["threshold"] = plan.Threshold.ValueString()
	}
	if !plan.Traffic.Equal(state.Traffic) && !plan.Traffic.IsUnknown() {
		body["traffic"] = plan.Traffic.ValueString()
	}
	if !plan.Trigger.Equal(state.Trigger) && !plan.Trigger.IsUnknown() {
		body["trigger"] = plan.Trigger.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/tool/traffic-monitor", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /tool/traffic-monitor failed", err.Error())
			return
		}
		toolTrafficMonitorApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolTrafficMonitorResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ToolTrafficMonitorModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/tool/traffic-monitor", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /tool/traffic-monitor failed", err.Error())
	}
}

func (r *ToolTrafficMonitorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := toolTrafficMonitorLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /tool/traffic-monitor matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// toolTrafficMonitorLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func toolTrafficMonitorLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/tool/traffic-monitor", id)
}

func toolTrafficMonitorApply(ctx context.Context, obj client.Object, m *ToolTrafficMonitorModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["interface"]; ok {
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["on-event"]; ok {
		if v != "" {
			m.OnEvent = types.StringValue(v)
		} else {
			m.OnEvent = types.StringNull()
		}
	}
	if v, ok := obj["threshold"]; ok {
		if v != "" {
			m.Threshold = types.StringValue(v)
		} else {
			m.Threshold = types.StringNull()
		}
	}
	if v, ok := obj["traffic"]; ok {
		if v != "" {
			m.Traffic = types.StringValue(v)
		} else {
			m.Traffic = types.StringNull()
		}
	}
	if v, ok := obj["trigger"]; ok {
		if v != "" {
			m.Trigger = types.StringValue(v)
		} else {
			m.Trigger = types.StringNull()
		}
	}
}
