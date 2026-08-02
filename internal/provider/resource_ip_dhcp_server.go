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
	_ resource.Resource                = &IPDHCPServerResource{}
	_ resource.ResourceWithImportState = &IPDHCPServerResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPDHCPServerResource struct {
	reg *client.Registry
}

type IPDHCPServerModel struct {
	ID                            types.String    `tfsdk:"id"`
	SupportBroadbandTr101         types.String    `tfsdk:"support_broadband_tr101"`
	AddressLists                  types.String    `tfsdk:"address_lists"`
	AddDnsEntriesSuffix           types.String    `tfsdk:"add_dns_entries_suffix"`
	AddDnsEntries                 types.String    `tfsdk:"add_dns_entries"`
	AddArp                        boolStringValue `tfsdk:"add_arp"`
	AddARPForLeases               types.Bool      `tfsdk:"add_arp_for_leases"`
	AddressList                   types.String    `tfsdk:"address_list"`
	AddressPool                   types.String    `tfsdk:"address_pool"`
	AllowDualStackQueue           types.Bool      `tfsdk:"allow_dual_stack_queue"`
	AlwaysBroadcast               types.Bool      `tfsdk:"always_broadcast"`
	Authoritative                 types.String    `tfsdk:"authoritative"`
	BootpLeaseTime                types.String    `tfsdk:"bootp_lease_time"`
	BootpSupport                  types.String    `tfsdk:"bootp_support"`
	ClientMACLimit                types.String    `tfsdk:"client_mac_limit"`
	Comment                       types.String    `tfsdk:"comment"`
	ConflictDetection             types.Bool      `tfsdk:"conflict_detection"`
	DelayThreshold                types.String    `tfsdk:"delay_threshold"`
	DHCPOptionSet                 types.String    `tfsdk:"dhcp_option_set"`
	Disabled                      types.Bool      `tfsdk:"disabled"`
	DynamicLeaseIdentifiers       types.String    `tfsdk:"dynamic_lease_identifiers"`
	Dynbootp                      types.String    `tfsdk:"dynbootp"`
	InsertQueueBefore             types.String    `tfsdk:"insert_queue_before"`
	Interface                     types.String    `tfsdk:"interface"`
	Invalid                       types.Bool      `tfsdk:"invalid"`
	LeaseScript                   types.String    `tfsdk:"lease_script"`
	LeaseTime                     types.String    `tfsdk:"lease_time"`
	Name                          types.String    `tfsdk:"name"`
	ParentQueue                   types.String    `tfsdk:"parent_queue"`
	Relay                         types.String    `tfsdk:"relay"`
	ServerAddress                 types.String    `tfsdk:"server_address"`
	SupportTheBroadbandForumTr101 types.Bool      `tfsdk:"support_the_broadband_forum_tr_101"`
	UseFramedAsClassless          types.Bool      `tfsdk:"use_framed_as_classless"`
	UseRADIUS                     types.String    `tfsdk:"use_radius"`
	UseReconfigure                types.Bool      `tfsdk:"use_reconfigure"`
	Router                        types.String    `tfsdk:"router"`
}

func NewIPDHCPServerResource() resource.Resource { return &IPDHCPServerResource{} }

func (r *IPDHCPServerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_dhcp_server"
}

func (r *IPDHCPServerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPDHCPServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/dhcp-server`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"support_broadband_tr101": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `support-broadband-tr101`.",
			},
			"address_lists": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `address-lists`.",
			},
			"add_dns_entries_suffix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `add-dns-entries-suffix`.",
			},
			"add_dns_entries": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `add-dns-entries`.",
			},
			"add_arp": schema.StringAttribute{
				CustomType:  boolStringType{},
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `add-arp`.",
			},
			"add_arp_for_leases": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"address_list": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"address_pool": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"allow_dual_stack_queue": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"always_broadcast": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"authoritative": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"yes", "after-2s-delay", "after-10s-delay", "no"}...)},
			},
			"bootp_lease_time": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"bootp_support": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"none", "static", "dynamic"}...)},
			},
			"client_mac_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "A number, or `unlimited`.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"conflict_detection": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"delay_threshold": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"dhcp_option_set": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"dynamic_lease_identifiers": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dynbootp": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"insert_queue_before": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interface": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"lease_script": schema.StringAttribute{
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
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"parent_queue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"relay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"server_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"support_the_broadband_forum_tr_101": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"use_framed_as_classless": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"use_radius": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "yes", "accounting"}...)},
			},
			"use_reconfigure": schema.BoolAttribute{
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

func (r *IPDHCPServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPDHCPServerModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AddressPool.IsNull() || plan.AddressPool.IsUnknown()) {
		body["address-pool"] = plan.AddressPool.ValueString()
	}
	if !(plan.AllowDualStackQueue.IsNull() || plan.AllowDualStackQueue.IsUnknown()) {
		body["allow-dual-stack-queue"] = client.FormatBool(plan.AllowDualStackQueue.ValueBool())
	}
	if !(plan.AlwaysBroadcast.IsNull() || plan.AlwaysBroadcast.IsUnknown()) {
		body["always-broadcast"] = client.FormatBool(plan.AlwaysBroadcast.ValueBool())
	}
	if !(plan.Authoritative.IsNull() || plan.Authoritative.IsUnknown()) {
		body["authoritative"] = plan.Authoritative.ValueString()
	}
	if !(plan.BootpLeaseTime.IsNull() || plan.BootpLeaseTime.IsUnknown()) {
		body["bootp-lease-time"] = plan.BootpLeaseTime.ValueString()
	}
	if !(plan.BootpSupport.IsNull() || plan.BootpSupport.IsUnknown()) {
		body["bootp-support"] = plan.BootpSupport.ValueString()
	}
	if !(plan.ClientMACLimit.IsNull() || plan.ClientMACLimit.IsUnknown()) {
		body["client-mac-limit"] = plan.ClientMACLimit.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.ConflictDetection.IsNull() || plan.ConflictDetection.IsUnknown()) {
		body["conflict-detection"] = client.FormatBool(plan.ConflictDetection.ValueBool())
	}
	if !(plan.DelayThreshold.IsNull() || plan.DelayThreshold.IsUnknown()) {
		body["delay-threshold"] = plan.DelayThreshold.ValueString()
	}
	if !(plan.DHCPOptionSet.IsNull() || plan.DHCPOptionSet.IsUnknown()) {
		body["dhcp-option-set"] = plan.DHCPOptionSet.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.DynamicLeaseIdentifiers.IsNull() || plan.DynamicLeaseIdentifiers.IsUnknown()) {
		body["dynamic-lease-identifiers"] = plan.DynamicLeaseIdentifiers.ValueString()
	}
	if !(plan.InsertQueueBefore.IsNull() || plan.InsertQueueBefore.IsUnknown()) {
		body["insert-queue-before"] = plan.InsertQueueBefore.ValueString()
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.LeaseScript.IsNull() || plan.LeaseScript.IsUnknown()) {
		body["lease-script"] = plan.LeaseScript.ValueString()
	}
	if !(plan.LeaseTime.IsNull() || plan.LeaseTime.IsUnknown()) {
		body["lease-time"] = plan.LeaseTime.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.ParentQueue.IsNull() || plan.ParentQueue.IsUnknown()) {
		body["parent-queue"] = plan.ParentQueue.ValueString()
	}
	if !(plan.Relay.IsNull() || plan.Relay.IsUnknown()) {
		body["relay"] = plan.Relay.ValueString()
	}
	if !(plan.ServerAddress.IsNull() || plan.ServerAddress.IsUnknown()) {
		body["server-address"] = plan.ServerAddress.ValueString()
	}
	if !(plan.UseFramedAsClassless.IsNull() || plan.UseFramedAsClassless.IsUnknown()) {
		body["use-framed-as-classless"] = client.FormatBool(plan.UseFramedAsClassless.ValueBool())
	}
	if !(plan.UseRADIUS.IsNull() || plan.UseRADIUS.IsUnknown()) {
		body["use-radius"] = plan.UseRADIUS.ValueString()
	}
	if !(plan.UseReconfigure.IsNull() || plan.UseReconfigure.IsUnknown()) {
		body["use-reconfigure"] = client.FormatBool(plan.UseReconfigure.ValueBool())
	}
	if !(plan.AddArp.IsNull() || plan.AddArp.IsUnknown()) {
		body["add-arp"] = plan.AddArp.ValueString()
	}
	if !(plan.AddDnsEntries.IsNull() || plan.AddDnsEntries.IsUnknown()) {
		body["add-dns-entries"] = plan.AddDnsEntries.ValueString()
	}
	if !(plan.AddDnsEntriesSuffix.IsNull() || plan.AddDnsEntriesSuffix.IsUnknown()) {
		body["add-dns-entries-suffix"] = plan.AddDnsEntriesSuffix.ValueString()
	}
	if !(plan.AddressLists.IsNull() || plan.AddressLists.IsUnknown()) {
		body["address-lists"] = plan.AddressLists.ValueString()
	}
	if !(plan.SupportBroadbandTr101.IsNull() || plan.SupportBroadbandTr101.IsUnknown()) {
		body["support-broadband-tr101"] = plan.SupportBroadbandTr101.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/dhcp-server", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/dhcp-server failed", err.Error())
		return
	}
	iPDHCPServerApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPDHCPServerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/dhcp-server", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/dhcp-server failed", err.Error())
		return
	}
	iPDHCPServerApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPDHCPServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPDHCPServerModel
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
	if !plan.AddressPool.Equal(state.AddressPool) && !plan.AddressPool.IsUnknown() {
		body["address-pool"] = plan.AddressPool.ValueString()
	}
	if !plan.AllowDualStackQueue.Equal(state.AllowDualStackQueue) && !plan.AllowDualStackQueue.IsUnknown() {
		body["allow-dual-stack-queue"] = client.FormatBool(plan.AllowDualStackQueue.ValueBool())
	}
	if !plan.AlwaysBroadcast.Equal(state.AlwaysBroadcast) && !plan.AlwaysBroadcast.IsUnknown() {
		body["always-broadcast"] = client.FormatBool(plan.AlwaysBroadcast.ValueBool())
	}
	if !plan.Authoritative.Equal(state.Authoritative) && !plan.Authoritative.IsUnknown() {
		body["authoritative"] = plan.Authoritative.ValueString()
	}
	if !plan.BootpLeaseTime.Equal(state.BootpLeaseTime) && !plan.BootpLeaseTime.IsUnknown() {
		body["bootp-lease-time"] = plan.BootpLeaseTime.ValueString()
	}
	if !plan.BootpSupport.Equal(state.BootpSupport) && !plan.BootpSupport.IsUnknown() {
		body["bootp-support"] = plan.BootpSupport.ValueString()
	}
	if !plan.ClientMACLimit.Equal(state.ClientMACLimit) && !plan.ClientMACLimit.IsUnknown() {
		body["client-mac-limit"] = plan.ClientMACLimit.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.ConflictDetection.Equal(state.ConflictDetection) && !plan.ConflictDetection.IsUnknown() {
		body["conflict-detection"] = client.FormatBool(plan.ConflictDetection.ValueBool())
	}
	if !plan.DelayThreshold.Equal(state.DelayThreshold) && !plan.DelayThreshold.IsUnknown() {
		body["delay-threshold"] = plan.DelayThreshold.ValueString()
	}
	if !plan.DHCPOptionSet.Equal(state.DHCPOptionSet) && !plan.DHCPOptionSet.IsUnknown() {
		body["dhcp-option-set"] = plan.DHCPOptionSet.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DynamicLeaseIdentifiers.Equal(state.DynamicLeaseIdentifiers) && !plan.DynamicLeaseIdentifiers.IsUnknown() {
		body["dynamic-lease-identifiers"] = plan.DynamicLeaseIdentifiers.ValueString()
	}
	if !plan.InsertQueueBefore.Equal(state.InsertQueueBefore) && !plan.InsertQueueBefore.IsUnknown() {
		body["insert-queue-before"] = plan.InsertQueueBefore.ValueString()
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.LeaseScript.Equal(state.LeaseScript) && !plan.LeaseScript.IsUnknown() {
		body["lease-script"] = plan.LeaseScript.ValueString()
	}
	if !plan.LeaseTime.Equal(state.LeaseTime) && !plan.LeaseTime.IsUnknown() {
		body["lease-time"] = plan.LeaseTime.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.ParentQueue.Equal(state.ParentQueue) && !plan.ParentQueue.IsUnknown() {
		body["parent-queue"] = plan.ParentQueue.ValueString()
	}
	if !plan.Relay.Equal(state.Relay) && !plan.Relay.IsUnknown() {
		body["relay"] = plan.Relay.ValueString()
	}
	if !plan.ServerAddress.Equal(state.ServerAddress) && !plan.ServerAddress.IsUnknown() {
		body["server-address"] = plan.ServerAddress.ValueString()
	}
	if !plan.UseFramedAsClassless.Equal(state.UseFramedAsClassless) && !plan.UseFramedAsClassless.IsUnknown() {
		body["use-framed-as-classless"] = client.FormatBool(plan.UseFramedAsClassless.ValueBool())
	}
	if !plan.UseRADIUS.Equal(state.UseRADIUS) && !plan.UseRADIUS.IsUnknown() {
		body["use-radius"] = plan.UseRADIUS.ValueString()
	}
	if !plan.UseReconfigure.Equal(state.UseReconfigure) && !plan.UseReconfigure.IsUnknown() {
		body["use-reconfigure"] = client.FormatBool(plan.UseReconfigure.ValueBool())
	}
	if !plan.AddArp.Equal(state.AddArp) && !plan.AddArp.IsUnknown() {
		body["add-arp"] = plan.AddArp.ValueString()
	}
	if !plan.AddDnsEntries.Equal(state.AddDnsEntries) && !plan.AddDnsEntries.IsUnknown() {
		body["add-dns-entries"] = plan.AddDnsEntries.ValueString()
	}
	if !plan.AddDnsEntriesSuffix.Equal(state.AddDnsEntriesSuffix) && !plan.AddDnsEntriesSuffix.IsUnknown() {
		body["add-dns-entries-suffix"] = plan.AddDnsEntriesSuffix.ValueString()
	}
	if !plan.AddressLists.Equal(state.AddressLists) && !plan.AddressLists.IsUnknown() {
		body["address-lists"] = plan.AddressLists.ValueString()
	}
	if !plan.SupportBroadbandTr101.Equal(state.SupportBroadbandTr101) && !plan.SupportBroadbandTr101.IsUnknown() {
		body["support-broadband-tr101"] = plan.SupportBroadbandTr101.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/dhcp-server", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/dhcp-server failed", err.Error())
			return
		}
		iPDHCPServerApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPDHCPServerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/dhcp-server", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/dhcp-server failed", err.Error())
	}
}

func (r *IPDHCPServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPDHCPServerLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/dhcp-server matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPDHCPServerLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPDHCPServerLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/dhcp-server", id)
}

func iPDHCPServerApply(ctx context.Context, obj client.Object, m *IPDHCPServerModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["support-broadband-tr101"]; ok && v != "" {
		m.SupportBroadbandTr101 = types.StringValue(v)
	} else {
		m.SupportBroadbandTr101 = types.StringNull()
	}
	if v, ok := obj["address-lists"]; ok && v != "" {
		m.AddressLists = types.StringValue(v)
	} else {
		m.AddressLists = types.StringNull()
	}
	if v, ok := obj["add-dns-entries-suffix"]; ok && v != "" {
		m.AddDnsEntriesSuffix = types.StringValue(v)
	} else {
		m.AddDnsEntriesSuffix = types.StringNull()
	}
	if v, ok := obj["add-dns-entries"]; ok && v != "" {
		m.AddDnsEntries = types.StringValue(v)
	} else {
		m.AddDnsEntries = types.StringNull()
	}
	if v, ok := obj["add-arp"]; ok && v != "" {
		m.AddArp = newBoolStringValue(v)
	} else {
		m.AddArp = newBoolStringNull()
	}
	if v, ok := obj["add-arp-for-leases"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.AddARPForLeases = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.AddARPForLeases = types.BoolValue(true)
		} else {
			m.AddARPForLeases = types.BoolNull()
		}
	}
	if v, ok := obj["address-list"]; ok {
		if v != "" {
			m.AddressList = types.StringValue(v)
		} else {
			m.AddressList = types.StringNull()
		}
	}
	if v, ok := obj["address-pool"]; ok {
		if v != "" {
			m.AddressPool = types.StringValue(v)
		} else {
			m.AddressPool = types.StringNull()
		}
	}
	if v, ok := obj["allow-dual-stack-queue"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.AllowDualStackQueue = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.AllowDualStackQueue = types.BoolValue(true)
		} else {
			m.AllowDualStackQueue = types.BoolNull()
		}
	}
	if v, ok := obj["always-broadcast"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.AlwaysBroadcast = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.AlwaysBroadcast = types.BoolValue(true)
		} else {
			m.AlwaysBroadcast = types.BoolNull()
		}
	}
	if v, ok := obj["authoritative"]; ok {
		_ = v
		if v != "" {
			m.Authoritative = types.StringValue(v)
		}
	}
	if v, ok := obj["bootp-lease-time"]; ok {
		if v != "" {
			m.BootpLeaseTime = types.StringValue(v)
		} else {
			m.BootpLeaseTime = types.StringNull()
		}
	}
	if v, ok := obj["bootp-support"]; ok {
		if v != "" {
			m.BootpSupport = types.StringValue(v)
		} else {
			m.BootpSupport = types.StringNull()
		}
	}
	if v, ok := obj["client-mac-limit"]; ok {
		if v != "" {
			m.ClientMACLimit = types.StringValue(v)
		} else {
			m.ClientMACLimit = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["conflict-detection"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.ConflictDetection = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.ConflictDetection = types.BoolValue(true)
		} else {
			m.ConflictDetection = types.BoolNull()
		}
	}
	if v, ok := obj["delay-threshold"]; ok {
		if v != "" {
			m.DelayThreshold = types.StringValue(v)
		} else {
			m.DelayThreshold = types.StringNull()
		}
	}
	if v, ok := obj["dhcp-option-set"]; ok {
		if v != "" {
			m.DHCPOptionSet = types.StringValue(v)
		} else {
			m.DHCPOptionSet = types.StringNull()
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
	if v, ok := obj["dynamic-lease-identifiers"]; ok {
		if v != "" {
			m.DynamicLeaseIdentifiers = types.StringValue(v)
		} else {
			m.DynamicLeaseIdentifiers = types.StringNull()
		}
	}
	if v, ok := obj["dynbootp"]; ok {
		if v != "" {
			m.Dynbootp = types.StringValue(v)
		} else {
			m.Dynbootp = types.StringNull()
		}
	}
	if v, ok := obj["insert-queue-before"]; ok {
		if v != "" {
			m.InsertQueueBefore = types.StringValue(v)
		} else {
			m.InsertQueueBefore = types.StringNull()
		}
	}
	if v, ok := obj["interface"]; ok {
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	}
	if v, ok := obj["invalid"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Invalid = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Invalid = types.BoolValue(true)
		} else {
			m.Invalid = types.BoolNull()
		}
	}
	if v, ok := obj["lease-script"]; ok {
		if v != "" {
			m.LeaseScript = types.StringValue(v)
		} else {
			m.LeaseScript = types.StringNull()
		}
	}
	if v, ok := obj["lease-time"]; ok {
		if v != "" {
			m.LeaseTime = types.StringValue(v)
		} else {
			m.LeaseTime = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["parent-queue"]; ok {
		if v != "" {
			m.ParentQueue = types.StringValue(v)
		} else {
			m.ParentQueue = types.StringNull()
		}
	}
	if v, ok := obj["relay"]; ok {
		if v != "" {
			m.Relay = types.StringValue(v)
		} else {
			m.Relay = types.StringNull()
		}
	}
	if v, ok := obj["server-address"]; ok {
		if v != "" {
			m.ServerAddress = types.StringValue(v)
		} else {
			m.ServerAddress = types.StringNull()
		}
	}
	if v, ok := obj["support-the-broadband-forum-tr-101"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SupportTheBroadbandForumTr101 = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.SupportTheBroadbandForumTr101 = types.BoolValue(true)
		} else {
			m.SupportTheBroadbandForumTr101 = types.BoolNull()
		}
	}
	if v, ok := obj["use-framed-as-classless"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.UseFramedAsClassless = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.UseFramedAsClassless = types.BoolValue(true)
		} else {
			m.UseFramedAsClassless = types.BoolNull()
		}
	}
	if v, ok := obj["use-radius"]; ok {
		if v != "" {
			m.UseRADIUS = types.StringValue(v)
		} else {
			m.UseRADIUS = types.StringNull()
		}
	}
	if v, ok := obj["use-reconfigure"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.UseReconfigure = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.UseReconfigure = types.BoolValue(true)
		} else {
			m.UseReconfigure = types.BoolNull()
		}
	}
}
