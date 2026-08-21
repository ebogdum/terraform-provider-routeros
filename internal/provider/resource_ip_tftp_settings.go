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
	_ resource.Resource                = &IPTftpSettingsResource{}
	_ resource.ResourceWithImportState = &IPTftpSettingsResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type IPTftpSettingsResource struct {
	reg *client.Registry
}

type IPTftpSettingsModel struct {
	ID           types.String `tfsdk:"id"`
	MaxBlockSize types.String `tfsdk:"max_block_size"`
	Router       types.String `tfsdk:"router"`
}

func NewIPTftpSettingsResource() resource.Resource { return &IPTftpSettingsResource{} }

func (r *IPTftpSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_tftp_settings"
}

func (r *IPTftpSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPTftpSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/tftp/settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"max_block_size": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *IPTftpSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPTftpSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPTftpSettingsUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPTftpSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPTftpSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state IPTftpSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPTftpSettingsUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPTftpSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPTftpSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/tftp/settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/tftp/settings failed", err.Error())
		return
	}
	iPTftpSettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/tftp/settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPTftpSettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPTftpSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/tftp/settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ip/tftp/settings", types.StringValue(routerName))))...)
}

func iPTftpSettingsUpsert(ctx context.Context, reg *client.Registry, plan, state *IPTftpSettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.MaxBlockSize.IsNull() || plan.MaxBlockSize.IsUnknown()) && (state == nil || !plan.MaxBlockSize.Equal(state.MaxBlockSize)) {
		body["max-block-size"] = plan.MaxBlockSize.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/ip/tftp/settings", body)
	if err != nil {
		diags.AddError("Upsert /ip/tftp/settings failed", err.Error())
		return
	}
	iPTftpSettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/tftp/settings", plan.Router))
}

func iPTftpSettingsApply(ctx context.Context, obj client.Object, m *IPTftpSettingsModel) {
	_ = ctx
	if v, ok := obj["max-block-size"]; ok {
		_ = v
		if v != "" {
			m.MaxBlockSize = types.StringValue(v)
		} else {
			m.MaxBlockSize = types.StringNull()
		}
	}
}
