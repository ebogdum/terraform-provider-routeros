---
subcategory: "Certificates"
page_title: "RouterOS: routeros_certificate_export_certificate"
description: |-
  Needs a cert .id. Skipped.
---

# Resource: routeros_certificate_export_certificate

Needs a cert .id. Skipped.

## Example Usage

```terraform
resource "routeros_certificate_export_certificate" "export_certificate_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

