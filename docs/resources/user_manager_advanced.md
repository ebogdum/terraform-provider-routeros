---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_manager_advanced"
description: |-
  RouterOS resource.
---

# Resource: routeros_user_manager_advanced

Manages the RouterOS `/user-manager/advanced` menu.

## Example Usage

```terraform
resource "routeros_user_manager_advanced" "advanced_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # paypal_allow = "replace-me"
  # paypal_currency = "replace-me"
  # paypal_password = "REDACTED"
  # paypal_signature = "replace-me"
  # paypal_use_sandbox = "replace-me"
  # paypal_user = "replace-me"
  # web_private_password = "REDACTED"
  # web_private_username = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `paypal_allow` - (Optional) Type: `string`.
* `paypal_currency` - (Optional) Type: `string`.
* `paypal_password` - (Optional) Type: `string`. **Sensitive.**
* `paypal_signature` - (Optional) Type: `string`.
* `paypal_use_sandbox` - (Optional) Type: `string`.
* `paypal_user` - (Optional) Type: `string`.
* `web_private_password` - (Optional) Type: `string`. **Sensitive.**
* `web_private_username` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_user_manager_advanced.this 'home'
```
