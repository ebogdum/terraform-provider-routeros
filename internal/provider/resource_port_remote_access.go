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
	_ resource.Resource                = &PortRemoteAccessResource{}
	_ resource.ResourceWithImportState = &PortRemoteAccessResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type PortRemoteAccessResource struct {
	reg *client.Registry
}

type PortRemoteAccessModel struct {
	ID              types.String `tfsdk:"id"`
	RemoteAddresses types.String `tfsdk:"remote_addresses"`
	LogFile         types.String `tfsdk:"log_file"`
	IpPort          types.String `tfsdk:"ip_port"`
	Channel         types.String `tfsdk:"channel"`
	Comment         types.String `tfsdk:"comment"`
	Disabled        types.Bool   `tfsdk:"disabled"`
	LocalAddress    types.String `tfsdk:"local_address"`
	Port            types.String `tfsdk:"port"`
	Protocol        types.String `tfsdk:"protocol"`
	Router          types.String `tfsdk:"router"`
}

func NewPortRemoteAccessResource() resource.Resource { return &PortRemoteAccessResource{} }

func (r *PortRemoteAccessResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_port_remote_access"
}

func (r *PortRemoteAccessResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *PortRemoteAccessResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/port/remote-access`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"remote_addresses": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `remote-addresses`.",
			},
			"log_file": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `log-file`.",
			},
			"ip_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ip-port`.",
			},
			"channel": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
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
			"local_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"protocol": schema.StringAttribute{
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

func (r *PortRemoteAccessResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan PortRemoteAccessModel
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
		body["channel"] = plan.Channel.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.LocalAddress.IsNull() || plan.LocalAddress.IsUnknown()) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = plan.Port.ValueString()
	}
	if !(plan.Protocol.IsNull() || plan.Protocol.IsUnknown()) {
		body["protocol"] = plan.Protocol.ValueString()
	}
	if !(plan.IpPort.IsNull() || plan.IpPort.IsUnknown()) {
		body["ip-port"] = plan.IpPort.ValueString()
	}
	if !(plan.LogFile.IsNull() || plan.LogFile.IsUnknown()) {
		body["log-file"] = plan.LogFile.ValueString()
	}
	if !(plan.RemoteAddresses.IsNull() || plan.RemoteAddresses.IsUnknown()) {
		body["remote-addresses"] = plan.RemoteAddresses.ValueString()
	}
	obj, err := c.Add(ctx, "/port/remote-access", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /port/remote-access failed", err.Error())
		return
	}
	portRemoteAccessApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PortRemoteAccessResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state PortRemoteAccessModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/port/remote-access", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /port/remote-access failed", err.Error())
		return
	}
	portRemoteAccessApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *PortRemoteAccessResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state PortRemoteAccessModel
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
		body["channel"] = plan.Channel.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.LocalAddress.Equal(state.LocalAddress) && !plan.LocalAddress.IsUnknown() {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = plan.Port.ValueString()
	}
	if !plan.Protocol.Equal(state.Protocol) && !plan.Protocol.IsUnknown() {
		body["protocol"] = plan.Protocol.ValueString()
	}
	if !plan.IpPort.Equal(state.IpPort) && !plan.IpPort.IsUnknown() {
		body["ip-port"] = plan.IpPort.ValueString()
	}
	if !plan.LogFile.Equal(state.LogFile) && !plan.LogFile.IsUnknown() {
		body["log-file"] = plan.LogFile.ValueString()
	}
	if !plan.RemoteAddresses.Equal(state.RemoteAddresses) && !plan.RemoteAddresses.IsUnknown() {
		body["remote-addresses"] = plan.RemoteAddresses.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/port/remote-access", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /port/remote-access failed", err.Error())
			return
		}
		portRemoteAccessApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *PortRemoteAccessResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state PortRemoteAccessModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/port/remote-access", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /port/remote-access failed", err.Error())
	}
}

func (r *PortRemoteAccessResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := portRemoteAccessLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /port/remote-access matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// portRemoteAccessLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func portRemoteAccessLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/port/remote-access", id)
}

func portRemoteAccessApply(ctx context.Context, obj client.Object, m *PortRemoteAccessModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["remote-addresses"]; ok && v != "" {
		m.RemoteAddresses = types.StringValue(v)
	} else {
		m.RemoteAddresses = types.StringNull()
	}
	if v, ok := obj["log-file"]; ok && v != "" {
		m.LogFile = types.StringValue(v)
	} else {
		m.LogFile = types.StringNull()
	}
	if v, ok := obj["ip-port"]; ok && v != "" {
		m.IpPort = types.StringValue(v)
	} else {
		m.IpPort = types.StringNull()
	}
	if v, ok := obj["channel"]; ok {
		_ = v
		if v != "" {
			m.Channel = types.StringValue(v)
		} else {
			m.Channel = types.StringNull()
		}
	} else {
		m.Channel = types.StringNull()
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
	if v, ok := obj["local-address"]; ok {
		_ = v
		if v != "" {
			m.LocalAddress = types.StringValue(v)
		} else {
			m.LocalAddress = types.StringNull()
		}
	} else {
		m.LocalAddress = types.StringNull()
	}
	if v, ok := obj["port"]; ok {
		_ = v
		if v != "" {
			m.Port = types.StringValue(v)
		} else {
			m.Port = types.StringNull()
		}
	} else {
		m.Port = types.StringNull()
	}
	if v, ok := obj["protocol"]; ok {
		_ = v
		if v != "" {
			m.Protocol = types.StringValue(v)
		} else {
			m.Protocol = types.StringNull()
		}
	} else {
		m.Protocol = types.StringNull()
	}
}
