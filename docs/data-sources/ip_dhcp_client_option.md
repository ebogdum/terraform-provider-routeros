---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_dhcp_client_option"
description: |-
  DHCP client option requires an existing dhcp-client + named option. Skipped.
---

# Data Source: routeros_ip_dhcp_client_option

DHCP client option requires an existing dhcp-client + named option. Skipped.

## Example Usage

```terraform
data "routeros_ip_dhcp_client_option" "option_example" {
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
* `code` - (Optional) Type: `int`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `name` - (Optional) Type: `string`.
* `value` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.
* `raw_value` - Type: `string`.

