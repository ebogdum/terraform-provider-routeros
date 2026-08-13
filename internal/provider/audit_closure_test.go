package provider

import (
	"context"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

func schemaOf(t *testing.T, r resource.Resource) resource.SchemaResponse {
	t.Helper()
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("schema diagnostics: %v", resp.Diagnostics)
	}
	return *resp
}

// queue_tree and queue_simple both declare place_before as Computed-only (no Optional), so the framework
// forbids ever setting it from config - plan.PlaceBefore is always Unknown pre-apply, making the Move() call
// that reads it permanently dead code. ip_firewall_filter's place_before must be genuinely settable from HCL.
func TestFirewallFilterPlaceBeforeIsSettable(t *testing.T) {
	attrs := schemaOf(t, NewIPFirewallFilterResource()).Schema.Attributes
	att, ok := attrs["place_before"]
	if !ok {
		t.Fatalf("routeros_ip_firewall_filter is missing place_before")
	}
	if !att.IsOptional() {
		t.Errorf("place_before must be Optional so it can be set from config; found Computed-only (matches the " +
			"dead-code bug in queue_tree/queue_simple)")
	}
}

// position and place_before order a rule against two different scopes and always run in a fixed internal
// sequence (position then place_before), so combining them would silently make place_before win with no
// indication that position's request was overridden. Test that setting both is rejected, and that setting
// either alone is not.
func TestFirewallFilterPositionAndPlaceBeforeAreMutuallyExclusive(t *testing.T) {
	cases := []struct {
		name         string
		position     types.Int64
		placeBefore  types.String
		wantConflict bool
	}{
		{"neither set", types.Int64Null(), types.StringNull(), false},
		{"position only", types.Int64Value(100), types.StringNull(), false},
		{"place_before only", types.Int64Null(), types.StringValue("*5"), false},
		{"both set", types.Int64Value(100), types.StringValue("*5"), true},
		{"position unknown, place_before set", types.Int64Unknown(), types.StringValue("*5"), false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			m := IPFirewallFilterModel{Position: tc.position, PlaceBefore: tc.placeBefore}
			got := firewallFilterPositionConflictsWithPlaceBefore(m)
			if got != tc.wantConflict {
				t.Errorf("conflict = %v, want %v", got, tc.wantConflict)
			}
		})
	}
}

// /caps-man/configuration was suspected of the same dotted sub-object defect as
// the wifi menus, because it exposes a `channel` section. It is not affected:
// it never flattened any sub-object member into a top-level attribute, so there
// is nothing addressing a wrong wire name. It simply does not offer the inline
// channel.*/datapath.*/security.* overrides, which is missing coverage rather
// than a bug. This test pins that, so the menu is not "fixed" by guesswork.
func TestCapsManConfigurationHasNoFlattenedSubObjects(t *testing.T) {
	attrs := schemaOf(t, NewCapsManConfigurationResource()).Schema.Attributes
	// Members that would only be reachable as channel.x / datapath.x / security.x.
	for _, member := range []string{
		"band", "frequency", "extension_channel", "control_channel_width",
		"bridge_cost", "client_to_client_forwarding", "local_forwarding",
		"authentication_types", "passphrase", "group_key_update",
	} {
		if _, ok := attrs[member]; ok {
			t.Errorf("caps_man_configuration exposes %q flat; if that is intended it must map to its dotted wire name", member)
		}
	}
	// The section references themselves are real top-level properties.
	for _, sec := range []string{"channel", "datapath", "security"} {
		if _, ok := attrs[sec]; !ok {
			t.Errorf("caps_man_configuration lost the %q profile reference", sec)
		}
	}
}

// The menu carries name/mirror-source/mirror-target/cpu-flow-control on the
// device; without them the resource had nothing to declare and emitted no
// usable configuration.
func TestEthernetSwitchAttrs(t *testing.T) {
	attrs := schemaOf(t, NewInterfaceEthernetSwitchResource()).Schema.Attributes
	for _, a := range []string{"name", "mirror_source", "mirror_target", "cpu_flow_control"} {
		if _, ok := attrs[a]; !ok {
			t.Errorf("routeros_interface_ethernet_switch is missing %q", a)
		}
	}
}

// Attributes holding read-only runtime state must not advertise themselves as
// settable: an Optional attribute that is never written is a silent no-op, and
// the user gets no feedback that their value was ignored.
func TestNoReadOnlyOptional(t *testing.T) {
	cases := []struct {
		newFn func() resource.Resource
		label string
		attrs []string
	}{
		{NewInterfaceListResource, "interface_list", []string{"dynamic", "builtin"}},
		{NewInterfaceBridgeVLANResource, "interface_bridge_vlan", []string{"current_tagged", "current_untagged", "dynamic"}},
		{NewInterfaceBridgePortResource, "interface_bridge_port", []string{"dynamic", "port_status"}},
		{NewInterfaceEoipResource, "interface_eoip", []string{"actual_mtu"}},
		{NewSystemScriptResource, "system_script", []string{"owner"}},
	}
	for _, tc := range cases {
		attrs := schemaOf(t, tc.newFn()).Schema.Attributes
		for _, a := range tc.attrs {
			att, ok := attrs[a]
			if !ok {
				continue
			}
			if att.IsOptional() {
				t.Errorf("%s.%s is read-only on the device but still advertised as Optional", tc.label, a)
			}
			if !att.IsComputed() {
				t.Errorf("%s.%s should be Computed so it still round-trips into state", tc.label, a)
			}
		}
	}
}

// Every menu the audit listed as having no provider resource now has one.
func TestUncoveredMenusHaveResources(t *testing.T) {
	want := []string{
		"routeros_interface_ethernet_switch_port",
		"routeros_interface_l2tp_server_server",
		"routeros_interface_lte_settings",
		"routeros_interface_macsec_profile",
		"routeros_interface_pptp_server_server",
		"routeros_interface_sstp_server_server",
		"routeros_ip_cloud_advanced",
		"routeros_ip_firewall_service_port",
		"routeros_ip_hotspot_service_port",
		"routeros_ip_ipsec_key_qkd",
		"routeros_ip_ipsec_policy_group",
		"routeros_ip_media_settings",
		"routeros_ip_neighbor_discovery_settings",
		"routeros_ip_traffic_flow_ipfix",
		"routeros_ipv6_dhcp_relay_option",
		"routeros_ipv6_nd_prefix_default",
		"routeros_queue_interface",
		"routeros_system_health_settings",
		"routeros_system_package_local_update_mirror",
		"routeros_system_resource_hardware_usb_settings",
		"routeros_system_resource_irq",
		"routeros_system_resource_irq_rps",
		"routeros_tool_mac_server_mac_winbox",
		// added earlier in the audit
		"routeros_ip_service",
		"routeros_system_routerboard_mode_button",
		"routeros_system_routerboard_wps_button",
		"routeros_system_routerboard_reset_button",
	}
	have := map[string]bool{}
	for _, f := range registryResources() {
		m := &resource.MetadataResponse{}
		f().Metadata(context.Background(), resource.MetadataRequest{ProviderTypeName: "routeros"}, m)
		have[m.TypeName] = true
	}
	for _, w := range want {
		if !have[w] {
			t.Errorf("%s is not registered", w)
		}
	}
}

// Verified against a live RouterOS 7.22 device by probing each wire key the
// provider writes: PATCH a nonexistent .id and read the error. RouterOS
// validates parameter names before the id lookup, so "unknown parameter <k>"
// proves the key does not exist while 404 proves it does. 335 keys the device
// rejected were corrected; these pin the two shapes that came out of it.
func TestBGPUsesDottedSubObjectKeys(t *testing.T) {
	src := readResource(t, "resource_routing_bgp_connection.go")
	for _, dotted := range []string{
		"input.accept-communities", "input.accept-nlri", "input.filter",
	} {
		if !strings.Contains(src, `body["`+dotted+`"]`) {
			t.Errorf("routing_bgp_connection does not write %q; the device rejects the flat spelling", dotted)
		}
	}
	// The flat spellings are rejected by RouterOS and must not be written.
	for _, flat := range []string{
		"input-accept-communities", "input-accept-nlri",
	} {
		if strings.Contains(src, `body["`+flat+`"]`) {
			t.Errorf("routing_bgp_connection still writes the rejected flat key %q", flat)
		}
	}
}

// Commands and read-only state modelled as properties: RouterOS rejects these
// outright, so writing one failed the whole apply.
func TestCommandsAreNotWrittenAsProperties(t *testing.T) {
	cases := map[string][]string{
		"resource_ip_arp.go":               {"ping", "torch", "make-static", "mac-telnet"},
		"resource_certificate.go":          {"sign", "import", "export", "revoke"},
		"resource_ip_dhcp_server_lease.go": {"make-static", "check-status"},
		"resource_system_script.go":        {"run-script"},
		"resource_user_group.go":           {"policies"},
	}
	for file, keys := range cases {
		src := readResource(t, file)
		for _, k := range keys {
			if strings.Contains(src, `body["`+k+`"]`) {
				t.Errorf("%s still writes %q, which RouterOS rejects", file, k)
			}
		}
	}
}

func readResource(t *testing.T, name string) string {
	t.Helper()
	b, err := os.ReadFile(name)
	if err != nil {
		t.Fatalf("read %s: %v", name, err)
	}
	return string(b)
}

// Exhaustive coverage pass: every settable argument the machine-generated CLI
// reference lists, cross-verified as accepted by a live device, is now
// expressible. These pin representative samples so the coverage cannot silently
// regress. Wire keys were confirmed on RouterOS 7.x by probing a nonexistent
// .id (a rejected key names itself in the error; an accepted one does not).
func TestExhaustiveCoverageSamples(t *testing.T) {
	cases := map[string][]string{
		// BGP dotted sub-objects the provider previously could not express
		"resource_routing_bgp_connection.go": {
			`body["input.accept-communities"]`, `body["output.no-early-cut"]`, `body["local.role"]`,
		},
		"resource_routing_isis_instance.go": {`body["l1.lsp-max-age"]`},
		// flat coverage gaps
		"resource_tool_netwatch.go":      {`body["down-script"]`, `body["check-certificate"]`},
		"resource_interface_bonding.go":  {`body["arp-interval"]`, `body["lacp-rate"]`},
		"resource_ip_hotspot_profile.go": {`body["radius-accounting"]`},
	}
	for file, keys := range cases {
		src := readResource(t, file)
		for _, k := range keys {
			if !strings.Contains(src, k) {
				t.Errorf("%s no longer writes %s (coverage regression)", file, k)
			}
		}
	}
}

// New secret-bearing fields added in the coverage pass must be redacted.
func TestNewSecretsAreSensitive(t *testing.T) {
	checks := map[string]string{
		"resource_caps_man_interface.go":           "passphrase",
		"resource_interface_l2tp_server_server.go": "ipsec_secret",
		"resource_ip_ipsec_identity.go":            "secret",
	}
	for file, attr := range checks {
		src := readResource(t, file)
		m := regexpFind(src, `"`+attr+`": schema\.StringAttribute\{[^}]*?\n\t\t\t\}`)
		if m == "" {
			t.Errorf("%s: attribute %q not found", file, attr)
			continue
		}
		if !strings.Contains(m, "Sensitive:") {
			t.Errorf("%s.%s is a secret but not marked Sensitive", file, attr)
		}
	}
}

func regexpFind(s, pat string) string {
	re := regexp.MustCompile(pat)
	return re.FindString(s)
}
