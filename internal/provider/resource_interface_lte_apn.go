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
	ID                   types.String `tfsdk:"id"`
	AddDefaultRoute      types.Bool   `tfsdk:"add_default_route"`
	Apn                  types.String `tfsdk:"apn"`
	Authentication       types.String `tfsdk:"authentication"`
	Comment              types.String `tfsdk:"comment"`
	Default              types.Bool   `tfsdk:"default"`
	DefaultRouteDistance types.Int64  `tfsdk:"default_route_distance"`
	IPType               types.String `tfsdk:"ip_type"`
	Name                 types.String `tfsdk:"name"`
	UseNetworkApn        types.Bool   `tfsdk:"use_network_apn"`
	UsePeerDNS           types.Bool   `tfsdk:"use_peer_dns"`
	Router               types.String `tfsdk:"router"`
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
	_ = fmt.Sprintf
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
				Optional:    true,
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
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
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
	obj, err := c.Add(ctx, "/interface/lte/apn", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/lte/apn failed", err.Error())
		return
	}
	interfaceLteApnApply(ctx, obj, &plan)
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
	if !plan.AddDefaultRoute.Equal(state.AddDefaultRoute) {
		body["add-default-route"] = client.FormatBool(plan.AddDefaultRoute.ValueBool())
	}
	if !plan.Apn.Equal(state.Apn) {
		body["apn"] = plan.Apn.ValueString()
	}
	if !plan.Authentication.Equal(state.Authentication) {
		body["authentication"] = plan.Authentication.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.DefaultRouteDistance.Equal(state.DefaultRouteDistance) {
		body["default-route-distance"] = client.FormatInt64(plan.DefaultRouteDistance.ValueInt64())
	}
	if !plan.IPType.Equal(state.IPType) {
		body["ip-type"] = plan.IPType.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.UseNetworkApn.Equal(state.UseNetworkApn) {
		body["use-network-apn"] = client.FormatBool(plan.UseNetworkApn.ValueBool())
	}
	if !plan.UsePeerDNS.Equal(state.UsePeerDNS) {
		body["use-peer-dns"] = client.FormatBool(plan.UsePeerDNS.ValueBool())
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
	if v, ok := obj["add-default-route"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AddDefaultRoute = types.BoolValue(b)
		} else {
			m.AddDefaultRoute = types.BoolNull()
		}
	} else {
		m.AddDefaultRoute = types.BoolNull()
	}
	if v, ok := obj["apn"]; ok {
		_ = v
		if v != "" {
			m.Apn = types.StringValue(v)
		} else {
			m.Apn = types.StringNull()
		}
	} else {
		m.Apn = types.StringNull()
	}
	if v, ok := obj["authentication"]; ok {
		_ = v
		if v != "" {
			m.Authentication = types.StringValue(v)
		} else {
			m.Authentication = types.StringNull()
		}
	} else {
		m.Authentication = types.StringNull()
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
		_ = v
		if v != "" {
			m.IPType = types.StringValue(v)
		} else {
			m.IPType = types.StringNull()
		}
	} else {
		m.IPType = types.StringNull()
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
	if v, ok := obj["use-network-apn"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.UseNetworkApn = types.BoolValue(b)
		} else {
			m.UseNetworkApn = types.BoolNull()
		}
	} else {
		m.UseNetworkApn = types.BoolNull()
	}
	if v, ok := obj["use-peer-dns"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.UsePeerDNS = types.BoolValue(b)
		} else {
			m.UsePeerDNS = types.BoolNull()
		}
	} else {
		m.UsePeerDNS = types.BoolNull()
	}
}
