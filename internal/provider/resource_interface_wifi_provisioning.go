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
	_ resource.Resource                = &InterfaceWifiProvisioningResource{}
	_ resource.ResourceWithImportState = &InterfaceWifiProvisioningResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceWifiProvisioningResource struct {
	reg *client.Registry
}

type InterfaceWifiProvisioningModel struct {
	ID                  types.String `tfsdk:"id"`
	Action              types.String `tfsdk:"action"`
	AddressRanges       types.String `tfsdk:"address_ranges"`
	Comment             types.String `tfsdk:"comment"`
	CommonNameRegexp    types.String `tfsdk:"common_name_regexp"`
	Disabled            types.Bool   `tfsdk:"disabled"`
	IdentityRegexp      types.String `tfsdk:"identity_regexp"`
	MasterConfiguration types.String `tfsdk:"master_configuration"`
	MultiLinkMode       types.String `tfsdk:"multi_link_mode"`
	NameFormat          types.String `tfsdk:"name_format"`
	RadioMAC            types.String `tfsdk:"radio_mac"`
	SlaveConfigurations types.String `tfsdk:"slave_configurations"`
	SlaveNameFormat     types.String `tfsdk:"slave_name_format"`
	SupportedBands      types.String `tfsdk:"supported_bands"`
	SupportedHwCaps     types.String `tfsdk:"supported_hw_caps"`
	Router              types.String `tfsdk:"router"`
}

func NewInterfaceWifiProvisioningResource() resource.Resource {
	return &InterfaceWifiProvisioningResource{}
}

func (r *InterfaceWifiProvisioningResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_wifi_provisioning"
}

func (r *InterfaceWifiProvisioningResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *InterfaceWifiProvisioningResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/wifi/provisioning`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"action": schema.StringAttribute{
				Required:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"none", "create-enabled", "create-disabled", "create-dynamic-enabled"}...)},
			},
			"address_ranges": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"common_name_regexp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"identity_regexp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"master_configuration": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"multi_link_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"radio_mac": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"slave_configurations": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"slave_name_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"supported_bands": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"supported_hw_caps": schema.StringAttribute{
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

func (r *InterfaceWifiProvisioningResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceWifiProvisioningModel
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
	if !(plan.AddressRanges.IsNull() || plan.AddressRanges.IsUnknown()) {
		body["address-ranges"] = plan.AddressRanges.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.CommonNameRegexp.IsNull() || plan.CommonNameRegexp.IsUnknown()) {
		body["common-name-regexp"] = plan.CommonNameRegexp.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.IdentityRegexp.IsNull() || plan.IdentityRegexp.IsUnknown()) {
		body["identity-regexp"] = plan.IdentityRegexp.ValueString()
	}
	if !(plan.MasterConfiguration.IsNull() || plan.MasterConfiguration.IsUnknown()) {
		body["master-configuration"] = plan.MasterConfiguration.ValueString()
	}
	if !(plan.MultiLinkMode.IsNull() || plan.MultiLinkMode.IsUnknown()) {
		body["multi-link-mode"] = plan.MultiLinkMode.ValueString()
	}
	if !(plan.NameFormat.IsNull() || plan.NameFormat.IsUnknown()) {
		body["name-format"] = plan.NameFormat.ValueString()
	}
	if !(plan.RadioMAC.IsNull() || plan.RadioMAC.IsUnknown()) {
		body["radio-mac"] = plan.RadioMAC.ValueString()
	}
	if !(plan.SlaveConfigurations.IsNull() || plan.SlaveConfigurations.IsUnknown()) {
		body["slave-configurations"] = plan.SlaveConfigurations.ValueString()
	}
	if !(plan.SlaveNameFormat.IsNull() || plan.SlaveNameFormat.IsUnknown()) {
		body["slave-name-format"] = plan.SlaveNameFormat.ValueString()
	}
	if !(plan.SupportedBands.IsNull() || plan.SupportedBands.IsUnknown()) {
		body["supported-bands"] = plan.SupportedBands.ValueString()
	}
	if !(plan.SupportedHwCaps.IsNull() || plan.SupportedHwCaps.IsUnknown()) {
		body["supported-hw-caps"] = plan.SupportedHwCaps.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/wifi/provisioning", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/wifi/provisioning failed", err.Error())
		return
	}
	interfaceWifiProvisioningApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiProvisioningResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceWifiProvisioningModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/wifi/provisioning", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/wifi/provisioning failed", err.Error())
		return
	}
	interfaceWifiProvisioningApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceWifiProvisioningResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceWifiProvisioningModel
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
	if !plan.Action.Equal(state.Action) {
		body["action"] = plan.Action.ValueString()
	}
	if !plan.AddressRanges.Equal(state.AddressRanges) {
		body["address-ranges"] = plan.AddressRanges.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.CommonNameRegexp.Equal(state.CommonNameRegexp) {
		body["common-name-regexp"] = plan.CommonNameRegexp.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.IdentityRegexp.Equal(state.IdentityRegexp) {
		body["identity-regexp"] = plan.IdentityRegexp.ValueString()
	}
	if !plan.MasterConfiguration.Equal(state.MasterConfiguration) {
		body["master-configuration"] = plan.MasterConfiguration.ValueString()
	}
	if !plan.MultiLinkMode.Equal(state.MultiLinkMode) {
		body["multi-link-mode"] = plan.MultiLinkMode.ValueString()
	}
	if !plan.NameFormat.Equal(state.NameFormat) {
		body["name-format"] = plan.NameFormat.ValueString()
	}
	if !plan.RadioMAC.Equal(state.RadioMAC) {
		body["radio-mac"] = plan.RadioMAC.ValueString()
	}
	if !plan.SlaveConfigurations.Equal(state.SlaveConfigurations) {
		body["slave-configurations"] = plan.SlaveConfigurations.ValueString()
	}
	if !plan.SlaveNameFormat.Equal(state.SlaveNameFormat) {
		body["slave-name-format"] = plan.SlaveNameFormat.ValueString()
	}
	if !plan.SupportedBands.Equal(state.SupportedBands) {
		body["supported-bands"] = plan.SupportedBands.ValueString()
	}
	if !plan.SupportedHwCaps.Equal(state.SupportedHwCaps) {
		body["supported-hw-caps"] = plan.SupportedHwCaps.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/wifi/provisioning", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/wifi/provisioning failed", err.Error())
			return
		}
		interfaceWifiProvisioningApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiProvisioningResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceWifiProvisioningModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/wifi/provisioning", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/wifi/provisioning failed", err.Error())
	}
}

func (r *InterfaceWifiProvisioningResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceWifiProvisioningLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/wifi/provisioning matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceWifiProvisioningLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceWifiProvisioningLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/wifi/provisioning", id)
}

func interfaceWifiProvisioningApply(ctx context.Context, obj client.Object, m *InterfaceWifiProvisioningModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["action"]; ok {
		_ = v
		if v != "" {
			m.Action = types.StringValue(v)
		} else {
			m.Action = types.StringNull()
		}
	} else {
		m.Action = types.StringNull()
	}
	if v, ok := obj["address-ranges"]; ok {
		_ = v
		if v != "" {
			m.AddressRanges = types.StringValue(v)
		} else {
			m.AddressRanges = types.StringNull()
		}
	} else {
		m.AddressRanges = types.StringNull()
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
	if v, ok := obj["common-name-regexp"]; ok {
		_ = v
		if v != "" {
			m.CommonNameRegexp = types.StringValue(v)
		} else {
			m.CommonNameRegexp = types.StringNull()
		}
	} else {
		m.CommonNameRegexp = types.StringNull()
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
	if v, ok := obj["identity-regexp"]; ok {
		_ = v
		if v != "" {
			m.IdentityRegexp = types.StringValue(v)
		} else {
			m.IdentityRegexp = types.StringNull()
		}
	} else {
		m.IdentityRegexp = types.StringNull()
	}
	if v, ok := obj["master-configuration"]; ok {
		_ = v
		if v != "" {
			m.MasterConfiguration = types.StringValue(v)
		} else {
			m.MasterConfiguration = types.StringNull()
		}
	} else {
		m.MasterConfiguration = types.StringNull()
	}
	if v, ok := obj["multi-link-mode"]; ok {
		_ = v
		if v != "" {
			m.MultiLinkMode = types.StringValue(v)
		} else {
			m.MultiLinkMode = types.StringNull()
		}
	} else {
		m.MultiLinkMode = types.StringNull()
	}
	if v, ok := obj["name-format"]; ok {
		_ = v
		if v != "" {
			m.NameFormat = types.StringValue(v)
		} else {
			m.NameFormat = types.StringNull()
		}
	} else {
		m.NameFormat = types.StringNull()
	}
	if v, ok := obj["radio-mac"]; ok {
		_ = v
		if v != "" {
			m.RadioMAC = types.StringValue(v)
		} else {
			m.RadioMAC = types.StringNull()
		}
	} else {
		m.RadioMAC = types.StringNull()
	}
	if v, ok := obj["slave-configurations"]; ok {
		_ = v
		if v != "" {
			m.SlaveConfigurations = types.StringValue(v)
		} else {
			m.SlaveConfigurations = types.StringNull()
		}
	} else {
		m.SlaveConfigurations = types.StringNull()
	}
	if v, ok := obj["slave-name-format"]; ok {
		_ = v
		if v != "" {
			m.SlaveNameFormat = types.StringValue(v)
		} else {
			m.SlaveNameFormat = types.StringNull()
		}
	} else {
		m.SlaveNameFormat = types.StringNull()
	}
	if v, ok := obj["supported-bands"]; ok {
		_ = v
		if v != "" {
			m.SupportedBands = types.StringValue(v)
		} else {
			m.SupportedBands = types.StringNull()
		}
	} else {
		m.SupportedBands = types.StringNull()
	}
	if v, ok := obj["supported-hw-caps"]; ok {
		_ = v
		if v != "" {
			m.SupportedHwCaps = types.StringValue(v)
		} else {
			m.SupportedHwCaps = types.StringNull()
		}
	} else {
		m.SupportedHwCaps = types.StringNull()
	}
}
