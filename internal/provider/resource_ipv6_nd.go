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
	_ resource.Resource                = &IPV6NdResource{}
	_ resource.ResourceWithImportState = &IPV6NdResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPV6NdResource struct {
	reg *client.Registry
}

type IPV6NdModel struct {
	ID                          types.String  `tfsdk:"id"`
	Pref64                      types.String  `tfsdk:"pref64"`
	Dns                         types.String  `tfsdk:"dns"`
	AdvertiseDNS                types.String  `tfsdk:"advertise_dns"`
	AdvertiseMACAddress         types.Bool    `tfsdk:"advertise_mac_address"`
	Comment                     types.String  `tfsdk:"comment"`
	Default                     types.Bool    `tfsdk:"default"`
	Disabled                    types.Bool    `tfsdk:"disabled"`
	DNSServers                  types.String  `tfsdk:"dns_servers"`
	HopLimit                    types.String  `tfsdk:"hop_limit"`
	Interface                   types.String  `tfsdk:"interface"`
	Invalid                     types.Bool    `tfsdk:"invalid"`
	ManagedAddressConfiguration types.Bool    `tfsdk:"managed_address_configuration"`
	MTU                         types.String  `tfsdk:"mtu"`
	OtherConfiguration          types.Bool    `tfsdk:"other_configuration"`
	Pref64Prefixes              types.String  `tfsdk:"pref64_prefixes"`
	RaDelay                     durationValue `tfsdk:"ra_delay"`
	RaInterval                  types.String  `tfsdk:"ra_interval"`
	RaLifetime                  durationValue `tfsdk:"ra_lifetime"`
	RaPreference                types.String  `tfsdk:"ra_preference"`
	ReachableTime               types.String  `tfsdk:"reachable_time"`
	RetransmitInterval          types.String  `tfsdk:"retransmit_interval"`
	Router                      types.String  `tfsdk:"router"`
}

func NewIPV6NdResource() resource.Resource { return &IPV6NdResource{} }

func (r *IPV6NdResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv6_nd"
}

func (r *IPV6NdResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPV6NdResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "ND config is per-interface and conflicts with defaults if the interface is already configured.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"pref64": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `pref64`.",
			},
			"dns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dns`.",
			},
			"advertise_dns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "yes", "self"}...)},
			},
			"advertise_mac_address": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"default": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dns_servers": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"hop_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Hop limit advertised in router advertisements. A number, or `unspecified` (the default).",
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
			"managed_address_configuration": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MTU advertised in router advertisements. A number, or `unspecified` (the default).",
			},
			"other_configuration": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pref64_prefixes": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"ra_delay": schema.StringAttribute{
				CustomType:  durationType{},
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"ra_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ra_lifetime": schema.StringAttribute{
				CustomType:  durationType{},
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"ra_preference": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"medium", "high", "low"}...)},
			},
			"reachable_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Reachable time advertised in router advertisements. A number, or `unspecified` (the default).",
			},
			"retransmit_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Retransmit interval advertised in router advertisements. A number, or `unspecified` (the default).",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *IPV6NdResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPV6NdModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AdvertiseDNS.IsNull() || plan.AdvertiseDNS.IsUnknown()) {
		body["advertise-dns"] = plan.AdvertiseDNS.ValueString()
	}
	if !(plan.AdvertiseMACAddress.IsNull() || plan.AdvertiseMACAddress.IsUnknown()) {
		body["advertise-mac-address"] = client.FormatBool(plan.AdvertiseMACAddress.ValueBool())
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.HopLimit.IsNull() || plan.HopLimit.IsUnknown()) {
		body["hop-limit"] = plan.HopLimit.ValueString()
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.ManagedAddressConfiguration.IsNull() || plan.ManagedAddressConfiguration.IsUnknown()) {
		body["managed-address-configuration"] = client.FormatBool(plan.ManagedAddressConfiguration.ValueBool())
	}
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !(plan.OtherConfiguration.IsNull() || plan.OtherConfiguration.IsUnknown()) {
		body["other-configuration"] = client.FormatBool(plan.OtherConfiguration.ValueBool())
	}
	if !(plan.RaDelay.IsNull() || plan.RaDelay.IsUnknown()) {
		body["ra-delay"] = plan.RaDelay.ValueString()
	}
	if !(plan.RaInterval.IsNull() || plan.RaInterval.IsUnknown()) {
		body["ra-interval"] = plan.RaInterval.ValueString()
	}
	if !(plan.RaLifetime.IsNull() || plan.RaLifetime.IsUnknown()) {
		body["ra-lifetime"] = plan.RaLifetime.ValueString()
	}
	if !(plan.RaPreference.IsNull() || plan.RaPreference.IsUnknown()) {
		body["ra-preference"] = plan.RaPreference.ValueString()
	}
	if !(plan.ReachableTime.IsNull() || plan.ReachableTime.IsUnknown()) {
		body["reachable-time"] = plan.ReachableTime.ValueString()
	}
	if !(plan.RetransmitInterval.IsNull() || plan.RetransmitInterval.IsUnknown()) {
		body["retransmit-interval"] = plan.RetransmitInterval.ValueString()
	}
	if !(plan.Dns.IsNull() || plan.Dns.IsUnknown()) {
		body["dns"] = plan.Dns.ValueString()
	}
	if !(plan.Pref64.IsNull() || plan.Pref64.IsUnknown()) {
		body["pref64"] = plan.Pref64.ValueString()
	}
	obj, err := c.Add(ctx, "/ipv6/nd", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ipv6/nd failed", err.Error())
		return
	}
	iPV6NdApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6NdResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPV6NdModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ipv6/nd", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ipv6/nd failed", err.Error())
		return
	}
	iPV6NdApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPV6NdResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPV6NdModel
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
	if !plan.AdvertiseDNS.Equal(state.AdvertiseDNS) && !plan.AdvertiseDNS.IsUnknown() {
		body["advertise-dns"] = plan.AdvertiseDNS.ValueString()
	}
	if !plan.AdvertiseMACAddress.Equal(state.AdvertiseMACAddress) && !plan.AdvertiseMACAddress.IsUnknown() {
		body["advertise-mac-address"] = client.FormatBool(plan.AdvertiseMACAddress.ValueBool())
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.HopLimit.Equal(state.HopLimit) && !plan.HopLimit.IsUnknown() {
		body["hop-limit"] = plan.HopLimit.ValueString()
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.ManagedAddressConfiguration.Equal(state.ManagedAddressConfiguration) && !plan.ManagedAddressConfiguration.IsUnknown() {
		body["managed-address-configuration"] = client.FormatBool(plan.ManagedAddressConfiguration.ValueBool())
	}
	if !plan.MTU.Equal(state.MTU) && !plan.MTU.IsUnknown() {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !plan.OtherConfiguration.Equal(state.OtherConfiguration) && !plan.OtherConfiguration.IsUnknown() {
		body["other-configuration"] = client.FormatBool(plan.OtherConfiguration.ValueBool())
	}
	if !plan.RaDelay.Equal(state.RaDelay) && !plan.RaDelay.IsUnknown() {
		body["ra-delay"] = plan.RaDelay.ValueString()
	}
	if !plan.RaInterval.Equal(state.RaInterval) && !plan.RaInterval.IsUnknown() {
		body["ra-interval"] = plan.RaInterval.ValueString()
	}
	if !plan.RaLifetime.Equal(state.RaLifetime) && !plan.RaLifetime.IsUnknown() {
		body["ra-lifetime"] = plan.RaLifetime.ValueString()
	}
	if !plan.RaPreference.Equal(state.RaPreference) && !plan.RaPreference.IsUnknown() {
		body["ra-preference"] = plan.RaPreference.ValueString()
	}
	if !plan.ReachableTime.Equal(state.ReachableTime) && !plan.ReachableTime.IsUnknown() {
		body["reachable-time"] = plan.ReachableTime.ValueString()
	}
	if !plan.RetransmitInterval.Equal(state.RetransmitInterval) && !plan.RetransmitInterval.IsUnknown() {
		body["retransmit-interval"] = plan.RetransmitInterval.ValueString()
	}
	if !plan.Dns.Equal(state.Dns) && !plan.Dns.IsUnknown() {
		body["dns"] = plan.Dns.ValueString()
	}
	if !plan.Pref64.Equal(state.Pref64) && !plan.Pref64.IsUnknown() {
		body["pref64"] = plan.Pref64.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ipv6/nd", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ipv6/nd failed", err.Error())
			return
		}
		iPV6NdApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6NdResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPV6NdModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ipv6/nd", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ipv6/nd failed", err.Error())
	}
}

func (r *IPV6NdResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPV6NdLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ipv6/nd matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPV6NdLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPV6NdLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ipv6/nd", id)
}

func iPV6NdApply(ctx context.Context, obj client.Object, m *IPV6NdModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["pref64"]; ok && v != "" {
		m.Pref64 = types.StringValue(v)
	} else {
		m.Pref64 = types.StringNull()
	}
	if v, ok := obj["dns"]; ok && v != "" {
		m.Dns = types.StringValue(v)
	} else {
		m.Dns = types.StringNull()
	}
	if v, ok := obj["advertise-dns"]; ok {
		if v != "" {
			m.AdvertiseDNS = types.StringValue(v)
		} else {
			m.AdvertiseDNS = types.StringNull()
		}
	}
	if v, ok := obj["advertise-mac-address"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.AdvertiseMACAddress = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.AdvertiseMACAddress = types.BoolValue(true)
		} else {
			m.AdvertiseMACAddress = types.BoolNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
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
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Disabled = types.BoolValue(true)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["dns-servers"]; ok {
		if v != "" {
			m.DNSServers = types.StringValue(v)
		} else {
			m.DNSServers = types.StringNull()
		}
	}
	if v, ok := obj["hop-limit"]; ok {
		if v != "" {
			m.HopLimit = types.StringValue(v)
		} else {
			m.HopLimit = types.StringNull()
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
	if v, ok := obj["managed-address-configuration"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.ManagedAddressConfiguration = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.ManagedAddressConfiguration = types.BoolValue(true)
		} else {
			m.ManagedAddressConfiguration = types.BoolNull()
		}
	}
	if v, ok := obj["mtu"]; ok {
		if v != "" {
			m.MTU = types.StringValue(v)
		} else {
			m.MTU = types.StringNull()
		}
	}
	if v, ok := obj["other-configuration"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.OtherConfiguration = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.OtherConfiguration = types.BoolValue(true)
		} else {
			m.OtherConfiguration = types.BoolNull()
		}
	}
	if v, ok := obj["pref64-prefixes"]; ok {
		if v != "" {
			m.Pref64Prefixes = types.StringValue(v)
		} else {
			m.Pref64Prefixes = types.StringNull()
		}
	}
	if v, ok := obj["ra-delay"]; ok {
		if v != "" {
			m.RaDelay = newDurationValue(v)
		} else {
			m.RaDelay = newDurationNull()
		}
	}
	if v, ok := obj["ra-interval"]; ok {
		if v != "" {
			m.RaInterval = types.StringValue(v)
		} else {
			m.RaInterval = types.StringNull()
		}
	}
	if v, ok := obj["ra-lifetime"]; ok {
		if v != "" {
			m.RaLifetime = newDurationValue(v)
		} else {
			m.RaLifetime = newDurationNull()
		}
	}
	if v, ok := obj["ra-preference"]; ok {
		if v != "" {
			m.RaPreference = types.StringValue(v)
		} else {
			m.RaPreference = types.StringNull()
		}
	}
	if v, ok := obj["reachable-time"]; ok {
		if v != "" {
			m.ReachableTime = types.StringValue(v)
		} else {
			m.ReachableTime = types.StringNull()
		}
	}
	if v, ok := obj["retransmit-interval"]; ok {
		if v != "" {
			m.RetransmitInterval = types.StringValue(v)
		} else {
			m.RetransmitInterval = types.StringNull()
		}
	}
}
