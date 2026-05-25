---
subcategory: "System & misc"
page_title: "RouterOS: routeros_caps_man_interface"
description: |-
  CAPsMAN virtual interfaces are typically created automatically; manual creation collides with the master.
---

# Resource: routeros_caps_man_interface

CAPsMAN virtual interfaces are typically created automatically; manual creation collides with the master.

## Example Usage

```terraform
resource "routeros_caps_man_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp_timeout = "replace-me"
  # mac_address = "10.99.0.0/24"
  # master_interface = "ether1"
  # name = "tf-example"
  # radio_mac = "02:00:00:00:00:01"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `arp_timeout` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `mac_address` - (Optional) Type: `mac`.
* `master_interface` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `radio_mac` - (Optional) Type: `mac`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_caps_man_interface.example '*3'

# Named router
terraform import routeros_caps_man_interface.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_caps_man_interface.example 'home/my-resource-name'
```
