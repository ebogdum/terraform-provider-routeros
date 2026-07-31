---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_l2tp_ether"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_l2tp_ether

Manages the RouterOS `/interface/l2tp-ether` menu.

## Example Usage

```terraform
resource "routeros_interface_l2tp_ether" "l2tp_ether_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # ipsec_secret = "REDACTED"
  # local_address = "10.99.0.1"
  # mac_address = "10.99.0.0/24"
  # mtu = "replace-me"
  # name = "tf-example"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `allow_fast_path` - (Optional) Type: `string`. RouterOS `allow-fast-path`.
* `circuit_id` - (Optional) Type: `string`. RouterOS `circuit-id`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connect_to` - (Optional) Type: `string`. RouterOS `connect-to`.
* `cookie_length` - (Optional) Type: `string`. RouterOS `cookie-length`.
* `digest_hash` - (Optional) Type: `string`. RouterOS `digest-hash`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ipsec_secret` - (Optional) Type: `string`. **Sensitive.**
* `l2tp_proto_version` - (Optional) Type: `string`. RouterOS `l2tp-proto-version`.
* `local_address` - (Optional) Type: `string`.
* `local_session_id` - (Optional) Type: `string`. RouterOS `local-session-id`.
* `local_tunnel_id` - (Optional) Type: `string`. RouterOS `local-tunnel-id`.
* `mac_address` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `peer_cookie` - (Optional) Type: `string`. RouterOS `peer-cookie`.
* `remote_session_id` - (Optional) Type: `string`. RouterOS `remote-session-id`.
* `remote_tunnel_id` - (Optional) Type: `string`. RouterOS `remote-tunnel-id`.
* `send_cookie` - (Optional) Type: `string`. RouterOS `send-cookie`.
* `unmanaged_mode` - (Optional) Type: `string`. RouterOS `unmanaged-mode`.
* `use_ipsec` - (Optional) Type: `string`. RouterOS `use-ipsec`.
* `use_l2_specific_sublayer` - (Optional) Type: `string`. RouterOS `use-l2-specific-sublayer`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_l2tp_ether.example '*3'

# Named router
terraform import routeros_interface_l2tp_ether.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_l2tp_ether.example 'home/my-resource-name'
```
