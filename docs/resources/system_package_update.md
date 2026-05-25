---
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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `channel` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `installed_version` - Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_package_update.this 'home'
```
