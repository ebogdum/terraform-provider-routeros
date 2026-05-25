---
page_title: "RouterOS: routeros_ip_dhcp_client"
description: |-
  DHCP client per interface. On most devices the default config already binds one to an interface, causing "dhcp-client on that interface already" on a fresh add. Skipped.
---

# Resource: routeros_ip_dhcp_client

DHCP client per interface. On most devices the default config already binds one to an interface, causing "dhcp-client on that interface already" on a fresh add. Skipped.

## Example Usage

```terraform
resource "routeros_ip_dhcp_client" "dhcp_client_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # add_default_route = "yes"
  # allow_reconfigure = false
  # check_gateway = "none"
  # default_route_distance = 0
  # default_route_tables = "replace-me"
  # dhcp_options = []
  # dscp = 0
  # interface = "ether1"
  # name = "tf-example"
  # script = "replace-me"
  # use_broadcast = "both"
  # use_peer_dns = true
  # use_peer_ntp = true
  # vlan_priority = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `add_default_route` - (Optional) Type: `enum(no|yes|special classless)`. Default: `1`.
* `allow_reconfigure` - (Optional) Type: `bool`.
* `check_gateway` - (Optional) Type: `enum(none|arp|ping|bfd)`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_route_distance` - (Optional) Type: `int`.
* `default_route_tables` - (Optional) Type: `string`.
* `dhcp_options` - (Optional) Type: `list`.
* `disabled` - (Optional) Type: `bool`.
* `dscp` - (Optional) Type: `int`.
* `interface` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `script` - (Optional) Type: `string`.
* `use_broadcast` - (Optional) Type: `enum(both|always|never)`.
* `use_peer_dns` - (Optional) Type: `bool`. Default: `1`.
* `use_peer_ntp` - (Optional) Type: `bool`. Default: `1`.
* `vlan_priority` - (Optional) Type: `int`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `address` - Type: `cidr`.
* `dhcp_server` - Type: `ip`.
* `dynamic` - Type: `bool`.
* `expires_after` - Type: `duration`.
* `gateway` - Type: `ip`.
* `invalid` - Type: `bool`.
* `primary_dns` - Type: `ip`.
* `primary_ntp` - Type: `ip`.
* `status` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_dhcp_client.example '*3'

# Named router
terraform import routeros_ip_dhcp_client.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_dhcp_client.example 'home/my-resource-name'
```
