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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `builtin` - (Read-only) Type: `bool`.
* `chain` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `exp` - (Optional) Type: `string`.
* `packets` - (Read-only) Type: `string`.
* `reset_counters` - (Read-only) Type: `string`.
* `reset_counters_all` - (Read-only) Type: `string`.
* `set_exp` - (Optional) Type: `string`.
* `set_mark` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
