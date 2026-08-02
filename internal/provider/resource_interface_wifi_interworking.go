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
	_ resource.Resource                = &InterfaceWifiInterworkingResource{}
	_ resource.ResourceWithImportState = &InterfaceWifiInterworkingResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceWifiInterworkingResource struct {
	reg *client.Registry
}

type InterfaceWifiInterworkingModel struct {
	ID                     types.String `tfsdk:"id"`
	Hotspot20Dgaf          types.String `tfsdk:"hotspot20_dgaf"`
	X3gppInfo              types.String `tfsdk:"x3gpp_info"`
	X3gppInfoRaw           types.String `tfsdk:"x3gpp_info_raw"`
	AuthenticationTypes    types.String `tfsdk:"authentication_types"`
	Comment                types.String `tfsdk:"comment"`
	ConnectionCapabilities types.String `tfsdk:"connection_capabilities"`
	Dgaf                   types.String `tfsdk:"dgaf"`
	Disabled               types.Bool   `tfsdk:"disabled"`
	DomainNames            types.String `tfsdk:"domain_names"`
	Esr                    types.String `tfsdk:"esr"`
	Hessid                 types.String `tfsdk:"hessid"`
	Hotspot20              types.String `tfsdk:"hotspot20"`
	Internet               types.String `tfsdk:"internet"`
	Ipv4Availability       types.String `tfsdk:"ipv4_availability"`
	IPV6Availability       types.String `tfsdk:"ipv6_availability"`
	Name                   types.String `tfsdk:"name"`
	NetworkType            types.String `tfsdk:"network_type"`
	OperationalClasses     types.String `tfsdk:"operational_classes"`
	OperatorNames          types.String `tfsdk:"operator_names"`
	Realms                 types.String `tfsdk:"realms"`
	RealmsRaw              types.String `tfsdk:"realms_raw"`
	RoamingOis             types.String `tfsdk:"roaming_ois"`
	Uesa                   types.String `tfsdk:"uesa"`
	Venue                  types.String `tfsdk:"venue"`
	VenueNames             types.String `tfsdk:"venue_names"`
	WanAtCapacity          types.String `tfsdk:"wan_at_capacity"`
	WanDownlink            types.String `tfsdk:"wan_downlink"`
	WanDownlinkLoad        types.String `tfsdk:"wan_downlink_load"`
	WanMeasurementDuration types.String `tfsdk:"wan_measurement_duration"`
	WanStatus              types.String `tfsdk:"wan_status"`
	WanSymmetric           types.String `tfsdk:"wan_symmetric"`
	WanUplink              types.String `tfsdk:"wan_uplink"`
	WanUplinkLoad          types.String `tfsdk:"wan_uplink_load"`
	Router                 types.String `tfsdk:"router"`
}

func NewInterfaceWifiInterworkingResource() resource.Resource {
	return &InterfaceWifiInterworkingResource{}
}

func (r *InterfaceWifiInterworkingResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_wifi_interworking"
}

func (r *InterfaceWifiInterworkingResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceWifiInterworkingResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/wifi/interworking`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"hotspot20_dgaf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `hotspot20-dgaf`.",
			},
			"x3gpp_info": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"x3gpp_info_raw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"authentication_types": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"connection_capabilities": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dgaf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"domain_names": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"esr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hessid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hotspot20": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"internet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ipv4_availability": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ipv6_availability": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"network_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"operational_classes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"operator_names": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"realms": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"realms_raw": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"roaming_ois": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"uesa": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"venue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"venue_names": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_at_capacity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_downlink": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_downlink_load": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_measurement_duration": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_symmetric": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_uplink": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"wan_uplink_load": schema.StringAttribute{
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

func (r *InterfaceWifiInterworkingResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceWifiInterworkingModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.X3gppInfo.IsNull() || plan.X3gppInfo.IsUnknown()) {
		body["3gpp-info"] = plan.X3gppInfo.ValueString()
	}
	if !(plan.X3gppInfoRaw.IsNull() || plan.X3gppInfoRaw.IsUnknown()) {
		body["3gpp-info-raw"] = plan.X3gppInfoRaw.ValueString()
	}
	if !(plan.AuthenticationTypes.IsNull() || plan.AuthenticationTypes.IsUnknown()) {
		body["authentication-types"] = plan.AuthenticationTypes.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.ConnectionCapabilities.IsNull() || plan.ConnectionCapabilities.IsUnknown()) {
		body["connection-capabilities"] = plan.ConnectionCapabilities.ValueString()
	}
	if !(plan.Dgaf.IsNull() || plan.Dgaf.IsUnknown()) {
		body["dgaf"] = plan.Dgaf.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.DomainNames.IsNull() || plan.DomainNames.IsUnknown()) {
		body["domain-names"] = plan.DomainNames.ValueString()
	}
	if !(plan.Esr.IsNull() || plan.Esr.IsUnknown()) {
		body["esr"] = plan.Esr.ValueString()
	}
	if !(plan.Hessid.IsNull() || plan.Hessid.IsUnknown()) {
		body["hessid"] = plan.Hessid.ValueString()
	}
	if !(plan.Hotspot20.IsNull() || plan.Hotspot20.IsUnknown()) {
		body["hotspot20"] = plan.Hotspot20.ValueString()
	}
	if !(plan.Internet.IsNull() || plan.Internet.IsUnknown()) {
		body["internet"] = plan.Internet.ValueString()
	}
	if !(plan.Ipv4Availability.IsNull() || plan.Ipv4Availability.IsUnknown()) {
		body["ipv4-availability"] = plan.Ipv4Availability.ValueString()
	}
	if !(plan.IPV6Availability.IsNull() || plan.IPV6Availability.IsUnknown()) {
		body["ipv6-availability"] = plan.IPV6Availability.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.NetworkType.IsNull() || plan.NetworkType.IsUnknown()) {
		body["network-type"] = plan.NetworkType.ValueString()
	}
	if !(plan.OperationalClasses.IsNull() || plan.OperationalClasses.IsUnknown()) {
		body["operational-classes"] = plan.OperationalClasses.ValueString()
	}
	if !(plan.OperatorNames.IsNull() || plan.OperatorNames.IsUnknown()) {
		body["operator-names"] = plan.OperatorNames.ValueString()
	}
	if !(plan.Realms.IsNull() || plan.Realms.IsUnknown()) {
		body["realms"] = plan.Realms.ValueString()
	}
	if !(plan.RealmsRaw.IsNull() || plan.RealmsRaw.IsUnknown()) {
		body["realms-raw"] = plan.RealmsRaw.ValueString()
	}
	if !(plan.RoamingOis.IsNull() || plan.RoamingOis.IsUnknown()) {
		body["roaming-ois"] = plan.RoamingOis.ValueString()
	}
	if !(plan.Uesa.IsNull() || plan.Uesa.IsUnknown()) {
		body["uesa"] = plan.Uesa.ValueString()
	}
	if !(plan.Venue.IsNull() || plan.Venue.IsUnknown()) {
		body["venue"] = plan.Venue.ValueString()
	}
	if !(plan.VenueNames.IsNull() || plan.VenueNames.IsUnknown()) {
		body["venue-names"] = plan.VenueNames.ValueString()
	}
	if !(plan.WanAtCapacity.IsNull() || plan.WanAtCapacity.IsUnknown()) {
		body["wan-at-capacity"] = plan.WanAtCapacity.ValueString()
	}
	if !(plan.WanDownlink.IsNull() || plan.WanDownlink.IsUnknown()) {
		body["wan-downlink"] = plan.WanDownlink.ValueString()
	}
	if !(plan.WanDownlinkLoad.IsNull() || plan.WanDownlinkLoad.IsUnknown()) {
		body["wan-downlink-load"] = plan.WanDownlinkLoad.ValueString()
	}
	if !(plan.WanMeasurementDuration.IsNull() || plan.WanMeasurementDuration.IsUnknown()) {
		body["wan-measurement-duration"] = plan.WanMeasurementDuration.ValueString()
	}
	if !(plan.WanStatus.IsNull() || plan.WanStatus.IsUnknown()) {
		body["wan-status"] = plan.WanStatus.ValueString()
	}
	if !(plan.WanSymmetric.IsNull() || plan.WanSymmetric.IsUnknown()) {
		body["wan-symmetric"] = plan.WanSymmetric.ValueString()
	}
	if !(plan.WanUplink.IsNull() || plan.WanUplink.IsUnknown()) {
		body["wan-uplink"] = plan.WanUplink.ValueString()
	}
	if !(plan.WanUplinkLoad.IsNull() || plan.WanUplinkLoad.IsUnknown()) {
		body["wan-uplink-load"] = plan.WanUplinkLoad.ValueString()
	}
	if !(plan.Hotspot20Dgaf.IsNull() || plan.Hotspot20Dgaf.IsUnknown()) {
		body["hotspot20-dgaf"] = plan.Hotspot20Dgaf.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/wifi/interworking", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/wifi/interworking failed", err.Error())
		return
	}
	interfaceWifiInterworkingApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiInterworkingResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceWifiInterworkingModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/wifi/interworking", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/wifi/interworking failed", err.Error())
		return
	}
	interfaceWifiInterworkingApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceWifiInterworkingResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceWifiInterworkingModel
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
	if !plan.X3gppInfo.Equal(state.X3gppInfo) && !plan.X3gppInfo.IsUnknown() {
		body["3gpp-info"] = plan.X3gppInfo.ValueString()
	}
	if !plan.X3gppInfoRaw.Equal(state.X3gppInfoRaw) && !plan.X3gppInfoRaw.IsUnknown() {
		body["3gpp-info-raw"] = plan.X3gppInfoRaw.ValueString()
	}
	if !plan.AuthenticationTypes.Equal(state.AuthenticationTypes) && !plan.AuthenticationTypes.IsUnknown() {
		body["authentication-types"] = plan.AuthenticationTypes.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.ConnectionCapabilities.Equal(state.ConnectionCapabilities) && !plan.ConnectionCapabilities.IsUnknown() {
		body["connection-capabilities"] = plan.ConnectionCapabilities.ValueString()
	}
	if !plan.Dgaf.Equal(state.Dgaf) && !plan.Dgaf.IsUnknown() {
		body["dgaf"] = plan.Dgaf.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DomainNames.Equal(state.DomainNames) && !plan.DomainNames.IsUnknown() {
		body["domain-names"] = plan.DomainNames.ValueString()
	}
	if !plan.Esr.Equal(state.Esr) && !plan.Esr.IsUnknown() {
		body["esr"] = plan.Esr.ValueString()
	}
	if !plan.Hessid.Equal(state.Hessid) && !plan.Hessid.IsUnknown() {
		body["hessid"] = plan.Hessid.ValueString()
	}
	if !plan.Hotspot20.Equal(state.Hotspot20) && !plan.Hotspot20.IsUnknown() {
		body["hotspot20"] = plan.Hotspot20.ValueString()
	}
	if !plan.Internet.Equal(state.Internet) && !plan.Internet.IsUnknown() {
		body["internet"] = plan.Internet.ValueString()
	}
	if !plan.Ipv4Availability.Equal(state.Ipv4Availability) && !plan.Ipv4Availability.IsUnknown() {
		body["ipv4-availability"] = plan.Ipv4Availability.ValueString()
	}
	if !plan.IPV6Availability.Equal(state.IPV6Availability) && !plan.IPV6Availability.IsUnknown() {
		body["ipv6-availability"] = plan.IPV6Availability.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.NetworkType.Equal(state.NetworkType) && !plan.NetworkType.IsUnknown() {
		body["network-type"] = plan.NetworkType.ValueString()
	}
	if !plan.OperationalClasses.Equal(state.OperationalClasses) && !plan.OperationalClasses.IsUnknown() {
		body["operational-classes"] = plan.OperationalClasses.ValueString()
	}
	if !plan.OperatorNames.Equal(state.OperatorNames) && !plan.OperatorNames.IsUnknown() {
		body["operator-names"] = plan.OperatorNames.ValueString()
	}
	if !plan.Realms.Equal(state.Realms) && !plan.Realms.IsUnknown() {
		body["realms"] = plan.Realms.ValueString()
	}
	if !plan.RealmsRaw.Equal(state.RealmsRaw) && !plan.RealmsRaw.IsUnknown() {
		body["realms-raw"] = plan.RealmsRaw.ValueString()
	}
	if !plan.RoamingOis.Equal(state.RoamingOis) && !plan.RoamingOis.IsUnknown() {
		body["roaming-ois"] = plan.RoamingOis.ValueString()
	}
	if !plan.Uesa.Equal(state.Uesa) && !plan.Uesa.IsUnknown() {
		body["uesa"] = plan.Uesa.ValueString()
	}
	if !plan.Venue.Equal(state.Venue) && !plan.Venue.IsUnknown() {
		body["venue"] = plan.Venue.ValueString()
	}
	if !plan.VenueNames.Equal(state.VenueNames) && !plan.VenueNames.IsUnknown() {
		body["venue-names"] = plan.VenueNames.ValueString()
	}
	if !plan.WanAtCapacity.Equal(state.WanAtCapacity) && !plan.WanAtCapacity.IsUnknown() {
		body["wan-at-capacity"] = plan.WanAtCapacity.ValueString()
	}
	if !plan.WanDownlink.Equal(state.WanDownlink) && !plan.WanDownlink.IsUnknown() {
		body["wan-downlink"] = plan.WanDownlink.ValueString()
	}
	if !plan.WanDownlinkLoad.Equal(state.WanDownlinkLoad) && !plan.WanDownlinkLoad.IsUnknown() {
		body["wan-downlink-load"] = plan.WanDownlinkLoad.ValueString()
	}
	if !plan.WanMeasurementDuration.Equal(state.WanMeasurementDuration) && !plan.WanMeasurementDuration.IsUnknown() {
		body["wan-measurement-duration"] = plan.WanMeasurementDuration.ValueString()
	}
	if !plan.WanStatus.Equal(state.WanStatus) && !plan.WanStatus.IsUnknown() {
		body["wan-status"] = plan.WanStatus.ValueString()
	}
	if !plan.WanSymmetric.Equal(state.WanSymmetric) && !plan.WanSymmetric.IsUnknown() {
		body["wan-symmetric"] = plan.WanSymmetric.ValueString()
	}
	if !plan.WanUplink.Equal(state.WanUplink) && !plan.WanUplink.IsUnknown() {
		body["wan-uplink"] = plan.WanUplink.ValueString()
	}
	if !plan.WanUplinkLoad.Equal(state.WanUplinkLoad) && !plan.WanUplinkLoad.IsUnknown() {
		body["wan-uplink-load"] = plan.WanUplinkLoad.ValueString()
	}
	if !plan.Hotspot20Dgaf.Equal(state.Hotspot20Dgaf) && !plan.Hotspot20Dgaf.IsUnknown() {
		body["hotspot20-dgaf"] = plan.Hotspot20Dgaf.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/wifi/interworking", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/wifi/interworking failed", err.Error())
			return
		}
		interfaceWifiInterworkingApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiInterworkingResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceWifiInterworkingModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/wifi/interworking", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/wifi/interworking failed", err.Error())
	}
}

func (r *InterfaceWifiInterworkingResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceWifiInterworkingLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/wifi/interworking matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceWifiInterworkingLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceWifiInterworkingLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/wifi/interworking", id)
}

func interfaceWifiInterworkingApply(ctx context.Context, obj client.Object, m *InterfaceWifiInterworkingModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["hotspot20-dgaf"]; ok && v != "" {
		m.Hotspot20Dgaf = types.StringValue(v)
	} else {
		m.Hotspot20Dgaf = types.StringNull()
	}
	if v, ok := obj["3gpp-info"]; ok {
		if v != "" {
			m.X3gppInfo = types.StringValue(v)
		} else {
			m.X3gppInfo = types.StringNull()
		}
	}
	if v, ok := obj["3gpp-info-raw"]; ok {
		if v != "" {
			m.X3gppInfoRaw = types.StringValue(v)
		} else {
			m.X3gppInfoRaw = types.StringNull()
		}
	}
	if v, ok := obj["authentication-types"]; ok {
		if v != "" {
			m.AuthenticationTypes = types.StringValue(v)
		} else {
			m.AuthenticationTypes = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["connection-capabilities"]; ok {
		if v != "" {
			m.ConnectionCapabilities = types.StringValue(v)
		} else {
			m.ConnectionCapabilities = types.StringNull()
		}
	}
	if v, ok := obj["dgaf"]; ok {
		if v != "" {
			m.Dgaf = types.StringValue(v)
		} else {
			m.Dgaf = types.StringNull()
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
	if v, ok := obj["domain-names"]; ok {
		if v != "" {
			m.DomainNames = types.StringValue(v)
		} else {
			m.DomainNames = types.StringNull()
		}
	}
	if v, ok := obj["esr"]; ok {
		if v != "" {
			m.Esr = types.StringValue(v)
		} else {
			m.Esr = types.StringNull()
		}
	}
	if v, ok := obj["hessid"]; ok {
		if v != "" {
			m.Hessid = types.StringValue(v)
		} else {
			m.Hessid = types.StringNull()
		}
	}
	if v, ok := obj["hotspot20"]; ok {
		if v != "" {
			m.Hotspot20 = types.StringValue(v)
		} else {
			m.Hotspot20 = types.StringNull()
		}
	}
	if v, ok := obj["internet"]; ok {
		if v != "" {
			m.Internet = types.StringValue(v)
		} else {
			m.Internet = types.StringNull()
		}
	}
	if v, ok := obj["ipv4-availability"]; ok {
		if v != "" {
			m.Ipv4Availability = types.StringValue(v)
		} else {
			m.Ipv4Availability = types.StringNull()
		}
	}
	if v, ok := obj["ipv6-availability"]; ok {
		if v != "" {
			m.IPV6Availability = types.StringValue(v)
		} else {
			m.IPV6Availability = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["network-type"]; ok {
		if v != "" {
			m.NetworkType = types.StringValue(v)
		} else {
			m.NetworkType = types.StringNull()
		}
	}
	if v, ok := obj["operational-classes"]; ok {
		if v != "" {
			m.OperationalClasses = types.StringValue(v)
		} else {
			m.OperationalClasses = types.StringNull()
		}
	}
	if v, ok := obj["operator-names"]; ok {
		if v != "" {
			m.OperatorNames = types.StringValue(v)
		} else {
			m.OperatorNames = types.StringNull()
		}
	}
	if v, ok := obj["realms"]; ok {
		if v != "" {
			m.Realms = types.StringValue(v)
		} else {
			m.Realms = types.StringNull()
		}
	}
	if v, ok := obj["realms-raw"]; ok {
		if v != "" {
			m.RealmsRaw = types.StringValue(v)
		} else {
			m.RealmsRaw = types.StringNull()
		}
	}
	if v, ok := obj["roaming-ois"]; ok {
		if v != "" {
			m.RoamingOis = types.StringValue(v)
		} else {
			m.RoamingOis = types.StringNull()
		}
	}
	if v, ok := obj["uesa"]; ok {
		if v != "" {
			m.Uesa = types.StringValue(v)
		} else {
			m.Uesa = types.StringNull()
		}
	}
	if v, ok := obj["venue"]; ok {
		if v != "" {
			m.Venue = types.StringValue(v)
		} else {
			m.Venue = types.StringNull()
		}
	}
	if v, ok := obj["venue-names"]; ok {
		if v != "" {
			m.VenueNames = types.StringValue(v)
		} else {
			m.VenueNames = types.StringNull()
		}
	}
	if v, ok := obj["wan-at-capacity"]; ok {
		if v != "" {
			m.WanAtCapacity = types.StringValue(v)
		} else {
			m.WanAtCapacity = types.StringNull()
		}
	}
	if v, ok := obj["wan-downlink"]; ok {
		if v != "" {
			m.WanDownlink = types.StringValue(v)
		} else {
			m.WanDownlink = types.StringNull()
		}
	}
	if v, ok := obj["wan-downlink-load"]; ok {
		if v != "" {
			m.WanDownlinkLoad = types.StringValue(v)
		} else {
			m.WanDownlinkLoad = types.StringNull()
		}
	}
	if v, ok := obj["wan-measurement-duration"]; ok {
		if v != "" {
			m.WanMeasurementDuration = types.StringValue(v)
		} else {
			m.WanMeasurementDuration = types.StringNull()
		}
	}
	if v, ok := obj["wan-status"]; ok {
		if v != "" {
			m.WanStatus = types.StringValue(v)
		} else {
			m.WanStatus = types.StringNull()
		}
	}
	if v, ok := obj["wan-symmetric"]; ok {
		if v != "" {
			m.WanSymmetric = types.StringValue(v)
		} else {
			m.WanSymmetric = types.StringNull()
		}
	}
	if v, ok := obj["wan-uplink"]; ok {
		if v != "" {
			m.WanUplink = types.StringValue(v)
		} else {
			m.WanUplink = types.StringNull()
		}
	}
	if v, ok := obj["wan-uplink-load"]; ok {
		if v != "" {
			m.WanUplinkLoad = types.StringValue(v)
		} else {
			m.WanUplinkLoad = types.StringNull()
		}
	}
}
