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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `page_refresh` - (Optional) Type: `int`.
* `store_every` - (Optional) Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_tool_graphing.this 'home'
```
