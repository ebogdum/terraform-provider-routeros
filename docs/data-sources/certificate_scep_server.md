---
subcategory: "Certificates"
page_title: "RouterOS: routeros_certificate_scep_server"
description: |-
  SCEP server references an existing CA cert. Skipped from automated acc tests.
---

# Data Source: routeros_certificate_scep_server

SCEP server references an existing CA cert. Skipped from automated acc tests.

## Example Usage

```terraform
data "routeros_certificate_scep_server" "scep_server_example" {
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
* `ca_certificate` - (Optional) Type: `string`.
* `days_valid` - (Optional) Type: `int`. Default: `1`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `next_ca_certificate` - (Optional) Type: `string`.
* `path` - (Optional) Type: `string`.
* `request_lifetime` - (Optional) Type: `duration`. Default: `3600`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

