---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_igmp_proxy_interface"
description: |-
  RouterOS resource.
---

# Resource: routeros_routing_igmp_proxy_interface

Manages the RouterOS `/routing/igmp-proxy/interface` menu.

## Example Usage

```terraform
resource "routeros_routing_igmp_proxy_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # alternative_subnets = "replace-me"
  # interface = "ether1"
  # threshold = 1
  # upstream = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `alternative_subnets` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `threshold` - (Optional) Type: `int`. Default: `1`.
* `upstream` - (Optional) Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_igmp_proxy_interface.example '*3'

# Named router
terraform import routeros_routing_igmp_proxy_interface.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_igmp_proxy_interface.example 'home/my-resource-name'
```
