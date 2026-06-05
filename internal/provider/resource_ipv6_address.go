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
	_ = fmt.Sprintf
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
			"actual_interface": schema.StringAttribute{
				Optional:    true,
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
			"dynglob": schema.StringAttribute{
				Optional:    true,
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
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"link_local": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"no_dad": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"preferred": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"scope": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"slave": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"valid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vrf": schema.StringAttribute{
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
	if !(plan.Dynglob.IsNull() || plan.Dynglob.IsUnknown()) {
		body["dynglob"] = plan.Dynglob.ValueString()
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
	obj, err := c.Add(ctx, "/ipv6/address", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ipv6/address failed", err.Error())
		return
	}
	iPV6AddressApply(ctx, obj, &plan)
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
	if !plan.Address.Equal(state.Address) {
		body["address"] = plan.Address.ValueString()
	}
	if !plan.Advertise.Equal(state.Advertise) {
		body["advertise"] = client.FormatBool(plan.Advertise.ValueBool())
	}
	if !plan.AutoLinkLocal.Equal(state.AutoLinkLocal) {
		body["auto-link-local"] = client.FormatBool(plan.AutoLinkLocal.ValueBool())
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Dynglob.Equal(state.Dynglob) {
		body["dynglob"] = plan.Dynglob.ValueString()
	}
	if !plan.Eui64.Equal(state.Eui64) {
		body["eui-64"] = client.FormatBool(plan.Eui64.ValueBool())
	}
	if !plan.FromPool.Equal(state.FromPool) {
		body["from-pool"] = plan.FromPool.ValueString()
	}
	if !plan.Interface.Equal(state.Interface) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.NoDad.Equal(state.NoDad) {
		body["no-dad"] = client.FormatBool(plan.NoDad.ValueBool())
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
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/ipv6/address", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func iPV6AddressApply(ctx context.Context, obj client.Object, m *IPV6AddressModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["actual-interface"]; ok {
		_ = v
		if v != "" {
			m.ActualInterface = types.StringValue(v)
		} else {
			m.ActualInterface = types.StringNull()
		}
	} else {
		m.ActualInterface = types.StringNull()
	}
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
	if v, ok := obj["advertise"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Advertise = types.BoolValue(b)
		} else {
			m.Advertise = types.BoolNull()
		}
	} else {
		m.Advertise = types.BoolNull()
	}
	if v, ok := obj["auto-link-local"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.AutoLinkLocal = types.BoolValue(b)
		} else {
			m.AutoLinkLocal = types.BoolNull()
		}
	} else {
		m.AutoLinkLocal = types.BoolNull()
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
	if v, ok := obj["deprecated"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Deprecated = types.BoolValue(b)
		} else {
			m.Deprecated = types.BoolNull()
		}
	} else {
		m.Deprecated = types.BoolNull()
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
	if v, ok := obj["dynglob"]; ok {
		_ = v
		if v != "" {
			m.Dynglob = types.StringValue(v)
		} else {
			m.Dynglob = types.StringNull()
		}
	} else {
		m.Dynglob = types.StringNull()
	}
	if v, ok := obj["eui-64"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Eui64 = types.BoolValue(b)
		} else {
			m.Eui64 = types.BoolNull()
		}
	} else {
		m.Eui64 = types.BoolNull()
	}
	if v, ok := obj["from-pool"]; ok {
		_ = v
		if v != "" {
			m.FromPool = types.StringValue(v)
		} else {
			m.FromPool = types.StringNull()
		}
	} else {
		m.FromPool = types.StringNull()
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
	if v, ok := obj["link-local"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.LinkLocal = types.BoolValue(b)
		} else {
			m.LinkLocal = types.BoolNull()
		}
	} else {
		m.LinkLocal = types.BoolNull()
	}
	if v, ok := obj["no-dad"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.NoDad = types.BoolValue(b)
		} else {
			m.NoDad = types.BoolNull()
		}
	} else {
		m.NoDad = types.BoolNull()
	}
	if v, ok := obj["preferred"]; ok {
		_ = v
		if v != "" {
			m.Preferred = types.StringValue(v)
		} else {
			m.Preferred = types.StringNull()
		}
	} else {
		m.Preferred = types.StringNull()
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
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Slave = types.BoolValue(b)
		} else {
			m.Slave = types.BoolNull()
		}
	} else {
		m.Slave = types.BoolNull()
	}
	if v, ok := obj["valid"]; ok {
		_ = v
		if v != "" {
			m.Valid = types.StringValue(v)
		} else {
			m.Valid = types.StringNull()
		}
	} else {
		m.Valid = types.StringNull()
	}
	if v, ok := obj["vrf"]; ok {
		_ = v
		if v != "" {
			m.Vrf = types.StringValue(v)
		} else {
			m.Vrf = types.StringNull()
		}
	} else {
		m.Vrf = types.StringNull()
	}
}
