---
subcategory: "System"
page_title: "RouterOS: routeros_system_logging_action"
description: |-
  RouterOS rejects hyphens AND underscores in action names on some 7.x versions; not portable.
---

# Data Source: routeros_system_logging_action

RouterOS rejects hyphens AND underscores in action names on some 7.x versions; not portable.

## Example Usage

```terraform
data "routeros_system_logging_action" "action_example" {
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
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disk_file_count` - (Optional) Type: `int`.
* `disk_file_name` - (Optional) Type: `string`.
* `disk_lines_per_file` - (Optional) Type: `int`.
* `disk_stop_on_full` - (Optional) Type: `bool`.
* `memory_lines` - (Optional) Type: `int`.
* `memory_stop_on_full` - (Optional) Type: `bool`.
* `name` - (Optional) Type: `string`.
* `remember` - (Optional) Type: `bool`.
* `remote` - (Optional) Type: `string`.
* `remote_log_format` - (Optional) Type: `string`.
* `remote_port` - (Optional) Type: `int`.
* `remote_protocol` - (Optional) Type: `string`.
* `src_address` - (Optional) Type: `ip`.
* `target` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.

