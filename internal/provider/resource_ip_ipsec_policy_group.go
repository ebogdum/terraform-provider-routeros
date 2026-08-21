package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

var (
	_ resource.Resource                = &IPIPsecPolicyGroupResource{}
	_ resource.ResourceWithImportState = &IPIPsecPolicyGroupResource{}
	_                                  = fmt.Sprintf
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPIPsecPolicyGroupResource struct {
	reg *client.Registry
}

type IPIPsecPolicyGroupModel struct {
	ID      types.String `tfsdk:"id"`
	Default types.Bool   `tfsdk:"default"`
	Name    types.String `tfsdk:"name"`
	Router  types.String `tfsdk:"router"`
}

func NewIPIPsecPolicyGroupResource() resource.Resource { return &IPIPsecPolicyGroupResource{} }

func (r *IPIPsecPolicyGroupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_ipsec_policy_group"
}

func (r *IPIPsecPolicyGroupResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPIPsecPolicyGroupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/ipsec/policy/group`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"default": schema.BoolAttribute{
				Computed:    true,
				Description: "RouterOS `default`.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `name`.",
			},
			"router": schema.StringAttribute{
				Optional:      true,
				Description:   "Name of the router (key in provider's `routers` map). Omit to use the default.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
		},
	}
}

func (r *IPIPsecPolicyGroupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPIPsecPolicyGroupModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/ipsec/policy/group", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/ipsec/policy/group failed", err.Error())
		return
	}
	iPIPsecPolicyGroupApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIPsecPolicyGroupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPIPsecPolicyGroupModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/ipsec/policy/group", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/ipsec/policy/group failed", err.Error())
		return
	}
	iPIPsecPolicyGroupApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPIPsecPolicyGroupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPIPsecPolicyGroupModel
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
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	var obj client.Object
	var err error
	if len(body) > 0 {
		obj, err = c.Set(ctx, "/ip/ipsec/policy/group", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/ipsec/policy/group failed", err.Error())
			return
		}
	} else {
		obj, err = c.GetByID(ctx, "/ip/ipsec/policy/group", state.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Read /ip/ipsec/policy/group failed", err.Error())
			return
		}
	}
	iPIPsecPolicyGroupApply(ctx, obj, &plan)
	plan.ID = state.ID
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPIPsecPolicyGroupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPIPsecPolicyGroupModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/ipsec/policy/group", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/ipsec/policy/group failed", err.Error())
	}
}

func (r *IPIPsecPolicyGroupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := lookupByNaturalKey(ctx, c, "/ip/ipsec/policy/group", id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/ipsec/policy/group matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

func iPIPsecPolicyGroupApply(ctx context.Context, obj client.Object, m *IPIPsecPolicyGroupModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["default"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Default = types.BoolValue(true)
		} else {
			m.Default = types.BoolNull()
		}
	} else {
		m.Default = types.BoolNull()
	}
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	} else {
		m.Name = types.StringNull()
	}
}
