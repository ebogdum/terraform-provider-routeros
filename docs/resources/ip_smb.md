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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Set comment for the server
* `domain` - (Optional) Type: `string`. Name of Windows Workgroup
* `enabled` - (Optional) Type: `string`. The default value is 'auto.' This means that the SMB server will automatically be enabled when the first non-disabled SMB share is configured under '/ip smb share'
* `interface` - (Optional) Type: `string`. List of interfaces on which SMB service will be running. all - SMB will be available on all interfaces.
* `interfaces` - (Optional) Type: `string`.
* `status` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_smb.this 'home'
```
