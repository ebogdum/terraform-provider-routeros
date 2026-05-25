---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_dns_forwarders"
description: |-
  Discovered; required address must be a valid resolvable IP
---

# Data Source: routeros_ip_dns_forwarders

Discovered; required address must be a valid resolvable IP

## Example Usage

```terraform
data "routeros_ip_dns_forwarders" "forwarders_example" {
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
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `dns_servers` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

