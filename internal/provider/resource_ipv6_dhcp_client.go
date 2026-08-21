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
	_ resource.Resource                = &IPV6DHCPClientResource{}
	_ resource.ResourceWithImportState = &IPV6DHCPClientResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPV6DHCPClientResource struct {
	reg *client.Registry
}

type IPV6DHCPClientModel struct {
	ID                         types.String `tfsdk:"id"`
	ValidateServerDuid         types.String `tfsdk:"validate_server_duid"`
	UsePeerDns                 types.String `tfsdk:"use_peer_dns"`
	UseInterfaceDuid           types.String `tfsdk:"use_interface_duid"`
	Script                     types.String `tfsdk:"script"`
	RapidCommit                types.String `tfsdk:"rapid_commit"`
	PrefixHint                 types.String `tfsdk:"prefix_hint"`
	PrefixAddressLists         types.String `tfsdk:"prefix_address_lists"`
	PoolPrefixLength           types.String `tfsdk:"pool_prefix_length"`
	PoolName                   types.String `tfsdk:"pool_name"`
	DhcpOptions                types.String `tfsdk:"dhcp_options"`
	DefaultRouteTables         types.String `tfsdk:"default_route_tables"`
	CustomIapdId               types.String `tfsdk:"custom_iapd_id"`
	CustomIanaId               types.String `tfsdk:"custom_iana_id"`
	CustomDuid                 types.String `tfsdk:"custom_duid"`
	CheckGateway               types.String `tfsdk:"check_gateway"`
	AllowReconfigure           types.String `tfsdk:"allow_reconfigure"`
	AddDefaultRoute            types.String `tfsdk:"add_default_route"`
	AcceptPrefixWithoutAddress types.String `tfsdk:"accept_prefix_without_address"`
	Comment                    types.String `tfsdk:"comment"`
	DefaultRouteDistance       types.String `tfsdk:"default_route_distance"`
	Disabled                   types.Bool   `tfsdk:"disabled"`
	Interface                  types.String `tfsdk:"interface"`
	Request                    types.String `tfsdk:"request"`
	Router                     types.String `tfsdk:"router"`
}

func NewIPV6DHCPClientResource() resource.Resource { return &IPV6DHCPClientResource{} }

func (r *IPV6DHCPClientResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv6_dhcp_client"
}

func (r *IPV6DHCPClientResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPV6DHCPClientResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ipv6/dhcp-client`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"validate_server_duid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `validate-server-duid`.",
			},
			"use_peer_dns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `use-peer-dns`.",
			},
			"use_interface_duid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `use-interface-duid`.",
			},
			"script": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `script`.",
			},
			"rapid_commit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `rapid-commit`.",
			},
			"prefix_hint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `prefix-hint`.",
			},
			"prefix_address_lists": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `prefix-address-lists`.",
			},
			"pool_prefix_length": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `pool-prefix-length`.",
			},
			"pool_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `pool-name`.",
			},
			"dhcp_options": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dhcp-options`.",
			},
			"default_route_tables": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `default-route-tables`.",
			},
			"custom_iapd_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `custom-iapd-id`.",
			},
			"custom_iana_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `custom-iana-id`.",
			},
			"custom_duid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `custom-duid`.",
			},
			"check_gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `check-gateway`.",
			},
			"allow_reconfigure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `allow-reconfigure`.",
			},
			"add_default_route": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `add-default-route`.",
			},
			"accept_prefix_without_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `accept-prefix-without-address`.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"default_route_distance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"interface": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"request": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *IPV6DHCPClientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPV6DHCPClientModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.DefaultRouteDistance.IsNull() || plan.DefaultRouteDistance.IsUnknown()) {
		body["default-route-distance"] = plan.DefaultRouteDistance.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.Request.IsNull() || plan.Request.IsUnknown()) {
		body["request"] = plan.Request.ValueString()
	}
	if !(plan.AcceptPrefixWithoutAddress.IsNull() || plan.AcceptPrefixWithoutAddress.IsUnknown()) {
		body["accept-prefix-without-address"] = plan.AcceptPrefixWithoutAddress.ValueString()
	}
	if !(plan.AddDefaultRoute.IsNull() || plan.AddDefaultRoute.IsUnknown()) {
		body["add-default-route"] = plan.AddDefaultRoute.ValueString()
	}
	if !(plan.AllowReconfigure.IsNull() || plan.AllowReconfigure.IsUnknown()) {
		body["allow-reconfigure"] = plan.AllowReconfigure.ValueString()
	}
	if !(plan.CheckGateway.IsNull() || plan.CheckGateway.IsUnknown()) {
		body["check-gateway"] = plan.CheckGateway.ValueString()
	}
	if !(plan.CustomDuid.IsNull() || plan.CustomDuid.IsUnknown()) {
		body["custom-duid"] = plan.CustomDuid.ValueString()
	}
	if !(plan.CustomIanaId.IsNull() || plan.CustomIanaId.IsUnknown()) {
		body["custom-iana-id"] = plan.CustomIanaId.ValueString()
	}
	if !(plan.CustomIapdId.IsNull() || plan.CustomIapdId.IsUnknown()) {
		body["custom-iapd-id"] = plan.CustomIapdId.ValueString()
	}
	if !(plan.DefaultRouteTables.IsNull() || plan.DefaultRouteTables.IsUnknown()) {
		body["default-route-tables"] = plan.DefaultRouteTables.ValueString()
	}
	if !(plan.DhcpOptions.IsNull() || plan.DhcpOptions.IsUnknown()) {
		body["dhcp-options"] = plan.DhcpOptions.ValueString()
	}
	if !(plan.PoolName.IsNull() || plan.PoolName.IsUnknown()) {
		body["pool-name"] = plan.PoolName.ValueString()
	}
	if !(plan.PoolPrefixLength.IsNull() || plan.PoolPrefixLength.IsUnknown()) {
		body["pool-prefix-length"] = plan.PoolPrefixLength.ValueString()
	}
	if !(plan.PrefixAddressLists.IsNull() || plan.PrefixAddressLists.IsUnknown()) {
		body["prefix-address-lists"] = plan.PrefixAddressLists.ValueString()
	}
	if !(plan.PrefixHint.IsNull() || plan.PrefixHint.IsUnknown()) {
		body["prefix-hint"] = plan.PrefixHint.ValueString()
	}
	if !(plan.RapidCommit.IsNull() || plan.RapidCommit.IsUnknown()) {
		body["rapid-commit"] = plan.RapidCommit.ValueString()
	}
	if !(plan.Script.IsNull() || plan.Script.IsUnknown()) {
		body["script"] = plan.Script.ValueString()
	}
	if !(plan.UseInterfaceDuid.IsNull() || plan.UseInterfaceDuid.IsUnknown()) {
		body["use-interface-duid"] = plan.UseInterfaceDuid.ValueString()
	}
	if !(plan.UsePeerDns.IsNull() || plan.UsePeerDns.IsUnknown()) {
		body["use-peer-dns"] = plan.UsePeerDns.ValueString()
	}
	if !(plan.ValidateServerDuid.IsNull() || plan.ValidateServerDuid.IsUnknown()) {
		body["validate-server-duid"] = plan.ValidateServerDuid.ValueString()
	}
	obj, err := c.Add(ctx, "/ipv6/dhcp-client", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ipv6/dhcp-client failed", err.Error())
		return
	}
	iPV6DHCPClientApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6DHCPClientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPV6DHCPClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ipv6/dhcp-client", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ipv6/dhcp-client failed", err.Error())
		return
	}
	iPV6DHCPClientApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPV6DHCPClientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPV6DHCPClientModel
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
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.DefaultRouteDistance.Equal(state.DefaultRouteDistance) && !plan.DefaultRouteDistance.IsUnknown() {
		body["default-route-distance"] = plan.DefaultRouteDistance.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.Request.Equal(state.Request) && !plan.Request.IsUnknown() {
		body["request"] = plan.Request.ValueString()
	}
	if !plan.AcceptPrefixWithoutAddress.Equal(state.AcceptPrefixWithoutAddress) && !plan.AcceptPrefixWithoutAddress.IsUnknown() {
		body["accept-prefix-without-address"] = plan.AcceptPrefixWithoutAddress.ValueString()
	}
	if !plan.AddDefaultRoute.Equal(state.AddDefaultRoute) && !plan.AddDefaultRoute.IsUnknown() {
		body["add-default-route"] = plan.AddDefaultRoute.ValueString()
	}
	if !plan.AllowReconfigure.Equal(state.AllowReconfigure) && !plan.AllowReconfigure.IsUnknown() {
		body["allow-reconfigure"] = plan.AllowReconfigure.ValueString()
	}
	if !plan.CheckGateway.Equal(state.CheckGateway) && !plan.CheckGateway.IsUnknown() {
		body["check-gateway"] = plan.CheckGateway.ValueString()
	}
	if !plan.CustomDuid.Equal(state.CustomDuid) && !plan.CustomDuid.IsUnknown() {
		body["custom-duid"] = plan.CustomDuid.ValueString()
	}
	if !plan.CustomIanaId.Equal(state.CustomIanaId) && !plan.CustomIanaId.IsUnknown() {
		body["custom-iana-id"] = plan.CustomIanaId.ValueString()
	}
	if !plan.CustomIapdId.Equal(state.CustomIapdId) && !plan.CustomIapdId.IsUnknown() {
		body["custom-iapd-id"] = plan.CustomIapdId.ValueString()
	}
	if !plan.DefaultRouteTables.Equal(state.DefaultRouteTables) && !plan.DefaultRouteTables.IsUnknown() {
		body["default-route-tables"] = plan.DefaultRouteTables.ValueString()
	}
	if !plan.DhcpOptions.Equal(state.DhcpOptions) && !plan.DhcpOptions.IsUnknown() {
		body["dhcp-options"] = plan.DhcpOptions.ValueString()
	}
	if !plan.PoolName.Equal(state.PoolName) && !plan.PoolName.IsUnknown() {
		body["pool-name"] = plan.PoolName.ValueString()
	}
	if !plan.PoolPrefixLength.Equal(state.PoolPrefixLength) && !plan.PoolPrefixLength.IsUnknown() {
		body["pool-prefix-length"] = plan.PoolPrefixLength.ValueString()
	}
	if !plan.PrefixAddressLists.Equal(state.PrefixAddressLists) && !plan.PrefixAddressLists.IsUnknown() {
		body["prefix-address-lists"] = plan.PrefixAddressLists.ValueString()
	}
	if !plan.PrefixHint.Equal(state.PrefixHint) && !plan.PrefixHint.IsUnknown() {
		body["prefix-hint"] = plan.PrefixHint.ValueString()
	}
	if !plan.RapidCommit.Equal(state.RapidCommit) && !plan.RapidCommit.IsUnknown() {
		body["rapid-commit"] = plan.RapidCommit.ValueString()
	}
	if !plan.Script.Equal(state.Script) && !plan.Script.IsUnknown() {
		body["script"] = plan.Script.ValueString()
	}
	if !plan.UseInterfaceDuid.Equal(state.UseInterfaceDuid) && !plan.UseInterfaceDuid.IsUnknown() {
		body["use-interface-duid"] = plan.UseInterfaceDuid.ValueString()
	}
	if !plan.UsePeerDns.Equal(state.UsePeerDns) && !plan.UsePeerDns.IsUnknown() {
		body["use-peer-dns"] = plan.UsePeerDns.ValueString()
	}
	if !plan.ValidateServerDuid.Equal(state.ValidateServerDuid) && !plan.ValidateServerDuid.IsUnknown() {
		body["validate-server-duid"] = plan.ValidateServerDuid.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ipv6/dhcp-client", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ipv6/dhcp-client failed", err.Error())
			return
		}
		iPV6DHCPClientApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6DHCPClientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPV6DHCPClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ipv6/dhcp-client", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ipv6/dhcp-client failed", err.Error())
	}
}

func (r *IPV6DHCPClientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPV6DHCPClientLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ipv6/dhcp-client matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPV6DHCPClientLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPV6DHCPClientLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ipv6/dhcp-client", id)
}

func iPV6DHCPClientApply(ctx context.Context, obj client.Object, m *IPV6DHCPClientModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["validate-server-duid"]; ok && v != "" {
		m.ValidateServerDuid = types.StringValue(v)
	} else {
		m.ValidateServerDuid = types.StringNull()
	}
	if v, ok := obj["use-peer-dns"]; ok && v != "" {
		m.UsePeerDns = types.StringValue(v)
	} else {
		m.UsePeerDns = types.StringNull()
	}
	if v, ok := obj["use-interface-duid"]; ok && v != "" {
		m.UseInterfaceDuid = types.StringValue(v)
	} else {
		m.UseInterfaceDuid = types.StringNull()
	}
	if v, ok := obj["script"]; ok && v != "" {
		m.Script = types.StringValue(v)
	} else {
		m.Script = types.StringNull()
	}
	if v, ok := obj["rapid-commit"]; ok && v != "" {
		m.RapidCommit = types.StringValue(v)
	} else {
		m.RapidCommit = types.StringNull()
	}
	if v, ok := obj["prefix-hint"]; ok && v != "" {
		m.PrefixHint = types.StringValue(v)
	} else {
		m.PrefixHint = types.StringNull()
	}
	if v, ok := obj["prefix-address-lists"]; ok && v != "" {
		m.PrefixAddressLists = types.StringValue(v)
	} else {
		m.PrefixAddressLists = types.StringNull()
	}
	if v, ok := obj["pool-prefix-length"]; ok && v != "" {
		m.PoolPrefixLength = types.StringValue(v)
	} else {
		m.PoolPrefixLength = types.StringNull()
	}
	if v, ok := obj["pool-name"]; ok && v != "" {
		m.PoolName = types.StringValue(v)
	} else {
		m.PoolName = types.StringNull()
	}
	if v, ok := obj["dhcp-options"]; ok && v != "" {
		m.DhcpOptions = types.StringValue(v)
	} else {
		m.DhcpOptions = types.StringNull()
	}
	if v, ok := obj["default-route-tables"]; ok && v != "" {
		m.DefaultRouteTables = types.StringValue(v)
	} else {
		m.DefaultRouteTables = types.StringNull()
	}
	if v, ok := obj["custom-iapd-id"]; ok && v != "" {
		m.CustomIapdId = types.StringValue(v)
	} else {
		m.CustomIapdId = types.StringNull()
	}
	if v, ok := obj["custom-iana-id"]; ok && v != "" {
		m.CustomIanaId = types.StringValue(v)
	} else {
		m.CustomIanaId = types.StringNull()
	}
	if v, ok := obj["custom-duid"]; ok && v != "" {
		m.CustomDuid = types.StringValue(v)
	} else {
		m.CustomDuid = types.StringNull()
	}
	if v, ok := obj["check-gateway"]; ok && v != "" {
		m.CheckGateway = types.StringValue(v)
	} else {
		m.CheckGateway = types.StringNull()
	}
	if v, ok := obj["allow-reconfigure"]; ok && v != "" {
		m.AllowReconfigure = types.StringValue(v)
	} else {
		m.AllowReconfigure = types.StringNull()
	}
	if v, ok := obj["add-default-route"]; ok && v != "" {
		m.AddDefaultRoute = types.StringValue(v)
	} else {
		m.AddDefaultRoute = types.StringNull()
	}
	if v, ok := obj["accept-prefix-without-address"]; ok && v != "" {
		m.AcceptPrefixWithoutAddress = types.StringValue(v)
	} else {
		m.AcceptPrefixWithoutAddress = types.StringNull()
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["default-route-distance"]; ok {
		if v != "" {
			m.DefaultRouteDistance = types.StringValue(v)
		} else {
			m.DefaultRouteDistance = types.StringNull()
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
	if v, ok := obj["request"]; ok {
		if v != "" {
			m.Request = types.StringValue(v)
		} else {
			m.Request = types.StringNull()
		}
	}
}
