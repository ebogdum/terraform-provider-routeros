---
subcategory: "System"
page_title: "RouterOS: routeros_system_identity"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_identity

Manages the RouterOS `/system/identity` menu.

## Example Usage

```terraform
resource "routeros_system_identity" "identity_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # name = "tf-example"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `name` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_identity.this 'home'
```
