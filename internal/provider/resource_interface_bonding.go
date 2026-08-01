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
	_ resource.Resource                = &InterfaceBondingResource{}
	_ resource.ResourceWithImportState = &InterfaceBondingResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceBondingResource struct {
	reg *client.Registry
}

type InterfaceBondingModel struct {
	ID                 types.String `tfsdk:"id"`
	UpDelay            types.String `tfsdk:"up_delay"`
	TransmitHashPolicy types.String `tfsdk:"transmit_hash_policy"`
	Slaves             types.String `tfsdk:"slaves"`
	Primary            types.String `tfsdk:"primary"`
	MlagId             types.String `tfsdk:"mlag_id"`
	MinLinks           types.String `tfsdk:"min_links"`
	MiiInterval        types.String `tfsdk:"mii_interval"`
	LinkMonitoring     types.String `tfsdk:"link_monitoring"`
	LacpUserKey        types.String `tfsdk:"lacp_user_key"`
	LacpSystemPriority types.String `tfsdk:"lacp_system_priority"`
	LacpSystemId       types.String `tfsdk:"lacp_system_id"`
	LacpRate           types.String `tfsdk:"lacp_rate"`
	LacpMode           types.String `tfsdk:"lacp_mode"`
	ForcedMacAddress   types.String `tfsdk:"forced_mac_address"`
	DownDelay          types.String `tfsdk:"down_delay"`
	ArpIpTargets       types.String `tfsdk:"arp_ip_targets"`
	ArpInterval        types.String `tfsdk:"arp_interval"`
	ARP                types.String `tfsdk:"arp"`
	ARPTimeout         types.String `tfsdk:"arp_timeout"`
	Comment            types.String `tfsdk:"comment"`
	Disabled           types.Bool   `tfsdk:"disabled"`
	Mode               types.String `tfsdk:"mode"`
	MTU                types.String `tfsdk:"mtu"`
	Name               types.String `tfsdk:"name"`
	Router             types.String `tfsdk:"router"`
}

func NewInterfaceBondingResource() resource.Resource { return &InterfaceBondingResource{} }

func (r *InterfaceBondingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_bonding"
}

func (r *InterfaceBondingResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceBondingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Auto-test requires a typed-reference precondition (e.g. an existing peer,\ninstance, bridge of the specific kind). The current acc-test generator's\ngeneric data.routeros_interface.all lookup can't satisfy these. Use this\nresource manually with explicit references to a precondition resource\nin your config.\n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"up_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `up-delay`.",
			},
			"transmit_hash_policy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `transmit-hash-policy`.",
			},
			"slaves": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `slaves`.",
			},
			"primary": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `primary`.",
			},
			"mlag_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mlag-id`.",
			},
			"min_links": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `min-links`.",
			},
			"mii_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mii-interval`.",
			},
			"link_monitoring": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `link-monitoring`.",
			},
			"lacp_user_key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `lacp-user-key`.",
			},
			"lacp_system_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `lacp-system-priority`.",
			},
			"lacp_system_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `lacp-system-id`.",
			},
			"lacp_rate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `lacp-rate`.",
			},
			"lacp_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `lacp-mode`.",
			},
			"forced_mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `forced-mac-address`.",
			},
			"down_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `down-delay`.",
			},
			"arp_ip_targets": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `arp-ip-targets`.",
			},
			"arp_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `arp-interval`.",
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
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
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

func (r *InterfaceBondingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceBondingModel
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
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Mode.IsNull() || plan.Mode.IsUnknown()) {
		body["mode"] = plan.Mode.ValueString()
	}
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.ArpInterval.IsNull() || plan.ArpInterval.IsUnknown()) {
		body["arp-interval"] = plan.ArpInterval.ValueString()
	}
	if !(plan.ArpIpTargets.IsNull() || plan.ArpIpTargets.IsUnknown()) {
		body["arp-ip-targets"] = plan.ArpIpTargets.ValueString()
	}
	if !(plan.DownDelay.IsNull() || plan.DownDelay.IsUnknown()) {
		body["down-delay"] = plan.DownDelay.ValueString()
	}
	if !(plan.ForcedMacAddress.IsNull() || plan.ForcedMacAddress.IsUnknown()) {
		body["forced-mac-address"] = plan.ForcedMacAddress.ValueString()
	}
	if !(plan.LacpMode.IsNull() || plan.LacpMode.IsUnknown()) {
		body["lacp-mode"] = plan.LacpMode.ValueString()
	}
	if !(plan.LacpRate.IsNull() || plan.LacpRate.IsUnknown()) {
		body["lacp-rate"] = plan.LacpRate.ValueString()
	}
	if !(plan.LacpSystemId.IsNull() || plan.LacpSystemId.IsUnknown()) {
		body["lacp-system-id"] = plan.LacpSystemId.ValueString()
	}
	if !(plan.LacpSystemPriority.IsNull() || plan.LacpSystemPriority.IsUnknown()) {
		body["lacp-system-priority"] = plan.LacpSystemPriority.ValueString()
	}
	if !(plan.LacpUserKey.IsNull() || plan.LacpUserKey.IsUnknown()) {
		body["lacp-user-key"] = plan.LacpUserKey.ValueString()
	}
	if !(plan.LinkMonitoring.IsNull() || plan.LinkMonitoring.IsUnknown()) {
		body["link-monitoring"] = plan.LinkMonitoring.ValueString()
	}
	if !(plan.MiiInterval.IsNull() || plan.MiiInterval.IsUnknown()) {
		body["mii-interval"] = plan.MiiInterval.ValueString()
	}
	if !(plan.MinLinks.IsNull() || plan.MinLinks.IsUnknown()) {
		body["min-links"] = plan.MinLinks.ValueString()
	}
	if !(plan.MlagId.IsNull() || plan.MlagId.IsUnknown()) {
		body["mlag-id"] = plan.MlagId.ValueString()
	}
	if !(plan.Primary.IsNull() || plan.Primary.IsUnknown()) {
		body["primary"] = plan.Primary.ValueString()
	}
	if !(plan.Slaves.IsNull() || plan.Slaves.IsUnknown()) {
		body["slaves"] = plan.Slaves.ValueString()
	}
	if !(plan.TransmitHashPolicy.IsNull() || plan.TransmitHashPolicy.IsUnknown()) {
		body["transmit-hash-policy"] = plan.TransmitHashPolicy.ValueString()
	}
	if !(plan.UpDelay.IsNull() || plan.UpDelay.IsUnknown()) {
		body["up-delay"] = plan.UpDelay.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/bonding", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/bonding failed", err.Error())
		return
	}
	interfaceBondingApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceBondingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceBondingModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/bonding", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/bonding failed", err.Error())
		return
	}
	interfaceBondingApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceBondingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceBondingModel
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
	if !plan.ARP.Equal(state.ARP) && !plan.ARP.IsUnknown() {
		body["arp"] = plan.ARP.ValueString()
	}
	if !plan.ARPTimeout.Equal(state.ARPTimeout) && !plan.ARPTimeout.IsUnknown() {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Mode.Equal(state.Mode) && !plan.Mode.IsUnknown() {
		body["mode"] = plan.Mode.ValueString()
	}
	if !plan.MTU.Equal(state.MTU) && !plan.MTU.IsUnknown() {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.ArpInterval.Equal(state.ArpInterval) && !plan.ArpInterval.IsUnknown() {
		body["arp-interval"] = plan.ArpInterval.ValueString()
	}
	if !plan.ArpIpTargets.Equal(state.ArpIpTargets) && !plan.ArpIpTargets.IsUnknown() {
		body["arp-ip-targets"] = plan.ArpIpTargets.ValueString()
	}
	if !plan.DownDelay.Equal(state.DownDelay) && !plan.DownDelay.IsUnknown() {
		body["down-delay"] = plan.DownDelay.ValueString()
	}
	if !plan.ForcedMacAddress.Equal(state.ForcedMacAddress) && !plan.ForcedMacAddress.IsUnknown() {
		body["forced-mac-address"] = plan.ForcedMacAddress.ValueString()
	}
	if !plan.LacpMode.Equal(state.LacpMode) && !plan.LacpMode.IsUnknown() {
		body["lacp-mode"] = plan.LacpMode.ValueString()
	}
	if !plan.LacpRate.Equal(state.LacpRate) && !plan.LacpRate.IsUnknown() {
		body["lacp-rate"] = plan.LacpRate.ValueString()
	}
	if !plan.LacpSystemId.Equal(state.LacpSystemId) && !plan.LacpSystemId.IsUnknown() {
		body["lacp-system-id"] = plan.LacpSystemId.ValueString()
	}
	if !plan.LacpSystemPriority.Equal(state.LacpSystemPriority) && !plan.LacpSystemPriority.IsUnknown() {
		body["lacp-system-priority"] = plan.LacpSystemPriority.ValueString()
	}
	if !plan.LacpUserKey.Equal(state.LacpUserKey) && !plan.LacpUserKey.IsUnknown() {
		body["lacp-user-key"] = plan.LacpUserKey.ValueString()
	}
	if !plan.LinkMonitoring.Equal(state.LinkMonitoring) && !plan.LinkMonitoring.IsUnknown() {
		body["link-monitoring"] = plan.LinkMonitoring.ValueString()
	}
	if !plan.MiiInterval.Equal(state.MiiInterval) && !plan.MiiInterval.IsUnknown() {
		body["mii-interval"] = plan.MiiInterval.ValueString()
	}
	if !plan.MinLinks.Equal(state.MinLinks) && !plan.MinLinks.IsUnknown() {
		body["min-links"] = plan.MinLinks.ValueString()
	}
	if !plan.MlagId.Equal(state.MlagId) && !plan.MlagId.IsUnknown() {
		body["mlag-id"] = plan.MlagId.ValueString()
	}
	if !plan.Primary.Equal(state.Primary) && !plan.Primary.IsUnknown() {
		body["primary"] = plan.Primary.ValueString()
	}
	if !plan.Slaves.Equal(state.Slaves) && !plan.Slaves.IsUnknown() {
		body["slaves"] = plan.Slaves.ValueString()
	}
	if !plan.TransmitHashPolicy.Equal(state.TransmitHashPolicy) && !plan.TransmitHashPolicy.IsUnknown() {
		body["transmit-hash-policy"] = plan.TransmitHashPolicy.ValueString()
	}
	if !plan.UpDelay.Equal(state.UpDelay) && !plan.UpDelay.IsUnknown() {
		body["up-delay"] = plan.UpDelay.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/bonding", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/bonding failed", err.Error())
			return
		}
		interfaceBondingApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceBondingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceBondingModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/bonding", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/bonding failed", err.Error())
	}
}

func (r *InterfaceBondingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceBondingLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/bonding matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceBondingLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceBondingLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/bonding", id)
}

func interfaceBondingApply(ctx context.Context, obj client.Object, m *InterfaceBondingModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["up-delay"]; ok && v != "" {
		m.UpDelay = types.StringValue(v)
	} else {
		m.UpDelay = types.StringNull()
	}
	if v, ok := obj["transmit-hash-policy"]; ok && v != "" {
		m.TransmitHashPolicy = types.StringValue(v)
	} else {
		m.TransmitHashPolicy = types.StringNull()
	}
	if v, ok := obj["slaves"]; ok && v != "" {
		m.Slaves = types.StringValue(v)
	} else {
		m.Slaves = types.StringNull()
	}
	if v, ok := obj["primary"]; ok && v != "" {
		m.Primary = types.StringValue(v)
	} else {
		m.Primary = types.StringNull()
	}
	if v, ok := obj["mlag-id"]; ok && v != "" {
		m.MlagId = types.StringValue(v)
	} else {
		m.MlagId = types.StringNull()
	}
	if v, ok := obj["min-links"]; ok && v != "" {
		m.MinLinks = types.StringValue(v)
	} else {
		m.MinLinks = types.StringNull()
	}
	if v, ok := obj["mii-interval"]; ok && v != "" {
		m.MiiInterval = types.StringValue(v)
	} else {
		m.MiiInterval = types.StringNull()
	}
	if v, ok := obj["link-monitoring"]; ok && v != "" {
		m.LinkMonitoring = types.StringValue(v)
	} else {
		m.LinkMonitoring = types.StringNull()
	}
	if v, ok := obj["lacp-user-key"]; ok && v != "" {
		m.LacpUserKey = types.StringValue(v)
	} else {
		m.LacpUserKey = types.StringNull()
	}
	if v, ok := obj["lacp-system-priority"]; ok && v != "" {
		m.LacpSystemPriority = types.StringValue(v)
	} else {
		m.LacpSystemPriority = types.StringNull()
	}
	if v, ok := obj["lacp-system-id"]; ok && v != "" {
		m.LacpSystemId = types.StringValue(v)
	} else {
		m.LacpSystemId = types.StringNull()
	}
	if v, ok := obj["lacp-rate"]; ok && v != "" {
		m.LacpRate = types.StringValue(v)
	} else {
		m.LacpRate = types.StringNull()
	}
	if v, ok := obj["lacp-mode"]; ok && v != "" {
		m.LacpMode = types.StringValue(v)
	} else {
		m.LacpMode = types.StringNull()
	}
	if v, ok := obj["forced-mac-address"]; ok && v != "" {
		m.ForcedMacAddress = types.StringValue(v)
	} else {
		m.ForcedMacAddress = types.StringNull()
	}
	if v, ok := obj["down-delay"]; ok && v != "" {
		m.DownDelay = types.StringValue(v)
	} else {
		m.DownDelay = types.StringNull()
	}
	if v, ok := obj["arp-ip-targets"]; ok && v != "" {
		m.ArpIpTargets = types.StringValue(v)
	} else {
		m.ArpIpTargets = types.StringNull()
	}
	if v, ok := obj["arp-interval"]; ok && v != "" {
		m.ArpInterval = types.StringValue(v)
	} else {
		m.ArpInterval = types.StringNull()
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
	if v, ok := obj["mode"]; ok {
		_ = v
		if v != "" {
			m.Mode = types.StringValue(v)
		} else {
			m.Mode = types.StringNull()
		}
	} else {
		m.Mode = types.StringNull()
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
}
