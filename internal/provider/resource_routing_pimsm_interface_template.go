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
	_ resource.Resource                = &RoutingPimsmInterfaceTemplateResource{}
	_ resource.ResourceWithImportState = &RoutingPimsmInterfaceTemplateResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingPimsmInterfaceTemplateResource struct {
	reg *client.Registry
}

type RoutingPimsmInterfaceTemplateModel struct {
	ID                  types.String `tfsdk:"id"`
	Disabled            types.Bool   `tfsdk:"disabled"`
	HelloDelay          types.String `tfsdk:"hello_delay"`
	HelloPeriod         types.String `tfsdk:"hello_period"`
	Instance            types.String `tfsdk:"instance"`
	Interfaces          types.String `tfsdk:"interfaces"`
	Invalid             types.Bool   `tfsdk:"invalid"`
	JoinPrunePeriod     types.String `tfsdk:"join_prune_period"`
	JoinTrackingSupport types.String `tfsdk:"join_tracking_support"`
	OverrideInterval    types.String `tfsdk:"override_interval"`
	Priority            types.String `tfsdk:"priority"`
	PropagationDelay    types.String `tfsdk:"propagation_delay"`
	SourceAddresses     types.String `tfsdk:"source_addresses"`
	Router              types.String `tfsdk:"router"`
}

func NewRoutingPimsmInterfaceTemplateResource() resource.Resource {
	return &RoutingPimsmInterfaceTemplateResource{}
}

func (r *RoutingPimsmInterfaceTemplateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_pimsm_interface_template"
}

func (r *RoutingPimsmInterfaceTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingPimsmInterfaceTemplateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Discovered; needs pimsm instance",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hello_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"hello_period": schema.StringAttribute{
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
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"join_prune_period": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"join_tracking_support": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"override_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"propagation_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"source_addresses": schema.StringAttribute{
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

func (r *RoutingPimsmInterfaceTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingPimsmInterfaceTemplateModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.HelloDelay.IsNull() || plan.HelloDelay.IsUnknown()) {
		body["hello-delay"] = plan.HelloDelay.ValueString()
	}
	if !(plan.HelloPeriod.IsNull() || plan.HelloPeriod.IsUnknown()) {
		body["hello-period"] = plan.HelloPeriod.ValueString()
	}
	if !(plan.Instance.IsNull() || plan.Instance.IsUnknown()) {
		body["instance"] = plan.Instance.ValueString()
	}
	if !(plan.Interfaces.IsNull() || plan.Interfaces.IsUnknown()) {
		body["interfaces"] = plan.Interfaces.ValueString()
	}
	if !(plan.JoinPrunePeriod.IsNull() || plan.JoinPrunePeriod.IsUnknown()) {
		body["join-prune-period"] = plan.JoinPrunePeriod.ValueString()
	}
	if !(plan.JoinTrackingSupport.IsNull() || plan.JoinTrackingSupport.IsUnknown()) {
		body["join-tracking-support"] = plan.JoinTrackingSupport.ValueString()
	}
	if !(plan.OverrideInterval.IsNull() || plan.OverrideInterval.IsUnknown()) {
		body["override-interval"] = plan.OverrideInterval.ValueString()
	}
	if !(plan.Priority.IsNull() || plan.Priority.IsUnknown()) {
		body["priority"] = plan.Priority.ValueString()
	}
	if !(plan.PropagationDelay.IsNull() || plan.PropagationDelay.IsUnknown()) {
		body["propagation-delay"] = plan.PropagationDelay.ValueString()
	}
	if !(plan.SourceAddresses.IsNull() || plan.SourceAddresses.IsUnknown()) {
		body["source-addresses"] = plan.SourceAddresses.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/pimsm/interface-template", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/pimsm/interface-template failed", err.Error())
		return
	}
	routingPimsmInterfaceTemplateApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingPimsmInterfaceTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingPimsmInterfaceTemplateModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/pimsm/interface-template", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/pimsm/interface-template failed", err.Error())
		return
	}
	routingPimsmInterfaceTemplateApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingPimsmInterfaceTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingPimsmInterfaceTemplateModel
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
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.HelloDelay.Equal(state.HelloDelay) && !plan.HelloDelay.IsUnknown() {
		body["hello-delay"] = plan.HelloDelay.ValueString()
	}
	if !plan.HelloPeriod.Equal(state.HelloPeriod) && !plan.HelloPeriod.IsUnknown() {
		body["hello-period"] = plan.HelloPeriod.ValueString()
	}
	if !plan.Instance.Equal(state.Instance) && !plan.Instance.IsUnknown() {
		body["instance"] = plan.Instance.ValueString()
	}
	if !plan.Interfaces.Equal(state.Interfaces) && !plan.Interfaces.IsUnknown() {
		body["interfaces"] = plan.Interfaces.ValueString()
	}
	if !plan.JoinPrunePeriod.Equal(state.JoinPrunePeriod) && !plan.JoinPrunePeriod.IsUnknown() {
		body["join-prune-period"] = plan.JoinPrunePeriod.ValueString()
	}
	if !plan.JoinTrackingSupport.Equal(state.JoinTrackingSupport) && !plan.JoinTrackingSupport.IsUnknown() {
		body["join-tracking-support"] = plan.JoinTrackingSupport.ValueString()
	}
	if !plan.OverrideInterval.Equal(state.OverrideInterval) && !plan.OverrideInterval.IsUnknown() {
		body["override-interval"] = plan.OverrideInterval.ValueString()
	}
	if !plan.Priority.Equal(state.Priority) && !plan.Priority.IsUnknown() {
		body["priority"] = plan.Priority.ValueString()
	}
	if !plan.PropagationDelay.Equal(state.PropagationDelay) && !plan.PropagationDelay.IsUnknown() {
		body["propagation-delay"] = plan.PropagationDelay.ValueString()
	}
	if !plan.SourceAddresses.Equal(state.SourceAddresses) && !plan.SourceAddresses.IsUnknown() {
		body["source-addresses"] = plan.SourceAddresses.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/pimsm/interface-template", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/pimsm/interface-template failed", err.Error())
			return
		}
		routingPimsmInterfaceTemplateApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingPimsmInterfaceTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingPimsmInterfaceTemplateModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/pimsm/interface-template", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/pimsm/interface-template failed", err.Error())
	}
}

func (r *RoutingPimsmInterfaceTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingPimsmInterfaceTemplateLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/pimsm/interface-template matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingPimsmInterfaceTemplateLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingPimsmInterfaceTemplateLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/pimsm/interface-template", id)
}

func routingPimsmInterfaceTemplateApply(ctx context.Context, obj client.Object, m *RoutingPimsmInterfaceTemplateModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
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
	if v, ok := obj["hello-delay"]; ok {
		_ = v
		if v != "" {
			m.HelloDelay = types.StringValue(v)
		} else {
			m.HelloDelay = types.StringNull()
		}
	} else {
		m.HelloDelay = types.StringNull()
	}
	if v, ok := obj["hello-period"]; ok {
		_ = v
		if v != "" {
			m.HelloPeriod = types.StringValue(v)
		} else {
			m.HelloPeriod = types.StringNull()
		}
	} else {
		m.HelloPeriod = types.StringNull()
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
	if v, ok := obj["join-prune-period"]; ok {
		_ = v
		if v != "" {
			m.JoinPrunePeriod = types.StringValue(v)
		} else {
			m.JoinPrunePeriod = types.StringNull()
		}
	} else {
		m.JoinPrunePeriod = types.StringNull()
	}
	if v, ok := obj["join-tracking-support"]; ok {
		_ = v
		if v != "" {
			m.JoinTrackingSupport = types.StringValue(v)
		} else {
			m.JoinTrackingSupport = types.StringNull()
		}
	} else {
		m.JoinTrackingSupport = types.StringNull()
	}
	if v, ok := obj["override-interval"]; ok {
		_ = v
		if v != "" {
			m.OverrideInterval = types.StringValue(v)
		} else {
			m.OverrideInterval = types.StringNull()
		}
	} else {
		m.OverrideInterval = types.StringNull()
	}
	if v, ok := obj["priority"]; ok {
		_ = v
		if v != "" {
			m.Priority = types.StringValue(v)
		} else {
			m.Priority = types.StringNull()
		}
	} else {
		m.Priority = types.StringNull()
	}
	if v, ok := obj["propagation-delay"]; ok {
		_ = v
		if v != "" {
			m.PropagationDelay = types.StringValue(v)
		} else {
			m.PropagationDelay = types.StringNull()
		}
	} else {
		m.PropagationDelay = types.StringNull()
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
}
