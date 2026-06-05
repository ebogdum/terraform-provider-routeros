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
	_ resource.Resource                = &IPUpnpResource{}
	_ resource.ResourceWithImportState = &IPUpnpResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type IPUpnpResource struct {
	reg *client.Registry
}

type IPUpnpModel struct {
	ID                            types.String `tfsdk:"id"`
	AllowDisableExternalInterface types.Bool   `tfsdk:"allow_disable_external_interface"`
	Enabled                       types.Bool   `tfsdk:"enabled"`
	ShowDummyRule                 types.Bool   `tfsdk:"show_dummy_rule"`
	Router                        types.String `tfsdk:"router"`
}

func NewIPUpnpResource() resource.Resource { return &IPUpnpResource{} }

func (r *IPUpnpResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_upnp"
}

func (r *IPUpnpResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPUpnpResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/upnp`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"allow_disable_external_interface": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"enabled": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"show_dummy_rule": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPUpnpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPUpnpModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPUpnpUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPUpnpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPUpnpModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPUpnpUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPUpnpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPUpnpModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/upnp")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/upnp failed", err.Error())
		return
	}
	iPUpnpApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/upnp", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPUpnpResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPUpnpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/upnp" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ip/upnp", types.StringValue(routerName))))...)
}

func iPUpnpUpsert(ctx context.Context, reg *client.Registry, plan *IPUpnpModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AllowDisableExternalInterface.IsNull() || plan.AllowDisableExternalInterface.IsUnknown()) {
		body["allow-disable-external-interface"] = client.FormatBool(plan.AllowDisableExternalInterface.ValueBool())
	}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	if !(plan.ShowDummyRule.IsNull() || plan.ShowDummyRule.IsUnknown()) {
		body["show-dummy-rule"] = client.FormatBool(plan.ShowDummyRule.ValueBool())
	}
	obj, err := c.SetSingleton(ctx, "/ip/upnp", body)
	if err != nil {
		diags.AddError("Upsert /ip/upnp failed", err.Error())
		return
	}
	iPUpnpApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/upnp", plan.Router))
}

func iPUpnpApply(ctx context.Context, obj client.Object, m *IPUpnpModel) {
	_ = ctx
	if v, ok := obj["allow-disable-external-interface"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AllowDisableExternalInterface = types.BoolValue(b)
		} else {
			m.AllowDisableExternalInterface = types.BoolNull()
		}
	}
	if v, ok := obj["enabled"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Enabled = types.BoolValue(b)
		} else {
			m.Enabled = types.BoolNull()
		}
	}
	if v, ok := obj["show-dummy-rule"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.ShowDummyRule = types.BoolValue(b)
		} else {
			m.ShowDummyRule = types.BoolNull()
		}
	}
}
