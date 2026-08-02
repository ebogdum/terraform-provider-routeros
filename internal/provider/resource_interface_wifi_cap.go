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
	_ resource.Resource                = &InterfaceWifiCapResource{}
	_ resource.ResourceWithImportState = &InterfaceWifiCapResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type InterfaceWifiCapResource struct {
	reg *client.Registry
}

type InterfaceWifiCapModel struct {
	ID                            types.String `tfsdk:"id"`
	SlavesDatapath                types.String `tfsdk:"slaves_datapath"`
	MldStatic                     types.String `tfsdk:"mld_static"`
	MldDatapath                   types.String `tfsdk:"mld_datapath"`
	Enabled                       types.Bool   `tfsdk:"enabled"`
	CapsManAddresses              types.String `tfsdk:"caps_man_addresses"`
	CapsManNames                  types.String `tfsdk:"caps_man_names"`
	CapsManCertificateCommonNames types.String `tfsdk:"caps_man_certificate_common_names"`
	DiscoveryInterfaces           types.String `tfsdk:"discovery_interfaces"`
	Certificate                   types.String `tfsdk:"certificate"`
	LockToCapsMan                 types.Bool   `tfsdk:"lock_to_caps_man"`
	SlavesStatic                  types.Bool   `tfsdk:"slaves_static"`
	Router                        types.String `tfsdk:"router"`
}

func NewInterfaceWifiCapResource() resource.Resource { return &InterfaceWifiCapResource{} }

func (r *InterfaceWifiCapResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_wifi_cap"
}

func (r *InterfaceWifiCapResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceWifiCapResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/wifi/cap` -- how this device finds and binds to its " +
			"CAPsMAN controller.\n\n" +
			"`enabled` alone is not enough to rebuild a CAP: without `caps_man_addresses` (or a discovery " +
			"interface that reaches the controller) the radios come up unmanaged and the provisioned SSIDs " +
			"never appear.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"slaves_datapath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `slaves-datapath`.",
			},
			"mld_static": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mld-static`.",
			},
			"mld_datapath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mld-datapath`.",
			},
			"enabled": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "Whether CAP mode is enabled on this device.",
			},
			"caps_man_addresses": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Comma-separated list of CAPsMAN controller addresses to connect to. " +
					"Without this the CAP relies on discovery over `discovery_interfaces`.",
			},
			"caps_man_names": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Comma-separated list of CAPsMAN controller names to accept.",
			},
			"caps_man_certificate_common_names": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Comma-separated list of controller certificate CommonNames to accept.",
			},
			"discovery_interfaces": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Comma-separated list of interfaces used to discover a CAPsMAN controller.",
			},
			"certificate": schema.StringAttribute{Optional: true, Computed: true,
				Description: "Certificate used towards the controller: a certificate name, `none`, or `request`.",
			},
			"lock_to_caps_man": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "Lock the CAP to the first controller it successfully connects to.",
			},
			"slaves_static": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "Create static rather than dynamic interfaces for provisioned radios.",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceWifiCapResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceWifiCapModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	interfaceWifiCapUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiCapResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan InterfaceWifiCapModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state InterfaceWifiCapModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	interfaceWifiCapUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiCapResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceWifiCapModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/interface/wifi/cap")
	if err != nil {
		resp.Diagnostics.AddError("Read /interface/wifi/cap failed", err.Error())
		return
	}
	interfaceWifiCapApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/interface/wifi/cap", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceWifiCapResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *InterfaceWifiCapResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/interface/wifi/cap" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/interface/wifi/cap", types.StringValue(routerName))))...)
}

func interfaceWifiCapUpsert(ctx context.Context, reg *client.Registry, plan, state *InterfaceWifiCapModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	if !(plan.CapsManAddresses.IsNull() || plan.CapsManAddresses.IsUnknown()) && (state == nil || !plan.CapsManAddresses.Equal(state.CapsManAddresses)) {
		body["caps-man-addresses"] = plan.CapsManAddresses.ValueString()
	}
	if !(plan.CapsManNames.IsNull() || plan.CapsManNames.IsUnknown()) && (state == nil || !plan.CapsManNames.Equal(state.CapsManNames)) {
		body["caps-man-names"] = plan.CapsManNames.ValueString()
	}
	if !(plan.CapsManCertificateCommonNames.IsNull() || plan.CapsManCertificateCommonNames.IsUnknown()) && (state == nil || !plan.CapsManCertificateCommonNames.Equal(state.CapsManCertificateCommonNames)) {
		body["caps-man-certificate-common-names"] = plan.CapsManCertificateCommonNames.ValueString()
	}
	if !(plan.DiscoveryInterfaces.IsNull() || plan.DiscoveryInterfaces.IsUnknown()) && (state == nil || !plan.DiscoveryInterfaces.Equal(state.DiscoveryInterfaces)) {
		body["discovery-interfaces"] = plan.DiscoveryInterfaces.ValueString()
	}
	if !(plan.Certificate.IsNull() || plan.Certificate.IsUnknown()) && (state == nil || !plan.Certificate.Equal(state.Certificate)) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !(plan.LockToCapsMan.IsNull() || plan.LockToCapsMan.IsUnknown()) && (state == nil || !plan.LockToCapsMan.Equal(state.LockToCapsMan)) {
		body["lock-to-caps-man"] = client.FormatBool(plan.LockToCapsMan.ValueBool())
	}
	if !(plan.SlavesStatic.IsNull() || plan.SlavesStatic.IsUnknown()) && (state == nil || !plan.SlavesStatic.Equal(state.SlavesStatic)) {
		body["slaves-static"] = client.FormatBool(plan.SlavesStatic.ValueBool())
	}
	if !(plan.MldDatapath.IsNull() || plan.MldDatapath.IsUnknown()) && (state == nil || !plan.MldDatapath.Equal(state.MldDatapath)) {
		body["mld-datapath"] = plan.MldDatapath.ValueString()
	}
	if !(plan.MldStatic.IsNull() || plan.MldStatic.IsUnknown()) && (state == nil || !plan.MldStatic.Equal(state.MldStatic)) {
		body["mld-static"] = plan.MldStatic.ValueString()
	}
	if !(plan.SlavesDatapath.IsNull() || plan.SlavesDatapath.IsUnknown()) && (state == nil || !plan.SlavesDatapath.Equal(state.SlavesDatapath)) {
		body["slaves-datapath"] = plan.SlavesDatapath.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/interface/wifi/cap", body)
	if err != nil {
		diags.AddError("Upsert /interface/wifi/cap failed", err.Error())
		return
	}
	interfaceWifiCapApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/interface/wifi/cap", plan.Router))
}

func interfaceWifiCapApply(ctx context.Context, obj client.Object, m *InterfaceWifiCapModel) {
	_ = ctx
	if v, ok := obj["slaves-datapath"]; ok && v != "" {
		m.SlavesDatapath = types.StringValue(v)
	} else {
		m.SlavesDatapath = types.StringNull()
	}
	if v, ok := obj["mld-static"]; ok && v != "" {
		m.MldStatic = types.StringValue(v)
	} else {
		m.MldStatic = types.StringNull()
	}
	if v, ok := obj["mld-datapath"]; ok && v != "" {
		m.MldDatapath = types.StringValue(v)
	} else {
		m.MldDatapath = types.StringNull()
	}
	if v, ok := obj["enabled"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Enabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Enabled = types.BoolValue(true)
		} else {
			m.Enabled = types.BoolNull()
		}
	}
	applyStr := func(key string, dst *types.String) {
		if v, ok := obj[key]; ok && v != "" {
			*dst = types.StringValue(v)
		} else {
			*dst = types.StringNull()
		}
	}
	applyBool := func(key string, dst *types.Bool) {
		if v, ok := obj[key]; ok {
			if b, err := client.ParseBool(v); err == nil {
				*dst = types.BoolValue(b)
				return
			}
		}
		*dst = types.BoolNull()
	}
	applyStr("caps-man-addresses", &m.CapsManAddresses)
	applyStr("caps-man-names", &m.CapsManNames)
	applyStr("caps-man-certificate-common-names", &m.CapsManCertificateCommonNames)
	applyStr("discovery-interfaces", &m.DiscoveryInterfaces)
	applyStr("certificate", &m.Certificate)
	applyBool("lock-to-caps-man", &m.LockToCapsMan)
	applyBool("slaves-static", &m.SlavesStatic)
}
