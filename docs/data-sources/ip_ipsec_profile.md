---
subcategory: "IPsec"
page_title: "RouterOS: routeros_ip_ipsec_profile"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_ipsec_profile

Manages the RouterOS `/ip/ipsec/profile` menu.

## Example Usage

```terraform
data "routeros_ip_ipsec_profile" "profile_example" {
  # router   = "my-router"  # omit for the default router
  # filter   = { name = "some-name" }
  # proplist = ["name", "address"]
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `proplist` - (Optional) List of property names to project; smaller payload.
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

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.

