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
	_ resource.Resource                = &IPKidControlDeviceResource{}
	_ resource.ResourceWithImportState = &IPKidControlDeviceResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPKidControlDeviceResource struct {
	reg *client.Registry
}

type IPKidControlDeviceModel struct {
	ID            types.String `tfsdk:"id"`
	Activity      types.String `tfsdk:"activity"`
	Blocked       types.Bool   `tfsdk:"blocked"`
	Bytes         types.String `tfsdk:"bytes"`
	Disabled      types.Bool   `tfsdk:"disabled"`
	Dynamic       types.Bool   `tfsdk:"dynamic"`
	IPAddress     types.String `tfsdk:"ip_address"`
	MACAddress    types.String `tfsdk:"mac_address"`
	Name          types.String `tfsdk:"name"`
	RateLimited   types.Bool   `tfsdk:"rate_limited"`
	RateUpDown    types.String `tfsdk:"rate_up_down"`
	ResetCounters types.String `tfsdk:"reset_counters"`
	User          types.String `tfsdk:"user"`
	Router        types.String `tfsdk:"router"`
}

func NewIPKidControlDeviceResource() resource.Resource { return &IPKidControlDeviceResource{} }

func (r *IPKidControlDeviceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_kid_control_device"
}

func (r *IPKidControlDeviceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *IPKidControlDeviceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Discovered; needs kid-control-name fixture",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"activity": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"blocked": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"bytes": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dynamic": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ip_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mac_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rate_limited": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"rate_up_down": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"reset_counters": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"user": schema.StringAttribute{
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

func (r *IPKidControlDeviceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPKidControlDeviceModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Blocked.IsNull() || plan.Blocked.IsUnknown()) {
		body["blocked"] = client.FormatBool(plan.Blocked.ValueBool())
	}
	if !(plan.Bytes.IsNull() || plan.Bytes.IsUnknown()) {
		body["bytes"] = plan.Bytes.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.MACAddress.IsNull() || plan.MACAddress.IsUnknown()) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.RateLimited.IsNull() || plan.RateLimited.IsUnknown()) {
		body["rate-limited"] = client.FormatBool(plan.RateLimited.ValueBool())
	}
	if !(plan.RateUpDown.IsNull() || plan.RateUpDown.IsUnknown()) {
		body["rate-up-down"] = plan.RateUpDown.ValueString()
	}
	if !(plan.ResetCounters.IsNull() || plan.ResetCounters.IsUnknown()) {
		body["reset-counters"] = plan.ResetCounters.ValueString()
	}
	if !(plan.User.IsNull() || plan.User.IsUnknown()) {
		body["user"] = plan.User.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/kid-control/device", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/kid-control/device failed", err.Error())
		return
	}
	iPKidControlDeviceApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPKidControlDeviceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPKidControlDeviceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/kid-control/device", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/kid-control/device failed", err.Error())
		return
	}
	iPKidControlDeviceApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPKidControlDeviceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPKidControlDeviceModel
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
	if !plan.Blocked.Equal(state.Blocked) {
		body["blocked"] = client.FormatBool(plan.Blocked.ValueBool())
	}
	if !plan.Bytes.Equal(state.Bytes) {
		body["bytes"] = plan.Bytes.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.MACAddress.Equal(state.MACAddress) {
		body["mac-address"] = plan.MACAddress.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.RateLimited.Equal(state.RateLimited) {
		body["rate-limited"] = client.FormatBool(plan.RateLimited.ValueBool())
	}
	if !plan.RateUpDown.Equal(state.RateUpDown) {
		body["rate-up-down"] = plan.RateUpDown.ValueString()
	}
	if !plan.ResetCounters.Equal(state.ResetCounters) {
		body["reset-counters"] = plan.ResetCounters.ValueString()
	}
	if !plan.User.Equal(state.User) {
		body["user"] = plan.User.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/kid-control/device", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/kid-control/device failed", err.Error())
			return
		}
		iPKidControlDeviceApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPKidControlDeviceResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPKidControlDeviceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/kid-control/device", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/kid-control/device failed", err.Error())
	}
}

func (r *IPKidControlDeviceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPKidControlDeviceLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/kid-control/device matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPKidControlDeviceLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPKidControlDeviceLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/kid-control/device", id)
}

func iPKidControlDeviceApply(ctx context.Context, obj client.Object, m *IPKidControlDeviceModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["activity"]; ok {
		_ = v
		if v != "" {
			m.Activity = types.StringValue(v)
		} else {
			m.Activity = types.StringNull()
		}
	} else {
		m.Activity = types.StringNull()
	}
	if v, ok := obj["blocked"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Blocked = types.BoolValue(b)
		} else {
			m.Blocked = types.BoolNull()
		}
	} else {
		m.Blocked = types.BoolNull()
	}
	if v, ok := obj["bytes"]; ok {
		_ = v
		if v != "" {
			m.Bytes = types.StringValue(v)
		} else {
			m.Bytes = types.StringNull()
		}
	} else {
		m.Bytes = types.StringNull()
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
	if v, ok := obj["dynamic"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else {
			m.Dynamic = types.BoolNull()
		}
	} else {
		m.Dynamic = types.BoolNull()
	}
	if v, ok := obj["ip-address"]; ok {
		_ = v
		if v != "" {
			m.IPAddress = types.StringValue(v)
		} else {
			m.IPAddress = types.StringNull()
		}
	} else {
		m.IPAddress = types.StringNull()
	}
	if v, ok := obj["mac-address"]; ok {
		_ = v
		if v != "" {
			m.MACAddress = types.StringValue(v)
		} else {
			m.MACAddress = types.StringNull()
		}
	} else {
		m.MACAddress = types.StringNull()
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
	if v, ok := obj["rate-limited"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.RateLimited = types.BoolValue(b)
		} else {
			m.RateLimited = types.BoolNull()
		}
	} else {
		m.RateLimited = types.BoolNull()
	}
	if v, ok := obj["rate-up-down"]; ok {
		_ = v
		if v != "" {
			m.RateUpDown = types.StringValue(v)
		} else {
			m.RateUpDown = types.StringNull()
		}
	} else {
		m.RateUpDown = types.StringNull()
	}
	if v, ok := obj["reset-counters"]; ok {
		_ = v
		if v != "" {
			m.ResetCounters = types.StringValue(v)
		} else {
			m.ResetCounters = types.StringNull()
		}
	} else {
		m.ResetCounters = types.StringNull()
	}
	if v, ok := obj["user"]; ok {
		_ = v
		if v != "" {
			m.User = types.StringValue(v)
		} else {
			m.User = types.StringNull()
		}
	} else {
		m.User = types.StringNull()
	}
}
