---
page_title: "RouterOS: routeros_certificate_card_reinstall"
description: |-
  Needs HSM card PIN. Skipped.
---

# Resource: routeros_certificate_card_reinstall

Needs HSM card PIN. Skipped.

## Example Usage

```terraform
resource "routeros_certificate_card_reinstall" "card_reinstall_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

