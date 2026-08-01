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

// boolStringType is a string attribute type for RouterOS fields that hold a
// boolean but are exposed as a string (because the menu also accepts a keyword
// like "auto"). RouterOS accepts "yes"/"no" on input but reads the value back
// as "true"/"false", so a config of "no" would otherwise never match the
// "false" read back. Semantic equality compares the two as booleans when both
// parse as one, and falls back to a literal comparison otherwise (so keyword
// values like "auto" still compare exactly).
type boolStringType struct {
	basetypes.StringType
}

var (
	_ basetypes.StringTypable                    = boolStringType{}
	_ basetypes.StringValuableWithSemanticEquals = boolStringValue{}
)

func (t boolStringType) Equal(o attr.Type) bool {
	other, ok := o.(boolStringType)
	return ok && t.StringType.Equal(other.StringType)
}

func (t boolStringType) String() string { return "boolStringType" }

func (t boolStringType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return boolStringValue{StringValue: in}, nil
}

func (t boolStringType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	av, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	sv, ok := av.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T from StringType", av)
	}
	return boolStringValue{StringValue: sv}, nil
}

func (t boolStringType) ValueType(_ context.Context) attr.Value { return boolStringValue{} }

type boolStringValue struct {
	basetypes.StringValue
}

func (v boolStringValue) Type(_ context.Context) attr.Type { return boolStringType{} }

func (v boolStringValue) Equal(o attr.Value) bool {
	other, ok := o.(boolStringValue)
	return ok && v.StringValue.Equal(other.StringValue)
}

func (v boolStringValue) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	newV, ok := newValuable.(boolStringValue)
	if !ok || v.IsNull() || v.IsUnknown() || newV.IsNull() || newV.IsUnknown() {
		return false, diags
	}
	a, aok := parseBoolish(v.ValueString())
	b, bok := parseBoolish(newV.ValueString())
	if aok && bok {
		return a == b, diags
	}
	return v.ValueString() == newV.ValueString(), diags
}

func parseBoolish(s string) (bool, bool) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "yes", "true", "1":
		return true, true
	case "no", "false", "0":
		return false, true
	}
	return false, false
}

func newBoolStringValue(s string) boolStringValue {
	return boolStringValue{StringValue: basetypes.NewStringValue(s)}
}

func newBoolStringNull() boolStringValue {
	return boolStringValue{StringValue: basetypes.NewStringNull()}
}
