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
	_ resource.Resource                = &IPDHCPServerNetworkResource{}
	_ resource.ResourceWithImportState = &IPDHCPServerNetworkResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPDHCPServerNetworkResource struct {
	reg *client.Registry
}

type IPDHCPServerNetworkModel struct {
	ID            types.String `tfsdk:"id"`
	NtpNone       types.String `tfsdk:"ntp_none"`
	DnsNone       types.String `tfsdk:"dns_none"`
	Address       types.String `tfsdk:"address"`
	BootFileName  types.String `tfsdk:"boot_file_name"`
	CapsManager   types.String `tfsdk:"caps_manager"`
	CapsManagers  types.String `tfsdk:"caps_managers"`
	Comment       types.String `tfsdk:"comment"`
	DHCPOption    types.String `tfsdk:"dhcp_option"`
	DHCPOptionSet types.String `tfsdk:"dhcp_option_set"`
	DHCPOptions   types.String `tfsdk:"dhcp_options"`
	DNSServer     types.String `tfsdk:"dns_server"`
	DNSServers    types.String `tfsdk:"dns_servers"`
	Domain        types.String `tfsdk:"domain"`
	Dynamic       types.String `tfsdk:"dynamic"`
	Gateway       types.String `tfsdk:"gateway"`
	Netmask       types.String `tfsdk:"netmask"`
	NextServer    types.String `tfsdk:"next_server"`
	Nndns         types.String `tfsdk:"nndns"`
	Nnntp         types.String `tfsdk:"nnntp"`
	NoDNS         types.Bool   `tfsdk:"no_dns"`
	NoNTP         types.Bool   `tfsdk:"no_ntp"`
	NTPServer     types.String `tfsdk:"ntp_server"`
	NTPServers    types.String `tfsdk:"ntp_servers"`
	WinsServer    types.String `tfsdk:"wins_server"`
	WinsServers   types.String `tfsdk:"wins_servers"`
	Router        types.String `tfsdk:"router"`
}

func NewIPDHCPServerNetworkResource() resource.Resource { return &IPDHCPServerNetworkResource{} }

func (r *IPDHCPServerNetworkResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_dhcp_server_network"
}

func (r *IPDHCPServerNetworkResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPDHCPServerNetworkResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/dhcp-server/network`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ntp_none": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ntp-none`.",
			},
			"dns_none": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dns-none`.",
			},
			"address": schema.StringAttribute{
				Required:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsCIDR()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeCIDR()},
			},
			"boot_file_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"caps_manager": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"caps_managers": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"dhcp_option": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dhcp_option_set": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dhcp_options": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"dns_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dns_servers": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dynamic": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"netmask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"next_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nndns": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"nnntp": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"no_dns": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"no_ntp": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"ntp_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ntp_servers": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"wins_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wins_servers": schema.StringAttribute{
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

func (r *IPDHCPServerNetworkResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPDHCPServerNetworkModel
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
	if !(plan.BootFileName.IsNull() || plan.BootFileName.IsUnknown()) {
		body["boot-file-name"] = plan.BootFileName.ValueString()
	}
	if !(plan.CapsManager.IsNull() || plan.CapsManager.IsUnknown()) {
		body["caps-manager"] = plan.CapsManager.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.DHCPOption.IsNull() || plan.DHCPOption.IsUnknown()) {
		body["dhcp-option"] = plan.DHCPOption.ValueString()
	}
	if !(plan.DHCPOptionSet.IsNull() || plan.DHCPOptionSet.IsUnknown()) {
		body["dhcp-option-set"] = plan.DHCPOptionSet.ValueString()
	}
	if !(plan.DNSServer.IsNull() || plan.DNSServer.IsUnknown()) {
		body["dns-server"] = plan.DNSServer.ValueString()
	}
	if !(plan.Domain.IsNull() || plan.Domain.IsUnknown()) {
		body["domain"] = plan.Domain.ValueString()
	}
	if !(plan.Gateway.IsNull() || plan.Gateway.IsUnknown()) {
		body["gateway"] = plan.Gateway.ValueString()
	}
	if !(plan.Netmask.IsNull() || plan.Netmask.IsUnknown()) {
		body["netmask"] = plan.Netmask.ValueString()
	}
	if !(plan.NextServer.IsNull() || plan.NextServer.IsUnknown()) {
		body["next-server"] = plan.NextServer.ValueString()
	}
	if !(plan.NTPServer.IsNull() || plan.NTPServer.IsUnknown()) {
		body["ntp-server"] = plan.NTPServer.ValueString()
	}
	if !(plan.WinsServer.IsNull() || plan.WinsServer.IsUnknown()) {
		body["wins-server"] = plan.WinsServer.ValueString()
	}
	if !(plan.DnsNone.IsNull() || plan.DnsNone.IsUnknown()) {
		body["dns-none"] = plan.DnsNone.ValueString()
	}
	if !(plan.NtpNone.IsNull() || plan.NtpNone.IsUnknown()) {
		body["ntp-none"] = plan.NtpNone.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/dhcp-server/network", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/dhcp-server/network failed", err.Error())
		return
	}
	iPDHCPServerNetworkApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPServerNetworkResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPDHCPServerNetworkModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/dhcp-server/network", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/dhcp-server/network failed", err.Error())
		return
	}
	iPDHCPServerNetworkApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPDHCPServerNetworkResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPDHCPServerNetworkModel
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
	if !plan.Address.Equal(state.Address) && !plan.Address.IsUnknown() {
		body["address"] = plan.Address.ValueString()
	}
	if !plan.BootFileName.Equal(state.BootFileName) && !plan.BootFileName.IsUnknown() {
		body["boot-file-name"] = plan.BootFileName.ValueString()
	}
	if !plan.CapsManager.Equal(state.CapsManager) && !plan.CapsManager.IsUnknown() {
		body["caps-manager"] = plan.CapsManager.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.DHCPOption.Equal(state.DHCPOption) && !plan.DHCPOption.IsUnknown() {
		body["dhcp-option"] = plan.DHCPOption.ValueString()
	}
	if !plan.DHCPOptionSet.Equal(state.DHCPOptionSet) && !plan.DHCPOptionSet.IsUnknown() {
		body["dhcp-option-set"] = plan.DHCPOptionSet.ValueString()
	}
	if !plan.DNSServer.Equal(state.DNSServer) && !plan.DNSServer.IsUnknown() {
		body["dns-server"] = plan.DNSServer.ValueString()
	}
	if !plan.Domain.Equal(state.Domain) && !plan.Domain.IsUnknown() {
		body["domain"] = plan.Domain.ValueString()
	}
	if !plan.Gateway.Equal(state.Gateway) && !plan.Gateway.IsUnknown() {
		body["gateway"] = plan.Gateway.ValueString()
	}
	if !plan.Netmask.Equal(state.Netmask) && !plan.Netmask.IsUnknown() {
		body["netmask"] = plan.Netmask.ValueString()
	}
	if !plan.NextServer.Equal(state.NextServer) && !plan.NextServer.IsUnknown() {
		body["next-server"] = plan.NextServer.ValueString()
	}
	if !plan.NTPServer.Equal(state.NTPServer) && !plan.NTPServer.IsUnknown() {
		body["ntp-server"] = plan.NTPServer.ValueString()
	}
	if !plan.WinsServer.Equal(state.WinsServer) && !plan.WinsServer.IsUnknown() {
		body["wins-server"] = plan.WinsServer.ValueString()
	}
	if !plan.DnsNone.Equal(state.DnsNone) && !plan.DnsNone.IsUnknown() {
		body["dns-none"] = plan.DnsNone.ValueString()
	}
	if !plan.NtpNone.Equal(state.NtpNone) && !plan.NtpNone.IsUnknown() {
		body["ntp-none"] = plan.NtpNone.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/dhcp-server/network", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/dhcp-server/network failed", err.Error())
			return
		}
		iPDHCPServerNetworkApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPServerNetworkResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPDHCPServerNetworkModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/dhcp-server/network", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/dhcp-server/network failed", err.Error())
	}
}

func (r *IPDHCPServerNetworkResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPDHCPServerNetworkLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/dhcp-server/network matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPDHCPServerNetworkLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPDHCPServerNetworkLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/dhcp-server/network", id)
}

func iPDHCPServerNetworkApply(ctx context.Context, obj client.Object, m *IPDHCPServerNetworkModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["ntp-none"]; ok && v != "" {
		m.NtpNone = types.StringValue(v)
	} else {
		m.NtpNone = types.StringNull()
	}
	if v, ok := obj["dns-none"]; ok && v != "" {
		m.DnsNone = types.StringValue(v)
	} else {
		m.DnsNone = types.StringNull()
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
	if v, ok := obj["boot-file-name"]; ok {
		_ = v
		if v != "" {
			m.BootFileName = types.StringValue(v)
		} else {
			m.BootFileName = types.StringNull()
		}
	} else {
		m.BootFileName = types.StringNull()
	}
	if v, ok := obj["caps-manager"]; ok {
		_ = v
		if v != "" {
			m.CapsManager = types.StringValue(v)
		} else {
			m.CapsManager = types.StringNull()
		}
	} else {
		m.CapsManager = types.StringNull()
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
	if v, ok := obj["dhcp-option"]; ok {
		_ = v
		if v != "" {
			m.DHCPOption = types.StringValue(v)
		} else {
			m.DHCPOption = types.StringNull()
		}
	} else {
		m.DHCPOption = types.StringNull()
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
	if v, ok := obj["dns-server"]; ok {
		_ = v
		if v != "" {
			m.DNSServer = types.StringValue(v)
		} else {
			m.DNSServer = types.StringNull()
		}
	} else {
		m.DNSServer = types.StringNull()
	}
	if v, ok := obj["dns-servers"]; ok {
		_ = v
		if v != "" {
			m.DNSServers = types.StringValue(v)
		} else {
			m.DNSServers = types.StringNull()
		}
	} else {
		m.DNSServers = types.StringNull()
	}
	if v, ok := obj["domain"]; ok {
		_ = v
		if v != "" {
			m.Domain = types.StringValue(v)
		} else {
			m.Domain = types.StringNull()
		}
	} else {
		m.Domain = types.StringNull()
	}
	if v, ok := obj["dynamic"]; ok {
		_ = v
		if v != "" {
			m.Dynamic = types.StringValue(v)
		} else {
			m.Dynamic = types.StringNull()
		}
	} else {
		m.Dynamic = types.StringNull()
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
	if v, ok := obj["netmask"]; ok {
		_ = v
		if v != "" {
			m.Netmask = types.StringValue(v)
		} else {
			m.Netmask = types.StringNull()
		}
	} else {
		m.Netmask = types.StringNull()
	}
	if v, ok := obj["next-server"]; ok {
		_ = v
		if v != "" {
			m.NextServer = types.StringValue(v)
		} else {
			m.NextServer = types.StringNull()
		}
	} else {
		m.NextServer = types.StringNull()
	}
	if v, ok := obj["nndns"]; ok {
		_ = v
		if v != "" {
			m.Nndns = types.StringValue(v)
		} else {
			m.Nndns = types.StringNull()
		}
	} else {
		m.Nndns = types.StringNull()
	}
	if v, ok := obj["nnntp"]; ok {
		_ = v
		if v != "" {
			m.Nnntp = types.StringValue(v)
		} else {
			m.Nnntp = types.StringNull()
		}
	} else {
		m.Nnntp = types.StringNull()
	}
	if v, ok := obj["no-dns"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.NoDNS = types.BoolValue(b)
		} else {
			m.NoDNS = types.BoolNull()
		}
	} else {
		m.NoDNS = types.BoolNull()
	}
	if v, ok := obj["no-ntp"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.NoNTP = types.BoolValue(b)
		} else {
			m.NoNTP = types.BoolNull()
		}
	} else {
		m.NoNTP = types.BoolNull()
	}
	if v, ok := obj["ntp-server"]; ok {
		_ = v
		if v != "" {
			m.NTPServer = types.StringValue(v)
		} else {
			m.NTPServer = types.StringNull()
		}
	} else {
		m.NTPServer = types.StringNull()
	}
	if v, ok := obj["ntp-servers"]; ok {
		_ = v
		if v != "" {
			m.NTPServers = types.StringValue(v)
		} else {
			m.NTPServers = types.StringNull()
		}
	} else {
		m.NTPServers = types.StringNull()
	}
	if v, ok := obj["wins-server"]; ok {
		_ = v
		if v != "" {
			m.WinsServer = types.StringValue(v)
		} else {
			m.WinsServer = types.StringNull()
		}
	} else {
		m.WinsServer = types.StringNull()
	}
	if v, ok := obj["wins-servers"]; ok {
		_ = v
		if v != "" {
			m.WinsServers = types.StringValue(v)
		} else {
			m.WinsServers = types.StringNull()
		}
	} else {
		m.WinsServers = types.StringNull()
	}
}
