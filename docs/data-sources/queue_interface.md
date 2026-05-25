---
subcategory: "Queues"
page_title: "RouterOS: routeros_queue_interface"
description: |-
  RouterOS resource.
---

# Data Source: routeros_queue_interface

Manages the RouterOS `/queue/interface` menu.

## Example Usage

```terraform
data "routeros_queue_interface" "interface_example" {
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
* `queue` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `active_queue` - Type: `string`.
* `default_queue` - Type: `string`.
* `interface` - Type: `string`.

