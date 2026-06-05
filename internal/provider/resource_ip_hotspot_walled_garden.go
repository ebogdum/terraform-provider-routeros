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
	_ resource.Resource                = &IPHotspotWalledGardenResource{}
	_ resource.ResourceWithImportState = &IPHotspotWalledGardenResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPHotspotWalledGardenResource struct {
	reg *client.Registry
}

type IPHotspotWalledGardenModel struct {
	ID         types.String `tfsdk:"id"`
	Action     types.String `tfsdk:"action"`
	Comment    types.String `tfsdk:"comment"`
	Disabled   types.Bool   `tfsdk:"disabled"`
	DstPort    types.String `tfsdk:"dst_port"`
	Path       types.String `tfsdk:"path"`
	Server     types.String `tfsdk:"server"`
	SrcAddress types.String `tfsdk:"src_address"`
	Router     types.String `tfsdk:"router"`
}

func NewIPHotspotWalledGardenResource() resource.Resource { return &IPHotspotWalledGardenResource{} }

func (r *IPHotspotWalledGardenResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_hotspot_walled_garden"
}

func (r *IPHotspotWalledGardenResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *IPHotspotWalledGardenResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/hotspot/walled-garden`.",
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
			"dst_port": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"path": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"server": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"src_address": schema.StringAttribute{
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

func (r *IPHotspotWalledGardenResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPHotspotWalledGardenModel
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
	if !(plan.DstPort.IsNull() || plan.DstPort.IsUnknown()) {
		body["dst-port"] = plan.DstPort.ValueString()
	}
	if !(plan.Path.IsNull() || plan.Path.IsUnknown()) {
		body["path"] = plan.Path.ValueString()
	}
	if !(plan.Server.IsNull() || plan.Server.IsUnknown()) {
		body["server"] = plan.Server.ValueString()
	}
	if !(plan.SrcAddress.IsNull() || plan.SrcAddress.IsUnknown()) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/hotspot/walled-garden", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/hotspot/walled-garden failed", err.Error())
		return
	}
	iPHotspotWalledGardenApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPHotspotWalledGardenResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPHotspotWalledGardenModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/hotspot/walled-garden", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/hotspot/walled-garden failed", err.Error())
		return
	}
	iPHotspotWalledGardenApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPHotspotWalledGardenResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPHotspotWalledGardenModel
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
	if !plan.Action.Equal(state.Action) {
		body["action"] = plan.Action.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.DstPort.Equal(state.DstPort) {
		body["dst-port"] = plan.DstPort.ValueString()
	}
	if !plan.Path.Equal(state.Path) {
		body["path"] = plan.Path.ValueString()
	}
	if !plan.Server.Equal(state.Server) {
		body["server"] = plan.Server.ValueString()
	}
	if !plan.SrcAddress.Equal(state.SrcAddress) {
		body["src-address"] = plan.SrcAddress.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/hotspot/walled-garden", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/hotspot/walled-garden failed", err.Error())
			return
		}
		iPHotspotWalledGardenApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPHotspotWalledGardenResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPHotspotWalledGardenModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/hotspot/walled-garden", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/hotspot/walled-garden failed", err.Error())
	}
}

func (r *IPHotspotWalledGardenResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPHotspotWalledGardenLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/hotspot/walled-garden matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPHotspotWalledGardenLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPHotspotWalledGardenLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/ip/hotspot/walled-garden", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func iPHotspotWalledGardenApply(ctx context.Context, obj client.Object, m *IPHotspotWalledGardenModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["action"]; ok {
		_ = v
		if v != "" {
			m.Action = types.StringValue(v)
		} else {
			m.Action = types.StringNull()
		}
	} else {
		m.Action = types.StringNull()
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
	if v, ok := obj["dst-port"]; ok {
		_ = v
		if v != "" {
			m.DstPort = types.StringValue(v)
		} else {
			m.DstPort = types.StringNull()
		}
	} else {
		m.DstPort = types.StringNull()
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
	if v, ok := obj["server"]; ok {
		_ = v
		if v != "" {
			m.Server = types.StringValue(v)
		} else {
			m.Server = types.StringNull()
		}
	} else {
		m.Server = types.StringNull()
	}
	if v, ok := obj["src-address"]; ok {
		_ = v
		if v != "" {
			m.SrcAddress = types.StringValue(v)
		} else {
			m.SrcAddress = types.StringNull()
		}
	} else {
		m.SrcAddress = types.StringNull()
	}
}
