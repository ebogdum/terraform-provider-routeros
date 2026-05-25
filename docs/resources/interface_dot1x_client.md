---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_dot1x_client"
description: |-
  802.1X client attaches to a specific Ethernet interface; values vary per device. Skipped.
---

# Resource: routeros_interface_dot1x_client

802.1X client attaches to a specific Ethernet interface; values vary per device. Skipped.

## Example Usage

```terraform
resource "routeros_interface_dot1x_client" "client_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # anon_identity = "replace-me"
  # certificate = "replace-me"
  # eap_methods = "replace-me"
  # identity = "replace-me"
  # interface = "ether1"
  # password = "REDACTED"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `anon_identity` - (Optional) Type: `string`.
* `certificate` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `eap_methods` - (Optional) Type: `string`.
* `identity` - (Optional) Type: `string`.
* `interface` - (Optional) Type: `string`.
* `password` - (Optional) Type: `string`. **Sensitive.**

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `invalid` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_dot1x_client.example '*3'

# Named router
terraform import routeros_interface_dot1x_client.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_dot1x_client.example 'home/my-resource-name'
```
