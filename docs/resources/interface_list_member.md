---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_list_member"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_interface_list_member

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_interface_list_member" "member_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # dynamic = "replace-me"
  # interface = "ether1"
  # list = "my-list"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic` - (Optional) Type: `string`.
* `interface` - (Optional) Type: `string`.
* `list` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_list_member.example '*3'

# Named router
terraform import routeros_interface_list_member.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_list_member.example 'home/my-resource-name'
```
