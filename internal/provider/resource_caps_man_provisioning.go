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
	_ resource.Resource                = &CapsManProvisioningResource{}
	_ resource.ResourceWithImportState = &CapsManProvisioningResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type CapsManProvisioningResource struct {
	reg *client.Registry
}

type CapsManProvisioningModel struct {
	ID                  types.String `tfsdk:"id"`
	Action              types.String `tfsdk:"action"`
	Comment             types.String `tfsdk:"comment"`
	CommonNameRegexp    types.String `tfsdk:"common_name_regexp"`
	Disabled            types.Bool   `tfsdk:"disabled"`
	HwSupportedModes    types.String `tfsdk:"hw_supported_modes"`
	IdentityRegexp      types.String `tfsdk:"identity_regexp"`
	IPAddressRanges     types.String `tfsdk:"ip_address_ranges"`
	MasterConfiguration types.String `tfsdk:"master_configuration"`
	NameFormat          types.String `tfsdk:"name_format"`
	NamePrefix          types.String `tfsdk:"name_prefix"`
	RadioMAC            types.String `tfsdk:"radio_mac"`
	SlaveConfiguration  types.String `tfsdk:"slave_configuration"`
	SlaveConfigurations types.String `tfsdk:"slave_configurations"`
	Router              types.String `tfsdk:"router"`
}

func NewCapsManProvisioningResource() resource.Resource { return &CapsManProvisioningResource{} }

func (r *CapsManProvisioningResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_caps_man_provisioning"
}

func (r *CapsManProvisioningResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *CapsManProvisioningResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/caps-man/provisioning`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"none", "create-enabled", "create-disabled", "create-dynamic-enabled"}...)},
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
			"hw_supported_modes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"identity_regexp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ip_address_ranges": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"master_configuration": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"cap", "prefix", "identity", "prefix-identity"}...)},
			},
			"name_prefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"radio_mac": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsMAC()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeMAC()},
			},
			"slave_configuration": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"slave_configurations": schema.StringAttribute{
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

func (r *CapsManProvisioningResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CapsManProvisioningModel
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
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.CommonNameRegexp.IsNull() || plan.CommonNameRegexp.IsUnknown()) {
		body["common-name-regexp"] = plan.CommonNameRegexp.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.HwSupportedModes.IsNull() || plan.HwSupportedModes.IsUnknown()) {
		body["hw-supported-modes"] = plan.HwSupportedModes.ValueString()
	}
	if !(plan.IdentityRegexp.IsNull() || plan.IdentityRegexp.IsUnknown()) {
		body["identity-regexp"] = plan.IdentityRegexp.ValueString()
	}
	if !(plan.IPAddressRanges.IsNull() || plan.IPAddressRanges.IsUnknown()) {
		body["ip-address-ranges"] = plan.IPAddressRanges.ValueString()
	}
	if !(plan.MasterConfiguration.IsNull() || plan.MasterConfiguration.IsUnknown()) {
		body["master-configuration"] = plan.MasterConfiguration.ValueString()
	}
	if !(plan.NameFormat.IsNull() || plan.NameFormat.IsUnknown()) {
		body["name-format"] = plan.NameFormat.ValueString()
	}
	if !(plan.NamePrefix.IsNull() || plan.NamePrefix.IsUnknown()) {
		body["name-prefix"] = plan.NamePrefix.ValueString()
	}
	if !(plan.RadioMAC.IsNull() || plan.RadioMAC.IsUnknown()) {
		body["radio-mac"] = plan.RadioMAC.ValueString()
	}
	if !(plan.SlaveConfiguration.IsNull() || plan.SlaveConfiguration.IsUnknown()) {
		body["slave-configuration"] = plan.SlaveConfiguration.ValueString()
	}
	if !(plan.SlaveConfigurations.IsNull() || plan.SlaveConfigurations.IsUnknown()) {
		body["slave-configurations"] = plan.SlaveConfigurations.ValueString()
	}
	obj, err := c.Add(ctx, "/caps-man/provisioning", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /caps-man/provisioning failed", err.Error())
		return
	}
	capsManProvisioningApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CapsManProvisioningResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CapsManProvisioningModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/caps-man/provisioning", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /caps-man/provisioning failed", err.Error())
		return
	}
	capsManProvisioningApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CapsManProvisioningResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CapsManProvisioningModel
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
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.CommonNameRegexp.Equal(state.CommonNameRegexp) {
		body["common-name-regexp"] = plan.CommonNameRegexp.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.HwSupportedModes.Equal(state.HwSupportedModes) {
		body["hw-supported-modes"] = plan.HwSupportedModes.ValueString()
	}
	if !plan.IdentityRegexp.Equal(state.IdentityRegexp) {
		body["identity-regexp"] = plan.IdentityRegexp.ValueString()
	}
	if !plan.IPAddressRanges.Equal(state.IPAddressRanges) {
		body["ip-address-ranges"] = plan.IPAddressRanges.ValueString()
	}
	if !plan.MasterConfiguration.Equal(state.MasterConfiguration) {
		body["master-configuration"] = plan.MasterConfiguration.ValueString()
	}
	if !plan.NameFormat.Equal(state.NameFormat) {
		body["name-format"] = plan.NameFormat.ValueString()
	}
	if !plan.NamePrefix.Equal(state.NamePrefix) {
		body["name-prefix"] = plan.NamePrefix.ValueString()
	}
	if !plan.RadioMAC.Equal(state.RadioMAC) {
		body["radio-mac"] = plan.RadioMAC.ValueString()
	}
	if !plan.SlaveConfiguration.Equal(state.SlaveConfiguration) {
		body["slave-configuration"] = plan.SlaveConfiguration.ValueString()
	}
	if !plan.SlaveConfigurations.Equal(state.SlaveConfigurations) {
		body["slave-configurations"] = plan.SlaveConfigurations.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/caps-man/provisioning", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /caps-man/provisioning failed", err.Error())
			return
		}
		capsManProvisioningApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CapsManProvisioningResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CapsManProvisioningModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/caps-man/provisioning", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /caps-man/provisioning failed", err.Error())
	}
}

func (r *CapsManProvisioningResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := capsManProvisioningLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /caps-man/provisioning matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// capsManProvisioningLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func capsManProvisioningLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/caps-man/provisioning", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func capsManProvisioningApply(ctx context.Context, obj client.Object, m *CapsManProvisioningModel) {
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
	if v, ok := obj["hw-supported-modes"]; ok {
		_ = v
		if v != "" {
			m.HwSupportedModes = types.StringValue(v)
		} else {
			m.HwSupportedModes = types.StringNull()
		}
	} else {
		m.HwSupportedModes = types.StringNull()
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
	if v, ok := obj["ip-address-ranges"]; ok {
		_ = v
		if v != "" {
			m.IPAddressRanges = types.StringValue(v)
		} else {
			m.IPAddressRanges = types.StringNull()
		}
	} else {
		m.IPAddressRanges = types.StringNull()
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
	if v, ok := obj["name-prefix"]; ok {
		_ = v
		if v != "" {
			m.NamePrefix = types.StringValue(v)
		} else {
			m.NamePrefix = types.StringNull()
		}
	} else {
		m.NamePrefix = types.StringNull()
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
	if v, ok := obj["slave-configuration"]; ok {
		_ = v
		if v != "" {
			m.SlaveConfiguration = types.StringValue(v)
		} else {
			m.SlaveConfiguration = types.StringNull()
		}
	} else {
		m.SlaveConfiguration = types.StringNull()
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
}
