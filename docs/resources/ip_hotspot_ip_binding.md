---
subcategory: "Hotspot"
page_title: "RouterOS: routeros_ip_hotspot_ip_binding"
description: |-
  Hotspot ip-binding requires existing hotspot.
---

# Resource: routeros_ip_hotspot_ip_binding

Hotspot ip-binding requires existing hotspot.

## Example Usage

```terraform
resource "routeros_ip_hotspot_ip_binding" "ip_binding_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "10.99.0.0/24"
  # mac_address = "10.99.0.0/24"
  # server = "replace-me"
  # to_address = "10.99.0.0/24"
  # type = "regular"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `mac_address` - (Optional) Type: `string`.
* `server` - (Optional) Type: `string`.
* `to_address` - (Optional) Type: `string`.
* `type` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_hotspot_ip_binding.example '*3'

# Named router
terraform import routeros_ip_hotspot_ip_binding.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_hotspot_ip_binding.example 'home/my-resource-name'
```
