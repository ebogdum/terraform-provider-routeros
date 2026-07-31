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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `max_block_size` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_tftp_settings.this 'home'
```
