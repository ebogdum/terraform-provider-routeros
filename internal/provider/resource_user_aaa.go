package provider

import (
	"context"
	"fmt"
	"strings"

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
	_ resource.Resource                = &UserAaaResource{}
	_ resource.ResourceWithImportState = &UserAaaResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type UserAaaResource struct {
	reg *client.Registry
}

type UserAaaModel struct {
	ID            types.String `tfsdk:"id"`
	Accounting    types.Bool   `tfsdk:"accounting"`
	DefaultGroup  types.String `tfsdk:"default_group"`
	ExcludeGroups types.String `tfsdk:"exclude_groups"`
	InterimUpdate types.String `tfsdk:"interim_update"`
	UseRADIUS     types.Bool   `tfsdk:"use_radius"`
	Router        types.String `tfsdk:"router"`
}

func NewUserAaaResource() resource.Resource { return &UserAaaResource{} }

func (r *UserAaaResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_aaa"
}

func (r *UserAaaResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *UserAaaResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/user/aaa`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"accounting": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"default_group": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"exclude_groups": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"interim_update": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"use_radius": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *UserAaaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserAaaModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	userAaaUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserAaaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan UserAaaModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state UserAaaModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	userAaaUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserAaaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserAaaModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/user/aaa")
	if err != nil {
		resp.Diagnostics.AddError("Read /user/aaa failed", err.Error())
		return
	}
	userAaaApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/user/aaa", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserAaaResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *UserAaaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/user/aaa" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/user/aaa", types.StringValue(routerName))))...)
}

func userAaaUpsert(ctx context.Context, reg *client.Registry, plan, state *UserAaaModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Accounting.IsNull() || plan.Accounting.IsUnknown()) && (state == nil || !plan.Accounting.Equal(state.Accounting)) {
		body["accounting"] = client.FormatBool(plan.Accounting.ValueBool())
	}
	if !(plan.DefaultGroup.IsNull() || plan.DefaultGroup.IsUnknown()) && (state == nil || !plan.DefaultGroup.Equal(state.DefaultGroup)) {
		body["default-group"] = plan.DefaultGroup.ValueString()
	}
	if !(plan.ExcludeGroups.IsNull() || plan.ExcludeGroups.IsUnknown()) && (state == nil || !plan.ExcludeGroups.Equal(state.ExcludeGroups)) {
		body["exclude-groups"] = plan.ExcludeGroups.ValueString()
	}
	if !(plan.InterimUpdate.IsNull() || plan.InterimUpdate.IsUnknown()) && (state == nil || !plan.InterimUpdate.Equal(state.InterimUpdate)) {
		body["interim-update"] = plan.InterimUpdate.ValueString()
	}
	if !(plan.UseRADIUS.IsNull() || plan.UseRADIUS.IsUnknown()) && (state == nil || !plan.UseRADIUS.Equal(state.UseRADIUS)) {
		body["use-radius"] = client.FormatBool(plan.UseRADIUS.ValueBool())
	}
	obj, err := c.SetSingleton(ctx, "/user/aaa", body)
	if err != nil {
		diags.AddError("Upsert /user/aaa failed", err.Error())
		return
	}
	userAaaApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/user/aaa", plan.Router))
}

func userAaaApply(ctx context.Context, obj client.Object, m *UserAaaModel) {
	_ = ctx
	if v, ok := obj["accounting"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Accounting = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Accounting = types.BoolValue(true)
		} else {
			m.Accounting = types.BoolNull()
		}
	}
	if v, ok := obj["default-group"]; ok {
		_ = v
		if v != "" {
			m.DefaultGroup = types.StringValue(v)
		} else {
			m.DefaultGroup = types.StringNull()
		}
	}
	if v, ok := obj["exclude-groups"]; ok {
		_ = v
		if v != "" {
			m.ExcludeGroups = types.StringValue(v)
		} else {
			m.ExcludeGroups = types.StringNull()
		}
	}
	if v, ok := obj["interim-update"]; ok {
		_ = v
		if v != "" {
			m.InterimUpdate = types.StringValue(v)
		} else {
			m.InterimUpdate = types.StringNull()
		}
	}
	if v, ok := obj["use-radius"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.UseRADIUS = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.UseRADIUS = types.BoolValue(true)
		} else {
			m.UseRADIUS = types.BoolNull()
		}
	}
}
