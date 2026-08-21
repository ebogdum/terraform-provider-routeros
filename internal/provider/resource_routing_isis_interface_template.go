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
	_ resource.Resource                = &RoutingIsisInterfaceTemplateResource{}
	_ resource.ResourceWithImportState = &RoutingIsisInterfaceTemplateResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingIsisInterfaceTemplateResource struct {
	reg *client.Registry
}

type RoutingIsisInterfaceTemplateModel struct {
	ID                     types.String `tfsdk:"id"`
	Ptp                    types.String `tfsdk:"ptp"`
	Passive                types.String `tfsdk:"passive"`
	Levels                 types.String `tfsdk:"levels"`
	Interfaces             types.String `tfsdk:"interfaces"`
	Instance               types.String `tfsdk:"instance"`
	PtpL2PsnpInterval      types.String `tfsdk:"ptp_l2_psnp_interval"`
	PtpL2Metric            types.String `tfsdk:"ptp_l2_metric"`
	PtpL2CsnpInterval      types.String `tfsdk:"ptp_l2_csnp_interval"`
	PtpL1PsnpInterval      types.String `tfsdk:"ptp_l1_psnp_interval"`
	PtpL1Metric            types.String `tfsdk:"ptp_l1_metric"`
	PtpL1CsnpInterval      types.String `tfsdk:"ptp_l1_csnp_interval"`
	PtpHelloMultiplier     types.String `tfsdk:"ptp_hello_multiplier"`
	PtpHelloInterval       types.String `tfsdk:"ptp_hello_interval"`
	PtpHello3way           types.String `tfsdk:"ptp_hello_3way"`
	BcastL2PsnpInterval    types.String `tfsdk:"bcast_l2_psnp_interval"`
	BcastL2Priority        types.String `tfsdk:"bcast_l2_priority"`
	BcastL2Metric          types.String `tfsdk:"bcast_l2_metric"`
	BcastL2HelloMultiplier types.String `tfsdk:"bcast_l2_hello_multiplier"`
	BcastL2HelloIntervalDr types.String `tfsdk:"bcast_l2_hello_interval_dr"`
	BcastL2HelloInterval   types.String `tfsdk:"bcast_l2_hello_interval"`
	BcastL2CsnpInterval    types.String `tfsdk:"bcast_l2_csnp_interval"`
	BcastL1PsnpInterval    types.String `tfsdk:"bcast_l1_psnp_interval"`
	BcastL1Priority        types.String `tfsdk:"bcast_l1_priority"`
	BcastL1Metric          types.String `tfsdk:"bcast_l1_metric"`
	BcastL1HelloMultiplier types.String `tfsdk:"bcast_l1_hello_multiplier"`
	BcastL1HelloIntervalDr types.String `tfsdk:"bcast_l1_hello_interval_dr"`
	BcastL1HelloInterval   types.String `tfsdk:"bcast_l1_hello_interval"`
	BcastL1CsnpInterval    types.String `tfsdk:"bcast_l1_csnp_interval"`
	Comment                types.String `tfsdk:"comment"`
	Disabled               types.Bool   `tfsdk:"disabled"`
	Router                 types.String `tfsdk:"router"`
}

func NewRoutingIsisInterfaceTemplateResource() resource.Resource {
	return &RoutingIsisInterfaceTemplateResource{}
}

func (r *RoutingIsisInterfaceTemplateResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_isis_interface_template"
}

func (r *RoutingIsisInterfaceTemplateResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingIsisInterfaceTemplateResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "References an existing isis instance; auto-test can't synthesise.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ptp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ptp`.",
			},
			"passive": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `passive`.",
			},
			"levels": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `levels`.",
			},
			"interfaces": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `interfaces`.",
			},
			"instance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `instance`.",
			},
			"ptp_l2_psnp_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ptp.l2.psnp-interval`.",
			},
			"ptp_l2_metric": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ptp.l2.metric`.",
			},
			"ptp_l2_csnp_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ptp.l2.csnp-interval`.",
			},
			"ptp_l1_psnp_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ptp.l1.psnp-interval`.",
			},
			"ptp_l1_metric": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ptp.l1.metric`.",
			},
			"ptp_l1_csnp_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ptp.l1.csnp-interval`.",
			},
			"ptp_hello_multiplier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ptp.hello-multiplier`.",
			},
			"ptp_hello_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ptp.hello-interval`.",
			},
			"ptp_hello_3way": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ptp.hello-3way`.",
			},
			"bcast_l2_psnp_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bcast.l2.psnp-interval`.",
			},
			"bcast_l2_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bcast.l2.priority`.",
			},
			"bcast_l2_metric": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bcast.l2.metric`.",
			},
			"bcast_l2_hello_multiplier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bcast.l2.hello-multiplier`.",
			},
			"bcast_l2_hello_interval_dr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bcast.l2.hello-interval-dr`.",
			},
			"bcast_l2_hello_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bcast.l2.hello-interval`.",
			},
			"bcast_l2_csnp_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bcast.l2.csnp-interval`.",
			},
			"bcast_l1_psnp_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bcast.l1.psnp-interval`.",
			},
			"bcast_l1_priority": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bcast.l1.priority`.",
			},
			"bcast_l1_metric": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bcast.l1.metric`.",
			},
			"bcast_l1_hello_multiplier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bcast.l1.hello-multiplier`.",
			},
			"bcast_l1_hello_interval_dr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bcast.l1.hello-interval-dr`.",
			},
			"bcast_l1_hello_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bcast.l1.hello-interval`.",
			},
			"bcast_l1_csnp_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bcast.l1.csnp-interval`.",
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
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *RoutingIsisInterfaceTemplateResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingIsisInterfaceTemplateModel
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
	if !(plan.BcastL1CsnpInterval.IsNull() || plan.BcastL1CsnpInterval.IsUnknown()) {
		body["bcast.l1.csnp-interval"] = plan.BcastL1CsnpInterval.ValueString()
	}
	if !(plan.BcastL1HelloInterval.IsNull() || plan.BcastL1HelloInterval.IsUnknown()) {
		body["bcast.l1.hello-interval"] = plan.BcastL1HelloInterval.ValueString()
	}
	if !(plan.BcastL1HelloIntervalDr.IsNull() || plan.BcastL1HelloIntervalDr.IsUnknown()) {
		body["bcast.l1.hello-interval-dr"] = plan.BcastL1HelloIntervalDr.ValueString()
	}
	if !(plan.BcastL1HelloMultiplier.IsNull() || plan.BcastL1HelloMultiplier.IsUnknown()) {
		body["bcast.l1.hello-multiplier"] = plan.BcastL1HelloMultiplier.ValueString()
	}
	if !(plan.BcastL1Metric.IsNull() || plan.BcastL1Metric.IsUnknown()) {
		body["bcast.l1.metric"] = plan.BcastL1Metric.ValueString()
	}
	if !(plan.BcastL1Priority.IsNull() || plan.BcastL1Priority.IsUnknown()) {
		body["bcast.l1.priority"] = plan.BcastL1Priority.ValueString()
	}
	if !(plan.BcastL1PsnpInterval.IsNull() || plan.BcastL1PsnpInterval.IsUnknown()) {
		body["bcast.l1.psnp-interval"] = plan.BcastL1PsnpInterval.ValueString()
	}
	if !(plan.BcastL2CsnpInterval.IsNull() || plan.BcastL2CsnpInterval.IsUnknown()) {
		body["bcast.l2.csnp-interval"] = plan.BcastL2CsnpInterval.ValueString()
	}
	if !(plan.BcastL2HelloInterval.IsNull() || plan.BcastL2HelloInterval.IsUnknown()) {
		body["bcast.l2.hello-interval"] = plan.BcastL2HelloInterval.ValueString()
	}
	if !(plan.BcastL2HelloIntervalDr.IsNull() || plan.BcastL2HelloIntervalDr.IsUnknown()) {
		body["bcast.l2.hello-interval-dr"] = plan.BcastL2HelloIntervalDr.ValueString()
	}
	if !(plan.BcastL2HelloMultiplier.IsNull() || plan.BcastL2HelloMultiplier.IsUnknown()) {
		body["bcast.l2.hello-multiplier"] = plan.BcastL2HelloMultiplier.ValueString()
	}
	if !(plan.BcastL2Metric.IsNull() || plan.BcastL2Metric.IsUnknown()) {
		body["bcast.l2.metric"] = plan.BcastL2Metric.ValueString()
	}
	if !(plan.BcastL2Priority.IsNull() || plan.BcastL2Priority.IsUnknown()) {
		body["bcast.l2.priority"] = plan.BcastL2Priority.ValueString()
	}
	if !(plan.BcastL2PsnpInterval.IsNull() || plan.BcastL2PsnpInterval.IsUnknown()) {
		body["bcast.l2.psnp-interval"] = plan.BcastL2PsnpInterval.ValueString()
	}
	if !(plan.PtpHello3way.IsNull() || plan.PtpHello3way.IsUnknown()) {
		body["ptp.hello-3way"] = plan.PtpHello3way.ValueString()
	}
	if !(plan.PtpHelloInterval.IsNull() || plan.PtpHelloInterval.IsUnknown()) {
		body["ptp.hello-interval"] = plan.PtpHelloInterval.ValueString()
	}
	if !(plan.PtpHelloMultiplier.IsNull() || plan.PtpHelloMultiplier.IsUnknown()) {
		body["ptp.hello-multiplier"] = plan.PtpHelloMultiplier.ValueString()
	}
	if !(plan.PtpL1CsnpInterval.IsNull() || plan.PtpL1CsnpInterval.IsUnknown()) {
		body["ptp.l1.csnp-interval"] = plan.PtpL1CsnpInterval.ValueString()
	}
	if !(plan.PtpL1Metric.IsNull() || plan.PtpL1Metric.IsUnknown()) {
		body["ptp.l1.metric"] = plan.PtpL1Metric.ValueString()
	}
	if !(plan.PtpL1PsnpInterval.IsNull() || plan.PtpL1PsnpInterval.IsUnknown()) {
		body["ptp.l1.psnp-interval"] = plan.PtpL1PsnpInterval.ValueString()
	}
	if !(plan.PtpL2CsnpInterval.IsNull() || plan.PtpL2CsnpInterval.IsUnknown()) {
		body["ptp.l2.csnp-interval"] = plan.PtpL2CsnpInterval.ValueString()
	}
	if !(plan.PtpL2Metric.IsNull() || plan.PtpL2Metric.IsUnknown()) {
		body["ptp.l2.metric"] = plan.PtpL2Metric.ValueString()
	}
	if !(plan.PtpL2PsnpInterval.IsNull() || plan.PtpL2PsnpInterval.IsUnknown()) {
		body["ptp.l2.psnp-interval"] = plan.PtpL2PsnpInterval.ValueString()
	}
	if !(plan.Instance.IsNull() || plan.Instance.IsUnknown()) {
		body["instance"] = plan.Instance.ValueString()
	}
	if !(plan.Interfaces.IsNull() || plan.Interfaces.IsUnknown()) {
		body["interfaces"] = plan.Interfaces.ValueString()
	}
	if !(plan.Levels.IsNull() || plan.Levels.IsUnknown()) {
		body["levels"] = plan.Levels.ValueString()
	}
	if !(plan.Passive.IsNull() || plan.Passive.IsUnknown()) {
		body["passive"] = plan.Passive.ValueString()
	}
	if !(plan.Ptp.IsNull() || plan.Ptp.IsUnknown()) {
		body["ptp"] = plan.Ptp.ValueString()
	}
	obj, err := c.Add(ctx, "/routing/isis/interface-template", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/isis/interface-template failed", err.Error())
		return
	}
	routingIsisInterfaceTemplateApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingIsisInterfaceTemplateResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingIsisInterfaceTemplateModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/isis/interface-template", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/isis/interface-template failed", err.Error())
		return
	}
	routingIsisInterfaceTemplateApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingIsisInterfaceTemplateResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingIsisInterfaceTemplateModel
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
	if !plan.BcastL1CsnpInterval.Equal(state.BcastL1CsnpInterval) && !plan.BcastL1CsnpInterval.IsUnknown() {
		body["bcast.l1.csnp-interval"] = plan.BcastL1CsnpInterval.ValueString()
	}
	if !plan.BcastL1HelloInterval.Equal(state.BcastL1HelloInterval) && !plan.BcastL1HelloInterval.IsUnknown() {
		body["bcast.l1.hello-interval"] = plan.BcastL1HelloInterval.ValueString()
	}
	if !plan.BcastL1HelloIntervalDr.Equal(state.BcastL1HelloIntervalDr) && !plan.BcastL1HelloIntervalDr.IsUnknown() {
		body["bcast.l1.hello-interval-dr"] = plan.BcastL1HelloIntervalDr.ValueString()
	}
	if !plan.BcastL1HelloMultiplier.Equal(state.BcastL1HelloMultiplier) && !plan.BcastL1HelloMultiplier.IsUnknown() {
		body["bcast.l1.hello-multiplier"] = plan.BcastL1HelloMultiplier.ValueString()
	}
	if !plan.BcastL1Metric.Equal(state.BcastL1Metric) && !plan.BcastL1Metric.IsUnknown() {
		body["bcast.l1.metric"] = plan.BcastL1Metric.ValueString()
	}
	if !plan.BcastL1Priority.Equal(state.BcastL1Priority) && !plan.BcastL1Priority.IsUnknown() {
		body["bcast.l1.priority"] = plan.BcastL1Priority.ValueString()
	}
	if !plan.BcastL1PsnpInterval.Equal(state.BcastL1PsnpInterval) && !plan.BcastL1PsnpInterval.IsUnknown() {
		body["bcast.l1.psnp-interval"] = plan.BcastL1PsnpInterval.ValueString()
	}
	if !plan.BcastL2CsnpInterval.Equal(state.BcastL2CsnpInterval) && !plan.BcastL2CsnpInterval.IsUnknown() {
		body["bcast.l2.csnp-interval"] = plan.BcastL2CsnpInterval.ValueString()
	}
	if !plan.BcastL2HelloInterval.Equal(state.BcastL2HelloInterval) && !plan.BcastL2HelloInterval.IsUnknown() {
		body["bcast.l2.hello-interval"] = plan.BcastL2HelloInterval.ValueString()
	}
	if !plan.BcastL2HelloIntervalDr.Equal(state.BcastL2HelloIntervalDr) && !plan.BcastL2HelloIntervalDr.IsUnknown() {
		body["bcast.l2.hello-interval-dr"] = plan.BcastL2HelloIntervalDr.ValueString()
	}
	if !plan.BcastL2HelloMultiplier.Equal(state.BcastL2HelloMultiplier) && !plan.BcastL2HelloMultiplier.IsUnknown() {
		body["bcast.l2.hello-multiplier"] = plan.BcastL2HelloMultiplier.ValueString()
	}
	if !plan.BcastL2Metric.Equal(state.BcastL2Metric) && !plan.BcastL2Metric.IsUnknown() {
		body["bcast.l2.metric"] = plan.BcastL2Metric.ValueString()
	}
	if !plan.BcastL2Priority.Equal(state.BcastL2Priority) && !plan.BcastL2Priority.IsUnknown() {
		body["bcast.l2.priority"] = plan.BcastL2Priority.ValueString()
	}
	if !plan.BcastL2PsnpInterval.Equal(state.BcastL2PsnpInterval) && !plan.BcastL2PsnpInterval.IsUnknown() {
		body["bcast.l2.psnp-interval"] = plan.BcastL2PsnpInterval.ValueString()
	}
	if !plan.PtpHello3way.Equal(state.PtpHello3way) && !plan.PtpHello3way.IsUnknown() {
		body["ptp.hello-3way"] = plan.PtpHello3way.ValueString()
	}
	if !plan.PtpHelloInterval.Equal(state.PtpHelloInterval) && !plan.PtpHelloInterval.IsUnknown() {
		body["ptp.hello-interval"] = plan.PtpHelloInterval.ValueString()
	}
	if !plan.PtpHelloMultiplier.Equal(state.PtpHelloMultiplier) && !plan.PtpHelloMultiplier.IsUnknown() {
		body["ptp.hello-multiplier"] = plan.PtpHelloMultiplier.ValueString()
	}
	if !plan.PtpL1CsnpInterval.Equal(state.PtpL1CsnpInterval) && !plan.PtpL1CsnpInterval.IsUnknown() {
		body["ptp.l1.csnp-interval"] = plan.PtpL1CsnpInterval.ValueString()
	}
	if !plan.PtpL1Metric.Equal(state.PtpL1Metric) && !plan.PtpL1Metric.IsUnknown() {
		body["ptp.l1.metric"] = plan.PtpL1Metric.ValueString()
	}
	if !plan.PtpL1PsnpInterval.Equal(state.PtpL1PsnpInterval) && !plan.PtpL1PsnpInterval.IsUnknown() {
		body["ptp.l1.psnp-interval"] = plan.PtpL1PsnpInterval.ValueString()
	}
	if !plan.PtpL2CsnpInterval.Equal(state.PtpL2CsnpInterval) && !plan.PtpL2CsnpInterval.IsUnknown() {
		body["ptp.l2.csnp-interval"] = plan.PtpL2CsnpInterval.ValueString()
	}
	if !plan.PtpL2Metric.Equal(state.PtpL2Metric) && !plan.PtpL2Metric.IsUnknown() {
		body["ptp.l2.metric"] = plan.PtpL2Metric.ValueString()
	}
	if !plan.PtpL2PsnpInterval.Equal(state.PtpL2PsnpInterval) && !plan.PtpL2PsnpInterval.IsUnknown() {
		body["ptp.l2.psnp-interval"] = plan.PtpL2PsnpInterval.ValueString()
	}
	if !plan.Instance.Equal(state.Instance) && !plan.Instance.IsUnknown() {
		body["instance"] = plan.Instance.ValueString()
	}
	if !plan.Interfaces.Equal(state.Interfaces) && !plan.Interfaces.IsUnknown() {
		body["interfaces"] = plan.Interfaces.ValueString()
	}
	if !plan.Levels.Equal(state.Levels) && !plan.Levels.IsUnknown() {
		body["levels"] = plan.Levels.ValueString()
	}
	if !plan.Passive.Equal(state.Passive) && !plan.Passive.IsUnknown() {
		body["passive"] = plan.Passive.ValueString()
	}
	if !plan.Ptp.Equal(state.Ptp) && !plan.Ptp.IsUnknown() {
		body["ptp"] = plan.Ptp.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/isis/interface-template", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/isis/interface-template failed", err.Error())
			return
		}
		routingIsisInterfaceTemplateApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingIsisInterfaceTemplateResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingIsisInterfaceTemplateModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/isis/interface-template", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/isis/interface-template failed", err.Error())
	}
}

func (r *RoutingIsisInterfaceTemplateResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingIsisInterfaceTemplateLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/isis/interface-template matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingIsisInterfaceTemplateLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingIsisInterfaceTemplateLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/isis/interface-template", id)
}

func routingIsisInterfaceTemplateApply(ctx context.Context, obj client.Object, m *RoutingIsisInterfaceTemplateModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["ptp"]; ok && v != "" {
		m.Ptp = types.StringValue(v)
	} else {
		m.Ptp = types.StringNull()
	}
	if v, ok := obj["passive"]; ok && v != "" {
		m.Passive = types.StringValue(v)
	} else {
		m.Passive = types.StringNull()
	}
	if v, ok := obj["levels"]; ok && v != "" {
		m.Levels = types.StringValue(v)
	} else {
		m.Levels = types.StringNull()
	}
	if v, ok := obj["interfaces"]; ok && v != "" {
		m.Interfaces = types.StringValue(v)
	} else {
		m.Interfaces = types.StringNull()
	}
	if v, ok := obj["instance"]; ok && v != "" {
		m.Instance = types.StringValue(v)
	} else {
		m.Instance = types.StringNull()
	}
	if v, ok := obj["ptp.l2.psnp-interval"]; ok && v != "" {
		m.PtpL2PsnpInterval = types.StringValue(v)
	} else {
		m.PtpL2PsnpInterval = types.StringNull()
	}
	if v, ok := obj["ptp.l2.metric"]; ok && v != "" {
		m.PtpL2Metric = types.StringValue(v)
	} else {
		m.PtpL2Metric = types.StringNull()
	}
	if v, ok := obj["ptp.l2.csnp-interval"]; ok && v != "" {
		m.PtpL2CsnpInterval = types.StringValue(v)
	} else {
		m.PtpL2CsnpInterval = types.StringNull()
	}
	if v, ok := obj["ptp.l1.psnp-interval"]; ok && v != "" {
		m.PtpL1PsnpInterval = types.StringValue(v)
	} else {
		m.PtpL1PsnpInterval = types.StringNull()
	}
	if v, ok := obj["ptp.l1.metric"]; ok && v != "" {
		m.PtpL1Metric = types.StringValue(v)
	} else {
		m.PtpL1Metric = types.StringNull()
	}
	if v, ok := obj["ptp.l1.csnp-interval"]; ok && v != "" {
		m.PtpL1CsnpInterval = types.StringValue(v)
	} else {
		m.PtpL1CsnpInterval = types.StringNull()
	}
	if v, ok := obj["ptp.hello-multiplier"]; ok && v != "" {
		m.PtpHelloMultiplier = types.StringValue(v)
	} else {
		m.PtpHelloMultiplier = types.StringNull()
	}
	if v, ok := obj["ptp.hello-interval"]; ok && v != "" {
		m.PtpHelloInterval = types.StringValue(v)
	} else {
		m.PtpHelloInterval = types.StringNull()
	}
	if v, ok := obj["ptp.hello-3way"]; ok && v != "" {
		m.PtpHello3way = types.StringValue(v)
	} else {
		m.PtpHello3way = types.StringNull()
	}
	if v, ok := obj["bcast.l2.psnp-interval"]; ok && v != "" {
		m.BcastL2PsnpInterval = types.StringValue(v)
	} else {
		m.BcastL2PsnpInterval = types.StringNull()
	}
	if v, ok := obj["bcast.l2.priority"]; ok && v != "" {
		m.BcastL2Priority = types.StringValue(v)
	} else {
		m.BcastL2Priority = types.StringNull()
	}
	if v, ok := obj["bcast.l2.metric"]; ok && v != "" {
		m.BcastL2Metric = types.StringValue(v)
	} else {
		m.BcastL2Metric = types.StringNull()
	}
	if v, ok := obj["bcast.l2.hello-multiplier"]; ok && v != "" {
		m.BcastL2HelloMultiplier = types.StringValue(v)
	} else {
		m.BcastL2HelloMultiplier = types.StringNull()
	}
	if v, ok := obj["bcast.l2.hello-interval-dr"]; ok && v != "" {
		m.BcastL2HelloIntervalDr = types.StringValue(v)
	} else {
		m.BcastL2HelloIntervalDr = types.StringNull()
	}
	if v, ok := obj["bcast.l2.hello-interval"]; ok && v != "" {
		m.BcastL2HelloInterval = types.StringValue(v)
	} else {
		m.BcastL2HelloInterval = types.StringNull()
	}
	if v, ok := obj["bcast.l2.csnp-interval"]; ok && v != "" {
		m.BcastL2CsnpInterval = types.StringValue(v)
	} else {
		m.BcastL2CsnpInterval = types.StringNull()
	}
	if v, ok := obj["bcast.l1.psnp-interval"]; ok && v != "" {
		m.BcastL1PsnpInterval = types.StringValue(v)
	} else {
		m.BcastL1PsnpInterval = types.StringNull()
	}
	if v, ok := obj["bcast.l1.priority"]; ok && v != "" {
		m.BcastL1Priority = types.StringValue(v)
	} else {
		m.BcastL1Priority = types.StringNull()
	}
	if v, ok := obj["bcast.l1.metric"]; ok && v != "" {
		m.BcastL1Metric = types.StringValue(v)
	} else {
		m.BcastL1Metric = types.StringNull()
	}
	if v, ok := obj["bcast.l1.hello-multiplier"]; ok && v != "" {
		m.BcastL1HelloMultiplier = types.StringValue(v)
	} else {
		m.BcastL1HelloMultiplier = types.StringNull()
	}
	if v, ok := obj["bcast.l1.hello-interval-dr"]; ok && v != "" {
		m.BcastL1HelloIntervalDr = types.StringValue(v)
	} else {
		m.BcastL1HelloIntervalDr = types.StringNull()
	}
	if v, ok := obj["bcast.l1.hello-interval"]; ok && v != "" {
		m.BcastL1HelloInterval = types.StringValue(v)
	} else {
		m.BcastL1HelloInterval = types.StringNull()
	}
	if v, ok := obj["bcast.l1.csnp-interval"]; ok && v != "" {
		m.BcastL1CsnpInterval = types.StringValue(v)
	} else {
		m.BcastL1CsnpInterval = types.StringNull()
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Disabled = types.BoolValue(true)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
}
