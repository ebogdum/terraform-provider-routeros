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
	_ resource.Resource                   = &IPDNSStaticResource{}
	_ resource.ResourceWithImportState    = &IPDNSStaticResource{}
	_ resource.ResourceWithValidateConfig = &IPDNSStaticResource{}
	_                                     = attr.Value(nil)
	_                                     = strings.TrimSpace
	_                                     = path.Root
)

type IPDNSStaticResource struct {
	reg *client.Registry
}

type IPDNSStaticModel struct {
	ID             types.String `tfsdk:"id"`
	Address        types.String `tfsdk:"address"`
	AddressList    types.String `tfsdk:"address_list"`
	Cname          types.String `tfsdk:"cname"`
	Comment        types.String `tfsdk:"comment"`
	Disabled       types.Bool   `tfsdk:"disabled"`
	ForwardTo      types.String `tfsdk:"forward_to"`
	MatchSubdomain types.String `tfsdk:"match_subdomain"`
	MxExchange     types.String `tfsdk:"mx_exchange"`
	MxPreference   types.String `tfsdk:"mx_preference"`
	Name           types.String `tfsdk:"name"`
	Ns             types.String `tfsdk:"ns"`
	Regexp         types.String `tfsdk:"regexp"`
	SrvPort        types.String `tfsdk:"srv_port"`
	SrvPriority    types.String `tfsdk:"srv_priority"`
	SrvTarget      types.String `tfsdk:"srv_target"`
	SrvWeight      types.String `tfsdk:"srv_weight"`
	Text           types.String `tfsdk:"text"`
	Ttl            types.String `tfsdk:"ttl"`
	Type           types.String `tfsdk:"type"`
	Router         types.String `tfsdk:"router"`
}

func NewIPDNSStaticResource() resource.Resource { return &IPDNSStaticResource{} }

func (r *IPDNSStaticResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_dns_static"
}

func (r *IPDNSStaticResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPDNSStaticResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A DNS A/AAAA/CNAME/MX/... static entry. Requires either `name` or `regexp`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"address": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Description: "IPv4/IPv6 address to return. Required when `type` is `\"A\"` or `\"AAAA\"` (the " +
					"default); must be left unset for other types (`CNAME`, `FWD`, `MX`, `NS`, `SRV`, `TXT`, or " +
					"a `regexp` entry).",
			},
			"address_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"cname": schema.StringAttribute{
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
			"forward_to": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"match_subdomain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mx_exchange": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mx_preference": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "FQDN matched against incoming queries. Provide this or `regexp`, not both.",
			},
			"ns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"regexp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"srv_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"srv_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"srv_target": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"srv_weight": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"text": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Record type. Defaults to `A`.",
				Validators: []validator.String{schemautil.OneOf(
					"A", "AAAA", "CNAME", "FWD", "MX", "NS", "NXDOMAIN", "SRV", "TXT")},
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

// ipDNSStaticTypeNeedsAddress reports whether typ stores an address ("" meaning
// the device default, "A"). Every other type has its own dedicated field.
func ipDNSStaticTypeNeedsAddress(typ string) bool {
	switch typ {
	case "", "A", "AAAA":
		return true
	default:
		return false
	}
}

// ipDNSStaticAddressTypeConflict checks address is required when type needs it and forbidden when it doesn't.
func ipDNSStaticAddressTypeConflict(typ string, hasAddress bool) (summary, detail string) {
	needsAddress := ipDNSStaticTypeNeedsAddress(typ)
	switch {
	case needsAddress && !hasAddress:
		return "Missing address", `routeros_ip_dns_static requires "address" when type is "A" or "AAAA" (the default)`
	case !needsAddress && hasAddress:
		return "Address not valid for this type",
			`routeros_ip_dns_static: "address" only applies when type is "A" or "AAAA". Setting it on any ` +
				`other type is silently dropped on Create and, worse, silently rewrites the record's type to ` +
				`"A" - destroying its other type-specific fields - on Update. Confirmed live on RouterOS 7.23.2.`
	default:
		return "", ""
	}
}

// ValidateConfig catches address/type conflicts at plan time.
func (r *IPDNSStaticResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var cfg IPDNSStaticModel
	if diags := req.Config.Get(ctx, &cfg); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	if summary, detail := ipDNSStaticNameOrRegexp(cfg); summary != "" {
		resp.Diagnostics.AddAttributeError(path.Root("name"), summary, detail)
	}
	if cfg.Type.IsUnknown() || cfg.Address.IsUnknown() {
		return
	}
	if summary, detail := ipDNSStaticAddressTypeConflict(cfg.Type.ValueString(), ipDNSStaticHasAddress(cfg)); summary != "" {
		resp.Diagnostics.AddAttributeError(path.Root("address"), summary, detail)
	}
}

func ipDNSStaticHasAddress(m IPDNSStaticModel) bool {
	return !m.Address.IsNull() && !m.Address.IsUnknown() && m.Address.ValueString() != ""
}

// A record is keyed by one or the other; RouterOS accepts a regexp entry with
// no name, and rejects an entry carrying both.
func ipDNSStaticNameOrRegexp(m IPDNSStaticModel) (summary, detail string) {
	if m.Name.IsUnknown() || m.Regexp.IsUnknown() {
		return "", ""
	}
	name := !m.Name.IsNull() && m.Name.ValueString() != ""
	re := !m.Regexp.IsNull() && m.Regexp.ValueString() != ""
	switch {
	case !name && !re:
		return "Missing name or regexp", `routeros_ip_dns_static requires either "name" or "regexp"`
	case name && re:
		return "Conflicting name and regexp",
			`routeros_ip_dns_static takes either "name" or "regexp", not both`
	}
	return "", ""
}

// ValidateConfig only sees config, so a value arriving from a variable or
// another resource is still Unknown there. Re-check once it is resolved.
func ipDNSStaticCheckResolved(m IPDNSStaticModel) (summary, detail string) {
	if summary, detail := ipDNSStaticNameOrRegexp(m); summary != "" {
		return summary, detail
	}
	if m.Type.IsUnknown() {
		return "", ""
	}
	return ipDNSStaticAddressTypeConflict(m.Type.ValueString(), ipDNSStaticHasAddress(m))
}

func (r *IPDNSStaticResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPDNSStaticModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if summary, detail := ipDNSStaticCheckResolved(plan); summary != "" {
		resp.Diagnostics.AddError(summary, detail)
		return
	}
	body := client.Object{}
	if ipDNSStaticHasAddress(plan) {
		body["address"] = plan.Address.ValueString()
	}
	if !(plan.AddressList.IsNull() || plan.AddressList.IsUnknown()) {
		body["address-list"] = plan.AddressList.ValueString()
	}
	if !(plan.Cname.IsNull() || plan.Cname.IsUnknown()) {
		body["cname"] = plan.Cname.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.ForwardTo.IsNull() || plan.ForwardTo.IsUnknown()) {
		body["forward-to"] = plan.ForwardTo.ValueString()
	}
	if !(plan.MatchSubdomain.IsNull() || plan.MatchSubdomain.IsUnknown()) {
		body["match-subdomain"] = plan.MatchSubdomain.ValueString()
	}
	if !(plan.MxExchange.IsNull() || plan.MxExchange.IsUnknown()) {
		body["mx-exchange"] = plan.MxExchange.ValueString()
	}
	if !(plan.MxPreference.IsNull() || plan.MxPreference.IsUnknown()) {
		body["mx-preference"] = plan.MxPreference.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Ns.IsNull() || plan.Ns.IsUnknown()) {
		body["ns"] = plan.Ns.ValueString()
	}
	if !(plan.Regexp.IsNull() || plan.Regexp.IsUnknown()) {
		body["regexp"] = plan.Regexp.ValueString()
	}
	if !(plan.SrvPort.IsNull() || plan.SrvPort.IsUnknown()) {
		body["srv-port"] = plan.SrvPort.ValueString()
	}
	if !(plan.SrvPriority.IsNull() || plan.SrvPriority.IsUnknown()) {
		body["srv-priority"] = plan.SrvPriority.ValueString()
	}
	if !(plan.SrvTarget.IsNull() || plan.SrvTarget.IsUnknown()) {
		body["srv-target"] = plan.SrvTarget.ValueString()
	}
	if !(plan.SrvWeight.IsNull() || plan.SrvWeight.IsUnknown()) {
		body["srv-weight"] = plan.SrvWeight.ValueString()
	}
	if !(plan.Text.IsNull() || plan.Text.IsUnknown()) {
		body["text"] = plan.Text.ValueString()
	}
	if !(plan.Ttl.IsNull() || plan.Ttl.IsUnknown()) {
		body["ttl"] = plan.Ttl.ValueString()
	}
	if !(plan.Type.IsNull() || plan.Type.IsUnknown()) {
		body["type"] = plan.Type.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/dns/static", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/dns/static failed", err.Error())
		return
	}
	iPDNSStaticApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDNSStaticResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPDNSStaticModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/dns/static", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/dns/static failed", err.Error())
		return
	}
	iPDNSStaticApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

// ipDNSStaticShouldWriteAddress reports whether Update should include address in the PATCH body.
func ipDNSStaticShouldWriteAddress(plan, state IPDNSStaticModel) bool {
	if plan.Address.IsUnknown() || plan.Address.Equal(state.Address) {
		return false
	}
	return !plan.Address.IsNull() && plan.Address.ValueString() != ""
}

func (r *IPDNSStaticResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPDNSStaticModel
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
	if summary, detail := ipDNSStaticCheckResolved(plan); summary != "" {
		resp.Diagnostics.AddError(summary, detail)
		return
	}
	body := client.Object{}
	if ipDNSStaticShouldWriteAddress(plan, state) {
		body["address"] = plan.Address.ValueString()
	}
	if !plan.AddressList.Equal(state.AddressList) && !plan.AddressList.IsUnknown() {
		body["address-list"] = plan.AddressList.ValueString()
	}
	if !plan.Cname.Equal(state.Cname) && !plan.Cname.IsUnknown() {
		body["cname"] = plan.Cname.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.ForwardTo.Equal(state.ForwardTo) && !plan.ForwardTo.IsUnknown() {
		body["forward-to"] = plan.ForwardTo.ValueString()
	}
	if !plan.MatchSubdomain.Equal(state.MatchSubdomain) && !plan.MatchSubdomain.IsUnknown() {
		body["match-subdomain"] = plan.MatchSubdomain.ValueString()
	}
	if !plan.MxExchange.Equal(state.MxExchange) && !plan.MxExchange.IsUnknown() {
		body["mx-exchange"] = plan.MxExchange.ValueString()
	}
	if !plan.MxPreference.Equal(state.MxPreference) && !plan.MxPreference.IsUnknown() {
		body["mx-preference"] = plan.MxPreference.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Ns.Equal(state.Ns) && !plan.Ns.IsUnknown() {
		body["ns"] = plan.Ns.ValueString()
	}
	if !plan.Regexp.Equal(state.Regexp) && !plan.Regexp.IsUnknown() {
		body["regexp"] = plan.Regexp.ValueString()
	}
	if !plan.SrvPort.Equal(state.SrvPort) && !plan.SrvPort.IsUnknown() {
		body["srv-port"] = plan.SrvPort.ValueString()
	}
	if !plan.SrvPriority.Equal(state.SrvPriority) && !plan.SrvPriority.IsUnknown() {
		body["srv-priority"] = plan.SrvPriority.ValueString()
	}
	if !plan.SrvTarget.Equal(state.SrvTarget) && !plan.SrvTarget.IsUnknown() {
		body["srv-target"] = plan.SrvTarget.ValueString()
	}
	if !plan.SrvWeight.Equal(state.SrvWeight) && !plan.SrvWeight.IsUnknown() {
		body["srv-weight"] = plan.SrvWeight.ValueString()
	}
	if !plan.Text.Equal(state.Text) && !plan.Text.IsUnknown() {
		body["text"] = plan.Text.ValueString()
	}
	if !plan.Ttl.Equal(state.Ttl) && !plan.Ttl.IsUnknown() {
		body["ttl"] = plan.Ttl.ValueString()
	}
	if !plan.Type.Equal(state.Type) && !plan.Type.IsUnknown() {
		body["type"] = plan.Type.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/dns/static", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/dns/static failed", err.Error())
			return
		}
		iPDNSStaticApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDNSStaticResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPDNSStaticModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/dns/static", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/dns/static failed", err.Error())
	}
}

func (r *IPDNSStaticResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPDNSStaticLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/dns/static matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPDNSStaticLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPDNSStaticLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/dns/static", id)
}

func iPDNSStaticApply(ctx context.Context, obj client.Object, m *IPDNSStaticModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["address"]; ok {
		if v != "" {
			m.Address = types.StringValue(v)
		} else {
			m.Address = types.StringNull()
		}
	}
	if v, ok := obj["address-list"]; ok {
		if v != "" {
			m.AddressList = types.StringValue(v)
		} else {
			m.AddressList = types.StringNull()
		}
	}
	if v, ok := obj["cname"]; ok {
		if v != "" {
			m.Cname = types.StringValue(v)
		} else {
			m.Cname = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
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
	if v, ok := obj["forward-to"]; ok {
		if v != "" {
			m.ForwardTo = types.StringValue(v)
		} else {
			m.ForwardTo = types.StringNull()
		}
	}
	if v, ok := obj["match-subdomain"]; ok {
		if v != "" {
			m.MatchSubdomain = types.StringValue(v)
		} else {
			m.MatchSubdomain = types.StringNull()
		}
	}
	if v, ok := obj["mx-exchange"]; ok {
		if v != "" {
			m.MxExchange = types.StringValue(v)
		} else {
			m.MxExchange = types.StringNull()
		}
	}
	if v, ok := obj["mx-preference"]; ok {
		if v != "" {
			m.MxPreference = types.StringValue(v)
		} else {
			m.MxPreference = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["ns"]; ok {
		if v != "" {
			m.Ns = types.StringValue(v)
		} else {
			m.Ns = types.StringNull()
		}
	}
	if v, ok := obj["regexp"]; ok {
		if v != "" {
			m.Regexp = types.StringValue(v)
		} else {
			m.Regexp = types.StringNull()
		}
	}
	if v, ok := obj["srv-port"]; ok {
		if v != "" {
			m.SrvPort = types.StringValue(v)
		} else {
			m.SrvPort = types.StringNull()
		}
	}
	if v, ok := obj["srv-priority"]; ok {
		if v != "" {
			m.SrvPriority = types.StringValue(v)
		} else {
			m.SrvPriority = types.StringNull()
		}
	}
	if v, ok := obj["srv-target"]; ok {
		if v != "" {
			m.SrvTarget = types.StringValue(v)
		} else {
			m.SrvTarget = types.StringNull()
		}
	}
	if v, ok := obj["srv-weight"]; ok {
		if v != "" {
			m.SrvWeight = types.StringValue(v)
		} else {
			m.SrvWeight = types.StringNull()
		}
	}
	if v, ok := obj["text"]; ok {
		if v != "" {
			m.Text = types.StringValue(v)
		} else {
			m.Text = types.StringNull()
		}
	}
	if v, ok := obj["ttl"]; ok {
		if v != "" {
			m.Ttl = types.StringValue(v)
		} else {
			m.Ttl = types.StringNull()
		}
	}
	if v, ok := obj["type"]; ok {
		if v != "" {
			m.Type = types.StringValue(v)
		} else {
			m.Type = types.StringNull()
		}
	}
}
