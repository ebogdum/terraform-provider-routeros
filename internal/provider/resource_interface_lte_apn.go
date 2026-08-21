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
	_ resource.Resource                = &InterfaceLteApnResource{}
	_ resource.ResourceWithImportState = &InterfaceLteApnResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceLteApnResource struct {
	reg *client.Registry
}

type InterfaceLteApnModel struct {
	ID                    types.String `tfsdk:"id"`
	User                  types.String `tfsdk:"user"`
	Password              types.String `tfsdk:"password"`
	PassthroughSubnetSize types.String `tfsdk:"passthrough_subnet_size"`
	PassthroughMac        macValue     `tfsdk:"passthrough_mac"`
	PassthroughInterface  types.String `tfsdk:"passthrough_interface"`
	Ipv6Interface         types.String `tfsdk:"ipv6_interface"`
	AddDefaultRoute       types.Bool   `tfsdk:"add_default_route"`
	Apn                   types.String `tfsdk:"apn"`
	Authentication        types.String `tfsdk:"authentication"`
	Comment               types.String `tfsdk:"comment"`
	Default               types.Bool   `tfsdk:"default"`
	DefaultRouteDistance  types.Int64  `tfsdk:"default_route_distance"`
	IPType                types.String `tfsdk:"ip_type"`
	Name                  types.String `tfsdk:"name"`
	UseNetworkApn         types.Bool   `tfsdk:"use_network_apn"`
	UsePeerDNS            types.Bool   `tfsdk:"use_peer_dns"`
	Router                types.String `tfsdk:"router"`
}

func NewInterfaceLteApnResource() resource.Resource { return &InterfaceLteApnResource{} }

func (r *InterfaceLteApnResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_lte_apn"
}

func (r *InterfaceLteApnResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceLteApnResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/lte/apn`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"user": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `user`.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "RouterOS `password`.",
			},
			"passthrough_subnet_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `passthrough-subnet-size`.",
			},
			"passthrough_mac": schema.StringAttribute{
				CustomType:  macType{},
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `passthrough-mac`.",
				Validators:  []validator.String{schemautil.IsMAC()},
			},
			"passthrough_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `passthrough-interface`.",
			},
			"ipv6_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ipv6-interface`.",
			},
			"add_default_route": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"apn": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"authentication": schema.StringAttribute{
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
			"default_route_distance": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ip_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"use_network_apn": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"use_peer_dns": schema.BoolAttribute{
				Optional:    true,
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

func (r *InterfaceLteApnResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceLteApnModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AddDefaultRoute.IsNull() || plan.AddDefaultRoute.IsUnknown()) {
		body["add-default-route"] = client.FormatBool(plan.AddDefaultRoute.ValueBool())
	}
	if !(plan.Apn.IsNull() || plan.Apn.IsUnknown()) {
		body["apn"] = plan.Apn.ValueString()
	}
	if !(plan.Authentication.IsNull() || plan.Authentication.IsUnknown()) {
		body["authentication"] = plan.Authentication.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.DefaultRouteDistance.IsNull() || plan.DefaultRouteDistance.IsUnknown()) {
		body["default-route-distance"] = client.FormatInt64(plan.DefaultRouteDistance.ValueInt64())
	}
	if !(plan.IPType.IsNull() || plan.IPType.IsUnknown()) {
		body["ip-type"] = plan.IPType.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.UseNetworkApn.IsNull() || plan.UseNetworkApn.IsUnknown()) {
		body["use-network-apn"] = client.FormatBool(plan.UseNetworkApn.ValueBool())
	}
	if !(plan.UsePeerDNS.IsNull() || plan.UsePeerDNS.IsUnknown()) {
		body["use-peer-dns"] = client.FormatBool(plan.UsePeerDNS.ValueBool())
	}
	if !(plan.Ipv6Interface.IsNull() || plan.Ipv6Interface.IsUnknown()) {
		body["ipv6-interface"] = plan.Ipv6Interface.ValueString()
	}
	if !(plan.PassthroughInterface.IsNull() || plan.PassthroughInterface.IsUnknown()) {
		body["passthrough-interface"] = plan.PassthroughInterface.ValueString()
	}
	if !(plan.PassthroughMac.IsNull() || plan.PassthroughMac.IsUnknown()) {
		body["passthrough-mac"] = plan.PassthroughMac.ValueString()
	}
	if !(plan.PassthroughSubnetSize.IsNull() || plan.PassthroughSubnetSize.IsUnknown()) {
		body["passthrough-subnet-size"] = plan.PassthroughSubnetSize.ValueString()
	}
	if !(plan.Password.IsNull() || plan.Password.IsUnknown()) {
		body["password"] = plan.Password.ValueString()
	}
	if !(plan.User.IsNull() || plan.User.IsUnknown()) {
		body["user"] = plan.User.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/lte/apn", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/lte/apn failed", err.Error())
		return
	}
	interfaceLteApnApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceLteApnResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceLteApnModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/lte/apn", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/lte/apn failed", err.Error())
		return
	}
	interfaceLteApnApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceLteApnResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceLteApnModel
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
	if !plan.AddDefaultRoute.Equal(state.AddDefaultRoute) && !plan.AddDefaultRoute.IsUnknown() {
		body["add-default-route"] = client.FormatBool(plan.AddDefaultRoute.ValueBool())
	}
	if !plan.Apn.Equal(state.Apn) && !plan.Apn.IsUnknown() {
		body["apn"] = plan.Apn.ValueString()
	}
	if !plan.Authentication.Equal(state.Authentication) && !plan.Authentication.IsUnknown() {
		body["authentication"] = plan.Authentication.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.DefaultRouteDistance.Equal(state.DefaultRouteDistance) && !plan.DefaultRouteDistance.IsUnknown() {
		body["default-route-distance"] = client.FormatInt64(plan.DefaultRouteDistance.ValueInt64())
	}
	if !plan.IPType.Equal(state.IPType) && !plan.IPType.IsUnknown() {
		body["ip-type"] = plan.IPType.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.UseNetworkApn.Equal(state.UseNetworkApn) && !plan.UseNetworkApn.IsUnknown() {
		body["use-network-apn"] = client.FormatBool(plan.UseNetworkApn.ValueBool())
	}
	if !plan.UsePeerDNS.Equal(state.UsePeerDNS) && !plan.UsePeerDNS.IsUnknown() {
		body["use-peer-dns"] = client.FormatBool(plan.UsePeerDNS.ValueBool())
	}
	if !plan.Ipv6Interface.Equal(state.Ipv6Interface) && !plan.Ipv6Interface.IsUnknown() {
		body["ipv6-interface"] = plan.Ipv6Interface.ValueString()
	}
	if !plan.PassthroughInterface.Equal(state.PassthroughInterface) && !plan.PassthroughInterface.IsUnknown() {
		body["passthrough-interface"] = plan.PassthroughInterface.ValueString()
	}
	if !plan.PassthroughMac.Equal(state.PassthroughMac) && !plan.PassthroughMac.IsUnknown() {
		body["passthrough-mac"] = plan.PassthroughMac.ValueString()
	}
	if !plan.PassthroughSubnetSize.Equal(state.PassthroughSubnetSize) && !plan.PassthroughSubnetSize.IsUnknown() {
		body["passthrough-subnet-size"] = plan.PassthroughSubnetSize.ValueString()
	}
	if !plan.Password.Equal(state.Password) && !plan.Password.IsUnknown() {
		body["password"] = plan.Password.ValueString()
	}
	if !plan.User.Equal(state.User) && !plan.User.IsUnknown() {
		body["user"] = plan.User.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/lte/apn", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/lte/apn failed", err.Error())
			return
		}
		interfaceLteApnApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceLteApnResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceLteApnModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/lte/apn", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/lte/apn failed", err.Error())
	}
}

func (r *InterfaceLteApnResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceLteApnLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/lte/apn matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceLteApnLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceLteApnLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/lte/apn", id)
}

func interfaceLteApnApply(ctx context.Context, obj client.Object, m *InterfaceLteApnModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["user"]; ok && v != "" {
		m.User = types.StringValue(v)
	} else {
		m.User = types.StringNull()
	}
	if v, ok := obj["password"]; ok && v != "" {
		m.Password = types.StringValue(v)
	} else {
		m.Password = types.StringNull()
	}
	if v, ok := obj["passthrough-subnet-size"]; ok && v != "" {
		m.PassthroughSubnetSize = types.StringValue(v)
	} else {
		m.PassthroughSubnetSize = types.StringNull()
	}
	if v, ok := obj["passthrough-mac"]; ok && v != "" {
		m.PassthroughMac = newMACValue(v)
	} else {
		m.PassthroughMac = newMACNull()
	}
	if v, ok := obj["passthrough-interface"]; ok && v != "" {
		m.PassthroughInterface = types.StringValue(v)
	} else {
		m.PassthroughInterface = types.StringNull()
	}
	if v, ok := obj["ipv6-interface"]; ok && v != "" {
		m.Ipv6Interface = types.StringValue(v)
	} else {
		m.Ipv6Interface = types.StringNull()
	}
	if v, ok := obj["add-default-route"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.AddDefaultRoute = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.AddDefaultRoute = types.BoolValue(true)
		} else {
			m.AddDefaultRoute = types.BoolNull()
		}
	}
	if v, ok := obj["apn"]; ok {
		if v != "" {
			m.Apn = types.StringValue(v)
		} else {
			m.Apn = types.StringNull()
		}
	}
	if v, ok := obj["authentication"]; ok {
		if v != "" {
			m.Authentication = types.StringValue(v)
		} else {
			m.Authentication = types.StringNull()
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
	if v, ok := obj["default-route-distance"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.DefaultRouteDistance = types.Int64Value(n)
		} else {
			m.DefaultRouteDistance = types.Int64Null()
		}
	} else {
		m.DefaultRouteDistance = types.Int64Null()
	}
	if v, ok := obj["ip-type"]; ok {
		if v != "" {
			m.IPType = types.StringValue(v)
		} else {
			m.IPType = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["use-network-apn"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.UseNetworkApn = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.UseNetworkApn = types.BoolValue(true)
		} else {
			m.UseNetworkApn = types.BoolNull()
		}
	}
	if v, ok := obj["use-peer-dns"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.UsePeerDNS = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.UsePeerDNS = types.BoolValue(true)
		} else {
			m.UsePeerDNS = types.BoolNull()
		}
	}
}
