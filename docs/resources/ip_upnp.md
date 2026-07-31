---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_upnp"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_upnp

Manages the RouterOS `/ip/upnp` menu.

## Example Usage

```terraform
resource "routeros_ip_upnp" "upnp_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # allow_disable_external_interface = false
  # enabled = false
  # show_dummy_rule = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `allow_disable_external_interface` - (Optional) Type: `bool`.
* `enabled` - (Optional) Type: `bool`.
* `show_dummy_rule` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_upnp.this 'home'
```
