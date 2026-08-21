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
	_ resource.Resource                = &UserManagerAdvancedResource{}
	_ resource.ResourceWithImportState = &UserManagerAdvancedResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type UserManagerAdvancedResource struct {
	reg *client.Registry
}

type UserManagerAdvancedModel struct {
	ID                 types.String `tfsdk:"id"`
	PaypalAllow        types.String `tfsdk:"paypal_allow"`
	PaypalCurrency     types.String `tfsdk:"paypal_currency"`
	PaypalPassword     types.String `tfsdk:"paypal_password"`
	PaypalSignature    types.String `tfsdk:"paypal_signature"`
	PaypalUseSandbox   types.String `tfsdk:"paypal_use_sandbox"`
	PaypalUser         types.String `tfsdk:"paypal_user"`
	WebPrivatePassword types.String `tfsdk:"web_private_password"`
	WebPrivateUsername types.String `tfsdk:"web_private_username"`
	Router             types.String `tfsdk:"router"`
}

func NewUserManagerAdvancedResource() resource.Resource { return &UserManagerAdvancedResource{} }

func (r *UserManagerAdvancedResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_manager_advanced"
}

func (r *UserManagerAdvancedResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *UserManagerAdvancedResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/user-manager/advanced`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"paypal_allow": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"paypal_currency": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"paypal_password": schema.StringAttribute{Optional: true,
				Sensitive: true, Computed: true,
				Description: "",
			},
			"paypal_signature": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"paypal_use_sandbox": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"paypal_user": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"web_private_password": schema.StringAttribute{Optional: true,
				Sensitive: true, Computed: true,
				Description: "",
			},
			"web_private_username": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *UserManagerAdvancedResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserManagerAdvancedModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	userManagerAdvancedUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserManagerAdvancedResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan UserManagerAdvancedModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state UserManagerAdvancedModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	userManagerAdvancedUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserManagerAdvancedResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserManagerAdvancedModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/user-manager/advanced")
	if err != nil {
		resp.Diagnostics.AddError("Read /user-manager/advanced failed", err.Error())
		return
	}
	userManagerAdvancedApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/user-manager/advanced", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserManagerAdvancedResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *UserManagerAdvancedResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/user-manager/advanced" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/user-manager/advanced", types.StringValue(routerName))))...)
}

func userManagerAdvancedUpsert(ctx context.Context, reg *client.Registry, plan, state *UserManagerAdvancedModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.PaypalAllow.IsNull() || plan.PaypalAllow.IsUnknown()) && (state == nil || !plan.PaypalAllow.Equal(state.PaypalAllow)) {
		body["paypal-allow"] = plan.PaypalAllow.ValueString()
	}
	if !(plan.PaypalCurrency.IsNull() || plan.PaypalCurrency.IsUnknown()) && (state == nil || !plan.PaypalCurrency.Equal(state.PaypalCurrency)) {
		body["paypal-currency"] = plan.PaypalCurrency.ValueString()
	}
	if !(plan.PaypalPassword.IsNull() || plan.PaypalPassword.IsUnknown()) && (state == nil || !plan.PaypalPassword.Equal(state.PaypalPassword)) {
		body["paypal-password"] = plan.PaypalPassword.ValueString()
	}
	if !(plan.PaypalSignature.IsNull() || plan.PaypalSignature.IsUnknown()) && (state == nil || !plan.PaypalSignature.Equal(state.PaypalSignature)) {
		body["paypal-signature"] = plan.PaypalSignature.ValueString()
	}
	if !(plan.PaypalUseSandbox.IsNull() || plan.PaypalUseSandbox.IsUnknown()) && (state == nil || !plan.PaypalUseSandbox.Equal(state.PaypalUseSandbox)) {
		body["paypal-use-sandbox"] = plan.PaypalUseSandbox.ValueString()
	}
	if !(plan.PaypalUser.IsNull() || plan.PaypalUser.IsUnknown()) && (state == nil || !plan.PaypalUser.Equal(state.PaypalUser)) {
		body["paypal-user"] = plan.PaypalUser.ValueString()
	}
	if !(plan.WebPrivatePassword.IsNull() || plan.WebPrivatePassword.IsUnknown()) && (state == nil || !plan.WebPrivatePassword.Equal(state.WebPrivatePassword)) {
		body["web-private-password"] = plan.WebPrivatePassword.ValueString()
	}
	if !(plan.WebPrivateUsername.IsNull() || plan.WebPrivateUsername.IsUnknown()) && (state == nil || !plan.WebPrivateUsername.Equal(state.WebPrivateUsername)) {
		body["web-private-username"] = plan.WebPrivateUsername.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/user-manager/advanced", body)
	if err != nil {
		diags.AddError("Upsert /user-manager/advanced failed", err.Error())
		return
	}
	userManagerAdvancedApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/user-manager/advanced", plan.Router))
}

func userManagerAdvancedApply(ctx context.Context, obj client.Object, m *UserManagerAdvancedModel) {
	_ = ctx
	if v, ok := obj["paypal-allow"]; ok {
		_ = v
		if v != "" {
			m.PaypalAllow = types.StringValue(v)
		} else {
			m.PaypalAllow = types.StringNull()
		}
	}
	if v, ok := obj["paypal-currency"]; ok {
		_ = v
		if v != "" {
			m.PaypalCurrency = types.StringValue(v)
		} else {
			m.PaypalCurrency = types.StringNull()
		}
	}
	if v, ok := obj["paypal-password"]; ok {
		_ = v
		if v != "" {
			m.PaypalPassword = types.StringValue(v)
		} else {
			m.PaypalPassword = types.StringNull()
		}
	}
	if v, ok := obj["paypal-signature"]; ok {
		_ = v
		if v != "" {
			m.PaypalSignature = types.StringValue(v)
		} else {
			m.PaypalSignature = types.StringNull()
		}
	}
	if v, ok := obj["paypal-use-sandbox"]; ok {
		_ = v
		if v != "" {
			m.PaypalUseSandbox = types.StringValue(v)
		} else {
			m.PaypalUseSandbox = types.StringNull()
		}
	}
	if v, ok := obj["paypal-user"]; ok {
		_ = v
		if v != "" {
			m.PaypalUser = types.StringValue(v)
		} else {
			m.PaypalUser = types.StringNull()
		}
	}
	if v, ok := obj["web-private-password"]; ok {
		_ = v
		if v != "" {
			m.WebPrivatePassword = types.StringValue(v)
		} else {
			m.WebPrivatePassword = types.StringNull()
		}
	}
	if v, ok := obj["web-private-username"]; ok {
		_ = v
		if v != "" {
			m.WebPrivateUsername = types.StringValue(v)
		} else {
			m.WebPrivateUsername = types.StringNull()
		}
	}
}
