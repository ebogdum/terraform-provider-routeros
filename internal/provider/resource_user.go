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
	_ resource.Resource                = &UserResource{}
	_ resource.ResourceWithImportState = &UserResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type UserResource struct {
	reg *client.Registry
}

type UserModel struct {
	ID                types.String `tfsdk:"id"`
	Address           types.String `tfsdk:"address"`
	Alias             types.String `tfsdk:"alias"`
	Comment           types.String `tfsdk:"comment"`
	Disabled          types.Bool   `tfsdk:"disabled"`
	Expired           types.Bool   `tfsdk:"expired"`
	Group             types.String `tfsdk:"group"`
	InactivityPolicy  types.String `tfsdk:"inactivity_policy"`
	InactivityTimeout types.String `tfsdk:"inactivity_timeout"`
	LastLoggedIn      types.String `tfsdk:"last_logged_in"`
	Name              types.String `tfsdk:"name"`
	Password          types.String `tfsdk:"password"`
	Type              types.Int64  `tfsdk:"type"`
	Router            types.String `tfsdk:"router"`
	LockoutAck        types.Bool   `tfsdk:"lockout_ack"`
}

func NewUserResource() resource.Resource { return &UserResource{} }

func (r *UserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user"
}

func (r *UserResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *UserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "User accounts. password is sensitive and not round-trippable (RouterOS scrubs it on read).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"alias": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"expired": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"group": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"inactivity_policy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"inactivity_timeout": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"last_logged_in": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"password": schema.StringAttribute{
				Required:    true,
				Sensitive:   true,
				Description: "",
			},
			"type": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"lockout_ack": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Acknowledge that this rule may sever management traffic (required for unconditional input/forward drop/reject/tarpit rules with no match).",
			},
		},
	}
}

func (r *UserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Address.IsNull() || plan.Address.IsUnknown()) {
		body["address"] = plan.Address.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Group.IsNull() || plan.Group.IsUnknown()) {
		body["group"] = plan.Group.ValueString()
	}
	if !(plan.InactivityPolicy.IsNull() || plan.InactivityPolicy.IsUnknown()) {
		body["inactivity-policy"] = plan.InactivityPolicy.ValueString()
	}
	if !(plan.InactivityTimeout.IsNull() || plan.InactivityTimeout.IsUnknown()) {
		body["inactivity-timeout"] = plan.InactivityTimeout.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Password.IsNull() || plan.Password.IsUnknown()) {
		body["password"] = plan.Password.ValueString()
	}
	obj, err := c.Add(ctx, "/user", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /user failed", err.Error())
		return
	}
	userApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/user", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /user failed", err.Error())
		return
	}
	userApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state UserModel
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
	if !plan.Address.Equal(state.Address) && !plan.Address.IsUnknown() {
		body["address"] = plan.Address.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Group.Equal(state.Group) && !plan.Group.IsUnknown() {
		body["group"] = plan.Group.ValueString()
	}
	if !plan.InactivityPolicy.Equal(state.InactivityPolicy) && !plan.InactivityPolicy.IsUnknown() {
		body["inactivity-policy"] = plan.InactivityPolicy.ValueString()
	}
	if !plan.InactivityTimeout.Equal(state.InactivityTimeout) && !plan.InactivityTimeout.IsUnknown() {
		body["inactivity-timeout"] = plan.InactivityTimeout.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Password.Equal(state.Password) && !plan.Password.IsUnknown() {
		body["password"] = plan.Password.ValueString()
	}
	// Block disabling the last admin via Update.
	if v, ok := body["disabled"]; ok && strings.EqualFold(v, "true") {
		userView := client.Object{"name": state.Name.ValueString(), "group": state.Group.ValueString(), "disabled": v}
		if err := schemautil.CheckUserDeleteLockout("/user", userView, "disable", !plan.LockoutAck.IsNull() && plan.LockoutAck.ValueBool()); err != nil {
			resp.Diagnostics.AddError("Refusing to disable admin user", err.Error())
			return
		}
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/user", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /user failed", err.Error())
			return
		}
		userApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state UserModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	// Block deletion of the last admin account unless the user opted in via
	// lockout_ack on the resource definition.
	userView := client.Object{"name": state.Name.ValueString(), "group": state.Group.ValueString()}
	if err := schemautil.CheckUserDeleteLockout("/user", userView, "delete", !state.LockoutAck.IsNull() && state.LockoutAck.ValueBool()); err != nil {
		resp.Diagnostics.AddError("Refusing to delete admin user", err.Error())
		return
	}
	if err := c.Remove(ctx, "/user", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /user failed", err.Error())
	}
}

func (r *UserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := userLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /user matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// userLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func userLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/user", id)
}

func userApply(ctx context.Context, obj client.Object, m *UserModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	// LockoutAck is local-only and not persisted on the wire. Carry the
	// plan's value through to state (Null if the user didn't set it).
	if m.LockoutAck.IsUnknown() {
		m.LockoutAck = types.BoolNull()
	}
	if v, ok := obj["address"]; ok {
		if v != "" {
			m.Address = types.StringValue(v)
		} else {
			m.Address = types.StringNull()
		}
	}
	if v, ok := obj["alias"]; ok {
		if v != "" {
			m.Alias = types.StringValue(v)
		} else {
			m.Alias = types.StringNull()
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
		} else if strings.TrimSpace(v) == "" {
			m.Disabled = types.BoolValue(true)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["expired"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Expired = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Expired = types.BoolValue(true)
		} else {
			m.Expired = types.BoolNull()
		}
	}
	if v, ok := obj["group"]; ok {
		if v != "" {
			m.Group = types.StringValue(v)
		} else {
			m.Group = types.StringNull()
		}
	}
	if v, ok := obj["inactivity-policy"]; ok {
		if v != "" {
			m.InactivityPolicy = types.StringValue(v)
		} else {
			m.InactivityPolicy = types.StringNull()
		}
	}
	if v, ok := obj["inactivity-timeout"]; ok {
		if v != "" {
			m.InactivityTimeout = types.StringValue(v)
		} else {
			m.InactivityTimeout = types.StringNull()
		}
	}
	if v, ok := obj["last-logged-in"]; ok {
		if v != "" {
			m.LastLoggedIn = types.StringValue(v)
		} else {
			m.LastLoggedIn = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.Password already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["password"]; ok && v != "" {
		_ = v
		if v != "" {
			m.Password = types.StringValue(v)
		} else {
			m.Password = types.StringNull()
		}
	} else if m.Password.IsUnknown() {
		m.Password = types.StringNull()
	}
	if v, ok := obj["type"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Type = types.Int64Value(n)
		} else {
			m.Type = types.Int64Null()
		}
	} else {
		m.Type = types.Int64Null()
	}
}
