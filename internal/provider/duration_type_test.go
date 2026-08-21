package provider

import (
	"context"
	"testing"
)

// RouterOS rewrites a duration to its own spelling -- "120" comes back "2m" --
// and Terraform will not let the plan rewrite a config value, so the two are
// reconciled by elapsed time. Every pairing here was measured on 7.23.2.
func TestDurationValueSemanticEquals(t *testing.T) {
	for _, tc := range []struct {
		name, config, device string
		want                 bool
	}{
		{"bare seconds", "120", "2m", true},
		{"clock form", "00:05:00", "5m", true},
		{"day clock form", "1d00:00:00", "1d", true},
		{"identical", "1h", "1h", true},
		{"equivalent spellings", "60s", "1m", true},
		{"genuinely different", "1h", "2h", false},
		{"keyword matches", "auto", "auto", true},
		{"keyword case differs", "Auto", "auto", true},
		{"keyword vs duration", "auto", "1m", false},
		{"different keywords", "auto", "disable-dpd", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, diags := newDurationValue(tc.config).StringSemanticEquals(context.Background(), newDurationValue(tc.device))
			if diags.HasError() {
				t.Fatalf("diags: %v", diags)
			}
			if got != tc.want {
				t.Errorf("semantic equality = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCIDRValueSemanticEquals(t *testing.T) {
	for _, tc := range []struct {
		name, config, device string
		want                 bool
	}{
		{"surrounding space", "  10.0.0.1/24  ", "10.0.0.1/24", true},
		{"identical", "10.0.0.1/24", "10.0.0.1/24", true},
		{"different prefix", "10.0.0.1/24", "10.0.0.1/25", false},
		{"different address", "10.0.0.1/24", "10.0.0.2/24", false},
		{"not an address", "not-an-ip", "10.0.0.1/24", false},
	} {
		t.Run(tc.name, func(t *testing.T) {
			got, diags := newCIDRValue(tc.config).StringSemanticEquals(context.Background(), newCIDRValue(tc.device))
			if diags.HasError() {
				t.Fatalf("diags: %v", diags)
			}
			if got != tc.want {
				t.Errorf("semantic equality = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestDurationAndCIDRNullsAreNeverEqual(t *testing.T) {
	if eq, _ := newDurationNull().StringSemanticEquals(context.Background(), newDurationValue("1m")); eq {
		t.Error("a null duration compared equal to a known one")
	}
	if eq, _ := newCIDRNull().StringSemanticEquals(context.Background(), newCIDRValue("10.0.0.1/24")); eq {
		t.Error("a null CIDR compared equal to a known one")
	}
}
