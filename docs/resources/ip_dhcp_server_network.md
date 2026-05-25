---
subcategory: "DHCP"
page_title: "RouterOS: routeros_ip_dhcp_server_network"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_dhcp_server_network

Manages the RouterOS `/ip/dhcp-server/network` menu.

## Example Usage

```terraform
resource "routeros_ip_dhcp_server_network" "network_example" {
  # router = "my-router"  # which router to target; omit for the default
  address = "10.255.255.0/30"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # boot_file_name = "replace-me"
  # caps_manager = "replace-me"
  # caps_managers = "replace-me"
  # dhcp_option = "replace-me"
  # dhcp_option_set = "4.294967295e+09"
  # dhcp_options = "replace-me"
  # dns_server = "1.1.1.1,8.8.8.8"
  # dns_servers = "replace-me"
  # domain = "example.local"
  # dynamic = "replace-me"
  # gateway = "10.255.255.1"
  # netmask = "255.255.255.0"
  # next_server = "replace-me"
  # nndns = "replace-me"
  # nnntp = "replace-me"
  # no_dns = false
  # no_ntp = false
  # ntp_server = "pool.ntp.org"
  # ntp_servers = "replace-me"
  # wins_server = "replace-me"
  # wins_servers = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Required) Type: `cidr`. Default: `10.255.255.0/30`.
* `boot_file_name` - (Optional) Type: `string`.
* `caps_manager` - (Optional) Type: `string`.
* `caps_managers` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dhcp_option` - (Optional) Type: `string`.
* `dhcp_option_set` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `dhcp_options` - (Optional) Type: `string`.
* `dns_server` - (Optional) Type: `string`.
* `dns_servers` - (Optional) Type: `string`.
* `domain` - (Optional) Type: `string`.
* `dynamic` - (Optional) Type: `string`.
* `gateway` - (Optional) Type: `string`. Default: `10.255.255.1`.
* `netmask` - (Optional) Type: `string`.
* `next_server` - (Optional) Type: `string`.
* `nndns` - (Optional) Type: `string`.
* `nnntp` - (Optional) Type: `string`.
* `no_dns` - (Optional) Type: `bool`.
* `no_ntp` - (Optional) Type: `bool`.
* `ntp_server` - (Optional) Type: `string`.
* `ntp_servers` - (Optional) Type: `string`.
* `wins_server` - (Optional) Type: `string`.
* `wins_servers` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_dhcp_server_network.example '*3'

# Named router
terraform import routeros_ip_dhcp_server_network.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_dhcp_server_network.example 'home/my-resource-name'
```
