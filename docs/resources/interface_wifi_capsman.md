---
page_title: "RouterOS: routeros_interface_wifi_capsman"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_capsman

Manages the RouterOS `/interface/wifi/capsman` menu.

## Example Usage

```terraform
resource "routeros_interface_wifi_capsman" "capsman_example" {
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
terraform import routeros_interface_wifi_capsman.this 'home'
```
