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
	ID               types.String `tfsdk:"id"`
	Comment          types.String `tfsdk:"comment"`
	Default          types.Bool   `tfsdk:"default"`
	DiskFileCount    types.Int64  `tfsdk:"disk_file_count"`
	DiskFileName     types.String `tfsdk:"disk_file_name"`
	DiskLinesPerFile types.Int64  `tfsdk:"disk_lines_per_file"`
	DiskStopOnFull   types.Bool   `tfsdk:"disk_stop_on_full"`
	MemoryLines      types.Int64  `tfsdk:"memory_lines"`
	MemoryStopOnFull types.Bool   `tfsdk:"memory_stop_on_full"`
	Name             types.String `tfsdk:"name"`
	Remember         types.Bool   `tfsdk:"remember"`
	Remote           types.String `tfsdk:"remote"`
	RemoteLogFormat  types.String `tfsdk:"remote_log_format"`
	RemotePort       types.Int64  `tfsdk:"remote_port"`
	RemoteProtocol   types.String `tfsdk:"remote_protocol"`
	SrcAddress       types.String `tfsdk:"src_address"`
	Target           types.String `tfsdk:"target"`
	Vrf              types.String `tfsdk:"vrf"`
	Router           types.String `tfsdk:"router"`
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
	_ = fmt.Sprintf
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
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"default": schema.BoolAttribute{
				Optional:    true,
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
	obj, err := c.Add(ctx, "/system/logging/action", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /system/logging/action failed", err.Error())
		return
	}
	systemLoggingActionApply(ctx, obj, &plan)
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
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.DiskFileCount.Equal(state.DiskFileCount) {
		body["disk-file-count"] = client.FormatInt64(plan.DiskFileCount.ValueInt64())
	}
	if !plan.DiskFileName.Equal(state.DiskFileName) {
		body["disk-file-name"] = plan.DiskFileName.ValueString()
	}
	if !plan.DiskLinesPerFile.Equal(state.DiskLinesPerFile) {
		body["disk-lines-per-file"] = client.FormatInt64(plan.DiskLinesPerFile.ValueInt64())
	}
	if !plan.DiskStopOnFull.Equal(state.DiskStopOnFull) {
		body["disk-stop-on-full"] = client.FormatBool(plan.DiskStopOnFull.ValueBool())
	}
	if !plan.MemoryLines.Equal(state.MemoryLines) {
		body["memory-lines"] = client.FormatInt64(plan.MemoryLines.ValueInt64())
	}
	if !plan.MemoryStopOnFull.Equal(state.MemoryStopOnFull) {
		body["memory-stop-on-full"] = client.FormatBool(plan.MemoryStopOnFull.ValueBool())
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Remember.Equal(state.Remember) {
		body["remember"] = client.FormatBool(plan.Remember.ValueBool())
	}
	if !plan.Remote.Equal(state.Remote) {
		body["remote"] = plan.Remote.ValueString()
	}
	if !plan.RemoteLogFormat.Equal(state.RemoteLogFormat) {
		body["remote-log-format"] = plan.RemoteLogFormat.ValueString()
	}
	if !plan.RemotePort.Equal(state.RemotePort) {
		body["remote-port"] = client.FormatInt64(plan.RemotePort.ValueInt64())
	}
	if !plan.RemoteProtocol.Equal(state.RemoteProtocol) {
		body["remote-protocol"] = plan.RemoteProtocol.ValueString()
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !plan.Target.Equal(state.Target) {
		body["target"] = plan.Target.ValueString()
	}
	if !plan.Vrf.Equal(state.Vrf) {
		body["vrf"] = plan.Vrf.ValueString()
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
	id := req.ID
	routerName := ""
	if i := strings.Index(id, "/"); i > 0 && !strings.HasPrefix(id, "*") {
		routerName, id = id[:i], id[i+1:]
	}
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
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/system/logging/action", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func systemLoggingActionApply(ctx context.Context, obj client.Object, m *SystemLoggingActionModel) {
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
	if v, ok := obj["default"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else {
			m.Default = types.BoolNull()
		}
	} else {
		m.Default = types.BoolNull()
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
		_ = v
		if v != "" {
			m.DiskFileName = types.StringValue(v)
		} else {
			m.DiskFileName = types.StringNull()
		}
	} else {
		m.DiskFileName = types.StringNull()
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
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.DiskStopOnFull = types.BoolValue(b)
		} else {
			m.DiskStopOnFull = types.BoolNull()
		}
	} else {
		m.DiskStopOnFull = types.BoolNull()
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
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.MemoryStopOnFull = types.BoolValue(b)
		} else {
			m.MemoryStopOnFull = types.BoolNull()
		}
	} else {
		m.MemoryStopOnFull = types.BoolNull()
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
	if v, ok := obj["remember"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Remember = types.BoolValue(b)
		} else {
			m.Remember = types.BoolNull()
		}
	} else {
		m.Remember = types.BoolNull()
	}
	if v, ok := obj["remote"]; ok {
		_ = v
		if v != "" {
			m.Remote = types.StringValue(v)
		} else {
			m.Remote = types.StringNull()
		}
	} else {
		m.Remote = types.StringNull()
	}
	if v, ok := obj["remote-log-format"]; ok {
		_ = v
		if v != "" {
			m.RemoteLogFormat = types.StringValue(v)
		} else {
			m.RemoteLogFormat = types.StringNull()
		}
	} else {
		m.RemoteLogFormat = types.StringNull()
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
		_ = v
		if v != "" {
			m.RemoteProtocol = types.StringValue(v)
		} else {
			m.RemoteProtocol = types.StringNull()
		}
	} else {
		m.RemoteProtocol = types.StringNull()
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
	if v, ok := obj["target"]; ok {
		_ = v
		if v != "" {
			m.Target = types.StringValue(v)
		} else {
			m.Target = types.StringNull()
		}
	} else {
		m.Target = types.StringNull()
	}
	if v, ok := obj["vrf"]; ok {
		_ = v
		if v != "" {
			m.Vrf = types.StringValue(v)
		} else {
			m.Vrf = types.StringNull()
		}
	} else {
		m.Vrf = types.StringNull()
	}
}
