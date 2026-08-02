package provider

import (
	"context"
	"fmt"
	"maps"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema/stringplanmodifier"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
	"github.com/ebogdum/terraform-provider-routeros/internal/schemautil"
)

var (
	_ resource.Resource                = &IPServiceResource{}
	_ resource.ResourceWithImportState = &IPServiceResource{}
)

// IPServiceResource manages one row of /ip/service (api, api-ssl, ftp, ssh,
// telnet, winbox, www, www-ssl).
//
// RouterOS ships a fixed set of service rows: they can be enabled, disabled and
// reconfigured, but never added or removed. The resource therefore *adopts* the
// existing row named by `name` on create, and Delete only drops the row from
// Terraform state -- it deliberately does not re-enable a service you disabled.
//
// On `disabled`: MikroTik's /ip/service property table omits it (the CLI shows
// it only as the X flag), which suggests the enable/disable console commands are
// the only way to toggle a service. They are not -- `disabled` is an ordinary
// settable property here, so the plain PATCH used for every other field works:
//
//	/export from RouterOS 7.22.3 emits `set ftp address="" disabled=yes ...`,
//	and menu introspection reports disabled as a writable bool (`name`, by
//	contrast, comes back read-only, so it is never sent in a PATCH body).
//
// Keeping disabled in the same PATCH as the other fields makes each apply a
// single atomic request instead of a property write plus a command call.
type IPServiceResource struct {
	reg *client.Registry
}

type IPServiceModel struct {
	ID          types.String `tfsdk:"id"`
	Name        types.String `tfsdk:"name"`
	Disabled    types.Bool   `tfsdk:"disabled"`
	Port        types.Int64  `tfsdk:"port"`
	Address     types.String `tfsdk:"address"`
	Certificate types.String `tfsdk:"certificate"`
	TLSVersion  types.String `tfsdk:"tls_version"`
	VRF         types.String `tfsdk:"vrf"`
	MaxSessions types.Int64  `tfsdk:"max_sessions"`
	Router      types.String `tfsdk:"router"`
	LockoutAck  types.Bool   `tfsdk:"lockout_ack"`
}

func NewIPServiceResource() resource.Resource { return &IPServiceResource{} }

func (r *IPServiceResource) Metadata(_ context.Context, req resource.MetadataRequest, resp *resource.MetadataResponse) {
	resp.TypeName = req.ProviderTypeName + "_ip_service"
}

func (r *IPServiceResource) Configure(_ context.Context, req resource.ConfigureRequest, resp *resource.ConfigureResponse) {
	reg, diags := configureRegistry(req.ProviderData)
	resp.Diagnostics.Append(diags...)
	if reg != nil {
		r.reg = reg
	}
}

func (r *IPServiceResource) Schema(_ context.Context, _ resource.SchemaRequest, resp *resource.SchemaResponse) {
	resp.Schema = schema.Schema{
		Description: "One row of RouterOS `/ip/service` -- the management services (`api`, `api-ssl`, " +
			"`ftp`, `ssh`, `telnet`, `winbox`, `www`, `www-ssl`). Rows cannot be created or destroyed; " +
			"this resource adopts the existing service named by `name` and manages its settings. " +
			"`terraform destroy` only forgets the row, it does not re-enable a disabled service.\n\n" +
			"Safety: refuses a change that would leave every management service disabled unless " +
			"`lockout_ack = true`.",
		Attributes: map[string]schema.Attribute{
			"id": schema.StringAttribute{
				Computed:      true,
				Description:   "RouterOS internal .id.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.UseStateForUnknown()},
			},
			"name": schema.StringAttribute{
				Required: true,
				Description: "Service name as listed by `/ip/service print`: api, api-ssl, ftp, ssh, telnet, " +
					"winbox, www, www-ssl, and reverse-proxy on newer releases. Deliberately not validated " +
					"against a fixed list -- the set varies by RouterOS version; an unknown name is reported " +
					"at apply time with the names the device actually exposes.",
				PlanModifiers: []planmodifier.String{stringplanmodifier.RequiresReplace()},
			},
			"disabled": schema.BoolAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Whether the service is disabled. Set to true to turn the service off.",
			},
			"port": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "TCP port the service listens on.",
			},
			"address": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "Comma-separated list of IP/IPv6 prefixes allowed to reach the service. Empty means any source.",
			},
			"certificate": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Description: "Certificate used by the TLS-enabled services (api-ssl, www-ssl, reverse-proxy). " +
					"`none` to unset. RouterOS rejects this on the plaintext services.",
			},
			"tls_version": schema.StringAttribute{
				Optional: true,
				Computed: true,
				Description: "Minimum accepted TLS version: `any` or `only-v1.2`. TLS-enabled services only " +
					"(api-ssl, www-ssl, reverse-proxy).",
			},
			"vrf": schema.StringAttribute{
				Optional:    true,
				Computed:    true,
				Description: "VRF the service is bound to.",
			},
			"max_sessions": schema.Int64Attribute{
				Optional:    true,
				Computed:    true,
				Description: "Maximum number of concurrent sessions (RouterOS 7.13+).",
			},
			"router": schema.StringAttribute{
				Optional:    true,
				Description: "Name of the router (key in provider's `routers` map). Omit to use the default.",
			},
			"lockout_ack": schema.BoolAttribute{
				Optional:    true,
				Description: "Acknowledge that this change may leave no enabled management service, locking you out of the router.",
			},
		},
	}
}

func (r *IPServiceResource) Create(ctx context.Context, req resource.CreateRequest, resp *resource.CreateResponse) {
	var plan IPServiceModel
	if diags := req.Plan.Get(ctx, &plan); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, plan.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	rows, err := c.List(ctx, "/ip/service")
	if err != nil {
		resp.Diagnostics.AddError("Read /ip/service failed", err.Error())
		return
	}
	name := plan.Name.ValueString()
	row := ipServiceFindByName(rows, name)
	if row == nil {
		resp.Diagnostics.AddError("Unknown service "+name,
			fmt.Sprintf("/ip/service has no row named %q. RouterOS services are a fixed set and cannot be "+
				"added; this router exposes: %s.", name, strings.Join(ipServiceNames(rows), ", ")))
		return
	}

	body := client.Object{}
	if !(plan.Disabled.IsNull() || plan.Disabled.IsUnknown()) {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !(plan.Port.IsNull() || plan.Port.IsUnknown()) {
		body["port"] = client.FormatInt64(plan.Port.ValueInt64())
	}
	if !(plan.Address.IsNull() || plan.Address.IsUnknown()) {
		body["address"] = plan.Address.ValueString()
	}
	if !(plan.Certificate.IsNull() || plan.Certificate.IsUnknown()) {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !(plan.TLSVersion.IsNull() || plan.TLSVersion.IsUnknown()) {
		body["tls-version"] = plan.TLSVersion.ValueString()
	}
	if !(plan.VRF.IsNull() || plan.VRF.IsUnknown()) {
		body["vrf"] = plan.VRF.ValueString()
	}
	if !(plan.MaxSessions.IsNull() || plan.MaxSessions.IsUnknown()) {
		body["max-sessions"] = client.FormatInt64(plan.MaxSessions.ValueInt64())
	}

	id := row[".id"]
	obj := row
	if len(body) > 0 {
		if err := ipServiceCheckLockout(rows, id, body, ipServiceAck(plan.LockoutAck)); err != nil {
			resp.Diagnostics.AddError("Refusing /ip/service change", err.Error())
			return
		}
		obj, err = c.Set(ctx, "/ip/service", id, body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/service failed", err.Error())
			return
		}
	}
	ipServiceApply(ctx, obj, &plan)
	plan.ID = types.StringValue(id)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPServiceResource) Read(ctx context.Context, req resource.ReadRequest, resp *resource.ReadResponse) {
	var state IPServiceModel
	if diags := req.State.Get(ctx, &state); diags.HasError() {
		resp.Diagnostics.Append(diags...)
		return
	}
	c := pickClient(r.reg, state.Router, &resp.Diagnostics)
	if c == nil {
		return
	}
	obj, err := c.GetByID(ctx, "/ip/service", state.ID.ValueString())
	if err != nil {
		if client.IsNotFound(err) {
			resp.State.RemoveResource(ctx)
			return
		}
		resp.Diagnostics.AddError("Read /ip/service failed", err.Error())
		return
	}
	ipServiceApply(ctx, obj, &state)
	resp.Diagnostics.Append(resp.State.Set(ctx, &state)...)
}

func (r *IPServiceResource) Update(ctx context.Context, req resource.UpdateRequest, resp *resource.UpdateResponse) {
	var plan, state IPServiceModel
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
	if !plan.Disabled.Equal(state.Disabled) && !plan.Disabled.IsUnknown() {
		body["disabled"] = client.FormatBool(plan.Disabled.ValueBool())
	}
	if !plan.Port.Equal(state.Port) && !plan.Port.IsUnknown() {
		body["port"] = client.FormatInt64(plan.Port.ValueInt64())
	}
	if !plan.Address.Equal(state.Address) && !plan.Address.IsUnknown() {
		body["address"] = plan.Address.ValueString()
	}
	if !plan.Certificate.Equal(state.Certificate) && !plan.Certificate.IsUnknown() {
		body["certificate"] = plan.Certificate.ValueString()
	}
	if !plan.TLSVersion.Equal(state.TLSVersion) && !plan.TLSVersion.IsUnknown() {
		body["tls-version"] = plan.TLSVersion.ValueString()
	}
	if !plan.VRF.Equal(state.VRF) && !plan.VRF.IsUnknown() {
		body["vrf"] = plan.VRF.ValueString()
	}
	if !plan.MaxSessions.Equal(state.MaxSessions) && !plan.MaxSessions.IsUnknown() {
		body["max-sessions"] = client.FormatInt64(plan.MaxSessions.ValueInt64())
	}

	id := state.ID.ValueString()
	var obj client.Object
	var err error
	if len(body) > 0 {
		rows, listErr := c.List(ctx, "/ip/service")
		if listErr != nil {
			resp.Diagnostics.AddError("Read /ip/service failed", listErr.Error())
			return
		}
		if err := ipServiceCheckLockout(rows, id, body, ipServiceAck(plan.LockoutAck)); err != nil {
			resp.Diagnostics.AddError("Refusing /ip/service change", err.Error())
			return
		}
		obj, err = c.Set(ctx, "/ip/service", id, body)
		if err != nil {
			resp.Diagnostics.AddError("Update /ip/service failed", err.Error())
			return
		}
	} else {
		// Nothing to push, but computed attributes may still be unknown in the
		// plan; re-read so state is fully known.
		obj, err = c.GetByID(ctx, "/ip/service", id)
		if err != nil {
			resp.Diagnostics.AddError("Read /ip/service failed", err.Error())
			return
		}
	}
	ipServiceApply(ctx, obj, &plan)
	plan.ID = types.StringValue(id)
	nullifyUnknownAttrs(&plan)
	resp.Diagnostics.Append(resp.State.Set(ctx, &plan)...)
}

func (r *IPServiceResource) Delete(_ context.Context, _ resource.DeleteRequest, _ *resource.DeleteResponse) {
	// /ip/service rows are fixed; RouterOS rejects remove. Dropping the row from
	// state is the only sensible destroy. A disabled service stays disabled --
	// silently re-enabling telnet/ftp on destroy would be a security regression.
}

func (r *IPServiceResource) ImportState(ctx context.Context, req resource.ImportStateRequest, resp *resource.ImportStateResponse) {
	// Import formats accepted:
	//   *<id>            -> bare RouterOS .id on the default router
	//   <router>/*<id>   -> .id on the named router
	//   <router>/<name>  -> service name on the named router
	//   <name>           -> service name on the default router
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
	rows, err := c.List(ctx, "/ip/service")
	if err != nil {
		resp.Diagnostics.AddError("Import lookup failed", err.Error())
		return
	}
	row := ipServiceFindByName(rows, id)
	if row == nil {
		resp.Diagnostics.AddError("Import not found",
			fmt.Sprintf("no /ip/service named %q; this router exposes: %s", id, strings.Join(ipServiceNames(rows), ", ")))
		return
	}
	resp.Diagnostics.Append(resp.State.SetAttribute(ctx, path.Root("id"), types.StringValue(row[".id"]))...)
}

// ipServiceAck reports whether the operator opted out of the lockout guard.
func ipServiceAck(ack types.Bool) bool {
	return !ack.IsNull() && !ack.IsUnknown() && ack.ValueBool()
}

// ipServiceFindByName returns the row whose name matches (case-insensitively).
func ipServiceFindByName(rows []client.Object, name string) client.Object {
	for _, row := range rows {
		if strings.EqualFold(row["name"], name) {
			return row
		}
	}
	return nil
}

// ipServiceNames lists the service names present on the device, sorted, for
// error messages.
func ipServiceNames(rows []client.Object) []string {
	out := make([]string, 0, len(rows))
	for _, row := range rows {
		if n := row["name"]; n != "" {
			out = append(out, n)
		}
	}
	sort.Strings(out)
	return out
}

// ipServiceCheckLockout projects body onto the row identified by id and asks the
// guard whether the resulting state leaves any management service reachable.
func ipServiceCheckLockout(rows []client.Object, id string, body client.Object, acknowledged bool) error {
	projected := make([]client.Object, 0, len(rows))
	for _, row := range rows {
		if row[".id"] != id {
			projected = append(projected, row)
			continue
		}
		merged := client.Object{}
		maps.Copy(merged, row)
		maps.Copy(merged, body)
		projected = append(projected, merged)
	}
	return schemautil.CheckIPServiceLockout("/ip/service", projected, acknowledged)
}

func ipServiceApply(ctx context.Context, obj client.Object, m *IPServiceModel) {
	_ = ctx
	m.ID = types.StringValue(obj[".id"])
	if v, ok := obj["name"]; ok && v != "" {
		m.Name = types.StringValue(v)
	}
	if v, ok := obj["disabled"]; ok {
		if b, err := client.ParseBool(v); err == nil {
			m.Disabled = types.BoolValue(b)
		} else if strings.TrimSpace(v) == "" {
			m.Disabled = types.BoolValue(true)
		} else {
			m.Disabled = types.BoolNull()
		}
	} else {
		m.Disabled = types.BoolNull()
	}
	if v, ok := obj["port"]; ok {
		if n, err := client.ParseInt64(v); err == nil {
			m.Port = types.Int64Value(n)
		} else {
			m.Port = types.Int64Null()
		}
	} else {
		m.Port = types.Int64Null()
	}
	if v, ok := obj["address"]; ok && v != "" {
		m.Address = types.StringValue(v)
	} else {
		m.Address = types.StringNull()
	}
	if v, ok := obj["certificate"]; ok && v != "" {
		m.Certificate = types.StringValue(v)
	} else {
		m.Certificate = types.StringNull()
	}
	if v, ok := obj["tls-version"]; ok && v != "" {
		m.TLSVersion = types.StringValue(v)
	} else {
		m.TLSVersion = types.StringNull()
	}
	if v, ok := obj["vrf"]; ok && v != "" {
		m.VRF = types.StringValue(v)
	} else {
		m.VRF = types.StringNull()
	}
	if v, ok := obj["max-sessions"]; ok {
		if n, err := client.ParseInt64(v); err == nil {
			m.MaxSessions = types.Int64Value(n)
		} else {
			m.MaxSessions = types.Int64Null()
		}
	} else {
		m.MaxSessions = types.Int64Null()
	}
}
