---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_wifi_provisioning"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_provisioning

Manages the RouterOS `/interface/wifi/provisioning` menu.

## Example Usage

```terraform
resource "routeros_interface_wifi_provisioning" "provisioning_example" {
  # router = "my-router"  # which router to target; omit for the default
  action = "none"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address_ranges = "replace-me"
  # common_name_regexp = "replace-me"
  # identity_regexp = "replace-me"
  # master_configuration = "replace-me"
  # multi_link_mode = "replace-me"
  # name_format = "replace-me"
  # radio_mac = "replace-me"
  # slave_configurations = "replace-me"
  # slave_name_format = "replace-me"
  # supported_bands = "replace-me"
  # supported_hw_caps = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `action` - (Required) Type: `enum(none|create enabled|create disabled|create dynamic enabled)`. Default: `create-dynamic-enabled`.
* `address_ranges` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `common_name_regexp` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `identity_regexp` - (Optional) Type: `string`.
* `master_configuration` - (Optional) Type: `string`.
* `multi_link_mode` - (Optional) Type: `string`.
* `name_format` - (Optional) Type: `string`.
* `radio_mac` - (Optional) Type: `string`.
* `slave_configurations` - (Optional) Type: `string`.
* `slave_name_format` - (Optional) Type: `string`.
* `supported_bands` - (Optional) Type: `string`.
* `supported_hw_caps` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wifi_provisioning.example '*3'

# Named router
terraform import routeros_interface_wifi_provisioning.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wifi_provisioning.example 'home/my-resource-name'
```
