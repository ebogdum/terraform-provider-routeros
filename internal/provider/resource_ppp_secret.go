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
	_ resource.Resource                = &PPPSecretResource{}
	_ resource.ResourceWithImportState = &PPPSecretResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type PPPSecretResource struct {
	reg *client.Registry
}

type PPPSecretModel struct {
	ID                   types.String  `tfsdk:"id"`
	CallerID             types.String  `tfsdk:"caller_id"`
	Comment              types.String  `tfsdk:"comment"`
	Disabled             types.Bool    `tfsdk:"disabled"`
	IPV6                 types.String  `tfsdk:"ipv6"`
	IPV6Routes           types.String  `tfsdk:"ipv6_routes"`
	LastCallerID         types.String  `tfsdk:"last_caller_id"`
	LastDisconnectReason types.String  `tfsdk:"last_disconnect_reason"`
	LastLoggedOut        types.String  `tfsdk:"last_logged_out"`
	LimitBytesIn         rosRateValue  `tfsdk:"limit_bytes_in"`
	LimitBytesOut        rosRateValue  `tfsdk:"limit_bytes_out"`
	LocalAddress         types.String  `tfsdk:"local_address"`
	Name                 types.String  `tfsdk:"name"`
	Password             types.String  `tfsdk:"password"`
	Profile              types.String  `tfsdk:"profile"`
	RemoteAddress        types.String  `tfsdk:"remote_address"`
	RemoteIPV6Prefix     hostAddrValue `tfsdk:"remote_ipv6_prefix"`
	Routes               types.String  `tfsdk:"routes"`
	Service              types.String  `tfsdk:"service"`
	Router               types.String  `tfsdk:"router"`
}

func NewPPPSecretResource() resource.Resource { return &PPPSecretResource{} }

func (r *PPPSecretResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ppp_secret"
}

func (r *PPPSecretResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *PPPSecretResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ppp/secret`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"caller_id": schema.StringAttribute{
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
			"ipv6": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"ipv6_routes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"last_caller_id": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"last_disconnect_reason": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "peer-request", "hung-up", "idle-timeout", "session-timeout", "reset", "reboot", "port-error", "nas-error", "nas-request"}...)},
			},
			"last_logged_out": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"limit_bytes_in": schema.StringAttribute{
				CustomType:  rosRateType{},
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"limit_bytes_out": schema.StringAttribute{
				CustomType:  rosRateType{},
				Optional:    true,
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
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"profile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_ipv6_prefix": schema.StringAttribute{
				CustomType:  hostAddrType{},
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"routes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"service": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"any", "async", "pptp", "pppoe", "l2tp", "ovpn", "sstp"}...)},
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *PPPSecretResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PPPSecretModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.CallerID.IsNull() || plan.CallerID.IsUnknown()) {
		body["caller-id"] = plan.CallerID.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.IPV6Routes.IsNull() || plan.IPV6Routes.IsUnknown()) {
		body["ipv6-routes"] = plan.IPV6Routes.ValueString()
	}
	if !(plan.LimitBytesIn.IsNull() || plan.LimitBytesIn.IsUnknown()) {
		body["limit-bytes-in"] = plan.LimitBytesIn.ValueString()
	}
	if !(plan.LimitBytesOut.IsNull() || plan.LimitBytesOut.IsUnknown()) {
		body["limit-bytes-out"] = plan.LimitBytesOut.ValueString()
	}
	if !(plan.LocalAddress.IsNull() || plan.LocalAddress.IsUnknown()) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Password.IsNull() || plan.Password.IsUnknown()) {
		body["password"] = plan.Password.ValueString()
	}
	if !(plan.Profile.IsNull() || plan.Profile.IsUnknown()) {
		body["profile"] = plan.Profile.ValueString()
	}
	if !(plan.RemoteAddress.IsNull() || plan.RemoteAddress.IsUnknown()) {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !(plan.RemoteIPV6Prefix.IsNull() || plan.RemoteIPV6Prefix.IsUnknown()) {
		body["remote-ipv6-prefix"] = plan.RemoteIPV6Prefix.ValueString()
	}
	if !(plan.Routes.IsNull() || plan.Routes.IsUnknown()) {
		body["routes"] = plan.Routes.ValueString()
	}
	if !(plan.Service.IsNull() || plan.Service.IsUnknown()) {
		body["service"] = plan.Service.ValueString()
	}
	obj, err := c.Add(ctx, "/ppp/secret", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ppp/secret failed", err.Error())
		return
	}
	pPPSecretApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PPPSecretResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PPPSecretModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ppp/secret", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ppp/secret failed", err.Error())
		return
	}
	pPPSecretApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *PPPSecretResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PPPSecretModel
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
	if !plan.CallerID.Equal(state.CallerID) && !plan.CallerID.IsUnknown() {
		body["caller-id"] = plan.CallerID.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.IPV6Routes.Equal(state.IPV6Routes) && !plan.IPV6Routes.IsUnknown() {
		body["ipv6-routes"] = plan.IPV6Routes.ValueString()
	}
	if !plan.LimitBytesIn.Equal(state.LimitBytesIn) && !plan.LimitBytesIn.IsUnknown() {
		body["limit-bytes-in"] = plan.LimitBytesIn.ValueString()
	}
	if !plan.LimitBytesOut.Equal(state.LimitBytesOut) && !plan.LimitBytesOut.IsUnknown() {
		body["limit-bytes-out"] = plan.LimitBytesOut.ValueString()
	}
	if !plan.LocalAddress.Equal(state.LocalAddress) && !plan.LocalAddress.IsUnknown() {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Password.Equal(state.Password) && !plan.Password.IsUnknown() {
		body["password"] = plan.Password.ValueString()
	}
	if !plan.Profile.Equal(state.Profile) && !plan.Profile.IsUnknown() {
		body["profile"] = plan.Profile.ValueString()
	}
	if !plan.RemoteAddress.Equal(state.RemoteAddress) && !plan.RemoteAddress.IsUnknown() {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !plan.RemoteIPV6Prefix.Equal(state.RemoteIPV6Prefix) && !plan.RemoteIPV6Prefix.IsUnknown() {
		body["remote-ipv6-prefix"] = plan.RemoteIPV6Prefix.ValueString()
	}
	if !plan.Routes.Equal(state.Routes) && !plan.Routes.IsUnknown() {
		body["routes"] = plan.Routes.ValueString()
	}
	if !plan.Service.Equal(state.Service) && !plan.Service.IsUnknown() {
		body["service"] = plan.Service.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ppp/secret", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ppp/secret failed", err.Error())
			return
		}
		pPPSecretApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PPPSecretResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PPPSecretModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ppp/secret", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ppp/secret failed", err.Error())
	}
}

func (r *PPPSecretResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := pPPSecretLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ppp/secret matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// pPPSecretLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func pPPSecretLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ppp/secret", id)
}

func pPPSecretApply(ctx context.Context, obj client.Object, m *PPPSecretModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["caller-id"]; ok {
		if v != "" {
			m.CallerID = types.StringValue(v)
		} else {
			m.CallerID = types.StringNull()
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
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["ipv6"]; ok {
		if v != "" {
			m.IPV6 = types.StringValue(v)
		} else {
			m.IPV6 = types.StringNull()
		}
	}
	if v, ok := obj["ipv6-routes"]; ok {
		if v != "" {
			m.IPV6Routes = types.StringValue(v)
		} else {
			m.IPV6Routes = types.StringNull()
		}
	}
	if v, ok := obj["last-caller-id"]; ok {
		if v != "" {
			m.LastCallerID = types.StringValue(v)
		} else {
			m.LastCallerID = types.StringNull()
		}
	}
	if v, ok := obj["last-disconnect-reason"]; ok {
		if v != "" {
			m.LastDisconnectReason = types.StringValue(v)
		} else {
			m.LastDisconnectReason = types.StringNull()
		}
	}
	if v, ok := obj["last-logged-out"]; ok {
		if v != "" {
			m.LastLoggedOut = types.StringValue(v)
		} else {
			m.LastLoggedOut = types.StringNull()
		}
	}
	if v, ok := obj["limit-bytes-in"]; ok {
		_ = v
		if v != "" {
			m.LimitBytesIn = newRosRateValue(v)
		} else {
			m.LimitBytesIn = newRosRateNull()
		}
	} else {
		m.LimitBytesIn = newRosRateNull()
	}
	if v, ok := obj["limit-bytes-out"]; ok {
		_ = v
		if v != "" {
			m.LimitBytesOut = newRosRateValue(v)
		} else {
			m.LimitBytesOut = newRosRateNull()
		}
	} else {
		m.LimitBytesOut = newRosRateNull()
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
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.Password already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["password"]; ok && v != "" {
		_ = v
		if v != "" {
			m.Password = types.StringValue(v)
		} else {
			m.Password = types.StringNull()
		}
	} else if m.Password.IsUnknown() {
		m.Password = types.StringNull()
	}
	if v, ok := obj["profile"]; ok {
		if v != "" {
			m.Profile = types.StringValue(v)
		} else {
			m.Profile = types.StringNull()
		}
	}
	if v, ok := obj["remote-address"]; ok {
		if v != "" {
			m.RemoteAddress = types.StringValue(v)
		} else {
			m.RemoteAddress = types.StringNull()
		}
	}
	if v, ok := obj["remote-ipv6-prefix"]; ok {
		_ = v
		if v != "" {
			m.RemoteIPV6Prefix = newHostAddrValue(v)
		} else {
			m.RemoteIPV6Prefix = newHostAddrNull()
		}
	} else {
		m.RemoteIPV6Prefix = newHostAddrNull()
	}
	if v, ok := obj["routes"]; ok {
		if v != "" {
			m.Routes = types.StringValue(v)
		} else {
			m.Routes = types.StringNull()
		}
	}
	if v, ok := obj["service"]; ok {
		if v != "" {
			m.Service = types.StringValue(v)
		} else {
			m.Service = types.StringNull()
		}
	}
}
