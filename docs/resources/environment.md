---
subcategory: "System & misc"
page_title: "RouterOS: routeros_environment"
description: |-
  Console environment variables -- read-only.
---

# Resource: routeros_environment

Console environment variables -- read-only.

## Example Usage

```terraform
resource "routeros_environment" "environment_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_environment.example '*3'

# Named router
terraform import routeros_environment.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_environment.example 'home/my-resource-name'
```
