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
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &InterfaceEthernetSwitchResource{}
	_ resource.ResourceWithImportState = &InterfaceEthernetSwitchResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceEthernetSwitchResource struct {
	reg *client.Registry
}

type InterfaceEthernetSwitchModel struct {
	ID               types.String `tfsdk:"id"`
	SwitchAllPorts   types.String `tfsdk:"switch_all_ports"`
	Name             types.String `tfsdk:"name"`
	MirrorSource     types.String `tfsdk:"mirror_source"`
	MirrorTarget     types.String `tfsdk:"mirror_target"`
	CPUFlowControl   types.String `tfsdk:"cpu_flow_control"`
	Autorestart      types.String `tfsdk:"autorestart"`
	FasttrackHw      types.String `tfsdk:"fasttrack_hw"`
	IcmpReplyOnError types.String `tfsdk:"icmp_reply_on_error"`
	IPV6Hw           types.String `tfsdk:"ipv6_hw"`
	Router           types.String `tfsdk:"router"`
}

func NewInterfaceEthernetSwitchResource() resource.Resource {
	return &InterfaceEthernetSwitchResource{}
}

func (r *InterfaceEthernetSwitchResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ethernet_switch"
}

func (r *InterfaceEthernetSwitchResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceEthernetSwitchResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Switch chip menu varies by hardware; not on hAP/CHR",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"switch_all_ports": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `switch-all-ports`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Switch name as reported by RouterOS, e.g. `switch1`.",
			},
			"mirror_source": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port whose traffic is mirrored, or `none` (the default).",
			},
			"mirror_target": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Port that receives mirrored traffic, or `none` (the default).",
			},
			"cpu_flow_control": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the switch sends pause frames to the CPU port.",
			},
			"autorestart": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Automatically restarts the l3hw driver in case of an error. Otherwise, if an error occurs, \u00a0 l3-hw-offloading \u00a0 gets disabled, and the error code is displayed in the switch settings and \u00a0 #monitor . Autorestart does not work for system failures, such as OOM (Out Of Memory).",
			},
			"fasttrack_hw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables or disables FastTrack HW Offloading. Keep it enabled unless HW TCAM memory reservation is required, e.g., for dynamic switch ACL rules creation. Not all switch chips support FastTrack HW Offloading (see \u00a0 hw-supports-fasttrack ).",
			},
			"icmp_reply_on_error": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Since the hardware cannot send ICMP messages, the packet must be redirected to the CPU to send an ICMP reply in case of an error (e.g., \"Time Exceeded\", \"Fragmentation required\", etc.). Enabling icmp-reply-on-error \u00a0 helps with network diagnostics but may open potential vulnerabilities for DDoS attacks. Disabling icmp-reply-on-error silently drops the packets on the hardware level in case of an error.",
			},
			"ipv6_hw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Enables or disables IPv6 Hardware Offloading. Since IPv6 routes occupy a lot of HW memory, enable it only if IPv6 traffic speed is significant enough to benefit from hardware routing.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceEthernetSwitchResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceEthernetSwitchModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.MirrorSource.IsNull() || plan.MirrorSource.IsUnknown()) {
		body["mirror-source"] = plan.MirrorSource.ValueString()
	}
	if !(plan.MirrorTarget.IsNull() || plan.MirrorTarget.IsUnknown()) {
		body["mirror-target"] = plan.MirrorTarget.ValueString()
	}
	if !(plan.CPUFlowControl.IsNull() || plan.CPUFlowControl.IsUnknown()) {
		body["cpu-flow-control"] = plan.CPUFlowControl.ValueString()
	}
	if !(plan.Autorestart.IsNull() || plan.Autorestart.IsUnknown()) {
		body["autorestart"] = plan.Autorestart.ValueString()
	}
	if !(plan.FasttrackHw.IsNull() || plan.FasttrackHw.IsUnknown()) {
		body["fasttrack-hw"] = plan.FasttrackHw.ValueString()
	}
	if !(plan.IcmpReplyOnError.IsNull() || plan.IcmpReplyOnError.IsUnknown()) {
		body["icmp-reply-on-error"] = plan.IcmpReplyOnError.ValueString()
	}
	if !(plan.IPV6Hw.IsNull() || plan.IPV6Hw.IsUnknown()) {
		body["ipv6-hw"] = plan.IPV6Hw.ValueString()
	}
	if !(plan.SwitchAllPorts.IsNull() || plan.SwitchAllPorts.IsUnknown()) {
		body["switch-all-ports"] = plan.SwitchAllPorts.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/ethernet/switch", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/ethernet/switch failed", err.Error())
		return
	}
	interfaceEthernetSwitchApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEthernetSwitchResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceEthernetSwitchModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/ethernet/switch", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/ethernet/switch failed", err.Error())
		return
	}
	interfaceEthernetSwitchApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceEthernetSwitchResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceEthernetSwitchModel
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
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.MirrorSource.Equal(state.MirrorSource) {
		body["mirror-source"] = plan.MirrorSource.ValueString()
	}
	if !plan.MirrorTarget.Equal(state.MirrorTarget) {
		body["mirror-target"] = plan.MirrorTarget.ValueString()
	}
	if !plan.CPUFlowControl.Equal(state.CPUFlowControl) {
		body["cpu-flow-control"] = plan.CPUFlowControl.ValueString()
	}
	if !plan.Autorestart.Equal(state.Autorestart) {
		body["autorestart"] = plan.Autorestart.ValueString()
	}
	if !plan.FasttrackHw.Equal(state.FasttrackHw) {
		body["fasttrack-hw"] = plan.FasttrackHw.ValueString()
	}
	if !plan.IcmpReplyOnError.Equal(state.IcmpReplyOnError) {
		body["icmp-reply-on-error"] = plan.IcmpReplyOnError.ValueString()
	}
	if !plan.IPV6Hw.Equal(state.IPV6Hw) {
		body["ipv6-hw"] = plan.IPV6Hw.ValueString()
	}
	if !plan.SwitchAllPorts.Equal(state.SwitchAllPorts) && !plan.SwitchAllPorts.IsUnknown() {
		body["switch-all-ports"] = plan.SwitchAllPorts.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/ethernet/switch", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/ethernet/switch failed", err.Error())
			return
		}
		interfaceEthernetSwitchApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEthernetSwitchResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceEthernetSwitchModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/ethernet/switch", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/ethernet/switch failed", err.Error())
	}
}

func (r *InterfaceEthernetSwitchResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceEthernetSwitchLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/ethernet/switch matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceEthernetSwitchLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceEthernetSwitchLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/ethernet/switch", id)
}

func interfaceEthernetSwitchApply(ctx context.Context, obj client.Object, m *InterfaceEthernetSwitchModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["switch-all-ports"]; ok && v != "" {
		m.SwitchAllPorts = types.StringValue(v)
	} else {
		m.SwitchAllPorts = types.StringNull()
	}
	if v, ok := obj["name"]; ok {
		_ = v
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["mirror-source"]; ok {
		_ = v
		if v != "" {
			m.MirrorSource = types.StringValue(v)
		} else {
			m.MirrorSource = types.StringNull()
		}
	} else {
		m.MirrorSource = types.StringNull()
	}
	if v, ok := obj["mirror-target"]; ok {
		_ = v
		if v != "" {
			m.MirrorTarget = types.StringValue(v)
		} else {
			m.MirrorTarget = types.StringNull()
		}
	} else {
		m.MirrorTarget = types.StringNull()
	}
	if v, ok := obj["cpu-flow-control"]; ok {
		_ = v
		if v != "" {
			m.CPUFlowControl = types.StringValue(v)
		} else {
			m.CPUFlowControl = types.StringNull()
		}
	} else {
		m.CPUFlowControl = types.StringNull()
	}
	if v, ok := obj["autorestart"]; ok {
		_ = v
		if v != "" {
			m.Autorestart = types.StringValue(v)
		} else {
			m.Autorestart = types.StringNull()
		}
	} else {
		m.Autorestart = types.StringNull()
	}
	if v, ok := obj["fasttrack-hw"]; ok {
		_ = v
		if v != "" {
			m.FasttrackHw = types.StringValue(v)
		} else {
			m.FasttrackHw = types.StringNull()
		}
	} else {
		m.FasttrackHw = types.StringNull()
	}
	if v, ok := obj["icmp-reply-on-error"]; ok {
		_ = v
		if v != "" {
			m.IcmpReplyOnError = types.StringValue(v)
		} else {
			m.IcmpReplyOnError = types.StringNull()
		}
	} else {
		m.IcmpReplyOnError = types.StringNull()
	}
	if v, ok := obj["ipv6-hw"]; ok {
		_ = v
		if v != "" {
			m.IPV6Hw = types.StringValue(v)
		} else {
			m.IPV6Hw = types.StringNull()
		}
	} else {
		m.IPV6Hw = types.StringNull()
	}
}
