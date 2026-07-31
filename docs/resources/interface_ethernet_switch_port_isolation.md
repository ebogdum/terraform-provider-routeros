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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `forward_to` - (Optional) Type: `string`.
* `forwarding_override` - (Optional) Type: `bool`.
* `invalid` - (Read-only) Type: `bool`.
* `name` - (Read-only) Type: `string`.
* `override` - (Optional) Type: `string`.
* `switch` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
