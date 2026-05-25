---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_bfd_session"
description: |-
  RouterOS resource.
---

# Data Source: routeros_routing_bfd_session

Manages the RouterOS `/routing/bfd/session` menu.

## Example Usage

```terraform
data "routeros_routing_bfd_session" "session_example" {
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
* `actual_tx_interval` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `desired_tx_interval` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `hold_time` - (Optional) Type: `string`.
* `inactive` - (Optional) Type: `bool`.
* `local_address` - (Optional) Type: `string`.
* `multiplier` - (Optional) Type: `int`.
* `packets_rx` - (Optional) Type: `int`.
* `packets_tx` - (Optional) Type: `int`.
* `remote_address` - (Optional) Type: `string`.
* `remote_min_rx` - (Optional) Type: `string`.
* `remote_min_tx` - (Optional) Type: `string`.
* `required_min_rx` - (Optional) Type: `string`.
* `state` - (Optional) Type: `enum(admin down|down|init|up)`.
* `state_changes` - (Optional) Type: `int`.
* `up` - (Optional) Type: `bool`.
* `uptime` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

