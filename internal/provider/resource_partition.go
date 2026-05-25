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
	_ resource.Resource                = &PartitionResource{}
	_ resource.ResourceWithImportState = &PartitionResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type PartitionResource struct {
	reg *client.Registry
}

type PartitionModel struct {
	ID         types.String `tfsdk:"id"`
	Activate   types.String `tfsdk:"activate"`
	Active     types.Bool   `tfsdk:"active"`
	Comment    types.String `tfsdk:"comment"`
	FallbackTo types.String `tfsdk:"fallback_to"`
	Name       types.String `tfsdk:"name"`
	Running    types.Bool   `tfsdk:"running"`
	Size       types.Int64  `tfsdk:"size"`
	Version    types.String `tfsdk:"version"`
	Router     types.String `tfsdk:"router"`
}

func NewPartitionResource() resource.Resource { return &PartitionResource{} }

func (r *PartitionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_partition"
}

func (r *PartitionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *PartitionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/partition`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"activate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"active": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"fallback_to": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"running": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"size": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"version": schema.StringAttribute{
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

func (r *PartitionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PartitionModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Activate.IsNull() || plan.Activate.IsUnknown()) {
		body["activate"] = plan.Activate.ValueString()
	}
	if !(plan.Active.IsNull() || plan.Active.IsUnknown()) {
		body["active"] = client.FormatBool(plan.Active.ValueBool())
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.FallbackTo.IsNull() || plan.FallbackTo.IsUnknown()) {
		body["fallback-to"] = plan.FallbackTo.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Running.IsNull() || plan.Running.IsUnknown()) {
		body["running"] = client.FormatBool(plan.Running.ValueBool())
	}
	obj, err := c.Add(ctx, "/partition", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /partition failed", err.Error())
		return
	}
	partitionApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PartitionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PartitionModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/partition", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /partition failed", err.Error())
		return
	}
	partitionApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *PartitionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PartitionModel
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
	if !plan.Activate.Equal(state.Activate) {
		body["activate"] = plan.Activate.ValueString()
	}
	if !plan.Active.Equal(state.Active) {
		body["active"] = client.FormatBool(plan.Active.ValueBool())
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.FallbackTo.Equal(state.FallbackTo) {
		body["fallback-to"] = plan.FallbackTo.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Running.Equal(state.Running) {
		body["running"] = client.FormatBool(plan.Running.ValueBool())
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/partition", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /partition failed", err.Error())
			return
		}
		partitionApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PartitionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PartitionModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/partition", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /partition failed", err.Error())
	}
}

func (r *PartitionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := partitionLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /partition matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// partitionLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func partitionLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/partition", id)
}

func partitionApply(ctx context.Context, obj client.Object, m *PartitionModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["activate"]; ok {
		_ = v
		if v != "" {
			m.Activate = types.StringValue(v)
		} else {
			m.Activate = types.StringNull()
		}
	} else {
		m.Activate = types.StringNull()
	}
	if v, ok := obj["active"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Active = types.BoolValue(b)
		} else {
			m.Active = types.BoolNull()
		}
	} else {
		m.Active = types.BoolNull()
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
	if v, ok := obj["fallback-to"]; ok {
		_ = v
		if v != "" {
			m.FallbackTo = types.StringValue(v)
		} else {
			m.FallbackTo = types.StringNull()
		}
	} else {
		m.FallbackTo = types.StringNull()
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
	if v, ok := obj["running"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Running = types.BoolValue(b)
		} else {
			m.Running = types.BoolNull()
		}
	} else {
		m.Running = types.BoolNull()
	}
	if v, ok := obj["size"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Size = types.Int64Value(n)
		} else {
			m.Size = types.Int64Null()
		}
	} else {
		m.Size = types.Int64Null()
	}
	if v, ok := obj["version"]; ok {
		_ = v
		if v != "" {
			m.Version = types.StringValue(v)
		} else {
			m.Version = types.StringNull()
		}
	} else {
		m.Version = types.StringNull()
	}
}
