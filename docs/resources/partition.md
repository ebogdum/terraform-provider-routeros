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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `activate` - (Optional) Type: `string`.
* `active` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`.
* `fallback_to` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `running` - (Optional) Type: `bool`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `size` - Type: `int`.
* `version` - Type: `string`.

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
