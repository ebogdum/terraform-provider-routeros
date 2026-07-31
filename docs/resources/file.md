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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `backup` - (Read-only) Type: `string`.
* `basename` - (Read-only) Type: `string`.
* `container` - (Read-only) Type: `int`.
* `contents` - (Optional) Type: `string`.
* `directory` - (Read-only) Type: `string`.
* `family` - (Read-only) Type: `int`.
* `file_name` - (Read-only) Type: `string`.
* `file_share_url` - (Read-only) Type: `string`.
* `hasvpn` - (Read-only) Type: `string`.
* `invalid` - (Read-only) Type: `string`.
* `invalidfile` - (Read-only) Type: `string`.
* `last_modified` - (Read-only) Type: `string`.
* `name` - (Optional) Type: `string`.
* `nondir` - (Read-only) Type: `string`.
* `restore` - (Read-only) Type: `string`.
* `share` - (Read-only) Type: `string`.
* `shared` - (Read-only) Type: `bool`.
* `size` - (Read-only) Type: `string`.
* `type` - (Read-only) Type: `string`.
* `unshare` - (Read-only) Type: `string`.
* `valid` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
