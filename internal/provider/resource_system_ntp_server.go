package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &SystemNTPServerResource{}
	_ resource.ResourceWithImportState = &SystemNTPServerResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type SystemNTPServerResource struct {
	reg *client.Registry
}

type SystemNTPServerModel struct {
	ID                 types.String `tfsdk:"id"`
	AuthKey            types.String `tfsdk:"auth_key"`
	Broadcast          types.Bool   `tfsdk:"broadcast"`
	BroadcastAddresses types.String `tfsdk:"broadcast_addresses"`
	Enabled            types.Bool   `tfsdk:"enabled"`
	LocalClockStratum  types.Int64  `tfsdk:"local_clock_stratum"`
	Manycast           types.Bool   `tfsdk:"manycast"`
	Multicast          types.Bool   `tfsdk:"multicast"`
	UseLocalClock      types.Bool   `tfsdk:"use_local_clock"`
	Vrf                types.String `tfsdk:"vrf"`
	Router             types.String `tfsdk:"router"`
}

func NewSystemNTPServerResource() resource.Resource { return &SystemNTPServerResource{} }

func (r *SystemNTPServerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_ntp_server"
}

func (r *SystemNTPServerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemNTPServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/system/ntp/server`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"auth_key": schema.StringAttribute{Optional: true, Computed: true, Sensitive: true,
				Description: "",
			},
			"broadcast": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"broadcast_addresses": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"enabled": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"local_clock_stratum": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"manycast": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"multicast": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"use_local_clock": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"vrf": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemNTPServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemNTPServerModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemNTPServerUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemNTPServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SystemNTPServerModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state SystemNTPServerModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemNTPServerUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemNTPServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemNTPServerModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/system/ntp/server")
	if err != nil {
		resp.Diagnostics.AddError("Read /system/ntp/server failed", err.Error())
		return
	}
	systemNTPServerApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/system/ntp/server", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemNTPServerResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *SystemNTPServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/system/ntp/server" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/system/ntp/server", types.StringValue(routerName))))...)
}

func systemNTPServerUpsert(ctx context.Context, reg *client.Registry, plan, state *SystemNTPServerModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AuthKey.IsNull() || plan.AuthKey.IsUnknown()) && (state == nil || !plan.AuthKey.Equal(state.AuthKey)) {
		body["auth-key"] = plan.AuthKey.ValueString()
	}
	if !(plan.Broadcast.IsNull() || plan.Broadcast.IsUnknown()) && (state == nil || !plan.Broadcast.Equal(state.Broadcast)) {
		body["broadcast"] = client.FormatBool(plan.Broadcast.ValueBool())
	}
	if !(plan.BroadcastAddresses.IsNull() || plan.BroadcastAddresses.IsUnknown()) && (state == nil || !plan.BroadcastAddresses.Equal(state.BroadcastAddresses)) {
		body["broadcast-addresses"] = plan.BroadcastAddresses.ValueString()
	}
	if !(plan.Enabled.IsNull() || plan.Enabled.IsUnknown()) && (state == nil || !plan.Enabled.Equal(state.Enabled)) {
		body["enabled"] = client.FormatBool(plan.Enabled.ValueBool())
	}
	if !(plan.LocalClockStratum.IsNull() || plan.LocalClockStratum.IsUnknown()) && (state == nil || !plan.LocalClockStratum.Equal(state.LocalClockStratum)) {
		body["local-clock-stratum"] = client.FormatInt64(plan.LocalClockStratum.ValueInt64())
	}
	if !(plan.Manycast.IsNull() || plan.Manycast.IsUnknown()) && (state == nil || !plan.Manycast.Equal(state.Manycast)) {
		body["manycast"] = client.FormatBool(plan.Manycast.ValueBool())
	}
	if !(plan.Multicast.IsNull() || plan.Multicast.IsUnknown()) && (state == nil || !plan.Multicast.Equal(state.Multicast)) {
		body["multicast"] = client.FormatBool(plan.Multicast.ValueBool())
	}
	if !(plan.UseLocalClock.IsNull() || plan.UseLocalClock.IsUnknown()) && (state == nil || !plan.UseLocalClock.Equal(state.UseLocalClock)) {
		body["use-local-clock"] = client.FormatBool(plan.UseLocalClock.ValueBool())
	}
	if !(plan.Vrf.IsNull() || plan.Vrf.IsUnknown()) && (state == nil || !plan.Vrf.Equal(state.Vrf)) {
		body["vrf"] = plan.Vrf.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/system/ntp/server", body)
	if err != nil {
		diags.AddError("Upsert /system/ntp/server failed", err.Error())
		return
	}
	systemNTPServerApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/system/ntp/server", plan.Router))
}

func systemNTPServerApply(ctx context.Context, obj client.Object, m *SystemNTPServerModel) {
	_ = ctx
	if v, ok := obj["auth-key"]; ok && v != "" {
		_ = v
		if v != "" {
			m.AuthKey = types.StringValue(v)
		} else {
			m.AuthKey = types.StringNull()
		}
	} else if m.AuthKey.IsUnknown() {
		m.AuthKey = types.StringNull()
	}
	if v, ok := obj["broadcast"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Broadcast = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Broadcast = types.BoolValue(true)
		} else {
			m.Broadcast = types.BoolNull()
		}
	}
	if v, ok := obj["broadcast-addresses"]; ok {
		_ = v
		if v != "" {
			m.BroadcastAddresses = types.StringValue(v)
		} else {
			m.BroadcastAddresses = types.StringNull()
		}
	}
	if v, ok := obj["enabled"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Enabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Enabled = types.BoolValue(true)
		} else {
			m.Enabled = types.BoolNull()
		}
	}
	if v, ok := obj["local-clock-stratum"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.LocalClockStratum = types.Int64Value(n)
		} else {
			m.LocalClockStratum = types.Int64Null()
		}
	}
	if v, ok := obj["manycast"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Manycast = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Manycast = types.BoolValue(true)
		} else {
			m.Manycast = types.BoolNull()
		}
	}
	if v, ok := obj["multicast"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Multicast = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Multicast = types.BoolValue(true)
		} else {
			m.Multicast = types.BoolNull()
		}
	}
	if v, ok := obj["use-local-clock"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.UseLocalClock = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.UseLocalClock = types.BoolValue(true)
		} else {
			m.UseLocalClock = types.BoolNull()
		}
	}
	if v, ok := obj["vrf"]; ok {
		_ = v
		if v != "" {
			m.Vrf = types.StringValue(v)
		} else {
			m.Vrf = types.StringNull()
		}
	}
}
