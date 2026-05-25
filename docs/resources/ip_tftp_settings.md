---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_tftp_settings"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_tftp_settings

Manages the RouterOS `/ip/tftp/settings` menu.

## Example Usage

```terraform
resource "routeros_ip_tftp_settings" "settings_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # max_block_size = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `max_block_size` - (Optional) Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_tftp_settings.this 'home'
```
