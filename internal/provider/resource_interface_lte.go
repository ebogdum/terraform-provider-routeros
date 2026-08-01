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
	_ resource.Resource                = &InterfaceLteResource{}
	_ resource.ResourceWithImportState = &InterfaceLteResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceLteResource struct {
	reg *client.Registry
}

type InterfaceLteModel struct {
	ID           types.String `tfsdk:"id"`
	SmsRead      types.String `tfsdk:"sms_read"`
	SmsProtocol  types.String `tfsdk:"sms_protocol"`
	Pin          types.String `tfsdk:"pin"`
	Operator     types.String `tfsdk:"operator"`
	NrBand       types.String `tfsdk:"nr_band"`
	NetworkMode  types.String `tfsdk:"network_mode"`
	Name         types.String `tfsdk:"name"`
	Mtu          types.String `tfsdk:"mtu"`
	ModemInit    types.String `tfsdk:"modem_init"`
	Band         types.String `tfsdk:"band"`
	ApnProfiles  types.String `tfsdk:"apn_profiles"`
	AllowRoaming types.String `tfsdk:"allow_roaming"`
	Comment      types.String `tfsdk:"comment"`
	Disabled     types.Bool   `tfsdk:"disabled"`
	Router       types.String `tfsdk:"router"`
}

func NewInterfaceLteResource() resource.Resource { return &InterfaceLteResource{} }

func (r *InterfaceLteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_lte"
}

func (r *InterfaceLteResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceLteResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "LTE interfaces are physical-device backed; skipped on virtual devices.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"sms_read": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `sms-read`.",
			},
			"sms_protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `sms-protocol`.",
			},
			"pin": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `pin`.",
			},
			"operator": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `operator`.",
			},
			"nr_band": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `nr-band`.",
			},
			"network_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `network-mode`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `name`.",
			},
			"mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mtu`.",
			},
			"modem_init": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `modem-init`.",
			},
			"band": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `band`.",
			},
			"apn_profiles": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `apn-profiles`.",
			},
			"allow_roaming": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `allow-roaming`.",
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

func (r *InterfaceLteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceLteModel
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
	if !(plan.AllowRoaming.IsNull() || plan.AllowRoaming.IsUnknown()) {
		body["allow-roaming"] = plan.AllowRoaming.ValueString()
	}
	if !(plan.ApnProfiles.IsNull() || plan.ApnProfiles.IsUnknown()) {
		body["apn-profiles"] = plan.ApnProfiles.ValueString()
	}
	if !(plan.Band.IsNull() || plan.Band.IsUnknown()) {
		body["band"] = plan.Band.ValueString()
	}
	if !(plan.ModemInit.IsNull() || plan.ModemInit.IsUnknown()) {
		body["modem-init"] = plan.ModemInit.ValueString()
	}
	if !(plan.Mtu.IsNull() || plan.Mtu.IsUnknown()) {
		body["mtu"] = plan.Mtu.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.NetworkMode.IsNull() || plan.NetworkMode.IsUnknown()) {
		body["network-mode"] = plan.NetworkMode.ValueString()
	}
	if !(plan.NrBand.IsNull() || plan.NrBand.IsUnknown()) {
		body["nr-band"] = plan.NrBand.ValueString()
	}
	if !(plan.Operator.IsNull() || plan.Operator.IsUnknown()) {
		body["operator"] = plan.Operator.ValueString()
	}
	if !(plan.Pin.IsNull() || plan.Pin.IsUnknown()) {
		body["pin"] = plan.Pin.ValueString()
	}
	if !(plan.SmsProtocol.IsNull() || plan.SmsProtocol.IsUnknown()) {
		body["sms-protocol"] = plan.SmsProtocol.ValueString()
	}
	if !(plan.SmsRead.IsNull() || plan.SmsRead.IsUnknown()) {
		body["sms-read"] = plan.SmsRead.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/lte", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/lte failed", err.Error())
		return
	}
	interfaceLteApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceLteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceLteModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/lte", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/lte failed", err.Error())
		return
	}
	interfaceLteApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceLteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceLteModel
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
	if !plan.AllowRoaming.Equal(state.AllowRoaming) && !plan.AllowRoaming.IsUnknown() {
		body["allow-roaming"] = plan.AllowRoaming.ValueString()
	}
	if !plan.ApnProfiles.Equal(state.ApnProfiles) && !plan.ApnProfiles.IsUnknown() {
		body["apn-profiles"] = plan.ApnProfiles.ValueString()
	}
	if !plan.Band.Equal(state.Band) && !plan.Band.IsUnknown() {
		body["band"] = plan.Band.ValueString()
	}
	if !plan.ModemInit.Equal(state.ModemInit) && !plan.ModemInit.IsUnknown() {
		body["modem-init"] = plan.ModemInit.ValueString()
	}
	if !plan.Mtu.Equal(state.Mtu) && !plan.Mtu.IsUnknown() {
		body["mtu"] = plan.Mtu.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.NetworkMode.Equal(state.NetworkMode) && !plan.NetworkMode.IsUnknown() {
		body["network-mode"] = plan.NetworkMode.ValueString()
	}
	if !plan.NrBand.Equal(state.NrBand) && !plan.NrBand.IsUnknown() {
		body["nr-band"] = plan.NrBand.ValueString()
	}
	if !plan.Operator.Equal(state.Operator) && !plan.Operator.IsUnknown() {
		body["operator"] = plan.Operator.ValueString()
	}
	if !plan.Pin.Equal(state.Pin) && !plan.Pin.IsUnknown() {
		body["pin"] = plan.Pin.ValueString()
	}
	if !plan.SmsProtocol.Equal(state.SmsProtocol) && !plan.SmsProtocol.IsUnknown() {
		body["sms-protocol"] = plan.SmsProtocol.ValueString()
	}
	if !plan.SmsRead.Equal(state.SmsRead) && !plan.SmsRead.IsUnknown() {
		body["sms-read"] = plan.SmsRead.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/lte", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/lte failed", err.Error())
			return
		}
		interfaceLteApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceLteResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceLteModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/lte", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/lte failed", err.Error())
	}
}

func (r *InterfaceLteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceLteLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/lte matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceLteLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceLteLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/lte", id)
}

func interfaceLteApply(ctx context.Context, obj client.Object, m *InterfaceLteModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["sms-read"]; ok && v != "" {
		m.SmsRead = types.StringValue(v)
	} else {
		m.SmsRead = types.StringNull()
	}
	if v, ok := obj["sms-protocol"]; ok && v != "" {
		m.SmsProtocol = types.StringValue(v)
	} else {
		m.SmsProtocol = types.StringNull()
	}
	if v, ok := obj["pin"]; ok && v != "" {
		m.Pin = types.StringValue(v)
	} else {
		m.Pin = types.StringNull()
	}
	if v, ok := obj["operator"]; ok && v != "" {
		m.Operator = types.StringValue(v)
	} else {
		m.Operator = types.StringNull()
	}
	if v, ok := obj["nr-band"]; ok && v != "" {
		m.NrBand = types.StringValue(v)
	} else {
		m.NrBand = types.StringNull()
	}
	if v, ok := obj["network-mode"]; ok && v != "" {
		m.NetworkMode = types.StringValue(v)
	} else {
		m.NetworkMode = types.StringNull()
	}
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["mtu"]; ok && v != "" {
		m.Mtu = types.StringValue(v)
	} else {
		m.Mtu = types.StringNull()
	}
	if v, ok := obj["modem-init"]; ok && v != "" {
		m.ModemInit = types.StringValue(v)
	} else {
		m.ModemInit = types.StringNull()
	}
	if v, ok := obj["band"]; ok && v != "" {
		m.Band = types.StringValue(v)
	} else {
		m.Band = types.StringNull()
	}
	if v, ok := obj["apn-profiles"]; ok && v != "" {
		m.ApnProfiles = types.StringValue(v)
	} else {
		m.ApnProfiles = types.StringNull()
	}
	if v, ok := obj["allow-roaming"]; ok && v != "" {
		m.AllowRoaming = types.StringValue(v)
	} else {
		m.AllowRoaming = types.StringNull()
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
}
