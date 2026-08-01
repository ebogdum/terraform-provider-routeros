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
	_ resource.Resource                = &MPLSMangleResource{}
	_ resource.ResourceWithImportState = &MPLSMangleResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type MPLSMangleResource struct {
	reg *client.Registry
}

type MPLSMangleModel struct {
	ID               types.String `tfsdk:"id"`
	Builtin          types.Bool   `tfsdk:"builtin"`
	Chain            types.String `tfsdk:"chain"`
	Comment          types.String `tfsdk:"comment"`
	Disabled         types.Bool   `tfsdk:"disabled"`
	Exp              types.String `tfsdk:"exp"`
	Packets          types.String `tfsdk:"packets"`
	ResetCounters    types.String `tfsdk:"reset_counters"`
	ResetCountersAll types.String `tfsdk:"reset_counters_all"`
	SetExp           types.String `tfsdk:"set_exp"`
	SetMark          types.String `tfsdk:"set_mark"`
	Router           types.String `tfsdk:"router"`
}

func NewMPLSMangleResource() resource.Resource { return &MPLSMangleResource{} }

func (r *MPLSMangleResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_mpls_mangle"
}

func (r *MPLSMangleResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *MPLSMangleResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "MPLS mangle schema differs across ROS versions and the audit can't determine the correct argument set without an active LDP. Skipped.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"builtin": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"chain": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "forward", "output"}...)},
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
			"exp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"0", "1", "2", "3", "4", "5", "6", "7"}...)},
			},
			"packets": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"reset_counters": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"reset_counters_all": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"set_exp": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"0", "1", "2", "3", "4", "5", "6", "7"}...)},
			},
			"set_mark": schema.StringAttribute{
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

func (r *MPLSMangleResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan MPLSMangleModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Chain.IsNull() || plan.Chain.IsUnknown()) {
		body["chain"] = plan.Chain.ValueString()
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Exp.IsNull() || plan.Exp.IsUnknown()) {
		body["exp"] = plan.Exp.ValueString()
	}
	if !(plan.SetExp.IsNull() || plan.SetExp.IsUnknown()) {
		body["set-exp"] = plan.SetExp.ValueString()
	}
	if !(plan.SetMark.IsNull() || plan.SetMark.IsUnknown()) {
		body["set-mark"] = plan.SetMark.ValueString()
	}
	obj, err := c.Add(ctx, "/mpls/mangle", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /mpls/mangle failed", err.Error())
		return
	}
	mPLSMangleApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MPLSMangleResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state MPLSMangleModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/mpls/mangle", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /mpls/mangle failed", err.Error())
		return
	}
	mPLSMangleApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *MPLSMangleResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state MPLSMangleModel
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
	if !plan.Chain.Equal(state.Chain) && !plan.Chain.IsUnknown() {
		body["chain"] = plan.Chain.ValueString()
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Exp.Equal(state.Exp) && !plan.Exp.IsUnknown() {
		body["exp"] = plan.Exp.ValueString()
	}
	if !plan.SetExp.Equal(state.SetExp) && !plan.SetExp.IsUnknown() {
		body["set-exp"] = plan.SetExp.ValueString()
	}
	if !plan.SetMark.Equal(state.SetMark) && !plan.SetMark.IsUnknown() {
		body["set-mark"] = plan.SetMark.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/mpls/mangle", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /mpls/mangle failed", err.Error())
			return
		}
		mPLSMangleApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *MPLSMangleResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state MPLSMangleModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/mpls/mangle", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /mpls/mangle failed", err.Error())
	}
}

func (r *MPLSMangleResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := mPLSMangleLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /mpls/mangle matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// mPLSMangleLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func mPLSMangleLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/mpls/mangle", id)
}

func mPLSMangleApply(ctx context.Context, obj client.Object, m *MPLSMangleModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["builtin"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Builtin = types.BoolValue(b)
		} else {
			m.Builtin = types.BoolNull()
		}
	} else {
		m.Builtin = types.BoolNull()
	}
	if v, ok := obj["chain"]; ok {
		_ = v
		if v != "" {
			m.Chain = types.StringValue(v)
		} else {
			m.Chain = types.StringNull()
		}
	} else {
		m.Chain = types.StringNull()
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
	if v, ok := obj["exp"]; ok {
		_ = v
		if v != "" {
			m.Exp = types.StringValue(v)
		} else {
			m.Exp = types.StringNull()
		}
	} else {
		m.Exp = types.StringNull()
	}
	if v, ok := obj["packets"]; ok {
		_ = v
		if v != "" {
			m.Packets = types.StringValue(v)
		} else {
			m.Packets = types.StringNull()
		}
	} else {
		m.Packets = types.StringNull()
	}
	if v, ok := obj["reset-counters"]; ok {
		_ = v
		if v != "" {
			m.ResetCounters = types.StringValue(v)
		} else {
			m.ResetCounters = types.StringNull()
		}
	} else {
		m.ResetCounters = types.StringNull()
	}
	if v, ok := obj["reset-counters-all"]; ok {
		_ = v
		if v != "" {
			m.ResetCountersAll = types.StringValue(v)
		} else {
			m.ResetCountersAll = types.StringNull()
		}
	} else {
		m.ResetCountersAll = types.StringNull()
	}
	if v, ok := obj["set-exp"]; ok {
		_ = v
		if v != "" {
			m.SetExp = types.StringValue(v)
		} else {
			m.SetExp = types.StringNull()
		}
	} else {
		m.SetExp = types.StringNull()
	}
	if v, ok := obj["set-mark"]; ok {
		_ = v
		if v != "" {
			m.SetMark = types.StringValue(v)
		} else {
			m.SetMark = types.StringNull()
		}
	} else {
		m.SetMark = types.StringNull()
	}
}
