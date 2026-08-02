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
	_ resource.Resource                = &SystemNoteResource{}
	_ resource.ResourceWithImportState = &SystemNoteResource{}
	_                                  = path.Root
	_                                  = fmt.Sprintf
)

type SystemNoteResource struct {
	reg *client.Registry
}

type SystemNoteModel struct {
	ID             types.String `tfsdk:"id"`
	Note           types.String `tfsdk:"note"`
	ShowAtCliLogin types.Bool   `tfsdk:"show_at_cli_login"`
	ShowAtLogin    types.Bool   `tfsdk:"show_at_login"`
	Router         types.String `tfsdk:"router"`
}

func NewSystemNoteResource() resource.Resource { return &SystemNoteResource{} }

func (r *SystemNoteResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_note"
}

func (r *SystemNoteResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemNoteResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/system/note`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "Stable identifier (the singleton's menu path, optionally namespaced by router).",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"note": schema.StringAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"show_at_cli_login": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"show_at_login": schema.BoolAttribute{Optional: true, Computed: true,
				Description: "",
			},
			"router": schema.StringAttribute{Optional: true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemNoteResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemNoteModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemNoteUpsert(ctx, r.reg, &plan, nil, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemNoteResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan SystemNoteModel
	if d := req.Plan.Get(ctx, &plan); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	var state SystemNoteModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	systemNoteUpsert(ctx, r.reg, &plan, &state, &resp.Diagnostics)
	if resp.Diagnostics.HasError() {
		return
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemNoteResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemNoteModel
	if d := req.State.Get(ctx, &state); d.HasError() {
		resp.Diagnostics.Append(d...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetSingleton(ctx, "/system/note")
	if err != nil {
		resp.Diagnostics.AddError("Read /system/note failed", err.Error())
		return
	}
	systemNoteApply(ctx, obj, &state)
	state.ID = types.StringValue(stateIDFor("/system/note", state.Router))
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemNoteResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// Singleton menus aren't removable; just drop the state.
}

func (r *SystemNoteResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import format: "<router>" or empty for default.
	routerName := req.ID
	if routerName == "/system/note" {
		routerName = ""
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(stateIDFor("/system/note", types.StringValue(routerName))))...)
}

func systemNoteUpsert(ctx context.Context, reg *client.Registry, plan, state *SystemNoteModel, diags *diagBuf) {
	c := pickClient(reg, plan.Router, diags)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Note.IsNull() || plan.Note.IsUnknown()) && (state == nil || !plan.Note.Equal(state.Note)) {
		body["note"] = plan.Note.ValueString()
	}
	if !(plan.ShowAtCliLogin.IsNull() || plan.ShowAtCliLogin.IsUnknown()) && (state == nil || !plan.ShowAtCliLogin.Equal(state.ShowAtCliLogin)) {
		body["show-at-cli-login"] = client.FormatBool(plan.ShowAtCliLogin.ValueBool())
	}
	if !(plan.ShowAtLogin.IsNull() || plan.ShowAtLogin.IsUnknown()) && (state == nil || !plan.ShowAtLogin.Equal(state.ShowAtLogin)) {
		body["show-at-login"] = client.FormatBool(plan.ShowAtLogin.ValueBool())
	}
	obj, err := c.SetSingleton(ctx, "/system/note", body)
	if err != nil {
		diags.AddError("Upsert /system/note failed", err.Error())
		return
	}
	systemNoteApply(ctx, obj, plan)
	plan.ID = types.StringValue(stateIDFor("/system/note", plan.Router))
}

func systemNoteApply(ctx context.Context, obj client.Object, m *SystemNoteModel) {
	_ = ctx
	if v, ok := obj["note"]; ok {
		_ = v
		if v != "" {
			m.Note = types.StringValue(v)
		} else {
			m.Note = types.StringNull()
		}
	}
	if v, ok := obj["show-at-cli-login"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.ShowAtCliLogin = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.ShowAtCliLogin = types.BoolValue(true)
		} else {
			m.ShowAtCliLogin = types.BoolNull()
		}
	}
	if v, ok := obj["show-at-login"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.ShowAtLogin = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.ShowAtLogin = types.BoolValue(true)
		} else {
			m.ShowAtLogin = types.BoolNull()
		}
	}
}
