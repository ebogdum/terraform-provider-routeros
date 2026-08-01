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
	_ resource.Resource                = &RoutingFantasyResource{}
	_ resource.ResourceWithImportState = &RoutingFantasyResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingFantasyResource struct {
	reg *client.Registry
}

type RoutingFantasyModel struct {
	ID           types.String `tfsdk:"id"`
	UseHold      types.String `tfsdk:"use_hold"`
	Seed         types.String `tfsdk:"seed"`
	PrivSize     types.String `tfsdk:"priv_size"`
	PrivOffs     types.String `tfsdk:"priv_offs"`
	PrefixLength types.String `tfsdk:"prefix_length"`
	Offset       types.String `tfsdk:"offset"`
	InstanceId   types.String `tfsdk:"instance_id"`
	DealerId     types.String `tfsdk:"dealer_id"`
	Count        types.String `tfsdk:"route_count"`
	Comment      types.String `tfsdk:"comment"`
	Disabled     types.Bool   `tfsdk:"disabled"`
	DstAddress   types.String `tfsdk:"dst_address"`
	Gateway      types.String `tfsdk:"gateway"`
	Name         types.String `tfsdk:"name"`
	Scope        types.String `tfsdk:"scope"`
	TargetScope  types.String `tfsdk:"target_scope"`
	Router       types.String `tfsdk:"router"`
}

func NewRoutingFantasyResource() resource.Resource { return &RoutingFantasyResource{} }

func (r *RoutingFantasyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_fantasy"
}

func (r *RoutingFantasyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingFantasyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/routing/fantasy`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"use_hold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `use-hold`.",
			},
			"seed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `seed`.",
			},
			"priv_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `priv-size`.",
			},
			"priv_offs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `priv-offs`.",
			},
			"prefix_length": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `prefix-length`.",
			},
			"offset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `offset`.",
			},
			"instance_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `instance-id`.",
			},
			"dealer_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dealer-id`.",
			},
			"route_count": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `count`.",
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
			"dst_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"gateway": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"scope": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"target_scope": schema.StringAttribute{
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

func (r *RoutingFantasyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingFantasyModel
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
	if !(plan.DstAddress.IsNull() || plan.DstAddress.IsUnknown()) {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !(plan.Gateway.IsNull() || plan.Gateway.IsUnknown()) {
		body["gateway"] = plan.Gateway.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Scope.IsNull() || plan.Scope.IsUnknown()) {
		body["scope"] = plan.Scope.ValueString()
	}
	if !(plan.TargetScope.IsNull() || plan.TargetScope.IsUnknown()) {
		body["target-scope"] = plan.TargetScope.ValueString()
	}
	if !(plan.Count.IsNull() || plan.Count.IsUnknown()) {
		body["count"] = plan.Count.ValueString()
	}
	if !(plan.DealerId.IsNull() || plan.DealerId.IsUnknown()) {
		body["dealer-id"] = plan.DealerId.ValueString()
	}
	if !(plan.InstanceId.IsNull() || plan.InstanceId.IsUnknown()) {
		body["instance-id"] = plan.InstanceId.ValueString()
	}
	if !(plan.Offset.IsNull() || plan.Offset.IsUnknown()) {
		body["offset"] = plan.Offset.ValueString()
	}
	if !(plan.PrefixLength.IsNull() || plan.PrefixLength.IsUnknown()) {
		body["prefix-length"] = plan.PrefixLength.ValueString()
	}
	if !(plan.PrivOffs.IsNull() || plan.PrivOffs.IsUnknown()) {
		body["priv-offs"] = plan.PrivOffs.ValueString()
	}
	if !(plan.PrivSize.IsNull() || plan.PrivSize.IsUnknown()) {
		body["priv-size"] = plan.PrivSize.ValueString()
	}
	if !(plan.Seed.IsNull() || plan.Seed.IsUnknown()) {
		body["seed"] = plan.Seed.ValueString()
	}
	if !(plan.UseHold.IsNull() || plan.UseHold.IsUnknown()) {
		body["use-hold"] = plan.UseHold.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/fantasy", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/fantasy failed", err.Error())
		return
	}
	routingFantasyApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingFantasyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingFantasyModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/fantasy", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/fantasy failed", err.Error())
		return
	}
	routingFantasyApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingFantasyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingFantasyModel
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
	if !plan.DstAddress.Equal(state.DstAddress) && !plan.DstAddress.IsUnknown() {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !plan.Gateway.Equal(state.Gateway) && !plan.Gateway.IsUnknown() {
		body["gateway"] = plan.Gateway.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Scope.Equal(state.Scope) && !plan.Scope.IsUnknown() {
		body["scope"] = plan.Scope.ValueString()
	}
	if !plan.TargetScope.Equal(state.TargetScope) && !plan.TargetScope.IsUnknown() {
		body["target-scope"] = plan.TargetScope.ValueString()
	}
	if !plan.Count.Equal(state.Count) && !plan.Count.IsUnknown() {
		body["count"] = plan.Count.ValueString()
	}
	if !plan.DealerId.Equal(state.DealerId) && !plan.DealerId.IsUnknown() {
		body["dealer-id"] = plan.DealerId.ValueString()
	}
	if !plan.InstanceId.Equal(state.InstanceId) && !plan.InstanceId.IsUnknown() {
		body["instance-id"] = plan.InstanceId.ValueString()
	}
	if !plan.Offset.Equal(state.Offset) && !plan.Offset.IsUnknown() {
		body["offset"] = plan.Offset.ValueString()
	}
	if !plan.PrefixLength.Equal(state.PrefixLength) && !plan.PrefixLength.IsUnknown() {
		body["prefix-length"] = plan.PrefixLength.ValueString()
	}
	if !plan.PrivOffs.Equal(state.PrivOffs) && !plan.PrivOffs.IsUnknown() {
		body["priv-offs"] = plan.PrivOffs.ValueString()
	}
	if !plan.PrivSize.Equal(state.PrivSize) && !plan.PrivSize.IsUnknown() {
		body["priv-size"] = plan.PrivSize.ValueString()
	}
	if !plan.Seed.Equal(state.Seed) && !plan.Seed.IsUnknown() {
		body["seed"] = plan.Seed.ValueString()
	}
	if !plan.UseHold.Equal(state.UseHold) && !plan.UseHold.IsUnknown() {
		body["use-hold"] = plan.UseHold.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/fantasy", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/fantasy failed", err.Error())
			return
		}
		routingFantasyApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingFantasyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingFantasyModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/fantasy", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/fantasy failed", err.Error())
	}
}

func (r *RoutingFantasyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingFantasyLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/fantasy matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingFantasyLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingFantasyLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/fantasy", id)
}

func routingFantasyApply(ctx context.Context, obj client.Object, m *RoutingFantasyModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["use-hold"]; ok && v != "" {
		m.UseHold = types.StringValue(v)
	} else {
		m.UseHold = types.StringNull()
	}
	if v, ok := obj["seed"]; ok && v != "" {
		m.Seed = types.StringValue(v)
	} else {
		m.Seed = types.StringNull()
	}
	if v, ok := obj["priv-size"]; ok && v != "" {
		m.PrivSize = types.StringValue(v)
	} else {
		m.PrivSize = types.StringNull()
	}
	if v, ok := obj["priv-offs"]; ok && v != "" {
		m.PrivOffs = types.StringValue(v)
	} else {
		m.PrivOffs = types.StringNull()
	}
	if v, ok := obj["prefix-length"]; ok && v != "" {
		m.PrefixLength = types.StringValue(v)
	} else {
		m.PrefixLength = types.StringNull()
	}
	if v, ok := obj["offset"]; ok && v != "" {
		m.Offset = types.StringValue(v)
	} else {
		m.Offset = types.StringNull()
	}
	if v, ok := obj["instance-id"]; ok && v != "" {
		m.InstanceId = types.StringValue(v)
	} else {
		m.InstanceId = types.StringNull()
	}
	if v, ok := obj["dealer-id"]; ok && v != "" {
		m.DealerId = types.StringValue(v)
	} else {
		m.DealerId = types.StringNull()
	}
	if v, ok := obj["count"]; ok && v != "" {
		m.Count = types.StringValue(v)
	} else {
		m.Count = types.StringNull()
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
	if v, ok := obj["scope"]; ok {
		_ = v
		if v != "" {
			m.Scope = types.StringValue(v)
		} else {
			m.Scope = types.StringNull()
		}
	} else {
		m.Scope = types.StringNull()
	}
	if v, ok := obj["target-scope"]; ok {
		_ = v
		if v != "" {
			m.TargetScope = types.StringValue(v)
		} else {
			m.TargetScope = types.StringNull()
		}
	} else {
		m.TargetScope = types.StringNull()
	}
}
