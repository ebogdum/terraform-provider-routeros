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
	_ resource.Resource                = &IPDHCPClientResource{}
	_ resource.ResourceWithImportState = &IPDHCPClientResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPDHCPClientResource struct {
	reg *client.Registry
}

type IPDHCPClientModel struct {
	ID                       types.String `tfsdk:"id"`
	AddDefaultRoute          types.String `tfsdk:"add_default_route"`
	Address                  types.String `tfsdk:"address"`
	AllowReconfigure         types.Bool   `tfsdk:"allow_reconfigure"`
	AllowReconfigureMessages types.Bool   `tfsdk:"allow_reconfigure_messages"`
	CapsManagers             types.String `tfsdk:"caps_managers"`
	CheckGateway             types.String `tfsdk:"check_gateway"`
	Comment                  types.String `tfsdk:"comment"`
	DefaultRouteDistance     types.Int64  `tfsdk:"default_route_distance"`
	DefaultRouteTables       types.String `tfsdk:"default_route_tables"`
	DHCPOptions              types.List   `tfsdk:"dhcp_options"`
	DHCPServer               types.String `tfsdk:"dhcp_server"`
	Disabled                 types.Bool   `tfsdk:"disabled"`
	Dscp                     types.Int64  `tfsdk:"dscp"`
	Dynamic                  types.Bool   `tfsdk:"dynamic"`
	ExpiresAfter             types.String `tfsdk:"expires_after"`
	Gateway                  types.String `tfsdk:"gateway"`
	Interface                types.String `tfsdk:"interface"`
	Invalid                  types.Bool   `tfsdk:"invalid"`
	IPAddress                types.String `tfsdk:"ip_address"`
	LastReceivedCounter      types.String `tfsdk:"last_received_counter"`
	Name                     types.String `tfsdk:"name"`
	PrimaryDNS               types.String `tfsdk:"primary_dns"`
	PrimaryNTP               types.String `tfsdk:"primary_ntp"`
	ReconfigureKey           types.String `tfsdk:"reconfigure_key"`
	Release                  types.String `tfsdk:"release"`
	Renew                    types.String `tfsdk:"renew"`
	Route                    types.String `tfsdk:"route"`
	RoutingTables            types.String `tfsdk:"routing_tables"`
	Script                   types.String `tfsdk:"script"`
	SecondaryDNS             types.String `tfsdk:"secondary_dns"`
	SecondaryNTP             types.String `tfsdk:"secondary_ntp"`
	Status                   types.String `tfsdk:"status"`
	UseBroadcast             types.String `tfsdk:"use_broadcast"`
	UsePeerDNS               types.Bool   `tfsdk:"use_peer_dns"`
	UsePeerNTP               types.Bool   `tfsdk:"use_peer_ntp"`
	VLANPriority             types.Int64  `tfsdk:"vlan_priority"`
	Router                   types.String `tfsdk:"router"`
}

func NewIPDHCPClientResource() resource.Resource { return &IPDHCPClientResource{} }

func (r *IPDHCPClientResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_dhcp_client"
}

func (r *IPDHCPClientResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPDHCPClientResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "DHCP client per interface. On most devices the default config already binds one to an interface, causing \"dhcp-client on that interface already\" on a fresh add. Skipped.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"add_default_route": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "yes", "special-classless"}...)},
			},
			"address": schema.StringAttribute{
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsCIDR()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeCIDR()},
			},
			"allow_reconfigure": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"allow_reconfigure_messages": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"caps_managers": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"check_gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"none", "arp", "ping", "bfd"}...)},
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"default_route_distance": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"default_route_tables": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dhcp_options": schema.ListAttribute{
				Optional:    true,
				Computed:    true,
				ElementType: types.StringType,
				Description: "",
			},
			"dhcp_server": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dscp": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"expires_after": schema.StringAttribute{
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"gateway": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"ip_address": schema.StringAttribute{
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsCIDR()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeCIDR()},
			},
			"last_received_counter": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"primary_dns": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"primary_ntp": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"reconfigure_key": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"release": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"renew": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"route": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"routing_tables": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"script": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"secondary_dns": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"secondary_ntp": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"status": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"use_broadcast": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"both", "always", "never"}...)},
			},
			"use_peer_dns": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"use_peer_ntp": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vlan_priority": schema.Int64Attribute{
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

func (r *IPDHCPClientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPDHCPClientModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AddDefaultRoute.IsNull() || plan.AddDefaultRoute.IsUnknown()) {
		body["add-default-route"] = plan.AddDefaultRoute.ValueString()
	}
	if !(plan.AllowReconfigure.IsNull() || plan.AllowReconfigure.IsUnknown()) {
		body["allow-reconfigure"] = client.FormatBool(plan.AllowReconfigure.ValueBool())
	}
	if !(plan.CheckGateway.IsNull() || plan.CheckGateway.IsUnknown()) {
		body["check-gateway"] = plan.CheckGateway.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.DefaultRouteDistance.IsNull() || plan.DefaultRouteDistance.IsUnknown()) {
		body["default-route-distance"] = client.FormatInt64(plan.DefaultRouteDistance.ValueInt64())
	}
	if !(plan.DefaultRouteTables.IsNull() || plan.DefaultRouteTables.IsUnknown()) {
		body["default-route-tables"] = plan.DefaultRouteTables.ValueString()
	}
	if !(plan.DHCPOptions.IsNull() || plan.DHCPOptions.IsUnknown()) {
		body["dhcp-options"] = encodeStringList(ctx, plan.DHCPOptions, &resp.Diagnostics)
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Dscp.IsNull() || plan.Dscp.IsUnknown()) {
		body["dscp"] = client.FormatInt64(plan.Dscp.ValueInt64())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Script.IsNull() || plan.Script.IsUnknown()) {
		body["script"] = plan.Script.ValueString()
	}
	if !(plan.UseBroadcast.IsNull() || plan.UseBroadcast.IsUnknown()) {
		body["use-broadcast"] = plan.UseBroadcast.ValueString()
	}
	if !(plan.UsePeerDNS.IsNull() || plan.UsePeerDNS.IsUnknown()) {
		body["use-peer-dns"] = client.FormatBool(plan.UsePeerDNS.ValueBool())
	}
	if !(plan.UsePeerNTP.IsNull() || plan.UsePeerNTP.IsUnknown()) {
		body["use-peer-ntp"] = client.FormatBool(plan.UsePeerNTP.ValueBool())
	}
	if !(plan.VLANPriority.IsNull() || plan.VLANPriority.IsUnknown()) {
		body["vlan-priority"] = client.FormatInt64(plan.VLANPriority.ValueInt64())
	}
	obj, err := c.Add(ctx, "/ip/dhcp-client", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/dhcp-client failed", err.Error())
		return
	}
	iPDHCPClientApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPClientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPDHCPClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/dhcp-client", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/dhcp-client failed", err.Error())
		return
	}
	iPDHCPClientApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPDHCPClientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPDHCPClientModel
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
	if !plan.AddDefaultRoute.Equal(state.AddDefaultRoute) && !plan.AddDefaultRoute.IsUnknown() {
		body["add-default-route"] = plan.AddDefaultRoute.ValueString()
	}
	if !plan.AllowReconfigure.Equal(state.AllowReconfigure) && !plan.AllowReconfigure.IsUnknown() {
		body["allow-reconfigure"] = client.FormatBool(plan.AllowReconfigure.ValueBool())
	}
	if !plan.CheckGateway.Equal(state.CheckGateway) && !plan.CheckGateway.IsUnknown() {
		body["check-gateway"] = plan.CheckGateway.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.DefaultRouteDistance.Equal(state.DefaultRouteDistance) && !plan.DefaultRouteDistance.IsUnknown() {
		body["default-route-distance"] = client.FormatInt64(plan.DefaultRouteDistance.ValueInt64())
	}
	if !plan.DefaultRouteTables.Equal(state.DefaultRouteTables) && !plan.DefaultRouteTables.IsUnknown() {
		body["default-route-tables"] = plan.DefaultRouteTables.ValueString()
	}
	if !plan.DHCPOptions.Equal(state.DHCPOptions) && !plan.DHCPOptions.IsUnknown() {
		body["dhcp-options"] = encodeStringList(ctx, plan.DHCPOptions, &resp.Diagnostics)
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Dscp.Equal(state.Dscp) && !plan.Dscp.IsUnknown() {
		body["dscp"] = client.FormatInt64(plan.Dscp.ValueInt64())
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Script.Equal(state.Script) && !plan.Script.IsUnknown() {
		body["script"] = plan.Script.ValueString()
	}
	if !plan.UseBroadcast.Equal(state.UseBroadcast) && !plan.UseBroadcast.IsUnknown() {
		body["use-broadcast"] = plan.UseBroadcast.ValueString()
	}
	if !plan.UsePeerDNS.Equal(state.UsePeerDNS) && !plan.UsePeerDNS.IsUnknown() {
		body["use-peer-dns"] = client.FormatBool(plan.UsePeerDNS.ValueBool())
	}
	if !plan.UsePeerNTP.Equal(state.UsePeerNTP) && !plan.UsePeerNTP.IsUnknown() {
		body["use-peer-ntp"] = client.FormatBool(plan.UsePeerNTP.ValueBool())
	}
	if !plan.VLANPriority.Equal(state.VLANPriority) && !plan.VLANPriority.IsUnknown() {
		body["vlan-priority"] = client.FormatInt64(plan.VLANPriority.ValueInt64())
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/dhcp-client", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/dhcp-client failed", err.Error())
			return
		}
		iPDHCPClientApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPClientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPDHCPClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/dhcp-client", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/dhcp-client failed", err.Error())
	}
}

func (r *IPDHCPClientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPDHCPClientLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/dhcp-client matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPDHCPClientLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPDHCPClientLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/dhcp-client", id)
}

func iPDHCPClientApply(ctx context.Context, obj client.Object, m *IPDHCPClientModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["add-default-route"]; ok {
		_ = v
		if v != "" {
			m.AddDefaultRoute = types.StringValue(v)
		} else {
			m.AddDefaultRoute = types.StringNull()
		}
	} else {
		m.AddDefaultRoute = types.StringNull()
	}
	if v, ok := obj["address"]; ok {
		_ = v
		if v != "" {
			m.Address = types.StringValue(v)
		} else {
			m.Address = types.StringNull()
		}
	} else {
		m.Address = types.StringNull()
	}
	if v, ok := obj["allow-reconfigure"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AllowReconfigure = types.BoolValue(b)
		} else {
			m.AllowReconfigure = types.BoolNull()
		}
	} else {
		m.AllowReconfigure = types.BoolNull()
	}
	if v, ok := obj["allow-reconfigure-messages"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AllowReconfigureMessages = types.BoolValue(b)
		} else {
			m.AllowReconfigureMessages = types.BoolNull()
		}
	} else {
		m.AllowReconfigureMessages = types.BoolNull()
	}
	if v, ok := obj["caps-managers"]; ok {
		_ = v
		if v != "" {
			m.CapsManagers = types.StringValue(v)
		} else {
			m.CapsManagers = types.StringNull()
		}
	} else {
		m.CapsManagers = types.StringNull()
	}
	if v, ok := obj["check-gateway"]; ok {
		_ = v
		if v != "" {
			m.CheckGateway = types.StringValue(v)
		} else {
			m.CheckGateway = types.StringNull()
		}
	} else {
		m.CheckGateway = types.StringNull()
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
	if v, ok := obj["default-route-distance"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.DefaultRouteDistance = types.Int64Value(n)
		} else {
			m.DefaultRouteDistance = types.Int64Null()
		}
	} else {
		m.DefaultRouteDistance = types.Int64Null()
	}
	if v, ok := obj["default-route-tables"]; ok {
		_ = v
		if v != "" {
			m.DefaultRouteTables = types.StringValue(v)
		} else {
			m.DefaultRouteTables = types.StringNull()
		}
	} else {
		m.DefaultRouteTables = types.StringNull()
	}
	if v, ok := obj["dhcp-options"]; ok {
		_ = v
		m.DHCPOptions = decodeStringList(ctx, v)
	} else {
		m.DHCPOptions = types.ListNull(types.StringType)
	}
	if v, ok := obj["dhcp-server"]; ok {
		_ = v
		if v != "" {
			m.DHCPServer = types.StringValue(v)
		} else {
			m.DHCPServer = types.StringNull()
		}
	} else {
		m.DHCPServer = types.StringNull()
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
	if v, ok := obj["dscp"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Dscp = types.Int64Value(n)
		} else {
			m.Dscp = types.Int64Null()
		}
	} else {
		m.Dscp = types.Int64Null()
	}
	if v, ok := obj["dynamic"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else {
			m.Dynamic = types.BoolNull()
		}
	} else {
		m.Dynamic = types.BoolNull()
	}
	if v, ok := obj["expires-after"]; ok {
		_ = v
		if v != "" {
			m.ExpiresAfter = types.StringValue(v)
		} else {
			m.ExpiresAfter = types.StringNull()
		}
	} else {
		m.ExpiresAfter = types.StringNull()
	}
	if v, ok := obj["gateway"]; ok {
		_ = v
		if v != "" {
			m.Gateway = types.StringValue(v)
		} else {
			m.Gateway = types.StringNull()
		}
	} else {
		m.Gateway = types.StringNull()
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
	if v, ok := obj["invalid"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Invalid = types.BoolValue(b)
		} else {
			m.Invalid = types.BoolNull()
		}
	} else {
		m.Invalid = types.BoolNull()
	}
	if v, ok := obj["ip-address"]; ok {
		_ = v
		if v != "" {
			m.IPAddress = types.StringValue(v)
		} else {
			m.IPAddress = types.StringNull()
		}
	} else {
		m.IPAddress = types.StringNull()
	}
	if v, ok := obj["last-received-counter"]; ok {
		_ = v
		if v != "" {
			m.LastReceivedCounter = types.StringValue(v)
		} else {
			m.LastReceivedCounter = types.StringNull()
		}
	} else {
		m.LastReceivedCounter = types.StringNull()
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
	if v, ok := obj["primary-dns"]; ok {
		_ = v
		if v != "" {
			m.PrimaryDNS = types.StringValue(v)
		} else {
			m.PrimaryDNS = types.StringNull()
		}
	} else {
		m.PrimaryDNS = types.StringNull()
	}
	if v, ok := obj["primary-ntp"]; ok {
		_ = v
		if v != "" {
			m.PrimaryNTP = types.StringValue(v)
		} else {
			m.PrimaryNTP = types.StringNull()
		}
	} else {
		m.PrimaryNTP = types.StringNull()
	}
	if v, ok := obj["reconfigure-key"]; ok {
		_ = v
		if v != "" {
			m.ReconfigureKey = types.StringValue(v)
		} else {
			m.ReconfigureKey = types.StringNull()
		}
	} else {
		m.ReconfigureKey = types.StringNull()
	}
	if v, ok := obj["release"]; ok {
		_ = v
		if v != "" {
			m.Release = types.StringValue(v)
		} else {
			m.Release = types.StringNull()
		}
	} else {
		m.Release = types.StringNull()
	}
	if v, ok := obj["renew"]; ok {
		_ = v
		if v != "" {
			m.Renew = types.StringValue(v)
		} else {
			m.Renew = types.StringNull()
		}
	} else {
		m.Renew = types.StringNull()
	}
	if v, ok := obj["route"]; ok {
		_ = v
		if v != "" {
			m.Route = types.StringValue(v)
		} else {
			m.Route = types.StringNull()
		}
	} else {
		m.Route = types.StringNull()
	}
	if v, ok := obj["routing-tables"]; ok {
		_ = v
		if v != "" {
			m.RoutingTables = types.StringValue(v)
		} else {
			m.RoutingTables = types.StringNull()
		}
	} else {
		m.RoutingTables = types.StringNull()
	}
	if v, ok := obj["script"]; ok {
		_ = v
		if v != "" {
			m.Script = types.StringValue(v)
		} else {
			m.Script = types.StringNull()
		}
	} else {
		m.Script = types.StringNull()
	}
	if v, ok := obj["secondary-dns"]; ok {
		_ = v
		if v != "" {
			m.SecondaryDNS = types.StringValue(v)
		} else {
			m.SecondaryDNS = types.StringNull()
		}
	} else {
		m.SecondaryDNS = types.StringNull()
	}
	if v, ok := obj["secondary-ntp"]; ok {
		_ = v
		if v != "" {
			m.SecondaryNTP = types.StringValue(v)
		} else {
			m.SecondaryNTP = types.StringNull()
		}
	} else {
		m.SecondaryNTP = types.StringNull()
	}
	if v, ok := obj["status"]; ok {
		_ = v
		if v != "" {
			m.Status = types.StringValue(v)
		} else {
			m.Status = types.StringNull()
		}
	} else {
		m.Status = types.StringNull()
	}
	if v, ok := obj["use-broadcast"]; ok {
		_ = v
		if v != "" {
			m.UseBroadcast = types.StringValue(v)
		} else {
			m.UseBroadcast = types.StringNull()
		}
	} else {
		m.UseBroadcast = types.StringNull()
	}
	if v, ok := obj["use-peer-dns"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.UsePeerDNS = types.BoolValue(b)
		} else {
			m.UsePeerDNS = types.BoolNull()
		}
	} else {
		m.UsePeerDNS = types.BoolNull()
	}
	if v, ok := obj["use-peer-ntp"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.UsePeerNTP = types.BoolValue(b)
		} else {
			m.UsePeerNTP = types.BoolNull()
		}
	} else {
		m.UsePeerNTP = types.BoolNull()
	}
	if v, ok := obj["vlan-priority"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.VLANPriority = types.Int64Value(n)
		} else {
			m.VLANPriority = types.Int64Null()
		}
	} else {
		m.VLANPriority = types.Int64Null()
	}
}
