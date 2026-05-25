---
subcategory: "Certificates"
page_title: "RouterOS: routeros_certificate_enable_ssl_certificate"
description: |-
  RouterOS resource.
---

# Resource: routeros_certificate_enable_ssl_certificate

Manages the RouterOS `/certificate/enable-ssl-certificate` menu.

## Example Usage

```terraform
resource "routeros_certificate_enable_ssl_certificate" "enable_ssl_certificate_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.

