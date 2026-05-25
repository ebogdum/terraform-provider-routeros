// Package schemautil -- lockout_guard.go
//
// LockoutGuard refuses to apply firewall changes that would obviously sever
// management access to the router. It's checked at Create / Update time by
// the generated firewall resources.
//
// Rules flagged:
//   - chain=input action=drop|reject|tarpit with no match conditions
//   - chain=input action=drop|reject|tarpit with only invert-everything matches
//   - chain=forward action=drop|reject|tarpit with no match (less likely to
//     lock you out but commonly nukes routed traffic)
//
// The provider returns a hard error suggesting either an explicit
// `lockout_ack = true` override or adding a match clause. A real lock-out
// would have cost me my router; nobody else should hit it.
package schemautil

import (
	"errors"
	"strings"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// MatchKeys are the firewall arguments that narrow a rule. If a rule has none
// of these set, it matches everything in its chain.
var MatchKeys = []string{
	"src-address", "dst-address", "src-address-list", "dst-address-list",
	"src-address-type", "dst-address-type",
	"in-interface", "out-interface", "in-interface-list", "out-interface-list",
	"in-bridge-port", "out-bridge-port",
	"protocol", "src-port", "dst-port", "port",
	"connection-state", "connection-nat-state", "connection-mark",
	"connection-bytes", "connection-rate", "connection-type",
	"packet-mark", "routing-mark", "routing-table",
	"icmp-options", "ipv4-options", "tcp-flags", "tls-host",
	"layer7-protocol", "content", "p2p", "dscp",
	"limit", "time", "random",
	"hotspot", "ipsec-policy", "psd", "tcp-mss",
	"src-mac-address", "src-prefix",
}

// CheckFirewallLockout returns a hard error if the given body would create
// a wide-open drop/reject in the input or forward chain.
//
// Pass acknowledged=true to bypass (corresponds to user-set `lockout_ack`).
func CheckFirewallLockout(menuPath string, body client.Object, acknowledged bool) error {
	if acknowledged {
		return nil
	}
	if !isFirewallFilter(menuPath) {
		return nil
	}
	chain := strings.ToLower(body["chain"])
	action := strings.ToLower(body["action"])
	if chain != "input" && chain != "forward" {
		return nil
	}
	switch action {
	case "drop", "reject", "tarpit":
	default:
		return nil
	}
	if hasAnyMatch(body) {
		return nil
	}
	return errors.New("refusing to add a " + chain + "/" + action +
		" rule with no match conditions: this would sever management traffic. " +
		"Add at least one of " + strings.Join(MatchKeys[:6], ", ") +
		", or set lockout_ack=true to override.")
}

func isFirewallFilter(p string) bool {
	switch p {
	case "/ip/firewall/filter", "/ipv6/firewall/filter":
		return true
	}
	return false
}

func hasAnyMatch(body client.Object) bool {
	for _, k := range MatchKeys {
		if v, ok := body[k]; ok && strings.TrimSpace(v) != "" {
			return true
		}
	}
	return false
}
