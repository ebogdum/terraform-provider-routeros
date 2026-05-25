---
page_title: "RouterOS: routeros_certificate_scep_renew"
description: |-
  Needs an existing SCEP cert .id. Skipped.
---

# Resource: routeros_certificate_scep_renew

Needs an existing SCEP cert .id. Skipped.

## Example Usage

```terraform
resource "routeros_certificate_scep_renew" "scep_renew_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

