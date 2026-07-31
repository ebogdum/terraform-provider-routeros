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
  # no_summaries = false
  # nssa_translator = "replace-me"
  # type = "default"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `area_id` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_cost` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic` - (Read-only) Type: `bool`.
* `instance` - (Optional) Type: `string`.
* `invalid` - (Read-only) Type: `bool`.
* `name` - (Optional) Type: `string`.
* `no_summaries` - (Optional) Type: `bool`.
* `nssa_translator` - (Optional) Type: `string`.
* `transit_capable` - (Read-only) Type: `bool`.
* `type` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
