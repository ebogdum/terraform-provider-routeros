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
	_ resource.Resource                = &IPDHCPServerAlertResource{}
	_ resource.ResourceWithImportState = &IPDHCPServerAlertResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPDHCPServerAlertResource struct {
	reg *client.Registry
}

type IPDHCPServerAlertModel struct {
	ID             types.String `tfsdk:"id"`
	ValidServer    types.String `tfsdk:"valid_server"`
	AlertTimeout   types.String `tfsdk:"alert_timeout"`
	Comment        types.String `tfsdk:"comment"`
	Disabled       types.Bool   `tfsdk:"disabled"`
	Interface      types.String `tfsdk:"interface"`
	OnAlert        types.String `tfsdk:"on_alert"`
	ResetAlert     types.String `tfsdk:"reset_alert"`
	UnknownServers types.String `tfsdk:"unknown_servers"`
	ValidServers   types.String `tfsdk:"valid_servers"`
	Router         types.String `tfsdk:"router"`
}

func NewIPDHCPServerAlertResource() resource.Resource { return &IPDHCPServerAlertResource{} }

func (r *IPDHCPServerAlertResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_dhcp_server_alert"
}

func (r *IPDHCPServerAlertResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPDHCPServerAlertResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/dhcp-server/alert`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"valid_server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `valid-server`.",
			},
			"alert_timeout": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsDurationRouterOS()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeDuration()},
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"interface": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"on_alert": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"reset_alert": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"unknown_servers": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"valid_servers": schema.StringAttribute{
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

func (r *IPDHCPServerAlertResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPDHCPServerAlertModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AlertTimeout.IsNull() || plan.AlertTimeout.IsUnknown()) {
		body["alert-timeout"] = plan.AlertTimeout.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.OnAlert.IsNull() || plan.OnAlert.IsUnknown()) {
		body["on-alert"] = plan.OnAlert.ValueString()
	}
	if !(plan.ValidServer.IsNull() || plan.ValidServer.IsUnknown()) {
		body["valid-server"] = plan.ValidServer.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/dhcp-server/alert", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/dhcp-server/alert failed", err.Error())
		return
	}
	iPDHCPServerAlertApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPServerAlertResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPDHCPServerAlertModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/dhcp-server/alert", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/dhcp-server/alert failed", err.Error())
		return
	}
	iPDHCPServerAlertApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPDHCPServerAlertResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPDHCPServerAlertModel
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
	if !plan.AlertTimeout.Equal(state.AlertTimeout) && !plan.AlertTimeout.IsUnknown() {
		body["alert-timeout"] = plan.AlertTimeout.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) && !plan.Interface.IsUnknown() {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.OnAlert.Equal(state.OnAlert) && !plan.OnAlert.IsUnknown() {
		body["on-alert"] = plan.OnAlert.ValueString()
	}
	if !plan.ValidServer.Equal(state.ValidServer) && !plan.ValidServer.IsUnknown() {
		body["valid-server"] = plan.ValidServer.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/dhcp-server/alert", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/dhcp-server/alert failed", err.Error())
			return
		}
		iPDHCPServerAlertApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDHCPServerAlertResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPDHCPServerAlertModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/dhcp-server/alert", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/dhcp-server/alert failed", err.Error())
	}
}

func (r *IPDHCPServerAlertResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPDHCPServerAlertLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/dhcp-server/alert matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPDHCPServerAlertLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPDHCPServerAlertLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/dhcp-server/alert", id)
}

func iPDHCPServerAlertApply(ctx context.Context, obj client.Object, m *IPDHCPServerAlertModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["valid-server"]; ok && v != "" {
		m.ValidServer = types.StringValue(v)
	} else {
		m.ValidServer = types.StringNull()
	}
	if v, ok := obj["alert-timeout"]; ok {
		_ = v
		if v != "" {
			m.AlertTimeout = types.StringValue(v)
		} else {
			m.AlertTimeout = types.StringNull()
		}
	} else {
		m.AlertTimeout = types.StringNull()
	}
	if v, ok := obj["comment"]; ok {
		_ = v
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	} else {
		m.Comment = types.StringNull()
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
	if v, ok := obj["interface"]; ok {
		_ = v
		if v != "" {
			m.Interface = types.StringValue(v)
		} else {
			m.Interface = types.StringNull()
		}
	} else {
		m.Interface = types.StringNull()
	}
	if v, ok := obj["on-alert"]; ok {
		_ = v
		if v != "" {
			m.OnAlert = types.StringValue(v)
		} else {
			m.OnAlert = types.StringNull()
		}
	} else {
		m.OnAlert = types.StringNull()
	}
	if v, ok := obj["reset-alert"]; ok {
		_ = v
		if v != "" {
			m.ResetAlert = types.StringValue(v)
		} else {
			m.ResetAlert = types.StringNull()
		}
	} else {
		m.ResetAlert = types.StringNull()
	}
	if v, ok := obj["unknown-servers"]; ok {
		_ = v
		if v != "" {
			m.UnknownServers = types.StringValue(v)
		} else {
			m.UnknownServers = types.StringNull()
		}
	} else {
		m.UnknownServers = types.StringNull()
	}
	if v, ok := obj["valid-servers"]; ok {
		_ = v
		if v != "" {
			m.ValidServers = types.StringValue(v)
		} else {
			m.ValidServers = types.StringNull()
		}
	} else {
		m.ValidServers = types.StringNull()
	}
}
