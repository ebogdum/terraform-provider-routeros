---
subcategory: "Users"
page_title: "RouterOS: routeros_user_expire_password"
description: |-
  Needs a target user .id. Skipped.
---

# Resource: routeros_user_expire_password

Needs a target user .id. Skipped.

## Example Usage

```terraform
resource "routeros_user_expire_password" "expire_password_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

