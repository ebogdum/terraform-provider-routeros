---
subcategory: "System & misc"
page_title: "RouterOS: routeros_caps_man_provisioning"
description: |-
  RouterOS resource.
---

# Resource: routeros_caps_man_provisioning

Manages the RouterOS `/caps-man/provisioning` menu.

## Example Usage

```terraform
resource "routeros_caps_man_provisioning" "provisioning_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "none"
  # common_name_regexp = "replace-me"
  # hw_supported_modes = "replace-me"
  # identity_regexp = "replace-me"
  # ip_address_ranges = "replace-me"
  # master_configuration = "replace-me"
  # name_format = "cap"
  # name_prefix = "replace-me"
  # radio_mac = "02:00:00:00:00:01"
  # slave_configuration = "replace-me"
  # slave_configurations = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `action` - (Optional) Type: `enum(none|create-enabled|create-disabled|create-dynamic-enabled)`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `common_name_regexp` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `hw_supported_modes` - (Optional) Type: `string`.
* `identity_regexp` - (Optional) Type: `string`.
* `ip_address_ranges` - (Optional) Type: `string`.
* `master_configuration` - (Optional) Type: `string`.
* `name_format` - (Optional) Type: `enum(cap|prefix|identity|prefix-identity)`.
* `name_prefix` - (Optional) Type: `string`.
* `radio_mac` - (Optional) Type: `mac`.
* `slave_configuration` - (Optional) Type: `string`.
* `slave_configurations` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_caps_man_provisioning.example '*3'

# Named router
terraform import routeros_caps_man_provisioning.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_caps_man_provisioning.example 'home/my-resource-name'
```
