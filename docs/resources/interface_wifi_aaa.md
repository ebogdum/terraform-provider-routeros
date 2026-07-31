---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_aaa"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_aaa

Manages the RouterOS `/interface/wifi/aaa` menu.

## Example Usage

```terraform
resource "routeros_interface_wifi_aaa" "aaa_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # called_format = "replace-me"
  # calling_format = "replace-me"
  # interim_update = "replace-me"
  # mac_caching = "replace-me"
  # name = "tf-example"
  # nas_identifier = "replace-me"
  # password_format = "replace-me"
  # username_format = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `called_format` - (Optional) Type: `string`.
* `calling_format` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `interim_update` - (Optional) Type: `string`.
* `mac_caching` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `nas_identifier` - (Optional) Type: `string`.
* `password_format` - (Optional) Type: `string`.
* `username_format` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wifi_aaa.example '*3'

# Named router
terraform import routeros_interface_wifi_aaa.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wifi_aaa.example 'home/my-resource-name'
```
