package provider

import (
	"regexp"
	"testing"
)

// Keys that a live RouterOS device rejects ("unknown parameter") when the
// provider writes them. Verified by probing both a 7.23 CHR and a 7.22.3
// hAP ax³ (PATCH to a nonexistent .id, which validates the parameter name
// without mutating anything). Each entry is a menu file -> wire keys that must
// NOT appear in a request body. Read-only state and console commands modelled
// as properties, plus mangled wire names, all land here.
var forbiddenWrites = map[string][]string{
	"resource_certificate.go":                            {"dynamic", "private-key", "type"},
	"resource_certificate_scep_server.go":                {"ca-certificate"},
	"resource_file.go":                                   {"directory", "file-name"},
	"resource_interface_bridge.go":                       {"type"},
	"resource_ip_dhcp_server.go":                         {"address-list"},
	"resource_ip_dhcp_server_lease.go":                   {"address-list", "dhcp-options"},
	"resource_ip_dhcp_server_network.go":                 {"dhcp-options", "dynamic"},
	"resource_ip_ipsec_peer.go":                          {"responder"},
	"resource_ip_ipsec_policy_group.go":                  {"default"},
	"resource_ip_pool.go":                                {"addresses"},
	"resource_ip_smb.go":                                 {"interface"},
	"resource_partition.go":                              {"activate", "active", "running"},
	"resource_routing_bgp_connection.go":                 {"template", "remote-port"},
	"resource_user.go":                                   {"type"},
	"resource_interface_wifi_interworking.go":            {"hotspot-2-0"},
	"resource_tool_traffic_generator_packet_template.go": {"dst-port", "src-port", "ttl", "protocol", "gateway", "priority", "hop-limit"},
}

func TestNoPhantomWrites(t *testing.T) {
	for file, keys := range forbiddenWrites {
		src := readResource(t, file)
		for _, k := range keys {
			re := regexp.MustCompile(`body\["` + regexp.QuoteMeta(k) + `"\]\s*=`)
			if re.MatchString(src) {
				t.Errorf("%s writes %q, which RouterOS rejects (device-verified phantom write)", file, k)
			}
		}
	}
}

// The dotted remaps must use the form the device accepts.
func TestDottedRemapsPresent(t *testing.T) {
	src := readResource(t, "resource_routing_bgp_connection.go")
	if !regexp.MustCompile(`body\["remote\.port"\]`).MatchString(src) {
		t.Error("routing_bgp_connection must write remote.port, not remote-port")
	}
	iw := readResource(t, "resource_interface_wifi_interworking.go")
	if !regexp.MustCompile(`body\["hotspot20"\]`).MatchString(iw) {
		t.Error("interface_wifi_interworking must write hotspot20, not hotspot-2-0")
	}
}
