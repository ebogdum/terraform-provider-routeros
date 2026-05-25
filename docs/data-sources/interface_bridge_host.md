---
subcategory: "Bridge"
page_title: "RouterOS: routeros_interface_bridge_host"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Data Source: routeros_interface_bridge_host

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
data "routeros_interface_bridge_host" "host_example" {
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
* `aged` - (Optional) Type: `bool`.
* `aged_on_peer` - (Optional) Type: `bool`.
* `bridge` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `external_fdb` - (Optional) Type: `bool`.
* `interface` - (Optional) Type: `string`.
* `local` - (Optional) Type: `bool`.
* `mac_address` - (Optional) Type: `string`.
* `vid` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `dynamic` - Type: `bool`.
* `on_interface` - Type: `string`.
* `remote_ip` - Type: `string`.

