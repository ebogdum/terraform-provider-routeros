---
subcategory: "Caps-man"
page_title: "RouterOS: routeros_caps_man_channel"
description: |-
  RouterOS resource.
---

# Resource: routeros_caps_man_channel

Manages the RouterOS `/caps-man/channel` menu.

## Example Usage

```terraform
resource "routeros_caps_man_channel" "channel_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # band = "replace-me"
  # frequency = "replace-me"
  # tx_power = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `band` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `frequency` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `ros_audit_20260523213235_3`.
* `tx_power` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_caps_man_channel.example '*3'

# Named router
terraform import routeros_caps_man_channel.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_caps_man_channel.example 'home/my-resource-name'
```
