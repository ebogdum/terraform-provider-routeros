---
page_title: "RouterOS: routeros_routing_filter_community_list"
description: |-
  7.x routing-filter community-list uses 'communities' field and rule chain semantics that vary across releases. Skipped.
---

# Data Source: routeros_routing_filter_community_list

7.x routing-filter community-list uses 'communities' field and rule chain semantics that vary across releases. Skipped.

## Example Usage

```terraform
data "routeros_routing_filter_community_list" "community_list_example" {
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

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

