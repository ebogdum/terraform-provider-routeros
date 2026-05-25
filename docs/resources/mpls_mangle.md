---
subcategory: "MPLS"
page_title: "RouterOS: routeros_mpls_mangle"
description: |-
  MPLS mangle schema differs across ROS versions and the audit can't determine the correct argument set without an active LDP. Skipped.
---

# Resource: routeros_mpls_mangle

MPLS mangle schema differs across ROS versions and the audit can't determine the correct argument set without an active LDP. Skipped.

## Example Usage

```terraform
resource "routeros_mpls_mangle" "mangle_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # chain = ""
  # exp = "0"
  # set_exp = "0"
  # set_mark = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `chain` - (Optional) Type: `enum(|forward|output)`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `exp` - (Optional) Type: `enum(0|1|2|3|4|5, ...)`.
* `set_exp` - (Optional) Type: `enum(0|1|2|3|4|5, ...)`.
* `set_mark` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_mpls_mangle.example '*3'

# Named router
terraform import routeros_mpls_mangle.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_mpls_mangle.example 'home/my-resource-name'
```
