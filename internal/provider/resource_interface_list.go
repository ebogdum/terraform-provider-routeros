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
	_ resource.Resource                = &InterfaceListResource{}
	_ resource.ResourceWithImportState = &InterfaceListResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceListResource struct {
	reg *client.Registry
}

type InterfaceListModel struct {
	ID      types.String `tfsdk:"id"`
	Builtin types.Bool   `tfsdk:"builtin"`
	Comment types.String `tfsdk:"comment"`
	Dynamic types.Bool   `tfsdk:"dynamic"`
	Exclude types.String `tfsdk:"exclude"`
	Include types.String `tfsdk:"include"`
	Name    types.String `tfsdk:"name"`
	Router  types.String `tfsdk:"router"`
}

func NewInterfaceListResource() resource.Resource { return &InterfaceListResource{} }

func (r *InterfaceListResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_list"
}

func (r *InterfaceListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceListResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/list`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"builtin": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"exclude": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Defines interface list which members are excluded from the list. It is possible to add multiple lists separated by commas",
			},
			"include": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Defines interface list which members are included in the list. It is possible to add multiple lists separated by commas",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "Name of the interface list",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *InterfaceListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceListModel
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
	if !(plan.Exclude.IsNull() || plan.Exclude.IsUnknown()) {
		body["exclude"] = plan.Exclude.ValueString()
	}
	if !(plan.Include.IsNull() || plan.Include.IsUnknown()) {
		body["include"] = plan.Include.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/list", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/list failed", err.Error())
		return
	}
	interfaceListApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceListModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/list", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/list failed", err.Error())
		return
	}
	interfaceListApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceListModel
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
	if !plan.Exclude.Equal(state.Exclude) && !plan.Exclude.IsUnknown() {
		body["exclude"] = plan.Exclude.ValueString()
	}
	if !plan.Include.Equal(state.Include) && !plan.Include.IsUnknown() {
		body["include"] = plan.Include.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/list", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/list failed", err.Error())
			return
		}
		interfaceListApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceListModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/list", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/list failed", err.Error())
	}
}

func (r *InterfaceListResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceListLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/list matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceListLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceListLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/list", id)
}

func interfaceListApply(ctx context.Context, obj client.Object, m *InterfaceListModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["builtin"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Builtin = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Builtin = types.BoolValue(true)
		} else {
			m.Builtin = types.BoolNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
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
	if v, ok := obj["exclude"]; ok {
		if v != "" {
			m.Exclude = types.StringValue(v)
		} else {
			m.Exclude = types.StringNull()
		}
	}
	if v, ok := obj["include"]; ok {
		if v != "" {
			m.Include = types.StringValue(v)
		} else {
			m.Include = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
}
