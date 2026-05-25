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
	_ resource.Resource                = &IPV6NdPrefixResource{}
	_ resource.ResourceWithImportState = &IPV6NdPrefixResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPV6NdPrefixResource struct {
	reg *client.Registry
}

type IPV6NdPrefixModel struct {
	ID                types.String `tfsdk:"id"`
	X6to4Interface    types.String `tfsdk:"x6to4_interface"`
	Autonomous        types.Bool   `tfsdk:"autonomous"`
	Dhcpv6PdPreferred types.Bool   `tfsdk:"dhcpv6_pd_preferred"`
	Disabled          types.Bool   `tfsdk:"disabled"`
	Dynamic           types.Bool   `tfsdk:"dynamic"`
	Interface         types.String `tfsdk:"interface"`
	Invalid           types.Bool   `tfsdk:"invalid"`
	No6to4            types.String `tfsdk:"no6to4"`
	OnLink            types.Bool   `tfsdk:"on_link"`
	PreferredLifetime types.String `tfsdk:"preferred_lifetime"`
	Prefix            types.String `tfsdk:"prefix"`
	ValidLifetime     types.String `tfsdk:"valid_lifetime"`
	Router            types.String `tfsdk:"router"`
}

func NewIPV6NdPrefixResource() resource.Resource { return &IPV6NdPrefixResource{} }

func (r *IPV6NdPrefixResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv6_nd_prefix"
}

func (r *IPV6NdPrefixResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *IPV6NdPrefixResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ipv6/nd/prefix`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"x6to4_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"autonomous": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dhcpv6_pd_preferred": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dynamic": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"no6to4": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"on_link": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"preferred_lifetime": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"prefix": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"valid_lifetime": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPV6NdPrefixResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPV6NdPrefixModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.X6to4Interface.IsNull() || plan.X6to4Interface.IsUnknown()) {
		body["6to4-interface"] = plan.X6to4Interface.ValueString()
	}
	if !(plan.Autonomous.IsNull() || plan.Autonomous.IsUnknown()) {
		body["autonomous"] = client.FormatBool(plan.Autonomous.ValueBool())
	}
	if !(plan.Dhcpv6PdPreferred.IsNull() || plan.Dhcpv6PdPreferred.IsUnknown()) {
		body["dhcpv6-pd-preferred"] = client.FormatBool(plan.Dhcpv6PdPreferred.ValueBool())
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.No6to4.IsNull() || plan.No6to4.IsUnknown()) {
		body["no6to4"] = plan.No6to4.ValueString()
	}
	if !(plan.OnLink.IsNull() || plan.OnLink.IsUnknown()) {
		body["on-link"] = client.FormatBool(plan.OnLink.ValueBool())
	}
	if !(plan.PreferredLifetime.IsNull() || plan.PreferredLifetime.IsUnknown()) {
		body["preferred-lifetime"] = plan.PreferredLifetime.ValueString()
	}
	if !(plan.Prefix.IsNull() || plan.Prefix.IsUnknown()) {
		body["prefix"] = plan.Prefix.ValueString()
	}
	if !(plan.ValidLifetime.IsNull() || plan.ValidLifetime.IsUnknown()) {
		body["valid-lifetime"] = plan.ValidLifetime.ValueString()
	}
	obj, err := c.Add(ctx, "/ipv6/nd/prefix", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ipv6/nd/prefix failed", err.Error())
		return
	}
	iPV6NdPrefixApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6NdPrefixResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPV6NdPrefixModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ipv6/nd/prefix", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ipv6/nd/prefix failed", err.Error())
		return
	}
	iPV6NdPrefixApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPV6NdPrefixResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPV6NdPrefixModel
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
	if !plan.X6to4Interface.Equal(state.X6to4Interface) {
		body["6to4-interface"] = plan.X6to4Interface.ValueString()
	}
	if !plan.Autonomous.Equal(state.Autonomous) {
		body["autonomous"] = client.FormatBool(plan.Autonomous.ValueBool())
	}
	if !plan.Dhcpv6PdPreferred.Equal(state.Dhcpv6PdPreferred) {
		body["dhcpv6-pd-preferred"] = client.FormatBool(plan.Dhcpv6PdPreferred.ValueBool())
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.No6to4.Equal(state.No6to4) {
		body["no6to4"] = plan.No6to4.ValueString()
	}
	if !plan.OnLink.Equal(state.OnLink) {
		body["on-link"] = client.FormatBool(plan.OnLink.ValueBool())
	}
	if !plan.PreferredLifetime.Equal(state.PreferredLifetime) {
		body["preferred-lifetime"] = plan.PreferredLifetime.ValueString()
	}
	if !plan.Prefix.Equal(state.Prefix) {
		body["prefix"] = plan.Prefix.ValueString()
	}
	if !plan.ValidLifetime.Equal(state.ValidLifetime) {
		body["valid-lifetime"] = plan.ValidLifetime.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ipv6/nd/prefix", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ipv6/nd/prefix failed", err.Error())
			return
		}
		iPV6NdPrefixApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6NdPrefixResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPV6NdPrefixModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ipv6/nd/prefix", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ipv6/nd/prefix failed", err.Error())
	}
}

func (r *IPV6NdPrefixResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPV6NdPrefixLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ipv6/nd/prefix matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPV6NdPrefixLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPV6NdPrefixLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ipv6/nd/prefix", id)
}

func iPV6NdPrefixApply(ctx context.Context, obj client.Object, m *IPV6NdPrefixModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["6to4-interface"]; ok {
		_ = v
		if v != "" {
			m.X6to4Interface = types.StringValue(v)
		} else {
			m.X6to4Interface = types.StringNull()
		}
	} else {
		m.X6to4Interface = types.StringNull()
	}
	if v, ok := obj["autonomous"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Autonomous = types.BoolValue(b)
		} else {
			m.Autonomous = types.BoolNull()
		}
	} else {
		m.Autonomous = types.BoolNull()
	}
	if v, ok := obj["dhcpv6-pd-preferred"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Dhcpv6PdPreferred = types.BoolValue(b)
		} else {
			m.Dhcpv6PdPreferred = types.BoolNull()
		}
	} else {
		m.Dhcpv6PdPreferred = types.BoolNull()
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
	if v, ok := obj["interface"]; ok {
		_ = v
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	} else {
		m.Interface = types.StringNull()
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
	if v, ok := obj["no6to4"]; ok {
		_ = v
		if v != "" {
			m.No6to4 = types.StringValue(v)
		} else {
			m.No6to4 = types.StringNull()
		}
	} else {
		m.No6to4 = types.StringNull()
	}
	if v, ok := obj["on-link"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.OnLink = types.BoolValue(b)
		} else {
			m.OnLink = types.BoolNull()
		}
	} else {
		m.OnLink = types.BoolNull()
	}
	if v, ok := obj["preferred-lifetime"]; ok {
		_ = v
		if v != "" {
			m.PreferredLifetime = types.StringValue(v)
		} else {
			m.PreferredLifetime = types.StringNull()
		}
	} else {
		m.PreferredLifetime = types.StringNull()
	}
	if v, ok := obj["prefix"]; ok {
		_ = v
		if v != "" {
			m.Prefix = types.StringValue(v)
		} else {
			m.Prefix = types.StringNull()
		}
	} else {
		m.Prefix = types.StringNull()
	}
	if v, ok := obj["valid-lifetime"]; ok {
		_ = v
		if v != "" {
			m.ValidLifetime = types.StringValue(v)
		} else {
			m.ValidLifetime = types.StringNull()
		}
	} else {
		m.ValidLifetime = types.StringNull()
	}
}
