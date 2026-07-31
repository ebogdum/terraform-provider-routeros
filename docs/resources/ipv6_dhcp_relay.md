---
subcategory: "DHCP"
page_title: "RouterOS: routeros_ipv6_dhcp_relay"
description: |-
  server address validator rejects literal addresses on this ROS version. Skipped.
---

# Resource: routeros_ipv6_dhcp_relay

server address validator rejects literal addresses on this ROS version. Skipped.

## Example Usage

```terraform
resource "routeros_ipv6_dhcp_relay" "dhcp_relay_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `delay_threshold` - (Optional) Type: `string`. RouterOS `delay-threshold`.
* `dhcp_options` - (Optional) Type: `string`. RouterOS `dhcp-options`.
* `dhcp_server` - (Optional) Type: `string`. RouterOS `dhcp-server`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Optional) Type: `string`. RouterOS `interface`.
* `link_address` - (Optional) Type: `string`. RouterOS `link-address`.
* `name` - (Optional) Type: `string`. RouterOS `name`.
* `store_relayed_bindings` - (Optional) Type: `string`. RouterOS `store-relayed-bindings`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_dhcp_relay.example '*3'

# Named router
terraform import routeros_ipv6_dhcp_relay.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_dhcp_relay.example 'home/my-resource-name'
```
