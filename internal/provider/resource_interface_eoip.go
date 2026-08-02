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
	_ resource.Resource                = &InterfaceEoipResource{}
	_ resource.ResourceWithImportState = &InterfaceEoipResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceEoipResource struct {
	reg *client.Registry
}

type InterfaceEoipModel struct {
	ID                      types.String `tfsdk:"id"`
	ActualMTU               types.Int64  `tfsdk:"actual_mtu"`
	AllowFastPath           types.Bool   `tfsdk:"allow_fast_path"`
	ARP                     types.String `tfsdk:"arp"`
	ARPTimeout              types.String `tfsdk:"arp_timeout"`
	ClampTCPMss             types.Bool   `tfsdk:"clamp_tcp_mss"`
	Comment                 types.String `tfsdk:"comment"`
	DisableTime             types.String `tfsdk:"disable_time"`
	Disabled                types.Bool   `tfsdk:"disabled"`
	DontFragment            types.String `tfsdk:"dont_fragment"`
	Dscp                    types.String `tfsdk:"dscp"`
	IpsecSecret             types.String `tfsdk:"ipsec_secret"`
	Keepalive               types.String `tfsdk:"keepalive"`
	LocalAddress            types.String `tfsdk:"local_address"`
	LoopProtect             types.String `tfsdk:"loop_protect"`
	LoopProtectDisableTime  types.String `tfsdk:"loop_protect_disable_time"`
	LoopProtectSendInterval types.String `tfsdk:"loop_protect_send_interval"`
	MACAddress              types.String `tfsdk:"mac_address"`
	MTU                     types.String `tfsdk:"mtu"`
	Name                    types.String `tfsdk:"name"`
	RemoteAddress           types.String `tfsdk:"remote_address"`
	SendInterval            types.String `tfsdk:"send_interval"`
	Status                  types.String `tfsdk:"status"`
	TunnelID                types.Int64  `tfsdk:"tunnel_id"`
	Router                  types.String `tfsdk:"router"`
}

func NewInterfaceEoipResource() resource.Resource { return &InterfaceEoipResource{} }

func (r *InterfaceEoipResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_eoip"
}

func (r *InterfaceEoipResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceEoipResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/eoip`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"actual_mtu": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"allow_fast_path": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether to allow FastPath processing. Must be disabled if IPsec tunneling is used.",
			},
			"arp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Address Resolution Protocol mode. disabled - the interface will not use ARP enabled - the interface will use ARP proxy-arp - the interface will use the ARP proxy feature reply-only - the interface will only reply to requests originated from matching IP address/MAC address combinations which are entered as static entries in the \"/ip arp\" table. No dynamic entries will be automatically stored in the \"/ip arp\" table. Therefore for communications to be successful, a valid static entry must already exist.",
				Validators:  []validator.String{schemautil.OneOf([]string{"disabled", "enabled", "proxy-arp", "reply-only", "local-proxy-arp"}...)},
			},
			"arp_timeout": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Time interval in which ARP entries should time out.",
				Validators:    []validator.String{schemautil.IsDurationOrKeyword("auto")},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"clamp_tcp_mss": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP overhead).The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is changed.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"disable_time": schema.StringAttribute{
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"dont_fragment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether to include DF bit in related packets: no \u00a0 - fragment if needed, \u00a0 inherit \u00a0 - use Dont Fragment flag of original packet. (Without Dont Fragment: inherit - packet may be fragmented).",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "inherit"}...)},
			},
			"dscp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "DSCP value of packet. Inherited option means that dscp value will be inherited from packet which is going to be encapsulated.",
				Validators:  []validator.String{schemautil.OneOf([]string{"inherit"}...)},
			},
			"ipsec_secret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "When secret is specified, router adds dynamic IPsec peer to remote-address with pre-shared key and policy (by default phase2 uses sha1/aes128cbc).",
			},
			"keepalive": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries - number of retry attempts. By default keepalive is set to 10 seconds and 10 retries.",
			},
			"local_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source address of the tunnel packets, local on the router.",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"loop_protect": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"default", "off", "on"}...)},
			},
			"loop_protect_disable_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"loop_protect_send_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Media Access Control number of an interface. The address numeration authority IANA allows the use of MAC addresses in the range from 00:00:5E:80:00:00 - 00:00:5E:FF:FF:FF freely",
			},
			"mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Layer3 Maximum transmission unit A number, or `auto`.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Interface name",
			},
			"remote_address": schema.StringAttribute{
				Required:    true,
				Description: "IP address of remote end of EoIP tunnel",
			},
			"send_interval": schema.StringAttribute{
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "off", "on", "disabled"}...)},
			},
			"tunnel_id": schema.Int64Attribute{
				Required:    true,
				Description: "Unique tunnel identifier, which must match other side of the tunnel",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceEoipResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceEoipModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AllowFastPath.IsNull() || plan.AllowFastPath.IsUnknown()) {
		body["allow-fast-path"] = client.FormatBool(plan.AllowFastPath.ValueBool())
	}
	if !(plan.ARP.IsNull() || plan.ARP.IsUnknown()) {
		body["arp"] = plan.ARP.ValueString()
	}
	if !(plan.ARPTimeout.IsNull() || plan.ARPTimeout.IsUnknown()) {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !(plan.ClampTCPMss.IsNull() || plan.ClampTCPMss.IsUnknown()) {
		body["clamp-tcp-mss"] = client.FormatBool(plan.ClampTCPMss.ValueBool())
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.DontFragment.IsNull() || plan.DontFragment.IsUnknown()) {
		body["dont-fragment"] = plan.DontFragment.ValueString()
	}
	if !(plan.Dscp.IsNull() || plan.Dscp.IsUnknown()) {
		body["dscp"] = plan.Dscp.ValueString()
	}
	if !(plan.IpsecSecret.IsNull() || plan.IpsecSecret.IsUnknown()) {
		body["ipsec-secret"] = plan.IpsecSecret.ValueString()
	}
	if !(plan.Keepalive.IsNull() || plan.Keepalive.IsUnknown()) {
		body["keepalive"] = plan.Keepalive.ValueString()
	}
	if !(plan.LocalAddress.IsNull() || plan.LocalAddress.IsUnknown()) {
		body["local-address"] = plan.LocalAddress.ValueString()
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
		body["tunnel-id"] = client.FormatInt64(plan.TunnelID.ValueInt64())
	}
	obj, err := c.Add(ctx, "/interface/eoip", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/eoip failed", err.Error())
		return
	}
	interfaceEoipApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEoipResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceEoipModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/eoip", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/eoip failed", err.Error())
		return
	}
	interfaceEoipApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceEoipResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceEoipModel
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
	if !plan.AllowFastPath.Equal(state.AllowFastPath) && !plan.AllowFastPath.IsUnknown() {
		body["allow-fast-path"] = client.FormatBool(plan.AllowFastPath.ValueBool())
	}
	if !plan.ARP.Equal(state.ARP) && !plan.ARP.IsUnknown() {
		body["arp"] = plan.ARP.ValueString()
	}
	if !plan.ARPTimeout.Equal(state.ARPTimeout) && !plan.ARPTimeout.IsUnknown() {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !plan.ClampTCPMss.Equal(state.ClampTCPMss) && !plan.ClampTCPMss.IsUnknown() {
		body["clamp-tcp-mss"] = client.FormatBool(plan.ClampTCPMss.ValueBool())
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DontFragment.Equal(state.DontFragment) && !plan.DontFragment.IsUnknown() {
		body["dont-fragment"] = plan.DontFragment.ValueString()
	}
	if !plan.Dscp.Equal(state.Dscp) && !plan.Dscp.IsUnknown() {
		body["dscp"] = plan.Dscp.ValueString()
	}
	if !plan.IpsecSecret.Equal(state.IpsecSecret) && !plan.IpsecSecret.IsUnknown() {
		body["ipsec-secret"] = plan.IpsecSecret.ValueString()
	}
	if !plan.Keepalive.Equal(state.Keepalive) && !plan.Keepalive.IsUnknown() {
		body["keepalive"] = plan.Keepalive.ValueString()
	}
	if !plan.LocalAddress.Equal(state.LocalAddress) && !plan.LocalAddress.IsUnknown() {
		body["local-address"] = plan.LocalAddress.ValueString()
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
		body["tunnel-id"] = client.FormatInt64(plan.TunnelID.ValueInt64())
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/eoip", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/eoip failed", err.Error())
			return
		}
		interfaceEoipApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEoipResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceEoipModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/eoip", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/eoip failed", err.Error())
	}
}

func (r *InterfaceEoipResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceEoipLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/eoip matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceEoipLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceEoipLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/eoip", id)
}

func interfaceEoipApply(ctx context.Context, obj client.Object, m *InterfaceEoipModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["actual-mtu"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.ActualMTU = types.Int64Value(n)
		} else {
			m.ActualMTU = types.Int64Null()
		}
	} else {
		m.ActualMTU = types.Int64Null()
	}
	if v, ok := obj["allow-fast-path"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.AllowFastPath = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.AllowFastPath = types.BoolValue(true)
		} else {
			m.AllowFastPath = types.BoolNull()
		}
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
	if v, ok := obj["clamp-tcp-mss"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.ClampTCPMss = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.ClampTCPMss = types.BoolValue(true)
		} else {
			m.ClampTCPMss = types.BoolNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["disable-time"]; ok {
		if v != "" {
			m.DisableTime = types.StringValue(v)
		} else {
			m.DisableTime = types.StringNull()
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
	if v, ok := obj["dont-fragment"]; ok {
		if v != "" {
			m.DontFragment = types.StringValue(v)
		} else {
			m.DontFragment = types.StringNull()
		}
	}
	if v, ok := obj["dscp"]; ok {
		if v != "" {
			m.Dscp = types.StringValue(v)
		} else {
			m.Dscp = types.StringNull()
		}
	}
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.IpsecSecret already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["ipsec-secret"]; ok && v != "" {
		_ = v
		if v != "" {
			m.IpsecSecret = types.StringValue(v)
		} else {
			m.IpsecSecret = types.StringNull()
		}
	} else if m.IpsecSecret.IsUnknown() {
		m.IpsecSecret = types.StringNull()
	}
	if v, ok := obj["keepalive"]; ok {
		if v != "" {
			m.Keepalive = types.StringValue(v)
		} else {
			m.Keepalive = types.StringNull()
		}
	}
	if v, ok := obj["local-address"]; ok {
		if v != "" {
			m.LocalAddress = types.StringValue(v)
		} else {
			m.LocalAddress = types.StringNull()
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
	if v, ok := obj["send-interval"]; ok {
		if v != "" {
			m.SendInterval = types.StringValue(v)
		} else {
			m.SendInterval = types.StringNull()
		}
	}
	if v, ok := obj["status"]; ok {
		if v != "" {
			m.Status = types.StringValue(v)
		} else {
			m.Status = types.StringNull()
		}
	}
	if v, ok := obj["tunnel-id"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TunnelID = types.Int64Value(n)
		} else {
			m.TunnelID = types.Int64Null()
		}
	} else {
		m.TunnelID = types.Int64Null()
	}
}
