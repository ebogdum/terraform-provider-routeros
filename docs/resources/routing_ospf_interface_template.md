---
subcategory: "Routing"
page_title: "RouterOS: routeros_routing_ospf_interface_template"
description: |-
  References an existing ospf area; auto-test can't synthesise.
---

# Resource: routeros_routing_ospf_interface_template

References an existing ospf area; auto-test can't synthesise.

## Example Usage

```terraform
resource "routeros_routing_ospf_interface_template" "interface_template_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # area = "replace-me"
  # auth_id = "replace-me"
  # auth_key = "REDACTED"
  # cost = 0
  # dead_interval = "1h"
  # hello_interval = "1h"
  # instance_id = 0
  # interfaces = "replace-me"
  # networks = "replace-me"
  # passive = "replace-me"
  # prefix_list = "replace-me"
  # priority = 0
  # retransmit_interval = "1h"
  # transmit_delay = 0
  # use_bfd = "replace-me"
  # vlink_neighbor_id = "replace-me"
  # vlink_transit_area = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `area` - (Optional) Type: `string`.
* `auth_id` - (Optional) Type: `string`.
* `auth_key` - (Optional) Type: `string`. **Sensitive.**
* `comment` - (Optional) Type: `string`. Free-form comment.
* `cost` - (Optional) Type: `int`.
* `dead_interval` - (Optional) Type: `duration`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `hello_interval` - (Optional) Type: `duration`.
* `instance_id` - (Optional) Type: `int`.
* `interfaces` - (Optional) Type: `string`.
* `networks` - (Optional) Type: `string`.
* `passive` - (Optional) Type: `string`.
* `prefix_list` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `int`.
* `retransmit_interval` - (Optional) Type: `duration`.
* `transmit_delay` - (Optional) Type: `int`.
* `use_bfd` - (Optional) Type: `string`.
* `vlink_neighbor_id` - (Optional) Type: `string`.
* `vlink_transit_area` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_ospf_interface_template.example '*3'

# Named router
terraform import routeros_routing_ospf_interface_template.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_ospf_interface_template.example 'home/my-resource-name'
```
