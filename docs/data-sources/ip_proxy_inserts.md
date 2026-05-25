---
page_title: "RouterOS: routeros_ip_proxy_inserts"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_proxy_inserts

Manages the RouterOS `/ip/proxy/inserts` menu.

## Example Usage

```terraform
data "routeros_ip_proxy_inserts" "inserts_example" {
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
* `denied` - (Optional) Type: `int`.
* `errors` - (Optional) Type: `int`.
* `no_memory` - (Optional) Type: `int`.
* `successes` - (Optional) Type: `int`.
* `too_large` - (Optional) Type: `int`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

