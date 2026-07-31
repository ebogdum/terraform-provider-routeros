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
	_ resource.Resource                = &QueueInterfaceResource{}
	_ resource.ResourceWithImportState = &QueueInterfaceResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type QueueInterfaceResource struct {
	reg *client.Registry
}

type QueueInterfaceModel struct {
	ID        types.String `tfsdk:"id"`
	Interface types.String `tfsdk:"interface"`
	Queue     types.String `tfsdk:"queue"`
	Router    types.String `tfsdk:"router"`
}

func NewQueueInterfaceResource() resource.Resource { return &QueueInterfaceResource{} }

func (r *QueueInterfaceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_queue_interface"
}

func (r *QueueInterfaceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *QueueInterfaceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/queue/interface`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"interface": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "RouterOS `interface`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"queue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `queue`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *QueueInterfaceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan QueueInterfaceModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	rows, err := c.List(ctx, "/queue/interface")
	if err != nil {
		resp.Diagnostics.AddError("Read /queue/interface failed", err.Error())
		return
	}
	want := plan.Interface.ValueString()
	var id string
	for _, row := range rows {
		if row["interface"] == want {
			id = row[".id"]
			break
		}
	}
	if id == "" {
		resp.Diagnostics.AddError("Unknown /queue/interface "+want, fmt.Sprintf("/queue/interface has no row with interface=%q; this is a fixed-row menu and rows cannot be added", want))
		return
	}
	body := client.Object{}
	if !(plan.Queue.IsNull() || plan.Queue.IsUnknown()) {
		body["queue"] = plan.Queue.ValueString()
	}
	obj, err := c.Set(ctx, "/queue/interface", id, body)
	if err != nil {
		resp.Diagnostics.AddError("Adopt /queue/interface failed", err.Error())
		return
	}
	queueInterfaceApply(ctx, obj, &plan)
	plan.ID = types.StringValue(id)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *QueueInterfaceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state QueueInterfaceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/queue/interface", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /queue/interface failed", err.Error())
		return
	}
	queueInterfaceApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *QueueInterfaceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state QueueInterfaceModel
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
	if !plan.Queue.Equal(state.Queue) && !plan.Queue.IsUnknown() {
		body["queue"] = plan.Queue.ValueString()
	}
	var obj client.Object
	var err error
	if len(body) > 0 {
		obj, err = c.Set(ctx, "/queue/interface", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /queue/interface failed", err.Error())
			return
		}
	} else {
		obj, err = c.GetByID(ctx, "/queue/interface", state.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Read /queue/interface failed", err.Error())
			return
		}
	}
	queueInterfaceApply(ctx, obj, &plan)
	plan.ID = state.ID
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *QueueInterfaceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Fixed-row menu: rows cannot be removed. Drop from state; the row keeps
	// its last-applied settings, matching /ip/service semantics.
	_ = ctx
	_ = req
	_ = resp
}

func (r *QueueInterfaceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := lookupByNaturalKey(ctx, c, "/queue/interface", id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /queue/interface matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

func queueInterfaceApply(ctx context.Context, obj client.Object, m *QueueInterfaceModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["interface"]; ok && v != "" {
		m.Interface = types.StringValue(v)
	} else {
		m.Interface = types.StringNull()
	}
	if v, ok := obj["queue"]; ok && v != "" {
		m.Queue = types.StringValue(v)
	} else {
		m.Queue = types.StringNull()
	}
}
