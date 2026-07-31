---
subcategory: "RIP"
page_title: "RouterOS: routeros_routing_rip_interface"
description: |-
  Discovered; needs rip instance
---

# Resource: routeros_routing_rip_interface

Discovered; needs rip instance

## Example Usage

```terraform
resource "routeros_routing_rip_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # cost = "replace-me"
  # instance = "replace-me"
  # interfaces = "replace-me"
  # key_chain = "replace-me"
  # mode = "replace-me"
  # name = "tf-example"
  # password = "REDACTED"
  # poison_reverse = "replace-me"
  # source_addresses = "replace-me"
  # split_horizon = "replace-me"
  # use_bfd = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `cost` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `instance` - (Optional) Type: `string`.
* `interfaces` - (Optional) Type: `string`.
* `key_chain` - (Optional) Type: `string`.
* `mode` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `password` - (Optional) Type: `string`. **Sensitive.**
* `poison_reverse` - (Optional) Type: `string`.
* `source_addresses` - (Optional) Type: `string`.
* `split_horizon` - (Optional) Type: `string`.
* `use_bfd` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_rip_interface.example '*3'

# Named router
terraform import routeros_routing_rip_interface.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_rip_interface.example 'home/my-resource-name'
```
