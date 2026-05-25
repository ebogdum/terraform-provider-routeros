---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_neighbor"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_neighbor

Manages the RouterOS `/ip/neighbor` menu.

## Example Usage

```terraform
data "routeros_ip_neighbor" "neighbor_example" {
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
* `address` - (Optional) Type: `ip`.
* `address4` - (Optional) Type: `ip`.
* `address6` - (Optional) Type: `ip`.
* `age` - (Optional) Type: `duration`.
* `board` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `discovered_by` - (Optional) Type: `list`.
* `identity` - (Optional) Type: `string`.
* `interface` - (Optional) Type: `string`.
* `interface_name` - (Optional) Type: `string`.
* `ipv6` - (Optional) Type: `bool`.
* `mac_address` - (Optional) Type: `mac`.
* `platform` - (Optional) Type: `string`.
* `running` - (Optional) Type: `string`.
* `software_id` - (Optional) Type: `string`.
* `system_caps` - (Optional) Type: `string`.
* `system_caps_enabled` - (Optional) Type: `string`.
* `unpack` - (Optional) Type: `string`.
* `uptime` - (Optional) Type: `duration`.
* `version` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

