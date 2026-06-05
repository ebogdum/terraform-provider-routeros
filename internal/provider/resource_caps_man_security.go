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
	_ resource.Resource                = &CapsManSecurityResource{}
	_ resource.ResourceWithImportState = &CapsManSecurityResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type CapsManSecurityResource struct {
	reg *client.Registry
}

type CapsManSecurityModel struct {
	ID         types.String `tfsdk:"id"`
	Comment    types.String `tfsdk:"comment"`
	Encryption types.String `tfsdk:"encryption"`
	Name       types.String `tfsdk:"name"`
	Passphrase types.String `tfsdk:"passphrase"`
	Router     types.String `tfsdk:"router"`
}

func NewCapsManSecurityResource() resource.Resource { return &CapsManSecurityResource{} }

func (r *CapsManSecurityResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_caps_man_security"
}

func (r *CapsManSecurityResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *CapsManSecurityResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/caps-man/security`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"encryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"passphrase": schema.StringAttribute{
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

func (r *CapsManSecurityResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CapsManSecurityModel
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
	if !(plan.Encryption.IsNull() || plan.Encryption.IsUnknown()) {
		body["encryption"] = plan.Encryption.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Passphrase.IsNull() || plan.Passphrase.IsUnknown()) {
		body["passphrase"] = plan.Passphrase.ValueString()
	}
	obj, err := c.Add(ctx, "/caps-man/security", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /caps-man/security failed", err.Error())
		return
	}
	capsManSecurityApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CapsManSecurityResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CapsManSecurityModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/caps-man/security", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /caps-man/security failed", err.Error())
		return
	}
	capsManSecurityApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CapsManSecurityResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CapsManSecurityModel
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
	if !plan.Encryption.Equal(state.Encryption) {
		body["encryption"] = plan.Encryption.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Passphrase.Equal(state.Passphrase) {
		body["passphrase"] = plan.Passphrase.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/caps-man/security", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /caps-man/security failed", err.Error())
			return
		}
		capsManSecurityApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CapsManSecurityResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CapsManSecurityModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/caps-man/security", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /caps-man/security failed", err.Error())
	}
}

func (r *CapsManSecurityResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := capsManSecurityLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /caps-man/security matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// capsManSecurityLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func capsManSecurityLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/caps-man/security", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func capsManSecurityApply(ctx context.Context, obj client.Object, m *CapsManSecurityModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
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
	if v, ok := obj["encryption"]; ok {
		_ = v
		if v != "" {
			m.Encryption = types.StringValue(v)
		} else {
			m.Encryption = types.StringNull()
		}
	} else {
		m.Encryption = types.StringNull()
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
	if v, ok := obj["passphrase"]; ok {
		_ = v
		if v != "" {
			m.Passphrase = types.StringValue(v)
		} else {
			m.Passphrase = types.StringNull()
		}
	} else {
		m.Passphrase = types.StringNull()
	}
}
