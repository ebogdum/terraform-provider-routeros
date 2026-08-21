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

// cidrType is a string attribute type for RouterOS addresses. The device
// canonicalises spacing and form, so equality is by parsed address.
type cidrType struct {
	basetypes.StringType
}

var (
	_ basetypes.StringTypable                    = cidrType{}
	_ basetypes.StringValuableWithSemanticEquals = cidrValue{}
)

func (t cidrType) Equal(o attr.Type) bool {
	other, ok := o.(cidrType)
	return ok && t.StringType.Equal(other.StringType)
}

func (t cidrType) String() string { return "cidrType" }

func (t cidrType) ValueFromString(_ context.Context, in basetypes.StringValue) (basetypes.StringValuable, diag.Diagnostics) {
	return cidrValue{StringValue: in}, nil
}

func (t cidrType) ValueFromTerraform(ctx context.Context, in tftypes.Value) (attr.Value, error) {
	av, err := t.StringType.ValueFromTerraform(ctx, in)
	if err != nil {
		return nil, err
	}
	sv, ok := av.(basetypes.StringValue)
	if !ok {
		return nil, fmt.Errorf("unexpected value type %T from StringType", av)
	}
	return cidrValue{StringValue: sv}, nil
}

func (t cidrType) ValueType(_ context.Context) attr.Value { return cidrValue{} }

type cidrValue struct {
	basetypes.StringValue
}

func (v cidrValue) Type(_ context.Context) attr.Type { return cidrType{} }

func (v cidrValue) Equal(o attr.Value) bool {
	other, ok := o.(cidrValue)
	return ok && v.StringValue.Equal(other.StringValue)
}

// A value that does not parse is compared literally.
func (v cidrValue) StringSemanticEquals(_ context.Context, newValuable basetypes.StringValuable) (bool, diag.Diagnostics) {
	var diags diag.Diagnostics
	newV, ok := newValuable.(cidrValue)
	if !ok || v.IsNull() || v.IsUnknown() || newV.IsNull() || newV.IsUnknown() {
		return false, diags
	}
	a, aerr := client.CanonicalCIDR(v.ValueString())
	b, berr := client.CanonicalCIDR(newV.ValueString())
	if aerr != nil || berr != nil {
		return strings.EqualFold(strings.TrimSpace(v.ValueString()), strings.TrimSpace(newV.ValueString())), diags
	}
	return a == b, diags
}

func newCIDRValue(s string) cidrValue {
	return cidrValue{StringValue: basetypes.NewStringValue(s)}
}

func newCIDRNull() cidrValue {
	return cidrValue{StringValue: basetypes.NewStringNull()}
}
