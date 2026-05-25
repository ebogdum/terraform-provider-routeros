---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_smb"
description: |-
  Requires SMB package/storage backend
---

# Resource: routeros_ip_smb

Requires SMB package/storage backend

## Example Usage

```terraform
resource "routeros_ip_smb" "smb_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # domain = "example.local"
  # enabled = "replace-me"
  # interface = "ether1"
  # interfaces = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Set comment for the server.
* `domain` - (Optional) Type: `string`. Name of Windows Workgroup.
* `enabled` - (Optional) Type: `string`. The default value is 'auto.' This means that the SMB server will automatically be enabled when the first non-disabled SMB share is configured under '/ip smb share'.
* `interface` - (Optional) Type: `string`. List of interfaces on which SMB service will be running. all - SMB will be available on all interfaces.
* `interfaces` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `status` - Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_smb.this 'home'
```
