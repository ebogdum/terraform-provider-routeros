package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// hostAddrType is a string attribute type whose values compare equal when they
// denote the same host address (IPv4 or IPv6). RouterOS stores a single-host address-list entry
// with an explicit "/128" suffix, so a config value of "2001:db8::1" would
// otherwise mismatch the value read back ("2001:db8::1/128") and trip
// "Provider produced inconsistent result after apply". Semantic equality lets a
// user write either form without a permanent diff. Ranges and prefixes compare
// literally (only a redundant "/128" on a single host is normalised away).
type hostAddrType struct {
	basetypes.StringType
}

var (
	_ basetypes.StringTypable                    = hostAddrType{}
	_ basetypes.StringValuableWithSemanticEquals = hostAddrValue{}
)

func (t hostAddrType) Equal(o attr.Type) bool {
	other, ok := o.(hostAddrType)
	if !ok {
		return false
	}
	return t.StringType.Equal(other.StringType)
}

func (t hostAddrType) String() string { return "hostAddrType" }

func (t hostAddrType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return hostAddrValue{StringValue: in}, nil
}

func (t hostAddrType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	sv, ok := attrValue.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T from StringType", attrValue)
	}
	return hostAddrValue{StringValue: sv}, nil
}

func (t hostAddrType) ValueType(_ context.Context) attr.Value {
	return hostAddrValue{}
}

// hostAddrValue is the value produced by hostAddrType.
type hostAddrValue struct {
	basetypes.StringValue
}

func (v hostAddrValue) Type(_ context.Context) attr.Type { return hostAddrType{} }

func (v hostAddrValue) Equal(o attr.Value) bool {
	other, ok := o.(hostAddrValue)
	if !ok {
		return false
	}
	return v.StringValue.Equal(other.StringValue)
}

func (v hostAddrValue) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	newV, ok := newValuable.(hostAddrValue)
	if !ok {
		return false, diags
	}
	if v.IsNull() || v.IsUnknown() || newV.IsNull() || newV.IsUnknown() {
		return false, diags
	}
	return canonHostAddr(v.ValueString()) == canonHostAddr(newV.ValueString()), diags
}

// canonHostAddr lower-cases and drops a redundant "/128" host-prefix so the two
// forms RouterOS accepts for a single host collapse to one.
func canonHostAddr(s string) string {
	s = strings.ToLower(strings.TrimSpace(s))
	s = strings.TrimSuffix(s, "/128")
	s = strings.TrimSuffix(s, "/32")
	return s
}

// newHostAddrValue wraps a plain string in an hostAddrValue for use in Apply.
func newHostAddrValue(s string) hostAddrValue {
	return hostAddrValue{StringValue: basetypes.NewStringValue(s)}
}

func newHostAddrNull() hostAddrValue {
	return hostAddrValue{StringValue: basetypes.NewStringNull()}
}
