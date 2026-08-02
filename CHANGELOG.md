# Changelog

All notable changes to this project are documented here.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [3.0.3] - 2026-08-02

### Fixed

Three schema defects confirmed against a live gateway:

- **`routeros_ip_route.gateway` accepts an interface name.** RouterOS allows a
  gateway of an IP, an interface name (e.g. a WireGuard peer `wgv-mac`), or
  `IP%interface`; the attribute was validated as IP-only, so a route via an
  interface was rejected at plan time. The IP-only validator is removed.
- **`routeros_interface_dot1x_server.radius_mac_format` accepts upper-case.**
  RouterOS stores `XX:XX:XX:XX:XX:XX`, but the enum only listed the lower-case
  forms. The upper-case `XX:XX:…`, `XX-XX-…` and `XXXXXXXXXXXX` are now valid.
- **`routeros_system_scheduler.on_event` is optional.** RouterOS runs a
  scheduler with an empty `on-event` (e.g. driven purely by interval), so the
  attribute is no longer Required.

## [3.0.2] - 2026-08-02

### Fixed

- **Presence-flag booleans read back correctly.** RouterOS returns a set
  boolean flag as a valueless key (`no-summaries: ""`); `ParseBool("")` failed
  and the attribute was read back as `null`, tainting the resource so every plan
  tried to replace it. `routing_ospf_area.no_summaries` was one such field. The
  whole class is now handled: every boolean Apply (652) treats a present but
  empty value as `true`, the same shape as the `fib`/`blackhole` fixes — this
  time applied across all boolean attributes rather than one at a time.

## [3.0.1] - 2026-08-01

### Fixed

A second-device functional lab (RouterOS 7.22.3) confirmed that v3.0.0 fixed
individual fields but not the whole class behind them. This release fixes the
classes:

- **Set-ordering.** RouterOS stores several attributes as sets and returns them
  in its own order, so a positional list comparison diffed forever. Seven list
  attributes are now `Set` (`ip_ipsec_profile` `dh_group`/`enc_algorithm`/
  `enc_algorithms`, `ip_hotspot_profile` `login_by`, `ip_dhcp_client`
  `dhcp_options`, `interface_ethernet` `advertise`, `certificate_builtin`
  `key_usage`), and 13 comma-string set attributes (`authentication` on the
  ppp/l2tp/pptp/sstp servers, `time` on every ip/ipv6 firewall chain + calea)
  compare order-insensitively.
- **Default-omission.** RouterOS omits an attribute left at its default, which
  the provider read back as `null` and reported as a perpetual diff. Every
  boolean Apply (488) and every standard string Apply (2699) now keeps the
  configured value when the device omits the key instead of nulling it. This
  covers `routing_ospf_interface_template.passive`, `ip_dhcp_server`
  `authoritative`, `certificate.trusted`, `ip_hotspot_profile.rate_limit`,
  `ip_dns_static.address`, and the rest of the class.

Not addressed (each has a real obstacle, none is a round-trip defect):
`/interface/ovpn-server/server` (the menu exposes no fields to model against),
certificate signing (a `/certificate/sign` action), and `/routing/rip`
(absent on RouterOS 7.2x firmware).

## [3.0.0] - 2026-08-01

### Fixed

Defects found by an exhaustive conformance campaign against a live RouterOS
7.23.2 hAP: a non-mutating audit of every writable attribute (4277 keys across
308 menus) plus create/apply/destroy of ~200 resources — all 63 applicable
singletons and ~135 collections spanning every value type, including the
privileged firewall/address/route/user menus and multi-resource IPsec, OSPF
and IS-IS chains. Every fix below was
reproduced against the device and re-verified:

- **Fresh creates no longer fail with "invalid result object after apply."** A
  Computed attribute the device never returns (e.g. `/ip/cloud dns_name`) was
  left at the plan's Unknown on create; the result is now swept so every
  unresolved attribute is null. Affects every resource.
- **`/tool/mac-server`** applied with `allowed-interface-list` unset no longer
  trips the lockout guard (the field is only guarded when actually written).
- **`routeros_routing_table`** `fib`, **`routeros_ip_route`** and
  **`routeros_ipv6_route`** `blackhole` round-trip: RouterOS returns these as
  valueless presence flags, so a row created with the flag set no longer reads
  back as null.
- **`routeros_ipv6_firewall_address_list`** accepts a bare host address: a
  config `2001:db8::1` and the `2001:db8::1/128` RouterOS stores now compare
  equal (semantic equality), ending the spurious diff / inconsistent result.
- **`routeros_queue_simple` / `routeros_queue_tree`** rate fields round-trip
  regardless of `k`/`M`/`G` suffix: `10M/10M` and `10000000/10000000` are
  equal.
- **`routeros_user_group` / `routeros_system_script`** `policy` round-trips.
  RouterOS returns a group's policy as the full permission set with everything
  not granted negated (`read,winbox,!ftp,!telnet,...`) in its own order, so the
  list a user wrote never matched what came back. `policy` is now an
  order-insensitive set of the granted permissions.
- **`routeros_ip_socks_users`** `rate_limit` (k/M/G suffix) and
  **`routeros_tool_graphing_resource`** `allow_address` (bare host `/32`) now
  round-trip.
- **`routeros_ip_socks_users` `rate_limit`** now round-trips a single rate: a
  single value is symmetric (`10M` = `10M/10M`), so the rate semantic type
  normalises a bare rate to its rx/tx pair. **`routeros_ip_socksify` `comment`**
  removed (the /ip/socks menu has no comment field; it was silently dropped).
- **`routeros_radius` `timeout`** was typed as an integer but RouterOS stores it
  as a duration (`1s`), so any value read back null. Retyped to a duration string.
- **`routeros_interface` `type`, `routeros_interface_bridge_port` `role`/`status`**
  were typed as integers but hold strings (`ether`, `designated-port`,
  `in-bridge`), so they always read back null. Retyped to (read-only) strings.
- **Full bool-string class swept from live-device data** — every string-typed
  attribute RouterOS reports as `true`/`false` now uses the semantic type and
  keeps its configured value when the device omits the default: `bridge.dhcpv6_snooping`,
  `bridge_port.hw`/`trusted_dhcpv6`, `interface_vlan.use_service_tag`/`mvrp`,
  `ip_neighbor_discovery_settings.add_dns_entries`/`lldp_med`,
  `system_routerboard_settings.auto_upgrade`/`silent_boot`/`force_backup_booter`, and more.
- **Boolean-valued string attributes now round-trip regardless of `yes`/`no`
  vs `true`/`false`.** RouterOS accepts `yes`/`no` on input but reads these
  fields back as `true`/`false`; a config of `no` never matched. A semantic
  type now compares them as booleans, fixing e.g. `interface_gre6`/`ipipv6`
  `clamp_tcp_mss`, `interface_ipip` `allow_fast_path`, `ip_socks_users`
  `only_one`, `mpls_ldp_interface` `accept_dynamic_neighbors`,
  `tool_graphing_resource` `store_on_disk`.
- **`routeros_tool_bandwidth_server`** `allowed_addresses4` / `allowed_addresses6`
  accept a bare host (stored `/32` / `/128`), and **`routeros_tool_traffic_generator`**
  / **`routeros_tool_sniffer`** `running` is read-only runtime state, no longer
  settable.
- **`routeros_ip_address`** `network` is read-only (derived from `address`);
  it was settable, so a config value conflicted with the computed one.
- **`routeros_ppp_secret`** `limit_bytes_in` / `limit_bytes_out` round-trip
  regardless of `k`/`M`/`G` suffix, and `remote_ipv6_prefix` accepts a bare
  IPv6 host (stored as `/128`).
- **`routeros_ppp_l2tp_secret`** `address` accepts a bare IPv4 host: the device
  stores it as `/32`, so `192.0.2.10` and `192.0.2.10/32` now compare equal
  (the host-address semantic type now covers IPv4 `/32` as well as IPv6 `/128`).
- **`routeros_ip_hotspot_ip_binding`** `address` is validated as a single host
  IP, not a CIDR, matching what RouterOS stores.
- **`routeros_interface_ethernet_switch_port`** applied cleanly again: it wrote
  its `name`/`switch` identifiers into the Set body, which the device rejects as
  unknown parameters, so every apply failed. The identifiers now only select
  the row to adopt.
- **Fixed-row menus that reject Add now adopt an existing row.**
  `routeros_interface_ethernet_switch_port_isolation`,
  `routeros_system_resource_irq_rps` and `routeros_system_resource_irq` had a
  Computed-only key, so Create always failed against the device's fixed rows.
  Their key (`name`, or `irq`) is now settable and the matching row is adopted
  and configured; Delete drops it from state.

### Removed

- **`routeros_tool_graphing_interface` / `_queue` / `_resource`** `comment` —
  the graphing menus silently accept but never store a comment, so it never
  round-tripped. Removed.
- **`routeros_interface_ethernet_poe`** `export`, `monitor`, `duration` — these
  are RouterOS CLI actions, not settable attributes; the device rejected them as
  unknown parameters. Removed.
- **Phantom attributes that mapped to non-existent RouterOS keys.** A generator
  quirk split acronyms/compounds with a stray dash, producing duplicate
  attributes whose device key the router rejects as an "unknown parameter".
  Each was already non-functional (always read null; the writable ones failed
  apply if set). Removed, keeping the correct sibling attribute:
  `l2_mtu` (→ `l2mtu`) on `interface`, `interface_ethernet`,
  `interface_wireguard`; `do_h_servers` (→ `doh_servers`) and
  `verify_do_h_certificate` (→ `verify_doh_cert`) on `ip_dns_forwarders`;
  `qo_s_classifier` (→ `qos_classifier`) and `ft_preserve_vlan_id`
  (→ `ft_preserve_vlanid`) on the `interface_wifi*` resources.
- **`routeros_system_backup`, `routeros_tool_sniffer_stop`,
  `routeros_tool_profile_cpu`** (resources) — these map to RouterOS action
  endpoints (`/system/backup`, `/tool/sniffer/stop`, `/tool/profile/cpu`), not
  manageable state; the device returns 400/500 on any read or add, so they could
  never apply. Removed (the matching action data sources were already dropped).
- **`routeros_system_resource_hardware`** (resource) — the device exposes
  `/system/resource/hardware` as read-only inventory that accepts neither Add
  nor Set, so the resource could never apply. Use the
  `routeros_system_resource_hardware` **data source**, which is unchanged.

## [2.0.2] - 2026-07-31

### Changed

**Breaking.** This release corrects attributes that addressed properties
RouterOS does not have, or that could not hold the values it reports. Every
change below affects a configuration that could not have worked before, but
some require an edit. See the
[v2 upgrade guide](guides/upgrading-to-v2) for the full list and the
mechanical fixes.

- Renamed attributes whose REST field names were mangled:
  `po_e_out` → `poe_out`, `po_e_priority` → `poe_priority`,
  `po_e_voltage` → `poe_voltage`, `po_e_out_*` → `poe_out_*`, and
  `don_t_require_permissions` → `dont_require_permissions`.
- Retyped attributes that hold a word as well as a number:
  `interface_bridge.mtu`, `interface_bridge_port.horizon` and
  `ipv6_nd`'s `hop_limit`, `mtu`, `reachable_time` and
  `retransmit_interval` are now strings. HCL converts a bare number
  automatically, so `mtu = 1500` keeps working.
- `interface_ethernet.auto_negotiation` changed from a string holding the
  read-only link state to a `bool` holding the settable toggle.
- 561 attributes across 108 resources that hold read-only runtime state
  are no longer `Optional`. Setting one used to be discarded silently and
  is now a configuration error.
- `autoneg`, `noautoneg`, `def`, `nondef`, `resp`, `nonresp` and
  `notemplate` are deprecated and no longer sent to the device.

### Added

- `routeros_ip_service` resource: enable, disable and reconfigure the
  RouterOS management services (`api`, `api-ssl`, `ftp`, `ssh`, `telnet`,
  `winbox`, `www`, `www-ssl`). Rows are adopted by `name` — RouterOS does
  not allow adding or removing them — and `terraform destroy` deliberately
  leaves a disabled service disabled. The existing lockout guard is wired
  in: a change that would disable every management service is refused
  unless `lockout_ack = true`.
- `routeros_ip_dns_adlist` resource: manage DNS blocklists in
  `/ip/dns/adlist` (RouterOS 7.15+), sourced from a `url` or an on-router
  `file`. Setting both or neither is rejected at plan time.
- `routeros_interface_wifi_cap` gained the CAPsMAN binding properties:
  `caps_man_addresses`, `caps_man_names`,
  `caps_man_certificate_common_names`, `discovery_interfaces`,
  `certificate`, `lock_to_caps_man` and `slaves_static`. Previously the
  resource exposed only `enabled`, so a CAP rebuilt from Terraform came up
  with no controller address and never adopted its provisioned SSIDs.
- `routeros_ipv6_firewall_filter` gained `hop_limit`. Without it the
  default rfc4890 rule could only be expressed as "drop all ICMPv6"
  instead of "drop ICMPv6 with hop-limit 1" — a silent widening of a
  firewall rule.
- `routeros_interface_ethernet` gained `bandwidth`, and
  `routeros_interface_vlan` gained `loop_protect`,
  `loop_protect_disable_time` and `loop_protect_send_interval`.
- Resources for the 23 menus that previously had none:
  `routeros_interface_ethernet_switch_port`,
  `routeros_interface_l2tp_server_server`,
  `routeros_interface_lte_settings`, `routeros_interface_macsec_profile`,
  `routeros_interface_pptp_server_server`,
  `routeros_interface_sstp_server_server`, `routeros_ip_cloud_advanced`,
  `routeros_ip_firewall_service_port`,
  `routeros_ip_hotspot_service_port`, `routeros_ip_ipsec_key_qkd`,
  `routeros_ip_ipsec_policy_group`, `routeros_ip_media_settings`,
  `routeros_ip_neighbor_discovery_settings`,
  `routeros_ip_traffic_flow_ipfix`, `routeros_ipv6_dhcp_relay_option`,
  `routeros_ipv6_nd_prefix_default`, `routeros_queue_interface`,
  `routeros_system_health_settings`,
  `routeros_system_package_local_update_mirror`,
  `routeros_system_resource_hardware_usb_settings`,
  `routeros_system_resource_irq`, `routeros_system_resource_irq_rps` and
  `routeros_tool_mac_server_mac_winbox`.
- `ip_dhcp_client_option.code` accepts a word (`hostname`) as well as a
  number and is now a string. Its sibling `ip_dhcp_server_option.code` does
  not, and was correctly left numeric — confirmed against the device.
- Eight more integer-typed attributes also accept a word on the device and
  could not hold their own valid values: `mtu` on `/interface`,
  `/interface/eoip`, `/interface/gre`, `/interface/ipip` and
  `/interface/wireguard` (`auto`),
  `interface_ethernet_switch_port.default_vlan_id` (`auto`),
  `ip_dhcp_server.client_mac_limit` and
  `ip_hotspot_user_profile.shared_users` (`unlimited`). All are strings now.
- Verified every wire key the provider writes against a live RouterOS 7.x
  device. 335 keys were rejected outright, so any configuration touching one
  failed the whole apply. Two shapes came out of it:
  `/routing/bgp/connection`, `/routing/bgp/template` and `/routing/bgp/vpn`
  use dotted sub-objects like the wifi menus (`input.filter`,
  `input.accept-communities`, `export.route-targets`) — 33 keys remapped;
  and 302 keys were console commands or read-only state modelled as
  properties (`/ip/arp` `ping`/`torch`/`make-static`, `/certificate`
  `sign`/`import`/`revoke`, `/ip/dhcp-server/lease` `check-status`,
  `/system/script` `run-script`, `/user/group` `policies`). Those are no
  longer sent and their attributes are `Computed`-only.
- Security: 17 secret-bearing attributes were not flagged sensitive and
  leaked in cleartext into plan output, state and CLI display — including
  `certificate.private_key`, `ip_cloud.vpn_private_key` /
  `vpn_peer_private_key`, the `ipsec_secret` on EoIPv6/GRE6/IPIPv6/L2TP
  tunnels, `eap_password` on the wifi security/configuration menus,
  `interface_ppp_client`/`interface_vrrp` `password`, the disk
  self-encryption and NVMe-TCP secrets, and User Manager
  paypal/web passwords. All are now marked sensitive, enforced by a
  `TestSecretsAreSensitive` regression test.
- Removed dead code: the never-called `CheckSSHKeyImportLockout` placeholder
  guard and 233 generator-emitted `_ = pkg.Sym` import-keeper lines across
  the resource files.
- Definitively verified every provider-written key against a live hAP ax³
  running a version-matched RouterOS 7.23.2: 217 keys are rejected, all
  explained (139 are 7.24 wifi features present only in newer firmware; 78
  need hardware the device lacks) and zero are defects. Fixed the 36 that
  were real: read-only state and console commands modelled as writable
  properties (`/certificate` `dynamic`/`private-key`/`type`, `/partition`
  `active`/`running`/`activate`, `/user type`, dhcp-server/lease/network
  read-only fields, traffic-generator template state, …) are no longer
  written; the mangled `interface_wifi_interworking` key `hotspot-2-0` is
  corrected to `hotspot20`; and `routing_bgp_connection` `remote-port`
  becomes the dotted `remote.port` the device accepts. A `TestNoPhantomWrites`
  regression test pins every one.
- Verified the provider against the machine-generated firmware CLI reference
  and a live 7.22.3 hAP ax³: built an authoritative spec of 545 menus and
  5072 device-confirmed settable parameters. Every cleanly-expressible
  settable parameter is now covered (the only unmapped ones have a literal
  `/` in the RouterOS name, which cannot be a Terraform identifier). Fixed a
  mangled wire key on `interface_wifi_interworking` (`hotspot-2-0` →
  `hotspot20`, confirmed on the device). Remaining rejected keys on 7.22.3
  are 7.24-only wifi grammar or hardware-gated (PoE/switch-chip) fields, not
  defects.
- Documentation is now a verified projection of the schema: all 389 resource
  docs regenerated so every attribute is documented with its real type,
  required/optional/read-only status, sensitivity and description — no
  missing, invented, or mistyped entries. A `TestDocsMatchSchema` build-time
  test keeps them in sync.
- Cross-checked every written key against the firmware's own CLI reference
  (unpacked from the 7.23 NPK) and the live device, which surfaced six keys
  the provider wrote but RouterOS rejects: read-only state on
  `interface_list_member` (`dynamic`), `interface_macsec_profile`
  (`default-name`) and `ip_cloud` (`dns-name`), now Computed-only; and the
  identity field on three fixed-row menus — `ip_firewall_service_port` and
  `ip_hotspot_service_port` (`name`), `queue_interface` (`interface`) —
  which are now adopted by that field (List+Set, like `ip_service`) instead
  of an `add` the device refuses, so those resources work at all.
- Exhaustive coverage pass. Diffed every resource against MikroTik's
  machine-generated CLI reference (auto-extracted from the firmware) and
  cross-verified each candidate against a live RouterOS 7.x device: 968
  settable arguments that no attribute could express are now present,
  including the dotted BGP and IS-IS sub-objects (`input.accept-communities`,
  `output.no-early-cut`, `l1.lsp-max-age`, …) and broad flat coverage on
  `/tool/netwatch`, `/interface/bonding`, `/ip/hotspot/profile`,
  `/ipv6/dhcp-client`, `/queue/type` and ~130 other menus. New secret-bearing
  fields (`passphrase`, `ipsec-secret`, ipsec `secret`/`password`, …) are
  marked sensitive. Every added key was confirmed accepted by the device;
  hardware-only keys the CHR cannot expose were deliberately excluded.
- `routeros_user_manager_user` and `routeros_ip_dhcp_server_option_sets`
  declared no manageable attributes at all — only `id` and `router` — so
  neither could express any configuration. The first now carries `name`,
  `password`, `group`, `shared_users`, `otp_secret`, `attributes` and
  `comment` (secrets marked sensitive); the second carries `name` and
  `options`.
- `routeros_interface_ethernet_switch` gained `name`, `mirror_source`,
  `mirror_target` and `cpu_flow_control`. Without them the resource had
  nothing declarable and produced no usable configuration.
- `routeros_caps_man_configuration` gained the `datapath`, `security` and
  `rates` profile references, which were missing alongside the existing
  `channel`.
- `routeros_system_routerboard_mode_button`,
  `routeros_system_routerboard_wps_button` and
  `routeros_system_routerboard_reset_button` resources for the RouterBOARD
  hardware buttons. The scripts these fire were already manageable as
  `routeros_system_script`, but the bindings that invoke them were not, so
  a rebuilt device came up with its buttons unbound.

### Fixed

- `routeros_ip_ipsec_key_qkd` is a singleton on the device (REST returns a
  bare object), but was modelled as a collection, so `import` could never
  resolve it. Converted to the singleton pattern.
- `routeros_ip_firewall_service_port.ports` was typed `int`, so it could not
  hold a multi-port value like `5060,5061` (the `sip` default). Now `string`.
- `routeros_file.type` was typed `int` but holds a category word (`backup`,
  `directory`, …); it is now a read-only string.

- `routeros_routing_fantasy` uses `route_count`, not the reserved Terraform
  meta-argument `count`. A `count` attribute makes Terraform reject the whole
  provider at load. Guarded by `TestNoReservedAttributeNames` and a real
  `terraform providers schema` load.

- Attributes that RouterOS reports as a word were typed as integers, so a
  device sitting on its own defaults produced an unfixable diff and the
  sentinel could never be written back. `interface_bridge.mtu` (`auto`),
  `interface_bridge_port.horizon` (`none`) and `ipv6_nd`'s `hop_limit`,
  `mtu`, `reachable_time` and `retransmit_interval` (`unspecified`) are
  now strings.
- `interface_ethernet.auto_negotiation` modelled the read-only link state
  (`done`, `incomplete`, …) rather than the settable yes/no toggle, and
  was never written to the wire at all. It is now a `bool` that maps to
  the real `auto-negotiation` property. The `autoneg` and `noautoneg`
  attributes are WebFig-only spellings that RouterOS rejects over REST;
  they are deprecated and no longer sent.
- `routeros_interface_wifi` and `routeros_interface_wifi_configuration`
  addressed sub-object properties by flat name, but RouterOS exposes them
  under dotted REST keys — `configuration.ssid`, `security.passphrase`,
  `channel.band` and so on. Neither read nor write matched what the device
  reports, so none of the wifi settings round-tripped and a CAP rebuilt
  from Terraform came up with no SSID and no passphrase. 292 field
  references across the two resources now use the dotted form. Terraform
  attribute names are unchanged, so existing configurations keep working
  — they simply reach the device now.
  Note the two menus differ: on `/interface/wifi` the `configuration.*`
  members are sub-objects, while on `/interface/wifi/configuration` the
  same names (`ssid`, `mode`, `country`, `manager`) are genuine top-level
  properties and are deliberately left flat.
- `interface_wifi.eap_password` is now marked sensitive, matching
  `passphrase`.
- 561 attributes across 108 resources advertised themselves as `Optional`
  while holding read-only runtime state (`dynamic`, `invalid`,
  `actual_mtu`, `fingerprint`, `owner`, …). Setting one was silently
  ignored. They are now `Computed`-only, so they still populate state but
  a value in configuration is reported as an error instead of vanishing.
- Several REST field names were derived by splitting Go identifiers, so
  they addressed properties RouterOS does not have and every write to
  them was silently wrong. The device reports `poe-out`, `poe-priority`,
  `poe-out-power`/`-current`/`-status`/`-voltage`, `poe-voltage` and
  `dont-require-permissions`; the provider sent `po-e-out`,
  `po-e-priority`, `po-e-*` and `don-t-require-permissions`. The wire
  names are corrected and the matching attributes renamed
  (`po_e_out` → `poe_out`, `po_e_priority` → `poe_priority`,
  `po_e_voltage` → `poe_voltage`, `po_e_out_*` → `poe_out_*`,
  `don_t_require_permissions` → `dont_require_permissions`). Configurations
  using the old names must be updated, but they could never have worked.
- A sweep for the same defect found five more WebFig-only spellings that
  shadow a real property and were being written to the REST API, so any
  configuration that set one failed to apply: `ip_hotspot_user.def` and
  `.nondef` (shadowing the read-only `default`),
  `ip_ipsec_mode_config.resp` and `.nonresp` (shadowing `responder`),
  `ip_ipsec_policy.notemplate` (shadowing `template`) and
  `ppp_profile.def`. All are deprecated and no longer sent; the real
  properties were already wired and are unaffected.
- `arp_timeout` rejected `auto` — the factory default on bridge, ethernet,
  EoIP, mesh and VPLS interfaces — because it was validated as a plain
  duration. It now accepts a duration or `auto`.
- `ip_ipsec_profile.dpd_interval` accepted only `disable-dpd`, rejecting
  the `8s` default. It now accepts any duration or the keyword.
- `snmp_community.authentication_protocol` and `encryption_protocol`
  rejected `MD5` and `DES` — the exact values RouterOS reports. Matching
  is now case-insensitive and normalised to the device's spelling.

### Security

- Updated `golang.org/x/crypto` (0.50.0 → 0.54.0), `golang.org/x/net`
  (0.52.0 → 0.57.0) and `google.golang.org/grpc` (1.79.3 → 1.82.1),
  clearing 15 advisories against transitive dependencies, several rated
  critical.

## [1.2.0] - 2026-06-05

### Added

- Import IDs accept an explicit `<router>::<key>` form, removing the
  ambiguity of the legacy `<router>/<key>` form.
- Natural-key import now matches across the common key columns (name,
  addresses, gateway, interface, host, comment, ...) instead of `name` only.
- Secret attributes (e.g. user SSH key material) are now marked
  `Sensitive`, so Terraform masks them in plan/apply output and state diffs.

### Fixed

- RouterOS `.id` URL encoding: special characters in keys (spaces, `#`,
  `?`, non-ASCII) are now encoded correctly while the literal `*` required
  by `/rest` is preserved.
- Import IDs shaped like CIDRs (e.g. `10.0.0.1/24`) are no longer
  misparsed as a `<router>/<key>` pair.
- Invalid values that fail normalization are now reported at plan time
  instead of being silently deferred.

### Changed

- Ordered writes are serialized per resource with internal locks, avoiding
  interleaving when multiple resources of the same type apply concurrently.
- HTTP client hardening: HTTP/2 is defensively disabled for the RouterOS
  REST endpoint, the response-header timeout is bounded, and in-flight
  requests stop retrying once the context is cancelled or times out.
- Internal: provider sources are maintained directly (generation markers
  and `_gen` filename suffixes removed); per-resource import/lookup
  boilerplate consolidated into shared helpers. No change to provider
  behavior or schema.

## [1.1.1] - 2026-06-05

### Fixed

- `routeros_interface_wifi_channel`: the `band` attribute now accepts the
  RouterOS 7 `wifi` package values, including `2ghz-ax`/`5ghz-ax` (Wi-Fi 6),
  `2ghz-be`/`5ghz-be` (Wi-Fi 7) and the 6 GHz bands (`6ghz-ax`, `6ghz-be`).
  Previously only the legacy wireless band values were allowed, rejecting
  valid Wi-Fi 6 configs on devices like the hAP ax3 (#1).

## [1.0.2] - 2026-05-24

- Provider landing page (`docs/index.md`) rewritten so the Terraform
  Registry renders it correctly: dropped the manually rolled `## Schema`
  block that collided with the Registry's auto-rendered provider schema.
- Per-resource and per-data-source docs now follow the standard
  `# Resource: name`, `## Argument Reference`, `## Attribute Reference`,
  `## Import` layout used by the AWS, Google, and Azure providers.
- Hand-written multi-resource workflow examples moved to `docs/guides/`
  (firewall, home-network, wireguard-vpn, certificate, system).

## [1.0.1] - 2026-05-24

- Release manifest checksum (`terraform-registry-manifest.json`) is now
  included in `SHA256SUMS`, unblocking Registry validation.
- Releases publish directly (no draft).

## [1.0.0] - 2026-05-24

First release.

- 186 resources, 277 data sources, 75 actions, 3522 properties.
- Multi-router from a single provider block via the `routers` map; every
  resource and data source takes an optional `router` attribute.
- Deterministic firewall rule ordering through a per-rule `position`
  integer; order survives destroy and recreate without churn.
- Per-row actions (sign, format, eject, reset-counters, ...) driven via
  `target_id` and a free-form `params` map.
- Lockout safety guards on firewall, user, user-group, and mac-server
  resources. Override per resource with `lockout_ack = true`.
- Sensitive properties (passwords, secrets, keys, OTP, SIM-PIN) flagged
  so Terraform redacts them in plan output, state, and CLI display.
- Drift detection: out-of-band REST changes detected on the next plan
  and reconciled by the following apply.
- Import support: `terraform import resource.x "<router>/<.id>"`.
- Five end-to-end workflow guides verified live against a CHR
  (firewall, home-network, wireguard-vpn, certificate, system).
