---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ppp_server"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_ppp_server

Manages the RouterOS `/interface/ppp-server` menu.

## Example Usage

```terraform
resource "routeros_interface_ppp_server" "ppp_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # authentication = "replace-me"
  # max_mru = "replace-me"
  # max_mtu = "replace-me"
  # mrru = "replace-me"
  # name = "example"
  # profile = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `authentication` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `max_mru` - (Optional) Type: `string`.
* `max_mtu` - (Optional) Type: `string`.
* `mrru` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `profile` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_ppp_server.example '*3'

# Named router
terraform import routeros_interface_ppp_server.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_ppp_server.example 'home/my-resource-name'
```
