package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/types/basetypes"
)

// RouterOS stores every MAC as upper-case colon form. Terraform will not let the
// plan rewrite a config value, so the two are reconciled here instead.
func TestMACValueSemanticEquals(t *testing.T) {
	for _, tc := range []struct {
		name, config, device string
		want                 bool
	}{
		{"case differs", "aa:bb:cc:dd:ee:ff", "AA:BB:CC:DD:EE:FF", true},
		{"hyphen separators", "aa-bb-cc-dd-ee-01", "AA:BB:CC:DD:EE:01", true},
		{"bare hex", "aabbccddee02", "AA:BB:CC:DD:EE:02", true},
		{"identical", "AA:BB:CC:DD:EE:FF", "AA:BB:CC:DD:EE:FF", true},
		{"genuinely different", "aa:bb:cc:dd:ee:ff", "11:22:33:44:55:66", false},
		{"config not a mac", "not-a-mac", "AA:BB:CC:DD:EE:FF", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, diags := newMACValue(tc.config).StringSemanticEquals(context.Background(), newMACValue(tc.device))
			if diags.HasError() {
				t.Fatalf("diags: %v", diags)
			}
			if got != tc.want {
				t.Errorf("semantic equality = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMACValueNullAndUnknownAreNeverSemanticallyEqual(t *testing.T) {
	known := newMACValue("AA:BB:CC:DD:EE:FF")
	for name, v := range map[string]macValue{
		"null":    newMACNull(),
		"unknown": {StringValue: basetypes.NewStringUnknown()},
	} {
		if eq, _ := v.StringSemanticEquals(context.Background(), known); eq {
			t.Errorf("%s compared equal to a known value", name)
		}
		if eq, _ := known.StringSemanticEquals(context.Background(), v); eq {
			t.Errorf("known value compared equal to %s", name)
		}
	}
}
