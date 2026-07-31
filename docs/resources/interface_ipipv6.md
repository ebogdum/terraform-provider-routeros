---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ipipv6"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_ipipv6

Manages the RouterOS `/interface/ipipv6` menu.

## Example Usage

```terraform
resource "routeros_interface_ipipv6" "ipipv6_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # ipsec_secret = "REDACTED"
  # local_address = "10.99.0.1"
  # mtu = "replace-me"
  # name = "tf-example"
  # remote_address = "10.99.0.1"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `clamp_tcp_mss` - (Optional) Type: `string`. RouterOS `clamp-tcp-mss`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dont_fragment` - (Optional) Type: `string`. RouterOS `dont-fragment`.
* `dscp` - (Optional) Type: `string`. RouterOS `dscp`.
* `ipsec_secret` - (Optional) Type: `string`. **Sensitive.**
* `keepalive` - (Optional) Type: `string`. RouterOS `keepalive`.
* `local_address` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `remote_address` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_ipipv6.example '*3'

# Named router
terraform import routeros_interface_ipipv6.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_ipipv6.example 'home/my-resource-name'
```
