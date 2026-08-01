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
	_ resource.Resource                = &IotLoraResource{}
	_ resource.ResourceWithImportState = &IotLoraResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IotLoraResource struct {
	reg *client.Registry
}

type IotLoraModel struct {
	ID                 types.String `tfsdk:"id"`
	TxImmediateDelayUs types.String `tfsdk:"tx_immediate_delay_us"`
	Long               types.String `tfsdk:"long"`
	Lat                types.String `tfsdk:"lat"`
	Antenna            types.String `tfsdk:"antenna"`
	Alt                types.String `tfsdk:"alt"`
	AntennaGain        types.String `tfsdk:"antenna_gain"`
	ChannelPlan        types.String `tfsdk:"channel_plan"`
	Disabled           types.String `tfsdk:"disabled"`
	Forward            types.String `tfsdk:"forward"`
	GatewayID          types.String `tfsdk:"gateway_id"`
	LbtEnabled         types.String `tfsdk:"lbt_enabled"`
	ListenTime         types.String `tfsdk:"listen_time"`
	Name               types.String `tfsdk:"name"`
	Network            types.String `tfsdk:"network"`
	RssiThreshold      types.String `tfsdk:"rssi_threshold"`
	Servers            types.String `tfsdk:"servers"`
	SpoofGps           types.String `tfsdk:"spoof_gps"`
	SrcAddress         types.String `tfsdk:"src_address"`
	Router             types.String `tfsdk:"router"`
}

func NewIotLoraResource() resource.Resource { return &IotLoraResource{} }

func (r *IotLoraResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_iot_lora"
}

func (r *IotLoraResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IotLoraResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Requires LoRa hardware/package",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"tx_immediate_delay_us": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tx-immediate-delay-us`.",
			},
			"long": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `long`.",
			},
			"lat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `lat`.",
			},
			"antenna": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `antenna`.",
			},
			"alt": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `alt`.",
			},
			"antenna_gain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Antenna gain in dBi. This value should be equal to\u00a0 setup-antenna-gain minus cable-loss . Using\u00a06.5 dBi antenna, 6.5 is the value to be configured (not taking into account cable loss).\u00a0 Output power of the gateway is dictated by the server. The gateway will calculate its actual output power by subtracting antenna-gain setting from server_value (value received in the downlink message).",
			},
			"channel_plan": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Frequency plans for various regions.",
			},
			"disabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether LoRaWAN gateway is disabled.",
			},
			"forward": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Defines what kind of packets should be forwarded to Network server: crc-validtaion - Forward valid packets with correct CRC. dev-addr-validtaion - Checks if DevAddr of the packet corresponds to the NetID and if not, drops the packet. The following sequence happens: 1) Dev. Addr value gets \"obtained\" from the received LoRa packet; 2) Dev. Addr is \"compared\" against \"valid\" Net IDs list; 3) If there is no Net ID for the Dev. Addr, the packet is not forwarded; 4) If Net ID is valid, Dev. Addr range is valid, the packet is forwarded. proprietary-traffic - Checks the content of the LoRa packet and if the \"type\" of the frame is \"proprietary\", the packet is not forwarded.",
			},
			"gateway_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Gateway ID or Gateway EUI, is used when registering the gateway with the server.",
			},
			"lbt_enabled": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether gateway should use LBT (Listen Before Talk) protocol.",
			},
			"listen_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Time in microseconds to track RSSI before TX (used when lbt-enabled=yes ).",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of LoRaWAN gateway.",
			},
			"network": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether sync word should (network=private) or should not (network=public) be used.",
			},
			"rssi_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RSSI value to determine whether forwarder may use specific channel to talk. If RSSI value is below rssi-threshold , channel could be used (used when lbt-enabled=yes ).",
			},
			"servers": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the server from the /iot lora servers section.",
			},
			"spoof_gps": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Set custom GPS location: Latitude [-90..90] Longitude [-180..180] Altitude( m ) [-2147483648..2147483647]",
			},
			"src_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specifies uplink packet source address if necessary (address should match an address configured on the RB).",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IotLoraResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IotLoraModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AntennaGain.IsNull() || plan.AntennaGain.IsUnknown()) {
		body["antenna-gain"] = plan.AntennaGain.ValueString()
	}
	if !(plan.ChannelPlan.IsNull() || plan.ChannelPlan.IsUnknown()) {
		body["channel-plan"] = plan.ChannelPlan.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = plan.Disabled.ValueString()
	}
	if !(plan.Forward.IsNull() || plan.Forward.IsUnknown()) {
		body["forward"] = plan.Forward.ValueString()
	}
	if !(plan.GatewayID.IsNull() || plan.GatewayID.IsUnknown()) {
		body["gateway-id"] = plan.GatewayID.ValueString()
	}
	if !(plan.LbtEnabled.IsNull() || plan.LbtEnabled.IsUnknown()) {
		body["lbt-enabled"] = plan.LbtEnabled.ValueString()
	}
	if !(plan.ListenTime.IsNull() || plan.ListenTime.IsUnknown()) {
		body["listen-time"] = plan.ListenTime.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Network.IsNull() || plan.Network.IsUnknown()) {
		body["network"] = plan.Network.ValueString()
	}
	if !(plan.RssiThreshold.IsNull() || plan.RssiThreshold.IsUnknown()) {
		body["rssi-threshold"] = plan.RssiThreshold.ValueString()
	}
	if !(plan.Servers.IsNull() || plan.Servers.IsUnknown()) {
		body["servers"] = plan.Servers.ValueString()
	}
	if !(plan.SpoofGps.IsNull() || plan.SpoofGps.IsUnknown()) {
		body["spoof-gps"] = plan.SpoofGps.ValueString()
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !(plan.Alt.IsNull() || plan.Alt.IsUnknown()) {
		body["alt"] = plan.Alt.ValueString()
	}
	if !(plan.Antenna.IsNull() || plan.Antenna.IsUnknown()) {
		body["antenna"] = plan.Antenna.ValueString()
	}
	if !(plan.Lat.IsNull() || plan.Lat.IsUnknown()) {
		body["lat"] = plan.Lat.ValueString()
	}
	if !(plan.Long.IsNull() || plan.Long.IsUnknown()) {
		body["long"] = plan.Long.ValueString()
	}
	if !(plan.TxImmediateDelayUs.IsNull() || plan.TxImmediateDelayUs.IsUnknown()) {
		body["tx-immediate-delay-us"] = plan.TxImmediateDelayUs.ValueString()
	}
	obj, err := c.Add(ctx, "/iot/lora", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /iot/lora failed", err.Error())
		return
	}
	iotLoraApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IotLoraResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IotLoraModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/iot/lora", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /iot/lora failed", err.Error())
		return
	}
	iotLoraApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IotLoraResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IotLoraModel
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
	if !plan.AntennaGain.Equal(state.AntennaGain) && !plan.AntennaGain.IsUnknown() {
		body["antenna-gain"] = plan.AntennaGain.ValueString()
	}
	if !plan.ChannelPlan.Equal(state.ChannelPlan) && !plan.ChannelPlan.IsUnknown() {
		body["channel-plan"] = plan.ChannelPlan.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = plan.Disabled.ValueString()
	}
	if !plan.Forward.Equal(state.Forward) && !plan.Forward.IsUnknown() {
		body["forward"] = plan.Forward.ValueString()
	}
	if !plan.GatewayID.Equal(state.GatewayID) && !plan.GatewayID.IsUnknown() {
		body["gateway-id"] = plan.GatewayID.ValueString()
	}
	if !plan.LbtEnabled.Equal(state.LbtEnabled) && !plan.LbtEnabled.IsUnknown() {
		body["lbt-enabled"] = plan.LbtEnabled.ValueString()
	}
	if !plan.ListenTime.Equal(state.ListenTime) && !plan.ListenTime.IsUnknown() {
		body["listen-time"] = plan.ListenTime.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Network.Equal(state.Network) && !plan.Network.IsUnknown() {
		body["network"] = plan.Network.ValueString()
	}
	if !plan.RssiThreshold.Equal(state.RssiThreshold) && !plan.RssiThreshold.IsUnknown() {
		body["rssi-threshold"] = plan.RssiThreshold.ValueString()
	}
	if !plan.Servers.Equal(state.Servers) && !plan.Servers.IsUnknown() {
		body["servers"] = plan.Servers.ValueString()
	}
	if !plan.SpoofGps.Equal(state.SpoofGps) && !plan.SpoofGps.IsUnknown() {
		body["spoof-gps"] = plan.SpoofGps.ValueString()
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) && !plan.SrcAddress.IsUnknown() {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !plan.Alt.Equal(state.Alt) && !plan.Alt.IsUnknown() {
		body["alt"] = plan.Alt.ValueString()
	}
	if !plan.Antenna.Equal(state.Antenna) && !plan.Antenna.IsUnknown() {
		body["antenna"] = plan.Antenna.ValueString()
	}
	if !plan.Lat.Equal(state.Lat) && !plan.Lat.IsUnknown() {
		body["lat"] = plan.Lat.ValueString()
	}
	if !plan.Long.Equal(state.Long) && !plan.Long.IsUnknown() {
		body["long"] = plan.Long.ValueString()
	}
	if !plan.TxImmediateDelayUs.Equal(state.TxImmediateDelayUs) && !plan.TxImmediateDelayUs.IsUnknown() {
		body["tx-immediate-delay-us"] = plan.TxImmediateDelayUs.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/iot/lora", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /iot/lora failed", err.Error())
			return
		}
		iotLoraApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IotLoraResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IotLoraModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/iot/lora", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /iot/lora failed", err.Error())
	}
}

func (r *IotLoraResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iotLoraLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /iot/lora matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iotLoraLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iotLoraLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/iot/lora", id)
}

func iotLoraApply(ctx context.Context, obj client.Object, m *IotLoraModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["tx-immediate-delay-us"]; ok && v != "" {
		m.TxImmediateDelayUs = types.StringValue(v)
	} else {
		m.TxImmediateDelayUs = types.StringNull()
	}
	if v, ok := obj["long"]; ok && v != "" {
		m.Long = types.StringValue(v)
	} else {
		m.Long = types.StringNull()
	}
	if v, ok := obj["lat"]; ok && v != "" {
		m.Lat = types.StringValue(v)
	} else {
		m.Lat = types.StringNull()
	}
	if v, ok := obj["antenna"]; ok && v != "" {
		m.Antenna = types.StringValue(v)
	} else {
		m.Antenna = types.StringNull()
	}
	if v, ok := obj["alt"]; ok && v != "" {
		m.Alt = types.StringValue(v)
	} else {
		m.Alt = types.StringNull()
	}
	if v, ok := obj["antenna-gain"]; ok {
		_ = v
		if v != "" {
			m.AntennaGain = types.StringValue(v)
		} else {
			m.AntennaGain = types.StringNull()
		}
	} else {
		m.AntennaGain = types.StringNull()
	}
	if v, ok := obj["channel-plan"]; ok {
		_ = v
		if v != "" {
			m.ChannelPlan = types.StringValue(v)
		} else {
			m.ChannelPlan = types.StringNull()
		}
	} else {
		m.ChannelPlan = types.StringNull()
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
	if v, ok := obj["forward"]; ok {
		_ = v
		if v != "" {
			m.Forward = types.StringValue(v)
		} else {
			m.Forward = types.StringNull()
		}
	} else {
		m.Forward = types.StringNull()
	}
	if v, ok := obj["gateway-id"]; ok {
		_ = v
		if v != "" {
			m.GatewayID = types.StringValue(v)
		} else {
			m.GatewayID = types.StringNull()
		}
	} else {
		m.GatewayID = types.StringNull()
	}
	if v, ok := obj["lbt-enabled"]; ok {
		_ = v
		if v != "" {
			m.LbtEnabled = types.StringValue(v)
		} else {
			m.LbtEnabled = types.StringNull()
		}
	} else {
		m.LbtEnabled = types.StringNull()
	}
	if v, ok := obj["listen-time"]; ok {
		_ = v
		if v != "" {
			m.ListenTime = types.StringValue(v)
		} else {
			m.ListenTime = types.StringNull()
		}
	} else {
		m.ListenTime = types.StringNull()
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
	if v, ok := obj["network"]; ok {
		_ = v
		if v != "" {
			m.Network = types.StringValue(v)
		} else {
			m.Network = types.StringNull()
		}
	} else {
		m.Network = types.StringNull()
	}
	if v, ok := obj["rssi-threshold"]; ok {
		_ = v
		if v != "" {
			m.RssiThreshold = types.StringValue(v)
		} else {
			m.RssiThreshold = types.StringNull()
		}
	} else {
		m.RssiThreshold = types.StringNull()
	}
	if v, ok := obj["servers"]; ok {
		_ = v
		if v != "" {
			m.Servers = types.StringValue(v)
		} else {
			m.Servers = types.StringNull()
		}
	} else {
		m.Servers = types.StringNull()
	}
	if v, ok := obj["spoof-gps"]; ok {
		_ = v
		if v != "" {
			m.SpoofGps = types.StringValue(v)
		} else {
			m.SpoofGps = types.StringNull()
		}
	} else {
		m.SpoofGps = types.StringNull()
	}
	if v, ok := obj["src-address"]; ok {
		_ = v
		if v != "" {
			m.SrcAddress = types.StringValue(v)
		} else {
			m.SrcAddress = types.StringNull()
		}
	} else {
		m.SrcAddress = types.StringNull()
	}
}
