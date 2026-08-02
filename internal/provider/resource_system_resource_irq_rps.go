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
	_ resource.Resource                = &SystemResourceIRQRPSResource{}
	_ resource.ResourceWithImportState = &SystemResourceIRQRPSResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemResourceIRQRPSResource struct {
	reg *client.Registry
}

type SystemResourceIRQRPSModel struct {
	ID       types.String `tfsdk:"id"`
	Disabled types.Bool   `tfsdk:"disabled"`
	Name     types.String `tfsdk:"name"`
	Router   types.String `tfsdk:"router"`
}

func NewSystemResourceIRQRPSResource() resource.Resource { return &SystemResourceIRQRPSResource{} }

func (r *SystemResourceIRQRPSResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_resource_irq_rps"
}

func (r *SystemResourceIRQRPSResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemResourceIRQRPSResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/system/resource/irq/rps`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `disabled`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Interface name of the fixed row to adopt (e.g. `ether1`).",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemResourceIRQRPSResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemResourceIRQRPSModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	// Fixed device menu: rows (one per interface) already exist and cannot be
	// added. Adopt the row whose name matches and configure it.
	want := plan.Name.ValueString()
	if want == "" {
		resp.Diagnostics.AddError("name is required",
			"/system/resource/irq/rps has fixed rows; set `name` to the interface to manage (e.g. ether1)")
		return
	}
	rows, err := c.List(ctx, "/system/resource/irq/rps")
	if err != nil {
		resp.Diagnostics.AddError("List /system/resource/irq/rps failed", err.Error())
		return
	}
	var id string
	for _, row := range rows {
		if row["name"] == want {
			id = row[".id"]
			break
		}
	}
	if id == "" {
		resp.Diagnostics.AddError("Unknown irq/rps row "+want,
			fmt.Sprintf("no /system/resource/irq/rps row named %q on the device", want))
		return
	}
	body := client.Object{}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	obj, err := c.Set(ctx, "/system/resource/irq/rps", id, body)
	if err != nil {
		resp.Diagnostics.AddError("Create /system/resource/irq/rps failed", err.Error())
		return
	}
	systemResourceIRQRPSApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemResourceIRQRPSResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemResourceIRQRPSModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/system/resource/irq/rps", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /system/resource/irq/rps failed", err.Error())
		return
	}
	systemResourceIRQRPSApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemResourceIRQRPSResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemResourceIRQRPSModel
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
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	var obj client.Object
	var err error
	if len(body) > 0 {
		obj, err = c.Set(ctx, "/system/resource/irq/rps", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /system/resource/irq/rps failed", err.Error())
			return
		}
	} else {
		obj, err = c.GetByID(ctx, "/system/resource/irq/rps", state.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Read /system/resource/irq/rps failed", err.Error())
			return
		}
	}
	systemResourceIRQRPSApply(ctx, obj, &plan)
	plan.ID = state.ID
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemResourceIRQRPSResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemResourceIRQRPSModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	// Fixed device row: it cannot be removed, only reconfigured. Drop it from
	// Terraform state and leave the row in place.
	_ = c
	resp.State.RemoveResource(ctx)
}

func (r *SystemResourceIRQRPSResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := lookupByNaturalKey(ctx, c, "/system/resource/irq/rps", id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /system/resource/irq/rps matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

func systemResourceIRQRPSApply(ctx context.Context, obj client.Object, m *SystemResourceIRQRPSModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Disabled = types.BoolValue(true)
		} else {
			m.Disabled = types.BoolNull()
		}
	} else {
		m.Disabled = types.BoolNull()
	}
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
}
