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
	_ resource.Resource                = &RoutingGmpResource{}
	_ resource.ResourceWithImportState = &RoutingGmpResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingGmpResource struct {
	reg *client.Registry
}

type RoutingGmpModel struct {
	ID         types.String `tfsdk:"id"`
	Groups     types.String `tfsdk:"groups"`
	Disabled   types.Bool   `tfsdk:"disabled"`
	Exclude    types.Bool   `tfsdk:"exclude"`
	Group      types.String `tfsdk:"group"`
	Interfaces types.String `tfsdk:"interfaces"`
	Sources    types.String `tfsdk:"sources"`
	Router     types.String `tfsdk:"router"`
}

func NewRoutingGmpResource() resource.Resource { return &RoutingGmpResource{} }

func (r *RoutingGmpResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_gmp"
}

func (r *RoutingGmpResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingGmpResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "GMP needs interface configuration. Skipped from automated acc tests.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"groups": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `groups`.",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"exclude": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"group": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"interfaces": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sources": schema.StringAttribute{
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

func (r *RoutingGmpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingGmpModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Exclude.IsNull() || plan.Exclude.IsUnknown()) {
		body["exclude"] = client.FormatBool(plan.Exclude.ValueBool())
	}
	if !(plan.Interfaces.IsNull() || plan.Interfaces.IsUnknown()) {
		body["interfaces"] = plan.Interfaces.ValueString()
	}
	if !(plan.Sources.IsNull() || plan.Sources.IsUnknown()) {
		body["sources"] = plan.Sources.ValueString()
	}
	if !(plan.Groups.IsNull() || plan.Groups.IsUnknown()) {
		body["groups"] = plan.Groups.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/gmp", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/gmp failed", err.Error())
		return
	}
	routingGmpApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingGmpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingGmpModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/gmp", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/gmp failed", err.Error())
		return
	}
	routingGmpApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingGmpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingGmpModel
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
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Exclude.Equal(state.Exclude) && !plan.Exclude.IsUnknown() {
		body["exclude"] = client.FormatBool(plan.Exclude.ValueBool())
	}
	if !plan.Interfaces.Equal(state.Interfaces) && !plan.Interfaces.IsUnknown() {
		body["interfaces"] = plan.Interfaces.ValueString()
	}
	if !plan.Sources.Equal(state.Sources) && !plan.Sources.IsUnknown() {
		body["sources"] = plan.Sources.ValueString()
	}
	if !plan.Groups.Equal(state.Groups) && !plan.Groups.IsUnknown() {
		body["groups"] = plan.Groups.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/gmp", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/gmp failed", err.Error())
			return
		}
		routingGmpApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingGmpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingGmpModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/gmp", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/gmp failed", err.Error())
	}
}

func (r *RoutingGmpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingGmpLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/gmp matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingGmpLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingGmpLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/gmp", id)
}

func routingGmpApply(ctx context.Context, obj client.Object, m *RoutingGmpModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["groups"]; ok && v != "" {
		m.Groups = types.StringValue(v)
	} else {
		m.Groups = types.StringNull()
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
	if v, ok := obj["exclude"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Exclude = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Exclude = types.BoolValue(true)
		} else {
			m.Exclude = types.BoolNull()
		}
	}
	if v, ok := obj["group"]; ok {
		if v != "" {
			m.Group = types.StringValue(v)
		} else {
			m.Group = types.StringNull()
		}
	}
	if v, ok := obj["interfaces"]; ok {
		if v != "" {
			m.Interfaces = types.StringValue(v)
		} else {
			m.Interfaces = types.StringNull()
		}
	}
	if v, ok := obj["sources"]; ok {
		if v != "" {
			m.Sources = types.StringValue(v)
		} else {
			m.Sources = types.StringNull()
		}
	}
}
