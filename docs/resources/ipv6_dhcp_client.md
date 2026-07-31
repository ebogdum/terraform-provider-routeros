---
subcategory: "DHCP"
page_title: "RouterOS: routeros_ipv6_dhcp_client"
description: |-
  RouterOS resource.
---

# Resource: routeros_ipv6_dhcp_client

Manages the RouterOS `/ipv6/dhcp-client` menu.

## Example Usage

```terraform
resource "routeros_ipv6_dhcp_client" "dhcp_client_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  request = "address"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # default_route_distance = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `accept_prefix_without_address` - (Optional) Type: `string`. RouterOS `accept-prefix-without-address`.
* `add_default_route` - (Optional) Type: `string`. RouterOS `add-default-route`.
* `allow_reconfigure` - (Optional) Type: `string`. RouterOS `allow-reconfigure`.
* `check_gateway` - (Optional) Type: `string`. RouterOS `check-gateway`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `custom_duid` - (Optional) Type: `string`. RouterOS `custom-duid`.
* `custom_iana_id` - (Optional) Type: `string`. RouterOS `custom-iana-id`.
* `custom_iapd_id` - (Optional) Type: `string`. RouterOS `custom-iapd-id`.
* `default_route_distance` - (Optional) Type: `string`.
* `default_route_tables` - (Optional) Type: `string`. RouterOS `default-route-tables`.
* `dhcp_options` - (Optional) Type: `string`. RouterOS `dhcp-options`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.
* `pool_name` - (Optional) Type: `string`. RouterOS `pool-name`.
* `pool_prefix_length` - (Optional) Type: `string`. RouterOS `pool-prefix-length`.
* `prefix_address_lists` - (Optional) Type: `string`. RouterOS `prefix-address-lists`.
* `prefix_hint` - (Optional) Type: `string`. RouterOS `prefix-hint`.
* `rapid_commit` - (Optional) Type: `string`. RouterOS `rapid-commit`.
* `request` - (Required) Type: `string`.
* `script` - (Optional) Type: `string`. RouterOS `script`.
* `use_interface_duid` - (Optional) Type: `string`. RouterOS `use-interface-duid`.
* `use_peer_dns` - (Optional) Type: `string`. RouterOS `use-peer-dns`.
* `validate_server_duid` - (Optional) Type: `string`. RouterOS `validate-server-duid`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_dhcp_client.example '*3'

# Named router
terraform import routeros_ipv6_dhcp_client.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_dhcp_client.example 'home/my-resource-name'
```
