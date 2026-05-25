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
  # builtin = false
  # chain = ""
  # exp = "0"
  # reset_counters = "replace-me"
  # reset_counters_all = "replace-me"
  # set_exp = "0"
  # set_mark = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `builtin` - (Optional) Type: `bool`.
* `chain` - (Optional) Type: `enum(|forward|output)`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `exp` - (Optional) Type: `enum(0|1|2|3|4|5, ...)`.
* `reset_counters` - (Optional) Type: `string`.
* `reset_counters_all` - (Optional) Type: `string`.
* `set_exp` - (Optional) Type: `enum(0|1|2|3|4|5, ...)`.
* `set_mark` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `packets` - Type: `string`.

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
