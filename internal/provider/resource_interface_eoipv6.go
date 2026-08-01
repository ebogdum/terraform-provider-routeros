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
	_ resource.Resource                = &InterfaceEoipv6Resource{}
	_ resource.ResourceWithImportState = &InterfaceEoipv6Resource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceEoipv6Resource struct {
	reg *client.Registry
}

type InterfaceEoipv6Model struct {
	ID                      types.String `tfsdk:"id"`
	LoopProtectSendInterval types.String `tfsdk:"loop_protect_send_interval"`
	LoopProtectDisableTime  types.String `tfsdk:"loop_protect_disable_time"`
	LoopProtect             types.String `tfsdk:"loop_protect"`
	Keepalive               types.String `tfsdk:"keepalive"`
	Dscp                    types.String `tfsdk:"dscp"`
	DontFragment            types.String `tfsdk:"dont_fragment"`
	ClampTcpMss             types.String `tfsdk:"clamp_tcp_mss"`
	ARP                     types.String `tfsdk:"arp"`
	ARPTimeout              types.String `tfsdk:"arp_timeout"`
	Comment                 types.String `tfsdk:"comment"`
	Disabled                types.Bool   `tfsdk:"disabled"`
	IpsecSecret             types.String `tfsdk:"ipsec_secret"`
	LocalAddress            types.String `tfsdk:"local_address"`
	MACAddress              types.String `tfsdk:"mac_address"`
	MTU                     types.String `tfsdk:"mtu"`
	Name                    types.String `tfsdk:"name"`
	RemoteAddress           types.String `tfsdk:"remote_address"`
	TunnelID                types.String `tfsdk:"tunnel_id"`
	Router                  types.String `tfsdk:"router"`
}

func NewInterfaceEoipv6Resource() resource.Resource { return &InterfaceEoipv6Resource{} }

func (r *InterfaceEoipv6Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_eoipv6"
}

func (r *InterfaceEoipv6Resource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceEoipv6Resource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/eoipv6`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"keepalive": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `keepalive`.",
			},
			"dscp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dscp`.",
			},
			"dont_fragment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dont-fragment`.",
			},
			"clamp_tcp_mss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `clamp-tcp-mss`.",
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
			"ipsec_secret": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
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
			"remote_address": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"tunnel_id": schema.StringAttribute{
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

func (r *InterfaceEoipv6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceEoipv6Model
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
	if !(plan.IpsecSecret.IsNull() || plan.IpsecSecret.IsUnknown()) {
		body["ipsec-secret"] = plan.IpsecSecret.ValueString()
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
	if !(plan.RemoteAddress.IsNull() || plan.RemoteAddress.IsUnknown()) {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !(plan.TunnelID.IsNull() || plan.TunnelID.IsUnknown()) {
		body["tunnel-id"] = plan.TunnelID.ValueString()
	}
	if !(plan.ClampTcpMss.IsNull() || plan.ClampTcpMss.IsUnknown()) {
		body["clamp-tcp-mss"] = plan.ClampTcpMss.ValueString()
	}
	if !(plan.DontFragment.IsNull() || plan.DontFragment.IsUnknown()) {
		body["dont-fragment"] = plan.DontFragment.ValueString()
	}
	if !(plan.Dscp.IsNull() || plan.Dscp.IsUnknown()) {
		body["dscp"] = plan.Dscp.ValueString()
	}
	if !(plan.Keepalive.IsNull() || plan.Keepalive.IsUnknown()) {
		body["keepalive"] = plan.Keepalive.ValueString()
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
	obj, err := c.Add(ctx, "/interface/eoipv6", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/eoipv6 failed", err.Error())
		return
	}
	interfaceEoipv6Apply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEoipv6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceEoipv6Model
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/eoipv6", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/eoipv6 failed", err.Error())
		return
	}
	interfaceEoipv6Apply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceEoipv6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceEoipv6Model
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
	if !plan.IpsecSecret.Equal(state.IpsecSecret) && !plan.IpsecSecret.IsUnknown() {
		body["ipsec-secret"] = plan.IpsecSecret.ValueString()
	}
	if !plan.LocalAddress.Equal(state.LocalAddress) && !plan.LocalAddress.IsUnknown() {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !plan.MACAddress.Equal(state.MACAddress) && !plan.MACAddress.IsUnknown() {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.MTU.Equal(state.MTU) && !plan.MTU.IsUnknown() {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.RemoteAddress.Equal(state.RemoteAddress) && !plan.RemoteAddress.IsUnknown() {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !plan.TunnelID.Equal(state.TunnelID) && !plan.TunnelID.IsUnknown() {
		body["tunnel-id"] = plan.TunnelID.ValueString()
	}
	if !plan.ClampTcpMss.Equal(state.ClampTcpMss) && !plan.ClampTcpMss.IsUnknown() {
		body["clamp-tcp-mss"] = plan.ClampTcpMss.ValueString()
	}
	if !plan.DontFragment.Equal(state.DontFragment) && !plan.DontFragment.IsUnknown() {
		body["dont-fragment"] = plan.DontFragment.ValueString()
	}
	if !plan.Dscp.Equal(state.Dscp) && !plan.Dscp.IsUnknown() {
		body["dscp"] = plan.Dscp.ValueString()
	}
	if !plan.Keepalive.Equal(state.Keepalive) && !plan.Keepalive.IsUnknown() {
		body["keepalive"] = plan.Keepalive.ValueString()
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
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/eoipv6", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/eoipv6 failed", err.Error())
			return
		}
		interfaceEoipv6Apply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEoipv6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceEoipv6Model
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/eoipv6", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/eoipv6 failed", err.Error())
	}
}

func (r *InterfaceEoipv6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceEoipv6LookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/eoipv6 matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceEoipv6LookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceEoipv6LookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/eoipv6", id)
}

func interfaceEoipv6Apply(ctx context.Context, obj client.Object, m *InterfaceEoipv6Model) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
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
	if v, ok := obj["keepalive"]; ok && v != "" {
		m.Keepalive = types.StringValue(v)
	} else {
		m.Keepalive = types.StringNull()
	}
	if v, ok := obj["dscp"]; ok && v != "" {
		m.Dscp = types.StringValue(v)
	} else {
		m.Dscp = types.StringNull()
	}
	if v, ok := obj["dont-fragment"]; ok && v != "" {
		m.DontFragment = types.StringValue(v)
	} else {
		m.DontFragment = types.StringNull()
	}
	if v, ok := obj["clamp-tcp-mss"]; ok && v != "" {
		m.ClampTcpMss = types.StringValue(v)
	} else {
		m.ClampTcpMss = types.StringNull()
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
	if v, ok := obj["ipsec-secret"]; ok {
		if v != "" {
			m.IpsecSecret = types.StringValue(v)
		} else {
			m.IpsecSecret = types.StringNull()
		}
	}
	if v, ok := obj["local-address"]; ok {
		if v != "" {
			m.LocalAddress = types.StringValue(v)
		} else {
			m.LocalAddress = types.StringNull()
		}
	}
	if v, ok := obj["mac-address"]; ok {
		if v != "" {
			m.MACAddress = types.StringValue(v)
		} else {
			m.MACAddress = types.StringNull()
		}
	}
	if v, ok := obj["mtu"]; ok {
		if v != "" {
			m.MTU = types.StringValue(v)
		} else {
			m.MTU = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["remote-address"]; ok {
		if v != "" {
			m.RemoteAddress = types.StringValue(v)
		} else {
			m.RemoteAddress = types.StringNull()
		}
	}
	if v, ok := obj["tunnel-id"]; ok {
		if v != "" {
			m.TunnelID = types.StringValue(v)
		} else {
			m.TunnelID = types.StringNull()
		}
	}
}
