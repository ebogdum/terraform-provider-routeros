---
subcategory: "System & misc"
page_title: "RouterOS: routeros_caps_man_datapath"
description: |-
  RouterOS resource.
---

# Resource: routeros_caps_man_datapath

Manages the RouterOS `/caps-man/datapath` menu.

## Example Usage

```terraform
resource "routeros_caps_man_datapath" "datapath_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # bridge = "bridge1"
  # l2mtu = "replace-me"
  # mtu = "replace-me"
  # vlan_id = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `arp` - (Optional) Type: `string`.
* `bridge` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `l2mtu` - (Optional) Type: `string`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `vlan_id` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_caps_man_datapath.example '*3'

# Named router
terraform import routeros_caps_man_datapath.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_caps_man_datapath.example 'home/my-resource-name'
```
