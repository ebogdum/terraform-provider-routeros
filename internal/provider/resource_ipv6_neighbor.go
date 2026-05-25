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
	_ resource.Resource                = &IPV6NeighborResource{}
	_ resource.ResourceWithImportState = &IPV6NeighborResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPV6NeighborResource struct {
	reg *client.Registry
}

type IPV6NeighborModel struct {
	ID         types.String `tfsdk:"id"`
	Address    types.String `tfsdk:"address"`
	BridgePort types.String `tfsdk:"bridge_port"`
	Comment    types.String `tfsdk:"comment"`
	Disabled   types.Bool   `tfsdk:"disabled"`
	Dynamic    types.Bool   `tfsdk:"dynamic"`
	HostName   types.String `tfsdk:"host_name"`
	Interface  types.String `tfsdk:"interface"`
	MACAddress types.String `tfsdk:"mac_address"`
	MACPing    types.String `tfsdk:"mac_ping"`
	MACTelnet  types.String `tfsdk:"mac_telnet"`
	MakeStatic types.String `tfsdk:"make_static"`
	Ping       types.String `tfsdk:"ping"`
	RouterRos  types.Bool   `tfsdk:"router_ros"`
	Status     types.String `tfsdk:"status"`
	Telnet     types.String `tfsdk:"telnet"`
	Torch      types.String `tfsdk:"torch"`
	Vrf        types.String `tfsdk:"vrf"`
	Router     types.String `tfsdk:"router"`
}

func NewIPV6NeighborResource() resource.Resource { return &IPV6NeighborResource{} }

func (r *IPV6NeighborResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv6_neighbor"
}

func (r *IPV6NeighborResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *IPV6NeighborResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "IPv6 neighbor table — read-only on most devices.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"bridge_port": schema.StringAttribute{
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
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"host_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mac_ping": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mac_telnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"make_static": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ping": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"router_ros": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"status": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"telnet": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"torch": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vrf": schema.StringAttribute{
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

func (r *IPV6NeighborResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPV6NeighborModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Address.IsNull() || plan.Address.IsUnknown()) {
		body["address"] = plan.Address.ValueString()
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
	if !(plan.MACPing.IsNull() || plan.MACPing.IsUnknown()) {
		body["mac-ping"] = plan.MACPing.ValueString()
	}
	if !(plan.MACTelnet.IsNull() || plan.MACTelnet.IsUnknown()) {
		body["mac-telnet"] = plan.MACTelnet.ValueString()
	}
	if !(plan.MakeStatic.IsNull() || plan.MakeStatic.IsUnknown()) {
		body["make-static"] = plan.MakeStatic.ValueString()
	}
	if !(plan.Ping.IsNull() || plan.Ping.IsUnknown()) {
		body["ping"] = plan.Ping.ValueString()
	}
	if !(plan.RouterRos.IsNull() || plan.RouterRos.IsUnknown()) {
		body["router"] = client.FormatBool(plan.RouterRos.ValueBool())
	}
	if !(plan.Telnet.IsNull() || plan.Telnet.IsUnknown()) {
		body["telnet"] = plan.Telnet.ValueString()
	}
	if !(plan.Torch.IsNull() || plan.Torch.IsUnknown()) {
		body["torch"] = plan.Torch.ValueString()
	}
	obj, err := c.Add(ctx, "/ipv6/neighbor", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ipv6/neighbor failed", err.Error())
		return
	}
	iPV6NeighborApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6NeighborResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPV6NeighborModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ipv6/neighbor", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ipv6/neighbor failed", err.Error())
		return
	}
	iPV6NeighborApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPV6NeighborResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPV6NeighborModel
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
	if !plan.Address.Equal(state.Address) {
		body["address"] = plan.Address.ValueString()
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
	if !plan.MACPing.Equal(state.MACPing) {
		body["mac-ping"] = plan.MACPing.ValueString()
	}
	if !plan.MACTelnet.Equal(state.MACTelnet) {
		body["mac-telnet"] = plan.MACTelnet.ValueString()
	}
	if !plan.MakeStatic.Equal(state.MakeStatic) {
		body["make-static"] = plan.MakeStatic.ValueString()
	}
	if !plan.Ping.Equal(state.Ping) {
		body["ping"] = plan.Ping.ValueString()
	}
	if !plan.RouterRos.Equal(state.RouterRos) {
		body["router"] = client.FormatBool(plan.RouterRos.ValueBool())
	}
	if !plan.Telnet.Equal(state.Telnet) {
		body["telnet"] = plan.Telnet.ValueString()
	}
	if !plan.Torch.Equal(state.Torch) {
		body["torch"] = plan.Torch.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ipv6/neighbor", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ipv6/neighbor failed", err.Error())
			return
		}
		iPV6NeighborApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6NeighborResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPV6NeighborModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ipv6/neighbor", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ipv6/neighbor failed", err.Error())
	}
}

func (r *IPV6NeighborResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPV6NeighborLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ipv6/neighbor matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPV6NeighborLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPV6NeighborLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ipv6/neighbor", id)
}

func iPV6NeighborApply(ctx context.Context, obj client.Object, m *IPV6NeighborModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["address"]; ok {
		_ = v
		if v != "" {
			m.Address = types.StringValue(v)
		} else {
			m.Address = types.StringNull()
		}
	} else {
		m.Address = types.StringNull()
	}
	if v, ok := obj["bridge-port"]; ok {
		_ = v
		if v != "" {
			m.BridgePort = types.StringValue(v)
		} else {
			m.BridgePort = types.StringNull()
		}
	} else {
		m.BridgePort = types.StringNull()
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
	if v, ok := obj["host-name"]; ok {
		_ = v
		if v != "" {
			m.HostName = types.StringValue(v)
		} else {
			m.HostName = types.StringNull()
		}
	} else {
		m.HostName = types.StringNull()
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
	if v, ok := obj["mac-ping"]; ok {
		_ = v
		if v != "" {
			m.MACPing = types.StringValue(v)
		} else {
			m.MACPing = types.StringNull()
		}
	} else {
		m.MACPing = types.StringNull()
	}
	if v, ok := obj["mac-telnet"]; ok {
		_ = v
		if v != "" {
			m.MACTelnet = types.StringValue(v)
		} else {
			m.MACTelnet = types.StringNull()
		}
	} else {
		m.MACTelnet = types.StringNull()
	}
	if v, ok := obj["make-static"]; ok {
		_ = v
		if v != "" {
			m.MakeStatic = types.StringValue(v)
		} else {
			m.MakeStatic = types.StringNull()
		}
	} else {
		m.MakeStatic = types.StringNull()
	}
	if v, ok := obj["ping"]; ok {
		_ = v
		if v != "" {
			m.Ping = types.StringValue(v)
		} else {
			m.Ping = types.StringNull()
		}
	} else {
		m.Ping = types.StringNull()
	}
	if v, ok := obj["router"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.RouterRos = types.BoolValue(b)
		} else {
			m.RouterRos = types.BoolNull()
		}
	} else {
		m.RouterRos = types.BoolNull()
	}
	if v, ok := obj["status"]; ok {
		_ = v
		if v != "" {
			m.Status = types.StringValue(v)
		} else {
			m.Status = types.StringNull()
		}
	} else {
		m.Status = types.StringNull()
	}
	if v, ok := obj["telnet"]; ok {
		_ = v
		if v != "" {
			m.Telnet = types.StringValue(v)
		} else {
			m.Telnet = types.StringNull()
		}
	} else {
		m.Telnet = types.StringNull()
	}
	if v, ok := obj["torch"]; ok {
		_ = v
		if v != "" {
			m.Torch = types.StringValue(v)
		} else {
			m.Torch = types.StringNull()
		}
	} else {
		m.Torch = types.StringNull()
	}
	if v, ok := obj["vrf"]; ok {
		_ = v
		if v != "" {
			m.Vrf = types.StringValue(v)
		} else {
			m.Vrf = types.StringNull()
		}
	} else {
		m.Vrf = types.StringNull()
	}
}
