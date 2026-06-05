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
	_ resource.Resource                = &SystemBackupResource{}
	_ resource.ResourceWithImportState = &SystemBackupResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type SystemBackupResource struct {
	reg *client.Registry
}

type SystemBackupModel struct {
	ID          types.String `tfsdk:"id"`
	DontEncrypt types.String `tfsdk:"dont_encrypt"`
	Encryption  types.String `tfsdk:"encryption"`
	Name        types.String `tfsdk:"name"`
	Password    types.String `tfsdk:"password"`
	Router      types.String `tfsdk:"router"`
}

func NewSystemBackupResource() resource.Resource { return &SystemBackupResource{} }

func (r *SystemBackupResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_system_backup"
}

func (r *SystemBackupResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *SystemBackupResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Action-only menu (save/load); not CRUD",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"dont_encrypt": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Disable backup file encryption. Note that since RouterOS v6.43 without a provided \u00a0 password, \u00a0 the backup file is unencrypted.",
			},
			"encryption": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The encryption algorithm to use for encrypting the backup file. Note that is not considered a secure encryption method and is only available for compatibility reasons with older RouterOS versions.",
			},
			"name": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "The filename for the backup file.",
			},
			"password": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Sensitive:   true,
				Description: "Password for the encrypted backup file. Note that since RouterOS v6.43 without a provided \u00a0 password, \u00a0 the backup file is unencrypted.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *SystemBackupResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan SystemBackupModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.DontEncrypt.IsNull() || plan.DontEncrypt.IsUnknown()) {
		body["dont-encrypt"] = plan.DontEncrypt.ValueString()
	}
	if !(plan.Encryption.IsNull() || plan.Encryption.IsUnknown()) {
		body["encryption"] = plan.Encryption.ValueString()
	}
	if !(plan.Name.IsNull() || plan.Name.IsUnknown()) {
		body["name"] = plan.Name.ValueString()
	}
	if !(plan.Password.IsNull() || plan.Password.IsUnknown()) {
		body["password"] = plan.Password.ValueString()
	}
	obj, err := c.Add(ctx, "/system/backup", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /system/backup failed", err.Error())
		return
	}
	systemBackupApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemBackupResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state SystemBackupModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/system/backup", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /system/backup failed", err.Error())
		return
	}
	systemBackupApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *SystemBackupResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state SystemBackupModel
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
	if !plan.DontEncrypt.Equal(state.DontEncrypt) {
		body["dont-encrypt"] = plan.DontEncrypt.ValueString()
	}
	if !plan.Encryption.Equal(state.Encryption) {
		body["encryption"] = plan.Encryption.ValueString()
	}
	if !plan.Name.Equal(state.Name) {
		body["name"] = plan.Name.ValueString()
	}
	if !plan.Password.Equal(state.Password) {
		body["password"] = plan.Password.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/system/backup", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /system/backup failed", err.Error())
			return
		}
		systemBackupApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *SystemBackupResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state SystemBackupModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/system/backup", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /system/backup failed", err.Error())
	}
}

func (r *SystemBackupResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := systemBackupLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /system/backup matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// systemBackupLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func systemBackupLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	keys := []string{}
	if len(keys) == 0 {
		keys = []string{"name"}
	}
	for _, k := range keys {
		rows, err := c.List(ctx, "/system/backup", client.WithFilter(k, id))
		if err != nil {
			return nil, err
		}
		if len(rows) > 0 {
			return rows, nil
		}
	}
	return nil, nil
}

func systemBackupApply(ctx context.Context, obj client.Object, m *SystemBackupModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["dont-encrypt"]; ok {
		_ = v
		if v != "" {
			m.DontEncrypt = types.StringValue(v)
		} else {
			m.DontEncrypt = types.StringNull()
		}
	} else {
		m.DontEncrypt = types.StringNull()
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
	// Sensitive: RouterOS scrubs the value on read. If the server returned
	// a value, decode it. Otherwise the plan value (user input) is what's
	// in m.Password already -- but if the user left it unset, resolve
	// the unknown to null so the framework accepts the state.
	if v, ok := obj["password"]; ok && v != "" {
		_ = v
		if v != "" {
			m.Password = types.StringValue(v)
		} else {
			m.Password = types.StringNull()
		}
	} else if m.Password.IsUnknown() {
		m.Password = types.StringNull()
	}
}
