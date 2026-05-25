---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_nd_proxy"
description: |-
  RouterOS resource.
---

# Resource: routeros_ipv6_nd_proxy

Manages the RouterOS `/ipv6/nd/proxy` menu.

## Example Usage

```terraform
resource "routeros_ipv6_nd_proxy" "proxy_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.1"
  # interface = "ether1"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Optional) Type: `ipv6`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `interface` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_nd_proxy.example '*3'

# Named router
terraform import routeros_ipv6_nd_proxy.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_nd_proxy.example 'home/my-resource-name'
```
