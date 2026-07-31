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
	_ resource.Resource                = &MPLSLdpInterfaceResource{}
	_ resource.ResourceWithImportState = &MPLSLdpInterfaceResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type MPLSLdpInterfaceResource struct {
	reg *client.Registry
}

type MPLSLdpInterfaceModel struct {
	ID                     types.String `tfsdk:"id"`
	TransportAddresses     types.String `tfsdk:"transport_addresses"`
	HoldTime               types.String `tfsdk:"hold_time"`
	HelloInterval          types.String `tfsdk:"hello_interval"`
	Afi                    types.String `tfsdk:"afi"`
	AcceptDynamicNeighbors types.String `tfsdk:"accept_dynamic_neighbors"`
	Comment                types.String `tfsdk:"comment"`
	Disabled               types.Bool   `tfsdk:"disabled"`
	Interface              types.String `tfsdk:"interface"`
	Router                 types.String `tfsdk:"router"`
}

func NewMPLSLdpInterfaceResource() resource.Resource { return &MPLSLdpInterfaceResource{} }

func (r *MPLSLdpInterfaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mpls_ldp_interface"
}

func (r *MPLSLdpInterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *MPLSLdpInterfaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/mpls/ldp/interface`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"transport_addresses": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `transport-addresses`.",
			},
			"hold_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hold-time`.",
			},
			"hello_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hello-interval`.",
			},
			"afi": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `afi`.",
			},
			"accept_dynamic_neighbors": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `accept-dynamic-neighbors`.",
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
			"interface": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *MPLSLdpInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan MPLSLdpInterfaceModel
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
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.AcceptDynamicNeighbors.IsNull() || plan.AcceptDynamicNeighbors.IsUnknown()) {
		body["accept-dynamic-neighbors"] = plan.AcceptDynamicNeighbors.ValueString()
	}
	if !(plan.Afi.IsNull() || plan.Afi.IsUnknown()) {
		body["afi"] = plan.Afi.ValueString()
	}
	if !(plan.HelloInterval.IsNull() || plan.HelloInterval.IsUnknown()) {
		body["hello-interval"] = plan.HelloInterval.ValueString()
	}
	if !(plan.HoldTime.IsNull() || plan.HoldTime.IsUnknown()) {
		body["hold-time"] = plan.HoldTime.ValueString()
	}
	if !(plan.TransportAddresses.IsNull() || plan.TransportAddresses.IsUnknown()) {
		body["transport-addresses"] = plan.TransportAddresses.ValueString()
	}
	obj, err := c.Add(ctx, "/mpls/ldp/interface", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /mpls/ldp/interface failed", err.Error())
		return
	}
	mPLSLdpInterfaceApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MPLSLdpInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MPLSLdpInterfaceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/mpls/ldp/interface", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /mpls/ldp/interface failed", err.Error())
		return
	}
	mPLSLdpInterfaceApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *MPLSLdpInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state MPLSLdpInterfaceModel
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
	if !plan.Interface.Equal(state.Interface) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.AcceptDynamicNeighbors.Equal(state.AcceptDynamicNeighbors) && !plan.AcceptDynamicNeighbors.IsUnknown() {
		body["accept-dynamic-neighbors"] = plan.AcceptDynamicNeighbors.ValueString()
	}
	if !plan.Afi.Equal(state.Afi) && !plan.Afi.IsUnknown() {
		body["afi"] = plan.Afi.ValueString()
	}
	if !plan.HelloInterval.Equal(state.HelloInterval) && !plan.HelloInterval.IsUnknown() {
		body["hello-interval"] = plan.HelloInterval.ValueString()
	}
	if !plan.HoldTime.Equal(state.HoldTime) && !plan.HoldTime.IsUnknown() {
		body["hold-time"] = plan.HoldTime.ValueString()
	}
	if !plan.TransportAddresses.Equal(state.TransportAddresses) && !plan.TransportAddresses.IsUnknown() {
		body["transport-addresses"] = plan.TransportAddresses.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/mpls/ldp/interface", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /mpls/ldp/interface failed", err.Error())
			return
		}
		mPLSLdpInterfaceApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MPLSLdpInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state MPLSLdpInterfaceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/mpls/ldp/interface", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /mpls/ldp/interface failed", err.Error())
	}
}

func (r *MPLSLdpInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := mPLSLdpInterfaceLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /mpls/ldp/interface matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// mPLSLdpInterfaceLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func mPLSLdpInterfaceLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/mpls/ldp/interface", id)
}

func mPLSLdpInterfaceApply(ctx context.Context, obj client.Object, m *MPLSLdpInterfaceModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["transport-addresses"]; ok && v != "" {
		m.TransportAddresses = types.StringValue(v)
	} else {
		m.TransportAddresses = types.StringNull()
	}
	if v, ok := obj["hold-time"]; ok && v != "" {
		m.HoldTime = types.StringValue(v)
	} else {
		m.HoldTime = types.StringNull()
	}
	if v, ok := obj["hello-interval"]; ok && v != "" {
		m.HelloInterval = types.StringValue(v)
	} else {
		m.HelloInterval = types.StringNull()
	}
	if v, ok := obj["afi"]; ok && v != "" {
		m.Afi = types.StringValue(v)
	} else {
		m.Afi = types.StringNull()
	}
	if v, ok := obj["accept-dynamic-neighbors"]; ok && v != "" {
		m.AcceptDynamicNeighbors = types.StringValue(v)
	} else {
		m.AcceptDynamicNeighbors = types.StringNull()
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
}
