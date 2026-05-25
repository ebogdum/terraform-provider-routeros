---
subcategory: "System"
page_title: "RouterOS: routeros_system_backup"
description: |-
  Action-only menu (save/load); not CRUD
---

# Resource: routeros_system_backup

Action-only menu (save/load); not CRUD

## Example Usage

```terraform
resource "routeros_system_backup" "backup_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # dont_encrypt = "replace-me"
  # encryption = "replace-me"
  # name = "example"
  # password = "REDACTED"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `dont_encrypt` - (Optional) Type: `string`. Disable backup file encryption. Note that since RouterOS v6.43 without a provided   password,   the backup file is unencrypted.
* `encryption` - (Optional) Type: `string`. The encryption algorithm to use for encrypting the backup file. Note that is not considered a secure encryption method and is only available for compatibility reasons with older RouterOS versions.
* `name` - (Optional) Type: `string`. The filename for the backup file.
* `password` - (Optional) Type: `string`. Password for the encrypted backup file. Note that since RouterOS v6.43 without a provided   password,   the backup file is unencrypted. **Sensitive.**

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_system_backup.example '*3'

# Named router
terraform import routeros_system_backup.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_system_backup.example 'home/my-resource-name'
```
