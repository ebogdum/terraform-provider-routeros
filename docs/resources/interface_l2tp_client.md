---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_l2tp_client"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_l2tp_client

Manages the RouterOS `/interface/l2tp-client` menu.

## Example Usage

```terraform
resource "routeros_interface_l2tp_client" "l2tp_client_example" {
  # router = "my-router"  # which router to target; omit for the default
  connect_to = "127.0.0.1"
  name = "tf-example"
  user = "myuser"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # allow = "replace-me"
  # default_route_distance = "replace-me"
  # dial_on_demand = "replace-me"
  # ipsec_secret = "REDACTED"
  # keepalive_timeout = "replace-me"
  # max_mru = "replace-me"
  # max_mtu = "replace-me"
  # mrru = "replace-me"
  # password = "REDACTED"
  # profile = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `add_default_route` - (Optional) Type: `string`. RouterOS `add-default-route`.
* `allow` - (Optional) Type: `string`.
* `allow_fast_path` - (Optional) Type: `string`. RouterOS `allow-fast-path`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connect_to` - (Required) Type: `string`.
* `default_route_distance` - (Optional) Type: `string`.
* `dial_on_demand` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ipsec_secret` - (Optional) Type: `string`. **Sensitive.**
* `keepalive_timeout` - (Optional) Type: `string`.
* `l2tp_proto_version` - (Optional) Type: `string`. RouterOS `l2tp-proto-version`.
* `l2tpv3_circuit_id` - (Optional) Type: `string`. RouterOS `l2tpv3-circuit-id`.
* `l2tpv3_cookie_length` - (Optional) Type: `string`. RouterOS `l2tpv3-cookie-length`.
* `l2tpv3_digest_hash` - (Optional) Type: `string`. RouterOS `l2tpv3-digest-hash`.
* `max_mru` - (Optional) Type: `string`.
* `max_mtu` - (Optional) Type: `string`.
* `mrru` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `password` - (Optional) Type: `string`. **Sensitive.**
* `profile` - (Optional) Type: `string`.
* `random_source_port` - (Optional) Type: `string`. RouterOS `random-source-port`.
* `src_address` - (Optional) Type: `string`. RouterOS `src-address`.
* `use_ipsec` - (Optional) Type: `string`. RouterOS `use-ipsec`.
* `use_peer_dns` - (Optional) Type: `string`. RouterOS `use-peer-dns`.
* `user` - (Required) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_l2tp_client.example '*3'

# Named router
terraform import routeros_interface_l2tp_client.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_l2tp_client.example 'home/my-resource-name'
```
