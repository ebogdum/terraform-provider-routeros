---
page_title: "RouterOS: routeros_ip_hotspot_user_profile"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_hotspot_user_profile

Manages the RouterOS `/ip/hotspot/user/profile` menu.

## Example Usage

```terraform
data "routeros_ip_hotspot_user_profile" "profile_example" {
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
* `add_mac_cookie` - (Optional) Type: `bool`.
* `address_list` - (Optional) Type: `string`.
* `idle_timeout` - (Optional) Type: `string`.
* `keepalive_timeout` - (Optional) Type: `duration`.
* `mac_cookie_timeout` - (Optional) Type: `duration`.
* `name` - (Optional) Type: `string`.
* `shared_users` - (Optional) Type: `int`.
* `status_autorefresh` - (Optional) Type: `duration`.
* `transparent_proxy` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.

