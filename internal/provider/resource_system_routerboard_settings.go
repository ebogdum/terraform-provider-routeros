package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &SystemRouterboardSettingsResource{}
	_ resource.ResourceWithImportState = &SystemRouterboardSettingsResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type SystemRouterboardSettingsResource struct {
	reg *client.Registry
}

type SystemRouterboardSettingsModel struct {
	ID                     types.String `tfsdk:"id"`
	PreferredArchitecture  types.String `tfsdk:"preferred_architecture"`
	InitDelay              types.String `tfsdk:"init_delay"`
	GpioFunction           types.String `tfsdk:"gpio_function"`
	EtherbootPort          types.String `tfsdk:"etherboot_port"`
	EnterSetupOn           types.String `tfsdk:"enter_setup_on"`
	EnableJumperReset      types.String `tfsdk:"enable_jumper_reset"`
	DisablePci             types.String `tfsdk:"disable_pci"`
	BootOs                 types.String `tfsdk:"boot_os"`
	BootDelay              types.String `tfsdk:"boot_delay"`
	BaudRate               types.String `tfsdk:"baud_rate"`
	AutoUpgrade            types.String `tfsdk:"auto_upgrade"`
	BootDevice             types.String `tfsdk:"boot_device"`
	BootProtocol           types.String `tfsdk:"boot_protocol"`
	CpuFrequency           types.String `tfsdk:"cpu_frequency"`
	ForceBackupBooter      types.String `tfsdk:"force_backup_booter"`
	PrebootEtherboot       types.String `tfsdk:"preboot_etherboot"`
	PrebootEtherbootServer types.String `tfsdk:"preboot_etherboot_server"`
	ProtectedRouterboot    types.String `tfsdk:"protected_routerboot"`
	ReformatHoldButton     types.String `tfsdk:"reformat_hold_button"`
	ReformatHoldButtonMax  types.String `tfsdk:"reformat_hold_button_max"`
	SilentBoot             types.String `tfsdk:"silent_boot"`
	Router                 types.String `tfsdk:"router"`
}

func NewSystemRouterboardSettingsResource() resource.Resource {
	return &SystemRouterboardSettingsResource{}
}

func (r *SystemRouterboardSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_routerboard_settings"
}

func (r *SystemRouterboardSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemRouterboardSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/system/routerboard/settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"preferred_architecture": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `preferred-architecture`.",
			},
			"init_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `init-delay`.",
			},
			"gpio_function": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `gpio-function`.",
			},
			"etherboot_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `etherboot-port`.",
			},
			"enter_setup_on": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `enter-setup-on`.",
			},
			"enable_jumper_reset": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `enable-jumper-reset`.",
			},
			"disable_pci": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `disable-pci`.",
			},
			"boot_os": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `boot-os`.",
			},
			"boot_delay": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `boot-delay`.",
			},
			"baud_rate": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `baud-rate`.",
			},
			"auto_upgrade": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"boot_device": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"boot_protocol": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"cpu_frequency": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"force_backup_booter": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"preboot_etherboot": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"preboot_etherboot_server": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"protected_routerboot": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"reformat_hold_button": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"reformat_hold_button_max": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"silent_boot": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemRouterboardSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemRouterboardSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemRouterboardSettingsUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemRouterboardSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SystemRouterboardSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemRouterboardSettingsUpsert(ctx, r.reg, &plan, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemRouterboardSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemRouterboardSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/system/routerboard/settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /system/routerboard/settings failed", err.Error())
		return
	}
	systemRouterboardSettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/system/routerboard/settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemRouterboardSettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *SystemRouterboardSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/system/routerboard/settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/system/routerboard/settings", types.StringValue(routerName))))...)
}

func systemRouterboardSettingsUpsert(ctx context.Context, reg *client.Registry, plan *SystemRouterboardSettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AutoUpgrade.IsNull() || plan.AutoUpgrade.IsUnknown()) {
		body["auto-upgrade"] = plan.AutoUpgrade.ValueString()
	}
	if !(plan.BootDevice.IsNull() || plan.BootDevice.IsUnknown()) {
		body["boot-device"] = plan.BootDevice.ValueString()
	}
	if !(plan.BootProtocol.IsNull() || plan.BootProtocol.IsUnknown()) {
		body["boot-protocol"] = plan.BootProtocol.ValueString()
	}
	if !(plan.CpuFrequency.IsNull() || plan.CpuFrequency.IsUnknown()) {
		body["cpu-frequency"] = plan.CpuFrequency.ValueString()
	}
	if !(plan.ForceBackupBooter.IsNull() || plan.ForceBackupBooter.IsUnknown()) {
		body["force-backup-booter"] = plan.ForceBackupBooter.ValueString()
	}
	if !(plan.PrebootEtherboot.IsNull() || plan.PrebootEtherboot.IsUnknown()) {
		body["preboot-etherboot"] = plan.PrebootEtherboot.ValueString()
	}
	if !(plan.PrebootEtherbootServer.IsNull() || plan.PrebootEtherbootServer.IsUnknown()) {
		body["preboot-etherboot-server"] = plan.PrebootEtherbootServer.ValueString()
	}
	if !(plan.ProtectedRouterboot.IsNull() || plan.ProtectedRouterboot.IsUnknown()) {
		body["protected-routerboot"] = plan.ProtectedRouterboot.ValueString()
	}
	if !(plan.ReformatHoldButton.IsNull() || plan.ReformatHoldButton.IsUnknown()) {
		body["reformat-hold-button"] = plan.ReformatHoldButton.ValueString()
	}
	if !(plan.ReformatHoldButtonMax.IsNull() || plan.ReformatHoldButtonMax.IsUnknown()) {
		body["reformat-hold-button-max"] = plan.ReformatHoldButtonMax.ValueString()
	}
	if !(plan.SilentBoot.IsNull() || plan.SilentBoot.IsUnknown()) {
		body["silent-boot"] = plan.SilentBoot.ValueString()
	}
	if !(plan.BaudRate.IsNull() || plan.BaudRate.IsUnknown()) {
		body["baud-rate"] = plan.BaudRate.ValueString()
	}
	if !(plan.BootDelay.IsNull() || plan.BootDelay.IsUnknown()) {
		body["boot-delay"] = plan.BootDelay.ValueString()
	}
	if !(plan.BootOs.IsNull() || plan.BootOs.IsUnknown()) {
		body["boot-os"] = plan.BootOs.ValueString()
	}
	if !(plan.DisablePci.IsNull() || plan.DisablePci.IsUnknown()) {
		body["disable-pci"] = plan.DisablePci.ValueString()
	}
	if !(plan.EnableJumperReset.IsNull() || plan.EnableJumperReset.IsUnknown()) {
		body["enable-jumper-reset"] = plan.EnableJumperReset.ValueString()
	}
	if !(plan.EnterSetupOn.IsNull() || plan.EnterSetupOn.IsUnknown()) {
		body["enter-setup-on"] = plan.EnterSetupOn.ValueString()
	}
	if !(plan.EtherbootPort.IsNull() || plan.EtherbootPort.IsUnknown()) {
		body["etherboot-port"] = plan.EtherbootPort.ValueString()
	}
	if !(plan.GpioFunction.IsNull() || plan.GpioFunction.IsUnknown()) {
		body["gpio-function"] = plan.GpioFunction.ValueString()
	}
	if !(plan.InitDelay.IsNull() || plan.InitDelay.IsUnknown()) {
		body["init-delay"] = plan.InitDelay.ValueString()
	}
	if !(plan.PreferredArchitecture.IsNull() || plan.PreferredArchitecture.IsUnknown()) {
		body["preferred-architecture"] = plan.PreferredArchitecture.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/system/routerboard/settings", body)
	if err != nil {
		diags.AddError("Upsert /system/routerboard/settings failed", err.Error())
		return
	}
	systemRouterboardSettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/system/routerboard/settings", plan.Router))
}

func systemRouterboardSettingsApply(ctx context.Context, obj client.Object, m *SystemRouterboardSettingsModel) {
	_ = ctx
	if v, ok := obj["preferred-architecture"]; ok && v != "" {
		m.PreferredArchitecture = types.StringValue(v)
	} else {
		m.PreferredArchitecture = types.StringNull()
	}
	if v, ok := obj["init-delay"]; ok && v != "" {
		m.InitDelay = types.StringValue(v)
	} else {
		m.InitDelay = types.StringNull()
	}
	if v, ok := obj["gpio-function"]; ok && v != "" {
		m.GpioFunction = types.StringValue(v)
	} else {
		m.GpioFunction = types.StringNull()
	}
	if v, ok := obj["etherboot-port"]; ok && v != "" {
		m.EtherbootPort = types.StringValue(v)
	} else {
		m.EtherbootPort = types.StringNull()
	}
	if v, ok := obj["enter-setup-on"]; ok && v != "" {
		m.EnterSetupOn = types.StringValue(v)
	} else {
		m.EnterSetupOn = types.StringNull()
	}
	if v, ok := obj["enable-jumper-reset"]; ok && v != "" {
		m.EnableJumperReset = types.StringValue(v)
	} else {
		m.EnableJumperReset = types.StringNull()
	}
	if v, ok := obj["disable-pci"]; ok && v != "" {
		m.DisablePci = types.StringValue(v)
	} else {
		m.DisablePci = types.StringNull()
	}
	if v, ok := obj["boot-os"]; ok && v != "" {
		m.BootOs = types.StringValue(v)
	} else {
		m.BootOs = types.StringNull()
	}
	if v, ok := obj["boot-delay"]; ok && v != "" {
		m.BootDelay = types.StringValue(v)
	} else {
		m.BootDelay = types.StringNull()
	}
	if v, ok := obj["baud-rate"]; ok && v != "" {
		m.BaudRate = types.StringValue(v)
	} else {
		m.BaudRate = types.StringNull()
	}
	if v, ok := obj["auto-upgrade"]; ok {
		_ = v
		if v != "" {
			m.AutoUpgrade = types.StringValue(v)
		} else {
			m.AutoUpgrade = types.StringNull()
		}
	}
	if v, ok := obj["boot-device"]; ok {
		_ = v
		if v != "" {
			m.BootDevice = types.StringValue(v)
		} else {
			m.BootDevice = types.StringNull()
		}
	}
	if v, ok := obj["boot-protocol"]; ok {
		_ = v
		if v != "" {
			m.BootProtocol = types.StringValue(v)
		} else {
			m.BootProtocol = types.StringNull()
		}
	}
	if v, ok := obj["cpu-frequency"]; ok {
		_ = v
		if v != "" {
			m.CpuFrequency = types.StringValue(v)
		} else {
			m.CpuFrequency = types.StringNull()
		}
	}
	if v, ok := obj["force-backup-booter"]; ok {
		_ = v
		if v != "" {
			m.ForceBackupBooter = types.StringValue(v)
		} else {
			m.ForceBackupBooter = types.StringNull()
		}
	}
	if v, ok := obj["preboot-etherboot"]; ok {
		_ = v
		if v != "" {
			m.PrebootEtherboot = types.StringValue(v)
		} else {
			m.PrebootEtherboot = types.StringNull()
		}
	}
	if v, ok := obj["preboot-etherboot-server"]; ok {
		_ = v
		if v != "" {
			m.PrebootEtherbootServer = types.StringValue(v)
		} else {
			m.PrebootEtherbootServer = types.StringNull()
		}
	}
	if v, ok := obj["protected-routerboot"]; ok {
		_ = v
		if v != "" {
			m.ProtectedRouterboot = types.StringValue(v)
		} else {
			m.ProtectedRouterboot = types.StringNull()
		}
	}
	if v, ok := obj["reformat-hold-button"]; ok {
		_ = v
		if v != "" {
			m.ReformatHoldButton = types.StringValue(v)
		} else {
			m.ReformatHoldButton = types.StringNull()
		}
	}
	if v, ok := obj["reformat-hold-button-max"]; ok {
		_ = v
		if v != "" {
			m.ReformatHoldButtonMax = types.StringValue(v)
		} else {
			m.ReformatHoldButtonMax = types.StringNull()
		}
	}
	if v, ok := obj["silent-boot"]; ok {
		_ = v
		if v != "" {
			m.SilentBoot = types.StringValue(v)
		} else {
			m.SilentBoot = types.StringNull()
		}
	}
}
