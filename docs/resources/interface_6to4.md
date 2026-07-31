---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_6to4"
description: |-
  6to4 tunnel deletion races on CHR (DELETE returns errors even after success). Skipped from automated acc tests.
---

# Resource: routeros_interface_6to4

6to4 tunnel deletion races on CHR (DELETE returns errors even after success). Skipped from automated acc tests.

## Example Usage

```terraform
resource "routeros_interface_6to4" "ex_6to4_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # clamp_tcp_mss = false
  # dont_fragment = false
  # dscp = "replace-me"
  # local_address = "10.99.0.1"
  # mtu = "replace-me"
  # name = "tf-example"
  # remote_address = "10.99.0.1"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `clamp_tcp_mss` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dont_fragment` - (Optional) Type: `bool`.
* `dscp` - (Optional) Type: `string`.
* `ipsec_secret` - (Optional) Type: `string`. RouterOS `ipsec-secret`. **Sensitive.**
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
terraform import routeros_interface_6to4.example '*3'

# Named router
terraform import routeros_interface_6to4.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_6to4.example 'home/my-resource-name'
```
