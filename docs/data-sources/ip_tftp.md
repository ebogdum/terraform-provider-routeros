---
page_title: "RouterOS: routeros_ip_tftp"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_tftp

Manages the RouterOS `/ip/tftp` menu.

## Example Usage

```terraform
data "routeros_ip_tftp" "tftp_example" {
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
* `allow` - (Optional) Type: `bool`. Default: `1`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ip_addresses` - (Optional) Type: `string`.
* `read_only` - (Optional) Type: `bool`. Default: `1`.
* `real_filename` - (Optional) Type: `string`.
* `req_filename` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

