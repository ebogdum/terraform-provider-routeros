---
subcategory: "Firewall"
page_title: "RouterOS: routeros_ip_firewall_layer7_protocol"
description: |-
  Layer7 patterns are unique by name; test re-runs trip "already have such name" unless previous artefact is cleaned. Skipped to keep the suite idempotent.
---

# Resource: routeros_ip_firewall_layer7_protocol

Layer7 patterns are unique by name; test re-runs trip "already have such name" unless previous artefact is cleaned. Skipped to keep the suite idempotent.

## Example Usage

```terraform
resource "routeros_ip_firewall_layer7_protocol" "layer7_protocol_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_firewall_layer7_protocol.example '*3'

# Named router
terraform import routeros_ip_firewall_layer7_protocol.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_firewall_layer7_protocol.example 'home/my-resource-name'
```
