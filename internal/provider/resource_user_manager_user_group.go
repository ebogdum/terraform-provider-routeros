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
	_ resource.Resource                = &UserManagerUserGroupResource{}
	_ resource.ResourceWithImportState = &UserManagerUserGroupResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type UserManagerUserGroupResource struct {
	reg *client.Registry
}

type UserManagerUserGroupModel struct {
	ID          types.String `tfsdk:"id"`
	Attributes  types.String `tfsdk:"attributes"`
	Default     types.String `tfsdk:"default"`
	DefaultName types.String `tfsdk:"default_name"`
	InnerAuths  types.String `tfsdk:"inner_auths"`
	Name        types.String `tfsdk:"name"`
	OuterAuths  types.String `tfsdk:"outer_auths"`
	Router      types.String `tfsdk:"router"`
}

func NewUserManagerUserGroupResource() resource.Resource { return &UserManagerUserGroupResource{} }

func (r *UserManagerUserGroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_manager_user_group"
}

func (r *UserManagerUserGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *UserManagerUserGroupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/user-manager/user/group`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"attributes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"default": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"default_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"inner_auths": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"outer_auths": schema.StringAttribute{
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

func (r *UserManagerUserGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserManagerUserGroupModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Attributes.IsNull() || plan.Attributes.IsUnknown()) {
		body["attributes"] = plan.Attributes.ValueString()
	}
	if !(plan.Default.IsNull() || plan.Default.IsUnknown()) {
		body["default"] = plan.Default.ValueString()
	}
	if !(plan.DefaultName.IsNull() || plan.DefaultName.IsUnknown()) {
		body["default-name"] = plan.DefaultName.ValueString()
	}
	if !(plan.InnerAuths.IsNull() || plan.InnerAuths.IsUnknown()) {
		body["inner-auths"] = plan.InnerAuths.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.OuterAuths.IsNull() || plan.OuterAuths.IsUnknown()) {
		body["outer-auths"] = plan.OuterAuths.ValueString()
	}
	obj, err := c.Add(ctx, "/user-manager/user/group", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /user-manager/user/group failed", err.Error())
		return
	}
	userManagerUserGroupApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserManagerUserGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserManagerUserGroupModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/user-manager/user/group", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /user-manager/user/group failed", err.Error())
		return
	}
	userManagerUserGroupApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserManagerUserGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state UserManagerUserGroupModel
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
	if !plan.Attributes.Equal(state.Attributes) && !plan.Attributes.IsUnknown() {
		body["attributes"] = plan.Attributes.ValueString()
	}
	if !plan.Default.Equal(state.Default) && !plan.Default.IsUnknown() {
		body["default"] = plan.Default.ValueString()
	}
	if !plan.DefaultName.Equal(state.DefaultName) && !plan.DefaultName.IsUnknown() {
		body["default-name"] = plan.DefaultName.ValueString()
	}
	if !plan.InnerAuths.Equal(state.InnerAuths) && !plan.InnerAuths.IsUnknown() {
		body["inner-auths"] = plan.InnerAuths.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.OuterAuths.Equal(state.OuterAuths) && !plan.OuterAuths.IsUnknown() {
		body["outer-auths"] = plan.OuterAuths.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/user-manager/user/group", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /user-manager/user/group failed", err.Error())
			return
		}
		userManagerUserGroupApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserManagerUserGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state UserManagerUserGroupModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/user-manager/user/group", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /user-manager/user/group failed", err.Error())
	}
}

func (r *UserManagerUserGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := userManagerUserGroupLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /user-manager/user/group matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// userManagerUserGroupLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func userManagerUserGroupLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/user-manager/user/group", id)
}

func userManagerUserGroupApply(ctx context.Context, obj client.Object, m *UserManagerUserGroupModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["attributes"]; ok {
		if v != "" {
			m.Attributes = types.StringValue(v)
		} else {
			m.Attributes = types.StringNull()
		}
	}
	if v, ok := obj["default"]; ok {
		if v != "" {
			m.Default = types.StringValue(v)
		} else {
			m.Default = types.StringNull()
		}
	}
	if v, ok := obj["default-name"]; ok {
		if v != "" {
			m.DefaultName = types.StringValue(v)
		} else {
			m.DefaultName = types.StringNull()
		}
	}
	if v, ok := obj["inner-auths"]; ok {
		if v != "" {
			m.InnerAuths = types.StringValue(v)
		} else {
			m.InnerAuths = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["outer-auths"]; ok {
		if v != "" {
			m.OuterAuths = types.StringValue(v)
		} else {
			m.OuterAuths = types.StringNull()
		}
	}
}
