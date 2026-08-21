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
	_ resource.Resource                = &UserManagerUserResource{}
	_ resource.ResourceWithImportState = &UserManagerUserResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type UserManagerUserResource struct {
	reg *client.Registry
}

type UserManagerUserModel struct {
	ID          types.String `tfsdk:"id"`
	CallerId    types.String `tfsdk:"caller_id"`
	Name        types.String `tfsdk:"name"`
	Password    types.String `tfsdk:"password"`
	Group       types.String `tfsdk:"group"`
	SharedUsers types.String `tfsdk:"shared_users"`
	OtpSecret   types.String `tfsdk:"otp_secret"`
	Attributes  types.String `tfsdk:"attributes"`
	Comment     types.String `tfsdk:"comment"`
	Router      types.String `tfsdk:"router"`
}

func NewUserManagerUserResource() resource.Resource { return &UserManagerUserResource{} }

func (r *UserManagerUserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_manager_user"
}

func (r *UserManagerUserResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *UserManagerUserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/user-manager/user`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"caller_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `caller-id`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Username.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "User password.",
			},
			"group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "User Manager group, e.g. `default`.",
			},
			"shared_users": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Number of simultaneous sessions permitted for this user.",
			},
			"otp_secret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Base32 TOTP secret used for one-time-password authentication.",
			},
			"attributes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RADIUS attributes returned on authentication, e.g. `Framed-IP-Address:10.0.0.5`.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *UserManagerUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserManagerUserModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Password.IsNull() || plan.Password.IsUnknown()) {
		body["password"] = plan.Password.ValueString()
	}
	if !(plan.Group.IsNull() || plan.Group.IsUnknown()) {
		body["group"] = plan.Group.ValueString()
	}
	if !(plan.SharedUsers.IsNull() || plan.SharedUsers.IsUnknown()) {
		body["shared-users"] = plan.SharedUsers.ValueString()
	}
	if !(plan.OtpSecret.IsNull() || plan.OtpSecret.IsUnknown()) {
		body["otp-secret"] = plan.OtpSecret.ValueString()
	}
	if !(plan.Attributes.IsNull() || plan.Attributes.IsUnknown()) {
		body["attributes"] = plan.Attributes.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.CallerId.IsNull() || plan.CallerId.IsUnknown()) {
		body["caller-id"] = plan.CallerId.ValueString()
	}
	obj, err := c.Add(ctx, "/user-manager/user", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /user-manager/user failed", err.Error())
		return
	}
	userManagerUserApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserManagerUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserManagerUserModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/user-manager/user", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /user-manager/user failed", err.Error())
		return
	}
	userManagerUserApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserManagerUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state UserManagerUserModel
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
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Password.Equal(state.Password) && !plan.Password.IsUnknown() {
		body["password"] = plan.Password.ValueString()
	}
	if !plan.Group.Equal(state.Group) && !plan.Group.IsUnknown() {
		body["group"] = plan.Group.ValueString()
	}
	if !plan.SharedUsers.Equal(state.SharedUsers) && !plan.SharedUsers.IsUnknown() {
		body["shared-users"] = plan.SharedUsers.ValueString()
	}
	if !plan.OtpSecret.Equal(state.OtpSecret) && !plan.OtpSecret.IsUnknown() {
		body["otp-secret"] = plan.OtpSecret.ValueString()
	}
	if !plan.Attributes.Equal(state.Attributes) && !plan.Attributes.IsUnknown() {
		body["attributes"] = plan.Attributes.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.CallerId.Equal(state.CallerId) && !plan.CallerId.IsUnknown() {
		body["caller-id"] = plan.CallerId.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/user-manager/user", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /user-manager/user failed", err.Error())
			return
		}
		userManagerUserApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserManagerUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state UserManagerUserModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/user-manager/user", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /user-manager/user failed", err.Error())
	}
}

func (r *UserManagerUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := userManagerUserLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /user-manager/user matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// userManagerUserLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func userManagerUserLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/user-manager/user", id)
}

func userManagerUserApply(ctx context.Context, obj client.Object, m *UserManagerUserModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["caller-id"]; ok && v != "" {
		m.CallerId = types.StringValue(v)
	} else {
		m.CallerId = types.StringNull()
	}
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["password"]; ok && v != "" {
		m.Password = types.StringValue(v)
	} else {
		m.Password = types.StringNull()
	}
	if v, ok := obj["group"]; ok && v != "" {
		m.Group = types.StringValue(v)
	} else {
		m.Group = types.StringNull()
	}
	if v, ok := obj["shared-users"]; ok && v != "" {
		m.SharedUsers = types.StringValue(v)
	} else {
		m.SharedUsers = types.StringNull()
	}
	if v, ok := obj["otp-secret"]; ok && v != "" {
		m.OtpSecret = types.StringValue(v)
	} else {
		m.OtpSecret = types.StringNull()
	}
	if v, ok := obj["attributes"]; ok && v != "" {
		m.Attributes = types.StringValue(v)
	} else {
		m.Attributes = types.StringNull()
	}
	if v, ok := obj["comment"]; ok && v != "" {
		m.Comment = types.StringValue(v)
	} else {
		m.Comment = types.StringNull()
	}
}
