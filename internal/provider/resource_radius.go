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
	_ resource.Resource                = &RADIUSResource{}
	_ resource.ResourceWithImportState = &RADIUSResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RADIUSResource struct {
	reg *client.Registry
}

type RADIUSModel struct {
	ID                 types.String `tfsdk:"id"`
	AccountingBackup   types.Bool   `tfsdk:"accounting_backup"`
	AccountingPort     types.Int64  `tfsdk:"accounting_port"`
	Address            types.String `tfsdk:"address"`
	AuthenticationPort types.Int64  `tfsdk:"authentication_port"`
	CalledID           types.String `tfsdk:"called_id"`
	Certificate        types.String `tfsdk:"certificate"`
	Comment            types.String `tfsdk:"comment"`
	Disabled           types.Bool   `tfsdk:"disabled"`
	Domain             types.String `tfsdk:"domain"`
	Protocol           types.String `tfsdk:"protocol"`
	Radsec             types.String `tfsdk:"radsec"`
	RadsecTimeout      types.String `tfsdk:"radsec_timeout"`
	Realm              types.String `tfsdk:"realm"`
	RequireMessageAuth types.String `tfsdk:"require_message_auth"`
	ResetStatus        types.String `tfsdk:"reset_status"`
	Secret             types.String `tfsdk:"secret"`
	Service            types.String `tfsdk:"service"`
	SrcAddress         types.String `tfsdk:"src_address"`
	Timeout            types.String `tfsdk:"timeout"`
	UDP                types.String `tfsdk:"udp"`
	Router             types.String `tfsdk:"router"`
}

func NewRADIUSResource() resource.Resource { return &RADIUSResource{} }

func (r *RADIUSResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_radius"
}

func (r *RADIUSResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RADIUSResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Auto-test requires a typed-reference precondition (e.g. an existing peer,\ninstance, bridge of the specific kind). The current acc-test generator's\ngeneric data.routeros_interface.all lookup can't satisfy these. Use this\nresource manually with explicit references to a precondition resource\nin your config.\n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"accounting_backup": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the configuration is for the backup RADIUS server",
			},
			"accounting_port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RADIUS server port used for accounting",
			},
			"address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IPv4 or IPv6 address of RADIUS server. The following formats are accepted: - \u00a0 ipv4 - \u00a0 ipv4 @ vrf - \u00a0 ipv6 - \u00a0 ipv6 @ vrf",
			},
			"authentication_port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RADIUS server port used for authentication.",
			},
			"called_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Value depends on Point-to-Point protocol: PPPoE - service name, PPTP - server's IP address, L2TP - server's IP address.",
			},
			"certificate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Certificate file to use for communicating with RADIUS Server with RadSec enabled.",
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
			"domain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Microsoft Windows domain of client passed to RADIUS servers that require domain validation.",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies the protocol to use when communicating with the RADIUS Server.",
				Validators:  []validator.String{schemautil.OneOf([]string{"udp", "radsec"}...)},
			},
			"radsec": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"radsec_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Timeout after which the request should be resent over RadSec protocol.",
			},
			"realm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Explicitly stated realm (user domain), so the users do not have to provide proper ISP domain name in the user name.",
			},
			"require_message_auth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies if Message-Authenticator attributes are required.",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "no", "yes-for-request-resp"}...)},
			},
			"reset_status": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"secret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "The shared secret used to access the RADIUS server.",
			},
			"service": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Router services that will use this RADIUS server: hotspot - HotSpot authentication service login - router's local user authentication ppp - Point-to-Point clients authentication wireless - wireless client authentication dhcp - DHCP protocol client authentication (client's MAC address is sent as User-Name) ipsec - ipsec client authentification dot1x - dot1x authentification",
				Validators:  []validator.String{schemautil.OneOf([]string{"ppp", "login", "hotspot", "wireless", "dhcp", "ipsec", "dot1x"}...)},
			},
			"src_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Source IP/IPv6 address of the packets sent to the RADIUS server",
			},
			"timeout": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "Timeout after which the request should be resent.",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"udp": schema.StringAttribute{
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

func (r *RADIUSResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RADIUSModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AccountingBackup.IsNull() || plan.AccountingBackup.IsUnknown()) {
		body["accounting-backup"] = client.FormatBool(plan.AccountingBackup.ValueBool())
	}
	if !(plan.AccountingPort.IsNull() || plan.AccountingPort.IsUnknown()) {
		body["accounting-port"] = client.FormatInt64(plan.AccountingPort.ValueInt64())
	}
	if !(plan.Address.IsNull() || plan.Address.IsUnknown()) {
		body["address"] = plan.Address.ValueString()
	}
	if !(plan.AuthenticationPort.IsNull() || plan.AuthenticationPort.IsUnknown()) {
		body["authentication-port"] = client.FormatInt64(plan.AuthenticationPort.ValueInt64())
	}
	if !(plan.CalledID.IsNull() || plan.CalledID.IsUnknown()) {
		body["called-id"] = plan.CalledID.ValueString()
	}
	if !(plan.Certificate.IsNull() || plan.Certificate.IsUnknown()) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Domain.IsNull() || plan.Domain.IsUnknown()) {
		body["domain"] = plan.Domain.ValueString()
	}
	if !(plan.Protocol.IsNull() || plan.Protocol.IsUnknown()) {
		body["protocol"] = plan.Protocol.ValueString()
	}
	if !(plan.RadsecTimeout.IsNull() || plan.RadsecTimeout.IsUnknown()) {
		body["radsec-timeout"] = plan.RadsecTimeout.ValueString()
	}
	if !(plan.Realm.IsNull() || plan.Realm.IsUnknown()) {
		body["realm"] = plan.Realm.ValueString()
	}
	if !(plan.RequireMessageAuth.IsNull() || plan.RequireMessageAuth.IsUnknown()) {
		body["require-message-auth"] = plan.RequireMessageAuth.ValueString()
	}
	if !(plan.Secret.IsNull() || plan.Secret.IsUnknown()) {
		body["secret"] = plan.Secret.ValueString()
	}
	if !(plan.Service.IsNull() || plan.Service.IsUnknown()) {
		body["service"] = plan.Service.ValueString()
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !(plan.Timeout.IsNull() || plan.Timeout.IsUnknown()) {
		body["timeout"] = plan.Timeout.ValueString()
	}
	obj, err := c.Add(ctx, "/radius", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /radius failed", err.Error())
		return
	}
	rADIUSApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RADIUSResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RADIUSModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/radius", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /radius failed", err.Error())
		return
	}
	rADIUSApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RADIUSResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RADIUSModel
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
	if !plan.AccountingBackup.Equal(state.AccountingBackup) && !plan.AccountingBackup.IsUnknown() {
		body["accounting-backup"] = client.FormatBool(plan.AccountingBackup.ValueBool())
	}
	if !plan.AccountingPort.Equal(state.AccountingPort) && !plan.AccountingPort.IsUnknown() {
		body["accounting-port"] = client.FormatInt64(plan.AccountingPort.ValueInt64())
	}
	if !plan.Address.Equal(state.Address) && !plan.Address.IsUnknown() {
		body["address"] = plan.Address.ValueString()
	}
	if !plan.AuthenticationPort.Equal(state.AuthenticationPort) && !plan.AuthenticationPort.IsUnknown() {
		body["authentication-port"] = client.FormatInt64(plan.AuthenticationPort.ValueInt64())
	}
	if !plan.CalledID.Equal(state.CalledID) && !plan.CalledID.IsUnknown() {
		body["called-id"] = plan.CalledID.ValueString()
	}
	if !plan.Certificate.Equal(state.Certificate) && !plan.Certificate.IsUnknown() {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Domain.Equal(state.Domain) && !plan.Domain.IsUnknown() {
		body["domain"] = plan.Domain.ValueString()
	}
	if !plan.Protocol.Equal(state.Protocol) && !plan.Protocol.IsUnknown() {
		body["protocol"] = plan.Protocol.ValueString()
	}
	if !plan.RadsecTimeout.Equal(state.RadsecTimeout) && !plan.RadsecTimeout.IsUnknown() {
		body["radsec-timeout"] = plan.RadsecTimeout.ValueString()
	}
	if !plan.Realm.Equal(state.Realm) && !plan.Realm.IsUnknown() {
		body["realm"] = plan.Realm.ValueString()
	}
	if !plan.RequireMessageAuth.Equal(state.RequireMessageAuth) && !plan.RequireMessageAuth.IsUnknown() {
		body["require-message-auth"] = plan.RequireMessageAuth.ValueString()
	}
	if !plan.Secret.Equal(state.Secret) && !plan.Secret.IsUnknown() {
		body["secret"] = plan.Secret.ValueString()
	}
	if !plan.Service.Equal(state.Service) && !plan.Service.IsUnknown() {
		body["service"] = plan.Service.ValueString()
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) && !plan.SrcAddress.IsUnknown() {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !plan.Timeout.Equal(state.Timeout) && !plan.Timeout.IsUnknown() {
		body["timeout"] = plan.Timeout.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/radius", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /radius failed", err.Error())
			return
		}
		rADIUSApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RADIUSResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RADIUSModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/radius", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /radius failed", err.Error())
	}
}

func (r *RADIUSResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := rADIUSLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /radius matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// rADIUSLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func rADIUSLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/radius", id)
}

func rADIUSApply(ctx context.Context, obj client.Object, m *RADIUSModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["accounting-backup"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.AccountingBackup = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.AccountingBackup = types.BoolValue(true)
		} else {
			m.AccountingBackup = types.BoolNull()
		}
	}
	if v, ok := obj["accounting-port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.AccountingPort = types.Int64Value(n)
		} else {
			m.AccountingPort = types.Int64Null()
		}
	} else {
		m.AccountingPort = types.Int64Null()
	}
	if v, ok := obj["address"]; ok {
		if v != "" {
			m.Address = types.StringValue(v)
		} else {
			m.Address = types.StringNull()
		}
	}
	if v, ok := obj["authentication-port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.AuthenticationPort = types.Int64Value(n)
		} else {
			m.AuthenticationPort = types.Int64Null()
		}
	} else {
		m.AuthenticationPort = types.Int64Null()
	}
	if v, ok := obj["called-id"]; ok {
		if v != "" {
			m.CalledID = types.StringValue(v)
		} else {
			m.CalledID = types.StringNull()
		}
	}
	if v, ok := obj["certificate"]; ok {
		if v != "" {
			m.Certificate = types.StringValue(v)
		} else {
			m.Certificate = types.StringNull()
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
	if v, ok := obj["domain"]; ok {
		if v != "" {
			m.Domain = types.StringValue(v)
		} else {
			m.Domain = types.StringNull()
		}
	}
	if v, ok := obj["protocol"]; ok {
		if v != "" {
			m.Protocol = types.StringValue(v)
		} else {
			m.Protocol = types.StringNull()
		}
	}
	if v, ok := obj["radsec"]; ok {
		if v != "" {
			m.Radsec = types.StringValue(v)
		} else {
			m.Radsec = types.StringNull()
		}
	}
	if v, ok := obj["radsec-timeout"]; ok {
		if v != "" {
			m.RadsecTimeout = types.StringValue(v)
		} else {
			m.RadsecTimeout = types.StringNull()
		}
	}
	if v, ok := obj["realm"]; ok {
		if v != "" {
			m.Realm = types.StringValue(v)
		} else {
			m.Realm = types.StringNull()
		}
	}
	if v, ok := obj["require-message-auth"]; ok {
		if v != "" {
			m.RequireMessageAuth = types.StringValue(v)
		} else {
			m.RequireMessageAuth = types.StringNull()
		}
	}
	if v, ok := obj["reset-status"]; ok {
		if v != "" {
			m.ResetStatus = types.StringValue(v)
		} else {
			m.ResetStatus = types.StringNull()
		}
	}
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.Secret already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["secret"]; ok && v != "" {
		_ = v
		if v != "" {
			m.Secret = types.StringValue(v)
		} else {
			m.Secret = types.StringNull()
		}
	} else if m.Secret.IsUnknown() {
		m.Secret = types.StringNull()
	}
	if v, ok := obj["service"]; ok {
		if v != "" {
			m.Service = types.StringValue(v)
		} else {
			m.Service = types.StringNull()
		}
	}
	if v, ok := obj["src-address"]; ok {
		if v != "" {
			m.SrcAddress = types.StringValue(v)
		} else {
			m.SrcAddress = types.StringNull()
		}
	}
	if v, ok := obj["timeout"]; ok && v != "" {
		m.Timeout = types.StringValue(v)
	} else {
		m.Timeout = types.StringNull()
	}
	if v, ok := obj["udp"]; ok {
		if v != "" {
			m.UDP = types.StringValue(v)
		} else {
			m.UDP = types.StringNull()
		}
	}
}
