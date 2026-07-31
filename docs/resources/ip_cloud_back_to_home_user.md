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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `active` - (Read-only) Type: `bool`.
* `allow_lan` - (Optional) Type: `bool`.
* `client_address` - (Read-only) Type: `string`.
* `client_config` - (Read-only) Type: `string`.
* `client_qr` - (Read-only) Type: `string`.
* `comment` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `expires` - (Optional) Type: `string`.
* `file_access` - (Optional) Type: `string`. RouterOS `file-access`.
* `file_access_mode` - (Read-only) Type: `string`.
* `file_access_path` - (Optional) Type: `string`. RouterOS `file-access-path`.
* `file_access_token` - (Read-only) Type: `string`.
* `files` - (Read-only) Type: `string`.
* `name` - (Optional) Type: `string`.
* `newe` - (Read-only) Type: `string`.
* `newfileman` - (Read-only) Type: `string`.
* `notnew` - (Read-only) Type: `string`.
* `oldfileman` - (Read-only) Type: `string`.
* `private_key` - (Optional) Type: `string`. **Sensitive.**
* `public_key` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
