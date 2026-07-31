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
	_ resource.Resource                = &IPMediaResource{}
	_ resource.ResourceWithImportState = &IPMediaResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPMediaResource struct {
	reg *client.Registry
}

type IPMediaModel struct {
	ID              types.String `tfsdk:"id"`
	FriendlyName    types.String `tfsdk:"friendly_name"`
	AllowedIp       types.String `tfsdk:"allowed_ip"`
	AllowedHostname types.String `tfsdk:"allowed_hostname"`
	Disabled        types.Bool   `tfsdk:"disabled"`
	Interface       types.String `tfsdk:"interface"`
	Path            types.String `tfsdk:"path"`
	Router          types.String `tfsdk:"router"`
}

func NewIPMediaResource() resource.Resource { return &IPMediaResource{} }

func (r *IPMediaResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_media"
}

func (r *IPMediaResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPMediaResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/media`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"friendly_name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `friendly-name`.",
			},
			"allowed_ip": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `allowed-ip`.",
			},
			"allowed_hostname": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `allowed-hostname`.",
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
			"path": schema.StringAttribute{
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

func (r *IPMediaResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPMediaModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Interface.IsNull() || plan.Interface.IsUnknown()) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !(plan.Path.IsNull() || plan.Path.IsUnknown()) {
		body["path"] = plan.Path.ValueString()
	}
	if !(plan.AllowedHostname.IsNull() || plan.AllowedHostname.IsUnknown()) {
		body["allowed-hostname"] = plan.AllowedHostname.ValueString()
	}
	if !(plan.AllowedIp.IsNull() || plan.AllowedIp.IsUnknown()) {
		body["allowed-ip"] = plan.AllowedIp.ValueString()
	}
	if !(plan.FriendlyName.IsNull() || plan.FriendlyName.IsUnknown()) {
		body["friendly-name"] = plan.FriendlyName.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/media", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/media failed", err.Error())
		return
	}
	iPMediaApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPMediaResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPMediaModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/media", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/media failed", err.Error())
		return
	}
	iPMediaApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPMediaResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPMediaModel
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
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Interface.Equal(state.Interface) {
		body["interface"] = plan.Interface.ValueString()
	}
	if !plan.Path.Equal(state.Path) {
		body["path"] = plan.Path.ValueString()
	}
	if !plan.AllowedHostname.Equal(state.AllowedHostname) && !plan.AllowedHostname.IsUnknown() {
		body["allowed-hostname"] = plan.AllowedHostname.ValueString()
	}
	if !plan.AllowedIp.Equal(state.AllowedIp) && !plan.AllowedIp.IsUnknown() {
		body["allowed-ip"] = plan.AllowedIp.ValueString()
	}
	if !plan.FriendlyName.Equal(state.FriendlyName) && !plan.FriendlyName.IsUnknown() {
		body["friendly-name"] = plan.FriendlyName.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/media", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/media failed", err.Error())
			return
		}
		iPMediaApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPMediaResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPMediaModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/media", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/media failed", err.Error())
	}
}

func (r *IPMediaResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPMediaLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/media matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPMediaLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPMediaLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/media", id)
}

func iPMediaApply(ctx context.Context, obj client.Object, m *IPMediaModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["friendly-name"]; ok && v != "" {
		m.FriendlyName = types.StringValue(v)
	} else {
		m.FriendlyName = types.StringNull()
	}
	if v, ok := obj["allowed-ip"]; ok && v != "" {
		m.AllowedIp = types.StringValue(v)
	} else {
		m.AllowedIp = types.StringNull()
	}
	if v, ok := obj["allowed-hostname"]; ok && v != "" {
		m.AllowedHostname = types.StringValue(v)
	} else {
		m.AllowedHostname = types.StringNull()
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
	if v, ok := obj["path"]; ok {
		_ = v
		if v != "" {
			m.Path = types.StringValue(v)
		} else {
			m.Path = types.StringNull()
		}
	} else {
		m.Path = types.StringNull()
	}
}
