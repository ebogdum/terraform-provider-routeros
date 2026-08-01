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
	_ resource.Resource                = &InterfaceMeshPortResource{}
	_ resource.ResourceWithImportState = &InterfaceMeshPortResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceMeshPortResource struct {
	reg *client.Registry
}

type InterfaceMeshPortModel struct {
	ID            types.String `tfsdk:"id"`
	Comment       types.String `tfsdk:"comment"`
	Disabled      types.Bool   `tfsdk:"disabled"`
	DrAddress     types.String `tfsdk:"dr_address"`
	Dynamic       types.Bool   `tfsdk:"dynamic"`
	HelloInterval types.Int64  `tfsdk:"hello_interval"`
	Inactive      types.Bool   `tfsdk:"inactive"`
	Interface     types.String `tfsdk:"interface"`
	Mesh          types.String `tfsdk:"mesh"`
	PathCost      types.Int64  `tfsdk:"path_cost"`
	PortType      types.String `tfsdk:"port_type"`
	Router        types.String `tfsdk:"router"`
}

func NewInterfaceMeshPortResource() resource.Resource { return &InterfaceMeshPortResource{} }

func (r *InterfaceMeshPortResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_mesh_port"
}

func (r *InterfaceMeshPortResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceMeshPortResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Discovered via WebFig; needs mesh-interface fixture",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dr_address": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"hello_interval": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"inactive": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mesh": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"path_cost": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"port_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"auto", "wds", "wireless", "ethernet"}...)},
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceMeshPortResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceMeshPortModel
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
	if !(plan.HelloInterval.IsNull() || plan.HelloInterval.IsUnknown()) {
		body["hello-interval"] = client.FormatInt64(plan.HelloInterval.ValueInt64())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.Mesh.IsNull() || plan.Mesh.IsUnknown()) {
		body["mesh"] = plan.Mesh.ValueString()
	}
	if !(plan.PathCost.IsNull() || plan.PathCost.IsUnknown()) {
		body["path-cost"] = client.FormatInt64(plan.PathCost.ValueInt64())
	}
	if !(plan.PortType.IsNull() || plan.PortType.IsUnknown()) {
		body["port-type"] = plan.PortType.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/mesh/port", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/mesh/port failed", err.Error())
		return
	}
	interfaceMeshPortApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceMeshPortResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceMeshPortModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/mesh/port", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/mesh/port failed", err.Error())
		return
	}
	interfaceMeshPortApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceMeshPortResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceMeshPortModel
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
	if !plan.HelloInterval.Equal(state.HelloInterval) && !plan.HelloInterval.IsUnknown() {
		body["hello-interval"] = client.FormatInt64(plan.HelloInterval.ValueInt64())
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.Mesh.Equal(state.Mesh) && !plan.Mesh.IsUnknown() {
		body["mesh"] = plan.Mesh.ValueString()
	}
	if !plan.PathCost.Equal(state.PathCost) && !plan.PathCost.IsUnknown() {
		body["path-cost"] = client.FormatInt64(plan.PathCost.ValueInt64())
	}
	if !plan.PortType.Equal(state.PortType) && !plan.PortType.IsUnknown() {
		body["port-type"] = plan.PortType.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/mesh/port", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/mesh/port failed", err.Error())
			return
		}
		interfaceMeshPortApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceMeshPortResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceMeshPortModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/mesh/port", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/mesh/port failed", err.Error())
	}
}

func (r *InterfaceMeshPortResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceMeshPortLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/mesh/port matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceMeshPortLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceMeshPortLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/mesh/port", id)
}

func interfaceMeshPortApply(ctx context.Context, obj client.Object, m *InterfaceMeshPortModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
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
	if v, ok := obj["dr-address"]; ok {
		_ = v
		if v != "" {
			m.DrAddress = types.StringValue(v)
		} else {
			m.DrAddress = types.StringNull()
		}
	} else {
		m.DrAddress = types.StringNull()
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
	if v, ok := obj["hello-interval"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.HelloInterval = types.Int64Value(n)
		} else {
			m.HelloInterval = types.Int64Null()
		}
	} else {
		m.HelloInterval = types.Int64Null()
	}
	if v, ok := obj["inactive"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Inactive = types.BoolValue(b)
		} else {
			m.Inactive = types.BoolNull()
		}
	} else {
		m.Inactive = types.BoolNull()
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
	if v, ok := obj["mesh"]; ok {
		_ = v
		if v != "" {
			m.Mesh = types.StringValue(v)
		} else {
			m.Mesh = types.StringNull()
		}
	} else {
		m.Mesh = types.StringNull()
	}
	if v, ok := obj["path-cost"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PathCost = types.Int64Value(n)
		} else {
			m.PathCost = types.Int64Null()
		}
	} else {
		m.PathCost = types.Int64Null()
	}
	if v, ok := obj["port-type"]; ok {
		_ = v
		if v != "" {
			m.PortType = types.StringValue(v)
		} else {
			m.PortType = types.StringNull()
		}
	} else {
		m.PortType = types.StringNull()
	}
}
