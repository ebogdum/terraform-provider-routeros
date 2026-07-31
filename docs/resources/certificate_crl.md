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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `akid` - (Read-only) Type: `string`.
* `certificate` - (Read-only) Type: `string`.
* `download` - (Read-only) Type: `string`.
* `dynamic` - (Read-only) Type: `bool`.
* `expired` - (Read-only) Type: `bool`.
* `flush` - (Read-only) Type: `string`.
* `invalid` - (Read-only) Type: `bool`.
* `last_update` - (Read-only) Type: `string`.
* `next_update` - (Read-only) Type: `string`.
* `num` - (Read-only) Type: `int`.
* `revoked` - (Read-only) Type: `int`.
* `signature` - (Read-only) Type: `string`.
* `url` - (Required) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
