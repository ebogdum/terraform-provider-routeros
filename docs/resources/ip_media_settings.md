---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_media_settings"
description: |-
  Mirrors RouterOS /ip/media/settings.
---

# Resource: routeros_ip_media_settings

Mirrors RouterOS `/ip/media/settings`.

## Example Usage

```terraform
resource "routeros_ip_media_settings" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # thumbnails = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `thumbnails` - (Optional) Type: `string`. RouterOS `thumbnails`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_media_settings.this 'home'
```
