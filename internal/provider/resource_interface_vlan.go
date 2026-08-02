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
	_ resource.Resource                = &InterfaceVLANResource{}
	_ resource.ResourceWithImportState = &InterfaceVLANResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceVLANResource struct {
	reg *client.Registry
}

type InterfaceVLANModel struct {
	ID                      types.String    `tfsdk:"id"`
	ARP                     types.String    `tfsdk:"arp"`
	ARPTimeout              types.String    `tfsdk:"arp_timeout"`
	Comment                 types.String    `tfsdk:"comment"`
	Disabled                types.Bool      `tfsdk:"disabled"`
	Interface               types.String    `tfsdk:"interface"`
	LoopProtect             types.String    `tfsdk:"loop_protect"`
	LoopProtectDisableTime  types.String    `tfsdk:"loop_protect_disable_time"`
	LoopProtectSendInterval types.String    `tfsdk:"loop_protect_send_interval"`
	MTU                     types.String    `tfsdk:"mtu"`
	Mvrp                    boolStringValue `tfsdk:"mvrp"`
	Name                    types.String    `tfsdk:"name"`
	UseServiceTag           boolStringValue `tfsdk:"use_service_tag"`
	VLANID                  types.String    `tfsdk:"vlan_id"`
	Router                  types.String    `tfsdk:"router"`
}

func NewInterfaceVLANResource() resource.Resource { return &InterfaceVLANResource{} }

func (r *InterfaceVLANResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_vlan"
}

func (r *InterfaceVLANResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceVLANResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/vlan`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"arp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Address Resolution Protocol setting disabled \u00a0 - the interface will not use ARP enabled \u00a0 - the interface will use ARP local-proxy-arp \u00a0 - \u00a0 \u00a0 the router performs proxy ARP on the interface and sends replies to the same interface proxy-arp \u00a0 - \u00a0 the router performs proxy ARP on the interface and sends replies to other interfaces reply-only \u00a0 - the interface will only reply to requests originated from matching IP address/MAC address combinations which are entered as static entries in the \u00a0 IP/ARP \u00a0 table. No dynamic entries will be automatically stored in the \u00a0 IP/ARP \u00a0 table. Therefore for communications to be successful, a valid static entry must already exist.",
			},
			"arp_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "How long the ARP record is kept in the ARP table after no packets are received from IP. Value \u00a0 auto \u00a0 equals to the value of \u00a0 arp-timeout \u00a0 in \u00a0 IP/Settings, default is 30s.",
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
			"interface": schema.StringAttribute{
				Required:    true,
				Description: "Name of the interface on top of which VLAN will work. Adding a VLAN interface to a bridge with vlan-filtering enabled will automatically tag the bridge interface as a member port. A dynamic entry with the comment \"added by vlan on bridge\" will appear under the /interface/bridge/vlan menu.",
			},
			"loop_protect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Loop protection: `default`, `on` or `off`.",
			},
			"loop_protect_disable_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "How long the interface stays disabled after a loop is detected. `0s` disables the timeout.",
			},
			"loop_protect_send_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interval between loop-protect probe packets.",
			},
			"mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Layer3 Maximum transmission unit",
			},
			"mvrp": schema.StringAttribute{
				CustomType:  boolStringType{},
				Optional:    true,
				Computed:    true,
				Description: "Specifies whether this VLAN should declare its attributes through Multiple VLAN Registration Protocol (MVRP) as an applicant. Its main use case is for VLANs that is created on Ethernet interface (such as a \"router on a stick\" setup) that is connected to a bridge supporting MVRP . Enabling this option on a VLAN interface that is already part of an MVRP-enabled bridge has no effect, as the bridge manages MVRP in that case. \u00a0 This property only has an effect when use-service-tag \u00a0 is disabled .",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Interface name",
			},
			"use_service_tag": schema.StringAttribute{
				CustomType:  boolStringType{},
				Optional:    true,
				Computed:    true,
				Description: "IEEE 802.1ad compatible Service Tag",
			},
			"vlan_id": schema.StringAttribute{
				Required:    true,
				Description: "Virtual LAN identifier or tag that is used to distinguish VLANs. Must be equal for all computers that belong to the same VLAN.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceVLANResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceVLANModel
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
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.LoopProtect.IsNull() || plan.LoopProtect.IsUnknown()) {
		body["loop-protect"] = plan.LoopProtect.ValueString()
	}
	if !(plan.LoopProtectDisableTime.IsNull() || plan.LoopProtectDisableTime.IsUnknown()) {
		body["loop-protect-disable-time"] = plan.LoopProtectDisableTime.ValueString()
	}
	if !(plan.LoopProtectSendInterval.IsNull() || plan.LoopProtectSendInterval.IsUnknown()) {
		body["loop-protect-send-interval"] = plan.LoopProtectSendInterval.ValueString()
	}
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !(plan.Mvrp.IsNull() || plan.Mvrp.IsUnknown()) {
		body["mvrp"] = plan.Mvrp.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.UseServiceTag.IsNull() || plan.UseServiceTag.IsUnknown()) {
		body["use-service-tag"] = plan.UseServiceTag.ValueString()
	}
	if !(plan.VLANID.IsNull() || plan.VLANID.IsUnknown()) {
		body["vlan-id"] = plan.VLANID.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/vlan", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/vlan failed", err.Error())
		return
	}
	interfaceVLANApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceVLANResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceVLANModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/vlan", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/vlan failed", err.Error())
		return
	}
	interfaceVLANApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceVLANResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceVLANModel
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
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.LoopProtect.Equal(state.LoopProtect) && !plan.LoopProtect.IsUnknown() {
		body["loop-protect"] = plan.LoopProtect.ValueString()
	}
	if !plan.LoopProtectDisableTime.Equal(state.LoopProtectDisableTime) && !plan.LoopProtectDisableTime.IsUnknown() {
		body["loop-protect-disable-time"] = plan.LoopProtectDisableTime.ValueString()
	}
	if !plan.LoopProtectSendInterval.Equal(state.LoopProtectSendInterval) && !plan.LoopProtectSendInterval.IsUnknown() {
		body["loop-protect-send-interval"] = plan.LoopProtectSendInterval.ValueString()
	}
	if !plan.MTU.Equal(state.MTU) && !plan.MTU.IsUnknown() {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !plan.Mvrp.Equal(state.Mvrp) && !plan.Mvrp.IsUnknown() {
		body["mvrp"] = plan.Mvrp.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.UseServiceTag.Equal(state.UseServiceTag) && !plan.UseServiceTag.IsUnknown() {
		body["use-service-tag"] = plan.UseServiceTag.ValueString()
	}
	if !plan.VLANID.Equal(state.VLANID) && !plan.VLANID.IsUnknown() {
		body["vlan-id"] = plan.VLANID.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/vlan", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/vlan failed", err.Error())
			return
		}
		interfaceVLANApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceVLANResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceVLANModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/vlan", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/vlan failed", err.Error())
	}
}

func (r *InterfaceVLANResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceVLANLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/vlan matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceVLANLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceVLANLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/vlan", id)
}

func interfaceVLANApply(ctx context.Context, obj client.Object, m *InterfaceVLANModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["arp"]; ok {
		if v != "" {
			m.ARP = types.StringValue(v)
		} else {
			m.ARP = types.StringNull()
		}
	}
	if v, ok := obj["arp-timeout"]; ok {
		if v != "" {
			m.ARPTimeout = types.StringValue(v)
		} else {
			m.ARPTimeout = types.StringNull()
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
		} else if strings.TrimSpace(v) == "" {
			m.Disabled = types.BoolValue(true)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["interface"]; ok {
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	}
	if v, ok := obj["loop-protect"]; ok {
		if v != "" {
			m.LoopProtect = types.StringValue(v)
		} else {
			m.LoopProtect = types.StringNull()
		}
	}
	if v, ok := obj["loop-protect-disable-time"]; ok {
		if v != "" {
			m.LoopProtectDisableTime = types.StringValue(v)
		} else {
			m.LoopProtectDisableTime = types.StringNull()
		}
	}
	if v, ok := obj["loop-protect-send-interval"]; ok {
		if v != "" {
			m.LoopProtectSendInterval = types.StringValue(v)
		} else {
			m.LoopProtectSendInterval = types.StringNull()
		}
	}
	if v, ok := obj["mtu"]; ok {
		if v != "" {
			m.MTU = types.StringValue(v)
		} else {
			m.MTU = types.StringNull()
		}
	}
	if v, ok := obj["mvrp"]; ok {
		_ = v
		if v != "" {
			m.Mvrp = newBoolStringValue(v)
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["use-service-tag"]; ok {
		_ = v
		if v != "" {
			m.UseServiceTag = newBoolStringValue(v)
		}
	}
	if v, ok := obj["vlan-id"]; ok {
		if v != "" {
			m.VLANID = types.StringValue(v)
		} else {
			m.VLANID = types.StringNull()
		}
	}
}
