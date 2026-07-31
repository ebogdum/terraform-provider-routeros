---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_arp"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_arp

Manages the RouterOS `/ip/arp` menu.

## Example Usage

```terraform
resource "routeros_ip_arp" "arp_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "10.99.0.1"
  interface = "ether1"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # ip_address = "10.99.0.0/24"
  # mac_address = "10.99.0.0/24"
  # mac_ping = "replace-me"
  # mac_telnet = "replace-me"
  # make_static = "replace-me"
  # ping = "replace-me"
  # published = false
  # telnet = "replace-me"
  # torch = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Required) Type: `string`.
* `bridge_port` - (Read-only) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `complete` - (Read-only) Type: `bool`.
* `dhcp` - (Read-only) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `dynamic` - (Read-only) Type: `bool`.
* `host_name` - (Read-only) Type: `string`.
* `interface` - (Required) Type: `string`.
* `invalid` - (Read-only) Type: `bool`.
* `ip_address` - (Read-only) Type: `string`.
* `mac_address` - (Optional) Type: `string`.
* `mac_ping` - (Read-only) Type: `string`.
* `mac_telnet` - (Read-only) Type: `string`.
* `make_static` - (Read-only) Type: `string`.
* `ping` - (Read-only) Type: `string`.
* `published` - (Optional) Type: `bool`.
* `status` - (Read-only) Type: `string`.
* `telnet` - (Read-only) Type: `string`.
* `torch` - (Read-only) Type: `string`.
* `vrf` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_arp.example '*3'

# Named router
terraform import routeros_ip_arp.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_arp.example 'home/my-resource-name'
```
