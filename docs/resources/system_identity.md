---
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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `name` - (Optional) Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_identity.this 'home'
```
