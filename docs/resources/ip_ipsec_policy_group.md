---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_ipsec_policy_group"
description: |-
  Mirrors RouterOS /ip/ipsec/policy/group.
---

# Resource: routeros_ip_ipsec_policy_group

Mirrors RouterOS `/ip/ipsec/policy/group`.

## Example Usage

```terraform
resource "routeros_ip_ipsec_policy_group" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # default = true
  # name = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `default` - (Read-only) Type: `bool`. RouterOS `default`.
* `name` - (Optional) Type: `string`. RouterOS `name`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
terraform import routeros_ip_ipsec_policy_group.example 'home::*3'
```
