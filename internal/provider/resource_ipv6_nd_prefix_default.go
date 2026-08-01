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
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &IPv6NdPrefixDefaultResource{}
	_ resource.ResourceWithImportState = &IPv6NdPrefixDefaultResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPv6NdPrefixDefaultResource struct {
	reg *client.Registry
}

type IPv6NdPrefixDefaultModel struct {
	ID                types.String `tfsdk:"id"`
	Autonomous        types.Bool   `tfsdk:"autonomous"`
	Dhcp6PdPreferred  types.Bool   `tfsdk:"dhcp6_pd_preferred"`
	PreferredLifetime types.String `tfsdk:"preferred_lifetime"`
	ValidLifetime     types.String `tfsdk:"valid_lifetime"`
	Router            types.String `tfsdk:"router"`
}

func NewIPv6NdPrefixDefaultResource() resource.Resource { return &IPv6NdPrefixDefaultResource{} }

func (r *IPv6NdPrefixDefaultResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv6_nd_prefix_default"
}

func (r *IPv6NdPrefixDefaultResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPv6NdPrefixDefaultResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ipv6/nd/prefix/default`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"autonomous": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `autonomous`.",
			},
			"dhcp6_pd_preferred": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dhcp6-pd-preferred`.",
			},
			"preferred_lifetime": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `preferred-lifetime`.",
			},
			"valid_lifetime": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `valid-lifetime`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPv6NdPrefixDefaultResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPv6NdPrefixDefaultModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPv6NdPrefixDefaultUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPv6NdPrefixDefaultResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPv6NdPrefixDefaultModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state IPv6NdPrefixDefaultModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPv6NdPrefixDefaultUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPv6NdPrefixDefaultResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPv6NdPrefixDefaultModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ipv6/nd/prefix/default")
	if err != nil {
		resp.Diagnostics.AddError("Read /ipv6/nd/prefix/default failed", err.Error())
		return
	}
	iPv6NdPrefixDefaultApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ipv6/nd/prefix/default", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPv6NdPrefixDefaultResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPv6NdPrefixDefaultResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ipv6/nd/prefix/default" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ipv6/nd/prefix/default", types.StringValue(routerName))))...)
}

func iPv6NdPrefixDefaultUpsert(ctx context.Context, reg *client.Registry, plan, state *IPv6NdPrefixDefaultModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Autonomous.IsNull() || plan.Autonomous.IsUnknown()) && (state == nil || !plan.Autonomous.Equal(state.Autonomous)) {
		body["autonomous"] = client.FormatBool(plan.Autonomous.ValueBool())
	}
	if !(plan.Dhcp6PdPreferred.IsNull() || plan.Dhcp6PdPreferred.IsUnknown()) && (state == nil || !plan.Dhcp6PdPreferred.Equal(state.Dhcp6PdPreferred)) {
		body["dhcp6-pd-preferred"] = client.FormatBool(plan.Dhcp6PdPreferred.ValueBool())
	}
	if !(plan.PreferredLifetime.IsNull() || plan.PreferredLifetime.IsUnknown()) && (state == nil || !plan.PreferredLifetime.Equal(state.PreferredLifetime)) {
		body["preferred-lifetime"] = plan.PreferredLifetime.ValueString()
	}
	if !(plan.ValidLifetime.IsNull() || plan.ValidLifetime.IsUnknown()) && (state == nil || !plan.ValidLifetime.Equal(state.ValidLifetime)) {
		body["valid-lifetime"] = plan.ValidLifetime.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/ipv6/nd/prefix/default", body)
	if err != nil {
		diags.AddError("Upsert /ipv6/nd/prefix/default failed", err.Error())
		return
	}
	iPv6NdPrefixDefaultApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ipv6/nd/prefix/default", plan.Router))
}

func iPv6NdPrefixDefaultApply(ctx context.Context, obj client.Object, m *IPv6NdPrefixDefaultModel) {
	_ = ctx
	if v, ok := obj["autonomous"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Autonomous = types.BoolValue(b)
		} else {
			m.Autonomous = types.BoolNull()
		}
	} else {
		m.Autonomous = types.BoolNull()
	}
	if v, ok := obj["dhcp6-pd-preferred"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Dhcp6PdPreferred = types.BoolValue(b)
		} else {
			m.Dhcp6PdPreferred = types.BoolNull()
		}
	} else {
		m.Dhcp6PdPreferred = types.BoolNull()
	}
	if v, ok := obj["preferred-lifetime"]; ok && v != "" {
		m.PreferredLifetime = types.StringValue(v)
	} else {
		m.PreferredLifetime = types.StringNull()
	}
	if v, ok := obj["valid-lifetime"]; ok && v != "" {
		m.ValidLifetime = types.StringValue(v)
	} else {
		m.ValidLifetime = types.StringNull()
	}
}
