---
subcategory: "System & misc"
page_title: "RouterOS: routeros_caps_man_configuration"
description: |-
  RouterOS resource.
---

# Resource: routeros_caps_man_configuration

Manages the RouterOS `/caps-man/configuration` menu.

## Example Usage

```terraform
resource "routeros_caps_man_configuration" "configuration_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # channel = "replace-me"
  # country = "replace-me"
  # distance = "replace-me"
  # mode = "replace-me"
  # ssid = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `channel` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `country` - (Optional) Type: `string`.
* `distance` - (Optional) Type: `string`.
* `mode` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `ros_audit_20260523213235_4`.
* `ssid` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_caps_man_configuration.example '*3'

# Named router
terraform import routeros_caps_man_configuration.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_caps_man_configuration.example 'home/my-resource-name'
```
