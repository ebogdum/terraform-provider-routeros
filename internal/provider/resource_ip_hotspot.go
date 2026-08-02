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
	_ resource.Resource                = &IPHotspotResource{}
	_ resource.ResourceWithImportState = &IPHotspotResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPHotspotResource struct {
	reg *client.Registry
}

type IPHotspotModel struct {
	ID               types.String `tfsdk:"id"`
	LoginTimeout     types.String `tfsdk:"login_timeout"`
	IdleTimeout      types.String `tfsdk:"idle_timeout"`
	AddressesPerMac  types.String `tfsdk:"addresses_per_mac"`
	AddressPool      types.String `tfsdk:"address_pool"`
	Disabled         types.Bool   `tfsdk:"disabled"`
	Interface        types.String `tfsdk:"interface"`
	KeepaliveTimeout types.String `tfsdk:"keepalive_timeout"`
	Name             types.String `tfsdk:"name"`
	Profile          types.String `tfsdk:"profile"`
	Router           types.String `tfsdk:"router"`
}

func NewIPHotspotResource() resource.Resource { return &IPHotspotResource{} }

func (r *IPHotspotResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_hotspot"
}

func (r *IPHotspotResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPHotspotResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/hotspot`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"login_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `login-timeout`.",
			},
			"idle_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `idle-timeout`.",
			},
			"addresses_per_mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `addresses-per-mac`.",
			},
			"address_pool": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `address-pool`.",
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
			"keepalive_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Descriptive name of the profile",
			},
			"profile": schema.StringAttribute{
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

func (r *IPHotspotResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPHotspotModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.KeepaliveTimeout.IsNull() || plan.KeepaliveTimeout.IsUnknown()) {
		body["keepalive-timeout"] = plan.KeepaliveTimeout.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Profile.IsNull() || plan.Profile.IsUnknown()) {
		body["profile"] = plan.Profile.ValueString()
	}
	if !(plan.AddressPool.IsNull() || plan.AddressPool.IsUnknown()) {
		body["address-pool"] = plan.AddressPool.ValueString()
	}
	if !(plan.AddressesPerMac.IsNull() || plan.AddressesPerMac.IsUnknown()) {
		body["addresses-per-mac"] = plan.AddressesPerMac.ValueString()
	}
	if !(plan.IdleTimeout.IsNull() || plan.IdleTimeout.IsUnknown()) {
		body["idle-timeout"] = plan.IdleTimeout.ValueString()
	}
	if !(plan.LoginTimeout.IsNull() || plan.LoginTimeout.IsUnknown()) {
		body["login-timeout"] = plan.LoginTimeout.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/hotspot", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/hotspot failed", err.Error())
		return
	}
	iPHotspotApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPHotspotResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPHotspotModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/hotspot", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/hotspot failed", err.Error())
		return
	}
	iPHotspotApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPHotspotResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPHotspotModel
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
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.KeepaliveTimeout.Equal(state.KeepaliveTimeout) && !plan.KeepaliveTimeout.IsUnknown() {
		body["keepalive-timeout"] = plan.KeepaliveTimeout.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Profile.Equal(state.Profile) && !plan.Profile.IsUnknown() {
		body["profile"] = plan.Profile.ValueString()
	}
	if !plan.AddressPool.Equal(state.AddressPool) && !plan.AddressPool.IsUnknown() {
		body["address-pool"] = plan.AddressPool.ValueString()
	}
	if !plan.AddressesPerMac.Equal(state.AddressesPerMac) && !plan.AddressesPerMac.IsUnknown() {
		body["addresses-per-mac"] = plan.AddressesPerMac.ValueString()
	}
	if !plan.IdleTimeout.Equal(state.IdleTimeout) && !plan.IdleTimeout.IsUnknown() {
		body["idle-timeout"] = plan.IdleTimeout.ValueString()
	}
	if !plan.LoginTimeout.Equal(state.LoginTimeout) && !plan.LoginTimeout.IsUnknown() {
		body["login-timeout"] = plan.LoginTimeout.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/hotspot", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/hotspot failed", err.Error())
			return
		}
		iPHotspotApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPHotspotResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPHotspotModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/hotspot", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/hotspot failed", err.Error())
	}
}

func (r *IPHotspotResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPHotspotLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/hotspot matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPHotspotLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPHotspotLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/hotspot", id)
}

func iPHotspotApply(ctx context.Context, obj client.Object, m *IPHotspotModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["login-timeout"]; ok && v != "" {
		m.LoginTimeout = types.StringValue(v)
	} else {
		m.LoginTimeout = types.StringNull()
	}
	if v, ok := obj["idle-timeout"]; ok && v != "" {
		m.IdleTimeout = types.StringValue(v)
	} else {
		m.IdleTimeout = types.StringNull()
	}
	if v, ok := obj["addresses-per-mac"]; ok && v != "" {
		m.AddressesPerMac = types.StringValue(v)
	} else {
		m.AddressesPerMac = types.StringNull()
	}
	if v, ok := obj["address-pool"]; ok && v != "" {
		m.AddressPool = types.StringValue(v)
	} else {
		m.AddressPool = types.StringNull()
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
	if v, ok := obj["interface"]; ok {
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	}
	if v, ok := obj["keepalive-timeout"]; ok {
		if v != "" {
			m.KeepaliveTimeout = types.StringValue(v)
		} else {
			m.KeepaliveTimeout = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["profile"]; ok {
		if v != "" {
			m.Profile = types.StringValue(v)
		} else {
			m.Profile = types.StringNull()
		}
	}
}
