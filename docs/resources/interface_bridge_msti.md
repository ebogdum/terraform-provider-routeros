---
subcategory: "Bridge"
page_title: "RouterOS: routeros_interface_bridge_msti"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_interface_bridge_msti

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_interface_bridge_msti" "msti_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # bridge = "bridge1"
  # identifier = 0
  # priority = 32768
  # status = 0
  # vlan_mapping = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `bridge` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic` - (Read-only) Type: `bool`.
* `identifier` - (Optional) Type: `int`.
* `priority` - (Optional) Type: `int`.
* `status` - (Read-only) Type: `int`.
* `vlan_mapping` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_bridge_msti.example '*3'

# Named router
terraform import routeros_interface_bridge_msti.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_bridge_msti.example 'home/my-resource-name'
```
