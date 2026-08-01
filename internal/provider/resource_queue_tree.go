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
	_ resource.Resource                = &QueueTreeResource{}
	_ resource.ResourceWithImportState = &QueueTreeResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type QueueTreeResource struct {
	reg *client.Registry
}

type QueueTreeModel struct {
	ID             types.String `tfsdk:"id"`
	BucketSize     types.String `tfsdk:"bucket_size"`
	BurstLimit     rosRateValue `tfsdk:"burst_limit"`
	BurstThreshold rosRateValue `tfsdk:"burst_threshold"`
	BurstTime      types.String `tfsdk:"burst_time"`
	Comment        types.String `tfsdk:"comment"`
	Disabled       types.Bool   `tfsdk:"disabled"`
	LimitAt        rosRateValue `tfsdk:"limit_at"`
	MaxLimit       rosRateValue `tfsdk:"max_limit"`
	Name           types.String `tfsdk:"name"`
	PacketMark     types.String `tfsdk:"packet_mark"`
	Parent         types.String `tfsdk:"parent"`
	PlaceBeforeRos types.String `tfsdk:"place_before_ros"`
	Priority       types.String `tfsdk:"priority"`
	Queue          types.String `tfsdk:"queue"`
	Router         types.String `tfsdk:"router"`
	PlaceBefore    types.String `tfsdk:"place_before"`
}

func NewQueueTreeResource() resource.Resource { return &QueueTreeResource{} }

func (r *QueueTreeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_queue_tree"
}

func (r *QueueTreeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *QueueTreeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/queue/tree`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"bucket_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bucket-size`.",
			},
			"burst_limit": schema.StringAttribute{
				CustomType:  rosRateType{},
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"burst_threshold": schema.StringAttribute{
				CustomType:  rosRateType{},
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
				CustomType:  rosRateType{},
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"max_limit": schema.StringAttribute{
				CustomType:  rosRateType{},
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"packet_mark": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"parent": schema.StringAttribute{
				Required:    true,
				Description: "Parent queue (\"global\" for top-level, or another queue's name/id).",
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
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
			"place_before": schema.StringAttribute{
				Computed:    true,
				Description: "RouterOS .id (e.g. *3) of the entry this one should be moved before. Use to enforce explicit ordering on ordered menus.",
			},
		},
	}
}

func (r *QueueTreeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan QueueTreeModel
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
	if !(plan.PacketMark.IsNull() || plan.PacketMark.IsUnknown()) {
		body["packet-mark"] = plan.PacketMark.ValueString()
	}
	if !(plan.Parent.IsNull() || plan.Parent.IsUnknown()) {
		body["parent"] = plan.Parent.ValueString()
	}
	if !(plan.Priority.IsNull() || plan.Priority.IsUnknown()) {
		body["priority"] = plan.Priority.ValueString()
	}
	if !(plan.Queue.IsNull() || plan.Queue.IsUnknown()) {
		body["queue"] = plan.Queue.ValueString()
	}
	if !(plan.BucketSize.IsNull() || plan.BucketSize.IsUnknown()) {
		body["bucket-size"] = plan.BucketSize.ValueString()
	}
	obj, err := c.Add(ctx, "/queue/tree", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /queue/tree failed", err.Error())
		return
	}
	if !(plan.PlaceBefore.IsNull() || plan.PlaceBefore.IsUnknown()) {
		if err := c.Move(ctx, "/queue/tree", obj[".id"], plan.PlaceBefore.ValueString()); err != nil {
			resp.Diagnostics.AddError("Move /queue/tree failed", err.Error())
			return
		}
		obj, err = c.GetByID(ctx, "/queue/tree", obj[".id"])
		if err != nil {
			resp.Diagnostics.AddError("Re-read after move failed", err.Error())
			return
		}
	}
	queueTreeApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *QueueTreeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state QueueTreeModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/queue/tree", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /queue/tree failed", err.Error())
		return
	}
	queueTreeApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *QueueTreeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state QueueTreeModel
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
	if !plan.BurstLimit.Equal(state.BurstLimit) && !plan.BurstLimit.IsUnknown() {
		body["burst-limit"] = plan.BurstLimit.ValueString()
	}
	if !plan.BurstThreshold.Equal(state.BurstThreshold) && !plan.BurstThreshold.IsUnknown() {
		body["burst-threshold"] = plan.BurstThreshold.ValueString()
	}
	if !plan.BurstTime.Equal(state.BurstTime) && !plan.BurstTime.IsUnknown() {
		body["burst-time"] = plan.BurstTime.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.LimitAt.Equal(state.LimitAt) && !plan.LimitAt.IsUnknown() {
		body["limit-at"] = plan.LimitAt.ValueString()
	}
	if !plan.MaxLimit.Equal(state.MaxLimit) && !plan.MaxLimit.IsUnknown() {
		body["max-limit"] = plan.MaxLimit.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.PacketMark.Equal(state.PacketMark) && !plan.PacketMark.IsUnknown() {
		body["packet-mark"] = plan.PacketMark.ValueString()
	}
	if !plan.Parent.Equal(state.Parent) && !plan.Parent.IsUnknown() {
		body["parent"] = plan.Parent.ValueString()
	}
	if !plan.Priority.Equal(state.Priority) && !plan.Priority.IsUnknown() {
		body["priority"] = plan.Priority.ValueString()
	}
	if !plan.Queue.Equal(state.Queue) && !plan.Queue.IsUnknown() {
		body["queue"] = plan.Queue.ValueString()
	}
	if !plan.BucketSize.Equal(state.BucketSize) && !plan.BucketSize.IsUnknown() {
		body["bucket-size"] = plan.BucketSize.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/queue/tree", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /queue/tree failed", err.Error())
			return
		}
		queueTreeApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	if !plan.PlaceBefore.Equal(state.PlaceBefore) && !(plan.PlaceBefore.IsNull() || plan.PlaceBefore.IsUnknown()) {
		if err := c.Move(ctx, "/queue/tree", plan.ID.ValueString(), plan.PlaceBefore.ValueString()); err != nil {
			resp.Diagnostics.AddError("Move /queue/tree failed", err.Error())
			return
		}
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *QueueTreeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state QueueTreeModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/queue/tree", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /queue/tree failed", err.Error())
	}
}

func (r *QueueTreeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := queueTreeLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /queue/tree matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// queueTreeLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func queueTreeLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/queue/tree", id)
}

func queueTreeApply(ctx context.Context, obj client.Object, m *QueueTreeModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["bucket-size"]; ok && v != "" {
		m.BucketSize = types.StringValue(v)
	} else {
		m.BucketSize = types.StringNull()
	}
	if v, ok := obj["burst-limit"]; ok {
		_ = v
		if v != "" {
			m.BurstLimit = newRosRateValue(v)
		} else {
			m.BurstLimit = newRosRateNull()
		}
	} else {
		m.BurstLimit = newRosRateNull()
	}
	if v, ok := obj["burst-threshold"]; ok {
		_ = v
		if v != "" {
			m.BurstThreshold = newRosRateValue(v)
		} else {
			m.BurstThreshold = newRosRateNull()
		}
	} else {
		m.BurstThreshold = newRosRateNull()
	}
	if v, ok := obj["burst-time"]; ok {
		if v != "" {
			m.BurstTime = types.StringValue(v)
		} else {
			m.BurstTime = types.StringNull()
		}
	}
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
	if v, ok := obj["limit-at"]; ok {
		_ = v
		if v != "" {
			m.LimitAt = newRosRateValue(v)
		} else {
			m.LimitAt = newRosRateNull()
		}
	} else {
		m.LimitAt = newRosRateNull()
	}
	if v, ok := obj["max-limit"]; ok {
		_ = v
		if v != "" {
			m.MaxLimit = newRosRateValue(v)
		} else {
			m.MaxLimit = newRosRateNull()
		}
	} else {
		m.MaxLimit = newRosRateNull()
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["packet-mark"]; ok {
		if v != "" {
			m.PacketMark = types.StringValue(v)
		} else {
			m.PacketMark = types.StringNull()
		}
	}
	if v, ok := obj["parent"]; ok {
		if v != "" {
			m.Parent = types.StringValue(v)
		} else {
			m.Parent = types.StringNull()
		}
	}
	if v, ok := obj["place_before"]; ok {
		if v != "" {
			m.PlaceBeforeRos = types.StringValue(v)
		} else {
			m.PlaceBeforeRos = types.StringNull()
		}
	}
	if v, ok := obj["priority"]; ok {
		if v != "" {
			m.Priority = types.StringValue(v)
		} else {
			m.Priority = types.StringNull()
		}
	}
	if v, ok := obj["queue"]; ok {
		if v != "" {
			m.Queue = types.StringValue(v)
		} else {
			m.Queue = types.StringNull()
		}
	}
}
