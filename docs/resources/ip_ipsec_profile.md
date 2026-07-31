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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `default` - (Read-only) Type: `bool`.
* `dh_group` - (Optional) Type: `list`.
* `dpd_interval` - (Optional) Type: `string`.
* `dpd_maximum_failures` - (Optional) Type: `int`.
* `enc_algorithm` - (Optional) Type: `list`.
* `encryption_algorithm` - (Read-only) Type: `string`.
* `hash_algorithm` - (Optional) Type: `string`.
* `hash_algorithms` - (Read-only) Type: `string`.
* `lifebytes` - (Optional) Type: `int`.
* `lifetime` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `nat_traversal` - (Optional) Type: `bool`.
* `ppk` - (Optional) Type: `string`.
* `prf_algorithm` - (Optional) Type: `string`. RouterOS `prf-algorithm`.
* `prf_algorithms` - (Read-only) Type: `string`.
* `proposal_check` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
