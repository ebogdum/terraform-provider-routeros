---
page_title: "RouterOS: routeros_ip_socksify"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_socksify

Manages the RouterOS `/ip/socksify` menu.

## Example Usage

```terraform
resource "routeros_ip_socksify" "socksify_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # port = "443"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `name` - (Required) Type: `string`. Default: `tf-acc-socksify`.
* `port` - (Optional) Type: `int`. Default: `1080`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_socksify.example '*3'

# Named router
terraform import routeros_ip_socksify.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_socksify.example 'home/my-resource-name'
```
