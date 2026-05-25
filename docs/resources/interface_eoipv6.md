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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `arp` - (Optional) Type: `string`.
* `arp_timeout` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ipsec_secret` - (Optional) Type: `string`.
* `local_address` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_eoip6`.
* `remote_address` - (Required) Type: `string`. Default: `fd00:db8::1`.
* `tunnel_id` - (Required) Type: `string`. Default: `1`.

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
