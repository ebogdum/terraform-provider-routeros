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
	_ resource.Resource                = &IPHotspotUserResource{}
	_ resource.ResourceWithImportState = &IPHotspotUserResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPHotspotUserResource struct {
	reg *client.Registry
}

type IPHotspotUserModel struct {
	ID               types.String `tfsdk:"id"`
	Address          types.String `tfsdk:"address"`
	BytesIn          types.Int64  `tfsdk:"bytes_in"`
	BytesOut         types.Int64  `tfsdk:"bytes_out"`
	Comment          types.String `tfsdk:"comment"`
	Def              types.Bool   `tfsdk:"def"`
	Default          types.Bool   `tfsdk:"default"`
	Disabled         types.Bool   `tfsdk:"disabled"`
	Dynamic          types.Bool   `tfsdk:"dynamic"`
	Email            types.String `tfsdk:"email"`
	LimitBytesIn     types.String `tfsdk:"limit_bytes_in"`
	LimitBytesOut    types.String `tfsdk:"limit_bytes_out"`
	LimitBytesTotal  types.String `tfsdk:"limit_bytes_total"`
	LimitUptime      types.String `tfsdk:"limit_uptime"`
	MACAddress       types.String `tfsdk:"mac_address"`
	Name             types.String `tfsdk:"name"`
	Nondef           types.String `tfsdk:"nondef"`
	Nondefro         types.String `tfsdk:"nondefro"`
	OtpSecret        types.String `tfsdk:"otp_secret"`
	PacketsIn        types.Int64  `tfsdk:"packets_in"`
	PacketsOut       types.Int64  `tfsdk:"packets_out"`
	Password         types.String `tfsdk:"password"`
	Profile          types.String `tfsdk:"profile"`
	ResetAllCounters types.String `tfsdk:"reset_all_counters"`
	ResetCounters    types.String `tfsdk:"reset_counters"`
	Routes           types.String `tfsdk:"routes"`
	Server           types.String `tfsdk:"server"`
	Uptime           types.String `tfsdk:"uptime"`
	Router           types.String `tfsdk:"router"`
}

func NewIPHotspotUserResource() resource.Resource { return &IPHotspotUserResource{} }

func (r *IPHotspotUserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_hotspot_user"
}

func (r *IPHotspotUserResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPHotspotUserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/hotspot/user`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"bytes_in": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"bytes_out": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"def": schema.BoolAttribute{
				Optional:           true,
				Computed:           true,
				Description:        "",
				DeprecationMessage: "Not a RouterOS REST property (WebFig-only spelling of `default`); RouterOS rejects it. Read-only and ignored on write - use `default`.",
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
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"email": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"limit_bytes_in": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"limit_bytes_out": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"limit_bytes_total": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"limit_uptime": schema.StringAttribute{
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
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"nondef": schema.StringAttribute{
				Optional:           true,
				Computed:           true,
				Description:        "",
				DeprecationMessage: "Not a RouterOS REST property (WebFig-only spelling of `default`); RouterOS rejects it. Read-only and ignored on write - use `default`.",
			},
			"nondefro": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"otp_secret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"packets_in": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"packets_out": schema.Int64Attribute{
				Computed:    true,
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
			"reset_all_counters": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"reset_counters": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"routes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"uptime": schema.StringAttribute{
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPHotspotUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPHotspotUserModel
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
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Email.IsNull() || plan.Email.IsUnknown()) {
		body["email"] = plan.Email.ValueString()
	}
	if !(plan.LimitBytesIn.IsNull() || plan.LimitBytesIn.IsUnknown()) {
		body["limit-bytes-in"] = plan.LimitBytesIn.ValueString()
	}
	if !(plan.LimitBytesOut.IsNull() || plan.LimitBytesOut.IsUnknown()) {
		body["limit-bytes-out"] = plan.LimitBytesOut.ValueString()
	}
	if !(plan.LimitBytesTotal.IsNull() || plan.LimitBytesTotal.IsUnknown()) {
		body["limit-bytes-total"] = plan.LimitBytesTotal.ValueString()
	}
	if !(plan.LimitUptime.IsNull() || plan.LimitUptime.IsUnknown()) {
		body["limit-uptime"] = plan.LimitUptime.ValueString()
	}
	if !(plan.MACAddress.IsNull() || plan.MACAddress.IsUnknown()) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.OtpSecret.IsNull() || plan.OtpSecret.IsUnknown()) {
		body["otp-secret"] = plan.OtpSecret.ValueString()
	}
	if !(plan.Password.IsNull() || plan.Password.IsUnknown()) {
		body["password"] = plan.Password.ValueString()
	}
	if !(plan.Profile.IsNull() || plan.Profile.IsUnknown()) {
		body["profile"] = plan.Profile.ValueString()
	}
	if !(plan.Routes.IsNull() || plan.Routes.IsUnknown()) {
		body["routes"] = plan.Routes.ValueString()
	}
	if !(plan.Server.IsNull() || plan.Server.IsUnknown()) {
		body["server"] = plan.Server.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/hotspot/user", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/hotspot/user failed", err.Error())
		return
	}
	iPHotspotUserApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPHotspotUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPHotspotUserModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/hotspot/user", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/hotspot/user failed", err.Error())
		return
	}
	iPHotspotUserApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPHotspotUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPHotspotUserModel
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
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Email.Equal(state.Email) {
		body["email"] = plan.Email.ValueString()
	}
	if !plan.LimitBytesIn.Equal(state.LimitBytesIn) {
		body["limit-bytes-in"] = plan.LimitBytesIn.ValueString()
	}
	if !plan.LimitBytesOut.Equal(state.LimitBytesOut) {
		body["limit-bytes-out"] = plan.LimitBytesOut.ValueString()
	}
	if !plan.LimitBytesTotal.Equal(state.LimitBytesTotal) {
		body["limit-bytes-total"] = plan.LimitBytesTotal.ValueString()
	}
	if !plan.LimitUptime.Equal(state.LimitUptime) {
		body["limit-uptime"] = plan.LimitUptime.ValueString()
	}
	if !plan.MACAddress.Equal(state.MACAddress) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.OtpSecret.Equal(state.OtpSecret) {
		body["otp-secret"] = plan.OtpSecret.ValueString()
	}
	if !plan.Password.Equal(state.Password) {
		body["password"] = plan.Password.ValueString()
	}
	if !plan.Profile.Equal(state.Profile) {
		body["profile"] = plan.Profile.ValueString()
	}
	if !plan.Routes.Equal(state.Routes) {
		body["routes"] = plan.Routes.ValueString()
	}
	if !plan.Server.Equal(state.Server) {
		body["server"] = plan.Server.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/hotspot/user", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/hotspot/user failed", err.Error())
			return
		}
		iPHotspotUserApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPHotspotUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPHotspotUserModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/hotspot/user", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/hotspot/user failed", err.Error())
	}
}

func (r *IPHotspotUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPHotspotUserLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/hotspot/user matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPHotspotUserLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPHotspotUserLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/hotspot/user", id)
}

func iPHotspotUserApply(ctx context.Context, obj client.Object, m *IPHotspotUserModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
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
	if v, ok := obj["bytes-in"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.BytesIn = types.Int64Value(n)
		} else {
			m.BytesIn = types.Int64Null()
		}
	} else {
		m.BytesIn = types.Int64Null()
	}
	if v, ok := obj["bytes-out"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.BytesOut = types.Int64Value(n)
		} else {
			m.BytesOut = types.Int64Null()
		}
	} else {
		m.BytesOut = types.Int64Null()
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
	if v, ok := obj["def"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Def = types.BoolValue(b)
		} else {
			m.Def = types.BoolNull()
		}
	} else {
		m.Def = types.BoolNull()
	}
	if v, ok := obj["default"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else {
			m.Default = types.BoolNull()
		}
	} else {
		m.Default = types.BoolNull()
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
	if v, ok := obj["email"]; ok {
		_ = v
		if v != "" {
			m.Email = types.StringValue(v)
		} else {
			m.Email = types.StringNull()
		}
	} else {
		m.Email = types.StringNull()
	}
	if v, ok := obj["limit-bytes-in"]; ok {
		_ = v
		if v != "" {
			m.LimitBytesIn = types.StringValue(v)
		} else {
			m.LimitBytesIn = types.StringNull()
		}
	} else {
		m.LimitBytesIn = types.StringNull()
	}
	if v, ok := obj["limit-bytes-out"]; ok {
		_ = v
		if v != "" {
			m.LimitBytesOut = types.StringValue(v)
		} else {
			m.LimitBytesOut = types.StringNull()
		}
	} else {
		m.LimitBytesOut = types.StringNull()
	}
	if v, ok := obj["limit-bytes-total"]; ok {
		_ = v
		if v != "" {
			m.LimitBytesTotal = types.StringValue(v)
		} else {
			m.LimitBytesTotal = types.StringNull()
		}
	} else {
		m.LimitBytesTotal = types.StringNull()
	}
	if v, ok := obj["limit-uptime"]; ok {
		_ = v
		if v != "" {
			m.LimitUptime = types.StringValue(v)
		} else {
			m.LimitUptime = types.StringNull()
		}
	} else {
		m.LimitUptime = types.StringNull()
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
	if v, ok := obj["nondef"]; ok {
		_ = v
		if v != "" {
			m.Nondef = types.StringValue(v)
		} else {
			m.Nondef = types.StringNull()
		}
	} else {
		m.Nondef = types.StringNull()
	}
	if v, ok := obj["nondefro"]; ok {
		_ = v
		if v != "" {
			m.Nondefro = types.StringValue(v)
		} else {
			m.Nondefro = types.StringNull()
		}
	} else {
		m.Nondefro = types.StringNull()
	}
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.OtpSecret already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["otp-secret"]; ok && v != "" {
		_ = v
		if v != "" {
			m.OtpSecret = types.StringValue(v)
		} else {
			m.OtpSecret = types.StringNull()
		}
	} else if m.OtpSecret.IsUnknown() {
		m.OtpSecret = types.StringNull()
	}
	if v, ok := obj["packets-in"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PacketsIn = types.Int64Value(n)
		} else {
			m.PacketsIn = types.Int64Null()
		}
	} else {
		m.PacketsIn = types.Int64Null()
	}
	if v, ok := obj["packets-out"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PacketsOut = types.Int64Value(n)
		} else {
			m.PacketsOut = types.Int64Null()
		}
	} else {
		m.PacketsOut = types.Int64Null()
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
		_ = v
		if v != "" {
			m.Profile = types.StringValue(v)
		} else {
			m.Profile = types.StringNull()
		}
	} else {
		m.Profile = types.StringNull()
	}
	if v, ok := obj["reset-all-counters"]; ok {
		_ = v
		if v != "" {
			m.ResetAllCounters = types.StringValue(v)
		} else {
			m.ResetAllCounters = types.StringNull()
		}
	} else {
		m.ResetAllCounters = types.StringNull()
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
	if v, ok := obj["uptime"]; ok {
		_ = v
		if v != "" {
			m.Uptime = types.StringValue(v)
		} else {
			m.Uptime = types.StringNull()
		}
	} else {
		m.Uptime = types.StringNull()
	}
}
