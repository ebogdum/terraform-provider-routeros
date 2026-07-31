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
	_ resource.Resource                = &InterfaceDot1xClientResource{}
	_ resource.ResourceWithImportState = &InterfaceDot1xClientResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceDot1xClientResource struct {
	reg *client.Registry
}

type InterfaceDot1xClientModel struct {
	ID           types.String `tfsdk:"id"`
	AnonIdentity types.String `tfsdk:"anon_identity"`
	Certificate  types.String `tfsdk:"certificate"`
	Comment      types.String `tfsdk:"comment"`
	Disabled     types.Bool   `tfsdk:"disabled"`
	EAPMethods   types.String `tfsdk:"eap_methods"`
	Identity     types.String `tfsdk:"identity"`
	Interface    types.String `tfsdk:"interface"`
	Invalid      types.Bool   `tfsdk:"invalid"`
	Password     types.String `tfsdk:"password"`
	Router       types.String `tfsdk:"router"`
}

func NewInterfaceDot1xClientResource() resource.Resource { return &InterfaceDot1xClientResource{} }

func (r *InterfaceDot1xClientResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_dot1x_client"
}

func (r *InterfaceDot1xClientResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceDot1xClientResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "802.1X client attaches to a specific Ethernet interface; values vary per device. Skipped.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"anon_identity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"certificate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
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
			"eap_methods": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"identity": schema.StringAttribute{
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
				Computed:    true,
				Description: "",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceDot1xClientResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceDot1xClientModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AnonIdentity.IsNull() || plan.AnonIdentity.IsUnknown()) {
		body["anon-identity"] = plan.AnonIdentity.ValueString()
	}
	if !(plan.Certificate.IsNull() || plan.Certificate.IsUnknown()) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.EAPMethods.IsNull() || plan.EAPMethods.IsUnknown()) {
		body["eap-methods"] = plan.EAPMethods.ValueString()
	}
	if !(plan.Identity.IsNull() || plan.Identity.IsUnknown()) {
		body["identity"] = plan.Identity.ValueString()
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.Password.IsNull() || plan.Password.IsUnknown()) {
		body["password"] = plan.Password.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/dot1x/client", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/dot1x/client failed", err.Error())
		return
	}
	interfaceDot1xClientApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceDot1xClientResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceDot1xClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/dot1x/client", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/dot1x/client failed", err.Error())
		return
	}
	interfaceDot1xClientApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceDot1xClientResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceDot1xClientModel
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
	if !plan.AnonIdentity.Equal(state.AnonIdentity) {
		body["anon-identity"] = plan.AnonIdentity.ValueString()
	}
	if !plan.Certificate.Equal(state.Certificate) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.EAPMethods.Equal(state.EAPMethods) {
		body["eap-methods"] = plan.EAPMethods.ValueString()
	}
	if !plan.Identity.Equal(state.Identity) {
		body["identity"] = plan.Identity.ValueString()
	}
	if !plan.Interface.Equal(state.Interface) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.Password.Equal(state.Password) {
		body["password"] = plan.Password.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/dot1x/client", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/dot1x/client failed", err.Error())
			return
		}
		interfaceDot1xClientApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceDot1xClientResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceDot1xClientModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/dot1x/client", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/dot1x/client failed", err.Error())
	}
}

func (r *InterfaceDot1xClientResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := interfaceDot1xClientLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/dot1x/client matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceDot1xClientLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceDot1xClientLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/interface/dot1x/client", id)
}

func interfaceDot1xClientApply(ctx context.Context, obj client.Object, m *InterfaceDot1xClientModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["anon-identity"]; ok {
		_ = v
		if v != "" {
			m.AnonIdentity = types.StringValue(v)
		} else {
			m.AnonIdentity = types.StringNull()
		}
	} else {
		m.AnonIdentity = types.StringNull()
	}
	if v, ok := obj["certificate"]; ok {
		_ = v
		if v != "" {
			m.Certificate = types.StringValue(v)
		} else {
			m.Certificate = types.StringNull()
		}
	} else {
		m.Certificate = types.StringNull()
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
	if v, ok := obj["eap-methods"]; ok {
		_ = v
		if v != "" {
			m.EAPMethods = types.StringValue(v)
		} else {
			m.EAPMethods = types.StringNull()
		}
	} else {
		m.EAPMethods = types.StringNull()
	}
	if v, ok := obj["identity"]; ok {
		_ = v
		if v != "" {
			m.Identity = types.StringValue(v)
		} else {
			m.Identity = types.StringNull()
		}
	} else {
		m.Identity = types.StringNull()
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
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.Password already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["password"]; ok && v != "" {
		_ = v
		if v != "" {
			m.Password = types.StringValue(v)
		} else {
			m.Password = types.StringNull()
		}
	} else if m.Password.IsUnknown() {
		m.Password = types.StringNull()
	}
}
