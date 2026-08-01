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
	_ resource.Resource                = &InterfacePPPServerResource{}
	_ resource.ResourceWithImportState = &InterfacePPPServerResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfacePPPServerResource struct {
	reg *client.Registry
}

type InterfacePPPServerModel struct {
	ID             types.String `tfsdk:"id"`
	RingCount      types.String `tfsdk:"ring_count"`
	Port           types.String `tfsdk:"port"`
	NullModem      types.String `tfsdk:"null_modem"`
	ModemInit      types.String `tfsdk:"modem_init"`
	DataChannel    types.String `tfsdk:"data_channel"`
	Authentication types.String `tfsdk:"authentication"`
	Comment        types.String `tfsdk:"comment"`
	Disabled       types.Bool   `tfsdk:"disabled"`
	MaxMru         types.String `tfsdk:"max_mru"`
	MaxMTU         types.String `tfsdk:"max_mtu"`
	Mrru           types.String `tfsdk:"mrru"`
	Name           types.String `tfsdk:"name"`
	Profile        types.String `tfsdk:"profile"`
	Router         types.String `tfsdk:"router"`
}

func NewInterfacePPPServerResource() resource.Resource { return &InterfacePPPServerResource{} }

func (r *InterfacePPPServerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ppp_server"
}

func (r *InterfacePPPServerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfacePPPServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/ppp-server`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ring_count": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ring-count`.",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `port`.",
			},
			"null_modem": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `null-modem`.",
			},
			"modem_init": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `modem-init`.",
			},
			"data_channel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `data-channel`.",
			},
			"authentication": schema.StringAttribute{
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
			"max_mru": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"max_mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mrru": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"profile": schema.StringAttribute{
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

func (r *InterfacePPPServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfacePPPServerModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Authentication.IsNull() || plan.Authentication.IsUnknown()) {
		body["authentication"] = plan.Authentication.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.MaxMru.IsNull() || plan.MaxMru.IsUnknown()) {
		body["max-mru"] = plan.MaxMru.ValueString()
	}
	if !(plan.MaxMTU.IsNull() || plan.MaxMTU.IsUnknown()) {
		body["max-mtu"] = plan.MaxMTU.ValueString()
	}
	if !(plan.Mrru.IsNull() || plan.Mrru.IsUnknown()) {
		body["mrru"] = plan.Mrru.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Profile.IsNull() || plan.Profile.IsUnknown()) {
		body["profile"] = plan.Profile.ValueString()
	}
	if !(plan.DataChannel.IsNull() || plan.DataChannel.IsUnknown()) {
		body["data-channel"] = plan.DataChannel.ValueString()
	}
	if !(plan.ModemInit.IsNull() || plan.ModemInit.IsUnknown()) {
		body["modem-init"] = plan.ModemInit.ValueString()
	}
	if !(plan.NullModem.IsNull() || plan.NullModem.IsUnknown()) {
		body["null-modem"] = plan.NullModem.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.RingCount.IsNull() || plan.RingCount.IsUnknown()) {
		body["ring-count"] = plan.RingCount.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/ppp-server", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/ppp-server failed", err.Error())
		return
	}
	interfacePPPServerApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfacePPPServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfacePPPServerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/ppp-server", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/ppp-server failed", err.Error())
		return
	}
	interfacePPPServerApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfacePPPServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfacePPPServerModel
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
	if !plan.Authentication.Equal(state.Authentication) && !plan.Authentication.IsUnknown() {
		body["authentication"] = plan.Authentication.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.MaxMru.Equal(state.MaxMru) && !plan.MaxMru.IsUnknown() {
		body["max-mru"] = plan.MaxMru.ValueString()
	}
	if !plan.MaxMTU.Equal(state.MaxMTU) && !plan.MaxMTU.IsUnknown() {
		body["max-mtu"] = plan.MaxMTU.ValueString()
	}
	if !plan.Mrru.Equal(state.Mrru) && !plan.Mrru.IsUnknown() {
		body["mrru"] = plan.Mrru.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Profile.Equal(state.Profile) && !plan.Profile.IsUnknown() {
		body["profile"] = plan.Profile.ValueString()
	}
	if !plan.DataChannel.Equal(state.DataChannel) && !plan.DataChannel.IsUnknown() {
		body["data-channel"] = plan.DataChannel.ValueString()
	}
	if !plan.ModemInit.Equal(state.ModemInit) && !plan.ModemInit.IsUnknown() {
		body["modem-init"] = plan.ModemInit.ValueString()
	}
	if !plan.NullModem.Equal(state.NullModem) && !plan.NullModem.IsUnknown() {
		body["null-modem"] = plan.NullModem.ValueString()
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.RingCount.Equal(state.RingCount) && !plan.RingCount.IsUnknown() {
		body["ring-count"] = plan.RingCount.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/ppp-server", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/ppp-server failed", err.Error())
			return
		}
		interfacePPPServerApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfacePPPServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfacePPPServerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/ppp-server", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/ppp-server failed", err.Error())
	}
}

func (r *InterfacePPPServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfacePPPServerLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/ppp-server matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfacePPPServerLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfacePPPServerLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/ppp-server", id)
}

func interfacePPPServerApply(ctx context.Context, obj client.Object, m *InterfacePPPServerModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["ring-count"]; ok && v != "" {
		m.RingCount = types.StringValue(v)
	} else {
		m.RingCount = types.StringNull()
	}
	if v, ok := obj["port"]; ok && v != "" {
		m.Port = types.StringValue(v)
	} else {
		m.Port = types.StringNull()
	}
	if v, ok := obj["null-modem"]; ok && v != "" {
		m.NullModem = types.StringValue(v)
	} else {
		m.NullModem = types.StringNull()
	}
	if v, ok := obj["modem-init"]; ok && v != "" {
		m.ModemInit = types.StringValue(v)
	} else {
		m.ModemInit = types.StringNull()
	}
	if v, ok := obj["data-channel"]; ok && v != "" {
		m.DataChannel = types.StringValue(v)
	} else {
		m.DataChannel = types.StringNull()
	}
	if v, ok := obj["authentication"]; ok {
		_ = v
		if v != "" {
			m.Authentication = types.StringValue(v)
		} else {
			m.Authentication = types.StringNull()
		}
	} else {
		m.Authentication = types.StringNull()
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
	if v, ok := obj["max-mru"]; ok {
		_ = v
		if v != "" {
			m.MaxMru = types.StringValue(v)
		} else {
			m.MaxMru = types.StringNull()
		}
	} else {
		m.MaxMru = types.StringNull()
	}
	if v, ok := obj["max-mtu"]; ok {
		_ = v
		if v != "" {
			m.MaxMTU = types.StringValue(v)
		} else {
			m.MaxMTU = types.StringNull()
		}
	} else {
		m.MaxMTU = types.StringNull()
	}
	if v, ok := obj["mrru"]; ok {
		_ = v
		if v != "" {
			m.Mrru = types.StringValue(v)
		} else {
			m.Mrru = types.StringNull()
		}
	} else {
		m.Mrru = types.StringNull()
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
	if v, ok := obj["profile"]; ok {
		_ = v
		if v != "" {
			m.Profile = types.StringValue(v)
		} else {
			m.Profile = types.StringNull()
		}
	} else {
		m.Profile = types.StringNull()
	}
}
