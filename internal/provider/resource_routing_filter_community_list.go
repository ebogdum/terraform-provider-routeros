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
	_ resource.Resource                = &RoutingFilterCommunityListResource{}
	_ resource.ResourceWithImportState = &RoutingFilterCommunityListResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingFilterCommunityListResource struct {
	reg *client.Registry
}

type RoutingFilterCommunityListModel struct {
	ID          types.String `tfsdk:"id"`
	Regexp      types.String `tfsdk:"regexp"`
	List        types.String `tfsdk:"list"`
	Communities types.String `tfsdk:"communities"`
	Comment     types.String `tfsdk:"comment"`
	Disabled    types.Bool   `tfsdk:"disabled"`
	Router      types.String `tfsdk:"router"`
}

func NewRoutingFilterCommunityListResource() resource.Resource {
	return &RoutingFilterCommunityListResource{}
}

func (r *RoutingFilterCommunityListResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_filter_community_list"
}

func (r *RoutingFilterCommunityListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingFilterCommunityListResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "7.x routing-filter community-list uses 'communities' field and rule chain semantics that vary across releases. Skipped.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"regexp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `regexp`.",
			},
			"list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `list`.",
			},
			"communities": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `communities`.",
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

func (r *RoutingFilterCommunityListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingFilterCommunityListModel
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
	if !(plan.Communities.IsNull() || plan.Communities.IsUnknown()) {
		body["communities"] = plan.Communities.ValueString()
	}
	if !(plan.List.IsNull() || plan.List.IsUnknown()) {
		body["list"] = plan.List.ValueString()
	}
	if !(plan.Regexp.IsNull() || plan.Regexp.IsUnknown()) {
		body["regexp"] = plan.Regexp.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/filter/community-list", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/filter/community-list failed", err.Error())
		return
	}
	routingFilterCommunityListApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingFilterCommunityListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingFilterCommunityListModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/filter/community-list", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/filter/community-list failed", err.Error())
		return
	}
	routingFilterCommunityListApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingFilterCommunityListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingFilterCommunityListModel
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
	if !plan.Communities.Equal(state.Communities) && !plan.Communities.IsUnknown() {
		body["communities"] = plan.Communities.ValueString()
	}
	if !plan.List.Equal(state.List) && !plan.List.IsUnknown() {
		body["list"] = plan.List.ValueString()
	}
	if !plan.Regexp.Equal(state.Regexp) && !plan.Regexp.IsUnknown() {
		body["regexp"] = plan.Regexp.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/filter/community-list", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/filter/community-list failed", err.Error())
			return
		}
		routingFilterCommunityListApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingFilterCommunityListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingFilterCommunityListModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/filter/community-list", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/filter/community-list failed", err.Error())
	}
}

func (r *RoutingFilterCommunityListResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingFilterCommunityListLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/filter/community-list matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingFilterCommunityListLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingFilterCommunityListLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/filter/community-list", id)
}

func routingFilterCommunityListApply(ctx context.Context, obj client.Object, m *RoutingFilterCommunityListModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["regexp"]; ok && v != "" {
		m.Regexp = types.StringValue(v)
	} else {
		m.Regexp = types.StringNull()
	}
	if v, ok := obj["list"]; ok && v != "" {
		m.List = types.StringValue(v)
	} else {
		m.List = types.StringNull()
	}
	if v, ok := obj["communities"]; ok && v != "" {
		m.Communities = types.StringValue(v)
	} else {
		m.Communities = types.StringNull()
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
}
