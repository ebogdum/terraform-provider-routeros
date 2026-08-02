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
	_ resource.Resource                = &InterfaceBridgeVLANResource{}
	_ resource.ResourceWithImportState = &InterfaceBridgeVLANResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceBridgeVLANResource struct {
	reg *client.Registry
}

type InterfaceBridgeVLANModel struct {
	ID              types.String `tfsdk:"id"`
	Bridge          types.String `tfsdk:"bridge"`
	Comment         types.String `tfsdk:"comment"`
	CurrentTagged   types.String `tfsdk:"current_tagged"`
	CurrentUntagged types.String `tfsdk:"current_untagged"`
	Disabled        types.Bool   `tfsdk:"disabled"`
	Dynamic         types.Bool   `tfsdk:"dynamic"`
	MvrpAttributes  types.String `tfsdk:"mvrp_attributes"`
	MvrpForbidden   types.String `tfsdk:"mvrp_forbidden"`
	Tagged          types.String `tfsdk:"tagged"`
	Untagged        types.String `tfsdk:"untagged"`
	VLANIds         types.String `tfsdk:"vlan_ids"`
	Router          types.String `tfsdk:"router"`
}

func NewInterfaceBridgeVLANResource() resource.Resource { return &InterfaceBridgeVLANResource{} }

func (r *InterfaceBridgeVLANResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_bridge_vlan"
}

func (r *InterfaceBridgeVLANResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceBridgeVLANResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Auto-test requires a typed-reference precondition (e.g. an existing peer,\ninstance, bridge of the specific kind). The current acc-test generator's\ngeneric data.routeros_interface.all lookup can't satisfy these. Use this\nresource manually with explicit references to a precondition resource\nin your config.\n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"bridge": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"current_tagged": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"current_untagged": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"mvrp_attributes": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"mvrp_forbidden": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tagged": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"untagged": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vlan_ids": schema.StringAttribute{
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

func (r *InterfaceBridgeVLANResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceBridgeVLANModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Bridge.IsNull() || plan.Bridge.IsUnknown()) {
		body["bridge"] = plan.Bridge.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.MvrpForbidden.IsNull() || plan.MvrpForbidden.IsUnknown()) {
		body["mvrp-forbidden"] = plan.MvrpForbidden.ValueString()
	}
	if !(plan.Tagged.IsNull() || plan.Tagged.IsUnknown()) {
		body["tagged"] = plan.Tagged.ValueString()
	}
	if !(plan.Untagged.IsNull() || plan.Untagged.IsUnknown()) {
		body["untagged"] = plan.Untagged.ValueString()
	}
	if !(plan.VLANIds.IsNull() || plan.VLANIds.IsUnknown()) {
		body["vlan-ids"] = plan.VLANIds.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/bridge/vlan", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/bridge/vlan failed", err.Error())
		return
	}
	interfaceBridgeVLANApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceBridgeVLANResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceBridgeVLANModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/bridge/vlan", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/bridge/vlan failed", err.Error())
		return
	}
	interfaceBridgeVLANApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceBridgeVLANResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceBridgeVLANModel
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
	if !plan.Bridge.Equal(state.Bridge) && !plan.Bridge.IsUnknown() {
		body["bridge"] = plan.Bridge.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.MvrpForbidden.Equal(state.MvrpForbidden) && !plan.MvrpForbidden.IsUnknown() {
		body["mvrp-forbidden"] = plan.MvrpForbidden.ValueString()
	}
	if !plan.Tagged.Equal(state.Tagged) && !plan.Tagged.IsUnknown() {
		body["tagged"] = plan.Tagged.ValueString()
	}
	if !plan.Untagged.Equal(state.Untagged) && !plan.Untagged.IsUnknown() {
		body["untagged"] = plan.Untagged.ValueString()
	}
	if !plan.VLANIds.Equal(state.VLANIds) && !plan.VLANIds.IsUnknown() {
		body["vlan-ids"] = plan.VLANIds.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/bridge/vlan", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/bridge/vlan failed", err.Error())
			return
		}
		interfaceBridgeVLANApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceBridgeVLANResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceBridgeVLANModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/bridge/vlan", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/bridge/vlan failed", err.Error())
	}
}

func (r *InterfaceBridgeVLANResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceBridgeVLANLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/bridge/vlan matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceBridgeVLANLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceBridgeVLANLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/bridge/vlan", id)
}

func interfaceBridgeVLANApply(ctx context.Context, obj client.Object, m *InterfaceBridgeVLANModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["bridge"]; ok {
		if v != "" {
			m.Bridge = types.StringValue(v)
		} else {
			m.Bridge = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["current-tagged"]; ok {
		if v != "" {
			m.CurrentTagged = types.StringValue(v)
		} else {
			m.CurrentTagged = types.StringNull()
		}
	}
	if v, ok := obj["current-untagged"]; ok {
		if v != "" {
			m.CurrentUntagged = types.StringValue(v)
		} else {
			m.CurrentUntagged = types.StringNull()
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
	if v, ok := obj["dynamic"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Dynamic = types.BoolValue(true)
		} else {
			m.Dynamic = types.BoolNull()
		}
	}
	if v, ok := obj["mvrp-attributes"]; ok {
		if v != "" {
			m.MvrpAttributes = types.StringValue(v)
		} else {
			m.MvrpAttributes = types.StringNull()
		}
	}
	if v, ok := obj["mvrp-forbidden"]; ok {
		if v != "" {
			m.MvrpForbidden = types.StringValue(v)
		} else {
			m.MvrpForbidden = types.StringNull()
		}
	}
	if v, ok := obj["tagged"]; ok {
		if v != "" {
			m.Tagged = types.StringValue(v)
		} else {
			m.Tagged = types.StringNull()
		}
	}
	if v, ok := obj["untagged"]; ok {
		if v != "" {
			m.Untagged = types.StringValue(v)
		} else {
			m.Untagged = types.StringNull()
		}
	}
	if v, ok := obj["vlan-ids"]; ok {
		if v != "" {
			m.VLANIds = types.StringValue(v)
		} else {
			m.VLANIds = types.StringNull()
		}
	}
}
