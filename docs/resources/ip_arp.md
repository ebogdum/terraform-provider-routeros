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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Required) Type: `ip`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `interface` - (Required) Type: `string`.
* `ip_address` - (Optional) Type: `ip`.
* `mac_address` - (Optional) Type: `mac`.
* `mac_ping` - (Optional) Type: `string`.
* `mac_telnet` - (Optional) Type: `string`.
* `make_static` - (Optional) Type: `string`.
* `ping` - (Optional) Type: `string`.
* `published` - (Optional) Type: `bool`.
* `telnet` - (Optional) Type: `string`.
* `torch` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `bridge_port` - Type: `string`.
* `complete` - Type: `bool`.
* `dhcp` - Type: `bool`.
* `dynamic` - Type: `bool`.
* `host_name` - Type: `string`.
* `invalid` - Type: `bool`.
* `status` - Type: `string`.
* `vrf` - Type: `string`.

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
