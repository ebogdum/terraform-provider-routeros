package client

import (
	"fmt"
	"net"
	"regexp"
	"strconv"
	"strings"
	"time"
)

// RouterOS encodes all JSON values as strings. These helpers convert between
// Go-typed values and the canonical RouterOS string form, with round-trip
// stability -- Marshal(Unmarshal(s)) == s for any RouterOS-emitted value.

// --- bool ---

// ParseBool accepts true/false/yes/no (case-insensitive). RouterOS emits
// "true"/"false" in JSON responses but accepts "yes"/"no" on input.
func ParseBool(s string) (bool, error) {
	switch strings.ToLower(strings.TrimSpace(s)) {
	case "true", "yes":
		return true, nil
	case "false", "no":
		return false, nil
	}
	return false, fmt.Errorf("routeros: %q is not a bool", s)
}

// FormatBool renders a bool for a RouterOS write body as "yes"/"no". RouterOS
// reads bools back as "true"/"false" in JSON but several menus (bridge-port
// `hw`, dns-static `match-subdomain`, ...) reject "true"/"false" on write with
// "must be either yes or no". "yes"/"no" is RouterOS's native input form and is
// accepted for every bool, so it is the safe universal choice. Reads use
// ParseBool, which accepts both forms.
func FormatBool(b bool) string {
	if b {
		return "yes"
	}
	return "no"
}

// --- int ---

// ParseInt64 accepts decimal, 0x-prefixed hex, and 0-prefixed octal, matching
// what RouterOS accepts on input (per REST docs).
func ParseInt64(s string) (int64, error) {
	s = strings.TrimSpace(s)
	if s == "" {
		return 0, fmt.Errorf("routeros: empty int")
	}
	switch {
	case strings.HasPrefix(s, "0x") || strings.HasPrefix(s, "0X"):
		return strconv.ParseInt(s[2:], 16, 64)
	case len(s) > 1 && s[0] == '0' && s[1] >= '0' && s[1] <= '9':
		return strconv.ParseInt(s, 8, 64)
	}
	return strconv.ParseInt(s, 10, 64)
}

func FormatInt64(v int64) string { return strconv.FormatInt(v, 10) }

// --- duration ---

// RouterOS time format: combination of <n>w<n>d<n>h<n>m<n>s, e.g. "1w2d3h4m5s",
// "30m", "1d", or just "120" (= 120 seconds). Both directions canonicalise to
// the most compact form composed of the largest non-zero units used in the
// input, with seconds as the smallest unit. Bare integers without a unit are
// treated as seconds (matching RouterOS).
var durationRe = regexp.MustCompile(`^(?:(\d+)w)?(?:(\d+)d)?(?:(\d+)h)?(?:(\d+)m)?(?:(\d+)s)?$`)

func ParseDuration(s string) (time.Duration, error) {
	s = strings.TrimSpace(strings.ToLower(s))
	if s == "" || s == "none" || s == "never" || s == "0s" || s == "0" {
		return 0, nil
	}
	if onlyDigits(s) {
		secs, err := strconv.ParseInt(s, 10, 64)
		if err != nil {
			return 0, fmt.Errorf("routeros: bad duration %q: %w", s, err)
		}
		return time.Duration(secs) * time.Second, nil
	}
	m := durationRe.FindStringSubmatch(s)
	if m == nil {
		// Fallback: Go's time.ParseDuration handles "1h2m3s", "500ms" etc.
		// RouterOS doesn't use ms, but be permissive on input.
		d, err := time.ParseDuration(s)
		if err != nil {
			return 0, fmt.Errorf("routeros: bad duration %q", s)
		}
		return d, nil
	}
	var d time.Duration
	weeks, _ := strconv.Atoi(m[1])
	days, _ := strconv.Atoi(m[2])
	hours, _ := strconv.Atoi(m[3])
	mins, _ := strconv.Atoi(m[4])
	secs, _ := strconv.Atoi(m[5])
	d += time.Duration(weeks) * 7 * 24 * time.Hour
	d += time.Duration(days) * 24 * time.Hour
	d += time.Duration(hours) * time.Hour
	d += time.Duration(mins) * time.Minute
	d += time.Duration(secs) * time.Second
	return d, nil
}

func FormatDuration(d time.Duration) string {
	if d <= 0 {
		return "0s"
	}
	secs := int64(d.Seconds())
	weeks := secs / (7 * 24 * 3600)
	secs -= weeks * 7 * 24 * 3600
	days := secs / (24 * 3600)
	secs -= days * 24 * 3600
	hours := secs / 3600
	secs -= hours * 3600
	mins := secs / 60
	secs -= mins * 60

	var b strings.Builder
	if weeks > 0 {
		fmt.Fprintf(&b, "%dw", weeks)
	}
	if days > 0 {
		fmt.Fprintf(&b, "%dd", days)
	}
	if hours > 0 {
		fmt.Fprintf(&b, "%dh", hours)
	}
	if mins > 0 {
		fmt.Fprintf(&b, "%dm", mins)
	}
	if secs > 0 || b.Len() == 0 {
		fmt.Fprintf(&b, "%ds", secs)
	}
	return b.String()
}

func onlyDigits(s string) bool {
	if s == "" {
		return false
	}
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}

// --- IP / CIDR / MAC ---

func ParseIP(s string) (net.IP, error) {
	ip := net.ParseIP(strings.TrimSpace(s))
	if ip == nil {
		return nil, fmt.Errorf("routeros: %q is not an IP", s)
	}
	return ip, nil
}

func ParseCIDR(s string) (*net.IPNet, error) {
	_, n, err := net.ParseCIDR(strings.TrimSpace(s))
	if err != nil {
		return nil, fmt.Errorf("routeros: %q is not a CIDR: %w", s, err)
	}
	return n, nil
}

// CanonicalCIDR returns s in "ip/prefixlen" form. The IP part is the address
// as given (not zeroed), matching RouterOS behaviour where /ip/address stores
// the host address, not the network.
func CanonicalCIDR(s string) (string, error) {
	s = strings.TrimSpace(s)
	idx := strings.IndexByte(s, '/')
	if idx < 0 {
		ip, err := ParseIP(s)
		if err != nil {
			return "", err
		}
		return ip.String(), nil
	}
	ipStr, mask := s[:idx], s[idx+1:]
	ip := net.ParseIP(ipStr)
	if ip == nil {
		return "", fmt.Errorf("routeros: %q is not an IP", ipStr)
	}
	m, err := strconv.Atoi(mask)
	if err != nil {
		return "", fmt.Errorf("routeros: %q is not a prefix length", mask)
	}
	bits := 32
	if ip.To4() == nil {
		bits = 128
	}
	if m < 0 || m > bits {
		return "", fmt.Errorf("routeros: prefix /%d out of range for %s", m, ip)
	}
	return fmt.Sprintf("%s/%d", ip.String(), m), nil
}

var macRe = regexp.MustCompile(`^[0-9A-Fa-f]{2}(:[0-9A-Fa-f]{2}){5}$`)

func ParseMAC(s string) (net.HardwareAddr, error) {
	s = strings.TrimSpace(s)
	if !macRe.MatchString(s) {
		return nil, fmt.Errorf("routeros: %q is not a MAC address", s)
	}
	return net.ParseMAC(strings.ToUpper(s))
}

func CanonicalMAC(s string) (string, error) {
	m, err := ParseMAC(s)
	if err != nil {
		return "", err
	}
	return strings.ToUpper(m.String()), nil
}

// --- list[T] ---
//
// RouterOS comma-separated lists. Whitespace around items is trimmed; empty
// items are dropped.

func ParseList(s string) []string {
	s = strings.TrimSpace(s)
	if s == "" {
		return nil
	}
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" {
			out = append(out, p)
		}
	}
	return out
}

func FormatList(items []string) string {
	return strings.Join(items, ",")
}
