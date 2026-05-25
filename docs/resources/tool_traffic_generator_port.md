---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_traffic_generator_port"
description: |-
  Discovered; needs tgen config
---

# Resource: routeros_tool_traffic_generator_port

Discovered; needs tgen config

## Example Usage

```terraform
resource "routeros_tool_traffic_generator_port" "port_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # interface = "ether1"
  # name = "tf-example"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `disabled` - (Optional) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_tool_traffic_generator_port.example '*3'

# Named router
terraform import routeros_tool_traffic_generator_port.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_tool_traffic_generator_port.example 'home/my-resource-name'
```
