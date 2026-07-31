---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_graphing"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_graphing

Manages the RouterOS `/tool/graphing` menu.

## Example Usage

```terraform
resource "routeros_tool_graphing" "graphing_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # page_refresh = 0
  # store_every = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `page_refresh` - (Optional) Type: `int`.
* `store_every` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_tool_graphing.this 'home'
```
