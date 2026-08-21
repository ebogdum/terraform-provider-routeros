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
	_ resource.Resource                = &InterfaceVplsResource{}
	_ resource.ResourceWithImportState = &InterfaceVplsResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceVplsResource struct {
	reg *client.Registry
}

type InterfaceVplsModel struct {
	ID                  types.String `tfsdk:"id"`
	Peer                types.String `tfsdk:"peer"`
	Name                types.String `tfsdk:"name"`
	DisableRunningCheck types.String `tfsdk:"disable_running_check"`
	ARP                 types.String `tfsdk:"arp"`
	ARPTimeout          types.String `tfsdk:"arp_timeout"`
	BGPSignaled         types.Bool   `tfsdk:"bgp_signaled"`
	BGPVpls             types.String `tfsdk:"bgp_vpls"`
	BGPVplsPrefix       types.String `tfsdk:"bgp_vpls_prefix"`
	Bridge              types.String `tfsdk:"bridge"`
	BridgeCost          types.String `tfsdk:"bridge_cost"`
	BridgeHorizon       types.String `tfsdk:"bridge_horizon"`
	BridgePvid          types.String `tfsdk:"bridge_pvid"`
	CiscoBGPSignaled    types.Bool   `tfsdk:"cisco_bgp_signaled"`
	CiscoStaticID       types.String `tfsdk:"cisco_static_id"`
	Comment             types.String `tfsdk:"comment"`
	Disabled            types.Bool   `tfsdk:"disabled"`
	LocalLabel          types.Int64  `tfsdk:"local_label"`
	MACAddress          types.String `tfsdk:"mac_address"`
	MTU                 types.Int64  `tfsdk:"mtu"`
	PwControlWord       types.String `tfsdk:"pw_control_word"`
	PwL2mtu             types.String `tfsdk:"pw_l2mtu"`
	PwType              types.String `tfsdk:"pw_type"`
	RemoteGroup         types.Int64  `tfsdk:"remote_group"`
	RemoteLabel         types.Int64  `tfsdk:"remote_label"`
	RemotePeer          types.String `tfsdk:"remote_peer"`
	RemoteStatus        types.String `tfsdk:"remote_status"`
	TeTunnel            types.Int64  `tfsdk:"te_tunnel"`
	VplsID              types.String `tfsdk:"vpls_id"`
	Router              types.String `tfsdk:"router"`
}

func NewInterfaceVplsResource() resource.Resource { return &InterfaceVplsResource{} }

func (r *InterfaceVplsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_vpls"
}

func (r *InterfaceVplsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceVplsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "VPLS needs MPLS/LDP setup; skipped from automated acc tests.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"peer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `peer`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `name`.",
			},
			"disable_running_check": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `disable-running-check`.",
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
			"bgp_signaled": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"bgp_vpls": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"bgp_vpls_prefix": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"bridge": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge_cost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge_horizon": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge_pvid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"cisco_bgp_signaled": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"cisco_static_id": schema.StringAttribute{
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
			"local_label": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"mac_address": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsMAC()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeMAC()},
			},
			"mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pw_control_word": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pw_l2mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pw_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_group": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"remote_label": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"remote_peer": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"remote_status": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"te_tunnel": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"vpls_id": schema.StringAttribute{
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

func (r *InterfaceVplsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceVplsModel
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
	if !(plan.Bridge.IsNull() || plan.Bridge.IsUnknown()) {
		body["bridge"] = plan.Bridge.ValueString()
	}
	if !(plan.BridgeCost.IsNull() || plan.BridgeCost.IsUnknown()) {
		body["bridge-cost"] = plan.BridgeCost.ValueString()
	}
	if !(plan.BridgeHorizon.IsNull() || plan.BridgeHorizon.IsUnknown()) {
		body["bridge-horizon"] = plan.BridgeHorizon.ValueString()
	}
	if !(plan.BridgePvid.IsNull() || plan.BridgePvid.IsUnknown()) {
		body["bridge-pvid"] = plan.BridgePvid.ValueString()
	}
	if !(plan.CiscoStaticID.IsNull() || plan.CiscoStaticID.IsUnknown()) {
		body["cisco-static-id"] = plan.CiscoStaticID.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.MACAddress.IsNull() || plan.MACAddress.IsUnknown()) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = client.FormatInt64(plan.MTU.ValueInt64())
	}
	if !(plan.PwControlWord.IsNull() || plan.PwControlWord.IsUnknown()) {
		body["pw-control-word"] = plan.PwControlWord.ValueString()
	}
	if !(plan.PwL2mtu.IsNull() || plan.PwL2mtu.IsUnknown()) {
		body["pw-l2mtu"] = plan.PwL2mtu.ValueString()
	}
	if !(plan.PwType.IsNull() || plan.PwType.IsUnknown()) {
		body["pw-type"] = plan.PwType.ValueString()
	}
	if !(plan.VplsID.IsNull() || plan.VplsID.IsUnknown()) {
		body["vpls-id"] = plan.VplsID.ValueString()
	}
	if !(plan.DisableRunningCheck.IsNull() || plan.DisableRunningCheck.IsUnknown()) {
		body["disable-running-check"] = plan.DisableRunningCheck.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Peer.IsNull() || plan.Peer.IsUnknown()) {
		body["peer"] = plan.Peer.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/vpls", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/vpls failed", err.Error())
		return
	}
	interfaceVplsApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceVplsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceVplsModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/vpls", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/vpls failed", err.Error())
		return
	}
	interfaceVplsApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceVplsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceVplsModel
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
	if !plan.Bridge.Equal(state.Bridge) && !plan.Bridge.IsUnknown() {
		body["bridge"] = plan.Bridge.ValueString()
	}
	if !plan.BridgeCost.Equal(state.BridgeCost) && !plan.BridgeCost.IsUnknown() {
		body["bridge-cost"] = plan.BridgeCost.ValueString()
	}
	if !plan.BridgeHorizon.Equal(state.BridgeHorizon) && !plan.BridgeHorizon.IsUnknown() {
		body["bridge-horizon"] = plan.BridgeHorizon.ValueString()
	}
	if !plan.BridgePvid.Equal(state.BridgePvid) && !plan.BridgePvid.IsUnknown() {
		body["bridge-pvid"] = plan.BridgePvid.ValueString()
	}
	if !plan.CiscoStaticID.Equal(state.CiscoStaticID) && !plan.CiscoStaticID.IsUnknown() {
		body["cisco-static-id"] = plan.CiscoStaticID.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.MACAddress.Equal(state.MACAddress) && !plan.MACAddress.IsUnknown() {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.MTU.Equal(state.MTU) && !plan.MTU.IsUnknown() {
		body["mtu"] = client.FormatInt64(plan.MTU.ValueInt64())
	}
	if !plan.PwControlWord.Equal(state.PwControlWord) && !plan.PwControlWord.IsUnknown() {
		body["pw-control-word"] = plan.PwControlWord.ValueString()
	}
	if !plan.PwL2mtu.Equal(state.PwL2mtu) && !plan.PwL2mtu.IsUnknown() {
		body["pw-l2mtu"] = plan.PwL2mtu.ValueString()
	}
	if !plan.PwType.Equal(state.PwType) && !plan.PwType.IsUnknown() {
		body["pw-type"] = plan.PwType.ValueString()
	}
	if !plan.VplsID.Equal(state.VplsID) && !plan.VplsID.IsUnknown() {
		body["vpls-id"] = plan.VplsID.ValueString()
	}
	if !plan.DisableRunningCheck.Equal(state.DisableRunningCheck) && !plan.DisableRunningCheck.IsUnknown() {
		body["disable-running-check"] = plan.DisableRunningCheck.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Peer.Equal(state.Peer) && !plan.Peer.IsUnknown() {
		body["peer"] = plan.Peer.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/vpls", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/vpls failed", err.Error())
			return
		}
		interfaceVplsApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceVplsResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceVplsModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/vpls", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/vpls failed", err.Error())
	}
}

func (r *InterfaceVplsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceVplsLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/vpls matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceVplsLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceVplsLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/vpls", id)
}

func interfaceVplsApply(ctx context.Context, obj client.Object, m *InterfaceVplsModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["peer"]; ok && v != "" {
		m.Peer = types.StringValue(v)
	} else {
		m.Peer = types.StringNull()
	}
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["disable-running-check"]; ok && v != "" {
		m.DisableRunningCheck = types.StringValue(v)
	} else {
		m.DisableRunningCheck = types.StringNull()
	}
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
	if v, ok := obj["bgp-signaled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.BGPSignaled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.BGPSignaled = types.BoolValue(true)
		} else {
			m.BGPSignaled = types.BoolNull()
		}
	}
	if v, ok := obj["bgp-vpls"]; ok {
		if v != "" {
			m.BGPVpls = types.StringValue(v)
		} else {
			m.BGPVpls = types.StringNull()
		}
	}
	if v, ok := obj["bgp-vpls-prefix"]; ok {
		if v != "" {
			m.BGPVplsPrefix = types.StringValue(v)
		} else {
			m.BGPVplsPrefix = types.StringNull()
		}
	}
	if v, ok := obj["bridge"]; ok {
		if v != "" {
			m.Bridge = types.StringValue(v)
		} else {
			m.Bridge = types.StringNull()
		}
	}
	if v, ok := obj["bridge-cost"]; ok {
		if v != "" {
			m.BridgeCost = types.StringValue(v)
		} else {
			m.BridgeCost = types.StringNull()
		}
	}
	if v, ok := obj["bridge-horizon"]; ok {
		if v != "" {
			m.BridgeHorizon = types.StringValue(v)
		} else {
			m.BridgeHorizon = types.StringNull()
		}
	}
	if v, ok := obj["bridge-pvid"]; ok {
		if v != "" {
			m.BridgePvid = types.StringValue(v)
		} else {
			m.BridgePvid = types.StringNull()
		}
	}
	if v, ok := obj["cisco-bgp-signaled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.CiscoBGPSignaled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.CiscoBGPSignaled = types.BoolValue(true)
		} else {
			m.CiscoBGPSignaled = types.BoolNull()
		}
	}
	if v, ok := obj["cisco-static-id"]; ok {
		if v != "" {
			m.CiscoStaticID = types.StringValue(v)
		} else {
			m.CiscoStaticID = types.StringNull()
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
	if v, ok := obj["local-label"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.LocalLabel = types.Int64Value(n)
		} else {
			m.LocalLabel = types.Int64Null()
		}
	} else {
		m.LocalLabel = types.Int64Null()
	}
	if v, ok := obj["mac-address"]; ok {
		if v != "" {
			m.MACAddress = types.StringValue(v)
		} else {
			m.MACAddress = types.StringNull()
		}
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
	if v, ok := obj["pw-control-word"]; ok {
		if v != "" {
			m.PwControlWord = types.StringValue(v)
		} else {
			m.PwControlWord = types.StringNull()
		}
	}
	if v, ok := obj["pw-l2mtu"]; ok {
		if v != "" {
			m.PwL2mtu = types.StringValue(v)
		} else {
			m.PwL2mtu = types.StringNull()
		}
	}
	if v, ok := obj["pw-type"]; ok {
		if v != "" {
			m.PwType = types.StringValue(v)
		} else {
			m.PwType = types.StringNull()
		}
	}
	if v, ok := obj["remote-group"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RemoteGroup = types.Int64Value(n)
		} else {
			m.RemoteGroup = types.Int64Null()
		}
	} else {
		m.RemoteGroup = types.Int64Null()
	}
	if v, ok := obj["remote-label"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RemoteLabel = types.Int64Value(n)
		} else {
			m.RemoteLabel = types.Int64Null()
		}
	} else {
		m.RemoteLabel = types.Int64Null()
	}
	if v, ok := obj["remote-peer"]; ok {
		if v != "" {
			m.RemotePeer = types.StringValue(v)
		} else {
			m.RemotePeer = types.StringNull()
		}
	}
	if v, ok := obj["remote-status"]; ok {
		if v != "" {
			m.RemoteStatus = types.StringValue(v)
		} else {
			m.RemoteStatus = types.StringNull()
		}
	}
	if v, ok := obj["te-tunnel"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TeTunnel = types.Int64Value(n)
		} else {
			m.TeTunnel = types.Int64Null()
		}
	} else {
		m.TeTunnel = types.Int64Null()
	}
	if v, ok := obj["vpls-id"]; ok {
		if v != "" {
			m.VplsID = types.StringValue(v)
		} else {
			m.VplsID = types.StringNull()
		}
	}
}
