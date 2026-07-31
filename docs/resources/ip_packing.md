---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_packing"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_packing

Manages the RouterOS `/ip/packing` menu.

## Example Usage

```terraform
resource "routeros_ip_packing" "packing_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"

  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `aggregated_size` - (Optional) Type: `string`. RouterOS `aggregated-size`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.
* `packing` - (Optional) Type: `string`. RouterOS `packing`.
* `unpacking` - (Optional) Type: `string`. RouterOS `unpacking`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_packing.example '*3'

# Named router
terraform import routeros_ip_packing.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_packing.example 'home/my-resource-name'
```
