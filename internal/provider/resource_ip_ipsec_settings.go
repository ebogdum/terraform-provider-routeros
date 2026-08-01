package provider

import (
	"context"
	"fmt"

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
	_ resource.Resource                = &IPIpsecSettingsResource{}
	_ resource.ResourceWithImportState = &IPIpsecSettingsResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type IPIpsecSettingsResource struct {
	reg *client.Registry
}

type IPIpsecSettingsModel struct {
	ID                  types.String `tfsdk:"id"`
	Accounting          types.Bool   `tfsdk:"accounting"`
	DdosCookieThreshold types.Int64  `tfsdk:"ddos_cookie_threshold"`
	InterimUpdate       types.String `tfsdk:"interim_update"`
	XauthUseRADIUS      types.Bool   `tfsdk:"xauth_use_radius"`
	Router              types.String `tfsdk:"router"`
}

func NewIPIpsecSettingsResource() resource.Resource { return &IPIpsecSettingsResource{} }

func (r *IPIpsecSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_ipsec_settings"
}

func (r *IPIpsecSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPIpsecSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/ipsec/settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"accounting": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"ddos_cookie_threshold": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"interim_update": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"xauth_use_radius": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPIpsecSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPIpsecSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPIpsecSettingsUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIpsecSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPIpsecSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state IPIpsecSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPIpsecSettingsUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIpsecSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPIpsecSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/ipsec/settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/ipsec/settings failed", err.Error())
		return
	}
	iPIpsecSettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/ipsec/settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPIpsecSettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPIpsecSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/ipsec/settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ip/ipsec/settings", types.StringValue(routerName))))...)
}

func iPIpsecSettingsUpsert(ctx context.Context, reg *client.Registry, plan, state *IPIpsecSettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Accounting.IsNull() || plan.Accounting.IsUnknown()) && (state == nil || !plan.Accounting.Equal(state.Accounting)) {
		body["accounting"] = client.FormatBool(plan.Accounting.ValueBool())
	}
	if !(plan.DdosCookieThreshold.IsNull() || plan.DdosCookieThreshold.IsUnknown()) && (state == nil || !plan.DdosCookieThreshold.Equal(state.DdosCookieThreshold)) {
		body["ddos-cookie-threshold"] = client.FormatInt64(plan.DdosCookieThreshold.ValueInt64())
	}
	if !(plan.InterimUpdate.IsNull() || plan.InterimUpdate.IsUnknown()) && (state == nil || !plan.InterimUpdate.Equal(state.InterimUpdate)) {
		body["interim-update"] = plan.InterimUpdate.ValueString()
	}
	if !(plan.XauthUseRADIUS.IsNull() || plan.XauthUseRADIUS.IsUnknown()) && (state == nil || !plan.XauthUseRADIUS.Equal(state.XauthUseRADIUS)) {
		body["xauth-use-radius"] = client.FormatBool(plan.XauthUseRADIUS.ValueBool())
	}
	obj, err := c.SetSingleton(ctx, "/ip/ipsec/settings", body)
	if err != nil {
		diags.AddError("Upsert /ip/ipsec/settings failed", err.Error())
		return
	}
	iPIpsecSettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/ipsec/settings", plan.Router))
}

func iPIpsecSettingsApply(ctx context.Context, obj client.Object, m *IPIpsecSettingsModel) {
	_ = ctx
	if v, ok := obj["accounting"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Accounting = types.BoolValue(b)
		} else {
			m.Accounting = types.BoolNull()
		}
	}
	if v, ok := obj["ddos-cookie-threshold"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.DdosCookieThreshold = types.Int64Value(n)
		} else {
			m.DdosCookieThreshold = types.Int64Null()
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
	if v, ok := obj["xauth-use-radius"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.XauthUseRADIUS = types.BoolValue(b)
		} else {
			m.XauthUseRADIUS = types.BoolNull()
		}
	}
}
