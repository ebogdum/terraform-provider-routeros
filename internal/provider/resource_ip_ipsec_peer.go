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
	_ resource.Resource                = &IPIpsecPeerResource{}
	_ resource.ResourceWithImportState = &IPIpsecPeerResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPIpsecPeerResource struct {
	reg *client.Registry
}

type IPIpsecPeerModel struct {
	ID                 types.String `tfsdk:"id"`
	PpkSecret          types.String `tfsdk:"ppk_secret"`
	Address            types.String `tfsdk:"address"`
	Comment            types.String `tfsdk:"comment"`
	Disabled           types.Bool   `tfsdk:"disabled"`
	Dynamic            types.Bool   `tfsdk:"dynamic"`
	ExchangeMode       types.String `tfsdk:"exchange_mode"`
	LocalAddress       types.String `tfsdk:"local_address"`
	Name               types.String `tfsdk:"name"`
	Passive            types.Bool   `tfsdk:"passive"`
	Port               types.Int64  `tfsdk:"port"`
	Profile            types.String `tfsdk:"profile"`
	Responder          types.Bool   `tfsdk:"responder"`
	SendInitialContact types.Bool   `tfsdk:"send_initial_contact"`
	Router             types.String `tfsdk:"router"`
}

func NewIPIpsecPeerResource() resource.Resource { return &IPIpsecPeerResource{} }

func (r *IPIpsecPeerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_ipsec_peer"
}

func (r *IPIpsecPeerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPIpsecPeerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/ipsec/peer`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"ppk_secret": schema.StringAttribute{
				Optional:    true,
				Sensitive:   true,
				Computed:    true,
				Description: "RouterOS `ppk-secret`.",
			},
			"address": schema.StringAttribute{
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
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"exchange_mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"base", "main", "aggressive", "ike2"}...)},
			},
			"local_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"passive": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"profile": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"responder": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"send_initial_contact": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *IPIpsecPeerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPIpsecPeerModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Address.IsNull() || plan.Address.IsUnknown()) {
		body["address"] = plan.Address.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.ExchangeMode.IsNull() || plan.ExchangeMode.IsUnknown()) {
		body["exchange-mode"] = plan.ExchangeMode.ValueString()
	}
	if !(plan.LocalAddress.IsNull() || plan.LocalAddress.IsUnknown()) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Passive.IsNull() || plan.Passive.IsUnknown()) {
		body["passive"] = client.FormatBool(plan.Passive.ValueBool())
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = client.FormatInt64(plan.Port.ValueInt64())
	}
	if !(plan.Profile.IsNull() || plan.Profile.IsUnknown()) {
		body["profile"] = plan.Profile.ValueString()
	}
	if !(plan.SendInitialContact.IsNull() || plan.SendInitialContact.IsUnknown()) {
		body["send-initial-contact"] = client.FormatBool(plan.SendInitialContact.ValueBool())
	}
	if !(plan.PpkSecret.IsNull() || plan.PpkSecret.IsUnknown()) {
		body["ppk-secret"] = plan.PpkSecret.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/ipsec/peer", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/ipsec/peer failed", err.Error())
		return
	}
	iPIpsecPeerApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIpsecPeerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPIpsecPeerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/ipsec/peer", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/ipsec/peer failed", err.Error())
		return
	}
	iPIpsecPeerApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPIpsecPeerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPIpsecPeerModel
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
	if !plan.Address.Equal(state.Address) && !plan.Address.IsUnknown() {
		body["address"] = plan.Address.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.ExchangeMode.Equal(state.ExchangeMode) && !plan.ExchangeMode.IsUnknown() {
		body["exchange-mode"] = plan.ExchangeMode.ValueString()
	}
	if !plan.LocalAddress.Equal(state.LocalAddress) && !plan.LocalAddress.IsUnknown() {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Passive.Equal(state.Passive) && !plan.Passive.IsUnknown() {
		body["passive"] = client.FormatBool(plan.Passive.ValueBool())
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = client.FormatInt64(plan.Port.ValueInt64())
	}
	if !plan.Profile.Equal(state.Profile) && !plan.Profile.IsUnknown() {
		body["profile"] = plan.Profile.ValueString()
	}
	if !plan.SendInitialContact.Equal(state.SendInitialContact) && !plan.SendInitialContact.IsUnknown() {
		body["send-initial-contact"] = client.FormatBool(plan.SendInitialContact.ValueBool())
	}
	if !plan.PpkSecret.Equal(state.PpkSecret) && !plan.PpkSecret.IsUnknown() {
		body["ppk-secret"] = plan.PpkSecret.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/ipsec/peer", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/ipsec/peer failed", err.Error())
			return
		}
		iPIpsecPeerApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIpsecPeerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPIpsecPeerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/ipsec/peer", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/ipsec/peer failed", err.Error())
	}
}

func (r *IPIpsecPeerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPIpsecPeerLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/ipsec/peer matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPIpsecPeerLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPIpsecPeerLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/ipsec/peer", id)
}

func iPIpsecPeerApply(ctx context.Context, obj client.Object, m *IPIpsecPeerModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["ppk-secret"]; ok && v != "" {
		m.PpkSecret = types.StringValue(v)
	} else {
		m.PpkSecret = types.StringNull()
	}
	if v, ok := obj["address"]; ok {
		if v != "" {
			m.Address = types.StringValue(v)
		} else {
			m.Address = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
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
	if v, ok := obj["dynamic"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Dynamic = types.BoolValue(true)
		} else {
			m.Dynamic = types.BoolNull()
		}
	}
	if v, ok := obj["exchange-mode"]; ok {
		if v != "" {
			m.ExchangeMode = types.StringValue(v)
		} else {
			m.ExchangeMode = types.StringNull()
		}
	}
	if v, ok := obj["local-address"]; ok {
		if v != "" {
			m.LocalAddress = types.StringValue(v)
		} else {
			m.LocalAddress = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["passive"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Passive = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Passive = types.BoolValue(true)
		} else {
			m.Passive = types.BoolNull()
		}
	}
	if v, ok := obj["port"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Port = types.Int64Value(n)
		} else {
			m.Port = types.Int64Null()
		}
	} else {
		m.Port = types.Int64Null()
	}
	if v, ok := obj["profile"]; ok {
		if v != "" {
			m.Profile = types.StringValue(v)
		} else {
			m.Profile = types.StringNull()
		}
	}
	if v, ok := obj["responder"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Responder = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Responder = types.BoolValue(true)
		} else {
			m.Responder = types.BoolNull()
		}
	}
	if v, ok := obj["send-initial-contact"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SendInitialContact = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.SendInitialContact = types.BoolValue(true)
		} else {
			m.SendInitialContact = types.BoolNull()
		}
	}
}
