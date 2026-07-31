---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ethernet_switch_rule"
description: |-
  Requires switch chip presence
---

# Resource: routeros_interface_ethernet_switch_rule

Requires switch chip presence

## Example Usage

```terraform
resource "routeros_interface_ethernet_switch_rule" "rule_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # qos_hw_offloading = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `qos_hw_offloading` - (Optional) Type: `string`. Allows enabling QoS for the given switch chip (if the latter supports QoS). New generation devices force qos-hw-offloading=yes at all times.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_ethernet_switch_rule.example '*3'

# Named router
terraform import routeros_interface_ethernet_switch_rule.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_ethernet_switch_rule.example 'home/my-resource-name'
```
