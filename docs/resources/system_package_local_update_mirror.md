---
subcategory: "System"
page_title: "RouterOS: routeros_system_package_local_update_mirror"
description: |-
  Mirrors RouterOS /system/package/local-update/mirror.
---

# Resource: routeros_system_package_local_update_mirror

Mirrors RouterOS `/system/package/local-update/mirror`.

## Example Usage

```terraform
resource "routeros_system_package_local_update_mirror" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # check_interval = "replace-me"
  # enabled = true
  # primary_server = "replace-me"
  # secondary_server = "replace-me"
  # user = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `check_interval` - (Optional) Type: `string`. RouterOS `check-interval`.
* `enabled` - (Optional) Type: `bool`. RouterOS `enabled`.
* `primary_server` - (Optional) Type: `string`. RouterOS `primary-server`.
* `secondary_server` - (Optional) Type: `string`. RouterOS `secondary-server`.
* `user` - (Optional) Type: `string`. RouterOS `user`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_package_local_update_mirror.this 'home'
```
