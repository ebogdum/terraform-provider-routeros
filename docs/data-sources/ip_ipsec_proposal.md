---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_ipsec_proposal"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_ipsec_proposal

Manages the RouterOS `/ip/ipsec/proposal` menu.

## Example Usage

```terraform
data "routeros_ip_ipsec_proposal" "proposal_example" {
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
* `auth_algorithms` - (Optional) Type: `enum(md5|sha1|null|sha256|sha512)`. Default: `128`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `enc_algorithms` - (Optional) Type: `list`.
* `lifetime` - (Optional) Type: `duration`. Default: `1800`.
* `name` - (Optional) Type: `string`.
* `pfs_group` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.

