package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

func TestHostAddrSemanticEquals(t *testing.T) {
	ctx := context.Background()
	eq := func(a, b string) bool {
		ok, d := newHostAddrValue(a).StringSemanticEquals(ctx, newHostAddrValue(b))
		if d.HasError() {
			t.Fatalf("diags: %v", d)
		}
		return ok
	}
	// A bare host and its /128 form denote the same address.
	if !eq("192.0.2.5", "192.0.2.5/32") {
		t.Error("192.0.2.5 and /32 should be equal")
	}
	if !eq("2001:db8::1", "2001:db8::1/128") {
		t.Error("2001:db8::1 and 2001:db8::1/128 should be semantically equal")
	}
	if !eq("2001:DB8::1", "2001:db8::1/128") {
		t.Error("case should not matter")
	}
	// Distinct hosts, ranges and prefixes stay distinct.
	if eq("2001:db8::1", "2001:db8::2") {
		t.Error("different hosts must not be equal")
	}
	if eq("2001:db8::/64", "2001:db8::/128") {
		t.Error("prefix vs host must not be equal")
	}
	// Null/unknown never claim equality.
	nullV := hostAddrValue{StringValue: basetypes.NewStringNull()}
	if ok, _ := nullV.StringSemanticEquals(ctx, newHostAddrValue("2001:db8::1")); ok {
		t.Error("null must not be equal to a value")
	}
}
