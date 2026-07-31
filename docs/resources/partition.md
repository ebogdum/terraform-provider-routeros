---
subcategory: "Storage"
page_title: "RouterOS: routeros_partition"
description: |-
  RouterOS resource.
---

# Resource: routeros_partition

Manages the RouterOS `/partition` menu.

## Example Usage

```terraform
resource "routeros_partition" "partition_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # activate = "replace-me"
  # active = false
  # fallback_to = "replace-me"
  # name = "tf-example"
  # running = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `activate` - (Read-only) Type: `string`.
* `active` - (Read-only) Type: `bool`.
* `comment` - (Optional) Type: `string`.
* `fallback_to` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `running` - (Read-only) Type: `bool`.
* `size` - (Read-only) Type: `int`.
* `version` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_partition.example '*3'

# Named router
terraform import routeros_partition.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_partition.example 'home/my-resource-name'
```
