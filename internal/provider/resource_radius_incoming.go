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
	_ resource.Resource                = &RADIUSIncomingResource{}
	_ resource.ResourceWithImportState = &RADIUSIncomingResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type RADIUSIncomingResource struct {
	reg *client.Registry
}

type RADIUSIncomingModel struct {
	ID     types.String `tfsdk:"id"`
	Accept types.Bool   `tfsdk:"accept"`
	Port   types.Int64  `tfsdk:"port"`
	Vrf    types.String `tfsdk:"vrf"`
	Router types.String `tfsdk:"router"`
}

func NewRADIUSIncomingResource() resource.Resource { return &RADIUSIncomingResource{} }

func (r *RADIUSIncomingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_radius_incoming"
}

func (r *RADIUSIncomingResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RADIUSIncomingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/radius/incoming`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"accept": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"port": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"vrf": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *RADIUSIncomingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RADIUSIncomingModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	rADIUSIncomingUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RADIUSIncomingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan RADIUSIncomingModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state RADIUSIncomingModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	rADIUSIncomingUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RADIUSIncomingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RADIUSIncomingModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/radius/incoming")
	if err != nil {
		resp.Diagnostics.AddError("Read /radius/incoming failed", err.Error())
		return
	}
	rADIUSIncomingApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/radius/incoming", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RADIUSIncomingResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *RADIUSIncomingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/radius/incoming" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/radius/incoming", types.StringValue(routerName))))...)
}

func rADIUSIncomingUpsert(ctx context.Context, reg *client.Registry, plan, state *RADIUSIncomingModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Accept.IsNull() || plan.Accept.IsUnknown()) && (state == nil || !plan.Accept.Equal(state.Accept)) {
		body["accept"] = client.FormatBool(plan.Accept.ValueBool())
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) && (state == nil || !plan.Port.Equal(state.Port)) {
		body["port"] = client.FormatInt64(plan.Port.ValueInt64())
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) && (state == nil || !plan.Vrf.Equal(state.Vrf)) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/radius/incoming", body)
	if err != nil {
		diags.AddError("Upsert /radius/incoming failed", err.Error())
		return
	}
	rADIUSIncomingApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/radius/incoming", plan.Router))
}

func rADIUSIncomingApply(ctx context.Context, obj client.Object, m *RADIUSIncomingModel) {
	_ = ctx
	if v, ok := obj["accept"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Accept = types.BoolValue(b)
		} else {
			m.Accept = types.BoolNull()
		}
	}
	if v, ok := obj["port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Port = types.Int64Value(n)
		} else {
			m.Port = types.Int64Null()
		}
	}
	if v, ok := obj["vrf"]; ok {
		_ = v
		if v != "" {
			m.Vrf = types.StringValue(v)
		} else {
			m.Vrf = types.StringNull()
		}
	}
}
