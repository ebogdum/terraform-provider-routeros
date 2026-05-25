---
subcategory: "Firewall"
page_title: "RouterOS: routeros_ipv6_firewall_raw"
description: |-
  RouterOS resource.
---

# Resource: routeros_ipv6_firewall_raw

Manages the RouterOS `/ipv6/firewall/raw` menu.

## Example Usage

```terraform
resource "routeros_ipv6_firewall_raw" "raw_example" {
  # router = "my-router"  # which router to target; omit for the default
  action = "accept"
  chain = "prerouting"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address_list = "replace-me"
  # address_list_timeout = "replace-me"
  # content = "replace-me"
  # dscp = "replace-me"
  # dst_address = "10.99.0.0/24"
  # dst_address_list = "my-list"
  # dst_address_type = "replace-me"
  # dst_limit = "replace-me"
  # dst_port = "443"
  # icmp_options = "replace-me"
  # in_bridge_port = "443"
  # in_bridge_port_list = "replace-me"
  # in_interface = "ether1"
  # in_interface_list = "LAN"
  # ingress_priority = "replace-me"
  # ipsec_policy = "replace-me"
  # jump_target = "replace-me"
  # limit = "replace-me"
  # log = "replace-me"
  # log_prefix = "replace-me"
  # nth = "replace-me"
  # out_bridge_port = "443"
  # out_bridge_port_list = "replace-me"
  # out_interface = "ether1"
  # out_interface_list = "LAN"
  # packet_mark = "replace-me"
  # packet_size = "replace-me"
  # per_connection_classifier = "replace-me"
  # port = "443"
  # priority = "replace-me"
  # protocol = "replace-me"
  # random = "replace-me"
  # src_address = "10.99.0.0/24"
  # src_address_list = "my-list"
  # src_address_type = "replace-me"
  # src_mac_address = "10.99.0.0/24"
  # src_port = "443"
  # tcp_flags = "replace-me"
  # tcp_mss = "replace-me"
  # time = "replace-me"
  # tls_host = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `action` - (Required) Type: `string`. Default: `accept`.
* `address_list` - (Optional) Type: `string`.
* `address_list_timeout` - (Optional) Type: `string`.
* `chain` - (Required) Type: `string`. Default: `prerouting`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `content` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dscp` - (Optional) Type: `string`.
* `dst_address` - (Optional) Type: `string`.
* `dst_address_list` - (Optional) Type: `string`.
* `dst_address_type` - (Optional) Type: `string`.
* `dst_limit` - (Optional) Type: `string`.
* `dst_port` - (Optional) Type: `string`.
* `icmp_options` - (Optional) Type: `string`.
* `in_bridge_port` - (Optional) Type: `string`.
* `in_bridge_port_list` - (Optional) Type: `string`.
* `in_interface` - (Optional) Type: `string`.
* `in_interface_list` - (Optional) Type: `string`.
* `ingress_priority` - (Optional) Type: `string`.
* `ipsec_policy` - (Optional) Type: `string`.
* `jump_target` - (Optional) Type: `string`.
* `limit` - (Optional) Type: `string`.
* `log` - (Optional) Type: `string`.
* `log_prefix` - (Optional) Type: `string`.
* `nth` - (Optional) Type: `string`.
* `out_bridge_port` - (Optional) Type: `string`.
* `out_bridge_port_list` - (Optional) Type: `string`.
* `out_interface` - (Optional) Type: `string`.
* `out_interface_list` - (Optional) Type: `string`.
* `packet_mark` - (Optional) Type: `string`.
* `packet_size` - (Optional) Type: `string`.
* `per_connection_classifier` - (Optional) Type: `string`.
* `port` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `string`.
* `random` - (Optional) Type: `string`.
* `src_address` - (Optional) Type: `string`.
* `src_address_list` - (Optional) Type: `string`.
* `src_address_type` - (Optional) Type: `string`.
* `src_mac_address` - (Optional) Type: `string`.
* `src_port` - (Optional) Type: `string`.
* `tcp_flags` - (Optional) Type: `string`.
* `tcp_mss` - (Optional) Type: `string`.
* `time` - (Optional) Type: `string`.
* `tls_host` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `bytes` - Type: `string`.
* `packets` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_firewall_raw.example '*3'

# Named router
terraform import routeros_ipv6_firewall_raw.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_firewall_raw.example 'home/my-resource-name'
```
