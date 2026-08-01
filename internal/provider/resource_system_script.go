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
	_ resource.Resource                = &SystemScriptResource{}
	_ resource.ResourceWithImportState = &SystemScriptResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemScriptResource struct {
	reg *client.Registry
}

type SystemScriptModel struct {
	ID                     types.String `tfsdk:"id"`
	Comment                types.String `tfsdk:"comment"`
	DonTRequirePermissions types.Bool   `tfsdk:"dont_require_permissions"`
	Invalid                types.Bool   `tfsdk:"invalid"`
	LastTimeStarted        types.String `tfsdk:"last_time_started"`
	Name                   types.String `tfsdk:"name"`
	Owner                  types.String `tfsdk:"owner"`
	Policy                 types.Set    `tfsdk:"policy"`
	RunCount               types.Int64  `tfsdk:"run_count"`
	RunScript              types.String `tfsdk:"run_script"`
	Source                 types.String `tfsdk:"source"`
	Router                 types.String `tfsdk:"router"`
}

func NewSystemScriptResource() resource.Resource { return &SystemScriptResource{} }

func (r *SystemScriptResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_script"
}

func (r *SystemScriptResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemScriptResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/system/script`.",
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
			"dont_require_permissions": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"last_time_started": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"owner": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"policy": schema.SetAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "",
			},
			"run_count": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"run_script": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"source": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemScriptResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemScriptModel
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
	if !(plan.DonTRequirePermissions.IsNull() || plan.DonTRequirePermissions.IsUnknown()) {
		body["dont-require-permissions"] = client.FormatBool(plan.DonTRequirePermissions.ValueBool())
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Policy.IsNull() || plan.Policy.IsUnknown()) {
		body["policy"] = encodeStringSet(ctx, plan.Policy, &resp.Diagnostics)
	}
	if !(plan.Source.IsNull() || plan.Source.IsUnknown()) {
		body["source"] = plan.Source.ValueString()
	}
	obj, err := c.Add(ctx, "/system/script", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /system/script failed", err.Error())
		return
	}
	systemScriptApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemScriptResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemScriptModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/system/script", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /system/script failed", err.Error())
		return
	}
	systemScriptApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemScriptResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemScriptModel
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
	if !plan.DonTRequirePermissions.Equal(state.DonTRequirePermissions) && !plan.DonTRequirePermissions.IsUnknown() {
		body["dont-require-permissions"] = client.FormatBool(plan.DonTRequirePermissions.ValueBool())
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Policy.Equal(state.Policy) && !plan.Policy.IsUnknown() {
		body["policy"] = encodeStringSet(ctx, plan.Policy, &resp.Diagnostics)
	}
	if !plan.Source.Equal(state.Source) && !plan.Source.IsUnknown() {
		body["source"] = plan.Source.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/system/script", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /system/script failed", err.Error())
			return
		}
		systemScriptApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemScriptResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemScriptModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/system/script", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /system/script failed", err.Error())
	}
}

func (r *SystemScriptResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := systemScriptLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /system/script matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// systemScriptLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func systemScriptLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/system/script", id)
}

func systemScriptApply(ctx context.Context, obj client.Object, m *SystemScriptModel) {
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
	if v, ok := obj["dont-require-permissions"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.DonTRequirePermissions = types.BoolValue(b)
		} else {
			m.DonTRequirePermissions = types.BoolNull()
		}
	}
	if v, ok := obj["invalid"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Invalid = types.BoolValue(b)
		} else {
			m.Invalid = types.BoolNull()
		}
	}
	if v, ok := obj["last-time-started"]; ok {
		_ = v
		if v != "" {
			m.LastTimeStarted = types.StringValue(v)
		} else {
			m.LastTimeStarted = types.StringNull()
		}
	} else {
		m.LastTimeStarted = types.StringNull()
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
		m.Policy = decodePolicySet(ctx, v)
	} else {
		m.Policy = types.SetNull(types.StringType)
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
	if v, ok := obj["run-script"]; ok {
		_ = v
		if v != "" {
			m.RunScript = types.StringValue(v)
		} else {
			m.RunScript = types.StringNull()
		}
	} else {
		m.RunScript = types.StringNull()
	}
	if v, ok := obj["source"]; ok {
		_ = v
		if v != "" {
			m.Source = types.StringValue(v)
		} else {
			m.Source = types.StringNull()
		}
	} else {
		m.Source = types.StringNull()
	}
}
