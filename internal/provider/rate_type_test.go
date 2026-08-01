package provider

import (
	"context"
	"testing"
)

func TestCanonRate(t *testing.T) {
	cases := map[string]string{
		"10M":            "10000000/10000000",
		"10M/10M":        "10000000/10000000",
		"1k":             "1000/1000",
		"1G":             "1000000000/1000000000",
		"512":            "512/512",
		"1500000/300000": "1500000/300000",
		"2M/512k":        "2000000/512000",
	}
	for in, want := range cases {
		got, ok := canonRate(in)
		if !ok || got != want {
			t.Errorf("canonRate(%q) = %q,%v want %q", in, got, ok, want)
		}
	}
	// Non-rate strings do not parse.
	if _, ok := canonRate("unlimited"); ok {
		t.Error("unlimited should not parse as a rate")
	}
}

func TestRosRateSemanticEquals(t *testing.T) {
	ctx := context.Background()
	eq := func(a, b string) bool {
		ok, d := newRosRateValue(a).StringSemanticEquals(ctx, newRosRateValue(b))
		if d.HasError() {
			t.Fatalf("diags: %v", d)
		}
		return ok
	}
	if !eq("10M/10M", "10000000/10000000") {
		t.Error("10M/10M must equal its expanded form")
	}
	if !eq("1k", "1000") {
		t.Error("1k must equal 1000")
	}
	if eq("10M", "20M") {
		t.Error("distinct rates must not be equal")
	}
	// Two unparseable values fall back to literal comparison.
	if !eq("unlimited", "unlimited") {
		t.Error("identical literals must be equal")
	}
}
