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
  # backup = "replace-me"
  # basename = "replace-me"
  # container = 0
  # contents = "replace-me"
  # directory = "replace-me"
  # file_name = "replace-me"
  # hasvpn = "replace-me"
  # invalid = "replace-me"
  # invalidfile = "replace-me"
  # name = "tf-example"
  # nondir = "replace-me"
  # restore = "replace-me"
  # share = "replace-me"
  # shared = false
  # type = 0
  # unshare = "replace-me"
  # valid = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `backup` - (Optional) Type: `string`.
* `basename` - (Optional) Type: `string`.
* `container` - (Optional) Type: `int`. Default: `4.294967295e+09`.
* `contents` - (Optional) Type: `string`.
* `directory` - (Optional) Type: `string`.
* `file_name` - (Optional) Type: `string`.
* `hasvpn` - (Optional) Type: `string`.
* `invalid` - (Optional) Type: `string`.
* `invalidfile` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `nondir` - (Optional) Type: `string`.
* `restore` - (Optional) Type: `string`.
* `share` - (Optional) Type: `string`.
* `shared` - (Optional) Type: `bool`.
* `type` - (Optional) Type: `int`.
* `unshare` - (Optional) Type: `string`.
* `valid` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `family` - Type: `int`.
* `file_share_url` - Type: `string`.
* `last_modified` - Type: `string`.
* `size` - Type: `string`.

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
