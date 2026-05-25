---
subcategory: "Firewall"
page_title: "RouterOS: routeros_ip_firewall_layer7_protocol"
description: |-
  Layer7 patterns are unique by name; test re-runs trip "already have such name" unless previous artefact is cleaned. Skipped to keep the suite idempotent.
---

# Data Source: routeros_ip_firewall_layer7_protocol

Layer7 patterns are unique by name; test re-runs trip "already have such name" unless previous artefact is cleaned. Skipped to keep the suite idempotent.

## Example Usage

```terraform
data "routeros_ip_firewall_layer7_protocol" "layer7_protocol_example" {
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
* `comment` - (Optional) Type: `string`. Free-form comment.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

