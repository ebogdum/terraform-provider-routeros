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
	_ resource.Resource                = &PPPProfileResource{}
	_ resource.ResourceWithImportState = &PPPProfileResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type PPPProfileResource struct {
	reg *client.Registry
}

type PPPProfileModel struct {
	ID                    types.String `tfsdk:"id"`
	RateLimit             types.String `tfsdk:"rate_limit"`
	QueueType             types.String `tfsdk:"queue_type"`
	AddressList           types.String `tfsdk:"address_list"`
	Bridge                types.String `tfsdk:"bridge"`
	BridgeHorizon         types.String `tfsdk:"bridge_horizon"`
	BridgeLearning        types.String `tfsdk:"bridge_learning"`
	BridgePathCost        types.String `tfsdk:"bridge_path_cost"`
	BridgePortPriority    types.String `tfsdk:"bridge_port_priority"`
	BridgePortTrusted     types.String `tfsdk:"bridge_port_trusted"`
	BridgePortVid         types.String `tfsdk:"bridge_port_vid"`
	ChangeTCPMss          types.String `tfsdk:"change_tcp_mss"`
	Comment               types.String `tfsdk:"comment"`
	Def                   types.String `tfsdk:"def"`
	Default               types.Bool   `tfsdk:"default"`
	Dhcpv6LeaseTime       types.String `tfsdk:"dhcpv6_lease_time"`
	Dhcpv6PdPool          types.String `tfsdk:"dhcpv6_pd_pool"`
	Dhcpv6UseRADIUS       types.String `tfsdk:"dhcpv6_use_radius"`
	DNSServer             types.String `tfsdk:"dns_server"`
	IdleTimeout           types.String `tfsdk:"idle_timeout"`
	IncomingFilter        types.String `tfsdk:"incoming_filter"`
	InsertQueueBefore     types.String `tfsdk:"insert_queue_before"`
	InterfaceList         types.String `tfsdk:"interface_list"`
	IPV6                  types.String `tfsdk:"ipv6"`
	LocalAddress          types.String `tfsdk:"local_address"`
	Name                  types.String `tfsdk:"name"`
	OnDown                types.String `tfsdk:"on_down"`
	OnUp                  types.String `tfsdk:"on_up"`
	OnlyOne               types.String `tfsdk:"only_one"`
	OutgoingFilter        types.String `tfsdk:"outgoing_filter"`
	ParentQueue           types.String `tfsdk:"parent_queue"`
	QueueTypeRxTx         types.String `tfsdk:"queue_type_rx_tx"`
	RateLimitRxTx         types.String `tfsdk:"rate_limit_rx_tx"`
	RemoteAddress         types.String `tfsdk:"remote_address"`
	RemoteIPV6PrefixPool  types.String `tfsdk:"remote_ipv6_prefix_pool"`
	RemoteIPV6PrefixReuse types.String `tfsdk:"remote_ipv6_prefix_reuse"`
	SessionTimeout        types.String `tfsdk:"session_timeout"`
	UseCompression        types.String `tfsdk:"use_compression"`
	UseEncryption         types.String `tfsdk:"use_encryption"`
	UseIPV6               types.String `tfsdk:"use_ipv6"`
	UseMPLS               types.String `tfsdk:"use_mpls"`
	UseUpnp               types.String `tfsdk:"use_upnp"`
	WinsServer            types.String `tfsdk:"wins_server"`
	Router                types.String `tfsdk:"router"`
}

func NewPPPProfileResource() resource.Resource { return &PPPProfileResource{} }

func (r *PPPProfileResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ppp_profile"
}

func (r *PPPProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *PPPProfileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ppp/profile`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"rate_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `rate-limit`.",
			},
			"queue_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `queue-type`.",
			},
			"address_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge_horizon": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge_learning": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "yes", "default"}...)},
			},
			"bridge_path_cost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge_port_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge_port_trusted": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge_port_vid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"change_tcp_mss": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "yes", "default"}...)},
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"def": schema.StringAttribute{
				Optional:           true,
				Computed:           true,
				Description:        "",
				DeprecationMessage: "Not a RouterOS REST property (WebFig-only spelling of `default`); RouterOS rejects it. Read-only and ignored on write - use `default`.",
			},
			"default": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"dhcpv6_lease_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dhcpv6_pd_pool": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dhcpv6_use_radius": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dns_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"idle_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"incoming_filter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"insert_queue_before": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interface_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ipv6": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"local_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"on_down": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"on_up": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"only_one": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "yes", "default"}...)},
			},
			"outgoing_filter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"parent_queue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"queue_type_rx_tx": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rate_limit_rx_tx": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"remote_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_ipv6_prefix_pool": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_ipv6_prefix_reuse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"session_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"use_compression": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "yes", "default"}...)},
			},
			"use_encryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "yes", "required", "default"}...)},
			},
			"use_ipv6": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "yes", "required", "default"}...)},
			},
			"use_mpls": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "yes", "required", "default"}...)},
			},
			"use_upnp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wins_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
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

func (r *PPPProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PPPProfileModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AddressList.IsNull() || plan.AddressList.IsUnknown()) {
		body["address-list"] = plan.AddressList.ValueString()
	}
	if !(plan.Bridge.IsNull() || plan.Bridge.IsUnknown()) {
		body["bridge"] = plan.Bridge.ValueString()
	}
	if !(plan.BridgeHorizon.IsNull() || plan.BridgeHorizon.IsUnknown()) {
		body["bridge-horizon"] = plan.BridgeHorizon.ValueString()
	}
	if !(plan.BridgeLearning.IsNull() || plan.BridgeLearning.IsUnknown()) {
		body["bridge-learning"] = plan.BridgeLearning.ValueString()
	}
	if !(plan.BridgePathCost.IsNull() || plan.BridgePathCost.IsUnknown()) {
		body["bridge-path-cost"] = plan.BridgePathCost.ValueString()
	}
	if !(plan.BridgePortPriority.IsNull() || plan.BridgePortPriority.IsUnknown()) {
		body["bridge-port-priority"] = plan.BridgePortPriority.ValueString()
	}
	if !(plan.BridgePortTrusted.IsNull() || plan.BridgePortTrusted.IsUnknown()) {
		body["bridge-port-trusted"] = plan.BridgePortTrusted.ValueString()
	}
	if !(plan.BridgePortVid.IsNull() || plan.BridgePortVid.IsUnknown()) {
		body["bridge-port-vid"] = plan.BridgePortVid.ValueString()
	}
	if !(plan.ChangeTCPMss.IsNull() || plan.ChangeTCPMss.IsUnknown()) {
		body["change-tcp-mss"] = plan.ChangeTCPMss.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Dhcpv6LeaseTime.IsNull() || plan.Dhcpv6LeaseTime.IsUnknown()) {
		body["dhcpv6-lease-time"] = plan.Dhcpv6LeaseTime.ValueString()
	}
	if !(plan.Dhcpv6PdPool.IsNull() || plan.Dhcpv6PdPool.IsUnknown()) {
		body["dhcpv6-pd-pool"] = plan.Dhcpv6PdPool.ValueString()
	}
	if !(plan.Dhcpv6UseRADIUS.IsNull() || plan.Dhcpv6UseRADIUS.IsUnknown()) {
		body["dhcpv6-use-radius"] = plan.Dhcpv6UseRADIUS.ValueString()
	}
	if !(plan.DNSServer.IsNull() || plan.DNSServer.IsUnknown()) {
		body["dns-server"] = plan.DNSServer.ValueString()
	}
	if !(plan.IdleTimeout.IsNull() || plan.IdleTimeout.IsUnknown()) {
		body["idle-timeout"] = plan.IdleTimeout.ValueString()
	}
	if !(plan.IncomingFilter.IsNull() || plan.IncomingFilter.IsUnknown()) {
		body["incoming-filter"] = plan.IncomingFilter.ValueString()
	}
	if !(plan.InsertQueueBefore.IsNull() || plan.InsertQueueBefore.IsUnknown()) {
		body["insert-queue-before"] = plan.InsertQueueBefore.ValueString()
	}
	if !(plan.InterfaceList.IsNull() || plan.InterfaceList.IsUnknown()) {
		body["interface-list"] = plan.InterfaceList.ValueString()
	}
	if !(plan.LocalAddress.IsNull() || plan.LocalAddress.IsUnknown()) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.OnDown.IsNull() || plan.OnDown.IsUnknown()) {
		body["on-down"] = plan.OnDown.ValueString()
	}
	if !(plan.OnUp.IsNull() || plan.OnUp.IsUnknown()) {
		body["on-up"] = plan.OnUp.ValueString()
	}
	if !(plan.OnlyOne.IsNull() || plan.OnlyOne.IsUnknown()) {
		body["only-one"] = plan.OnlyOne.ValueString()
	}
	if !(plan.OutgoingFilter.IsNull() || plan.OutgoingFilter.IsUnknown()) {
		body["outgoing-filter"] = plan.OutgoingFilter.ValueString()
	}
	if !(plan.ParentQueue.IsNull() || plan.ParentQueue.IsUnknown()) {
		body["parent-queue"] = plan.ParentQueue.ValueString()
	}
	if !(plan.RemoteAddress.IsNull() || plan.RemoteAddress.IsUnknown()) {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !(plan.RemoteIPV6PrefixPool.IsNull() || plan.RemoteIPV6PrefixPool.IsUnknown()) {
		body["remote-ipv6-prefix-pool"] = plan.RemoteIPV6PrefixPool.ValueString()
	}
	if !(plan.RemoteIPV6PrefixReuse.IsNull() || plan.RemoteIPV6PrefixReuse.IsUnknown()) {
		body["remote-ipv6-prefix-reuse"] = plan.RemoteIPV6PrefixReuse.ValueString()
	}
	if !(plan.SessionTimeout.IsNull() || plan.SessionTimeout.IsUnknown()) {
		body["session-timeout"] = plan.SessionTimeout.ValueString()
	}
	if !(plan.UseCompression.IsNull() || plan.UseCompression.IsUnknown()) {
		body["use-compression"] = plan.UseCompression.ValueString()
	}
	if !(plan.UseEncryption.IsNull() || plan.UseEncryption.IsUnknown()) {
		body["use-encryption"] = plan.UseEncryption.ValueString()
	}
	if !(plan.UseIPV6.IsNull() || plan.UseIPV6.IsUnknown()) {
		body["use-ipv6"] = plan.UseIPV6.ValueString()
	}
	if !(plan.UseMPLS.IsNull() || plan.UseMPLS.IsUnknown()) {
		body["use-mpls"] = plan.UseMPLS.ValueString()
	}
	if !(plan.UseUpnp.IsNull() || plan.UseUpnp.IsUnknown()) {
		body["use-upnp"] = plan.UseUpnp.ValueString()
	}
	if !(plan.WinsServer.IsNull() || plan.WinsServer.IsUnknown()) {
		body["wins-server"] = plan.WinsServer.ValueString()
	}
	if !(plan.QueueType.IsNull() || plan.QueueType.IsUnknown()) {
		body["queue-type"] = plan.QueueType.ValueString()
	}
	if !(plan.RateLimit.IsNull() || plan.RateLimit.IsUnknown()) {
		body["rate-limit"] = plan.RateLimit.ValueString()
	}
	obj, err := c.Add(ctx, "/ppp/profile", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ppp/profile failed", err.Error())
		return
	}
	pPPProfileApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PPPProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PPPProfileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ppp/profile", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ppp/profile failed", err.Error())
		return
	}
	pPPProfileApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *PPPProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PPPProfileModel
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
	if !plan.AddressList.Equal(state.AddressList) && !plan.AddressList.IsUnknown() {
		body["address-list"] = plan.AddressList.ValueString()
	}
	if !plan.Bridge.Equal(state.Bridge) && !plan.Bridge.IsUnknown() {
		body["bridge"] = plan.Bridge.ValueString()
	}
	if !plan.BridgeHorizon.Equal(state.BridgeHorizon) && !plan.BridgeHorizon.IsUnknown() {
		body["bridge-horizon"] = plan.BridgeHorizon.ValueString()
	}
	if !plan.BridgeLearning.Equal(state.BridgeLearning) && !plan.BridgeLearning.IsUnknown() {
		body["bridge-learning"] = plan.BridgeLearning.ValueString()
	}
	if !plan.BridgePathCost.Equal(state.BridgePathCost) && !plan.BridgePathCost.IsUnknown() {
		body["bridge-path-cost"] = plan.BridgePathCost.ValueString()
	}
	if !plan.BridgePortPriority.Equal(state.BridgePortPriority) && !plan.BridgePortPriority.IsUnknown() {
		body["bridge-port-priority"] = plan.BridgePortPriority.ValueString()
	}
	if !plan.BridgePortTrusted.Equal(state.BridgePortTrusted) && !plan.BridgePortTrusted.IsUnknown() {
		body["bridge-port-trusted"] = plan.BridgePortTrusted.ValueString()
	}
	if !plan.BridgePortVid.Equal(state.BridgePortVid) && !plan.BridgePortVid.IsUnknown() {
		body["bridge-port-vid"] = plan.BridgePortVid.ValueString()
	}
	if !plan.ChangeTCPMss.Equal(state.ChangeTCPMss) && !plan.ChangeTCPMss.IsUnknown() {
		body["change-tcp-mss"] = plan.ChangeTCPMss.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Dhcpv6LeaseTime.Equal(state.Dhcpv6LeaseTime) && !plan.Dhcpv6LeaseTime.IsUnknown() {
		body["dhcpv6-lease-time"] = plan.Dhcpv6LeaseTime.ValueString()
	}
	if !plan.Dhcpv6PdPool.Equal(state.Dhcpv6PdPool) && !plan.Dhcpv6PdPool.IsUnknown() {
		body["dhcpv6-pd-pool"] = plan.Dhcpv6PdPool.ValueString()
	}
	if !plan.Dhcpv6UseRADIUS.Equal(state.Dhcpv6UseRADIUS) && !plan.Dhcpv6UseRADIUS.IsUnknown() {
		body["dhcpv6-use-radius"] = plan.Dhcpv6UseRADIUS.ValueString()
	}
	if !plan.DNSServer.Equal(state.DNSServer) && !plan.DNSServer.IsUnknown() {
		body["dns-server"] = plan.DNSServer.ValueString()
	}
	if !plan.IdleTimeout.Equal(state.IdleTimeout) && !plan.IdleTimeout.IsUnknown() {
		body["idle-timeout"] = plan.IdleTimeout.ValueString()
	}
	if !plan.IncomingFilter.Equal(state.IncomingFilter) && !plan.IncomingFilter.IsUnknown() {
		body["incoming-filter"] = plan.IncomingFilter.ValueString()
	}
	if !plan.InsertQueueBefore.Equal(state.InsertQueueBefore) && !plan.InsertQueueBefore.IsUnknown() {
		body["insert-queue-before"] = plan.InsertQueueBefore.ValueString()
	}
	if !plan.InterfaceList.Equal(state.InterfaceList) && !plan.InterfaceList.IsUnknown() {
		body["interface-list"] = plan.InterfaceList.ValueString()
	}
	if !plan.LocalAddress.Equal(state.LocalAddress) && !plan.LocalAddress.IsUnknown() {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.OnDown.Equal(state.OnDown) && !plan.OnDown.IsUnknown() {
		body["on-down"] = plan.OnDown.ValueString()
	}
	if !plan.OnUp.Equal(state.OnUp) && !plan.OnUp.IsUnknown() {
		body["on-up"] = plan.OnUp.ValueString()
	}
	if !plan.OnlyOne.Equal(state.OnlyOne) && !plan.OnlyOne.IsUnknown() {
		body["only-one"] = plan.OnlyOne.ValueString()
	}
	if !plan.OutgoingFilter.Equal(state.OutgoingFilter) && !plan.OutgoingFilter.IsUnknown() {
		body["outgoing-filter"] = plan.OutgoingFilter.ValueString()
	}
	if !plan.ParentQueue.Equal(state.ParentQueue) && !plan.ParentQueue.IsUnknown() {
		body["parent-queue"] = plan.ParentQueue.ValueString()
	}
	if !plan.RemoteAddress.Equal(state.RemoteAddress) && !plan.RemoteAddress.IsUnknown() {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !plan.RemoteIPV6PrefixPool.Equal(state.RemoteIPV6PrefixPool) && !plan.RemoteIPV6PrefixPool.IsUnknown() {
		body["remote-ipv6-prefix-pool"] = plan.RemoteIPV6PrefixPool.ValueString()
	}
	if !plan.RemoteIPV6PrefixReuse.Equal(state.RemoteIPV6PrefixReuse) && !plan.RemoteIPV6PrefixReuse.IsUnknown() {
		body["remote-ipv6-prefix-reuse"] = plan.RemoteIPV6PrefixReuse.ValueString()
	}
	if !plan.SessionTimeout.Equal(state.SessionTimeout) && !plan.SessionTimeout.IsUnknown() {
		body["session-timeout"] = plan.SessionTimeout.ValueString()
	}
	if !plan.UseCompression.Equal(state.UseCompression) && !plan.UseCompression.IsUnknown() {
		body["use-compression"] = plan.UseCompression.ValueString()
	}
	if !plan.UseEncryption.Equal(state.UseEncryption) && !plan.UseEncryption.IsUnknown() {
		body["use-encryption"] = plan.UseEncryption.ValueString()
	}
	if !plan.UseIPV6.Equal(state.UseIPV6) && !plan.UseIPV6.IsUnknown() {
		body["use-ipv6"] = plan.UseIPV6.ValueString()
	}
	if !plan.UseMPLS.Equal(state.UseMPLS) && !plan.UseMPLS.IsUnknown() {
		body["use-mpls"] = plan.UseMPLS.ValueString()
	}
	if !plan.UseUpnp.Equal(state.UseUpnp) && !plan.UseUpnp.IsUnknown() {
		body["use-upnp"] = plan.UseUpnp.ValueString()
	}
	if !plan.WinsServer.Equal(state.WinsServer) && !plan.WinsServer.IsUnknown() {
		body["wins-server"] = plan.WinsServer.ValueString()
	}
	if !plan.QueueType.Equal(state.QueueType) && !plan.QueueType.IsUnknown() {
		body["queue-type"] = plan.QueueType.ValueString()
	}
	if !plan.RateLimit.Equal(state.RateLimit) && !plan.RateLimit.IsUnknown() {
		body["rate-limit"] = plan.RateLimit.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ppp/profile", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ppp/profile failed", err.Error())
			return
		}
		pPPProfileApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PPPProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PPPProfileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ppp/profile", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ppp/profile failed", err.Error())
	}
}

func (r *PPPProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := pPPProfileLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ppp/profile matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// pPPProfileLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func pPPProfileLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ppp/profile", id)
}

func pPPProfileApply(ctx context.Context, obj client.Object, m *PPPProfileModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["rate-limit"]; ok && v != "" {
		m.RateLimit = types.StringValue(v)
	} else {
		m.RateLimit = types.StringNull()
	}
	if v, ok := obj["queue-type"]; ok && v != "" {
		m.QueueType = types.StringValue(v)
	} else {
		m.QueueType = types.StringNull()
	}
	if v, ok := obj["address-list"]; ok {
		if v != "" {
			m.AddressList = types.StringValue(v)
		} else {
			m.AddressList = types.StringNull()
		}
	}
	if v, ok := obj["bridge"]; ok {
		if v != "" {
			m.Bridge = types.StringValue(v)
		} else {
			m.Bridge = types.StringNull()
		}
	}
	if v, ok := obj["bridge-horizon"]; ok {
		if v != "" {
			m.BridgeHorizon = types.StringValue(v)
		} else {
			m.BridgeHorizon = types.StringNull()
		}
	}
	if v, ok := obj["bridge-learning"]; ok {
		if v != "" {
			m.BridgeLearning = types.StringValue(v)
		} else {
			m.BridgeLearning = types.StringNull()
		}
	}
	if v, ok := obj["bridge-path-cost"]; ok {
		if v != "" {
			m.BridgePathCost = types.StringValue(v)
		} else {
			m.BridgePathCost = types.StringNull()
		}
	}
	if v, ok := obj["bridge-port-priority"]; ok {
		if v != "" {
			m.BridgePortPriority = types.StringValue(v)
		} else {
			m.BridgePortPriority = types.StringNull()
		}
	}
	if v, ok := obj["bridge-port-trusted"]; ok {
		if v != "" {
			m.BridgePortTrusted = types.StringValue(v)
		} else {
			m.BridgePortTrusted = types.StringNull()
		}
	}
	if v, ok := obj["bridge-port-vid"]; ok {
		if v != "" {
			m.BridgePortVid = types.StringValue(v)
		} else {
			m.BridgePortVid = types.StringNull()
		}
	}
	if v, ok := obj["change-tcp-mss"]; ok {
		if v != "" {
			m.ChangeTCPMss = types.StringValue(v)
		} else {
			m.ChangeTCPMss = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["def"]; ok {
		if v != "" {
			m.Def = types.StringValue(v)
		} else {
			m.Def = types.StringNull()
		}
	}
	if v, ok := obj["default"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Default = types.BoolValue(true)
		} else {
			m.Default = types.BoolNull()
		}
	}
	if v, ok := obj["dhcpv6-lease-time"]; ok {
		if v != "" {
			m.Dhcpv6LeaseTime = types.StringValue(v)
		} else {
			m.Dhcpv6LeaseTime = types.StringNull()
		}
	}
	if v, ok := obj["dhcpv6-pd-pool"]; ok {
		if v != "" {
			m.Dhcpv6PdPool = types.StringValue(v)
		} else {
			m.Dhcpv6PdPool = types.StringNull()
		}
	}
	if v, ok := obj["dhcpv6-use-radius"]; ok {
		if v != "" {
			m.Dhcpv6UseRADIUS = types.StringValue(v)
		} else {
			m.Dhcpv6UseRADIUS = types.StringNull()
		}
	}
	if v, ok := obj["dns-server"]; ok {
		if v != "" {
			m.DNSServer = types.StringValue(v)
		} else {
			m.DNSServer = types.StringNull()
		}
	}
	if v, ok := obj["idle-timeout"]; ok {
		if v != "" {
			m.IdleTimeout = types.StringValue(v)
		} else {
			m.IdleTimeout = types.StringNull()
		}
	}
	if v, ok := obj["incoming-filter"]; ok {
		if v != "" {
			m.IncomingFilter = types.StringValue(v)
		} else {
			m.IncomingFilter = types.StringNull()
		}
	}
	if v, ok := obj["insert-queue-before"]; ok {
		if v != "" {
			m.InsertQueueBefore = types.StringValue(v)
		} else {
			m.InsertQueueBefore = types.StringNull()
		}
	}
	if v, ok := obj["interface-list"]; ok {
		if v != "" {
			m.InterfaceList = types.StringValue(v)
		} else {
			m.InterfaceList = types.StringNull()
		}
	}
	if v, ok := obj["ipv6"]; ok {
		if v != "" {
			m.IPV6 = types.StringValue(v)
		} else {
			m.IPV6 = types.StringNull()
		}
	}
	if v, ok := obj["local-address"]; ok {
		if v != "" {
			m.LocalAddress = types.StringValue(v)
		} else {
			m.LocalAddress = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["on-down"]; ok {
		if v != "" {
			m.OnDown = types.StringValue(v)
		} else {
			m.OnDown = types.StringNull()
		}
	}
	if v, ok := obj["on-up"]; ok {
		if v != "" {
			m.OnUp = types.StringValue(v)
		} else {
			m.OnUp = types.StringNull()
		}
	}
	if v, ok := obj["only-one"]; ok {
		if v != "" {
			m.OnlyOne = types.StringValue(v)
		} else {
			m.OnlyOne = types.StringNull()
		}
	}
	if v, ok := obj["outgoing-filter"]; ok {
		if v != "" {
			m.OutgoingFilter = types.StringValue(v)
		} else {
			m.OutgoingFilter = types.StringNull()
		}
	}
	if v, ok := obj["parent-queue"]; ok {
		if v != "" {
			m.ParentQueue = types.StringValue(v)
		} else {
			m.ParentQueue = types.StringNull()
		}
	}
	if v, ok := obj["queue-type-rx-tx"]; ok {
		if v != "" {
			m.QueueTypeRxTx = types.StringValue(v)
		} else {
			m.QueueTypeRxTx = types.StringNull()
		}
	}
	if v, ok := obj["rate-limit-rx-tx"]; ok {
		if v != "" {
			m.RateLimitRxTx = types.StringValue(v)
		} else {
			m.RateLimitRxTx = types.StringNull()
		}
	}
	if v, ok := obj["remote-address"]; ok {
		if v != "" {
			m.RemoteAddress = types.StringValue(v)
		} else {
			m.RemoteAddress = types.StringNull()
		}
	}
	if v, ok := obj["remote-ipv6-prefix-pool"]; ok {
		if v != "" {
			m.RemoteIPV6PrefixPool = types.StringValue(v)
		} else {
			m.RemoteIPV6PrefixPool = types.StringNull()
		}
	}
	if v, ok := obj["remote-ipv6-prefix-reuse"]; ok {
		if v != "" {
			m.RemoteIPV6PrefixReuse = types.StringValue(v)
		} else {
			m.RemoteIPV6PrefixReuse = types.StringNull()
		}
	}
	if v, ok := obj["session-timeout"]; ok {
		if v != "" {
			m.SessionTimeout = types.StringValue(v)
		} else {
			m.SessionTimeout = types.StringNull()
		}
	}
	if v, ok := obj["use-compression"]; ok {
		if v != "" {
			m.UseCompression = types.StringValue(v)
		} else {
			m.UseCompression = types.StringNull()
		}
	}
	if v, ok := obj["use-encryption"]; ok {
		if v != "" {
			m.UseEncryption = types.StringValue(v)
		} else {
			m.UseEncryption = types.StringNull()
		}
	}
	if v, ok := obj["use-ipv6"]; ok {
		if v != "" {
			m.UseIPV6 = types.StringValue(v)
		} else {
			m.UseIPV6 = types.StringNull()
		}
	}
	if v, ok := obj["use-mpls"]; ok {
		if v != "" {
			m.UseMPLS = types.StringValue(v)
		} else {
			m.UseMPLS = types.StringNull()
		}
	}
	if v, ok := obj["use-upnp"]; ok {
		if v != "" {
			m.UseUpnp = types.StringValue(v)
		} else {
			m.UseUpnp = types.StringNull()
		}
	}
	if v, ok := obj["wins-server"]; ok {
		if v != "" {
			m.WinsServer = types.StringValue(v)
		} else {
			m.WinsServer = types.StringNull()
		}
	}
}
