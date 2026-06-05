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
	_ resource.Resource                = &ToolNetwatchResource{}
	_ resource.ResourceWithImportState = &ToolNetwatchResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type ToolNetwatchResource struct {
	reg *client.Registry
}

type ToolNetwatchModel struct {
	ID          types.String `tfsdk:"id"`
	Certificate types.String `tfsdk:"certificate"`
	Comment     types.String `tfsdk:"comment"`
	Disabled    types.Bool   `tfsdk:"disabled"`
	DNSServer   types.String `tfsdk:"dns_server"`
	Host        types.String `tfsdk:"host"`
	Interval    types.String `tfsdk:"interval"`
	Name        types.String `tfsdk:"name"`
	Port        types.String `tfsdk:"port"`
	SrcAddress  types.String `tfsdk:"src_address"`
	Timeout     types.String `tfsdk:"timeout"`
	Ttl         types.String `tfsdk:"ttl"`
	Type        types.String `tfsdk:"type"`
	Router      types.String `tfsdk:"router"`
}

func NewToolNetwatchResource() resource.Resource { return &ToolNetwatchResource{} }

func (r *ToolNetwatchResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_netwatch"
}

func (r *ToolNetwatchResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *ToolNetwatchResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/tool/netwatch`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"certificate": schema.StringAttribute{
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
			"dns_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"host": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"interval": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"type": schema.StringAttribute{
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

func (r *ToolNetwatchResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolNetwatchModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Certificate.IsNull() || plan.Certificate.IsUnknown()) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.DNSServer.IsNull() || plan.DNSServer.IsUnknown()) {
		body["dns-server"] = plan.DNSServer.ValueString()
	}
	if !(plan.Host.IsNull() || plan.Host.IsUnknown()) {
		body["host"] = plan.Host.ValueString()
	}
	if !(plan.Interval.IsNull() || plan.Interval.IsUnknown()) {
		body["interval"] = plan.Interval.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !(plan.Timeout.IsNull() || plan.Timeout.IsUnknown()) {
		body["timeout"] = plan.Timeout.ValueString()
	}
	if !(plan.Ttl.IsNull() || plan.Ttl.IsUnknown()) {
		body["ttl"] = plan.Ttl.ValueString()
	}
	if !(plan.Type.IsNull() || plan.Type.IsUnknown()) {
		body["type"] = plan.Type.ValueString()
	}
	obj, err := c.Add(ctx, "/tool/netwatch", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /tool/netwatch failed", err.Error())
		return
	}
	toolNetwatchApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolNetwatchResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolNetwatchModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/tool/netwatch", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /tool/netwatch failed", err.Error())
		return
	}
	toolNetwatchApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolNetwatchResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ToolNetwatchModel
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
	if !plan.Certificate.Equal(state.Certificate) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DNSServer.Equal(state.DNSServer) {
		body["dns-server"] = plan.DNSServer.ValueString()
	}
	if !plan.Host.Equal(state.Host) {
		body["host"] = plan.Host.ValueString()
	}
	if !plan.Interval.Equal(state.Interval) {
		body["interval"] = plan.Interval.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Port.Equal(state.Port) {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !plan.Timeout.Equal(state.Timeout) {
		body["timeout"] = plan.Timeout.ValueString()
	}
	if !plan.Ttl.Equal(state.Ttl) {
		body["ttl"] = plan.Ttl.ValueString()
	}
	if !plan.Type.Equal(state.Type) {
		body["type"] = plan.Type.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/tool/netwatch", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /tool/netwatch failed", err.Error())
			return
		}
		toolNetwatchApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolNetwatchResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ToolNetwatchModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/tool/netwatch", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /tool/netwatch failed", err.Error())
	}
}

func (r *ToolNetwatchResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import formats accepted:
	//   *<id>                            -> bare RouterOS .id on the default router
	//   <router>/*<id>                   -> .id on the named router
	//   <router>/<naturalkey>            -> resolved via List + filter
	//   <naturalkey>                     -> resolved on the default router
	id := req.ID
	routerName := ""
	if i := strings.Index(id, "/"); i > 0 && !strings.HasPrefix(id, "*") {
		routerName, id = id[:i], id[i+1:]
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	if strings.HasPrefix(id, "*") {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(id))...)
		return
	}
	c := pickClient(r.reg, types.StringValue(routerName), &resp.Diagnostics)
	if c == nil {
		return
	}
	rows, err := toolNetwatchLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /tool/netwatch matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// toolNetwatchLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func toolNetwatchLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/tool/netwatch", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func toolNetwatchApply(ctx context.Context, obj client.Object, m *ToolNetwatchModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["certificate"]; ok {
		_ = v
		if v != "" {
			m.Certificate = types.StringValue(v)
		} else {
			m.Certificate = types.StringNull()
		}
	} else {
		m.Certificate = types.StringNull()
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
	if v, ok := obj["dns-server"]; ok {
		_ = v
		if v != "" {
			m.DNSServer = types.StringValue(v)
		} else {
			m.DNSServer = types.StringNull()
		}
	} else {
		m.DNSServer = types.StringNull()
	}
	if v, ok := obj["host"]; ok {
		_ = v
		if v != "" {
			m.Host = types.StringValue(v)
		} else {
			m.Host = types.StringNull()
		}
	} else {
		m.Host = types.StringNull()
	}
	if v, ok := obj["interval"]; ok {
		_ = v
		if v != "" {
			m.Interval = types.StringValue(v)
		} else {
			m.Interval = types.StringNull()
		}
	} else {
		m.Interval = types.StringNull()
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
	if v, ok := obj["port"]; ok {
		_ = v
		if v != "" {
			m.Port = types.StringValue(v)
		} else {
			m.Port = types.StringNull()
		}
	} else {
		m.Port = types.StringNull()
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
	if v, ok := obj["timeout"]; ok {
		_ = v
		if v != "" {
			m.Timeout = types.StringValue(v)
		} else {
			m.Timeout = types.StringNull()
		}
	} else {
		m.Timeout = types.StringNull()
	}
	if v, ok := obj["ttl"]; ok {
		_ = v
		if v != "" {
			m.Ttl = types.StringValue(v)
		} else {
			m.Ttl = types.StringNull()
		}
	} else {
		m.Ttl = types.StringNull()
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
