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
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
	"github.com/ebogdum/terraform-provider-routeros/internal/schemautil"
)

var (
	_ resource.Resource                = &IPDHCPServerConfigResource{}
	_ resource.ResourceWithImportState = &IPDHCPServerConfigResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type IPDHCPServerConfigResource struct {
	reg *client.Registry
}

type IPDHCPServerConfigModel struct {
	ID              types.String `tfsdk:"id"`
	Accounting      types.Bool   `tfsdk:"accounting"`
	InterimUpdate   types.String `tfsdk:"interim_update"`
	RADIUSPassword  types.String `tfsdk:"radius_password"`
	StoreLeasesDisk types.String `tfsdk:"store_leases_disk"`
	Router          types.String `tfsdk:"router"`
}

func NewIPDHCPServerConfigResource() resource.Resource { return &IPDHCPServerConfigResource{} }

func (r *IPDHCPServerConfigResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_dhcp_server_config"
}

func (r *IPDHCPServerConfigResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPDHCPServerConfigResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/dhcp-server/config`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"accounting": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"interim_update": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"radius_password": schema.StringAttribute{Optional: true, Computed: true, Sensitive: true,
				Description: "",
			},
			"store_leases_disk": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPDHCPServerConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPDHCPServerConfigModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPDHCPServerConfigUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPServerConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPDHCPServerConfigModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state IPDHCPServerConfigModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPDHCPServerConfigUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPServerConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPDHCPServerConfigModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/dhcp-server/config")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/dhcp-server/config failed", err.Error())
		return
	}
	iPDHCPServerConfigApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/dhcp-server/config", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPDHCPServerConfigResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPDHCPServerConfigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/dhcp-server/config" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ip/dhcp-server/config", types.StringValue(routerName))))...)
}

func iPDHCPServerConfigUpsert(ctx context.Context, reg *client.Registry, plan, state *IPDHCPServerConfigModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Accounting.IsNull() || plan.Accounting.IsUnknown()) && (state == nil || !plan.Accounting.Equal(state.Accounting)) {
		body["accounting"] = client.FormatBool(plan.Accounting.ValueBool())
	}
	if !(plan.InterimUpdate.IsNull() || plan.InterimUpdate.IsUnknown()) && (state == nil || !plan.InterimUpdate.Equal(state.InterimUpdate)) {
		body["interim-update"] = plan.InterimUpdate.ValueString()
	}
	if !(plan.RADIUSPassword.IsNull() || plan.RADIUSPassword.IsUnknown()) && (state == nil || !plan.RADIUSPassword.Equal(state.RADIUSPassword)) {
		body["radius-password"] = plan.RADIUSPassword.ValueString()
	}
	if !(plan.StoreLeasesDisk.IsNull() || plan.StoreLeasesDisk.IsUnknown()) && (state == nil || !plan.StoreLeasesDisk.Equal(state.StoreLeasesDisk)) {
		body["store-leases-disk"] = plan.StoreLeasesDisk.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/ip/dhcp-server/config", body)
	if err != nil {
		diags.AddError("Upsert /ip/dhcp-server/config failed", err.Error())
		return
	}
	iPDHCPServerConfigApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/dhcp-server/config", plan.Router))
}

func iPDHCPServerConfigApply(ctx context.Context, obj client.Object, m *IPDHCPServerConfigModel) {
	_ = ctx
	if v, ok := obj["accounting"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Accounting = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Accounting = types.BoolValue(true)
		} else {
			m.Accounting = types.BoolNull()
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
	if v, ok := obj["radius-password"]; ok && v != "" {
		_ = v
		if v != "" {
			m.RADIUSPassword = types.StringValue(v)
		} else {
			m.RADIUSPassword = types.StringNull()
		}
	} else if m.RADIUSPassword.IsUnknown() {
		m.RADIUSPassword = types.StringNull()
	}
	if v, ok := obj["store-leases-disk"]; ok {
		_ = v
		if v != "" {
			m.StoreLeasesDisk = types.StringValue(v)
		} else {
			m.StoreLeasesDisk = types.StringNull()
		}
	}
}
