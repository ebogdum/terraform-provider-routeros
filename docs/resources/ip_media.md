---
page_title: "RouterOS: routeros_ip_media"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_media

Manages the RouterOS `/ip/media` menu.

## Example Usage

```terraform
resource "routeros_ip_media" "media_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"

  disabled = false

  # Optional attributes (uncomment as needed):
  # path = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.
* `path` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_media.example '*3'

# Named router
terraform import routeros_ip_media.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_media.example 'home/my-resource-name'
```
