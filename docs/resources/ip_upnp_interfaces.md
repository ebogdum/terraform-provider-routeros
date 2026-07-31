---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_upnp_interfaces"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_upnp_interfaces

Manages the RouterOS `/ip/upnp/interfaces` menu.

## Example Usage

```terraform
resource "routeros_ip_upnp_interfaces" "interfaces_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  type = "internal"

  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `forced_ip` - (Optional) Type: `string`. RouterOS `forced-ip`.
* `interface` - (Required) Type: `string`.
* `type` - (Required) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_upnp_interfaces.example '*3'

# Named router
terraform import routeros_ip_upnp_interfaces.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_upnp_interfaces.example 'home/my-resource-name'
```
