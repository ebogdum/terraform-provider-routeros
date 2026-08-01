package provider

import (
	"context"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// encodeStringSet renders a set of strings as a RouterOS comma list.
func encodeStringSet(ctx context.Context, s types.Set, diags *diag.Diagnostics) string {
	if s.IsNull() || s.IsUnknown() {
		return ""
	}
	var items []string
	if d := s.ElementsAs(ctx, &items, false); d.HasError() {
		diags.Append(d...)
		return ""
	}
	return client.FormatList(items)
}

// decodeStringSet parses a RouterOS comma list into an order-insensitive Set.
// RouterOS returns many list-valued attributes (dh-group, enc-algorithms,
// login-by, ...) in its own order, so modelling them as a Set makes them
// round-trip regardless of the order the device echoes back.
func decodeStringSet(_ context.Context, wire string) types.Set {
	items := client.ParseList(wire)
	vals := make([]attr.Value, 0, len(items))
	for _, v := range items {
		vals = append(vals, types.StringValue(v))
	}
	s, _ := types.SetValue(types.StringType, vals)
	return s
}

// decodePolicySet parses a /user/group or /system/script policy value.
// RouterOS returns the full permission list with everything not granted negated
// ("read,write,!ftp,!telnet,..."), in its own order. Only the granted
// permissions are kept, as an order-insensitive Set, so the value round-trips
// regardless of order or the negated remainder.
func decodePolicySet(_ context.Context, wire string) types.Set {
	items := client.ParseList(wire)
	vals := make([]attr.Value, 0, len(items))
	for _, v := range items {
		if strings.HasPrefix(strings.TrimSpace(v), "!") {
			continue
		}
		vals = append(vals, types.StringValue(v))
	}
	s, _ := types.SetValue(types.StringType, vals)
	return s
}

func encodeStringList(ctx context.Context, l types.List, diags *diag.Diagnostics) string {
	if l.IsNull() || l.IsUnknown() {
		return ""
	}
	var items []string
	if d := l.ElementsAs(ctx, &items, false); d.HasError() {
		// Surface the conversion failure rather than silently writing an empty
		// list, which would clear the field on RouterOS.
		diags.Append(d...)
		return ""
	}
	return client.FormatList(items)
}

func decodeStringList(_ context.Context, wire string) types.List {
	items := client.ParseList(wire)
	vals := make([]attr.Value, 0, len(items))
	for _, v := range items {
		vals = append(vals, types.StringValue(v))
	}
	l, _ := types.ListValue(types.StringType, vals)
	return l
}

func dsRowsToList(_ context.Context, rows []client.Object) (types.List, diag.Diagnostics) {
	mapType := types.MapType{ElemType: types.StringType}
	out := make([]attr.Value, 0, len(rows))
	var diags diag.Diagnostics
	for _, r := range rows {
		m := map[string]attr.Value{}
		for k, v := range r {
			m[k] = types.StringValue(v)
		}
		mv, d := types.MapValue(types.StringType, m)
		diags.Append(d...)
		out = append(out, mv)
	}
	l, d := types.ListValue(mapType, out)
	diags.Append(d...)
	return l, diags
}

func actionRowsToList(ctx context.Context, rows []client.Object) types.List {
	l, _ := dsRowsToList(ctx, rows)
	return l
}

// stateIDFor builds a stable id for singletons / routerless resources.
func stateIDFor(menuPath string, router types.String) string {
	if router.IsNull() || router.IsUnknown() || router.ValueString() == "" {
		return menuPath
	}
	return router.ValueString() + ":" + menuPath
}
