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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `channel` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `country` - (Optional) Type: `string`.
* `datapath` - (Optional) Type: `string`. Name of a `/caps-man/datapath` profile to apply.
* `distance` - (Optional) Type: `string`.
* `mode` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `rates` - (Optional) Type: `string`. Name of a `/caps-man/rates` profile to apply.
* `security` - (Optional) Type: `string`. Name of a `/caps-man/security` profile to apply.
* `ssid` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
