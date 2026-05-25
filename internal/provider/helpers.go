package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// diagBuf is the diagnostics type used by generated Apply/Upsert helpers.
type diagBuf = diag.Diagnostics

// configureRegistry is called from every generated resource's and data source's
// Configure(). It pulls a *client.Registry out of ProviderData and stores it.
// Returns the registry plus a diagnostic if the provider isn't configured yet.
func configureRegistry(providerData any) (*client.Registry, diag.Diagnostics) {
	var d diag.Diagnostics
	if providerData == nil {
		return nil, d
	}
	reg, ok := providerData.(*client.Registry)
	if !ok {
		d.AddError("Unexpected provider data", fmt.Sprintf("got %T, want *client.Registry", providerData))
		return nil, d
	}
	return reg, d
}

// pickClient resolves a `router` attribute value to a *Client, falling back to
// the registry's default.
func pickClient(reg *client.Registry, routerAttr types.String, diags *diag.Diagnostics) *client.Client {
	if reg == nil {
		diags.AddError("Provider not configured", "the provider was not fully initialised before this resource was used")
		return nil
	}
	name := ""
	if !routerAttr.IsNull() && !routerAttr.IsUnknown() {
		name = routerAttr.ValueString()
	}
	c, err := reg.Get(name)
	if err != nil {
		diags.AddError("Unknown router", err.Error())
		return nil
	}
	return c
}

// naturalKeyCandidates lists the columns that, in priority order, are tested
// as the natural-key match during import. Many RouterOS menus pick one of
// these (name for /interface, address for /ip/address, dst-address for
// /ip/route, ...). Scanning a small, well-known list is safer than scanning
// every value of every row.
var naturalKeyCandidates = []string{
	"name",
	"address",
	"dst-address",
	"src-address",
	"gateway",
	"interface",
	"host",
	"comment",
	"target",
	"ipv6-address",
	"dns-name",
}

// lookupByNaturalKey scans rows at menuPath looking for one whose well-known
// natural-key column equals id. Used by every generated ImportState helper
// when the import target is not a bare *<id>.
func lookupByNaturalKey(ctx context.Context, c *client.Client, menuPath, id string) ([]client.Object, error) {
	rows, err := c.List(ctx, menuPath)
	if err != nil {
		return nil, err
	}
	for _, k := range naturalKeyCandidates {
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

// parseImportID parses a Terraform import ID. The accepted shapes are:
//
//	*<id>            -> bare RouterOS .id on the default router
//	<router>::<rest> -> explicit form (preferred when <rest> contains '/')
//	<router>/<rest>  -> legacy form, honoured only if <router> is a configured
//	                    name; otherwise the whole string is treated as <rest>
//	                    on the default router (so CIDR-shaped natural keys
//	                    like 10.0.0.1/24 are not misparsed as router=10.0.0.1)
//	<rest>           -> natural key on the default router
func parseImportID(reg *client.Registry, raw string) (router, key string) {
	if i := strings.Index(raw, "::"); i > 0 {
		return raw[:i], raw[i+2:]
	}
	if strings.HasPrefix(raw, "*") {
		return "", raw
	}
	if i := strings.Index(raw, "/"); i > 0 {
		candidate := raw[:i]
		if _, err := reg.Get(candidate); err == nil {
			return candidate, raw[i+1:]
		}
	}
	return "", raw
}
