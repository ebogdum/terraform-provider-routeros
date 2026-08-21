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
	_ resource.Resource                = &SystemPackageLocalUpdateMirrorResource{}
	_ resource.ResourceWithImportState = &SystemPackageLocalUpdateMirrorResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemPackageLocalUpdateMirrorResource struct {
	reg *client.Registry
}

type SystemPackageLocalUpdateMirrorModel struct {
	ID              types.String `tfsdk:"id"`
	CheckInterval   types.String `tfsdk:"check_interval"`
	Enabled         types.Bool   `tfsdk:"enabled"`
	PrimaryServer   types.String `tfsdk:"primary_server"`
	SecondaryServer types.String `tfsdk:"secondary_server"`
	User            types.String `tfsdk:"user"`
	Router          types.String `tfsdk:"router"`
}

func NewSystemPackageLocalUpdateMirrorResource() resource.Resource {
	return &SystemPackageLocalUpdateMirrorResource{}
}

func (r *SystemPackageLocalUpdateMirrorResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_package_local_update_mirror"
}

func (r *SystemPackageLocalUpdateMirrorResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemPackageLocalUpdateMirrorResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/system/package/local-update/mirror`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"check_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `check-interval`.",
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `enabled`.",
			},
			"primary_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `primary-server`.",
			},
			"secondary_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `secondary-server`.",
			},
			"user": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `user`.",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *SystemPackageLocalUpdateMirrorResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemPackageLocalUpdateMirrorModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemPackageLocalUpdateMirrorUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemPackageLocalUpdateMirrorResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SystemPackageLocalUpdateMirrorModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state SystemPackageLocalUpdateMirrorModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemPackageLocalUpdateMirrorUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemPackageLocalUpdateMirrorResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemPackageLocalUpdateMirrorModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/system/package/local-update/mirror")
	if err != nil {
		resp.Diagnostics.AddError("Read /system/package/local-update/mirror failed", err.Error())
		return
	}
	systemPackageLocalUpdateMirrorApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/system/package/local-update/mirror", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemPackageLocalUpdateMirrorResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *SystemPackageLocalUpdateMirrorResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/system/package/local-update/mirror" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/system/package/local-update/mirror", types.StringValue(routerName))))...)
}

func systemPackageLocalUpdateMirrorUpsert(ctx context.Context, reg *client.Registry, plan, state *SystemPackageLocalUpdateMirrorModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.CheckInterval.IsNull() || plan.CheckInterval.IsUnknown()) && (state == nil || !plan.CheckInterval.Equal(state.CheckInterval)) {
		body["check-interval"] = plan.CheckInterval.ValueString()
	}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	if !(plan.PrimaryServer.IsNull() || plan.PrimaryServer.IsUnknown()) && (state == nil || !plan.PrimaryServer.Equal(state.PrimaryServer)) {
		body["primary-server"] = plan.PrimaryServer.ValueString()
	}
	if !(plan.SecondaryServer.IsNull() || plan.SecondaryServer.IsUnknown()) && (state == nil || !plan.SecondaryServer.Equal(state.SecondaryServer)) {
		body["secondary-server"] = plan.SecondaryServer.ValueString()
	}
	if !(plan.User.IsNull() || plan.User.IsUnknown()) && (state == nil || !plan.User.Equal(state.User)) {
		body["user"] = plan.User.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/system/package/local-update/mirror", body)
	if err != nil {
		diags.AddError("Upsert /system/package/local-update/mirror failed", err.Error())
		return
	}
	systemPackageLocalUpdateMirrorApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/system/package/local-update/mirror", plan.Router))
}

func systemPackageLocalUpdateMirrorApply(ctx context.Context, obj client.Object, m *SystemPackageLocalUpdateMirrorModel) {
	_ = ctx
	if v, ok := obj["check-interval"]; ok && v != "" {
		m.CheckInterval = types.StringValue(v)
	} else {
		m.CheckInterval = types.StringNull()
	}
	if v, ok := obj["enabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Enabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Enabled = types.BoolValue(true)
		} else {
			m.Enabled = types.BoolNull()
		}
	} else {
		m.Enabled = types.BoolNull()
	}
	if v, ok := obj["primary-server"]; ok && v != "" {
		m.PrimaryServer = types.StringValue(v)
	} else {
		m.PrimaryServer = types.StringNull()
	}
	if v, ok := obj["secondary-server"]; ok && v != "" {
		m.SecondaryServer = types.StringValue(v)
	} else {
		m.SecondaryServer = types.StringNull()
	}
	if v, ok := obj["user"]; ok && v != "" {
		m.User = types.StringValue(v)
	} else {
		m.User = types.StringNull()
	}
}
