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
	_ resource.Resource                = &IPHotspotUserProfileResource{}
	_ resource.ResourceWithImportState = &IPHotspotUserProfileResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPHotspotUserProfileResource struct {
	reg *client.Registry
}

type IPHotspotUserProfileModel struct {
	ID                types.String `tfsdk:"id"`
	AddMACCookie      types.Bool   `tfsdk:"add_mac_cookie"`
	AddressList       types.String `tfsdk:"address_list"`
	Default           types.Bool   `tfsdk:"default"`
	IdleTimeout       types.String `tfsdk:"idle_timeout"`
	KeepaliveTimeout  types.String `tfsdk:"keepalive_timeout"`
	MACCookieTimeout  types.String `tfsdk:"mac_cookie_timeout"`
	Name              types.String `tfsdk:"name"`
	SharedUsers       types.Int64  `tfsdk:"shared_users"`
	StatusAutorefresh types.String `tfsdk:"status_autorefresh"`
	TransparentProxy  types.Bool   `tfsdk:"transparent_proxy"`
	Router            types.String `tfsdk:"router"`
}

func NewIPHotspotUserProfileResource() resource.Resource { return &IPHotspotUserProfileResource{} }

func (r *IPHotspotUserProfileResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_hotspot_user_profile"
}

func (r *IPHotspotUserProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *IPHotspotUserProfileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/hotspot/user/profile`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"add_mac_cookie": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"address_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"default": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"idle_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"keepalive_timeout": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"mac_cookie_timeout": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"shared_users": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"status_autorefresh": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"transparent_proxy": schema.BoolAttribute{
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

func (r *IPHotspotUserProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPHotspotUserProfileModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AddMACCookie.IsNull() || plan.AddMACCookie.IsUnknown()) {
		body["add-mac-cookie"] = client.FormatBool(plan.AddMACCookie.ValueBool())
	}
	if !(plan.AddressList.IsNull() || plan.AddressList.IsUnknown()) {
		body["address-list"] = plan.AddressList.ValueString()
	}
	if !(plan.IdleTimeout.IsNull() || plan.IdleTimeout.IsUnknown()) {
		body["idle-timeout"] = plan.IdleTimeout.ValueString()
	}
	if !(plan.KeepaliveTimeout.IsNull() || plan.KeepaliveTimeout.IsUnknown()) {
		body["keepalive-timeout"] = plan.KeepaliveTimeout.ValueString()
	}
	if !(plan.MACCookieTimeout.IsNull() || plan.MACCookieTimeout.IsUnknown()) {
		body["mac-cookie-timeout"] = plan.MACCookieTimeout.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.SharedUsers.IsNull() || plan.SharedUsers.IsUnknown()) {
		body["shared-users"] = client.FormatInt64(plan.SharedUsers.ValueInt64())
	}
	if !(plan.StatusAutorefresh.IsNull() || plan.StatusAutorefresh.IsUnknown()) {
		body["status-autorefresh"] = plan.StatusAutorefresh.ValueString()
	}
	if !(plan.TransparentProxy.IsNull() || plan.TransparentProxy.IsUnknown()) {
		body["transparent-proxy"] = client.FormatBool(plan.TransparentProxy.ValueBool())
	}
	obj, err := c.Add(ctx, "/ip/hotspot/user/profile", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/hotspot/user/profile failed", err.Error())
		return
	}
	iPHotspotUserProfileApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPHotspotUserProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPHotspotUserProfileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/hotspot/user/profile", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/hotspot/user/profile failed", err.Error())
		return
	}
	iPHotspotUserProfileApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPHotspotUserProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPHotspotUserProfileModel
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
	if !plan.AddMACCookie.Equal(state.AddMACCookie) {
		body["add-mac-cookie"] = client.FormatBool(plan.AddMACCookie.ValueBool())
	}
	if !plan.AddressList.Equal(state.AddressList) {
		body["address-list"] = plan.AddressList.ValueString()
	}
	if !plan.IdleTimeout.Equal(state.IdleTimeout) {
		body["idle-timeout"] = plan.IdleTimeout.ValueString()
	}
	if !plan.KeepaliveTimeout.Equal(state.KeepaliveTimeout) {
		body["keepalive-timeout"] = plan.KeepaliveTimeout.ValueString()
	}
	if !plan.MACCookieTimeout.Equal(state.MACCookieTimeout) {
		body["mac-cookie-timeout"] = plan.MACCookieTimeout.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.SharedUsers.Equal(state.SharedUsers) {
		body["shared-users"] = client.FormatInt64(plan.SharedUsers.ValueInt64())
	}
	if !plan.StatusAutorefresh.Equal(state.StatusAutorefresh) {
		body["status-autorefresh"] = plan.StatusAutorefresh.ValueString()
	}
	if !plan.TransparentProxy.Equal(state.TransparentProxy) {
		body["transparent-proxy"] = client.FormatBool(plan.TransparentProxy.ValueBool())
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/hotspot/user/profile", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/hotspot/user/profile failed", err.Error())
			return
		}
		iPHotspotUserProfileApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPHotspotUserProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPHotspotUserProfileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/hotspot/user/profile", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/hotspot/user/profile failed", err.Error())
	}
}

func (r *IPHotspotUserProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPHotspotUserProfileLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/hotspot/user/profile matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPHotspotUserProfileLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPHotspotUserProfileLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/hotspot/user/profile", id)
}

func iPHotspotUserProfileApply(ctx context.Context, obj client.Object, m *IPHotspotUserProfileModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["add-mac-cookie"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AddMACCookie = types.BoolValue(b)
		} else {
			m.AddMACCookie = types.BoolNull()
		}
	} else {
		m.AddMACCookie = types.BoolNull()
	}
	if v, ok := obj["address-list"]; ok {
		_ = v
		if v != "" {
			m.AddressList = types.StringValue(v)
		} else {
			m.AddressList = types.StringNull()
		}
	} else {
		m.AddressList = types.StringNull()
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
	if v, ok := obj["idle-timeout"]; ok {
		_ = v
		if v != "" {
			m.IdleTimeout = types.StringValue(v)
		} else {
			m.IdleTimeout = types.StringNull()
		}
	} else {
		m.IdleTimeout = types.StringNull()
	}
	if v, ok := obj["keepalive-timeout"]; ok {
		_ = v
		if v != "" {
			m.KeepaliveTimeout = types.StringValue(v)
		} else {
			m.KeepaliveTimeout = types.StringNull()
		}
	} else {
		m.KeepaliveTimeout = types.StringNull()
	}
	if v, ok := obj["mac-cookie-timeout"]; ok {
		_ = v
		if v != "" {
			m.MACCookieTimeout = types.StringValue(v)
		} else {
			m.MACCookieTimeout = types.StringNull()
		}
	} else {
		m.MACCookieTimeout = types.StringNull()
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
	if v, ok := obj["shared-users"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.SharedUsers = types.Int64Value(n)
		} else {
			m.SharedUsers = types.Int64Null()
		}
	} else {
		m.SharedUsers = types.Int64Null()
	}
	if v, ok := obj["status-autorefresh"]; ok {
		_ = v
		if v != "" {
			m.StatusAutorefresh = types.StringValue(v)
		} else {
			m.StatusAutorefresh = types.StringNull()
		}
	} else {
		m.StatusAutorefresh = types.StringNull()
	}
	if v, ok := obj["transparent-proxy"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.TransparentProxy = types.BoolValue(b)
		} else {
			m.TransparentProxy = types.BoolNull()
		}
	} else {
		m.TransparentProxy = types.BoolNull()
	}
}
