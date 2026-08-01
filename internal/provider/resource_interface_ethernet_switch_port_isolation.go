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
	_ resource.Resource                = &InterfaceEthernetSwitchPortIsolationResource{}
	_ resource.ResourceWithImportState = &InterfaceEthernetSwitchPortIsolationResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceEthernetSwitchPortIsolationResource struct {
	reg *client.Registry
}

type InterfaceEthernetSwitchPortIsolationModel struct {
	ID                 types.String `tfsdk:"id"`
	ForwardTo          types.String `tfsdk:"forward_to"`
	ForwardingOverride types.Bool   `tfsdk:"forwarding_override"`
	Invalid            types.Bool   `tfsdk:"invalid"`
	Name               types.String `tfsdk:"name"`
	Override           types.String `tfsdk:"override"`
	Switch             types.String `tfsdk:"switch"`
	Router             types.String `tfsdk:"router"`
}

func NewInterfaceEthernetSwitchPortIsolationResource() resource.Resource {
	return &InterfaceEthernetSwitchPortIsolationResource{}
}

func (r *InterfaceEthernetSwitchPortIsolationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ethernet_switch_port_isolation"
}

func (r *InterfaceEthernetSwitchPortIsolationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceEthernetSwitchPortIsolationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/ethernet/switch/port-isolation`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"forward_to": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"forwarding_override": schema.BoolAttribute{
				Optional:    true,
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
				Description: "Port name of the fixed row to adopt (e.g. `ether1`).",
			},
			"override": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"switch": schema.StringAttribute{
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

func (r *InterfaceEthernetSwitchPortIsolationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceEthernetSwitchPortIsolationModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	// Fixed device menu: the rows (one per switch port) already exist and
	// cannot be added. Adopt the row whose name matches and configure it.
	want := plan.Name.ValueString()
	if want == "" {
		resp.Diagnostics.AddError("name is required",
			"/interface/ethernet/switch/port-isolation has fixed rows; set `name` to the port to manage (e.g. ether1)")
		return
	}
	rows, err := c.List(ctx, "/interface/ethernet/switch/port-isolation")
	if err != nil {
		resp.Diagnostics.AddError("List /interface/ethernet/switch/port-isolation failed", err.Error())
		return
	}
	var id string
	for _, row := range rows {
		if row["name"] == want {
			id = row[".id"]
			break
		}
	}
	if id == "" {
		resp.Diagnostics.AddError("Unknown port-isolation row "+want,
			fmt.Sprintf("no /interface/ethernet/switch/port-isolation row named %q on the device", want))
		return
	}
	body := client.Object{}
	if !(plan.ForwardTo.IsNull() || plan.ForwardTo.IsUnknown()) {
		body["forward-to"] = plan.ForwardTo.ValueString()
	}
	if !(plan.ForwardingOverride.IsNull() || plan.ForwardingOverride.IsUnknown()) {
		body["forwarding-override"] = client.FormatBool(plan.ForwardingOverride.ValueBool())
	}
	if !(plan.Override.IsNull() || plan.Override.IsUnknown()) {
		body["override"] = plan.Override.ValueString()
	}
	obj, err := c.Set(ctx, "/interface/ethernet/switch/port-isolation", id, body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/ethernet/switch/port-isolation failed", err.Error())
		return
	}
	interfaceEthernetSwitchPortIsolationApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEthernetSwitchPortIsolationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceEthernetSwitchPortIsolationModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/ethernet/switch/port-isolation", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/ethernet/switch/port-isolation failed", err.Error())
		return
	}
	interfaceEthernetSwitchPortIsolationApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceEthernetSwitchPortIsolationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceEthernetSwitchPortIsolationModel
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
	if !plan.ForwardTo.Equal(state.ForwardTo) && !plan.ForwardTo.IsUnknown() {
		body["forward-to"] = plan.ForwardTo.ValueString()
	}
	if !plan.ForwardingOverride.Equal(state.ForwardingOverride) && !plan.ForwardingOverride.IsUnknown() {
		body["forwarding-override"] = client.FormatBool(plan.ForwardingOverride.ValueBool())
	}
	if !plan.Override.Equal(state.Override) && !plan.Override.IsUnknown() {
		body["override"] = plan.Override.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/ethernet/switch/port-isolation", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/ethernet/switch/port-isolation failed", err.Error())
			return
		}
		interfaceEthernetSwitchPortIsolationApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEthernetSwitchPortIsolationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceEthernetSwitchPortIsolationModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	// Fixed device row: it cannot be removed, only reconfigured. Drop it from
	// Terraform state and leave the row in place.
	_ = c
	resp.State.RemoveResource(ctx)
}

func (r *InterfaceEthernetSwitchPortIsolationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceEthernetSwitchPortIsolationLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/ethernet/switch/port-isolation matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceEthernetSwitchPortIsolationLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceEthernetSwitchPortIsolationLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/ethernet/switch/port-isolation", id)
}

func interfaceEthernetSwitchPortIsolationApply(ctx context.Context, obj client.Object, m *InterfaceEthernetSwitchPortIsolationModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["forward-to"]; ok {
		if v != "" {
			m.ForwardTo = types.StringValue(v)
		} else {
			m.ForwardTo = types.StringNull()
		}
	}
	if v, ok := obj["forwarding-override"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.ForwardingOverride = types.BoolValue(b)
		} else {
			m.ForwardingOverride = types.BoolNull()
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
	if v, ok := obj["override"]; ok {
		if v != "" {
			m.Override = types.StringValue(v)
		} else {
			m.Override = types.StringNull()
		}
	}
	if v, ok := obj["switch"]; ok {
		if v != "" {
			m.Switch = types.StringValue(v)
		} else {
			m.Switch = types.StringNull()
		}
	}
}
