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
	_ resource.Resource                = &MPLSLdpNeighborResource{}
	_ resource.ResourceWithImportState = &MPLSLdpNeighborResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type MPLSLdpNeighborResource struct {
	reg *client.Registry
}

type MPLSLdpNeighborModel struct {
	ID           types.String `tfsdk:"id"`
	Transport    types.String `tfsdk:"transport"`
	SendTargeted types.String `tfsdk:"send_targeted"`
	Comment      types.String `tfsdk:"comment"`
	Disabled     types.Bool   `tfsdk:"disabled"`
	Router       types.String `tfsdk:"router"`
}

func NewMPLSLdpNeighborResource() resource.Resource { return &MPLSLdpNeighborResource{} }

func (r *MPLSLdpNeighborResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mpls_ldp_neighbor"
}

func (r *MPLSLdpNeighborResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *MPLSLdpNeighborResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Requires an existing LDP instance. Skipped.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"transport": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `transport`.",
			},
			"send_targeted": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `send-targeted`.",
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
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *MPLSLdpNeighborResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan MPLSLdpNeighborModel
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
	if !(plan.SendTargeted.IsNull() || plan.SendTargeted.IsUnknown()) {
		body["send-targeted"] = plan.SendTargeted.ValueString()
	}
	if !(plan.Transport.IsNull() || plan.Transport.IsUnknown()) {
		body["transport"] = plan.Transport.ValueString()
	}
	obj, err := c.Add(ctx, "/mpls/ldp/neighbor", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /mpls/ldp/neighbor failed", err.Error())
		return
	}
	mPLSLdpNeighborApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MPLSLdpNeighborResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MPLSLdpNeighborModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/mpls/ldp/neighbor", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /mpls/ldp/neighbor failed", err.Error())
		return
	}
	mPLSLdpNeighborApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *MPLSLdpNeighborResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state MPLSLdpNeighborModel
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
	if !plan.SendTargeted.Equal(state.SendTargeted) && !plan.SendTargeted.IsUnknown() {
		body["send-targeted"] = plan.SendTargeted.ValueString()
	}
	if !plan.Transport.Equal(state.Transport) && !plan.Transport.IsUnknown() {
		body["transport"] = plan.Transport.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/mpls/ldp/neighbor", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /mpls/ldp/neighbor failed", err.Error())
			return
		}
		mPLSLdpNeighborApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MPLSLdpNeighborResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state MPLSLdpNeighborModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/mpls/ldp/neighbor", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /mpls/ldp/neighbor failed", err.Error())
	}
}

func (r *MPLSLdpNeighborResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := mPLSLdpNeighborLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /mpls/ldp/neighbor matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// mPLSLdpNeighborLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func mPLSLdpNeighborLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/mpls/ldp/neighbor", id)
}

func mPLSLdpNeighborApply(ctx context.Context, obj client.Object, m *MPLSLdpNeighborModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["transport"]; ok && v != "" {
		m.Transport = types.StringValue(v)
	} else {
		m.Transport = types.StringNull()
	}
	if v, ok := obj["send-targeted"]; ok && v != "" {
		m.SendTargeted = types.StringValue(v)
	} else {
		m.SendTargeted = types.StringNull()
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
}
