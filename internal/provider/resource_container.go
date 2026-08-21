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
	_ resource.Resource                = &ContainerResource{}
	_ resource.ResourceWithImportState = &ContainerResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type ContainerResource struct {
	reg *client.Registry
}

type ContainerModel struct {
	ID                  types.String `tfsdk:"id"`
	AutoRestartInterval types.String `tfsdk:"auto_restart_interval"`
	Cmd                 types.String `tfsdk:"cmd"`
	Comment             types.String `tfsdk:"comment"`
	CpuList             types.String `tfsdk:"cpu_list"`
	Devices             types.String `tfsdk:"devices"`
	DNS                 types.String `tfsdk:"dns"`
	DomainName          types.String `tfsdk:"domain_name"`
	Entrypoint          types.String `tfsdk:"entrypoint"`
	Envlist             types.String `tfsdk:"envlist"`
	Hostname            types.String `tfsdk:"hostname"`
	Interface           types.String `tfsdk:"interface"`
	Logging             types.String `tfsdk:"logging"`
	MemoryHigh          types.String `tfsdk:"memory_high"`
	MemoryMax           types.String `tfsdk:"memory_max"`
	Mount               types.String `tfsdk:"mount"`
	Mountlists          types.String `tfsdk:"mountlists"`
	RemoteImage         types.String `tfsdk:"remote_image"`
	RootDir             types.String `tfsdk:"root_dir"`
	StartOnBoot         types.String `tfsdk:"start_on_boot"`
	StopSignal          types.String `tfsdk:"stop_signal"`
	User                types.String `tfsdk:"user"`
	Workdir             types.String `tfsdk:"workdir"`
	Router              types.String `tfsdk:"router"`
}

func NewContainerResource() resource.Resource { return &ContainerResource{} }

func (r *ContainerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_container"
}

func (r *ContainerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ContainerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Requires container package + capable architecture",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"auto_restart_interval": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Specify an interval at which Container will be restarted on Container failure. Example: 10s",
			},
			"cmd": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The main purpose of a CMD is to provide defaults for an executing container. These defaults can include an executable, or they can omit the executable, in which case you must specify an ENTRYPOINT instruction as well.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Short description",
			},
			"cpu_list": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "specifies which CPU cores the container is allowed to run on",
			},
			"devices": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "passes through physical device to the container",
			},
			"dns": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "If container needs different DNS, it can be configured here",
			},
			"domain_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"entrypoint": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "An ENTRYPOINT allows to specify executable to run when starting container. Example: /bin/sh",
			},
			"envlist": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "list of environmental variables (configured under /container envs ) to be used with container",
			},
			"hostname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Assigning a hostname to a container helps in identifying and managing the container more easily",
			},
			"interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "veth interface to be used with the container",
			},
			"logging": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "if set to yes, all container-generated output will be shown in the RouterOS log",
			},
			"memory_high": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RAM usage limit in bytes for a specific container",
			},
			"memory_max": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "max RAM usage limit in bytes per container (The container process will be terminated if the memory-max value is smaller than the container memory-current.)",
			},
			"mount": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "specify directory to be used as a mount",
			},
			"mountlists": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "mounts from /container/mounts/ sub-menu to be used with this container",
			},
			"remote_image": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "the container image name to be installed if an external registry is used (configured under /container/config set registry-url=...)",
			},
			"root_dir": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "used to save container store outside main memory",
			},
			"start_on_boot": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "if set to yes, the container will be started automatically on device start-up.",
			},
			"stop_signal": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Type of Linux signal to send when container was not stopped after 10 seconds",
			},
			"user": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "sets the user and group the container process runs as before execution.",
			},
			"workdir": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "the working directory for cmd entrypoint",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *ContainerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ContainerModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AutoRestartInterval.IsNull() || plan.AutoRestartInterval.IsUnknown()) {
		body["auto-restart-interval"] = plan.AutoRestartInterval.ValueString()
	}
	if !(plan.Cmd.IsNull() || plan.Cmd.IsUnknown()) {
		body["cmd"] = plan.Cmd.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.CpuList.IsNull() || plan.CpuList.IsUnknown()) {
		body["cpu-list"] = plan.CpuList.ValueString()
	}
	if !(plan.Devices.IsNull() || plan.Devices.IsUnknown()) {
		body["devices"] = plan.Devices.ValueString()
	}
	if !(plan.DNS.IsNull() || plan.DNS.IsUnknown()) {
		body["dns"] = plan.DNS.ValueString()
	}
	if !(plan.DomainName.IsNull() || plan.DomainName.IsUnknown()) {
		body["domain-name"] = plan.DomainName.ValueString()
	}
	if !(plan.Entrypoint.IsNull() || plan.Entrypoint.IsUnknown()) {
		body["entrypoint"] = plan.Entrypoint.ValueString()
	}
	if !(plan.Envlist.IsNull() || plan.Envlist.IsUnknown()) {
		body["envlist"] = plan.Envlist.ValueString()
	}
	if !(plan.Hostname.IsNull() || plan.Hostname.IsUnknown()) {
		body["hostname"] = plan.Hostname.ValueString()
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.Logging.IsNull() || plan.Logging.IsUnknown()) {
		body["logging"] = plan.Logging.ValueString()
	}
	if !(plan.MemoryHigh.IsNull() || plan.MemoryHigh.IsUnknown()) {
		body["memory-high"] = plan.MemoryHigh.ValueString()
	}
	if !(plan.MemoryMax.IsNull() || plan.MemoryMax.IsUnknown()) {
		body["memory-max"] = plan.MemoryMax.ValueString()
	}
	if !(plan.Mount.IsNull() || plan.Mount.IsUnknown()) {
		body["mount"] = plan.Mount.ValueString()
	}
	if !(plan.Mountlists.IsNull() || plan.Mountlists.IsUnknown()) {
		body["mountlists"] = plan.Mountlists.ValueString()
	}
	if !(plan.RemoteImage.IsNull() || plan.RemoteImage.IsUnknown()) {
		body["remote-image"] = plan.RemoteImage.ValueString()
	}
	if !(plan.RootDir.IsNull() || plan.RootDir.IsUnknown()) {
		body["root-dir"] = plan.RootDir.ValueString()
	}
	if !(plan.StartOnBoot.IsNull() || plan.StartOnBoot.IsUnknown()) {
		body["start-on-boot"] = plan.StartOnBoot.ValueString()
	}
	if !(plan.StopSignal.IsNull() || plan.StopSignal.IsUnknown()) {
		body["stop-signal"] = plan.StopSignal.ValueString()
	}
	if !(plan.User.IsNull() || plan.User.IsUnknown()) {
		body["user"] = plan.User.ValueString()
	}
	if !(plan.Workdir.IsNull() || plan.Workdir.IsUnknown()) {
		body["workdir"] = plan.Workdir.ValueString()
	}
	obj, err := c.Add(ctx, "/container", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /container failed", err.Error())
		return
	}
	containerApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ContainerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ContainerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/container", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /container failed", err.Error())
		return
	}
	containerApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ContainerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state ContainerModel
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
	if !plan.AutoRestartInterval.Equal(state.AutoRestartInterval) && !plan.AutoRestartInterval.IsUnknown() {
		body["auto-restart-interval"] = plan.AutoRestartInterval.ValueString()
	}
	if !plan.Cmd.Equal(state.Cmd) && !plan.Cmd.IsUnknown() {
		body["cmd"] = plan.Cmd.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.CpuList.Equal(state.CpuList) && !plan.CpuList.IsUnknown() {
		body["cpu-list"] = plan.CpuList.ValueString()
	}
	if !plan.Devices.Equal(state.Devices) && !plan.Devices.IsUnknown() {
		body["devices"] = plan.Devices.ValueString()
	}
	if !plan.DNS.Equal(state.DNS) && !plan.DNS.IsUnknown() {
		body["dns"] = plan.DNS.ValueString()
	}
	if !plan.DomainName.Equal(state.DomainName) && !plan.DomainName.IsUnknown() {
		body["domain-name"] = plan.DomainName.ValueString()
	}
	if !plan.Entrypoint.Equal(state.Entrypoint) && !plan.Entrypoint.IsUnknown() {
		body["entrypoint"] = plan.Entrypoint.ValueString()
	}
	if !plan.Envlist.Equal(state.Envlist) && !plan.Envlist.IsUnknown() {
		body["envlist"] = plan.Envlist.ValueString()
	}
	if !plan.Hostname.Equal(state.Hostname) && !plan.Hostname.IsUnknown() {
		body["hostname"] = plan.Hostname.ValueString()
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.Logging.Equal(state.Logging) && !plan.Logging.IsUnknown() {
		body["logging"] = plan.Logging.ValueString()
	}
	if !plan.MemoryHigh.Equal(state.MemoryHigh) && !plan.MemoryHigh.IsUnknown() {
		body["memory-high"] = plan.MemoryHigh.ValueString()
	}
	if !plan.MemoryMax.Equal(state.MemoryMax) && !plan.MemoryMax.IsUnknown() {
		body["memory-max"] = plan.MemoryMax.ValueString()
	}
	if !plan.Mount.Equal(state.Mount) && !plan.Mount.IsUnknown() {
		body["mount"] = plan.Mount.ValueString()
	}
	if !plan.Mountlists.Equal(state.Mountlists) && !plan.Mountlists.IsUnknown() {
		body["mountlists"] = plan.Mountlists.ValueString()
	}
	if !plan.RemoteImage.Equal(state.RemoteImage) && !plan.RemoteImage.IsUnknown() {
		body["remote-image"] = plan.RemoteImage.ValueString()
	}
	if !plan.RootDir.Equal(state.RootDir) && !plan.RootDir.IsUnknown() {
		body["root-dir"] = plan.RootDir.ValueString()
	}
	if !plan.StartOnBoot.Equal(state.StartOnBoot) && !plan.StartOnBoot.IsUnknown() {
		body["start-on-boot"] = plan.StartOnBoot.ValueString()
	}
	if !plan.StopSignal.Equal(state.StopSignal) && !plan.StopSignal.IsUnknown() {
		body["stop-signal"] = plan.StopSignal.ValueString()
	}
	if !plan.User.Equal(state.User) && !plan.User.IsUnknown() {
		body["user"] = plan.User.ValueString()
	}
	if !plan.Workdir.Equal(state.Workdir) && !plan.Workdir.IsUnknown() {
		body["workdir"] = plan.Workdir.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/container", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /container failed", err.Error())
			return
		}
		containerApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ContainerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state ContainerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/container", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /container failed", err.Error())
	}
}

func (r *ContainerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := containerLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /container matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// containerLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func containerLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/container", id)
}

func containerApply(ctx context.Context, obj client.Object, m *ContainerModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["auto-restart-interval"]; ok {
		if v != "" {
			m.AutoRestartInterval = types.StringValue(v)
		} else {
			m.AutoRestartInterval = types.StringNull()
		}
	}
	if v, ok := obj["cmd"]; ok {
		if v != "" {
			m.Cmd = types.StringValue(v)
		} else {
			m.Cmd = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["cpu-list"]; ok {
		if v != "" {
			m.CpuList = types.StringValue(v)
		} else {
			m.CpuList = types.StringNull()
		}
	}
	if v, ok := obj["devices"]; ok {
		if v != "" {
			m.Devices = types.StringValue(v)
		} else {
			m.Devices = types.StringNull()
		}
	}
	if v, ok := obj["dns"]; ok {
		if v != "" {
			m.DNS = types.StringValue(v)
		} else {
			m.DNS = types.StringNull()
		}
	}
	if v, ok := obj["domain-name"]; ok {
		if v != "" {
			m.DomainName = types.StringValue(v)
		} else {
			m.DomainName = types.StringNull()
		}
	}
	if v, ok := obj["entrypoint"]; ok {
		if v != "" {
			m.Entrypoint = types.StringValue(v)
		} else {
			m.Entrypoint = types.StringNull()
		}
	}
	if v, ok := obj["envlist"]; ok {
		if v != "" {
			m.Envlist = types.StringValue(v)
		} else {
			m.Envlist = types.StringNull()
		}
	}
	if v, ok := obj["hostname"]; ok {
		if v != "" {
			m.Hostname = types.StringValue(v)
		} else {
			m.Hostname = types.StringNull()
		}
	}
	if v, ok := obj["interface"]; ok {
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	}
	if v, ok := obj["logging"]; ok {
		if v != "" {
			m.Logging = types.StringValue(v)
		} else {
			m.Logging = types.StringNull()
		}
	}
	if v, ok := obj["memory-high"]; ok {
		if v != "" {
			m.MemoryHigh = types.StringValue(v)
		} else {
			m.MemoryHigh = types.StringNull()
		}
	}
	if v, ok := obj["memory-max"]; ok {
		if v != "" {
			m.MemoryMax = types.StringValue(v)
		} else {
			m.MemoryMax = types.StringNull()
		}
	}
	if v, ok := obj["mount"]; ok {
		if v != "" {
			m.Mount = types.StringValue(v)
		} else {
			m.Mount = types.StringNull()
		}
	}
	if v, ok := obj["mountlists"]; ok {
		if v != "" {
			m.Mountlists = types.StringValue(v)
		} else {
			m.Mountlists = types.StringNull()
		}
	}
	if v, ok := obj["remote-image"]; ok {
		if v != "" {
			m.RemoteImage = types.StringValue(v)
		} else {
			m.RemoteImage = types.StringNull()
		}
	}
	if v, ok := obj["root-dir"]; ok {
		if v != "" {
			m.RootDir = types.StringValue(v)
		} else {
			m.RootDir = types.StringNull()
		}
	}
	if v, ok := obj["start-on-boot"]; ok {
		if v != "" {
			m.StartOnBoot = types.StringValue(v)
		} else {
			m.StartOnBoot = types.StringNull()
		}
	}
	if v, ok := obj["stop-signal"]; ok {
		if v != "" {
			m.StopSignal = types.StringValue(v)
		} else {
			m.StopSignal = types.StringNull()
		}
	}
	if v, ok := obj["user"]; ok {
		if v != "" {
			m.User = types.StringValue(v)
		} else {
			m.User = types.StringNull()
		}
	}
	if v, ok := obj["workdir"]; ok {
		if v != "" {
			m.Workdir = types.StringValue(v)
		} else {
			m.Workdir = types.StringNull()
		}
	}
}
