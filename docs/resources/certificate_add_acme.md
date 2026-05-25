---
page_title: "RouterOS: routeros_certificate_add_acme"
description: |-
  ACME enrolment needs a real domain name and reachable ACME server.
---

# Resource: routeros_certificate_add_acme

ACME enrolment needs a real domain name and reachable ACME server.

## Example Usage

```terraform
resource "routeros_certificate_add_acme" "add_acme_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

