---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_dns_static"
description: |-
  A DNS A/AAAA/CNAME/MX/... static entry. Requires either name OR regexp.
---

# Data Source: routeros_ip_dns_static

A DNS A/AAAA/CNAME/MX/... static entry. Requires either name OR regexp.

## Example Usage

```terraform
data "routeros_ip_dns_static" "static_example" {
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
* `address` - (Required) Type: `string`. Address to return. Default: `127.0.0.1`.
* `address_list` - (Optional) Type: `string`.
* `cname` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `forward_to` - (Optional) Type: `string`.
* `match_subdomain` - (Optional) Type: `string`.
* `mx_exchange` - (Optional) Type: `string`.
* `mx_preference` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. FQDN matched against incoming queries. Default: `tf-acc-test.invalid`.
* `ns` - (Optional) Type: `string`.
* `regexp` - (Optional) Type: `string`.
* `srv_port` - (Optional) Type: `string`.
* `srv_priority` - (Optional) Type: `string`.
* `srv_target` - (Optional) Type: `string`.
* `srv_weight` - (Optional) Type: `string`.
* `text` - (Optional) Type: `string`.
* `ttl` - (Optional) Type: `string`.
* `type` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

