// Package schemautil provides plan modifiers, validators, and helpers shared
// by hand-written and generated resources. It is the only Terraform-aware
// support library outside the provider package itself.
package schemautil

import (
	"context"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// --- string validators ---

type stringValidator struct {
	desc string
	fn   func(string) error
	// emptyOK marks a format validator, where "" means the attribute is unset
	// rather than malformed. Enumerations leave it false.
	emptyOK bool
}

func (v stringValidator) Description(_ context.Context) string         { return v.desc }
func (v stringValidator) MarkdownDescription(_ context.Context) string { return v.desc }
func (v stringValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	if v.emptyOK && req.ConfigValue.ValueString() == "" {
		return
	}
	if err := v.fn(req.ConfigValue.ValueString()); err != nil {
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid value", err.Error())
	}
}

// IsIP validates a bare IPv4 or IPv6 address.
func IsIP() validator.String {
	return stringValidator{desc: "must be an IPv4 or IPv6 address", fn: func(s string) error {
		_, err := client.ParseIP(s)
		return err
	}, emptyOK: true}
}

// IsCIDR validates an IP/prefixlen pair.
func IsCIDR() validator.String {
	return stringValidator{desc: "must be in CIDR form (e.g. 10.0.0.0/24)", fn: func(s string) error {
		_, err := client.ParseCIDR(s)
		return err
	}, emptyOK: true}
}

// IsMAC validates a colon-separated EUI-48 MAC address.
func IsMAC() validator.String {
	return stringValidator{desc: "must be a MAC address (aa:bb:cc:dd:ee:ff)", fn: func(s string) error {
		_, err := client.ParseMAC(s)
		return err
	}, emptyOK: true}
}

// IsDurationRouterOS validates a RouterOS-style duration string ("1w2d3h4m5s",
// "30m", "120"). Bare integer = seconds.
func IsDurationRouterOS() validator.String {
	return stringValidator{desc: "must be a RouterOS duration (e.g. 1w2d3h, 30m, 120)", fn: func(s string) error {
		_, err := client.ParseDuration(s)
		return err
	}, emptyOK: true}
}

// OneOf restricts a string to one of the listed values (case-sensitive).
func OneOf(values ...string) validator.String {
	set := make(map[string]struct{}, len(values))
	for _, v := range values {
		set[v] = struct{}{}
	}
	return stringValidator{desc: fmt.Sprintf("must be one of %v", values), fn: func(s string) error {
		if _, ok := set[s]; !ok {
			return fmt.Errorf("%q not in %v", s, values)
		}
		return nil
	}}
}

// OneOfFold restricts a string to one of the listed values, ignoring case.
//
// Several RouterOS menus report an enum in a different case than they accept it
// in -- /snmp/community, for instance, prints authentication-protocol=MD5 and
// encryption-protocol=DES while the CLI documents them lower-case. A
// case-sensitive OneOf rejects the device's own factory default, which makes the
// menu impossible to import or manage. Pair this with NormalizeCase so the
// canonical spelling reaches state and no permadiff appears.
func OneOfFold(values ...string) validator.String {
	set := make(map[string]struct{}, len(values))
	for _, v := range values {
		set[strings.ToLower(v)] = struct{}{}
	}
	return stringValidator{desc: fmt.Sprintf("must be one of %v (case-insensitive)", values), fn: func(s string) error {
		if _, ok := set[strings.ToLower(s)]; !ok {
			return fmt.Errorf("%q not in %v", s, values)
		}
		return nil
	}}
}

// IsDurationOrKeyword validates a RouterOS duration or one of the sentinel words
// the menu accepts in place of one.
//
// RouterOS mixes durations and magic words in the same property: arp-timeout is
// a duration or `auto`, dpd-interval is a duration or `disable-dpd`. Validating
// such a property as a pure duration rejects the router's own default and makes
// the attribute unusable. Keywords match case-insensitively.
func IsDurationOrKeyword(keywords ...string) validator.String {
	set := make(map[string]struct{}, len(keywords))
	for _, k := range keywords {
		set[strings.ToLower(k)] = struct{}{}
	}
	desc := fmt.Sprintf("must be a RouterOS duration (e.g. 1w2d3h, 30m, 120) or one of %v", keywords)
	return stringValidator{desc: desc, fn: func(s string) error {
		if _, ok := set[strings.ToLower(strings.TrimSpace(s))]; ok {
			return nil
		}
		if _, err := client.ParseDuration(s); err != nil {
			return fmt.Errorf("%q is neither a duration nor one of %v", s, keywords)
		}
		return nil
	}}
}

// RegexMatch checks the value against re.
func RegexMatch(re *regexp.Regexp, desc string) validator.String {
	if desc == "" {
		desc = "must match " + re.String()
	}
	return stringValidator{desc: desc, fn: func(s string) error {
		if !re.MatchString(s) {
			return fmt.Errorf("does not match %s", re.String())
		}
		return nil
	}}
}

// IsTimeOfDayOrStartup accepts what /system/scheduler start-time accepts: a
// clock time, or the keyword "startup".
func IsTimeOfDayOrStartup() validator.String {
	return stringValidator{desc: `must be a time of day (HH:MM:SS) or "startup"`, fn: func(s string) error {
		_, err := client.CanonicalTimeOfDay(s)
		return err
	}, emptyOK: true}
}

// IsDSCPOrInherit accepts a DSCP code point (0-63) or the keyword "inherit",
// which is what RouterOS's tunnel menus take.
func IsDSCPOrInherit() validator.String {
	return stringValidator{desc: `must be 0-63 or "inherit"`, fn: func(s string) error {
		if strings.EqualFold(strings.TrimSpace(s), "inherit") {
			return nil
		}
		n, err := strconv.Atoi(strings.TrimSpace(s))
		if err != nil || n < 0 || n > 63 {
			return fmt.Errorf("%q is not a DSCP value (0-63) or \"inherit\"", s)
		}
		return nil
	}, emptyOK: true}
}
