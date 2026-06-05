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
	_ resource.Resource                = &InterfaceResource{}
	_ resource.ResourceWithImportState = &InterfaceResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceResource struct {
	reg *client.Registry
}

type InterfaceModel struct {
	ID                   types.String `tfsdk:"id"`
	ActualMTU            types.Int64  `tfsdk:"actual_mtu"`
	AnswerTime           types.String `tfsdk:"answer_time"`
	Caps                 types.Int64  `tfsdk:"caps"`
	Comment              types.String `tfsdk:"comment"`
	DefaultName          types.String `tfsdk:"default_name"`
	Disabled             types.Bool   `tfsdk:"disabled"`
	Dynamic              types.Bool   `tfsdk:"dynamic"`
	FpRpsDrop            types.Int64  `tfsdk:"fp_rps_drop"`
	FpRxByte             types.Int64  `tfsdk:"fp_rx_byte"`
	FpRxPacket           types.Int64  `tfsdk:"fp_rx_packet"`
	FpTxByte             types.Int64  `tfsdk:"fp_tx_byte"`
	FpTxPacket           types.Int64  `tfsdk:"fp_tx_packet"`
	FpTxRxPacketRate     types.String `tfsdk:"fp_tx_rx_packet_rate"`
	FpTxRxRate           types.String `tfsdk:"fp_tx_rx_rate"`
	Inactive             types.Bool   `tfsdk:"inactive"`
	L2MTU                types.Int64  `tfsdk:"l2_mtu"`
	LastLinkDownTime     types.String `tfsdk:"last_link_down_time"`
	LastLinkUpTime       types.String `tfsdk:"last_link_up_time"`
	Link                 types.Int64  `tfsdk:"link"`
	LinkDowns            types.Int64  `tfsdk:"link_downs"`
	MACAddress           types.String `tfsdk:"mac_address"`
	MTU                  types.Int64  `tfsdk:"mtu"`
	Name                 types.String `tfsdk:"name"`
	Nodefname            types.String `tfsdk:"nodefname"`
	Notrunning           types.String `tfsdk:"notrunning"`
	Passthrough          types.Bool   `tfsdk:"passthrough"`
	ResetTrafficCounters types.String `tfsdk:"reset_traffic_counters"`
	Running              types.Bool   `tfsdk:"running"`
	RxByte               types.Int64  `tfsdk:"rx_byte"`
	RxDrop               types.Int64  `tfsdk:"rx_drop"`
	RxError              types.Int64  `tfsdk:"rx_error"`
	RxPacket             types.Int64  `tfsdk:"rx_packet"`
	Slave                types.Bool   `tfsdk:"slave"`
	Torch                types.String `tfsdk:"torch"`
	TxByte               types.Int64  `tfsdk:"tx_byte"`
	TxDrop               types.Int64  `tfsdk:"tx_drop"`
	TxError              types.Int64  `tfsdk:"tx_error"`
	TxPacket             types.Int64  `tfsdk:"tx_packet"`
	TxQueueDrop          types.Int64  `tfsdk:"tx_queue_drop"`
	TxQueueDrops         types.String `tfsdk:"tx_queue_drops"`
	TxRxBytes            types.String `tfsdk:"tx_rx_bytes"`
	TxRxDrops            types.String `tfsdk:"tx_rx_drops"`
	TxRxErrors           types.String `tfsdk:"tx_rx_errors"`
	TxRxPacketRate       types.String `tfsdk:"tx_rx_packet_rate"`
	TxRxPackets          types.String `tfsdk:"tx_rx_packets"`
	TxRxRate             types.String `tfsdk:"tx_rx_rate"`
	Type                 types.Int64  `tfsdk:"type"`
	Vrf                  types.String `tfsdk:"vrf"`
	Router               types.String `tfsdk:"router"`
}

func NewInterfaceResource() resource.Resource { return &InterfaceResource{} }

func (r *InterfaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface"
}

func (r *InterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *InterfaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "/interface is mostly read-only — interfaces are created by their specific subtypes (/interface/bridge, /interface/wireguard, etc.).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"actual_mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"answer_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"caps": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"default_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dynamic": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"fp_rps_drop": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"fp_rx_byte": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"fp_rx_packet": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"fp_tx_byte": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"fp_tx_packet": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"fp_tx_rx_packet_rate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"fp_tx_rx_rate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"inactive": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"l2_mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"last_link_down_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"last_link_up_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"link": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"link_downs": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mac_address": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsMAC()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeMAC()},
			},
			"mtu": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nodefname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"notrunning": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"passthrough": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"reset_traffic_counters": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"running": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_byte": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_drop": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_error": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rx_packet": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"slave": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"torch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_byte": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_drop": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_error": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_packet": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_queue_drop": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_queue_drops": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_bytes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_drops": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_errors": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_packet_rate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_packets": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tx_rx_rate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"type": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vrf": schema.StringAttribute{
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

func (r *InterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceModel
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
	if !(plan.FpTxRxPacketRate.IsNull() || plan.FpTxRxPacketRate.IsUnknown()) {
		body["fp-tx-rx-packet-rate"] = plan.FpTxRxPacketRate.ValueString()
	}
	if !(plan.FpTxRxRate.IsNull() || plan.FpTxRxRate.IsUnknown()) {
		body["fp-tx-rx-rate"] = plan.FpTxRxRate.ValueString()
	}
	if !(plan.Inactive.IsNull() || plan.Inactive.IsUnknown()) {
		body["inactive"] = client.FormatBool(plan.Inactive.ValueBool())
	}
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = client.FormatInt64(plan.MTU.ValueInt64())
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Nodefname.IsNull() || plan.Nodefname.IsUnknown()) {
		body["nodefname"] = plan.Nodefname.ValueString()
	}
	if !(plan.Notrunning.IsNull() || plan.Notrunning.IsUnknown()) {
		body["notrunning"] = plan.Notrunning.ValueString()
	}
	if !(plan.Passthrough.IsNull() || plan.Passthrough.IsUnknown()) {
		body["passthrough"] = client.FormatBool(plan.Passthrough.ValueBool())
	}
	if !(plan.ResetTrafficCounters.IsNull() || plan.ResetTrafficCounters.IsUnknown()) {
		body["reset-traffic-counters"] = plan.ResetTrafficCounters.ValueString()
	}
	if !(plan.Slave.IsNull() || plan.Slave.IsUnknown()) {
		body["slave"] = client.FormatBool(plan.Slave.ValueBool())
	}
	if !(plan.Torch.IsNull() || plan.Torch.IsUnknown()) {
		body["torch"] = plan.Torch.ValueString()
	}
	if !(plan.TxRxBytes.IsNull() || plan.TxRxBytes.IsUnknown()) {
		body["tx-rx-bytes"] = plan.TxRxBytes.ValueString()
	}
	if !(plan.TxRxDrops.IsNull() || plan.TxRxDrops.IsUnknown()) {
		body["tx-rx-drops"] = plan.TxRxDrops.ValueString()
	}
	if !(plan.TxRxErrors.IsNull() || plan.TxRxErrors.IsUnknown()) {
		body["tx-rx-errors"] = plan.TxRxErrors.ValueString()
	}
	if !(plan.TxRxPacketRate.IsNull() || plan.TxRxPacketRate.IsUnknown()) {
		body["tx-rx-packet-rate"] = plan.TxRxPacketRate.ValueString()
	}
	if !(plan.TxRxPackets.IsNull() || plan.TxRxPackets.IsUnknown()) {
		body["tx-rx-packets"] = plan.TxRxPackets.ValueString()
	}
	if !(plan.TxRxRate.IsNull() || plan.TxRxRate.IsUnknown()) {
		body["tx-rx-rate"] = plan.TxRxRate.ValueString()
	}
	obj, err := c.Add(ctx, "/interface", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface failed", err.Error())
		return
	}
	interfaceApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface failed", err.Error())
		return
	}
	interfaceApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceModel
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
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.FpTxRxPacketRate.Equal(state.FpTxRxPacketRate) {
		body["fp-tx-rx-packet-rate"] = plan.FpTxRxPacketRate.ValueString()
	}
	if !plan.FpTxRxRate.Equal(state.FpTxRxRate) {
		body["fp-tx-rx-rate"] = plan.FpTxRxRate.ValueString()
	}
	if !plan.Inactive.Equal(state.Inactive) {
		body["inactive"] = client.FormatBool(plan.Inactive.ValueBool())
	}
	if !plan.MTU.Equal(state.MTU) {
		body["mtu"] = client.FormatInt64(plan.MTU.ValueInt64())
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Nodefname.Equal(state.Nodefname) {
		body["nodefname"] = plan.Nodefname.ValueString()
	}
	if !plan.Notrunning.Equal(state.Notrunning) {
		body["notrunning"] = plan.Notrunning.ValueString()
	}
	if !plan.Passthrough.Equal(state.Passthrough) {
		body["passthrough"] = client.FormatBool(plan.Passthrough.ValueBool())
	}
	if !plan.ResetTrafficCounters.Equal(state.ResetTrafficCounters) {
		body["reset-traffic-counters"] = plan.ResetTrafficCounters.ValueString()
	}
	if !plan.Slave.Equal(state.Slave) {
		body["slave"] = client.FormatBool(plan.Slave.ValueBool())
	}
	if !plan.Torch.Equal(state.Torch) {
		body["torch"] = plan.Torch.ValueString()
	}
	if !plan.TxRxBytes.Equal(state.TxRxBytes) {
		body["tx-rx-bytes"] = plan.TxRxBytes.ValueString()
	}
	if !plan.TxRxDrops.Equal(state.TxRxDrops) {
		body["tx-rx-drops"] = plan.TxRxDrops.ValueString()
	}
	if !plan.TxRxErrors.Equal(state.TxRxErrors) {
		body["tx-rx-errors"] = plan.TxRxErrors.ValueString()
	}
	if !plan.TxRxPacketRate.Equal(state.TxRxPacketRate) {
		body["tx-rx-packet-rate"] = plan.TxRxPacketRate.ValueString()
	}
	if !plan.TxRxPackets.Equal(state.TxRxPackets) {
		body["tx-rx-packets"] = plan.TxRxPackets.ValueString()
	}
	if !plan.TxRxRate.Equal(state.TxRxRate) {
		body["tx-rx-rate"] = plan.TxRxRate.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface failed", err.Error())
			return
		}
		interfaceApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface failed", err.Error())
	}
}

func (r *InterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/interface", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func interfaceApply(ctx context.Context, obj client.Object, m *InterfaceModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["actual-mtu"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.ActualMTU = types.Int64Value(n)
		} else {
			m.ActualMTU = types.Int64Null()
		}
	} else {
		m.ActualMTU = types.Int64Null()
	}
	if v, ok := obj["answer-time"]; ok {
		_ = v
		if v != "" {
			m.AnswerTime = types.StringValue(v)
		} else {
			m.AnswerTime = types.StringNull()
		}
	} else {
		m.AnswerTime = types.StringNull()
	}
	if v, ok := obj["caps"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Caps = types.Int64Value(n)
		} else {
			m.Caps = types.Int64Null()
		}
	} else {
		m.Caps = types.Int64Null()
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
	if v, ok := obj["default-name"]; ok {
		_ = v
		if v != "" {
			m.DefaultName = types.StringValue(v)
		} else {
			m.DefaultName = types.StringNull()
		}
	} else {
		m.DefaultName = types.StringNull()
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
	if v, ok := obj["dynamic"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else {
			m.Dynamic = types.BoolNull()
		}
	} else {
		m.Dynamic = types.BoolNull()
	}
	if v, ok := obj["fp-rps-drop"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.FpRpsDrop = types.Int64Value(n)
		} else {
			m.FpRpsDrop = types.Int64Null()
		}
	} else {
		m.FpRpsDrop = types.Int64Null()
	}
	if v, ok := obj["fp-rx-byte"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.FpRxByte = types.Int64Value(n)
		} else {
			m.FpRxByte = types.Int64Null()
		}
	} else {
		m.FpRxByte = types.Int64Null()
	}
	if v, ok := obj["fp-rx-packet"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.FpRxPacket = types.Int64Value(n)
		} else {
			m.FpRxPacket = types.Int64Null()
		}
	} else {
		m.FpRxPacket = types.Int64Null()
	}
	if v, ok := obj["fp-tx-byte"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.FpTxByte = types.Int64Value(n)
		} else {
			m.FpTxByte = types.Int64Null()
		}
	} else {
		m.FpTxByte = types.Int64Null()
	}
	if v, ok := obj["fp-tx-packet"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.FpTxPacket = types.Int64Value(n)
		} else {
			m.FpTxPacket = types.Int64Null()
		}
	} else {
		m.FpTxPacket = types.Int64Null()
	}
	if v, ok := obj["fp-tx-rx-packet-rate"]; ok {
		_ = v
		if v != "" {
			m.FpTxRxPacketRate = types.StringValue(v)
		} else {
			m.FpTxRxPacketRate = types.StringNull()
		}
	} else {
		m.FpTxRxPacketRate = types.StringNull()
	}
	if v, ok := obj["fp-tx-rx-rate"]; ok {
		_ = v
		if v != "" {
			m.FpTxRxRate = types.StringValue(v)
		} else {
			m.FpTxRxRate = types.StringNull()
		}
	} else {
		m.FpTxRxRate = types.StringNull()
	}
	if v, ok := obj["inactive"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Inactive = types.BoolValue(b)
		} else {
			m.Inactive = types.BoolNull()
		}
	} else {
		m.Inactive = types.BoolNull()
	}
	if v, ok := obj["l2-mtu"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.L2MTU = types.Int64Value(n)
		} else {
			m.L2MTU = types.Int64Null()
		}
	} else {
		m.L2MTU = types.Int64Null()
	}
	if v, ok := obj["last-link-down-time"]; ok {
		_ = v
		if v != "" {
			m.LastLinkDownTime = types.StringValue(v)
		} else {
			m.LastLinkDownTime = types.StringNull()
		}
	} else {
		m.LastLinkDownTime = types.StringNull()
	}
	if v, ok := obj["last-link-up-time"]; ok {
		_ = v
		if v != "" {
			m.LastLinkUpTime = types.StringValue(v)
		} else {
			m.LastLinkUpTime = types.StringNull()
		}
	} else {
		m.LastLinkUpTime = types.StringNull()
	}
	if v, ok := obj["link"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Link = types.Int64Value(n)
		} else {
			m.Link = types.Int64Null()
		}
	} else {
		m.Link = types.Int64Null()
	}
	if v, ok := obj["link-downs"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.LinkDowns = types.Int64Value(n)
		} else {
			m.LinkDowns = types.Int64Null()
		}
	} else {
		m.LinkDowns = types.Int64Null()
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
	if v, ok := obj["mtu"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MTU = types.Int64Value(n)
		} else {
			m.MTU = types.Int64Null()
		}
	} else {
		m.MTU = types.Int64Null()
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
	if v, ok := obj["nodefname"]; ok {
		_ = v
		if v != "" {
			m.Nodefname = types.StringValue(v)
		} else {
			m.Nodefname = types.StringNull()
		}
	} else {
		m.Nodefname = types.StringNull()
	}
	if v, ok := obj["notrunning"]; ok {
		_ = v
		if v != "" {
			m.Notrunning = types.StringValue(v)
		} else {
			m.Notrunning = types.StringNull()
		}
	} else {
		m.Notrunning = types.StringNull()
	}
	if v, ok := obj["passthrough"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Passthrough = types.BoolValue(b)
		} else {
			m.Passthrough = types.BoolNull()
		}
	} else {
		m.Passthrough = types.BoolNull()
	}
	if v, ok := obj["reset-traffic-counters"]; ok {
		_ = v
		if v != "" {
			m.ResetTrafficCounters = types.StringValue(v)
		} else {
			m.ResetTrafficCounters = types.StringNull()
		}
	} else {
		m.ResetTrafficCounters = types.StringNull()
	}
	if v, ok := obj["running"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Running = types.BoolValue(b)
		} else {
			m.Running = types.BoolNull()
		}
	} else {
		m.Running = types.BoolNull()
	}
	if v, ok := obj["rx-byte"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RxByte = types.Int64Value(n)
		} else {
			m.RxByte = types.Int64Null()
		}
	} else {
		m.RxByte = types.Int64Null()
	}
	if v, ok := obj["rx-drop"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RxDrop = types.Int64Value(n)
		} else {
			m.RxDrop = types.Int64Null()
		}
	} else {
		m.RxDrop = types.Int64Null()
	}
	if v, ok := obj["rx-error"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RxError = types.Int64Value(n)
		} else {
			m.RxError = types.Int64Null()
		}
	} else {
		m.RxError = types.Int64Null()
	}
	if v, ok := obj["rx-packet"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RxPacket = types.Int64Value(n)
		} else {
			m.RxPacket = types.Int64Null()
		}
	} else {
		m.RxPacket = types.Int64Null()
	}
	if v, ok := obj["slave"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Slave = types.BoolValue(b)
		} else {
			m.Slave = types.BoolNull()
		}
	} else {
		m.Slave = types.BoolNull()
	}
	if v, ok := obj["torch"]; ok {
		_ = v
		if v != "" {
			m.Torch = types.StringValue(v)
		} else {
			m.Torch = types.StringNull()
		}
	} else {
		m.Torch = types.StringNull()
	}
	if v, ok := obj["tx-byte"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TxByte = types.Int64Value(n)
		} else {
			m.TxByte = types.Int64Null()
		}
	} else {
		m.TxByte = types.Int64Null()
	}
	if v, ok := obj["tx-drop"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TxDrop = types.Int64Value(n)
		} else {
			m.TxDrop = types.Int64Null()
		}
	} else {
		m.TxDrop = types.Int64Null()
	}
	if v, ok := obj["tx-error"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TxError = types.Int64Value(n)
		} else {
			m.TxError = types.Int64Null()
		}
	} else {
		m.TxError = types.Int64Null()
	}
	if v, ok := obj["tx-packet"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TxPacket = types.Int64Value(n)
		} else {
			m.TxPacket = types.Int64Null()
		}
	} else {
		m.TxPacket = types.Int64Null()
	}
	if v, ok := obj["tx-queue-drop"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TxQueueDrop = types.Int64Value(n)
		} else {
			m.TxQueueDrop = types.Int64Null()
		}
	} else {
		m.TxQueueDrop = types.Int64Null()
	}
	if v, ok := obj["tx-queue-drops"]; ok {
		_ = v
		if v != "" {
			m.TxQueueDrops = types.StringValue(v)
		} else {
			m.TxQueueDrops = types.StringNull()
		}
	} else {
		m.TxQueueDrops = types.StringNull()
	}
	if v, ok := obj["tx-rx-bytes"]; ok {
		_ = v
		if v != "" {
			m.TxRxBytes = types.StringValue(v)
		} else {
			m.TxRxBytes = types.StringNull()
		}
	} else {
		m.TxRxBytes = types.StringNull()
	}
	if v, ok := obj["tx-rx-drops"]; ok {
		_ = v
		if v != "" {
			m.TxRxDrops = types.StringValue(v)
		} else {
			m.TxRxDrops = types.StringNull()
		}
	} else {
		m.TxRxDrops = types.StringNull()
	}
	if v, ok := obj["tx-rx-errors"]; ok {
		_ = v
		if v != "" {
			m.TxRxErrors = types.StringValue(v)
		} else {
			m.TxRxErrors = types.StringNull()
		}
	} else {
		m.TxRxErrors = types.StringNull()
	}
	if v, ok := obj["tx-rx-packet-rate"]; ok {
		_ = v
		if v != "" {
			m.TxRxPacketRate = types.StringValue(v)
		} else {
			m.TxRxPacketRate = types.StringNull()
		}
	} else {
		m.TxRxPacketRate = types.StringNull()
	}
	if v, ok := obj["tx-rx-packets"]; ok {
		_ = v
		if v != "" {
			m.TxRxPackets = types.StringValue(v)
		} else {
			m.TxRxPackets = types.StringNull()
		}
	} else {
		m.TxRxPackets = types.StringNull()
	}
	if v, ok := obj["tx-rx-rate"]; ok {
		_ = v
		if v != "" {
			m.TxRxRate = types.StringValue(v)
		} else {
			m.TxRxRate = types.StringNull()
		}
	} else {
		m.TxRxRate = types.StringNull()
	}
	if v, ok := obj["type"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Type = types.Int64Value(n)
		} else {
			m.Type = types.Int64Null()
		}
	} else {
		m.Type = types.Int64Null()
	}
	if v, ok := obj["vrf"]; ok {
		_ = v
		if v != "" {
			m.Vrf = types.StringValue(v)
		} else {
			m.Vrf = types.StringNull()
		}
	} else {
		m.Vrf = types.StringNull()
	}
}
