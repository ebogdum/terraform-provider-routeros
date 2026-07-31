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
  # encr_algorithms = "4"
  # lifetime = "1800"
  # name = "tf-example"
  # pfs_group = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `auth_algorithms` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default` - (Read-only) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `enc_algorithms` - (Optional) Type: `list`.
* `encr_algorithms` - (Read-only) Type: `string`.
* `lifetime` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `pfs_group` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
