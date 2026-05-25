// Package schemautil provides plan modifiers, validators, and helpers shared
// by hand-written and generated resources. It is the only Terraform-aware
// support library outside the provider package itself.
package schemautil

import (
	"context"
	"fmt"
	"regexp"

	"github.com/hashicorp/terraform-plugin-framework/schema/validator"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// --- string validators ---

type stringValidator struct {
	desc string
	fn   func(string) error
}

func (v stringValidator) Description(_ context.Context) string         { return v.desc }
func (v stringValidator) MarkdownDescription(_ context.Context) string { return v.desc }
func (v stringValidator) ValidateString(_ context.Context, req validator.StringRequest, resp *validator.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
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
	}}
}

// IsCIDR validates an IP/prefixlen pair.
func IsCIDR() validator.String {
	return stringValidator{desc: "must be in CIDR form (e.g. 10.0.0.0/24)", fn: func(s string) error {
		_, err := client.ParseCIDR(s)
		return err
	}}
}

// IsMAC validates a colon-separated EUI-48 MAC address.
func IsMAC() validator.String {
	return stringValidator{desc: "must be a MAC address (aa:bb:cc:dd:ee:ff)", fn: func(s string) error {
		_, err := client.ParseMAC(s)
		return err
	}}
}

// IsDurationRouterOS validates a RouterOS-style duration string ("1w2d3h4m5s",
// "30m", "120"). Bare integer = seconds.
func IsDurationRouterOS() validator.String {
	return stringValidator{desc: "must be a RouterOS duration (e.g. 1w2d3h, 30m, 120)", fn: func(s string) error {
		_, err := client.ParseDuration(s)
		return err
	}}
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
