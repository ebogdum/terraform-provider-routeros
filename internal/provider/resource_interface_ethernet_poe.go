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
	_ resource.Resource                = &InterfaceEthernetPoeResource{}
	_ resource.ResourceWithImportState = &InterfaceEthernetPoeResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceEthernetPoeResource struct {
	reg *client.Registry
}

type InterfaceEthernetPoeModel struct {
	ID         types.String `tfsdk:"id"`
	Export     types.String `tfsdk:"export"`
	Monitor    types.String `tfsdk:"monitor"`
	PowerCycle types.String `tfsdk:"power_cycle"`
	Print      types.String `tfsdk:"print"`
	Router     types.String `tfsdk:"router"`
}

func NewInterfaceEthernetPoeResource() resource.Resource { return &InterfaceEthernetPoeResource{} }

func (r *InterfaceEthernetPoeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ethernet_poe"
}

func (r *InterfaceEthernetPoeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *InterfaceEthernetPoeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Requires PoE-capable ethernet port",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"export": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "export is displayed under \u00a0 /interface ethernet \u00a0 menu.",
			},
			"monitor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Shows poe-out-status of a specified port, or all ports with \u00a0 /interface ethernet poe monitor [find] \u00a0 command.",
			},
			"power_cycle": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Disables PoE-Out power for a specified period of time.",
			},
			"print": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Prints PoE-Out related settings.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceEthernetPoeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceEthernetPoeModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Export.IsNull() || plan.Export.IsUnknown()) {
		body["export"] = plan.Export.ValueString()
	}
	if !(plan.Monitor.IsNull() || plan.Monitor.IsUnknown()) {
		body["monitor"] = plan.Monitor.ValueString()
	}
	if !(plan.PowerCycle.IsNull() || plan.PowerCycle.IsUnknown()) {
		body["power-cycle"] = plan.PowerCycle.ValueString()
	}
	if !(plan.Print.IsNull() || plan.Print.IsUnknown()) {
		body["print"] = plan.Print.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/ethernet/poe", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/ethernet/poe failed", err.Error())
		return
	}
	interfaceEthernetPoeApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEthernetPoeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceEthernetPoeModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/ethernet/poe", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/ethernet/poe failed", err.Error())
		return
	}
	interfaceEthernetPoeApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceEthernetPoeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceEthernetPoeModel
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
	if !plan.Export.Equal(state.Export) {
		body["export"] = plan.Export.ValueString()
	}
	if !plan.Monitor.Equal(state.Monitor) {
		body["monitor"] = plan.Monitor.ValueString()
	}
	if !plan.PowerCycle.Equal(state.PowerCycle) {
		body["power-cycle"] = plan.PowerCycle.ValueString()
	}
	if !plan.Print.Equal(state.Print) {
		body["print"] = plan.Print.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/ethernet/poe", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/ethernet/poe failed", err.Error())
			return
		}
		interfaceEthernetPoeApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceEthernetPoeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceEthernetPoeModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/ethernet/poe", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/ethernet/poe failed", err.Error())
	}
}

func (r *InterfaceEthernetPoeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceEthernetPoeLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/ethernet/poe matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceEthernetPoeLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceEthernetPoeLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/ethernet/poe", id)
}

func interfaceEthernetPoeApply(ctx context.Context, obj client.Object, m *InterfaceEthernetPoeModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["export"]; ok {
		_ = v
		if v != "" {
			m.Export = types.StringValue(v)
		} else {
			m.Export = types.StringNull()
		}
	} else {
		m.Export = types.StringNull()
	}
	if v, ok := obj["monitor"]; ok {
		_ = v
		if v != "" {
			m.Monitor = types.StringValue(v)
		} else {
			m.Monitor = types.StringNull()
		}
	} else {
		m.Monitor = types.StringNull()
	}
	if v, ok := obj["power-cycle"]; ok {
		_ = v
		if v != "" {
			m.PowerCycle = types.StringValue(v)
		} else {
			m.PowerCycle = types.StringNull()
		}
	} else {
		m.PowerCycle = types.StringNull()
	}
	if v, ok := obj["print"]; ok {
		_ = v
		if v != "" {
			m.Print = types.StringValue(v)
		} else {
			m.Print = types.StringNull()
		}
	} else {
		m.Print = types.StringNull()
	}
}
