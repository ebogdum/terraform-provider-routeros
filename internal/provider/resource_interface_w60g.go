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
	_ resource.Resource                = &InterfaceW60gResource{}
	_ resource.ResourceWithImportState = &InterfaceW60gResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceW60gResource struct {
	reg *client.Registry
}

type InterfaceW60gModel struct {
	ID                  types.String `tfsdk:"id"`
	ARP                 types.String `tfsdk:"arp"`
	ARPTimeout          types.String `tfsdk:"arp_timeout"`
	Comment             types.String `tfsdk:"comment"`
	Disabled            types.String `tfsdk:"disabled"`
	Frequency           types.String `tfsdk:"frequency"`
	IsolateStations     types.String `tfsdk:"isolate_stations"`
	L2mtu               types.String `tfsdk:"l2mtu"`
	MACAddress          types.String `tfsdk:"mac_address"`
	MdmgFix             types.String `tfsdk:"mdmg_fix"`
	Mode                types.String `tfsdk:"mode"`
	MTU                 types.String `tfsdk:"mtu"`
	Name                types.String `tfsdk:"name"`
	Password            types.String `tfsdk:"password"`
	PutStationsInBridge types.String `tfsdk:"put_stations_in_bridge"`
	Region              types.String `tfsdk:"region"`
	ScanList            types.String `tfsdk:"scan_list"`
	Ssid                types.String `tfsdk:"ssid"`
	TxSector            types.String `tfsdk:"tx_sector"`
	Router              types.String `tfsdk:"router"`
}

func NewInterfaceW60gResource() resource.Resource { return &InterfaceW60gResource{} }

func (r *InterfaceW60gResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_w60g"
}

func (r *InterfaceW60gResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceW60gResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Requires 60GHz wAP60G hardware",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"arp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Read more >>",
			},
			"arp_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "ARP timeout is time how long ARP record is kept in ARP table after no packets are received from IP. Value auto equals to the value of arp-timeout in /ip settings , default is 30s",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Short description of the interface",
			},
			"disabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether interface is disabled",
			},
			"frequency": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Frequency used in communication (Only active on bridge device)",
			},
			"isolate_stations": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Don't allow communication between connected clients (from RouterOS 6.41)",
			},
			"l2mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Layer2 Maximum transmission unit",
			},
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "MAC address of the radio interface",
			},
			"mdmg_fix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Experimental feature working only on wAP60Gx3 devices, providing better point to multi point stability in some cases",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Operation mode",
			},
			"mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Layer3 Maximum transmission unit",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the interface",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Password used for AES encryption",
			},
			"put_stations_in_bridge": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Put newly created station device interfaces in this bridge",
			},
			"region": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Parameter to limit frequency use",
			},
			"scan_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Scan list to limit connectivity over frequencies in station mode",
			},
			"ssid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "SSID (service set identifier) is a name that identifies wireless network",
			},
			"tx_sector": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Disables beamforming and locks to selected radiation pattern",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceW60gResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceW60gModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.ARP.IsNull() || plan.ARP.IsUnknown()) {
		body["arp"] = plan.ARP.ValueString()
	}
	if !(plan.ARPTimeout.IsNull() || plan.ARPTimeout.IsUnknown()) {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = plan.Disabled.ValueString()
	}
	if !(plan.Frequency.IsNull() || plan.Frequency.IsUnknown()) {
		body["frequency"] = plan.Frequency.ValueString()
	}
	if !(plan.IsolateStations.IsNull() || plan.IsolateStations.IsUnknown()) {
		body["isolate-stations"] = plan.IsolateStations.ValueString()
	}
	if !(plan.L2mtu.IsNull() || plan.L2mtu.IsUnknown()) {
		body["l2mtu"] = plan.L2mtu.ValueString()
	}
	if !(plan.MACAddress.IsNull() || plan.MACAddress.IsUnknown()) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !(plan.MdmgFix.IsNull() || plan.MdmgFix.IsUnknown()) {
		body["mdmg-fix"] = plan.MdmgFix.ValueString()
	}
	if !(plan.Mode.IsNull() || plan.Mode.IsUnknown()) {
		body["mode"] = plan.Mode.ValueString()
	}
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Password.IsNull() || plan.Password.IsUnknown()) {
		body["password"] = plan.Password.ValueString()
	}
	if !(plan.PutStationsInBridge.IsNull() || plan.PutStationsInBridge.IsUnknown()) {
		body["put-stations-in-bridge"] = plan.PutStationsInBridge.ValueString()
	}
	if !(plan.Region.IsNull() || plan.Region.IsUnknown()) {
		body["region"] = plan.Region.ValueString()
	}
	if !(plan.ScanList.IsNull() || plan.ScanList.IsUnknown()) {
		body["scan-list"] = plan.ScanList.ValueString()
	}
	if !(plan.Ssid.IsNull() || plan.Ssid.IsUnknown()) {
		body["ssid"] = plan.Ssid.ValueString()
	}
	if !(plan.TxSector.IsNull() || plan.TxSector.IsUnknown()) {
		body["tx-sector"] = plan.TxSector.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/w60g", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/w60g failed", err.Error())
		return
	}
	interfaceW60gApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceW60gResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceW60gModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/w60g", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/w60g failed", err.Error())
		return
	}
	interfaceW60gApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceW60gResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceW60gModel
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
	if !plan.ARP.Equal(state.ARP) && !plan.ARP.IsUnknown() {
		body["arp"] = plan.ARP.ValueString()
	}
	if !plan.ARPTimeout.Equal(state.ARPTimeout) && !plan.ARPTimeout.IsUnknown() {
		body["arp-timeout"] = plan.ARPTimeout.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = plan.Disabled.ValueString()
	}
	if !plan.Frequency.Equal(state.Frequency) && !plan.Frequency.IsUnknown() {
		body["frequency"] = plan.Frequency.ValueString()
	}
	if !plan.IsolateStations.Equal(state.IsolateStations) && !plan.IsolateStations.IsUnknown() {
		body["isolate-stations"] = plan.IsolateStations.ValueString()
	}
	if !plan.L2mtu.Equal(state.L2mtu) && !plan.L2mtu.IsUnknown() {
		body["l2mtu"] = plan.L2mtu.ValueString()
	}
	if !plan.MACAddress.Equal(state.MACAddress) && !plan.MACAddress.IsUnknown() {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.MdmgFix.Equal(state.MdmgFix) && !plan.MdmgFix.IsUnknown() {
		body["mdmg-fix"] = plan.MdmgFix.ValueString()
	}
	if !plan.Mode.Equal(state.Mode) && !plan.Mode.IsUnknown() {
		body["mode"] = plan.Mode.ValueString()
	}
	if !plan.MTU.Equal(state.MTU) && !plan.MTU.IsUnknown() {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Password.Equal(state.Password) && !plan.Password.IsUnknown() {
		body["password"] = plan.Password.ValueString()
	}
	if !plan.PutStationsInBridge.Equal(state.PutStationsInBridge) && !plan.PutStationsInBridge.IsUnknown() {
		body["put-stations-in-bridge"] = plan.PutStationsInBridge.ValueString()
	}
	if !plan.Region.Equal(state.Region) && !plan.Region.IsUnknown() {
		body["region"] = plan.Region.ValueString()
	}
	if !plan.ScanList.Equal(state.ScanList) && !plan.ScanList.IsUnknown() {
		body["scan-list"] = plan.ScanList.ValueString()
	}
	if !plan.Ssid.Equal(state.Ssid) && !plan.Ssid.IsUnknown() {
		body["ssid"] = plan.Ssid.ValueString()
	}
	if !plan.TxSector.Equal(state.TxSector) && !plan.TxSector.IsUnknown() {
		body["tx-sector"] = plan.TxSector.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/w60g", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/w60g failed", err.Error())
			return
		}
		interfaceW60gApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceW60gResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceW60gModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/w60g", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/w60g failed", err.Error())
	}
}

func (r *InterfaceW60gResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceW60gLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/w60g matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceW60gLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceW60gLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/w60g", id)
}

func interfaceW60gApply(ctx context.Context, obj client.Object, m *InterfaceW60gModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["arp"]; ok {
		_ = v
		if v != "" {
			m.ARP = types.StringValue(v)
		} else {
			m.ARP = types.StringNull()
		}
	} else {
		m.ARP = types.StringNull()
	}
	if v, ok := obj["arp-timeout"]; ok {
		_ = v
		if v != "" {
			m.ARPTimeout = types.StringValue(v)
		} else {
			m.ARPTimeout = types.StringNull()
		}
	} else {
		m.ARPTimeout = types.StringNull()
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
		if v != "" {
			m.Disabled = types.StringValue(v)
		} else {
			m.Disabled = types.StringNull()
		}
	} else {
		m.Disabled = types.StringNull()
	}
	if v, ok := obj["frequency"]; ok {
		_ = v
		if v != "" {
			m.Frequency = types.StringValue(v)
		} else {
			m.Frequency = types.StringNull()
		}
	} else {
		m.Frequency = types.StringNull()
	}
	if v, ok := obj["isolate-stations"]; ok {
		_ = v
		if v != "" {
			m.IsolateStations = types.StringValue(v)
		} else {
			m.IsolateStations = types.StringNull()
		}
	} else {
		m.IsolateStations = types.StringNull()
	}
	if v, ok := obj["l2mtu"]; ok {
		_ = v
		if v != "" {
			m.L2mtu = types.StringValue(v)
		} else {
			m.L2mtu = types.StringNull()
		}
	} else {
		m.L2mtu = types.StringNull()
	}
	if v, ok := obj["mac-address"]; ok {
		_ = v
		if v != "" {
			m.MACAddress = types.StringValue(v)
		} else {
			m.MACAddress = types.StringNull()
		}
	} else {
		m.MACAddress = types.StringNull()
	}
	if v, ok := obj["mdmg-fix"]; ok {
		_ = v
		if v != "" {
			m.MdmgFix = types.StringValue(v)
		} else {
			m.MdmgFix = types.StringNull()
		}
	} else {
		m.MdmgFix = types.StringNull()
	}
	if v, ok := obj["mode"]; ok {
		_ = v
		if v != "" {
			m.Mode = types.StringValue(v)
		} else {
			m.Mode = types.StringNull()
		}
	} else {
		m.Mode = types.StringNull()
	}
	if v, ok := obj["mtu"]; ok {
		_ = v
		if v != "" {
			m.MTU = types.StringValue(v)
		} else {
			m.MTU = types.StringNull()
		}
	} else {
		m.MTU = types.StringNull()
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
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.Password already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["password"]; ok && v != "" {
		_ = v
		if v != "" {
			m.Password = types.StringValue(v)
		} else {
			m.Password = types.StringNull()
		}
	} else if m.Password.IsUnknown() {
		m.Password = types.StringNull()
	}
	if v, ok := obj["put-stations-in-bridge"]; ok {
		_ = v
		if v != "" {
			m.PutStationsInBridge = types.StringValue(v)
		} else {
			m.PutStationsInBridge = types.StringNull()
		}
	} else {
		m.PutStationsInBridge = types.StringNull()
	}
	if v, ok := obj["region"]; ok {
		_ = v
		if v != "" {
			m.Region = types.StringValue(v)
		} else {
			m.Region = types.StringNull()
		}
	} else {
		m.Region = types.StringNull()
	}
	if v, ok := obj["scan-list"]; ok {
		_ = v
		if v != "" {
			m.ScanList = types.StringValue(v)
		} else {
			m.ScanList = types.StringNull()
		}
	} else {
		m.ScanList = types.StringNull()
	}
	if v, ok := obj["ssid"]; ok {
		_ = v
		if v != "" {
			m.Ssid = types.StringValue(v)
		} else {
			m.Ssid = types.StringNull()
		}
	} else {
		m.Ssid = types.StringNull()
	}
	if v, ok := obj["tx-sector"]; ok {
		_ = v
		if v != "" {
			m.TxSector = types.StringValue(v)
		} else {
			m.TxSector = types.StringNull()
		}
	} else {
		m.TxSector = types.StringNull()
	}
}
