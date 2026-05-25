---
subcategory: "Firewall"
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
  # parent = 0
  # timeout = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Required) Type: `string`. Default: `fd00:db8::/64`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `list` - (Required) Type: `string`. Default: `tf_acc_list6`.
* `parent` - (Optional) Type: `int`. Default: `4.294967295e+09`.
* `timeout` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `creation_time` - Type: `string`.
* `dynamic` - Type: `bool`.

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
