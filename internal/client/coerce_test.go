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

// Every accepted form was PUT to a ROS 7.23.2 /ip/dhcp-server/lease and came
// back as upper-case colon form.
func TestCanonicalMAC(t *testing.T) {
	for in, want := range map[string]string{
		"aa:bb:cc:dd:ee:ff":   "AA:BB:CC:DD:EE:FF",
		"AA:BB:CC:DD:EE:FF":   "AA:BB:CC:DD:EE:FF",
		"Aa-Bb-Cc-Dd-Ee-01":   "AA:BB:CC:DD:EE:01",
		"aabbccddee02":        "AA:BB:CC:DD:EE:02",
		" AA:BB:CC:DD:EE:FF ": "AA:BB:CC:DD:EE:FF",
	} {
		got, err := CanonicalMAC(in)
		if err != nil || got != want {
			t.Errorf("CanonicalMAC(%q) = %q,%v want %q", in, got, err, want)
		}
	}
	for _, bad := range []string{"nope", "", "aa:bb:cc:dd:ee", "aa:bb:cc:dd:ee:gg", "aabbccddee", "AA:BB:CC:DD:EE:FF/FF:FF:FF:FF:FF:FF"} {
		if _, err := CanonicalMAC(bad); err == nil {
			t.Errorf("CanonicalMAC(%q) accepted, want error", bad)
		}
	}
}

// RouterOS prints and accepts a clock form alongside the unit form; the parser
// rejected both, so an interval copied off the router would not plan.
func TestParseDurationClockForm(t *testing.T) {
	for in, want := range map[string]time.Duration{
		"00:05:00":   5 * time.Minute,
		"1d00:00:00": 24 * time.Hour,
		"23:59:59":   23*time.Hour + 59*time.Minute + 59*time.Second,
		"2d03:04:05": 2*24*time.Hour + 3*time.Hour + 4*time.Minute + 5*time.Second,
	} {
		got, err := ParseDuration(in)
		if err != nil || got != want {
			t.Errorf("ParseDuration(%q) = %v,%v want %v", in, got, err, want)
		}
	}
	for _, bad := range []string{"00:60:00", "00:00:60", "nonsense"} {
		if _, err := ParseDuration(bad); err == nil {
			t.Errorf("ParseDuration(%q) accepted, want error", bad)
		}
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
