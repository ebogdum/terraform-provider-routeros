package provider

import (
	"context"
	"os"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/datasource"
	"github.com/hashicorp/terraform-plugin-framework/provider"
	"github.com/hashicorp/terraform-plugin-framework/provider/schema"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

type rosProvider struct {
	version string
}

// rosProviderModel mirrors the provider block. Two coexisting forms:
//
//  1. Single-router (back-compat):
//     provider "routeros" { host = "..." username = "..." password = "..." }
//     Becomes one registered router named "default".
//
//  2. Multi-router:
//     provider "routeros" {
//     routers = {
//     core    = { host = "...", username = "...", password = "..." }
//     edge_se = { host = "...", username = "...", password = "..." }
//     }
//     }
//     Every resource may set `router = "core"` to pick which one to act on;
//     omitting the attribute uses the router named "default", or the first
//     router in alphabetical order if "default" isn't defined.
//
// The two forms can be combined: a `routers` map registers many; the loose
// host/username/password registers (or overrides) the "default" entry.
type rosProviderModel struct {
	Host       types.String         `tfsdk:"host"`
	Username   types.String         `tfsdk:"username"`
	Password   types.String         `tfsdk:"password"`
	CACert     types.String         `tfsdk:"ca_cert"`
	Insecure   types.Bool           `tfsdk:"insecure"`
	ROSVersion types.String         `tfsdk:"ros_version"`
	Routers    map[string]routerCfg `tfsdk:"routers"`
}

type routerCfg struct {
	Host       types.String `tfsdk:"host"`
	Username   types.String `tfsdk:"username"`
	Password   types.String `tfsdk:"password"`
	CACert     types.String `tfsdk:"ca_cert"`
	Insecure   types.Bool   `tfsdk:"insecure"`
	ROSVersion types.String `tfsdk:"ros_version"`
}

// New returns a provider factory for the given build version.
func New(version string) func() provider.Provider {
	return func() provider.Provider {
		return &rosProvider{version: version}
	}
}

func (p *rosProvider) Metadata(_ context.Context, _ provider.MetadataRequest, resp *provider.MetadataResponse) {
	resp.TypeName = "routeros"
	resp.Version = p.version
}

func (p *rosProvider) Schema(_ context.Context, _ provider.SchemaRequest, resp *provider.SchemaResponse) {
	routerAttrs := map[string]schema.Attribute{
		"host":        schema.StringAttribute{Required: true, Description: "RouterOS base URL (https://ip or http://ip)."},
		"username":    schema.StringAttribute{Required: true, Description: "API user."},
		"password":    schema.StringAttribute{Required: true, Sensitive: true, Description: "API password."},
		"ca_cert":     schema.StringAttribute{Optional: true, Description: "PEM-encoded CA bundle."},
		"insecure":    schema.BoolAttribute{Optional: true, Description: "Skip TLS verification."},
		"ros_version": schema.StringAttribute{Optional: true, Description: "RouterOS version override."},
	}
	resp.Schema = schema.Schema{
		Description: "Complete Terraform provider for MikroTik RouterOS 7.x via the REST API. Supports managing multiple routers from a single provider block via the `routers` map.",
		Attributes: map[string]schema.Attribute{
			"host":        schema.StringAttribute{Optional: true, Description: "Single-router host. Falls back to env ROUTEROS_HOST."},
			"username":    schema.StringAttribute{Optional: true, Description: "Single-router user. Falls back to env ROUTEROS_USER."},
			"password":    schema.StringAttribute{Optional: true, Sensitive: true, Description: "Single-router password. Falls back to env ROUTEROS_PASSWORD."},
			"ca_cert":     schema.StringAttribute{Optional: true, Description: "Single-router CA bundle. Falls back to env ROUTEROS_CA_CERT."},
			"insecure":    schema.BoolAttribute{Optional: true, Description: "Single-router TLS skip. Falls back to env ROUTEROS_INSECURE."},
			"ros_version": schema.StringAttribute{Optional: true, Description: "Single-router version override. Falls back to env ROUTEROS_VERSION."},
			"routers": schema.MapNestedAttribute{
				Optional:     true,
				Description:  "Named routers. Resources set `router = \"<name>\"` to pick one; omitting uses `default` or the first name in sorted order.",
				NestedObject: schema.NestedAttributeObject{Attributes: routerAttrs},
			},
		},
	}
}

func (p *rosProvider) Configure(ctx context.Context, req provider.ConfigureRequest, resp *provider.ConfigureResponse) {
	var data rosProviderModel
	diags := req.Config.Get(ctx, &data)
	resp.Diagnostics.Append(diags...)
	if resp.Diagnostics.HasError() {
		return
	}

	// Defer configuration if any required value is unknown (e.g. supplied by a
	// variable that hasn't been resolved at plan time). terraform-plugin-
	// framework will call Configure again at apply time with concrete values.
	if data.Host.IsUnknown() || data.Username.IsUnknown() || data.Password.IsUnknown() {
		return
	}
	for _, r := range data.Routers {
		if r.Host.IsUnknown() || r.Username.IsUnknown() || r.Password.IsUnknown() {
			return
		}
	}

	configs := map[string]client.Config{}
	for name, r := range data.Routers {
		if r.Host.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(pathRouters.AtMapKey(name).AtName("host"),
				"router host required",
				"every entry under `routers` must set host")
			continue
		}
		if r.Username.ValueString() == "" {
			resp.Diagnostics.AddAttributeError(pathRouters.AtMapKey(name).AtName("username"),
				"router username required",
				"every entry under `routers` must set username")
			continue
		}
		configs[name] = client.Config{
			Host:       r.Host.ValueString(),
			Username:   r.Username.ValueString(),
			Password:   r.Password.ValueString(),
			CACertPEM:  r.CACert.ValueString(),
			Insecure:   r.Insecure.ValueBool(),
			ROSVersion: r.ROSVersion.ValueString(),
		}
	}
	if resp.Diagnostics.HasError() {
		return
	}

	// Single-router back-compat: derive a "default" entry from the loose
	// attributes and the env. If `routers` also defines "default", the loose
	// values override only the empty fields of the explicit definition.
	loose := client.Config{
		Host:       firstNonEmpty(data.Host.ValueString(), os.Getenv("ROUTEROS_HOST")),
		Username:   firstNonEmpty(data.Username.ValueString(), os.Getenv("ROUTEROS_USER")),
		Password:   firstNonEmpty(data.Password.ValueString(), os.Getenv("ROUTEROS_PASSWORD")),
		CACertPEM:  firstNonEmpty(data.CACert.ValueString(), os.Getenv("ROUTEROS_CA_CERT")),
		Insecure:   data.Insecure.ValueBool() || strings.EqualFold(os.Getenv("ROUTEROS_INSECURE"), "true"),
		ROSVersion: firstNonEmpty(data.ROSVersion.ValueString(), os.Getenv("ROUTEROS_VERSION")),
	}
	if loose.Host != "" {
		existing, has := configs["default"]
		if !has {
			configs["default"] = loose
		} else {
			configs["default"] = mergeConfig(existing, loose)
		}
	}

	if len(configs) == 0 {
		resp.Diagnostics.AddError(
			"No RouterOS routers configured",
			"Set the `host`/`username`/`password` attributes (or the matching ROUTEROS_HOST/USER/PASSWORD env vars), or define one or more entries under `routers = { name = { host = ..., username = ..., password = ... } }`.",
		)
		return
	}

	reg, err := client.NewRegistry(configs)
	if err != nil {
		resp.Diagnostics.AddError("RouterOS registry init failed", err.Error())
		return
	}

	resp.DataSourceData = reg
	resp.ResourceData = reg
	resp.EphemeralResourceData = reg
}

func (p *rosProvider) Resources(_ context.Context) []func() resource.Resource {
	return registryResources()
}

func (p *rosProvider) DataSources(_ context.Context) []func() datasource.DataSource {
	return registryDataSources()
}

func firstNonEmpty(vals ...string) string {
	for _, v := range vals {
		if v != "" {
			return v
		}
	}
	return ""
}

// mergeConfig fills empty fields of explicit with values from loose. Explicit
// always wins when it has a non-zero value.
func mergeConfig(explicit, loose client.Config) client.Config {
	if explicit.Host == "" {
		explicit.Host = loose.Host
	}
	if explicit.Username == "" {
		explicit.Username = loose.Username
	}
	if explicit.Password == "" {
		explicit.Password = loose.Password
	}
	if explicit.CACertPEM == "" {
		explicit.CACertPEM = loose.CACertPEM
	}
	if explicit.ROSVersion == "" {
		explicit.ROSVersion = loose.ROSVersion
	}
	// Insecure: explicit-true is honored; otherwise inherit.
	if !explicit.Insecure {
		explicit.Insecure = loose.Insecure
	}
	return explicit
}
