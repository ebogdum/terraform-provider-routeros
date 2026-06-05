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
	_ resource.Resource                = &InterfaceIpipv6Resource{}
	_ resource.ResourceWithImportState = &InterfaceIpipv6Resource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type InterfaceIpipv6Resource struct {
	reg *client.Registry
}

type InterfaceIpipv6Model struct {
	ID            types.String `tfsdk:"id"`
	Comment       types.String `tfsdk:"comment"`
	Disabled      types.Bool   `tfsdk:"disabled"`
	IpsecSecret   types.String `tfsdk:"ipsec_secret"`
	LocalAddress  types.String `tfsdk:"local_address"`
	MTU           types.String `tfsdk:"mtu"`
	Name          types.String `tfsdk:"name"`
	RemoteAddress types.String `tfsdk:"remote_address"`
	Router        types.String `tfsdk:"router"`
}

func NewInterfaceIpipv6Resource() resource.Resource { return &InterfaceIpipv6Resource{} }

func (r *InterfaceIpipv6Resource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_interface_ipipv6"
}

func (r *InterfaceIpipv6Resource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *InterfaceIpipv6Resource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/interface/ipipv6`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
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
			"ipsec_secret": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"local_address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mtu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"remote_address": schema.StringAttribute{
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

func (r *InterfaceIpipv6Resource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan InterfaceIpipv6Model
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.IpsecSecret.IsNull() || plan.IpsecSecret.IsUnknown()) {
		body["ipsec-secret"] = plan.IpsecSecret.ValueString()
	}
	if !(plan.LocalAddress.IsNull() || plan.LocalAddress.IsUnknown()) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !(plan.MTU.IsNull() || plan.MTU.IsUnknown()) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.RemoteAddress.IsNull() || plan.RemoteAddress.IsUnknown()) {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	obj, err := c.Add(ctx, "/interface/ipipv6", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /interface/ipipv6 failed", err.Error())
		return
	}
	interfaceIpipv6Apply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceIpipv6Resource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state InterfaceIpipv6Model
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/interface/ipipv6", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /interface/ipipv6 failed", err.Error())
		return
	}
	interfaceIpipv6Apply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *InterfaceIpipv6Resource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state InterfaceIpipv6Model
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
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.IpsecSecret.Equal(state.IpsecSecret) {
		body["ipsec-secret"] = plan.IpsecSecret.ValueString()
	}
	if !plan.LocalAddress.Equal(state.LocalAddress) {
		body["local-address"] = plan.LocalAddress.ValueString()
	}
	if !plan.MTU.Equal(state.MTU) {
		body["mtu"] = plan.MTU.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.RemoteAddress.Equal(state.RemoteAddress) {
		body["remote-address"] = plan.RemoteAddress.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/interface/ipipv6", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /interface/ipipv6 failed", err.Error())
			return
		}
		interfaceIpipv6Apply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *InterfaceIpipv6Resource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state InterfaceIpipv6Model
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/interface/ipipv6", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /interface/ipipv6 failed", err.Error())
	}
}

func (r *InterfaceIpipv6Resource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import formats accepted:
	//   *<id>                            -> bare RouterOS .id on the default router
	//   <router>/*<id>                   -> .id on the named router
	//   <router>/<naturalkey>            -> resolved via List + filter
	//   <naturalkey>                     -> resolved on the default router
	id := req.ID
	routerName := ""
	if i := strings.Index(id, "/"); i > 0 && !strings.HasPrefix(id, "*") {
		routerName, id = id[:i], id[i+1:]
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("router"), types.StringValue(routerName))...)
	if strings.HasPrefix(id, "*") {
		resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(id))...)
		return
	}
	c := pickClient(r.reg, types.StringValue(routerName), &resp.Diagnostics)
	if c == nil {
		return
	}
	rows, err := interfaceIpipv6LookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /interface/ipipv6 matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// interfaceIpipv6LookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func interfaceIpipv6LookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/interface/ipipv6", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func interfaceIpipv6Apply(ctx context.Context, obj client.Object, m *InterfaceIpipv6Model) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
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
	if v, ok := obj["ipsec-secret"]; ok {
		_ = v
		if v != "" {
			m.IpsecSecret = types.StringValue(v)
		} else {
			m.IpsecSecret = types.StringNull()
		}
	} else {
		m.IpsecSecret = types.StringNull()
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
	if v, ok := obj["mtu"]; ok {
		_ = v
		if v != "" {
			m.MTU = types.StringValue(v)
		} else {
			m.MTU = types.StringNull()
		}
	} else {
		m.MTU = types.StringNull()
	}
	if v, ok := obj["name"]; ok {
		_ = v
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	} else {
		m.Name = types.StringNull()
	}
	if v, ok := obj["remote-address"]; ok {
		_ = v
		if v != "" {
			m.RemoteAddress = types.StringValue(v)
		} else {
			m.RemoteAddress = types.StringNull()
		}
	} else {
		m.RemoteAddress = types.StringNull()
	}
}
