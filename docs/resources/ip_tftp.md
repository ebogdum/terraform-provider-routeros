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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `allow` - (Optional) Type: `bool`. Default: `1`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ip_addresses` - (Optional) Type: `string`.
* `read_only` - (Optional) Type: `bool`. Default: `1`.
* `real_filename` - (Optional) Type: `string`.
* `req_filename` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `hits` - Type: `int`.

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
