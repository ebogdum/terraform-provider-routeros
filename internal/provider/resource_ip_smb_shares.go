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
	_ resource.Resource                = &IPSmbSharesResource{}
	_ resource.ResourceWithImportState = &IPSmbSharesResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPSmbSharesResource struct {
	reg *client.Registry
}

type IPSmbSharesModel struct {
	ID                types.String `tfsdk:"id"`
	Comment           types.String `tfsdk:"comment"`
	Default           types.Bool   `tfsdk:"default"`
	Directory         types.String `tfsdk:"directory"`
	Disabled          types.Bool   `tfsdk:"disabled"`
	Dynamic           types.Bool   `tfsdk:"dynamic"`
	InvalidUsers      types.String `tfsdk:"invalid_users"`
	Name              types.String `tfsdk:"name"`
	Newfileman        types.String `tfsdk:"newfileman"`
	OldDirectory      types.String `tfsdk:"old_directory"`
	Oldfileman        types.String `tfsdk:"oldfileman"`
	ReadOnly          types.Bool   `tfsdk:"read_only"`
	RequireEncryption types.Bool   `tfsdk:"require_encryption"`
	ValidUsers        types.String `tfsdk:"valid_users"`
	Router            types.String `tfsdk:"router"`
}

func NewIPSmbSharesResource() resource.Resource { return &IPSmbSharesResource{} }

func (r *IPSmbSharesResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_smb_shares"
}

func (r *IPSmbSharesResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPSmbSharesResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/smb/shares`.",
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
			"default": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"directory": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"invalid_users": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"newfileman": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"old_directory": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"oldfileman": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"read_only": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"require_encryption": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"valid_users": schema.StringAttribute{
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

func (r *IPSmbSharesResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPSmbSharesModel
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
	if !(plan.Directory.IsNull() || plan.Directory.IsUnknown()) {
		body["directory"] = plan.Directory.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.InvalidUsers.IsNull() || plan.InvalidUsers.IsUnknown()) {
		body["invalid-users"] = plan.InvalidUsers.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.ReadOnly.IsNull() || plan.ReadOnly.IsUnknown()) {
		body["read-only"] = client.FormatBool(plan.ReadOnly.ValueBool())
	}
	if !(plan.RequireEncryption.IsNull() || plan.RequireEncryption.IsUnknown()) {
		body["require-encryption"] = client.FormatBool(plan.RequireEncryption.ValueBool())
	}
	if !(plan.ValidUsers.IsNull() || plan.ValidUsers.IsUnknown()) {
		body["valid-users"] = plan.ValidUsers.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/smb/shares", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/smb/shares failed", err.Error())
		return
	}
	iPSmbSharesApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPSmbSharesResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPSmbSharesModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/smb/shares", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/smb/shares failed", err.Error())
		return
	}
	iPSmbSharesApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPSmbSharesResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPSmbSharesModel
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
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Directory.Equal(state.Directory) && !plan.Directory.IsUnknown() {
		body["directory"] = plan.Directory.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.InvalidUsers.Equal(state.InvalidUsers) && !plan.InvalidUsers.IsUnknown() {
		body["invalid-users"] = plan.InvalidUsers.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.ReadOnly.Equal(state.ReadOnly) && !plan.ReadOnly.IsUnknown() {
		body["read-only"] = client.FormatBool(plan.ReadOnly.ValueBool())
	}
	if !plan.RequireEncryption.Equal(state.RequireEncryption) && !plan.RequireEncryption.IsUnknown() {
		body["require-encryption"] = client.FormatBool(plan.RequireEncryption.ValueBool())
	}
	if !plan.ValidUsers.Equal(state.ValidUsers) && !plan.ValidUsers.IsUnknown() {
		body["valid-users"] = plan.ValidUsers.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/smb/shares", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/smb/shares failed", err.Error())
			return
		}
		iPSmbSharesApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPSmbSharesResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPSmbSharesModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/smb/shares", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/smb/shares failed", err.Error())
	}
}

func (r *IPSmbSharesResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPSmbSharesLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/smb/shares matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPSmbSharesLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPSmbSharesLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/smb/shares", id)
}

func iPSmbSharesApply(ctx context.Context, obj client.Object, m *IPSmbSharesModel) {
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
	if v, ok := obj["default"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Default = types.BoolValue(b)
		} else {
			m.Default = types.BoolNull()
		}
	}
	if v, ok := obj["directory"]; ok {
		_ = v
		if v != "" {
			m.Directory = types.StringValue(v)
		} else {
			m.Directory = types.StringNull()
		}
	} else {
		m.Directory = types.StringNull()
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["dynamic"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else {
			m.Dynamic = types.BoolNull()
		}
	}
	if v, ok := obj["invalid-users"]; ok {
		_ = v
		if v != "" {
			m.InvalidUsers = types.StringValue(v)
		} else {
			m.InvalidUsers = types.StringNull()
		}
	} else {
		m.InvalidUsers = types.StringNull()
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
	if v, ok := obj["newfileman"]; ok {
		_ = v
		if v != "" {
			m.Newfileman = types.StringValue(v)
		} else {
			m.Newfileman = types.StringNull()
		}
	} else {
		m.Newfileman = types.StringNull()
	}
	if v, ok := obj["old-directory"]; ok {
		_ = v
		if v != "" {
			m.OldDirectory = types.StringValue(v)
		} else {
			m.OldDirectory = types.StringNull()
		}
	} else {
		m.OldDirectory = types.StringNull()
	}
	if v, ok := obj["oldfileman"]; ok {
		_ = v
		if v != "" {
			m.Oldfileman = types.StringValue(v)
		} else {
			m.Oldfileman = types.StringNull()
		}
	} else {
		m.Oldfileman = types.StringNull()
	}
	if v, ok := obj["read-only"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.ReadOnly = types.BoolValue(b)
		} else {
			m.ReadOnly = types.BoolNull()
		}
	}
	if v, ok := obj["require-encryption"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.RequireEncryption = types.BoolValue(b)
		} else {
			m.RequireEncryption = types.BoolNull()
		}
	}
	if v, ok := obj["valid-users"]; ok {
		_ = v
		if v != "" {
			m.ValidUsers = types.StringValue(v)
		} else {
			m.ValidUsers = types.StringNull()
		}
	} else {
		m.ValidUsers = types.StringNull()
	}
}
