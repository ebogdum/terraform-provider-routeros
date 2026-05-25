---
subcategory: "Scripts & scheduler"
page_title: "RouterOS: routeros_system_scheduler"
description: |-
  RouterOS resource.
---

# Data Source: routeros_system_scheduler

Manages the RouterOS `/system/scheduler` menu.

## Example Usage

```terraform
data "routeros_system_scheduler" "scheduler_example" {
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
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interval` - (Optional) Type: `duration`. Default: `1h`.
* `name` - (Required) Type: `string`. Default: `tf-acc-sched`.
* `on_event` - (Required) Type: `string`. Default: `:put "tick"`.
* `policy` - (Optional) Type: `string`.
* `start_date` - (Optional) Type: `string`.
* `start_time` - (Optional) Type: `enum(startup)`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `next_run` - Type: `string`.
* `owner` - Type: `string`.
* `run_count` - Type: `int`.

