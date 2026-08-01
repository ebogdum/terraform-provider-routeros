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
	_ resource.Resource                = &IPKidControlResource{}
	_ resource.ResourceWithImportState = &IPKidControlResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPKidControlResource struct {
	reg *client.Registry
}

type IPKidControlModel struct {
	ID        types.String `tfsdk:"id"`
	Wed       types.String `tfsdk:"wed"`
	TurWed    types.String `tfsdk:"tur_wed"`
	TurTue    types.String `tfsdk:"tur_tue"`
	TurThu    types.String `tfsdk:"tur_thu"`
	TurSun    types.String `tfsdk:"tur_sun"`
	TurSat    types.String `tfsdk:"tur_sat"`
	TurMon    types.String `tfsdk:"tur_mon"`
	TurFri    types.String `tfsdk:"tur_fri"`
	Tue       types.String `tfsdk:"tue"`
	Thu       types.String `tfsdk:"thu"`
	Sun       types.String `tfsdk:"sun"`
	Sat       types.String `tfsdk:"sat"`
	Mon       types.String `tfsdk:"mon"`
	Fri       types.String `tfsdk:"fri"`
	Disabled  types.Bool   `tfsdk:"disabled"`
	Name      types.String `tfsdk:"name"`
	RateLimit types.String `tfsdk:"rate_limit"`
	Router    types.String `tfsdk:"router"`
}

func NewIPKidControlResource() resource.Resource { return &IPKidControlResource{} }

func (r *IPKidControlResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_kid_control"
}

func (r *IPKidControlResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPKidControlResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/kid-control`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"wed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `wed`.",
			},
			"tur_wed": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tur-wed`.",
			},
			"tur_tue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tur-tue`.",
			},
			"tur_thu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tur-thu`.",
			},
			"tur_sun": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tur-sun`.",
			},
			"tur_sat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tur-sat`.",
			},
			"tur_mon": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tur-mon`.",
			},
			"tur_fri": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tur-fri`.",
			},
			"tue": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `tue`.",
			},
			"thu": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `thu`.",
			},
			"sun": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `sun`.",
			},
			"sat": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `sat`.",
			},
			"mon": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `mon`.",
			},
			"fri": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `fri`.",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the entry is disabled.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Name of the Kid's profile",
			},
			"rate_limit": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The maximum available data rate for flow",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *IPKidControlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPKidControlModel
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
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.RateLimit.IsNull() || plan.RateLimit.IsUnknown()) {
		body["rate-limit"] = plan.RateLimit.ValueString()
	}
	if !(plan.Fri.IsNull() || plan.Fri.IsUnknown()) {
		body["fri"] = plan.Fri.ValueString()
	}
	if !(plan.Mon.IsNull() || plan.Mon.IsUnknown()) {
		body["mon"] = plan.Mon.ValueString()
	}
	if !(plan.Sat.IsNull() || plan.Sat.IsUnknown()) {
		body["sat"] = plan.Sat.ValueString()
	}
	if !(plan.Sun.IsNull() || plan.Sun.IsUnknown()) {
		body["sun"] = plan.Sun.ValueString()
	}
	if !(plan.Thu.IsNull() || plan.Thu.IsUnknown()) {
		body["thu"] = plan.Thu.ValueString()
	}
	if !(plan.Tue.IsNull() || plan.Tue.IsUnknown()) {
		body["tue"] = plan.Tue.ValueString()
	}
	if !(plan.TurFri.IsNull() || plan.TurFri.IsUnknown()) {
		body["tur-fri"] = plan.TurFri.ValueString()
	}
	if !(plan.TurMon.IsNull() || plan.TurMon.IsUnknown()) {
		body["tur-mon"] = plan.TurMon.ValueString()
	}
	if !(plan.TurSat.IsNull() || plan.TurSat.IsUnknown()) {
		body["tur-sat"] = plan.TurSat.ValueString()
	}
	if !(plan.TurSun.IsNull() || plan.TurSun.IsUnknown()) {
		body["tur-sun"] = plan.TurSun.ValueString()
	}
	if !(plan.TurThu.IsNull() || plan.TurThu.IsUnknown()) {
		body["tur-thu"] = plan.TurThu.ValueString()
	}
	if !(plan.TurTue.IsNull() || plan.TurTue.IsUnknown()) {
		body["tur-tue"] = plan.TurTue.ValueString()
	}
	if !(plan.TurWed.IsNull() || plan.TurWed.IsUnknown()) {
		body["tur-wed"] = plan.TurWed.ValueString()
	}
	if !(plan.Wed.IsNull() || plan.Wed.IsUnknown()) {
		body["wed"] = plan.Wed.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/kid-control", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/kid-control failed", err.Error())
		return
	}
	iPKidControlApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPKidControlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPKidControlModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/kid-control", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/kid-control failed", err.Error())
		return
	}
	iPKidControlApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPKidControlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPKidControlModel
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
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.RateLimit.Equal(state.RateLimit) && !plan.RateLimit.IsUnknown() {
		body["rate-limit"] = plan.RateLimit.ValueString()
	}
	if !plan.Fri.Equal(state.Fri) && !plan.Fri.IsUnknown() {
		body["fri"] = plan.Fri.ValueString()
	}
	if !plan.Mon.Equal(state.Mon) && !plan.Mon.IsUnknown() {
		body["mon"] = plan.Mon.ValueString()
	}
	if !plan.Sat.Equal(state.Sat) && !plan.Sat.IsUnknown() {
		body["sat"] = plan.Sat.ValueString()
	}
	if !plan.Sun.Equal(state.Sun) && !plan.Sun.IsUnknown() {
		body["sun"] = plan.Sun.ValueString()
	}
	if !plan.Thu.Equal(state.Thu) && !plan.Thu.IsUnknown() {
		body["thu"] = plan.Thu.ValueString()
	}
	if !plan.Tue.Equal(state.Tue) && !plan.Tue.IsUnknown() {
		body["tue"] = plan.Tue.ValueString()
	}
	if !plan.TurFri.Equal(state.TurFri) && !plan.TurFri.IsUnknown() {
		body["tur-fri"] = plan.TurFri.ValueString()
	}
	if !plan.TurMon.Equal(state.TurMon) && !plan.TurMon.IsUnknown() {
		body["tur-mon"] = plan.TurMon.ValueString()
	}
	if !plan.TurSat.Equal(state.TurSat) && !plan.TurSat.IsUnknown() {
		body["tur-sat"] = plan.TurSat.ValueString()
	}
	if !plan.TurSun.Equal(state.TurSun) && !plan.TurSun.IsUnknown() {
		body["tur-sun"] = plan.TurSun.ValueString()
	}
	if !plan.TurThu.Equal(state.TurThu) && !plan.TurThu.IsUnknown() {
		body["tur-thu"] = plan.TurThu.ValueString()
	}
	if !plan.TurTue.Equal(state.TurTue) && !plan.TurTue.IsUnknown() {
		body["tur-tue"] = plan.TurTue.ValueString()
	}
	if !plan.TurWed.Equal(state.TurWed) && !plan.TurWed.IsUnknown() {
		body["tur-wed"] = plan.TurWed.ValueString()
	}
	if !plan.Wed.Equal(state.Wed) && !plan.Wed.IsUnknown() {
		body["wed"] = plan.Wed.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/kid-control", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/kid-control failed", err.Error())
			return
		}
		iPKidControlApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPKidControlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPKidControlModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/kid-control", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/kid-control failed", err.Error())
	}
}

func (r *IPKidControlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPKidControlLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/kid-control matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPKidControlLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPKidControlLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/kid-control", id)
}

func iPKidControlApply(ctx context.Context, obj client.Object, m *IPKidControlModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["wed"]; ok && v != "" {
		m.Wed = types.StringValue(v)
	} else {
		m.Wed = types.StringNull()
	}
	if v, ok := obj["tur-wed"]; ok && v != "" {
		m.TurWed = types.StringValue(v)
	} else {
		m.TurWed = types.StringNull()
	}
	if v, ok := obj["tur-tue"]; ok && v != "" {
		m.TurTue = types.StringValue(v)
	} else {
		m.TurTue = types.StringNull()
	}
	if v, ok := obj["tur-thu"]; ok && v != "" {
		m.TurThu = types.StringValue(v)
	} else {
		m.TurThu = types.StringNull()
	}
	if v, ok := obj["tur-sun"]; ok && v != "" {
		m.TurSun = types.StringValue(v)
	} else {
		m.TurSun = types.StringNull()
	}
	if v, ok := obj["tur-sat"]; ok && v != "" {
		m.TurSat = types.StringValue(v)
	} else {
		m.TurSat = types.StringNull()
	}
	if v, ok := obj["tur-mon"]; ok && v != "" {
		m.TurMon = types.StringValue(v)
	} else {
		m.TurMon = types.StringNull()
	}
	if v, ok := obj["tur-fri"]; ok && v != "" {
		m.TurFri = types.StringValue(v)
	} else {
		m.TurFri = types.StringNull()
	}
	if v, ok := obj["tue"]; ok && v != "" {
		m.Tue = types.StringValue(v)
	} else {
		m.Tue = types.StringNull()
	}
	if v, ok := obj["thu"]; ok && v != "" {
		m.Thu = types.StringValue(v)
	} else {
		m.Thu = types.StringNull()
	}
	if v, ok := obj["sun"]; ok && v != "" {
		m.Sun = types.StringValue(v)
	} else {
		m.Sun = types.StringNull()
	}
	if v, ok := obj["sat"]; ok && v != "" {
		m.Sat = types.StringValue(v)
	} else {
		m.Sat = types.StringNull()
	}
	if v, ok := obj["mon"]; ok && v != "" {
		m.Mon = types.StringValue(v)
	} else {
		m.Mon = types.StringNull()
	}
	if v, ok := obj["fri"]; ok && v != "" {
		m.Fri = types.StringValue(v)
	} else {
		m.Fri = types.StringNull()
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
	if v, ok := obj["rate-limit"]; ok {
		_ = v
		if v != "" {
			m.RateLimit = types.StringValue(v)
		} else {
			m.RateLimit = types.StringNull()
		}
	} else {
		m.RateLimit = types.StringNull()
	}
}
