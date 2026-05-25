---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_cloud_back_to_home_user"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_cloud_back_to_home_user

Manages the RouterOS `/ip/cloud/back-to-home-user` menu.

## Example Usage

```terraform
resource "routeros_ip_cloud_back_to_home_user" "back_to_home_user_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # active = false
  # allow_lan = false
  # expires = "replace-me"
  # file_access_mode = ""
  # files = "replace-me"
  # name = "tf-example"
  # newe = "replace-me"
  # newfileman = "replace-me"
  # notnew = "replace-me"
  # oldfileman = "replace-me"
  # private_key = "REDACTED"
  # public_key = "REDACTED"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `active` - (Optional) Type: `bool`.
* `allow_lan` - (Optional) Type: `bool`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `expires` - (Optional) Type: `string`.
* `file_access_mode` - (Optional) Type: `enum(|disabled|read-only|full)`.
* `files` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `newe` - (Optional) Type: `string`.
* `newfileman` - (Optional) Type: `string`.
* `notnew` - (Optional) Type: `string`.
* `oldfileman` - (Optional) Type: `string`.
* `private_key` - (Optional) Type: `string`.
* `public_key` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `client_address` - Type: `string`.
* `client_config` - Type: `string`.
* `client_qr` - Type: `string`.
* `file_access_token` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_cloud_back_to_home_user.example '*3'

# Named router
terraform import routeros_ip_cloud_back_to_home_user.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_cloud_back_to_home_user.example 'home/my-resource-name'
```
