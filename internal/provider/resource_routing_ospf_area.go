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
	_ resource.Resource                = &RoutingOSPFAreaResource{}
	_ resource.ResourceWithImportState = &RoutingOSPFAreaResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingOSPFAreaResource struct {
	reg *client.Registry
}

type RoutingOSPFAreaModel struct {
	ID             types.String `tfsdk:"id"`
	AreaID         types.String `tfsdk:"area_id"`
	Comment        types.String `tfsdk:"comment"`
	DefaultCost    types.String `tfsdk:"default_cost"`
	Disabled       types.Bool   `tfsdk:"disabled"`
	Dynamic        types.Bool   `tfsdk:"dynamic"`
	Instance       types.String `tfsdk:"instance"`
	Invalid        types.Bool   `tfsdk:"invalid"`
	Name           types.String `tfsdk:"name"`
	NoSummaries    types.Bool   `tfsdk:"no_summaries"`
	NssaTranslator types.String `tfsdk:"nssa_translator"`
	TransitCapable types.Bool   `tfsdk:"transit_capable"`
	Type           types.String `tfsdk:"type"`
	Router         types.String `tfsdk:"router"`
}

func NewRoutingOSPFAreaResource() resource.Resource { return &RoutingOSPFAreaResource{} }

func (r *RoutingOSPFAreaResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_ospf_area"
}

func (r *RoutingOSPFAreaResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *RoutingOSPFAreaResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Auto-test requires a typed-reference precondition (e.g. an existing peer,\ninstance, bridge of the specific kind). The current acc-test generator's\ngeneric data.routeros_interface.all lookup can't satisfy these. Use this\nresource manually with explicit references to a precondition resource\nin your config.\n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"area_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"default_cost": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
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
			"instance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"no_summaries": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nssa_translator": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"transit_capable": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"default", "stub", "nssa"}...)},
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *RoutingOSPFAreaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingOSPFAreaModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AreaID.IsNull() || plan.AreaID.IsUnknown()) {
		body["area-id"] = plan.AreaID.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.DefaultCost.IsNull() || plan.DefaultCost.IsUnknown()) {
		body["default-cost"] = plan.DefaultCost.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Instance.IsNull() || plan.Instance.IsUnknown()) {
		body["instance"] = plan.Instance.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.NoSummaries.IsNull() || plan.NoSummaries.IsUnknown()) {
		body["no-summaries"] = client.FormatBool(plan.NoSummaries.ValueBool())
	}
	if !(plan.NssaTranslator.IsNull() || plan.NssaTranslator.IsUnknown()) {
		body["nssa-translator"] = plan.NssaTranslator.ValueString()
	}
	if !(plan.Type.IsNull() || plan.Type.IsUnknown()) {
		body["type"] = plan.Type.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/ospf/area", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/ospf/area failed", err.Error())
		return
	}
	routingOSPFAreaApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingOSPFAreaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingOSPFAreaModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/ospf/area", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/ospf/area failed", err.Error())
		return
	}
	routingOSPFAreaApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingOSPFAreaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingOSPFAreaModel
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
	if !plan.AreaID.Equal(state.AreaID) {
		body["area-id"] = plan.AreaID.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.DefaultCost.Equal(state.DefaultCost) {
		body["default-cost"] = plan.DefaultCost.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Instance.Equal(state.Instance) {
		body["instance"] = plan.Instance.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.NoSummaries.Equal(state.NoSummaries) {
		body["no-summaries"] = client.FormatBool(plan.NoSummaries.ValueBool())
	}
	if !plan.NssaTranslator.Equal(state.NssaTranslator) {
		body["nssa-translator"] = plan.NssaTranslator.ValueString()
	}
	if !plan.Type.Equal(state.Type) {
		body["type"] = plan.Type.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/ospf/area", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/ospf/area failed", err.Error())
			return
		}
		routingOSPFAreaApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingOSPFAreaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingOSPFAreaModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/ospf/area", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/ospf/area failed", err.Error())
	}
}

func (r *RoutingOSPFAreaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingOSPFAreaLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/ospf/area matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingOSPFAreaLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingOSPFAreaLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/ospf/area", id)
}

func routingOSPFAreaApply(ctx context.Context, obj client.Object, m *RoutingOSPFAreaModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["area-id"]; ok {
		_ = v
		if v != "" {
			m.AreaID = types.StringValue(v)
		} else {
			m.AreaID = types.StringNull()
		}
	} else {
		m.AreaID = types.StringNull()
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
	if v, ok := obj["default-cost"]; ok {
		_ = v
		if v != "" {
			m.DefaultCost = types.StringValue(v)
		} else {
			m.DefaultCost = types.StringNull()
		}
	} else {
		m.DefaultCost = types.StringNull()
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
	if v, ok := obj["no-summaries"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.NoSummaries = types.BoolValue(b)
		} else {
			m.NoSummaries = types.BoolNull()
		}
	} else {
		m.NoSummaries = types.BoolNull()
	}
	if v, ok := obj["nssa-translator"]; ok {
		_ = v
		if v != "" {
			m.NssaTranslator = types.StringValue(v)
		} else {
			m.NssaTranslator = types.StringNull()
		}
	} else {
		m.NssaTranslator = types.StringNull()
	}
	if v, ok := obj["transit-capable"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.TransitCapable = types.BoolValue(b)
		} else {
			m.TransitCapable = types.BoolNull()
		}
	} else {
		m.TransitCapable = types.BoolNull()
	}
	if v, ok := obj["type"]; ok {
		_ = v
		if v != "" {
			m.Type = types.StringValue(v)
		} else {
			m.Type = types.StringNull()
		}
	} else {
		m.Type = types.StringNull()
	}
}
