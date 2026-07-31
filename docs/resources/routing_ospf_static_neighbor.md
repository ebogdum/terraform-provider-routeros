---
subcategory: "OSPF"
page_title: "RouterOS: routeros_routing_ospf_static_neighbor"
description: |-
  References an existing ospf area; auto-test can't synthesise.
---

# Resource: routeros_routing_ospf_static_neighbor

References an existing ospf area; auto-test can't synthesise.

## Example Usage

```terraform
resource "routeros_routing_ospf_static_neighbor" "static_neighbor_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # area = "replace-me"
  # instance_id = 0
  # poll_interval = "1h"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Optional) Type: `string`.
* `area` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `instance_id` - (Optional) Type: `int`.
* `invalid` - (Read-only) Type: `bool`.
* `poll_interval` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_ospf_static_neighbor.example '*3'

# Named router
terraform import routeros_routing_ospf_static_neighbor.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_ospf_static_neighbor.example 'home/my-resource-name'
```
