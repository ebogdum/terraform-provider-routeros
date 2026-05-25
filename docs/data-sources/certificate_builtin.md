---
page_title: "RouterOS: routeros_certificate_builtin"
description: |-
  System-generated certificates -- read-only.
---

# Data Source: routeros_certificate_builtin

System-generated certificates -- read-only.

## Example Usage

```terraform
data "routeros_certificate_builtin" "builtin_example" {
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
* `akid` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `common_name` - (Optional) Type: `string`.
* `country` - (Optional) Type: `string`.
* `days_valid` - (Optional) Type: `int`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `invalid_after` - (Optional) Type: `string`.
* `invalid_before` - (Optional) Type: `string`.
* `issuer` - (Optional) Type: `string`.
* `key_type` - (Optional) Type: `string`.
* `key_usage` - (Optional) Type: `list`.
* `locality` - (Optional) Type: `string`.
* `organization` - (Optional) Type: `string`.
* `serial_number` - (Optional) Type: `string`.
* `skid` - (Optional) Type: `string`.
* `state` - (Optional) Type: `string`.
* `subject_alt_name` - (Optional) Type: `string`.
* `unit` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

