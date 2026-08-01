package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"
)

// rosRateType is a string attribute type for RouterOS rate values (queue
// limits, burst rates). RouterOS accepts human suffixes (k/M/G, decimal 1000)
// but stores and reads back the expanded integer: "10M/10M" becomes
// "10000000/10000000". Semantic equality compares the expanded forms so a user
// may write either notation without a permanent diff or an "inconsistent result
// after apply" error. Values that do not parse as rates compare literally.
type rosRateType struct {
	basetypes.StringType
}

var (
	_ basetypes.StringTypable                    = rosRateType{}
	_ basetypes.StringValuableWithSemanticEquals = rosRateValue{}
)

func (t rosRateType) Equal(o attr.Type) bool {
	other, ok := o.(rosRateType)
	if !ok {
		return false
	}
	return t.StringType.Equal(other.StringType)
}

func (t rosRateType) String() string { return "rosRateType" }

func (t rosRateType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return rosRateValue{StringValue: in}, nil
}

func (t rosRateType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	attrValue, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	sv, ok := attrValue.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T from StringType", attrValue)
	}
	return rosRateValue{StringValue: sv}, nil
}

func (t rosRateType) ValueType(_ context.Context) attr.Value { return rosRateValue{} }

// rosRateValue is the value produced by rosRateType.
type rosRateValue struct {
	basetypes.StringValue
}

func (v rosRateValue) Type(_ context.Context) attr.Type { return rosRateType{} }

func (v rosRateValue) Equal(o attr.Value) bool {
	other, ok := o.(rosRateValue)
	if !ok {
		return false
	}
	return v.StringValue.Equal(other.StringValue)
}

func (v rosRateValue) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	newV, ok := newValuable.(rosRateValue)
	if !ok {
		return false, diags
	}
	if v.IsNull() || v.IsUnknown() || newV.IsNull() || newV.IsUnknown() {
		return false, diags
	}
	a, aok := canonRate(v.ValueString())
	b, bok := canonRate(newV.ValueString())
	if !aok || !bok {
		return v.ValueString() == newV.ValueString(), diags
	}
	return a == b, diags
}

func newRosRateValue(s string) rosRateValue {
	return rosRateValue{StringValue: basetypes.NewStringValue(s)}
}

func newRosRateNull() rosRateValue {
	return rosRateValue{StringValue: basetypes.NewStringNull()}
}
