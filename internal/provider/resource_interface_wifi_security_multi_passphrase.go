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
	_ resource.Resource                = &InterfaceWifiSecurityMultiPassphraseResource{}
	_ resource.ResourceWithImportState = &InterfaceWifiSecurityMultiPassphraseResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceWifiSecurityMultiPassphraseResource struct {
	reg *client.Registry
}

type InterfaceWifiSecurityMultiPassphraseModel struct {
	ID         types.String `tfsdk:"id"`
	Comment    types.String `tfsdk:"comment"`
	Disabled   types.Bool   `tfsdk:"disabled"`
	Expired    types.Bool   `tfsdk:"expired"`
	Expires    types.String `tfsdk:"expires"`
	Group      types.String `tfsdk:"group"`
	Isolation  types.Bool   `tfsdk:"isolation"`
	Passphrase types.String `tfsdk:"passphrase"`
	VLANID     types.String `tfsdk:"vlan_id"`
	Router     types.String `tfsdk:"router"`
}

func NewInterfaceWifiSecurityMultiPassphraseResource() resource.Resource {
	return &InterfaceWifiSecurityMultiPassphraseResource{}
}

func (r *InterfaceWifiSecurityMultiPassphraseResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_wifi_security_multi_passphrase"
}

func (r *InterfaceWifiSecurityMultiPassphraseResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceWifiSecurityMultiPassphraseResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/wifi/security/multi-passphrase`. Each row is one passphrase " +
			"entry in a PPSK group; assign the group name to `multi_passphrase_group` on " +
			"`routeros_interface_wifi`, `routeros_interface_wifi_configuration`, `routeros_interface_wifi_security` " +
			"or `routeros_interface_wifi_access_list` to put it to use.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"expired": schema.BoolAttribute{
				Computed:    true,
				Description: "Whether this passphrase's `expires` date/time has passed.",
			},
			"expires": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Description: "Expiration date and time for this passphrase; doesn't affect the rest of the " +
					"group. Once reached, existing clients using it are disconnected and new clients can't use " +
					"it. Leave unset for no expiry.",
			},
			"group": schema.StringAttribute{
				Required: true,
				Description: "The PPSK group name. Assigning it to a security profile or an access list " +
					"enables use of every passphrase defined under it.",
			},
			"isolation": schema.BoolAttribute{
				Optional: true,
				Computed: true,
				Description: "Whether a client using this passphrase is isolated from other clients on the " +
					"AP: traffic from an isolated client is not forwarded to other clients, and unicast traffic " +
					"from a non-isolated client is not forwarded to an isolated one.",
			},
			"passphrase": schema.StringAttribute{
				Optional:  true,
				Computed:  true,
				Sensitive: true,
				Description: "The PSK passphrase. Multiple entries may share a passphrase. Not compatible " +
					"with WPA3-PSK.",
			},
			"vlan_id": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Description: "VLAN ID assigned to clients using this passphrase. Only supported on " +
					"wifi-qcom interfaces; a wifi-qcom-ac AP will refuse a client whose passphrase carries one.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceWifiSecurityMultiPassphraseResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceWifiSecurityMultiPassphraseModel
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
	if !(plan.Expires.IsNull() || plan.Expires.IsUnknown()) {
		body["expires"] = plan.Expires.ValueString()
	}
	if !(plan.Group.IsNull() || plan.Group.IsUnknown()) {
		body["group"] = plan.Group.ValueString()
	}
	if !(plan.Isolation.IsNull() || plan.Isolation.IsUnknown()) {
		body["isolation"] = client.FormatBool(plan.Isolation.ValueBool())
	}
	if !(plan.Passphrase.IsNull() || plan.Passphrase.IsUnknown()) {
		body["passphrase"] = plan.Passphrase.ValueString()
	}
	if !(plan.VLANID.IsNull() || plan.VLANID.IsUnknown()) {
		body["vlan-id"] = plan.VLANID.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/wifi/security/multi-passphrase", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/wifi/security/multi-passphrase failed", err.Error())
		return
	}
	interfaceWifiSecurityMultiPassphraseApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiSecurityMultiPassphraseResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceWifiSecurityMultiPassphraseModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/wifi/security/multi-passphrase", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/wifi/security/multi-passphrase failed", err.Error())
		return
	}
	interfaceWifiSecurityMultiPassphraseApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceWifiSecurityMultiPassphraseResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceWifiSecurityMultiPassphraseModel
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
	if !plan.Expires.Equal(state.Expires) && !plan.Expires.IsUnknown() {
		body["expires"] = plan.Expires.ValueString()
	}
	if !plan.Group.Equal(state.Group) && !plan.Group.IsUnknown() {
		body["group"] = plan.Group.ValueString()
	}
	if !plan.Isolation.Equal(state.Isolation) && !plan.Isolation.IsUnknown() {
		body["isolation"] = client.FormatBool(plan.Isolation.ValueBool())
	}
	if !plan.Passphrase.Equal(state.Passphrase) && !plan.Passphrase.IsUnknown() {
		body["passphrase"] = plan.Passphrase.ValueString()
	}
	if !plan.VLANID.Equal(state.VLANID) && !plan.VLANID.IsUnknown() {
		body["vlan-id"] = plan.VLANID.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/wifi/security/multi-passphrase", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/wifi/security/multi-passphrase failed", err.Error())
			return
		}
		interfaceWifiSecurityMultiPassphraseApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiSecurityMultiPassphraseResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceWifiSecurityMultiPassphraseModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/wifi/security/multi-passphrase", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/wifi/security/multi-passphrase failed", err.Error())
	}
}

func (r *InterfaceWifiSecurityMultiPassphraseResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceWifiSecurityMultiPassphraseLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/wifi/security/multi-passphrase matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceWifiSecurityMultiPassphraseLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceWifiSecurityMultiPassphraseLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/wifi/security/multi-passphrase", id)
}

func interfaceWifiSecurityMultiPassphraseApply(ctx context.Context, obj client.Object, m *InterfaceWifiSecurityMultiPassphraseModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
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
	if v, ok := obj["expired"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Expired = types.BoolValue(b)
		} else {
			m.Expired = types.BoolNull()
		}
	}
	if v, ok := obj["expires"]; ok {
		if v != "" {
			m.Expires = types.StringValue(v)
		} else {
			m.Expires = types.StringNull()
		}
	}
	if v, ok := obj["group"]; ok {
		if v != "" {
			m.Group = types.StringValue(v)
		} else {
			m.Group = types.StringNull()
		}
	}
	if v, ok := obj["isolation"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Isolation = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Isolation = types.BoolValue(false)
		} else {
			m.Isolation = types.BoolNull()
		}
	}
	// The device returns the passphrase in plain text on read (confirmed against a live
	// 7.23.3 router), so this isn't a scrub-on-read case like some other sensitive fields.
	if v, ok := obj["passphrase"]; ok {
		if v != "" {
			m.Passphrase = types.StringValue(v)
		} else {
			m.Passphrase = types.StringNull()
		}
	} else if m.Passphrase.IsUnknown() {
		m.Passphrase = types.StringNull()
	}
	if v, ok := obj["vlan-id"]; ok {
		if v != "" {
			m.VLANID = types.StringValue(v)
		} else {
			m.VLANID = types.StringNull()
		}
	}
}
