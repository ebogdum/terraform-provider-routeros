---
page_title: "RouterOS: routeros_certificate_scep_server"
description: |-
  SCEP server references an existing CA cert. Skipped from automated acc tests.
---

# Resource: routeros_certificate_scep_server

SCEP server references an existing CA cert. Skipped from automated acc tests.

## Example Usage

```terraform
resource "routeros_certificate_scep_server" "scep_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # days_valid = 1
  # path = "replace-me"
  # request_lifetime = "3600"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `days_valid` - (Optional) Type: `int`. Default: `1`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `path` - (Optional) Type: `string`.
* `request_lifetime` - (Optional) Type: `duration`. Default: `3600`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_certificate_scep_server.example '*3'

# Named router
terraform import routeros_certificate_scep_server.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_certificate_scep_server.example 'home/my-resource-name'
```
