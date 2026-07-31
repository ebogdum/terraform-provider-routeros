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
	_ resource.Resource                = &UserManagerProfileResource{}
	_ resource.ResourceWithImportState = &UserManagerProfileResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type UserManagerProfileResource struct {
	reg *client.Registry
}

type UserManagerProfileModel struct {
	ID                  types.String `tfsdk:"id"`
	Validity            types.String `tfsdk:"validity"`
	StartsWhen          types.String `tfsdk:"starts_when"`
	Price               types.String `tfsdk:"price"`
	OverrideSharedUsers types.String `tfsdk:"override_shared_users"`
	NameForUsers        types.String `tfsdk:"name_for_users"`
	Name                types.String `tfsdk:"name"`
	Router              types.String `tfsdk:"router"`
}

func NewUserManagerProfileResource() resource.Resource { return &UserManagerProfileResource{} }

func (r *UserManagerProfileResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_manager_profile"
}

func (r *UserManagerProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *UserManagerProfileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/user-manager/profile`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"validity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `validity`.",
			},
			"starts_when": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `starts-when`.",
			},
			"price": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `price`.",
			},
			"override_shared_users": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `override-shared-users`.",
			},
			"name_for_users": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `name-for-users`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `name`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *UserManagerProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserManagerProfileModel
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
	if !(plan.NameForUsers.IsNull() || plan.NameForUsers.IsUnknown()) {
		body["name-for-users"] = plan.NameForUsers.ValueString()
	}
	if !(plan.OverrideSharedUsers.IsNull() || plan.OverrideSharedUsers.IsUnknown()) {
		body["override-shared-users"] = plan.OverrideSharedUsers.ValueString()
	}
	if !(plan.Price.IsNull() || plan.Price.IsUnknown()) {
		body["price"] = plan.Price.ValueString()
	}
	if !(plan.StartsWhen.IsNull() || plan.StartsWhen.IsUnknown()) {
		body["starts-when"] = plan.StartsWhen.ValueString()
	}
	if !(plan.Validity.IsNull() || plan.Validity.IsUnknown()) {
		body["validity"] = plan.Validity.ValueString()
	}
	obj, err := c.Add(ctx, "/user-manager/profile", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /user-manager/profile failed", err.Error())
		return
	}
	userManagerProfileApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserManagerProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserManagerProfileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/user-manager/profile", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /user-manager/profile failed", err.Error())
		return
	}
	userManagerProfileApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserManagerProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state UserManagerProfileModel
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
		obj, err := c.Set(ctx, "/user-manager/profile", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /user-manager/profile failed", err.Error())
			return
		}
		userManagerProfileApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserManagerProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state UserManagerProfileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/user-manager/profile", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /user-manager/profile failed", err.Error())
	}
}

func (r *UserManagerProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := userManagerProfileLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /user-manager/profile matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// userManagerProfileLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func userManagerProfileLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/user-manager/profile", id)
}

func userManagerProfileApply(ctx context.Context, obj client.Object, m *UserManagerProfileModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["validity"]; ok && v != "" {
		m.Validity = types.StringValue(v)
	} else {
		m.Validity = types.StringNull()
	}
	if v, ok := obj["starts-when"]; ok && v != "" {
		m.StartsWhen = types.StringValue(v)
	} else {
		m.StartsWhen = types.StringNull()
	}
	if v, ok := obj["price"]; ok && v != "" {
		m.Price = types.StringValue(v)
	} else {
		m.Price = types.StringNull()
	}
	if v, ok := obj["override-shared-users"]; ok && v != "" {
		m.OverrideSharedUsers = types.StringValue(v)
	} else {
		m.OverrideSharedUsers = types.StringNull()
	}
	if v, ok := obj["name-for-users"]; ok && v != "" {
		m.NameForUsers = types.StringValue(v)
	} else {
		m.NameForUsers = types.StringNull()
	}
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
}
