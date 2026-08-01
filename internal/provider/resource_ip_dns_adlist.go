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
	_ resource.Resource                   = &IPDNSAdlistResource{}
	_ resource.ResourceWithImportState    = &IPDNSAdlistResource{}
	_ resource.ResourceWithValidateConfig = &IPDNSAdlistResource{}
)

// IPDNSAdlistResource manages one entry of /ip/dns/adlist -- a DNS blocklist
// pulled from a URL or read from a file on the router (RouterOS 7.15+).
type IPDNSAdlistResource struct {
	reg *client.Registry
}

type IPDNSAdlistModel struct {
	ID         types.String `tfsdk:"id"`
	URL        types.String `tfsdk:"url"`
	File       types.String `tfsdk:"file"`
	SSLVerify  types.Bool   `tfsdk:"ssl_verify"`
	Disabled   types.Bool   `tfsdk:"disabled"`
	Comment    types.String `tfsdk:"comment"`
	MatchCount types.Int64  `tfsdk:"match_count"`
	NameCount  types.Int64  `tfsdk:"name_count"`
	Router     types.String `tfsdk:"router"`
}

func NewIPDNSAdlistResource() resource.Resource { return &IPDNSAdlistResource{} }

func (r *IPDNSAdlistResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_dns_adlist"
}

func (r *IPDNSAdlistResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPDNSAdlistResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "A DNS adlist (blocklist) entry in RouterOS `/ip/dns/adlist`. Requires RouterOS 7.15 or newer. " +
			"Supply exactly one of `url` (downloaded and refreshed by the router) or `file` (a hosts-format file " +
			"already stored on the router).",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"url": schema.StringAttribute{
				Optional:    true,
				Description: "HTTP(S) URL of a hosts-format or adblock-format blocklist. Mutually exclusive with `file`.",
			},
			"file": schema.StringAttribute{
				Optional:    true,
				Description: "Name of a blocklist file already present on the router. Mutually exclusive with `url`.",
			},
			"ssl_verify": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Verify the TLS certificate when downloading `url`.",
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the adlist entry is disabled.",
			},
			"comment": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Free-form comment.",
			},
			"match_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Read-only: number of queries blocked by this list.",
			},
			"name_count": schema.Int64Attribute{
				Computed:    true,
				Description: "Read-only: number of names loaded from this list.",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
		},
	}
}

// ValidateConfig enforces the url/file exclusivity RouterOS itself requires, so
// the failure surfaces at plan time rather than as an opaque API error.
func (r *IPDNSAdlistResource) ValidateConfig(ctx context.Context, req resource.ValidateConfigRequest, resp *resource.ValidateConfigResponse) {
	var cfg IPDNSAdlistModel
	if diags := req.Config.Get(ctx, &cfg); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	hasURL := !cfg.URL.IsNull() && !cfg.URL.IsUnknown() && cfg.URL.ValueString() != ""
	hasFile := !cfg.File.IsNull() && !cfg.File.IsUnknown() && cfg.File.ValueString() != ""
	// Unknown values are only resolvable at apply time; don't guess.
	if cfg.URL.IsUnknown() || cfg.File.IsUnknown() {
		return
	}
	if hasURL && hasFile {
		resp.Diagnostics.AddError("Conflicting adlist source",
			"set either `url` or `file` on routeros_ip_dns_adlist, not both")
		return
	}
	if !hasURL && !hasFile {
		resp.Diagnostics.AddError("Missing adlist source",
			"routeros_ip_dns_adlist requires one of `url` or `file`")
	}
}

func (r *IPDNSAdlistResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPDNSAdlistModel
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
	if !(plan.File.IsNull() || plan.File.IsUnknown()) {
		body["file"] = plan.File.ValueString()
	}
	if !(plan.SSLVerify.IsNull() || plan.SSLVerify.IsUnknown()) {
		body["ssl-verify"] = client.FormatBool(plan.SSLVerify.ValueBool())
	}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Comment.IsNull() || plan.Comment.IsUnknown()) {
		body["comment"] = plan.Comment.ValueString()
	}
	obj, err := c.Add(ctx, "/ip/dns/adlist", body)
	if err != nil {
		resp.Diagnostics.AddError("Create /ip/dns/adlist failed", err.Error())
		return
	}
	iPDNSAdlistApply(ctx, obj, &plan)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDNSAdlistResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPDNSAdlistModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/dns/adlist", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/dns/adlist failed", err.Error())
		return
	}
	iPDNSAdlistApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPDNSAdlistResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPDNSAdlistModel
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
	if !plan.File.Equal(state.File) && !plan.File.IsUnknown() {
		body["file"] = plan.File.ValueString()
	}
	if !plan.SSLVerify.Equal(state.SSLVerify) && !plan.SSLVerify.IsUnknown() {
		body["ssl-verify"] = client.FormatBool(plan.SSLVerify.ValueBool())
	}
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Comment.Equal(state.Comment) && !plan.Comment.IsUnknown() {
		body["comment"] = plan.Comment.ValueString()
	}
	if len(body) > 0 {
		obj, err := c.Set(ctx, "/ip/dns/adlist", state.ID.ValueString(), body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/dns/adlist failed", err.Error())
			return
		}
		iPDNSAdlistApply(ctx, obj, &plan)
	} else {
		obj, err := c.GetByID(ctx, "/ip/dns/adlist", state.ID.ValueString())
		if err != nil {
			resp.Diagnostics.AddError("Read /ip/dns/adlist failed", err.Error())
			return
		}
		iPDNSAdlistApply(ctx, obj, &plan)
	}
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPDNSAdlistResource) Delete(ctx context.Context, req resource.DeleteRequest, resp *resource.DeleteResponse) {
	var state IPDNSAdlistModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	if err := c.Remove(ctx, "/ip/dns/adlist", state.ID.ValueString()); err != nil {
		resp.Diagnostics.AddError("Delete /ip/dns/adlist failed", err.Error())
	}
}

func (r *IPDNSAdlistResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import formats accepted:
	//   *<id>            -> bare RouterOS .id on the default router
	//   <router>::*<id>  -> .id on the named router
	//   <router>::<url>  -> resolved by matching url, file or comment
	//   <url>            -> resolved on the default router
	//
	// Adlist rows have no `name`, and a URL contains '/', so the '::' form is
	// the reliable way to name a router explicitly.
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
	rows, err := iPDNSAdlistLookupByNaturalKey(ctx, c, id)
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	if len(rows) == 0 {
		resp.Diagnostics.AddError("Import not found", fmt.Sprintf("no /ip/dns/adlist matches %q", id))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(rows[0][".id"]))...)
}

// iPDNSAdlistLookupByNaturalKey resolves an import key against the columns that
// actually identify an adlist row: url, file, then comment.
func iPDNSAdlistLookupByNaturalKey(ctx context.Context, c *client.Client, id string) ([]client.Object, error) {
	rows, err := c.List(ctx, "/ip/dns/adlist")
	if err != nil {
		return nil, err
	}
	for _, k := range []string{"url", "file", "comment"} {
		var hits []client.Object
		for _, r := range rows {
			if v, ok := r[k]; ok && v == id {
				hits = append(hits, r)
			}
		}
		if len(hits) > 0 {
			return hits, nil
		}
	}
	return nil, nil
}

func iPDNSAdlistApply(ctx context.Context, obj client.Object, m *IPDNSAdlistModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["url"]; ok && v != "" {
		m.URL = types.StringValue(v)
	} else {
		m.URL = types.StringNull()
	}
	if v, ok := obj["file"]; ok && v != "" {
		m.File = types.StringValue(v)
	} else {
		m.File = types.StringNull()
	}
	if v, ok := obj["ssl-verify"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.SSLVerify = types.BoolValue(b)
		} else {
			m.SSLVerify = types.BoolNull()
		}
	} else {
		m.SSLVerify = types.BoolNull()
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else {
			m.Disabled = types.BoolNull()
		}
	} else {
		m.Disabled = types.BoolNull()
	}
	if v, ok := obj["comment"]; ok && v != "" {
		m.Comment = types.StringValue(v)
	} else {
		m.Comment = types.StringNull()
	}
	if v, ok := obj["match-count"]; ok {
		if n, err := client.ParseInt64(v); err == nil {
			m.MatchCount = types.Int64Value(n)
		} else {
			m.MatchCount = types.Int64Null()
		}
	} else {
		m.MatchCount = types.Int64Null()
	}
	if v, ok := obj["name-count"]; ok {
		if n, err := client.ParseInt64(v); err == nil {
			m.NameCount = types.Int64Value(n)
		} else {
			m.NameCount = types.Int64Null()
		}
	} else {
		m.NameCount = types.Int64Null()
	}
}
