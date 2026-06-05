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
	_ resource.Resource                = &SystemSchedulerResource{}
	_ resource.ResourceWithImportState = &SystemSchedulerResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemSchedulerResource struct {
	reg *client.Registry
}

type SystemSchedulerModel struct {
	ID        types.String `tfsdk:"id"`
	Comment   types.String `tfsdk:"comment"`
	Disabled  types.Bool   `tfsdk:"disabled"`
	Interval  types.String `tfsdk:"interval"`
	Name      types.String `tfsdk:"name"`
	NextRun   types.String `tfsdk:"next_run"`
	OnEvent   types.String `tfsdk:"on_event"`
	Owner     types.String `tfsdk:"owner"`
	Policy    types.String `tfsdk:"policy"`
	RunCount  types.Int64  `tfsdk:"run_count"`
	StartDate types.String `tfsdk:"start_date"`
	StartTime types.String `tfsdk:"start_time"`
	Router    types.String `tfsdk:"router"`
}

func NewSystemSchedulerResource() resource.Resource { return &SystemSchedulerResource{} }

func (r *SystemSchedulerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_scheduler"
}

func (r *SystemSchedulerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *SystemSchedulerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/system/scheduler`.",
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
			"interval": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"next_run": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"on_event": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"owner": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"policy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"run_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"start_date": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"start_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"startup"}...)},
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemSchedulerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemSchedulerModel
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
	if !(plan.Interval.IsNull() || plan.Interval.IsUnknown()) {
		body["interval"] = plan.Interval.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.OnEvent.IsNull() || plan.OnEvent.IsUnknown()) {
		body["on-event"] = plan.OnEvent.ValueString()
	}
	if !(plan.Policy.IsNull() || plan.Policy.IsUnknown()) {
		body["policy"] = plan.Policy.ValueString()
	}
	if !(plan.StartDate.IsNull() || plan.StartDate.IsUnknown()) {
		body["start-date"] = plan.StartDate.ValueString()
	}
	if !(plan.StartTime.IsNull() || plan.StartTime.IsUnknown()) {
		body["start-time"] = plan.StartTime.ValueString()
	}
	obj, err := c.Add(ctx, "/system/scheduler", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /system/scheduler failed", err.Error())
		return
	}
	systemSchedulerApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemSchedulerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemSchedulerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/system/scheduler", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /system/scheduler failed", err.Error())
		return
	}
	systemSchedulerApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemSchedulerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemSchedulerModel
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
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interval.Equal(state.Interval) {
		body["interval"] = plan.Interval.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.OnEvent.Equal(state.OnEvent) {
		body["on-event"] = plan.OnEvent.ValueString()
	}
	if !plan.Policy.Equal(state.Policy) {
		body["policy"] = plan.Policy.ValueString()
	}
	if !plan.StartDate.Equal(state.StartDate) {
		body["start-date"] = plan.StartDate.ValueString()
	}
	if !plan.StartTime.Equal(state.StartTime) {
		body["start-time"] = plan.StartTime.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/system/scheduler", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /system/scheduler failed", err.Error())
			return
		}
		systemSchedulerApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemSchedulerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemSchedulerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/system/scheduler", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /system/scheduler failed", err.Error())
	}
}

func (r *SystemSchedulerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import formats accepted:
	//   *<id>                            -> bare RouterOS .id on the default router
	//   <router>/*<id>                   -> .id on the named router
	//   <router>/<naturalkey>            -> resolved via List + filter
	//   <naturalkey>                     -> resolved on the default router
	id := req.ID
	routerName := ""
	if i := strings.Index(id, "/"); i > 0 && !strings.HasPrefix(id, "*") {
		routerName, id = id[:i], id[i+1:]
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	if strings.HasPrefix(id, "*") {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(id))...)
		return
	}
	c := pickClient(r.reg, types.StringValue(routerName), &resp.Diagnostics)
	if c == nil {
		return
	}
	rows, err := systemSchedulerLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /system/scheduler matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// systemSchedulerLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func systemSchedulerLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/system/scheduler", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func systemSchedulerApply(ctx context.Context, obj client.Object, m *SystemSchedulerModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
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
	if v, ok := obj["interval"]; ok {
		_ = v
		if v != "" {
			m.Interval = types.StringValue(v)
		} else {
			m.Interval = types.StringNull()
		}
	} else {
		m.Interval = types.StringNull()
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
	if v, ok := obj["next-run"]; ok {
		_ = v
		if v != "" {
			m.NextRun = types.StringValue(v)
		} else {
			m.NextRun = types.StringNull()
		}
	} else {
		m.NextRun = types.StringNull()
	}
	if v, ok := obj["on-event"]; ok {
		_ = v
		if v != "" {
			m.OnEvent = types.StringValue(v)
		} else {
			m.OnEvent = types.StringNull()
		}
	} else {
		m.OnEvent = types.StringNull()
	}
	if v, ok := obj["owner"]; ok {
		_ = v
		if v != "" {
			m.Owner = types.StringValue(v)
		} else {
			m.Owner = types.StringNull()
		}
	} else {
		m.Owner = types.StringNull()
	}
	if v, ok := obj["policy"]; ok {
		_ = v
		if v != "" {
			m.Policy = types.StringValue(v)
		} else {
			m.Policy = types.StringNull()
		}
	} else {
		m.Policy = types.StringNull()
	}
	if v, ok := obj["run-count"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RunCount = types.Int64Value(n)
		} else {
			m.RunCount = types.Int64Null()
		}
	} else {
		m.RunCount = types.Int64Null()
	}
	if v, ok := obj["start-date"]; ok {
		_ = v
		if v != "" {
			m.StartDate = types.StringValue(v)
		} else {
			m.StartDate = types.StringNull()
		}
	} else {
		m.StartDate = types.StringNull()
	}
	if v, ok := obj["start-time"]; ok {
		_ = v
		if v != "" {
			m.StartTime = types.StringValue(v)
		} else {
			m.StartTime = types.StringNull()
		}
	} else {
		m.StartTime = types.StringNull()
	}
}
