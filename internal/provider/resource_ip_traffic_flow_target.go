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
	_ resource.Resource                = &IPTrafficFlowTargetResource{}
	_ resource.ResourceWithImportState = &IPTrafficFlowTargetResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPTrafficFlowTargetResource struct {
	reg *client.Registry
}

type IPTrafficFlowTargetModel struct {
	ID                     types.String `tfsdk:"id"`
	V9TemplateTimeout      types.String `tfsdk:"v9_template_timeout"`
	V9TemplateRefresh      types.String `tfsdk:"v9_template_refresh"`
	Disabled               types.Bool   `tfsdk:"disabled"`
	DstAddress             types.String `tfsdk:"dst_address"`
	Port                   types.Int64  `tfsdk:"port"`
	SrcAddress             types.String `tfsdk:"src_address"`
	V9                     types.String `tfsdk:"v9"`
	V9IpfixTemplateRefresh types.Int64  `tfsdk:"v9_ipfix_template_refresh"`
	V9IpfixTemplateTimeout types.Int64  `tfsdk:"v9_ipfix_template_timeout"`
	Version                types.String `tfsdk:"version"`
	Router                 types.String `tfsdk:"router"`
}

func NewIPTrafficFlowTargetResource() resource.Resource { return &IPTrafficFlowTargetResource{} }

func (r *IPTrafficFlowTargetResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_traffic_flow_target"
}

func (r *IPTrafficFlowTargetResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPTrafficFlowTargetResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Discovered; required dst-address must be valid",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"v9_template_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `v9-template-timeout`.",
			},
			"v9_template_refresh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `v9-template-refresh`.",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dst_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"v9": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"v9_ipfix_template_refresh": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"v9_ipfix_template_timeout": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"version": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"1", "5", "9", "ipfix"}...)},
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPTrafficFlowTargetResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPTrafficFlowTargetModel
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
	if !(plan.DstAddress.IsNull() || plan.DstAddress.IsUnknown()) {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = client.FormatInt64(plan.Port.ValueInt64())
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !(plan.Version.IsNull() || plan.Version.IsUnknown()) {
		body["version"] = plan.Version.ValueString()
	}
	if !(plan.V9TemplateRefresh.IsNull() || plan.V9TemplateRefresh.IsUnknown()) {
		body["v9-template-refresh"] = plan.V9TemplateRefresh.ValueString()
	}
	if !(plan.V9TemplateTimeout.IsNull() || plan.V9TemplateTimeout.IsUnknown()) {
		body["v9-template-timeout"] = plan.V9TemplateTimeout.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/traffic-flow/target", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/traffic-flow/target failed", err.Error())
		return
	}
	iPTrafficFlowTargetApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPTrafficFlowTargetResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPTrafficFlowTargetModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/traffic-flow/target", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/traffic-flow/target failed", err.Error())
		return
	}
	iPTrafficFlowTargetApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPTrafficFlowTargetResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPTrafficFlowTargetModel
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
	if !plan.DstAddress.Equal(state.DstAddress) && !plan.DstAddress.IsUnknown() {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = client.FormatInt64(plan.Port.ValueInt64())
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) && !plan.SrcAddress.IsUnknown() {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !plan.Version.Equal(state.Version) && !plan.Version.IsUnknown() {
		body["version"] = plan.Version.ValueString()
	}
	if !plan.V9TemplateRefresh.Equal(state.V9TemplateRefresh) && !plan.V9TemplateRefresh.IsUnknown() {
		body["v9-template-refresh"] = plan.V9TemplateRefresh.ValueString()
	}
	if !plan.V9TemplateTimeout.Equal(state.V9TemplateTimeout) && !plan.V9TemplateTimeout.IsUnknown() {
		body["v9-template-timeout"] = plan.V9TemplateTimeout.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/traffic-flow/target", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/traffic-flow/target failed", err.Error())
			return
		}
		iPTrafficFlowTargetApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPTrafficFlowTargetResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPTrafficFlowTargetModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/traffic-flow/target", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/traffic-flow/target failed", err.Error())
	}
}

func (r *IPTrafficFlowTargetResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPTrafficFlowTargetLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/traffic-flow/target matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPTrafficFlowTargetLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPTrafficFlowTargetLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/traffic-flow/target", id)
}

func iPTrafficFlowTargetApply(ctx context.Context, obj client.Object, m *IPTrafficFlowTargetModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["v9-template-timeout"]; ok && v != "" {
		m.V9TemplateTimeout = types.StringValue(v)
	} else {
		m.V9TemplateTimeout = types.StringNull()
	}
	if v, ok := obj["v9-template-refresh"]; ok && v != "" {
		m.V9TemplateRefresh = types.StringValue(v)
	} else {
		m.V9TemplateRefresh = types.StringNull()
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["dst-address"]; ok {
		_ = v
		if v != "" {
			m.DstAddress = types.StringValue(v)
		} else {
			m.DstAddress = types.StringNull()
		}
	} else {
		m.DstAddress = types.StringNull()
	}
	if v, ok := obj["port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Port = types.Int64Value(n)
		} else {
			m.Port = types.Int64Null()
		}
	} else {
		m.Port = types.Int64Null()
	}
	if v, ok := obj["src-address"]; ok {
		_ = v
		if v != "" {
			m.SrcAddress = types.StringValue(v)
		} else {
			m.SrcAddress = types.StringNull()
		}
	} else {
		m.SrcAddress = types.StringNull()
	}
	if v, ok := obj["v9"]; ok {
		_ = v
		if v != "" {
			m.V9 = types.StringValue(v)
		} else {
			m.V9 = types.StringNull()
		}
	} else {
		m.V9 = types.StringNull()
	}
	if v, ok := obj["v9-ipfix-template-refresh"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.V9IpfixTemplateRefresh = types.Int64Value(n)
		} else {
			m.V9IpfixTemplateRefresh = types.Int64Null()
		}
	} else {
		m.V9IpfixTemplateRefresh = types.Int64Null()
	}
	if v, ok := obj["v9-ipfix-template-timeout"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.V9IpfixTemplateTimeout = types.Int64Value(n)
		} else {
			m.V9IpfixTemplateTimeout = types.Int64Null()
		}
	} else {
		m.V9IpfixTemplateTimeout = types.Int64Null()
	}
	if v, ok := obj["version"]; ok {
		_ = v
		if v != "" {
			m.Version = types.StringValue(v)
		} else {
			m.Version = types.StringNull()
		}
	} else {
		m.Version = types.StringNull()
	}
}
