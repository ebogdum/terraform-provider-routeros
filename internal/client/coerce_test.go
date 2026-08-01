package client

import (
	"testing"
	"time"
)

func TestParseFormatBool(t *testing.T) {
	for _, tc := range []struct {
		in   string
		want bool
	}{
		{"true", true}, {"yes", true}, {"YES", true},
		{"false", false}, {"no", false}, {"  False  ", false},
	} {
		got, err := ParseBool(tc.in)
		if err != nil || got != tc.want {
			t.Fatalf("ParseBool(%q) = %v,%v want %v", tc.in, got, err, tc.want)
		}
	}
	// RouterOS writes use yes/no (true/false is rejected by several menus).
	if FormatBool(true) != "yes" || FormatBool(false) != "no" {
		t.Fatal("FormatBool must emit yes/no for RouterOS writes")
	}
	if _, err := ParseBool("maybe"); err == nil {
		t.Fatal("want error on bogus bool")
	}
}

func TestParseInt64(t *testing.T) {
	cases := map[string]int64{
		"0":     0,
		"123":   123,
		"-7":    -7,
		"0x10":  16,
		"0X10":  16,
		"010":   8,
		"01000": 512,
	}
	for in, want := range cases {
		got, err := ParseInt64(in)
		if err != nil || got != want {
			t.Fatalf("ParseInt64(%q) = %d,%v want %d", in, got, err, want)
		}
	}
}

func TestParseFormatDurationRoundTrip(t *testing.T) {
	cases := map[string]time.Duration{
		"0s":         0,
		"30s":        30 * time.Second,
		"5m":         5 * time.Minute,
		"1h":         time.Hour,
		"1d":         24 * time.Hour,
		"1w":         7 * 24 * time.Hour,
		"1w2d3h4m5s": 7*24*time.Hour + 2*24*time.Hour + 3*time.Hour + 4*time.Minute + 5*time.Second,
		"120":        120 * time.Second, // bare integer = seconds
	}
	for in, wantD := range cases {
		gotD, err := ParseDuration(in)
		if err != nil {
			t.Fatalf("ParseDuration(%q): %v", in, err)
		}
		if gotD != wantD {
			t.Fatalf("ParseDuration(%q) = %v want %v", in, gotD, wantD)
		}
		// Round-trip via canonical form:
		canon := FormatDuration(gotD)
		gotD2, err := ParseDuration(canon)
		if err != nil || gotD2 != wantD {
			t.Fatalf("round trip %q -> %q -> %v (err %v)", in, canon, gotD2, err)
		}
	}
}

func TestCanonicalCIDR(t *testing.T) {
	cases := map[string]string{
		"10.0.0.1/24":     "10.0.0.1/24",
		"  192.168.1.1  ": "192.168.1.1",
		"::1/128":         "::1/128",
	}
	for in, want := range cases {
		got, err := CanonicalCIDR(in)
		if err != nil || got != want {
			t.Fatalf("CanonicalCIDR(%q) = %q,%v want %q", in, got, err, want)
		}
	}
	if _, err := CanonicalCIDR("not-an-ip/24"); err == nil {
		t.Fatal("want error")
	}
}

func TestCanonicalMAC(t *testing.T) {
	got, err := CanonicalMAC("aa:bb:cc:dd:ee:ff")
	if err != nil || got != "AA:BB:CC:DD:EE:FF" {
		t.Fatalf("CanonicalMAC = %q,%v", got, err)
	}
	if _, err := CanonicalMAC("nope"); err == nil {
		t.Fatal("want error")
	}
}

func TestParseFormatList(t *testing.T) {
	got := ParseList("a, b ,,c")
	want := []string{"a", "b", "c"}
	if len(got) != len(want) {
		t.Fatalf("len %d want %d", len(got), len(want))
	}
	for i, v := range got {
		if v != want[i] {
			t.Fatalf("idx %d = %q want %q", i, v, want[i])
		}
	}
	if FormatList(want) != "a,b,c" {
		t.Fatalf("FormatList = %q", FormatList(want))
	}
}
