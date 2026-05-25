---
page_title: "RouterOS: routeros_interface_lte_apn"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_lte_apn

Manages the RouterOS `/interface/lte/apn` menu.

## Example Usage

```terraform
resource "routeros_interface_lte_apn" "apn_example" {
  # router = "my-router"  # which router to target; omit for the default
  apn = "internet"
  name = "tf-example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # add_default_route = false
  # authentication = "replace-me"
  # default_route_distance = 0
  # ip_type = "replace-me"
  # use_network_apn = false
  # use_peer_dns = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `add_default_route` - (Optional) Type: `bool`.
* `apn` - (Required) Type: `string`. Default: `internet`.
* `authentication` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_route_distance` - (Optional) Type: `int`.
* `ip_type` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf-acc-apn`.
* `use_network_apn` - (Optional) Type: `bool`.
* `use_peer_dns` - (Optional) Type: `bool`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_lte_apn.example '*3'

# Named router
terraform import routeros_interface_lte_apn.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_lte_apn.example 'home/my-resource-name'
```
