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
	_ resource.Resource                = &InterfaceVxlanResource{}
	_ resource.ResourceWithImportState = &InterfaceVxlanResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceVxlanResource struct {
	reg *client.Registry
}

type InterfaceVxlanModel struct {
	ID                      types.String `tfsdk:"id"`
	VtepsIpVersion          types.String `tfsdk:"vteps_ip_version"`
	VtepVrf                 types.String `tfsdk:"vtep_vrf"`
	RemCsum                 types.String `tfsdk:"rem_csum"`
	MaxFdbSize              types.String `tfsdk:"max_fdb_size"`
	LoopProtectSendInterval types.String `tfsdk:"loop_protect_send_interval"`
	LoopProtectDisableTime  types.String `tfsdk:"loop_protect_disable_time"`
	LoopProtect             types.String `tfsdk:"loop_protect"`
	Learning                types.String `tfsdk:"learning"`
	Hw                      types.String `tfsdk:"hw"`
	Group                   types.String `tfsdk:"group"`
	DontFragment            types.String `tfsdk:"dont_fragment"`
	Checksum                types.String `tfsdk:"checksum"`
	BridgePvid              types.String `tfsdk:"bridge_pvid"`
	AllowFastPath           types.String `tfsdk:"allow_fast_path"`
	ARP                     types.String `tfsdk:"arp"`
	ARPTimeout              types.String `tfsdk:"arp_timeout"`
	Bridge                  types.String `tfsdk:"bridge"`
	Comment                 types.String `tfsdk:"comment"`
	Disabled                types.Bool   `tfsdk:"disabled"`
	Interface               types.String `tfsdk:"interface"`
	LocalAddress            types.String `tfsdk:"local_address"`
	MACAddress              types.String `tfsdk:"mac_address"`
	MTU                     types.String `tfsdk:"mtu"`
	Name                    types.String `tfsdk:"name"`
	Port                    types.String `tfsdk:"port"`
	Ttl                     types.String `tfsdk:"ttl"`
	Vni                     types.String `tfsdk:"vni"`
	Router                  types.String `tfsdk:"router"`
}

func NewInterfaceVxlanResource() resource.Resource { return &InterfaceVxlanResource{} }

func (r *InterfaceVxlanResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_vxlan"
}

func (r *InterfaceVxlanResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceVxlanResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/vxlan`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"vteps_ip_version": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `vteps-ip-version`.",
			},
			"vtep_vrf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `vtep-vrf`.",
			},
			"rem_csum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `rem-csum`.",
			},
			"max_fdb_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `max-fdb-size`.",
			},
			"loop_protect_send_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `loop-protect-send-interval`.",
			},
			"loop_protect_disable_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `loop-protect-disable-time`.",
			},
			"loop_protect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `loop-protect`.",
			},
			"learning": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `learning`.",
			},
			"hw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hw`.",
			},
			"group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `group`.",
			},
			"dont_fragment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dont-fragment`.",
			},
			"checksum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `checksum`.",
			},
			"bridge_pvid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bridge-pvid`.",
			},
			"allow_fast_path": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `allow-fast-path`.",
			},
			"arp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"arp_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge": schema.StringAttribute{
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
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"local_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vni": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceVxlanResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceVxlanModel
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
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.LocalAddress.IsNull() || plan.LocalAddress.IsUnknown()) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !(plan.MACAddress.IsNull() || plan.MACAddress.IsUnknown()) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.Ttl.IsNull() || plan.Ttl.IsUnknown()) {
		body["ttl"] = plan.Ttl.ValueString()
	}
	if !(plan.Vni.IsNull() || plan.Vni.IsUnknown()) {
		body["vni"] = plan.Vni.ValueString()
	}
	if !(plan.AllowFastPath.IsNull() || plan.AllowFastPath.IsUnknown()) {
		body["allow-fast-path"] = plan.AllowFastPath.ValueString()
	}
	if !(plan.BridgePvid.IsNull() || plan.BridgePvid.IsUnknown()) {
		body["bridge-pvid"] = plan.BridgePvid.ValueString()
	}
	if !(plan.Checksum.IsNull() || plan.Checksum.IsUnknown()) {
		body["checksum"] = plan.Checksum.ValueString()
	}
	if !(plan.DontFragment.IsNull() || plan.DontFragment.IsUnknown()) {
		body["dont-fragment"] = plan.DontFragment.ValueString()
	}
	if !(plan.Group.IsNull() || plan.Group.IsUnknown()) {
		body["group"] = plan.Group.ValueString()
	}
	if !(plan.Hw.IsNull() || plan.Hw.IsUnknown()) {
		body["hw"] = plan.Hw.ValueString()
	}
	if !(plan.Learning.IsNull() || plan.Learning.IsUnknown()) {
		body["learning"] = plan.Learning.ValueString()
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
	if !(plan.MaxFdbSize.IsNull() || plan.MaxFdbSize.IsUnknown()) {
		body["max-fdb-size"] = plan.MaxFdbSize.ValueString()
	}
	if !(plan.RemCsum.IsNull() || plan.RemCsum.IsUnknown()) {
		body["rem-csum"] = plan.RemCsum.ValueString()
	}
	if !(plan.VtepVrf.IsNull() || plan.VtepVrf.IsUnknown()) {
		body["vtep-vrf"] = plan.VtepVrf.ValueString()
	}
	if !(plan.VtepsIpVersion.IsNull() || plan.VtepsIpVersion.IsUnknown()) {
		body["vteps-ip-version"] = plan.VtepsIpVersion.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/vxlan", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/vxlan failed", err.Error())
		return
	}
	interfaceVxlanApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceVxlanResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceVxlanModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/vxlan", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/vxlan failed", err.Error())
		return
	}
	interfaceVxlanApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceVxlanResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceVxlanModel
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
	if !plan.ARP.Equal(state.ARP) {
		body["arp"] = plan.ARP.ValueString()
	}
	if !plan.ARPTimeout.Equal(state.ARPTimeout) {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !plan.Bridge.Equal(state.Bridge) {
		body["bridge"] = plan.Bridge.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.LocalAddress.Equal(state.LocalAddress) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !plan.MACAddress.Equal(state.MACAddress) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.MTU.Equal(state.MTU) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Port.Equal(state.Port) {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.Ttl.Equal(state.Ttl) {
		body["ttl"] = plan.Ttl.ValueString()
	}
	if !plan.Vni.Equal(state.Vni) {
		body["vni"] = plan.Vni.ValueString()
	}
	if !plan.AllowFastPath.Equal(state.AllowFastPath) && !plan.AllowFastPath.IsUnknown() {
		body["allow-fast-path"] = plan.AllowFastPath.ValueString()
	}
	if !plan.BridgePvid.Equal(state.BridgePvid) && !plan.BridgePvid.IsUnknown() {
		body["bridge-pvid"] = plan.BridgePvid.ValueString()
	}
	if !plan.Checksum.Equal(state.Checksum) && !plan.Checksum.IsUnknown() {
		body["checksum"] = plan.Checksum.ValueString()
	}
	if !plan.DontFragment.Equal(state.DontFragment) && !plan.DontFragment.IsUnknown() {
		body["dont-fragment"] = plan.DontFragment.ValueString()
	}
	if !plan.Group.Equal(state.Group) && !plan.Group.IsUnknown() {
		body["group"] = plan.Group.ValueString()
	}
	if !plan.Hw.Equal(state.Hw) && !plan.Hw.IsUnknown() {
		body["hw"] = plan.Hw.ValueString()
	}
	if !plan.Learning.Equal(state.Learning) && !plan.Learning.IsUnknown() {
		body["learning"] = plan.Learning.ValueString()
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
	if !plan.MaxFdbSize.Equal(state.MaxFdbSize) && !plan.MaxFdbSize.IsUnknown() {
		body["max-fdb-size"] = plan.MaxFdbSize.ValueString()
	}
	if !plan.RemCsum.Equal(state.RemCsum) && !plan.RemCsum.IsUnknown() {
		body["rem-csum"] = plan.RemCsum.ValueString()
	}
	if !plan.VtepVrf.Equal(state.VtepVrf) && !plan.VtepVrf.IsUnknown() {
		body["vtep-vrf"] = plan.VtepVrf.ValueString()
	}
	if !plan.VtepsIpVersion.Equal(state.VtepsIpVersion) && !plan.VtepsIpVersion.IsUnknown() {
		body["vteps-ip-version"] = plan.VtepsIpVersion.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/vxlan", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/vxlan failed", err.Error())
			return
		}
		interfaceVxlanApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceVxlanResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceVxlanModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/vxlan", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/vxlan failed", err.Error())
	}
}

func (r *InterfaceVxlanResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceVxlanLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/vxlan matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceVxlanLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceVxlanLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/vxlan", id)
}

func interfaceVxlanApply(ctx context.Context, obj client.Object, m *InterfaceVxlanModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["vteps-ip-version"]; ok && v != "" {
		m.VtepsIpVersion = types.StringValue(v)
	} else {
		m.VtepsIpVersion = types.StringNull()
	}
	if v, ok := obj["vtep-vrf"]; ok && v != "" {
		m.VtepVrf = types.StringValue(v)
	} else {
		m.VtepVrf = types.StringNull()
	}
	if v, ok := obj["rem-csum"]; ok && v != "" {
		m.RemCsum = types.StringValue(v)
	} else {
		m.RemCsum = types.StringNull()
	}
	if v, ok := obj["max-fdb-size"]; ok && v != "" {
		m.MaxFdbSize = types.StringValue(v)
	} else {
		m.MaxFdbSize = types.StringNull()
	}
	if v, ok := obj["loop-protect-send-interval"]; ok && v != "" {
		m.LoopProtectSendInterval = types.StringValue(v)
	} else {
		m.LoopProtectSendInterval = types.StringNull()
	}
	if v, ok := obj["loop-protect-disable-time"]; ok && v != "" {
		m.LoopProtectDisableTime = types.StringValue(v)
	} else {
		m.LoopProtectDisableTime = types.StringNull()
	}
	if v, ok := obj["loop-protect"]; ok && v != "" {
		m.LoopProtect = types.StringValue(v)
	} else {
		m.LoopProtect = types.StringNull()
	}
	if v, ok := obj["learning"]; ok && v != "" {
		m.Learning = types.StringValue(v)
	} else {
		m.Learning = types.StringNull()
	}
	if v, ok := obj["hw"]; ok && v != "" {
		m.Hw = types.StringValue(v)
	} else {
		m.Hw = types.StringNull()
	}
	if v, ok := obj["group"]; ok && v != "" {
		m.Group = types.StringValue(v)
	} else {
		m.Group = types.StringNull()
	}
	if v, ok := obj["dont-fragment"]; ok && v != "" {
		m.DontFragment = types.StringValue(v)
	} else {
		m.DontFragment = types.StringNull()
	}
	if v, ok := obj["checksum"]; ok && v != "" {
		m.Checksum = types.StringValue(v)
	} else {
		m.Checksum = types.StringNull()
	}
	if v, ok := obj["bridge-pvid"]; ok && v != "" {
		m.BridgePvid = types.StringValue(v)
	} else {
		m.BridgePvid = types.StringNull()
	}
	if v, ok := obj["allow-fast-path"]; ok && v != "" {
		m.AllowFastPath = types.StringValue(v)
	} else {
		m.AllowFastPath = types.StringNull()
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
	if v, ok := obj["bridge"]; ok {
		_ = v
		if v != "" {
			m.Bridge = types.StringValue(v)
		} else {
			m.Bridge = types.StringNull()
		}
	} else {
		m.Bridge = types.StringNull()
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
	if v, ok := obj["interface"]; ok {
		_ = v
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	} else {
		m.Interface = types.StringNull()
	}
	if v, ok := obj["local-address"]; ok {
		_ = v
		if v != "" {
			m.LocalAddress = types.StringValue(v)
		} else {
			m.LocalAddress = types.StringNull()
		}
	} else {
		m.LocalAddress = types.StringNull()
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
	if v, ok := obj["mtu"]; ok {
		_ = v
		if v != "" {
			m.MTU = types.StringValue(v)
		} else {
			m.MTU = types.StringNull()
		}
	} else {
		m.MTU = types.StringNull()
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
	if v, ok := obj["port"]; ok {
		_ = v
		if v != "" {
			m.Port = types.StringValue(v)
		} else {
			m.Port = types.StringNull()
		}
	} else {
		m.Port = types.StringNull()
	}
	if v, ok := obj["ttl"]; ok {
		_ = v
		if v != "" {
			m.Ttl = types.StringValue(v)
		} else {
			m.Ttl = types.StringNull()
		}
	} else {
		m.Ttl = types.StringNull()
	}
	if v, ok := obj["vni"]; ok {
		_ = v
		if v != "" {
			m.Vni = types.StringValue(v)
		} else {
			m.Vni = types.StringNull()
		}
	} else {
		m.Vni = types.StringNull()
	}
}
