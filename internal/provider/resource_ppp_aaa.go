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
	_ resource.Resource                = &PPPAaaResource{}
	_ resource.ResourceWithImportState = &PPPAaaResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type PPPAaaResource struct {
	reg *client.Registry
}

type PPPAaaModel struct {
	ID                      types.String `tfsdk:"id"`
	Accounting              types.Bool   `tfsdk:"accounting"`
	EnableIPV6Accounting    types.Bool   `tfsdk:"enable_ipv6_accounting"`
	InterimUpdate           types.String `tfsdk:"interim_update"`
	UseCircuitIDInNasPortID types.Bool   `tfsdk:"use_circuit_id_in_nas_port_id"`
	UseRADIUS               types.Bool   `tfsdk:"use_radius"`
	Router                  types.String `tfsdk:"router"`
}

func NewPPPAaaResource() resource.Resource { return &PPPAaaResource{} }

func (r *PPPAaaResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ppp_aaa"
}

func (r *PPPAaaResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *PPPAaaResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ppp/aaa`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"accounting": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"enable_ipv6_accounting": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"interim_update": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"use_circuit_id_in_nas_port_id": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"use_radius": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *PPPAaaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PPPAaaModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	pPPAaaUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PPPAaaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan PPPAaaModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	pPPAaaUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PPPAaaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PPPAaaModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ppp/aaa")
	if err != nil {
		resp.Diagnostics.AddError("Read /ppp/aaa failed", err.Error())
		return
	}
	pPPAaaApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ppp/aaa", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *PPPAaaResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *PPPAaaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ppp/aaa" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ppp/aaa", types.StringValue(routerName))))...)
}

func pPPAaaUpsert(ctx context.Context, reg *client.Registry, plan *PPPAaaModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Accounting.IsNull() || plan.Accounting.IsUnknown()) {
		body["accounting"] = client.FormatBool(plan.Accounting.ValueBool())
	}
	if !(plan.EnableIPV6Accounting.IsNull() || plan.EnableIPV6Accounting.IsUnknown()) {
		body["enable-ipv6-accounting"] = client.FormatBool(plan.EnableIPV6Accounting.ValueBool())
	}
	if !(plan.InterimUpdate.IsNull() || plan.InterimUpdate.IsUnknown()) {
		body["interim-update"] = plan.InterimUpdate.ValueString()
	}
	if !(plan.UseCircuitIDInNasPortID.IsNull() || plan.UseCircuitIDInNasPortID.IsUnknown()) {
		body["use-circuit-id-in-nas-port-id"] = client.FormatBool(plan.UseCircuitIDInNasPortID.ValueBool())
	}
	if !(plan.UseRADIUS.IsNull() || plan.UseRADIUS.IsUnknown()) {
		body["use-radius"] = client.FormatBool(plan.UseRADIUS.ValueBool())
	}
	obj, err := c.SetSingleton(ctx, "/ppp/aaa", body)
	if err != nil {
		diags.AddError("Upsert /ppp/aaa failed", err.Error())
		return
	}
	pPPAaaApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ppp/aaa", plan.Router))
}

func pPPAaaApply(ctx context.Context, obj client.Object, m *PPPAaaModel) {
	_ = ctx
	if v, ok := obj["accounting"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Accounting = types.BoolValue(b)
		} else {
			m.Accounting = types.BoolNull()
		}
	}
	if v, ok := obj["enable-ipv6-accounting"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.EnableIPV6Accounting = types.BoolValue(b)
		} else {
			m.EnableIPV6Accounting = types.BoolNull()
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
	if v, ok := obj["use-circuit-id-in-nas-port-id"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.UseCircuitIDInNasPortID = types.BoolValue(b)
		} else {
			m.UseCircuitIDInNasPortID = types.BoolNull()
		}
	}
	if v, ok := obj["use-radius"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.UseRADIUS = types.BoolValue(b)
		} else {
			m.UseRADIUS = types.BoolNull()
		}
	}
}
