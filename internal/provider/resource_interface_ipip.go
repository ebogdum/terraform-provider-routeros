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
	_ resource.Resource                = &InterfaceIpipResource{}
	_ resource.ResourceWithImportState = &InterfaceIpipResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceIpipResource struct {
	reg *client.Registry
}

type InterfaceIpipModel struct {
	ID            types.String `tfsdk:"id"`
	AllowFastPath types.String `tfsdk:"allow_fast_path"`
	ClampTCPMss   types.Bool   `tfsdk:"clamp_tcp_mss"`
	Comment       types.String `tfsdk:"comment"`
	Disabled      types.Bool   `tfsdk:"disabled"`
	DontFragment  types.String `tfsdk:"dont_fragment"`
	Dscp          types.String `tfsdk:"dscp"`
	IpsecSecret   types.String `tfsdk:"ipsec_secret"`
	Keepalive     types.String `tfsdk:"keepalive"`
	LocalAddress  types.String `tfsdk:"local_address"`
	MTU           types.String `tfsdk:"mtu"`
	Name          types.String `tfsdk:"name"`
	RemoteAddress types.String `tfsdk:"remote_address"`
	Router        types.String `tfsdk:"router"`
}

func NewInterfaceIpipResource() resource.Resource { return &InterfaceIpipResource{} }

func (r *InterfaceIpipResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ipip"
}

func (r *InterfaceIpipResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceIpipResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/ipip`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"allow_fast_path": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `allow-fast-path`.",
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
				Description: "Set dscp value in IPIP header to a fixed value or inherit from dscp value taken from tunnelled traffic.",
				Validators:  []validator.String{schemautil.OneOf([]string{"inherit"}...)},
			},
			"ipsec_secret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "When secret is specified, router adds dynamic ipsec peer to remote-address with pre-shared key and policy with default values (by default phase2 uses sha1/aes128cbc).",
			},
			"keepalive": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries - number of retry attempts. By default keepalive is set to 10 seconds and 10 retries. To disable set *set ipipv6-tunnel1 !keepalive.",
			},
			"local_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address on a router that will be used by IPIP tunnel.",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Layer3 Maximum transmission unit. A number, or `auto`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interface name.",
			},
			"remote_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of remote end of IPIP tunnel.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceIpipResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceIpipModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
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
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.RemoteAddress.IsNull() || plan.RemoteAddress.IsUnknown()) {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !(plan.AllowFastPath.IsNull() || plan.AllowFastPath.IsUnknown()) {
		body["allow-fast-path"] = plan.AllowFastPath.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/ipip", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/ipip failed", err.Error())
		return
	}
	interfaceIpipApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceIpipResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceIpipModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/ipip", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/ipip failed", err.Error())
		return
	}
	interfaceIpipApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceIpipResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceIpipModel
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
	if !plan.MTU.Equal(state.MTU) && !plan.MTU.IsUnknown() {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.RemoteAddress.Equal(state.RemoteAddress) && !plan.RemoteAddress.IsUnknown() {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !plan.AllowFastPath.Equal(state.AllowFastPath) && !plan.AllowFastPath.IsUnknown() {
		body["allow-fast-path"] = plan.AllowFastPath.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/ipip", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/ipip failed", err.Error())
			return
		}
		interfaceIpipApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceIpipResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceIpipModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/ipip", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/ipip failed", err.Error())
	}
}

func (r *InterfaceIpipResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceIpipLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/ipip matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceIpipLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceIpipLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/ipip", id)
}

func interfaceIpipApply(ctx context.Context, obj client.Object, m *InterfaceIpipModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["allow-fast-path"]; ok && v != "" {
		m.AllowFastPath = types.StringValue(v)
	} else {
		m.AllowFastPath = types.StringNull()
	}
	if v, ok := obj["clamp-tcp-mss"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.ClampTCPMss = types.BoolValue(b)
		} else {
			m.ClampTCPMss = types.BoolNull()
		}
	} else {
		m.ClampTCPMss = types.BoolNull()
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
	if v, ok := obj["dont-fragment"]; ok {
		_ = v
		if v != "" {
			m.DontFragment = types.StringValue(v)
		} else {
			m.DontFragment = types.StringNull()
		}
	} else {
		m.DontFragment = types.StringNull()
	}
	if v, ok := obj["dscp"]; ok {
		_ = v
		if v != "" {
			m.Dscp = types.StringValue(v)
		} else {
			m.Dscp = types.StringNull()
		}
	} else {
		m.Dscp = types.StringNull()
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
		_ = v
		if v != "" {
			m.Keepalive = types.StringValue(v)
		} else {
			m.Keepalive = types.StringNull()
		}
	} else {
		m.Keepalive = types.StringNull()
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
	if v, ok := obj["remote-address"]; ok {
		_ = v
		if v != "" {
			m.RemoteAddress = types.StringValue(v)
		} else {
			m.RemoteAddress = types.StringNull()
		}
	} else {
		m.RemoteAddress = types.StringNull()
	}
}
