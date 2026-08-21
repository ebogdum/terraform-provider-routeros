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
	_ resource.Resource                = &DiskResource{}
	_ resource.ResourceWithImportState = &DiskResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type DiskResource struct {
	reg *client.Registry
}

type DiskModel struct {
	ID                         types.String `tfsdk:"id"`
	Acquired                   types.Bool   `tfsdk:"acquired"`
	ActiveTime                 types.String `tfsdk:"active_time"`
	AvailableSpare             types.Int64  `tfsdk:"available_spare"`
	AvailableSpareThreshold    types.Int64  `tfsdk:"available_spare_threshold"`
	BlockDevice                types.Bool   `tfsdk:"block_device"`
	Btrfs                      types.String `tfsdk:"btrfs"`
	Comment                    types.String `tfsdk:"comment"`
	ControllerBurstTime        types.String `tfsdk:"controller_burst_time"`
	CriticalTemperature        types.Int64  `tfsdk:"critical_temperature"`
	CriticalTemperatureTime    types.String `tfsdk:"critical_temperature_time"`
	CriticalWarning            types.String `tfsdk:"critical_warning"`
	DefaultSlot                types.String `tfsdk:"default_slot"`
	Disabled                   types.Bool   `tfsdk:"disabled"`
	DiscardBytes               types.String `tfsdk:"discard_bytes"`
	DiscardMerges              types.String `tfsdk:"discard_merges"`
	DiscardOps                 types.String `tfsdk:"discard_ops"`
	DiscardTime                types.String `tfsdk:"discard_time"`
	EjectDrive                 types.String `tfsdk:"eject_drive"`
	Empty                      types.Bool   `tfsdk:"empty"`
	Encrypted                  types.Bool   `tfsdk:"encrypted"`
	FlushOps                   types.String `tfsdk:"flush_ops"`
	FlushTime                  types.String `tfsdk:"flush_time"`
	Formatting                 types.Bool   `tfsdk:"formatting"`
	Free                       types.String `tfsdk:"free"`
	Fs                         types.String `tfsdk:"fs"`
	FwVersion                  types.String `tfsdk:"fw_version"`
	GuidPartitionTable         types.Bool   `tfsdk:"guid_partition_table"`
	HostReadBytes              types.String `tfsdk:"host_read_bytes"`
	HostReadCommands           types.String `tfsdk:"host_read_commands"`
	HostWriteBytes             types.String `tfsdk:"host_write_bytes"`
	HostWriteCommands          types.String `tfsdk:"host_write_commands"`
	IScsiExport                types.Bool   `tfsdk:"i_scsi_export"`
	IScsiServerIqn             types.String `tfsdk:"i_scsi_server_iqn"`
	IScsiServerPort            types.Int64  `tfsdk:"i_scsi_server_port"`
	InFlightOps                types.String `tfsdk:"in_flight_ops"`
	Interface                  types.String `tfsdk:"interface"`
	InterfaceSpeed             types.String `tfsdk:"interface_speed"`
	IscsiSharing               types.String `tfsdk:"iscsi_sharing"`
	Label                      types.String `tfsdk:"label"`
	MediaInterface             types.String `tfsdk:"media_interface"`
	MediaSharing               types.Bool   `tfsdk:"media_sharing"`
	Model                      types.String `tfsdk:"model"`
	MountCompress              types.Bool   `tfsdk:"mount_compress"`
	MountFilesystem            types.Bool   `tfsdk:"mount_filesystem"`
	MountPoint                 types.String `tfsdk:"mount_point"`
	MountPointTemplate         types.String `tfsdk:"mount_point_template"`
	MountReadOnly              types.Bool   `tfsdk:"mount_read_only"`
	Mounted                    types.Bool   `tfsdk:"mounted"`
	Newfileman                 types.String `tfsdk:"newfileman"`
	NfsSharing                 types.Bool   `tfsdk:"nfs_sharing"`
	Nvme                       types.String `tfsdk:"nvme"`
	NvmeTCPExport              types.Bool   `tfsdk:"nvme_tcp_export"`
	NvmeTCPServerAllowHostName types.String `tfsdk:"nvme_tcp_server_allow_host_name"`
	NvmeTCPServerNqn           types.String `tfsdk:"nvme_tcp_server_nqn"`
	NvmeTCPServerPassword      types.String `tfsdk:"nvme_tcp_server_password"`
	NvmeTCPServerPort          types.Int64  `tfsdk:"nvme_tcp_server_port"`
	NvmeTCPServerSecret        types.String `tfsdk:"nvme_tcp_server_secret"`
	Oldfileman                 types.String `tfsdk:"oldfileman"`
	Parent                     types.String `tfsdk:"parent"`
	Part                       types.String `tfsdk:"part"`
	Partition                  types.Bool   `tfsdk:"partition"`
	PartitionNumber            types.Int64  `tfsdk:"partition_number"`
	PartitionOffset            types.String `tfsdk:"partition_offset"`
	PartitionSize              types.String `tfsdk:"partition_size"`
	PercentageUsed             types.Int64  `tfsdk:"percentage_used"`
	PowerCycles                types.Int64  `tfsdk:"power_cycles"`
	PowerOnTime                types.String `tfsdk:"power_on_time"`
	Raid                       types.String `tfsdk:"raid"`
	RaidAndMaster              types.String `tfsdk:"raid_and_master"`
	RaidAndType                types.String `tfsdk:"raid_and_type"`
	RaidMaster                 types.String `tfsdk:"raid_master"`
	RaidMember                 types.Bool   `tfsdk:"raid_member"`
	RaidMemberFailed           types.Bool   `tfsdk:"raid_member_failed"`
	RaidRole                   types.String `tfsdk:"raid_role"`
	RaidScrub                  types.String `tfsdk:"raid_scrub"`
	ReadBytes                  types.String `tfsdk:"read_bytes"`
	ReadMerges                 types.String `tfsdk:"read_merges"`
	ReadOnly                   types.Bool   `tfsdk:"read_only"`
	ReadOps                    types.String `tfsdk:"read_ops"`
	ReadOpsPerSecond           types.String `tfsdk:"read_ops_per_second"`
	ReadRate                   types.String `tfsdk:"read_rate"`
	ReadTime                   types.String `tfsdk:"read_time"`
	ResetCounters              types.String `tfsdk:"reset_counters"`
	Rose                       types.String `tfsdk:"rose"`
	Scan                       types.String `tfsdk:"scan"`
	SelfEncryptedAndLocked     types.Bool   `tfsdk:"self_encrypted_and_locked"`
	SelfEncryptionEnabled      types.Bool   `tfsdk:"self_encryption_enabled"`
	SelfEncryptionPassword     types.String `tfsdk:"self_encryption_password"`
	SelfEncryptionSupported    types.Bool   `tfsdk:"self_encryption_supported"`
	Serial                     types.String `tfsdk:"serial"`
	Size                       types.String `tfsdk:"size"`
	Slot                       types.String `tfsdk:"slot"`
	SmbServerEncryption        types.Bool   `tfsdk:"smb_server_encryption"`
	SmbServerPassword          types.String `tfsdk:"smb_server_password"`
	SmbServerUser              types.String `tfsdk:"smb_server_user"`
	SmbSharing                 types.Bool   `tfsdk:"smb_sharing"`
	State                      types.String `tfsdk:"state"`
	Swap                       types.Bool   `tfsdk:"swap"`
	SwapEnabled                types.Bool   `tfsdk:"swap_enabled"`
	Temperature                types.Int64  `tfsdk:"temperature"`
	Temperatures               types.String `tfsdk:"temperatures"`
	Tmpfs                      types.String `tfsdk:"tmpfs"`
	TmpfsMaxSize               types.String `tfsdk:"tmpfs_max_size"`
	Trim                       types.String `tfsdk:"trim"`
	Type                       types.String `tfsdk:"type"`
	UnrecoveredIntegrityErrors types.Int64  `tfsdk:"unrecovered_integrity_errors"`
	UnsafeShutdown             types.Int64  `tfsdk:"unsafe_shutdown"`
	Use                        types.Int64  `tfsdk:"use"`
	Uuid                       types.String `tfsdk:"uuid"`
	WaitTime                   types.String `tfsdk:"wait_time"`
	WarningTemperature         types.Int64  `tfsdk:"warning_temperature"`
	WarningTemperatureTime     types.String `tfsdk:"warning_temperature_time"`
	WriteBytes                 types.String `tfsdk:"write_bytes"`
	WriteMerges                types.String `tfsdk:"write_merges"`
	WriteOps                   types.String `tfsdk:"write_ops"`
	WriteOpsPerSecond          types.String `tfsdk:"write_ops_per_second"`
	WriteRate                  types.String `tfsdk:"write_rate"`
	WriteTime                  types.String `tfsdk:"write_time"`
	Router                     types.String `tfsdk:"router"`
}

func NewDiskResource() resource.Resource { return &DiskResource{} }

func (r *DiskResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_disk"
}

func (r *DiskResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *DiskResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Storage volumes. Creating one usually requires a backing device.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"acquired": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"active_time": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"available_spare": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"available_spare_threshold": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"block_device": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"btrfs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"controller_burst_time": schema.StringAttribute{
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"critical_temperature": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"critical_temperature_time": schema.StringAttribute{
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"critical_warning": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"default_slot": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"discard_bytes": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"discard_merges": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"discard_ops": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"discard_time": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"eject_drive": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"empty": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"encrypted": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"flush_ops": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"flush_time": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"formatting": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"free": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"fs": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"fat32", "ext4", "btrfs", "nfs", "smb", "wipe", "tmpfs", "exfat", "ntfs", "wipe-quck", "sshfs", "squashfs", "iso", "discard", "discard-secure", "xfs", "unknown"}...)},
			},
			"fw_version": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"guid_partition_table": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"host_read_bytes": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"host_read_commands": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"host_write_bytes": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"host_write_commands": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"i_scsi_export": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"i_scsi_server_iqn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"i_scsi_server_port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"in_flight_ops": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"interface": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"interface_speed": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"iscsi_sharing": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"label": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"media_interface": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"media_sharing": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"model": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"mount_compress": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mount_filesystem": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mount_point": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"mount_point_template": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Sets the mounting point for the file system. It is possible to set the mount point as the following parameters based on the disk: [slot] (default) - sets the mount point as the slot name. [model] - sets the mount point as the device's model name. [serial] - sets the mount point as the device serial [fw-version] - sets the mount point as the device's firmware version. [fs-label] - sets the mount point as the device's file system label. [fs-uuid] - sets the mount point as the device's UUID [fs] - sets the mount point as the device's file system ros Additionally, it is possible to combine multiple variables to create a single mount point: ros",
			},
			"mount_read_only": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Sets the mounted disk in read only mode when set to yes .",
			},
			"mounted": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"newfileman": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nfs_sharing": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nvme": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nvme_tcp_export": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nvme_tcp_server_allow_host_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nvme_tcp_server_nqn": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nvme_tcp_server_password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "",
			},
			"nvme_tcp_server_port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nvme_tcp_server_secret": schema.StringAttribute{
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"oldfileman": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"parent": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"part": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"partition": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"partition_number": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"partition_offset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"partition_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"percentage_used": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"power_cycles": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"power_on_time": schema.StringAttribute{
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"raid": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"raid_and_master": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"raid_and_type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"raid_master": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"raid_member": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"raid_member_failed": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"raid_role": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"spare"}...)},
			},
			"raid_scrub": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"read_bytes": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"read_merges": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"read_only": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"read_ops": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"read_ops_per_second": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"read_rate": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"read_time": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"reset_counters": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rose": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"scan": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"self_encrypted_and_locked": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"self_encryption_enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"self_encryption_password": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "",
			},
			"self_encryption_supported": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"serial": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"size": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"slot": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"smb_server_encryption": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"smb_server_password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"smb_server_user": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"smb_sharing": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"state": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"swap": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"swap_enabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"temperature": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"temperatures": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"tmpfs": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tmpfs_max_size": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"trim": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"unrecovered_integrity_errors": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"unsafe_shutdown": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"use": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"uuid": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"wait_time": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"warning_temperature": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"warning_temperature_time": schema.StringAttribute{
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"write_bytes": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"write_merges": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"write_ops": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"write_ops_per_second": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"write_rate": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"write_time": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *DiskResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan DiskModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.BlockDevice.IsNull() || plan.BlockDevice.IsUnknown()) {
		body["block-device"] = client.FormatBool(plan.BlockDevice.ValueBool())
	}
	if !(plan.Btrfs.IsNull() || plan.Btrfs.IsUnknown()) {
		body["btrfs"] = plan.Btrfs.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.EjectDrive.IsNull() || plan.EjectDrive.IsUnknown()) {
		body["eject-drive"] = plan.EjectDrive.ValueString()
	}
	if !(plan.Empty.IsNull() || plan.Empty.IsUnknown()) {
		body["empty"] = client.FormatBool(plan.Empty.ValueBool())
	}
	if !(plan.Encrypted.IsNull() || plan.Encrypted.IsUnknown()) {
		body["encrypted"] = client.FormatBool(plan.Encrypted.ValueBool())
	}
	if !(plan.Formatting.IsNull() || plan.Formatting.IsUnknown()) {
		body["formatting"] = client.FormatBool(plan.Formatting.ValueBool())
	}
	if !(plan.GuidPartitionTable.IsNull() || plan.GuidPartitionTable.IsUnknown()) {
		body["guid-partition-table"] = client.FormatBool(plan.GuidPartitionTable.ValueBool())
	}
	if !(plan.IScsiExport.IsNull() || plan.IScsiExport.IsUnknown()) {
		body["i-scsi-export"] = client.FormatBool(plan.IScsiExport.ValueBool())
	}
	if !(plan.IScsiServerIqn.IsNull() || plan.IScsiServerIqn.IsUnknown()) {
		body["i-scsi-server-iqn"] = plan.IScsiServerIqn.ValueString()
	}
	if !(plan.IScsiServerPort.IsNull() || plan.IScsiServerPort.IsUnknown()) {
		body["i-scsi-server-port"] = client.FormatInt64(plan.IScsiServerPort.ValueInt64())
	}
	if !(plan.IscsiSharing.IsNull() || plan.IscsiSharing.IsUnknown()) {
		body["iscsi-sharing"] = plan.IscsiSharing.ValueString()
	}
	if !(plan.MediaInterface.IsNull() || plan.MediaInterface.IsUnknown()) {
		body["media-interface"] = plan.MediaInterface.ValueString()
	}
	if !(plan.MediaSharing.IsNull() || plan.MediaSharing.IsUnknown()) {
		body["media-sharing"] = client.FormatBool(plan.MediaSharing.ValueBool())
	}
	if !(plan.MountCompress.IsNull() || plan.MountCompress.IsUnknown()) {
		body["mount-compress"] = client.FormatBool(plan.MountCompress.ValueBool())
	}
	if !(plan.MountFilesystem.IsNull() || plan.MountFilesystem.IsUnknown()) {
		body["mount-filesystem"] = client.FormatBool(plan.MountFilesystem.ValueBool())
	}
	if !(plan.MountPointTemplate.IsNull() || plan.MountPointTemplate.IsUnknown()) {
		body["mount-point-template"] = plan.MountPointTemplate.ValueString()
	}
	if !(plan.MountReadOnly.IsNull() || plan.MountReadOnly.IsUnknown()) {
		body["mount-read-only"] = client.FormatBool(plan.MountReadOnly.ValueBool())
	}
	if !(plan.Mounted.IsNull() || plan.Mounted.IsUnknown()) {
		body["mounted"] = client.FormatBool(plan.Mounted.ValueBool())
	}
	if !(plan.Newfileman.IsNull() || plan.Newfileman.IsUnknown()) {
		body["newfileman"] = plan.Newfileman.ValueString()
	}
	if !(plan.NfsSharing.IsNull() || plan.NfsSharing.IsUnknown()) {
		body["nfs-sharing"] = client.FormatBool(plan.NfsSharing.ValueBool())
	}
	if !(plan.Nvme.IsNull() || plan.Nvme.IsUnknown()) {
		body["nvme"] = plan.Nvme.ValueString()
	}
	if !(plan.NvmeTCPExport.IsNull() || plan.NvmeTCPExport.IsUnknown()) {
		body["nvme-tcp-export"] = client.FormatBool(plan.NvmeTCPExport.ValueBool())
	}
	if !(plan.NvmeTCPServerAllowHostName.IsNull() || plan.NvmeTCPServerAllowHostName.IsUnknown()) {
		body["nvme-tcp-server-allow-host-name"] = plan.NvmeTCPServerAllowHostName.ValueString()
	}
	if !(plan.NvmeTCPServerNqn.IsNull() || plan.NvmeTCPServerNqn.IsUnknown()) {
		body["nvme-tcp-server-nqn"] = plan.NvmeTCPServerNqn.ValueString()
	}
	if !(plan.NvmeTCPServerPassword.IsNull() || plan.NvmeTCPServerPassword.IsUnknown()) {
		body["nvme-tcp-server-password"] = plan.NvmeTCPServerPassword.ValueString()
	}
	if !(plan.NvmeTCPServerPort.IsNull() || plan.NvmeTCPServerPort.IsUnknown()) {
		body["nvme-tcp-server-port"] = client.FormatInt64(plan.NvmeTCPServerPort.ValueInt64())
	}
	if !(plan.Oldfileman.IsNull() || plan.Oldfileman.IsUnknown()) {
		body["oldfileman"] = plan.Oldfileman.ValueString()
	}
	if !(plan.Parent.IsNull() || plan.Parent.IsUnknown()) {
		body["parent"] = plan.Parent.ValueString()
	}
	if !(plan.Part.IsNull() || plan.Part.IsUnknown()) {
		body["part"] = plan.Part.ValueString()
	}
	if !(plan.Partition.IsNull() || plan.Partition.IsUnknown()) {
		body["partition"] = client.FormatBool(plan.Partition.ValueBool())
	}
	if !(plan.PartitionOffset.IsNull() || plan.PartitionOffset.IsUnknown()) {
		body["partition-offset"] = plan.PartitionOffset.ValueString()
	}
	if !(plan.PartitionSize.IsNull() || plan.PartitionSize.IsUnknown()) {
		body["partition-size"] = plan.PartitionSize.ValueString()
	}
	if !(plan.Raid.IsNull() || plan.Raid.IsUnknown()) {
		body["raid"] = plan.Raid.ValueString()
	}
	if !(plan.RaidAndMaster.IsNull() || plan.RaidAndMaster.IsUnknown()) {
		body["raid-and-master"] = plan.RaidAndMaster.ValueString()
	}
	if !(plan.RaidAndType.IsNull() || plan.RaidAndType.IsUnknown()) {
		body["raid-and-type"] = plan.RaidAndType.ValueString()
	}
	if !(plan.RaidMaster.IsNull() || plan.RaidMaster.IsUnknown()) {
		body["raid-master"] = plan.RaidMaster.ValueString()
	}
	if !(plan.RaidMember.IsNull() || plan.RaidMember.IsUnknown()) {
		body["raid-member"] = client.FormatBool(plan.RaidMember.ValueBool())
	}
	if !(plan.RaidMemberFailed.IsNull() || plan.RaidMemberFailed.IsUnknown()) {
		body["raid-member-failed"] = client.FormatBool(plan.RaidMemberFailed.ValueBool())
	}
	if !(plan.RaidRole.IsNull() || plan.RaidRole.IsUnknown()) {
		body["raid-role"] = plan.RaidRole.ValueString()
	}
	if !(plan.RaidScrub.IsNull() || plan.RaidScrub.IsUnknown()) {
		body["raid-scrub"] = plan.RaidScrub.ValueString()
	}
	if !(plan.ReadOnly.IsNull() || plan.ReadOnly.IsUnknown()) {
		body["read-only"] = client.FormatBool(plan.ReadOnly.ValueBool())
	}
	if !(plan.ResetCounters.IsNull() || plan.ResetCounters.IsUnknown()) {
		body["reset-counters"] = plan.ResetCounters.ValueString()
	}
	if !(plan.Rose.IsNull() || plan.Rose.IsUnknown()) {
		body["rose"] = plan.Rose.ValueString()
	}
	if !(plan.Scan.IsNull() || plan.Scan.IsUnknown()) {
		body["scan"] = plan.Scan.ValueString()
	}
	if !(plan.SelfEncryptedAndLocked.IsNull() || plan.SelfEncryptedAndLocked.IsUnknown()) {
		body["self-encrypted-and-locked"] = client.FormatBool(plan.SelfEncryptedAndLocked.ValueBool())
	}
	if !(plan.SelfEncryptionEnabled.IsNull() || plan.SelfEncryptionEnabled.IsUnknown()) {
		body["self-encryption-enabled"] = client.FormatBool(plan.SelfEncryptionEnabled.ValueBool())
	}
	if !(plan.SelfEncryptionPassword.IsNull() || plan.SelfEncryptionPassword.IsUnknown()) {
		body["self-encryption-password"] = plan.SelfEncryptionPassword.ValueString()
	}
	if !(plan.SelfEncryptionSupported.IsNull() || plan.SelfEncryptionSupported.IsUnknown()) {
		body["self-encryption-supported"] = client.FormatBool(plan.SelfEncryptionSupported.ValueBool())
	}
	if !(plan.Slot.IsNull() || plan.Slot.IsUnknown()) {
		body["slot"] = plan.Slot.ValueString()
	}
	if !(plan.SmbServerEncryption.IsNull() || plan.SmbServerEncryption.IsUnknown()) {
		body["smb-server-encryption"] = client.FormatBool(plan.SmbServerEncryption.ValueBool())
	}
	if !(plan.SmbServerPassword.IsNull() || plan.SmbServerPassword.IsUnknown()) {
		body["smb-server-password"] = plan.SmbServerPassword.ValueString()
	}
	if !(plan.SmbServerUser.IsNull() || plan.SmbServerUser.IsUnknown()) {
		body["smb-server-user"] = plan.SmbServerUser.ValueString()
	}
	if !(plan.SmbSharing.IsNull() || plan.SmbSharing.IsUnknown()) {
		body["smb-sharing"] = client.FormatBool(plan.SmbSharing.ValueBool())
	}
	if !(plan.Swap.IsNull() || plan.Swap.IsUnknown()) {
		body["swap"] = client.FormatBool(plan.Swap.ValueBool())
	}
	if !(plan.SwapEnabled.IsNull() || plan.SwapEnabled.IsUnknown()) {
		body["swap-enabled"] = client.FormatBool(plan.SwapEnabled.ValueBool())
	}
	if !(plan.Tmpfs.IsNull() || plan.Tmpfs.IsUnknown()) {
		body["tmpfs"] = plan.Tmpfs.ValueString()
	}
	if !(plan.TmpfsMaxSize.IsNull() || plan.TmpfsMaxSize.IsUnknown()) {
		body["tmpfs-max-size"] = plan.TmpfsMaxSize.ValueString()
	}
	if !(plan.Trim.IsNull() || plan.Trim.IsUnknown()) {
		body["trim"] = plan.Trim.ValueString()
	}
	if !(plan.Type.IsNull() || plan.Type.IsUnknown()) {
		body["type"] = plan.Type.ValueString()
	}
	obj, err := c.Add(ctx, "/disk", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /disk failed", err.Error())
		return
	}
	diskApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *DiskResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state DiskModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/disk", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /disk failed", err.Error())
		return
	}
	diskApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *DiskResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state DiskModel
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
	if !plan.BlockDevice.Equal(state.BlockDevice) && !plan.BlockDevice.IsUnknown() {
		body["block-device"] = client.FormatBool(plan.BlockDevice.ValueBool())
	}
	if !plan.Btrfs.Equal(state.Btrfs) && !plan.Btrfs.IsUnknown() {
		body["btrfs"] = plan.Btrfs.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.EjectDrive.Equal(state.EjectDrive) && !plan.EjectDrive.IsUnknown() {
		body["eject-drive"] = plan.EjectDrive.ValueString()
	}
	if !plan.Empty.Equal(state.Empty) && !plan.Empty.IsUnknown() {
		body["empty"] = client.FormatBool(plan.Empty.ValueBool())
	}
	if !plan.Encrypted.Equal(state.Encrypted) && !plan.Encrypted.IsUnknown() {
		body["encrypted"] = client.FormatBool(plan.Encrypted.ValueBool())
	}
	if !plan.Formatting.Equal(state.Formatting) && !plan.Formatting.IsUnknown() {
		body["formatting"] = client.FormatBool(plan.Formatting.ValueBool())
	}
	if !plan.GuidPartitionTable.Equal(state.GuidPartitionTable) && !plan.GuidPartitionTable.IsUnknown() {
		body["guid-partition-table"] = client.FormatBool(plan.GuidPartitionTable.ValueBool())
	}
	if !plan.IScsiExport.Equal(state.IScsiExport) && !plan.IScsiExport.IsUnknown() {
		body["i-scsi-export"] = client.FormatBool(plan.IScsiExport.ValueBool())
	}
	if !plan.IScsiServerIqn.Equal(state.IScsiServerIqn) && !plan.IScsiServerIqn.IsUnknown() {
		body["i-scsi-server-iqn"] = plan.IScsiServerIqn.ValueString()
	}
	if !plan.IScsiServerPort.Equal(state.IScsiServerPort) && !plan.IScsiServerPort.IsUnknown() {
		body["i-scsi-server-port"] = client.FormatInt64(plan.IScsiServerPort.ValueInt64())
	}
	if !plan.IscsiSharing.Equal(state.IscsiSharing) && !plan.IscsiSharing.IsUnknown() {
		body["iscsi-sharing"] = plan.IscsiSharing.ValueString()
	}
	if !plan.MediaInterface.Equal(state.MediaInterface) && !plan.MediaInterface.IsUnknown() {
		body["media-interface"] = plan.MediaInterface.ValueString()
	}
	if !plan.MediaSharing.Equal(state.MediaSharing) && !plan.MediaSharing.IsUnknown() {
		body["media-sharing"] = client.FormatBool(plan.MediaSharing.ValueBool())
	}
	if !plan.MountCompress.Equal(state.MountCompress) && !plan.MountCompress.IsUnknown() {
		body["mount-compress"] = client.FormatBool(plan.MountCompress.ValueBool())
	}
	if !plan.MountFilesystem.Equal(state.MountFilesystem) && !plan.MountFilesystem.IsUnknown() {
		body["mount-filesystem"] = client.FormatBool(plan.MountFilesystem.ValueBool())
	}
	if !plan.MountPointTemplate.Equal(state.MountPointTemplate) && !plan.MountPointTemplate.IsUnknown() {
		body["mount-point-template"] = plan.MountPointTemplate.ValueString()
	}
	if !plan.MountReadOnly.Equal(state.MountReadOnly) && !plan.MountReadOnly.IsUnknown() {
		body["mount-read-only"] = client.FormatBool(plan.MountReadOnly.ValueBool())
	}
	if !plan.Mounted.Equal(state.Mounted) && !plan.Mounted.IsUnknown() {
		body["mounted"] = client.FormatBool(plan.Mounted.ValueBool())
	}
	if !plan.Newfileman.Equal(state.Newfileman) && !plan.Newfileman.IsUnknown() {
		body["newfileman"] = plan.Newfileman.ValueString()
	}
	if !plan.NfsSharing.Equal(state.NfsSharing) && !plan.NfsSharing.IsUnknown() {
		body["nfs-sharing"] = client.FormatBool(plan.NfsSharing.ValueBool())
	}
	if !plan.Nvme.Equal(state.Nvme) && !plan.Nvme.IsUnknown() {
		body["nvme"] = plan.Nvme.ValueString()
	}
	if !plan.NvmeTCPExport.Equal(state.NvmeTCPExport) && !plan.NvmeTCPExport.IsUnknown() {
		body["nvme-tcp-export"] = client.FormatBool(plan.NvmeTCPExport.ValueBool())
	}
	if !plan.NvmeTCPServerAllowHostName.Equal(state.NvmeTCPServerAllowHostName) && !plan.NvmeTCPServerAllowHostName.IsUnknown() {
		body["nvme-tcp-server-allow-host-name"] = plan.NvmeTCPServerAllowHostName.ValueString()
	}
	if !plan.NvmeTCPServerNqn.Equal(state.NvmeTCPServerNqn) && !plan.NvmeTCPServerNqn.IsUnknown() {
		body["nvme-tcp-server-nqn"] = plan.NvmeTCPServerNqn.ValueString()
	}
	if !plan.NvmeTCPServerPassword.Equal(state.NvmeTCPServerPassword) && !plan.NvmeTCPServerPassword.IsUnknown() {
		body["nvme-tcp-server-password"] = plan.NvmeTCPServerPassword.ValueString()
	}
	if !plan.NvmeTCPServerPort.Equal(state.NvmeTCPServerPort) && !plan.NvmeTCPServerPort.IsUnknown() {
		body["nvme-tcp-server-port"] = client.FormatInt64(plan.NvmeTCPServerPort.ValueInt64())
	}
	if !plan.Oldfileman.Equal(state.Oldfileman) && !plan.Oldfileman.IsUnknown() {
		body["oldfileman"] = plan.Oldfileman.ValueString()
	}
	if !plan.Parent.Equal(state.Parent) && !plan.Parent.IsUnknown() {
		body["parent"] = plan.Parent.ValueString()
	}
	if !plan.Part.Equal(state.Part) && !plan.Part.IsUnknown() {
		body["part"] = plan.Part.ValueString()
	}
	if !plan.Partition.Equal(state.Partition) && !plan.Partition.IsUnknown() {
		body["partition"] = client.FormatBool(plan.Partition.ValueBool())
	}
	if !plan.PartitionOffset.Equal(state.PartitionOffset) && !plan.PartitionOffset.IsUnknown() {
		body["partition-offset"] = plan.PartitionOffset.ValueString()
	}
	if !plan.PartitionSize.Equal(state.PartitionSize) && !plan.PartitionSize.IsUnknown() {
		body["partition-size"] = plan.PartitionSize.ValueString()
	}
	if !plan.Raid.Equal(state.Raid) && !plan.Raid.IsUnknown() {
		body["raid"] = plan.Raid.ValueString()
	}
	if !plan.RaidAndMaster.Equal(state.RaidAndMaster) && !plan.RaidAndMaster.IsUnknown() {
		body["raid-and-master"] = plan.RaidAndMaster.ValueString()
	}
	if !plan.RaidAndType.Equal(state.RaidAndType) && !plan.RaidAndType.IsUnknown() {
		body["raid-and-type"] = plan.RaidAndType.ValueString()
	}
	if !plan.RaidMaster.Equal(state.RaidMaster) && !plan.RaidMaster.IsUnknown() {
		body["raid-master"] = plan.RaidMaster.ValueString()
	}
	if !plan.RaidMember.Equal(state.RaidMember) && !plan.RaidMember.IsUnknown() {
		body["raid-member"] = client.FormatBool(plan.RaidMember.ValueBool())
	}
	if !plan.RaidMemberFailed.Equal(state.RaidMemberFailed) && !plan.RaidMemberFailed.IsUnknown() {
		body["raid-member-failed"] = client.FormatBool(plan.RaidMemberFailed.ValueBool())
	}
	if !plan.RaidRole.Equal(state.RaidRole) && !plan.RaidRole.IsUnknown() {
		body["raid-role"] = plan.RaidRole.ValueString()
	}
	if !plan.RaidScrub.Equal(state.RaidScrub) && !plan.RaidScrub.IsUnknown() {
		body["raid-scrub"] = plan.RaidScrub.ValueString()
	}
	if !plan.ReadOnly.Equal(state.ReadOnly) && !plan.ReadOnly.IsUnknown() {
		body["read-only"] = client.FormatBool(plan.ReadOnly.ValueBool())
	}
	if !plan.ResetCounters.Equal(state.ResetCounters) && !plan.ResetCounters.IsUnknown() {
		body["reset-counters"] = plan.ResetCounters.ValueString()
	}
	if !plan.Rose.Equal(state.Rose) && !plan.Rose.IsUnknown() {
		body["rose"] = plan.Rose.ValueString()
	}
	if !plan.Scan.Equal(state.Scan) && !plan.Scan.IsUnknown() {
		body["scan"] = plan.Scan.ValueString()
	}
	if !plan.SelfEncryptedAndLocked.Equal(state.SelfEncryptedAndLocked) && !plan.SelfEncryptedAndLocked.IsUnknown() {
		body["self-encrypted-and-locked"] = client.FormatBool(plan.SelfEncryptedAndLocked.ValueBool())
	}
	if !plan.SelfEncryptionEnabled.Equal(state.SelfEncryptionEnabled) && !plan.SelfEncryptionEnabled.IsUnknown() {
		body["self-encryption-enabled"] = client.FormatBool(plan.SelfEncryptionEnabled.ValueBool())
	}
	if !plan.SelfEncryptionPassword.Equal(state.SelfEncryptionPassword) && !plan.SelfEncryptionPassword.IsUnknown() {
		body["self-encryption-password"] = plan.SelfEncryptionPassword.ValueString()
	}
	if !plan.SelfEncryptionSupported.Equal(state.SelfEncryptionSupported) && !plan.SelfEncryptionSupported.IsUnknown() {
		body["self-encryption-supported"] = client.FormatBool(plan.SelfEncryptionSupported.ValueBool())
	}
	if !plan.Slot.Equal(state.Slot) && !plan.Slot.IsUnknown() {
		body["slot"] = plan.Slot.ValueString()
	}
	if !plan.SmbServerEncryption.Equal(state.SmbServerEncryption) && !plan.SmbServerEncryption.IsUnknown() {
		body["smb-server-encryption"] = client.FormatBool(plan.SmbServerEncryption.ValueBool())
	}
	if !plan.SmbServerPassword.Equal(state.SmbServerPassword) && !plan.SmbServerPassword.IsUnknown() {
		body["smb-server-password"] = plan.SmbServerPassword.ValueString()
	}
	if !plan.SmbServerUser.Equal(state.SmbServerUser) && !plan.SmbServerUser.IsUnknown() {
		body["smb-server-user"] = plan.SmbServerUser.ValueString()
	}
	if !plan.SmbSharing.Equal(state.SmbSharing) && !plan.SmbSharing.IsUnknown() {
		body["smb-sharing"] = client.FormatBool(plan.SmbSharing.ValueBool())
	}
	if !plan.Swap.Equal(state.Swap) && !plan.Swap.IsUnknown() {
		body["swap"] = client.FormatBool(plan.Swap.ValueBool())
	}
	if !plan.SwapEnabled.Equal(state.SwapEnabled) && !plan.SwapEnabled.IsUnknown() {
		body["swap-enabled"] = client.FormatBool(plan.SwapEnabled.ValueBool())
	}
	if !plan.Tmpfs.Equal(state.Tmpfs) && !plan.Tmpfs.IsUnknown() {
		body["tmpfs"] = plan.Tmpfs.ValueString()
	}
	if !plan.TmpfsMaxSize.Equal(state.TmpfsMaxSize) && !plan.TmpfsMaxSize.IsUnknown() {
		body["tmpfs-max-size"] = plan.TmpfsMaxSize.ValueString()
	}
	if !plan.Trim.Equal(state.Trim) && !plan.Trim.IsUnknown() {
		body["trim"] = plan.Trim.ValueString()
	}
	if !plan.Type.Equal(state.Type) && !plan.Type.IsUnknown() {
		body["type"] = plan.Type.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/disk", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /disk failed", err.Error())
			return
		}
		diskApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *DiskResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state DiskModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/disk", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /disk failed", err.Error())
	}
}

func (r *DiskResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := diskLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /disk matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// diskLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func diskLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/disk", id)
}

func diskApply(ctx context.Context, obj client.Object, m *DiskModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["acquired"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Acquired = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Acquired = types.BoolValue(true)
		} else {
			m.Acquired = types.BoolNull()
		}
	}
	if v, ok := obj["active-time"]; ok {
		if v != "" {
			m.ActiveTime = types.StringValue(v)
		} else {
			m.ActiveTime = types.StringNull()
		}
	}
	if v, ok := obj["available-spare"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.AvailableSpare = types.Int64Value(n)
		} else {
			m.AvailableSpare = types.Int64Null()
		}
	} else {
		m.AvailableSpare = types.Int64Null()
	}
	if v, ok := obj["available-spare-threshold"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.AvailableSpareThreshold = types.Int64Value(n)
		} else {
			m.AvailableSpareThreshold = types.Int64Null()
		}
	} else {
		m.AvailableSpareThreshold = types.Int64Null()
	}
	if v, ok := obj["block-device"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.BlockDevice = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.BlockDevice = types.BoolValue(true)
		} else {
			m.BlockDevice = types.BoolNull()
		}
	}
	if v, ok := obj["btrfs"]; ok {
		if v != "" {
			m.Btrfs = types.StringValue(v)
		} else {
			m.Btrfs = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["controller-burst-time"]; ok {
		if v != "" {
			m.ControllerBurstTime = types.StringValue(v)
		} else {
			m.ControllerBurstTime = types.StringNull()
		}
	}
	if v, ok := obj["critical-temperature"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.CriticalTemperature = types.Int64Value(n)
		} else {
			m.CriticalTemperature = types.Int64Null()
		}
	} else {
		m.CriticalTemperature = types.Int64Null()
	}
	if v, ok := obj["critical-temperature-time"]; ok {
		if v != "" {
			m.CriticalTemperatureTime = types.StringValue(v)
		} else {
			m.CriticalTemperatureTime = types.StringNull()
		}
	}
	if v, ok := obj["critical-warning"]; ok {
		if v != "" {
			m.CriticalWarning = types.StringValue(v)
		} else {
			m.CriticalWarning = types.StringNull()
		}
	}
	if v, ok := obj["default-slot"]; ok {
		if v != "" {
			m.DefaultSlot = types.StringValue(v)
		} else {
			m.DefaultSlot = types.StringNull()
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
	if v, ok := obj["discard-bytes"]; ok {
		if v != "" {
			m.DiscardBytes = types.StringValue(v)
		} else {
			m.DiscardBytes = types.StringNull()
		}
	}
	if v, ok := obj["discard-merges"]; ok {
		if v != "" {
			m.DiscardMerges = types.StringValue(v)
		} else {
			m.DiscardMerges = types.StringNull()
		}
	}
	if v, ok := obj["discard-ops"]; ok {
		if v != "" {
			m.DiscardOps = types.StringValue(v)
		} else {
			m.DiscardOps = types.StringNull()
		}
	}
	if v, ok := obj["discard-time"]; ok {
		if v != "" {
			m.DiscardTime = types.StringValue(v)
		} else {
			m.DiscardTime = types.StringNull()
		}
	}
	if v, ok := obj["eject-drive"]; ok {
		if v != "" {
			m.EjectDrive = types.StringValue(v)
		} else {
			m.EjectDrive = types.StringNull()
		}
	}
	if v, ok := obj["empty"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Empty = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Empty = types.BoolValue(true)
		} else {
			m.Empty = types.BoolNull()
		}
	}
	if v, ok := obj["encrypted"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Encrypted = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Encrypted = types.BoolValue(true)
		} else {
			m.Encrypted = types.BoolNull()
		}
	}
	if v, ok := obj["flush-ops"]; ok {
		if v != "" {
			m.FlushOps = types.StringValue(v)
		} else {
			m.FlushOps = types.StringNull()
		}
	}
	if v, ok := obj["flush-time"]; ok {
		if v != "" {
			m.FlushTime = types.StringValue(v)
		} else {
			m.FlushTime = types.StringNull()
		}
	}
	if v, ok := obj["formatting"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Formatting = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Formatting = types.BoolValue(true)
		} else {
			m.Formatting = types.BoolNull()
		}
	}
	if v, ok := obj["free"]; ok {
		if v != "" {
			m.Free = types.StringValue(v)
		} else {
			m.Free = types.StringNull()
		}
	}
	if v, ok := obj["fs"]; ok {
		if v != "" {
			m.Fs = types.StringValue(v)
		} else {
			m.Fs = types.StringNull()
		}
	}
	if v, ok := obj["fw-version"]; ok {
		if v != "" {
			m.FwVersion = types.StringValue(v)
		} else {
			m.FwVersion = types.StringNull()
		}
	}
	if v, ok := obj["guid-partition-table"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.GuidPartitionTable = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.GuidPartitionTable = types.BoolValue(true)
		} else {
			m.GuidPartitionTable = types.BoolNull()
		}
	}
	if v, ok := obj["host-read-bytes"]; ok {
		if v != "" {
			m.HostReadBytes = types.StringValue(v)
		} else {
			m.HostReadBytes = types.StringNull()
		}
	}
	if v, ok := obj["host-read-commands"]; ok {
		if v != "" {
			m.HostReadCommands = types.StringValue(v)
		} else {
			m.HostReadCommands = types.StringNull()
		}
	}
	if v, ok := obj["host-write-bytes"]; ok {
		if v != "" {
			m.HostWriteBytes = types.StringValue(v)
		} else {
			m.HostWriteBytes = types.StringNull()
		}
	}
	if v, ok := obj["host-write-commands"]; ok {
		if v != "" {
			m.HostWriteCommands = types.StringValue(v)
		} else {
			m.HostWriteCommands = types.StringNull()
		}
	}
	if v, ok := obj["i-scsi-export"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.IScsiExport = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.IScsiExport = types.BoolValue(true)
		} else {
			m.IScsiExport = types.BoolNull()
		}
	}
	if v, ok := obj["i-scsi-server-iqn"]; ok {
		if v != "" {
			m.IScsiServerIqn = types.StringValue(v)
		} else {
			m.IScsiServerIqn = types.StringNull()
		}
	}
	if v, ok := obj["i-scsi-server-port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.IScsiServerPort = types.Int64Value(n)
		} else {
			m.IScsiServerPort = types.Int64Null()
		}
	} else {
		m.IScsiServerPort = types.Int64Null()
	}
	if v, ok := obj["in-flight-ops"]; ok {
		if v != "" {
			m.InFlightOps = types.StringValue(v)
		} else {
			m.InFlightOps = types.StringNull()
		}
	}
	if v, ok := obj["interface"]; ok {
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	}
	if v, ok := obj["interface-speed"]; ok {
		if v != "" {
			m.InterfaceSpeed = types.StringValue(v)
		} else {
			m.InterfaceSpeed = types.StringNull()
		}
	}
	if v, ok := obj["iscsi-sharing"]; ok {
		if v != "" {
			m.IscsiSharing = types.StringValue(v)
		} else {
			m.IscsiSharing = types.StringNull()
		}
	}
	if v, ok := obj["label"]; ok {
		if v != "" {
			m.Label = types.StringValue(v)
		} else {
			m.Label = types.StringNull()
		}
	}
	if v, ok := obj["media-interface"]; ok {
		if v != "" {
			m.MediaInterface = types.StringValue(v)
		} else {
			m.MediaInterface = types.StringNull()
		}
	}
	if v, ok := obj["media-sharing"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.MediaSharing = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.MediaSharing = types.BoolValue(true)
		} else {
			m.MediaSharing = types.BoolNull()
		}
	}
	if v, ok := obj["model"]; ok {
		if v != "" {
			m.Model = types.StringValue(v)
		} else {
			m.Model = types.StringNull()
		}
	}
	if v, ok := obj["mount-compress"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.MountCompress = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.MountCompress = types.BoolValue(true)
		} else {
			m.MountCompress = types.BoolNull()
		}
	}
	if v, ok := obj["mount-filesystem"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.MountFilesystem = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.MountFilesystem = types.BoolValue(true)
		} else {
			m.MountFilesystem = types.BoolNull()
		}
	}
	if v, ok := obj["mount-point"]; ok {
		if v != "" {
			m.MountPoint = types.StringValue(v)
		} else {
			m.MountPoint = types.StringNull()
		}
	}
	if v, ok := obj["mount-point-template"]; ok {
		if v != "" {
			m.MountPointTemplate = types.StringValue(v)
		} else {
			m.MountPointTemplate = types.StringNull()
		}
	}
	if v, ok := obj["mount-read-only"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.MountReadOnly = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.MountReadOnly = types.BoolValue(true)
		} else {
			m.MountReadOnly = types.BoolNull()
		}
	}
	if v, ok := obj["mounted"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Mounted = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Mounted = types.BoolValue(true)
		} else {
			m.Mounted = types.BoolNull()
		}
	}
	if v, ok := obj["newfileman"]; ok {
		if v != "" {
			m.Newfileman = types.StringValue(v)
		} else {
			m.Newfileman = types.StringNull()
		}
	}
	if v, ok := obj["nfs-sharing"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.NfsSharing = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.NfsSharing = types.BoolValue(true)
		} else {
			m.NfsSharing = types.BoolNull()
		}
	}
	if v, ok := obj["nvme"]; ok {
		if v != "" {
			m.Nvme = types.StringValue(v)
		} else {
			m.Nvme = types.StringNull()
		}
	}
	if v, ok := obj["nvme-tcp-export"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.NvmeTCPExport = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.NvmeTCPExport = types.BoolValue(true)
		} else {
			m.NvmeTCPExport = types.BoolNull()
		}
	}
	if v, ok := obj["nvme-tcp-server-allow-host-name"]; ok {
		if v != "" {
			m.NvmeTCPServerAllowHostName = types.StringValue(v)
		} else {
			m.NvmeTCPServerAllowHostName = types.StringNull()
		}
	}
	if v, ok := obj["nvme-tcp-server-nqn"]; ok {
		if v != "" {
			m.NvmeTCPServerNqn = types.StringValue(v)
		} else {
			m.NvmeTCPServerNqn = types.StringNull()
		}
	}
	if v, ok := obj["nvme-tcp-server-password"]; ok {
		if v != "" {
			m.NvmeTCPServerPassword = types.StringValue(v)
		} else {
			m.NvmeTCPServerPassword = types.StringNull()
		}
	}
	if v, ok := obj["nvme-tcp-server-port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.NvmeTCPServerPort = types.Int64Value(n)
		} else {
			m.NvmeTCPServerPort = types.Int64Null()
		}
	} else {
		m.NvmeTCPServerPort = types.Int64Null()
	}
	if v, ok := obj["nvme-tcp-server-secret"]; ok {
		if v != "" {
			m.NvmeTCPServerSecret = types.StringValue(v)
		} else {
			m.NvmeTCPServerSecret = types.StringNull()
		}
	}
	if v, ok := obj["oldfileman"]; ok {
		if v != "" {
			m.Oldfileman = types.StringValue(v)
		} else {
			m.Oldfileman = types.StringNull()
		}
	}
	if v, ok := obj["parent"]; ok {
		if v != "" {
			m.Parent = types.StringValue(v)
		} else {
			m.Parent = types.StringNull()
		}
	}
	if v, ok := obj["part"]; ok {
		if v != "" {
			m.Part = types.StringValue(v)
		} else {
			m.Part = types.StringNull()
		}
	}
	if v, ok := obj["partition"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Partition = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Partition = types.BoolValue(true)
		} else {
			m.Partition = types.BoolNull()
		}
	}
	if v, ok := obj["partition-number"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PartitionNumber = types.Int64Value(n)
		} else {
			m.PartitionNumber = types.Int64Null()
		}
	} else {
		m.PartitionNumber = types.Int64Null()
	}
	if v, ok := obj["partition-offset"]; ok {
		if v != "" {
			m.PartitionOffset = types.StringValue(v)
		} else {
			m.PartitionOffset = types.StringNull()
		}
	}
	if v, ok := obj["partition-size"]; ok {
		if v != "" {
			m.PartitionSize = types.StringValue(v)
		} else {
			m.PartitionSize = types.StringNull()
		}
	}
	if v, ok := obj["percentage-used"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PercentageUsed = types.Int64Value(n)
		} else {
			m.PercentageUsed = types.Int64Null()
		}
	} else {
		m.PercentageUsed = types.Int64Null()
	}
	if v, ok := obj["power-cycles"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.PowerCycles = types.Int64Value(n)
		} else {
			m.PowerCycles = types.Int64Null()
		}
	} else {
		m.PowerCycles = types.Int64Null()
	}
	if v, ok := obj["power-on-time"]; ok {
		if v != "" {
			m.PowerOnTime = types.StringValue(v)
		} else {
			m.PowerOnTime = types.StringNull()
		}
	}
	if v, ok := obj["raid"]; ok {
		if v != "" {
			m.Raid = types.StringValue(v)
		} else {
			m.Raid = types.StringNull()
		}
	}
	if v, ok := obj["raid-and-master"]; ok {
		if v != "" {
			m.RaidAndMaster = types.StringValue(v)
		} else {
			m.RaidAndMaster = types.StringNull()
		}
	}
	if v, ok := obj["raid-and-type"]; ok {
		if v != "" {
			m.RaidAndType = types.StringValue(v)
		} else {
			m.RaidAndType = types.StringNull()
		}
	}
	if v, ok := obj["raid-master"]; ok {
		if v != "" {
			m.RaidMaster = types.StringValue(v)
		} else {
			m.RaidMaster = types.StringNull()
		}
	}
	if v, ok := obj["raid-member"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.RaidMember = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.RaidMember = types.BoolValue(true)
		} else {
			m.RaidMember = types.BoolNull()
		}
	}
	if v, ok := obj["raid-member-failed"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.RaidMemberFailed = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.RaidMemberFailed = types.BoolValue(true)
		} else {
			m.RaidMemberFailed = types.BoolNull()
		}
	}
	if v, ok := obj["raid-role"]; ok {
		if v != "" {
			m.RaidRole = types.StringValue(v)
		} else {
			m.RaidRole = types.StringNull()
		}
	}
	if v, ok := obj["raid-scrub"]; ok {
		if v != "" {
			m.RaidScrub = types.StringValue(v)
		} else {
			m.RaidScrub = types.StringNull()
		}
	}
	if v, ok := obj["read-bytes"]; ok {
		if v != "" {
			m.ReadBytes = types.StringValue(v)
		} else {
			m.ReadBytes = types.StringNull()
		}
	}
	if v, ok := obj["read-merges"]; ok {
		if v != "" {
			m.ReadMerges = types.StringValue(v)
		} else {
			m.ReadMerges = types.StringNull()
		}
	}
	if v, ok := obj["read-only"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.ReadOnly = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.ReadOnly = types.BoolValue(true)
		} else {
			m.ReadOnly = types.BoolNull()
		}
	}
	if v, ok := obj["read-ops"]; ok {
		if v != "" {
			m.ReadOps = types.StringValue(v)
		} else {
			m.ReadOps = types.StringNull()
		}
	}
	if v, ok := obj["read-ops-per-second"]; ok {
		if v != "" {
			m.ReadOpsPerSecond = types.StringValue(v)
		} else {
			m.ReadOpsPerSecond = types.StringNull()
		}
	}
	if v, ok := obj["read-rate"]; ok {
		if v != "" {
			m.ReadRate = types.StringValue(v)
		} else {
			m.ReadRate = types.StringNull()
		}
	}
	if v, ok := obj["read-time"]; ok {
		if v != "" {
			m.ReadTime = types.StringValue(v)
		} else {
			m.ReadTime = types.StringNull()
		}
	}
	if v, ok := obj["reset-counters"]; ok {
		if v != "" {
			m.ResetCounters = types.StringValue(v)
		} else {
			m.ResetCounters = types.StringNull()
		}
	}
	if v, ok := obj["rose"]; ok {
		if v != "" {
			m.Rose = types.StringValue(v)
		} else {
			m.Rose = types.StringNull()
		}
	}
	if v, ok := obj["scan"]; ok {
		if v != "" {
			m.Scan = types.StringValue(v)
		} else {
			m.Scan = types.StringNull()
		}
	}
	if v, ok := obj["self-encrypted-and-locked"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SelfEncryptedAndLocked = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.SelfEncryptedAndLocked = types.BoolValue(true)
		} else {
			m.SelfEncryptedAndLocked = types.BoolNull()
		}
	}
	if v, ok := obj["self-encryption-enabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SelfEncryptionEnabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.SelfEncryptionEnabled = types.BoolValue(true)
		} else {
			m.SelfEncryptionEnabled = types.BoolNull()
		}
	}
	if v, ok := obj["self-encryption-password"]; ok {
		if v != "" {
			m.SelfEncryptionPassword = types.StringValue(v)
		} else {
			m.SelfEncryptionPassword = types.StringNull()
		}
	}
	if v, ok := obj["self-encryption-supported"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SelfEncryptionSupported = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.SelfEncryptionSupported = types.BoolValue(true)
		} else {
			m.SelfEncryptionSupported = types.BoolNull()
		}
	}
	if v, ok := obj["serial"]; ok {
		if v != "" {
			m.Serial = types.StringValue(v)
		} else {
			m.Serial = types.StringNull()
		}
	}
	if v, ok := obj["size"]; ok {
		if v != "" {
			m.Size = types.StringValue(v)
		} else {
			m.Size = types.StringNull()
		}
	}
	if v, ok := obj["slot"]; ok {
		if v != "" {
			m.Slot = types.StringValue(v)
		} else {
			m.Slot = types.StringNull()
		}
	}
	if v, ok := obj["smb-server-encryption"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SmbServerEncryption = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.SmbServerEncryption = types.BoolValue(true)
		} else {
			m.SmbServerEncryption = types.BoolNull()
		}
	}
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.SmbServerPassword already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["smb-server-password"]; ok && v != "" {
		_ = v
		if v != "" {
			m.SmbServerPassword = types.StringValue(v)
		} else {
			m.SmbServerPassword = types.StringNull()
		}
	} else if m.SmbServerPassword.IsUnknown() {
		m.SmbServerPassword = types.StringNull()
	}
	if v, ok := obj["smb-server-user"]; ok {
		if v != "" {
			m.SmbServerUser = types.StringValue(v)
		} else {
			m.SmbServerUser = types.StringNull()
		}
	}
	if v, ok := obj["smb-sharing"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SmbSharing = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.SmbSharing = types.BoolValue(true)
		} else {
			m.SmbSharing = types.BoolNull()
		}
	}
	if v, ok := obj["state"]; ok {
		if v != "" {
			m.State = types.StringValue(v)
		} else {
			m.State = types.StringNull()
		}
	}
	if v, ok := obj["swap"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Swap = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Swap = types.BoolValue(true)
		} else {
			m.Swap = types.BoolNull()
		}
	}
	if v, ok := obj["swap-enabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SwapEnabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.SwapEnabled = types.BoolValue(true)
		} else {
			m.SwapEnabled = types.BoolNull()
		}
	}
	if v, ok := obj["temperature"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Temperature = types.Int64Value(n)
		} else {
			m.Temperature = types.Int64Null()
		}
	} else {
		m.Temperature = types.Int64Null()
	}
	if v, ok := obj["temperatures"]; ok {
		if v != "" {
			m.Temperatures = types.StringValue(v)
		} else {
			m.Temperatures = types.StringNull()
		}
	}
	if v, ok := obj["tmpfs"]; ok {
		if v != "" {
			m.Tmpfs = types.StringValue(v)
		} else {
			m.Tmpfs = types.StringNull()
		}
	}
	if v, ok := obj["tmpfs-max-size"]; ok {
		if v != "" {
			m.TmpfsMaxSize = types.StringValue(v)
		} else {
			m.TmpfsMaxSize = types.StringNull()
		}
	}
	if v, ok := obj["trim"]; ok {
		if v != "" {
			m.Trim = types.StringValue(v)
		} else {
			m.Trim = types.StringNull()
		}
	}
	if v, ok := obj["type"]; ok {
		if v != "" {
			m.Type = types.StringValue(v)
		} else {
			m.Type = types.StringNull()
		}
	}
	if v, ok := obj["unrecovered-integrity-errors"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.UnrecoveredIntegrityErrors = types.Int64Value(n)
		} else {
			m.UnrecoveredIntegrityErrors = types.Int64Null()
		}
	} else {
		m.UnrecoveredIntegrityErrors = types.Int64Null()
	}
	if v, ok := obj["unsafe-shutdown"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.UnsafeShutdown = types.Int64Value(n)
		} else {
			m.UnsafeShutdown = types.Int64Null()
		}
	} else {
		m.UnsafeShutdown = types.Int64Null()
	}
	if v, ok := obj["use"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Use = types.Int64Value(n)
		} else {
			m.Use = types.Int64Null()
		}
	} else {
		m.Use = types.Int64Null()
	}
	if v, ok := obj["uuid"]; ok {
		if v != "" {
			m.Uuid = types.StringValue(v)
		} else {
			m.Uuid = types.StringNull()
		}
	}
	if v, ok := obj["wait-time"]; ok {
		if v != "" {
			m.WaitTime = types.StringValue(v)
		} else {
			m.WaitTime = types.StringNull()
		}
	}
	if v, ok := obj["warning-temperature"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.WarningTemperature = types.Int64Value(n)
		} else {
			m.WarningTemperature = types.Int64Null()
		}
	} else {
		m.WarningTemperature = types.Int64Null()
	}
	if v, ok := obj["warning-temperature-time"]; ok {
		if v != "" {
			m.WarningTemperatureTime = types.StringValue(v)
		} else {
			m.WarningTemperatureTime = types.StringNull()
		}
	}
	if v, ok := obj["write-bytes"]; ok {
		if v != "" {
			m.WriteBytes = types.StringValue(v)
		} else {
			m.WriteBytes = types.StringNull()
		}
	}
	if v, ok := obj["write-merges"]; ok {
		if v != "" {
			m.WriteMerges = types.StringValue(v)
		} else {
			m.WriteMerges = types.StringNull()
		}
	}
	if v, ok := obj["write-ops"]; ok {
		if v != "" {
			m.WriteOps = types.StringValue(v)
		} else {
			m.WriteOps = types.StringNull()
		}
	}
	if v, ok := obj["write-ops-per-second"]; ok {
		if v != "" {
			m.WriteOpsPerSecond = types.StringValue(v)
		} else {
			m.WriteOpsPerSecond = types.StringNull()
		}
	}
	if v, ok := obj["write-rate"]; ok {
		if v != "" {
			m.WriteRate = types.StringValue(v)
		} else {
			m.WriteRate = types.StringNull()
		}
	}
	if v, ok := obj["write-time"]; ok {
		if v != "" {
			m.WriteTime = types.StringValue(v)
		} else {
			m.WriteTime = types.StringNull()
		}
	}
}
