---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_firewall_address_list"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_firewall_address_list

Manages the RouterOS `/ip/firewall/address-list` menu.

## Example Usage

```terraform
resource "routeros_ip_firewall_address_list" "address_list_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "10.99.0.1"
  list = "my-list"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # timeout = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Required) Type: `string`. Default: `10.255.255.0/30`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `list` - (Required) Type: `string`. Default: `tf_acc_list`.
* `timeout` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `dynamic` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_firewall_address_list.example '*3'

# Named router
terraform import routeros_ip_firewall_address_list.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_firewall_address_list.example 'home/my-resource-name'
```
