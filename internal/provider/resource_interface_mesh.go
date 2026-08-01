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
	_ resource.Resource                = &InterfaceMeshResource{}
	_ resource.ResourceWithImportState = &InterfaceMeshResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceMeshResource struct {
	reg *client.Registry
}

type InterfaceMeshModel struct {
	ID                       types.String `tfsdk:"id"`
	Name                     types.String `tfsdk:"name"`
	HwmpRannPropagationDelay types.String `tfsdk:"hwmp_rann_propagation_delay"`
	HwmpRannLifetime         types.String `tfsdk:"hwmp_rann_lifetime"`
	HwmpRannInterval         types.String `tfsdk:"hwmp_rann_interval"`
	HwmpPreqWaitingTime      types.String `tfsdk:"hwmp_preq_waiting_time"`
	HwmpPreqRetries          types.String `tfsdk:"hwmp_preq_retries"`
	HwmpPreqReplyAndForward  types.String `tfsdk:"hwmp_preq_reply_and_forward"`
	HwmpPreqDestinationOnly  types.String `tfsdk:"hwmp_preq_destination_only"`
	HwmpPrepLifetime         types.String `tfsdk:"hwmp_prep_lifetime"`
	HwmpDefaultHoplimit      types.String `tfsdk:"hwmp_default_hoplimit"`
	AutoMac                  types.String `tfsdk:"auto_mac"`
	AdminMac                 types.String `tfsdk:"admin_mac"`
	AdminMACAddress          types.String `tfsdk:"admin_mac_address"`
	ARP                      types.String `tfsdk:"arp"`
	ARPTimeout               types.String `tfsdk:"arp_timeout"`
	Comment                  types.String `tfsdk:"comment"`
	DefaultHoplimit          types.Int64  `tfsdk:"default_hoplimit"`
	Disabled                 types.Bool   `tfsdk:"disabled"`
	MACAddress               types.String `tfsdk:"mac_address"`
	MeshPortal               types.Bool   `tfsdk:"mesh_portal"`
	MeshTraceroute           types.String `tfsdk:"mesh_traceroute"`
	MTU                      types.Int64  `tfsdk:"mtu"`
	PrepLifetime             types.String `tfsdk:"prep_lifetime"`
	PreqDestinationOnly      types.Bool   `tfsdk:"preq_destination_only"`
	PreqReplyAndForward      types.Bool   `tfsdk:"preq_reply_and_forward"`
	PreqRetries              types.Int64  `tfsdk:"preq_retries"`
	PreqWaitingTime          types.Int64  `tfsdk:"preq_waiting_time"`
	RannInterval             types.String `tfsdk:"rann_interval"`
	RannLifetime             types.String `tfsdk:"rann_lifetime"`
	RannPropagationDelay     types.Int64  `tfsdk:"rann_propagation_delay"`
	ReoptimizePaths          types.Bool   `tfsdk:"reoptimize_paths"`
	Router                   types.String `tfsdk:"router"`
}

func NewInterfaceMeshResource() resource.Resource { return &InterfaceMeshResource{} }

func (r *InterfaceMeshResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_mesh"
}

func (r *InterfaceMeshResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceMeshResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/mesh`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `name`.",
			},
			"hwmp_rann_propagation_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hwmp-rann-propagation-delay`.",
			},
			"hwmp_rann_lifetime": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hwmp-rann-lifetime`.",
			},
			"hwmp_rann_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hwmp-rann-interval`.",
			},
			"hwmp_preq_waiting_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hwmp-preq-waiting-time`.",
			},
			"hwmp_preq_retries": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hwmp-preq-retries`.",
			},
			"hwmp_preq_reply_and_forward": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hwmp-preq-reply-and-forward`.",
			},
			"hwmp_preq_destination_only": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hwmp-preq-destination-only`.",
			},
			"hwmp_prep_lifetime": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hwmp-prep-lifetime`.",
			},
			"hwmp_default_hoplimit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hwmp-default-hoplimit`.",
			},
			"auto_mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `auto-mac`.",
			},
			"admin_mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `admin-mac`.",
			},
			"admin_mac_address": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"arp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"disabled", "enabled", "proxy-arp", "reply-only", "local-proxy-arp"}...)},
			},
			"arp_timeout": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationOrKeyword("auto")},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"default_hoplimit": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"mac_address": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"mesh_portal": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mesh_traceroute": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"prep_lifetime": schema.StringAttribute{
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"preq_destination_only": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"preq_reply_and_forward": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"preq_retries": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"preq_waiting_time": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"rann_interval": schema.StringAttribute{
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"rann_lifetime": schema.StringAttribute{
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"rann_propagation_delay": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"reoptimize_paths": schema.BoolAttribute{
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

func (r *InterfaceMeshResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceMeshModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.ARP.IsNull() || plan.ARP.IsUnknown()) {
		body["arp"] = plan.ARP.ValueString()
	}
	if !(plan.ARPTimeout.IsNull() || plan.ARPTimeout.IsUnknown()) {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.MeshPortal.IsNull() || plan.MeshPortal.IsUnknown()) {
		body["mesh-portal"] = client.FormatBool(plan.MeshPortal.ValueBool())
	}
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = client.FormatInt64(plan.MTU.ValueInt64())
	}
	if !(plan.ReoptimizePaths.IsNull() || plan.ReoptimizePaths.IsUnknown()) {
		body["reoptimize-paths"] = client.FormatBool(plan.ReoptimizePaths.ValueBool())
	}
	if !(plan.AdminMac.IsNull() || plan.AdminMac.IsUnknown()) {
		body["admin-mac"] = plan.AdminMac.ValueString()
	}
	if !(plan.AutoMac.IsNull() || plan.AutoMac.IsUnknown()) {
		body["auto-mac"] = plan.AutoMac.ValueString()
	}
	if !(plan.HwmpDefaultHoplimit.IsNull() || plan.HwmpDefaultHoplimit.IsUnknown()) {
		body["hwmp-default-hoplimit"] = plan.HwmpDefaultHoplimit.ValueString()
	}
	if !(plan.HwmpPrepLifetime.IsNull() || plan.HwmpPrepLifetime.IsUnknown()) {
		body["hwmp-prep-lifetime"] = plan.HwmpPrepLifetime.ValueString()
	}
	if !(plan.HwmpPreqDestinationOnly.IsNull() || plan.HwmpPreqDestinationOnly.IsUnknown()) {
		body["hwmp-preq-destination-only"] = plan.HwmpPreqDestinationOnly.ValueString()
	}
	if !(plan.HwmpPreqReplyAndForward.IsNull() || plan.HwmpPreqReplyAndForward.IsUnknown()) {
		body["hwmp-preq-reply-and-forward"] = plan.HwmpPreqReplyAndForward.ValueString()
	}
	if !(plan.HwmpPreqRetries.IsNull() || plan.HwmpPreqRetries.IsUnknown()) {
		body["hwmp-preq-retries"] = plan.HwmpPreqRetries.ValueString()
	}
	if !(plan.HwmpPreqWaitingTime.IsNull() || plan.HwmpPreqWaitingTime.IsUnknown()) {
		body["hwmp-preq-waiting-time"] = plan.HwmpPreqWaitingTime.ValueString()
	}
	if !(plan.HwmpRannInterval.IsNull() || plan.HwmpRannInterval.IsUnknown()) {
		body["hwmp-rann-interval"] = plan.HwmpRannInterval.ValueString()
	}
	if !(plan.HwmpRannLifetime.IsNull() || plan.HwmpRannLifetime.IsUnknown()) {
		body["hwmp-rann-lifetime"] = plan.HwmpRannLifetime.ValueString()
	}
	if !(plan.HwmpRannPropagationDelay.IsNull() || plan.HwmpRannPropagationDelay.IsUnknown()) {
		body["hwmp-rann-propagation-delay"] = plan.HwmpRannPropagationDelay.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/mesh", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/mesh failed", err.Error())
		return
	}
	interfaceMeshApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceMeshResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceMeshModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/mesh", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/mesh failed", err.Error())
		return
	}
	interfaceMeshApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceMeshResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceMeshModel
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
	if !plan.ARP.Equal(state.ARP) && !plan.ARP.IsUnknown() {
		body["arp"] = plan.ARP.ValueString()
	}
	if !plan.ARPTimeout.Equal(state.ARPTimeout) && !plan.ARPTimeout.IsUnknown() {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.MeshPortal.Equal(state.MeshPortal) && !plan.MeshPortal.IsUnknown() {
		body["mesh-portal"] = client.FormatBool(plan.MeshPortal.ValueBool())
	}
	if !plan.MTU.Equal(state.MTU) && !plan.MTU.IsUnknown() {
		body["mtu"] = client.FormatInt64(plan.MTU.ValueInt64())
	}
	if !plan.ReoptimizePaths.Equal(state.ReoptimizePaths) && !plan.ReoptimizePaths.IsUnknown() {
		body["reoptimize-paths"] = client.FormatBool(plan.ReoptimizePaths.ValueBool())
	}
	if !plan.AdminMac.Equal(state.AdminMac) && !plan.AdminMac.IsUnknown() {
		body["admin-mac"] = plan.AdminMac.ValueString()
	}
	if !plan.AutoMac.Equal(state.AutoMac) && !plan.AutoMac.IsUnknown() {
		body["auto-mac"] = plan.AutoMac.ValueString()
	}
	if !plan.HwmpDefaultHoplimit.Equal(state.HwmpDefaultHoplimit) && !plan.HwmpDefaultHoplimit.IsUnknown() {
		body["hwmp-default-hoplimit"] = plan.HwmpDefaultHoplimit.ValueString()
	}
	if !plan.HwmpPrepLifetime.Equal(state.HwmpPrepLifetime) && !plan.HwmpPrepLifetime.IsUnknown() {
		body["hwmp-prep-lifetime"] = plan.HwmpPrepLifetime.ValueString()
	}
	if !plan.HwmpPreqDestinationOnly.Equal(state.HwmpPreqDestinationOnly) && !plan.HwmpPreqDestinationOnly.IsUnknown() {
		body["hwmp-preq-destination-only"] = plan.HwmpPreqDestinationOnly.ValueString()
	}
	if !plan.HwmpPreqReplyAndForward.Equal(state.HwmpPreqReplyAndForward) && !plan.HwmpPreqReplyAndForward.IsUnknown() {
		body["hwmp-preq-reply-and-forward"] = plan.HwmpPreqReplyAndForward.ValueString()
	}
	if !plan.HwmpPreqRetries.Equal(state.HwmpPreqRetries) && !plan.HwmpPreqRetries.IsUnknown() {
		body["hwmp-preq-retries"] = plan.HwmpPreqRetries.ValueString()
	}
	if !plan.HwmpPreqWaitingTime.Equal(state.HwmpPreqWaitingTime) && !plan.HwmpPreqWaitingTime.IsUnknown() {
		body["hwmp-preq-waiting-time"] = plan.HwmpPreqWaitingTime.ValueString()
	}
	if !plan.HwmpRannInterval.Equal(state.HwmpRannInterval) && !plan.HwmpRannInterval.IsUnknown() {
		body["hwmp-rann-interval"] = plan.HwmpRannInterval.ValueString()
	}
	if !plan.HwmpRannLifetime.Equal(state.HwmpRannLifetime) && !plan.HwmpRannLifetime.IsUnknown() {
		body["hwmp-rann-lifetime"] = plan.HwmpRannLifetime.ValueString()
	}
	if !plan.HwmpRannPropagationDelay.Equal(state.HwmpRannPropagationDelay) && !plan.HwmpRannPropagationDelay.IsUnknown() {
		body["hwmp-rann-propagation-delay"] = plan.HwmpRannPropagationDelay.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/mesh", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/mesh failed", err.Error())
			return
		}
		interfaceMeshApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceMeshResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceMeshModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/mesh", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/mesh failed", err.Error())
	}
}

func (r *InterfaceMeshResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceMeshLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/mesh matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceMeshLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceMeshLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/mesh", id)
}

func interfaceMeshApply(ctx context.Context, obj client.Object, m *InterfaceMeshModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["hwmp-rann-propagation-delay"]; ok && v != "" {
		m.HwmpRannPropagationDelay = types.StringValue(v)
	} else {
		m.HwmpRannPropagationDelay = types.StringNull()
	}
	if v, ok := obj["hwmp-rann-lifetime"]; ok && v != "" {
		m.HwmpRannLifetime = types.StringValue(v)
	} else {
		m.HwmpRannLifetime = types.StringNull()
	}
	if v, ok := obj["hwmp-rann-interval"]; ok && v != "" {
		m.HwmpRannInterval = types.StringValue(v)
	} else {
		m.HwmpRannInterval = types.StringNull()
	}
	if v, ok := obj["hwmp-preq-waiting-time"]; ok && v != "" {
		m.HwmpPreqWaitingTime = types.StringValue(v)
	} else {
		m.HwmpPreqWaitingTime = types.StringNull()
	}
	if v, ok := obj["hwmp-preq-retries"]; ok && v != "" {
		m.HwmpPreqRetries = types.StringValue(v)
	} else {
		m.HwmpPreqRetries = types.StringNull()
	}
	if v, ok := obj["hwmp-preq-reply-and-forward"]; ok && v != "" {
		m.HwmpPreqReplyAndForward = types.StringValue(v)
	} else {
		m.HwmpPreqReplyAndForward = types.StringNull()
	}
	if v, ok := obj["hwmp-preq-destination-only"]; ok && v != "" {
		m.HwmpPreqDestinationOnly = types.StringValue(v)
	} else {
		m.HwmpPreqDestinationOnly = types.StringNull()
	}
	if v, ok := obj["hwmp-prep-lifetime"]; ok && v != "" {
		m.HwmpPrepLifetime = types.StringValue(v)
	} else {
		m.HwmpPrepLifetime = types.StringNull()
	}
	if v, ok := obj["hwmp-default-hoplimit"]; ok && v != "" {
		m.HwmpDefaultHoplimit = types.StringValue(v)
	} else {
		m.HwmpDefaultHoplimit = types.StringNull()
	}
	if v, ok := obj["auto-mac"]; ok && v != "" {
		m.AutoMac = types.StringValue(v)
	} else {
		m.AutoMac = types.StringNull()
	}
	if v, ok := obj["admin-mac"]; ok && v != "" {
		m.AdminMac = types.StringValue(v)
	} else {
		m.AdminMac = types.StringNull()
	}
	if v, ok := obj["admin-mac-address"]; ok {
		_ = v
		if v != "" {
			m.AdminMACAddress = types.StringValue(v)
		} else {
			m.AdminMACAddress = types.StringNull()
		}
	} else {
		m.AdminMACAddress = types.StringNull()
	}
	if v, ok := obj["arp"]; ok {
		_ = v
		if v != "" {
			m.ARP = types.StringValue(v)
		} else {
			m.ARP = types.StringNull()
		}
	} else {
		m.ARP = types.StringNull()
	}
	if v, ok := obj["arp-timeout"]; ok {
		_ = v
		if v != "" {
			m.ARPTimeout = types.StringValue(v)
		} else {
			m.ARPTimeout = types.StringNull()
		}
	} else {
		m.ARPTimeout = types.StringNull()
	}
	if v, ok := obj["comment"]; ok {
		_ = v
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	} else {
		m.Comment = types.StringNull()
	}
	if v, ok := obj["default-hoplimit"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.DefaultHoplimit = types.Int64Value(n)
		} else {
			m.DefaultHoplimit = types.Int64Null()
		}
	} else {
		m.DefaultHoplimit = types.Int64Null()
	}
	if v, ok := obj["disabled"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	} else {
		m.Disabled = types.BoolNull()
	}
	if v, ok := obj["mac-address"]; ok {
		_ = v
		if v != "" {
			m.MACAddress = types.StringValue(v)
		} else {
			m.MACAddress = types.StringNull()
		}
	} else {
		m.MACAddress = types.StringNull()
	}
	if v, ok := obj["mesh-portal"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.MeshPortal = types.BoolValue(b)
		} else {
			m.MeshPortal = types.BoolNull()
		}
	} else {
		m.MeshPortal = types.BoolNull()
	}
	if v, ok := obj["mesh-traceroute"]; ok {
		_ = v
		if v != "" {
			m.MeshTraceroute = types.StringValue(v)
		} else {
			m.MeshTraceroute = types.StringNull()
		}
	} else {
		m.MeshTraceroute = types.StringNull()
	}
	if v, ok := obj["mtu"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MTU = types.Int64Value(n)
		} else {
			m.MTU = types.Int64Null()
		}
	} else {
		m.MTU = types.Int64Null()
	}
	if v, ok := obj["prep-lifetime"]; ok {
		_ = v
		if v != "" {
			m.PrepLifetime = types.StringValue(v)
		} else {
			m.PrepLifetime = types.StringNull()
		}
	} else {
		m.PrepLifetime = types.StringNull()
	}
	if v, ok := obj["preq-destination-only"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.PreqDestinationOnly = types.BoolValue(b)
		} else {
			m.PreqDestinationOnly = types.BoolNull()
		}
	} else {
		m.PreqDestinationOnly = types.BoolNull()
	}
	if v, ok := obj["preq-reply-and-forward"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.PreqReplyAndForward = types.BoolValue(b)
		} else {
			m.PreqReplyAndForward = types.BoolNull()
		}
	} else {
		m.PreqReplyAndForward = types.BoolNull()
	}
	if v, ok := obj["preq-retries"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PreqRetries = types.Int64Value(n)
		} else {
			m.PreqRetries = types.Int64Null()
		}
	} else {
		m.PreqRetries = types.Int64Null()
	}
	if v, ok := obj["preq-waiting-time"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PreqWaitingTime = types.Int64Value(n)
		} else {
			m.PreqWaitingTime = types.Int64Null()
		}
	} else {
		m.PreqWaitingTime = types.Int64Null()
	}
	if v, ok := obj["rann-interval"]; ok {
		_ = v
		if v != "" {
			m.RannInterval = types.StringValue(v)
		} else {
			m.RannInterval = types.StringNull()
		}
	} else {
		m.RannInterval = types.StringNull()
	}
	if v, ok := obj["rann-lifetime"]; ok {
		_ = v
		if v != "" {
			m.RannLifetime = types.StringValue(v)
		} else {
			m.RannLifetime = types.StringNull()
		}
	} else {
		m.RannLifetime = types.StringNull()
	}
	if v, ok := obj["rann-propagation-delay"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RannPropagationDelay = types.Int64Value(n)
		} else {
			m.RannPropagationDelay = types.Int64Null()
		}
	} else {
		m.RannPropagationDelay = types.Int64Null()
	}
	if v, ok := obj["reoptimize-paths"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.ReoptimizePaths = types.BoolValue(b)
		} else {
			m.ReoptimizePaths = types.BoolNull()
		}
	} else {
		m.ReoptimizePaths = types.BoolNull()
	}
}
