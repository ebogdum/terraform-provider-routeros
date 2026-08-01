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
	_ resource.Resource                = &IPPoolResource{}
	_ resource.ResourceWithImportState = &IPPoolResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPPoolResource struct {
	reg *client.Registry
}

type IPPoolModel struct {
	ID        types.String `tfsdk:"id"`
	Addresses types.String `tfsdk:"addresses"`
	Available types.String `tfsdk:"available"`
	Comment   types.String `tfsdk:"comment"`
	Name      types.String `tfsdk:"name"`
	NextPool  types.String `tfsdk:"next_pool"`
	Ranges    types.String `tfsdk:"ranges"`
	Total     types.String `tfsdk:"total"`
	Used      types.String `tfsdk:"used"`
	Router    types.String `tfsdk:"router"`
}

func NewIPPoolResource() resource.Resource { return &IPPoolResource{} }

func (r *IPPoolResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_pool"
}

func (r *IPPoolResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPPoolResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/pool`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"addresses": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"available": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Unique identifier of the pool",
			},
			"next_pool": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "When IP address acquisition is performed a pool that has no free addresses, and the next-pool property is set, then IP address will be acquired from next-pool",
			},
			"ranges": schema.StringAttribute{
				Required:    true,
				Description: "IP address list of non-overlapping IP address ranges in the form of: from1-to1,from2-to2,...,fromN-toN. For example, 10.0.0.1-10.0.0.27,10.0.0.32-10.0.0.47",
			},
			"total": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"used": schema.StringAttribute{
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

func (r *IPPoolResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPPoolModel
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
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.NextPool.IsNull() || plan.NextPool.IsUnknown()) {
		body["next-pool"] = plan.NextPool.ValueString()
	}
	if !(plan.Ranges.IsNull() || plan.Ranges.IsUnknown()) {
		body["ranges"] = plan.Ranges.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/pool", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/pool failed", err.Error())
		return
	}
	iPPoolApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPPoolResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPPoolModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/pool", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/pool failed", err.Error())
		return
	}
	iPPoolApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPPoolResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPPoolModel
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
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.NextPool.Equal(state.NextPool) && !plan.NextPool.IsUnknown() {
		body["next-pool"] = plan.NextPool.ValueString()
	}
	if !plan.Ranges.Equal(state.Ranges) && !plan.Ranges.IsUnknown() {
		body["ranges"] = plan.Ranges.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/pool", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/pool failed", err.Error())
			return
		}
		iPPoolApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPPoolResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPPoolModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/pool", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/pool failed", err.Error())
	}
}

func (r *IPPoolResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPPoolLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/pool matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPPoolLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPPoolLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/pool", id)
}

func iPPoolApply(ctx context.Context, obj client.Object, m *IPPoolModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["addresses"]; ok {
		_ = v
		if v != "" {
			m.Addresses = types.StringValue(v)
		} else {
			m.Addresses = types.StringNull()
		}
	} else {
		m.Addresses = types.StringNull()
	}
	if v, ok := obj["available"]; ok {
		_ = v
		if v != "" {
			m.Available = types.StringValue(v)
		} else {
			m.Available = types.StringNull()
		}
	} else {
		m.Available = types.StringNull()
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
	if v, ok := obj["next-pool"]; ok {
		_ = v
		if v != "" {
			m.NextPool = types.StringValue(v)
		} else {
			m.NextPool = types.StringNull()
		}
	} else {
		m.NextPool = types.StringNull()
	}
	if v, ok := obj["ranges"]; ok {
		_ = v
		if v != "" {
			m.Ranges = types.StringValue(v)
		} else {
			m.Ranges = types.StringNull()
		}
	} else {
		m.Ranges = types.StringNull()
	}
	if v, ok := obj["total"]; ok {
		_ = v
		if v != "" {
			m.Total = types.StringValue(v)
		} else {
			m.Total = types.StringNull()
		}
	} else {
		m.Total = types.StringNull()
	}
	if v, ok := obj["used"]; ok {
		_ = v
		if v != "" {
			m.Used = types.StringValue(v)
		} else {
			m.Used = types.StringNull()
		}
	} else {
		m.Used = types.StringNull()
	}
}
