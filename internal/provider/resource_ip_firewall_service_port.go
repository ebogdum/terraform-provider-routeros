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
	_ resource.Resource                = &IPFirewallServicePortResource{}
	_ resource.ResourceWithImportState = &IPFirewallServicePortResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPFirewallServicePortResource struct {
	reg *client.Registry
}

type IPFirewallServicePortModel struct {
	ID             types.String `tfsdk:"id"`
	Disabled       types.Bool   `tfsdk:"disabled"`
	Name           types.String `tfsdk:"name"`
	Ports          types.String `tfsdk:"ports"`
	SipDirectMedia types.Bool   `tfsdk:"sip_direct_media"`
	SipTimeout     types.String `tfsdk:"sip_timeout"`
	Router         types.String `tfsdk:"router"`
}

func NewIPFirewallServicePortResource() resource.Resource { return &IPFirewallServicePortResource{} }

func (r *IPFirewallServicePortResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_firewall_service_port"
}

func (r *IPFirewallServicePortResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPFirewallServicePortResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/firewall/service-port`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `disabled`.",
			},
			"name": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "RouterOS `name`.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"ports": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ports`.",
			},
			"sip_direct_media": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `sip-direct-media`.",
			},
			"sip_timeout": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `sip-timeout`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPFirewallServicePortResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPFirewallServicePortModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	rows, err := c.List(ctx, "/ip/firewall/service-port")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/firewall/service-port failed", err.Error())
		return
	}
	want := plan.Name.ValueString()
	var id string
	for _, row := range rows {
		if row["name"] == want {
			id = row[".id"]
			break
		}
	}
	if id == "" {
		resp.Diagnostics.AddError("Unknown /ip/firewall/service-port "+want, fmt.Sprintf("/ip/firewall/service-port has no row with name=%q; this is a fixed-row menu and rows cannot be added", want))
		return
	}
	body := client.Object{}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Ports.IsNull() || plan.Ports.IsUnknown()) {
		body["ports"] = plan.Ports.ValueString()
	}
	if !(plan.SipDirectMedia.IsNull() || plan.SipDirectMedia.IsUnknown()) {
		body["sip-direct-media"] = client.FormatBool(plan.SipDirectMedia.ValueBool())
	}
	if !(plan.SipTimeout.IsNull() || plan.SipTimeout.IsUnknown()) {
		body["sip-timeout"] = plan.SipTimeout.ValueString()
	}
	obj, err := c.Set(ctx, "/ip/firewall/service-port", id, body)
	if err != nil {
		resp.Diagnostics.AddError("Adopt /ip/firewall/service-port failed", err.Error())
		return
	}
	iPFirewallServicePortApply(ctx, obj, &plan)
	plan.ID = types.StringValue(id)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPFirewallServicePortResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPFirewallServicePortModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/firewall/service-port", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/firewall/service-port failed", err.Error())
		return
	}
	iPFirewallServicePortApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPFirewallServicePortResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPFirewallServicePortModel
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
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Ports.Equal(state.Ports) && !plan.Ports.IsUnknown() {
		body["ports"] = plan.Ports.ValueString()
	}
	if !plan.SipDirectMedia.Equal(state.SipDirectMedia) && !plan.SipDirectMedia.IsUnknown() {
		body["sip-direct-media"] = client.FormatBool(plan.SipDirectMedia.ValueBool())
	}
	if !plan.SipTimeout.Equal(state.SipTimeout) && !plan.SipTimeout.IsUnknown() {
		body["sip-timeout"] = plan.SipTimeout.ValueString()
	}
	var obj client.Object
	var err error
	if len(body) > 0 {
		obj, err = c.Set(ctx, "/ip/firewall/service-port", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/firewall/service-port failed", err.Error())
			return
		}
	} else {
		obj, err = c.GetByID(ctx, "/ip/firewall/service-port", state.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Read /ip/firewall/service-port failed", err.Error())
			return
		}
	}
	iPFirewallServicePortApply(ctx, obj, &plan)
	plan.ID = state.ID
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPFirewallServicePortResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	// Fixed-row menu: rows cannot be removed. Drop from state; the row keeps
	// its last-applied settings, matching /ip/service semantics.
	_ = ctx
	_ = req
	_ = resp
}

func (r *IPFirewallServicePortResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := lookupByNaturalKey(ctx, c, "/ip/firewall/service-port", id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/firewall/service-port matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

func iPFirewallServicePortApply(ctx context.Context, obj client.Object, m *IPFirewallServicePortModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	} else {
		m.Disabled = types.BoolNull()
	}
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["ports"]; ok && v != "" {
		m.Ports = types.StringValue(v)
	} else {
		m.Ports = types.StringNull()
	}
	if v, ok := obj["sip-direct-media"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SipDirectMedia = types.BoolValue(b)
		} else {
			m.SipDirectMedia = types.BoolNull()
		}
	} else {
		m.SipDirectMedia = types.BoolNull()
	}
	if v, ok := obj["sip-timeout"]; ok && v != "" {
		m.SipTimeout = types.StringValue(v)
	} else {
		m.SipTimeout = types.StringNull()
	}
}
