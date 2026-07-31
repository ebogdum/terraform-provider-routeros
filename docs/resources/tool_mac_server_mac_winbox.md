---
subcategory: "Tool"
page_title: "RouterOS: routeros_tool_mac_server_mac_winbox"
description: |-
  Mirrors RouterOS /tool/mac-server/mac-winbox.
---

# Resource: routeros_tool_mac_server_mac_winbox

Mirrors RouterOS `/tool/mac-server/mac-winbox`.

## Example Usage

```terraform
resource "routeros_tool_mac_server_mac_winbox" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # allowed_interface_list = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `allowed_interface_list` - (Optional) Type: `string`. RouterOS `allowed-interface-list`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_tool_mac_server_mac_winbox.this 'home'
```
