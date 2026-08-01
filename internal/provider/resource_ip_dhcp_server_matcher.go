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
	_ resource.Resource                = &IPDHCPServerMatcherResource{}
	_ resource.ResourceWithImportState = &IPDHCPServerMatcherResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPDHCPServerMatcherResource struct {
	reg *client.Registry
}

type IPDHCPServerMatcherModel struct {
	ID           types.String `tfsdk:"id"`
	Value        types.String `tfsdk:"value"`
	Server       types.String `tfsdk:"server"`
	OptionSet    types.String `tfsdk:"option_set"`
	Name         types.String `tfsdk:"name"`
	MatchingType types.String `tfsdk:"matching_type"`
	Code         types.String `tfsdk:"code"`
	AddressPool  types.String `tfsdk:"address_pool"`
	Comment      types.String `tfsdk:"comment"`
	Disabled     types.Bool   `tfsdk:"disabled"`
	Router       types.String `tfsdk:"router"`
}

func NewIPDHCPServerMatcherResource() resource.Resource { return &IPDHCPServerMatcherResource{} }

func (r *IPDHCPServerMatcherResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_dhcp_server_matcher"
}

func (r *IPDHCPServerMatcherResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPDHCPServerMatcherResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "matching-type field has version-specific accepted values; needs hand-tuning per ROS.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"value": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `value`.",
			},
			"server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `server`.",
			},
			"option_set": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `option-set`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `name`.",
			},
			"matching_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `matching-type`.",
			},
			"code": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `code`.",
			},
			"address_pool": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `address-pool`.",
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

func (r *IPDHCPServerMatcherResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPDHCPServerMatcherModel
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
	if !(plan.AddressPool.IsNull() || plan.AddressPool.IsUnknown()) {
		body["address-pool"] = plan.AddressPool.ValueString()
	}
	if !(plan.Code.IsNull() || plan.Code.IsUnknown()) {
		body["code"] = plan.Code.ValueString()
	}
	if !(plan.MatchingType.IsNull() || plan.MatchingType.IsUnknown()) {
		body["matching-type"] = plan.MatchingType.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.OptionSet.IsNull() || plan.OptionSet.IsUnknown()) {
		body["option-set"] = plan.OptionSet.ValueString()
	}
	if !(plan.Server.IsNull() || plan.Server.IsUnknown()) {
		body["server"] = plan.Server.ValueString()
	}
	if !(plan.Value.IsNull() || plan.Value.IsUnknown()) {
		body["value"] = plan.Value.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/dhcp-server/matcher", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/dhcp-server/matcher failed", err.Error())
		return
	}
	iPDHCPServerMatcherApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPServerMatcherResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPDHCPServerMatcherModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/dhcp-server/matcher", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/dhcp-server/matcher failed", err.Error())
		return
	}
	iPDHCPServerMatcherApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPDHCPServerMatcherResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPDHCPServerMatcherModel
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
	if !plan.AddressPool.Equal(state.AddressPool) && !plan.AddressPool.IsUnknown() {
		body["address-pool"] = plan.AddressPool.ValueString()
	}
	if !plan.Code.Equal(state.Code) && !plan.Code.IsUnknown() {
		body["code"] = plan.Code.ValueString()
	}
	if !plan.MatchingType.Equal(state.MatchingType) && !plan.MatchingType.IsUnknown() {
		body["matching-type"] = plan.MatchingType.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.OptionSet.Equal(state.OptionSet) && !plan.OptionSet.IsUnknown() {
		body["option-set"] = plan.OptionSet.ValueString()
	}
	if !plan.Server.Equal(state.Server) && !plan.Server.IsUnknown() {
		body["server"] = plan.Server.ValueString()
	}
	if !plan.Value.Equal(state.Value) && !plan.Value.IsUnknown() {
		body["value"] = plan.Value.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/dhcp-server/matcher", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/dhcp-server/matcher failed", err.Error())
			return
		}
		iPDHCPServerMatcherApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPServerMatcherResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPDHCPServerMatcherModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/dhcp-server/matcher", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/dhcp-server/matcher failed", err.Error())
	}
}

func (r *IPDHCPServerMatcherResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPDHCPServerMatcherLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/dhcp-server/matcher matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPDHCPServerMatcherLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPDHCPServerMatcherLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/dhcp-server/matcher", id)
}

func iPDHCPServerMatcherApply(ctx context.Context, obj client.Object, m *IPDHCPServerMatcherModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["value"]; ok && v != "" {
		m.Value = types.StringValue(v)
	} else {
		m.Value = types.StringNull()
	}
	if v, ok := obj["server"]; ok && v != "" {
		m.Server = types.StringValue(v)
	} else {
		m.Server = types.StringNull()
	}
	if v, ok := obj["option-set"]; ok && v != "" {
		m.OptionSet = types.StringValue(v)
	} else {
		m.OptionSet = types.StringNull()
	}
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["matching-type"]; ok && v != "" {
		m.MatchingType = types.StringValue(v)
	} else {
		m.MatchingType = types.StringNull()
	}
	if v, ok := obj["code"]; ok && v != "" {
		m.Code = types.StringValue(v)
	} else {
		m.Code = types.StringNull()
	}
	if v, ok := obj["address-pool"]; ok && v != "" {
		m.AddressPool = types.StringValue(v)
	} else {
		m.AddressPool = types.StringNull()
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
		} else {
			m.Disabled = types.BoolNull()
		}
	}
}
