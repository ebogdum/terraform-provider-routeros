package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &InterfaceMacsecProfileResource{}
	_ resource.ResourceWithImportState = &InterfaceMacsecProfileResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceMacsecProfileResource struct {
	reg *client.Registry
}

type InterfaceMacsecProfileModel struct {
	ID             types.String `tfsdk:"id"`
	Ciphers        types.String `tfsdk:"ciphers"`
	DefaultName    types.String `tfsdk:"default_name"`
	Name           types.String `tfsdk:"name"`
	ServerPriority types.Int64  `tfsdk:"server_priority"`
	Router         types.String `tfsdk:"router"`
}

func NewInterfaceMacsecProfileResource() resource.Resource { return &InterfaceMacsecProfileResource{} }

func (r *InterfaceMacsecProfileResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_macsec_profile"
}

func (r *InterfaceMacsecProfileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *InterfaceMacsecProfileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/macsec/profile`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ciphers": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ciphers`.",
			},
			"default_name": schema.StringAttribute{
				Computed:    true,
				Description: "RouterOS `default-name`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `name`.",
			},
			"server_priority": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `server-priority`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *InterfaceMacsecProfileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceMacsecProfileModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.ServerPriority.IsNull() || plan.ServerPriority.IsUnknown()) {
		body["server-priority"] = client.FormatInt64(plan.ServerPriority.ValueInt64())
	}
	if !(plan.Ciphers.IsNull() || plan.Ciphers.IsUnknown()) {
		body["ciphers"] = plan.Ciphers.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/macsec/profile", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/macsec/profile failed", err.Error())
		return
	}
	interfaceMacsecProfileApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceMacsecProfileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceMacsecProfileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/macsec/profile", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/macsec/profile failed", err.Error())
		return
	}
	interfaceMacsecProfileApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceMacsecProfileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceMacsecProfileModel
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
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.ServerPriority.Equal(state.ServerPriority) && !plan.ServerPriority.IsUnknown() {
		body["server-priority"] = client.FormatInt64(plan.ServerPriority.ValueInt64())
	}
	if !plan.Ciphers.Equal(state.Ciphers) && !plan.Ciphers.IsUnknown() {
		body["ciphers"] = plan.Ciphers.ValueString()
	}
	var obj client.Object
	var err error
	if len(body) > 0 {
		obj, err = c.Set(ctx, "/interface/macsec/profile", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/macsec/profile failed", err.Error())
			return
		}
	} else {
		obj, err = c.GetByID(ctx, "/interface/macsec/profile", state.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Read /interface/macsec/profile failed", err.Error())
			return
		}
	}
	interfaceMacsecProfileApply(ctx, obj, &plan)
	plan.ID = state.ID
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceMacsecProfileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceMacsecProfileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/macsec/profile", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/macsec/profile failed", err.Error())
	}
}

func (r *InterfaceMacsecProfileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := lookupByNaturalKey(ctx, c, "/interface/macsec/profile", id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/macsec/profile matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

func interfaceMacsecProfileApply(ctx context.Context, obj client.Object, m *InterfaceMacsecProfileModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["ciphers"]; ok && v != "" {
		m.Ciphers = types.StringValue(v)
	} else {
		m.Ciphers = types.StringNull()
	}
	if v, ok := obj["default-name"]; ok && v != "" {
		m.DefaultName = types.StringValue(v)
	} else {
		m.DefaultName = types.StringNull()
	}
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["server-priority"]; ok {
		if n, err := client.ParseInt64(v); err == nil {
			m.ServerPriority = types.Int64Value(n)
		} else {
			m.ServerPriority = types.Int64Null()
		}
	} else {
		m.ServerPriority = types.Int64Null()
	}
}
