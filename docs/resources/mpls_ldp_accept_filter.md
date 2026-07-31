---
subcategory: "MPLS"
page_title: "RouterOS: routeros_mpls_ldp_accept_filter"
description: |-
  RouterOS resource.
---

# Resource: routeros_mpls_ldp_accept_filter

Manages the RouterOS `/mpls/ldp/accept-filter` menu.

## Example Usage

```terraform
resource "routeros_mpls_ldp_accept_filter" "accept_filter_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # prefix = "replace-me"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `accept` - (Optional) Type: `string`. RouterOS `accept`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `neighbor` - (Optional) Type: `string`. RouterOS `neighbor`.
* `prefix` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_mpls_ldp_accept_filter.example '*3'

# Named router
terraform import routeros_mpls_ldp_accept_filter.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_mpls_ldp_accept_filter.example 'home/my-resource-name'
```
