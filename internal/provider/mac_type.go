package provider

import (
	"context"
	"fmt"

	"github.com/hashicorp/terraform-plugin-framework/attr"
	"github.com/hashicorp/terraform-plugin-framework/diag"
	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
	"github.com/hashicorp/terraform-plugin-go/tftypes"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// macType is a string attribute type for RouterOS MAC addresses. The device
// stores every MAC as upper-case colon form whatever the config wrote, so
// equality is semantic rather than literal.
type macType struct {
	basetypes.StringType
}

var (
	_ basetypes.StringTypable                    = macType{}
	_ basetypes.StringValuableWithSemanticEquals = macValue{}
)

func (t macType) Equal(o attr.Type) bool {
	other, ok := o.(macType)
	return ok && t.StringType.Equal(other.StringType)
}

func (t macType) String() string { return "macType" }

func (t macType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return macValue{StringValue: in}, nil
}

func (t macType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	av, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	sv, ok := av.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T from StringType", av)
	}
	return macValue{StringValue: sv}, nil
}

func (t macType) ValueType(_ context.Context) attr.Value { return macValue{} }

type macValue struct {
	basetypes.StringValue
}

func (v macValue) Type(_ context.Context) attr.Type { return macType{} }

func (v macValue) Equal(o attr.Value) bool {
	other, ok := o.(macValue)
	return ok && v.StringValue.Equal(other.StringValue)
}

func (v macValue) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	newV, ok := newValuable.(macValue)
	if !ok || v.IsNull() || v.IsUnknown() || newV.IsNull() || newV.IsUnknown() {
		return false, diags
	}
	a, err := client.CanonicalMAC(v.ValueString())
	if err != nil {
		return false, diags
	}
	b, err := client.CanonicalMAC(newV.ValueString())
	if err != nil {
		return false, diags
	}
	return a == b, diags
}

func newMACValue(s string) macValue {
	return macValue{StringValue: basetypes.NewStringValue(s)}
}

func newMACNull() macValue {
	return macValue{StringValue: basetypes.NewStringNull()}
}
