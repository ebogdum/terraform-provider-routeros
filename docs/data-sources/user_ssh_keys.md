---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_ssh_keys"
description: |-
  SSH key import requires a real key file already uploaded to /file. Skipped from automated acc tests.
---

# Data Source: routeros_user_ssh_keys

SSH key import requires a real key file already uploaded to /file. Skipped from automated acc tests.

## Example Usage

```terraform
data "routeros_user_ssh_keys" "ssh_keys_example" {
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
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `import_ssh_key` - (Optional) Type: `string`.
* `key` - (Optional) Type: `string`.
* `newk` - (Optional) Type: `string`.
* `oldk` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `bits` - Type: `int`.
* `fingerprint` - Type: `string`.
* `info` - Type: `string`.
* `key_type` - Type: `enum(rsa|ed25519|ed25519-sk)`.
* `user` - Type: `string`.

