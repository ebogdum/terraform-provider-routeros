package provider

import (
	"context"
	"fmt"

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
	_ resource.Resource                = &IPProxyResource{}
	_ resource.ResourceWithImportState = &IPProxyResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type IPProxyResource struct {
	reg *client.Registry
}

type IPProxyModel struct {
	ID                   types.String `tfsdk:"id"`
	AlwaysFromCache      types.Bool   `tfsdk:"always_from_cache"`
	Anonymous            types.Bool   `tfsdk:"anonymous"`
	CacheAdministrator   types.String `tfsdk:"cache_administrator"`
	CacheHitDscp         types.Int64  `tfsdk:"cache_hit_dscp"`
	CacheOnDisk          types.Bool   `tfsdk:"cache_on_disk"`
	CachePath            types.String `tfsdk:"cache_path"`
	Enabled              types.Bool   `tfsdk:"enabled"`
	MaxCacheObjectSize   types.Int64  `tfsdk:"max_cache_object_size"`
	MaxCacheSize         types.String `tfsdk:"max_cache_size"`
	MaxClientConnections types.Int64  `tfsdk:"max_client_connections"`
	MaxFreshTime         types.String `tfsdk:"max_fresh_time"`
	MaxServerConnections types.Int64  `tfsdk:"max_server_connections"`
	ParentProxy          types.String `tfsdk:"parent_proxy"`
	ParentProxyPort      types.Int64  `tfsdk:"parent_proxy_port"`
	Port                 types.Int64  `tfsdk:"port"`
	SerializeConnections types.Bool   `tfsdk:"serialize_connections"`
	SrcAddress           types.String `tfsdk:"src_address"`
	Router               types.String `tfsdk:"router"`
}

func NewIPProxyResource() resource.Resource { return &IPProxyResource{} }

func (r *IPProxyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_proxy"
}

func (r *IPProxyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPProxyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/proxy`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"always_from_cache": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"anonymous": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"cache_administrator": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"cache_hit_dscp": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"cache_on_disk": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"cache_path": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"enabled": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"max_cache_object_size": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"max_cache_size": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"max_client_connections": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"max_fresh_time": schema.StringAttribute{Optional: true, Computed: true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"max_server_connections": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"parent_proxy": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"parent_proxy_port": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"port": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"serialize_connections": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"src_address": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPProxyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPProxyModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPProxyUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPProxyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPProxyModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	iPProxyUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPProxyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPProxyModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/proxy")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/proxy failed", err.Error())
		return
	}
	iPProxyApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/proxy", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPProxyResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *IPProxyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/proxy" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/ip/proxy", types.StringValue(routerName))))...)
}

func iPProxyUpsert(ctx context.Context, reg *client.Registry, plan *IPProxyModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AlwaysFromCache.IsNull() || plan.AlwaysFromCache.IsUnknown()) {
		body["always-from-cache"] = client.FormatBool(plan.AlwaysFromCache.ValueBool())
	}
	if !(plan.Anonymous.IsNull() || plan.Anonymous.IsUnknown()) {
		body["anonymous"] = client.FormatBool(plan.Anonymous.ValueBool())
	}
	if !(plan.CacheAdministrator.IsNull() || plan.CacheAdministrator.IsUnknown()) {
		body["cache-administrator"] = plan.CacheAdministrator.ValueString()
	}
	if !(plan.CacheHitDscp.IsNull() || plan.CacheHitDscp.IsUnknown()) {
		body["cache-hit-dscp"] = client.FormatInt64(plan.CacheHitDscp.ValueInt64())
	}
	if !(plan.CacheOnDisk.IsNull() || plan.CacheOnDisk.IsUnknown()) {
		body["cache-on-disk"] = client.FormatBool(plan.CacheOnDisk.ValueBool())
	}
	if !(plan.CachePath.IsNull() || plan.CachePath.IsUnknown()) {
		body["cache-path"] = plan.CachePath.ValueString()
	}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	if !(plan.MaxCacheObjectSize.IsNull() || plan.MaxCacheObjectSize.IsUnknown()) {
		body["max-cache-object-size"] = client.FormatInt64(plan.MaxCacheObjectSize.ValueInt64())
	}
	if !(plan.MaxCacheSize.IsNull() || plan.MaxCacheSize.IsUnknown()) {
		body["max-cache-size"] = plan.MaxCacheSize.ValueString()
	}
	if !(plan.MaxClientConnections.IsNull() || plan.MaxClientConnections.IsUnknown()) {
		body["max-client-connections"] = client.FormatInt64(plan.MaxClientConnections.ValueInt64())
	}
	if !(plan.MaxFreshTime.IsNull() || plan.MaxFreshTime.IsUnknown()) {
		body["max-fresh-time"] = plan.MaxFreshTime.ValueString()
	}
	if !(plan.MaxServerConnections.IsNull() || plan.MaxServerConnections.IsUnknown()) {
		body["max-server-connections"] = client.FormatInt64(plan.MaxServerConnections.ValueInt64())
	}
	if !(plan.ParentProxy.IsNull() || plan.ParentProxy.IsUnknown()) {
		body["parent-proxy"] = plan.ParentProxy.ValueString()
	}
	if !(plan.ParentProxyPort.IsNull() || plan.ParentProxyPort.IsUnknown()) {
		body["parent-proxy-port"] = client.FormatInt64(plan.ParentProxyPort.ValueInt64())
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = client.FormatInt64(plan.Port.ValueInt64())
	}
	if !(plan.SerializeConnections.IsNull() || plan.SerializeConnections.IsUnknown()) {
		body["serialize-connections"] = client.FormatBool(plan.SerializeConnections.ValueBool())
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/ip/proxy", body)
	if err != nil {
		diags.AddError("Upsert /ip/proxy failed", err.Error())
		return
	}
	iPProxyApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/proxy", plan.Router))
}

func iPProxyApply(ctx context.Context, obj client.Object, m *IPProxyModel) {
	_ = ctx
	if v, ok := obj["always-from-cache"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AlwaysFromCache = types.BoolValue(b)
		} else {
			m.AlwaysFromCache = types.BoolNull()
		}
	}
	if v, ok := obj["anonymous"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Anonymous = types.BoolValue(b)
		} else {
			m.Anonymous = types.BoolNull()
		}
	}
	if v, ok := obj["cache-administrator"]; ok {
		_ = v
		if v != "" {
			m.CacheAdministrator = types.StringValue(v)
		} else {
			m.CacheAdministrator = types.StringNull()
		}
	}
	if v, ok := obj["cache-hit-dscp"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.CacheHitDscp = types.Int64Value(n)
		} else {
			m.CacheHitDscp = types.Int64Null()
		}
	}
	if v, ok := obj["cache-on-disk"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.CacheOnDisk = types.BoolValue(b)
		} else {
			m.CacheOnDisk = types.BoolNull()
		}
	}
	if v, ok := obj["cache-path"]; ok {
		_ = v
		if v != "" {
			m.CachePath = types.StringValue(v)
		} else {
			m.CachePath = types.StringNull()
		}
	}
	if v, ok := obj["enabled"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Enabled = types.BoolValue(b)
		} else {
			m.Enabled = types.BoolNull()
		}
	}
	if v, ok := obj["max-cache-object-size"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MaxCacheObjectSize = types.Int64Value(n)
		} else {
			m.MaxCacheObjectSize = types.Int64Null()
		}
	}
	if v, ok := obj["max-cache-size"]; ok {
		_ = v
		if v != "" {
			m.MaxCacheSize = types.StringValue(v)
		} else {
			m.MaxCacheSize = types.StringNull()
		}
	}
	if v, ok := obj["max-client-connections"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MaxClientConnections = types.Int64Value(n)
		} else {
			m.MaxClientConnections = types.Int64Null()
		}
	}
	if v, ok := obj["max-fresh-time"]; ok {
		_ = v
		if v != "" {
			m.MaxFreshTime = types.StringValue(v)
		} else {
			m.MaxFreshTime = types.StringNull()
		}
	}
	if v, ok := obj["max-server-connections"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MaxServerConnections = types.Int64Value(n)
		} else {
			m.MaxServerConnections = types.Int64Null()
		}
	}
	if v, ok := obj["parent-proxy"]; ok {
		_ = v
		if v != "" {
			m.ParentProxy = types.StringValue(v)
		} else {
			m.ParentProxy = types.StringNull()
		}
	}
	if v, ok := obj["parent-proxy-port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.ParentProxyPort = types.Int64Value(n)
		} else {
			m.ParentProxyPort = types.Int64Null()
		}
	}
	if v, ok := obj["port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Port = types.Int64Value(n)
		} else {
			m.Port = types.Int64Null()
		}
	}
	if v, ok := obj["serialize-connections"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.SerializeConnections = types.BoolValue(b)
		} else {
			m.SerializeConnections = types.BoolNull()
		}
	}
	if v, ok := obj["src-address"]; ok {
		_ = v
		if v != "" {
			m.SrcAddress = types.StringValue(v)
		} else {
			m.SrcAddress = types.StringNull()
		}
	}
}
