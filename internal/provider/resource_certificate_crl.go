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
	_ resource.Resource                = &CertificateCrlResource{}
	_ resource.ResourceWithImportState = &CertificateCrlResource{}
	_                                  = attr.Value(nil)
	_                                  = strings.TrimSpace
	_                                  = path.Root
)

type CertificateCrlResource struct {
	reg *client.Registry
}

type CertificateCrlModel struct {
	ID          types.String `tfsdk:"id"`
	Akid        types.String `tfsdk:"akid"`
	Certificate types.String `tfsdk:"certificate"`
	Download    types.String `tfsdk:"download"`
	Dynamic     types.Bool   `tfsdk:"dynamic"`
	Expired     types.Bool   `tfsdk:"expired"`
	Flush       types.String `tfsdk:"flush"`
	Invalid     types.Bool   `tfsdk:"invalid"`
	LastUpdate  types.String `tfsdk:"last_update"`
	NextUpdate  types.String `tfsdk:"next_update"`
	Num         types.Int64  `tfsdk:"num"`
	Revoked     types.Int64  `tfsdk:"revoked"`
	Signature   types.String `tfsdk:"signature"`
	URL         types.String `tfsdk:"url"`
	Router      types.String `tfsdk:"router"`
}

func NewCertificateCrlResource() resource.Resource { return &CertificateCrlResource{} }

func (r *CertificateCrlResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_certificate_crl"
}

func (r *CertificateCrlResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *CertificateCrlResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "Mirrors RouterOS `/certificate/crl`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"akid": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"certificate": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"download": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"dynamic": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"expired": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"flush": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"invalid": schema.BoolAttribute{
				Computed:    true,
				Description: "",
			},
			"last_update": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"next_update": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"num": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"revoked": schema.Int64Attribute{
				Computed:    true,
				Description: "",
			},
			"signature": schema.StringAttribute{
				Computed:    true,
				Description: "",
			},
			"url": schema.StringAttribute{
				Required:    true,
				Description: "",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

func (r *CertificateCrlResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan CertificateCrlModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	body := client.Object{}
	if !(plan.URL.IsNull() || plan.URL.IsUnknown()) {
		body["url"] = plan.URL.ValueString()
	}
	obj, err := c.Add(ctx, "/certificate/crl", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /certificate/crl failed", err.Error())
		return
	}
	certificateCrlApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CertificateCrlResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state CertificateCrlModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/certificate/crl", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /certificate/crl failed", err.Error())
		return
	}
	certificateCrlApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *CertificateCrlResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state CertificateCrlModel
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
	if !plan.URL.Equal(state.URL) && !plan.URL.IsUnknown() {
		body["url"] = plan.URL.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/certificate/crl", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /certificate/crl failed", err.Error())
			return
		}
		certificateCrlApply(ctx, obj, &plan)
	} else {
		plan.ID = state.ID
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *CertificateCrlResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state CertificateCrlModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/certificate/crl", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /certificate/crl failed", err.Error())
	}
}

func (r *CertificateCrlResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
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
	rows, err := certificateCrlLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /certificate/crl matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// certificateCrlLookupByNaturalKey searches for a record whose natural
// keys match id. The strategy: try every key declared in the schema overlay's
// natural_keys list (or fall back to "name") with equality matching.
func certificateCrlLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	return lookupByNaturalKey(ctx, c, "/certificate/crl", id)
}

func certificateCrlApply(ctx context.Context, obj client.Object, m *CertificateCrlModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["akid"]; ok {
		if v != "" {
			m.Akid = types.StringValue(v)
		} else {
			m.Akid = types.StringNull()
		}
	}
	if v, ok := obj["certificate"]; ok {
		if v != "" {
			m.Certificate = types.StringValue(v)
		} else {
			m.Certificate = types.StringNull()
		}
	}
	if v, ok := obj["download"]; ok {
		if v != "" {
			m.Download = types.StringValue(v)
		} else {
			m.Download = types.StringNull()
		}
	}
	if v, ok := obj["dynamic"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Dynamic = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Dynamic = types.BoolValue(true)
		} else {
			m.Dynamic = types.BoolNull()
		}
	}
	if v, ok := obj["expired"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Expired = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Expired = types.BoolValue(true)
		} else {
			m.Expired = types.BoolNull()
		}
	}
	if v, ok := obj["flush"]; ok {
		if v != "" {
			m.Flush = types.StringValue(v)
		} else {
			m.Flush = types.StringNull()
		}
	}
	if v, ok := obj["invalid"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Invalid = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Invalid = types.BoolValue(true)
		} else {
			m.Invalid = types.BoolNull()
		}
	}
	if v, ok := obj["last-update"]; ok {
		if v != "" {
			m.LastUpdate = types.StringValue(v)
		} else {
			m.LastUpdate = types.StringNull()
		}
	}
	if v, ok := obj["next-update"]; ok {
		if v != "" {
			m.NextUpdate = types.StringValue(v)
		} else {
			m.NextUpdate = types.StringNull()
		}
	}
	if v, ok := obj["num"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Num = types.Int64Value(n)
		} else {
			m.Num = types.Int64Null()
		}
	} else {
		m.Num = types.Int64Null()
	}
	if v, ok := obj["revoked"]; ok {
		_ = v
		if n, err := client.ParseInt64(v); err == nil {
			m.Revoked = types.Int64Value(n)
		} else {
			m.Revoked = types.Int64Null()
		}
	} else {
		m.Revoked = types.Int64Null()
	}
	if v, ok := obj["signature"]; ok {
		if v != "" {
			m.Signature = types.StringValue(v)
		} else {
			m.Signature = types.StringNull()
		}
	}
	if v, ok := obj["url"]; ok {
		if v != "" {
			m.URL = types.StringValue(v)
		} else {
			m.URL = types.StringNull()
		}
	}
}
