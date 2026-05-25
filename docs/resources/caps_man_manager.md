---
subcategory: "Caps-man"
page_title: "RouterOS: routeros_caps_man_manager"
description: |-
  RouterOS resource.
---

# Resource: routeros_caps_man_manager

Manages the RouterOS `/caps-man/manager` menu.

## Example Usage

```terraform
resource "routeros_caps_man_manager" "manager_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # ca_certificate = "replace-me"
  # certificate = "replace-me"
  # enabled = false
  # package_path = "replace-me"
  # require_peer_certificate = false
  # upgrade_policy = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `ca_certificate` - (Optional) Type: `string`.
* `certificate` - (Optional) Type: `string`.
* `enabled` - (Optional) Type: `bool`.
* `package_path` - (Optional) Type: `string`.
* `require_peer_certificate` - (Optional) Type: `bool`.
* `upgrade_policy` - (Optional) Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_caps_man_manager.this 'home'
```
