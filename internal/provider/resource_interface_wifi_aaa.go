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
	_ resource.Resource                = &InterfaceWifiAaaResource{}
	_ resource.ResourceWithImportState = &InterfaceWifiAaaResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceWifiAaaResource struct {
	reg *client.Registry
}

type InterfaceWifiAaaModel struct {
	ID             types.String `tfsdk:"id"`
	CalledFormat   types.String `tfsdk:"called_format"`
	CallingFormat  types.String `tfsdk:"calling_format"`
	Comment        types.String `tfsdk:"comment"`
	Disabled       types.Bool   `tfsdk:"disabled"`
	InterimUpdate  types.String `tfsdk:"interim_update"`
	MACCaching     types.String `tfsdk:"mac_caching"`
	Name           types.String `tfsdk:"name"`
	NasIdentifier  types.String `tfsdk:"nas_identifier"`
	PasswordFormat types.String `tfsdk:"password_format"`
	UsernameFormat types.String `tfsdk:"username_format"`
	Router         types.String `tfsdk:"router"`
}

func NewInterfaceWifiAaaResource() resource.Resource { return &InterfaceWifiAaaResource{} }

func (r *InterfaceWifiAaaResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_wifi_aaa"
}

func (r *InterfaceWifiAaaResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceWifiAaaResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/wifi/aaa`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"called_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"calling_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interim_update": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mac_caching": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nas_identifier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"password_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"username_format": schema.StringAttribute{
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

func (r *InterfaceWifiAaaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceWifiAaaModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.CalledFormat.IsNull() || plan.CalledFormat.IsUnknown()) {
		body["called-format"] = plan.CalledFormat.ValueString()
	}
	if !(plan.CallingFormat.IsNull() || plan.CallingFormat.IsUnknown()) {
		body["calling-format"] = plan.CallingFormat.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.InterimUpdate.IsNull() || plan.InterimUpdate.IsUnknown()) {
		body["interim-update"] = plan.InterimUpdate.ValueString()
	}
	if !(plan.MACCaching.IsNull() || plan.MACCaching.IsUnknown()) {
		body["mac-caching"] = plan.MACCaching.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.NasIdentifier.IsNull() || plan.NasIdentifier.IsUnknown()) {
		body["nas-identifier"] = plan.NasIdentifier.ValueString()
	}
	if !(plan.PasswordFormat.IsNull() || plan.PasswordFormat.IsUnknown()) {
		body["password-format"] = plan.PasswordFormat.ValueString()
	}
	if !(plan.UsernameFormat.IsNull() || plan.UsernameFormat.IsUnknown()) {
		body["username-format"] = plan.UsernameFormat.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/wifi/aaa", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/wifi/aaa failed", err.Error())
		return
	}
	interfaceWifiAaaApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiAaaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceWifiAaaModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/wifi/aaa", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/wifi/aaa failed", err.Error())
		return
	}
	interfaceWifiAaaApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceWifiAaaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceWifiAaaModel
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
	if !plan.CalledFormat.Equal(state.CalledFormat) && !plan.CalledFormat.IsUnknown() {
		body["called-format"] = plan.CalledFormat.ValueString()
	}
	if !plan.CallingFormat.Equal(state.CallingFormat) && !plan.CallingFormat.IsUnknown() {
		body["calling-format"] = plan.CallingFormat.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.InterimUpdate.Equal(state.InterimUpdate) && !plan.InterimUpdate.IsUnknown() {
		body["interim-update"] = plan.InterimUpdate.ValueString()
	}
	if !plan.MACCaching.Equal(state.MACCaching) && !plan.MACCaching.IsUnknown() {
		body["mac-caching"] = plan.MACCaching.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.NasIdentifier.Equal(state.NasIdentifier) && !plan.NasIdentifier.IsUnknown() {
		body["nas-identifier"] = plan.NasIdentifier.ValueString()
	}
	if !plan.PasswordFormat.Equal(state.PasswordFormat) && !plan.PasswordFormat.IsUnknown() {
		body["password-format"] = plan.PasswordFormat.ValueString()
	}
	if !plan.UsernameFormat.Equal(state.UsernameFormat) && !plan.UsernameFormat.IsUnknown() {
		body["username-format"] = plan.UsernameFormat.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/wifi/aaa", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/wifi/aaa failed", err.Error())
			return
		}
		interfaceWifiAaaApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceWifiAaaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceWifiAaaModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/wifi/aaa", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/wifi/aaa failed", err.Error())
	}
}

func (r *InterfaceWifiAaaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceWifiAaaLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/wifi/aaa matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceWifiAaaLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceWifiAaaLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/wifi/aaa", id)
}

func interfaceWifiAaaApply(ctx context.Context, obj client.Object, m *InterfaceWifiAaaModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["called-format"]; ok {
		_ = v
		if v != "" {
			m.CalledFormat = types.StringValue(v)
		} else {
			m.CalledFormat = types.StringNull()
		}
	} else {
		m.CalledFormat = types.StringNull()
	}
	if v, ok := obj["calling-format"]; ok {
		_ = v
		if v != "" {
			m.CallingFormat = types.StringValue(v)
		} else {
			m.CallingFormat = types.StringNull()
		}
	} else {
		m.CallingFormat = types.StringNull()
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
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["interim-update"]; ok {
		_ = v
		if v != "" {
			m.InterimUpdate = types.StringValue(v)
		} else {
			m.InterimUpdate = types.StringNull()
		}
	} else {
		m.InterimUpdate = types.StringNull()
	}
	if v, ok := obj["mac-caching"]; ok {
		_ = v
		if v != "" {
			m.MACCaching = types.StringValue(v)
		} else {
			m.MACCaching = types.StringNull()
		}
	} else {
		m.MACCaching = types.StringNull()
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
	if v, ok := obj["nas-identifier"]; ok {
		_ = v
		if v != "" {
			m.NasIdentifier = types.StringValue(v)
		} else {
			m.NasIdentifier = types.StringNull()
		}
	} else {
		m.NasIdentifier = types.StringNull()
	}
	if v, ok := obj["password-format"]; ok {
		_ = v
		if v != "" {
			m.PasswordFormat = types.StringValue(v)
		} else {
			m.PasswordFormat = types.StringNull()
		}
	} else {
		m.PasswordFormat = types.StringNull()
	}
	if v, ok := obj["username-format"]; ok {
		_ = v
		if v != "" {
			m.UsernameFormat = types.StringValue(v)
		} else {
			m.UsernameFormat = types.StringNull()
		}
	} else {
		m.UsernameFormat = types.StringNull()
	}
}
