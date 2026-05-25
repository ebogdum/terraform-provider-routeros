---
page_title: "RouterOS: routeros_system_backup"
description: |-
  Action-only menu (save/load); not CRUD
---

# Data Source: routeros_system_backup

Action-only menu (save/load); not CRUD

## Example Usage

```terraform
data "routeros_system_backup" "backup_example" {
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
* `dont_encrypt` - (Optional) Type: `string`. Disable backup file encryption. Note that since RouterOS v6.43 without a provided   password,   the backup file is unencrypted.
* `encryption` - (Optional) Type: `string`. The encryption algorithm to use for encrypting the backup file. Note that is not considered a secure encryption method and is only available for compatibility reasons with older RouterOS versions.
* `name` - (Optional) Type: `string`. The filename for the backup file.
* `password` - (Optional) Type: `string`. Password for the encrypted backup file. Note that since RouterOS v6.43 without a provided   password,   the backup file is unencrypted. **Sensitive.**

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

