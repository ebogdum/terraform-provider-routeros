---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_graphing_resource"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_graphing_resource

Manages the RouterOS `/tool/graphing/resource` menu.

## Example Usage

```terraform
resource "routeros_tool_graphing_resource" "resource_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `allow_address` - (Optional) Type: `string`. RouterOS `allow-address`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `store_on_disk` - (Optional) Type: `string`. RouterOS `store-on-disk`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_tool_graphing_resource.example '*3'

# Named router
terraform import routeros_tool_graphing_resource.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_tool_graphing_resource.example 'home/my-resource-name'
```
