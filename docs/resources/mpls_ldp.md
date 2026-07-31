---
subcategory: "MPLS"
page_title: "RouterOS: routeros_mpls_ldp"
description: |-
  LDP allows one active instance per VRF; if one already exists this fails.
---

# Resource: routeros_mpls_ldp

LDP allows one active instance per VRF; if one already exists this fails.

## Example Usage

```terraform
resource "routeros_mpls_ldp" "ldp_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `afi` - (Optional) Type: `string`. RouterOS `afi`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `distribute_for_default` - (Optional) Type: `string`. RouterOS `distribute-for-default`.
* `hop_limit` - (Optional) Type: `string`. RouterOS `hop-limit`.
* `loop_detect` - (Optional) Type: `string`. RouterOS `loop-detect`.
* `lsr_id` - (Optional) Type: `string`. RouterOS `lsr-id`.
* `path_vector_limit` - (Optional) Type: `string`. RouterOS `path-vector-limit`.
* `preferred_afi` - (Optional) Type: `string`. RouterOS `preferred-afi`.
* `transport_addresses` - (Optional) Type: `string`. RouterOS `transport-addresses`.
* `use_explicit_null` - (Optional) Type: `string`. RouterOS `use-explicit-null`.
* `vrf` - (Optional) Type: `string`. RouterOS `vrf`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_mpls_ldp.example '*3'

# Named router
terraform import routeros_mpls_ldp.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_mpls_ldp.example 'home/my-resource-name'
```
