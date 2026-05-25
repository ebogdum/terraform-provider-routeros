---
subcategory: "System & misc"
page_title: "RouterOS: routeros_caps_man_registration_table"
description: |-
  CAPsMAN registration table; read-only.
---

# Resource: routeros_caps_man_registration_table

CAPsMAN registration table; read-only.

## Example Usage

```terraform
resource "routeros_caps_man_registration_table" "registration_table_example" {
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
terraform import routeros_caps_man_registration_table.example '*3'

# Named router
terraform import routeros_caps_man_registration_table.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_caps_man_registration_table.example 'home/my-resource-name'
```
