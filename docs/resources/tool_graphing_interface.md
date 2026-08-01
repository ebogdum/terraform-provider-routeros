---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_graphing_interface"
description: |-
  Graphing rules are unique per (interface, dest-addr) — running tests repeatedly hits "already exists" without explicit cleanup.
---

# Resource: routeros_tool_graphing_interface

Graphing rules are unique per (interface, dest-addr) — running tests repeatedly hits "already exists" without explicit cleanup.

## Example Usage

```terraform
resource "routeros_tool_graphing_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_tool_graphing_interface.example '*3'

# Named router
terraform import routeros_tool_graphing_interface.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_tool_graphing_interface.example 'home/my-resource-name'
```
