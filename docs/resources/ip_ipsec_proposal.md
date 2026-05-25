---
subcategory: "IPsec"
page_title: "RouterOS: routeros_ip_ipsec_proposal"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_ipsec_proposal

Manages the RouterOS `/ip/ipsec/proposal` menu.

## Example Usage

```terraform
resource "routeros_ip_ipsec_proposal" "proposal_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # auth_algorithms = "md5"
  # enc_algorithms = []
  # lifetime = "1800"
  # name = "tf-example"
  # pfs_group = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `auth_algorithms` - (Optional) Type: `enum(md5|sha1|null|sha256|sha512)`. Default: `128`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `enc_algorithms` - (Optional) Type: `list`.
* `lifetime` - (Optional) Type: `duration`. Default: `1800`.
* `name` - (Optional) Type: `string`.
* `pfs_group` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_ipsec_proposal.example '*3'

# Named router
terraform import routeros_ip_ipsec_proposal.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_ipsec_proposal.example 'home/my-resource-name'
```
