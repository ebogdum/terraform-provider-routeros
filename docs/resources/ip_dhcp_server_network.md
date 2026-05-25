---
subcategory: "IP"
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
  # dhcp_option_set = "4.294967295e+09"
  # domain = "example.local"
  # gateway = "10.255.255.1"
  # netmask = "255.255.255.0"
  # next_server = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Required) Type: `cidr`. Default: `10.255.255.0/30`.
* `boot_file_name` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dhcp_option_set` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `domain` - (Optional) Type: `string`.
* `gateway` - (Optional) Type: `string`. Default: `10.255.255.1`.
* `netmask` - (Optional) Type: `string`.
* `next_server` - (Optional) Type: `string`.

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
