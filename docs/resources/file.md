---
subcategory: "Files"
page_title: "RouterOS: routeros_file"
description: |-
  Creating a file via REST requires writing the contents in a follow-up call; not in the acc-test fast path.
---

# Resource: routeros_file

Creating a file via REST requires writing the contents in a follow-up call; not in the acc-test fast path.

## Example Usage

```terraform
resource "routeros_file" "file_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # contents = "replace-me"
  # name = "example"
  # type = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `contents` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `type` - (Optional) Type: `int`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `last_modified` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_file.example '*3'

# Named router
terraform import routeros_file.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_file.example 'home/my-resource-name'
```
