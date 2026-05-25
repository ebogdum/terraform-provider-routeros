---
page_title: "RouterOS: routeros_interface_pppoe_server"
description: |-
  Needs an interface to bind to, plus auth-stack setup.
---

# Resource: routeros_interface_pppoe_server

Needs an interface to bind to, plus auth-stack setup.

## Example Usage

```terraform
resource "routeros_interface_pppoe_server" "pppoe_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # name = "tf-example"
  # service = "replace-me"
  # user = "myuser"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `name` - (Optional) Type: `string`.
* `service` - (Optional) Type: `string`.
* `user` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_pppoe_server.example '*3'

# Named router
terraform import routeros_interface_pppoe_server.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_pppoe_server.example 'home/my-resource-name'
```
