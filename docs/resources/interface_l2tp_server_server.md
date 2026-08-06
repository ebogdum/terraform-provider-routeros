---
subcategory: "Interface"
page_title: "RouterOS: routeros_interface_l2tp_server_server"
description: |-
  Mirrors RouterOS /interface/l2tp-server/server.
---

# Resource: routeros_interface_l2tp_server_server

Mirrors RouterOS `/interface/l2tp-server/server`.

## Example Usage

```terraform
resource "routeros_interface_l2tp_server_server" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # accept_proto_version = "replace-me"
  # accept_pseudowire_type = "replace-me"
  # allow_fast_path = true
  # authentication = "replace-me"
  # caller_id_type = "replace-me"
  # default_profile = "replace-me"
  # enabled = true
  # ipsec_secret = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `accept_proto_version` - (Optional) Type: `string`. RouterOS `accept-proto-version`.
* `accept_pseudowire_type` - (Optional) Type: `string`. RouterOS `accept-pseudowire-type`.
* `allow_fast_path` - (Optional) Type: `bool`. RouterOS `allow-fast-path`.
* `authentication` - (Optional) Type: `string`. RouterOS `authentication`.
* `caller_id_type` - (Optional) Type: `string`. RouterOS `caller-id-type`.
* `default_profile` - (Optional) Type: `string`. RouterOS `default-profile`.
* `enabled` - (Optional) Type: `bool`. RouterOS `enabled`.
* `ipsec_secret` - (Optional) Type: `string`. RouterOS `ipsec-secret`. **Sensitive.**
* `keepalive_timeout` - (Optional) Type: `int`. RouterOS `keepalive-timeout`.
* `l2tpv3_circuit_id` - (Optional) Type: `string`. RouterOS `l2tpv3-circuit-id`.
* `l2tpv3_cookie_length` - (Optional) Type: `int`. RouterOS `l2tpv3-cookie-length`.
* `l2tpv3_digest_hash` - (Optional) Type: `string`. RouterOS `l2tpv3-digest-hash`.
* `l2tpv3_ether_interface_list` - (Optional) Type: `string`. RouterOS `l2tpv3-ether-interface-list`.
* `max_mru` - (Optional) Type: `int`. RouterOS `max-mru`.
* `max_mtu` - (Optional) Type: `int`. RouterOS `max-mtu`.
* `max_sessions` - (Optional) Type: `string`. RouterOS `max-sessions`.
* `mrru` - (Optional) Type: `string`. RouterOS `mrru`.
* `one_session_per_host` - (Optional) Type: `bool`. RouterOS `one-session-per-host`.
* `use_ipsec` - (Optional) Type: `string`. IPsec usage for the L2TP server: `no`, `yes` (offer IPsec) or `required` (refuse plain L2TP).

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_interface_l2tp_server_server.this 'home'
```
