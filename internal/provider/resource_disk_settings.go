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
	_ resource.Resource                = &DiskSettingsResource{}
	_ resource.ResourceWithImportState = &DiskSettingsResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type DiskSettingsResource struct {
	reg *client.Registry
}

type DiskSettingsModel struct {
	ID                        types.String `tfsdk:"id"`
	AutoMediaInterface        types.String `tfsdk:"auto_media_interface"`
	AutoMediaSharing          types.Bool   `tfsdk:"auto_media_sharing"`
	AutoSmbSharing            types.Bool   `tfsdk:"auto_smb_sharing"`
	AutoSmbUser               types.String `tfsdk:"auto_smb_user"`
	DefaultMountPointTemplate types.String `tfsdk:"default_mount_point_template"`
	Router                    types.String `tfsdk:"router"`
}

func NewDiskSettingsResource() resource.Resource { return &DiskSettingsResource{} }

func (r *DiskSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_disk_settings"
}

func (r *DiskSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *DiskSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/disk/settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"auto_media_interface": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"auto_media_sharing": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"auto_smb_sharing": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"auto_smb_user": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"default_mount_point_template": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *DiskSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DiskSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	diskSettingsUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *DiskSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan DiskSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state DiskSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	diskSettingsUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *DiskSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DiskSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/disk/settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /disk/settings failed", err.Error())
		return
	}
	diskSettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/disk/settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *DiskSettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *DiskSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/disk/settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/disk/settings", types.StringValue(routerName))))...)
}

func diskSettingsUpsert(ctx context.Context, reg *client.Registry, plan, state *DiskSettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AutoMediaInterface.IsNull() || plan.AutoMediaInterface.IsUnknown()) && (state == nil || !plan.AutoMediaInterface.Equal(state.AutoMediaInterface)) {
		body["auto-media-interface"] = plan.AutoMediaInterface.ValueString()
	}
	if !(plan.AutoMediaSharing.IsNull() || plan.AutoMediaSharing.IsUnknown()) && (state == nil || !plan.AutoMediaSharing.Equal(state.AutoMediaSharing)) {
		body["auto-media-sharing"] = client.FormatBool(plan.AutoMediaSharing.ValueBool())
	}
	if !(plan.AutoSmbSharing.IsNull() || plan.AutoSmbSharing.IsUnknown()) && (state == nil || !plan.AutoSmbSharing.Equal(state.AutoSmbSharing)) {
		body["auto-smb-sharing"] = client.FormatBool(plan.AutoSmbSharing.ValueBool())
	}
	if !(plan.AutoSmbUser.IsNull() || plan.AutoSmbUser.IsUnknown()) && (state == nil || !plan.AutoSmbUser.Equal(state.AutoSmbUser)) {
		body["auto-smb-user"] = plan.AutoSmbUser.ValueString()
	}
	if !(plan.DefaultMountPointTemplate.IsNull() || plan.DefaultMountPointTemplate.IsUnknown()) && (state == nil || !plan.DefaultMountPointTemplate.Equal(state.DefaultMountPointTemplate)) {
		body["default-mount-point-template"] = plan.DefaultMountPointTemplate.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/disk/settings", body)
	if err != nil {
		diags.AddError("Upsert /disk/settings failed", err.Error())
		return
	}
	diskSettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/disk/settings", plan.Router))
}

func diskSettingsApply(ctx context.Context, obj client.Object, m *DiskSettingsModel) {
	_ = ctx
	if v, ok := obj["auto-media-interface"]; ok {
		_ = v
		if v != "" {
			m.AutoMediaInterface = types.StringValue(v)
		} else {
			m.AutoMediaInterface = types.StringNull()
		}
	}
	if v, ok := obj["auto-media-sharing"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AutoMediaSharing = types.BoolValue(b)
		} else {
			m.AutoMediaSharing = types.BoolNull()
		}
	}
	if v, ok := obj["auto-smb-sharing"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AutoSmbSharing = types.BoolValue(b)
		} else {
			m.AutoSmbSharing = types.BoolNull()
		}
	}
	if v, ok := obj["auto-smb-user"]; ok {
		_ = v
		if v != "" {
			m.AutoSmbUser = types.StringValue(v)
		} else {
			m.AutoSmbUser = types.StringNull()
		}
	}
	if v, ok := obj["default-mount-point-template"]; ok {
		_ = v
		if v != "" {
			m.DefaultMountPointTemplate = types.StringValue(v)
		} else {
			m.DefaultMountPointTemplate = types.StringNull()
		}
	}
}
