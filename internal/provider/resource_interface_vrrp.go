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
	_ resource.Resource                = &InterfaceVRRPResource{}
	_ resource.ResourceWithImportState = &InterfaceVRRPResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceVRRPResource struct {
	reg *client.Registry
}

type InterfaceVRRPModel struct {
	ID                     types.String `tfsdk:"id"`
	Vrid                   types.String `tfsdk:"vrid"`
	Version                types.String `tfsdk:"version"`
	V3Protocol             types.String `tfsdk:"v3_protocol"`
	SyncConnectionTracking types.String `tfsdk:"sync_connection_tracking"`
	PreemptionMode         types.String `tfsdk:"preemption_mode"`
	OnMaster               types.String `tfsdk:"on_master"`
	OnFail                 types.String `tfsdk:"on_fail"`
	OnBackup               types.String `tfsdk:"on_backup"`
	GroupMaster            types.String `tfsdk:"group_master"`
	GroupAuthority         types.String `tfsdk:"group_authority"`
	ConnectionTrackingPort types.String `tfsdk:"connection_tracking_port"`
	ConnectionTrackingMode types.String `tfsdk:"connection_tracking_mode"`
	ARP                    types.String `tfsdk:"arp"`
	ARPTimeout             types.String `tfsdk:"arp_timeout"`
	Authentication         types.String `tfsdk:"authentication"`
	Comment                types.String `tfsdk:"comment"`
	Disabled               types.Bool   `tfsdk:"disabled"`
	Interface              types.String `tfsdk:"interface"`
	Interval               types.String `tfsdk:"interval"`
	Name                   types.String `tfsdk:"name"`
	Password               types.String `tfsdk:"password"`
	Priority               types.String `tfsdk:"priority"`
	RemoteAddress          types.String `tfsdk:"remote_address"`
	Router                 types.String `tfsdk:"router"`
}

func NewInterfaceVRRPResource() resource.Resource { return &InterfaceVRRPResource{} }

func (r *InterfaceVRRPResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_vrrp"
}

func (r *InterfaceVRRPResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceVRRPResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/vrrp`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"vrid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `vrid`.",
			},
			"version": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `version`.",
			},
			"v3_protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `v3-protocol`.",
			},
			"sync_connection_tracking": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `sync-connection-tracking`.",
			},
			"preemption_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `preemption-mode`.",
			},
			"on_master": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `on-master`.",
			},
			"on_fail": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `on-fail`.",
			},
			"on_backup": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `on-backup`.",
			},
			"group_master": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `group-master`.",
			},
			"group_authority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `group-authority`.",
			},
			"connection_tracking_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `connection-tracking-port`.",
			},
			"connection_tracking_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `connection-tracking-mode`.",
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
			"authentication": schema.StringAttribute{
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
				Required:    true,
				Description: "",
			},
			"interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "",
			},
			"priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_address": schema.StringAttribute{
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

func (r *InterfaceVRRPResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceVRRPModel
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
	if !(plan.Authentication.IsNull() || plan.Authentication.IsUnknown()) {
		body["authentication"] = plan.Authentication.ValueString()
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
	if !(plan.Interval.IsNull() || plan.Interval.IsUnknown()) {
		body["interval"] = plan.Interval.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Password.IsNull() || plan.Password.IsUnknown()) {
		body["password"] = plan.Password.ValueString()
	}
	if !(plan.Priority.IsNull() || plan.Priority.IsUnknown()) {
		body["priority"] = plan.Priority.ValueString()
	}
	if !(plan.RemoteAddress.IsNull() || plan.RemoteAddress.IsUnknown()) {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !(plan.ConnectionTrackingMode.IsNull() || plan.ConnectionTrackingMode.IsUnknown()) {
		body["connection-tracking-mode"] = plan.ConnectionTrackingMode.ValueString()
	}
	if !(plan.ConnectionTrackingPort.IsNull() || plan.ConnectionTrackingPort.IsUnknown()) {
		body["connection-tracking-port"] = plan.ConnectionTrackingPort.ValueString()
	}
	if !(plan.GroupAuthority.IsNull() || plan.GroupAuthority.IsUnknown()) {
		body["group-authority"] = plan.GroupAuthority.ValueString()
	}
	if !(plan.GroupMaster.IsNull() || plan.GroupMaster.IsUnknown()) {
		body["group-master"] = plan.GroupMaster.ValueString()
	}
	if !(plan.OnBackup.IsNull() || plan.OnBackup.IsUnknown()) {
		body["on-backup"] = plan.OnBackup.ValueString()
	}
	if !(plan.OnFail.IsNull() || plan.OnFail.IsUnknown()) {
		body["on-fail"] = plan.OnFail.ValueString()
	}
	if !(plan.OnMaster.IsNull() || plan.OnMaster.IsUnknown()) {
		body["on-master"] = plan.OnMaster.ValueString()
	}
	if !(plan.PreemptionMode.IsNull() || plan.PreemptionMode.IsUnknown()) {
		body["preemption-mode"] = plan.PreemptionMode.ValueString()
	}
	if !(plan.SyncConnectionTracking.IsNull() || plan.SyncConnectionTracking.IsUnknown()) {
		body["sync-connection-tracking"] = plan.SyncConnectionTracking.ValueString()
	}
	if !(plan.V3Protocol.IsNull() || plan.V3Protocol.IsUnknown()) {
		body["v3-protocol"] = plan.V3Protocol.ValueString()
	}
	if !(plan.Version.IsNull() || plan.Version.IsUnknown()) {
		body["version"] = plan.Version.ValueString()
	}
	if !(plan.Vrid.IsNull() || plan.Vrid.IsUnknown()) {
		body["vrid"] = plan.Vrid.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/vrrp", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/vrrp failed", err.Error())
		return
	}
	interfaceVRRPApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceVRRPResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceVRRPModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/vrrp", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/vrrp failed", err.Error())
		return
	}
	interfaceVRRPApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceVRRPResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceVRRPModel
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
	if !plan.Authentication.Equal(state.Authentication) {
		body["authentication"] = plan.Authentication.ValueString()
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
	if !plan.Interval.Equal(state.Interval) {
		body["interval"] = plan.Interval.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Password.Equal(state.Password) {
		body["password"] = plan.Password.ValueString()
	}
	if !plan.Priority.Equal(state.Priority) {
		body["priority"] = plan.Priority.ValueString()
	}
	if !plan.RemoteAddress.Equal(state.RemoteAddress) {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !plan.ConnectionTrackingMode.Equal(state.ConnectionTrackingMode) && !plan.ConnectionTrackingMode.IsUnknown() {
		body["connection-tracking-mode"] = plan.ConnectionTrackingMode.ValueString()
	}
	if !plan.ConnectionTrackingPort.Equal(state.ConnectionTrackingPort) && !plan.ConnectionTrackingPort.IsUnknown() {
		body["connection-tracking-port"] = plan.ConnectionTrackingPort.ValueString()
	}
	if !plan.GroupAuthority.Equal(state.GroupAuthority) && !plan.GroupAuthority.IsUnknown() {
		body["group-authority"] = plan.GroupAuthority.ValueString()
	}
	if !plan.GroupMaster.Equal(state.GroupMaster) && !plan.GroupMaster.IsUnknown() {
		body["group-master"] = plan.GroupMaster.ValueString()
	}
	if !plan.OnBackup.Equal(state.OnBackup) && !plan.OnBackup.IsUnknown() {
		body["on-backup"] = plan.OnBackup.ValueString()
	}
	if !plan.OnFail.Equal(state.OnFail) && !plan.OnFail.IsUnknown() {
		body["on-fail"] = plan.OnFail.ValueString()
	}
	if !plan.OnMaster.Equal(state.OnMaster) && !plan.OnMaster.IsUnknown() {
		body["on-master"] = plan.OnMaster.ValueString()
	}
	if !plan.PreemptionMode.Equal(state.PreemptionMode) && !plan.PreemptionMode.IsUnknown() {
		body["preemption-mode"] = plan.PreemptionMode.ValueString()
	}
	if !plan.SyncConnectionTracking.Equal(state.SyncConnectionTracking) && !plan.SyncConnectionTracking.IsUnknown() {
		body["sync-connection-tracking"] = plan.SyncConnectionTracking.ValueString()
	}
	if !plan.V3Protocol.Equal(state.V3Protocol) && !plan.V3Protocol.IsUnknown() {
		body["v3-protocol"] = plan.V3Protocol.ValueString()
	}
	if !plan.Version.Equal(state.Version) && !plan.Version.IsUnknown() {
		body["version"] = plan.Version.ValueString()
	}
	if !plan.Vrid.Equal(state.Vrid) && !plan.Vrid.IsUnknown() {
		body["vrid"] = plan.Vrid.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/vrrp", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/vrrp failed", err.Error())
			return
		}
		interfaceVRRPApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceVRRPResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceVRRPModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/vrrp", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/vrrp failed", err.Error())
	}
}

func (r *InterfaceVRRPResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceVRRPLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/vrrp matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceVRRPLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceVRRPLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/vrrp", id)
}

func interfaceVRRPApply(ctx context.Context, obj client.Object, m *InterfaceVRRPModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["vrid"]; ok && v != "" {
		m.Vrid = types.StringValue(v)
	} else {
		m.Vrid = types.StringNull()
	}
	if v, ok := obj["version"]; ok && v != "" {
		m.Version = types.StringValue(v)
	} else {
		m.Version = types.StringNull()
	}
	if v, ok := obj["v3-protocol"]; ok && v != "" {
		m.V3Protocol = types.StringValue(v)
	} else {
		m.V3Protocol = types.StringNull()
	}
	if v, ok := obj["sync-connection-tracking"]; ok && v != "" {
		m.SyncConnectionTracking = types.StringValue(v)
	} else {
		m.SyncConnectionTracking = types.StringNull()
	}
	if v, ok := obj["preemption-mode"]; ok && v != "" {
		m.PreemptionMode = types.StringValue(v)
	} else {
		m.PreemptionMode = types.StringNull()
	}
	if v, ok := obj["on-master"]; ok && v != "" {
		m.OnMaster = types.StringValue(v)
	} else {
		m.OnMaster = types.StringNull()
	}
	if v, ok := obj["on-fail"]; ok && v != "" {
		m.OnFail = types.StringValue(v)
	} else {
		m.OnFail = types.StringNull()
	}
	if v, ok := obj["on-backup"]; ok && v != "" {
		m.OnBackup = types.StringValue(v)
	} else {
		m.OnBackup = types.StringNull()
	}
	if v, ok := obj["group-master"]; ok && v != "" {
		m.GroupMaster = types.StringValue(v)
	} else {
		m.GroupMaster = types.StringNull()
	}
	if v, ok := obj["group-authority"]; ok && v != "" {
		m.GroupAuthority = types.StringValue(v)
	} else {
		m.GroupAuthority = types.StringNull()
	}
	if v, ok := obj["connection-tracking-port"]; ok && v != "" {
		m.ConnectionTrackingPort = types.StringValue(v)
	} else {
		m.ConnectionTrackingPort = types.StringNull()
	}
	if v, ok := obj["connection-tracking-mode"]; ok && v != "" {
		m.ConnectionTrackingMode = types.StringValue(v)
	} else {
		m.ConnectionTrackingMode = types.StringNull()
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
	if v, ok := obj["authentication"]; ok {
		_ = v
		if v != "" {
			m.Authentication = types.StringValue(v)
		} else {
			m.Authentication = types.StringNull()
		}
	} else {
		m.Authentication = types.StringNull()
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
	if v, ok := obj["interval"]; ok {
		_ = v
		if v != "" {
			m.Interval = types.StringValue(v)
		} else {
			m.Interval = types.StringNull()
		}
	} else {
		m.Interval = types.StringNull()
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
	if v, ok := obj["password"]; ok {
		_ = v
		if v != "" {
			m.Password = types.StringValue(v)
		} else {
			m.Password = types.StringNull()
		}
	} else {
		m.Password = types.StringNull()
	}
	if v, ok := obj["priority"]; ok {
		_ = v
		if v != "" {
			m.Priority = types.StringValue(v)
		} else {
			m.Priority = types.StringNull()
		}
	} else {
		m.Priority = types.StringNull()
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
