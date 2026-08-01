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
	_ resource.Resource                = &Interface6to4Resource{}
	_ resource.ResourceWithImportState = &Interface6to4Resource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type Interface6to4Resource struct {
	reg *client.Registry
}

type Interface6to4Model struct {
	ID            types.String `tfsdk:"id"`
	Keepalive     types.String `tfsdk:"keepalive"`
	IpsecSecret   types.String `tfsdk:"ipsec_secret"`
	ClampTCPMss   types.Bool   `tfsdk:"clamp_tcp_mss"`
	Comment       types.String `tfsdk:"comment"`
	Disabled      types.Bool   `tfsdk:"disabled"`
	DontFragment  types.Bool   `tfsdk:"dont_fragment"`
	Dscp          types.String `tfsdk:"dscp"`
	LocalAddress  types.String `tfsdk:"local_address"`
	MTU           types.String `tfsdk:"mtu"`
	Name          types.String `tfsdk:"name"`
	RemoteAddress types.String `tfsdk:"remote_address"`
	Router        types.String `tfsdk:"router"`
}

func NewInterface6to4Resource() resource.Resource { return &Interface6to4Resource{} }

func (r *Interface6to4Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_6to4"
}

func (r *Interface6to4Resource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *Interface6to4Resource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "6to4 tunnel deletion races on CHR (DELETE returns errors even after success). Skipped from automated acc tests.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"keepalive": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `keepalive`.",
			},
			"ipsec_secret": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "RouterOS `ipsec-secret`.",
			},
			"clamp_tcp_mss": schema.BoolAttribute{
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
			"dont_fragment": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dscp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"local_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_address": schema.StringAttribute{
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

func (r *Interface6to4Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan Interface6to4Model
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.ClampTCPMss.IsNull() || plan.ClampTCPMss.IsUnknown()) {
		body["clamp-tcp-mss"] = client.FormatBool(plan.ClampTCPMss.ValueBool())
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.DontFragment.IsNull() || plan.DontFragment.IsUnknown()) {
		body["dont-fragment"] = client.FormatBool(plan.DontFragment.ValueBool())
	}
	if !(plan.Dscp.IsNull() || plan.Dscp.IsUnknown()) {
		body["dscp"] = plan.Dscp.ValueString()
	}
	if !(plan.LocalAddress.IsNull() || plan.LocalAddress.IsUnknown()) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.RemoteAddress.IsNull() || plan.RemoteAddress.IsUnknown()) {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !(plan.IpsecSecret.IsNull() || plan.IpsecSecret.IsUnknown()) {
		body["ipsec-secret"] = plan.IpsecSecret.ValueString()
	}
	if !(plan.Keepalive.IsNull() || plan.Keepalive.IsUnknown()) {
		body["keepalive"] = plan.Keepalive.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/6to4", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/6to4 failed", err.Error())
		return
	}
	interface6to4Apply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *Interface6to4Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state Interface6to4Model
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/6to4", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/6to4 failed", err.Error())
		return
	}
	interface6to4Apply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *Interface6to4Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state Interface6to4Model
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
	if !plan.ClampTCPMss.Equal(state.ClampTCPMss) && !plan.ClampTCPMss.IsUnknown() {
		body["clamp-tcp-mss"] = client.FormatBool(plan.ClampTCPMss.ValueBool())
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DontFragment.Equal(state.DontFragment) && !plan.DontFragment.IsUnknown() {
		body["dont-fragment"] = client.FormatBool(plan.DontFragment.ValueBool())
	}
	if !plan.Dscp.Equal(state.Dscp) && !plan.Dscp.IsUnknown() {
		body["dscp"] = plan.Dscp.ValueString()
	}
	if !plan.LocalAddress.Equal(state.LocalAddress) && !plan.LocalAddress.IsUnknown() {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !plan.MTU.Equal(state.MTU) && !plan.MTU.IsUnknown() {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.RemoteAddress.Equal(state.RemoteAddress) && !plan.RemoteAddress.IsUnknown() {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if !plan.IpsecSecret.Equal(state.IpsecSecret) && !plan.IpsecSecret.IsUnknown() {
		body["ipsec-secret"] = plan.IpsecSecret.ValueString()
	}
	if !plan.Keepalive.Equal(state.Keepalive) && !plan.Keepalive.IsUnknown() {
		body["keepalive"] = plan.Keepalive.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/6to4", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/6to4 failed", err.Error())
			return
		}
		interface6to4Apply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *Interface6to4Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state Interface6to4Model
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/6to4", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/6to4 failed", err.Error())
	}
}

func (r *Interface6to4Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interface6to4LookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/6to4 matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interface6to4LookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interface6to4LookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/6to4", id)
}

func interface6to4Apply(ctx context.Context, obj client.Object, m *Interface6to4Model) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["keepalive"]; ok && v != "" {
		m.Keepalive = types.StringValue(v)
	} else {
		m.Keepalive = types.StringNull()
	}
	if v, ok := obj["ipsec-secret"]; ok && v != "" {
		m.IpsecSecret = types.StringValue(v)
	} else {
		m.IpsecSecret = types.StringNull()
	}
	if v, ok := obj["clamp-tcp-mss"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.ClampTCPMss = types.BoolValue(b)
		} else {
			m.ClampTCPMss = types.BoolNull()
		}
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
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["dont-fragment"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.DontFragment = types.BoolValue(b)
		} else {
			m.DontFragment = types.BoolNull()
		}
	}
	if v, ok := obj["dscp"]; ok {
		_ = v
		if v != "" {
			m.Dscp = types.StringValue(v)
		} else {
			m.Dscp = types.StringNull()
		}
	} else {
		m.Dscp = types.StringNull()
	}
	if v, ok := obj["local-address"]; ok {
		_ = v
		if v != "" {
			m.LocalAddress = types.StringValue(v)
		} else {
			m.LocalAddress = types.StringNull()
		}
	} else {
		m.LocalAddress = types.StringNull()
	}
	if v, ok := obj["mtu"]; ok {
		_ = v
		if v != "" {
			m.MTU = types.StringValue(v)
		} else {
			m.MTU = types.StringNull()
		}
	} else {
		m.MTU = types.StringNull()
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
	if v, ok := obj["remote-address"]; ok {
		_ = v
		if v != "" {
			m.RemoteAddress = types.StringValue(v)
		} else {
			m.RemoteAddress = types.StringNull()
		}
	} else {
		m.RemoteAddress = types.StringNull()
	}
}
