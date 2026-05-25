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
	_ resource.Resource                = &QueueSimpleResource{}
	_ resource.ResourceWithImportState = &QueueSimpleResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type QueueSimpleResource struct {
	reg *client.Registry
}

type QueueSimpleModel struct {
	ID             types.String `tfsdk:"id"`
	BurstLimit     types.String `tfsdk:"burst_limit"`
	BurstThreshold types.String `tfsdk:"burst_threshold"`
	BurstTime      types.String `tfsdk:"burst_time"`
	Comment        types.String `tfsdk:"comment"`
	Disabled       types.Bool   `tfsdk:"disabled"`
	LimitAt        types.String `tfsdk:"limit_at"`
	MaxLimit       types.String `tfsdk:"max_limit"`
	Name           types.String `tfsdk:"name"`
	PacketMarks    types.String `tfsdk:"packet_marks"`
	Parent         types.String `tfsdk:"parent"`
	PlaceBeforeRos types.String `tfsdk:"place_before_ros"`
	Priority       types.String `tfsdk:"priority"`
	Queue          types.String `tfsdk:"queue"`
	Target         types.String `tfsdk:"target"`
	Router         types.String `tfsdk:"router"`
	PlaceBefore    types.String `tfsdk:"place_before"`
}

func NewQueueSimpleResource() resource.Resource { return &QueueSimpleResource{} }

func (r *QueueSimpleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_queue_simple"
}

func (r *QueueSimpleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *QueueSimpleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/queue/simple`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"burst_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"burst_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"burst_time": schema.StringAttribute{
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
			"limit_at": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"max_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"packet_marks": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"parent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"place_before_ros": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"queue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"target": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
			"place_before": schema.StringAttribute{
				Optional:    true,
				Description: "RouterOS .id (e.g. *3) of the entry this one should be moved before. Use to enforce explicit ordering on ordered menus.",
			},
		},
	}
}

func (r *QueueSimpleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan QueueSimpleModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.BurstLimit.IsNull() || plan.BurstLimit.IsUnknown()) {
		body["burst-limit"] = plan.BurstLimit.ValueString()
	}
	if !(plan.BurstThreshold.IsNull() || plan.BurstThreshold.IsUnknown()) {
		body["burst-threshold"] = plan.BurstThreshold.ValueString()
	}
	if !(plan.BurstTime.IsNull() || plan.BurstTime.IsUnknown()) {
		body["burst-time"] = plan.BurstTime.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.LimitAt.IsNull() || plan.LimitAt.IsUnknown()) {
		body["limit-at"] = plan.LimitAt.ValueString()
	}
	if !(plan.MaxLimit.IsNull() || plan.MaxLimit.IsUnknown()) {
		body["max-limit"] = plan.MaxLimit.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.PacketMarks.IsNull() || plan.PacketMarks.IsUnknown()) {
		body["packet-marks"] = plan.PacketMarks.ValueString()
	}
	if !(plan.Parent.IsNull() || plan.Parent.IsUnknown()) {
		body["parent"] = plan.Parent.ValueString()
	}
	if !(plan.PlaceBeforeRos.IsNull() || plan.PlaceBeforeRos.IsUnknown()) {
		body["place_before"] = plan.PlaceBeforeRos.ValueString()
	}
	if !(plan.Priority.IsNull() || plan.Priority.IsUnknown()) {
		body["priority"] = plan.Priority.ValueString()
	}
	if !(plan.Queue.IsNull() || plan.Queue.IsUnknown()) {
		body["queue"] = plan.Queue.ValueString()
	}
	if !(plan.Target.IsNull() || plan.Target.IsUnknown()) {
		body["target"] = plan.Target.ValueString()
	}
	obj, err := c.Add(ctx, "/queue/simple", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /queue/simple failed", err.Error())
		return
	}
	if !(plan.PlaceBefore.IsNull() || plan.PlaceBefore.IsUnknown()) {
		if err := c.Move(ctx, "/queue/simple", obj[".id"], plan.PlaceBefore.ValueString()); err != nil {
			resp.Diagnostics.AddError("Move /queue/simple failed", err.Error())
			return
		}
		obj, err = c.GetByID(ctx, "/queue/simple", obj[".id"])
		if err != nil {
			resp.Diagnostics.AddError("Re-read after move failed", err.Error())
			return
		}
	}
	queueSimpleApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *QueueSimpleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state QueueSimpleModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/queue/simple", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /queue/simple failed", err.Error())
		return
	}
	queueSimpleApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *QueueSimpleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state QueueSimpleModel
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
	if !plan.BurstLimit.Equal(state.BurstLimit) {
		body["burst-limit"] = plan.BurstLimit.ValueString()
	}
	if !plan.BurstThreshold.Equal(state.BurstThreshold) {
		body["burst-threshold"] = plan.BurstThreshold.ValueString()
	}
	if !plan.BurstTime.Equal(state.BurstTime) {
		body["burst-time"] = plan.BurstTime.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.LimitAt.Equal(state.LimitAt) {
		body["limit-at"] = plan.LimitAt.ValueString()
	}
	if !plan.MaxLimit.Equal(state.MaxLimit) {
		body["max-limit"] = plan.MaxLimit.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.PacketMarks.Equal(state.PacketMarks) {
		body["packet-marks"] = plan.PacketMarks.ValueString()
	}
	if !plan.Parent.Equal(state.Parent) {
		body["parent"] = plan.Parent.ValueString()
	}
	if !plan.PlaceBeforeRos.Equal(state.PlaceBeforeRos) {
		body["place_before"] = plan.PlaceBeforeRos.ValueString()
	}
	if !plan.Priority.Equal(state.Priority) {
		body["priority"] = plan.Priority.ValueString()
	}
	if !plan.Queue.Equal(state.Queue) {
		body["queue"] = plan.Queue.ValueString()
	}
	if !plan.Target.Equal(state.Target) {
		body["target"] = plan.Target.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/queue/simple", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /queue/simple failed", err.Error())
			return
		}
		queueSimpleApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	if !plan.PlaceBefore.Equal(state.PlaceBefore) && !(plan.PlaceBefore.IsNull() || plan.PlaceBefore.IsUnknown()) {
		if err := c.Move(ctx, "/queue/simple", plan.ID.ValueString(), plan.PlaceBefore.ValueString()); err != nil {
			resp.Diagnostics.AddError("Move /queue/simple failed", err.Error())
			return
		}
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *QueueSimpleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state QueueSimpleModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/queue/simple", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /queue/simple failed", err.Error())
	}
}

func (r *QueueSimpleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := queueSimpleLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /queue/simple matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// queueSimpleLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func queueSimpleLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/queue/simple", id)
}

func queueSimpleApply(ctx context.Context, obj client.Object, m *QueueSimpleModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["burst-limit"]; ok {
		_ = v
		if v != "" {
			m.BurstLimit = types.StringValue(v)
		} else {
			m.BurstLimit = types.StringNull()
		}
	} else {
		m.BurstLimit = types.StringNull()
	}
	if v, ok := obj["burst-threshold"]; ok {
		_ = v
		if v != "" {
			m.BurstThreshold = types.StringValue(v)
		} else {
			m.BurstThreshold = types.StringNull()
		}
	} else {
		m.BurstThreshold = types.StringNull()
	}
	if v, ok := obj["burst-time"]; ok {
		_ = v
		if v != "" {
			m.BurstTime = types.StringValue(v)
		} else {
			m.BurstTime = types.StringNull()
		}
	} else {
		m.BurstTime = types.StringNull()
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
	if v, ok := obj["limit-at"]; ok {
		_ = v
		if v != "" {
			m.LimitAt = types.StringValue(v)
		} else {
			m.LimitAt = types.StringNull()
		}
	} else {
		m.LimitAt = types.StringNull()
	}
	if v, ok := obj["max-limit"]; ok {
		_ = v
		if v != "" {
			m.MaxLimit = types.StringValue(v)
		} else {
			m.MaxLimit = types.StringNull()
		}
	} else {
		m.MaxLimit = types.StringNull()
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
	if v, ok := obj["packet-marks"]; ok {
		_ = v
		if v != "" {
			m.PacketMarks = types.StringValue(v)
		} else {
			m.PacketMarks = types.StringNull()
		}
	} else {
		m.PacketMarks = types.StringNull()
	}
	if v, ok := obj["parent"]; ok {
		_ = v
		if v != "" {
			m.Parent = types.StringValue(v)
		} else {
			m.Parent = types.StringNull()
		}
	} else {
		m.Parent = types.StringNull()
	}
	if v, ok := obj["place_before"]; ok {
		_ = v
		if v != "" {
			m.PlaceBeforeRos = types.StringValue(v)
		} else {
			m.PlaceBeforeRos = types.StringNull()
		}
	} else {
		m.PlaceBeforeRos = types.StringNull()
	}
	if v, ok := obj["priority"]; ok {
		_ = v
		if v != "" {
			m.Priority = types.StringValue(v)
		} else {
			m.Priority = types.StringNull()
		}
	} else {
		m.Priority = types.StringNull()
	}
	if v, ok := obj["queue"]; ok {
		_ = v
		if v != "" {
			m.Queue = types.StringValue(v)
		} else {
			m.Queue = types.StringNull()
		}
	} else {
		m.Queue = types.StringNull()
	}
	if v, ok := obj["target"]; ok {
		_ = v
		if v != "" {
			m.Target = types.StringValue(v)
		} else {
			m.Target = types.StringNull()
		}
	} else {
		m.Target = types.StringNull()
	}
}
