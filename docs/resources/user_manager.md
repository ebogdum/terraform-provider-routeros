---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_manager"
description: |-
  RouterOS resource.
---

# Resource: routeros_user_manager

Manages the RouterOS `/user-manager` menu.

## Example Usage

```terraform
resource "routeros_user_manager" "user_manager_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # accounting_port = "443"
  # authentication_port = "443"
  # certificate = "replace-me"
  # enabled = "replace-me"
  # radsec_certificate = "replace-me"
  # require_message_auth = "replace-me"
  # use_profiles = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `accounting_port` - (Optional) Type: `string`.
* `authentication_port` - (Optional) Type: `string`.
* `certificate` - (Optional) Type: `string`.
* `enabled` - (Optional) Type: `string`.
* `radsec_certificate` - (Optional) Type: `string`.
* `require_message_auth` - (Optional) Type: `string`.
* `use_profiles` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_user_manager.this 'home'
```
