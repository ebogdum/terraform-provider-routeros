---
subcategory: "Certificates"
page_title: "RouterOS: routeros_certificate_crl"
description: |-
  RouterOS resource.
---

# Resource: routeros_certificate_crl

Manages the RouterOS `/certificate/crl` menu.

## Example Usage

```terraform
resource "routeros_certificate_crl" "crl_example" {
  # router = "my-router"  # which router to target; omit for the default
  url = "https://example.com"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `url` - (Required) Type: `string`. Default: `http://invalid.example/crl`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_certificate_crl.example '*3'

# Named router
terraform import routeros_certificate_crl.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_certificate_crl.example 'home/my-resource-name'
```
