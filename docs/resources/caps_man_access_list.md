---
subcategory: "System & misc"
page_title: "RouterOS: routeros_caps_man_access_list"
description: |-
  RouterOS resource.
---

# Resource: routeros_caps_man_access_list

Manages the RouterOS `/caps-man/access-list` menu.

## Example Usage

```terraform
resource "routeros_caps_man_access_list" "access_list_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "replace-me"
  # interface = "ether1"
  # mac_address = "10.99.0.0/24"
  # vlan_id = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `action` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`.
* `vlan_id` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_caps_man_access_list.example '*3'

# Named router
terraform import routeros_caps_man_access_list.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_caps_man_access_list.example 'home/my-resource-name'
```
