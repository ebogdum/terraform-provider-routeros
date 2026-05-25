---
page_title: "RouterOS: routeros_tool_romon_port"
description: |-
  RoMON port references a specific interface; values vary per device. Skipped.
---

# Resource: routeros_tool_romon_port

RoMON port references a specific interface; values vary per device. Skipped.

## Example Usage

```terraform
resource "routeros_tool_romon_port" "port_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # cost = 100
  # forbid = false
  # interface = "ether1"
  # secrets = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `cost` - (Optional) Type: `int`. Default: `100`.
* `disabled` - (Optional) Type: `bool`.
* `forbid` - (Optional) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `secrets` - (Optional) Type: `string`. **Sensitive.**

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.
* `dynamic` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_tool_romon_port.example '*3'

# Named router
terraform import routeros_tool_romon_port.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_tool_romon_port.example 'home/my-resource-name'
```
