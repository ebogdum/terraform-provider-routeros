---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_password"
description: |-
  Changes the logged-in user's password. Skipped from acc tests — would lock subsequent tests out.
---

# Resource: routeros_password

Changes the logged-in user's password. Skipped from acc tests — would lock subsequent tests out.

## Example Usage

```terraform
resource "routeros_password" "password_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

