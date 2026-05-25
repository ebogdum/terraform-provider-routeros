---
subcategory: "Certificates"
page_title: "RouterOS: routeros_certificate_crl"
description: |-
  RouterOS resource.
---

# Data Source: routeros_certificate_crl

Manages the RouterOS `/certificate/crl` menu.

## Example Usage

```terraform
data "routeros_certificate_crl" "crl_example" {
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
* `download` - (Optional) Type: `string`.
* `expired` - (Optional) Type: `bool`.
* `flush` - (Optional) Type: `string`.
* `url` - (Required) Type: `string`. Default: `http://invalid.example/crl`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `akid` - Type: `string`.
* `certificate` - Type: `string`.
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.
* `last_update` - Type: `string`.
* `next_update` - Type: `string`.
* `num` - Type: `int`.
* `revoked` - Type: `int`.
* `signature` - Type: `string`.

