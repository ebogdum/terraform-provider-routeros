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
	_ resource.Resource                = &IPV6DHCPRelayResource{}
	_ resource.ResourceWithImportState = &IPV6DHCPRelayResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPV6DHCPRelayResource struct {
	reg *client.Registry
}

type IPV6DHCPRelayModel struct {
	ID                   types.String `tfsdk:"id"`
	StoreRelayedBindings types.String `tfsdk:"store_relayed_bindings"`
	Name                 types.String `tfsdk:"name"`
	LinkAddress          types.String `tfsdk:"link_address"`
	Interface            types.String `tfsdk:"interface"`
	DhcpServer           types.String `tfsdk:"dhcp_server"`
	DhcpOptions          types.String `tfsdk:"dhcp_options"`
	DelayThreshold       types.String `tfsdk:"delay_threshold"`
	Comment              types.String `tfsdk:"comment"`
	Disabled             types.Bool   `tfsdk:"disabled"`
	Router               types.String `tfsdk:"router"`
}

func NewIPV6DHCPRelayResource() resource.Resource { return &IPV6DHCPRelayResource{} }

func (r *IPV6DHCPRelayResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv6_dhcp_relay"
}

func (r *IPV6DHCPRelayResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPV6DHCPRelayResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "server address validator rejects literal addresses on this ROS version. Skipped.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"store_relayed_bindings": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `store-relayed-bindings`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `name`.",
			},
			"link_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `link-address`.",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `interface`.",
			},
			"dhcp_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dhcp-server`.",
			},
			"dhcp_options": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dhcp-options`.",
			},
			"delay_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `delay-threshold`.",
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
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPV6DHCPRelayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPV6DHCPRelayModel
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
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.DelayThreshold.IsNull() || plan.DelayThreshold.IsUnknown()) {
		body["delay-threshold"] = plan.DelayThreshold.ValueString()
	}
	if !(plan.DhcpOptions.IsNull() || plan.DhcpOptions.IsUnknown()) {
		body["dhcp-options"] = plan.DhcpOptions.ValueString()
	}
	if !(plan.DhcpServer.IsNull() || plan.DhcpServer.IsUnknown()) {
		body["dhcp-server"] = plan.DhcpServer.ValueString()
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.LinkAddress.IsNull() || plan.LinkAddress.IsUnknown()) {
		body["link-address"] = plan.LinkAddress.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.StoreRelayedBindings.IsNull() || plan.StoreRelayedBindings.IsUnknown()) {
		body["store-relayed-bindings"] = plan.StoreRelayedBindings.ValueString()
	}
	obj, err := c.Add(ctx, "/ipv6/dhcp-relay", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ipv6/dhcp-relay failed", err.Error())
		return
	}
	iPV6DHCPRelayApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6DHCPRelayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPV6DHCPRelayModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ipv6/dhcp-relay", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ipv6/dhcp-relay failed", err.Error())
		return
	}
	iPV6DHCPRelayApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPV6DHCPRelayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPV6DHCPRelayModel
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
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DelayThreshold.Equal(state.DelayThreshold) && !plan.DelayThreshold.IsUnknown() {
		body["delay-threshold"] = plan.DelayThreshold.ValueString()
	}
	if !plan.DhcpOptions.Equal(state.DhcpOptions) && !plan.DhcpOptions.IsUnknown() {
		body["dhcp-options"] = plan.DhcpOptions.ValueString()
	}
	if !plan.DhcpServer.Equal(state.DhcpServer) && !plan.DhcpServer.IsUnknown() {
		body["dhcp-server"] = plan.DhcpServer.ValueString()
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.LinkAddress.Equal(state.LinkAddress) && !plan.LinkAddress.IsUnknown() {
		body["link-address"] = plan.LinkAddress.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.StoreRelayedBindings.Equal(state.StoreRelayedBindings) && !plan.StoreRelayedBindings.IsUnknown() {
		body["store-relayed-bindings"] = plan.StoreRelayedBindings.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ipv6/dhcp-relay", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ipv6/dhcp-relay failed", err.Error())
			return
		}
		iPV6DHCPRelayApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6DHCPRelayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPV6DHCPRelayModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ipv6/dhcp-relay", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ipv6/dhcp-relay failed", err.Error())
	}
}

func (r *IPV6DHCPRelayResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPV6DHCPRelayLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ipv6/dhcp-relay matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPV6DHCPRelayLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPV6DHCPRelayLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ipv6/dhcp-relay", id)
}

func iPV6DHCPRelayApply(ctx context.Context, obj client.Object, m *IPV6DHCPRelayModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["store-relayed-bindings"]; ok && v != "" {
		m.StoreRelayedBindings = types.StringValue(v)
	} else {
		m.StoreRelayedBindings = types.StringNull()
	}
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["link-address"]; ok && v != "" {
		m.LinkAddress = types.StringValue(v)
	} else {
		m.LinkAddress = types.StringNull()
	}
	if v, ok := obj["interface"]; ok && v != "" {
		m.Interface = types.StringValue(v)
	} else {
		m.Interface = types.StringNull()
	}
	if v, ok := obj["dhcp-server"]; ok && v != "" {
		m.DhcpServer = types.StringValue(v)
	} else {
		m.DhcpServer = types.StringNull()
	}
	if v, ok := obj["dhcp-options"]; ok && v != "" {
		m.DhcpOptions = types.StringValue(v)
	} else {
		m.DhcpOptions = types.StringNull()
	}
	if v, ok := obj["delay-threshold"]; ok && v != "" {
		m.DelayThreshold = types.StringValue(v)
	} else {
		m.DelayThreshold = types.StringNull()
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
}
