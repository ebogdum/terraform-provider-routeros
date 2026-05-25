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
	_ resource.Resource                = &IPTftpResource{}
	_ resource.ResourceWithImportState = &IPTftpResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type IPTftpResource struct {
	reg *client.Registry
}

type IPTftpModel struct {
	ID           types.String `tfsdk:"id"`
	Allow        types.Bool   `tfsdk:"allow"`
	Comment      types.String `tfsdk:"comment"`
	Disabled     types.Bool   `tfsdk:"disabled"`
	Hits         types.Int64  `tfsdk:"hits"`
	IPAddresses  types.String `tfsdk:"ip_addresses"`
	ReadOnly     types.Bool   `tfsdk:"read_only"`
	RealFilename types.String `tfsdk:"real_filename"`
	ReqFilename  types.String `tfsdk:"req_filename"`
	Router       types.String `tfsdk:"router"`
}

func NewIPTftpResource() resource.Resource { return &IPTftpResource{} }

func (r *IPTftpResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_tftp"
}

func (r *IPTftpResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
	_ = fmt.Sprintf
}

func (r *IPTftpResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/ip/tftp`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"allow": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
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
			"hits": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"ip_addresses": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"read_only": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"real_filename": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "",
			},
			"req_filename": schema.StringAttribute{
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

func (r *IPTftpResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPTftpModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.Allow.IsNull() || plan.Allow.IsUnknown()) {
		body["allow"] = client.FormatBool(plan.Allow.ValueBool())
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.IPAddresses.IsNull() || plan.IPAddresses.IsUnknown()) {
		body["ip-addresses"] = plan.IPAddresses.ValueString()
	}
	if !(plan.ReadOnly.IsNull() || plan.ReadOnly.IsUnknown()) {
		body["read-only"] = client.FormatBool(plan.ReadOnly.ValueBool())
	}
	if !(plan.RealFilename.IsNull() || plan.RealFilename.IsUnknown()) {
		body["real-filename"] = plan.RealFilename.ValueString()
	}
	if !(plan.ReqFilename.IsNull() || plan.ReqFilename.IsUnknown()) {
		body["req-filename"] = plan.ReqFilename.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/tftp", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/tftp failed", err.Error())
		return
	}
	iPTftpApply(ctx, obj, &plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPTftpResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPTftpModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/tftp", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/tftp failed", err.Error())
		return
	}
	iPTftpApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPTftpResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPTftpModel
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
	if !plan.Allow.Equal(state.Allow) {
		body["allow"] = client.FormatBool(plan.Allow.ValueBool())
	}
	if !plan.Comment.Equal(state.Comment) {
		body["comment"] = plan.Comment.ValueString()
	}
	if !plan.Disabled.Equal(state.Disabled) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.IPAddresses.Equal(state.IPAddresses) {
		body["ip-addresses"] = plan.IPAddresses.ValueString()
	}
	if !plan.ReadOnly.Equal(state.ReadOnly) {
		body["read-only"] = client.FormatBool(plan.ReadOnly.ValueBool())
	}
	if !plan.RealFilename.Equal(state.RealFilename) {
		body["real-filename"] = plan.RealFilename.ValueString()
	}
	if !plan.ReqFilename.Equal(state.ReqFilename) {
		body["req-filename"] = plan.ReqFilename.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/tftp", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/tftp failed", err.Error())
			return
		}
		iPTftpApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPTftpResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPTftpModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/tftp", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/tftp failed", err.Error())
	}
}

func (r *IPTftpResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := iPTftpLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/tftp matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPTftpLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func iPTftpLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/ip/tftp", id)
}

func iPTftpApply(ctx context.Context, obj client.Object, m *IPTftpModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["allow"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.Allow = types.BoolValue(b)
		} else {
			m.Allow = types.BoolNull()
		}
	} else {
		m.Allow = types.BoolNull()
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
	if v, ok := obj["hits"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Hits = types.Int64Value(n)
		} else {
			m.Hits = types.Int64Null()
		}
	} else {
		m.Hits = types.Int64Null()
	}
	if v, ok := obj["ip-addresses"]; ok {
		_ = v
		if v != "" {
			m.IPAddresses = types.StringValue(v)
		} else {
			m.IPAddresses = types.StringNull()
		}
	} else {
		m.IPAddresses = types.StringNull()
	}
	if v, ok := obj["read-only"]; ok {
		_ = v
		if b, err := client.ParseBool(v); err == nil {
			m.ReadOnly = types.BoolValue(b)
		} else {
			m.ReadOnly = types.BoolNull()
		}
	} else {
		m.ReadOnly = types.BoolNull()
	}
	if v, ok := obj["real-filename"]; ok {
		_ = v
		if v != "" {
			m.RealFilename = types.StringValue(v)
		} else {
			m.RealFilename = types.StringNull()
		}
	} else {
		m.RealFilename = types.StringNull()
	}
	if v, ok := obj["req-filename"]; ok {
		_ = v
		if v != "" {
			m.ReqFilename = types.StringValue(v)
		} else {
			m.ReqFilename = types.StringNull()
		}
	} else {
		m.ReqFilename = types.StringNull()
	}
}
