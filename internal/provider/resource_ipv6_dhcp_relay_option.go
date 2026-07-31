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
	_ resource.Resource                = &IPv6DHCPRelayOptionResource{}
	_ resource.ResourceWithImportState = &IPv6DHCPRelayOptionResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPv6DHCPRelayOptionResource struct {
	reg *client.Registry
}

type IPv6DHCPRelayOptionModel struct {
	ID                 types.String `tfsdk:"id"`
	Code               types.Int64  `tfsdk:"code"`
	Name               types.String `tfsdk:"name"`
	OnlyIfMACAvailable types.Bool   `tfsdk:"only_if_mac_available"`
	Value              types.String `tfsdk:"value"`
	Router             types.String `tfsdk:"router"`
}

func NewIPv6DHCPRelayOptionResource() resource.Resource { return &IPv6DHCPRelayOptionResource{} }

func (r *IPv6DHCPRelayOptionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv6_dhcp_relay_option"
}

func (r *IPv6DHCPRelayOptionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPv6DHCPRelayOptionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ipv6/dhcp-relay/option`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"code": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `code`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `name`.",
			},
			"only_if_mac_available": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `only-if-mac-available`.",
			},
			"value": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `value`.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPv6DHCPRelayOptionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPv6DHCPRelayOptionModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Code.IsNull() || plan.Code.IsUnknown()) {
		body["code"] = client.FormatInt64(plan.Code.ValueInt64())
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.OnlyIfMACAvailable.IsNull() || plan.OnlyIfMACAvailable.IsUnknown()) {
		body["only-if-mac-available"] = client.FormatBool(plan.OnlyIfMACAvailable.ValueBool())
	}
	if !(plan.Value.IsNull() || plan.Value.IsUnknown()) {
		body["value"] = plan.Value.ValueString()
	}
	obj, err := c.Add(ctx, "/ipv6/dhcp-relay/option", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ipv6/dhcp-relay/option failed", err.Error())
		return
	}
	iPv6DHCPRelayOptionApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPv6DHCPRelayOptionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPv6DHCPRelayOptionModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ipv6/dhcp-relay/option", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ipv6/dhcp-relay/option failed", err.Error())
		return
	}
	iPv6DHCPRelayOptionApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPv6DHCPRelayOptionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPv6DHCPRelayOptionModel
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
	if !plan.Code.Equal(state.Code) && !plan.Code.IsUnknown() {
		body["code"] = client.FormatInt64(plan.Code.ValueInt64())
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.OnlyIfMACAvailable.Equal(state.OnlyIfMACAvailable) && !plan.OnlyIfMACAvailable.IsUnknown() {
		body["only-if-mac-available"] = client.FormatBool(plan.OnlyIfMACAvailable.ValueBool())
	}
	if !plan.Value.Equal(state.Value) && !plan.Value.IsUnknown() {
		body["value"] = plan.Value.ValueString()
	}
	var obj client.Object
	var err error
	if len(body) > 0 {
		obj, err = c.Set(ctx, "/ipv6/dhcp-relay/option", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ipv6/dhcp-relay/option failed", err.Error())
			return
		}
	} else {
		obj, err = c.GetByID(ctx, "/ipv6/dhcp-relay/option", state.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Read /ipv6/dhcp-relay/option failed", err.Error())
			return
		}
	}
	iPv6DHCPRelayOptionApply(ctx, obj, &plan)
	plan.ID = state.ID
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPv6DHCPRelayOptionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPv6DHCPRelayOptionModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ipv6/dhcp-relay/option", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ipv6/dhcp-relay/option failed", err.Error())
	}
}

func (r *IPv6DHCPRelayOptionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := lookupByNaturalKey(ctx, c, "/ipv6/dhcp-relay/option", id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ipv6/dhcp-relay/option matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

func iPv6DHCPRelayOptionApply(ctx context.Context, obj client.Object, m *IPv6DHCPRelayOptionModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["code"]; ok {
		if n, err := client.ParseInt64(v); err == nil {
			m.Code = types.Int64Value(n)
		} else {
			m.Code = types.Int64Null()
		}
	} else {
		m.Code = types.Int64Null()
	}
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["only-if-mac-available"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.OnlyIfMACAvailable = types.BoolValue(b)
		} else {
			m.OnlyIfMACAvailable = types.BoolNull()
		}
	} else {
		m.OnlyIfMACAvailable = types.BoolNull()
	}
	if v, ok := obj["value"]; ok && v != "" {
		m.Value = types.StringValue(v)
	} else {
		m.Value = types.StringNull()
	}
}
