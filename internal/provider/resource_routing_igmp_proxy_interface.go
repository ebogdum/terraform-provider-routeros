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
	_ resource.Resource                = &RoutingIgmpProxyInterfaceResource{}
	_ resource.ResourceWithImportState = &RoutingIgmpProxyInterfaceResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingIgmpProxyInterfaceResource struct {
	reg *client.Registry
}

type RoutingIgmpProxyInterfaceModel struct {
	ID                 types.String `tfsdk:"id"`
	AlternativeSubnets types.String `tfsdk:"alternative_subnets"`
	Comment            types.String `tfsdk:"comment"`
	Disabled           types.Bool   `tfsdk:"disabled"`
	Dynamic            types.Bool   `tfsdk:"dynamic"`
	Inactive           types.Bool   `tfsdk:"inactive"`
	Interface          types.String `tfsdk:"interface"`
	Querier            types.Bool   `tfsdk:"querier"`
	RxBytes            types.String `tfsdk:"rx_bytes"`
	RxPackets          types.String `tfsdk:"rx_packets"`
	SourceIPAddress    types.String `tfsdk:"source_ip_address"`
	Threshold          types.Int64  `tfsdk:"threshold"`
	TxBytes            types.String `tfsdk:"tx_bytes"`
	TxPackets          types.String `tfsdk:"tx_packets"`
	Upstream           types.Bool   `tfsdk:"upstream"`
	Router             types.String `tfsdk:"router"`
}

func NewRoutingIgmpProxyInterfaceResource() resource.Resource {
	return &RoutingIgmpProxyInterfaceResource{}
}

func (r *RoutingIgmpProxyInterfaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_igmp_proxy_interface"
}

func (r *RoutingIgmpProxyInterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingIgmpProxyInterfaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/routing/igmp-proxy/interface`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"alternative_subnets": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
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
			"inactive": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"querier": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_bytes": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"rx_packets": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"source_ip_address": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"threshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_bytes": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tx_packets": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"upstream": schema.BoolAttribute{
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

func (r *RoutingIgmpProxyInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingIgmpProxyInterfaceModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AlternativeSubnets.IsNull() || plan.AlternativeSubnets.IsUnknown()) {
		body["alternative-subnets"] = plan.AlternativeSubnets.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.Threshold.IsNull() || plan.Threshold.IsUnknown()) {
		body["threshold"] = client.FormatInt64(plan.Threshold.ValueInt64())
	}
	if !(plan.Upstream.IsNull() || plan.Upstream.IsUnknown()) {
		body["upstream"] = client.FormatBool(plan.Upstream.ValueBool())
	}
	obj, err := c.Add(ctx, "/routing/igmp-proxy/interface", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/igmp-proxy/interface failed", err.Error())
		return
	}
	routingIgmpProxyInterfaceApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingIgmpProxyInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingIgmpProxyInterfaceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/igmp-proxy/interface", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/igmp-proxy/interface failed", err.Error())
		return
	}
	routingIgmpProxyInterfaceApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingIgmpProxyInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingIgmpProxyInterfaceModel
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
	if !plan.AlternativeSubnets.Equal(state.AlternativeSubnets) && !plan.AlternativeSubnets.IsUnknown() {
		body["alternative-subnets"] = plan.AlternativeSubnets.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.Threshold.Equal(state.Threshold) && !plan.Threshold.IsUnknown() {
		body["threshold"] = client.FormatInt64(plan.Threshold.ValueInt64())
	}
	if !plan.Upstream.Equal(state.Upstream) && !plan.Upstream.IsUnknown() {
		body["upstream"] = client.FormatBool(plan.Upstream.ValueBool())
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/igmp-proxy/interface", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/igmp-proxy/interface failed", err.Error())
			return
		}
		routingIgmpProxyInterfaceApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingIgmpProxyInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingIgmpProxyInterfaceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/igmp-proxy/interface", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/igmp-proxy/interface failed", err.Error())
	}
}

func (r *RoutingIgmpProxyInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingIgmpProxyInterfaceLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/igmp-proxy/interface matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingIgmpProxyInterfaceLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingIgmpProxyInterfaceLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/igmp-proxy/interface", id)
}

func routingIgmpProxyInterfaceApply(ctx context.Context, obj client.Object, m *RoutingIgmpProxyInterfaceModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["alternative-subnets"]; ok {
		if v != "" {
			m.AlternativeSubnets = types.StringValue(v)
		} else {
			m.AlternativeSubnets = types.StringNull()
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
	if v, ok := obj["dynamic"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Dynamic = types.BoolValue(true)
		} else {
			m.Dynamic = types.BoolNull()
		}
	}
	if v, ok := obj["inactive"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Inactive = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Inactive = types.BoolValue(true)
		} else {
			m.Inactive = types.BoolNull()
		}
	}
	if v, ok := obj["interface"]; ok {
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	}
	if v, ok := obj["querier"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Querier = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Querier = types.BoolValue(true)
		} else {
			m.Querier = types.BoolNull()
		}
	}
	if v, ok := obj["rx-bytes"]; ok {
		if v != "" {
			m.RxBytes = types.StringValue(v)
		} else {
			m.RxBytes = types.StringNull()
		}
	}
	if v, ok := obj["rx-packets"]; ok {
		if v != "" {
			m.RxPackets = types.StringValue(v)
		} else {
			m.RxPackets = types.StringNull()
		}
	}
	if v, ok := obj["source-ip-address"]; ok {
		if v != "" {
			m.SourceIPAddress = types.StringValue(v)
		} else {
			m.SourceIPAddress = types.StringNull()
		}
	}
	if v, ok := obj["threshold"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Threshold = types.Int64Value(n)
		} else {
			m.Threshold = types.Int64Null()
		}
	} else {
		m.Threshold = types.Int64Null()
	}
	if v, ok := obj["tx-bytes"]; ok {
		if v != "" {
			m.TxBytes = types.StringValue(v)
		} else {
			m.TxBytes = types.StringNull()
		}
	}
	if v, ok := obj["tx-packets"]; ok {
		if v != "" {
			m.TxPackets = types.StringValue(v)
		} else {
			m.TxPackets = types.StringNull()
		}
	}
	if v, ok := obj["upstream"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Upstream = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Upstream = types.BoolValue(true)
		} else {
			m.Upstream = types.BoolNull()
		}
	}
}
