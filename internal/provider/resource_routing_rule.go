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
	_ resource.Resource                = &RoutingRuleResource{}
	_ resource.ResourceWithImportState = &RoutingRuleResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingRuleResource struct {
	reg *client.Registry
}

type RoutingRuleModel struct {
	ID          types.String `tfsdk:"id"`
	Table       types.String `tfsdk:"table"`
	Action      types.String `tfsdk:"action"`
	Chain       types.String `tfsdk:"chain"`
	Comment     types.String `tfsdk:"comment"`
	Disabled    types.Bool   `tfsdk:"disabled"`
	DstAddress  types.String `tfsdk:"dst_address"`
	Interface   types.String `tfsdk:"interface"`
	Invalid     types.Bool   `tfsdk:"invalid"`
	MinPrefix   types.String `tfsdk:"min_prefix"`
	Realm       types.String `tfsdk:"realm"`
	RoutingMark types.String `tfsdk:"routing_mark"`
	SrcAddress  types.String `tfsdk:"src_address"`
	Vrf         types.Bool   `tfsdk:"vrf"`
	Router      types.String `tfsdk:"router"`
}

func NewRoutingRuleResource() resource.Resource { return &RoutingRuleResource{} }

func (r *RoutingRuleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_rule"
}

func (r *RoutingRuleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingRuleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/routing/rule`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"table": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `table`.",
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"lookup", "unreachable", "drop", "lookup-only-in-table", "mangle"}...)},
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
			"dst_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"min_prefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"realm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"routing_mark": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vrf": schema.BoolAttribute{
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

func (r *RoutingRuleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingRuleModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Action.IsNull() || plan.Action.IsUnknown()) {
		body["action"] = plan.Action.ValueString()
	}
	if !(plan.Chain.IsNull() || plan.Chain.IsUnknown()) {
		body["chain"] = plan.Chain.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.DstAddress.IsNull() || plan.DstAddress.IsUnknown()) {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.MinPrefix.IsNull() || plan.MinPrefix.IsUnknown()) {
		body["min-prefix"] = plan.MinPrefix.ValueString()
	}
	if !(plan.Realm.IsNull() || plan.Realm.IsUnknown()) {
		body["realm"] = plan.Realm.ValueString()
	}
	if !(plan.RoutingMark.IsNull() || plan.RoutingMark.IsUnknown()) {
		body["routing-mark"] = plan.RoutingMark.ValueString()
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) {
		body["vrf"] = client.FormatBool(plan.Vrf.ValueBool())
	}
	if !(plan.Table.IsNull() || plan.Table.IsUnknown()) {
		body["table"] = plan.Table.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/rule", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/rule failed", err.Error())
		return
	}
	routingRuleApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingRuleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingRuleModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/rule", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/rule failed", err.Error())
		return
	}
	routingRuleApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingRuleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingRuleModel
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
	if !plan.Action.Equal(state.Action) && !plan.Action.IsUnknown() {
		body["action"] = plan.Action.ValueString()
	}
	if !plan.Chain.Equal(state.Chain) && !plan.Chain.IsUnknown() {
		body["chain"] = plan.Chain.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DstAddress.Equal(state.DstAddress) && !plan.DstAddress.IsUnknown() {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.MinPrefix.Equal(state.MinPrefix) && !plan.MinPrefix.IsUnknown() {
		body["min-prefix"] = plan.MinPrefix.ValueString()
	}
	if !plan.Realm.Equal(state.Realm) && !plan.Realm.IsUnknown() {
		body["realm"] = plan.Realm.ValueString()
	}
	if !plan.RoutingMark.Equal(state.RoutingMark) && !plan.RoutingMark.IsUnknown() {
		body["routing-mark"] = plan.RoutingMark.ValueString()
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) && !plan.SrcAddress.IsUnknown() {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) && !plan.Vrf.IsUnknown() {
		body["vrf"] = client.FormatBool(plan.Vrf.ValueBool())
	}
	if !plan.Table.Equal(state.Table) && !plan.Table.IsUnknown() {
		body["table"] = plan.Table.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/rule", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/rule failed", err.Error())
			return
		}
		routingRuleApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingRuleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingRuleModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/rule", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/rule failed", err.Error())
	}
}

func (r *RoutingRuleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingRuleLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/rule matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingRuleLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingRuleLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/rule", id)
}

func routingRuleApply(ctx context.Context, obj client.Object, m *RoutingRuleModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["table"]; ok && v != "" {
		m.Table = types.StringValue(v)
	} else {
		m.Table = types.StringNull()
	}
	if v, ok := obj["action"]; ok {
		if v != "" {
			m.Action = types.StringValue(v)
		} else {
			m.Action = types.StringNull()
		}
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
	if v, ok := obj["dst-address"]; ok {
		if v != "" {
			m.DstAddress = types.StringValue(v)
		} else {
			m.DstAddress = types.StringNull()
		}
	}
	if v, ok := obj["interface"]; ok {
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
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
	if v, ok := obj["min-prefix"]; ok {
		if v != "" {
			m.MinPrefix = types.StringValue(v)
		} else {
			m.MinPrefix = types.StringNull()
		}
	}
	if v, ok := obj["realm"]; ok {
		if v != "" {
			m.Realm = types.StringValue(v)
		} else {
			m.Realm = types.StringNull()
		}
	}
	if v, ok := obj["routing-mark"]; ok {
		if v != "" {
			m.RoutingMark = types.StringValue(v)
		} else {
			m.RoutingMark = types.StringNull()
		}
	}
	if v, ok := obj["src-address"]; ok {
		if v != "" {
			m.SrcAddress = types.StringValue(v)
		} else {
			m.SrcAddress = types.StringNull()
		}
	}
	if v, ok := obj["vrf"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Vrf = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Vrf = types.BoolValue(true)
		} else {
			m.Vrf = types.BoolNull()
		}
	}
}
