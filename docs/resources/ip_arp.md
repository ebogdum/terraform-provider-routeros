---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_arp"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_arp

Manages the RouterOS `/ip/arp` menu.

## Example Usage

```terraform
resource "routeros_ip_arp" "arp_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "10.99.0.1"
  interface = "ether1"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # mac_address = "10.99.0.0/24"
  # published = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Required) Type: `ip`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `interface` - (Required) Type: `string`.
* `mac_address` - (Optional) Type: `mac`.
* `published` - (Optional) Type: `bool`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `complete` - Type: `bool`.
* `dhcp` - Type: `bool`.
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.
* `status` - Type: `string`.
* `vrf` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_arp.example '*3'

# Named router
terraform import routeros_ip_arp.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_arp.example 'home/my-resource-name'
```
