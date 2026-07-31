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
	_ resource.Resource                = &InterfaceGreResource{}
	_ resource.ResourceWithImportState = &InterfaceGreResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceGreResource struct {
	reg *client.Registry
}

type InterfaceGreModel struct {
	ID            types.String `tfsdk:"id"`
	ActualMTU     types.Int64  `tfsdk:"actual_mtu"`
	AllowFastPath types.Bool   `tfsdk:"allow_fast_path"`
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

func NewInterfaceGreResource() resource.Resource { return &InterfaceGreResource{} }

func (r *InterfaceGreResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_gre"
}

func (r *InterfaceGreResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceGreResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "GRE tunnel — needs reachable remote address and unused name. Skipped.",
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
			"clamp_tcp_mss": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP overhead). The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is changed.",
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
				Description: "Set dscp value in Gre header to a fixed value or inherit from dscp value taken from tunnelled traffic",
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
				Description: "IP address that will be used for local tunnel end. If set to 0.0.0.0 then IP address of outgoing interface will be used.",
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
				Description: "Name of the tunnel.",
			},
			"remote_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IP address of remote tunnel end.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceGreResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceGreModel
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
	obj, err := c.Add(ctx, "/interface/gre", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/gre failed", err.Error())
		return
	}
	interfaceGreApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceGreResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceGreModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/gre", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/gre failed", err.Error())
		return
	}
	interfaceGreApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceGreResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceGreModel
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
	if !plan.AllowFastPath.Equal(state.AllowFastPath) {
		body["allow-fast-path"] = client.FormatBool(plan.AllowFastPath.ValueBool())
	}
	if !plan.ClampTCPMss.Equal(state.ClampTCPMss) {
		body["clamp-tcp-mss"] = client.FormatBool(plan.ClampTCPMss.ValueBool())
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DontFragment.Equal(state.DontFragment) {
		body["dont-fragment"] = plan.DontFragment.ValueString()
	}
	if !plan.Dscp.Equal(state.Dscp) {
		body["dscp"] = plan.Dscp.ValueString()
	}
	if !plan.IpsecSecret.Equal(state.IpsecSecret) {
		body["ipsec-secret"] = plan.IpsecSecret.ValueString()
	}
	if !plan.Keepalive.Equal(state.Keepalive) {
		body["keepalive"] = plan.Keepalive.ValueString()
	}
	if !plan.LocalAddress.Equal(state.LocalAddress) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !plan.MTU.Equal(state.MTU) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.RemoteAddress.Equal(state.RemoteAddress) {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/gre", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/gre failed", err.Error())
			return
		}
		interfaceGreApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceGreResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceGreModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/gre", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/gre failed", err.Error())
	}
}

func (r *InterfaceGreResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceGreLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/gre matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceGreLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceGreLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/gre", id)
}

func interfaceGreApply(ctx context.Context, obj client.Object, m *InterfaceGreModel) {
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
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AllowFastPath = types.BoolValue(b)
		} else {
			m.AllowFastPath = types.BoolNull()
		}
	} else {
		m.AllowFastPath = types.BoolNull()
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
