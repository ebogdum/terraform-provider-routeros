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
	_ resource.Resource                = &InterfaceBridgeHostResource{}
	_ resource.ResourceWithImportState = &InterfaceBridgeHostResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceBridgeHostResource struct {
	reg *client.Registry
}

type InterfaceBridgeHostModel struct {
	ID          types.String `tfsdk:"id"`
	Aged        types.Bool   `tfsdk:"aged"`
	AgedOnPeer  types.Bool   `tfsdk:"aged_on_peer"`
	Bridge      types.String `tfsdk:"bridge"`
	Comment     types.String `tfsdk:"comment"`
	Disabled    types.Bool   `tfsdk:"disabled"`
	Dynamic     types.Bool   `tfsdk:"dynamic"`
	ExternalFdb types.Bool   `tfsdk:"external_fdb"`
	Interface   types.String `tfsdk:"interface"`
	Local       types.Bool   `tfsdk:"local"`
	MACAddress  types.String `tfsdk:"mac_address"`
	OnInterface types.String `tfsdk:"on_interface"`
	RemoteIP    types.String `tfsdk:"remote_ip"`
	Vid         types.String `tfsdk:"vid"`
	Router      types.String `tfsdk:"router"`
}

func NewInterfaceBridgeHostResource() resource.Resource { return &InterfaceBridgeHostResource{} }

func (r *InterfaceBridgeHostResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_bridge_host"
}

func (r *InterfaceBridgeHostResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceBridgeHostResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Auto-test requires a typed-reference precondition (e.g. an existing peer,\ninstance, bridge of the specific kind). The current acc-test generator's\ngeneric data.routeros_interface.all lookup can't satisfy these. Use this\nresource manually with explicit references to a precondition resource\nin your config.\n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"aged": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"aged_on_peer": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"bridge": schema.StringAttribute{
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
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"external_fdb": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"local": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"on_interface": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"remote_ip": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"vid": schema.StringAttribute{
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

func (r *InterfaceBridgeHostResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceBridgeHostModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Bridge.IsNull() || plan.Bridge.IsUnknown()) {
		body["bridge"] = plan.Bridge.ValueString()
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
	if !(plan.Vid.IsNull() || plan.Vid.IsUnknown()) {
		body["vid"] = plan.Vid.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/bridge/host", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/bridge/host failed", err.Error())
		return
	}
	interfaceBridgeHostApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceBridgeHostResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceBridgeHostModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/bridge/host", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/bridge/host failed", err.Error())
		return
	}
	interfaceBridgeHostApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceBridgeHostResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceBridgeHostModel
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
	if !plan.Bridge.Equal(state.Bridge) {
		body["bridge"] = plan.Bridge.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.MACAddress.Equal(state.MACAddress) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.Vid.Equal(state.Vid) {
		body["vid"] = plan.Vid.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/bridge/host", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/bridge/host failed", err.Error())
			return
		}
		interfaceBridgeHostApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceBridgeHostResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceBridgeHostModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/bridge/host", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/bridge/host failed", err.Error())
	}
}

func (r *InterfaceBridgeHostResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceBridgeHostLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/bridge/host matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceBridgeHostLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceBridgeHostLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/bridge/host", id)
}

func interfaceBridgeHostApply(ctx context.Context, obj client.Object, m *InterfaceBridgeHostModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["aged"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Aged = types.BoolValue(b)
		} else {
			m.Aged = types.BoolNull()
		}
	} else {
		m.Aged = types.BoolNull()
	}
	if v, ok := obj["aged-on-peer"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AgedOnPeer = types.BoolValue(b)
		} else {
			m.AgedOnPeer = types.BoolNull()
		}
	} else {
		m.AgedOnPeer = types.BoolNull()
	}
	if v, ok := obj["bridge"]; ok {
		_ = v
		if v != "" {
			m.Bridge = types.StringValue(v)
		} else {
			m.Bridge = types.StringNull()
		}
	} else {
		m.Bridge = types.StringNull()
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
	if v, ok := obj["external-fdb"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.ExternalFdb = types.BoolValue(b)
		} else {
			m.ExternalFdb = types.BoolNull()
		}
	} else {
		m.ExternalFdb = types.BoolNull()
	}
	if v, ok := obj["interface"]; ok {
		_ = v
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	} else {
		m.Interface = types.StringNull()
	}
	if v, ok := obj["local"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Local = types.BoolValue(b)
		} else {
			m.Local = types.BoolNull()
		}
	} else {
		m.Local = types.BoolNull()
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
	if v, ok := obj["on-interface"]; ok {
		_ = v
		if v != "" {
			m.OnInterface = types.StringValue(v)
		} else {
			m.OnInterface = types.StringNull()
		}
	} else {
		m.OnInterface = types.StringNull()
	}
	if v, ok := obj["remote-ip"]; ok {
		_ = v
		if v != "" {
			m.RemoteIP = types.StringValue(v)
		} else {
			m.RemoteIP = types.StringNull()
		}
	} else {
		m.RemoteIP = types.StringNull()
	}
	if v, ok := obj["vid"]; ok {
		_ = v
		if v != "" {
			m.Vid = types.StringValue(v)
		} else {
			m.Vid = types.StringNull()
		}
	} else {
		m.Vid = types.StringNull()
	}
}
