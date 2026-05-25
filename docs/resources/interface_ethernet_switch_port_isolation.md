---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ethernet_switch_port_isolation"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_ethernet_switch_port_isolation

Manages the RouterOS `/interface/ethernet/switch/port-isolation` menu.

## Example Usage

```terraform
resource "routeros_interface_ethernet_switch_port_isolation" "port_isolation_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # forward_to = "replace-me"
  # forwarding_override = false
  # override = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `forward_to` - (Optional) Type: `string`.
* `forwarding_override` - (Optional) Type: `bool`.
* `override` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `invalid` - Type: `bool`.
* `name` - Type: `string`.
* `switch` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_ethernet_switch_port_isolation.example '*3'

# Named router
terraform import routeros_interface_ethernet_switch_port_isolation.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_ethernet_switch_port_isolation.example 'home/my-resource-name'
```
