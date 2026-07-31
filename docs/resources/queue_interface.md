---
subcategory: "Queue"
page_title: "RouterOS: routeros_queue_interface"
description: |-
  Mirrors RouterOS /queue/interface.
---

# Resource: routeros_queue_interface

Mirrors RouterOS `/queue/interface`.

## Example Usage

```terraform
resource "routeros_queue_interface" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # interface = "replace-me"
  # queue = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `interface` - (Optional) Type: `string`. RouterOS `interface`.
* `queue` - (Optional) Type: `string`. RouterOS `queue`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
terraform import routeros_queue_interface.example 'home::*3'
```
