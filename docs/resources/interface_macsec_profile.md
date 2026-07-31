---
subcategory: "Interface"
page_title: "RouterOS: routeros_interface_macsec_profile"
description: |-
  Mirrors RouterOS /interface/macsec/profile.
---

# Resource: routeros_interface_macsec_profile

Mirrors RouterOS `/interface/macsec/profile`.

## Example Usage

```terraform
resource "routeros_interface_macsec_profile" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # default_name = "replace-me"
  # name = "replace-me"
  # server_priority = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `ciphers` - (Optional) Type: `string`. RouterOS `ciphers`.
* `default_name` - (Read-only) Type: `string`. RouterOS `default-name`.
* `name` - (Optional) Type: `string`. RouterOS `name`.
* `server_priority` - (Optional) Type: `int`. RouterOS `server-priority`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
terraform import routeros_interface_macsec_profile.example 'home::*3'
```
