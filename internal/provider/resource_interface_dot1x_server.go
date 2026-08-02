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
	_ resource.Resource                = &InterfaceDot1xServerResource{}
	_ resource.ResourceWithImportState = &InterfaceDot1xServerResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceDot1xServerResource struct {
	reg *client.Registry
}

type InterfaceDot1xServerModel struct {
	ID               types.String `tfsdk:"id"`
	Accounting       types.Bool   `tfsdk:"accounting"`
	AuthTimeout      types.String `tfsdk:"auth_timeout"`
	AuthTypes        types.String `tfsdk:"auth_types"`
	Comment          types.String `tfsdk:"comment"`
	Disabled         types.Bool   `tfsdk:"disabled"`
	GuestVLANID      types.String `tfsdk:"guest_vlan_id"`
	Interface        types.String `tfsdk:"interface"`
	InterimUpdate    types.String `tfsdk:"interim_update"`
	Invalid          types.Bool   `tfsdk:"invalid"`
	MAC              types.String `tfsdk:"mac"`
	MACAuthMode      types.String `tfsdk:"mac_auth_mode"`
	RADIUSMACFormat  types.String `tfsdk:"radius_mac_format"`
	ReauthTimeout    types.String `tfsdk:"reauth_timeout"`
	RejectVLANID     types.String `tfsdk:"reject_vlan_id"`
	RetransTimeout   types.String `tfsdk:"retrans_timeout"`
	ServerFailVLANID types.String `tfsdk:"server_fail_vlan_id"`
	Router           types.String `tfsdk:"router"`
}

func NewInterfaceDot1xServerResource() resource.Resource { return &InterfaceDot1xServerResource{} }

func (r *InterfaceDot1xServerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_dot1x_server"
}

func (r *InterfaceDot1xServerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceDot1xServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "802.1X server attaches to a specific Ethernet interface; values vary per device. Skipped.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"accounting": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"auth_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"auth_types": schema.StringAttribute{
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
			"guest_vlan_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interim_update": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"mac": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"mac_auth_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"mac-as-username", "mac-as-username-and-password"}...)},
			},
			"radius_mac_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"xx:xx:xx:xx:xx:xx", "xx-xx-xx-xx-xx-xx", "xxxxxxxxxxxx"}...)},
			},
			"reauth_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"reject_vlan_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"retrans_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"server_fail_vlan_id": schema.StringAttribute{
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

func (r *InterfaceDot1xServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceDot1xServerModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Accounting.IsNull() || plan.Accounting.IsUnknown()) {
		body["accounting"] = client.FormatBool(plan.Accounting.ValueBool())
	}
	if !(plan.AuthTimeout.IsNull() || plan.AuthTimeout.IsUnknown()) {
		body["auth-timeout"] = plan.AuthTimeout.ValueString()
	}
	if !(plan.AuthTypes.IsNull() || plan.AuthTypes.IsUnknown()) {
		body["auth-types"] = plan.AuthTypes.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.GuestVLANID.IsNull() || plan.GuestVLANID.IsUnknown()) {
		body["guest-vlan-id"] = plan.GuestVLANID.ValueString()
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.InterimUpdate.IsNull() || plan.InterimUpdate.IsUnknown()) {
		body["interim-update"] = plan.InterimUpdate.ValueString()
	}
	if !(plan.MACAuthMode.IsNull() || plan.MACAuthMode.IsUnknown()) {
		body["mac-auth-mode"] = plan.MACAuthMode.ValueString()
	}
	if !(plan.RADIUSMACFormat.IsNull() || plan.RADIUSMACFormat.IsUnknown()) {
		body["radius-mac-format"] = plan.RADIUSMACFormat.ValueString()
	}
	if !(plan.ReauthTimeout.IsNull() || plan.ReauthTimeout.IsUnknown()) {
		body["reauth-timeout"] = plan.ReauthTimeout.ValueString()
	}
	if !(plan.RejectVLANID.IsNull() || plan.RejectVLANID.IsUnknown()) {
		body["reject-vlan-id"] = plan.RejectVLANID.ValueString()
	}
	if !(plan.RetransTimeout.IsNull() || plan.RetransTimeout.IsUnknown()) {
		body["retrans-timeout"] = plan.RetransTimeout.ValueString()
	}
	if !(plan.ServerFailVLANID.IsNull() || plan.ServerFailVLANID.IsUnknown()) {
		body["server-fail-vlan-id"] = plan.ServerFailVLANID.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/dot1x/server", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/dot1x/server failed", err.Error())
		return
	}
	interfaceDot1xServerApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceDot1xServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceDot1xServerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/dot1x/server", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/dot1x/server failed", err.Error())
		return
	}
	interfaceDot1xServerApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceDot1xServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceDot1xServerModel
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
	if !plan.Accounting.Equal(state.Accounting) && !plan.Accounting.IsUnknown() {
		body["accounting"] = client.FormatBool(plan.Accounting.ValueBool())
	}
	if !plan.AuthTimeout.Equal(state.AuthTimeout) && !plan.AuthTimeout.IsUnknown() {
		body["auth-timeout"] = plan.AuthTimeout.ValueString()
	}
	if !plan.AuthTypes.Equal(state.AuthTypes) && !plan.AuthTypes.IsUnknown() {
		body["auth-types"] = plan.AuthTypes.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.GuestVLANID.Equal(state.GuestVLANID) && !plan.GuestVLANID.IsUnknown() {
		body["guest-vlan-id"] = plan.GuestVLANID.ValueString()
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.InterimUpdate.Equal(state.InterimUpdate) && !plan.InterimUpdate.IsUnknown() {
		body["interim-update"] = plan.InterimUpdate.ValueString()
	}
	if !plan.MACAuthMode.Equal(state.MACAuthMode) && !plan.MACAuthMode.IsUnknown() {
		body["mac-auth-mode"] = plan.MACAuthMode.ValueString()
	}
	if !plan.RADIUSMACFormat.Equal(state.RADIUSMACFormat) && !plan.RADIUSMACFormat.IsUnknown() {
		body["radius-mac-format"] = plan.RADIUSMACFormat.ValueString()
	}
	if !plan.ReauthTimeout.Equal(state.ReauthTimeout) && !plan.ReauthTimeout.IsUnknown() {
		body["reauth-timeout"] = plan.ReauthTimeout.ValueString()
	}
	if !plan.RejectVLANID.Equal(state.RejectVLANID) && !plan.RejectVLANID.IsUnknown() {
		body["reject-vlan-id"] = plan.RejectVLANID.ValueString()
	}
	if !plan.RetransTimeout.Equal(state.RetransTimeout) && !plan.RetransTimeout.IsUnknown() {
		body["retrans-timeout"] = plan.RetransTimeout.ValueString()
	}
	if !plan.ServerFailVLANID.Equal(state.ServerFailVLANID) && !plan.ServerFailVLANID.IsUnknown() {
		body["server-fail-vlan-id"] = plan.ServerFailVLANID.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/dot1x/server", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/dot1x/server failed", err.Error())
			return
		}
		interfaceDot1xServerApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceDot1xServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceDot1xServerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/dot1x/server", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/dot1x/server failed", err.Error())
	}
}

func (r *InterfaceDot1xServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceDot1xServerLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/dot1x/server matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceDot1xServerLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceDot1xServerLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/dot1x/server", id)
}

func interfaceDot1xServerApply(ctx context.Context, obj client.Object, m *InterfaceDot1xServerModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["accounting"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Accounting = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Accounting = types.BoolValue(true)
		} else {
			m.Accounting = types.BoolNull()
		}
	}
	if v, ok := obj["auth-timeout"]; ok {
		if v != "" {
			m.AuthTimeout = types.StringValue(v)
		} else {
			m.AuthTimeout = types.StringNull()
		}
	}
	if v, ok := obj["auth-types"]; ok {
		if v != "" {
			m.AuthTypes = types.StringValue(v)
		} else {
			m.AuthTypes = types.StringNull()
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
	if v, ok := obj["guest-vlan-id"]; ok {
		if v != "" {
			m.GuestVLANID = types.StringValue(v)
		} else {
			m.GuestVLANID = types.StringNull()
		}
	}
	if v, ok := obj["interface"]; ok {
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	}
	if v, ok := obj["interim-update"]; ok {
		if v != "" {
			m.InterimUpdate = types.StringValue(v)
		} else {
			m.InterimUpdate = types.StringNull()
		}
	}
	if v, ok := obj["invalid"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Invalid = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Invalid = types.BoolValue(true)
		} else {
			m.Invalid = types.BoolNull()
		}
	}
	if v, ok := obj["mac"]; ok {
		if v != "" {
			m.MAC = types.StringValue(v)
		} else {
			m.MAC = types.StringNull()
		}
	}
	if v, ok := obj["mac-auth-mode"]; ok {
		if v != "" {
			m.MACAuthMode = types.StringValue(v)
		} else {
			m.MACAuthMode = types.StringNull()
		}
	}
	if v, ok := obj["radius-mac-format"]; ok {
		if v != "" {
			m.RADIUSMACFormat = types.StringValue(v)
		} else {
			m.RADIUSMACFormat = types.StringNull()
		}
	}
	if v, ok := obj["reauth-timeout"]; ok {
		if v != "" {
			m.ReauthTimeout = types.StringValue(v)
		} else {
			m.ReauthTimeout = types.StringNull()
		}
	}
	if v, ok := obj["reject-vlan-id"]; ok {
		if v != "" {
			m.RejectVLANID = types.StringValue(v)
		} else {
			m.RejectVLANID = types.StringNull()
		}
	}
	if v, ok := obj["retrans-timeout"]; ok {
		if v != "" {
			m.RetransTimeout = types.StringValue(v)
		} else {
			m.RetransTimeout = types.StringNull()
		}
	}
	if v, ok := obj["server-fail-vlan-id"]; ok {
		if v != "" {
			m.ServerFailVLANID = types.StringValue(v)
		} else {
			m.ServerFailVLANID = types.StringNull()
		}
	}
}
