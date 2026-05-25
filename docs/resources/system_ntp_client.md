---
subcategory: "System"
page_title: "RouterOS: routeros_system_ntp_client"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_ntp_client

Manages the RouterOS `/system/ntp/client` menu.

## Example Usage

```terraform
resource "routeros_system_ntp_client" "client_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # enabled = false
  # mode = "replace-me"
  # servers = "replace-me"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `enabled` - (Optional) Type: `bool`.
* `mode` - (Optional) Type: `string`.
* `servers` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `freq_drift` - Type: `int`.
* `status` - Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_ntp_client.this 'home'
```
