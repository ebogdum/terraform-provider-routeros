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
	_ resource.Resource                = &IPRouteResource{}
	_ resource.ResourceWithImportState = &IPRouteResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPRouteResource struct {
	reg *client.Registry
}

type IPRouteModel struct {
	ID           types.String `tfsdk:"id"`
	PrefSrc      types.String `tfsdk:"pref_src"`
	CheckGateway types.String `tfsdk:"check_gateway"`
	Blackhole    types.String `tfsdk:"blackhole"`
	Active       types.Bool   `tfsdk:"active"`
	Comment      types.String `tfsdk:"comment"`
	Connect      types.Bool   `tfsdk:"connect"`
	DHCP         types.Bool   `tfsdk:"dhcp"`
	Disabled     types.Bool   `tfsdk:"disabled"`
	Distance     types.Int64  `tfsdk:"distance"`
	DstAddress   types.String `tfsdk:"dst_address"`
	Dynamic      types.Bool   `tfsdk:"dynamic"`
	Ecmp         types.Bool   `tfsdk:"ecmp"`
	Gateway      types.String `tfsdk:"gateway"`
	HwOffloaded  types.Bool   `tfsdk:"hw_offloaded"`
	ImmediateGw  types.String `tfsdk:"immediate_gw"`
	Inactive     types.Bool   `tfsdk:"inactive"`
	LocalAddress types.String `tfsdk:"local_address"`
	RoutingTable types.String `tfsdk:"routing_table"`
	Rtype        types.Int64  `tfsdk:"rtype"`
	Scope        types.Int64  `tfsdk:"scope"`
	TargetScope  types.Int64  `tfsdk:"target_scope"`
	Type         types.Int64  `tfsdk:"type"`
	VrfInterface types.String `tfsdk:"vrf_interface"`
	Router       types.String `tfsdk:"router"`
}

func NewIPRouteResource() resource.Resource { return &IPRouteResource{} }

func (r *IPRouteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_route"
}

func (r *IPRouteResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPRouteResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/route`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"pref_src": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `pref-src`.",
			},
			"check_gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `check-gateway`.",
			},
			"blackhole": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `blackhole`.",
			},
			"active": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"connect": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"dhcp": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"distance": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dst_address": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsCIDR()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeCIDR()},
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"ecmp": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"hw_offloaded": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"immediate_gw": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"inactive": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"local_address": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"routing_table": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rtype": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"scope": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"target_scope": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"type": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"vrf_interface": schema.StringAttribute{
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

func (r *IPRouteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPRouteModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Distance.IsNull() || plan.Distance.IsUnknown()) {
		body["distance"] = client.FormatInt64(plan.Distance.ValueInt64())
	}
	if !(plan.DstAddress.IsNull() || plan.DstAddress.IsUnknown()) {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !(plan.Gateway.IsNull() || plan.Gateway.IsUnknown()) {
		body["gateway"] = plan.Gateway.ValueString()
	}
	if !(plan.RoutingTable.IsNull() || plan.RoutingTable.IsUnknown()) {
		body["routing-table"] = plan.RoutingTable.ValueString()
	}
	if !(plan.Scope.IsNull() || plan.Scope.IsUnknown()) {
		body["scope"] = client.FormatInt64(plan.Scope.ValueInt64())
	}
	if !(plan.TargetScope.IsNull() || plan.TargetScope.IsUnknown()) {
		body["target-scope"] = client.FormatInt64(plan.TargetScope.ValueInt64())
	}
	if !(plan.VrfInterface.IsNull() || plan.VrfInterface.IsUnknown()) {
		body["vrf-interface"] = plan.VrfInterface.ValueString()
	}
	if !(plan.Blackhole.IsNull() || plan.Blackhole.IsUnknown()) {
		body["blackhole"] = plan.Blackhole.ValueString()
	}
	if !(plan.CheckGateway.IsNull() || plan.CheckGateway.IsUnknown()) {
		body["check-gateway"] = plan.CheckGateway.ValueString()
	}
	if !(plan.PrefSrc.IsNull() || plan.PrefSrc.IsUnknown()) {
		body["pref-src"] = plan.PrefSrc.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/route", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/route failed", err.Error())
		return
	}
	iPRouteApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPRouteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPRouteModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/route", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/route failed", err.Error())
		return
	}
	iPRouteApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPRouteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPRouteModel
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
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Distance.Equal(state.Distance) && !plan.Distance.IsUnknown() {
		body["distance"] = client.FormatInt64(plan.Distance.ValueInt64())
	}
	if !plan.DstAddress.Equal(state.DstAddress) && !plan.DstAddress.IsUnknown() {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !plan.Gateway.Equal(state.Gateway) && !plan.Gateway.IsUnknown() {
		body["gateway"] = plan.Gateway.ValueString()
	}
	if !plan.RoutingTable.Equal(state.RoutingTable) && !plan.RoutingTable.IsUnknown() {
		body["routing-table"] = plan.RoutingTable.ValueString()
	}
	if !plan.Scope.Equal(state.Scope) && !plan.Scope.IsUnknown() {
		body["scope"] = client.FormatInt64(plan.Scope.ValueInt64())
	}
	if !plan.TargetScope.Equal(state.TargetScope) && !plan.TargetScope.IsUnknown() {
		body["target-scope"] = client.FormatInt64(plan.TargetScope.ValueInt64())
	}
	if !plan.VrfInterface.Equal(state.VrfInterface) && !plan.VrfInterface.IsUnknown() {
		body["vrf-interface"] = plan.VrfInterface.ValueString()
	}
	if !plan.Blackhole.Equal(state.Blackhole) && !plan.Blackhole.IsUnknown() {
		body["blackhole"] = plan.Blackhole.ValueString()
	}
	if !plan.CheckGateway.Equal(state.CheckGateway) && !plan.CheckGateway.IsUnknown() {
		body["check-gateway"] = plan.CheckGateway.ValueString()
	}
	if !plan.PrefSrc.Equal(state.PrefSrc) && !plan.PrefSrc.IsUnknown() {
		body["pref-src"] = plan.PrefSrc.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/route", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/route failed", err.Error())
			return
		}
		iPRouteApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPRouteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPRouteModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/route", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/route failed", err.Error())
	}
}

func (r *IPRouteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPRouteLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/route matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPRouteLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPRouteLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/route", id)
}

func iPRouteApply(ctx context.Context, obj client.Object, m *IPRouteModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["pref-src"]; ok && v != "" {
		m.PrefSrc = types.StringValue(v)
	} else {
		m.PrefSrc = types.StringNull()
	}
	if v, ok := obj["check-gateway"]; ok && v != "" {
		m.CheckGateway = types.StringValue(v)
	} else {
		m.CheckGateway = types.StringNull()
	}
	if v, ok := obj["blackhole"]; ok {
		// RouterOS returns blackhole as a valueless presence flag: the key is
		// present with an empty value when the route is a blackhole route.
		if strings.TrimSpace(v) == "" {
			m.Blackhole = types.StringValue("true")
		} else {
			m.Blackhole = types.StringValue(v)
		}
	} else {
		m.Blackhole = types.StringValue("false")
	}
	if v, ok := obj["active"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Active = types.BoolValue(b)
		} else {
			m.Active = types.BoolNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["connect"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Connect = types.BoolValue(b)
		} else {
			m.Connect = types.BoolNull()
		}
	}
	if v, ok := obj["dhcp"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.DHCP = types.BoolValue(b)
		} else {
			m.DHCP = types.BoolNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["distance"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Distance = types.Int64Value(n)
		} else {
			m.Distance = types.Int64Null()
		}
	} else {
		m.Distance = types.Int64Null()
	}
	if v, ok := obj["dst-address"]; ok {
		if v != "" {
			m.DstAddress = types.StringValue(v)
		} else {
			m.DstAddress = types.StringNull()
		}
	}
	if v, ok := obj["dynamic"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else {
			m.Dynamic = types.BoolNull()
		}
	}
	if v, ok := obj["ecmp"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Ecmp = types.BoolValue(b)
		} else {
			m.Ecmp = types.BoolNull()
		}
	}
	if v, ok := obj["gateway"]; ok {
		if v != "" {
			m.Gateway = types.StringValue(v)
		} else {
			m.Gateway = types.StringNull()
		}
	}
	if v, ok := obj["hw-offloaded"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.HwOffloaded = types.BoolValue(b)
		} else {
			m.HwOffloaded = types.BoolNull()
		}
	}
	if v, ok := obj["immediate-gw"]; ok {
		if v != "" {
			m.ImmediateGw = types.StringValue(v)
		} else {
			m.ImmediateGw = types.StringNull()
		}
	}
	if v, ok := obj["inactive"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Inactive = types.BoolValue(b)
		} else {
			m.Inactive = types.BoolNull()
		}
	}
	if v, ok := obj["local-address"]; ok {
		if v != "" {
			m.LocalAddress = types.StringValue(v)
		} else {
			m.LocalAddress = types.StringNull()
		}
	}
	if v, ok := obj["routing-table"]; ok {
		if v != "" {
			m.RoutingTable = types.StringValue(v)
		} else {
			m.RoutingTable = types.StringNull()
		}
	}
	if v, ok := obj["rtype"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Rtype = types.Int64Value(n)
		} else {
			m.Rtype = types.Int64Null()
		}
	} else {
		m.Rtype = types.Int64Null()
	}
	if v, ok := obj["scope"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Scope = types.Int64Value(n)
		} else {
			m.Scope = types.Int64Null()
		}
	} else {
		m.Scope = types.Int64Null()
	}
	if v, ok := obj["target-scope"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TargetScope = types.Int64Value(n)
		} else {
			m.TargetScope = types.Int64Null()
		}
	} else {
		m.TargetScope = types.Int64Null()
	}
	if v, ok := obj["type"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Type = types.Int64Value(n)
		} else {
			m.Type = types.Int64Null()
		}
	} else {
		m.Type = types.Int64Null()
	}
	if v, ok := obj["vrf-interface"]; ok {
		if v != "" {
			m.VrfInterface = types.StringValue(v)
		} else {
			m.VrfInterface = types.StringNull()
		}
	}
}
