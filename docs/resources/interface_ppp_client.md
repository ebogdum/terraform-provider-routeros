---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ppp_client"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_ppp_client

Manages the RouterOS `/interface/ppp-client` menu.

## Example Usage

```terraform
resource "routeros_interface_ppp_client" "ppp_client_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # allow = "replace-me"
  # default_route_distance = "replace-me"
  # dial_on_demand = "replace-me"
  # keepalive_timeout = "replace-me"
  # max_mru = "replace-me"
  # max_mtu = "replace-me"
  # mrru = "replace-me"
  # name = "tf-example"
  # password = "REDACTED"
  # profile = "replace-me"
  # remote_address = "10.99.0.1"
  # user = "myuser"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `allow` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_route_distance` - (Optional) Type: `string`.
* `dial_on_demand` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `keepalive_timeout` - (Optional) Type: `string`.
* `max_mru` - (Optional) Type: `string`.
* `max_mtu` - (Optional) Type: `string`.
* `mrru` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `password` - (Optional) Type: `string`.
* `profile` - (Optional) Type: `string`.
* `remote_address` - (Optional) Type: `string`.
* `user` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_ppp_client.example '*3'

# Named router
terraform import routeros_interface_ppp_client.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_ppp_client.example 'home/my-resource-name'
```
