---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_tftp"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_tftp

Manages the RouterOS `/ip/tftp` menu.

## Example Usage

```terraform
resource "routeros_ip_tftp" "tftp_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # allow = true
  # ip_addresses = "replace-me"
  # read_only = true
  # real_filename = "replace-me"
  # req_filename = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `allow` - (Optional) Type: `bool`.
* `allow_overwrite` - (Optional) Type: `string`. RouterOS `allow-overwrite`.
* `allow_rollover` - (Optional) Type: `string`. RouterOS `allow-rollover`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `hits` - (Read-only) Type: `int`.
* `ip_addresses` - (Optional) Type: `string`.
* `read_only` - (Optional) Type: `bool`.
* `reading_window_size` - (Optional) Type: `string`. RouterOS `reading-window-size`.
* `real_filename` - (Optional) Type: `string`.
* `req_filename` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_tftp.example '*3'

# Named router
terraform import routeros_ip_tftp.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_tftp.example 'home/my-resource-name'
```
