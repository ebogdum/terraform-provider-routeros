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
	_ resource.Resource                = &UserManagerProfileLimitationResource{}
	_ resource.ResourceWithImportState = &UserManagerProfileLimitationResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type UserManagerProfileLimitationResource struct {
	reg *client.Registry
}

type UserManagerProfileLimitationModel struct {
	ID         types.String `tfsdk:"id"`
	Weekdays   types.String `tfsdk:"weekdays"`
	TillTime   types.String `tfsdk:"till_time"`
	Profile    types.String `tfsdk:"profile"`
	Limitation types.String `tfsdk:"limitation"`
	FromTime   types.String `tfsdk:"from_time"`
	Router     types.String `tfsdk:"router"`
}

func NewUserManagerProfileLimitationResource() resource.Resource {
	return &UserManagerProfileLimitationResource{}
}

func (r *UserManagerProfileLimitationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_manager_profile_limitation"
}

func (r *UserManagerProfileLimitationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *UserManagerProfileLimitationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/user-manager/profile-limitation`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"weekdays": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `weekdays`.",
			},
			"till_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `till-time`.",
			},
			"profile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `profile`.",
			},
			"limitation": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `limitation`.",
			},
			"from_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `from-time`.",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *UserManagerProfileLimitationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserManagerProfileLimitationModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.FromTime.IsNull() || plan.FromTime.IsUnknown()) {
		body["from-time"] = plan.FromTime.ValueString()
	}
	if !(plan.Limitation.IsNull() || plan.Limitation.IsUnknown()) {
		body["limitation"] = plan.Limitation.ValueString()
	}
	if !(plan.Profile.IsNull() || plan.Profile.IsUnknown()) {
		body["profile"] = plan.Profile.ValueString()
	}
	if !(plan.TillTime.IsNull() || plan.TillTime.IsUnknown()) {
		body["till-time"] = plan.TillTime.ValueString()
	}
	if !(plan.Weekdays.IsNull() || plan.Weekdays.IsUnknown()) {
		body["weekdays"] = plan.Weekdays.ValueString()
	}
	obj, err := c.Add(ctx, "/user-manager/profile-limitation", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /user-manager/profile-limitation failed", err.Error())
		return
	}
	userManagerProfileLimitationApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserManagerProfileLimitationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserManagerProfileLimitationModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/user-manager/profile-limitation", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /user-manager/profile-limitation failed", err.Error())
		return
	}
	userManagerProfileLimitationApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserManagerProfileLimitationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state UserManagerProfileLimitationModel
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
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/user-manager/profile-limitation", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /user-manager/profile-limitation failed", err.Error())
			return
		}
		userManagerProfileLimitationApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserManagerProfileLimitationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state UserManagerProfileLimitationModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/user-manager/profile-limitation", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /user-manager/profile-limitation failed", err.Error())
	}
}

func (r *UserManagerProfileLimitationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := userManagerProfileLimitationLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /user-manager/profile-limitation matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// userManagerProfileLimitationLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func userManagerProfileLimitationLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/user-manager/profile-limitation", id)
}

func userManagerProfileLimitationApply(ctx context.Context, obj client.Object, m *UserManagerProfileLimitationModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["weekdays"]; ok && v != "" {
		m.Weekdays = types.StringValue(v)
	} else {
		m.Weekdays = types.StringNull()
	}
	if v, ok := obj["till-time"]; ok && v != "" {
		m.TillTime = types.StringValue(v)
	} else {
		m.TillTime = types.StringNull()
	}
	if v, ok := obj["profile"]; ok && v != "" {
		m.Profile = types.StringValue(v)
	} else {
		m.Profile = types.StringNull()
	}
	if v, ok := obj["limitation"]; ok && v != "" {
		m.Limitation = types.StringValue(v)
	} else {
		m.Limitation = types.StringNull()
	}
	if v, ok := obj["from-time"]; ok && v != "" {
		m.FromTime = types.StringValue(v)
	} else {
		m.FromTime = types.StringNull()
	}
}
