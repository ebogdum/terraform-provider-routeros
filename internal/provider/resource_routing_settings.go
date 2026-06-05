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
	_ resource.Resource                = &RoutingSettingsResource{}
	_ resource.ResourceWithImportState = &RoutingSettingsResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type RoutingSettingsResource struct {
	reg *client.Registry
}

type RoutingSettingsModel struct {
	ID                       types.String `tfsdk:"id"`
	CheckGatewayPingCount    types.Int64  `tfsdk:"check_gateway_ping_count"`
	CheckGatewayPingInterval types.String `tfsdk:"check_gateway_ping_interval"`
	CheckGatewayPingTimeout  types.String `tfsdk:"check_gateway_ping_timeout"`
	PolicyRules              types.List   `tfsdk:"policy_rules"`
	SingleProcess            types.Bool   `tfsdk:"single_process"`
	Router                   types.String `tfsdk:"router"`
}

func NewRoutingSettingsResource() resource.Resource { return &RoutingSettingsResource{} }

func (r *RoutingSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_settings"
}

func (r *RoutingSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/routing/settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"check_gateway_ping_count": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"check_gateway_ping_interval": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"check_gateway_ping_timeout": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"policy_rules": schema.ListAttribute{Optional: true, Computed: true,
				ElementType: types.StringType,
				Description: "",
			},
			"single_process": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *RoutingSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	routingSettingsUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan RoutingSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	routingSettingsUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/routing/settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /routing/settings failed", err.Error())
		return
	}
	routingSettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/routing/settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingSettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *RoutingSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/routing/settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/routing/settings", types.StringValue(routerName))))...)
}

func routingSettingsUpsert(ctx context.Context, reg *client.Registry, plan *RoutingSettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.CheckGatewayPingCount.IsNull() || plan.CheckGatewayPingCount.IsUnknown()) {
		body["check-gateway-ping-count"] = client.FormatInt64(plan.CheckGatewayPingCount.ValueInt64())
	}
	if !(plan.CheckGatewayPingInterval.IsNull() || plan.CheckGatewayPingInterval.IsUnknown()) {
		body["check-gateway-ping-interval"] = plan.CheckGatewayPingInterval.ValueString()
	}
	if !(plan.CheckGatewayPingTimeout.IsNull() || plan.CheckGatewayPingTimeout.IsUnknown()) {
		body["check-gateway-ping-timeout"] = plan.CheckGatewayPingTimeout.ValueString()
	}
	if !(plan.PolicyRules.IsNull() || plan.PolicyRules.IsUnknown()) {
		body["policy-rules"] = encodeStringList(ctx, plan.PolicyRules)
	}
	if !(plan.SingleProcess.IsNull() || plan.SingleProcess.IsUnknown()) {
		body["single-process"] = client.FormatBool(plan.SingleProcess.ValueBool())
	}
	obj, err := c.SetSingleton(ctx, "/routing/settings", body)
	if err != nil {
		diags.AddError("Upsert /routing/settings failed", err.Error())
		return
	}
	routingSettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/routing/settings", plan.Router))
}

func routingSettingsApply(ctx context.Context, obj client.Object, m *RoutingSettingsModel) {
	_ = ctx
	if v, ok := obj["check-gateway-ping-count"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.CheckGatewayPingCount = types.Int64Value(n)
		} else {
			m.CheckGatewayPingCount = types.Int64Null()
		}
	}
	if v, ok := obj["check-gateway-ping-interval"]; ok {
		_ = v
		if v != "" {
			m.CheckGatewayPingInterval = types.StringValue(v)
		} else {
			m.CheckGatewayPingInterval = types.StringNull()
		}
	}
	if v, ok := obj["check-gateway-ping-timeout"]; ok {
		_ = v
		if v != "" {
			m.CheckGatewayPingTimeout = types.StringValue(v)
		} else {
			m.CheckGatewayPingTimeout = types.StringNull()
		}
	}
	if v, ok := obj["policy-rules"]; ok {
		_ = v
		m.PolicyRules = decodeStringList(ctx, v)
	}
	if v, ok := obj["single-process"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.SingleProcess = types.BoolValue(b)
		} else {
			m.SingleProcess = types.BoolNull()
		}
	}
}
