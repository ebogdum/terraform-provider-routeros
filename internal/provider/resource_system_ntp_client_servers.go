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
	_ resource.Resource                = &SystemNTPClientServersResource{}
	_ resource.ResourceWithImportState = &SystemNTPClientServersResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemNTPClientServersResource struct {
	reg *client.Registry
}

type SystemNTPClientServersModel struct {
	ID              types.String `tfsdk:"id"`
	Address         types.String `tfsdk:"address"`
	AuthKey         types.String `tfsdk:"auth_key"`
	Comment         types.String `tfsdk:"comment"`
	Disabled        types.Bool   `tfsdk:"disabled"`
	Dynamic         types.Bool   `tfsdk:"dynamic"`
	Iburst          types.Bool   `tfsdk:"iburst"`
	Keys            types.String `tfsdk:"keys"`
	MaxPoll         types.Int64  `tfsdk:"max_poll"`
	MinPoll         types.Int64  `tfsdk:"min_poll"`
	ResolvedAddress types.String `tfsdk:"resolved_address"`
	Router          types.String `tfsdk:"router"`
}

func NewSystemNTPClientServersResource() resource.Resource { return &SystemNTPClientServersResource{} }

func (r *SystemNTPClientServersResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_ntp_client_servers"
}

func (r *SystemNTPClientServersResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *SystemNTPClientServersResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "NTP server list — accepts add but validator differs per ROS. Skipped from acc tests.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"auth_key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
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
				Description: "",
			},
			"dynamic": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"iburst": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"keys": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"max_poll": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"min_poll": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"resolved_address": schema.StringAttribute{
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

func (r *SystemNTPClientServersResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemNTPClientServersModel
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
	if !(plan.AuthKey.IsNull() || plan.AuthKey.IsUnknown()) {
		body["auth-key"] = plan.AuthKey.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Iburst.IsNull() || plan.Iburst.IsUnknown()) {
		body["iburst"] = client.FormatBool(plan.Iburst.ValueBool())
	}
	if !(plan.Keys.IsNull() || plan.Keys.IsUnknown()) {
		body["keys"] = plan.Keys.ValueString()
	}
	if !(plan.MaxPoll.IsNull() || plan.MaxPoll.IsUnknown()) {
		body["max-poll"] = client.FormatInt64(plan.MaxPoll.ValueInt64())
	}
	if !(plan.MinPoll.IsNull() || plan.MinPoll.IsUnknown()) {
		body["min-poll"] = client.FormatInt64(plan.MinPoll.ValueInt64())
	}
	obj, err := c.Add(ctx, "/system/ntp/client/servers", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /system/ntp/client/servers failed", err.Error())
		return
	}
	systemNTPClientServersApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemNTPClientServersResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemNTPClientServersModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/system/ntp/client/servers", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /system/ntp/client/servers failed", err.Error())
		return
	}
	systemNTPClientServersApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemNTPClientServersResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemNTPClientServersModel
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
	if !plan.AuthKey.Equal(state.AuthKey) {
		body["auth-key"] = plan.AuthKey.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Iburst.Equal(state.Iburst) {
		body["iburst"] = client.FormatBool(plan.Iburst.ValueBool())
	}
	if !plan.Keys.Equal(state.Keys) {
		body["keys"] = plan.Keys.ValueString()
	}
	if !plan.MaxPoll.Equal(state.MaxPoll) {
		body["max-poll"] = client.FormatInt64(plan.MaxPoll.ValueInt64())
	}
	if !plan.MinPoll.Equal(state.MinPoll) {
		body["min-poll"] = client.FormatInt64(plan.MinPoll.ValueInt64())
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/system/ntp/client/servers", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /system/ntp/client/servers failed", err.Error())
			return
		}
		systemNTPClientServersApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemNTPClientServersResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemNTPClientServersModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/system/ntp/client/servers", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /system/ntp/client/servers failed", err.Error())
	}
}

func (r *SystemNTPClientServersResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := systemNTPClientServersLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /system/ntp/client/servers matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// systemNTPClientServersLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func systemNTPClientServersLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/system/ntp/client/servers", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func systemNTPClientServersApply(ctx context.Context, obj client.Object, m *SystemNTPClientServersModel) {
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
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.AuthKey already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["auth-key"]; ok && v != "" {
		_ = v
		if v != "" {
			m.AuthKey = types.StringValue(v)
		} else {
			m.AuthKey = types.StringNull()
		}
	} else if m.AuthKey.IsUnknown() {
		m.AuthKey = types.StringNull()
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
	if v, ok := obj["iburst"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Iburst = types.BoolValue(b)
		} else {
			m.Iburst = types.BoolNull()
		}
	} else {
		m.Iburst = types.BoolNull()
	}
	if v, ok := obj["keys"]; ok {
		_ = v
		if v != "" {
			m.Keys = types.StringValue(v)
		} else {
			m.Keys = types.StringNull()
		}
	} else {
		m.Keys = types.StringNull()
	}
	if v, ok := obj["max-poll"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MaxPoll = types.Int64Value(n)
		} else {
			m.MaxPoll = types.Int64Null()
		}
	} else {
		m.MaxPoll = types.Int64Null()
	}
	if v, ok := obj["min-poll"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.MinPoll = types.Int64Value(n)
		} else {
			m.MinPoll = types.Int64Null()
		}
	} else {
		m.MinPoll = types.Int64Null()
	}
	if v, ok := obj["resolved-address"]; ok {
		_ = v
		if v != "" {
			m.ResolvedAddress = types.StringValue(v)
		} else {
			m.ResolvedAddress = types.StringNull()
		}
	} else {
		m.ResolvedAddress = types.StringNull()
	}
}
