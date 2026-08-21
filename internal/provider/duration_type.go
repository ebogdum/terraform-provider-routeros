package provider

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// durationType is a string attribute type for RouterOS time values. The device
// rewrites what it is given -- "120" comes back "2m", "00:05:00" comes back
// "5m" -- so equality is by elapsed time, not by spelling.
type durationType struct {
	basetypes.StringType
}

var (
	_ basetypes.StringTypable                    = durationType{}
	_ basetypes.StringValuableWithSemanticEquals = durationValue{}
)

func (t durationType) Equal(o attr.Type) bool {
	other, ok := o.(durationType)
	return ok && t.StringType.Equal(other.StringType)
}

func (t durationType) String() string { return "durationType" }

func (t durationType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return durationValue{StringValue: in}, nil
}

func (t durationType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	av, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	sv, ok := av.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T from StringType", av)
	}
	return durationValue{StringValue: sv}, nil
}

func (t durationType) ValueType(_ context.Context) attr.Value { return durationValue{} }

type durationValue struct {
	basetypes.StringValue
}

func (v durationValue) Type(_ context.Context) attr.Type { return durationType{} }

func (v durationValue) Equal(o attr.Value) bool {
	other, ok := o.(durationValue)
	return ok && v.StringValue.Equal(other.StringValue)
}

// A value that does not parse is a keyword such as "auto" or "disable-dpd",
// which the menus that accept one compare literally.
func (v durationValue) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	newV, ok := newValuable.(durationValue)
	if !ok || v.IsNull() || v.IsUnknown() || newV.IsNull() || newV.IsUnknown() {
		return false, diags
	}
	a, aerr := client.ParseDuration(v.ValueString())
	b, berr := client.ParseDuration(newV.ValueString())
	if aerr != nil || berr != nil {
		return strings.EqualFold(strings.TrimSpace(v.ValueString()), strings.TrimSpace(newV.ValueString())), diags
	}
	return a == b, diags
}

func newDurationValue(s string) durationValue {
	return durationValue{StringValue: basetypes.NewStringValue(s)}
}

func newDurationNull() durationValue {
	return durationValue{StringValue: basetypes.NewStringNull()}
}
