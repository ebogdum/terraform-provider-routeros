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
	_ resource.Resource                = &ToolSmsResource{}
	_ resource.ResourceWithImportState = &ToolSmsResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type ToolSmsResource struct {
	reg *client.Registry
}

type ToolSmsModel struct {
	ID                     types.String `tfsdk:"id"`
	AllowedNumber          types.String `tfsdk:"allowed_number"`
	Channel                types.Int64  `tfsdk:"channel"`
	LastUssd               types.String `tfsdk:"last_ussd"`
	Polling                types.Bool   `tfsdk:"polling"`
	Port                   types.String `tfsdk:"port"`
	ReceiveEnabled         types.Bool   `tfsdk:"receive_enabled"`
	RemoveSentSmsAfterSend types.Bool   `tfsdk:"remove_sent_sms_after_send"`
	Secret                 types.String `tfsdk:"secret"`
	SimPin                 types.String `tfsdk:"sim_pin"`
	SmsStorage             types.String `tfsdk:"sms_storage"`
	Status                 types.String `tfsdk:"status"`
	Router                 types.String `tfsdk:"router"`
}

func NewToolSmsResource() resource.Resource { return &ToolSmsResource{} }

func (r *ToolSmsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_tool_sms"
}

func (r *ToolSmsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ToolSmsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/tool/sms`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"allowed_number": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"channel": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"last_ussd": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"polling": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"port": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"receive_enabled": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"remove_sent_sms_after_send": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"secret": schema.StringAttribute{Optional: true, Computed: true, Sensitive: true,
				Description: "",
			},
			"sim_pin": schema.StringAttribute{Optional: true, Computed: true, Sensitive: true,
				Description: "",
			},
			"sms_storage": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"status": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *ToolSmsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ToolSmsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolSmsUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolSmsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ToolSmsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state ToolSmsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	toolSmsUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ToolSmsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ToolSmsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/tool/sms")
	if err != nil {
		resp.Diagnostics.AddError("Read /tool/sms failed", err.Error())
		return
	}
	toolSmsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/tool/sms", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ToolSmsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *ToolSmsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/tool/sms" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/tool/sms", types.StringValue(routerName))))...)
}

func toolSmsUpsert(ctx context.Context, reg *client.Registry, plan, state *ToolSmsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AllowedNumber.IsNull() || plan.AllowedNumber.IsUnknown()) && (state == nil || !plan.AllowedNumber.Equal(state.AllowedNumber)) {
		body["allowed-number"] = plan.AllowedNumber.ValueString()
	}
	if !(plan.Channel.IsNull() || plan.Channel.IsUnknown()) && (state == nil || !plan.Channel.Equal(state.Channel)) {
		body["channel"] = client.FormatInt64(plan.Channel.ValueInt64())
	}
	if !(plan.Polling.IsNull() || plan.Polling.IsUnknown()) && (state == nil || !plan.Polling.Equal(state.Polling)) {
		body["polling"] = client.FormatBool(plan.Polling.ValueBool())
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) && (state == nil || !plan.Port.Equal(state.Port)) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.ReceiveEnabled.IsNull() || plan.ReceiveEnabled.IsUnknown()) && (state == nil || !plan.ReceiveEnabled.Equal(state.ReceiveEnabled)) {
		body["receive-enabled"] = client.FormatBool(plan.ReceiveEnabled.ValueBool())
	}
	if !(plan.RemoveSentSmsAfterSend.IsNull() || plan.RemoveSentSmsAfterSend.IsUnknown()) && (state == nil || !plan.RemoveSentSmsAfterSend.Equal(state.RemoveSentSmsAfterSend)) {
		body["remove-sent-sms-after-send"] = client.FormatBool(plan.RemoveSentSmsAfterSend.ValueBool())
	}
	if !(plan.Secret.IsNull() || plan.Secret.IsUnknown()) && (state == nil || !plan.Secret.Equal(state.Secret)) {
		body["secret"] = plan.Secret.ValueString()
	}
	if !(plan.SimPin.IsNull() || plan.SimPin.IsUnknown()) && (state == nil || !plan.SimPin.Equal(state.SimPin)) {
		body["sim-pin"] = plan.SimPin.ValueString()
	}
	if !(plan.SmsStorage.IsNull() || plan.SmsStorage.IsUnknown()) && (state == nil || !plan.SmsStorage.Equal(state.SmsStorage)) {
		body["sms-storage"] = plan.SmsStorage.ValueString()
	}
	obj, err := c.SetSingleton(ctx, "/tool/sms", body)
	if err != nil {
		diags.AddError("Upsert /tool/sms failed", err.Error())
		return
	}
	toolSmsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/tool/sms", plan.Router))
}

func toolSmsApply(ctx context.Context, obj client.Object, m *ToolSmsModel) {
	_ = ctx
	if v, ok := obj["allowed-number"]; ok {
		_ = v
		if v != "" {
			m.AllowedNumber = types.StringValue(v)
		} else {
			m.AllowedNumber = types.StringNull()
		}
	}
	if v, ok := obj["channel"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Channel = types.Int64Value(n)
		} else {
			m.Channel = types.Int64Null()
		}
	}
	if v, ok := obj["last-ussd"]; ok {
		_ = v
		if v != "" {
			m.LastUssd = types.StringValue(v)
		} else {
			m.LastUssd = types.StringNull()
		}
	}
	if v, ok := obj["polling"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Polling = types.BoolValue(b)
		} else {
			m.Polling = types.BoolNull()
		}
	}
	if v, ok := obj["port"]; ok {
		_ = v
		if v != "" {
			m.Port = types.StringValue(v)
		} else {
			m.Port = types.StringNull()
		}
	}
	if v, ok := obj["receive-enabled"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.ReceiveEnabled = types.BoolValue(b)
		} else {
			m.ReceiveEnabled = types.BoolNull()
		}
	}
	if v, ok := obj["remove-sent-sms-after-send"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.RemoveSentSmsAfterSend = types.BoolValue(b)
		} else {
			m.RemoveSentSmsAfterSend = types.BoolNull()
		}
	}
	if v, ok := obj["secret"]; ok && v != "" {
		_ = v
		if v != "" {
			m.Secret = types.StringValue(v)
		} else {
			m.Secret = types.StringNull()
		}
	} else if m.Secret.IsUnknown() {
		m.Secret = types.StringNull()
	}
	if v, ok := obj["sim-pin"]; ok && v != "" {
		_ = v
		if v != "" {
			m.SimPin = types.StringValue(v)
		} else {
			m.SimPin = types.StringNull()
		}
	} else if m.SimPin.IsUnknown() {
		m.SimPin = types.StringNull()
	}
	if v, ok := obj["sms-storage"]; ok {
		_ = v
		if v != "" {
			m.SmsStorage = types.StringValue(v)
		} else {
			m.SmsStorage = types.StringNull()
		}
	}
	if v, ok := obj["status"]; ok {
		_ = v
		if v != "" {
			m.Status = types.StringValue(v)
		} else {
			m.Status = types.StringNull()
		}
	}
}
