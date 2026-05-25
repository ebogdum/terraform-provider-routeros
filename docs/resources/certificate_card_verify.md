---
page_title: "RouterOS: routeros_certificate_card_verify"
description: |-
  Needs HSM card PIN. Skipped.
---

# Resource: routeros_certificate_card_verify

Needs HSM card PIN. Skipped.

## Example Usage

```terraform
resource "routeros_certificate_card_verify" "card_verify_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

