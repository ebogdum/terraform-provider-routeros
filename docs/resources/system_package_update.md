---
subcategory: "System"
page_title: "RouterOS: routeros_system_package_update"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_package_update

Manages the RouterOS `/system/package/update` menu.

## Example Usage

```terraform
resource "routeros_system_package_update" "update_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # channel = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `channel` - (Optional) Type: `string`.
* `check_certificate` - (Optional) Type: `string`. RouterOS `check-certificate`.
* `installed_version` - (Optional) Type: `string`.
* `ip_version` - (Optional) Type: `string`. RouterOS `ip-version`.
* `mode` - (Optional) Type: `string`. RouterOS `mode`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_package_update.this 'home'
```
