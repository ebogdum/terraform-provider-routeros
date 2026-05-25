---
subcategory: "SNMP"
page_title: "RouterOS: routeros_snmp_community"
description: |-
  RouterOS resource.
---

# Data Source: routeros_snmp_community

Manages the RouterOS `/snmp/community` menu.

## Example Usage

```terraform
data "routeros_snmp_community" "community_example" {
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
* `addresses` - (Optional) Type: `cidr`.
* `authentication_password` - (Optional) Type: `string`. **Sensitive.**
* `authentication_protocol` - (Optional) Type: `enum(MD5|SHA1)`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `encryption_password` - (Optional) Type: `string`. **Sensitive.**
* `encryption_protocol` - (Optional) Type: `enum(DES|AES)`.
* `name` - (Required) Type: `string`. Default: `tf-acc-snmp-community`.
* `read_access` - (Optional) Type: `bool`. Default: `1`.
* `security` - (Optional) Type: `enum(none|authorized|private)`.
* `write_access` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.

