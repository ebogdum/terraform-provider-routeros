package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &IPIPsecKeyQKDResource{}
	_ resource.ResourceWithImportState = &IPIPsecKeyQKDResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPIPsecKeyQKDResource struct {
	reg *client.Registry
}

type IPIPsecKeyQKDModel struct {
	ID          types.String `tfsdk:"id"`
	Address     types.String `tfsdk:"address"`
	CacheSize   types.Int64  `tfsdk:"cache_size"`
	Certificate types.String `tfsdk:"certificate"`
	Enabled     types.Bool   `tfsdk:"enabled"`
	KeySize     types.Int64  `tfsdk:"key_size"`
	KmeID       types.String `tfsdk:"kme_id"`
	PeerSaeID   types.String `tfsdk:"peer_sae_id"`
	Router      types.String `tfsdk:"router"`
}

func NewIPIPsecKeyQKDResource() resource.Resource { return &IPIPsecKeyQKDResource{} }

func (r *IPIPsecKeyQKDResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_ipsec_key_qkd"
}

func (r *IPIPsecKeyQKDResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPIPsecKeyQKDResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/ipsec/key/qkd`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `address`.",
			},
			"cache_size": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cache-size`.",
			},
			"certificate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `certificate`.",
			},
			"enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `enabled`.",
			},
			"key_size": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `key-size`.",
			},
			"kme_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `kme-id`.",
			},
			"peer_sae_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `peer-sae-id`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPIPsecKeyQKDResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPIPsecKeyQKDModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	iPIPsecKeyQKDUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIPsecKeyQKDResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPIPsecKeyQKDModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/ip/ipsec/key/qkd")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/ipsec/key/qkd failed", err.Error())
		return
	}
	iPIPsecKeyQKDApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/ip/ipsec/key/qkd", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPIPsecKeyQKDResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan IPIPsecKeyQKDModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	var state IPIPsecKeyQKDModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	iPIPsecKeyQKDUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIPsecKeyQKDResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menu: not removable. Just drop the state.
}

func (r *IPIPsecKeyQKDResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/ip/ipsec/key/qkd" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"),
		types.StringValue(stateIDFor("/ip/ipsec/key/qkd", types.StringValue(routerName))))...)
}

func iPIPsecKeyQKDUpsert(ctx context.Context, reg *client.Registry, plan, state *IPIPsecKeyQKDModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Address.IsNull() || plan.Address.IsUnknown()) && (state == nil || !plan.Address.Equal(state.Address)) {
		body["address"] = plan.Address.ValueString()
	}
	if !(plan.CacheSize.IsNull() || plan.CacheSize.IsUnknown()) && (state == nil || !plan.CacheSize.Equal(state.CacheSize)) {
		body["cache-size"] = client.FormatInt64(plan.CacheSize.ValueInt64())
	}
	if !(plan.Certificate.IsNull() || plan.Certificate.IsUnknown()) && (state == nil || !plan.Certificate.Equal(state.Certificate)) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	if !(plan.KeySize.IsNull() || plan.KeySize.IsUnknown()) && (state == nil || !plan.KeySize.Equal(state.KeySize)) {
		body["key-size"] = client.FormatInt64(plan.KeySize.ValueInt64())
	}
	if !(plan.KmeID.IsNull() || plan.KmeID.IsUnknown()) && (state == nil || !plan.KmeID.Equal(state.KmeID)) {
		body["kme-id"] = plan.KmeID.ValueString()
	}
	if !(plan.PeerSaeID.IsNull() || plan.PeerSaeID.IsUnknown()) && (state == nil || !plan.PeerSaeID.Equal(state.PeerSaeID)) {
		body["peer-sae-id"] = plan.PeerSaeID.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/ip/ipsec/key/qkd", body)
	if err != nil {
		diags.AddError("Upsert /ip/ipsec/key/qkd failed", err.Error())
		return
	}
	iPIPsecKeyQKDApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/ip/ipsec/key/qkd", plan.Router))
}

func iPIPsecKeyQKDApply(ctx context.Context, obj client.Object, m *IPIPsecKeyQKDModel) {
	_ = ctx
	if v, ok := obj["address"]; ok && v != "" {
		m.Address = types.StringValue(v)
	} else {
		m.Address = types.StringNull()
	}
	if v, ok := obj["cache-size"]; ok {
		if n, err := client.ParseInt64(v); err == nil {
			m.CacheSize = types.Int64Value(n)
		} else {
			m.CacheSize = types.Int64Null()
		}
	} else {
		m.CacheSize = types.Int64Null()
	}
	if v, ok := obj["certificate"]; ok && v != "" {
		m.Certificate = types.StringValue(v)
	} else {
		m.Certificate = types.StringNull()
	}
	if v, ok := obj["enabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Enabled = types.BoolValue(b)
		} else {
			m.Enabled = types.BoolNull()
		}
	} else {
		m.Enabled = types.BoolNull()
	}
	if v, ok := obj["key-size"]; ok {
		if n, err := client.ParseInt64(v); err == nil {
			m.KeySize = types.Int64Value(n)
		} else {
			m.KeySize = types.Int64Null()
		}
	} else {
		m.KeySize = types.Int64Null()
	}
	if v, ok := obj["kme-id"]; ok && v != "" {
		m.KmeID = types.StringValue(v)
	} else {
		m.KmeID = types.StringNull()
	}
	if v, ok := obj["peer-sae-id"]; ok && v != "" {
		m.PeerSaeID = types.StringValue(v)
	} else {
		m.PeerSaeID = types.StringNull()
	}
}
