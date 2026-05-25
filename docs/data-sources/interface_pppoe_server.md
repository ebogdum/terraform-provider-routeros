---
page_title: "RouterOS: routeros_interface_pppoe_server"
description: |-
  Needs an interface to bind to, plus auth-stack setup.
---

# Data Source: routeros_interface_pppoe_server

Needs an interface to bind to, plus auth-stack setup.

## Example Usage

```terraform
data "routeros_interface_pppoe_server" "pppoe_server_example" {
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
* `name` - (Optional) Type: `string`.
* `service` - (Optional) Type: `string`.
* `user` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

