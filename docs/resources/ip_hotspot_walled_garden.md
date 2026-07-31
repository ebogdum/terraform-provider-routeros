---
subcategory: "Hotspot"
page_title: "RouterOS: routeros_ip_hotspot_walled_garden"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_hotspot_walled_garden

Manages the RouterOS `/ip/hotspot/walled-garden` menu.

## Example Usage

```terraform
resource "routeros_ip_hotspot_walled_garden" "walled_garden_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "replace-me"
  # dst_port = "443"
  # path = "replace-me"
  # server = "replace-me"
  # src_address = "10.99.0.0/24"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `action` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dst_host` - (Optional) Type: `string`. RouterOS `dst-host`.
* `dst_port` - (Optional) Type: `string`.
* `method` - (Optional) Type: `string`. RouterOS `method`.
* `path` - (Optional) Type: `string`.
* `server` - (Optional) Type: `string`.
* `src_address` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_hotspot_walled_garden.example '*3'

# Named router
terraform import routeros_ip_hotspot_walled_garden.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_hotspot_walled_garden.example 'home/my-resource-name'
```
