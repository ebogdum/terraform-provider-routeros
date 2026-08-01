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
	_ resource.Resource                = &RoutingRipInterfaceResource{}
	_ resource.ResourceWithImportState = &RoutingRipInterfaceResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingRipInterfaceResource struct {
	reg *client.Registry
}

type RoutingRipInterfaceModel struct {
	ID              types.String `tfsdk:"id"`
	Cost            types.String `tfsdk:"cost"`
	Disabled        types.Bool   `tfsdk:"disabled"`
	Instance        types.String `tfsdk:"instance"`
	Interfaces      types.String `tfsdk:"interfaces"`
	KeyChain        types.String `tfsdk:"key_chain"`
	Mode            types.String `tfsdk:"mode"`
	Name            types.String `tfsdk:"name"`
	Password        types.String `tfsdk:"password"`
	PoisonReverse   types.String `tfsdk:"poison_reverse"`
	SourceAddresses types.String `tfsdk:"source_addresses"`
	SplitHorizon    types.String `tfsdk:"split_horizon"`
	UseBfd          types.String `tfsdk:"use_bfd"`
	Router          types.String `tfsdk:"router"`
}

func NewRoutingRipInterfaceResource() resource.Resource { return &RoutingRipInterfaceResource{} }

func (r *RoutingRipInterfaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_rip_interface"
}

func (r *RoutingRipInterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingRipInterfaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Discovered; needs rip instance",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"cost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"instance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interfaces": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"key_chain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"poison_reverse": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"source_addresses": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"split_horizon": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"use_bfd": schema.StringAttribute{
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

func (r *RoutingRipInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingRipInterfaceModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Cost.IsNull() || plan.Cost.IsUnknown()) {
		body["cost"] = plan.Cost.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Instance.IsNull() || plan.Instance.IsUnknown()) {
		body["instance"] = plan.Instance.ValueString()
	}
	if !(plan.Interfaces.IsNull() || plan.Interfaces.IsUnknown()) {
		body["interfaces"] = plan.Interfaces.ValueString()
	}
	if !(plan.KeyChain.IsNull() || plan.KeyChain.IsUnknown()) {
		body["key-chain"] = plan.KeyChain.ValueString()
	}
	if !(plan.Mode.IsNull() || plan.Mode.IsUnknown()) {
		body["mode"] = plan.Mode.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Password.IsNull() || plan.Password.IsUnknown()) {
		body["password"] = plan.Password.ValueString()
	}
	if !(plan.PoisonReverse.IsNull() || plan.PoisonReverse.IsUnknown()) {
		body["poison-reverse"] = plan.PoisonReverse.ValueString()
	}
	if !(plan.SourceAddresses.IsNull() || plan.SourceAddresses.IsUnknown()) {
		body["source-addresses"] = plan.SourceAddresses.ValueString()
	}
	if !(plan.SplitHorizon.IsNull() || plan.SplitHorizon.IsUnknown()) {
		body["split-horizon"] = plan.SplitHorizon.ValueString()
	}
	if !(plan.UseBfd.IsNull() || plan.UseBfd.IsUnknown()) {
		body["use-bfd"] = plan.UseBfd.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/rip/interface", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/rip/interface failed", err.Error())
		return
	}
	routingRipInterfaceApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingRipInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingRipInterfaceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/rip/interface", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/rip/interface failed", err.Error())
		return
	}
	routingRipInterfaceApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingRipInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingRipInterfaceModel
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
	if !plan.Cost.Equal(state.Cost) && !plan.Cost.IsUnknown() {
		body["cost"] = plan.Cost.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Instance.Equal(state.Instance) && !plan.Instance.IsUnknown() {
		body["instance"] = plan.Instance.ValueString()
	}
	if !plan.Interfaces.Equal(state.Interfaces) && !plan.Interfaces.IsUnknown() {
		body["interfaces"] = plan.Interfaces.ValueString()
	}
	if !plan.KeyChain.Equal(state.KeyChain) && !plan.KeyChain.IsUnknown() {
		body["key-chain"] = plan.KeyChain.ValueString()
	}
	if !plan.Mode.Equal(state.Mode) && !plan.Mode.IsUnknown() {
		body["mode"] = plan.Mode.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Password.Equal(state.Password) && !plan.Password.IsUnknown() {
		body["password"] = plan.Password.ValueString()
	}
	if !plan.PoisonReverse.Equal(state.PoisonReverse) && !plan.PoisonReverse.IsUnknown() {
		body["poison-reverse"] = plan.PoisonReverse.ValueString()
	}
	if !plan.SourceAddresses.Equal(state.SourceAddresses) && !plan.SourceAddresses.IsUnknown() {
		body["source-addresses"] = plan.SourceAddresses.ValueString()
	}
	if !plan.SplitHorizon.Equal(state.SplitHorizon) && !plan.SplitHorizon.IsUnknown() {
		body["split-horizon"] = plan.SplitHorizon.ValueString()
	}
	if !plan.UseBfd.Equal(state.UseBfd) && !plan.UseBfd.IsUnknown() {
		body["use-bfd"] = plan.UseBfd.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/rip/interface", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/rip/interface failed", err.Error())
			return
		}
		routingRipInterfaceApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingRipInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingRipInterfaceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/rip/interface", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/rip/interface failed", err.Error())
	}
}

func (r *RoutingRipInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingRipInterfaceLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/rip/interface matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingRipInterfaceLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingRipInterfaceLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/rip/interface", id)
}

func routingRipInterfaceApply(ctx context.Context, obj client.Object, m *RoutingRipInterfaceModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["cost"]; ok {
		_ = v
		if v != "" {
			m.Cost = types.StringValue(v)
		} else {
			m.Cost = types.StringNull()
		}
	} else {
		m.Cost = types.StringNull()
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
	if v, ok := obj["instance"]; ok {
		_ = v
		if v != "" {
			m.Instance = types.StringValue(v)
		} else {
			m.Instance = types.StringNull()
		}
	} else {
		m.Instance = types.StringNull()
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
	if v, ok := obj["key-chain"]; ok {
		_ = v
		if v != "" {
			m.KeyChain = types.StringValue(v)
		} else {
			m.KeyChain = types.StringNull()
		}
	} else {
		m.KeyChain = types.StringNull()
	}
	if v, ok := obj["mode"]; ok {
		_ = v
		if v != "" {
			m.Mode = types.StringValue(v)
		} else {
			m.Mode = types.StringNull()
		}
	} else {
		m.Mode = types.StringNull()
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
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.Password already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["password"]; ok && v != "" {
		_ = v
		if v != "" {
			m.Password = types.StringValue(v)
		} else {
			m.Password = types.StringNull()
		}
	} else if m.Password.IsUnknown() {
		m.Password = types.StringNull()
	}
	if v, ok := obj["poison-reverse"]; ok {
		_ = v
		if v != "" {
			m.PoisonReverse = types.StringValue(v)
		} else {
			m.PoisonReverse = types.StringNull()
		}
	} else {
		m.PoisonReverse = types.StringNull()
	}
	if v, ok := obj["source-addresses"]; ok {
		_ = v
		if v != "" {
			m.SourceAddresses = types.StringValue(v)
		} else {
			m.SourceAddresses = types.StringNull()
		}
	} else {
		m.SourceAddresses = types.StringNull()
	}
	if v, ok := obj["split-horizon"]; ok {
		_ = v
		if v != "" {
			m.SplitHorizon = types.StringValue(v)
		} else {
			m.SplitHorizon = types.StringNull()
		}
	} else {
		m.SplitHorizon = types.StringNull()
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
}
