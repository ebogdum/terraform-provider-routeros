package schemautil

import (
	"context"
	"fmt"
	"strings"

	"github.com/hashicorp/terraform-plugin-framework/resource/schema/planmodifier"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// Plan modifiers that suppress cosmetic-only diffs introduced by RouterOS
// canonicalising values on the wire. Each one is type-specific; attach the
// matching modifier to attributes that hold that semantic type.

type normalizeStringPM struct {
	desc string
	fn   func(string) (string, error)
}

func (m normalizeStringPM) Description(_ context.Context) string         { return m.desc }
func (m normalizeStringPM) MarkdownDescription(_ context.Context) string { return m.desc }
func (m normalizeStringPM) PlanModifyString(_ context.Context, req planmodifier.StringRequest, resp *planmodifier.StringResponse) {
	if req.ConfigValue.IsNull() || req.ConfigValue.IsUnknown() {
		return
	}
	canon, err := m.fn(req.ConfigValue.ValueString())
	if err != nil {
		// Refuse to plan: the value cannot be normalised. A validator may
		// also reject it, but plan modifiers can run independently of (or
		// before) validators, so we surface the error here too.
		resp.Diagnostics.AddAttributeError(req.Path, "Invalid value", err.Error())
		return
	}
	// If state already equals the canonical form, keep state to avoid a diff.
	if !req.StateValue.IsNull() && req.StateValue.ValueString() == canon {
		resp.PlanValue = req.StateValue
	}
}

// NormalizeCIDR canonicalises "10.0.0.1/24" / spaces / case so that
// equivalent values don't produce a diff.
func NormalizeCIDR() planmodifier.String {
	return normalizeStringPM{desc: "normalize CIDR form", fn: client.CanonicalCIDR}
}

// NormalizeMAC upper-cases and validates a MAC.
func NormalizeMAC() planmodifier.String {
	return normalizeStringPM{desc: "normalize MAC to upper-case colon form", fn: client.CanonicalMAC}
}

// NormalizeDuration converts equivalent duration strings to RouterOS canonical
// (1w2d3h4m5s).
func NormalizeDuration() planmodifier.String {
	return NormalizeDurationExcept()
}

// NormalizeDurationExcept is NormalizeDuration that leaves the given sentinel
// words alone.
//
// Attributes like arp-timeout and dpd-interval hold either a duration or a magic
// word (`auto`, `disable-dpd`). Running those words through the duration parser
// raises "Invalid value" at plan time, so a router sitting on its own default
// cannot be planned at all. Keywords are matched case-insensitively and passed
// through verbatim; anything else is normalised as usual.
func NormalizeDurationExcept(keywords ...string) planmodifier.String {
	set := make(map[string]struct{}, len(keywords))
	for _, k := range keywords {
		set[strings.ToLower(k)] = struct{}{}
	}
	desc := "normalize RouterOS duration"
	if len(keywords) > 0 {
		desc += fmt.Sprintf(" (passing through %v)", keywords)
	}
	return normalizeStringPM{desc: desc, fn: func(s string) (string, error) {
		if _, ok := set[strings.ToLower(strings.TrimSpace(s))]; ok {
			return s, nil
		}
		d, err := client.ParseDuration(s)
		if err != nil {
			return "", err
		}
		return client.FormatDuration(d), nil
	}}
}

// NormalizeCase rewrites a value to the spelling RouterOS itself reports, so a
// config written as `md5` does not permadiff against a device that answers
// `MD5`. Unlisted values are left untouched -- a validator, not this modifier,
// is responsible for rejecting them.
func NormalizeCase(canonical ...string) planmodifier.String {
	byLower := make(map[string]string, len(canonical))
	for _, v := range canonical {
		byLower[strings.ToLower(v)] = v
	}
	return normalizeStringPM{
		desc: fmt.Sprintf("normalize case to one of %v", canonical),
		fn: func(s string) (string, error) {
			if c, ok := byLower[strings.ToLower(strings.TrimSpace(s))]; ok {
				return c, nil
			}
			return s, nil
		},
	}
}
