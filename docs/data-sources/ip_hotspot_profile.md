---
subcategory: "Hotspot"
page_title: "RouterOS: routeros_ip_hotspot_profile"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_hotspot_profile

Manages the RouterOS `/ip/hotspot/profile` menu.

## Example Usage

```terraform
data "routeros_ip_hotspot_profile" "profile_example" {
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
* `dns_name` - (Optional) Type: `string`.
* `hotspot_address` - (Optional) Type: `ip`.
* `html_directory` - (Optional) Type: `string`.
* `html_directory_override` - (Optional) Type: `string`.
* `http_cookie_lifetime` - (Optional) Type: `duration`.
* `http_proxy` - (Optional) Type: `string`.
* `install_hotspot_queue` - (Optional) Type: `bool`.
* `login_by` - (Optional) Type: `list`.
* `name` - (Optional) Type: `string`.
* `smtp_server` - (Optional) Type: `ip`.
* `split_user_domain` - (Optional) Type: `bool`.
* `use_radius` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.
* `default` - Type: `bool`.

