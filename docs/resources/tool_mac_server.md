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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `allowed_interface_list` - (Optional) Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_tool_mac_server.this 'home'
```
