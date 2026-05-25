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

  # Optional attributes (uncomment as needed):
  # download = "replace-me"
  # expired = false
  # flush = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `download` - (Optional) Type: `string`.
* `expired` - (Optional) Type: `bool`.
* `flush` - (Optional) Type: `string`.
* `url` - (Required) Type: `string`. Default: `http://invalid.example/crl`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `akid` - Type: `string`.
* `certificate` - Type: `string`.
* `dynamic` - Type: `bool`.
* `invalid` - Type: `bool`.
* `last_update` - Type: `string`.
* `next_update` - Type: `string`.
* `num` - Type: `int`.
* `revoked` - Type: `int`.
* `signature` - Type: `string`.

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
