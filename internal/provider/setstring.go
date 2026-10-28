package provider

import (
	"sort"
	"strings"
)

// canonicalSetString normalises a comma-separated RouterOS list value.
//
// RouterOS returns list-valued properties (`/interface/bridge/vlan` tagged,
// untagged, vlan-ids, ...) in whatever order it stores them, which is not
// necessarily the order the operator wrote. Because the provider models these
// properties as strings, a plain string comparison reports a diff that no
// apply can settle.
//
// Sorting the members gives both sides a stable comparison.
func canonicalSetString(s string) string {
	parts := strings.Split(s, ",")
	out := make([]string, 0, len(parts))
	for _, p := range parts {
		if t := strings.TrimSpace(p); t != "" {
			out = append(out, t)
		}
	}
	sort.Strings(out)
	return strings.Join(out, ",")
}

// preferSetSpelling returns the value to store in state for a list-valued
// property given what the configuration/prior state holds and what the device
// reports.
//
// If the two describe the same set, the configured spelling is kept: the
// device is free to reorder the members, and adopting its order would make the
// next plan report a change that no apply can settle. If the sets differ the
// device value wins — that is real drift.
//
// A plan modifier that rewrites the configured value is deliberately avoided:
// OpenTofu then reports an in-place update on every plan (the planned value
// never equals the configuration), even when nothing about the resource
// changes.
func preferSetSpelling(prior, device string) string {
	if prior != "" && device != "" && canonicalSetString(prior) == canonicalSetString(device) {
		return prior
	}
	return device
}
