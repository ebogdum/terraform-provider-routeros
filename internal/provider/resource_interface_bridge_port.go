package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
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
	_ resource.Resource                = &InterfaceBridgePortResource{}
	_ resource.ResourceWithImportState = &InterfaceBridgePortResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceBridgePortResource struct {
	reg *client.Registry
}

type InterfaceBridgePortModel struct {
	ID                    types.String    `tfsdk:"id"`
	TrustedDhcpv6         boolStringValue `tfsdk:"trusted_dhcpv6"`
	AutoIsolate           types.Bool      `tfsdk:"auto_isolate"`
	BpduGuard             types.Bool      `tfsdk:"bpdu_guard"`
	Bridge                types.String    `tfsdk:"bridge"`
	BroadcastFlood        types.Bool      `tfsdk:"broadcast_flood"`
	Comment               types.String    `tfsdk:"comment"`
	Disabled              types.Bool      `tfsdk:"disabled"`
	Dynamic               types.Bool      `tfsdk:"dynamic"`
	Edge                  types.String    `tfsdk:"edge"`
	FastLeave             types.Bool      `tfsdk:"fast_leave"`
	FrameTypes            types.String    `tfsdk:"frame_types"`
	HardwareOffload       types.Bool      `tfsdk:"hardware_offload"`
	Horizon               types.String    `tfsdk:"horizon"`
	Hw                    boolStringValue `tfsdk:"hw"`
	HwOffload             types.Bool      `tfsdk:"hw_offload"`
	HwOffloadGroup        types.String    `tfsdk:"hw_offload_group"`
	Inactive              types.Bool      `tfsdk:"inactive"`
	IngressFiltering      types.Bool      `tfsdk:"ingress_filtering"`
	Interface             types.String    `tfsdk:"interface"`
	InternalPathCost      types.String    `tfsdk:"internal_path_cost"`
	Learn                 types.String    `tfsdk:"learn"`
	MulticastRouter       types.String    `tfsdk:"multicast_router"`
	MvrpApplicantState    types.String    `tfsdk:"mvrp_applicant_state"`
	MvrpRegistrarState    types.String    `tfsdk:"mvrp_registrar_state"`
	Parent                types.Int64     `tfsdk:"parent"`
	PathCost              types.String    `tfsdk:"path_cost"`
	PointToPoint          types.String    `tfsdk:"point_to_point"`
	PortStatus            types.String    `tfsdk:"port_status"`
	Priority              types.Int64     `tfsdk:"priority"`
	Pvid                  types.Int64     `tfsdk:"pvid"`
	RestrictedRole        types.Bool      `tfsdk:"restricted_role"`
	RestrictedTcn         types.Bool      `tfsdk:"restricted_tcn"`
	Role                  types.String    `tfsdk:"role"`
	Status                types.String    `tfsdk:"status"`
	TagStacking           types.Bool      `tfsdk:"tag_stacking"`
	Trusted               types.Bool      `tfsdk:"trusted"`
	TrustedRa             types.Bool      `tfsdk:"trusted_ra"`
	UnknownMulticastFlood types.Bool      `tfsdk:"unknown_multicast_flood"`
	UnknownUnicastFlood   types.Bool      `tfsdk:"unknown_unicast_flood"`
	Router                types.String    `tfsdk:"router"`
}

func NewInterfaceBridgePortResource() resource.Resource { return &InterfaceBridgePortResource{} }

func (r *InterfaceBridgePortResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_bridge_port"
}

func (r *InterfaceBridgePortResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceBridgePortResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Auto-test requires a typed-reference precondition (e.g. an existing peer,\ninstance, bridge of the specific kind). The current acc-test generator's\ngeneric data.routeros_interface.all lookup can't satisfy these. Use this\nresource manually with explicit references to a precondition resource\nin your config.\n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"trusted_dhcpv6": schema.StringAttribute{
				CustomType:  boolStringType{},
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `trusted-dhcpv6`.",
			},
			"auto_isolate": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bpdu_guard": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"broadcast_flood": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"edge": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"auto", "yes", "no", "yes-discover", "no-discover"}...)},
			},
			"fast_leave": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"frame_types": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"admit-all", "admit-only-vlan-tagged", "admit-only-untagged-and-priority-tagged"}...)},
			},
			"hardware_offload": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"horizon": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Split-horizon group used to isolate ports. A number, or `none` (the default).",
			},
			"hw": schema.StringAttribute{
				CustomType:  boolStringType{},
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hw_offload": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"hw_offload_group": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"inactive": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"ingress_filtering": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"internal_path_cost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"learn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"auto", "no", "yes"}...)},
			},
			"multicast_router": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"disabled", "temporary-query", "permanent"}...)},
			},
			"mvrp_applicant_state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"normal-participant", "non-participant"}...)},
			},
			"mvrp_registrar_state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"normal", "fixed"}...)},
			},
			"parent": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"path_cost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"point_to_point": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"auto", "yes", "no"}...)},
			},
			"port_status": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "inactive", "active", "disabled"}...)},
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pvid": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"restricted_role": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"restricted_tcn": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"role": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tag_stacking": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"trusted": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"trusted_ra": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"unknown_multicast_flood": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"unknown_unicast_flood": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceBridgePortResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceBridgePortModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AutoIsolate.IsNull() || plan.AutoIsolate.IsUnknown()) {
		body["auto-isolate"] = client.FormatBool(plan.AutoIsolate.ValueBool())
	}
	if !(plan.BpduGuard.IsNull() || plan.BpduGuard.IsUnknown()) {
		body["bpdu-guard"] = client.FormatBool(plan.BpduGuard.ValueBool())
	}
	if !(plan.Bridge.IsNull() || plan.Bridge.IsUnknown()) {
		body["bridge"] = plan.Bridge.ValueString()
	}
	if !(plan.BroadcastFlood.IsNull() || plan.BroadcastFlood.IsUnknown()) {
		body["broadcast-flood"] = client.FormatBool(plan.BroadcastFlood.ValueBool())
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Edge.IsNull() || plan.Edge.IsUnknown()) {
		body["edge"] = plan.Edge.ValueString()
	}
	if !(plan.FastLeave.IsNull() || plan.FastLeave.IsUnknown()) {
		body["fast-leave"] = client.FormatBool(plan.FastLeave.ValueBool())
	}
	if !(plan.FrameTypes.IsNull() || plan.FrameTypes.IsUnknown()) {
		body["frame-types"] = plan.FrameTypes.ValueString()
	}
	if !(plan.Horizon.IsNull() || plan.Horizon.IsUnknown()) {
		body["horizon"] = plan.Horizon.ValueString()
	}
	if !(plan.Hw.IsNull() || plan.Hw.IsUnknown()) {
		body["hw"] = plan.Hw.ValueString()
	}
	if !(plan.IngressFiltering.IsNull() || plan.IngressFiltering.IsUnknown()) {
		body["ingress-filtering"] = client.FormatBool(plan.IngressFiltering.ValueBool())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.InternalPathCost.IsNull() || plan.InternalPathCost.IsUnknown()) {
		body["internal-path-cost"] = plan.InternalPathCost.ValueString()
	}
	if !(plan.Learn.IsNull() || plan.Learn.IsUnknown()) {
		body["learn"] = plan.Learn.ValueString()
	}
	if !(plan.MulticastRouter.IsNull() || plan.MulticastRouter.IsUnknown()) {
		body["multicast-router"] = plan.MulticastRouter.ValueString()
	}
	if !(plan.MvrpApplicantState.IsNull() || plan.MvrpApplicantState.IsUnknown()) {
		body["mvrp-applicant-state"] = plan.MvrpApplicantState.ValueString()
	}
	if !(plan.MvrpRegistrarState.IsNull() || plan.MvrpRegistrarState.IsUnknown()) {
		body["mvrp-registrar-state"] = plan.MvrpRegistrarState.ValueString()
	}
	if !(plan.PathCost.IsNull() || plan.PathCost.IsUnknown()) {
		body["path-cost"] = plan.PathCost.ValueString()
	}
	if !(plan.PointToPoint.IsNull() || plan.PointToPoint.IsUnknown()) {
		body["point-to-point"] = plan.PointToPoint.ValueString()
	}
	if !(plan.Priority.IsNull() || plan.Priority.IsUnknown()) {
		body["priority"] = client.FormatInt64(plan.Priority.ValueInt64())
	}
	if !(plan.Pvid.IsNull() || plan.Pvid.IsUnknown()) {
		body["pvid"] = client.FormatInt64(plan.Pvid.ValueInt64())
	}
	if !(plan.RestrictedRole.IsNull() || plan.RestrictedRole.IsUnknown()) {
		body["restricted-role"] = client.FormatBool(plan.RestrictedRole.ValueBool())
	}
	if !(plan.RestrictedTcn.IsNull() || plan.RestrictedTcn.IsUnknown()) {
		body["restricted-tcn"] = client.FormatBool(plan.RestrictedTcn.ValueBool())
	}
	if !(plan.TagStacking.IsNull() || plan.TagStacking.IsUnknown()) {
		body["tag-stacking"] = client.FormatBool(plan.TagStacking.ValueBool())
	}
	if !(plan.Trusted.IsNull() || plan.Trusted.IsUnknown()) {
		body["trusted"] = client.FormatBool(plan.Trusted.ValueBool())
	}
	if !(plan.TrustedRa.IsNull() || plan.TrustedRa.IsUnknown()) {
		body["trusted-ra"] = client.FormatBool(plan.TrustedRa.ValueBool())
	}
	if !(plan.UnknownMulticastFlood.IsNull() || plan.UnknownMulticastFlood.IsUnknown()) {
		body["unknown-multicast-flood"] = client.FormatBool(plan.UnknownMulticastFlood.ValueBool())
	}
	if !(plan.UnknownUnicastFlood.IsNull() || plan.UnknownUnicastFlood.IsUnknown()) {
		body["unknown-unicast-flood"] = client.FormatBool(plan.UnknownUnicastFlood.ValueBool())
	}
	if !(plan.TrustedDhcpv6.IsNull() || plan.TrustedDhcpv6.IsUnknown()) {
		body["trusted-dhcpv6"] = plan.TrustedDhcpv6.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/bridge/port", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/bridge/port failed", err.Error())
		return
	}
	interfaceBridgePortApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceBridgePortResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceBridgePortModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/bridge/port", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/bridge/port failed", err.Error())
		return
	}
	interfaceBridgePortApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceBridgePortResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceBridgePortModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !plan.AutoIsolate.Equal(state.AutoIsolate) && !plan.AutoIsolate.IsUnknown() {
		body["auto-isolate"] = client.FormatBool(plan.AutoIsolate.ValueBool())
	}
	if !plan.BpduGuard.Equal(state.BpduGuard) && !plan.BpduGuard.IsUnknown() {
		body["bpdu-guard"] = client.FormatBool(plan.BpduGuard.ValueBool())
	}
	if !plan.Bridge.Equal(state.Bridge) && !plan.Bridge.IsUnknown() {
		body["bridge"] = plan.Bridge.ValueString()
	}
	if !plan.BroadcastFlood.Equal(state.BroadcastFlood) && !plan.BroadcastFlood.IsUnknown() {
		body["broadcast-flood"] = client.FormatBool(plan.BroadcastFlood.ValueBool())
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Edge.Equal(state.Edge) && !plan.Edge.IsUnknown() {
		body["edge"] = plan.Edge.ValueString()
	}
	if !plan.FastLeave.Equal(state.FastLeave) && !plan.FastLeave.IsUnknown() {
		body["fast-leave"] = client.FormatBool(plan.FastLeave.ValueBool())
	}
	if !plan.FrameTypes.Equal(state.FrameTypes) && !plan.FrameTypes.IsUnknown() {
		body["frame-types"] = plan.FrameTypes.ValueString()
	}
	if !plan.Horizon.Equal(state.Horizon) && !plan.Horizon.IsUnknown() {
		body["horizon"] = plan.Horizon.ValueString()
	}
	if !plan.Hw.Equal(state.Hw) && !plan.Hw.IsUnknown() {
		body["hw"] = plan.Hw.ValueString()
	}
	if !plan.IngressFiltering.Equal(state.IngressFiltering) && !plan.IngressFiltering.IsUnknown() {
		body["ingress-filtering"] = client.FormatBool(plan.IngressFiltering.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.InternalPathCost.Equal(state.InternalPathCost) && !plan.InternalPathCost.IsUnknown() {
		body["internal-path-cost"] = plan.InternalPathCost.ValueString()
	}
	if !plan.Learn.Equal(state.Learn) && !plan.Learn.IsUnknown() {
		body["learn"] = plan.Learn.ValueString()
	}
	if !plan.MulticastRouter.Equal(state.MulticastRouter) && !plan.MulticastRouter.IsUnknown() {
		body["multicast-router"] = plan.MulticastRouter.ValueString()
	}
	if !plan.MvrpApplicantState.Equal(state.MvrpApplicantState) && !plan.MvrpApplicantState.IsUnknown() {
		body["mvrp-applicant-state"] = plan.MvrpApplicantState.ValueString()
	}
	if !plan.MvrpRegistrarState.Equal(state.MvrpRegistrarState) && !plan.MvrpRegistrarState.IsUnknown() {
		body["mvrp-registrar-state"] = plan.MvrpRegistrarState.ValueString()
	}
	if !plan.PathCost.Equal(state.PathCost) && !plan.PathCost.IsUnknown() {
		body["path-cost"] = plan.PathCost.ValueString()
	}
	if !plan.PointToPoint.Equal(state.PointToPoint) && !plan.PointToPoint.IsUnknown() {
		body["point-to-point"] = plan.PointToPoint.ValueString()
	}
	if !plan.Priority.Equal(state.Priority) && !plan.Priority.IsUnknown() {
		body["priority"] = client.FormatInt64(plan.Priority.ValueInt64())
	}
	if !plan.Pvid.Equal(state.Pvid) && !plan.Pvid.IsUnknown() {
		body["pvid"] = client.FormatInt64(plan.Pvid.ValueInt64())
	}
	if !plan.RestrictedRole.Equal(state.RestrictedRole) && !plan.RestrictedRole.IsUnknown() {
		body["restricted-role"] = client.FormatBool(plan.RestrictedRole.ValueBool())
	}
	if !plan.RestrictedTcn.Equal(state.RestrictedTcn) && !plan.RestrictedTcn.IsUnknown() {
		body["restricted-tcn"] = client.FormatBool(plan.RestrictedTcn.ValueBool())
	}
	if !plan.TagStacking.Equal(state.TagStacking) && !plan.TagStacking.IsUnknown() {
		body["tag-stacking"] = client.FormatBool(plan.TagStacking.ValueBool())
	}
	if !plan.Trusted.Equal(state.Trusted) && !plan.Trusted.IsUnknown() {
		body["trusted"] = client.FormatBool(plan.Trusted.ValueBool())
	}
	if !plan.TrustedRa.Equal(state.TrustedRa) && !plan.TrustedRa.IsUnknown() {
		body["trusted-ra"] = client.FormatBool(plan.TrustedRa.ValueBool())
	}
	if !plan.UnknownMulticastFlood.Equal(state.UnknownMulticastFlood) && !plan.UnknownMulticastFlood.IsUnknown() {
		body["unknown-multicast-flood"] = client.FormatBool(plan.UnknownMulticastFlood.ValueBool())
	}
	if !plan.UnknownUnicastFlood.Equal(state.UnknownUnicastFlood) && !plan.UnknownUnicastFlood.IsUnknown() {
		body["unknown-unicast-flood"] = client.FormatBool(plan.UnknownUnicastFlood.ValueBool())
	}
	if !plan.TrustedDhcpv6.Equal(state.TrustedDhcpv6) && !plan.TrustedDhcpv6.IsUnknown() {
		body["trusted-dhcpv6"] = plan.TrustedDhcpv6.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/bridge/port", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/bridge/port failed", err.Error())
			return
		}
		interfaceBridgePortApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceBridgePortResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceBridgePortModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/bridge/port", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/bridge/port failed", err.Error())
	}
}

func (r *InterfaceBridgePortResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import formats accepted:
	//   *<id>                            -> bare RouterOS .id on the default router
	//   <router>/*<id>                   -> .id on the named router
	//   <router>/<naturalkey>            -> resolved via List + filter
	//   <naturalkey>                     -> resolved on the default router
	routerName, id := parseImportID(r.reg, req.ID)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	if strings.HasPrefix(id, "*") {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(id))...)
		return
	}
	c := pickClient(r.reg, types.StringValue(routerName), &resp.Diagnostics)
	if c == nil {
		return
	}
	rows, err := interfaceBridgePortLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/bridge/port matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceBridgePortLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceBridgePortLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/bridge/port", id)
}

func interfaceBridgePortApply(ctx context.Context, obj client.Object, m *InterfaceBridgePortModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["trusted-dhcpv6"]; ok && v != "" {
		m.TrustedDhcpv6 = newBoolStringValue(v)
	}
	if v, ok := obj["auto-isolate"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.AutoIsolate = types.BoolValue(b)
		} else {
			m.AutoIsolate = types.BoolNull()
		}
	}
	if v, ok := obj["bpdu-guard"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.BpduGuard = types.BoolValue(b)
		} else {
			m.BpduGuard = types.BoolNull()
		}
	}
	if v, ok := obj["bridge"]; ok {
		if v != "" {
			m.Bridge = types.StringValue(v)
		} else {
			m.Bridge = types.StringNull()
		}
	}
	if v, ok := obj["broadcast-flood"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.BroadcastFlood = types.BoolValue(b)
		} else {
			m.BroadcastFlood = types.BoolNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["dynamic"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else {
			m.Dynamic = types.BoolNull()
		}
	}
	if v, ok := obj["edge"]; ok {
		if v != "" {
			m.Edge = types.StringValue(v)
		} else {
			m.Edge = types.StringNull()
		}
	}
	if v, ok := obj["fast-leave"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.FastLeave = types.BoolValue(b)
		} else {
			m.FastLeave = types.BoolNull()
		}
	}
	if v, ok := obj["frame-types"]; ok {
		if v != "" {
			m.FrameTypes = types.StringValue(v)
		} else {
			m.FrameTypes = types.StringNull()
		}
	}
	if v, ok := obj["hardware-offload"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.HardwareOffload = types.BoolValue(b)
		} else {
			m.HardwareOffload = types.BoolNull()
		}
	}
	if v, ok := obj["horizon"]; ok {
		if v != "" {
			m.Horizon = types.StringValue(v)
		} else {
			m.Horizon = types.StringNull()
		}
	}
	if v, ok := obj["hw"]; ok {
		_ = v
		if v != "" {
			m.Hw = newBoolStringValue(v)
		}
	}
	if v, ok := obj["hw-offload"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.HwOffload = types.BoolValue(b)
		} else {
			m.HwOffload = types.BoolNull()
		}
	}
	if v, ok := obj["hw-offload-group"]; ok {
		if v != "" {
			m.HwOffloadGroup = types.StringValue(v)
		} else {
			m.HwOffloadGroup = types.StringNull()
		}
	}
	if v, ok := obj["inactive"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Inactive = types.BoolValue(b)
		} else {
			m.Inactive = types.BoolNull()
		}
	}
	if v, ok := obj["ingress-filtering"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.IngressFiltering = types.BoolValue(b)
		} else {
			m.IngressFiltering = types.BoolNull()
		}
	}
	if v, ok := obj["interface"]; ok {
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	}
	if v, ok := obj["internal-path-cost"]; ok {
		if v != "" {
			m.InternalPathCost = types.StringValue(v)
		} else {
			m.InternalPathCost = types.StringNull()
		}
	}
	if v, ok := obj["learn"]; ok {
		if v != "" {
			m.Learn = types.StringValue(v)
		} else {
			m.Learn = types.StringNull()
		}
	}
	if v, ok := obj["multicast-router"]; ok {
		if v != "" {
			m.MulticastRouter = types.StringValue(v)
		} else {
			m.MulticastRouter = types.StringNull()
		}
	}
	if v, ok := obj["mvrp-applicant-state"]; ok {
		if v != "" {
			m.MvrpApplicantState = types.StringValue(v)
		} else {
			m.MvrpApplicantState = types.StringNull()
		}
	}
	if v, ok := obj["mvrp-registrar-state"]; ok {
		if v != "" {
			m.MvrpRegistrarState = types.StringValue(v)
		} else {
			m.MvrpRegistrarState = types.StringNull()
		}
	}
	if v, ok := obj["parent"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Parent = types.Int64Value(n)
		} else {
			m.Parent = types.Int64Null()
		}
	} else {
		m.Parent = types.Int64Null()
	}
	if v, ok := obj["path-cost"]; ok {
		if v != "" {
			m.PathCost = types.StringValue(v)
		} else {
			m.PathCost = types.StringNull()
		}
	}
	if v, ok := obj["point-to-point"]; ok {
		if v != "" {
			m.PointToPoint = types.StringValue(v)
		} else {
			m.PointToPoint = types.StringNull()
		}
	}
	if v, ok := obj["port-status"]; ok {
		if v != "" {
			m.PortStatus = types.StringValue(v)
		} else {
			m.PortStatus = types.StringNull()
		}
	}
	if v, ok := obj["priority"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Priority = types.Int64Value(n)
		} else {
			m.Priority = types.Int64Null()
		}
	} else {
		m.Priority = types.Int64Null()
	}
	if v, ok := obj["pvid"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Pvid = types.Int64Value(n)
		} else {
			m.Pvid = types.Int64Null()
		}
	} else {
		m.Pvid = types.Int64Null()
	}
	if v, ok := obj["restricted-role"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.RestrictedRole = types.BoolValue(b)
		} else {
			m.RestrictedRole = types.BoolNull()
		}
	}
	if v, ok := obj["restricted-tcn"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.RestrictedTcn = types.BoolValue(b)
		} else {
			m.RestrictedTcn = types.BoolNull()
		}
	}
	if v, ok := obj["role"]; ok && v != "" {
		m.Role = types.StringValue(v)
	} else {
		m.Role = types.StringNull()
	}
	if v, ok := obj["status"]; ok && v != "" {
		m.Status = types.StringValue(v)
	} else {
		m.Status = types.StringNull()
	}
	if v, ok := obj["tag-stacking"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.TagStacking = types.BoolValue(b)
		} else {
			m.TagStacking = types.BoolNull()
		}
	}
	if v, ok := obj["trusted"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Trusted = types.BoolValue(b)
		} else {
			m.Trusted = types.BoolNull()
		}
	}
	if v, ok := obj["trusted-ra"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.TrustedRa = types.BoolValue(b)
		} else {
			m.TrustedRa = types.BoolNull()
		}
	}
	if v, ok := obj["unknown-multicast-flood"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.UnknownMulticastFlood = types.BoolValue(b)
		} else {
			m.UnknownMulticastFlood = types.BoolNull()
		}
	}
	if v, ok := obj["unknown-unicast-flood"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.UnknownUnicastFlood = types.BoolValue(b)
		} else {
			m.UnknownUnicastFlood = types.BoolNull()
		}
	}
}
