package provider

import (
	"context"
	"testing"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// RouterOS reports a word where the menu documents a number: bridge mtu=auto,
// bridge/port horizon=none, /ipv6/nd hop-limit=unspecified. When these were
// modelled as int64 the value failed to parse, landed in state as null, and the
// sentinel could never be written back -- so a device sitting on its factory
// defaults produced a permanent diff.
func TestSentinelIntBridgeMTU(t *testing.T) {
	var m InterfaceBridgeModel
	interfaceBridgeApply(context.Background(), client.Object{
		".id": "*1", "name": "iot-lan", "mtu": "auto", "max-learned-entries": "auto",
	}, &m)
	if got := m.MTU.ValueString(); got != "auto" {
		t.Fatalf("bridge mtu = %q, want %q", got, "auto")
	}
	if got := m.MaxLearnedEntries.ValueString(); got != "auto" {
		t.Fatalf("bridge max-learned-entries = %q, want %q", got, "auto")
	}
	// A numeric value must still round-trip unchanged.
	interfaceBridgeApply(context.Background(), client.Object{".id": "*1", "mtu": "1500"}, &m)
	if got := m.MTU.ValueString(); got != "1500" {
		t.Fatalf("bridge mtu = %q, want %q", got, "1500")
	}
}

func TestSentinelIntBridgePortHorizon(t *testing.T) {
	var m InterfaceBridgePortModel
	interfaceBridgePortApply(context.Background(), client.Object{
		".id": "*1", "bridge": "iot-lan", "horizon": "none",
	}, &m)
	if got := m.Horizon.ValueString(); got != "none" {
		t.Fatalf("horizon = %q, want %q", got, "none")
	}
	interfaceBridgePortApply(context.Background(), client.Object{".id": "*1", "horizon": "1"}, &m)
	if got := m.Horizon.ValueString(); got != "1" {
		t.Fatalf("horizon = %q, want %q", got, "1")
	}
}

func TestSentinelIntIPv6NdUnspecified(t *testing.T) {
	var m IPV6NdModel
	iPV6NdApply(context.Background(), client.Object{
		".id":            "*0",
		"hop-limit":      "unspecified",
		"mtu":            "unspecified",
		"reachable-time": "unspecified",
		// retransmit-interval intentionally absent: must become null, not "".
	}, &m)
	for name, got := range map[string]string{
		"hop-limit":      m.HopLimit.ValueString(),
		"mtu":            m.MTU.ValueString(),
		"reachable-time": m.ReachableTime.ValueString(),
	} {
		if got != "unspecified" {
			t.Errorf("ipv6 nd %s = %q, want %q", name, got, "unspecified")
		}
	}
	if !m.RetransmitInterval.IsNull() {
		t.Errorf("absent retransmit-interval = %q, want null", m.RetransmitInterval.ValueString())
	}
}
