---
page_title: "RouterOS: routeros_port_remote_access"
description: |-
  RouterOS resource.
---

# Data Source: routeros_port_remote_access

Manages the RouterOS `/port/remote-access` menu.

## Example Usage

```terraform
data "routeros_port_remote_access" "remote_access_example" {
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
* `channel` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `local_address` - (Optional) Type: `string`.
* `port` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

