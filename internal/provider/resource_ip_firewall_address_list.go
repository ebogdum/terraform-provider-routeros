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
	_ resource.Resource                = &IPFirewallAddressListResource{}
	_ resource.ResourceWithImportState = &IPFirewallAddressListResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPFirewallAddressListResource struct {
	reg *client.Registry
}

type IPFirewallAddressListModel struct {
	ID           types.String `tfsdk:"id"`
	Address      types.String `tfsdk:"address"`
	Comment      types.String `tfsdk:"comment"`
	CreationTime types.String `tfsdk:"creation_time"`
	Disabled     types.Bool   `tfsdk:"disabled"`
	Dynamic      types.Bool   `tfsdk:"dynamic"`
	List         types.String `tfsdk:"list"`
	Parent       types.Int64  `tfsdk:"parent"`
	Timeout      types.String `tfsdk:"timeout"`
	Router       types.String `tfsdk:"router"`
}

func NewIPFirewallAddressListResource() resource.Resource { return &IPFirewallAddressListResource{} }

func (r *IPFirewallAddressListResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_firewall_address_list"
}

func (r *IPFirewallAddressListResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *IPFirewallAddressListResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/firewall/address-list`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"address": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"creation_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"dynamic": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"list": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"parent": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"timeout": schema.StringAttribute{
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

func (r *IPFirewallAddressListResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPFirewallAddressListModel
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
	if !(plan.List.IsNull() || plan.List.IsUnknown()) {
		body["list"] = plan.List.ValueString()
	}
	if !(plan.Parent.IsNull() || plan.Parent.IsUnknown()) {
		body["parent"] = client.FormatInt64(plan.Parent.ValueInt64())
	}
	if !(plan.Timeout.IsNull() || plan.Timeout.IsUnknown()) {
		body["timeout"] = plan.Timeout.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/firewall/address-list", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/firewall/address-list failed", err.Error())
		return
	}
	iPFirewallAddressListApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPFirewallAddressListResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPFirewallAddressListModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/firewall/address-list", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/firewall/address-list failed", err.Error())
		return
	}
	iPFirewallAddressListApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPFirewallAddressListResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPFirewallAddressListModel
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
	if !plan.Address.Equal(state.Address) {
		body["address"] = plan.Address.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.List.Equal(state.List) {
		body["list"] = plan.List.ValueString()
	}
	if !plan.Parent.Equal(state.Parent) {
		body["parent"] = client.FormatInt64(plan.Parent.ValueInt64())
	}
	if !plan.Timeout.Equal(state.Timeout) {
		body["timeout"] = plan.Timeout.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/firewall/address-list", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/firewall/address-list failed", err.Error())
			return
		}
		iPFirewallAddressListApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPFirewallAddressListResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPFirewallAddressListModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/firewall/address-list", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/firewall/address-list failed", err.Error())
	}
}

func (r *IPFirewallAddressListResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPFirewallAddressListLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/firewall/address-list matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPFirewallAddressListLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPFirewallAddressListLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/ip/firewall/address-list", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func iPFirewallAddressListApply(ctx context.Context, obj client.Object, m *IPFirewallAddressListModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["address"]; ok {
		_ = v
		if v != "" {
			m.Address = types.StringValue(v)
		} else {
			m.Address = types.StringNull()
		}
	} else {
		m.Address = types.StringNull()
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
	if v, ok := obj["creation-time"]; ok {
		_ = v
		if v != "" {
			m.CreationTime = types.StringValue(v)
		} else {
			m.CreationTime = types.StringNull()
		}
	} else {
		m.CreationTime = types.StringNull()
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
	if v, ok := obj["dynamic"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else {
			m.Dynamic = types.BoolNull()
		}
	} else {
		m.Dynamic = types.BoolNull()
	}
	if v, ok := obj["list"]; ok {
		_ = v
		if v != "" {
			m.List = types.StringValue(v)
		} else {
			m.List = types.StringNull()
		}
	} else {
		m.List = types.StringNull()
	}
	if v, ok := obj["parent"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Parent = types.Int64Value(n)
		} else {
			m.Parent = types.Int64Null()
		}
	} else {
		m.Parent = types.Int64Null()
	}
	if v, ok := obj["timeout"]; ok {
		_ = v
		if v != "" {
			m.Timeout = types.StringValue(v)
		} else {
			m.Timeout = types.StringNull()
		}
	} else {
		m.Timeout = types.StringNull()
	}
}
