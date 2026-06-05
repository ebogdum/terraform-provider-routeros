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
	_ resource.Resource                = &IPDHCPServerLeaseResource{}
	_ resource.ResourceWithImportState = &IPDHCPServerLeaseResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPDHCPServerLeaseResource struct {
	reg *client.Registry
}

type IPDHCPServerLeaseModel struct {
	ID                   types.String `tfsdk:"id"`
	ActiveAddress        types.String `tfsdk:"active_address"`
	ActiveAgentCircuitID types.String `tfsdk:"active_agent_circuit_id"`
	ActiveAgentRemoteID  types.String `tfsdk:"active_agent_remote_id"`
	ActiveClassID        types.String `tfsdk:"active_class_id"`
	ActiveClientID       types.String `tfsdk:"active_client_id"`
	ActiveHostName       types.String `tfsdk:"active_host_name"`
	ActiveMACAddress     types.String `tfsdk:"active_mac_address"`
	ActiveServer         types.String `tfsdk:"active_server"`
	Address              types.String `tfsdk:"address"`
	AddressList          types.String `tfsdk:"address_list"`
	Age                  types.String `tfsdk:"age"`
	AgentCircuitID       types.String `tfsdk:"agent_circuit_id"`
	AgentRemoteID        types.String `tfsdk:"agent_remote_id"`
	AllowDualStackQueue  types.String `tfsdk:"allow_dual_stack_queue"`
	AlwaysBroadcast      types.Bool   `tfsdk:"always_broadcast"`
	BlockAccess          types.Bool   `tfsdk:"block_access"`
	Blocked              types.Bool   `tfsdk:"blocked"`
	BridgePort           types.String `tfsdk:"bridge_port"`
	CheckStatus          types.String `tfsdk:"check_status"`
	ClientID             types.String `tfsdk:"client_id"`
	Comment              types.String `tfsdk:"comment"`
	DHCPOptionSet        types.String `tfsdk:"dhcp_option_set"`
	DHCPOptions          types.String `tfsdk:"dhcp_options"`
	Disabled             types.Bool   `tfsdk:"disabled"`
	Dyn                  types.String `tfsdk:"dyn"`
	Dynamic              types.Bool   `tfsdk:"dynamic"`
	ExpiresAfter         types.String `tfsdk:"expires_after"`
	InsertQueueBefore    types.String `tfsdk:"insert_queue_before"`
	LastSeen             types.String `tfsdk:"last_seen"`
	LastSentCounter      types.String `tfsdk:"last_sent_counter"`
	LeaseTime            types.String `tfsdk:"lease_time"`
	MACAddress           types.String `tfsdk:"mac_address"`
	MakeStatic           types.String `tfsdk:"make_static"`
	ParentQueue          types.String `tfsdk:"parent_queue"`
	Ping                 types.String `tfsdk:"ping"`
	QueueType            types.String `tfsdk:"queue_type"`
	RADIUS               types.Bool   `tfsdk:"radius"`
	RateLimit            types.String `tfsdk:"rate_limit"`
	ReconfigureKey       types.String `tfsdk:"reconfigure_key"`
	ReconfigureStatus    types.String `tfsdk:"reconfigure_status"`
	Rostat               types.String `tfsdk:"rostat"`
	Routes               types.String `tfsdk:"routes"`
	SendReconfigure      types.String `tfsdk:"send_reconfigure"`
	Server               types.String `tfsdk:"server"`
	SrcMACAddress        types.String `tfsdk:"src_mac_address"`
	Stat                 types.String `tfsdk:"stat"`
	UseSrcMACAddress     types.Bool   `tfsdk:"use_src_mac_address"`
	Router               types.String `tfsdk:"router"`
}

func NewIPDHCPServerLeaseResource() resource.Resource { return &IPDHCPServerLeaseResource{} }

func (r *IPDHCPServerLeaseResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_dhcp_server_lease"
}

func (r *IPDHCPServerLeaseResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *IPDHCPServerLeaseResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Lease entries reference an existing dhcp-server; auto-test can't synthesise the precondition reliably.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"active_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"active_agent_circuit_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"active_agent_remote_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"active_class_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"active_client_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"active_host_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"active_mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"active_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"address_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"age": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"agent_circuit_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"agent_remote_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"allow_dual_stack_queue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"always_broadcast": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"block_access": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"blocked": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"check_status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"client_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"dhcp_option_set": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dhcp_options": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"dyn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dynamic": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"expires_after": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"insert_queue_before": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"last_seen": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"last_sent_counter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"lease_time": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"make_static": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"parent_queue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ping": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"queue_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"radius": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rate_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"reconfigure_key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"reconfigure_status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rostat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"routes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"send_reconfigure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"stat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"use_src_mac_address": schema.BoolAttribute{
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

func (r *IPDHCPServerLeaseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPDHCPServerLeaseModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Address.IsNull() || plan.Address.IsUnknown()) {
		body["address"] = plan.Address.ValueString()
	}
	if !(plan.AddressList.IsNull() || plan.AddressList.IsUnknown()) {
		body["address-list"] = plan.AddressList.ValueString()
	}
	if !(plan.AgentCircuitID.IsNull() || plan.AgentCircuitID.IsUnknown()) {
		body["agent-circuit-id"] = plan.AgentCircuitID.ValueString()
	}
	if !(plan.AgentRemoteID.IsNull() || plan.AgentRemoteID.IsUnknown()) {
		body["agent-remote-id"] = plan.AgentRemoteID.ValueString()
	}
	if !(plan.AllowDualStackQueue.IsNull() || plan.AllowDualStackQueue.IsUnknown()) {
		body["allow-dual-stack-queue"] = plan.AllowDualStackQueue.ValueString()
	}
	if !(plan.AlwaysBroadcast.IsNull() || plan.AlwaysBroadcast.IsUnknown()) {
		body["always-broadcast"] = client.FormatBool(plan.AlwaysBroadcast.ValueBool())
	}
	if !(plan.BlockAccess.IsNull() || plan.BlockAccess.IsUnknown()) {
		body["block-access"] = client.FormatBool(plan.BlockAccess.ValueBool())
	}
	if !(plan.Blocked.IsNull() || plan.Blocked.IsUnknown()) {
		body["blocked"] = client.FormatBool(plan.Blocked.ValueBool())
	}
	if !(plan.CheckStatus.IsNull() || plan.CheckStatus.IsUnknown()) {
		body["check-status"] = plan.CheckStatus.ValueString()
	}
	if !(plan.ClientID.IsNull() || plan.ClientID.IsUnknown()) {
		body["client-id"] = plan.ClientID.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.DHCPOptionSet.IsNull() || plan.DHCPOptionSet.IsUnknown()) {
		body["dhcp-option-set"] = plan.DHCPOptionSet.ValueString()
	}
	if !(plan.DHCPOptions.IsNull() || plan.DHCPOptions.IsUnknown()) {
		body["dhcp-options"] = plan.DHCPOptions.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Dyn.IsNull() || plan.Dyn.IsUnknown()) {
		body["dyn"] = plan.Dyn.ValueString()
	}
	if !(plan.InsertQueueBefore.IsNull() || plan.InsertQueueBefore.IsUnknown()) {
		body["insert-queue-before"] = plan.InsertQueueBefore.ValueString()
	}
	if !(plan.LeaseTime.IsNull() || plan.LeaseTime.IsUnknown()) {
		body["lease-time"] = plan.LeaseTime.ValueString()
	}
	if !(plan.MACAddress.IsNull() || plan.MACAddress.IsUnknown()) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !(plan.MakeStatic.IsNull() || plan.MakeStatic.IsUnknown()) {
		body["make-static"] = plan.MakeStatic.ValueString()
	}
	if !(plan.ParentQueue.IsNull() || plan.ParentQueue.IsUnknown()) {
		body["parent-queue"] = plan.ParentQueue.ValueString()
	}
	if !(plan.Ping.IsNull() || plan.Ping.IsUnknown()) {
		body["ping"] = plan.Ping.ValueString()
	}
	if !(plan.QueueType.IsNull() || plan.QueueType.IsUnknown()) {
		body["queue-type"] = plan.QueueType.ValueString()
	}
	if !(plan.RADIUS.IsNull() || plan.RADIUS.IsUnknown()) {
		body["radius"] = client.FormatBool(plan.RADIUS.ValueBool())
	}
	if !(plan.RateLimit.IsNull() || plan.RateLimit.IsUnknown()) {
		body["rate-limit"] = plan.RateLimit.ValueString()
	}
	if !(plan.Rostat.IsNull() || plan.Rostat.IsUnknown()) {
		body["rostat"] = plan.Rostat.ValueString()
	}
	if !(plan.Routes.IsNull() || plan.Routes.IsUnknown()) {
		body["routes"] = plan.Routes.ValueString()
	}
	if !(plan.SendReconfigure.IsNull() || plan.SendReconfigure.IsUnknown()) {
		body["send-reconfigure"] = plan.SendReconfigure.ValueString()
	}
	if !(plan.Server.IsNull() || plan.Server.IsUnknown()) {
		body["server"] = plan.Server.ValueString()
	}
	if !(plan.Stat.IsNull() || plan.Stat.IsUnknown()) {
		body["stat"] = plan.Stat.ValueString()
	}
	if !(plan.UseSrcMACAddress.IsNull() || plan.UseSrcMACAddress.IsUnknown()) {
		body["use-src-mac-address"] = client.FormatBool(plan.UseSrcMACAddress.ValueBool())
	}
	obj, err := c.Add(ctx, "/ip/dhcp-server/lease", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/dhcp-server/lease failed", err.Error())
		return
	}
	iPDHCPServerLeaseApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPServerLeaseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPDHCPServerLeaseModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/dhcp-server/lease", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/dhcp-server/lease failed", err.Error())
		return
	}
	iPDHCPServerLeaseApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPDHCPServerLeaseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPDHCPServerLeaseModel
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
	if !plan.Address.Equal(state.Address) {
		body["address"] = plan.Address.ValueString()
	}
	if !plan.AddressList.Equal(state.AddressList) {
		body["address-list"] = plan.AddressList.ValueString()
	}
	if !plan.AgentCircuitID.Equal(state.AgentCircuitID) {
		body["agent-circuit-id"] = plan.AgentCircuitID.ValueString()
	}
	if !plan.AgentRemoteID.Equal(state.AgentRemoteID) {
		body["agent-remote-id"] = plan.AgentRemoteID.ValueString()
	}
	if !plan.AllowDualStackQueue.Equal(state.AllowDualStackQueue) {
		body["allow-dual-stack-queue"] = plan.AllowDualStackQueue.ValueString()
	}
	if !plan.AlwaysBroadcast.Equal(state.AlwaysBroadcast) {
		body["always-broadcast"] = client.FormatBool(plan.AlwaysBroadcast.ValueBool())
	}
	if !plan.BlockAccess.Equal(state.BlockAccess) {
		body["block-access"] = client.FormatBool(plan.BlockAccess.ValueBool())
	}
	if !plan.Blocked.Equal(state.Blocked) {
		body["blocked"] = client.FormatBool(plan.Blocked.ValueBool())
	}
	if !plan.CheckStatus.Equal(state.CheckStatus) {
		body["check-status"] = plan.CheckStatus.ValueString()
	}
	if !plan.ClientID.Equal(state.ClientID) {
		body["client-id"] = plan.ClientID.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.DHCPOptionSet.Equal(state.DHCPOptionSet) {
		body["dhcp-option-set"] = plan.DHCPOptionSet.ValueString()
	}
	if !plan.DHCPOptions.Equal(state.DHCPOptions) {
		body["dhcp-options"] = plan.DHCPOptions.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Dyn.Equal(state.Dyn) {
		body["dyn"] = plan.Dyn.ValueString()
	}
	if !plan.InsertQueueBefore.Equal(state.InsertQueueBefore) {
		body["insert-queue-before"] = plan.InsertQueueBefore.ValueString()
	}
	if !plan.LeaseTime.Equal(state.LeaseTime) {
		body["lease-time"] = plan.LeaseTime.ValueString()
	}
	if !plan.MACAddress.Equal(state.MACAddress) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.MakeStatic.Equal(state.MakeStatic) {
		body["make-static"] = plan.MakeStatic.ValueString()
	}
	if !plan.ParentQueue.Equal(state.ParentQueue) {
		body["parent-queue"] = plan.ParentQueue.ValueString()
	}
	if !plan.Ping.Equal(state.Ping) {
		body["ping"] = plan.Ping.ValueString()
	}
	if !plan.QueueType.Equal(state.QueueType) {
		body["queue-type"] = plan.QueueType.ValueString()
	}
	if !plan.RADIUS.Equal(state.RADIUS) {
		body["radius"] = client.FormatBool(plan.RADIUS.ValueBool())
	}
	if !plan.RateLimit.Equal(state.RateLimit) {
		body["rate-limit"] = plan.RateLimit.ValueString()
	}
	if !plan.Rostat.Equal(state.Rostat) {
		body["rostat"] = plan.Rostat.ValueString()
	}
	if !plan.Routes.Equal(state.Routes) {
		body["routes"] = plan.Routes.ValueString()
	}
	if !plan.SendReconfigure.Equal(state.SendReconfigure) {
		body["send-reconfigure"] = plan.SendReconfigure.ValueString()
	}
	if !plan.Server.Equal(state.Server) {
		body["server"] = plan.Server.ValueString()
	}
	if !plan.Stat.Equal(state.Stat) {
		body["stat"] = plan.Stat.ValueString()
	}
	if !plan.UseSrcMACAddress.Equal(state.UseSrcMACAddress) {
		body["use-src-mac-address"] = client.FormatBool(plan.UseSrcMACAddress.ValueBool())
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/dhcp-server/lease", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/dhcp-server/lease failed", err.Error())
			return
		}
		iPDHCPServerLeaseApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPServerLeaseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPDHCPServerLeaseModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/dhcp-server/lease", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/dhcp-server/lease failed", err.Error())
	}
}

func (r *IPDHCPServerLeaseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import formats accepted:
	//   *<id>                            -> bare RouterOS .id on the default router
	//   <router>/*<id>                   -> .id on the named router
	//   <router>/<naturalkey>            -> resolved via List + filter
	//   <naturalkey>                     -> resolved on the default router
	id := req.ID
	routerName := ""
	if i := strings.Index(id, "/"); i > 0 && !strings.HasPrefix(id, "*") {
		routerName, id = id[:i], id[i+1:]
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	if strings.HasPrefix(id, "*") {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(id))...)
		return
	}
	c := pickClient(r.reg, types.StringValue(routerName), &resp.Diagnostics)
	if c == nil {
		return
	}
	rows, err := iPDHCPServerLeaseLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/dhcp-server/lease matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPDHCPServerLeaseLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPDHCPServerLeaseLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/ip/dhcp-server/lease", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func iPDHCPServerLeaseApply(ctx context.Context, obj client.Object, m *IPDHCPServerLeaseModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["active-address"]; ok {
		_ = v
		if v != "" {
			m.ActiveAddress = types.StringValue(v)
		} else {
			m.ActiveAddress = types.StringNull()
		}
	} else {
		m.ActiveAddress = types.StringNull()
	}
	if v, ok := obj["active-agent-circuit-id"]; ok {
		_ = v
		if v != "" {
			m.ActiveAgentCircuitID = types.StringValue(v)
		} else {
			m.ActiveAgentCircuitID = types.StringNull()
		}
	} else {
		m.ActiveAgentCircuitID = types.StringNull()
	}
	if v, ok := obj["active-agent-remote-id"]; ok {
		_ = v
		if v != "" {
			m.ActiveAgentRemoteID = types.StringValue(v)
		} else {
			m.ActiveAgentRemoteID = types.StringNull()
		}
	} else {
		m.ActiveAgentRemoteID = types.StringNull()
	}
	if v, ok := obj["active-class-id"]; ok {
		_ = v
		if v != "" {
			m.ActiveClassID = types.StringValue(v)
		} else {
			m.ActiveClassID = types.StringNull()
		}
	} else {
		m.ActiveClassID = types.StringNull()
	}
	if v, ok := obj["active-client-id"]; ok {
		_ = v
		if v != "" {
			m.ActiveClientID = types.StringValue(v)
		} else {
			m.ActiveClientID = types.StringNull()
		}
	} else {
		m.ActiveClientID = types.StringNull()
	}
	if v, ok := obj["active-host-name"]; ok {
		_ = v
		if v != "" {
			m.ActiveHostName = types.StringValue(v)
		} else {
			m.ActiveHostName = types.StringNull()
		}
	} else {
		m.ActiveHostName = types.StringNull()
	}
	if v, ok := obj["active-mac-address"]; ok {
		_ = v
		if v != "" {
			m.ActiveMACAddress = types.StringValue(v)
		} else {
			m.ActiveMACAddress = types.StringNull()
		}
	} else {
		m.ActiveMACAddress = types.StringNull()
	}
	if v, ok := obj["active-server"]; ok {
		_ = v
		if v != "" {
			m.ActiveServer = types.StringValue(v)
		} else {
			m.ActiveServer = types.StringNull()
		}
	} else {
		m.ActiveServer = types.StringNull()
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
	if v, ok := obj["address-list"]; ok {
		_ = v
		if v != "" {
			m.AddressList = types.StringValue(v)
		} else {
			m.AddressList = types.StringNull()
		}
	} else {
		m.AddressList = types.StringNull()
	}
	if v, ok := obj["age"]; ok {
		_ = v
		if v != "" {
			m.Age = types.StringValue(v)
		} else {
			m.Age = types.StringNull()
		}
	} else {
		m.Age = types.StringNull()
	}
	if v, ok := obj["agent-circuit-id"]; ok {
		_ = v
		if v != "" {
			m.AgentCircuitID = types.StringValue(v)
		} else {
			m.AgentCircuitID = types.StringNull()
		}
	} else {
		m.AgentCircuitID = types.StringNull()
	}
	if v, ok := obj["agent-remote-id"]; ok {
		_ = v
		if v != "" {
			m.AgentRemoteID = types.StringValue(v)
		} else {
			m.AgentRemoteID = types.StringNull()
		}
	} else {
		m.AgentRemoteID = types.StringNull()
	}
	if v, ok := obj["allow-dual-stack-queue"]; ok {
		_ = v
		if v != "" {
			m.AllowDualStackQueue = types.StringValue(v)
		} else {
			m.AllowDualStackQueue = types.StringNull()
		}
	} else {
		m.AllowDualStackQueue = types.StringNull()
	}
	if v, ok := obj["always-broadcast"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AlwaysBroadcast = types.BoolValue(b)
		} else {
			m.AlwaysBroadcast = types.BoolNull()
		}
	} else {
		m.AlwaysBroadcast = types.BoolNull()
	}
	if v, ok := obj["block-access"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.BlockAccess = types.BoolValue(b)
		} else {
			m.BlockAccess = types.BoolNull()
		}
	} else {
		m.BlockAccess = types.BoolNull()
	}
	if v, ok := obj["blocked"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Blocked = types.BoolValue(b)
		} else {
			m.Blocked = types.BoolNull()
		}
	} else {
		m.Blocked = types.BoolNull()
	}
	if v, ok := obj["bridge-port"]; ok {
		_ = v
		if v != "" {
			m.BridgePort = types.StringValue(v)
		} else {
			m.BridgePort = types.StringNull()
		}
	} else {
		m.BridgePort = types.StringNull()
	}
	if v, ok := obj["check-status"]; ok {
		_ = v
		if v != "" {
			m.CheckStatus = types.StringValue(v)
		} else {
			m.CheckStatus = types.StringNull()
		}
	} else {
		m.CheckStatus = types.StringNull()
	}
	if v, ok := obj["client-id"]; ok {
		_ = v
		if v != "" {
			m.ClientID = types.StringValue(v)
		} else {
			m.ClientID = types.StringNull()
		}
	} else {
		m.ClientID = types.StringNull()
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
	if v, ok := obj["dhcp-option-set"]; ok {
		_ = v
		if v != "" {
			m.DHCPOptionSet = types.StringValue(v)
		} else {
			m.DHCPOptionSet = types.StringNull()
		}
	} else {
		m.DHCPOptionSet = types.StringNull()
	}
	if v, ok := obj["dhcp-options"]; ok {
		_ = v
		if v != "" {
			m.DHCPOptions = types.StringValue(v)
		} else {
			m.DHCPOptions = types.StringNull()
		}
	} else {
		m.DHCPOptions = types.StringNull()
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
	if v, ok := obj["dyn"]; ok {
		_ = v
		if v != "" {
			m.Dyn = types.StringValue(v)
		} else {
			m.Dyn = types.StringNull()
		}
	} else {
		m.Dyn = types.StringNull()
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
	if v, ok := obj["insert-queue-before"]; ok {
		_ = v
		if v != "" {
			m.InsertQueueBefore = types.StringValue(v)
		} else {
			m.InsertQueueBefore = types.StringNull()
		}
	} else {
		m.InsertQueueBefore = types.StringNull()
	}
	if v, ok := obj["last-seen"]; ok {
		_ = v
		if v != "" {
			m.LastSeen = types.StringValue(v)
		} else {
			m.LastSeen = types.StringNull()
		}
	} else {
		m.LastSeen = types.StringNull()
	}
	if v, ok := obj["last-sent-counter"]; ok {
		_ = v
		if v != "" {
			m.LastSentCounter = types.StringValue(v)
		} else {
			m.LastSentCounter = types.StringNull()
		}
	} else {
		m.LastSentCounter = types.StringNull()
	}
	if v, ok := obj["lease-time"]; ok {
		_ = v
		if v != "" {
			m.LeaseTime = types.StringValue(v)
		} else {
			m.LeaseTime = types.StringNull()
		}
	} else {
		m.LeaseTime = types.StringNull()
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
	if v, ok := obj["make-static"]; ok {
		_ = v
		if v != "" {
			m.MakeStatic = types.StringValue(v)
		} else {
			m.MakeStatic = types.StringNull()
		}
	} else {
		m.MakeStatic = types.StringNull()
	}
	if v, ok := obj["parent-queue"]; ok {
		_ = v
		if v != "" {
			m.ParentQueue = types.StringValue(v)
		} else {
			m.ParentQueue = types.StringNull()
		}
	} else {
		m.ParentQueue = types.StringNull()
	}
	if v, ok := obj["ping"]; ok {
		_ = v
		if v != "" {
			m.Ping = types.StringValue(v)
		} else {
			m.Ping = types.StringNull()
		}
	} else {
		m.Ping = types.StringNull()
	}
	if v, ok := obj["queue-type"]; ok {
		_ = v
		if v != "" {
			m.QueueType = types.StringValue(v)
		} else {
			m.QueueType = types.StringNull()
		}
	} else {
		m.QueueType = types.StringNull()
	}
	if v, ok := obj["radius"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.RADIUS = types.BoolValue(b)
		} else {
			m.RADIUS = types.BoolNull()
		}
	} else {
		m.RADIUS = types.BoolNull()
	}
	if v, ok := obj["rate-limit"]; ok {
		_ = v
		if v != "" {
			m.RateLimit = types.StringValue(v)
		} else {
			m.RateLimit = types.StringNull()
		}
	} else {
		m.RateLimit = types.StringNull()
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
	if v, ok := obj["reconfigure-status"]; ok {
		_ = v
		if v != "" {
			m.ReconfigureStatus = types.StringValue(v)
		} else {
			m.ReconfigureStatus = types.StringNull()
		}
	} else {
		m.ReconfigureStatus = types.StringNull()
	}
	if v, ok := obj["rostat"]; ok {
		_ = v
		if v != "" {
			m.Rostat = types.StringValue(v)
		} else {
			m.Rostat = types.StringNull()
		}
	} else {
		m.Rostat = types.StringNull()
	}
	if v, ok := obj["routes"]; ok {
		_ = v
		if v != "" {
			m.Routes = types.StringValue(v)
		} else {
			m.Routes = types.StringNull()
		}
	} else {
		m.Routes = types.StringNull()
	}
	if v, ok := obj["send-reconfigure"]; ok {
		_ = v
		if v != "" {
			m.SendReconfigure = types.StringValue(v)
		} else {
			m.SendReconfigure = types.StringNull()
		}
	} else {
		m.SendReconfigure = types.StringNull()
	}
	if v, ok := obj["server"]; ok {
		_ = v
		if v != "" {
			m.Server = types.StringValue(v)
		} else {
			m.Server = types.StringNull()
		}
	} else {
		m.Server = types.StringNull()
	}
	if v, ok := obj["src-mac-address"]; ok {
		_ = v
		if v != "" {
			m.SrcMACAddress = types.StringValue(v)
		} else {
			m.SrcMACAddress = types.StringNull()
		}
	} else {
		m.SrcMACAddress = types.StringNull()
	}
	if v, ok := obj["stat"]; ok {
		_ = v
		if v != "" {
			m.Stat = types.StringValue(v)
		} else {
			m.Stat = types.StringNull()
		}
	} else {
		m.Stat = types.StringNull()
	}
	if v, ok := obj["use-src-mac-address"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.UseSrcMACAddress = types.BoolValue(b)
		} else {
			m.UseSrcMACAddress = types.BoolNull()
		}
	} else {
		m.UseSrcMACAddress = types.BoolNull()
	}
}
