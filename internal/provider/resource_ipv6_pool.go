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
	_ resource.Resource                = &IPV6PoolResource{}
	_ resource.ResourceWithImportState = &IPV6PoolResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPV6PoolResource struct {
	reg *client.Registry
}

type IPV6PoolModel struct {
	ID                types.String `tfsdk:"id"`
	ActualPrefix      types.String `tfsdk:"actual_prefix"`
	Comment           types.String `tfsdk:"comment"`
	Dynamic           types.Bool   `tfsdk:"dynamic"`
	FromPool          types.String `tfsdk:"from_pool"`
	Invalid           types.Bool   `tfsdk:"invalid"`
	Name              types.String `tfsdk:"name"`
	PreferredLifetime types.String `tfsdk:"preferred_lifetime"`
	Prefix            types.String `tfsdk:"prefix"`
	PrefixLength      types.Int64  `tfsdk:"prefix_length"`
	ValidLifetime     types.String `tfsdk:"valid_lifetime"`
	Router            types.String `tfsdk:"router"`
}

func NewIPV6PoolResource() resource.Resource { return &IPV6PoolResource{} }

func (r *IPV6PoolResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv6_pool"
}

func (r *IPV6PoolResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPV6PoolResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ipv6/pool`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"actual_prefix": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"from_pool": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"preferred_lifetime": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"prefix": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"prefix_length": schema.Int64Attribute{
				Required:    true,
				Description: "",
			},
			"valid_lifetime": schema.StringAttribute{
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

func (r *IPV6PoolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPV6PoolModel
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
	if !(plan.FromPool.IsNull() || plan.FromPool.IsUnknown()) {
		body["from-pool"] = plan.FromPool.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Prefix.IsNull() || plan.Prefix.IsUnknown()) {
		body["prefix"] = plan.Prefix.ValueString()
	}
	if !(plan.PrefixLength.IsNull() || plan.PrefixLength.IsUnknown()) {
		body["prefix-length"] = client.FormatInt64(plan.PrefixLength.ValueInt64())
	}
	obj, err := c.Add(ctx, "/ipv6/pool", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ipv6/pool failed", err.Error())
		return
	}
	iPV6PoolApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6PoolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPV6PoolModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ipv6/pool", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ipv6/pool failed", err.Error())
		return
	}
	iPV6PoolApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPV6PoolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPV6PoolModel
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
	if !plan.FromPool.Equal(state.FromPool) && !plan.FromPool.IsUnknown() {
		body["from-pool"] = plan.FromPool.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Prefix.Equal(state.Prefix) && !plan.Prefix.IsUnknown() {
		body["prefix"] = plan.Prefix.ValueString()
	}
	if !plan.PrefixLength.Equal(state.PrefixLength) && !plan.PrefixLength.IsUnknown() {
		body["prefix-length"] = client.FormatInt64(plan.PrefixLength.ValueInt64())
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ipv6/pool", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ipv6/pool failed", err.Error())
			return
		}
		iPV6PoolApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6PoolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPV6PoolModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ipv6/pool", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ipv6/pool failed", err.Error())
	}
}

func (r *IPV6PoolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPV6PoolLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ipv6/pool matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPV6PoolLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPV6PoolLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ipv6/pool", id)
}

func iPV6PoolApply(ctx context.Context, obj client.Object, m *IPV6PoolModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["actual-prefix"]; ok {
		_ = v
		if v != "" {
			m.ActualPrefix = types.StringValue(v)
		} else {
			m.ActualPrefix = types.StringNull()
		}
	} else {
		m.ActualPrefix = types.StringNull()
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
	if v, ok := obj["prefix-length"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PrefixLength = types.Int64Value(n)
		} else {
			m.PrefixLength = types.Int64Null()
		}
	} else {
		m.PrefixLength = types.Int64Null()
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
