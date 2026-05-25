---
page_title: "RouterOS: routeros_ppp_aaa"
description: |-
  RouterOS resource.
---

# Resource: routeros_ppp_aaa

Manages the RouterOS `/ppp/aaa` menu.

## Example Usage

```terraform
resource "routeros_ppp_aaa" "aaa_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # accounting = false
  # enable_ipv6_accounting = false
  # interim_update = "1h"
  # use_circuit_id_in_nas_port_id = false
  # use_radius = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `accounting` - (Optional) Type: `bool`.
* `enable_ipv6_accounting` - (Optional) Type: `bool`.
* `interim_update` - (Optional) Type: `duration`.
* `use_circuit_id_in_nas_port_id` - (Optional) Type: `bool`.
* `use_radius` - (Optional) Type: `bool`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_ppp_aaa.this 'home'
```
