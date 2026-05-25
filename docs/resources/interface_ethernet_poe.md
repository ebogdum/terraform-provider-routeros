---
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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `export` - (Optional) Type: `string`. export is displayed under   /interface ethernet   menu.
* `monitor` - (Optional) Type: `string`. Shows poe-out-status of a specified port, or all ports with   /interface ethernet poe monitor [find]   command.
* `power_cycle` - (Optional) Type: `string`. Disables PoE-Out power for a specified period of time.
* `print` - (Optional) Type: `string`. Prints PoE-Out related settings.

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
