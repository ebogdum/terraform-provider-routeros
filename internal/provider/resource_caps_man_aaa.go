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
	_ resource.Resource                = &CapsManAaaResource{}
	_ resource.ResourceWithImportState = &CapsManAaaResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type CapsManAaaResource struct {
	reg *client.Registry
}

type CapsManAaaModel struct {
	ID            types.String `tfsdk:"id"`
	CalledFormat  types.String `tfsdk:"called_format"`
	InterimUpdate types.String `tfsdk:"interim_update"`
	MACCaching    types.String `tfsdk:"mac_caching"`
	MACFormat     types.String `tfsdk:"mac_format"`
	MACMode       types.String `tfsdk:"mac_mode"`
	Router        types.String `tfsdk:"router"`
}

func NewCapsManAaaResource() resource.Resource { return &CapsManAaaResource{} }

func (r *CapsManAaaResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_caps_man_aaa"
}

func (r *CapsManAaaResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *CapsManAaaResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/caps-man/aaa`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"called_format": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"interim_update": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"mac_caching": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"mac_format": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"mac_mode": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *CapsManAaaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CapsManAaaModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	capsManAaaUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CapsManAaaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan CapsManAaaModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state CapsManAaaModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	capsManAaaUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CapsManAaaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CapsManAaaModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/caps-man/aaa")
	if err != nil {
		resp.Diagnostics.AddError("Read /caps-man/aaa failed", err.Error())
		return
	}
	capsManAaaApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/caps-man/aaa", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CapsManAaaResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *CapsManAaaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/caps-man/aaa" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/caps-man/aaa", types.StringValue(routerName))))...)
}

func capsManAaaUpsert(ctx context.Context, reg *client.Registry, plan, state *CapsManAaaModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.CalledFormat.IsNull() || plan.CalledFormat.IsUnknown()) && (state == nil || !plan.CalledFormat.Equal(state.CalledFormat)) {
		body["called-format"] = plan.CalledFormat.ValueString()
	}
	if !(plan.InterimUpdate.IsNull() || plan.InterimUpdate.IsUnknown()) && (state == nil || !plan.InterimUpdate.Equal(state.InterimUpdate)) {
		body["interim-update"] = plan.InterimUpdate.ValueString()
	}
	if !(plan.MACCaching.IsNull() || plan.MACCaching.IsUnknown()) && (state == nil || !plan.MACCaching.Equal(state.MACCaching)) {
		body["mac-caching"] = plan.MACCaching.ValueString()
	}
	if !(plan.MACFormat.IsNull() || plan.MACFormat.IsUnknown()) && (state == nil || !plan.MACFormat.Equal(state.MACFormat)) {
		body["mac-format"] = plan.MACFormat.ValueString()
	}
	if !(plan.MACMode.IsNull() || plan.MACMode.IsUnknown()) && (state == nil || !plan.MACMode.Equal(state.MACMode)) {
		body["mac-mode"] = plan.MACMode.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/caps-man/aaa", body)
	if err != nil {
		diags.AddError("Upsert /caps-man/aaa failed", err.Error())
		return
	}
	capsManAaaApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/caps-man/aaa", plan.Router))
}

func capsManAaaApply(ctx context.Context, obj client.Object, m *CapsManAaaModel) {
	_ = ctx
	if v, ok := obj["called-format"]; ok {
		_ = v
		if v != "" {
			m.CalledFormat = types.StringValue(v)
		} else {
			m.CalledFormat = types.StringNull()
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
	if v, ok := obj["mac-caching"]; ok {
		_ = v
		if v != "" {
			m.MACCaching = types.StringValue(v)
		} else {
			m.MACCaching = types.StringNull()
		}
	}
	if v, ok := obj["mac-format"]; ok {
		_ = v
		if v != "" {
			m.MACFormat = types.StringValue(v)
		} else {
			m.MACFormat = types.StringNull()
		}
	}
	if v, ok := obj["mac-mode"]; ok {
		_ = v
		if v != "" {
			m.MACMode = types.StringValue(v)
		} else {
			m.MACMode = types.StringNull()
		}
	}
}
