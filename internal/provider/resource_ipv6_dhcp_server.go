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
	_ resource.Resource                = &IPV6DHCPServerResource{}
	_ resource.ResourceWithImportState = &IPV6DHCPServerResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPV6DHCPServerResource struct {
	reg *client.Registry
}

type IPV6DHCPServerModel struct {
	ID                  types.String `tfsdk:"id"`
	UseReconfigure      types.String `tfsdk:"use_reconfigure"`
	UseRadius           types.String `tfsdk:"use_radius"`
	RouteDistance       types.String `tfsdk:"route_distance"`
	RapidCommit         types.String `tfsdk:"rapid_commit"`
	PrefixPool          types.String `tfsdk:"prefix_pool"`
	Preference          types.String `tfsdk:"preference"`
	ParentQueue         types.String `tfsdk:"parent_queue"`
	InsertQueueBefore   types.String `tfsdk:"insert_queue_before"`
	IgnoreIaNaBindings  types.String `tfsdk:"ignore_ia_na_bindings"`
	BindingScript       types.String `tfsdk:"binding_script"`
	AllowDualStackQueue types.String `tfsdk:"allow_dual_stack_queue"`
	AddressPool         types.String `tfsdk:"address_pool"`
	AddressLists        types.String `tfsdk:"address_lists"`
	Comment             types.String `tfsdk:"comment"`
	DHCPOption          types.String `tfsdk:"dhcp_option"`
	Disabled            types.Bool   `tfsdk:"disabled"`
	Interface           types.String `tfsdk:"interface"`
	LeaseTime           types.String `tfsdk:"lease_time"`
	Name                types.String `tfsdk:"name"`
	Router              types.String `tfsdk:"router"`
}

func NewIPV6DHCPServerResource() resource.Resource { return &IPV6DHCPServerResource{} }

func (r *IPV6DHCPServerResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ipv6_dhcp_server"
}

func (r *IPV6DHCPServerResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPV6DHCPServerResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ipv6/dhcp-server`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"use_reconfigure": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `use-reconfigure`.",
			},
			"use_radius": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `use-radius`.",
			},
			"route_distance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `route-distance`.",
			},
			"rapid_commit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `rapid-commit`.",
			},
			"prefix_pool": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `prefix-pool`.",
			},
			"preference": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `preference`.",
			},
			"parent_queue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `parent-queue`.",
			},
			"insert_queue_before": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `insert-queue-before`.",
			},
			"ignore_ia_na_bindings": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `ignore-ia-na-bindings`.",
			},
			"binding_script": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `binding-script`.",
			},
			"allow_dual_stack_queue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `allow-dual-stack-queue`.",
			},
			"address_pool": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `address-pool`.",
			},
			"address_lists": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `address-lists`.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"dhcp_option": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
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
			"lease_time": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPV6DHCPServerResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPV6DHCPServerModel
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
	if !(plan.DHCPOption.IsNull() || plan.DHCPOption.IsUnknown()) {
		body["dhcp-option"] = plan.DHCPOption.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.LeaseTime.IsNull() || plan.LeaseTime.IsUnknown()) {
		body["lease-time"] = plan.LeaseTime.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.AddressLists.IsNull() || plan.AddressLists.IsUnknown()) {
		body["address-lists"] = plan.AddressLists.ValueString()
	}
	if !(plan.AddressPool.IsNull() || plan.AddressPool.IsUnknown()) {
		body["address-pool"] = plan.AddressPool.ValueString()
	}
	if !(plan.AllowDualStackQueue.IsNull() || plan.AllowDualStackQueue.IsUnknown()) {
		body["allow-dual-stack-queue"] = plan.AllowDualStackQueue.ValueString()
	}
	if !(plan.BindingScript.IsNull() || plan.BindingScript.IsUnknown()) {
		body["binding-script"] = plan.BindingScript.ValueString()
	}
	if !(plan.IgnoreIaNaBindings.IsNull() || plan.IgnoreIaNaBindings.IsUnknown()) {
		body["ignore-ia-na-bindings"] = plan.IgnoreIaNaBindings.ValueString()
	}
	if !(plan.InsertQueueBefore.IsNull() || plan.InsertQueueBefore.IsUnknown()) {
		body["insert-queue-before"] = plan.InsertQueueBefore.ValueString()
	}
	if !(plan.ParentQueue.IsNull() || plan.ParentQueue.IsUnknown()) {
		body["parent-queue"] = plan.ParentQueue.ValueString()
	}
	if !(plan.Preference.IsNull() || plan.Preference.IsUnknown()) {
		body["preference"] = plan.Preference.ValueString()
	}
	if !(plan.PrefixPool.IsNull() || plan.PrefixPool.IsUnknown()) {
		body["prefix-pool"] = plan.PrefixPool.ValueString()
	}
	if !(plan.RapidCommit.IsNull() || plan.RapidCommit.IsUnknown()) {
		body["rapid-commit"] = plan.RapidCommit.ValueString()
	}
	if !(plan.RouteDistance.IsNull() || plan.RouteDistance.IsUnknown()) {
		body["route-distance"] = plan.RouteDistance.ValueString()
	}
	if !(plan.UseRadius.IsNull() || plan.UseRadius.IsUnknown()) {
		body["use-radius"] = plan.UseRadius.ValueString()
	}
	if !(plan.UseReconfigure.IsNull() || plan.UseReconfigure.IsUnknown()) {
		body["use-reconfigure"] = plan.UseReconfigure.ValueString()
	}
	obj, err := c.Add(ctx, "/ipv6/dhcp-server", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ipv6/dhcp-server failed", err.Error())
		return
	}
	iPV6DHCPServerApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6DHCPServerResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPV6DHCPServerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ipv6/dhcp-server", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ipv6/dhcp-server failed", err.Error())
		return
	}
	iPV6DHCPServerApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPV6DHCPServerResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPV6DHCPServerModel
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
	if !plan.DHCPOption.Equal(state.DHCPOption) {
		body["dhcp-option"] = plan.DHCPOption.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.LeaseTime.Equal(state.LeaseTime) {
		body["lease-time"] = plan.LeaseTime.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.AddressLists.Equal(state.AddressLists) && !plan.AddressLists.IsUnknown() {
		body["address-lists"] = plan.AddressLists.ValueString()
	}
	if !plan.AddressPool.Equal(state.AddressPool) && !plan.AddressPool.IsUnknown() {
		body["address-pool"] = plan.AddressPool.ValueString()
	}
	if !plan.AllowDualStackQueue.Equal(state.AllowDualStackQueue) && !plan.AllowDualStackQueue.IsUnknown() {
		body["allow-dual-stack-queue"] = plan.AllowDualStackQueue.ValueString()
	}
	if !plan.BindingScript.Equal(state.BindingScript) && !plan.BindingScript.IsUnknown() {
		body["binding-script"] = plan.BindingScript.ValueString()
	}
	if !plan.IgnoreIaNaBindings.Equal(state.IgnoreIaNaBindings) && !plan.IgnoreIaNaBindings.IsUnknown() {
		body["ignore-ia-na-bindings"] = plan.IgnoreIaNaBindings.ValueString()
	}
	if !plan.InsertQueueBefore.Equal(state.InsertQueueBefore) && !plan.InsertQueueBefore.IsUnknown() {
		body["insert-queue-before"] = plan.InsertQueueBefore.ValueString()
	}
	if !plan.ParentQueue.Equal(state.ParentQueue) && !plan.ParentQueue.IsUnknown() {
		body["parent-queue"] = plan.ParentQueue.ValueString()
	}
	if !plan.Preference.Equal(state.Preference) && !plan.Preference.IsUnknown() {
		body["preference"] = plan.Preference.ValueString()
	}
	if !plan.PrefixPool.Equal(state.PrefixPool) && !plan.PrefixPool.IsUnknown() {
		body["prefix-pool"] = plan.PrefixPool.ValueString()
	}
	if !plan.RapidCommit.Equal(state.RapidCommit) && !plan.RapidCommit.IsUnknown() {
		body["rapid-commit"] = plan.RapidCommit.ValueString()
	}
	if !plan.RouteDistance.Equal(state.RouteDistance) && !plan.RouteDistance.IsUnknown() {
		body["route-distance"] = plan.RouteDistance.ValueString()
	}
	if !plan.UseRadius.Equal(state.UseRadius) && !plan.UseRadius.IsUnknown() {
		body["use-radius"] = plan.UseRadius.ValueString()
	}
	if !plan.UseReconfigure.Equal(state.UseReconfigure) && !plan.UseReconfigure.IsUnknown() {
		body["use-reconfigure"] = plan.UseReconfigure.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ipv6/dhcp-server", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ipv6/dhcp-server failed", err.Error())
			return
		}
		iPV6DHCPServerApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPV6DHCPServerResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPV6DHCPServerModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ipv6/dhcp-server", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ipv6/dhcp-server failed", err.Error())
	}
}

func (r *IPV6DHCPServerResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPV6DHCPServerLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ipv6/dhcp-server matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPV6DHCPServerLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPV6DHCPServerLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ipv6/dhcp-server", id)
}

func iPV6DHCPServerApply(ctx context.Context, obj client.Object, m *IPV6DHCPServerModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["use-reconfigure"]; ok && v != "" {
		m.UseReconfigure = types.StringValue(v)
	} else {
		m.UseReconfigure = types.StringNull()
	}
	if v, ok := obj["use-radius"]; ok && v != "" {
		m.UseRadius = types.StringValue(v)
	} else {
		m.UseRadius = types.StringNull()
	}
	if v, ok := obj["route-distance"]; ok && v != "" {
		m.RouteDistance = types.StringValue(v)
	} else {
		m.RouteDistance = types.StringNull()
	}
	if v, ok := obj["rapid-commit"]; ok && v != "" {
		m.RapidCommit = types.StringValue(v)
	} else {
		m.RapidCommit = types.StringNull()
	}
	if v, ok := obj["prefix-pool"]; ok && v != "" {
		m.PrefixPool = types.StringValue(v)
	} else {
		m.PrefixPool = types.StringNull()
	}
	if v, ok := obj["preference"]; ok && v != "" {
		m.Preference = types.StringValue(v)
	} else {
		m.Preference = types.StringNull()
	}
	if v, ok := obj["parent-queue"]; ok && v != "" {
		m.ParentQueue = types.StringValue(v)
	} else {
		m.ParentQueue = types.StringNull()
	}
	if v, ok := obj["insert-queue-before"]; ok && v != "" {
		m.InsertQueueBefore = types.StringValue(v)
	} else {
		m.InsertQueueBefore = types.StringNull()
	}
	if v, ok := obj["ignore-ia-na-bindings"]; ok && v != "" {
		m.IgnoreIaNaBindings = types.StringValue(v)
	} else {
		m.IgnoreIaNaBindings = types.StringNull()
	}
	if v, ok := obj["binding-script"]; ok && v != "" {
		m.BindingScript = types.StringValue(v)
	} else {
		m.BindingScript = types.StringNull()
	}
	if v, ok := obj["allow-dual-stack-queue"]; ok && v != "" {
		m.AllowDualStackQueue = types.StringValue(v)
	} else {
		m.AllowDualStackQueue = types.StringNull()
	}
	if v, ok := obj["address-pool"]; ok && v != "" {
		m.AddressPool = types.StringValue(v)
	} else {
		m.AddressPool = types.StringNull()
	}
	if v, ok := obj["address-lists"]; ok && v != "" {
		m.AddressLists = types.StringValue(v)
	} else {
		m.AddressLists = types.StringNull()
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
	if v, ok := obj["dhcp-option"]; ok {
		_ = v
		if v != "" {
			m.DHCPOption = types.StringValue(v)
		} else {
			m.DHCPOption = types.StringNull()
		}
	} else {
		m.DHCPOption = types.StringNull()
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
	if v, ok := obj["lease-time"]; ok {
		_ = v
		if v != "" {
			m.LeaseTime = types.StringValue(v)
		} else {
			m.LeaseTime = types.StringNull()
		}
	} else {
		m.LeaseTime = types.StringNull()
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
}
