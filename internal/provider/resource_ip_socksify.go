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
	_ resource.Resource                = &IPSocksifyResource{}
	_ resource.ResourceWithImportState = &IPSocksifyResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPSocksifyResource struct {
	reg *client.Registry
}

type IPSocksifyModel struct {
	ID                types.String `tfsdk:"id"`
	Socks5User        types.String `tfsdk:"socks5_user"`
	Socks5Server      types.String `tfsdk:"socks5_server"`
	Socks5Port        types.String `tfsdk:"socks5_port"`
	Socks5Password    types.String `tfsdk:"socks5_password"`
	ConnectionTimeout types.String `tfsdk:"connection_timeout"`
	Disabled          types.Bool   `tfsdk:"disabled"`
	Name              types.String `tfsdk:"name"`
	Port              types.Int64  `tfsdk:"port"`
	Router            types.String `tfsdk:"router"`
}

func NewIPSocksifyResource() resource.Resource { return &IPSocksifyResource{} }

func (r *IPSocksifyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_socksify"
}

func (r *IPSocksifyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPSocksifyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/socksify`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"socks5_user": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `socks5-user`.",
			},
			"socks5_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `socks5-server`.",
			},
			"socks5_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `socks5-port`.",
			},
			"socks5_password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "RouterOS `socks5-password`.",
			},
			"connection_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `connection-timeout`.",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"port": schema.Int64Attribute{
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

func (r *IPSocksifyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPSocksifyModel
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
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = client.FormatInt64(plan.Port.ValueInt64())
	}
	if !(plan.ConnectionTimeout.IsNull() || plan.ConnectionTimeout.IsUnknown()) {
		body["connection-timeout"] = plan.ConnectionTimeout.ValueString()
	}
	if !(plan.Socks5Password.IsNull() || plan.Socks5Password.IsUnknown()) {
		body["socks5-password"] = plan.Socks5Password.ValueString()
	}
	if !(plan.Socks5Port.IsNull() || plan.Socks5Port.IsUnknown()) {
		body["socks5-port"] = plan.Socks5Port.ValueString()
	}
	if !(plan.Socks5Server.IsNull() || plan.Socks5Server.IsUnknown()) {
		body["socks5-server"] = plan.Socks5Server.ValueString()
	}
	if !(plan.Socks5User.IsNull() || plan.Socks5User.IsUnknown()) {
		body["socks5-user"] = plan.Socks5User.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/socksify", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/socksify failed", err.Error())
		return
	}
	iPSocksifyApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPSocksifyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPSocksifyModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/socksify", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/socksify failed", err.Error())
		return
	}
	iPSocksifyApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPSocksifyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPSocksifyModel
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
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = client.FormatInt64(plan.Port.ValueInt64())
	}
	if !plan.ConnectionTimeout.Equal(state.ConnectionTimeout) && !plan.ConnectionTimeout.IsUnknown() {
		body["connection-timeout"] = plan.ConnectionTimeout.ValueString()
	}
	if !plan.Socks5Password.Equal(state.Socks5Password) && !plan.Socks5Password.IsUnknown() {
		body["socks5-password"] = plan.Socks5Password.ValueString()
	}
	if !plan.Socks5Port.Equal(state.Socks5Port) && !plan.Socks5Port.IsUnknown() {
		body["socks5-port"] = plan.Socks5Port.ValueString()
	}
	if !plan.Socks5Server.Equal(state.Socks5Server) && !plan.Socks5Server.IsUnknown() {
		body["socks5-server"] = plan.Socks5Server.ValueString()
	}
	if !plan.Socks5User.Equal(state.Socks5User) && !plan.Socks5User.IsUnknown() {
		body["socks5-user"] = plan.Socks5User.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/socksify", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/socksify failed", err.Error())
			return
		}
		iPSocksifyApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPSocksifyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPSocksifyModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/socksify", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/socksify failed", err.Error())
	}
}

func (r *IPSocksifyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPSocksifyLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/socksify matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPSocksifyLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPSocksifyLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/socksify", id)
}

func iPSocksifyApply(ctx context.Context, obj client.Object, m *IPSocksifyModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["socks5-user"]; ok && v != "" {
		m.Socks5User = types.StringValue(v)
	} else {
		m.Socks5User = types.StringNull()
	}
	if v, ok := obj["socks5-server"]; ok && v != "" {
		m.Socks5Server = types.StringValue(v)
	} else {
		m.Socks5Server = types.StringNull()
	}
	if v, ok := obj["socks5-port"]; ok && v != "" {
		m.Socks5Port = types.StringValue(v)
	} else {
		m.Socks5Port = types.StringNull()
	}
	if v, ok := obj["socks5-password"]; ok && v != "" {
		m.Socks5Password = types.StringValue(v)
	} else {
		m.Socks5Password = types.StringNull()
	}
	if v, ok := obj["connection-timeout"]; ok && v != "" {
		m.ConnectionTimeout = types.StringValue(v)
	} else {
		m.ConnectionTimeout = types.StringNull()
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
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
}
