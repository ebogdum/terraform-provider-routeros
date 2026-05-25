---
page_title: "RouterOS: routeros_ip_nat_pmp"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_nat_pmp

Manages the RouterOS `/ip/nat-pmp` menu.

## Example Usage

```terraform
resource "routeros_ip_nat_pmp" "nat_pmp_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # enabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `enabled` - (Optional) Type: `bool`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_nat_pmp.this 'home'
```
