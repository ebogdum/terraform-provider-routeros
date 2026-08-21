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
	_ resource.Resource                = &UserManagerResource{}
	_ resource.ResourceWithImportState = &UserManagerResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type UserManagerResource struct {
	reg *client.Registry
}

type UserManagerModel struct {
	ID                 types.String `tfsdk:"id"`
	AccountingPort     types.String `tfsdk:"accounting_port"`
	AuthenticationPort types.String `tfsdk:"authentication_port"`
	Certificate        types.String `tfsdk:"certificate"`
	Enabled            types.String `tfsdk:"enabled"`
	RadsecCertificate  types.String `tfsdk:"radsec_certificate"`
	RequireMessageAuth types.String `tfsdk:"require_message_auth"`
	UseProfiles        types.String `tfsdk:"use_profiles"`
	Router             types.String `tfsdk:"router"`
}

func NewUserManagerResource() resource.Resource { return &UserManagerResource{} }

func (r *UserManagerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_manager"
}

func (r *UserManagerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *UserManagerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/user-manager`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"accounting_port": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"authentication_port": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"certificate": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"enabled": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"radsec_certificate": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"require_message_auth": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"use_profiles": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *UserManagerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserManagerModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	userManagerUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserManagerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan UserManagerModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state UserManagerModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	userManagerUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserManagerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserManagerModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/user-manager")
	if err != nil {
		resp.Diagnostics.AddError("Read /user-manager failed", err.Error())
		return
	}
	userManagerApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/user-manager", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserManagerResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *UserManagerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/user-manager" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/user-manager", types.StringValue(routerName))))...)
}

func userManagerUpsert(ctx context.Context, reg *client.Registry, plan, state *UserManagerModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AccountingPort.IsNull() || plan.AccountingPort.IsUnknown()) && (state == nil || !plan.AccountingPort.Equal(state.AccountingPort)) {
		body["accounting-port"] = plan.AccountingPort.ValueString()
	}
	if !(plan.AuthenticationPort.IsNull() || plan.AuthenticationPort.IsUnknown()) && (state == nil || !plan.AuthenticationPort.Equal(state.AuthenticationPort)) {
		body["authentication-port"] = plan.AuthenticationPort.ValueString()
	}
	if !(plan.Certificate.IsNull() || plan.Certificate.IsUnknown()) && (state == nil || !plan.Certificate.Equal(state.Certificate)) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = plan.Enabled.ValueString()
	}
	if !(plan.RadsecCertificate.IsNull() || plan.RadsecCertificate.IsUnknown()) && (state == nil || !plan.RadsecCertificate.Equal(state.RadsecCertificate)) {
		body["radsec-certificate"] = plan.RadsecCertificate.ValueString()
	}
	if !(plan.RequireMessageAuth.IsNull() || plan.RequireMessageAuth.IsUnknown()) && (state == nil || !plan.RequireMessageAuth.Equal(state.RequireMessageAuth)) {
		body["require-message-auth"] = plan.RequireMessageAuth.ValueString()
	}
	if !(plan.UseProfiles.IsNull() || plan.UseProfiles.IsUnknown()) && (state == nil || !plan.UseProfiles.Equal(state.UseProfiles)) {
		body["use-profiles"] = plan.UseProfiles.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/user-manager", body)
	if err != nil {
		diags.AddError("Upsert /user-manager failed", err.Error())
		return
	}
	userManagerApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/user-manager", plan.Router))
}

func userManagerApply(ctx context.Context, obj client.Object, m *UserManagerModel) {
	_ = ctx
	if v, ok := obj["accounting-port"]; ok {
		_ = v
		if v != "" {
			m.AccountingPort = types.StringValue(v)
		} else {
			m.AccountingPort = types.StringNull()
		}
	}
	if v, ok := obj["authentication-port"]; ok {
		_ = v
		if v != "" {
			m.AuthenticationPort = types.StringValue(v)
		} else {
			m.AuthenticationPort = types.StringNull()
		}
	}
	if v, ok := obj["certificate"]; ok {
		_ = v
		if v != "" {
			m.Certificate = types.StringValue(v)
		} else {
			m.Certificate = types.StringNull()
		}
	}
	if v, ok := obj["enabled"]; ok {
		_ = v
		if v != "" {
			m.Enabled = types.StringValue(v)
		} else {
			m.Enabled = types.StringNull()
		}
	}
	if v, ok := obj["radsec-certificate"]; ok {
		_ = v
		if v != "" {
			m.RadsecCertificate = types.StringValue(v)
		} else {
			m.RadsecCertificate = types.StringNull()
		}
	}
	if v, ok := obj["require-message-auth"]; ok {
		_ = v
		if v != "" {
			m.RequireMessageAuth = types.StringValue(v)
		} else {
			m.RequireMessageAuth = types.StringNull()
		}
	}
	if v, ok := obj["use-profiles"]; ok {
		_ = v
		if v != "" {
			m.UseProfiles = types.StringValue(v)
		} else {
			m.UseProfiles = types.StringNull()
		}
	}
}
