---
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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `import_ssh_key` - (Optional) Type: `string`.
* `key` - (Optional) Type: `string`.
* `newk` - (Optional) Type: `string`.
* `oldk` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `bits` - Type: `int`.
* `fingerprint` - Type: `string`.
* `info` - Type: `string`.
* `key_type` - Type: `enum(RSA|Ed25519|Ed25519-sk)`.
* `user` - Type: `string`.

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
