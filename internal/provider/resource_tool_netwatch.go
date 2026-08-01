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
	_ resource.Resource                = &ToolNetwatchResource{}
	_ resource.ResourceWithImportState = &ToolNetwatchResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type ToolNetwatchResource struct {
	reg *client.Registry
}

type ToolNetwatchModel struct {
	ID                     types.String `tfsdk:"id"`
	UpScript               types.String `tfsdk:"up_script"`
	ThrTcpConnTime         types.String `tfsdk:"thr_tcp_conn_time"`
	ThrStdev               types.String `tfsdk:"thr_stdev"`
	ThrMax                 types.String `tfsdk:"thr_max"`
	ThrLossPercent         types.String `tfsdk:"thr_loss_percent"`
	ThrLossCount           types.String `tfsdk:"thr_loss_count"`
	ThrJitter              types.String `tfsdk:"thr_jitter"`
	ThrHttpTime            types.String `tfsdk:"thr_http_time"`
	ThrAvg                 types.String `tfsdk:"thr_avg"`
	TestScript             types.String `tfsdk:"test_script"`
	StartupDelay           types.String `tfsdk:"startup_delay"`
	StartDelay             types.String `tfsdk:"start_delay"`
	RecordType             types.String `tfsdk:"record_type"`
	PacketSize             types.String `tfsdk:"packet_size"`
	PacketInterval         types.String `tfsdk:"packet_interval"`
	PacketCount            types.String `tfsdk:"packet_count"`
	IgnoreInitialUp        types.String `tfsdk:"ignore_initial_up"`
	IgnoreInitialDown      types.String `tfsdk:"ignore_initial_down"`
	HttpCodes              types.String `tfsdk:"http_codes"`
	EarlySuccessDetection  types.String `tfsdk:"early_success_detection"`
	EarlyFailureDetection  types.String `tfsdk:"early_failure_detection"`
	DownScript             types.String `tfsdk:"down_script"`
	CheckCertificate       types.String `tfsdk:"check_certificate"`
	AcceptIcmpTimeExceeded types.String `tfsdk:"accept_icmp_time_exceeded"`
	Certificate            types.String `tfsdk:"certificate"`
	Comment                types.String `tfsdk:"comment"`
	Disabled               types.Bool   `tfsdk:"disabled"`
	DNSServer              types.String `tfsdk:"dns_server"`
	Host                   types.String `tfsdk:"host"`
	Interval               types.String `tfsdk:"interval"`
	Name                   types.String `tfsdk:"name"`
	Port                   types.String `tfsdk:"port"`
	SrcAddress             types.String `tfsdk:"src_address"`
	Timeout                types.String `tfsdk:"timeout"`
	Ttl                    types.String `tfsdk:"ttl"`
	Type                   types.String `tfsdk:"type"`
	Router                 types.String `tfsdk:"router"`
}

func NewToolNetwatchResource() resource.Resource { return &ToolNetwatchResource{} }

func (r *ToolNetwatchResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_netwatch"
}

func (r *ToolNetwatchResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolNetwatchResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/tool/netwatch`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"up_script": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `up-script`.",
			},
			"thr_tcp_conn_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `thr-tcp-conn-time`.",
			},
			"thr_stdev": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `thr-stdev`.",
			},
			"thr_max": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `thr-max`.",
			},
			"thr_loss_percent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `thr-loss-percent`.",
			},
			"thr_loss_count": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `thr-loss-count`.",
			},
			"thr_jitter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `thr-jitter`.",
			},
			"thr_http_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `thr-http-time`.",
			},
			"thr_avg": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `thr-avg`.",
			},
			"test_script": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `test-script`.",
			},
			"startup_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `startup-delay`.",
			},
			"start_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `start-delay`.",
			},
			"record_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `record-type`.",
			},
			"packet_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `packet-size`.",
			},
			"packet_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `packet-interval`.",
			},
			"packet_count": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `packet-count`.",
			},
			"ignore_initial_up": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ignore-initial-up`.",
			},
			"ignore_initial_down": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ignore-initial-down`.",
			},
			"http_codes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `http-codes`.",
			},
			"early_success_detection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `early-success-detection`.",
			},
			"early_failure_detection": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `early-failure-detection`.",
			},
			"down_script": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `down-script`.",
			},
			"check_certificate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `check-certificate`.",
			},
			"accept_icmp_time_exceeded": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `accept-icmp-time-exceeded`.",
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
			"dns_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"host": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"interval": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ttl": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"type": schema.StringAttribute{
				Optional:    true,
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

func (r *ToolNetwatchResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolNetwatchModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Certificate.IsNull() || plan.Certificate.IsUnknown()) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.DNSServer.IsNull() || plan.DNSServer.IsUnknown()) {
		body["dns-server"] = plan.DNSServer.ValueString()
	}
	if !(plan.Host.IsNull() || plan.Host.IsUnknown()) {
		body["host"] = plan.Host.ValueString()
	}
	if !(plan.Interval.IsNull() || plan.Interval.IsUnknown()) {
		body["interval"] = plan.Interval.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !(plan.Timeout.IsNull() || plan.Timeout.IsUnknown()) {
		body["timeout"] = plan.Timeout.ValueString()
	}
	if !(plan.Ttl.IsNull() || plan.Ttl.IsUnknown()) {
		body["ttl"] = plan.Ttl.ValueString()
	}
	if !(plan.Type.IsNull() || plan.Type.IsUnknown()) {
		body["type"] = plan.Type.ValueString()
	}
	if !(plan.AcceptIcmpTimeExceeded.IsNull() || plan.AcceptIcmpTimeExceeded.IsUnknown()) {
		body["accept-icmp-time-exceeded"] = plan.AcceptIcmpTimeExceeded.ValueString()
	}
	if !(plan.CheckCertificate.IsNull() || plan.CheckCertificate.IsUnknown()) {
		body["check-certificate"] = plan.CheckCertificate.ValueString()
	}
	if !(plan.DownScript.IsNull() || plan.DownScript.IsUnknown()) {
		body["down-script"] = plan.DownScript.ValueString()
	}
	if !(plan.EarlyFailureDetection.IsNull() || plan.EarlyFailureDetection.IsUnknown()) {
		body["early-failure-detection"] = plan.EarlyFailureDetection.ValueString()
	}
	if !(plan.EarlySuccessDetection.IsNull() || plan.EarlySuccessDetection.IsUnknown()) {
		body["early-success-detection"] = plan.EarlySuccessDetection.ValueString()
	}
	if !(plan.HttpCodes.IsNull() || plan.HttpCodes.IsUnknown()) {
		body["http-codes"] = plan.HttpCodes.ValueString()
	}
	if !(plan.IgnoreInitialDown.IsNull() || plan.IgnoreInitialDown.IsUnknown()) {
		body["ignore-initial-down"] = plan.IgnoreInitialDown.ValueString()
	}
	if !(plan.IgnoreInitialUp.IsNull() || plan.IgnoreInitialUp.IsUnknown()) {
		body["ignore-initial-up"] = plan.IgnoreInitialUp.ValueString()
	}
	if !(plan.PacketCount.IsNull() || plan.PacketCount.IsUnknown()) {
		body["packet-count"] = plan.PacketCount.ValueString()
	}
	if !(plan.PacketInterval.IsNull() || plan.PacketInterval.IsUnknown()) {
		body["packet-interval"] = plan.PacketInterval.ValueString()
	}
	if !(plan.PacketSize.IsNull() || plan.PacketSize.IsUnknown()) {
		body["packet-size"] = plan.PacketSize.ValueString()
	}
	if !(plan.RecordType.IsNull() || plan.RecordType.IsUnknown()) {
		body["record-type"] = plan.RecordType.ValueString()
	}
	if !(plan.StartDelay.IsNull() || plan.StartDelay.IsUnknown()) {
		body["start-delay"] = plan.StartDelay.ValueString()
	}
	if !(plan.StartupDelay.IsNull() || plan.StartupDelay.IsUnknown()) {
		body["startup-delay"] = plan.StartupDelay.ValueString()
	}
	if !(plan.TestScript.IsNull() || plan.TestScript.IsUnknown()) {
		body["test-script"] = plan.TestScript.ValueString()
	}
	if !(plan.ThrAvg.IsNull() || plan.ThrAvg.IsUnknown()) {
		body["thr-avg"] = plan.ThrAvg.ValueString()
	}
	if !(plan.ThrHttpTime.IsNull() || plan.ThrHttpTime.IsUnknown()) {
		body["thr-http-time"] = plan.ThrHttpTime.ValueString()
	}
	if !(plan.ThrJitter.IsNull() || plan.ThrJitter.IsUnknown()) {
		body["thr-jitter"] = plan.ThrJitter.ValueString()
	}
	if !(plan.ThrLossCount.IsNull() || plan.ThrLossCount.IsUnknown()) {
		body["thr-loss-count"] = plan.ThrLossCount.ValueString()
	}
	if !(plan.ThrLossPercent.IsNull() || plan.ThrLossPercent.IsUnknown()) {
		body["thr-loss-percent"] = plan.ThrLossPercent.ValueString()
	}
	if !(plan.ThrMax.IsNull() || plan.ThrMax.IsUnknown()) {
		body["thr-max"] = plan.ThrMax.ValueString()
	}
	if !(plan.ThrStdev.IsNull() || plan.ThrStdev.IsUnknown()) {
		body["thr-stdev"] = plan.ThrStdev.ValueString()
	}
	if !(plan.ThrTcpConnTime.IsNull() || plan.ThrTcpConnTime.IsUnknown()) {
		body["thr-tcp-conn-time"] = plan.ThrTcpConnTime.ValueString()
	}
	if !(plan.UpScript.IsNull() || plan.UpScript.IsUnknown()) {
		body["up-script"] = plan.UpScript.ValueString()
	}
	obj, err := c.Add(ctx, "/tool/netwatch", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /tool/netwatch failed", err.Error())
		return
	}
	toolNetwatchApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolNetwatchResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolNetwatchModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/tool/netwatch", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /tool/netwatch failed", err.Error())
		return
	}
	toolNetwatchApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolNetwatchResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ToolNetwatchModel
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
	if !plan.Certificate.Equal(state.Certificate) && !plan.Certificate.IsUnknown() {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DNSServer.Equal(state.DNSServer) && !plan.DNSServer.IsUnknown() {
		body["dns-server"] = plan.DNSServer.ValueString()
	}
	if !plan.Host.Equal(state.Host) && !plan.Host.IsUnknown() {
		body["host"] = plan.Host.ValueString()
	}
	if !plan.Interval.Equal(state.Interval) && !plan.Interval.IsUnknown() {
		body["interval"] = plan.Interval.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) && !plan.SrcAddress.IsUnknown() {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !plan.Timeout.Equal(state.Timeout) && !plan.Timeout.IsUnknown() {
		body["timeout"] = plan.Timeout.ValueString()
	}
	if !plan.Ttl.Equal(state.Ttl) && !plan.Ttl.IsUnknown() {
		body["ttl"] = plan.Ttl.ValueString()
	}
	if !plan.Type.Equal(state.Type) && !plan.Type.IsUnknown() {
		body["type"] = plan.Type.ValueString()
	}
	if !plan.AcceptIcmpTimeExceeded.Equal(state.AcceptIcmpTimeExceeded) && !plan.AcceptIcmpTimeExceeded.IsUnknown() {
		body["accept-icmp-time-exceeded"] = plan.AcceptIcmpTimeExceeded.ValueString()
	}
	if !plan.CheckCertificate.Equal(state.CheckCertificate) && !plan.CheckCertificate.IsUnknown() {
		body["check-certificate"] = plan.CheckCertificate.ValueString()
	}
	if !plan.DownScript.Equal(state.DownScript) && !plan.DownScript.IsUnknown() {
		body["down-script"] = plan.DownScript.ValueString()
	}
	if !plan.EarlyFailureDetection.Equal(state.EarlyFailureDetection) && !plan.EarlyFailureDetection.IsUnknown() {
		body["early-failure-detection"] = plan.EarlyFailureDetection.ValueString()
	}
	if !plan.EarlySuccessDetection.Equal(state.EarlySuccessDetection) && !plan.EarlySuccessDetection.IsUnknown() {
		body["early-success-detection"] = plan.EarlySuccessDetection.ValueString()
	}
	if !plan.HttpCodes.Equal(state.HttpCodes) && !plan.HttpCodes.IsUnknown() {
		body["http-codes"] = plan.HttpCodes.ValueString()
	}
	if !plan.IgnoreInitialDown.Equal(state.IgnoreInitialDown) && !plan.IgnoreInitialDown.IsUnknown() {
		body["ignore-initial-down"] = plan.IgnoreInitialDown.ValueString()
	}
	if !plan.IgnoreInitialUp.Equal(state.IgnoreInitialUp) && !plan.IgnoreInitialUp.IsUnknown() {
		body["ignore-initial-up"] = plan.IgnoreInitialUp.ValueString()
	}
	if !plan.PacketCount.Equal(state.PacketCount) && !plan.PacketCount.IsUnknown() {
		body["packet-count"] = plan.PacketCount.ValueString()
	}
	if !plan.PacketInterval.Equal(state.PacketInterval) && !plan.PacketInterval.IsUnknown() {
		body["packet-interval"] = plan.PacketInterval.ValueString()
	}
	if !plan.PacketSize.Equal(state.PacketSize) && !plan.PacketSize.IsUnknown() {
		body["packet-size"] = plan.PacketSize.ValueString()
	}
	if !plan.RecordType.Equal(state.RecordType) && !plan.RecordType.IsUnknown() {
		body["record-type"] = plan.RecordType.ValueString()
	}
	if !plan.StartDelay.Equal(state.StartDelay) && !plan.StartDelay.IsUnknown() {
		body["start-delay"] = plan.StartDelay.ValueString()
	}
	if !plan.StartupDelay.Equal(state.StartupDelay) && !plan.StartupDelay.IsUnknown() {
		body["startup-delay"] = plan.StartupDelay.ValueString()
	}
	if !plan.TestScript.Equal(state.TestScript) && !plan.TestScript.IsUnknown() {
		body["test-script"] = plan.TestScript.ValueString()
	}
	if !plan.ThrAvg.Equal(state.ThrAvg) && !plan.ThrAvg.IsUnknown() {
		body["thr-avg"] = plan.ThrAvg.ValueString()
	}
	if !plan.ThrHttpTime.Equal(state.ThrHttpTime) && !plan.ThrHttpTime.IsUnknown() {
		body["thr-http-time"] = plan.ThrHttpTime.ValueString()
	}
	if !plan.ThrJitter.Equal(state.ThrJitter) && !plan.ThrJitter.IsUnknown() {
		body["thr-jitter"] = plan.ThrJitter.ValueString()
	}
	if !plan.ThrLossCount.Equal(state.ThrLossCount) && !plan.ThrLossCount.IsUnknown() {
		body["thr-loss-count"] = plan.ThrLossCount.ValueString()
	}
	if !plan.ThrLossPercent.Equal(state.ThrLossPercent) && !plan.ThrLossPercent.IsUnknown() {
		body["thr-loss-percent"] = plan.ThrLossPercent.ValueString()
	}
	if !plan.ThrMax.Equal(state.ThrMax) && !plan.ThrMax.IsUnknown() {
		body["thr-max"] = plan.ThrMax.ValueString()
	}
	if !plan.ThrStdev.Equal(state.ThrStdev) && !plan.ThrStdev.IsUnknown() {
		body["thr-stdev"] = plan.ThrStdev.ValueString()
	}
	if !plan.ThrTcpConnTime.Equal(state.ThrTcpConnTime) && !plan.ThrTcpConnTime.IsUnknown() {
		body["thr-tcp-conn-time"] = plan.ThrTcpConnTime.ValueString()
	}
	if !plan.UpScript.Equal(state.UpScript) && !plan.UpScript.IsUnknown() {
		body["up-script"] = plan.UpScript.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/tool/netwatch", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /tool/netwatch failed", err.Error())
			return
		}
		toolNetwatchApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolNetwatchResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ToolNetwatchModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/tool/netwatch", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /tool/netwatch failed", err.Error())
	}
}

func (r *ToolNetwatchResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := toolNetwatchLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /tool/netwatch matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// toolNetwatchLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func toolNetwatchLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/tool/netwatch", id)
}

func toolNetwatchApply(ctx context.Context, obj client.Object, m *ToolNetwatchModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["up-script"]; ok && v != "" {
		m.UpScript = types.StringValue(v)
	} else {
		m.UpScript = types.StringNull()
	}
	if v, ok := obj["thr-tcp-conn-time"]; ok && v != "" {
		m.ThrTcpConnTime = types.StringValue(v)
	} else {
		m.ThrTcpConnTime = types.StringNull()
	}
	if v, ok := obj["thr-stdev"]; ok && v != "" {
		m.ThrStdev = types.StringValue(v)
	} else {
		m.ThrStdev = types.StringNull()
	}
	if v, ok := obj["thr-max"]; ok && v != "" {
		m.ThrMax = types.StringValue(v)
	} else {
		m.ThrMax = types.StringNull()
	}
	if v, ok := obj["thr-loss-percent"]; ok && v != "" {
		m.ThrLossPercent = types.StringValue(v)
	} else {
		m.ThrLossPercent = types.StringNull()
	}
	if v, ok := obj["thr-loss-count"]; ok && v != "" {
		m.ThrLossCount = types.StringValue(v)
	} else {
		m.ThrLossCount = types.StringNull()
	}
	if v, ok := obj["thr-jitter"]; ok && v != "" {
		m.ThrJitter = types.StringValue(v)
	} else {
		m.ThrJitter = types.StringNull()
	}
	if v, ok := obj["thr-http-time"]; ok && v != "" {
		m.ThrHttpTime = types.StringValue(v)
	} else {
		m.ThrHttpTime = types.StringNull()
	}
	if v, ok := obj["thr-avg"]; ok && v != "" {
		m.ThrAvg = types.StringValue(v)
	} else {
		m.ThrAvg = types.StringNull()
	}
	if v, ok := obj["test-script"]; ok && v != "" {
		m.TestScript = types.StringValue(v)
	} else {
		m.TestScript = types.StringNull()
	}
	if v, ok := obj["startup-delay"]; ok && v != "" {
		m.StartupDelay = types.StringValue(v)
	} else {
		m.StartupDelay = types.StringNull()
	}
	if v, ok := obj["start-delay"]; ok && v != "" {
		m.StartDelay = types.StringValue(v)
	} else {
		m.StartDelay = types.StringNull()
	}
	if v, ok := obj["record-type"]; ok && v != "" {
		m.RecordType = types.StringValue(v)
	} else {
		m.RecordType = types.StringNull()
	}
	if v, ok := obj["packet-size"]; ok && v != "" {
		m.PacketSize = types.StringValue(v)
	} else {
		m.PacketSize = types.StringNull()
	}
	if v, ok := obj["packet-interval"]; ok && v != "" {
		m.PacketInterval = types.StringValue(v)
	} else {
		m.PacketInterval = types.StringNull()
	}
	if v, ok := obj["packet-count"]; ok && v != "" {
		m.PacketCount = types.StringValue(v)
	} else {
		m.PacketCount = types.StringNull()
	}
	if v, ok := obj["ignore-initial-up"]; ok && v != "" {
		m.IgnoreInitialUp = types.StringValue(v)
	} else {
		m.IgnoreInitialUp = types.StringNull()
	}
	if v, ok := obj["ignore-initial-down"]; ok && v != "" {
		m.IgnoreInitialDown = types.StringValue(v)
	} else {
		m.IgnoreInitialDown = types.StringNull()
	}
	if v, ok := obj["http-codes"]; ok && v != "" {
		m.HttpCodes = types.StringValue(v)
	} else {
		m.HttpCodes = types.StringNull()
	}
	if v, ok := obj["early-success-detection"]; ok && v != "" {
		m.EarlySuccessDetection = types.StringValue(v)
	} else {
		m.EarlySuccessDetection = types.StringNull()
	}
	if v, ok := obj["early-failure-detection"]; ok && v != "" {
		m.EarlyFailureDetection = types.StringValue(v)
	} else {
		m.EarlyFailureDetection = types.StringNull()
	}
	if v, ok := obj["down-script"]; ok && v != "" {
		m.DownScript = types.StringValue(v)
	} else {
		m.DownScript = types.StringNull()
	}
	if v, ok := obj["check-certificate"]; ok && v != "" {
		m.CheckCertificate = types.StringValue(v)
	} else {
		m.CheckCertificate = types.StringNull()
	}
	if v, ok := obj["accept-icmp-time-exceeded"]; ok && v != "" {
		m.AcceptIcmpTimeExceeded = types.StringValue(v)
	} else {
		m.AcceptIcmpTimeExceeded = types.StringNull()
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
	if v, ok := obj["dns-server"]; ok {
		_ = v
		if v != "" {
			m.DNSServer = types.StringValue(v)
		} else {
			m.DNSServer = types.StringNull()
		}
	} else {
		m.DNSServer = types.StringNull()
	}
	if v, ok := obj["host"]; ok {
		_ = v
		if v != "" {
			m.Host = types.StringValue(v)
		} else {
			m.Host = types.StringNull()
		}
	} else {
		m.Host = types.StringNull()
	}
	if v, ok := obj["interval"]; ok {
		_ = v
		if v != "" {
			m.Interval = types.StringValue(v)
		} else {
			m.Interval = types.StringNull()
		}
	} else {
		m.Interval = types.StringNull()
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
	if v, ok := obj["port"]; ok {
		_ = v
		if v != "" {
			m.Port = types.StringValue(v)
		} else {
			m.Port = types.StringNull()
		}
	} else {
		m.Port = types.StringNull()
	}
	if v, ok := obj["src-address"]; ok {
		_ = v
		if v != "" {
			m.SrcAddress = types.StringValue(v)
		} else {
			m.SrcAddress = types.StringNull()
		}
	} else {
		m.SrcAddress = types.StringNull()
	}
	if v, ok := obj["timeout"]; ok {
		_ = v
		if v != "" {
			m.Timeout = types.StringValue(v)
		} else {
			m.Timeout = types.StringNull()
		}
	} else {
		m.Timeout = types.StringNull()
	}
	if v, ok := obj["ttl"]; ok {
		_ = v
		if v != "" {
			m.Ttl = types.StringValue(v)
		} else {
			m.Ttl = types.StringNull()
		}
	} else {
		m.Ttl = types.StringNull()
	}
	if v, ok := obj["type"]; ok {
		_ = v
		if v != "" {
			m.Type = types.StringValue(v)
		} else {
			m.Type = types.StringNull()
		}
	} else {
		m.Type = types.StringNull()
	}
}
