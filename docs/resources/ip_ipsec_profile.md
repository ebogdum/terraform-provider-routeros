---
subcategory: "IPsec"
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
  name = "tf-example"

  # Optional attributes (uncomment as needed):
  # dh_group = ["16388"]
  # dpd_interval = "disable-dpd"
  # dpd_maximum_failures = 4
  # enc_algorithm = []
  # encryption_algorithm = "12"
  # hash_algorithm = "replace-me"
  # hash_algorithms = "sha256"
  # lifebytes = 0
  # lifetime = "86400"
  # nat_traversal = true
  # ppk = "no"
  # prf_algorithms = "auto"
  # proposal_check = "obey"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `dh_group` - (Optional) Type: `list`. Default: `16388`.
* `dpd_interval` - (Optional) Type: `enum(disable-dpd)`. Default: `8`.
* `dpd_maximum_failures` - (Optional) Type: `int`. Default: `4`.
* `enc_algorithm` - (Optional) Type: `list`.
* `encryption_algorithm` - (Optional) Type: `string`. Default: `12`.
* `hash_algorithm` - (Optional) Type: `string`.
* `hash_algorithms` - (Optional) Type: `enum(md5|sha1|sha256|sha384|sha512)`. Default: `2`.
* `lifebytes` - (Optional) Type: `int`.
* `lifetime` - (Optional) Type: `duration`. Default: `86400`.
* `name` - (Required) Type: `string`. Default: `tf_acc_ipsecprof`.
* `nat_traversal` - (Optional) Type: `bool`. Default: `1`.
* `ppk` - (Optional) Type: `enum(no|psk|qkd|psk-ike-initial)`.
* `prf_algorithms` - (Optional) Type: `enum(auto|sha1|sha256|sha384|sha512)`.
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
