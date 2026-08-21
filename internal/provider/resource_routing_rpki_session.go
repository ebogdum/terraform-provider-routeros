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
	_ resource.Resource                = &RoutingRpkiSessionResource{}
	_ resource.ResourceWithImportState = &RoutingRpkiSessionResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type RoutingRpkiSessionResource struct {
	reg *client.Registry
}

type RoutingRpkiSessionModel struct {
	ID      types.String  `tfsdk:"id"`
	Address types.String  `tfsdk:"address"`
	Expires durationValue `tfsdk:"expires"`
	Group   types.String  `tfsdk:"group"`
	Port    types.Int64   `tfsdk:"port"`
	Serial  types.Int64   `tfsdk:"serial"`
	Session types.Int64   `tfsdk:"session"`
	State   types.String  `tfsdk:"state"`
	Version types.Int64   `tfsdk:"version"`
	Router  types.String  `tfsdk:"router"`
}

func NewRoutingRpkiSessionResource() resource.Resource { return &RoutingRpkiSessionResource{} }

func (r *RoutingRpkiSessionResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_routing_rpki_session"
}

func (r *RoutingRpkiSessionResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *RoutingRpkiSessionResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Discovered; needs rpki backend",
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
			"expires": schema.StringAttribute{
				CustomType:  durationType{},
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.IsDurationRouterOS()},
			},
			"group": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"serial": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"session": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"state": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"idle", "connecting", "prepare", "loading", "sync", "error"}...)},
			},
			"version": schema.Int64Attribute{
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

func (r *RoutingRpkiSessionResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan RoutingRpkiSessionModel
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
	if !(plan.Expires.IsNull() || plan.Expires.IsUnknown()) {
		body["expires"] = plan.Expires.ValueString()
	}
	if !(plan.Group.IsNull() || plan.Group.IsUnknown()) {
		body["group"] = plan.Group.ValueString()
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = client.FormatInt64(plan.Port.ValueInt64())
	}
	if !(plan.Serial.IsNull() || plan.Serial.IsUnknown()) {
		body["serial"] = client.FormatInt64(plan.Serial.ValueInt64())
	}
	if !(plan.Session.IsNull() || plan.Session.IsUnknown()) {
		body["session"] = client.FormatInt64(plan.Session.ValueInt64())
	}
	if !(plan.State.IsNull() || plan.State.IsUnknown()) {
		body["state"] = plan.State.ValueString()
	}
	if !(plan.Version.IsNull() || plan.Version.IsUnknown()) {
		body["version"] = client.FormatInt64(plan.Version.ValueInt64())
	}
	obj, err := c.Add(ctx, "/routing/rpki/session", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /routing/rpki/session failed", err.Error())
		return
	}
	routingRpkiSessionApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingRpkiSessionResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state RoutingRpkiSessionModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/routing/rpki/session", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /routing/rpki/session failed", err.Error())
		return
	}
	routingRpkiSessionApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *RoutingRpkiSessionResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state RoutingRpkiSessionModel
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
	if !plan.Expires.Equal(state.Expires) && !plan.Expires.IsUnknown() {
		body["expires"] = plan.Expires.ValueString()
	}
	if !plan.Group.Equal(state.Group) && !plan.Group.IsUnknown() {
		body["group"] = plan.Group.ValueString()
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = client.FormatInt64(plan.Port.ValueInt64())
	}
	if !plan.Serial.Equal(state.Serial) && !plan.Serial.IsUnknown() {
		body["serial"] = client.FormatInt64(plan.Serial.ValueInt64())
	}
	if !plan.Session.Equal(state.Session) && !plan.Session.IsUnknown() {
		body["session"] = client.FormatInt64(plan.Session.ValueInt64())
	}
	if !plan.State.Equal(state.State) && !plan.State.IsUnknown() {
		body["state"] = plan.State.ValueString()
	}
	if !plan.Version.Equal(state.Version) && !plan.Version.IsUnknown() {
		body["version"] = client.FormatInt64(plan.Version.ValueInt64())
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/routing/rpki/session", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /routing/rpki/session failed", err.Error())
			return
		}
		routingRpkiSessionApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *RoutingRpkiSessionResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state RoutingRpkiSessionModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/routing/rpki/session", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /routing/rpki/session failed", err.Error())
	}
}

func (r *RoutingRpkiSessionResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := routingRpkiSessionLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /routing/rpki/session matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// routingRpkiSessionLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func routingRpkiSessionLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/routing/rpki/session", id)
}

func routingRpkiSessionApply(ctx context.Context, obj client.Object, m *RoutingRpkiSessionModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["address"]; ok {
		if v != "" {
			m.Address = types.StringValue(v)
		} else {
			m.Address = types.StringNull()
		}
	}
	if v, ok := obj["expires"]; ok {
		if v != "" {
			m.Expires = newDurationValue(v)
		} else {
			m.Expires = newDurationNull()
		}
	}
	if v, ok := obj["group"]; ok {
		if v != "" {
			m.Group = types.StringValue(v)
		} else {
			m.Group = types.StringNull()
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
	if v, ok := obj["serial"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Serial = types.Int64Value(n)
		} else {
			m.Serial = types.Int64Null()
		}
	} else {
		m.Serial = types.Int64Null()
	}
	if v, ok := obj["session"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Session = types.Int64Value(n)
		} else {
			m.Session = types.Int64Null()
		}
	} else {
		m.Session = types.Int64Null()
	}
	if v, ok := obj["state"]; ok {
		if v != "" {
			m.State = types.StringValue(v)
		} else {
			m.State = types.StringNull()
		}
	}
	if v, ok := obj["version"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Version = types.Int64Value(n)
		} else {
			m.Version = types.Int64Null()
		}
	} else {
		m.Version = types.Int64Null()
	}
}
