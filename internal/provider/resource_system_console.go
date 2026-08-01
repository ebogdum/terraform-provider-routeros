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
	_ resource.Resource                = &SystemConsoleResource{}
	_ resource.ResourceWithImportState = &SystemConsoleResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemConsoleResource struct {
	reg *client.Registry
}

type SystemConsoleModel struct {
	ID       types.String `tfsdk:"id"`
	Channel  types.Int64  `tfsdk:"channel"`
	Default  types.Bool   `tfsdk:"default"`
	Disabled types.Bool   `tfsdk:"disabled"`
	Free     types.Bool   `tfsdk:"free"`
	Port     types.String `tfsdk:"port"`
	Term     types.String `tfsdk:"term"`
	Used     types.Bool   `tfsdk:"used"`
	Vc       types.Int64  `tfsdk:"vc"`
	Vcno     types.Int64  `tfsdk:"vcno"`
	Wedged   types.Bool   `tfsdk:"wedged"`
	Router   types.String `tfsdk:"router"`
}

func NewSystemConsoleResource() resource.Resource { return &SystemConsoleResource{} }

func (r *SystemConsoleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_console"
}

func (r *SystemConsoleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *SystemConsoleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Active console sessions — RouterOS-managed; PUT EOFs because the endpoint isn't add-able.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"channel": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"default": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"free": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"term": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"used": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"vc": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"vcno": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"wedged": schema.BoolAttribute{
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

func (r *SystemConsoleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemConsoleModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Channel.IsNull() || plan.Channel.IsUnknown()) {
		body["channel"] = client.FormatInt64(plan.Channel.ValueInt64())
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.Term.IsNull() || plan.Term.IsUnknown()) {
		body["term"] = plan.Term.ValueString()
	}
	obj, err := c.Add(ctx, "/system/console", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /system/console failed", err.Error())
		return
	}
	systemConsoleApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemConsoleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemConsoleModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/system/console", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /system/console failed", err.Error())
		return
	}
	systemConsoleApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemConsoleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemConsoleModel
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
	if !plan.Channel.Equal(state.Channel) && !plan.Channel.IsUnknown() {
		body["channel"] = client.FormatInt64(plan.Channel.ValueInt64())
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.Term.Equal(state.Term) && !plan.Term.IsUnknown() {
		body["term"] = plan.Term.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/system/console", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /system/console failed", err.Error())
			return
		}
		systemConsoleApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemConsoleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemConsoleModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/system/console", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /system/console failed", err.Error())
	}
}

func (r *SystemConsoleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := systemConsoleLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /system/console matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// systemConsoleLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func systemConsoleLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/system/console", id)
}

func systemConsoleApply(ctx context.Context, obj client.Object, m *SystemConsoleModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["channel"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Channel = types.Int64Value(n)
		} else {
			m.Channel = types.Int64Null()
		}
	} else {
		m.Channel = types.Int64Null()
	}
	if v, ok := obj["default"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else {
			m.Default = types.BoolNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["free"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Free = types.BoolValue(b)
		} else {
			m.Free = types.BoolNull()
		}
	}
	if v, ok := obj["port"]; ok {
		if v != "" {
			m.Port = types.StringValue(v)
		} else {
			m.Port = types.StringNull()
		}
	}
	if v, ok := obj["term"]; ok {
		if v != "" {
			m.Term = types.StringValue(v)
		} else {
			m.Term = types.StringNull()
		}
	}
	if v, ok := obj["used"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Used = types.BoolValue(b)
		} else {
			m.Used = types.BoolNull()
		}
	}
	if v, ok := obj["vc"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Vc = types.Int64Value(n)
		} else {
			m.Vc = types.Int64Null()
		}
	} else {
		m.Vc = types.Int64Null()
	}
	if v, ok := obj["vcno"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Vcno = types.Int64Value(n)
		} else {
			m.Vcno = types.Int64Null()
		}
	} else {
		m.Vcno = types.Int64Null()
	}
	if v, ok := obj["wedged"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Wedged = types.BoolValue(b)
		} else {
			m.Wedged = types.BoolNull()
		}
	}
}
