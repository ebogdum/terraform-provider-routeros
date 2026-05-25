---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_pimsm_interface_template"
description: |-
  Discovered; needs pimsm instance
---

# Data Source: routeros_routing_pimsm_interface_template

Discovered; needs pimsm instance

## Example Usage

```terraform
data "routeros_routing_pimsm_interface_template" "interface_template_example" {
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
* `disabled` - (Optional) Type: `bool`.
* `hello_delay` - (Optional) Type: `string`.
* `hello_period` - (Optional) Type: `string`.
* `instance` - (Optional) Type: `string`.
* `interfaces` - (Optional) Type: `string`.
* `join_prune_period` - (Optional) Type: `string`.
* `join_tracking_support` - (Optional) Type: `string`.
* `override_interval` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `string`.
* `propagation_delay` - (Optional) Type: `string`.
* `source_addresses` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

