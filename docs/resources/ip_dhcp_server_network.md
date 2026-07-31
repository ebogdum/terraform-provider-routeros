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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Required) Type: `string`.
* `boot_file_name` - (Optional) Type: `string`.
* `caps_manager` - (Optional) Type: `string`.
* `caps_managers` - (Read-only) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dhcp_option` - (Optional) Type: `string`.
* `dhcp_option_set` - (Optional) Type: `string`.
* `dhcp_options` - (Read-only) Type: `string`.
* `dns_none` - (Optional) Type: `string`. RouterOS `dns-none`.
* `dns_server` - (Optional) Type: `string`.
* `dns_servers` - (Read-only) Type: `string`.
* `domain` - (Optional) Type: `string`.
* `dynamic` - (Read-only) Type: `string`.
* `gateway` - (Optional) Type: `string`.
* `netmask` - (Optional) Type: `string`.
* `next_server` - (Optional) Type: `string`.
* `nndns` - (Read-only) Type: `string`.
* `nnntp` - (Read-only) Type: `string`.
* `no_dns` - (Read-only) Type: `bool`.
* `no_ntp` - (Read-only) Type: `bool`.
* `ntp_none` - (Optional) Type: `string`. RouterOS `ntp-none`.
* `ntp_server` - (Optional) Type: `string`.
* `ntp_servers` - (Read-only) Type: `string`.
* `wins_server` - (Optional) Type: `string`.
* `wins_servers` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
