---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_smb_shares"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_smb_shares

Manages the RouterOS `/ip/smb/shares` menu.

## Example Usage

```terraform
data "routeros_ip_smb_shares" "shares_example" {
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
* `directory` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `invalid_users` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_smbshare`.
* `newfileman` - (Optional) Type: `string`.
* `old_directory` - (Optional) Type: `string`.
* `oldfileman` - (Optional) Type: `string`.
* `read_only` - (Optional) Type: `bool`.
* `require_encryption` - (Optional) Type: `bool`.
* `valid_users` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.
* `dynamic` - Type: `bool`.

