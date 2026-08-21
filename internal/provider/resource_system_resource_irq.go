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
	_ resource.Resource                = &SystemResourceIRQResource{}
	_ resource.ResourceWithImportState = &SystemResourceIRQResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemResourceIRQResource struct {
	reg *client.Registry
}

type SystemResourceIRQModel struct {
	ID     types.String `tfsdk:"id"`
	CPU    types.String `tfsdk:"cpu"`
	IRQ    types.String `tfsdk:"irq"`
	Router types.String `tfsdk:"router"`
}

func NewSystemResourceIRQResource() resource.Resource { return &SystemResourceIRQResource{} }

func (r *SystemResourceIRQResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_resource_irq"
}

func (r *SystemResourceIRQResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemResourceIRQResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/system/resource/irq`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"cpu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cpu`.",
			},
			"irq": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "IRQ number of the fixed row to adopt (sets its CPU affinity).",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *SystemResourceIRQResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemResourceIRQModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	// Fixed device menu: one row per hardware IRQ; it rejects Add. Adopt the
	// row whose irq number matches and set its CPU affinity.
	want := plan.IRQ.ValueString()
	if want == "" {
		resp.Diagnostics.AddError("irq is required",
			"/system/resource/irq has fixed rows; set `irq` to the IRQ number to manage")
		return
	}
	rows, err := c.List(ctx, "/system/resource/irq")
	if err != nil {
		resp.Diagnostics.AddError("List /system/resource/irq failed", err.Error())
		return
	}
	var id string
	for _, row := range rows {
		if row["irq"] == want {
			id = row[".id"]
			break
		}
	}
	if id == "" {
		resp.Diagnostics.AddError("Unknown irq row "+want,
			fmt.Sprintf("no /system/resource/irq row with irq=%q on the device", want))
		return
	}
	body := client.Object{}
	if !(plan.CPU.IsNull() || plan.CPU.IsUnknown()) {
		body["cpu"] = plan.CPU.ValueString()
	}
	obj, err := c.Set(ctx, "/system/resource/irq", id, body)
	if err != nil {
		resp.Diagnostics.AddError("Create /system/resource/irq failed", err.Error())
		return
	}
	systemResourceIRQApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemResourceIRQResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemResourceIRQModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/system/resource/irq", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /system/resource/irq failed", err.Error())
		return
	}
	systemResourceIRQApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemResourceIRQResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemResourceIRQModel
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
	if !plan.CPU.Equal(state.CPU) && !plan.CPU.IsUnknown() {
		body["cpu"] = plan.CPU.ValueString()
	}
	var obj client.Object
	var err error
	if len(body) > 0 {
		obj, err = c.Set(ctx, "/system/resource/irq", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /system/resource/irq failed", err.Error())
			return
		}
	} else {
		obj, err = c.GetByID(ctx, "/system/resource/irq", state.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Read /system/resource/irq failed", err.Error())
			return
		}
	}
	systemResourceIRQApply(ctx, obj, &plan)
	plan.ID = state.ID
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemResourceIRQResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemResourceIRQModel
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

func (r *SystemResourceIRQResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := lookupByNaturalKey(ctx, c, "/system/resource/irq", id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /system/resource/irq matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

func systemResourceIRQApply(ctx context.Context, obj client.Object, m *SystemResourceIRQModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["cpu"]; ok && v != "" {
		m.CPU = types.StringValue(v)
	} else {
		m.CPU = types.StringNull()
	}
	if v, ok := obj["irq"]; ok && v != "" {
		m.IRQ = types.StringValue(v)
	} else {
		m.IRQ = types.StringNull()
	}
}
