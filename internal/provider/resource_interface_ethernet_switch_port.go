package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &InterfaceEthernetSwitchPortResource{}
	_ resource.ResourceWithImportState = &InterfaceEthernetSwitchPortResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceEthernetSwitchPortResource struct {
	reg *client.Registry
}

type InterfaceEthernetSwitchPortModel struct {
	ID            types.String `tfsdk:"id"`
	DefaultVlanID types.String `tfsdk:"default_vlan_id"`
	Name          types.String `tfsdk:"name"`
	Switch        types.String `tfsdk:"switch"`
	Router        types.String `tfsdk:"router"`
}

func NewInterfaceEthernetSwitchPortResource() resource.Resource {
	return &InterfaceEthernetSwitchPortResource{}
}

func (r *InterfaceEthernetSwitchPortResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ethernet_switch_port"
}

func (r *InterfaceEthernetSwitchPortResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceEthernetSwitchPortResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/ethernet/switch/port`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"default_vlan_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `default-vlan-id`. A number, or `auto`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `name`.",
			},
			"switch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `switch`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceEthernetSwitchPortResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceEthernetSwitchPortModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.DefaultVlanID.IsNull() || plan.DefaultVlanID.IsUnknown()) {
		body["default-vlan-id"] = plan.DefaultVlanID.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Switch.IsNull() || plan.Switch.IsUnknown()) {
		body["switch"] = plan.Switch.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/ethernet/switch/port", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/ethernet/switch/port failed", err.Error())
		return
	}
	interfaceEthernetSwitchPortApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEthernetSwitchPortResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceEthernetSwitchPortModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/ethernet/switch/port", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/ethernet/switch/port failed", err.Error())
		return
	}
	interfaceEthernetSwitchPortApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceEthernetSwitchPortResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceEthernetSwitchPortModel
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
	if !plan.DefaultVlanID.Equal(state.DefaultVlanID) && !plan.DefaultVlanID.IsUnknown() {
		body["default-vlan-id"] = plan.DefaultVlanID.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Switch.Equal(state.Switch) && !plan.Switch.IsUnknown() {
		body["switch"] = plan.Switch.ValueString()
	}
	var obj client.Object
	var err error
	if len(body) > 0 {
		obj, err = c.Set(ctx, "/interface/ethernet/switch/port", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/ethernet/switch/port failed", err.Error())
			return
		}
	} else {
		obj, err = c.GetByID(ctx, "/interface/ethernet/switch/port", state.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Read /interface/ethernet/switch/port failed", err.Error())
			return
		}
	}
	interfaceEthernetSwitchPortApply(ctx, obj, &plan)
	plan.ID = state.ID
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEthernetSwitchPortResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceEthernetSwitchPortModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/ethernet/switch/port", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/ethernet/switch/port failed", err.Error())
	}
}

func (r *InterfaceEthernetSwitchPortResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := lookupByNaturalKey(ctx, c, "/interface/ethernet/switch/port", id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/ethernet/switch/port matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

func interfaceEthernetSwitchPortApply(ctx context.Context, obj client.Object, m *InterfaceEthernetSwitchPortModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["default-vlan-id"]; ok {
		if v != "" {
			m.DefaultVlanID = types.StringValue(v)
		} else {
			m.DefaultVlanID = types.StringNull()
		}
	} else {
		m.DefaultVlanID = types.StringNull()
	}
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["switch"]; ok && v != "" {
		m.Switch = types.StringValue(v)
	} else {
		m.Switch = types.StringNull()
	}
}
