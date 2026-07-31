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
	_ resource.Resource                = &CapsManConfigurationResource{}
	_ resource.ResourceWithImportState = &CapsManConfigurationResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type CapsManConfigurationResource struct {
	reg *client.Registry
}

type CapsManConfigurationModel struct {
	ID       types.String `tfsdk:"id"`
	Datapath types.String `tfsdk:"datapath"`
	Security types.String `tfsdk:"security"`
	Rates    types.String `tfsdk:"rates"`
	Channel  types.String `tfsdk:"channel"`
	Comment  types.String `tfsdk:"comment"`
	Country  types.String `tfsdk:"country"`
	Distance types.String `tfsdk:"distance"`
	Mode     types.String `tfsdk:"mode"`
	Name     types.String `tfsdk:"name"`
	Ssid     types.String `tfsdk:"ssid"`
	Router   types.String `tfsdk:"router"`
}

func NewCapsManConfigurationResource() resource.Resource { return &CapsManConfigurationResource{} }

func (r *CapsManConfigurationResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_caps_man_configuration"
}

func (r *CapsManConfigurationResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *CapsManConfigurationResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/caps-man/configuration`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"datapath": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of a `/caps-man/datapath` profile to apply.",
			},
			"security": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of a `/caps-man/security` profile to apply.",
			},
			"rates": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of a `/caps-man/rates` profile to apply.",
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
			"country": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"distance": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"mode": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"ssid": schema.StringAttribute{
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

func (r *CapsManConfigurationResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CapsManConfigurationModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Datapath.IsNull() || plan.Datapath.IsUnknown()) {
		body["datapath"] = plan.Datapath.ValueString()
	}
	if !(plan.Security.IsNull() || plan.Security.IsUnknown()) {
		body["security"] = plan.Security.ValueString()
	}
	if !(plan.Rates.IsNull() || plan.Rates.IsUnknown()) {
		body["rates"] = plan.Rates.ValueString()
	}
	if !(plan.Channel.IsNull() || plan.Channel.IsUnknown()) {
		body["channel"] = plan.Channel.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Country.IsNull() || plan.Country.IsUnknown()) {
		body["country"] = plan.Country.ValueString()
	}
	if !(plan.Distance.IsNull() || plan.Distance.IsUnknown()) {
		body["distance"] = plan.Distance.ValueString()
	}
	if !(plan.Mode.IsNull() || plan.Mode.IsUnknown()) {
		body["mode"] = plan.Mode.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Ssid.IsNull() || plan.Ssid.IsUnknown()) {
		body["ssid"] = plan.Ssid.ValueString()
	}
	obj, err := c.Add(ctx, "/caps-man/configuration", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /caps-man/configuration failed", err.Error())
		return
	}
	capsManConfigurationApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CapsManConfigurationResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CapsManConfigurationModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/caps-man/configuration", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /caps-man/configuration failed", err.Error())
		return
	}
	capsManConfigurationApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CapsManConfigurationResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CapsManConfigurationModel
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
	if !plan.Datapath.Equal(state.Datapath) {
		body["datapath"] = plan.Datapath.ValueString()
	}
	if !plan.Security.Equal(state.Security) {
		body["security"] = plan.Security.ValueString()
	}
	if !plan.Rates.Equal(state.Rates) {
		body["rates"] = plan.Rates.ValueString()
	}
	if !plan.Channel.Equal(state.Channel) {
		body["channel"] = plan.Channel.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Country.Equal(state.Country) {
		body["country"] = plan.Country.ValueString()
	}
	if !plan.Distance.Equal(state.Distance) {
		body["distance"] = plan.Distance.ValueString()
	}
	if !plan.Mode.Equal(state.Mode) {
		body["mode"] = plan.Mode.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Ssid.Equal(state.Ssid) {
		body["ssid"] = plan.Ssid.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/caps-man/configuration", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /caps-man/configuration failed", err.Error())
			return
		}
		capsManConfigurationApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CapsManConfigurationResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CapsManConfigurationModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/caps-man/configuration", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /caps-man/configuration failed", err.Error())
	}
}

func (r *CapsManConfigurationResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := capsManConfigurationLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /caps-man/configuration matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// capsManConfigurationLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func capsManConfigurationLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/caps-man/configuration", id)
}

func capsManConfigurationApply(ctx context.Context, obj client.Object, m *CapsManConfigurationModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["datapath"]; ok && v != "" {
		m.Datapath = types.StringValue(v)
	} else {
		m.Datapath = types.StringNull()
	}
	if v, ok := obj["security"]; ok && v != "" {
		m.Security = types.StringValue(v)
	} else {
		m.Security = types.StringNull()
	}
	if v, ok := obj["rates"]; ok && v != "" {
		m.Rates = types.StringValue(v)
	} else {
		m.Rates = types.StringNull()
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
	if v, ok := obj["country"]; ok {
		_ = v
		if v != "" {
			m.Country = types.StringValue(v)
		} else {
			m.Country = types.StringNull()
		}
	} else {
		m.Country = types.StringNull()
	}
	if v, ok := obj["distance"]; ok {
		_ = v
		if v != "" {
			m.Distance = types.StringValue(v)
		} else {
			m.Distance = types.StringNull()
		}
	} else {
		m.Distance = types.StringNull()
	}
	if v, ok := obj["mode"]; ok {
		_ = v
		if v != "" {
			m.Mode = types.StringValue(v)
		} else {
			m.Mode = types.StringNull()
		}
	} else {
		m.Mode = types.StringNull()
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
	if v, ok := obj["ssid"]; ok {
		_ = v
		if v != "" {
			m.Ssid = types.StringValue(v)
		} else {
			m.Ssid = types.StringNull()
		}
	} else {
		m.Ssid = types.StringNull()
	}
}
