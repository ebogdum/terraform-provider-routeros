package provider

import (
	"context"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// canonRate expands every "/"-separated part of a rate value, returning the
// canonical string and whether the whole value parsed.
func canonRate(s string) (string, bool) {
	s = strings.TrimSpace(s)
	if s == "" {
		return s, false
	}
	parts := strings.Split(s, "/")
	out := make([]string, len(parts))
	for i, p := range parts {
		n, ok := expandRateSuffix(strings.TrimSpace(p))
		if !ok {
			return s, false
		}
		out[i] = n
	}
	// A single rate is symmetric (rx=tx); RouterOS stores it as the pair "X/X".
	// Normalise to a pair so "10M" and "10M/10M" compare equal.
	if len(out) == 1 {
		out = append(out, out[0])
	}
	return strings.Join(out, "/"), true
}

func expandRateSuffix(p string) (string, bool) {
	if p == "" {
		return "", false
	}
	mult := int64(1)
	num := p
	switch p[len(p)-1] {
	case 'k', 'K':
		mult, num = 1000, p[:len(p)-1]
	case 'M':
		mult, num = 1000000, p[:len(p)-1]
	case 'G':
		mult, num = 1000000000, p[:len(p)-1]
	}
	if iv, err := strconv.ParseInt(num, 10, 64); err == nil {
		return strconv.FormatInt(iv*mult, 10), true
	}
	if mult != 1 {
		if fv, err := strconv.ParseFloat(num, 64); err == nil {
			return strconv.FormatInt(int64(fv*float64(mult)), 10), true
		}
	}
	return "", false
}

// diagBuf is the diagnostics type used by generated Apply/Upsert helpers.
type diagBuf = diag.Diagnostics

// nullifyUnknownAttrs resolves any attribute still marked Unknown on the model
// to its typed Null value. It is called before writing plan-derived state back
// to Terraform (Create/Update). An Apply helper only assigns an attribute when
// the device response actually carries that key; a Computed attribute the
// device never returns (e.g. /ip/cloud dns-name on a board with DDNS off) is
// therefore left at its incoming value. On Create that incoming value is the
// plan's Unknown, and Terraform rejects a result that still has Unknowns
// ("Provider returned invalid result object after apply"). Resolving only
// Unknown -> Null is safe in every path: user-set Optional values and prior
// state are already known, so they are untouched.
func nullifyUnknownAttrs(m any) {
	rv := reflect.ValueOf(m)
	if rv.Kind() != reflect.Pointer || rv.IsNil() {
		return
	}
	rv = rv.Elem()
	if rv.Kind() != reflect.Struct {
		return
	}
	for i := 0; i < rv.NumField(); i++ {
		f := rv.Field(i)
		if !f.CanSet() {
			continue
		}
		av, ok := f.Interface().(attr.Value)
		if !ok || !av.IsUnknown() {
			continue
		}
		switch av.(type) {
		case types.String:
			f.Set(reflect.ValueOf(types.StringNull()))
		case types.Bool:
			f.Set(reflect.ValueOf(types.BoolNull()))
		case types.Int64:
			f.Set(reflect.ValueOf(types.Int64Null()))
		case types.Float64:
			f.Set(reflect.ValueOf(types.Float64Null()))
		case types.Number:
			f.Set(reflect.ValueOf(types.NumberNull()))
		}
	}
}

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
