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
	_ resource.Resource                = &SystemLoggingActionResource{}
	_ resource.ResourceWithImportState = &SystemLoggingActionResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemLoggingActionResource struct {
	reg *client.Registry
}

type SystemLoggingActionModel struct {
	ID                types.String `tfsdk:"id"`
	SyslogTimeFormat  types.String `tfsdk:"syslog_time_format"`
	SyslogSeverity    types.String `tfsdk:"syslog_severity"`
	SyslogFacility    types.String `tfsdk:"syslog_facility"`
	Script            types.String `tfsdk:"script"`
	EmailTo           types.String `tfsdk:"email_to"`
	EmailStartTls     types.String `tfsdk:"email_start_tls"`
	EmailCc           types.String `tfsdk:"email_cc"`
	CheckCertificate  types.String `tfsdk:"check_certificate"`
	CefEventDelimiter types.String `tfsdk:"cef_event_delimiter"`
	AddTopicsString   types.String `tfsdk:"add_topics_string"`
	Comment           types.String `tfsdk:"comment"`
	Default           types.Bool   `tfsdk:"default"`
	DiskFileCount     types.Int64  `tfsdk:"disk_file_count"`
	DiskFileName      types.String `tfsdk:"disk_file_name"`
	DiskLinesPerFile  types.Int64  `tfsdk:"disk_lines_per_file"`
	DiskStopOnFull    types.Bool   `tfsdk:"disk_stop_on_full"`
	MemoryLines       types.Int64  `tfsdk:"memory_lines"`
	MemoryStopOnFull  types.Bool   `tfsdk:"memory_stop_on_full"`
	Name              types.String `tfsdk:"name"`
	Remember          types.Bool   `tfsdk:"remember"`
	Remote            types.String `tfsdk:"remote"`
	RemoteLogFormat   types.String `tfsdk:"remote_log_format"`
	RemotePort        types.Int64  `tfsdk:"remote_port"`
	RemoteProtocol    types.String `tfsdk:"remote_protocol"`
	SrcAddress        types.String `tfsdk:"src_address"`
	Target            types.String `tfsdk:"target"`
	Vrf               types.String `tfsdk:"vrf"`
	Router            types.String `tfsdk:"router"`
}

func NewSystemLoggingActionResource() resource.Resource { return &SystemLoggingActionResource{} }

func (r *SystemLoggingActionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_logging_action"
}

func (r *SystemLoggingActionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemLoggingActionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "RouterOS rejects hyphens AND underscores in action names on some 7.x versions; not portable.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"syslog_time_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `syslog-time-format`.",
			},
			"syslog_severity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `syslog-severity`.",
			},
			"syslog_facility": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `syslog-facility`.",
			},
			"script": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `script`.",
			},
			"email_to": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `email-to`.",
			},
			"email_start_tls": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `email-start-tls`.",
			},
			"email_cc": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `email-cc`.",
			},
			"check_certificate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `check-certificate`.",
			},
			"cef_event_delimiter": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `cef-event-delimiter`.",
			},
			"add_topics_string": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `add-topics-string`.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"default": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"disk_file_count": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disk_file_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disk_lines_per_file": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disk_stop_on_full": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"memory_lines": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"memory_stop_on_full": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remember": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_log_format": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsIP()},
			},
			"target": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vrf": schema.StringAttribute{
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

func (r *SystemLoggingActionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemLoggingActionModel
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
	if !(plan.DiskFileCount.IsNull() || plan.DiskFileCount.IsUnknown()) {
		body["disk-file-count"] = client.FormatInt64(plan.DiskFileCount.ValueInt64())
	}
	if !(plan.DiskFileName.IsNull() || plan.DiskFileName.IsUnknown()) {
		body["disk-file-name"] = plan.DiskFileName.ValueString()
	}
	if !(plan.DiskLinesPerFile.IsNull() || plan.DiskLinesPerFile.IsUnknown()) {
		body["disk-lines-per-file"] = client.FormatInt64(plan.DiskLinesPerFile.ValueInt64())
	}
	if !(plan.DiskStopOnFull.IsNull() || plan.DiskStopOnFull.IsUnknown()) {
		body["disk-stop-on-full"] = client.FormatBool(plan.DiskStopOnFull.ValueBool())
	}
	if !(plan.MemoryLines.IsNull() || plan.MemoryLines.IsUnknown()) {
		body["memory-lines"] = client.FormatInt64(plan.MemoryLines.ValueInt64())
	}
	if !(plan.MemoryStopOnFull.IsNull() || plan.MemoryStopOnFull.IsUnknown()) {
		body["memory-stop-on-full"] = client.FormatBool(plan.MemoryStopOnFull.ValueBool())
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Remember.IsNull() || plan.Remember.IsUnknown()) {
		body["remember"] = client.FormatBool(plan.Remember.ValueBool())
	}
	if !(plan.Remote.IsNull() || plan.Remote.IsUnknown()) {
		body["remote"] = plan.Remote.ValueString()
	}
	if !(plan.RemoteLogFormat.IsNull() || plan.RemoteLogFormat.IsUnknown()) {
		body["remote-log-format"] = plan.RemoteLogFormat.ValueString()
	}
	if !(plan.RemotePort.IsNull() || plan.RemotePort.IsUnknown()) {
		body["remote-port"] = client.FormatInt64(plan.RemotePort.ValueInt64())
	}
	if !(plan.RemoteProtocol.IsNull() || plan.RemoteProtocol.IsUnknown()) {
		body["remote-protocol"] = plan.RemoteProtocol.ValueString()
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !(plan.Target.IsNull() || plan.Target.IsUnknown()) {
		body["target"] = plan.Target.ValueString()
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if !(plan.AddTopicsString.IsNull() || plan.AddTopicsString.IsUnknown()) {
		body["add-topics-string"] = plan.AddTopicsString.ValueString()
	}
	if !(plan.CefEventDelimiter.IsNull() || plan.CefEventDelimiter.IsUnknown()) {
		body["cef-event-delimiter"] = plan.CefEventDelimiter.ValueString()
	}
	if !(plan.CheckCertificate.IsNull() || plan.CheckCertificate.IsUnknown()) {
		body["check-certificate"] = plan.CheckCertificate.ValueString()
	}
	if !(plan.EmailCc.IsNull() || plan.EmailCc.IsUnknown()) {
		body["email-cc"] = plan.EmailCc.ValueString()
	}
	if !(plan.EmailStartTls.IsNull() || plan.EmailStartTls.IsUnknown()) {
		body["email-start-tls"] = plan.EmailStartTls.ValueString()
	}
	if !(plan.EmailTo.IsNull() || plan.EmailTo.IsUnknown()) {
		body["email-to"] = plan.EmailTo.ValueString()
	}
	if !(plan.Script.IsNull() || plan.Script.IsUnknown()) {
		body["script"] = plan.Script.ValueString()
	}
	if !(plan.SyslogFacility.IsNull() || plan.SyslogFacility.IsUnknown()) {
		body["syslog-facility"] = plan.SyslogFacility.ValueString()
	}
	if !(plan.SyslogSeverity.IsNull() || plan.SyslogSeverity.IsUnknown()) {
		body["syslog-severity"] = plan.SyslogSeverity.ValueString()
	}
	if !(plan.SyslogTimeFormat.IsNull() || plan.SyslogTimeFormat.IsUnknown()) {
		body["syslog-time-format"] = plan.SyslogTimeFormat.ValueString()
	}
	obj, err := c.Add(ctx, "/system/logging/action", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /system/logging/action failed", err.Error())
		return
	}
	systemLoggingActionApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemLoggingActionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemLoggingActionModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/system/logging/action", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /system/logging/action failed", err.Error())
		return
	}
	systemLoggingActionApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemLoggingActionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemLoggingActionModel
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
	if !plan.DiskFileCount.Equal(state.DiskFileCount) && !plan.DiskFileCount.IsUnknown() {
		body["disk-file-count"] = client.FormatInt64(plan.DiskFileCount.ValueInt64())
	}
	if !plan.DiskFileName.Equal(state.DiskFileName) && !plan.DiskFileName.IsUnknown() {
		body["disk-file-name"] = plan.DiskFileName.ValueString()
	}
	if !plan.DiskLinesPerFile.Equal(state.DiskLinesPerFile) && !plan.DiskLinesPerFile.IsUnknown() {
		body["disk-lines-per-file"] = client.FormatInt64(plan.DiskLinesPerFile.ValueInt64())
	}
	if !plan.DiskStopOnFull.Equal(state.DiskStopOnFull) && !plan.DiskStopOnFull.IsUnknown() {
		body["disk-stop-on-full"] = client.FormatBool(plan.DiskStopOnFull.ValueBool())
	}
	if !plan.MemoryLines.Equal(state.MemoryLines) && !plan.MemoryLines.IsUnknown() {
		body["memory-lines"] = client.FormatInt64(plan.MemoryLines.ValueInt64())
	}
	if !plan.MemoryStopOnFull.Equal(state.MemoryStopOnFull) && !plan.MemoryStopOnFull.IsUnknown() {
		body["memory-stop-on-full"] = client.FormatBool(plan.MemoryStopOnFull.ValueBool())
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Remember.Equal(state.Remember) && !plan.Remember.IsUnknown() {
		body["remember"] = client.FormatBool(plan.Remember.ValueBool())
	}
	if !plan.Remote.Equal(state.Remote) && !plan.Remote.IsUnknown() {
		body["remote"] = plan.Remote.ValueString()
	}
	if !plan.RemoteLogFormat.Equal(state.RemoteLogFormat) && !plan.RemoteLogFormat.IsUnknown() {
		body["remote-log-format"] = plan.RemoteLogFormat.ValueString()
	}
	if !plan.RemotePort.Equal(state.RemotePort) && !plan.RemotePort.IsUnknown() {
		body["remote-port"] = client.FormatInt64(plan.RemotePort.ValueInt64())
	}
	if !plan.RemoteProtocol.Equal(state.RemoteProtocol) && !plan.RemoteProtocol.IsUnknown() {
		body["remote-protocol"] = plan.RemoteProtocol.ValueString()
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) && !plan.SrcAddress.IsUnknown() {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !plan.Target.Equal(state.Target) && !plan.Target.IsUnknown() {
		body["target"] = plan.Target.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) && !plan.Vrf.IsUnknown() {
		body["vrf"] = plan.Vrf.ValueString()
	}
	if !plan.AddTopicsString.Equal(state.AddTopicsString) && !plan.AddTopicsString.IsUnknown() {
		body["add-topics-string"] = plan.AddTopicsString.ValueString()
	}
	if !plan.CefEventDelimiter.Equal(state.CefEventDelimiter) && !plan.CefEventDelimiter.IsUnknown() {
		body["cef-event-delimiter"] = plan.CefEventDelimiter.ValueString()
	}
	if !plan.CheckCertificate.Equal(state.CheckCertificate) && !plan.CheckCertificate.IsUnknown() {
		body["check-certificate"] = plan.CheckCertificate.ValueString()
	}
	if !plan.EmailCc.Equal(state.EmailCc) && !plan.EmailCc.IsUnknown() {
		body["email-cc"] = plan.EmailCc.ValueString()
	}
	if !plan.EmailStartTls.Equal(state.EmailStartTls) && !plan.EmailStartTls.IsUnknown() {
		body["email-start-tls"] = plan.EmailStartTls.ValueString()
	}
	if !plan.EmailTo.Equal(state.EmailTo) && !plan.EmailTo.IsUnknown() {
		body["email-to"] = plan.EmailTo.ValueString()
	}
	if !plan.Script.Equal(state.Script) && !plan.Script.IsUnknown() {
		body["script"] = plan.Script.ValueString()
	}
	if !plan.SyslogFacility.Equal(state.SyslogFacility) && !plan.SyslogFacility.IsUnknown() {
		body["syslog-facility"] = plan.SyslogFacility.ValueString()
	}
	if !plan.SyslogSeverity.Equal(state.SyslogSeverity) && !plan.SyslogSeverity.IsUnknown() {
		body["syslog-severity"] = plan.SyslogSeverity.ValueString()
	}
	if !plan.SyslogTimeFormat.Equal(state.SyslogTimeFormat) && !plan.SyslogTimeFormat.IsUnknown() {
		body["syslog-time-format"] = plan.SyslogTimeFormat.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/system/logging/action", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /system/logging/action failed", err.Error())
			return
		}
		systemLoggingActionApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemLoggingActionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemLoggingActionModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/system/logging/action", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /system/logging/action failed", err.Error())
	}
}

func (r *SystemLoggingActionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := systemLoggingActionLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /system/logging/action matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// systemLoggingActionLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func systemLoggingActionLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/system/logging/action", id)
}

func systemLoggingActionApply(ctx context.Context, obj client.Object, m *SystemLoggingActionModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["syslog-time-format"]; ok && v != "" {
		m.SyslogTimeFormat = types.StringValue(v)
	} else {
		m.SyslogTimeFormat = types.StringNull()
	}
	if v, ok := obj["syslog-severity"]; ok && v != "" {
		m.SyslogSeverity = types.StringValue(v)
	} else {
		m.SyslogSeverity = types.StringNull()
	}
	if v, ok := obj["syslog-facility"]; ok && v != "" {
		m.SyslogFacility = types.StringValue(v)
	} else {
		m.SyslogFacility = types.StringNull()
	}
	if v, ok := obj["script"]; ok && v != "" {
		m.Script = types.StringValue(v)
	} else {
		m.Script = types.StringNull()
	}
	if v, ok := obj["email-to"]; ok && v != "" {
		m.EmailTo = types.StringValue(v)
	} else {
		m.EmailTo = types.StringNull()
	}
	if v, ok := obj["email-start-tls"]; ok && v != "" {
		m.EmailStartTls = types.StringValue(v)
	} else {
		m.EmailStartTls = types.StringNull()
	}
	if v, ok := obj["email-cc"]; ok && v != "" {
		m.EmailCc = types.StringValue(v)
	} else {
		m.EmailCc = types.StringNull()
	}
	if v, ok := obj["check-certificate"]; ok && v != "" {
		m.CheckCertificate = types.StringValue(v)
	} else {
		m.CheckCertificate = types.StringNull()
	}
	if v, ok := obj["cef-event-delimiter"]; ok && v != "" {
		m.CefEventDelimiter = types.StringValue(v)
	} else {
		m.CefEventDelimiter = types.StringNull()
	}
	if v, ok := obj["add-topics-string"]; ok && v != "" {
		m.AddTopicsString = types.StringValue(v)
	} else {
		m.AddTopicsString = types.StringNull()
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
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
	if v, ok := obj["disk-file-count"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.DiskFileCount = types.Int64Value(n)
		} else {
			m.DiskFileCount = types.Int64Null()
		}
	} else {
		m.DiskFileCount = types.Int64Null()
	}
	if v, ok := obj["disk-file-name"]; ok {
		if v != "" {
			m.DiskFileName = types.StringValue(v)
		} else {
			m.DiskFileName = types.StringNull()
		}
	}
	if v, ok := obj["disk-lines-per-file"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.DiskLinesPerFile = types.Int64Value(n)
		} else {
			m.DiskLinesPerFile = types.Int64Null()
		}
	} else {
		m.DiskLinesPerFile = types.Int64Null()
	}
	if v, ok := obj["disk-stop-on-full"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.DiskStopOnFull = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.DiskStopOnFull = types.BoolValue(true)
		} else {
			m.DiskStopOnFull = types.BoolNull()
		}
	}
	if v, ok := obj["memory-lines"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MemoryLines = types.Int64Value(n)
		} else {
			m.MemoryLines = types.Int64Null()
		}
	} else {
		m.MemoryLines = types.Int64Null()
	}
	if v, ok := obj["memory-stop-on-full"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.MemoryStopOnFull = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.MemoryStopOnFull = types.BoolValue(true)
		} else {
			m.MemoryStopOnFull = types.BoolNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["remember"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Remember = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Remember = types.BoolValue(true)
		} else {
			m.Remember = types.BoolNull()
		}
	}
	if v, ok := obj["remote"]; ok {
		if v != "" {
			m.Remote = types.StringValue(v)
		} else {
			m.Remote = types.StringNull()
		}
	}
	if v, ok := obj["remote-log-format"]; ok {
		if v != "" {
			m.RemoteLogFormat = types.StringValue(v)
		} else {
			m.RemoteLogFormat = types.StringNull()
		}
	}
	if v, ok := obj["remote-port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.RemotePort = types.Int64Value(n)
		} else {
			m.RemotePort = types.Int64Null()
		}
	} else {
		m.RemotePort = types.Int64Null()
	}
	if v, ok := obj["remote-protocol"]; ok {
		if v != "" {
			m.RemoteProtocol = types.StringValue(v)
		} else {
			m.RemoteProtocol = types.StringNull()
		}
	}
	if v, ok := obj["src-address"]; ok {
		if v != "" {
			m.SrcAddress = types.StringValue(v)
		} else {
			m.SrcAddress = types.StringNull()
		}
	}
	if v, ok := obj["target"]; ok {
		if v != "" {
			m.Target = types.StringValue(v)
		} else {
			m.Target = types.StringNull()
		}
	}
	if v, ok := obj["vrf"]; ok {
		if v != "" {
			m.Vrf = types.StringValue(v)
		} else {
			m.Vrf = types.StringNull()
		}
	}
}
