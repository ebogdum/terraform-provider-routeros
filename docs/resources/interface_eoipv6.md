---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_eoipv6"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_eoipv6

Manages the RouterOS `/interface/eoipv6` menu.

## Example Usage

```terraform
resource "routeros_interface_eoipv6" "eoipv6_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"
  remote_address = "10.99.0.1"
  tunnel_id = "1"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # ipsec_secret = "REDACTED"
  # local_address = "10.99.0.1"
  # mac_address = "10.99.0.0/24"
  # mtu = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `arp` - (Optional) Type: `string`.
* `arp_timeout` - (Optional) Type: `string`.
* `clamp_tcp_mss` - (Optional) Type: `string`. RouterOS `clamp-tcp-mss`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dont_fragment` - (Optional) Type: `string`. RouterOS `dont-fragment`.
* `dscp` - (Optional) Type: `string`. RouterOS `dscp`.
* `ipsec_secret` - (Optional) Type: `string`. **Sensitive.**
* `keepalive` - (Optional) Type: `string`. RouterOS `keepalive`.
* `local_address` - (Optional) Type: `string`.
* `loop_protect` - (Optional) Type: `string`. RouterOS `loop-protect`.
* `loop_protect_disable_time` - (Optional) Type: `string`. RouterOS `loop-protect-disable-time`.
* `loop_protect_send_interval` - (Optional) Type: `string`. RouterOS `loop-protect-send-interval`.
* `mac_address` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `remote_address` - (Required) Type: `string`.
* `tunnel_id` - (Required) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_eoipv6.example '*3'

# Named router
terraform import routeros_interface_eoipv6.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_eoipv6.example 'home/my-resource-name'
```
