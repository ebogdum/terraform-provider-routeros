---
subcategory: "Bridge"
page_title: "RouterOS: routeros_interface_bridge_vlan"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_interface_bridge_vlan

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_interface_bridge_vlan" "vlan_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # bridge = "bridge1"
  # mvrp_attributes = "replace-me"
  # mvrp_forbidden = "replace-me"
  # tagged = "replace-me"
  # untagged = "replace-me"
  # vlan_ids = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `bridge` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `mvrp_attributes` - (Optional) Type: `string`.
* `mvrp_forbidden` - (Optional) Type: `string`.
* `tagged` - (Optional) Type: `string`.
* `untagged` - (Optional) Type: `string`.
* `vlan_ids` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `current_tagged` - Type: `string`.
* `current_untagged` - Type: `string`.
* `dynamic` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_bridge_vlan.example '*3'

# Named router
terraform import routeros_interface_bridge_vlan.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_bridge_vlan.example 'home/my-resource-name'
```
