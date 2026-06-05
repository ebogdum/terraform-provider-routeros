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
	_ resource.Resource                = &SystemPackageUpdateResource{}
	_ resource.ResourceWithImportState = &SystemPackageUpdateResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type SystemPackageUpdateResource struct {
	reg *client.Registry
}

type SystemPackageUpdateModel struct {
	ID               types.String `tfsdk:"id"`
	Channel          types.String `tfsdk:"channel"`
	InstalledVersion types.String `tfsdk:"installed_version"`
	Router           types.String `tfsdk:"router"`
}

func NewSystemPackageUpdateResource() resource.Resource { return &SystemPackageUpdateResource{} }

func (r *SystemPackageUpdateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_package_update"
}

func (r *SystemPackageUpdateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemPackageUpdateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/system/package/update`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"channel": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"installed_version": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemPackageUpdateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemPackageUpdateModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemPackageUpdateUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemPackageUpdateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SystemPackageUpdateModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemPackageUpdateUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemPackageUpdateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemPackageUpdateModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/system/package/update")
	if err != nil {
		resp.Diagnostics.AddError("Read /system/package/update failed", err.Error())
		return
	}
	systemPackageUpdateApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/system/package/update", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemPackageUpdateResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *SystemPackageUpdateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/system/package/update" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/system/package/update", types.StringValue(routerName))))...)
}

func systemPackageUpdateUpsert(ctx context.Context, reg *client.Registry, plan *SystemPackageUpdateModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Channel.IsNull() || plan.Channel.IsUnknown()) {
		body["channel"] = plan.Channel.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/system/package/update", body)
	if err != nil {
		diags.AddError("Upsert /system/package/update failed", err.Error())
		return
	}
	systemPackageUpdateApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/system/package/update", plan.Router))
}

func systemPackageUpdateApply(ctx context.Context, obj client.Object, m *SystemPackageUpdateModel) {
	_ = ctx
	if v, ok := obj["channel"]; ok {
		_ = v
		if v != "" {
			m.Channel = types.StringValue(v)
		} else {
			m.Channel = types.StringNull()
		}
	}
	if v, ok := obj["installed-version"]; ok {
		_ = v
		if v != "" {
			m.InstalledVersion = types.StringValue(v)
		} else {
			m.InstalledVersion = types.StringNull()
		}
	}
}
