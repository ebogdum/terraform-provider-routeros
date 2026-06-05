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
	_ resource.Resource                = &InterfaceWifiDatapathResource{}
	_ resource.ResourceWithImportState = &InterfaceWifiDatapathResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceWifiDatapathResource struct {
	reg *client.Registry
}

type InterfaceWifiDatapathModel struct {
	ID                types.String `tfsdk:"id"`
	Bridge            types.String `tfsdk:"bridge"`
	BridgeCost        types.String `tfsdk:"bridge_cost"`
	BridgeHorizon     types.String `tfsdk:"bridge_horizon"`
	ClientIsolation   types.String `tfsdk:"client_isolation"`
	Comment           types.String `tfsdk:"comment"`
	Disabled          types.Bool   `tfsdk:"disabled"`
	InterfaceList     types.String `tfsdk:"interface_list"`
	Name              types.String `tfsdk:"name"`
	OpenFlowSwitch    types.String `tfsdk:"open_flow_switch"`
	Openflow          types.String `tfsdk:"openflow"`
	TrafficProcessing types.String `tfsdk:"traffic_processing"`
	VLANID            types.String `tfsdk:"vlan_id"`
	Router            types.String `tfsdk:"router"`
}

func NewInterfaceWifiDatapathResource() resource.Resource { return &InterfaceWifiDatapathResource{} }

func (r *InterfaceWifiDatapathResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_wifi_datapath"
}

func (r *InterfaceWifiDatapathResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *InterfaceWifiDatapathResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/wifi/datapath`.",
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
			"bridge_cost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bridge_horizon": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"client_isolation": schema.StringAttribute{
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
			"interface_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"open_flow_switch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"openflow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"traffic_processing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vlan_id": schema.StringAttribute{
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

func (r *InterfaceWifiDatapathResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceWifiDatapathModel
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
	if !(plan.BridgeCost.IsNull() || plan.BridgeCost.IsUnknown()) {
		body["bridge-cost"] = plan.BridgeCost.ValueString()
	}
	if !(plan.BridgeHorizon.IsNull() || plan.BridgeHorizon.IsUnknown()) {
		body["bridge-horizon"] = plan.BridgeHorizon.ValueString()
	}
	if !(plan.ClientIsolation.IsNull() || plan.ClientIsolation.IsUnknown()) {
		body["client-isolation"] = plan.ClientIsolation.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.InterfaceList.IsNull() || plan.InterfaceList.IsUnknown()) {
		body["interface-list"] = plan.InterfaceList.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.OpenFlowSwitch.IsNull() || plan.OpenFlowSwitch.IsUnknown()) {
		body["open-flow-switch"] = plan.OpenFlowSwitch.ValueString()
	}
	if !(plan.Openflow.IsNull() || plan.Openflow.IsUnknown()) {
		body["openflow"] = plan.Openflow.ValueString()
	}
	if !(plan.TrafficProcessing.IsNull() || plan.TrafficProcessing.IsUnknown()) {
		body["traffic-processing"] = plan.TrafficProcessing.ValueString()
	}
	if !(plan.VLANID.IsNull() || plan.VLANID.IsUnknown()) {
		body["vlan-id"] = plan.VLANID.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/wifi/datapath", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/wifi/datapath failed", err.Error())
		return
	}
	interfaceWifiDatapathApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiDatapathResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceWifiDatapathModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/wifi/datapath", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/wifi/datapath failed", err.Error())
		return
	}
	interfaceWifiDatapathApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceWifiDatapathResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceWifiDatapathModel
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
	if !plan.Bridge.Equal(state.Bridge) {
		body["bridge"] = plan.Bridge.ValueString()
	}
	if !plan.BridgeCost.Equal(state.BridgeCost) {
		body["bridge-cost"] = plan.BridgeCost.ValueString()
	}
	if !plan.BridgeHorizon.Equal(state.BridgeHorizon) {
		body["bridge-horizon"] = plan.BridgeHorizon.ValueString()
	}
	if !plan.ClientIsolation.Equal(state.ClientIsolation) {
		body["client-isolation"] = plan.ClientIsolation.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.InterfaceList.Equal(state.InterfaceList) {
		body["interface-list"] = plan.InterfaceList.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.OpenFlowSwitch.Equal(state.OpenFlowSwitch) {
		body["open-flow-switch"] = plan.OpenFlowSwitch.ValueString()
	}
	if !plan.Openflow.Equal(state.Openflow) {
		body["openflow"] = plan.Openflow.ValueString()
	}
	if !plan.TrafficProcessing.Equal(state.TrafficProcessing) {
		body["traffic-processing"] = plan.TrafficProcessing.ValueString()
	}
	if !plan.VLANID.Equal(state.VLANID) {
		body["vlan-id"] = plan.VLANID.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/wifi/datapath", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/wifi/datapath failed", err.Error())
			return
		}
		interfaceWifiDatapathApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiDatapathResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceWifiDatapathModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/wifi/datapath", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/wifi/datapath failed", err.Error())
	}
}

func (r *InterfaceWifiDatapathResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceWifiDatapathLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/wifi/datapath matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceWifiDatapathLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceWifiDatapathLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/interface/wifi/datapath", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func interfaceWifiDatapathApply(ctx context.Context, obj client.Object, m *InterfaceWifiDatapathModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["bridge"]; ok {
		_ = v
		if v != "" {
			m.Bridge = types.StringValue(v)
		} else {
			m.Bridge = types.StringNull()
		}
	} else {
		m.Bridge = types.StringNull()
	}
	if v, ok := obj["bridge-cost"]; ok {
		_ = v
		if v != "" {
			m.BridgeCost = types.StringValue(v)
		} else {
			m.BridgeCost = types.StringNull()
		}
	} else {
		m.BridgeCost = types.StringNull()
	}
	if v, ok := obj["bridge-horizon"]; ok {
		_ = v
		if v != "" {
			m.BridgeHorizon = types.StringValue(v)
		} else {
			m.BridgeHorizon = types.StringNull()
		}
	} else {
		m.BridgeHorizon = types.StringNull()
	}
	if v, ok := obj["client-isolation"]; ok {
		_ = v
		if v != "" {
			m.ClientIsolation = types.StringValue(v)
		} else {
			m.ClientIsolation = types.StringNull()
		}
	} else {
		m.ClientIsolation = types.StringNull()
	}
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
	if v, ok := obj["interface-list"]; ok {
		_ = v
		if v != "" {
			m.InterfaceList = types.StringValue(v)
		} else {
			m.InterfaceList = types.StringNull()
		}
	} else {
		m.InterfaceList = types.StringNull()
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
	if v, ok := obj["open-flow-switch"]; ok {
		_ = v
		if v != "" {
			m.OpenFlowSwitch = types.StringValue(v)
		} else {
			m.OpenFlowSwitch = types.StringNull()
		}
	} else {
		m.OpenFlowSwitch = types.StringNull()
	}
	if v, ok := obj["openflow"]; ok {
		_ = v
		if v != "" {
			m.Openflow = types.StringValue(v)
		} else {
			m.Openflow = types.StringNull()
		}
	} else {
		m.Openflow = types.StringNull()
	}
	if v, ok := obj["traffic-processing"]; ok {
		_ = v
		if v != "" {
			m.TrafficProcessing = types.StringValue(v)
		} else {
			m.TrafficProcessing = types.StringNull()
		}
	} else {
		m.TrafficProcessing = types.StringNull()
	}
	if v, ok := obj["vlan-id"]; ok {
		_ = v
		if v != "" {
			m.VLANID = types.StringValue(v)
		} else {
			m.VLANID = types.StringNull()
		}
	} else {
		m.VLANID = types.StringNull()
	}
}
