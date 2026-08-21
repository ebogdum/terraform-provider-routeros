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
	_ resource.Resource                = &InterfaceWifiAccessListResource{}
	_ resource.ResourceWithImportState = &InterfaceWifiAccessListResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceWifiAccessListResource struct {
	reg *client.Registry
}

type InterfaceWifiAccessListModel struct {
	ID                    types.String `tfsdk:"id"`
	Days                  types.String `tfsdk:"days"`
	Action                types.String `tfsdk:"action"`
	AllowSignalOutOfRange types.String `tfsdk:"allow_signal_out_of_range"`
	ClientIsolation       types.String `tfsdk:"client_isolation"`
	Comment               types.String `tfsdk:"comment"`
	Disabled              types.Bool   `tfsdk:"disabled"`
	Interface             types.String `tfsdk:"interface"`
	LastLoggedIn          types.String `tfsdk:"last_logged_in"`
	LastLoggedOut         types.String `tfsdk:"last_logged_out"`
	MACAddress            macValue     `tfsdk:"mac_address"`
	MACAddressMask        types.String `tfsdk:"mac_address_mask"`
	MultiPassphraseGroup  types.String `tfsdk:"multi_passphrase_group"`
	Passphrase            types.String `tfsdk:"passphrase"`
	RADIUSAccounting      types.String `tfsdk:"radius_accounting"`
	SignalRange           types.String `tfsdk:"signal_range"`
	SsidRegexp            types.String `tfsdk:"ssid_regexp"`
	Time                  types.String `tfsdk:"time"`
	TimesMatched          types.String `tfsdk:"times_matched"`
	VLANID                types.String `tfsdk:"vlan_id"`
	Weekdays              types.String `tfsdk:"weekdays"`
	Router                types.String `tfsdk:"router"`
}

func NewInterfaceWifiAccessListResource() resource.Resource {
	return &InterfaceWifiAccessListResource{}
}

func (r *InterfaceWifiAccessListResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_wifi_access_list"
}

func (r *InterfaceWifiAccessListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceWifiAccessListResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/wifi/access-list`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"days": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `days`.",
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"allow_signal_out_of_range": schema.StringAttribute{
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
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"last_logged_in": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"last_logged_out": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"mac_address": schema.StringAttribute{
				CustomType:  macType{},
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsMAC()},
			},
			"mac_address_mask": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"multi_passphrase_group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"passphrase": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"radius_accounting": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"signal_range": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ssid_regexp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"times_matched": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"vlan_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"weekdays": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *InterfaceWifiAccessListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceWifiAccessListModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Action.IsNull() || plan.Action.IsUnknown()) {
		body["action"] = plan.Action.ValueString()
	}
	if !(plan.AllowSignalOutOfRange.IsNull() || plan.AllowSignalOutOfRange.IsUnknown()) {
		body["allow-signal-out-of-range"] = plan.AllowSignalOutOfRange.ValueString()
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
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.MACAddress.IsNull() || plan.MACAddress.IsUnknown()) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !(plan.MACAddressMask.IsNull() || plan.MACAddressMask.IsUnknown()) {
		body["mac-address-mask"] = plan.MACAddressMask.ValueString()
	}
	if !(plan.MultiPassphraseGroup.IsNull() || plan.MultiPassphraseGroup.IsUnknown()) {
		body["multi-passphrase-group"] = plan.MultiPassphraseGroup.ValueString()
	}
	if !(plan.Passphrase.IsNull() || plan.Passphrase.IsUnknown()) {
		body["passphrase"] = plan.Passphrase.ValueString()
	}
	if !(plan.RADIUSAccounting.IsNull() || plan.RADIUSAccounting.IsUnknown()) {
		body["radius-accounting"] = plan.RADIUSAccounting.ValueString()
	}
	if !(plan.SignalRange.IsNull() || plan.SignalRange.IsUnknown()) {
		body["signal-range"] = plan.SignalRange.ValueString()
	}
	if !(plan.SsidRegexp.IsNull() || plan.SsidRegexp.IsUnknown()) {
		body["ssid-regexp"] = plan.SsidRegexp.ValueString()
	}
	if !(plan.Time.IsNull() || plan.Time.IsUnknown()) {
		body["time"] = plan.Time.ValueString()
	}
	if !(plan.VLANID.IsNull() || plan.VLANID.IsUnknown()) {
		body["vlan-id"] = plan.VLANID.ValueString()
	}
	if !(plan.Weekdays.IsNull() || plan.Weekdays.IsUnknown()) {
		body["days"] = plan.Weekdays.ValueString()
	}
	if !(plan.Days.IsNull() || plan.Days.IsUnknown()) {
		body["days"] = plan.Days.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/wifi/access-list", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/wifi/access-list failed", err.Error())
		return
	}
	interfaceWifiAccessListApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiAccessListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceWifiAccessListModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/wifi/access-list", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/wifi/access-list failed", err.Error())
		return
	}
	interfaceWifiAccessListApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceWifiAccessListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceWifiAccessListModel
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
	if !plan.Action.Equal(state.Action) && !plan.Action.IsUnknown() {
		body["action"] = plan.Action.ValueString()
	}
	if !plan.AllowSignalOutOfRange.Equal(state.AllowSignalOutOfRange) && !plan.AllowSignalOutOfRange.IsUnknown() {
		body["allow-signal-out-of-range"] = plan.AllowSignalOutOfRange.ValueString()
	}
	if !plan.ClientIsolation.Equal(state.ClientIsolation) && !plan.ClientIsolation.IsUnknown() {
		body["client-isolation"] = plan.ClientIsolation.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.MACAddress.Equal(state.MACAddress) && !plan.MACAddress.IsUnknown() {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.MACAddressMask.Equal(state.MACAddressMask) && !plan.MACAddressMask.IsUnknown() {
		body["mac-address-mask"] = plan.MACAddressMask.ValueString()
	}
	if !plan.MultiPassphraseGroup.Equal(state.MultiPassphraseGroup) && !plan.MultiPassphraseGroup.IsUnknown() {
		body["multi-passphrase-group"] = plan.MultiPassphraseGroup.ValueString()
	}
	if !plan.Passphrase.Equal(state.Passphrase) && !plan.Passphrase.IsUnknown() {
		body["passphrase"] = plan.Passphrase.ValueString()
	}
	if !plan.RADIUSAccounting.Equal(state.RADIUSAccounting) && !plan.RADIUSAccounting.IsUnknown() {
		body["radius-accounting"] = plan.RADIUSAccounting.ValueString()
	}
	if !plan.SignalRange.Equal(state.SignalRange) && !plan.SignalRange.IsUnknown() {
		body["signal-range"] = plan.SignalRange.ValueString()
	}
	if !plan.SsidRegexp.Equal(state.SsidRegexp) && !plan.SsidRegexp.IsUnknown() {
		body["ssid-regexp"] = plan.SsidRegexp.ValueString()
	}
	if !plan.Time.Equal(state.Time) && !plan.Time.IsUnknown() {
		body["time"] = plan.Time.ValueString()
	}
	if !plan.VLANID.Equal(state.VLANID) && !plan.VLANID.IsUnknown() {
		body["vlan-id"] = plan.VLANID.ValueString()
	}
	if !plan.Weekdays.Equal(state.Weekdays) && !plan.Weekdays.IsUnknown() {
		body["days"] = plan.Weekdays.ValueString()
	}
	if !plan.Days.Equal(state.Days) && !plan.Days.IsUnknown() {
		body["days"] = plan.Days.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/wifi/access-list", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/wifi/access-list failed", err.Error())
			return
		}
		interfaceWifiAccessListApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiAccessListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceWifiAccessListModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/wifi/access-list", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/wifi/access-list failed", err.Error())
	}
}

func (r *InterfaceWifiAccessListResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceWifiAccessListLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/wifi/access-list matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceWifiAccessListLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceWifiAccessListLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/wifi/access-list", id)
}

func interfaceWifiAccessListApply(ctx context.Context, obj client.Object, m *InterfaceWifiAccessListModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["days"]; ok && v != "" {
		m.Days = types.StringValue(v)
	} else {
		m.Days = types.StringNull()
	}
	if v, ok := obj["action"]; ok {
		if v != "" {
			m.Action = types.StringValue(v)
		} else {
			m.Action = types.StringNull()
		}
	}
	if v, ok := obj["allow-signal-out-of-range"]; ok {
		if v != "" {
			m.AllowSignalOutOfRange = types.StringValue(v)
		} else {
			m.AllowSignalOutOfRange = types.StringNull()
		}
	}
	if v, ok := obj["client-isolation"]; ok {
		if v != "" {
			m.ClientIsolation = types.StringValue(v)
		} else {
			m.ClientIsolation = types.StringNull()
		}
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
		} else if strings.TrimSpace(v) == "" {
			m.Disabled = types.BoolValue(true)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["interface"]; ok {
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	}
	if v, ok := obj["last-logged-in"]; ok {
		if v != "" {
			m.LastLoggedIn = types.StringValue(v)
		} else {
			m.LastLoggedIn = types.StringNull()
		}
	}
	if v, ok := obj["last-logged-out"]; ok {
		if v != "" {
			m.LastLoggedOut = types.StringValue(v)
		} else {
			m.LastLoggedOut = types.StringNull()
		}
	}
	if v, ok := obj["mac-address"]; ok {
		if v != "" {
			m.MACAddress = newMACValue(v)
		} else {
			m.MACAddress = newMACNull()
		}
	}
	if v, ok := obj["mac-address-mask"]; ok {
		if v != "" {
			m.MACAddressMask = types.StringValue(v)
		} else {
			m.MACAddressMask = types.StringNull()
		}
	}
	if v, ok := obj["multi-passphrase-group"]; ok {
		if v != "" {
			m.MultiPassphraseGroup = types.StringValue(v)
		} else {
			m.MultiPassphraseGroup = types.StringNull()
		}
	}
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.Passphrase already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["passphrase"]; ok && v != "" {
		_ = v
		if v != "" {
			m.Passphrase = types.StringValue(v)
		} else {
			m.Passphrase = types.StringNull()
		}
	} else if m.Passphrase.IsUnknown() {
		m.Passphrase = types.StringNull()
	}
	if v, ok := obj["radius-accounting"]; ok {
		if v != "" {
			m.RADIUSAccounting = types.StringValue(v)
		} else {
			m.RADIUSAccounting = types.StringNull()
		}
	}
	if v, ok := obj["signal-range"]; ok {
		if v != "" {
			m.SignalRange = types.StringValue(v)
		} else {
			m.SignalRange = types.StringNull()
		}
	}
	if v, ok := obj["ssid-regexp"]; ok {
		if v != "" {
			m.SsidRegexp = types.StringValue(v)
		} else {
			m.SsidRegexp = types.StringNull()
		}
	}
	if v, ok := obj["time"]; ok {
		if v != "" {
			m.Time = types.StringValue(v)
		} else {
			m.Time = types.StringNull()
		}
	}
	if v, ok := obj["times-matched"]; ok {
		if v != "" {
			m.TimesMatched = types.StringValue(v)
		} else {
			m.TimesMatched = types.StringNull()
		}
	}
	if v, ok := obj["vlan-id"]; ok {
		if v != "" {
			m.VLANID = types.StringValue(v)
		} else {
			m.VLANID = types.StringNull()
		}
	}
	if v, ok := obj["days"]; ok {
		if v != "" {
			m.Weekdays = types.StringValue(v)
		} else {
			m.Weekdays = types.StringNull()
		}
	}
}
