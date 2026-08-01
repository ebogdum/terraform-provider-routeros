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
	_ resource.Resource                = &RoutingBGPInstanceResource{}
	_ resource.ResourceWithImportState = &RoutingBGPInstanceResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingBGPInstanceResource struct {
	reg *client.Registry
}

type RoutingBGPInstanceModel struct {
	ID              types.String `tfsdk:"id"`
	Multipath       types.String `tfsdk:"multipath"`
	IgnoreAsPathLen types.String `tfsdk:"ignore_as_path_len"`
	As              types.String `tfsdk:"as"`
	ClusterID       types.String `tfsdk:"cluster_id"`
	Disabled        types.Bool   `tfsdk:"disabled"`
	IgnoreAsPath    types.String `tfsdk:"ignore_as_path"`
	Invalid         types.Bool   `tfsdk:"invalid"`
	Name            types.String `tfsdk:"name"`
	RouterID        types.String `tfsdk:"router_id"`
	RoutingTable    types.String `tfsdk:"routing_table"`
	Vrf             types.String `tfsdk:"vrf"`
	Router          types.String `tfsdk:"router"`
}

func NewRoutingBGPInstanceResource() resource.Resource { return &RoutingBGPInstanceResource{} }

func (r *RoutingBGPInstanceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_bgp_instance"
}

func (r *RoutingBGPInstanceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingBGPInstanceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/routing/bgp/instance`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"multipath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `multipath`.",
			},
			"ignore_as_path_len": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ignore-as-path-len`.",
			},
			"as": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"cluster_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ignore_as_path": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"router_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"routing_table": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vrf": schema.StringAttribute{
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

func (r *RoutingBGPInstanceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingBGPInstanceModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.As.IsNull() || plan.As.IsUnknown()) {
		body["as"] = plan.As.ValueString()
	}
	if !(plan.ClusterID.IsNull() || plan.ClusterID.IsUnknown()) {
		body["cluster-id"] = plan.ClusterID.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.RouterID.IsNull() || plan.RouterID.IsUnknown()) {
		body["router-id"] = plan.RouterID.ValueString()
	}
	if !(plan.RoutingTable.IsNull() || plan.RoutingTable.IsUnknown()) {
		body["routing-table"] = plan.RoutingTable.ValueString()
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if !(plan.IgnoreAsPathLen.IsNull() || plan.IgnoreAsPathLen.IsUnknown()) {
		body["ignore-as-path-len"] = plan.IgnoreAsPathLen.ValueString()
	}
	if !(plan.Multipath.IsNull() || plan.Multipath.IsUnknown()) {
		body["multipath"] = plan.Multipath.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/bgp/instance", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/bgp/instance failed", err.Error())
		return
	}
	routingBGPInstanceApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingBGPInstanceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingBGPInstanceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/bgp/instance", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/bgp/instance failed", err.Error())
		return
	}
	routingBGPInstanceApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingBGPInstanceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingBGPInstanceModel
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
	if !plan.As.Equal(state.As) && !plan.As.IsUnknown() {
		body["as"] = plan.As.ValueString()
	}
	if !plan.ClusterID.Equal(state.ClusterID) && !plan.ClusterID.IsUnknown() {
		body["cluster-id"] = plan.ClusterID.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.RouterID.Equal(state.RouterID) && !plan.RouterID.IsUnknown() {
		body["router-id"] = plan.RouterID.ValueString()
	}
	if !plan.RoutingTable.Equal(state.RoutingTable) && !plan.RoutingTable.IsUnknown() {
		body["routing-table"] = plan.RoutingTable.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) && !plan.Vrf.IsUnknown() {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if !plan.IgnoreAsPathLen.Equal(state.IgnoreAsPathLen) && !plan.IgnoreAsPathLen.IsUnknown() {
		body["ignore-as-path-len"] = plan.IgnoreAsPathLen.ValueString()
	}
	if !plan.Multipath.Equal(state.Multipath) && !plan.Multipath.IsUnknown() {
		body["multipath"] = plan.Multipath.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/bgp/instance", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/bgp/instance failed", err.Error())
			return
		}
		routingBGPInstanceApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingBGPInstanceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingBGPInstanceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/bgp/instance", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/bgp/instance failed", err.Error())
	}
}

func (r *RoutingBGPInstanceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingBGPInstanceLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/bgp/instance matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingBGPInstanceLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingBGPInstanceLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/bgp/instance", id)
}

func routingBGPInstanceApply(ctx context.Context, obj client.Object, m *RoutingBGPInstanceModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["multipath"]; ok && v != "" {
		m.Multipath = types.StringValue(v)
	} else {
		m.Multipath = types.StringNull()
	}
	if v, ok := obj["ignore-as-path-len"]; ok && v != "" {
		m.IgnoreAsPathLen = types.StringValue(v)
	} else {
		m.IgnoreAsPathLen = types.StringNull()
	}
	if v, ok := obj["as"]; ok {
		if v != "" {
			m.As = types.StringValue(v)
		} else {
			m.As = types.StringNull()
		}
	}
	if v, ok := obj["cluster-id"]; ok {
		if v != "" {
			m.ClusterID = types.StringValue(v)
		} else {
			m.ClusterID = types.StringNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["ignore-as-path"]; ok {
		if v != "" {
			m.IgnoreAsPath = types.StringValue(v)
		} else {
			m.IgnoreAsPath = types.StringNull()
		}
	}
	if v, ok := obj["invalid"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Invalid = types.BoolValue(b)
		} else {
			m.Invalid = types.BoolNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["router-id"]; ok {
		if v != "" {
			m.RouterID = types.StringValue(v)
		} else {
			m.RouterID = types.StringNull()
		}
	}
	if v, ok := obj["routing-table"]; ok {
		if v != "" {
			m.RoutingTable = types.StringValue(v)
		} else {
			m.RoutingTable = types.StringNull()
		}
	}
	if v, ok := obj["vrf"]; ok {
		if v != "" {
			m.Vrf = types.StringValue(v)
		} else {
			m.Vrf = types.StringNull()
		}
	}
}
