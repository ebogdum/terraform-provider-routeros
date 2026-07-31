---
subcategory: "PPP"
page_title: "RouterOS: routeros_ppp_profile"
description: |-
  RouterOS resource.
---

# Resource: routeros_ppp_profile

Manages the RouterOS `/ppp/profile` menu.

## Example Usage

```terraform
resource "routeros_ppp_profile" "profile_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # address_list = "replace-me"
  # bridge = "bridge1"
  # bridge_horizon = "replace-me"
  # bridge_learning = "no"
  # bridge_path_cost = "replace-me"
  # bridge_port_priority = "replace-me"
  # bridge_port_trusted = "replace-me"
  # bridge_port_vid = "replace-me"
  # change_tcp_mss = "no"
  # def = "replace-me"
  # dhcpv6_lease_time = "replace-me"
  # dhcpv6_pd_pool = "replace-me"
  # dhcpv6_use_radius = "replace-me"
  # dns_server = "1.1.1.1,8.8.8.8"
  # idle_timeout = "replace-me"
  # incoming_filter = "replace-me"
  # insert_queue_before = "replace-me"
  # interface_list = "replace-me"
  # ipv6 = "replace-me"
  # local_address = "10.99.0.1"
  # on_down = "replace-me"
  # on_up = "replace-me"
  # only_one = "no"
  # outgoing_filter = "replace-me"
  # parent_queue = "replace-me"
  # queue_type_rx_tx = "replace-me"
  # rate_limit_rx_tx = "replace-me"
  # remote_address = "10.99.0.1"
  # remote_ipv6_prefix_pool = "replace-me"
  # remote_ipv6_prefix_reuse = "replace-me"
  # session_timeout = "replace-me"
  # use_compression = "no"
  # use_encryption = "no"
  # use_ipv6 = "yes"
  # use_mpls = "no"
  # use_upnp = "replace-me"
  # wins_server = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address_list` - (Optional) Type: `string`.
* `bridge` - (Optional) Type: `string`.
* `bridge_horizon` - (Optional) Type: `string`.
* `bridge_learning` - (Optional) Type: `string`.
* `bridge_path_cost` - (Optional) Type: `string`.
* `bridge_port_priority` - (Optional) Type: `string`.
* `bridge_port_trusted` - (Optional) Type: `string`.
* `bridge_port_vid` - (Optional) Type: `string`.
* `change_tcp_mss` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `def` - (Optional) Type: `string`.
* `default` - (Read-only) Type: `bool`.
* `dhcpv6_lease_time` - (Optional) Type: `string`.
* `dhcpv6_pd_pool` - (Optional) Type: `string`.
* `dhcpv6_use_radius` - (Optional) Type: `string`.
* `dns_server` - (Optional) Type: `string`.
* `idle_timeout` - (Optional) Type: `string`.
* `incoming_filter` - (Optional) Type: `string`.
* `insert_queue_before` - (Optional) Type: `string`.
* `interface_list` - (Optional) Type: `string`.
* `ipv6` - (Read-only) Type: `string`.
* `local_address` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `on_down` - (Optional) Type: `string`.
* `on_up` - (Optional) Type: `string`.
* `only_one` - (Optional) Type: `string`.
* `outgoing_filter` - (Optional) Type: `string`.
* `parent_queue` - (Optional) Type: `string`.
* `queue_type` - (Optional) Type: `string`. RouterOS `queue-type`.
* `queue_type_rx_tx` - (Read-only) Type: `string`.
* `rate_limit` - (Optional) Type: `string`. RouterOS `rate-limit`.
* `rate_limit_rx_tx` - (Read-only) Type: `string`.
* `remote_address` - (Optional) Type: `string`.
* `remote_ipv6_prefix_pool` - (Optional) Type: `string`.
* `remote_ipv6_prefix_reuse` - (Optional) Type: `string`.
* `session_timeout` - (Optional) Type: `string`.
* `use_compression` - (Optional) Type: `string`.
* `use_encryption` - (Optional) Type: `string`.
* `use_ipv6` - (Optional) Type: `string`.
* `use_mpls` - (Optional) Type: `string`.
* `use_upnp` - (Optional) Type: `string`.
* `wins_server` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ppp_profile.example '*3'

# Named router
terraform import routeros_ppp_profile.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ppp_profile.example 'home/my-resource-name'
```
