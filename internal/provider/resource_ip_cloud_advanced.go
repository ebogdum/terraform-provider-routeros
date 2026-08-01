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
	_ resource.Resource                = &IPCloudAdvancedResource{}
	_ resource.ResourceWithImportState = &IPCloudAdvancedResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPCloudAdvancedResource struct {
	reg *client.Registry
}

type IPCloudAdvancedModel struct {
	ID              types.String `tfsdk:"id"`
	UseLocalAddress types.Bool   `tfsdk:"use_local_address"`
	Router          types.String `tfsdk:"router"`
}

func NewIPCloudAdvancedResource() resource.Resource { return &IPCloudAdvancedResource{} }

func (r *IPCloudAdvancedResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_cloud_advanced"
}

func (r *IPCloudAdvancedResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPCloudAdvancedResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/cloud/advanced`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"use_local_address": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `use-local-address`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPCloudAdvancedResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPCloudAdvancedModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPCloudAdvancedUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPCloudAdvancedResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPCloudAdvancedModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state IPCloudAdvancedModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPCloudAdvancedUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPCloudAdvancedResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPCloudAdvancedModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/cloud/advanced")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/cloud/advanced failed", err.Error())
		return
	}
	iPCloudAdvancedApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/cloud/advanced", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPCloudAdvancedResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPCloudAdvancedResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/cloud/advanced" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ip/cloud/advanced", types.StringValue(routerName))))...)
}

func iPCloudAdvancedUpsert(ctx context.Context, reg *client.Registry, plan, state *IPCloudAdvancedModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.UseLocalAddress.IsNull() || plan.UseLocalAddress.IsUnknown()) && (state == nil || !plan.UseLocalAddress.Equal(state.UseLocalAddress)) {
		body["use-local-address"] = client.FormatBool(plan.UseLocalAddress.ValueBool())
	}
	obj, err := c.SetSingleton(ctx, "/ip/cloud/advanced", body)
	if err != nil {
		diags.AddError("Upsert /ip/cloud/advanced failed", err.Error())
		return
	}
	iPCloudAdvancedApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/cloud/advanced", plan.Router))
}

func iPCloudAdvancedApply(ctx context.Context, obj client.Object, m *IPCloudAdvancedModel) {
	_ = ctx
	if v, ok := obj["use-local-address"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.UseLocalAddress = types.BoolValue(b)
		} else {
			m.UseLocalAddress = types.BoolNull()
		}
	} else {
		m.UseLocalAddress = types.BoolNull()
	}
}
