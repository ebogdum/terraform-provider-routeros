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
	_ resource.Resource                = &IPIpsecModeConfigResource{}
	_ resource.ResourceWithImportState = &IPIpsecModeConfigResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPIpsecModeConfigResource struct {
	reg *client.Registry
}

type IPIpsecModeConfigModel struct {
	ID                  types.String `tfsdk:"id"`
	Address             types.String `tfsdk:"address"`
	AddressPool         types.String `tfsdk:"address_pool"`
	AddressPrefixLength types.Int64  `tfsdk:"address_prefix_length"`
	ConnectionMark      types.String `tfsdk:"connection_mark"`
	Default             types.Bool   `tfsdk:"default"`
	Name                types.String `tfsdk:"name"`
	Nonresp             types.String `tfsdk:"nonresp"`
	Pool                types.String `tfsdk:"pool"`
	Resp                types.String `tfsdk:"resp"`
	Responder           types.Bool   `tfsdk:"responder"`
	Sdns                types.String `tfsdk:"sdns"`
	SplitDNS            types.String `tfsdk:"split_dns"`
	SplitInclude        types.String `tfsdk:"split_include"`
	SrcAddressList      types.String `tfsdk:"src_address_list"`
	StaticDNS           types.String `tfsdk:"static_dns"`
	SystemDNS           types.Bool   `tfsdk:"system_dns"`
	UseResponderDNS     types.String `tfsdk:"use_responder_dns"`
	Router              types.String `tfsdk:"router"`
}

func NewIPIpsecModeConfigResource() resource.Resource { return &IPIpsecModeConfigResource{} }

func (r *IPIpsecModeConfigResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_ipsec_mode_config"
}

func (r *IPIpsecModeConfigResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPIpsecModeConfigResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/ipsec/mode-config`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"address_pool": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"address_prefix_length": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"connection_mark": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"default": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"nonresp": schema.StringAttribute{
				Optional:           true,
				Computed:           true,
				Description:        "",
				DeprecationMessage: "Not a RouterOS REST property (WebFig-only spelling of `responder`); RouterOS rejects it. Read-only and ignored on write - use `responder`.",
			},
			"pool": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"resp": schema.StringAttribute{
				Optional:           true,
				Computed:           true,
				Description:        "",
				DeprecationMessage: "Not a RouterOS REST property (WebFig-only spelling of `responder`); RouterOS rejects it. Read-only and ignored on write - use `responder`.",
			},
			"responder": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sdns": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"split_dns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"split_include": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_address_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"static_dns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"system_dns": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"use_responder_dns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"no", "yes", "exclusively"}...)},
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPIpsecModeConfigResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPIpsecModeConfigModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Address.IsNull() || plan.Address.IsUnknown()) {
		body["address"] = plan.Address.ValueString()
	}
	if !(plan.AddressPool.IsNull() || plan.AddressPool.IsUnknown()) {
		body["address-pool"] = plan.AddressPool.ValueString()
	}
	if !(plan.AddressPrefixLength.IsNull() || plan.AddressPrefixLength.IsUnknown()) {
		body["address-prefix-length"] = client.FormatInt64(plan.AddressPrefixLength.ValueInt64())
	}
	if !(plan.ConnectionMark.IsNull() || plan.ConnectionMark.IsUnknown()) {
		body["connection-mark"] = plan.ConnectionMark.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Responder.IsNull() || plan.Responder.IsUnknown()) {
		body["responder"] = client.FormatBool(plan.Responder.ValueBool())
	}
	if !(plan.SplitDNS.IsNull() || plan.SplitDNS.IsUnknown()) {
		body["split-dns"] = plan.SplitDNS.ValueString()
	}
	if !(plan.SplitInclude.IsNull() || plan.SplitInclude.IsUnknown()) {
		body["split-include"] = plan.SplitInclude.ValueString()
	}
	if !(plan.SrcAddressList.IsNull() || plan.SrcAddressList.IsUnknown()) {
		body["src-address-list"] = plan.SrcAddressList.ValueString()
	}
	if !(plan.StaticDNS.IsNull() || plan.StaticDNS.IsUnknown()) {
		body["static-dns"] = plan.StaticDNS.ValueString()
	}
	if !(plan.SystemDNS.IsNull() || plan.SystemDNS.IsUnknown()) {
		body["system-dns"] = client.FormatBool(plan.SystemDNS.ValueBool())
	}
	if !(plan.UseResponderDNS.IsNull() || plan.UseResponderDNS.IsUnknown()) {
		body["use-responder-dns"] = plan.UseResponderDNS.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/ipsec/mode-config", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/ipsec/mode-config failed", err.Error())
		return
	}
	iPIpsecModeConfigApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIpsecModeConfigResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPIpsecModeConfigModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/ipsec/mode-config", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/ipsec/mode-config failed", err.Error())
		return
	}
	iPIpsecModeConfigApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPIpsecModeConfigResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPIpsecModeConfigModel
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
	if !plan.Address.Equal(state.Address) && !plan.Address.IsUnknown() {
		body["address"] = plan.Address.ValueString()
	}
	if !plan.AddressPool.Equal(state.AddressPool) && !plan.AddressPool.IsUnknown() {
		body["address-pool"] = plan.AddressPool.ValueString()
	}
	if !plan.AddressPrefixLength.Equal(state.AddressPrefixLength) && !plan.AddressPrefixLength.IsUnknown() {
		body["address-prefix-length"] = client.FormatInt64(plan.AddressPrefixLength.ValueInt64())
	}
	if !plan.ConnectionMark.Equal(state.ConnectionMark) && !plan.ConnectionMark.IsUnknown() {
		body["connection-mark"] = plan.ConnectionMark.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Responder.Equal(state.Responder) && !plan.Responder.IsUnknown() {
		body["responder"] = client.FormatBool(plan.Responder.ValueBool())
	}
	if !plan.SplitDNS.Equal(state.SplitDNS) && !plan.SplitDNS.IsUnknown() {
		body["split-dns"] = plan.SplitDNS.ValueString()
	}
	if !plan.SplitInclude.Equal(state.SplitInclude) && !plan.SplitInclude.IsUnknown() {
		body["split-include"] = plan.SplitInclude.ValueString()
	}
	if !plan.SrcAddressList.Equal(state.SrcAddressList) && !plan.SrcAddressList.IsUnknown() {
		body["src-address-list"] = plan.SrcAddressList.ValueString()
	}
	if !plan.StaticDNS.Equal(state.StaticDNS) && !plan.StaticDNS.IsUnknown() {
		body["static-dns"] = plan.StaticDNS.ValueString()
	}
	if !plan.SystemDNS.Equal(state.SystemDNS) && !plan.SystemDNS.IsUnknown() {
		body["system-dns"] = client.FormatBool(plan.SystemDNS.ValueBool())
	}
	if !plan.UseResponderDNS.Equal(state.UseResponderDNS) && !plan.UseResponderDNS.IsUnknown() {
		body["use-responder-dns"] = plan.UseResponderDNS.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/ipsec/mode-config", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/ipsec/mode-config failed", err.Error())
			return
		}
		iPIpsecModeConfigApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIpsecModeConfigResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPIpsecModeConfigModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/ipsec/mode-config", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/ipsec/mode-config failed", err.Error())
	}
}

func (r *IPIpsecModeConfigResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPIpsecModeConfigLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/ipsec/mode-config matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPIpsecModeConfigLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPIpsecModeConfigLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/ipsec/mode-config", id)
}

func iPIpsecModeConfigApply(ctx context.Context, obj client.Object, m *IPIpsecModeConfigModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["address"]; ok {
		_ = v
		if v != "" {
			m.Address = types.StringValue(v)
		} else {
			m.Address = types.StringNull()
		}
	} else {
		m.Address = types.StringNull()
	}
	if v, ok := obj["address-pool"]; ok {
		_ = v
		if v != "" {
			m.AddressPool = types.StringValue(v)
		} else {
			m.AddressPool = types.StringNull()
		}
	} else {
		m.AddressPool = types.StringNull()
	}
	if v, ok := obj["address-prefix-length"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.AddressPrefixLength = types.Int64Value(n)
		} else {
			m.AddressPrefixLength = types.Int64Null()
		}
	} else {
		m.AddressPrefixLength = types.Int64Null()
	}
	if v, ok := obj["connection-mark"]; ok {
		_ = v
		if v != "" {
			m.ConnectionMark = types.StringValue(v)
		} else {
			m.ConnectionMark = types.StringNull()
		}
	} else {
		m.ConnectionMark = types.StringNull()
	}
	if v, ok := obj["default"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else {
			m.Default = types.BoolNull()
		}
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
	if v, ok := obj["nonresp"]; ok {
		_ = v
		if v != "" {
			m.Nonresp = types.StringValue(v)
		} else {
			m.Nonresp = types.StringNull()
		}
	} else {
		m.Nonresp = types.StringNull()
	}
	if v, ok := obj["pool"]; ok {
		_ = v
		if v != "" {
			m.Pool = types.StringValue(v)
		} else {
			m.Pool = types.StringNull()
		}
	} else {
		m.Pool = types.StringNull()
	}
	if v, ok := obj["resp"]; ok {
		_ = v
		if v != "" {
			m.Resp = types.StringValue(v)
		} else {
			m.Resp = types.StringNull()
		}
	} else {
		m.Resp = types.StringNull()
	}
	if v, ok := obj["responder"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Responder = types.BoolValue(b)
		} else {
			m.Responder = types.BoolNull()
		}
	}
	if v, ok := obj["sdns"]; ok {
		_ = v
		if v != "" {
			m.Sdns = types.StringValue(v)
		} else {
			m.Sdns = types.StringNull()
		}
	} else {
		m.Sdns = types.StringNull()
	}
	if v, ok := obj["split-dns"]; ok {
		_ = v
		if v != "" {
			m.SplitDNS = types.StringValue(v)
		} else {
			m.SplitDNS = types.StringNull()
		}
	} else {
		m.SplitDNS = types.StringNull()
	}
	if v, ok := obj["split-include"]; ok {
		_ = v
		if v != "" {
			m.SplitInclude = types.StringValue(v)
		} else {
			m.SplitInclude = types.StringNull()
		}
	} else {
		m.SplitInclude = types.StringNull()
	}
	if v, ok := obj["src-address-list"]; ok {
		_ = v
		if v != "" {
			m.SrcAddressList = types.StringValue(v)
		} else {
			m.SrcAddressList = types.StringNull()
		}
	} else {
		m.SrcAddressList = types.StringNull()
	}
	if v, ok := obj["static-dns"]; ok {
		_ = v
		if v != "" {
			m.StaticDNS = types.StringValue(v)
		} else {
			m.StaticDNS = types.StringNull()
		}
	} else {
		m.StaticDNS = types.StringNull()
	}
	if v, ok := obj["system-dns"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SystemDNS = types.BoolValue(b)
		} else {
			m.SystemDNS = types.BoolNull()
		}
	}
	if v, ok := obj["use-responder-dns"]; ok {
		_ = v
		if v != "" {
			m.UseResponderDNS = types.StringValue(v)
		} else {
			m.UseResponderDNS = types.StringNull()
		}
	} else {
		m.UseResponderDNS = types.StringNull()
	}
}
