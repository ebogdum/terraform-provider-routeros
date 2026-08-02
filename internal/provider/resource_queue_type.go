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
	_ resource.Resource                = &QueueTypeResource{}
	_ resource.ResourceWithImportState = &QueueTypeResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type QueueTypeResource struct {
	reg *client.Registry
}

type QueueTypeModel struct {
	ID                  types.String `tfsdk:"id"`
	FqCodelTarget       types.String `tfsdk:"fq_codel_target"`
	FqCodelQuantum      types.String `tfsdk:"fq_codel_quantum"`
	FqCodelMemlimit     types.String `tfsdk:"fq_codel_memlimit"`
	FqCodelLimit        types.String `tfsdk:"fq_codel_limit"`
	FqCodelInterval     types.String `tfsdk:"fq_codel_interval"`
	FqCodelFlows        types.String `tfsdk:"fq_codel_flows"`
	FqCodelEcn          types.String `tfsdk:"fq_codel_ecn"`
	FqCodelCeThreshold  types.String `tfsdk:"fq_codel_ce_threshold"`
	CodelTarget         types.String `tfsdk:"codel_target"`
	CodelLimit          types.String `tfsdk:"codel_limit"`
	CodelInterval       types.String `tfsdk:"codel_interval"`
	CodelEcn            types.String `tfsdk:"codel_ecn"`
	CodelCeThreshold    types.String `tfsdk:"codel_ce_threshold"`
	CakeWash            types.String `tfsdk:"cake_wash"`
	CakeRttScheme       types.String `tfsdk:"cake_rtt_scheme"`
	CakeRtt             types.String `tfsdk:"cake_rtt"`
	CakeOverheadScheme  types.String `tfsdk:"cake_overhead_scheme"`
	CakeOverhead        types.String `tfsdk:"cake_overhead"`
	CakeNat             types.String `tfsdk:"cake_nat"`
	CakeMpu             types.String `tfsdk:"cake_mpu"`
	CakeMemlimit        types.String `tfsdk:"cake_memlimit"`
	CakeFlowmode        types.String `tfsdk:"cake_flowmode"`
	CakeDiffserv        types.String `tfsdk:"cake_diffserv"`
	CakeBandwidth       types.String `tfsdk:"cake_bandwidth"`
	CakeAutorateIngress types.String `tfsdk:"cake_autorate_ingress"`
	CakeAtm             types.String `tfsdk:"cake_atm"`
	CakeAckFilter       types.String `tfsdk:"cake_ack_filter"`
	BfifoLimit          types.String `tfsdk:"bfifo_limit"`
	Default             types.Bool   `tfsdk:"default"`
	Kind                types.String `tfsdk:"kind"`
	MqPfifoLimit        types.Int64  `tfsdk:"mq_pfifo_limit"`
	Name                types.String `tfsdk:"name"`
	PcqBurstRate        types.Int64  `tfsdk:"pcq_burst_rate"`
	PcqBurstThreshold   types.Int64  `tfsdk:"pcq_burst_threshold"`
	PcqBurstTime        types.String `tfsdk:"pcq_burst_time"`
	PcqClassifier       types.String `tfsdk:"pcq_classifier"`
	PcqDstAddressMask   types.Int64  `tfsdk:"pcq_dst_address_mask"`
	PcqDstAddress6Mask  types.Int64  `tfsdk:"pcq_dst_address6_mask"`
	PcqLimit            types.Int64  `tfsdk:"pcq_limit"`
	PcqRate             types.Int64  `tfsdk:"pcq_rate"`
	PcqSrcAddressMask   types.Int64  `tfsdk:"pcq_src_address_mask"`
	PcqSrcAddress6Mask  types.Int64  `tfsdk:"pcq_src_address6_mask"`
	PcqTotalLimit       types.Int64  `tfsdk:"pcq_total_limit"`
	PfifoLimit          types.Int64  `tfsdk:"pfifo_limit"`
	RedAvgPacket        types.Int64  `tfsdk:"red_avg_packet"`
	RedBurst            types.Int64  `tfsdk:"red_burst"`
	RedLimit            types.Int64  `tfsdk:"red_limit"`
	RedMaxThreshold     types.Int64  `tfsdk:"red_max_threshold"`
	RedMinThreshold     types.Int64  `tfsdk:"red_min_threshold"`
	SfqAllot            types.Int64  `tfsdk:"sfq_allot"`
	SfqPerturb          types.Int64  `tfsdk:"sfq_perturb"`
	TypeName            types.String `tfsdk:"type_name"`
	Router              types.String `tfsdk:"router"`
}

func NewQueueTypeResource() resource.Resource { return &QueueTypeResource{} }

func (r *QueueTypeResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_queue_type"
}

func (r *QueueTypeResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *QueueTypeResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/queue/type`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"fq_codel_target": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fq-codel-target`.",
			},
			"fq_codel_quantum": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fq-codel-quantum`.",
			},
			"fq_codel_memlimit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fq-codel-memlimit`.",
			},
			"fq_codel_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fq-codel-limit`.",
			},
			"fq_codel_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fq-codel-interval`.",
			},
			"fq_codel_flows": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fq-codel-flows`.",
			},
			"fq_codel_ecn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fq-codel-ecn`.",
			},
			"fq_codel_ce_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fq-codel-ce-threshold`.",
			},
			"codel_target": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `codel-target`.",
			},
			"codel_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `codel-limit`.",
			},
			"codel_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `codel-interval`.",
			},
			"codel_ecn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `codel-ecn`.",
			},
			"codel_ce_threshold": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `codel-ce-threshold`.",
			},
			"cake_wash": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cake-wash`.",
			},
			"cake_rtt_scheme": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cake-rtt-scheme`.",
			},
			"cake_rtt": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cake-rtt`.",
			},
			"cake_overhead_scheme": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cake-overhead-scheme`.",
			},
			"cake_overhead": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cake-overhead`.",
			},
			"cake_nat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cake-nat`.",
			},
			"cake_mpu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cake-mpu`.",
			},
			"cake_memlimit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cake-memlimit`.",
			},
			"cake_flowmode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cake-flowmode`.",
			},
			"cake_diffserv": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cake-diffserv`.",
			},
			"cake_bandwidth": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cake-bandwidth`.",
			},
			"cake_autorate_ingress": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cake-autorate-ingress`.",
			},
			"cake_atm": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cake-atm`.",
			},
			"cake_ack_filter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cake-ack-filter`.",
			},
			"bfifo_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bfifo-limit`.",
			},
			"default": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"kind": schema.StringAttribute{
				Required:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "bfifo", "pfifo", "red", "sfq", "pcq", "mq-pfifo", "none", "codel", "fq-codel", "cake"}...)},
			},
			"mq_pfifo_limit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"pcq_burst_rate": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pcq_burst_threshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pcq_burst_time": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"pcq_classifier": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pcq_dst_address_mask": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pcq_dst_address6_mask": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pcq_limit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pcq_rate": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pcq_src_address_mask": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pcq_src_address6_mask": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pcq_total_limit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pfifo_limit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"red_avg_packet": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"red_burst": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"red_limit": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"red_max_threshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"red_min_threshold": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sfq_allot": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"sfq_perturb": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"type_name": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *QueueTypeResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan QueueTypeModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Kind.IsNull() || plan.Kind.IsUnknown()) {
		body["kind"] = plan.Kind.ValueString()
	}
	if !(plan.MqPfifoLimit.IsNull() || plan.MqPfifoLimit.IsUnknown()) {
		body["mq-pfifo-limit"] = client.FormatInt64(plan.MqPfifoLimit.ValueInt64())
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.PcqBurstRate.IsNull() || plan.PcqBurstRate.IsUnknown()) {
		body["pcq-burst-rate"] = client.FormatInt64(plan.PcqBurstRate.ValueInt64())
	}
	if !(plan.PcqBurstThreshold.IsNull() || plan.PcqBurstThreshold.IsUnknown()) {
		body["pcq-burst-threshold"] = client.FormatInt64(plan.PcqBurstThreshold.ValueInt64())
	}
	if !(plan.PcqBurstTime.IsNull() || plan.PcqBurstTime.IsUnknown()) {
		body["pcq-burst-time"] = plan.PcqBurstTime.ValueString()
	}
	if !(plan.PcqClassifier.IsNull() || plan.PcqClassifier.IsUnknown()) {
		body["pcq-classifier"] = plan.PcqClassifier.ValueString()
	}
	if !(plan.PcqDstAddressMask.IsNull() || plan.PcqDstAddressMask.IsUnknown()) {
		body["pcq-dst-address-mask"] = client.FormatInt64(plan.PcqDstAddressMask.ValueInt64())
	}
	if !(plan.PcqDstAddress6Mask.IsNull() || plan.PcqDstAddress6Mask.IsUnknown()) {
		body["pcq-dst-address6-mask"] = client.FormatInt64(plan.PcqDstAddress6Mask.ValueInt64())
	}
	if !(plan.PcqLimit.IsNull() || plan.PcqLimit.IsUnknown()) {
		body["pcq-limit"] = client.FormatInt64(plan.PcqLimit.ValueInt64())
	}
	if !(plan.PcqRate.IsNull() || plan.PcqRate.IsUnknown()) {
		body["pcq-rate"] = client.FormatInt64(plan.PcqRate.ValueInt64())
	}
	if !(plan.PcqSrcAddressMask.IsNull() || plan.PcqSrcAddressMask.IsUnknown()) {
		body["pcq-src-address-mask"] = client.FormatInt64(plan.PcqSrcAddressMask.ValueInt64())
	}
	if !(plan.PcqSrcAddress6Mask.IsNull() || plan.PcqSrcAddress6Mask.IsUnknown()) {
		body["pcq-src-address6-mask"] = client.FormatInt64(plan.PcqSrcAddress6Mask.ValueInt64())
	}
	if !(plan.PcqTotalLimit.IsNull() || plan.PcqTotalLimit.IsUnknown()) {
		body["pcq-total-limit"] = client.FormatInt64(plan.PcqTotalLimit.ValueInt64())
	}
	if !(plan.PfifoLimit.IsNull() || plan.PfifoLimit.IsUnknown()) {
		body["pfifo-limit"] = client.FormatInt64(plan.PfifoLimit.ValueInt64())
	}
	if !(plan.RedAvgPacket.IsNull() || plan.RedAvgPacket.IsUnknown()) {
		body["red-avg-packet"] = client.FormatInt64(plan.RedAvgPacket.ValueInt64())
	}
	if !(plan.RedBurst.IsNull() || plan.RedBurst.IsUnknown()) {
		body["red-burst"] = client.FormatInt64(plan.RedBurst.ValueInt64())
	}
	if !(plan.RedLimit.IsNull() || plan.RedLimit.IsUnknown()) {
		body["red-limit"] = client.FormatInt64(plan.RedLimit.ValueInt64())
	}
	if !(plan.RedMaxThreshold.IsNull() || plan.RedMaxThreshold.IsUnknown()) {
		body["red-max-threshold"] = client.FormatInt64(plan.RedMaxThreshold.ValueInt64())
	}
	if !(plan.RedMinThreshold.IsNull() || plan.RedMinThreshold.IsUnknown()) {
		body["red-min-threshold"] = client.FormatInt64(plan.RedMinThreshold.ValueInt64())
	}
	if !(plan.SfqAllot.IsNull() || plan.SfqAllot.IsUnknown()) {
		body["sfq-allot"] = client.FormatInt64(plan.SfqAllot.ValueInt64())
	}
	if !(plan.SfqPerturb.IsNull() || plan.SfqPerturb.IsUnknown()) {
		body["sfq-perturb"] = client.FormatInt64(plan.SfqPerturb.ValueInt64())
	}
	if !(plan.BfifoLimit.IsNull() || plan.BfifoLimit.IsUnknown()) {
		body["bfifo-limit"] = plan.BfifoLimit.ValueString()
	}
	if !(plan.CakeAckFilter.IsNull() || plan.CakeAckFilter.IsUnknown()) {
		body["cake-ack-filter"] = plan.CakeAckFilter.ValueString()
	}
	if !(plan.CakeAtm.IsNull() || plan.CakeAtm.IsUnknown()) {
		body["cake-atm"] = plan.CakeAtm.ValueString()
	}
	if !(plan.CakeAutorateIngress.IsNull() || plan.CakeAutorateIngress.IsUnknown()) {
		body["cake-autorate-ingress"] = plan.CakeAutorateIngress.ValueString()
	}
	if !(plan.CakeBandwidth.IsNull() || plan.CakeBandwidth.IsUnknown()) {
		body["cake-bandwidth"] = plan.CakeBandwidth.ValueString()
	}
	if !(plan.CakeDiffserv.IsNull() || plan.CakeDiffserv.IsUnknown()) {
		body["cake-diffserv"] = plan.CakeDiffserv.ValueString()
	}
	if !(plan.CakeFlowmode.IsNull() || plan.CakeFlowmode.IsUnknown()) {
		body["cake-flowmode"] = plan.CakeFlowmode.ValueString()
	}
	if !(plan.CakeMemlimit.IsNull() || plan.CakeMemlimit.IsUnknown()) {
		body["cake-memlimit"] = plan.CakeMemlimit.ValueString()
	}
	if !(plan.CakeMpu.IsNull() || plan.CakeMpu.IsUnknown()) {
		body["cake-mpu"] = plan.CakeMpu.ValueString()
	}
	if !(plan.CakeNat.IsNull() || plan.CakeNat.IsUnknown()) {
		body["cake-nat"] = plan.CakeNat.ValueString()
	}
	if !(plan.CakeOverhead.IsNull() || plan.CakeOverhead.IsUnknown()) {
		body["cake-overhead"] = plan.CakeOverhead.ValueString()
	}
	if !(plan.CakeOverheadScheme.IsNull() || plan.CakeOverheadScheme.IsUnknown()) {
		body["cake-overhead-scheme"] = plan.CakeOverheadScheme.ValueString()
	}
	if !(plan.CakeRtt.IsNull() || plan.CakeRtt.IsUnknown()) {
		body["cake-rtt"] = plan.CakeRtt.ValueString()
	}
	if !(plan.CakeRttScheme.IsNull() || plan.CakeRttScheme.IsUnknown()) {
		body["cake-rtt-scheme"] = plan.CakeRttScheme.ValueString()
	}
	if !(plan.CakeWash.IsNull() || plan.CakeWash.IsUnknown()) {
		body["cake-wash"] = plan.CakeWash.ValueString()
	}
	if !(plan.CodelCeThreshold.IsNull() || plan.CodelCeThreshold.IsUnknown()) {
		body["codel-ce-threshold"] = plan.CodelCeThreshold.ValueString()
	}
	if !(plan.CodelEcn.IsNull() || plan.CodelEcn.IsUnknown()) {
		body["codel-ecn"] = plan.CodelEcn.ValueString()
	}
	if !(plan.CodelInterval.IsNull() || plan.CodelInterval.IsUnknown()) {
		body["codel-interval"] = plan.CodelInterval.ValueString()
	}
	if !(plan.CodelLimit.IsNull() || plan.CodelLimit.IsUnknown()) {
		body["codel-limit"] = plan.CodelLimit.ValueString()
	}
	if !(plan.CodelTarget.IsNull() || plan.CodelTarget.IsUnknown()) {
		body["codel-target"] = plan.CodelTarget.ValueString()
	}
	if !(plan.FqCodelCeThreshold.IsNull() || plan.FqCodelCeThreshold.IsUnknown()) {
		body["fq-codel-ce-threshold"] = plan.FqCodelCeThreshold.ValueString()
	}
	if !(plan.FqCodelEcn.IsNull() || plan.FqCodelEcn.IsUnknown()) {
		body["fq-codel-ecn"] = plan.FqCodelEcn.ValueString()
	}
	if !(plan.FqCodelFlows.IsNull() || plan.FqCodelFlows.IsUnknown()) {
		body["fq-codel-flows"] = plan.FqCodelFlows.ValueString()
	}
	if !(plan.FqCodelInterval.IsNull() || plan.FqCodelInterval.IsUnknown()) {
		body["fq-codel-interval"] = plan.FqCodelInterval.ValueString()
	}
	if !(plan.FqCodelLimit.IsNull() || plan.FqCodelLimit.IsUnknown()) {
		body["fq-codel-limit"] = plan.FqCodelLimit.ValueString()
	}
	if !(plan.FqCodelMemlimit.IsNull() || plan.FqCodelMemlimit.IsUnknown()) {
		body["fq-codel-memlimit"] = plan.FqCodelMemlimit.ValueString()
	}
	if !(plan.FqCodelQuantum.IsNull() || plan.FqCodelQuantum.IsUnknown()) {
		body["fq-codel-quantum"] = plan.FqCodelQuantum.ValueString()
	}
	if !(plan.FqCodelTarget.IsNull() || plan.FqCodelTarget.IsUnknown()) {
		body["fq-codel-target"] = plan.FqCodelTarget.ValueString()
	}
	obj, err := c.Add(ctx, "/queue/type", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /queue/type failed", err.Error())
		return
	}
	queueTypeApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *QueueTypeResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state QueueTypeModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/queue/type", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /queue/type failed", err.Error())
		return
	}
	queueTypeApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *QueueTypeResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state QueueTypeModel
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
	if !plan.Kind.Equal(state.Kind) && !plan.Kind.IsUnknown() {
		body["kind"] = plan.Kind.ValueString()
	}
	if !plan.MqPfifoLimit.Equal(state.MqPfifoLimit) && !plan.MqPfifoLimit.IsUnknown() {
		body["mq-pfifo-limit"] = client.FormatInt64(plan.MqPfifoLimit.ValueInt64())
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.PcqBurstRate.Equal(state.PcqBurstRate) && !plan.PcqBurstRate.IsUnknown() {
		body["pcq-burst-rate"] = client.FormatInt64(plan.PcqBurstRate.ValueInt64())
	}
	if !plan.PcqBurstThreshold.Equal(state.PcqBurstThreshold) && !plan.PcqBurstThreshold.IsUnknown() {
		body["pcq-burst-threshold"] = client.FormatInt64(plan.PcqBurstThreshold.ValueInt64())
	}
	if !plan.PcqBurstTime.Equal(state.PcqBurstTime) && !plan.PcqBurstTime.IsUnknown() {
		body["pcq-burst-time"] = plan.PcqBurstTime.ValueString()
	}
	if !plan.PcqClassifier.Equal(state.PcqClassifier) && !plan.PcqClassifier.IsUnknown() {
		body["pcq-classifier"] = plan.PcqClassifier.ValueString()
	}
	if !plan.PcqDstAddressMask.Equal(state.PcqDstAddressMask) && !plan.PcqDstAddressMask.IsUnknown() {
		body["pcq-dst-address-mask"] = client.FormatInt64(plan.PcqDstAddressMask.ValueInt64())
	}
	if !plan.PcqDstAddress6Mask.Equal(state.PcqDstAddress6Mask) && !plan.PcqDstAddress6Mask.IsUnknown() {
		body["pcq-dst-address6-mask"] = client.FormatInt64(plan.PcqDstAddress6Mask.ValueInt64())
	}
	if !plan.PcqLimit.Equal(state.PcqLimit) && !plan.PcqLimit.IsUnknown() {
		body["pcq-limit"] = client.FormatInt64(plan.PcqLimit.ValueInt64())
	}
	if !plan.PcqRate.Equal(state.PcqRate) && !plan.PcqRate.IsUnknown() {
		body["pcq-rate"] = client.FormatInt64(plan.PcqRate.ValueInt64())
	}
	if !plan.PcqSrcAddressMask.Equal(state.PcqSrcAddressMask) && !plan.PcqSrcAddressMask.IsUnknown() {
		body["pcq-src-address-mask"] = client.FormatInt64(plan.PcqSrcAddressMask.ValueInt64())
	}
	if !plan.PcqSrcAddress6Mask.Equal(state.PcqSrcAddress6Mask) && !plan.PcqSrcAddress6Mask.IsUnknown() {
		body["pcq-src-address6-mask"] = client.FormatInt64(plan.PcqSrcAddress6Mask.ValueInt64())
	}
	if !plan.PcqTotalLimit.Equal(state.PcqTotalLimit) && !plan.PcqTotalLimit.IsUnknown() {
		body["pcq-total-limit"] = client.FormatInt64(plan.PcqTotalLimit.ValueInt64())
	}
	if !plan.PfifoLimit.Equal(state.PfifoLimit) && !plan.PfifoLimit.IsUnknown() {
		body["pfifo-limit"] = client.FormatInt64(plan.PfifoLimit.ValueInt64())
	}
	if !plan.RedAvgPacket.Equal(state.RedAvgPacket) && !plan.RedAvgPacket.IsUnknown() {
		body["red-avg-packet"] = client.FormatInt64(plan.RedAvgPacket.ValueInt64())
	}
	if !plan.RedBurst.Equal(state.RedBurst) && !plan.RedBurst.IsUnknown() {
		body["red-burst"] = client.FormatInt64(plan.RedBurst.ValueInt64())
	}
	if !plan.RedLimit.Equal(state.RedLimit) && !plan.RedLimit.IsUnknown() {
		body["red-limit"] = client.FormatInt64(plan.RedLimit.ValueInt64())
	}
	if !plan.RedMaxThreshold.Equal(state.RedMaxThreshold) && !plan.RedMaxThreshold.IsUnknown() {
		body["red-max-threshold"] = client.FormatInt64(plan.RedMaxThreshold.ValueInt64())
	}
	if !plan.RedMinThreshold.Equal(state.RedMinThreshold) && !plan.RedMinThreshold.IsUnknown() {
		body["red-min-threshold"] = client.FormatInt64(plan.RedMinThreshold.ValueInt64())
	}
	if !plan.SfqAllot.Equal(state.SfqAllot) && !plan.SfqAllot.IsUnknown() {
		body["sfq-allot"] = client.FormatInt64(plan.SfqAllot.ValueInt64())
	}
	if !plan.SfqPerturb.Equal(state.SfqPerturb) && !plan.SfqPerturb.IsUnknown() {
		body["sfq-perturb"] = client.FormatInt64(plan.SfqPerturb.ValueInt64())
	}
	if !plan.BfifoLimit.Equal(state.BfifoLimit) && !plan.BfifoLimit.IsUnknown() {
		body["bfifo-limit"] = plan.BfifoLimit.ValueString()
	}
	if !plan.CakeAckFilter.Equal(state.CakeAckFilter) && !plan.CakeAckFilter.IsUnknown() {
		body["cake-ack-filter"] = plan.CakeAckFilter.ValueString()
	}
	if !plan.CakeAtm.Equal(state.CakeAtm) && !plan.CakeAtm.IsUnknown() {
		body["cake-atm"] = plan.CakeAtm.ValueString()
	}
	if !plan.CakeAutorateIngress.Equal(state.CakeAutorateIngress) && !plan.CakeAutorateIngress.IsUnknown() {
		body["cake-autorate-ingress"] = plan.CakeAutorateIngress.ValueString()
	}
	if !plan.CakeBandwidth.Equal(state.CakeBandwidth) && !plan.CakeBandwidth.IsUnknown() {
		body["cake-bandwidth"] = plan.CakeBandwidth.ValueString()
	}
	if !plan.CakeDiffserv.Equal(state.CakeDiffserv) && !plan.CakeDiffserv.IsUnknown() {
		body["cake-diffserv"] = plan.CakeDiffserv.ValueString()
	}
	if !plan.CakeFlowmode.Equal(state.CakeFlowmode) && !plan.CakeFlowmode.IsUnknown() {
		body["cake-flowmode"] = plan.CakeFlowmode.ValueString()
	}
	if !plan.CakeMemlimit.Equal(state.CakeMemlimit) && !plan.CakeMemlimit.IsUnknown() {
		body["cake-memlimit"] = plan.CakeMemlimit.ValueString()
	}
	if !plan.CakeMpu.Equal(state.CakeMpu) && !plan.CakeMpu.IsUnknown() {
		body["cake-mpu"] = plan.CakeMpu.ValueString()
	}
	if !plan.CakeNat.Equal(state.CakeNat) && !plan.CakeNat.IsUnknown() {
		body["cake-nat"] = plan.CakeNat.ValueString()
	}
	if !plan.CakeOverhead.Equal(state.CakeOverhead) && !plan.CakeOverhead.IsUnknown() {
		body["cake-overhead"] = plan.CakeOverhead.ValueString()
	}
	if !plan.CakeOverheadScheme.Equal(state.CakeOverheadScheme) && !plan.CakeOverheadScheme.IsUnknown() {
		body["cake-overhead-scheme"] = plan.CakeOverheadScheme.ValueString()
	}
	if !plan.CakeRtt.Equal(state.CakeRtt) && !plan.CakeRtt.IsUnknown() {
		body["cake-rtt"] = plan.CakeRtt.ValueString()
	}
	if !plan.CakeRttScheme.Equal(state.CakeRttScheme) && !plan.CakeRttScheme.IsUnknown() {
		body["cake-rtt-scheme"] = plan.CakeRttScheme.ValueString()
	}
	if !plan.CakeWash.Equal(state.CakeWash) && !plan.CakeWash.IsUnknown() {
		body["cake-wash"] = plan.CakeWash.ValueString()
	}
	if !plan.CodelCeThreshold.Equal(state.CodelCeThreshold) && !plan.CodelCeThreshold.IsUnknown() {
		body["codel-ce-threshold"] = plan.CodelCeThreshold.ValueString()
	}
	if !plan.CodelEcn.Equal(state.CodelEcn) && !plan.CodelEcn.IsUnknown() {
		body["codel-ecn"] = plan.CodelEcn.ValueString()
	}
	if !plan.CodelInterval.Equal(state.CodelInterval) && !plan.CodelInterval.IsUnknown() {
		body["codel-interval"] = plan.CodelInterval.ValueString()
	}
	if !plan.CodelLimit.Equal(state.CodelLimit) && !plan.CodelLimit.IsUnknown() {
		body["codel-limit"] = plan.CodelLimit.ValueString()
	}
	if !plan.CodelTarget.Equal(state.CodelTarget) && !plan.CodelTarget.IsUnknown() {
		body["codel-target"] = plan.CodelTarget.ValueString()
	}
	if !plan.FqCodelCeThreshold.Equal(state.FqCodelCeThreshold) && !plan.FqCodelCeThreshold.IsUnknown() {
		body["fq-codel-ce-threshold"] = plan.FqCodelCeThreshold.ValueString()
	}
	if !plan.FqCodelEcn.Equal(state.FqCodelEcn) && !plan.FqCodelEcn.IsUnknown() {
		body["fq-codel-ecn"] = plan.FqCodelEcn.ValueString()
	}
	if !plan.FqCodelFlows.Equal(state.FqCodelFlows) && !plan.FqCodelFlows.IsUnknown() {
		body["fq-codel-flows"] = plan.FqCodelFlows.ValueString()
	}
	if !plan.FqCodelInterval.Equal(state.FqCodelInterval) && !plan.FqCodelInterval.IsUnknown() {
		body["fq-codel-interval"] = plan.FqCodelInterval.ValueString()
	}
	if !plan.FqCodelLimit.Equal(state.FqCodelLimit) && !plan.FqCodelLimit.IsUnknown() {
		body["fq-codel-limit"] = plan.FqCodelLimit.ValueString()
	}
	if !plan.FqCodelMemlimit.Equal(state.FqCodelMemlimit) && !plan.FqCodelMemlimit.IsUnknown() {
		body["fq-codel-memlimit"] = plan.FqCodelMemlimit.ValueString()
	}
	if !plan.FqCodelQuantum.Equal(state.FqCodelQuantum) && !plan.FqCodelQuantum.IsUnknown() {
		body["fq-codel-quantum"] = plan.FqCodelQuantum.ValueString()
	}
	if !plan.FqCodelTarget.Equal(state.FqCodelTarget) && !plan.FqCodelTarget.IsUnknown() {
		body["fq-codel-target"] = plan.FqCodelTarget.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/queue/type", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /queue/type failed", err.Error())
			return
		}
		queueTypeApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *QueueTypeResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state QueueTypeModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/queue/type", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /queue/type failed", err.Error())
	}
}

func (r *QueueTypeResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := queueTypeLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /queue/type matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// queueTypeLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func queueTypeLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/queue/type", id)
}

func queueTypeApply(ctx context.Context, obj client.Object, m *QueueTypeModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["fq-codel-target"]; ok && v != "" {
		m.FqCodelTarget = types.StringValue(v)
	} else {
		m.FqCodelTarget = types.StringNull()
	}
	if v, ok := obj["fq-codel-quantum"]; ok && v != "" {
		m.FqCodelQuantum = types.StringValue(v)
	} else {
		m.FqCodelQuantum = types.StringNull()
	}
	if v, ok := obj["fq-codel-memlimit"]; ok && v != "" {
		m.FqCodelMemlimit = types.StringValue(v)
	} else {
		m.FqCodelMemlimit = types.StringNull()
	}
	if v, ok := obj["fq-codel-limit"]; ok && v != "" {
		m.FqCodelLimit = types.StringValue(v)
	} else {
		m.FqCodelLimit = types.StringNull()
	}
	if v, ok := obj["fq-codel-interval"]; ok && v != "" {
		m.FqCodelInterval = types.StringValue(v)
	} else {
		m.FqCodelInterval = types.StringNull()
	}
	if v, ok := obj["fq-codel-flows"]; ok && v != "" {
		m.FqCodelFlows = types.StringValue(v)
	} else {
		m.FqCodelFlows = types.StringNull()
	}
	if v, ok := obj["fq-codel-ecn"]; ok && v != "" {
		m.FqCodelEcn = types.StringValue(v)
	} else {
		m.FqCodelEcn = types.StringNull()
	}
	if v, ok := obj["fq-codel-ce-threshold"]; ok && v != "" {
		m.FqCodelCeThreshold = types.StringValue(v)
	} else {
		m.FqCodelCeThreshold = types.StringNull()
	}
	if v, ok := obj["codel-target"]; ok && v != "" {
		m.CodelTarget = types.StringValue(v)
	} else {
		m.CodelTarget = types.StringNull()
	}
	if v, ok := obj["codel-limit"]; ok && v != "" {
		m.CodelLimit = types.StringValue(v)
	} else {
		m.CodelLimit = types.StringNull()
	}
	if v, ok := obj["codel-interval"]; ok && v != "" {
		m.CodelInterval = types.StringValue(v)
	} else {
		m.CodelInterval = types.StringNull()
	}
	if v, ok := obj["codel-ecn"]; ok && v != "" {
		m.CodelEcn = types.StringValue(v)
	} else {
		m.CodelEcn = types.StringNull()
	}
	if v, ok := obj["codel-ce-threshold"]; ok && v != "" {
		m.CodelCeThreshold = types.StringValue(v)
	} else {
		m.CodelCeThreshold = types.StringNull()
	}
	if v, ok := obj["cake-wash"]; ok && v != "" {
		m.CakeWash = types.StringValue(v)
	} else {
		m.CakeWash = types.StringNull()
	}
	if v, ok := obj["cake-rtt-scheme"]; ok && v != "" {
		m.CakeRttScheme = types.StringValue(v)
	} else {
		m.CakeRttScheme = types.StringNull()
	}
	if v, ok := obj["cake-rtt"]; ok && v != "" {
		m.CakeRtt = types.StringValue(v)
	} else {
		m.CakeRtt = types.StringNull()
	}
	if v, ok := obj["cake-overhead-scheme"]; ok && v != "" {
		m.CakeOverheadScheme = types.StringValue(v)
	} else {
		m.CakeOverheadScheme = types.StringNull()
	}
	if v, ok := obj["cake-overhead"]; ok && v != "" {
		m.CakeOverhead = types.StringValue(v)
	} else {
		m.CakeOverhead = types.StringNull()
	}
	if v, ok := obj["cake-nat"]; ok && v != "" {
		m.CakeNat = types.StringValue(v)
	} else {
		m.CakeNat = types.StringNull()
	}
	if v, ok := obj["cake-mpu"]; ok && v != "" {
		m.CakeMpu = types.StringValue(v)
	} else {
		m.CakeMpu = types.StringNull()
	}
	if v, ok := obj["cake-memlimit"]; ok && v != "" {
		m.CakeMemlimit = types.StringValue(v)
	} else {
		m.CakeMemlimit = types.StringNull()
	}
	if v, ok := obj["cake-flowmode"]; ok && v != "" {
		m.CakeFlowmode = types.StringValue(v)
	} else {
		m.CakeFlowmode = types.StringNull()
	}
	if v, ok := obj["cake-diffserv"]; ok && v != "" {
		m.CakeDiffserv = types.StringValue(v)
	} else {
		m.CakeDiffserv = types.StringNull()
	}
	if v, ok := obj["cake-bandwidth"]; ok && v != "" {
		m.CakeBandwidth = types.StringValue(v)
	} else {
		m.CakeBandwidth = types.StringNull()
	}
	if v, ok := obj["cake-autorate-ingress"]; ok && v != "" {
		m.CakeAutorateIngress = types.StringValue(v)
	} else {
		m.CakeAutorateIngress = types.StringNull()
	}
	if v, ok := obj["cake-atm"]; ok && v != "" {
		m.CakeAtm = types.StringValue(v)
	} else {
		m.CakeAtm = types.StringNull()
	}
	if v, ok := obj["cake-ack-filter"]; ok && v != "" {
		m.CakeAckFilter = types.StringValue(v)
	} else {
		m.CakeAckFilter = types.StringNull()
	}
	if v, ok := obj["bfifo-limit"]; ok && v != "" {
		m.BfifoLimit = types.StringValue(v)
	} else {
		m.BfifoLimit = types.StringNull()
	}
	if v, ok := obj["default"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Default = types.BoolValue(true)
		} else {
			m.Default = types.BoolNull()
		}
	}
	if v, ok := obj["kind"]; ok {
		if v != "" {
			m.Kind = types.StringValue(v)
		} else {
			m.Kind = types.StringNull()
		}
	}
	if v, ok := obj["mq-pfifo-limit"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MqPfifoLimit = types.Int64Value(n)
		} else {
			m.MqPfifoLimit = types.Int64Null()
		}
	} else {
		m.MqPfifoLimit = types.Int64Null()
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["pcq-burst-rate"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PcqBurstRate = types.Int64Value(n)
		} else {
			m.PcqBurstRate = types.Int64Null()
		}
	} else {
		m.PcqBurstRate = types.Int64Null()
	}
	if v, ok := obj["pcq-burst-threshold"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PcqBurstThreshold = types.Int64Value(n)
		} else {
			m.PcqBurstThreshold = types.Int64Null()
		}
	} else {
		m.PcqBurstThreshold = types.Int64Null()
	}
	if v, ok := obj["pcq-burst-time"]; ok {
		if v != "" {
			m.PcqBurstTime = types.StringValue(v)
		} else {
			m.PcqBurstTime = types.StringNull()
		}
	}
	if v, ok := obj["pcq-classifier"]; ok {
		if v != "" {
			m.PcqClassifier = types.StringValue(v)
		} else {
			m.PcqClassifier = types.StringNull()
		}
	}
	if v, ok := obj["pcq-dst-address-mask"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PcqDstAddressMask = types.Int64Value(n)
		} else {
			m.PcqDstAddressMask = types.Int64Null()
		}
	} else {
		m.PcqDstAddressMask = types.Int64Null()
	}
	if v, ok := obj["pcq-dst-address6-mask"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PcqDstAddress6Mask = types.Int64Value(n)
		} else {
			m.PcqDstAddress6Mask = types.Int64Null()
		}
	} else {
		m.PcqDstAddress6Mask = types.Int64Null()
	}
	if v, ok := obj["pcq-limit"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PcqLimit = types.Int64Value(n)
		} else {
			m.PcqLimit = types.Int64Null()
		}
	} else {
		m.PcqLimit = types.Int64Null()
	}
	if v, ok := obj["pcq-rate"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PcqRate = types.Int64Value(n)
		} else {
			m.PcqRate = types.Int64Null()
		}
	} else {
		m.PcqRate = types.Int64Null()
	}
	if v, ok := obj["pcq-src-address-mask"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PcqSrcAddressMask = types.Int64Value(n)
		} else {
			m.PcqSrcAddressMask = types.Int64Null()
		}
	} else {
		m.PcqSrcAddressMask = types.Int64Null()
	}
	if v, ok := obj["pcq-src-address6-mask"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PcqSrcAddress6Mask = types.Int64Value(n)
		} else {
			m.PcqSrcAddress6Mask = types.Int64Null()
		}
	} else {
		m.PcqSrcAddress6Mask = types.Int64Null()
	}
	if v, ok := obj["pcq-total-limit"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PcqTotalLimit = types.Int64Value(n)
		} else {
			m.PcqTotalLimit = types.Int64Null()
		}
	} else {
		m.PcqTotalLimit = types.Int64Null()
	}
	if v, ok := obj["pfifo-limit"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PfifoLimit = types.Int64Value(n)
		} else {
			m.PfifoLimit = types.Int64Null()
		}
	} else {
		m.PfifoLimit = types.Int64Null()
	}
	if v, ok := obj["red-avg-packet"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RedAvgPacket = types.Int64Value(n)
		} else {
			m.RedAvgPacket = types.Int64Null()
		}
	} else {
		m.RedAvgPacket = types.Int64Null()
	}
	if v, ok := obj["red-burst"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RedBurst = types.Int64Value(n)
		} else {
			m.RedBurst = types.Int64Null()
		}
	} else {
		m.RedBurst = types.Int64Null()
	}
	if v, ok := obj["red-limit"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RedLimit = types.Int64Value(n)
		} else {
			m.RedLimit = types.Int64Null()
		}
	} else {
		m.RedLimit = types.Int64Null()
	}
	if v, ok := obj["red-max-threshold"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RedMaxThreshold = types.Int64Value(n)
		} else {
			m.RedMaxThreshold = types.Int64Null()
		}
	} else {
		m.RedMaxThreshold = types.Int64Null()
	}
	if v, ok := obj["red-min-threshold"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RedMinThreshold = types.Int64Value(n)
		} else {
			m.RedMinThreshold = types.Int64Null()
		}
	} else {
		m.RedMinThreshold = types.Int64Null()
	}
	if v, ok := obj["sfq-allot"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.SfqAllot = types.Int64Value(n)
		} else {
			m.SfqAllot = types.Int64Null()
		}
	} else {
		m.SfqAllot = types.Int64Null()
	}
	if v, ok := obj["sfq-perturb"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.SfqPerturb = types.Int64Value(n)
		} else {
			m.SfqPerturb = types.Int64Null()
		}
	} else {
		m.SfqPerturb = types.Int64Null()
	}
	if v, ok := obj["type-name"]; ok {
		if v != "" {
			m.TypeName = types.StringValue(v)
		} else {
			m.TypeName = types.StringNull()
		}
	}
}
