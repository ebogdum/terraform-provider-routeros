---
subcategory: "System"
page_title: "RouterOS: routeros_system_history"
description: |-
  RouterOS resource.
---

# Data Source: routeros_system_history

Manages the RouterOS `/system/history` menu.

## Example Usage

```terraform
data "routeros_system_history" "history_example" {
  # router   = "my-router"  # omit for the default router
  # filter   = { name = "some-name" }
  # proplist = ["name", "address"]
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `proplist` - (Optional) List of property names to project; smaller payload.
* `action` - (Optional) Type: `string`.
* `by` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `policy` - (Optional) Type: `list`.
* `redo` - (Optional) Type: `string`.
* `time` - (Optional) Type: `string`.
* `trace` - (Optional) Type: `string`.
* `undo` - (Optional) Type: `string`.
* `undoable` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `redo_cmds` - Type: `string`.
* `undo_cmds` - Type: `string`.

