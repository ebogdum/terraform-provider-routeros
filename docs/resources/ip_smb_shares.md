---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_smb_shares"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_smb_shares

Manages the RouterOS `/ip/smb/shares` menu.

## Example Usage

```terraform
resource "routeros_ip_smb_shares" "shares_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # directory = "replace-me"
  # invalid_users = "replace-me"
  # read_only = false
  # require_encryption = false
  # valid_users = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `directory` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `invalid_users` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_smbshare`.
* `read_only` - (Optional) Type: `bool`.
* `require_encryption` - (Optional) Type: `bool`.
* `valid_users` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.
* `dynamic` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_smb_shares.example '*3'

# Named router
terraform import routeros_ip_smb_shares.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_smb_shares.example 'home/my-resource-name'
```
