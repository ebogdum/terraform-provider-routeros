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
	_ resource.Resource                = &IPCloudBackToHomeUserResource{}
	_ resource.ResourceWithImportState = &IPCloudBackToHomeUserResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPCloudBackToHomeUserResource struct {
	reg *client.Registry
}

type IPCloudBackToHomeUserModel struct {
	ID              types.String `tfsdk:"id"`
	FileAccessPath  types.String `tfsdk:"file_access_path"`
	FileAccess      types.String `tfsdk:"file_access"`
	Active          types.Bool   `tfsdk:"active"`
	AllowLan        types.Bool   `tfsdk:"allow_lan"`
	ClientAddress   types.String `tfsdk:"client_address"`
	ClientConfig    types.String `tfsdk:"client_config"`
	ClientQr        types.String `tfsdk:"client_qr"`
	Comment         types.String `tfsdk:"comment"`
	Disabled        types.Bool   `tfsdk:"disabled"`
	Expires         types.String `tfsdk:"expires"`
	FileAccessMode  types.String `tfsdk:"file_access_mode"`
	FileAccessToken types.String `tfsdk:"file_access_token"`
	Files           types.String `tfsdk:"files"`
	Name            types.String `tfsdk:"name"`
	Newe            types.String `tfsdk:"newe"`
	Newfileman      types.String `tfsdk:"newfileman"`
	Notnew          types.String `tfsdk:"notnew"`
	Oldfileman      types.String `tfsdk:"oldfileman"`
	PrivateKey      types.String `tfsdk:"private_key"`
	PublicKey       types.String `tfsdk:"public_key"`
	Router          types.String `tfsdk:"router"`
}

func NewIPCloudBackToHomeUserResource() resource.Resource { return &IPCloudBackToHomeUserResource{} }

func (r *IPCloudBackToHomeUserResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_cloud_back_to_home_user"
}

func (r *IPCloudBackToHomeUserResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPCloudBackToHomeUserResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/cloud/back-to-home-user`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"file_access_path": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `file-access-path`.",
			},
			"file_access": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "RouterOS `file-access`.",
			},
			"active": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"allow_lan": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"client_address": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"client_config": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"client_qr": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"expires": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"file_access_mode": schema.StringAttribute{
				Computed:    true,
				Description: "",
				Validators:  []validator.String{schemautil.OneOf([]string{"", "disabled", "read-only", "full"}...)},
			},
			"file_access_token": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"files": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"newe": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"newfileman": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"notnew": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"oldfileman": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"private_key": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "",
			},
			"public_key": schema.StringAttribute{
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

func (r *IPCloudBackToHomeUserResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPCloudBackToHomeUserModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.AllowLan.IsNull() || plan.AllowLan.IsUnknown()) {
		body["allow-lan"] = client.FormatBool(plan.AllowLan.ValueBool())
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Expires.IsNull() || plan.Expires.IsUnknown()) {
		body["expires"] = plan.Expires.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.PrivateKey.IsNull() || plan.PrivateKey.IsUnknown()) {
		body["private-key"] = plan.PrivateKey.ValueString()
	}
	if !(plan.PublicKey.IsNull() || plan.PublicKey.IsUnknown()) {
		body["public-key"] = plan.PublicKey.ValueString()
	}
	if !(plan.FileAccess.IsNull() || plan.FileAccess.IsUnknown()) {
		body["file-access"] = plan.FileAccess.ValueString()
	}
	if !(plan.FileAccessPath.IsNull() || plan.FileAccessPath.IsUnknown()) {
		body["file-access-path"] = plan.FileAccessPath.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/cloud/back-to-home-user", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/cloud/back-to-home-user failed", err.Error())
		return
	}
	iPCloudBackToHomeUserApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPCloudBackToHomeUserResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPCloudBackToHomeUserModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/cloud/back-to-home-user", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/cloud/back-to-home-user failed", err.Error())
		return
	}
	iPCloudBackToHomeUserApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPCloudBackToHomeUserResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPCloudBackToHomeUserModel
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
	if !plan.AllowLan.Equal(state.AllowLan) && !plan.AllowLan.IsUnknown() {
		body["allow-lan"] = client.FormatBool(plan.AllowLan.ValueBool())
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Expires.Equal(state.Expires) && !plan.Expires.IsUnknown() {
		body["expires"] = plan.Expires.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.PrivateKey.Equal(state.PrivateKey) && !plan.PrivateKey.IsUnknown() {
		body["private-key"] = plan.PrivateKey.ValueString()
	}
	if !plan.PublicKey.Equal(state.PublicKey) && !plan.PublicKey.IsUnknown() {
		body["public-key"] = plan.PublicKey.ValueString()
	}
	if !plan.FileAccess.Equal(state.FileAccess) && !plan.FileAccess.IsUnknown() {
		body["file-access"] = plan.FileAccess.ValueString()
	}
	if !plan.FileAccessPath.Equal(state.FileAccessPath) && !plan.FileAccessPath.IsUnknown() {
		body["file-access-path"] = plan.FileAccessPath.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/cloud/back-to-home-user", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/cloud/back-to-home-user failed", err.Error())
			return
		}
		iPCloudBackToHomeUserApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPCloudBackToHomeUserResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPCloudBackToHomeUserModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/cloud/back-to-home-user", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/cloud/back-to-home-user failed", err.Error())
	}
}

func (r *IPCloudBackToHomeUserResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPCloudBackToHomeUserLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/cloud/back-to-home-user matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPCloudBackToHomeUserLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPCloudBackToHomeUserLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/cloud/back-to-home-user", id)
}

func iPCloudBackToHomeUserApply(ctx context.Context, obj client.Object, m *IPCloudBackToHomeUserModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["file-access-path"]; ok && v != "" {
		m.FileAccessPath = types.StringValue(v)
	} else {
		m.FileAccessPath = types.StringNull()
	}
	if v, ok := obj["file-access"]; ok && v != "" {
		m.FileAccess = types.StringValue(v)
	} else {
		m.FileAccess = types.StringNull()
	}
	if v, ok := obj["active"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Active = types.BoolValue(b)
		} else {
			m.Active = types.BoolNull()
		}
	}
	if v, ok := obj["allow-lan"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.AllowLan = types.BoolValue(b)
		} else {
			m.AllowLan = types.BoolNull()
		}
	}
	if v, ok := obj["client-address"]; ok {
		if v != "" {
			m.ClientAddress = types.StringValue(v)
		} else {
			m.ClientAddress = types.StringNull()
		}
	}
	if v, ok := obj["client-config"]; ok {
		if v != "" {
			m.ClientConfig = types.StringValue(v)
		} else {
			m.ClientConfig = types.StringNull()
		}
	}
	if v, ok := obj["client-qr"]; ok {
		if v != "" {
			m.ClientQr = types.StringValue(v)
		} else {
			m.ClientQr = types.StringNull()
		}
	}
	if v, ok := obj["comment"]; ok {
		if v != "" {
			m.Comment = types.StringValue(v)
		} else {
			m.Comment = types.StringNull()
		}
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	}
	if v, ok := obj["expires"]; ok {
		if v != "" {
			m.Expires = types.StringValue(v)
		} else {
			m.Expires = types.StringNull()
		}
	}
	if v, ok := obj["file-access-mode"]; ok {
		if v != "" {
			m.FileAccessMode = types.StringValue(v)
		} else {
			m.FileAccessMode = types.StringNull()
		}
	}
	if v, ok := obj["file-access-token"]; ok {
		if v != "" {
			m.FileAccessToken = types.StringValue(v)
		} else {
			m.FileAccessToken = types.StringNull()
		}
	}
	if v, ok := obj["files"]; ok {
		if v != "" {
			m.Files = types.StringValue(v)
		} else {
			m.Files = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["newe"]; ok {
		if v != "" {
			m.Newe = types.StringValue(v)
		} else {
			m.Newe = types.StringNull()
		}
	}
	if v, ok := obj["newfileman"]; ok {
		if v != "" {
			m.Newfileman = types.StringValue(v)
		} else {
			m.Newfileman = types.StringNull()
		}
	}
	if v, ok := obj["notnew"]; ok {
		if v != "" {
			m.Notnew = types.StringValue(v)
		} else {
			m.Notnew = types.StringNull()
		}
	}
	if v, ok := obj["oldfileman"]; ok {
		if v != "" {
			m.Oldfileman = types.StringValue(v)
		} else {
			m.Oldfileman = types.StringNull()
		}
	}
	if v, ok := obj["private-key"]; ok {
		if v != "" {
			m.PrivateKey = types.StringValue(v)
		} else {
			m.PrivateKey = types.StringNull()
		}
	}
	if v, ok := obj["public-key"]; ok {
		if v != "" {
			m.PublicKey = types.StringValue(v)
		} else {
			m.PublicKey = types.StringNull()
		}
	}
}
