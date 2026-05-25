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
	_ resource.Resource                = &SystemLoggingResource{}
	_ resource.ResourceWithImportState = &SystemLoggingResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemLoggingResource struct {
	reg *client.Registry
}

type SystemLoggingModel struct {
	ID       types.String `tfsdk:"id"`
	Action   types.String `tfsdk:"action"`
	Comment  types.String `tfsdk:"comment"`
	Default  types.Bool   `tfsdk:"default"`
	Disabled types.Bool   `tfsdk:"disabled"`
	Invalid  types.Bool   `tfsdk:"invalid"`
	Name     types.String `tfsdk:"name"`
	Prefix   types.String `tfsdk:"prefix"`
	Regex    types.String `tfsdk:"regex"`
	Topics   types.String `tfsdk:"topics"`
	Router   types.String `tfsdk:"router"`
}

func NewSystemLoggingResource() resource.Resource { return &SystemLoggingResource{} }

func (r *SystemLoggingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_logging"
}

func (r *SystemLoggingResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *SystemLoggingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/system/logging`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"default": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"prefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"regex": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"topics": schema.StringAttribute{
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

func (r *SystemLoggingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemLoggingModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Action.IsNull() || plan.Action.IsUnknown()) {
		body["action"] = plan.Action.ValueString()
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
	if !(plan.Prefix.IsNull() || plan.Prefix.IsUnknown()) {
		body["prefix"] = plan.Prefix.ValueString()
	}
	if !(plan.Regex.IsNull() || plan.Regex.IsUnknown()) {
		body["regex"] = plan.Regex.ValueString()
	}
	if !(plan.Topics.IsNull() || plan.Topics.IsUnknown()) {
		body["topics"] = plan.Topics.ValueString()
	}
	obj, err := c.Add(ctx, "/system/logging", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /system/logging failed", err.Error())
		return
	}
	systemLoggingApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemLoggingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemLoggingModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/system/logging", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /system/logging failed", err.Error())
		return
	}
	systemLoggingApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemLoggingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemLoggingModel
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
	if !plan.Action.Equal(state.Action) {
		body["action"] = plan.Action.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Prefix.Equal(state.Prefix) {
		body["prefix"] = plan.Prefix.ValueString()
	}
	if !plan.Regex.Equal(state.Regex) {
		body["regex"] = plan.Regex.ValueString()
	}
	if !plan.Topics.Equal(state.Topics) {
		body["topics"] = plan.Topics.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/system/logging", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /system/logging failed", err.Error())
			return
		}
		systemLoggingApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemLoggingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemLoggingModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/system/logging", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /system/logging failed", err.Error())
	}
}

func (r *SystemLoggingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := systemLoggingLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /system/logging matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// systemLoggingLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func systemLoggingLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/system/logging", id)
}

func systemLoggingApply(ctx context.Context, obj client.Object, m *SystemLoggingModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["action"]; ok {
		_ = v
		if v != "" {
			m.Action = types.StringValue(v)
		} else {
			m.Action = types.StringNull()
		}
	} else {
		m.Action = types.StringNull()
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
	if v, ok := obj["default"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else {
			m.Default = types.BoolNull()
		}
	} else {
		m.Default = types.BoolNull()
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
	if v, ok := obj["invalid"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Invalid = types.BoolValue(b)
		} else {
			m.Invalid = types.BoolNull()
		}
	} else {
		m.Invalid = types.BoolNull()
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
	if v, ok := obj["prefix"]; ok {
		_ = v
		if v != "" {
			m.Prefix = types.StringValue(v)
		} else {
			m.Prefix = types.StringNull()
		}
	} else {
		m.Prefix = types.StringNull()
	}
	if v, ok := obj["regex"]; ok {
		_ = v
		if v != "" {
			m.Regex = types.StringValue(v)
		} else {
			m.Regex = types.StringNull()
		}
	} else {
		m.Regex = types.StringNull()
	}
	if v, ok := obj["topics"]; ok {
		_ = v
		if v != "" {
			m.Topics = types.StringValue(v)
		} else {
			m.Topics = types.StringNull()
		}
	} else {
		m.Topics = types.StringNull()
	}
}
