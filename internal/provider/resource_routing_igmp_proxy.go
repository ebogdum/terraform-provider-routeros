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
	_ resource.Resource                = &RoutingIgmpProxyResource{}
	_ resource.ResourceWithImportState = &RoutingIgmpProxyResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type RoutingIgmpProxyResource struct {
	reg *client.Registry
}

type RoutingIgmpProxyModel struct {
	ID                    types.String `tfsdk:"id"`
	QueryInterval         types.String `tfsdk:"query_interval"`
	QueryResponseInterval types.String `tfsdk:"query_response_interval"`
	QuickLeave            types.Bool   `tfsdk:"quick_leave"`
	Router                types.String `tfsdk:"router"`
}

func NewRoutingIgmpProxyResource() resource.Resource { return &RoutingIgmpProxyResource{} }

func (r *RoutingIgmpProxyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_igmp_proxy"
}

func (r *RoutingIgmpProxyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingIgmpProxyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/routing/igmp-proxy`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"query_interval": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"query_response_interval": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"quick_leave": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *RoutingIgmpProxyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingIgmpProxyModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	routingIgmpProxyUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingIgmpProxyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan RoutingIgmpProxyModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state RoutingIgmpProxyModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	routingIgmpProxyUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingIgmpProxyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingIgmpProxyModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/routing/igmp-proxy")
	if err != nil {
		resp.Diagnostics.AddError("Read /routing/igmp-proxy failed", err.Error())
		return
	}
	routingIgmpProxyApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/routing/igmp-proxy", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingIgmpProxyResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *RoutingIgmpProxyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/routing/igmp-proxy" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/routing/igmp-proxy", types.StringValue(routerName))))...)
}

func routingIgmpProxyUpsert(ctx context.Context, reg *client.Registry, plan, state *RoutingIgmpProxyModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.QueryInterval.IsNull() || plan.QueryInterval.IsUnknown()) && (state == nil || !plan.QueryInterval.Equal(state.QueryInterval)) {
		body["query-interval"] = plan.QueryInterval.ValueString()
	}
	if !(plan.QueryResponseInterval.IsNull() || plan.QueryResponseInterval.IsUnknown()) && (state == nil || !plan.QueryResponseInterval.Equal(state.QueryResponseInterval)) {
		body["query-response-interval"] = plan.QueryResponseInterval.ValueString()
	}
	if !(plan.QuickLeave.IsNull() || plan.QuickLeave.IsUnknown()) && (state == nil || !plan.QuickLeave.Equal(state.QuickLeave)) {
		body["quick-leave"] = client.FormatBool(plan.QuickLeave.ValueBool())
	}
	obj, err := c.SetSingleton(ctx, "/routing/igmp-proxy", body)
	if err != nil {
		diags.AddError("Upsert /routing/igmp-proxy failed", err.Error())
		return
	}
	routingIgmpProxyApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/routing/igmp-proxy", plan.Router))
}

func routingIgmpProxyApply(ctx context.Context, obj client.Object, m *RoutingIgmpProxyModel) {
	_ = ctx
	if v, ok := obj["query-interval"]; ok {
		_ = v
		if v != "" {
			m.QueryInterval = types.StringValue(v)
		} else {
			m.QueryInterval = types.StringNull()
		}
	}
	if v, ok := obj["query-response-interval"]; ok {
		_ = v
		if v != "" {
			m.QueryResponseInterval = types.StringValue(v)
		} else {
			m.QueryResponseInterval = types.StringNull()
		}
	}
	if v, ok := obj["quick-leave"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.QuickLeave = types.BoolValue(b)
		} else {
			m.QuickLeave = types.BoolNull()
		}
	}
}
