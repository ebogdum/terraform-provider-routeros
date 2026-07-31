---
subcategory: "DHCP"
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
  # allow_reconfigure_messages = false
  # check_gateway = "none"
  # default_route_distance = 0
  # default_route_tables = "replace-me"
  # dhcp_options = []
  # dscp = 0
  # interface = "ether1"
  # name = "tf-example"
  # release = "replace-me"
  # renew = "replace-me"
  # route = "replace-me"
  # routing_tables = "replace-me"
  # script = "replace-me"
  # use_broadcast = "both"
  # use_peer_dns = true
  # use_peer_ntp = true
  # vlan_priority = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `add_default_route` - (Optional) Type: `string`.
* `address` - (Read-only) Type: `string`.
* `allow_reconfigure` - (Optional) Type: `bool`.
* `allow_reconfigure_messages` - (Read-only) Type: `bool`.
* `caps_managers` - (Read-only) Type: `string`.
* `check_gateway` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_route_distance` - (Optional) Type: `int`.
* `default_route_tables` - (Optional) Type: `string`.
* `dhcp_options` - (Optional) Type: `list`.
* `dhcp_server` - (Read-only) Type: `string`.
* `disabled` - (Optional) Type: `bool`.
* `dscp` - (Optional) Type: `int`.
* `dynamic` - (Read-only) Type: `bool`.
* `expires_after` - (Read-only) Type: `string`.
* `gateway` - (Read-only) Type: `string`.
* `interface` - (Optional) Type: `string`.
* `invalid` - (Read-only) Type: `bool`.
* `ip_address` - (Read-only) Type: `string`.
* `last_received_counter` - (Read-only) Type: `string`.
* `name` - (Optional) Type: `string`.
* `primary_dns` - (Read-only) Type: `string`.
* `primary_ntp` - (Read-only) Type: `string`.
* `reconfigure_key` - (Read-only) Type: `string`.
* `release` - (Read-only) Type: `string`.
* `renew` - (Read-only) Type: `string`.
* `route` - (Read-only) Type: `string`.
* `routing_tables` - (Read-only) Type: `string`.
* `script` - (Optional) Type: `string`.
* `secondary_dns` - (Read-only) Type: `string`.
* `secondary_ntp` - (Read-only) Type: `string`.
* `status` - (Read-only) Type: `string`.
* `use_broadcast` - (Optional) Type: `string`.
* `use_peer_dns` - (Optional) Type: `bool`.
* `use_peer_ntp` - (Optional) Type: `bool`.
* `vlan_priority` - (Optional) Type: `int`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
