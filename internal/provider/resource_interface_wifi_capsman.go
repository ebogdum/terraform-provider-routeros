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
	_ resource.Resource                = &InterfaceWifiCapsmanResource{}
	_ resource.ResourceWithImportState = &InterfaceWifiCapsmanResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type InterfaceWifiCapsmanResource struct {
	reg *client.Registry
}

type InterfaceWifiCapsmanModel struct {
	ID      types.String `tfsdk:"id"`
	Enabled types.Bool   `tfsdk:"enabled"`
	Router  types.String `tfsdk:"router"`
}

func NewInterfaceWifiCapsmanResource() resource.Resource { return &InterfaceWifiCapsmanResource{} }

func (r *InterfaceWifiCapsmanResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_wifi_capsman"
}

func (r *InterfaceWifiCapsmanResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceWifiCapsmanResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/wifi/capsman`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"enabled": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceWifiCapsmanResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceWifiCapsmanModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	interfaceWifiCapsmanUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiCapsmanResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InterfaceWifiCapsmanModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	interfaceWifiCapsmanUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiCapsmanResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceWifiCapsmanModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/interface/wifi/capsman")
	if err != nil {
		resp.Diagnostics.AddError("Read /interface/wifi/capsman failed", err.Error())
		return
	}
	interfaceWifiCapsmanApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/interface/wifi/capsman", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceWifiCapsmanResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *InterfaceWifiCapsmanResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/interface/wifi/capsman" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/interface/wifi/capsman", types.StringValue(routerName))))...)
}

func interfaceWifiCapsmanUpsert(ctx context.Context, reg *client.Registry, plan *InterfaceWifiCapsmanModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	obj, err := c.SetSingleton(ctx, "/interface/wifi/capsman", body)
	if err != nil {
		diags.AddError("Upsert /interface/wifi/capsman failed", err.Error())
		return
	}
	interfaceWifiCapsmanApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/interface/wifi/capsman", plan.Router))
}

func interfaceWifiCapsmanApply(ctx context.Context, obj client.Object, m *InterfaceWifiCapsmanModel) {
	_ = ctx
	if v, ok := obj["enabled"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Enabled = types.BoolValue(b)
		} else {
			m.Enabled = types.BoolNull()
		}
	}
}
