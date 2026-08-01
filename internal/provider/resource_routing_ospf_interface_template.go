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
	_ resource.Resource                = &RoutingOSPFInterfaceTemplateResource{}
	_ resource.ResourceWithImportState = &RoutingOSPFInterfaceTemplateResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingOSPFInterfaceTemplateResource struct {
	reg *client.Registry
}

type RoutingOSPFInterfaceTemplateModel struct {
	ID                 types.String `tfsdk:"id"`
	Type               types.String `tfsdk:"type"`
	Auth               types.String `tfsdk:"auth"`
	Area               types.String `tfsdk:"area"`
	AuthID             types.String `tfsdk:"auth_id"`
	AuthKey            types.String `tfsdk:"auth_key"`
	Authentication     types.String `tfsdk:"authentication"`
	Comment            types.String `tfsdk:"comment"`
	Cost               types.Int64  `tfsdk:"cost"`
	DeadInterval       types.String `tfsdk:"dead_interval"`
	Disabled           types.Bool   `tfsdk:"disabled"`
	HelloInterval      types.String `tfsdk:"hello_interval"`
	InstanceID         types.Int64  `tfsdk:"instance_id"`
	Interfaces         types.String `tfsdk:"interfaces"`
	Invalid            types.Bool   `tfsdk:"invalid"`
	NetworkType        types.String `tfsdk:"network_type"`
	Networks           types.String `tfsdk:"networks"`
	Passive            types.Bool   `tfsdk:"passive"`
	PrefixList         types.String `tfsdk:"prefix_list"`
	Priority           types.Int64  `tfsdk:"priority"`
	RetransmitInterval types.String `tfsdk:"retransmit_interval"`
	TransmitDelay      types.Int64  `tfsdk:"transmit_delay"`
	UseBfd             types.String `tfsdk:"use_bfd"`
	VlinkNeighborID    types.String `tfsdk:"vlink_neighbor_id"`
	VlinkTransitArea   types.String `tfsdk:"vlink_transit_area"`
	Router             types.String `tfsdk:"router"`
}

func NewRoutingOSPFInterfaceTemplateResource() resource.Resource {
	return &RoutingOSPFInterfaceTemplateResource{}
}

func (r *RoutingOSPFInterfaceTemplateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_ospf_interface_template"
}

func (r *RoutingOSPFInterfaceTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingOSPFInterfaceTemplateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "References an existing ospf area; auto-test can't synthesise.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `type`.",
			},
			"auth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `auth`.",
			},
			"area": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"auth_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"auth_key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"authentication": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"cost": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dead_interval": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"hello_interval": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"instance_id": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interfaces": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"network_type": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"broadcast", "nbma", "ptp", "ptp-unnumbered", "ptmp", "virtual-link", "ptmp-broadcast"}...)},
			},
			"networks": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"passive": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"prefix_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"retransmit_interval": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"transmit_delay": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"use_bfd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vlink_neighbor_id": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"vlink_transit_area": schema.StringAttribute{
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

func (r *RoutingOSPFInterfaceTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingOSPFInterfaceTemplateModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Area.IsNull() || plan.Area.IsUnknown()) {
		body["area"] = plan.Area.ValueString()
	}
	if !(plan.AuthID.IsNull() || plan.AuthID.IsUnknown()) {
		body["auth-id"] = plan.AuthID.ValueString()
	}
	if !(plan.AuthKey.IsNull() || plan.AuthKey.IsUnknown()) {
		body["auth-key"] = plan.AuthKey.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Cost.IsNull() || plan.Cost.IsUnknown()) {
		body["cost"] = client.FormatInt64(plan.Cost.ValueInt64())
	}
	if !(plan.DeadInterval.IsNull() || plan.DeadInterval.IsUnknown()) {
		body["dead-interval"] = plan.DeadInterval.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.HelloInterval.IsNull() || plan.HelloInterval.IsUnknown()) {
		body["hello-interval"] = plan.HelloInterval.ValueString()
	}
	if !(plan.InstanceID.IsNull() || plan.InstanceID.IsUnknown()) {
		body["instance-id"] = client.FormatInt64(plan.InstanceID.ValueInt64())
	}
	if !(plan.Interfaces.IsNull() || plan.Interfaces.IsUnknown()) {
		body["interfaces"] = plan.Interfaces.ValueString()
	}
	if !(plan.Networks.IsNull() || plan.Networks.IsUnknown()) {
		body["networks"] = plan.Networks.ValueString()
	}
	if !(plan.Passive.IsNull() || plan.Passive.IsUnknown()) {
		body["passive"] = client.FormatBool(plan.Passive.ValueBool())
	}
	if !(plan.PrefixList.IsNull() || plan.PrefixList.IsUnknown()) {
		body["prefix-list"] = plan.PrefixList.ValueString()
	}
	if !(plan.Priority.IsNull() || plan.Priority.IsUnknown()) {
		body["priority"] = client.FormatInt64(plan.Priority.ValueInt64())
	}
	if !(plan.RetransmitInterval.IsNull() || plan.RetransmitInterval.IsUnknown()) {
		body["retransmit-interval"] = plan.RetransmitInterval.ValueString()
	}
	if !(plan.TransmitDelay.IsNull() || plan.TransmitDelay.IsUnknown()) {
		body["transmit-delay"] = client.FormatInt64(plan.TransmitDelay.ValueInt64())
	}
	if !(plan.UseBfd.IsNull() || plan.UseBfd.IsUnknown()) {
		body["use-bfd"] = plan.UseBfd.ValueString()
	}
	if !(plan.Auth.IsNull() || plan.Auth.IsUnknown()) {
		body["auth"] = plan.Auth.ValueString()
	}
	if !(plan.Type.IsNull() || plan.Type.IsUnknown()) {
		body["type"] = plan.Type.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/ospf/interface-template", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/ospf/interface-template failed", err.Error())
		return
	}
	routingOSPFInterfaceTemplateApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingOSPFInterfaceTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingOSPFInterfaceTemplateModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/ospf/interface-template", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/ospf/interface-template failed", err.Error())
		return
	}
	routingOSPFInterfaceTemplateApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingOSPFInterfaceTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingOSPFInterfaceTemplateModel
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
	if !plan.Area.Equal(state.Area) && !plan.Area.IsUnknown() {
		body["area"] = plan.Area.ValueString()
	}
	if !plan.AuthID.Equal(state.AuthID) && !plan.AuthID.IsUnknown() {
		body["auth-id"] = plan.AuthID.ValueString()
	}
	if !plan.AuthKey.Equal(state.AuthKey) && !plan.AuthKey.IsUnknown() {
		body["auth-key"] = plan.AuthKey.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Cost.Equal(state.Cost) && !plan.Cost.IsUnknown() {
		body["cost"] = client.FormatInt64(plan.Cost.ValueInt64())
	}
	if !plan.DeadInterval.Equal(state.DeadInterval) && !plan.DeadInterval.IsUnknown() {
		body["dead-interval"] = plan.DeadInterval.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.HelloInterval.Equal(state.HelloInterval) && !plan.HelloInterval.IsUnknown() {
		body["hello-interval"] = plan.HelloInterval.ValueString()
	}
	if !plan.InstanceID.Equal(state.InstanceID) && !plan.InstanceID.IsUnknown() {
		body["instance-id"] = client.FormatInt64(plan.InstanceID.ValueInt64())
	}
	if !plan.Interfaces.Equal(state.Interfaces) && !plan.Interfaces.IsUnknown() {
		body["interfaces"] = plan.Interfaces.ValueString()
	}
	if !plan.Networks.Equal(state.Networks) && !plan.Networks.IsUnknown() {
		body["networks"] = plan.Networks.ValueString()
	}
	if !plan.Passive.Equal(state.Passive) && !plan.Passive.IsUnknown() {
		body["passive"] = client.FormatBool(plan.Passive.ValueBool())
	}
	if !plan.PrefixList.Equal(state.PrefixList) && !plan.PrefixList.IsUnknown() {
		body["prefix-list"] = plan.PrefixList.ValueString()
	}
	if !plan.Priority.Equal(state.Priority) && !plan.Priority.IsUnknown() {
		body["priority"] = client.FormatInt64(plan.Priority.ValueInt64())
	}
	if !plan.RetransmitInterval.Equal(state.RetransmitInterval) && !plan.RetransmitInterval.IsUnknown() {
		body["retransmit-interval"] = plan.RetransmitInterval.ValueString()
	}
	if !plan.TransmitDelay.Equal(state.TransmitDelay) && !plan.TransmitDelay.IsUnknown() {
		body["transmit-delay"] = client.FormatInt64(plan.TransmitDelay.ValueInt64())
	}
	if !plan.UseBfd.Equal(state.UseBfd) && !plan.UseBfd.IsUnknown() {
		body["use-bfd"] = plan.UseBfd.ValueString()
	}
	if !plan.Auth.Equal(state.Auth) && !plan.Auth.IsUnknown() {
		body["auth"] = plan.Auth.ValueString()
	}
	if !plan.Type.Equal(state.Type) && !plan.Type.IsUnknown() {
		body["type"] = plan.Type.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/ospf/interface-template", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/ospf/interface-template failed", err.Error())
			return
		}
		routingOSPFInterfaceTemplateApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingOSPFInterfaceTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingOSPFInterfaceTemplateModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/ospf/interface-template", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/ospf/interface-template failed", err.Error())
	}
}

func (r *RoutingOSPFInterfaceTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingOSPFInterfaceTemplateLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/ospf/interface-template matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingOSPFInterfaceTemplateLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingOSPFInterfaceTemplateLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/ospf/interface-template", id)
}

func routingOSPFInterfaceTemplateApply(ctx context.Context, obj client.Object, m *RoutingOSPFInterfaceTemplateModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["type"]; ok && v != "" {
		m.Type = types.StringValue(v)
	} else {
		m.Type = types.StringNull()
	}
	if v, ok := obj["auth"]; ok && v != "" {
		m.Auth = types.StringValue(v)
	} else {
		m.Auth = types.StringNull()
	}
	if v, ok := obj["area"]; ok {
		_ = v
		if v != "" {
			m.Area = types.StringValue(v)
		} else {
			m.Area = types.StringNull()
		}
	} else {
		m.Area = types.StringNull()
	}
	if v, ok := obj["auth-id"]; ok {
		_ = v
		if v != "" {
			m.AuthID = types.StringValue(v)
		} else {
			m.AuthID = types.StringNull()
		}
	} else {
		m.AuthID = types.StringNull()
	}
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.AuthKey already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["auth-key"]; ok && v != "" {
		_ = v
		if v != "" {
			m.AuthKey = types.StringValue(v)
		} else {
			m.AuthKey = types.StringNull()
		}
	} else if m.AuthKey.IsUnknown() {
		m.AuthKey = types.StringNull()
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
	if v, ok := obj["cost"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Cost = types.Int64Value(n)
		} else {
			m.Cost = types.Int64Null()
		}
	} else {
		m.Cost = types.Int64Null()
	}
	if v, ok := obj["dead-interval"]; ok {
		_ = v
		if v != "" {
			m.DeadInterval = types.StringValue(v)
		} else {
			m.DeadInterval = types.StringNull()
		}
	} else {
		m.DeadInterval = types.StringNull()
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
	if v, ok := obj["hello-interval"]; ok {
		_ = v
		if v != "" {
			m.HelloInterval = types.StringValue(v)
		} else {
			m.HelloInterval = types.StringNull()
		}
	} else {
		m.HelloInterval = types.StringNull()
	}
	if v, ok := obj["instance-id"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.InstanceID = types.Int64Value(n)
		} else {
			m.InstanceID = types.Int64Null()
		}
	} else {
		m.InstanceID = types.Int64Null()
	}
	if v, ok := obj["interfaces"]; ok {
		_ = v
		if v != "" {
			m.Interfaces = types.StringValue(v)
		} else {
			m.Interfaces = types.StringNull()
		}
	} else {
		m.Interfaces = types.StringNull()
	}
	if v, ok := obj["invalid"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Invalid = types.BoolValue(b)
		} else {
			m.Invalid = types.BoolNull()
		}
	} else {
		m.Invalid = types.BoolNull()
	}
	if v, ok := obj["network-type"]; ok {
		_ = v
		if v != "" {
			m.NetworkType = types.StringValue(v)
		} else {
			m.NetworkType = types.StringNull()
		}
	} else {
		m.NetworkType = types.StringNull()
	}
	if v, ok := obj["networks"]; ok {
		_ = v
		if v != "" {
			m.Networks = types.StringValue(v)
		} else {
			m.Networks = types.StringNull()
		}
	} else {
		m.Networks = types.StringNull()
	}
	if v, ok := obj["passive"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Passive = types.BoolValue(b)
		} else {
			m.Passive = types.BoolNull()
		}
	} else {
		m.Passive = types.BoolNull()
	}
	if v, ok := obj["prefix-list"]; ok {
		_ = v
		if v != "" {
			m.PrefixList = types.StringValue(v)
		} else {
			m.PrefixList = types.StringNull()
		}
	} else {
		m.PrefixList = types.StringNull()
	}
	if v, ok := obj["priority"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Priority = types.Int64Value(n)
		} else {
			m.Priority = types.Int64Null()
		}
	} else {
		m.Priority = types.Int64Null()
	}
	if v, ok := obj["retransmit-interval"]; ok {
		_ = v
		if v != "" {
			m.RetransmitInterval = types.StringValue(v)
		} else {
			m.RetransmitInterval = types.StringNull()
		}
	} else {
		m.RetransmitInterval = types.StringNull()
	}
	if v, ok := obj["transmit-delay"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TransmitDelay = types.Int64Value(n)
		} else {
			m.TransmitDelay = types.Int64Null()
		}
	} else {
		m.TransmitDelay = types.Int64Null()
	}
	if v, ok := obj["use-bfd"]; ok {
		_ = v
		if v != "" {
			m.UseBfd = types.StringValue(v)
		} else {
			m.UseBfd = types.StringNull()
		}
	} else {
		m.UseBfd = types.StringNull()
	}
	if v, ok := obj["vlink-neighbor-id"]; ok {
		_ = v
		if v != "" {
			m.VlinkNeighborID = types.StringValue(v)
		} else {
			m.VlinkNeighborID = types.StringNull()
		}
	} else {
		m.VlinkNeighborID = types.StringNull()
	}
	if v, ok := obj["vlink-transit-area"]; ok {
		_ = v
		if v != "" {
			m.VlinkTransitArea = types.StringValue(v)
		} else {
			m.VlinkTransitArea = types.StringNull()
		}
	} else {
		m.VlinkTransitArea = types.StringNull()
	}
}
