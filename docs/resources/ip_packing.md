---
page_title: "RouterOS: routeros_ip_packing"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_packing

Manages the RouterOS `/ip/packing` menu.

## Example Usage

```terraform
resource "routeros_ip_packing" "packing_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"

  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_packing.example '*3'

# Named router
terraform import routeros_ip_packing.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_packing.example 'home/my-resource-name'
```
