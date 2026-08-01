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
	_ resource.Resource                = &UserManagerAttributeResource{}
	_ resource.ResourceWithImportState = &UserManagerAttributeResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type UserManagerAttributeResource struct {
	reg *client.Registry
}

type UserManagerAttributeModel struct {
	ID           types.String `tfsdk:"id"`
	Default      types.String `tfsdk:"default"`
	DefaultName  types.String `tfsdk:"default_name"`
	Name         types.String `tfsdk:"name"`
	PacketTypes  types.String `tfsdk:"packet_types"`
	StandardName types.String `tfsdk:"standard_name"`
	TypeID       types.String `tfsdk:"type_id"`
	ValueType    types.String `tfsdk:"value_type"`
	VendorID     types.String `tfsdk:"vendor_id"`
	Router       types.String `tfsdk:"router"`
}

func NewUserManagerAttributeResource() resource.Resource { return &UserManagerAttributeResource{} }

func (r *UserManagerAttributeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_user_manager_attribute"
}

func (r *UserManagerAttributeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *UserManagerAttributeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Requires user-manager package",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"default": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"default_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the attribute.",
			},
			"packet_types": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "access-accept - use this attribute in RADIUS Access-Accept messages access-challenge - use this attribute in RADIUS Access-Challenge messages",
			},
			"standard_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"type_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Attribute identification number from the specific vendor's attribute database.",
			},
			"value_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "hex ip-address - IPv4 or IPv6 IP address ip6-prefix - IPv6 prefix macro string uint32",
			},
			"vendor_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IANA allocated a specific enterprise identification number.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *UserManagerAttributeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan UserManagerAttributeModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Default.IsNull() || plan.Default.IsUnknown()) {
		body["default"] = plan.Default.ValueString()
	}
	if !(plan.DefaultName.IsNull() || plan.DefaultName.IsUnknown()) {
		body["default-name"] = plan.DefaultName.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.PacketTypes.IsNull() || plan.PacketTypes.IsUnknown()) {
		body["packet-types"] = plan.PacketTypes.ValueString()
	}
	if !(plan.StandardName.IsNull() || plan.StandardName.IsUnknown()) {
		body["standard-name"] = plan.StandardName.ValueString()
	}
	if !(plan.TypeID.IsNull() || plan.TypeID.IsUnknown()) {
		body["type-id"] = plan.TypeID.ValueString()
	}
	if !(plan.ValueType.IsNull() || plan.ValueType.IsUnknown()) {
		body["value-type"] = plan.ValueType.ValueString()
	}
	if !(plan.VendorID.IsNull() || plan.VendorID.IsUnknown()) {
		body["vendor-id"] = plan.VendorID.ValueString()
	}
	obj, err := c.Add(ctx, "/user-manager/attribute", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /user-manager/attribute failed", err.Error())
		return
	}
	userManagerAttributeApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserManagerAttributeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state UserManagerAttributeModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/user-manager/attribute", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /user-manager/attribute failed", err.Error())
		return
	}
	userManagerAttributeApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *UserManagerAttributeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state UserManagerAttributeModel
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
	if !plan.Default.Equal(state.Default) && !plan.Default.IsUnknown() {
		body["default"] = plan.Default.ValueString()
	}
	if !plan.DefaultName.Equal(state.DefaultName) && !plan.DefaultName.IsUnknown() {
		body["default-name"] = plan.DefaultName.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.PacketTypes.Equal(state.PacketTypes) && !plan.PacketTypes.IsUnknown() {
		body["packet-types"] = plan.PacketTypes.ValueString()
	}
	if !plan.StandardName.Equal(state.StandardName) && !plan.StandardName.IsUnknown() {
		body["standard-name"] = plan.StandardName.ValueString()
	}
	if !plan.TypeID.Equal(state.TypeID) && !plan.TypeID.IsUnknown() {
		body["type-id"] = plan.TypeID.ValueString()
	}
	if !plan.ValueType.Equal(state.ValueType) && !plan.ValueType.IsUnknown() {
		body["value-type"] = plan.ValueType.ValueString()
	}
	if !plan.VendorID.Equal(state.VendorID) && !plan.VendorID.IsUnknown() {
		body["vendor-id"] = plan.VendorID.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/user-manager/attribute", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /user-manager/attribute failed", err.Error())
			return
		}
		userManagerAttributeApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *UserManagerAttributeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state UserManagerAttributeModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/user-manager/attribute", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /user-manager/attribute failed", err.Error())
	}
}

func (r *UserManagerAttributeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := userManagerAttributeLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /user-manager/attribute matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// userManagerAttributeLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func userManagerAttributeLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/user-manager/attribute", id)
}

func userManagerAttributeApply(ctx context.Context, obj client.Object, m *UserManagerAttributeModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["default"]; ok {
		_ = v
		if v != "" {
			m.Default = types.StringValue(v)
		} else {
			m.Default = types.StringNull()
		}
	} else {
		m.Default = types.StringNull()
	}
	if v, ok := obj["default-name"]; ok {
		_ = v
		if v != "" {
			m.DefaultName = types.StringValue(v)
		} else {
			m.DefaultName = types.StringNull()
		}
	} else {
		m.DefaultName = types.StringNull()
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
	if v, ok := obj["packet-types"]; ok {
		_ = v
		if v != "" {
			m.PacketTypes = types.StringValue(v)
		} else {
			m.PacketTypes = types.StringNull()
		}
	} else {
		m.PacketTypes = types.StringNull()
	}
	if v, ok := obj["standard-name"]; ok {
		_ = v
		if v != "" {
			m.StandardName = types.StringValue(v)
		} else {
			m.StandardName = types.StringNull()
		}
	} else {
		m.StandardName = types.StringNull()
	}
	if v, ok := obj["type-id"]; ok {
		_ = v
		if v != "" {
			m.TypeID = types.StringValue(v)
		} else {
			m.TypeID = types.StringNull()
		}
	} else {
		m.TypeID = types.StringNull()
	}
	if v, ok := obj["value-type"]; ok {
		_ = v
		if v != "" {
			m.ValueType = types.StringValue(v)
		} else {
			m.ValueType = types.StringNull()
		}
	} else {
		m.ValueType = types.StringNull()
	}
	if v, ok := obj["vendor-id"]; ok {
		_ = v
		if v != "" {
			m.VendorID = types.StringValue(v)
		} else {
			m.VendorID = types.StringNull()
		}
	} else {
		m.VendorID = types.StringNull()
	}
}
