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
	_ resource.Resource                = &MPLSSettingsResource{}
	_ resource.ResourceWithImportState = &MPLSSettingsResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type MPLSSettingsResource struct {
	reg *client.Registry
}

type MPLSSettingsModel struct {
	ID                  types.String `tfsdk:"id"`
	AllowFastPath       types.Bool   `tfsdk:"allow_fast_path"`
	DynamicLabelRange   types.String `tfsdk:"dynamic_label_range"`
	MPLSFastPathBytes   types.Int64  `tfsdk:"mpls_fast_path_bytes"`
	MPLSFastPathPackets types.Int64  `tfsdk:"mpls_fast_path_packets"`
	PropagateTtl        types.Bool   `tfsdk:"propagate_ttl"`
	Router              types.String `tfsdk:"router"`
}

func NewMPLSSettingsResource() resource.Resource { return &MPLSSettingsResource{} }

func (r *MPLSSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mpls_settings"
}

func (r *MPLSSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *MPLSSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/mpls/settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"allow_fast_path": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"dynamic_label_range": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"mpls_fast_path_bytes": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"mpls_fast_path_packets": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"propagate_ttl": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *MPLSSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan MPLSSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	mPLSSettingsUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MPLSSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan MPLSSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state MPLSSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	mPLSSettingsUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MPLSSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MPLSSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/mpls/settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /mpls/settings failed", err.Error())
		return
	}
	mPLSSettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/mpls/settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *MPLSSettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *MPLSSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/mpls/settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/mpls/settings", types.StringValue(routerName))))...)
}

func mPLSSettingsUpsert(ctx context.Context, reg *client.Registry, plan, state *MPLSSettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AllowFastPath.IsNull() || plan.AllowFastPath.IsUnknown()) && (state == nil || !plan.AllowFastPath.Equal(state.AllowFastPath)) {
		body["allow-fast-path"] = client.FormatBool(plan.AllowFastPath.ValueBool())
	}
	if !(plan.DynamicLabelRange.IsNull() || plan.DynamicLabelRange.IsUnknown()) && (state == nil || !plan.DynamicLabelRange.Equal(state.DynamicLabelRange)) {
		body["dynamic-label-range"] = plan.DynamicLabelRange.ValueString()
	}
	if !(plan.PropagateTtl.IsNull() || plan.PropagateTtl.IsUnknown()) && (state == nil || !plan.PropagateTtl.Equal(state.PropagateTtl)) {
		body["propagate-ttl"] = client.FormatBool(plan.PropagateTtl.ValueBool())
	}
	obj, err := c.SetSingleton(ctx, "/mpls/settings", body)
	if err != nil {
		diags.AddError("Upsert /mpls/settings failed", err.Error())
		return
	}
	mPLSSettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/mpls/settings", plan.Router))
}

func mPLSSettingsApply(ctx context.Context, obj client.Object, m *MPLSSettingsModel) {
	_ = ctx
	if v, ok := obj["allow-fast-path"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AllowFastPath = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.AllowFastPath = types.BoolValue(true)
		} else {
			m.AllowFastPath = types.BoolNull()
		}
	}
	if v, ok := obj["dynamic-label-range"]; ok {
		_ = v
		if v != "" {
			m.DynamicLabelRange = types.StringValue(v)
		} else {
			m.DynamicLabelRange = types.StringNull()
		}
	}
	if v, ok := obj["mpls-fast-path-bytes"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MPLSFastPathBytes = types.Int64Value(n)
		} else {
			m.MPLSFastPathBytes = types.Int64Null()
		}
	}
	if v, ok := obj["mpls-fast-path-packets"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MPLSFastPathPackets = types.Int64Value(n)
		} else {
			m.MPLSFastPathPackets = types.Int64Null()
		}
	}
	if v, ok := obj["propagate-ttl"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.PropagateTtl = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.PropagateTtl = types.BoolValue(true)
		} else {
			m.PropagateTtl = types.BoolNull()
		}
	}
}
