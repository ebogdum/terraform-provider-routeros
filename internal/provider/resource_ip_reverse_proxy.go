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
	_ resource.Resource                = &IPReverseProxyResource{}
	_ resource.ResourceWithImportState = &IPReverseProxyResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPReverseProxyResource struct {
	reg *client.Registry
}

type IPReverseProxyModel struct {
	ID          types.String `tfsdk:"id"`
	Vrf         types.String `tfsdk:"vrf"`
	Sni         types.String `tfsdk:"sni"`
	Port        types.String `tfsdk:"port"`
	IpAddress   types.String `tfsdk:"ip_address"`
	Certificate types.String `tfsdk:"certificate"`
	Comment     types.String `tfsdk:"comment"`
	Disabled    types.Bool   `tfsdk:"disabled"`
	Router      types.String `tfsdk:"router"`
}

func NewIPReverseProxyResource() resource.Resource { return &IPReverseProxyResource{} }

func (r *IPReverseProxyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_reverse_proxy"
}

func (r *IPReverseProxyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPReverseProxyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Reverse proxy listener — properties differ across ROS versions; safe defaults rejected on 7.x. Skipped.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"vrf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `vrf`.",
			},
			"sni": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `sni`.",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `port`.",
			},
			"ip_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ip-address`.",
			},
			"certificate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `certificate`.",
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
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPReverseProxyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPReverseProxyModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Certificate.IsNull() || plan.Certificate.IsUnknown()) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !(plan.IpAddress.IsNull() || plan.IpAddress.IsUnknown()) {
		body["ip-address"] = plan.IpAddress.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.Sni.IsNull() || plan.Sni.IsUnknown()) {
		body["sni"] = plan.Sni.ValueString()
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/reverse-proxy", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/reverse-proxy failed", err.Error())
		return
	}
	iPReverseProxyApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPReverseProxyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPReverseProxyModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/reverse-proxy", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/reverse-proxy failed", err.Error())
		return
	}
	iPReverseProxyApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPReverseProxyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPReverseProxyModel
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
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Certificate.Equal(state.Certificate) && !plan.Certificate.IsUnknown() {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !plan.IpAddress.Equal(state.IpAddress) && !plan.IpAddress.IsUnknown() {
		body["ip-address"] = plan.IpAddress.ValueString()
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.Sni.Equal(state.Sni) && !plan.Sni.IsUnknown() {
		body["sni"] = plan.Sni.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) && !plan.Vrf.IsUnknown() {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/reverse-proxy", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/reverse-proxy failed", err.Error())
			return
		}
		iPReverseProxyApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPReverseProxyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPReverseProxyModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/reverse-proxy", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/reverse-proxy failed", err.Error())
	}
}

func (r *IPReverseProxyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPReverseProxyLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/reverse-proxy matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPReverseProxyLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPReverseProxyLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/reverse-proxy", id)
}

func iPReverseProxyApply(ctx context.Context, obj client.Object, m *IPReverseProxyModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["vrf"]; ok && v != "" {
		m.Vrf = types.StringValue(v)
	} else {
		m.Vrf = types.StringNull()
	}
	if v, ok := obj["sni"]; ok && v != "" {
		m.Sni = types.StringValue(v)
	} else {
		m.Sni = types.StringNull()
	}
	if v, ok := obj["port"]; ok && v != "" {
		m.Port = types.StringValue(v)
	} else {
		m.Port = types.StringNull()
	}
	if v, ok := obj["ip-address"]; ok && v != "" {
		m.IpAddress = types.StringValue(v)
	} else {
		m.IpAddress = types.StringNull()
	}
	if v, ok := obj["certificate"]; ok && v != "" {
		m.Certificate = types.StringValue(v)
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
}
