---
subcategory: "OSPF"
page_title: "RouterOS: routeros_routing_ospf_area"
description: |-
  Auto-test requires a typed-reference precondition (e.g. an existing peer,
---

# Resource: routeros_routing_ospf_area

Auto-test requires a typed-reference precondition (e.g. an existing peer,
instance, bridge of the specific kind). The current acc-test generator's
generic data.routeros_interface.all lookup can't satisfy these. Use this
resource manually with explicit references to a precondition resource
in your config.


## Example Usage

```terraform
resource "routeros_routing_ospf_area" "area_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # area_id = "10.99.0.1"
  # default_cost = "replace-me"
  # instance = "replace-me"
  # name = "tf-example"
  # no_summaries = "replace-me"
  # nssa_translator = "replace-me"
  # type = "default"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `area_id` - (Optional) Type: `ip`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_cost` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `instance` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `no_summaries` - (Optional) Type: `string`.
* `nssa_translator` - (Optional) Type: `string`.
* `type` - (Optional) Type: `enum(default|stub|nssa)`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_ospf_area.example '*3'

# Named router
terraform import routeros_routing_ospf_area.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_ospf_area.example 'home/my-resource-name'
```
