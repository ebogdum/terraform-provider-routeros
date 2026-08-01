package schemautil

import (
	"testing"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

func TestWatchdogGuard(t *testing.T) {
	// Unroutable / documentation addresses reboot-loop the device — must be refused.
	for _, addr := range []string{
		"192.0.2.1",    // TEST-NET-1 (the address that bricked the device)
		"198.51.100.5", // TEST-NET-2
		"203.0.113.9",  // TEST-NET-3
		"127.0.0.1",    // loopback
		"0.0.0.0",      // unspecified
		"169.254.1.1",  // link-local
		"224.0.0.1",    // multicast
		"240.0.0.1",    // reserved
	} {
		if err := CheckWatchdogLockout(client.Object{"watch-address": addr}, false); err == nil {
			t.Errorf("watch-address=%s was allowed; it can never answer pings and reboot-loops the device", addr)
		}
		// ack overrides
		if err := CheckWatchdogLockout(client.Object{"watch-address": addr}, true); err != nil {
			t.Errorf("lockout_ack did not override for %s: %v", addr, err)
		}
	}
	// Reachable addresses and 'none'/empty are allowed.
	for _, addr := range []string{"8.8.8.8", "192.168.10.1", "1.1.1.1", "none", ""} {
		if err := CheckWatchdogLockout(client.Object{"watch-address": addr}, false); err != nil {
			t.Errorf("reachable/none watch-address=%q was refused: %v", addr, err)
		}
	}
	// A non-IP (name) is not judged here.
	if err := CheckWatchdogLockout(client.Object{"watch-address": "gateway.local"}, false); err != nil {
		t.Errorf("hostname watch-address refused: %v", err)
	}
}
