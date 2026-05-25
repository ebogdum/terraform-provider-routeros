---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_cloud_back_to_home_user"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_cloud_back_to_home_user

Manages the RouterOS `/ip/cloud/back-to-home-user` menu.

## Example Usage

```terraform
data "routeros_ip_cloud_back_to_home_user" "back_to_home_user_example" {
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
* `active` - (Optional) Type: `bool`.
* `allow_lan` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `expires` - (Optional) Type: `string`.
* `file_access_mode` - (Optional) Type: `enum(|disabled|read-only|full)`.
* `files` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `newe` - (Optional) Type: `string`.
* `newfileman` - (Optional) Type: `string`.
* `notnew` - (Optional) Type: `string`.
* `oldfileman` - (Optional) Type: `string`.
* `private_key` - (Optional) Type: `string`.
* `public_key` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `client_address` - Type: `string`.
* `client_config` - Type: `string`.
* `client_qr` - Type: `string`.
* `file_access_token` - Type: `string`.

