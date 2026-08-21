package provider

import (
	"context"
	"os"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/path"
	"github.com/hashicorp/terraform-plugin-framework/resource"
	"github.com/hashicorp/terraform-plugin-framework/resource/schema"
	"github.com/hashicorp/terraform-plugin-framework/schema/validator"
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

func TestIPDNSStaticAddressIsOptional(t *testing.T) {
	attrs := schemaOf(t, NewIPDNSStaticResource()).Schema.Attributes
	att, ok := attrs["address"].(schema.StringAttribute)
	if !ok {
		t.Fatalf("routeros_ip_dns_static.address missing or not a StringAttribute")
	}
	if att.Required {
		t.Error("address is Required; should be Optional so non-A/AAAA entries can omit it")
	}
	if !att.Optional || !att.Computed {
		t.Error("address should be Optional and Computed, matching the resource's other type-specific fields")
	}
}

// Test which /ip/dns/static types support address. Ie 'A' or 'AAAA' records versus every other type, which have their own
// dedicated field.
func TestIPDNSStaticTypeNeedsAddress(t *testing.T) {
	cases := []struct {
		typ  string
		want bool
	}{
		{"", true}, // unset means the device default, "A"
		{"A", true},
		{"AAAA", true},
		{"NXDOMAIN", false},
		{"CNAME", false},
		{"FWD", false},
		{"MX", false},
		{"NS", false},
		{"SRV", false},
		{"TXT", false},
	}
	for _, tc := range cases {
		if got := ipDNSStaticTypeNeedsAddress(tc.typ); got != tc.want {
			t.Errorf("ipDNSStaticTypeNeedsAddress(%q) = %v, want %v", tc.typ, got, tc.want)
		}
	}
}

// Confirmed on RouterOS 7.23.2: Create silently drops address on a non-A/AAAA
// type, and Update silently rewrites the record to type A, losing the srv fields.
func TestIPDNSStaticAddressTypeConflict(t *testing.T) {
	cases := []struct {
		name       string
		typ        string
		hasAddress bool
		wantErr    bool
	}{
		{"A with address", "A", true, false},
		{"default type with address", "", true, false},
		{"AAAA with address", "AAAA", true, false},
		{"A without address", "A", false, true},
		{"default type without address", "", false, true},
		{"SRV without address", "SRV", false, false},
		{"SRV with address", "SRV", true, true},
		{"TXT with address", "TXT", true, true},
		{"CNAME with address", "CNAME", true, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			summary, detail := ipDNSStaticAddressTypeConflict(tc.typ, tc.hasAddress)
			gotErr := summary != ""
			if gotErr != tc.wantErr {
				t.Errorf("conflict(%q, %v) = (%q, %q), want error=%v", tc.typ, tc.hasAddress, summary, detail, tc.wantErr)
			}
		})
	}
}

// RouterOS rejects an empty address on both Create and Update with "invalid value
// for argument ip/ipv6", so neither path may write one.
func TestIPDNSStaticShouldWriteAddress(t *testing.T) {
	cases := []struct {
		name        string
		planAddr    types.String
		stateAddr   types.String
		wantWritten bool
	}{
		{"both null (typical non-A/AAAA record)", types.StringNull(), types.StringNull(), false},
		{"empty unchanged", types.StringValue(""), types.StringValue(""), false},
		{"empty, was null", types.StringValue(""), types.StringNull(), false},
		{"null, was empty (pre-fix state converging)", types.StringNull(), types.StringValue(""), false},
		{"real change", types.StringValue("10.0.0.1"), types.StringValue("10.0.0.2"), true},
		{"real value unchanged", types.StringValue("10.0.0.1"), types.StringValue("10.0.0.1"), false},
		{"unknown", types.StringUnknown(), types.StringValue(""), false},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			plan := IPDNSStaticModel{Address: tc.planAddr}
			state := IPDNSStaticModel{Address: tc.stateAddr}
			got := ipDNSStaticShouldWriteAddress(plan, state)
			if got != tc.wantWritten {
				t.Errorf("shouldWriteAddress = %v, want %v", got, tc.wantWritten)
			}
		})
	}
}

// The menu carries name/mirror-source/mirror-target/cpu-flow-control/l3-hw-offloading on the
// device; without them the resource had nothing to declare and emitted no
// usable configuration.
func TestEthernetSwitchAttrs(t *testing.T) {
	attrs := schemaOf(t, NewInterfaceEthernetSwitchResource()).Schema.Attributes
	for _, a := range []string{"name", "mirror_source", "mirror_target", "cpu_flow_control", "l3_hw_offloading"} {
		if _, ok := attrs[a]; !ok {
			t.Errorf("routeros_interface_ethernet_switch is missing %q", a)
		}
	}
}

// start_time previously carried a OneOf(["startup"]) validator, rejecting every HH:MM:SS
// time that ROS accepts. Check validator accepts a time, so the restriction can't return.
func TestSchedulerStartTimeAcceptsConcreteTime(t *testing.T) {
	attrs := schemaOf(t, NewSystemSchedulerResource()).Schema.Attributes
	att, ok := attrs["start_time"].(schema.StringAttribute)
	if !ok {
		t.Fatalf("routeros_system_scheduler.start_time missing or not a StringAttribute")
	}
	if len(att.Validators) == 0 {
		t.Fatal("start_time has no validator: garbage now reaches the device and fails mid-apply")
	}
	check := func(v string) bool {
		req := validator.StringRequest{Path: path.Root("start_time"), ConfigValue: types.StringValue(v)}
		for _, val := range att.Validators {
			resp := &validator.StringResponse{}
			val.ValidateString(context.Background(), req, resp)
			if resp.Diagnostics.HasError() {
				return false
			}
		}
		return true
	}
	// Accepted by ROS 7.23.2; "24:00:00" and "1d00:00:00" are rejected here
	// because the device silently rewrites both to "00:00:00".
	for _, ok := range []string{"23:57:05", "startup", "00:00:00", "0:0:0", "23:57"} {
		if !check(ok) {
			t.Errorf("start_time rejected %q, want accepted", ok)
		}
	}
	for _, bad := range []string{"nonsense", "24:00:00", "1d00:00:00", "25:00:00", "12:60:00"} {
		if check(bad) {
			t.Errorf("start_time accepted %q, want rejected", bad)
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

// ValidateConfig only sees config, so an interpolated type or address is still
// Unknown there. The resolved plan has to be re-checked or the type rewrite is
// reachable through any computed reference.
func TestIPDNSStaticResolvedCheckCatchesWhatValidateConfigCannot(t *testing.T) {
	srvWithAddress := IPDNSStaticModel{
		Name:    types.StringValue("_sip._tcp.example.lan"),
		Type:    types.StringValue("SRV"),
		Address: types.StringValue("10.0.0.1"),
	}
	if summary, _ := ipDNSStaticCheckResolved(srvWithAddress); summary == "" {
		t.Error("a resolved SRV record carrying an address was accepted")
	}

	// At plan time the same config has address Unknown, which is exactly why
	// ValidateConfig cannot catch it and Create/Update must.
	atPlanTime := srvWithAddress
	atPlanTime.Address = types.StringUnknown()
	if ipDNSStaticHasAddress(atPlanTime) {
		t.Error("an unknown address counted as set")
	}

	ok := IPDNSStaticModel{Name: types.StringValue("a.example.lan"), Type: types.StringValue("A"), Address: types.StringValue("10.0.0.1")}
	if summary, _ := ipDNSStaticCheckResolved(ok); summary != "" {
		t.Errorf("a valid A record was rejected: %s", summary)
	}
}

// Requires either name OR regexp -- what the docs always claimed and the schema
// did not allow. RouterOS accepts a regexp entry with no name.
func TestIPDNSStaticNameOrRegexp(t *testing.T) {
	for _, tc := range []struct {
		name, nameVal, reVal string
		wantErr              bool
	}{
		{"name only", "a.example.lan", "", false},
		{"regexp only", "", `.*\.example\.lan`, false},
		{"neither", "", "", true},
		{"both", "a.example.lan", `.*\.example\.lan`, true},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := IPDNSStaticModel{Name: types.StringNull(), Regexp: types.StringNull()}
			if tc.nameVal != "" {
				m.Name = types.StringValue(tc.nameVal)
			}
			if tc.reVal != "" {
				m.Regexp = types.StringValue(tc.reVal)
			}
			summary, _ := ipDNSStaticNameOrRegexp(m)
			if (summary != "") != tc.wantErr {
				t.Errorf("got %q, wantErr=%v", summary, tc.wantErr)
			}
		})
	}
}

// The device rejects a lower-case type outright, so nothing may rely on
// case-folding it.
func TestIPDNSStaticTypeIsCaseSensitive(t *testing.T) {
	attrs := schemaOf(t, NewIPDNSStaticResource()).Schema.Attributes
	att, ok := attrs["type"].(schema.StringAttribute)
	if !ok || len(att.Validators) == 0 {
		t.Fatal("type has no validator")
	}
	check := func(v string) bool {
		req := validator.StringRequest{Path: path.Root("type"), ConfigValue: types.StringValue(v)}
		for _, val := range att.Validators {
			resp := &validator.StringResponse{}
			val.ValidateString(context.Background(), req, resp)
			if resp.Diagnostics.HasError() {
				return false
			}
		}
		return true
	}
	for _, good := range []string{"A", "AAAA", "CNAME", "FWD", "MX", "NS", "NXDOMAIN", "SRV", "TXT"} {
		if !check(good) {
			t.Errorf("type rejected %q", good)
		}
	}
	for _, bad := range []string{"a", "srv", "bogus"} {
		if check(bad) {
			t.Errorf("type accepted %q; the device answers 400 for it", bad)
		}
	}
}
