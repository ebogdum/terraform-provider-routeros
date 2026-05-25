package schemautil

import (
	"context"

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
		// Leave for the validator to reject.
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
	return normalizeStringPM{desc: "normalize RouterOS duration", fn: func(s string) (string, error) {
		d, err := client.ParseDuration(s)
		if err != nil {
			return "", err
		}
		return client.FormatDuration(d), nil
	}}
}
