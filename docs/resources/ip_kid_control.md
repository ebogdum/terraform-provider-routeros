---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_kid_control"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_kid_control

Manages the RouterOS `/ip/kid-control` menu.

## Example Usage

```terraform
resource "routeros_ip_kid_control" "kid_control_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # name = "example"
  # rate_limit = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `name` - (Optional) Type: `string`. Name of the Kid's profile.
* `rate_limit` - (Optional) Type: `string`. The maximum available data rate for flow.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_kid_control.example '*3'

# Named router
terraform import routeros_ip_kid_control.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_kid_control.example 'home/my-resource-name'
```
