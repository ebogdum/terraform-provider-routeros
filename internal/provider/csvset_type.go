package provider

import (
	"context"
	"fmt"
	"sort"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// csvSetType is a string attribute type for RouterOS comma-separated values that
// are semantically an unordered set (auth methods, firewall time days, ...).
// RouterOS returns the members in its own order, so "mschap2,mschap1" and
// "mschap1,mschap2" denote the same value. Semantic equality compares the
// comma-separated members as a set, ending the perpetual re-order diff.
type csvSetType struct {
	basetypes.StringType
}

var (
	_ basetypes.StringTypable                    = csvSetType{}
	_ basetypes.StringValuableWithSemanticEquals = csvSetValue{}
)

func (t csvSetType) Equal(o attr.Type) bool {
	other, ok := o.(csvSetType)
	return ok && t.StringType.Equal(other.StringType)
}

func (t csvSetType) String() string { return "csvSetType" }

func (t csvSetType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return csvSetValue{StringValue: in}, nil
}

func (t csvSetType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	av, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	sv, ok := av.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T from StringType", av)
	}
	return csvSetValue{StringValue: sv}, nil
}

func (t csvSetType) ValueType(_ context.Context) attr.Value { return csvSetValue{} }

type csvSetValue struct {
	basetypes.StringValue
}

func (v csvSetValue) Type(_ context.Context) attr.Type { return csvSetType{} }

func (v csvSetValue) Equal(o attr.Value) bool {
	other, ok := o.(csvSetValue)
	return ok && v.StringValue.Equal(other.StringValue)
}

func (v csvSetValue) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	newV, ok := newValuable.(csvSetValue)
	if !ok || v.IsNull() || v.IsUnknown() || newV.IsNull() || newV.IsUnknown() {
		return false, diags
	}
	return canonCSVSet(v.ValueString()) == canonCSVSet(newV.ValueString()), diags
}

// canonCSVSet lower-cases, trims, drops empties and sorts the members so any
// ordering of the same comma-separated set collapses to one canonical string.
func canonCSVSet(s string) string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	sort.Strings(out)
	return strings.Join(out, ",")
}

func newCSVSetValue(s string) csvSetValue {
	return csvSetValue{StringValue: basetypes.NewStringValue(s)}
}

func newCSVSetNull() csvSetValue {
	return csvSetValue{StringValue: basetypes.NewStringNull()}
}
