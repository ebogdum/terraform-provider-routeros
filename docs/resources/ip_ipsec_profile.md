---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_ipsec_profile"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_ipsec_profile

Manages the RouterOS `/ip/ipsec/profile` menu.

## Example Usage

```terraform
resource "routeros_ip_ipsec_profile" "profile_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"

  # Optional attributes (uncomment as needed):
  # dh_group = ["16388"]
  # dpd_interval = "disable DPD"
  # dpd_maximum_failures = 4
  # enc_algorithm = []
  # hash_algorithm = "replace-me"
  # lifebytes = 0
  # lifetime = "86400"
  # nat_traversal = true
  # ppk = "No"
  # proposal_check = "obey"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `dh_group` - (Optional) Type: `list`. Default: `16388`.
* `dpd_interval` - (Optional) Type: `enum(disable DPD)`. Default: `8`.
* `dpd_maximum_failures` - (Optional) Type: `int`. Default: `4`.
* `enc_algorithm` - (Optional) Type: `list`.
* `hash_algorithm` - (Optional) Type: `string`.
* `lifebytes` - (Optional) Type: `int`.
* `lifetime` - (Optional) Type: `duration`. Default: `86400`.
* `name` - (Required) Type: `string`. Default: `tf_acc_ipsecprof`.
* `nat_traversal` - (Optional) Type: `bool`. Default: `1`.
* `ppk` - (Optional) Type: `enum(No|PSK|QKD|PSK IKE Initial)`.
* `proposal_check` - (Optional) Type: `enum(|obey|strict|claim|exact)`. Default: `1`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_ipsec_profile.example '*3'

# Named router
terraform import routeros_ip_ipsec_profile.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_ipsec_profile.example 'home/my-resource-name'
```
