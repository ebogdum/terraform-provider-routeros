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
	_ resource.Resource                = &InterfacePppoeClientResource{}
	_ resource.ResourceWithImportState = &InterfacePppoeClientResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfacePppoeClientResource struct {
	reg *client.Registry
}

type InterfacePppoeClientModel struct {
	ID                   types.String `tfsdk:"id"`
	HostUniq             types.String `tfsdk:"host_uniq"`
	AcName               types.String `tfsdk:"ac_name"`
	AddDefaultRoute      types.String `tfsdk:"add_default_route"`
	Allow                types.String `tfsdk:"allow"`
	Comment              types.String `tfsdk:"comment"`
	DefaultRouteDistance types.String `tfsdk:"default_route_distance"`
	DialOnDemand         types.String `tfsdk:"dial_on_demand"`
	Disabled             types.Bool   `tfsdk:"disabled"`
	Interface            types.String `tfsdk:"interface"`
	KeepaliveTimeout     types.String `tfsdk:"keepalive_timeout"`
	MaxMru               types.String `tfsdk:"max_mru"`
	MaxMTU               types.String `tfsdk:"max_mtu"`
	Mrru                 types.String `tfsdk:"mrru"`
	Name                 types.String `tfsdk:"name"`
	Password             types.String `tfsdk:"password"`
	Profile              types.String `tfsdk:"profile"`
	ServiceName          types.String `tfsdk:"service_name"`
	UsePeerDNS           types.String `tfsdk:"use_peer_dns"`
	User                 types.String `tfsdk:"user"`
	Router               types.String `tfsdk:"router"`
}

func NewInterfacePppoeClientResource() resource.Resource { return &InterfacePppoeClientResource{} }

func (r *InterfacePppoeClientResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_pppoe_client"
}

func (r *InterfacePppoeClientResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfacePppoeClientResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "PPPoE client needs at least one interface in 'interfaces'. Skipped.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"host_uniq": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `host-uniq`.",
			},
			"ac_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"add_default_route": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"allow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"default_route_distance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dial_on_demand": schema.StringAttribute{
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
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"keepalive_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"max_mru": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"max_mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mrru": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
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
			"service_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"use_peer_dns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"user": schema.StringAttribute{
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

func (r *InterfacePppoeClientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfacePppoeClientModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AcName.IsNull() || plan.AcName.IsUnknown()) {
		body["ac-name"] = plan.AcName.ValueString()
	}
	if !(plan.AddDefaultRoute.IsNull() || plan.AddDefaultRoute.IsUnknown()) {
		body["add-default-route"] = plan.AddDefaultRoute.ValueString()
	}
	if !(plan.Allow.IsNull() || plan.Allow.IsUnknown()) {
		body["allow"] = plan.Allow.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.DefaultRouteDistance.IsNull() || plan.DefaultRouteDistance.IsUnknown()) {
		body["default-route-distance"] = plan.DefaultRouteDistance.ValueString()
	}
	if !(plan.DialOnDemand.IsNull() || plan.DialOnDemand.IsUnknown()) {
		body["dial-on-demand"] = plan.DialOnDemand.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.KeepaliveTimeout.IsNull() || plan.KeepaliveTimeout.IsUnknown()) {
		body["keepalive-timeout"] = plan.KeepaliveTimeout.ValueString()
	}
	if !(plan.MaxMru.IsNull() || plan.MaxMru.IsUnknown()) {
		body["max-mru"] = plan.MaxMru.ValueString()
	}
	if !(plan.MaxMTU.IsNull() || plan.MaxMTU.IsUnknown()) {
		body["max-mtu"] = plan.MaxMTU.ValueString()
	}
	if !(plan.Mrru.IsNull() || plan.Mrru.IsUnknown()) {
		body["mrru"] = plan.Mrru.ValueString()
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
	if !(plan.ServiceName.IsNull() || plan.ServiceName.IsUnknown()) {
		body["service-name"] = plan.ServiceName.ValueString()
	}
	if !(plan.UsePeerDNS.IsNull() || plan.UsePeerDNS.IsUnknown()) {
		body["use-peer-dns"] = plan.UsePeerDNS.ValueString()
	}
	if !(plan.User.IsNull() || plan.User.IsUnknown()) {
		body["user"] = plan.User.ValueString()
	}
	if !(plan.HostUniq.IsNull() || plan.HostUniq.IsUnknown()) {
		body["host-uniq"] = plan.HostUniq.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/pppoe-client", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/pppoe-client failed", err.Error())
		return
	}
	interfacePppoeClientApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfacePppoeClientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfacePppoeClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/pppoe-client", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/pppoe-client failed", err.Error())
		return
	}
	interfacePppoeClientApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfacePppoeClientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfacePppoeClientModel
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
	if !plan.AcName.Equal(state.AcName) && !plan.AcName.IsUnknown() {
		body["ac-name"] = plan.AcName.ValueString()
	}
	if !plan.AddDefaultRoute.Equal(state.AddDefaultRoute) && !plan.AddDefaultRoute.IsUnknown() {
		body["add-default-route"] = plan.AddDefaultRoute.ValueString()
	}
	if !plan.Allow.Equal(state.Allow) && !plan.Allow.IsUnknown() {
		body["allow"] = plan.Allow.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.DefaultRouteDistance.Equal(state.DefaultRouteDistance) && !plan.DefaultRouteDistance.IsUnknown() {
		body["default-route-distance"] = plan.DefaultRouteDistance.ValueString()
	}
	if !plan.DialOnDemand.Equal(state.DialOnDemand) && !plan.DialOnDemand.IsUnknown() {
		body["dial-on-demand"] = plan.DialOnDemand.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.KeepaliveTimeout.Equal(state.KeepaliveTimeout) && !plan.KeepaliveTimeout.IsUnknown() {
		body["keepalive-timeout"] = plan.KeepaliveTimeout.ValueString()
	}
	if !plan.MaxMru.Equal(state.MaxMru) && !plan.MaxMru.IsUnknown() {
		body["max-mru"] = plan.MaxMru.ValueString()
	}
	if !plan.MaxMTU.Equal(state.MaxMTU) && !plan.MaxMTU.IsUnknown() {
		body["max-mtu"] = plan.MaxMTU.ValueString()
	}
	if !plan.Mrru.Equal(state.Mrru) && !plan.Mrru.IsUnknown() {
		body["mrru"] = plan.Mrru.ValueString()
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
	if !plan.ServiceName.Equal(state.ServiceName) && !plan.ServiceName.IsUnknown() {
		body["service-name"] = plan.ServiceName.ValueString()
	}
	if !plan.UsePeerDNS.Equal(state.UsePeerDNS) && !plan.UsePeerDNS.IsUnknown() {
		body["use-peer-dns"] = plan.UsePeerDNS.ValueString()
	}
	if !plan.User.Equal(state.User) && !plan.User.IsUnknown() {
		body["user"] = plan.User.ValueString()
	}
	if !plan.HostUniq.Equal(state.HostUniq) && !plan.HostUniq.IsUnknown() {
		body["host-uniq"] = plan.HostUniq.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/pppoe-client", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/pppoe-client failed", err.Error())
			return
		}
		interfacePppoeClientApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfacePppoeClientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfacePppoeClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/pppoe-client", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/pppoe-client failed", err.Error())
	}
}

func (r *InterfacePppoeClientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfacePppoeClientLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/pppoe-client matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfacePppoeClientLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfacePppoeClientLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/pppoe-client", id)
}

func interfacePppoeClientApply(ctx context.Context, obj client.Object, m *InterfacePppoeClientModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["host-uniq"]; ok && v != "" {
		m.HostUniq = types.StringValue(v)
	} else {
		m.HostUniq = types.StringNull()
	}
	if v, ok := obj["ac-name"]; ok {
		_ = v
		if v != "" {
			m.AcName = types.StringValue(v)
		} else {
			m.AcName = types.StringNull()
		}
	} else {
		m.AcName = types.StringNull()
	}
	if v, ok := obj["add-default-route"]; ok {
		_ = v
		if v != "" {
			m.AddDefaultRoute = types.StringValue(v)
		} else {
			m.AddDefaultRoute = types.StringNull()
		}
	} else {
		m.AddDefaultRoute = types.StringNull()
	}
	if v, ok := obj["allow"]; ok {
		_ = v
		if v != "" {
			m.Allow = types.StringValue(v)
		} else {
			m.Allow = types.StringNull()
		}
	} else {
		m.Allow = types.StringNull()
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
	if v, ok := obj["default-route-distance"]; ok {
		_ = v
		if v != "" {
			m.DefaultRouteDistance = types.StringValue(v)
		} else {
			m.DefaultRouteDistance = types.StringNull()
		}
	} else {
		m.DefaultRouteDistance = types.StringNull()
	}
	if v, ok := obj["dial-on-demand"]; ok {
		_ = v
		if v != "" {
			m.DialOnDemand = types.StringValue(v)
		} else {
			m.DialOnDemand = types.StringNull()
		}
	} else {
		m.DialOnDemand = types.StringNull()
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
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
	if v, ok := obj["max-mru"]; ok {
		_ = v
		if v != "" {
			m.MaxMru = types.StringValue(v)
		} else {
			m.MaxMru = types.StringNull()
		}
	} else {
		m.MaxMru = types.StringNull()
	}
	if v, ok := obj["max-mtu"]; ok {
		_ = v
		if v != "" {
			m.MaxMTU = types.StringValue(v)
		} else {
			m.MaxMTU = types.StringNull()
		}
	} else {
		m.MaxMTU = types.StringNull()
	}
	if v, ok := obj["mrru"]; ok {
		_ = v
		if v != "" {
			m.Mrru = types.StringValue(v)
		} else {
			m.Mrru = types.StringNull()
		}
	} else {
		m.Mrru = types.StringNull()
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
	if v, ok := obj["password"]; ok {
		_ = v
		if v != "" {
			m.Password = types.StringValue(v)
		} else {
			m.Password = types.StringNull()
		}
	} else {
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
	if v, ok := obj["service-name"]; ok {
		_ = v
		if v != "" {
			m.ServiceName = types.StringValue(v)
		} else {
			m.ServiceName = types.StringNull()
		}
	} else {
		m.ServiceName = types.StringNull()
	}
	if v, ok := obj["use-peer-dns"]; ok {
		_ = v
		if v != "" {
			m.UsePeerDNS = types.StringValue(v)
		} else {
			m.UsePeerDNS = types.StringNull()
		}
	} else {
		m.UsePeerDNS = types.StringNull()
	}
	if v, ok := obj["user"]; ok {
		_ = v
		if v != "" {
			m.User = types.StringValue(v)
		} else {
			m.User = types.StringNull()
		}
	} else {
		m.User = types.StringNull()
	}
}
