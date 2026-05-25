---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_manager_attribute"
description: |-
  Requires user-manager package
---

# Data Source: routeros_user_manager_attribute

Requires user-manager package

## Example Usage

```terraform
data "routeros_user_manager_attribute" "attribute_example" {
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
* `name` - (Optional) Type: `string`. Name of the attribute.
* `packet_types` - (Optional) Type: `string`. access-accept - use this attribute in RADIUS Access-Accept messages access-challenge - use this attribute in RADIUS Access-Challenge messages.
* `type_id` - (Optional) Type: `string`. Attribute identification number from the specific vendor's attribute database.
* `value_type` - (Optional) Type: `string`. hex ip-address - IPv4 or IPv6 IP address ip6-prefix - IPv6 prefix macro string uint32.
* `vendor_id` - (Optional) Type: `string`. IANA allocated a specific enterprise identification number.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

