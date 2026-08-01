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
	_ resource.Resource                = &FileResource{}
	_ resource.ResourceWithImportState = &FileResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type FileResource struct {
	reg *client.Registry
}

type FileModel struct {
	ID           types.String `tfsdk:"id"`
	Backup       types.String `tfsdk:"backup"`
	Basename     types.String `tfsdk:"basename"`
	Container    types.Int64  `tfsdk:"container"`
	Contents     types.String `tfsdk:"contents"`
	Directory    types.String `tfsdk:"directory"`
	Family       types.Int64  `tfsdk:"family"`
	FileName     types.String `tfsdk:"file_name"`
	FileShareURL types.String `tfsdk:"file_share_url"`
	Hasvpn       types.String `tfsdk:"hasvpn"`
	Invalid      types.String `tfsdk:"invalid"`
	Invalidfile  types.String `tfsdk:"invalidfile"`
	LastModified types.String `tfsdk:"last_modified"`
	Name         types.String `tfsdk:"name"`
	Nondir       types.String `tfsdk:"nondir"`
	Restore      types.String `tfsdk:"restore"`
	Share        types.String `tfsdk:"share"`
	Shared       types.Bool   `tfsdk:"shared"`
	Size         types.String `tfsdk:"size"`
	Type         types.String `tfsdk:"type"`
	Unshare      types.String `tfsdk:"unshare"`
	Valid        types.String `tfsdk:"valid"`
	Router       types.String `tfsdk:"router"`
}

func NewFileResource() resource.Resource { return &FileResource{} }

func (r *FileResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_file"
}

func (r *FileResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *FileResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Creating a file via REST requires writing the contents in a follow-up call; not in the acc-test fast path.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"backup": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"basename": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"container": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"contents": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"directory": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"family": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"file_name": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"file_share_url": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"hasvpn": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"invalid": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"invalidfile": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"last_modified": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"nondir": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"restore": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"share": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"shared": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"size": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"type": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"unshare": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"valid": schema.StringAttribute{
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

func (r *FileResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan FileModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Contents.IsNull() || plan.Contents.IsUnknown()) {
		body["contents"] = plan.Contents.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	obj, err := c.Add(ctx, "/file", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /file failed", err.Error())
		return
	}
	fileApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *FileResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state FileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/file", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /file failed", err.Error())
		return
	}
	fileApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *FileResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state FileModel
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
	if !plan.Contents.Equal(state.Contents) && !plan.Contents.IsUnknown() {
		body["contents"] = plan.Contents.ValueString()
	}
	if !plan.Name.Equal(state.Name) && !plan.Name.IsUnknown() {
		body["name"] = plan.Name.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/file", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /file failed", err.Error())
			return
		}
		fileApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *FileResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state FileModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/file", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /file failed", err.Error())
	}
}

func (r *FileResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := fileLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /file matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// fileLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func fileLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/file", id)
}

func fileApply(ctx context.Context, obj client.Object, m *FileModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["backup"]; ok {
		if v != "" {
			m.Backup = types.StringValue(v)
		} else {
			m.Backup = types.StringNull()
		}
	}
	if v, ok := obj["basename"]; ok {
		if v != "" {
			m.Basename = types.StringValue(v)
		} else {
			m.Basename = types.StringNull()
		}
	}
	if v, ok := obj["container"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Container = types.Int64Value(n)
		} else {
			m.Container = types.Int64Null()
		}
	} else {
		m.Container = types.Int64Null()
	}
	if v, ok := obj["contents"]; ok {
		if v != "" {
			m.Contents = types.StringValue(v)
		} else {
			m.Contents = types.StringNull()
		}
	}
	if v, ok := obj["directory"]; ok {
		if v != "" {
			m.Directory = types.StringValue(v)
		} else {
			m.Directory = types.StringNull()
		}
	}
	if v, ok := obj["family"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Family = types.Int64Value(n)
		} else {
			m.Family = types.Int64Null()
		}
	} else {
		m.Family = types.Int64Null()
	}
	if v, ok := obj["file-name"]; ok {
		if v != "" {
			m.FileName = types.StringValue(v)
		} else {
			m.FileName = types.StringNull()
		}
	}
	if v, ok := obj["file-share-url"]; ok {
		if v != "" {
			m.FileShareURL = types.StringValue(v)
		} else {
			m.FileShareURL = types.StringNull()
		}
	}
	if v, ok := obj["hasvpn"]; ok {
		if v != "" {
			m.Hasvpn = types.StringValue(v)
		} else {
			m.Hasvpn = types.StringNull()
		}
	}
	if v, ok := obj["invalid"]; ok {
		if v != "" {
			m.Invalid = types.StringValue(v)
		} else {
			m.Invalid = types.StringNull()
		}
	}
	if v, ok := obj["invalidfile"]; ok {
		if v != "" {
			m.Invalidfile = types.StringValue(v)
		} else {
			m.Invalidfile = types.StringNull()
		}
	}
	if v, ok := obj["last-modified"]; ok {
		if v != "" {
			m.LastModified = types.StringValue(v)
		} else {
			m.LastModified = types.StringNull()
		}
	}
	if v, ok := obj["name"]; ok {
		if v != "" {
			m.Name = types.StringValue(v)
		} else {
			m.Name = types.StringNull()
		}
	}
	if v, ok := obj["nondir"]; ok {
		if v != "" {
			m.Nondir = types.StringValue(v)
		} else {
			m.Nondir = types.StringNull()
		}
	}
	if v, ok := obj["restore"]; ok {
		if v != "" {
			m.Restore = types.StringValue(v)
		} else {
			m.Restore = types.StringNull()
		}
	}
	if v, ok := obj["share"]; ok {
		if v != "" {
			m.Share = types.StringValue(v)
		} else {
			m.Share = types.StringNull()
		}
	}
	if v, ok := obj["shared"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Shared = types.BoolValue(b)
		} else {
			m.Shared = types.BoolNull()
		}
	}
	if v, ok := obj["size"]; ok {
		if v != "" {
			m.Size = types.StringValue(v)
		} else {
			m.Size = types.StringNull()
		}
	}
	if v, ok := obj["type"]; ok && v != "" {
		m.Type = types.StringValue(v)
	} else {
		m.Type = types.StringNull()
	}
	if v, ok := obj["unshare"]; ok {
		if v != "" {
			m.Unshare = types.StringValue(v)
		} else {
			m.Unshare = types.StringNull()
		}
	}
	if v, ok := obj["valid"]; ok {
		if v != "" {
			m.Valid = types.StringValue(v)
		} else {
			m.Valid = types.StringNull()
		}
	}
}
