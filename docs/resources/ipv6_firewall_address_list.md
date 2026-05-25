---
page_title: "RouterOS: routeros_ipv6_firewall_address_list"
description: |-
  RouterOS resource.
---

# Resource: routeros_ipv6_firewall_address_list

Manages the RouterOS `/ipv6/firewall/address-list` menu.

## Example Usage

```terraform
resource "routeros_ipv6_firewall_address_list" "address_list_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "fd00:db8::/64"
  list = "my-list"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # dynamic = "replace-me"
  # timeout = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Required) Type: `string`. Default: `fd00:db8::/64`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic` - (Optional) Type: `string`.
* `list` - (Required) Type: `string`. Default: `tf_acc_list6`.
* `timeout` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_firewall_address_list.example '*3'

# Named router
terraform import routeros_ipv6_firewall_address_list.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_firewall_address_list.example 'home/my-resource-name'
```
