---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_mac_server"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_mac_server

Manages the RouterOS `/tool/mac-server` menu.

## Example Usage

```terraform
resource "routeros_tool_mac_server" "mac_server_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # allowed_interface_list = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `allowed_interface_list` - (Optional) Type: `string`.
* `lockout_ack` - (Optional) Type: `bool`. Acknowledge that this rule may sever management traffic (required for unconditional input/forward drop/reject/tarpit rules with no match).

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_tool_mac_server.this 'home'
```
