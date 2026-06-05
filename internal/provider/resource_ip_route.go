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
	_ = fmt.Sprintf
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
			"active": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"connect": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dhcp": schema.BoolAttribute{
				Optional:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ecmp": schema.BoolAttribute{
				Optional:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"immediate_gw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"inactive": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"local_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"routing_table": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rtype": schema.Int64Attribute{
				Optional:    true,
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
				Optional:    true,
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
	if !(plan.Ecmp.IsNull() || plan.Ecmp.IsUnknown()) {
		body["ecmp"] = client.FormatBool(plan.Ecmp.ValueBool())
	}
	if !(plan.Gateway.IsNull() || plan.Gateway.IsUnknown()) {
		body["gateway"] = plan.Gateway.ValueString()
	}
	if !(plan.HwOffloaded.IsNull() || plan.HwOffloaded.IsUnknown()) {
		body["hw-offloaded"] = client.FormatBool(plan.HwOffloaded.ValueBool())
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
	obj, err := c.Add(ctx, "/ip/route", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/route failed", err.Error())
		return
	}
	iPRouteApply(ctx, obj, &plan)
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
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Distance.Equal(state.Distance) {
		body["distance"] = client.FormatInt64(plan.Distance.ValueInt64())
	}
	if !plan.DstAddress.Equal(state.DstAddress) {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !plan.Ecmp.Equal(state.Ecmp) {
		body["ecmp"] = client.FormatBool(plan.Ecmp.ValueBool())
	}
	if !plan.Gateway.Equal(state.Gateway) {
		body["gateway"] = plan.Gateway.ValueString()
	}
	if !plan.HwOffloaded.Equal(state.HwOffloaded) {
		body["hw-offloaded"] = client.FormatBool(plan.HwOffloaded.ValueBool())
	}
	if !plan.RoutingTable.Equal(state.RoutingTable) {
		body["routing-table"] = plan.RoutingTable.ValueString()
	}
	if !plan.Scope.Equal(state.Scope) {
		body["scope"] = client.FormatInt64(plan.Scope.ValueInt64())
	}
	if !plan.TargetScope.Equal(state.TargetScope) {
		body["target-scope"] = client.FormatInt64(plan.TargetScope.ValueInt64())
	}
	if !plan.VrfInterface.Equal(state.VrfInterface) {
		body["vrf-interface"] = plan.VrfInterface.ValueString()
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
	id := req.ID
	routerName := ""
	if i := strings.Index(id, "/"); i > 0 && !strings.HasPrefix(id, "*") {
		routerName, id = id[:i], id[i+1:]
	}
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
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/ip/route", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func iPRouteApply(ctx context.Context, obj client.Object, m *IPRouteModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["active"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Active = types.BoolValue(b)
		} else {
			m.Active = types.BoolNull()
		}
	} else {
		m.Active = types.BoolNull()
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
	if v, ok := obj["connect"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Connect = types.BoolValue(b)
		} else {
			m.Connect = types.BoolNull()
		}
	} else {
		m.Connect = types.BoolNull()
	}
	if v, ok := obj["dhcp"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.DHCP = types.BoolValue(b)
		} else {
			m.DHCP = types.BoolNull()
		}
	} else {
		m.DHCP = types.BoolNull()
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
		_ = v
		if v != "" {
			m.DstAddress = types.StringValue(v)
		} else {
			m.DstAddress = types.StringNull()
		}
	} else {
		m.DstAddress = types.StringNull()
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
	if v, ok := obj["ecmp"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Ecmp = types.BoolValue(b)
		} else {
			m.Ecmp = types.BoolNull()
		}
	} else {
		m.Ecmp = types.BoolNull()
	}
	if v, ok := obj["gateway"]; ok {
		_ = v
		if v != "" {
			m.Gateway = types.StringValue(v)
		} else {
			m.Gateway = types.StringNull()
		}
	} else {
		m.Gateway = types.StringNull()
	}
	if v, ok := obj["hw-offloaded"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.HwOffloaded = types.BoolValue(b)
		} else {
			m.HwOffloaded = types.BoolNull()
		}
	} else {
		m.HwOffloaded = types.BoolNull()
	}
	if v, ok := obj["immediate-gw"]; ok {
		_ = v
		if v != "" {
			m.ImmediateGw = types.StringValue(v)
		} else {
			m.ImmediateGw = types.StringNull()
		}
	} else {
		m.ImmediateGw = types.StringNull()
	}
	if v, ok := obj["inactive"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Inactive = types.BoolValue(b)
		} else {
			m.Inactive = types.BoolNull()
		}
	} else {
		m.Inactive = types.BoolNull()
	}
	if v, ok := obj["local-address"]; ok {
		_ = v
		if v != "" {
			m.LocalAddress = types.StringValue(v)
		} else {
			m.LocalAddress = types.StringNull()
		}
	} else {
		m.LocalAddress = types.StringNull()
	}
	if v, ok := obj["routing-table"]; ok {
		_ = v
		if v != "" {
			m.RoutingTable = types.StringValue(v)
		} else {
			m.RoutingTable = types.StringNull()
		}
	} else {
		m.RoutingTable = types.StringNull()
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
		_ = v
		if v != "" {
			m.VrfInterface = types.StringValue(v)
		} else {
			m.VrfInterface = types.StringNull()
		}
	} else {
		m.VrfInterface = types.StringNull()
	}
}
