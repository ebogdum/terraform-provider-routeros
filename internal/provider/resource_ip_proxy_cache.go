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
	_ resource.Resource                = &IPProxyCacheResource{}
	_ resource.ResourceWithImportState = &IPProxyCacheResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPProxyCacheResource struct {
	reg *client.Registry
}

type IPProxyCacheModel struct {
	ID         types.String `tfsdk:"id"`
	Method     types.String `tfsdk:"method"`
	LocalPort  types.String `tfsdk:"local_port"`
	DstHost    types.String `tfsdk:"dst_host"`
	Action     types.String `tfsdk:"action"`
	Comment    types.String `tfsdk:"comment"`
	Disabled   types.Bool   `tfsdk:"disabled"`
	DstAddress types.String `tfsdk:"dst_address"`
	DstPort    types.String `tfsdk:"dst_port"`
	Path       types.String `tfsdk:"path"`
	SrcAddress types.String `tfsdk:"src_address"`
	Router     types.String `tfsdk:"router"`
}

func NewIPProxyCacheResource() resource.Resource { return &IPProxyCacheResource{} }

func (r *IPProxyCacheResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_proxy_cache"
}

func (r *IPProxyCacheResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPProxyCacheResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/proxy/cache`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"method": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `method`.",
			},
			"local_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `local-port`.",
			},
			"dst_host": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `dst-host`.",
			},
			"action": schema.StringAttribute{
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
			"dst_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dst_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"path": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_address": schema.StringAttribute{
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

func (r *IPProxyCacheResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPProxyCacheModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Action.IsNull() || plan.Action.IsUnknown()) {
		body["action"] = plan.Action.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.DstAddress.IsNull() || plan.DstAddress.IsUnknown()) {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !(plan.DstPort.IsNull() || plan.DstPort.IsUnknown()) {
		body["dst-port"] = plan.DstPort.ValueString()
	}
	if !(plan.Path.IsNull() || plan.Path.IsUnknown()) {
		body["path"] = plan.Path.ValueString()
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !(plan.DstHost.IsNull() || plan.DstHost.IsUnknown()) {
		body["dst-host"] = plan.DstHost.ValueString()
	}
	if !(plan.LocalPort.IsNull() || plan.LocalPort.IsUnknown()) {
		body["local-port"] = plan.LocalPort.ValueString()
	}
	if !(plan.Method.IsNull() || plan.Method.IsUnknown()) {
		body["method"] = plan.Method.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/proxy/cache", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/proxy/cache failed", err.Error())
		return
	}
	iPProxyCacheApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPProxyCacheResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPProxyCacheModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/proxy/cache", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/proxy/cache failed", err.Error())
		return
	}
	iPProxyCacheApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPProxyCacheResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPProxyCacheModel
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
	if !plan.Action.Equal(state.Action) && !plan.Action.IsUnknown() {
		body["action"] = plan.Action.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DstAddress.Equal(state.DstAddress) && !plan.DstAddress.IsUnknown() {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !plan.DstPort.Equal(state.DstPort) && !plan.DstPort.IsUnknown() {
		body["dst-port"] = plan.DstPort.ValueString()
	}
	if !plan.Path.Equal(state.Path) && !plan.Path.IsUnknown() {
		body["path"] = plan.Path.ValueString()
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) && !plan.SrcAddress.IsUnknown() {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !plan.DstHost.Equal(state.DstHost) && !plan.DstHost.IsUnknown() {
		body["dst-host"] = plan.DstHost.ValueString()
	}
	if !plan.LocalPort.Equal(state.LocalPort) && !plan.LocalPort.IsUnknown() {
		body["local-port"] = plan.LocalPort.ValueString()
	}
	if !plan.Method.Equal(state.Method) && !plan.Method.IsUnknown() {
		body["method"] = plan.Method.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/proxy/cache", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/proxy/cache failed", err.Error())
			return
		}
		iPProxyCacheApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPProxyCacheResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPProxyCacheModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/proxy/cache", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/proxy/cache failed", err.Error())
	}
}

func (r *IPProxyCacheResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPProxyCacheLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/proxy/cache matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPProxyCacheLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPProxyCacheLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/proxy/cache", id)
}

func iPProxyCacheApply(ctx context.Context, obj client.Object, m *IPProxyCacheModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["method"]; ok && v != "" {
		m.Method = types.StringValue(v)
	} else {
		m.Method = types.StringNull()
	}
	if v, ok := obj["local-port"]; ok && v != "" {
		m.LocalPort = types.StringValue(v)
	} else {
		m.LocalPort = types.StringNull()
	}
	if v, ok := obj["dst-host"]; ok && v != "" {
		m.DstHost = types.StringValue(v)
	} else {
		m.DstHost = types.StringNull()
	}
	if v, ok := obj["action"]; ok {
		_ = v
		if v != "" {
			m.Action = types.StringValue(v)
		} else {
			m.Action = types.StringNull()
		}
	} else {
		m.Action = types.StringNull()
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
	if v, ok := obj["dst-port"]; ok {
		_ = v
		if v != "" {
			m.DstPort = types.StringValue(v)
		} else {
			m.DstPort = types.StringNull()
		}
	} else {
		m.DstPort = types.StringNull()
	}
	if v, ok := obj["path"]; ok {
		_ = v
		if v != "" {
			m.Path = types.StringValue(v)
		} else {
			m.Path = types.StringNull()
		}
	} else {
		m.Path = types.StringNull()
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
}
