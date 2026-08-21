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
	_ resource.Resource                = &ConsoleSettingsResource{}
	_ resource.ResourceWithImportState = &ConsoleSettingsResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type ConsoleSettingsResource struct {
	reg *client.Registry
}

type ConsoleSettingsModel struct {
	ID              types.String `tfsdk:"id"`
	LogScriptErrors types.Bool   `tfsdk:"log_script_errors"`
	SanitizeNames   types.Bool   `tfsdk:"sanitize_names"`
	TabWidth        types.Int64  `tfsdk:"tab_width"`
	Router          types.String `tfsdk:"router"`
}

func NewConsoleSettingsResource() resource.Resource { return &ConsoleSettingsResource{} }

func (r *ConsoleSettingsResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_console_settings"
}

func (r *ConsoleSettingsResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *ConsoleSettingsResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/console/settings`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"log_script_errors": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"sanitize_names": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"tab_width": schema.Int64Attribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *ConsoleSettingsResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan ConsoleSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	consoleSettingsUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ConsoleSettingsResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan ConsoleSettingsModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state ConsoleSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	consoleSettingsUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *ConsoleSettingsResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state ConsoleSettingsModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/console/settings")
	if err != nil {
		resp.Diagnostics.AddError("Read /console/settings failed", err.Error())
		return
	}
	consoleSettingsApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/console/settings", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *ConsoleSettingsResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *ConsoleSettingsResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/console/settings" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/console/settings", types.StringValue(routerName))))...)
}

func consoleSettingsUpsert(ctx context.Context, reg *client.Registry, plan, state *ConsoleSettingsModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.LogScriptErrors.IsNull() || plan.LogScriptErrors.IsUnknown()) && (state == nil || !plan.LogScriptErrors.Equal(state.LogScriptErrors)) {
		body["log-script-errors"] = client.FormatBool(plan.LogScriptErrors.ValueBool())
	}
	if !(plan.SanitizeNames.IsNull() || plan.SanitizeNames.IsUnknown()) && (state == nil || !plan.SanitizeNames.Equal(state.SanitizeNames)) {
		body["sanitize-names"] = client.FormatBool(plan.SanitizeNames.ValueBool())
	}
	if !(plan.TabWidth.IsNull() || plan.TabWidth.IsUnknown()) && (state == nil || !plan.TabWidth.Equal(state.TabWidth)) {
		body["tab-width"] = client.FormatInt64(plan.TabWidth.ValueInt64())
	}
	obj, err := c.SetSingleton(ctx, "/console/settings", body)
	if err != nil {
		diags.AddError("Upsert /console/settings failed", err.Error())
		return
	}
	consoleSettingsApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/console/settings", plan.Router))
}

func consoleSettingsApply(ctx context.Context, obj client.Object, m *ConsoleSettingsModel) {
	_ = ctx
	if v, ok := obj["log-script-errors"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.LogScriptErrors = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.LogScriptErrors = types.BoolValue(true)
		} else {
			m.LogScriptErrors = types.BoolNull()
		}
	}
	if v, ok := obj["sanitize-names"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.SanitizeNames = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.SanitizeNames = types.BoolValue(true)
		} else {
			m.SanitizeNames = types.BoolNull()
		}
	}
	if v, ok := obj["tab-width"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.TabWidth = types.Int64Value(n)
		} else {
			m.TabWidth = types.Int64Null()
		}
	}
}
