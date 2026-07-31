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
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # directory = "replace-me"
  # invalid_users = "replace-me"
  # newfileman = "replace-me"
  # old_directory = "replace-me"
  # oldfileman = "replace-me"
  # read_only = false
  # require_encryption = false
  # valid_users = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default` - (Read-only) Type: `bool`.
* `directory` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `dynamic` - (Read-only) Type: `bool`.
* `invalid_users` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `newfileman` - (Read-only) Type: `string`.
* `old_directory` - (Read-only) Type: `string`.
* `oldfileman` - (Read-only) Type: `string`.
* `read_only` - (Optional) Type: `bool`.
* `require_encryption` - (Optional) Type: `bool`.
* `valid_users` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
