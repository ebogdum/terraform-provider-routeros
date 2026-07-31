---
subcategory: "Interface"
page_title: "RouterOS: routeros_interface_ethernet_switch_port"
description: |-
  Mirrors RouterOS /interface/ethernet/switch/port.
---

# Resource: routeros_interface_ethernet_switch_port

Mirrors RouterOS `/interface/ethernet/switch/port`.

## Example Usage

```terraform
resource "routeros_interface_ethernet_switch_port" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # default_vlan_id = 0
  # name = "replace-me"
  # switch = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `default_vlan_id` - (Optional) Type: `string`. RouterOS `default-vlan-id`. A number, or `auto`.
* `name` - (Optional) Type: `string`. RouterOS `name`.
* `switch` - (Optional) Type: `string`. RouterOS `switch`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
terraform import routeros_interface_ethernet_switch_port.example 'home::*3'
```
