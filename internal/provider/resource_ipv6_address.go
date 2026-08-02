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
	_ resource.Resource                = &IPV6AddressResource{}
	_ resource.ResourceWithImportState = &IPV6AddressResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPV6AddressResource struct {
	reg *client.Registry
}

type IPV6AddressModel struct {
	ID              types.String `tfsdk:"id"`
	FromPoolPolicy  types.String `tfsdk:"from_pool_policy"`
	ActualInterface types.String `tfsdk:"actual_interface"`
	Address         types.String `tfsdk:"address"`
	Advertise       types.Bool   `tfsdk:"advertise"`
	AutoLinkLocal   types.Bool   `tfsdk:"auto_link_local"`
	Comment         types.String `tfsdk:"comment"`
	Deprecated      types.Bool   `tfsdk:"deprecated"`
	Disabled        types.Bool   `tfsdk:"disabled"`
	Dynamic         types.Bool   `tfsdk:"dynamic"`
	Dynglob         types.String `tfsdk:"dynglob"`
	Eui64           types.Bool   `tfsdk:"eui_64"`
	FromPool        types.String `tfsdk:"from_pool"`
	Interface       types.String `tfsdk:"interface"`
	Invalid         types.Bool   `tfsdk:"invalid"`
	LinkLocal       types.Bool   `tfsdk:"link_local"`
	NoDad           types.Bool   `tfsdk:"no_dad"`
	Preferred       types.String `tfsdk:"preferred"`
	Scope           types.Int64  `tfsdk:"scope"`
	Slave           types.Bool   `tfsdk:"slave"`
	Valid           types.String `tfsdk:"valid"`
	Vrf             types.String `tfsdk:"vrf"`
	Router          types.String `tfsdk:"router"`
}

func NewIPV6AddressResource() resource.Resource { return &IPV6AddressResource{} }

func (r *IPV6AddressResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv6_address"
}

func (r *IPV6AddressResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPV6AddressResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ipv6/address`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"from_pool_policy": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `from-pool-policy`.",
			},
			"actual_interface": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"address": schema.StringAttribute{
				Required:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsCIDR()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeCIDR()},
			},
			"advertise": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"auto_link_local": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"deprecated": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"dynglob": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"eui_64": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"from_pool": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"interface": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"link_local": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"no_dad": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"preferred": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"scope": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"slave": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"valid": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"vrf": schema.StringAttribute{
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

func (r *IPV6AddressResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPV6AddressModel
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
	if !(plan.Advertise.IsNull() || plan.Advertise.IsUnknown()) {
		body["advertise"] = client.FormatBool(plan.Advertise.ValueBool())
	}
	if !(plan.AutoLinkLocal.IsNull() || plan.AutoLinkLocal.IsUnknown()) {
		body["auto-link-local"] = client.FormatBool(plan.AutoLinkLocal.ValueBool())
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Eui64.IsNull() || plan.Eui64.IsUnknown()) {
		body["eui-64"] = client.FormatBool(plan.Eui64.ValueBool())
	}
	if !(plan.FromPool.IsNull() || plan.FromPool.IsUnknown()) {
		body["from-pool"] = plan.FromPool.ValueString()
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.NoDad.IsNull() || plan.NoDad.IsUnknown()) {
		body["no-dad"] = client.FormatBool(plan.NoDad.ValueBool())
	}
	if !(plan.FromPoolPolicy.IsNull() || plan.FromPoolPolicy.IsUnknown()) {
		body["from-pool-policy"] = plan.FromPoolPolicy.ValueString()
	}
	obj, err := c.Add(ctx, "/ipv6/address", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ipv6/address failed", err.Error())
		return
	}
	iPV6AddressApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6AddressResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPV6AddressModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ipv6/address", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ipv6/address failed", err.Error())
		return
	}
	iPV6AddressApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPV6AddressResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPV6AddressModel
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
	if !plan.Advertise.Equal(state.Advertise) && !plan.Advertise.IsUnknown() {
		body["advertise"] = client.FormatBool(plan.Advertise.ValueBool())
	}
	if !plan.AutoLinkLocal.Equal(state.AutoLinkLocal) && !plan.AutoLinkLocal.IsUnknown() {
		body["auto-link-local"] = client.FormatBool(plan.AutoLinkLocal.ValueBool())
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Eui64.Equal(state.Eui64) && !plan.Eui64.IsUnknown() {
		body["eui-64"] = client.FormatBool(plan.Eui64.ValueBool())
	}
	if !plan.FromPool.Equal(state.FromPool) && !plan.FromPool.IsUnknown() {
		body["from-pool"] = plan.FromPool.ValueString()
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.NoDad.Equal(state.NoDad) && !plan.NoDad.IsUnknown() {
		body["no-dad"] = client.FormatBool(plan.NoDad.ValueBool())
	}
	if !plan.FromPoolPolicy.Equal(state.FromPoolPolicy) && !plan.FromPoolPolicy.IsUnknown() {
		body["from-pool-policy"] = plan.FromPoolPolicy.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ipv6/address", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ipv6/address failed", err.Error())
			return
		}
		iPV6AddressApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6AddressResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPV6AddressModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ipv6/address", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ipv6/address failed", err.Error())
	}
}

func (r *IPV6AddressResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPV6AddressLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ipv6/address matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPV6AddressLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPV6AddressLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ipv6/address", id)
}

func iPV6AddressApply(ctx context.Context, obj client.Object, m *IPV6AddressModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["from-pool-policy"]; ok && v != "" {
		m.FromPoolPolicy = types.StringValue(v)
	} else {
		m.FromPoolPolicy = types.StringNull()
	}
	if v, ok := obj["actual-interface"]; ok {
		if v != "" {
			m.ActualInterface = types.StringValue(v)
		} else {
			m.ActualInterface = types.StringNull()
		}
	}
	if v, ok := obj["address"]; ok {
		if v != "" {
			m.Address = types.StringValue(v)
		} else {
			m.Address = types.StringNull()
		}
	}
	if v, ok := obj["advertise"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Advertise = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Advertise = types.BoolValue(true)
		} else {
			m.Advertise = types.BoolNull()
		}
	}
	if v, ok := obj["auto-link-local"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.AutoLinkLocal = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.AutoLinkLocal = types.BoolValue(true)
		} else {
			m.AutoLinkLocal = types.BoolNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["deprecated"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Deprecated = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Deprecated = types.BoolValue(true)
		} else {
			m.Deprecated = types.BoolNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Disabled = types.BoolValue(true)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["dynamic"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Dynamic = types.BoolValue(true)
		} else {
			m.Dynamic = types.BoolNull()
		}
	}
	if v, ok := obj["dynglob"]; ok {
		if v != "" {
			m.Dynglob = types.StringValue(v)
		} else {
			m.Dynglob = types.StringNull()
		}
	}
	if v, ok := obj["eui-64"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Eui64 = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Eui64 = types.BoolValue(true)
		} else {
			m.Eui64 = types.BoolNull()
		}
	}
	if v, ok := obj["from-pool"]; ok {
		if v != "" {
			m.FromPool = types.StringValue(v)
		} else {
			m.FromPool = types.StringNull()
		}
	}
	if v, ok := obj["interface"]; ok {
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	}
	if v, ok := obj["invalid"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Invalid = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Invalid = types.BoolValue(true)
		} else {
			m.Invalid = types.BoolNull()
		}
	}
	if v, ok := obj["link-local"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.LinkLocal = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.LinkLocal = types.BoolValue(true)
		} else {
			m.LinkLocal = types.BoolNull()
		}
	}
	if v, ok := obj["no-dad"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.NoDad = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.NoDad = types.BoolValue(true)
		} else {
			m.NoDad = types.BoolNull()
		}
	}
	if v, ok := obj["preferred"]; ok {
		if v != "" {
			m.Preferred = types.StringValue(v)
		} else {
			m.Preferred = types.StringNull()
		}
	}
	if v, ok := obj["scope"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Scope = types.Int64Value(n)
		} else {
			m.Scope = types.Int64Null()
		}
	} else {
		m.Scope = types.Int64Null()
	}
	if v, ok := obj["slave"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Slave = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Slave = types.BoolValue(true)
		} else {
			m.Slave = types.BoolNull()
		}
	}
	if v, ok := obj["valid"]; ok {
		if v != "" {
			m.Valid = types.StringValue(v)
		} else {
			m.Valid = types.StringNull()
		}
	}
	if v, ok := obj["vrf"]; ok {
		if v != "" {
			m.Vrf = types.StringValue(v)
		} else {
			m.Vrf = types.StringNull()
		}
	}
}
