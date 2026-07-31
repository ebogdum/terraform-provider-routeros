---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ethernet_poe"
description: |-
  Requires PoE-capable ethernet port
---

# Resource: routeros_interface_ethernet_poe

Requires PoE-capable ethernet port

## Example Usage

```terraform
resource "routeros_interface_ethernet_poe" "poe_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # export = "replace-me"
  # monitor = "replace-me"
  # power_cycle = "replace-me"
  # print = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `duration` - (Optional) Type: `string`. RouterOS `duration`.
* `ether1_poe_in_long_cable` - (Optional) Type: `string`. RouterOS `ether1-poe-in-long-cable`.
* `export` - (Optional) Type: `string`. export is displayed under /interface ethernet menu.
* `jack1_max_power` - (Optional) Type: `string`. RouterOS `jack1-max-power`.
* `jack2_max_power` - (Optional) Type: `string`. RouterOS `jack2-max-power`.
* `jack_max_power` - (Optional) Type: `string`. RouterOS `jack-max-power`.
* `monitor` - (Optional) Type: `string`. Shows poe-out-status of a specified port, or all ports with /interface ethernet poe monitor [find] command.
* `name` - (Optional) Type: `string`. RouterOS `name`.
* `poe_in_max_power` - (Optional) Type: `string`. RouterOS `poe-in-max-power`.
* `poe_out` - (Optional) Type: `string`. RouterOS `poe-out`.
* `poe_priority` - (Optional) Type: `string`. RouterOS `poe-priority`.
* `poe_voltage` - (Optional) Type: `string`. RouterOS `poe-voltage`.
* `power_cycle` - (Optional) Type: `string`. Disables PoE-Out power for a specified period of time.
* `power_cycle_interval` - (Optional) Type: `string`. RouterOS `power-cycle-interval`.
* `power_cycle_ping_address` - (Optional) Type: `string`. RouterOS `power-cycle-ping-address`.
* `power_cycle_ping_enabled` - (Optional) Type: `string`. RouterOS `power-cycle-ping-enabled`.
* `power_cycle_ping_timeout` - (Optional) Type: `string`. RouterOS `power-cycle-ping-timeout`.
* `print` - (Optional) Type: `string`. Prints PoE-Out related settings.
* `psu1_max_power` - (Optional) Type: `string`. RouterOS `psu1-max-power`.
* `psu2_max_power` - (Optional) Type: `string`. RouterOS `psu2-max-power`.
* `psu_max_power` - (Optional) Type: `string`. RouterOS `psu-max-power`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_ethernet_poe.example '*3'

# Named router
terraform import routeros_interface_ethernet_poe.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_ethernet_poe.example 'home/my-resource-name'
```
