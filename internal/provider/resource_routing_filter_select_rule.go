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
	_ resource.Resource                = &RoutingFilterSelectRuleResource{}
	_ resource.ResourceWithImportState = &RoutingFilterSelectRuleResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingFilterSelectRuleResource struct {
	reg *client.Registry
}

type RoutingFilterSelectRuleModel struct {
	ID           types.String `tfsdk:"id"`
	DoWhere      types.String `tfsdk:"do_where"`
	DoTake       types.String `tfsdk:"do_take"`
	DoSelectPrfx types.String `tfsdk:"do_select_prfx"`
	DoSelectNum  types.String `tfsdk:"do_select_num"`
	DoJump       types.String `tfsdk:"do_jump"`
	DoGroupPrfx  types.String `tfsdk:"do_group_prfx"`
	DoGroupNum   types.String `tfsdk:"do_group_num"`
	Chain        types.String `tfsdk:"chain"`
	Comment      types.String `tfsdk:"comment"`
	Disabled     types.Bool   `tfsdk:"disabled"`
	Do           types.String `tfsdk:"do"`
	Invalid      types.Bool   `tfsdk:"invalid"`
	Type         types.String `tfsdk:"type"`
	Router       types.String `tfsdk:"router"`
}

func NewRoutingFilterSelectRuleResource() resource.Resource {
	return &RoutingFilterSelectRuleResource{}
}

func (r *RoutingFilterSelectRuleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_filter_select_rule"
}

func (r *RoutingFilterSelectRuleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingFilterSelectRuleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "References a /routing/filter rule. Skipped.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"do_where": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `do-where`.",
			},
			"do_take": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `do-take`.",
			},
			"do_select_prfx": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `do-select-prfx`.",
			},
			"do_select_num": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `do-select-num`.",
			},
			"do_jump": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `do-jump`.",
			},
			"do_group_prfx": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `do-group-prfx`.",
			},
			"do_group_num": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `do-group-num`.",
			},
			"chain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
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
			"do": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"where", "group-num", "group-prfx", "select-num", "select-prfx", "take", "jump"}...)},
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *RoutingFilterSelectRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingFilterSelectRuleModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Chain.IsNull() || plan.Chain.IsUnknown()) {
		body["chain"] = plan.Chain.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.DoGroupNum.IsNull() || plan.DoGroupNum.IsUnknown()) {
		body["do-group-num"] = plan.DoGroupNum.ValueString()
	}
	if !(plan.DoGroupPrfx.IsNull() || plan.DoGroupPrfx.IsUnknown()) {
		body["do-group-prfx"] = plan.DoGroupPrfx.ValueString()
	}
	if !(plan.DoJump.IsNull() || plan.DoJump.IsUnknown()) {
		body["do-jump"] = plan.DoJump.ValueString()
	}
	if !(plan.DoSelectNum.IsNull() || plan.DoSelectNum.IsUnknown()) {
		body["do-select-num"] = plan.DoSelectNum.ValueString()
	}
	if !(plan.DoSelectPrfx.IsNull() || plan.DoSelectPrfx.IsUnknown()) {
		body["do-select-prfx"] = plan.DoSelectPrfx.ValueString()
	}
	if !(plan.DoTake.IsNull() || plan.DoTake.IsUnknown()) {
		body["do-take"] = plan.DoTake.ValueString()
	}
	if !(plan.DoWhere.IsNull() || plan.DoWhere.IsUnknown()) {
		body["do-where"] = plan.DoWhere.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/filter/select-rule", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/filter/select-rule failed", err.Error())
		return
	}
	routingFilterSelectRuleApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingFilterSelectRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingFilterSelectRuleModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/filter/select-rule", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/filter/select-rule failed", err.Error())
		return
	}
	routingFilterSelectRuleApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingFilterSelectRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingFilterSelectRuleModel
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
	if !plan.Chain.Equal(state.Chain) && !plan.Chain.IsUnknown() {
		body["chain"] = plan.Chain.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DoGroupNum.Equal(state.DoGroupNum) && !plan.DoGroupNum.IsUnknown() {
		body["do-group-num"] = plan.DoGroupNum.ValueString()
	}
	if !plan.DoGroupPrfx.Equal(state.DoGroupPrfx) && !plan.DoGroupPrfx.IsUnknown() {
		body["do-group-prfx"] = plan.DoGroupPrfx.ValueString()
	}
	if !plan.DoJump.Equal(state.DoJump) && !plan.DoJump.IsUnknown() {
		body["do-jump"] = plan.DoJump.ValueString()
	}
	if !plan.DoSelectNum.Equal(state.DoSelectNum) && !plan.DoSelectNum.IsUnknown() {
		body["do-select-num"] = plan.DoSelectNum.ValueString()
	}
	if !plan.DoSelectPrfx.Equal(state.DoSelectPrfx) && !plan.DoSelectPrfx.IsUnknown() {
		body["do-select-prfx"] = plan.DoSelectPrfx.ValueString()
	}
	if !plan.DoTake.Equal(state.DoTake) && !plan.DoTake.IsUnknown() {
		body["do-take"] = plan.DoTake.ValueString()
	}
	if !plan.DoWhere.Equal(state.DoWhere) && !plan.DoWhere.IsUnknown() {
		body["do-where"] = plan.DoWhere.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/filter/select-rule", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/filter/select-rule failed", err.Error())
			return
		}
		routingFilterSelectRuleApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingFilterSelectRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingFilterSelectRuleModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/filter/select-rule", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/filter/select-rule failed", err.Error())
	}
}

func (r *RoutingFilterSelectRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingFilterSelectRuleLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/filter/select-rule matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingFilterSelectRuleLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingFilterSelectRuleLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/filter/select-rule", id)
}

func routingFilterSelectRuleApply(ctx context.Context, obj client.Object, m *RoutingFilterSelectRuleModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["do-where"]; ok && v != "" {
		m.DoWhere = types.StringValue(v)
	} else {
		m.DoWhere = types.StringNull()
	}
	if v, ok := obj["do-take"]; ok && v != "" {
		m.DoTake = types.StringValue(v)
	} else {
		m.DoTake = types.StringNull()
	}
	if v, ok := obj["do-select-prfx"]; ok && v != "" {
		m.DoSelectPrfx = types.StringValue(v)
	} else {
		m.DoSelectPrfx = types.StringNull()
	}
	if v, ok := obj["do-select-num"]; ok && v != "" {
		m.DoSelectNum = types.StringValue(v)
	} else {
		m.DoSelectNum = types.StringNull()
	}
	if v, ok := obj["do-jump"]; ok && v != "" {
		m.DoJump = types.StringValue(v)
	} else {
		m.DoJump = types.StringNull()
	}
	if v, ok := obj["do-group-prfx"]; ok && v != "" {
		m.DoGroupPrfx = types.StringValue(v)
	} else {
		m.DoGroupPrfx = types.StringNull()
	}
	if v, ok := obj["do-group-num"]; ok && v != "" {
		m.DoGroupNum = types.StringValue(v)
	} else {
		m.DoGroupNum = types.StringNull()
	}
	if v, ok := obj["chain"]; ok {
		if v != "" {
			m.Chain = types.StringValue(v)
		} else {
			m.Chain = types.StringNull()
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
	if v, ok := obj["do"]; ok {
		if v != "" {
			m.Do = types.StringValue(v)
		} else {
			m.Do = types.StringNull()
		}
	}
	if v, ok := obj["invalid"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Invalid = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Invalid = types.BoolValue(true)
		} else {
			m.Invalid = types.BoolNull()
		}
	}
	if v, ok := obj["type"]; ok {
		if v != "" {
			m.Type = types.StringValue(v)
		} else {
			m.Type = types.StringNull()
		}
	}
}
