package provider

import (
	"context"
	"testing"

	"github.com/hashicorp/terraform-plugin-framework/resource"

	"github.com/ebogdum/terraform-provider-routeros/internal/client"
)

// resourceAttrs returns the attribute names a resource's schema declares.
func resourceAttrs(t *testing.T, r resource.Resource) map[string]struct{} {
	t.Helper()
	resp := &resource.SchemaResponse{}
	r.Schema(context.Background(), resource.SchemaRequest{}, resp)
	if resp.Diagnostics.HasError() {
		t.Fatalf("schema diagnostics: %v", resp.Diagnostics)
	}
	out := make(map[string]struct{}, len(resp.Schema.Attributes))
	for name := range resp.Schema.Attributes {
		out[name] = struct{}{}
	}
	return out
}

func requireAttrs(t *testing.T, r resource.Resource, label string, want ...string) {
	t.Helper()
	have := resourceAttrs(t, r)
	for _, w := range want {
		if _, ok := have[w]; !ok {
			t.Errorf("%s is missing attribute %q", label, w)
		}
	}
}

// A CAP that exposes only `enabled` cannot be rebuilt from Terraform: without
// the controller address the radios come up unmanaged and no SSID appears.
func TestWifiCapSchemaCoversCapsManBinding(t *testing.T) {
	requireAttrs(t, NewInterfaceWifiCapResource(), "routeros_interface_wifi_cap",
		"enabled", "caps_man_addresses", "caps_man_names",
		"caps_man_certificate_common_names", "discovery_interfaces",
		"certificate", "lock_to_caps_man", "slaves_static")
}

func TestWifiCapRoundTripsCapsManAddresses(t *testing.T) {
	var m InterfaceWifiCapModel
	interfaceWifiCapApply(context.Background(), client.Object{
		"enabled": "true", "caps-man-addresses": "192.168.10.1", "slaves-static": "false",
	}, &m)
	if got := m.CapsManAddresses.ValueString(); got != "192.168.10.1" {
		t.Errorf("caps_man_addresses = %q, want 192.168.10.1", got)
	}
	if !m.Enabled.ValueBool() {
		t.Error("enabled did not round-trip")
	}
	if m.SlavesStatic.IsNull() || m.SlavesStatic.ValueBool() {
		t.Errorf("slaves_static = %v, want false", m.SlavesStatic)
	}
	// An absent optional must be null rather than empty string.
	if !m.Certificate.IsNull() {
		t.Errorf("absent certificate = %q, want null", m.Certificate.ValueString())
	}
}

// Without hop_limit the defconf rfc4890 rule ("drop ICMPv6 with hop-limit 1")
// is emitted as "drop all ICMPv6" -- a silent widening of a firewall rule.
func TestIPv6FirewallHopLimit(t *testing.T) {
	requireAttrs(t, NewIPV6FirewallFilterResource(), "routeros_ipv6_firewall_filter", "hop_limit")

	var m IPV6FirewallFilterModel
	iPV6FirewallFilterApply(context.Background(), client.Object{
		".id": "*1", "chain": "forward", "action": "drop",
		"protocol": "icmpv6", "hop-limit": "equal:1",
	}, &m)
	if got := m.HopLimit.ValueString(); got != "equal:1" {
		t.Fatalf("hop_limit = %q, want equal:1", got)
	}

	var absent IPV6FirewallFilterModel
	iPV6FirewallFilterApply(context.Background(), client.Object{".id": "*2", "chain": "forward"}, &absent)
	if !absent.HopLimit.IsNull() {
		t.Errorf("absent hop-limit = %q, want null", absent.HopLimit.ValueString())
	}
}

// auto-negotiation is the settable yes/no toggle the device exports; the
// enum of link states (done/incomplete/failed) comes from `monitor`, not the
// menu. Validating the attribute against that enum made it unusable, and the
// value was never written to the wire at all.
func TestEthernetAutoNegotiation(t *testing.T) {
	requireAttrs(t, NewInterfaceEthernetResource(), "routeros_interface_ethernet",
		"auto_negotiation", "bandwidth")

	var m InterfaceEthernetModel
	interfaceEthernetApply(context.Background(), client.Object{
		".id": "*1", "name": "ether1",
		"auto-negotiation": "true", "bandwidth": "unlimited/unlimited",
	}, &m)
	if m.AutoNegotiation.IsNull() || !m.AutoNegotiation.ValueBool() {
		t.Errorf("auto_negotiation = %v, want true", m.AutoNegotiation)
	}
	if got := m.Bandwidth.ValueString(); got != "unlimited/unlimited" {
		t.Errorf("bandwidth = %q, want unlimited/unlimited", got)
	}
	// A device answering with a link state must not be coerced to false.
	interfaceEthernetApply(context.Background(), client.Object{".id": "*1", "auto-negotiation": "done"}, &m)
	if !m.AutoNegotiation.IsNull() {
		t.Errorf("link-state answer produced %v, want null", m.AutoNegotiation)
	}
}

func TestVlanLoopProtect(t *testing.T) {
	requireAttrs(t, NewInterfaceVLANResource(), "routeros_interface_vlan",
		"loop_protect", "loop_protect_disable_time", "loop_protect_send_interval")

	var m InterfaceVLANModel
	interfaceVLANApply(context.Background(), client.Object{
		".id": "*1", "name": "ether4-iot", "vlan-id": "120",
		"loop-protect": "default", "loop-protect-disable-time": "5m",
		"loop-protect-send-interval": "5s",
	}, &m)
	if got := m.LoopProtect.ValueString(); got != "default" {
		t.Errorf("loop_protect = %q, want default", got)
	}
	if got := m.LoopProtectDisableTime.ValueString(); got != "5m" {
		t.Errorf("loop_protect_disable_time = %q, want 5m", got)
	}
	if got := m.LoopProtectSendInterval.ValueString(); got != "5s" {
		t.Errorf("loop_protect_send_interval = %q, want 5s", got)
	}
}

// The button bindings hold real configuration (mode-button -> dark-mode,
// wps-button -> wps-accept) yet had no resource at all, so a device rebuilt
// from Terraform came up with its buttons unbound even though the scripts they
// invoke were managed.
func TestRouterboardButtons(t *testing.T) {
	cases := []struct {
		newFn    func() resource.Resource
		typeName string
	}{
		{NewSystemRouterboardModeButtonResource, "routeros_system_routerboard_mode_button"},
		{NewSystemRouterboardWPSButtonResource, "routeros_system_routerboard_wps_button"},
		{NewSystemRouterboardResetButtonResource, "routeros_system_routerboard_reset_button"},
	}
	registered := map[string]struct{}{}
	for _, f := range registryResources() {
		resp := &resource.MetadataResponse{}
		f().Metadata(context.Background(), resource.MetadataRequest{ProviderTypeName: "routeros"}, resp)
		registered[resp.TypeName] = struct{}{}
	}
	for _, tc := range cases {
		r := tc.newFn()
		resp := &resource.MetadataResponse{}
		r.Metadata(context.Background(), resource.MetadataRequest{ProviderTypeName: "routeros"}, resp)
		if resp.TypeName != tc.typeName {
			t.Errorf("type name = %q, want %q", resp.TypeName, tc.typeName)
		}
		if _, ok := registered[tc.typeName]; !ok {
			t.Errorf("%s is not registered with the provider", tc.typeName)
		}
		requireAttrs(t, r, tc.typeName, "enabled", "hold_time", "on_event", "router")
	}
}

func TestRouterboardButtonApply(t *testing.T) {
	var m SystemRouterboardButtonModel
	systemRouterboardButtonApply(client.Object{
		"enabled": "true", "hold-time": "0s..1m", "on-event": "dark-mode",
	}, &m)
	if !m.Enabled.ValueBool() {
		t.Error("enabled did not round-trip")
	}
	// hold-time is a range, not a duration; it must survive verbatim.
	if got := m.HoldTime.ValueString(); got != "0s..1m" {
		t.Errorf("hold_time = %q, want 0s..1m", got)
	}
	if got := m.OnEvent.ValueString(); got != "dark-mode" {
		t.Errorf("on_event = %q, want dark-mode", got)
	}

	// The reset button ships enabled=no with an empty binding: "" is a real
	// value here and must not collapse to null.
	var unbound SystemRouterboardButtonModel
	systemRouterboardButtonApply(client.Object{
		"enabled": "false", "hold-time": "0s..1m", "on-event": "",
	}, &unbound)
	if unbound.OnEvent.IsNull() {
		t.Error("empty on-event became null; it is a meaningful value")
	}
	if unbound.OnEvent.ValueString() != "" {
		t.Errorf("on_event = %q, want empty", unbound.OnEvent.ValueString())
	}
}

// Several menus carry a WebFig-only spelling that shadows a real property:
// autoneg/auto-negotiation, def/default, resp/responder, notemplate/template.
// RouterOS rejects the shadow over /rest, so any config that set one failed to
// apply. They are deprecated and must never reach a request body.
func TestShadowPropertiesAreNotWritten(t *testing.T) {
	shadowed := map[string][]string{
		"routeros_interface_ethernet":   {"autoneg", "noautoneg"},
		"routeros_ip_hotspot_user":      {"def", "nondef"},
		"routeros_ip_ipsec_mode_config": {"resp", "nonresp"},
		"routeros_ip_ipsec_policy":      {"notemplate"},
		"routeros_ppp_profile":          {"def"},
	}
	byName := map[string]resource.Resource{}
	for _, f := range registryResources() {
		r := f()
		resp := &resource.MetadataResponse{}
		r.Metadata(context.Background(), resource.MetadataRequest{ProviderTypeName: "routeros"}, resp)
		byName[resp.TypeName] = r
	}
	for typeName, attrs := range shadowed {
		r, ok := byName[typeName]
		if !ok {
			t.Fatalf("%s not registered", typeName)
		}
		sResp := &resource.SchemaResponse{}
		r.Schema(context.Background(), resource.SchemaRequest{}, sResp)
		for _, a := range attrs {
			att, ok := sResp.Schema.Attributes[a]
			if !ok {
				continue // dropping the attribute outright is also acceptable
			}
			if att.GetDeprecationMessage() == "" {
				t.Errorf("%s.%s shadows a real property but is not deprecated", typeName, a)
			}
		}
	}
}

// The real siblings must still be settable -- deprecating the shadow must not
// have taken the working property with it.
func TestRealSiblingsStillSettable(t *testing.T) {
	requireAttrs(t, NewIPIpsecModeConfigResource(), "routeros_ip_ipsec_mode_config", "responder")
	requireAttrs(t, NewIPIpsecPolicyResource(), "routeros_ip_ipsec_policy", "template")
}

// The provider's REST field names were derived by splitting Go identifiers, so
// PoEOut became "po-e-out" and DonTRequirePermissions became
// "don-t-require-permissions". RouterOS has neither: it reports poe-out and
// dont-require-permissions, so every write to these silently addressed a
// property that does not exist.
func TestPoEWireNamesMatchDevice(t *testing.T) {
	requireAttrs(t, NewInterfaceEthernetResource(), "routeros_interface_ethernet",
		"poe_out", "poe_priority", "poe_out_power")

	var m InterfaceEthernetModel
	interfaceEthernetApply(context.Background(), client.Object{
		".id": "*1", "name": "ether1",
		"poe-out": "auto-on", "poe-priority": "10", "poe-out-power": "12W",
	}, &m)
	if got := m.PoEOut.ValueString(); got != "auto-on" {
		t.Errorf("poe_out = %q, want auto-on", got)
	}
	if got := m.PoEPriority.ValueInt64(); got != 10 {
		t.Errorf("poe_priority = %d, want 10", got)
	}
	// The old mangled spelling must no longer be understood.
	var stale InterfaceEthernetModel
	interfaceEthernetApply(context.Background(), client.Object{".id": "*1", "po-e-out": "auto-on"}, &stale)
	if !stale.PoEOut.IsNull() {
		t.Errorf("the mangled key po-e-out still populates poe_out (%q)", stale.PoEOut.ValueString())
	}
}

func TestScriptPermissionsWireName(t *testing.T) {
	requireAttrs(t, NewSystemScriptResource(), "routeros_system_script", "dont_require_permissions")

	var m SystemScriptModel
	systemScriptApply(context.Background(), client.Object{
		".id": "*1", "name": "dark-mode", "dont-require-permissions": "true",
	}, &m)
	if !m.DonTRequirePermissions.ValueBool() {
		t.Error("dont-require-permissions did not round-trip")
	}
}

// RouterOS exposes /interface/wifi sub-objects under dotted REST keys
// (configuration.ssid, security.passphrase, ...). The resource declared those
// members as flat attributes and read/wrote them as flat keys, which RouterOS
// does not have -- so none of the wifi configuration round-tripped. The
// attribute names are unchanged; only the wire mapping is corrected.
func TestWifiDottedSubObjectFields(t *testing.T) {
	var m InterfaceWifiModel
	interfaceWifiApply(context.Background(), client.Object{
		".id": "*3", "name": "wifi-mb2", "disabled": "false",
		"configuration.manager":         "capsman",
		"configuration.ssid":            "mb2",
		"configuration.country":         "Latvia",
		"security.passphrase":           "hunter2hunter2",
		"security.authentication-types": "wpa2-psk,wpa3-psk",
		"datapath.vlan-id":              "120",
		"channel.band":                  "5ghz-ax",
	}, &m)

	for _, tc := range []struct{ name, got, want string }{
		{"manager", m.Manager.ValueString(), "capsman"},
		{"ssid", m.Ssid.ValueString(), "mb2"},
		{"country", m.Country.ValueString(), "Latvia"},
		{"passphrase", m.Passphrase.ValueString(), "hunter2hunter2"},
		{"authentication_types", m.AuthenticationTypes.ValueString(), "wpa2-psk,wpa3-psk"},
		{"band", m.Band.ValueString(), "5ghz-ax"},
	} {
		if tc.got != tc.want {
			t.Errorf("%s = %q, want %q", tc.name, tc.got, tc.want)
		}
	}

	// The flat spellings RouterOS does not have must no longer populate anything.
	var flat InterfaceWifiModel
	interfaceWifiApply(context.Background(), client.Object{
		".id": "*3", "manager": "capsman", "passphrase": "nope", "ssid": "nope",
	}, &flat)
	if !flat.Manager.IsNull() || !flat.Passphrase.IsNull() || !flat.Ssid.IsNull() {
		t.Error("flat sub-object keys still populate the model; the dotted mapping is not exclusive")
	}
}

// The section names are themselves real top-level properties (a reference to a
// named profile), and must not have been swallowed by the dotted rewrite.
func TestWifiProfileReferencesSurvive(t *testing.T) {
	var m InterfaceWifiModel
	interfaceWifiApply(context.Background(), client.Object{
		".id": "*3", "name": "cap-wifi1", "configuration": "cfg-mb5", "security": "sec-home",
	}, &m)
	if got := m.Configuration.ValueString(); got != "cfg-mb5" {
		t.Errorf("configuration = %q, want cfg-mb5", got)
	}
	if got := m.Security.ValueString(); got != "sec-home" {
		t.Errorf("security = %q, want sec-home", got)
	}
}

func TestWifiSecretsAreSensitive(t *testing.T) {
	resp := &resource.SchemaResponse{}
	NewInterfaceWifiResource().Schema(context.Background(), resource.SchemaRequest{}, resp)
	for _, a := range []string{"passphrase", "eap_password"} {
		att, ok := resp.Schema.Attributes[a]
		if !ok {
			t.Fatalf("%s missing", a)
		}
		if !att.IsSensitive() {
			t.Errorf("%s is not marked Sensitive", a)
		}
	}
}

// /interface/wifi/configuration is the mirror case: its own members (ssid,
// mode, country, manager) are genuinely top-level there, while channel,
// datapath, security, aaa and steering remain sub-objects. Mapping the wrong
// half would break a menu that currently works.
func TestWifiConfigurationDottedSplit(t *testing.T) {
	var m InterfaceWifiConfigurationModel
	interfaceWifiConfigurationApply(context.Background(), client.Object{
		".id": "*1", "name": "cfg-mb2",
		"ssid": "mb2", "mode": "ap", "country": "Latvia", // top-level here
		"security.passphrase": "hunter2hunter2",
		"channel.band":        "5ghz-ax",
		"datapath.vlan-id":    "120",
	}, &m)
	if got := m.Ssid.ValueString(); got != "mb2" {
		t.Errorf("ssid = %q, want mb2 (top-level in this menu)", got)
	}
	if got := m.Mode.ValueString(); got != "ap" {
		t.Errorf("mode = %q, want ap (top-level in this menu)", got)
	}
	if got := m.Passphrase.ValueString(); got != "hunter2hunter2" {
		t.Errorf("passphrase = %q, want the dotted security.passphrase value", got)
	}
	if got := m.Band.ValueString(); got != "5ghz-ax" {
		t.Errorf("band = %q, want the dotted channel.band value", got)
	}
}
