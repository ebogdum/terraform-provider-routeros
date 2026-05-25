---
subcategory: "Certificates"
page_title: "RouterOS: routeros_certificate_builtin"
description: |-
  System-generated certificates — read-only.
---

# Resource: routeros_certificate_builtin

System-generated certificates — read-only.

## Example Usage

```terraform
resource "routeros_certificate_builtin" "builtin_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # akid = "replace-me"
  # common_name = "replace-me"
  # country = "replace-me"
  # days_valid = 0
  # invalid_after = "replace-me"
  # invalid_before = "replace-me"
  # issuer = "replace-me"
  # key_type = "replace-me"
  # key_usage = []
  # locality = "replace-me"
  # organization = "replace-me"
  # serial_number = "replace-me"
  # skid = "replace-me"
  # state = "replace-me"
  # subject_alt_name = "replace-me"
  # unit = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `akid` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `common_name` - (Optional) Type: `string`.
* `country` - (Optional) Type: `string`.
* `days_valid` - (Optional) Type: `int`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `invalid_after` - (Optional) Type: `string`.
* `invalid_before` - (Optional) Type: `string`.
* `issuer` - (Optional) Type: `string`.
* `key_type` - (Optional) Type: `string`.
* `key_usage` - (Optional) Type: `list`.
* `locality` - (Optional) Type: `string`.
* `organization` - (Optional) Type: `string`.
* `serial_number` - (Optional) Type: `string`.
* `skid` - (Optional) Type: `string`.
* `state` - (Optional) Type: `string`.
* `subject_alt_name` - (Optional) Type: `string`.
* `unit` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_certificate_builtin.example '*3'

# Named router
terraform import routeros_certificate_builtin.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_certificate_builtin.example 'home/my-resource-name'
```
