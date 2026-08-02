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
	_ resource.Resource                = &IPIpsecPolicyResource{}
	_ resource.ResourceWithImportState = &IPIpsecPolicyResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPIpsecPolicyResource struct {
	reg *client.Registry
}

type IPIpsecPolicyModel struct {
	ID             types.String `tfsdk:"id"`
	Action         types.String `tfsdk:"action"`
	Active         types.Bool   `tfsdk:"active"`
	Comment        types.String `tfsdk:"comment"`
	Default        types.Bool   `tfsdk:"default"`
	Disabled       types.Bool   `tfsdk:"disabled"`
	DstAddress     types.String `tfsdk:"dst_address"`
	DstPort        types.Int64  `tfsdk:"dst_port"`
	Dynamic        types.Bool   `tfsdk:"dynamic"`
	Group          types.String `tfsdk:"group"`
	Invalid        types.Bool   `tfsdk:"invalid"`
	IpsecProtocols types.String `tfsdk:"ipsec_protocols"`
	Level          types.String `tfsdk:"level"`
	Nopeer         types.String `tfsdk:"nopeer"`
	Notemplate     types.String `tfsdk:"notemplate"`
	Peer           types.String `tfsdk:"peer"`
	Ph2Count       types.Int64  `tfsdk:"ph2_count"`
	Ph2State       types.String `tfsdk:"ph2_state"`
	Proposal       types.String `tfsdk:"proposal"`
	Protocol       types.String `tfsdk:"protocol"`
	SaDstAddress   types.String `tfsdk:"sa_dst_address"`
	SaSrcAddress   types.String `tfsdk:"sa_src_address"`
	SrcAddress     types.String `tfsdk:"src_address"`
	SrcPort        types.Int64  `tfsdk:"src_port"`
	Template       types.Bool   `tfsdk:"template"`
	Tunnel         types.Bool   `tfsdk:"tunnel"`
	Router         types.String `tfsdk:"router"`
}

func NewIPIpsecPolicyResource() resource.Resource { return &IPIpsecPolicyResource{} }

func (r *IPIpsecPolicyResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_ipsec_policy"
}

func (r *IPIpsecPolicyResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPIpsecPolicyResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Auto-test requires a typed-reference precondition (e.g. an existing peer,\ninstance, bridge of the specific kind). The current acc-test generator's\ngeneric data.routeros_interface.all lookup can't satisfy these. Use this\nresource manually with explicit references to a precondition resource\nin your config.\n",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"action": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"discard", "none", "encrypt"}...)},
			},
			"active": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
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
			"dst_address": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsCIDR()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeCIDR()},
			},
			"dst_port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"ipsec_protocols": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "ah", "esp"}...)},
			},
			"level": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"use", "require", "unique"}...)},
			},
			"nopeer": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"notemplate": schema.StringAttribute{
				Optional:           true,
				Computed:           true,
				Description:        "",
				DeprecationMessage: "Not a RouterOS REST property (WebFig-only spelling of `template`); RouterOS rejects it. Read-only and ignored on write - use `template`.",
			},
			"peer": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ph2_count": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"ph2_state": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"spawning", "starting", "message-1-received", "message-1-sent", "message-2-received", "message-2-sent", "message-3-received", "message-3-sent", "message-4-received", "established", "expired", "no-phase1", "eap", "crypto", "qkd"}...)},
			},
			"proposal": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"protocol": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"icmp", "igmp", "ggp", "ip-encap", "tcp", "egp", "udp", "ipsec", "all"}...)},
			},
			"sa_dst_address": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"sa_src_address": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"src_address": schema.StringAttribute{
				Optional:      true,
				Computed:      true,
				Description:   "",
				Validators:    []validator.String{schemautil.IsCIDR()},
				PlanModifiers: []planmodifier.String{schemautil.NormalizeCIDR()},
			},
			"src_port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"template": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"tunnel": schema.BoolAttribute{
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

func (r *IPIpsecPolicyResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPIpsecPolicyModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Action.IsNull() || plan.Action.IsUnknown()) {
		body["action"] = plan.Action.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.DstAddress.IsNull() || plan.DstAddress.IsUnknown()) {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !(plan.DstPort.IsNull() || plan.DstPort.IsUnknown()) {
		body["dst-port"] = client.FormatInt64(plan.DstPort.ValueInt64())
	}
	if !(plan.Group.IsNull() || plan.Group.IsUnknown()) {
		body["group"] = plan.Group.ValueString()
	}
	if !(plan.IpsecProtocols.IsNull() || plan.IpsecProtocols.IsUnknown()) {
		body["ipsec-protocols"] = plan.IpsecProtocols.ValueString()
	}
	if !(plan.Level.IsNull() || plan.Level.IsUnknown()) {
		body["level"] = plan.Level.ValueString()
	}
	if !(plan.Peer.IsNull() || plan.Peer.IsUnknown()) {
		body["peer"] = plan.Peer.ValueString()
	}
	if !(plan.Proposal.IsNull() || plan.Proposal.IsUnknown()) {
		body["proposal"] = plan.Proposal.ValueString()
	}
	if !(plan.Protocol.IsNull() || plan.Protocol.IsUnknown()) {
		body["protocol"] = plan.Protocol.ValueString()
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !(plan.SrcPort.IsNull() || plan.SrcPort.IsUnknown()) {
		body["src-port"] = client.FormatInt64(plan.SrcPort.ValueInt64())
	}
	if !(plan.Template.IsNull() || plan.Template.IsUnknown()) {
		body["template"] = client.FormatBool(plan.Template.ValueBool())
	}
	if !(plan.Tunnel.IsNull() || plan.Tunnel.IsUnknown()) {
		body["tunnel"] = client.FormatBool(plan.Tunnel.ValueBool())
	}
	obj, err := c.Add(ctx, "/ip/ipsec/policy", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/ipsec/policy failed", err.Error())
		return
	}
	iPIpsecPolicyApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIpsecPolicyResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPIpsecPolicyModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/ipsec/policy", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/ipsec/policy failed", err.Error())
		return
	}
	iPIpsecPolicyApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPIpsecPolicyResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPIpsecPolicyModel
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
	if !plan.Action.Equal(state.Action) && !plan.Action.IsUnknown() {
		body["action"] = plan.Action.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DstAddress.Equal(state.DstAddress) && !plan.DstAddress.IsUnknown() {
		body["dst-address"] = plan.DstAddress.ValueString()
	}
	if !plan.DstPort.Equal(state.DstPort) && !plan.DstPort.IsUnknown() {
		body["dst-port"] = client.FormatInt64(plan.DstPort.ValueInt64())
	}
	if !plan.Group.Equal(state.Group) && !plan.Group.IsUnknown() {
		body["group"] = plan.Group.ValueString()
	}
	if !plan.IpsecProtocols.Equal(state.IpsecProtocols) && !plan.IpsecProtocols.IsUnknown() {
		body["ipsec-protocols"] = plan.IpsecProtocols.ValueString()
	}
	if !plan.Level.Equal(state.Level) && !plan.Level.IsUnknown() {
		body["level"] = plan.Level.ValueString()
	}
	if !plan.Peer.Equal(state.Peer) && !plan.Peer.IsUnknown() {
		body["peer"] = plan.Peer.ValueString()
	}
	if !plan.Proposal.Equal(state.Proposal) && !plan.Proposal.IsUnknown() {
		body["proposal"] = plan.Proposal.ValueString()
	}
	if !plan.Protocol.Equal(state.Protocol) && !plan.Protocol.IsUnknown() {
		body["protocol"] = plan.Protocol.ValueString()
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) && !plan.SrcAddress.IsUnknown() {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if !plan.SrcPort.Equal(state.SrcPort) && !plan.SrcPort.IsUnknown() {
		body["src-port"] = client.FormatInt64(plan.SrcPort.ValueInt64())
	}
	if !plan.Template.Equal(state.Template) && !plan.Template.IsUnknown() {
		body["template"] = client.FormatBool(plan.Template.ValueBool())
	}
	if !plan.Tunnel.Equal(state.Tunnel) && !plan.Tunnel.IsUnknown() {
		body["tunnel"] = client.FormatBool(plan.Tunnel.ValueBool())
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/ipsec/policy", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/ipsec/policy failed", err.Error())
			return
		}
		iPIpsecPolicyApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIpsecPolicyResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPIpsecPolicyModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/ipsec/policy", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/ipsec/policy failed", err.Error())
	}
}

func (r *IPIpsecPolicyResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPIpsecPolicyLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/ipsec/policy matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPIpsecPolicyLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPIpsecPolicyLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/ipsec/policy", id)
}

func iPIpsecPolicyApply(ctx context.Context, obj client.Object, m *IPIpsecPolicyModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["action"]; ok {
		if v != "" {
			m.Action = types.StringValue(v)
		} else {
			m.Action = types.StringNull()
		}
	}
	if v, ok := obj["active"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Active = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Active = types.BoolValue(true)
		} else {
			m.Active = types.BoolNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["default"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Default = types.BoolValue(true)
		} else {
			m.Default = types.BoolNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Disabled = types.BoolValue(true)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["dst-address"]; ok {
		if v != "" {
			m.DstAddress = types.StringValue(v)
		} else {
			m.DstAddress = types.StringNull()
		}
	}
	if v, ok := obj["dst-port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.DstPort = types.Int64Value(n)
		} else {
			m.DstPort = types.Int64Null()
		}
	} else {
		m.DstPort = types.Int64Null()
	}
	if v, ok := obj["dynamic"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Dynamic = types.BoolValue(true)
		} else {
			m.Dynamic = types.BoolNull()
		}
	}
	if v, ok := obj["group"]; ok {
		if v != "" {
			m.Group = types.StringValue(v)
		} else {
			m.Group = types.StringNull()
		}
	}
	if v, ok := obj["invalid"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Invalid = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Invalid = types.BoolValue(true)
		} else {
			m.Invalid = types.BoolNull()
		}
	}
	if v, ok := obj["ipsec-protocols"]; ok {
		if v != "" {
			m.IpsecProtocols = types.StringValue(v)
		} else {
			m.IpsecProtocols = types.StringNull()
		}
	}
	if v, ok := obj["level"]; ok {
		if v != "" {
			m.Level = types.StringValue(v)
		} else {
			m.Level = types.StringNull()
		}
	}
	if v, ok := obj["nopeer"]; ok {
		if v != "" {
			m.Nopeer = types.StringValue(v)
		} else {
			m.Nopeer = types.StringNull()
		}
	}
	if v, ok := obj["notemplate"]; ok {
		if v != "" {
			m.Notemplate = types.StringValue(v)
		} else {
			m.Notemplate = types.StringNull()
		}
	}
	if v, ok := obj["peer"]; ok {
		if v != "" {
			m.Peer = types.StringValue(v)
		} else {
			m.Peer = types.StringNull()
		}
	}
	if v, ok := obj["ph2-count"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Ph2Count = types.Int64Value(n)
		} else {
			m.Ph2Count = types.Int64Null()
		}
	} else {
		m.Ph2Count = types.Int64Null()
	}
	if v, ok := obj["ph2-state"]; ok {
		if v != "" {
			m.Ph2State = types.StringValue(v)
		} else {
			m.Ph2State = types.StringNull()
		}
	}
	if v, ok := obj["proposal"]; ok {
		if v != "" {
			m.Proposal = types.StringValue(v)
		} else {
			m.Proposal = types.StringNull()
		}
	}
	if v, ok := obj["protocol"]; ok {
		if v != "" {
			m.Protocol = types.StringValue(v)
		} else {
			m.Protocol = types.StringNull()
		}
	}
	if v, ok := obj["sa-dst-address"]; ok {
		if v != "" {
			m.SaDstAddress = types.StringValue(v)
		} else {
			m.SaDstAddress = types.StringNull()
		}
	}
	if v, ok := obj["sa-src-address"]; ok {
		if v != "" {
			m.SaSrcAddress = types.StringValue(v)
		} else {
			m.SaSrcAddress = types.StringNull()
		}
	}
	if v, ok := obj["src-address"]; ok {
		if v != "" {
			m.SrcAddress = types.StringValue(v)
		} else {
			m.SrcAddress = types.StringNull()
		}
	}
	if v, ok := obj["src-port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.SrcPort = types.Int64Value(n)
		} else {
			m.SrcPort = types.Int64Null()
		}
	} else {
		m.SrcPort = types.Int64Null()
	}
	if v, ok := obj["template"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Template = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Template = types.BoolValue(true)
		} else {
			m.Template = types.BoolNull()
		}
	}
	if v, ok := obj["tunnel"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Tunnel = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Tunnel = types.BoolValue(true)
		} else {
			m.Tunnel = types.BoolNull()
		}
	}
}
