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
	_ resource.Resource                = &IPDHCPRelayResource{}
	_ resource.ResourceWithImportState = &IPDHCPRelayResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPDHCPRelayResource struct {
	reg *client.Registry
}

type IPDHCPRelayModel struct {
	ID                     types.String `tfsdk:"id"`
	AddRelayInfo           types.Bool   `tfsdk:"add_relay_info"`
	DelayThreshold         types.String `tfsdk:"delay_threshold"`
	DHCPServer             types.String `tfsdk:"dhcp_server"`
	DHCPServerVrf          types.String `tfsdk:"dhcp_server_vrf"`
	Disabled               types.Bool   `tfsdk:"disabled"`
	Interface              types.String `tfsdk:"interface"`
	Invalid                types.Bool   `tfsdk:"invalid"`
	LocalAddress           types.String `tfsdk:"local_address"`
	LocalAddressAsSourceIP types.Bool   `tfsdk:"local_address_as_source_ip"`
	Name                   types.String `tfsdk:"name"`
	RelayInfoRemoteID      types.String `tfsdk:"relay_info_remote_id"`
	ResetCounters          types.String `tfsdk:"reset_counters"`
	Router                 types.String `tfsdk:"router"`
}

func NewIPDHCPRelayResource() resource.Resource { return &IPDHCPRelayResource{} }

func (r *IPDHCPRelayResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_dhcp_relay"
}

func (r *IPDHCPRelayResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *IPDHCPRelayResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/dhcp-relay`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"add_relay_info": schema.BoolAttribute{
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
			"dhcp_server": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"dhcp_server_vrf": schema.StringAttribute{
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
			"invalid": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"local_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"local_address_as_source_ip": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"relay_info_remote_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"reset_counters": schema.StringAttribute{
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

func (r *IPDHCPRelayResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPDHCPRelayModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AddRelayInfo.IsNull() || plan.AddRelayInfo.IsUnknown()) {
		body["add-relay-info"] = client.FormatBool(plan.AddRelayInfo.ValueBool())
	}
	if !(plan.DelayThreshold.IsNull() || plan.DelayThreshold.IsUnknown()) {
		body["delay-threshold"] = plan.DelayThreshold.ValueString()
	}
	if !(plan.DHCPServer.IsNull() || plan.DHCPServer.IsUnknown()) {
		body["dhcp-server"] = plan.DHCPServer.ValueString()
	}
	if !(plan.DHCPServerVrf.IsNull() || plan.DHCPServerVrf.IsUnknown()) {
		body["dhcp-server-vrf"] = plan.DHCPServerVrf.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.LocalAddress.IsNull() || plan.LocalAddress.IsUnknown()) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !(plan.LocalAddressAsSourceIP.IsNull() || plan.LocalAddressAsSourceIP.IsUnknown()) {
		body["local-address-as-source-ip"] = client.FormatBool(plan.LocalAddressAsSourceIP.ValueBool())
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.RelayInfoRemoteID.IsNull() || plan.RelayInfoRemoteID.IsUnknown()) {
		body["relay-info-remote-id"] = plan.RelayInfoRemoteID.ValueString()
	}
	if !(plan.ResetCounters.IsNull() || plan.ResetCounters.IsUnknown()) {
		body["reset-counters"] = plan.ResetCounters.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/dhcp-relay", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/dhcp-relay failed", err.Error())
		return
	}
	iPDHCPRelayApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPRelayResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPDHCPRelayModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/dhcp-relay", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/dhcp-relay failed", err.Error())
		return
	}
	iPDHCPRelayApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPDHCPRelayResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPDHCPRelayModel
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
	if !plan.AddRelayInfo.Equal(state.AddRelayInfo) {
		body["add-relay-info"] = client.FormatBool(plan.AddRelayInfo.ValueBool())
	}
	if !plan.DelayThreshold.Equal(state.DelayThreshold) {
		body["delay-threshold"] = plan.DelayThreshold.ValueString()
	}
	if !plan.DHCPServer.Equal(state.DHCPServer) {
		body["dhcp-server"] = plan.DHCPServer.ValueString()
	}
	if !plan.DHCPServerVrf.Equal(state.DHCPServerVrf) {
		body["dhcp-server-vrf"] = plan.DHCPServerVrf.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.LocalAddress.Equal(state.LocalAddress) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !plan.LocalAddressAsSourceIP.Equal(state.LocalAddressAsSourceIP) {
		body["local-address-as-source-ip"] = client.FormatBool(plan.LocalAddressAsSourceIP.ValueBool())
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.RelayInfoRemoteID.Equal(state.RelayInfoRemoteID) {
		body["relay-info-remote-id"] = plan.RelayInfoRemoteID.ValueString()
	}
	if !plan.ResetCounters.Equal(state.ResetCounters) {
		body["reset-counters"] = plan.ResetCounters.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/dhcp-relay", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/dhcp-relay failed", err.Error())
			return
		}
		iPDHCPRelayApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPRelayResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPDHCPRelayModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/dhcp-relay", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/dhcp-relay failed", err.Error())
	}
}

func (r *IPDHCPRelayResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPDHCPRelayLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/dhcp-relay matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPDHCPRelayLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPDHCPRelayLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/dhcp-relay", id)
}

func iPDHCPRelayApply(ctx context.Context, obj client.Object, m *IPDHCPRelayModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["add-relay-info"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AddRelayInfo = types.BoolValue(b)
		} else {
			m.AddRelayInfo = types.BoolNull()
		}
	} else {
		m.AddRelayInfo = types.BoolNull()
	}
	if v, ok := obj["delay-threshold"]; ok {
		_ = v
		if v != "" {
			m.DelayThreshold = types.StringValue(v)
		} else {
			m.DelayThreshold = types.StringNull()
		}
	} else {
		m.DelayThreshold = types.StringNull()
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
	if v, ok := obj["dhcp-server-vrf"]; ok {
		_ = v
		if v != "" {
			m.DHCPServerVrf = types.StringValue(v)
		} else {
			m.DHCPServerVrf = types.StringNull()
		}
	} else {
		m.DHCPServerVrf = types.StringNull()
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
	if v, ok := obj["local-address-as-source-ip"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.LocalAddressAsSourceIP = types.BoolValue(b)
		} else {
			m.LocalAddressAsSourceIP = types.BoolNull()
		}
	} else {
		m.LocalAddressAsSourceIP = types.BoolNull()
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
	if v, ok := obj["relay-info-remote-id"]; ok {
		_ = v
		if v != "" {
			m.RelayInfoRemoteID = types.StringValue(v)
		} else {
			m.RelayInfoRemoteID = types.StringNull()
		}
	} else {
		m.RelayInfoRemoteID = types.StringNull()
	}
	if v, ok := obj["reset-counters"]; ok {
		_ = v
		if v != "" {
			m.ResetCounters = types.StringValue(v)
		} else {
			m.ResetCounters = types.StringNull()
		}
	} else {
		m.ResetCounters = types.StringNull()
	}
}
