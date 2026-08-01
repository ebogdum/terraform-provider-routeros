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
	_ resource.Resource                = &RoutingBfdConfigurationResource{}
	_ resource.ResourceWithImportState = &RoutingBfdConfigurationResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingBfdConfigurationResource struct {
	reg *client.Registry
}

type RoutingBfdConfigurationModel struct {
	ID          types.String `tfsdk:"id"`
	AddressList types.String `tfsdk:"address_list"`
	Addresses   types.String `tfsdk:"addresses"`
	Comment     types.String `tfsdk:"comment"`
	Disabled    types.Bool   `tfsdk:"disabled"`
	ForbidBfd   types.String `tfsdk:"forbid_bfd"`
	Inactive    types.Bool   `tfsdk:"inactive"`
	Interfaces  types.String `tfsdk:"interfaces"`
	MinRx       types.String `tfsdk:"min_rx"`
	MinTx       types.String `tfsdk:"min_tx"`
	Multiplier  types.String `tfsdk:"multiplier"`
	Vrf         types.String `tfsdk:"vrf"`
	Router      types.String `tfsdk:"router"`
}

func NewRoutingBfdConfigurationResource() resource.Resource {
	return &RoutingBfdConfigurationResource{}
}

func (r *RoutingBfdConfigurationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_bfd_configuration"
}

func (r *RoutingBfdConfigurationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingBfdConfigurationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/routing/bfd/configuration`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"address_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"addresses": schema.StringAttribute{
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
			"forbid_bfd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"inactive": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"interfaces": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"min_rx": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"min_tx": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"multiplier": schema.StringAttribute{
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

func (r *RoutingBfdConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingBfdConfigurationModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AddressList.IsNull() || plan.AddressList.IsUnknown()) {
		body["address-list"] = plan.AddressList.ValueString()
	}
	if !(plan.Addresses.IsNull() || plan.Addresses.IsUnknown()) {
		body["addresses"] = plan.Addresses.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.ForbidBfd.IsNull() || plan.ForbidBfd.IsUnknown()) {
		body["forbid-bfd"] = plan.ForbidBfd.ValueString()
	}
	if !(plan.Interfaces.IsNull() || plan.Interfaces.IsUnknown()) {
		body["interfaces"] = plan.Interfaces.ValueString()
	}
	if !(plan.MinRx.IsNull() || plan.MinRx.IsUnknown()) {
		body["min-rx"] = plan.MinRx.ValueString()
	}
	if !(plan.MinTx.IsNull() || plan.MinTx.IsUnknown()) {
		body["min-tx"] = plan.MinTx.ValueString()
	}
	if !(plan.Multiplier.IsNull() || plan.Multiplier.IsUnknown()) {
		body["multiplier"] = plan.Multiplier.ValueString()
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/bfd/configuration", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/bfd/configuration failed", err.Error())
		return
	}
	routingBfdConfigurationApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingBfdConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingBfdConfigurationModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/bfd/configuration", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/bfd/configuration failed", err.Error())
		return
	}
	routingBfdConfigurationApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingBfdConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingBfdConfigurationModel
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
	if !plan.AddressList.Equal(state.AddressList) && !plan.AddressList.IsUnknown() {
		body["address-list"] = plan.AddressList.ValueString()
	}
	if !plan.Addresses.Equal(state.Addresses) && !plan.Addresses.IsUnknown() {
		body["addresses"] = plan.Addresses.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.ForbidBfd.Equal(state.ForbidBfd) && !plan.ForbidBfd.IsUnknown() {
		body["forbid-bfd"] = plan.ForbidBfd.ValueString()
	}
	if !plan.Interfaces.Equal(state.Interfaces) && !plan.Interfaces.IsUnknown() {
		body["interfaces"] = plan.Interfaces.ValueString()
	}
	if !plan.MinRx.Equal(state.MinRx) && !plan.MinRx.IsUnknown() {
		body["min-rx"] = plan.MinRx.ValueString()
	}
	if !plan.MinTx.Equal(state.MinTx) && !plan.MinTx.IsUnknown() {
		body["min-tx"] = plan.MinTx.ValueString()
	}
	if !plan.Multiplier.Equal(state.Multiplier) && !plan.Multiplier.IsUnknown() {
		body["multiplier"] = plan.Multiplier.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) && !plan.Vrf.IsUnknown() {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/bfd/configuration", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/bfd/configuration failed", err.Error())
			return
		}
		routingBfdConfigurationApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingBfdConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingBfdConfigurationModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/bfd/configuration", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/bfd/configuration failed", err.Error())
	}
}

func (r *RoutingBfdConfigurationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingBfdConfigurationLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/bfd/configuration matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingBfdConfigurationLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingBfdConfigurationLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/bfd/configuration", id)
}

func routingBfdConfigurationApply(ctx context.Context, obj client.Object, m *RoutingBfdConfigurationModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["address-list"]; ok {
		_ = v
		if v != "" {
			m.AddressList = types.StringValue(v)
		} else {
			m.AddressList = types.StringNull()
		}
	} else {
		m.AddressList = types.StringNull()
	}
	if v, ok := obj["addresses"]; ok {
		_ = v
		if v != "" {
			m.Addresses = types.StringValue(v)
		} else {
			m.Addresses = types.StringNull()
		}
	} else {
		m.Addresses = types.StringNull()
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
	if v, ok := obj["forbid-bfd"]; ok {
		_ = v
		if v != "" {
			m.ForbidBfd = types.StringValue(v)
		} else {
			m.ForbidBfd = types.StringNull()
		}
	} else {
		m.ForbidBfd = types.StringNull()
	}
	if v, ok := obj["inactive"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Inactive = types.BoolValue(b)
		} else {
			m.Inactive = types.BoolNull()
		}
	} else {
		m.Inactive = types.BoolNull()
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
	if v, ok := obj["min-rx"]; ok {
		_ = v
		if v != "" {
			m.MinRx = types.StringValue(v)
		} else {
			m.MinRx = types.StringNull()
		}
	} else {
		m.MinRx = types.StringNull()
	}
	if v, ok := obj["min-tx"]; ok {
		_ = v
		if v != "" {
			m.MinTx = types.StringValue(v)
		} else {
			m.MinTx = types.StringNull()
		}
	} else {
		m.MinTx = types.StringNull()
	}
	if v, ok := obj["multiplier"]; ok {
		_ = v
		if v != "" {
			m.Multiplier = types.StringValue(v)
		} else {
			m.Multiplier = types.StringNull()
		}
	} else {
		m.Multiplier = types.StringNull()
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
