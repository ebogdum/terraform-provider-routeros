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
	_ resource.Resource                = &RoutingIDResource{}
	_ resource.ResourceWithImportState = &RoutingIDResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingIDResource struct {
	reg *client.Registry
}

type RoutingIDModel struct {
	ID              types.String `tfsdk:"id"`
	Comment         types.String `tfsdk:"comment"`
	Disabled        types.Bool   `tfsdk:"disabled"`
	Dynamic         types.Bool   `tfsdk:"dynamic"`
	DynamicID       types.String `tfsdk:"dynamic_id"`
	Inactive        types.Bool   `tfsdk:"inactive"`
	Invalid         types.Bool   `tfsdk:"invalid"`
	Name            types.String `tfsdk:"name"`
	SelectDynamicID types.String `tfsdk:"select_dynamic_id"`
	SelectFromVrf   types.String `tfsdk:"select_from_vrf"`
	Router          types.String `tfsdk:"router"`
}

func NewRoutingIDResource() resource.Resource { return &RoutingIDResource{} }

func (r *RoutingIDResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_id"
}

func (r *RoutingIDResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *RoutingIDResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/routing/id`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"dynamic": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dynamic_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"inactive": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"select_dynamic_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"only-static", "only-loopback", "only-vrf", "only-active", "any", "lowest"}...)},
			},
			"select_from_vrf": schema.StringAttribute{
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

func (r *RoutingIDResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingIDModel
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
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.SelectDynamicID.IsNull() || plan.SelectDynamicID.IsUnknown()) {
		body["select-dynamic-id"] = plan.SelectDynamicID.ValueString()
	}
	if !(plan.SelectFromVrf.IsNull() || plan.SelectFromVrf.IsUnknown()) {
		body["select-from-vrf"] = plan.SelectFromVrf.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/id", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/id failed", err.Error())
		return
	}
	routingIDApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingIDResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingIDModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/id", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/id failed", err.Error())
		return
	}
	routingIDApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingIDResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingIDModel
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
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.SelectDynamicID.Equal(state.SelectDynamicID) {
		body["select-dynamic-id"] = plan.SelectDynamicID.ValueString()
	}
	if !plan.SelectFromVrf.Equal(state.SelectFromVrf) {
		body["select-from-vrf"] = plan.SelectFromVrf.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/id", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/id failed", err.Error())
			return
		}
		routingIDApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingIDResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingIDModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/id", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/id failed", err.Error())
	}
}

func (r *RoutingIDResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingIDLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/id matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingIDLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingIDLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/routing/id", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func routingIDApply(ctx context.Context, obj client.Object, m *RoutingIDModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
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
	if v, ok := obj["dynamic-id"]; ok {
		_ = v
		if v != "" {
			m.DynamicID = types.StringValue(v)
		} else {
			m.DynamicID = types.StringNull()
		}
	} else {
		m.DynamicID = types.StringNull()
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
	if v, ok := obj["invalid"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Invalid = types.BoolValue(b)
		} else {
			m.Invalid = types.BoolNull()
		}
	} else {
		m.Invalid = types.BoolNull()
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
	if v, ok := obj["select-dynamic-id"]; ok {
		_ = v
		if v != "" {
			m.SelectDynamicID = types.StringValue(v)
		} else {
			m.SelectDynamicID = types.StringNull()
		}
	} else {
		m.SelectDynamicID = types.StringNull()
	}
	if v, ok := obj["select-from-vrf"]; ok {
		_ = v
		if v != "" {
			m.SelectFromVrf = types.StringValue(v)
		} else {
			m.SelectFromVrf = types.StringNull()
		}
	} else {
		m.SelectFromVrf = types.StringNull()
	}
}
