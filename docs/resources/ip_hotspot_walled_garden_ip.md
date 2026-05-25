---
page_title: "RouterOS: routeros_ip_hotspot_walled_garden_ip"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_hotspot_walled_garden_ip

Manages the RouterOS `/ip/hotspot/walled-garden/ip` menu.

## Example Usage

```terraform
resource "routeros_ip_hotspot_walled_garden_ip" "ip_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "replace-me"
  # dst_address = "10.99.0.0/24"
  # dst_address_list = "my-list"
  # dst_port = "443"
  # protocol = "replace-me"
  # server = "replace-me"
  # src_address = "10.99.0.0/24"
  # src_address_list = "my-list"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `action` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dst_address` - (Optional) Type: `string`.
* `dst_address_list` - (Optional) Type: `string`.
* `dst_port` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `string`.
* `server` - (Optional) Type: `string`.
* `src_address` - (Optional) Type: `string`.
* `src_address_list` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_hotspot_walled_garden_ip.example '*3'

# Named router
terraform import routeros_ip_hotspot_walled_garden_ip.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_hotspot_walled_garden_ip.example 'home/my-resource-name'
```
