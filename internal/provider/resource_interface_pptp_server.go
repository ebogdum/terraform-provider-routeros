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
	_ resource.Resource                = &InterfacePPTPServerResource{}
	_ resource.ResourceWithImportState = &InterfacePPTPServerResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfacePPTPServerResource struct {
	reg *client.Registry
}

type InterfacePPTPServerModel struct {
	ID       types.String `tfsdk:"id"`
	Comment  types.String `tfsdk:"comment"`
	Disabled types.Bool   `tfsdk:"disabled"`
	Name     types.String `tfsdk:"name"`
	User     types.String `tfsdk:"user"`
	Router   types.String `tfsdk:"router"`
}

func NewInterfacePPTPServerResource() resource.Resource { return &InterfacePPTPServerResource{} }

func (r *InterfacePPTPServerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_pptp_server"
}

func (r *InterfacePPTPServerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfacePPTPServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/pptp-server`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"user": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfacePPTPServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfacePPTPServerModel
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
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.User.IsNull() || plan.User.IsUnknown()) {
		body["user"] = plan.User.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/pptp-server", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/pptp-server failed", err.Error())
		return
	}
	interfacePPTPServerApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfacePPTPServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfacePPTPServerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/pptp-server", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/pptp-server failed", err.Error())
		return
	}
	interfacePPTPServerApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfacePPTPServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfacePPTPServerModel
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
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.User.Equal(state.User) && !plan.User.IsUnknown() {
		body["user"] = plan.User.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/pptp-server", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/pptp-server failed", err.Error())
			return
		}
		interfacePPTPServerApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfacePPTPServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfacePPTPServerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/pptp-server", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/pptp-server failed", err.Error())
	}
}

func (r *InterfacePPTPServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfacePPTPServerLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/pptp-server matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfacePPTPServerLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfacePPTPServerLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/pptp-server", id)
}

func interfacePPTPServerApply(ctx context.Context, obj client.Object, m *InterfacePPTPServerModel) {
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
	if v, ok := obj["user"]; ok {
		_ = v
		if v != "" {
			m.User = types.StringValue(v)
		} else {
			m.User = types.StringNull()
		}
	} else {
		m.User = types.StringNull()
	}
}
