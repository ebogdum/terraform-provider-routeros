---
subcategory: "Certificates"
page_title: "RouterOS: routeros_certificate_add_scep"
description: |-
  SCEP enrolment needs a template and CA; cannot be auto-tested without infrastructure.
---

# Resource: routeros_certificate_add_scep

SCEP enrolment needs a template and CA; cannot be auto-tested without infrastructure.

## Example Usage

```terraform
resource "routeros_certificate_add_scep" "add_scep_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

