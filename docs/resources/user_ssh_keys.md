---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_ssh_keys"
description: |-
  SSH key import requires a real key file already uploaded to /file. Skipped from automated acc tests.
---

# Resource: routeros_user_ssh_keys

SSH key import requires a real key file already uploaded to /file. Skipped from automated acc tests.

## Example Usage

```terraform
resource "routeros_user_ssh_keys" "ssh_keys_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # import_ssh_key = "REDACTED"
  # key = "replace-me"
  # newk = "replace-me"
  # oldk = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `bits` - (Read-only) Type: `int`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `fingerprint` - (Read-only) Type: `string`.
* `import_ssh_key` - (Optional) Type: `string`.
* `info` - (Read-only) Type: `string`.
* `key` - (Optional) Type: `string`. **Sensitive.**
* `key_type` - (Read-only) Type: `string`.
* `newk` - (Optional) Type: `string`.
* `oldk` - (Optional) Type: `string`.
* `user` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_user_ssh_keys.example '*3'

# Named router
terraform import routeros_user_ssh_keys.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_user_ssh_keys.example 'home/my-resource-name'
```
