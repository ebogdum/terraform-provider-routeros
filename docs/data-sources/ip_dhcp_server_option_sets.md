---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_dhcp_server_option_sets"
description: |-
  Sets reference existing /ip/dhcp-server/option entries; skipped from automated acc tests.
---

# Data Source: routeros_ip_dhcp_server_option_sets

Sets reference existing /ip/dhcp-server/option entries; skipped from automated acc tests.

## Example Usage

```terraform
data "routeros_ip_dhcp_server_option_sets" "sets_example" {
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

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

