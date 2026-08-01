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
	_ resource.Resource                = &SystemResourceHardwareResource{}
	_ resource.ResourceWithImportState = &SystemResourceHardwareResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemResourceHardwareResource struct {
	reg *client.Registry
}

type SystemResourceHardwareModel struct {
	ID            types.String `tfsdk:"id"`
	Slot          types.String `tfsdk:"slot"`
	Duration      types.String `tfsdk:"duration"`
	Bus           types.String `tfsdk:"bus"`
	Authorization types.String `tfsdk:"authorization"`
	Allow         types.String `tfsdk:"allow"`
	Category      types.String `tfsdk:"category"`
	DeviceID      types.String `tfsdk:"device_id"`
	Devices       types.String `tfsdk:"devices"`
	Io            types.String `tfsdk:"io"`
	Irq           types.Int64  `tfsdk:"irq"`
	Location      types.String `tfsdk:"location"`
	Memory        types.String `tfsdk:"memory"`
	Name          types.String `tfsdk:"name"`
	Owner         types.String `tfsdk:"owner"`
	Parent        types.Int64  `tfsdk:"parent"`
	Pci           types.String `tfsdk:"pci"`
	Ports         types.Int64  `tfsdk:"ports"`
	SerialNumber  types.String `tfsdk:"serial_number"`
	Speed         types.String `tfsdk:"speed"`
	StdDescr      types.String `tfsdk:"std_descr"`
	Type          types.String `tfsdk:"type"`
	Usb           types.String `tfsdk:"usb"`
	UsbVersion    types.String `tfsdk:"usb_version"`
	Vendor        types.String `tfsdk:"vendor"`
	VendorID      types.String `tfsdk:"vendor_id"`
	Router        types.String `tfsdk:"router"`
}

func NewSystemResourceHardwareResource() resource.Resource { return &SystemResourceHardwareResource{} }

func (r *SystemResourceHardwareResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_resource_hardware"
}

func (r *SystemResourceHardwareResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemResourceHardwareResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Discovered read-only menu",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"slot": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `slot`.",
			},
			"duration": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `duration`.",
			},
			"bus": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `bus`.",
			},
			"authorization": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `authorization`.",
			},
			"allow": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `allow`.",
			},
			"category": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"device_id": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"devices": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"io": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"irq": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"location": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"memory": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"owner": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"parent": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"pci": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ports": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"serial_number": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"speed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"std_descr": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"type": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"usb", "pci", "scsi", "serial"}...)},
			},
			"usb": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"usb_version": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vendor": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"vendor_id": schema.StringAttribute{
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

func (r *SystemResourceHardwareResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemResourceHardwareModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Category.IsNull() || plan.Category.IsUnknown()) {
		body["category"] = plan.Category.ValueString()
	}
	if !(plan.DeviceID.IsNull() || plan.DeviceID.IsUnknown()) {
		body["device-id"] = plan.DeviceID.ValueString()
	}
	if !(plan.Devices.IsNull() || plan.Devices.IsUnknown()) {
		body["devices"] = plan.Devices.ValueString()
	}
	if !(plan.Io.IsNull() || plan.Io.IsUnknown()) {
		body["io"] = plan.Io.ValueString()
	}
	if !(plan.Irq.IsNull() || plan.Irq.IsUnknown()) {
		body["irq"] = client.FormatInt64(plan.Irq.ValueInt64())
	}
	if !(plan.Location.IsNull() || plan.Location.IsUnknown()) {
		body["location"] = plan.Location.ValueString()
	}
	if !(plan.Memory.IsNull() || plan.Memory.IsUnknown()) {
		body["memory"] = plan.Memory.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Owner.IsNull() || plan.Owner.IsUnknown()) {
		body["owner"] = plan.Owner.ValueString()
	}
	if !(plan.Parent.IsNull() || plan.Parent.IsUnknown()) {
		body["parent"] = client.FormatInt64(plan.Parent.ValueInt64())
	}
	if !(plan.Pci.IsNull() || plan.Pci.IsUnknown()) {
		body["pci"] = plan.Pci.ValueString()
	}
	if !(plan.Ports.IsNull() || plan.Ports.IsUnknown()) {
		body["ports"] = client.FormatInt64(plan.Ports.ValueInt64())
	}
	if !(plan.SerialNumber.IsNull() || plan.SerialNumber.IsUnknown()) {
		body["serial-number"] = plan.SerialNumber.ValueString()
	}
	if !(plan.Speed.IsNull() || plan.Speed.IsUnknown()) {
		body["speed"] = plan.Speed.ValueString()
	}
	if !(plan.StdDescr.IsNull() || plan.StdDescr.IsUnknown()) {
		body["std-descr"] = plan.StdDescr.ValueString()
	}
	if !(plan.Type.IsNull() || plan.Type.IsUnknown()) {
		body["type"] = plan.Type.ValueString()
	}
	if !(plan.Usb.IsNull() || plan.Usb.IsUnknown()) {
		body["usb"] = plan.Usb.ValueString()
	}
	if !(plan.UsbVersion.IsNull() || plan.UsbVersion.IsUnknown()) {
		body["usb-version"] = plan.UsbVersion.ValueString()
	}
	if !(plan.Vendor.IsNull() || plan.Vendor.IsUnknown()) {
		body["vendor"] = plan.Vendor.ValueString()
	}
	if !(plan.VendorID.IsNull() || plan.VendorID.IsUnknown()) {
		body["vendor-id"] = plan.VendorID.ValueString()
	}
	if !(plan.Allow.IsNull() || plan.Allow.IsUnknown()) {
		body["allow"] = plan.Allow.ValueString()
	}
	if !(plan.Authorization.IsNull() || plan.Authorization.IsUnknown()) {
		body["authorization"] = plan.Authorization.ValueString()
	}
	if !(plan.Bus.IsNull() || plan.Bus.IsUnknown()) {
		body["bus"] = plan.Bus.ValueString()
	}
	if !(plan.Duration.IsNull() || plan.Duration.IsUnknown()) {
		body["duration"] = plan.Duration.ValueString()
	}
	if !(plan.Slot.IsNull() || plan.Slot.IsUnknown()) {
		body["slot"] = plan.Slot.ValueString()
	}
	obj, err := c.Add(ctx, "/system/resource/hardware", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /system/resource/hardware failed", err.Error())
		return
	}
	systemResourceHardwareApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemResourceHardwareResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemResourceHardwareModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/system/resource/hardware", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /system/resource/hardware failed", err.Error())
		return
	}
	systemResourceHardwareApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemResourceHardwareResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemResourceHardwareModel
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
	if !plan.Category.Equal(state.Category) && !plan.Category.IsUnknown() {
		body["category"] = plan.Category.ValueString()
	}
	if !plan.DeviceID.Equal(state.DeviceID) && !plan.DeviceID.IsUnknown() {
		body["device-id"] = plan.DeviceID.ValueString()
	}
	if !plan.Devices.Equal(state.Devices) && !plan.Devices.IsUnknown() {
		body["devices"] = plan.Devices.ValueString()
	}
	if !plan.Io.Equal(state.Io) && !plan.Io.IsUnknown() {
		body["io"] = plan.Io.ValueString()
	}
	if !plan.Irq.Equal(state.Irq) && !plan.Irq.IsUnknown() {
		body["irq"] = client.FormatInt64(plan.Irq.ValueInt64())
	}
	if !plan.Location.Equal(state.Location) && !plan.Location.IsUnknown() {
		body["location"] = plan.Location.ValueString()
	}
	if !plan.Memory.Equal(state.Memory) && !plan.Memory.IsUnknown() {
		body["memory"] = plan.Memory.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Owner.Equal(state.Owner) && !plan.Owner.IsUnknown() {
		body["owner"] = plan.Owner.ValueString()
	}
	if !plan.Parent.Equal(state.Parent) && !plan.Parent.IsUnknown() {
		body["parent"] = client.FormatInt64(plan.Parent.ValueInt64())
	}
	if !plan.Pci.Equal(state.Pci) && !plan.Pci.IsUnknown() {
		body["pci"] = plan.Pci.ValueString()
	}
	if !plan.Ports.Equal(state.Ports) && !plan.Ports.IsUnknown() {
		body["ports"] = client.FormatInt64(plan.Ports.ValueInt64())
	}
	if !plan.SerialNumber.Equal(state.SerialNumber) && !plan.SerialNumber.IsUnknown() {
		body["serial-number"] = plan.SerialNumber.ValueString()
	}
	if !plan.Speed.Equal(state.Speed) && !plan.Speed.IsUnknown() {
		body["speed"] = plan.Speed.ValueString()
	}
	if !plan.StdDescr.Equal(state.StdDescr) && !plan.StdDescr.IsUnknown() {
		body["std-descr"] = plan.StdDescr.ValueString()
	}
	if !plan.Type.Equal(state.Type) && !plan.Type.IsUnknown() {
		body["type"] = plan.Type.ValueString()
	}
	if !plan.Usb.Equal(state.Usb) && !plan.Usb.IsUnknown() {
		body["usb"] = plan.Usb.ValueString()
	}
	if !plan.UsbVersion.Equal(state.UsbVersion) && !plan.UsbVersion.IsUnknown() {
		body["usb-version"] = plan.UsbVersion.ValueString()
	}
	if !plan.Vendor.Equal(state.Vendor) && !plan.Vendor.IsUnknown() {
		body["vendor"] = plan.Vendor.ValueString()
	}
	if !plan.VendorID.Equal(state.VendorID) && !plan.VendorID.IsUnknown() {
		body["vendor-id"] = plan.VendorID.ValueString()
	}
	if !plan.Allow.Equal(state.Allow) && !plan.Allow.IsUnknown() {
		body["allow"] = plan.Allow.ValueString()
	}
	if !plan.Authorization.Equal(state.Authorization) && !plan.Authorization.IsUnknown() {
		body["authorization"] = plan.Authorization.ValueString()
	}
	if !plan.Bus.Equal(state.Bus) && !plan.Bus.IsUnknown() {
		body["bus"] = plan.Bus.ValueString()
	}
	if !plan.Duration.Equal(state.Duration) && !plan.Duration.IsUnknown() {
		body["duration"] = plan.Duration.ValueString()
	}
	if !plan.Slot.Equal(state.Slot) && !plan.Slot.IsUnknown() {
		body["slot"] = plan.Slot.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/system/resource/hardware", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /system/resource/hardware failed", err.Error())
			return
		}
		systemResourceHardwareApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemResourceHardwareResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemResourceHardwareModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/system/resource/hardware", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /system/resource/hardware failed", err.Error())
	}
}

func (r *SystemResourceHardwareResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := systemResourceHardwareLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /system/resource/hardware matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// systemResourceHardwareLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func systemResourceHardwareLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/system/resource/hardware", id)
}

func systemResourceHardwareApply(ctx context.Context, obj client.Object, m *SystemResourceHardwareModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["slot"]; ok && v != "" {
		m.Slot = types.StringValue(v)
	} else {
		m.Slot = types.StringNull()
	}
	if v, ok := obj["duration"]; ok && v != "" {
		m.Duration = types.StringValue(v)
	} else {
		m.Duration = types.StringNull()
	}
	if v, ok := obj["bus"]; ok && v != "" {
		m.Bus = types.StringValue(v)
	} else {
		m.Bus = types.StringNull()
	}
	if v, ok := obj["authorization"]; ok && v != "" {
		m.Authorization = types.StringValue(v)
	} else {
		m.Authorization = types.StringNull()
	}
	if v, ok := obj["allow"]; ok && v != "" {
		m.Allow = types.StringValue(v)
	} else {
		m.Allow = types.StringNull()
	}
	if v, ok := obj["category"]; ok {
		_ = v
		if v != "" {
			m.Category = types.StringValue(v)
		} else {
			m.Category = types.StringNull()
		}
	} else {
		m.Category = types.StringNull()
	}
	if v, ok := obj["device-id"]; ok {
		_ = v
		if v != "" {
			m.DeviceID = types.StringValue(v)
		} else {
			m.DeviceID = types.StringNull()
		}
	} else {
		m.DeviceID = types.StringNull()
	}
	if v, ok := obj["devices"]; ok {
		_ = v
		if v != "" {
			m.Devices = types.StringValue(v)
		} else {
			m.Devices = types.StringNull()
		}
	} else {
		m.Devices = types.StringNull()
	}
	if v, ok := obj["io"]; ok {
		_ = v
		if v != "" {
			m.Io = types.StringValue(v)
		} else {
			m.Io = types.StringNull()
		}
	} else {
		m.Io = types.StringNull()
	}
	if v, ok := obj["irq"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Irq = types.Int64Value(n)
		} else {
			m.Irq = types.Int64Null()
		}
	} else {
		m.Irq = types.Int64Null()
	}
	if v, ok := obj["location"]; ok {
		_ = v
		if v != "" {
			m.Location = types.StringValue(v)
		} else {
			m.Location = types.StringNull()
		}
	} else {
		m.Location = types.StringNull()
	}
	if v, ok := obj["memory"]; ok {
		_ = v
		if v != "" {
			m.Memory = types.StringValue(v)
		} else {
			m.Memory = types.StringNull()
		}
	} else {
		m.Memory = types.StringNull()
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
	if v, ok := obj["owner"]; ok {
		_ = v
		if v != "" {
			m.Owner = types.StringValue(v)
		} else {
			m.Owner = types.StringNull()
		}
	} else {
		m.Owner = types.StringNull()
	}
	if v, ok := obj["parent"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Parent = types.Int64Value(n)
		} else {
			m.Parent = types.Int64Null()
		}
	} else {
		m.Parent = types.Int64Null()
	}
	if v, ok := obj["pci"]; ok {
		_ = v
		if v != "" {
			m.Pci = types.StringValue(v)
		} else {
			m.Pci = types.StringNull()
		}
	} else {
		m.Pci = types.StringNull()
	}
	if v, ok := obj["ports"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Ports = types.Int64Value(n)
		} else {
			m.Ports = types.Int64Null()
		}
	} else {
		m.Ports = types.Int64Null()
	}
	if v, ok := obj["serial-number"]; ok {
		_ = v
		if v != "" {
			m.SerialNumber = types.StringValue(v)
		} else {
			m.SerialNumber = types.StringNull()
		}
	} else {
		m.SerialNumber = types.StringNull()
	}
	if v, ok := obj["speed"]; ok {
		_ = v
		if v != "" {
			m.Speed = types.StringValue(v)
		} else {
			m.Speed = types.StringNull()
		}
	} else {
		m.Speed = types.StringNull()
	}
	if v, ok := obj["std-descr"]; ok {
		_ = v
		if v != "" {
			m.StdDescr = types.StringValue(v)
		} else {
			m.StdDescr = types.StringNull()
		}
	} else {
		m.StdDescr = types.StringNull()
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
	if v, ok := obj["usb"]; ok {
		_ = v
		if v != "" {
			m.Usb = types.StringValue(v)
		} else {
			m.Usb = types.StringNull()
		}
	} else {
		m.Usb = types.StringNull()
	}
	if v, ok := obj["usb-version"]; ok {
		_ = v
		if v != "" {
			m.UsbVersion = types.StringValue(v)
		} else {
			m.UsbVersion = types.StringNull()
		}
	} else {
		m.UsbVersion = types.StringNull()
	}
	if v, ok := obj["vendor"]; ok {
		_ = v
		if v != "" {
			m.Vendor = types.StringValue(v)
		} else {
			m.Vendor = types.StringNull()
		}
	} else {
		m.Vendor = types.StringNull()
	}
	if v, ok := obj["vendor-id"]; ok {
		_ = v
		if v != "" {
			m.VendorID = types.StringValue(v)
		} else {
			m.VendorID = types.StringNull()
		}
	} else {
		m.VendorID = types.StringNull()
	}
}
