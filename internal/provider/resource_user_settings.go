package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &UserSettingsResource{}
	_ resource.ResourceWithImportState = &UserSettingsResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type UserSettingsResource struct {
	reg *client.Registry
}

type UserSettingsModel struct {
	ID                    types.String `tfsdk:"id"`
	MinimumCategories     types.Int64  `tfsdk:"minimum_categories"`
	MinimumPasswordLength types.Int64  `tfsdk:"minimum_password_length"`
	Router                types.String `tfsdk:"router"`
}

func NewUserSettingsResource() resource.Resource { return &UserSettingsResource{} }

func (r *UserSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_settings"
}

func (r *UserSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *UserSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/user/settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"minimum_categories": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"minimum_password_length": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *UserSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	userSettingsUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan UserSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state UserSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	userSettingsUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/user/settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /user/settings failed", err.Error())
		return
	}
	userSettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/user/settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserSettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *UserSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/user/settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/user/settings", types.StringValue(routerName))))...)
}

func userSettingsUpsert(ctx context.Context, reg *client.Registry, plan, state *UserSettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.MinimumCategories.IsNull() || plan.MinimumCategories.IsUnknown()) && (state == nil || !plan.MinimumCategories.Equal(state.MinimumCategories)) {
		body["minimum-categories"] = client.FormatInt64(plan.MinimumCategories.ValueInt64())
	}
	if !(plan.MinimumPasswordLength.IsNull() || plan.MinimumPasswordLength.IsUnknown()) && (state == nil || !plan.MinimumPasswordLength.Equal(state.MinimumPasswordLength)) {
		body["minimum-password-length"] = client.FormatInt64(plan.MinimumPasswordLength.ValueInt64())
	}
	obj, err := c.SetSingleton(ctx, "/user/settings", body)
	if err != nil {
		diags.AddError("Upsert /user/settings failed", err.Error())
		return
	}
	userSettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/user/settings", plan.Router))
}

func userSettingsApply(ctx context.Context, obj client.Object, m *UserSettingsModel) {
	_ = ctx
	if v, ok := obj["minimum-categories"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MinimumCategories = types.Int64Value(n)
		} else {
			m.MinimumCategories = types.Int64Null()
		}
	}
	if v, ok := obj["minimum-password-length"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MinimumPasswordLength = types.Int64Value(n)
		} else {
			m.MinimumPasswordLength = types.Int64Null()
		}
	}
}
