---
page_title: "RouterOS: routeros_ip_firewall_mangle"
description: |-
  IP firewall mangle rule. Ordered by `position` (sort key, not identity).
---

# Data Source: routeros_ip_firewall_mangle

IP firewall mangle rule. Ordered by `position` (sort key, not identity).

## Example Usage

```terraform
data "routeros_ip_firewall_mangle" "mangle_example" {
  # router   = "my-router"  # omit for the default router
  # filter   = { name = "some-name" }
  # proplist = ["name", "address"]
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `proplist` - (Optional) List of property names to project; smaller payload.
* `action` - (Required) Type: `string`. Default: `accept`.
* `address_list` - (Optional) Type: `string`.
* `address_list_timeout` - (Optional) Type: `string`.
* `chain` - (Required) Type: `string`. Default: `prerouting`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connection_bytes` - (Optional) Type: `string`.
* `connection_limit` - (Optional) Type: `string`.
* `connection_mark` - (Optional) Type: `string`.
* `connection_nat_state` - (Optional) Type: `string`.
* `connection_rate` - (Optional) Type: `string`.
* `connection_state` - (Optional) Type: `string`.
* `connection_type` - (Optional) Type: `string`.
* `content` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dscp` - (Optional) Type: `string`.
* `dst_address` - (Optional) Type: `string`.
* `dst_address_list` - (Optional) Type: `string`.
* `dst_address_type` - (Optional) Type: `string`.
* `dst_limit` - (Optional) Type: `string`.
* `dst_port` - (Optional) Type: `string`.
* `fragment` - (Optional) Type: `string`.
* `hotspot` - (Optional) Type: `string`.
* `icmp_options` - (Optional) Type: `string`.
* `in_bridge_port` - (Optional) Type: `string`.
* `in_bridge_port_list` - (Optional) Type: `string`.
* `in_interface` - (Optional) Type: `string`.
* `in_interface_list` - (Optional) Type: `string`.
* `ingress_priority` - (Optional) Type: `string`.
* `ipsec_policy` - (Optional) Type: `string`.
* `ipv4_options` - (Optional) Type: `string`.
* `jump_target` - (Optional) Type: `string`.
* `layer7_protocol` - (Optional) Type: `string`.
* `limit` - (Optional) Type: `string`.
* `log` - (Optional) Type: `string`.
* `log_prefix` - (Optional) Type: `string`.
* `new_connection_mark` - (Optional) Type: `string`.
* `new_dscp` - (Optional) Type: `string`.
* `new_mss` - (Optional) Type: `string`.
* `new_packet_mark` - (Optional) Type: `string`.
* `new_priority` - (Optional) Type: `string`.
* `new_routing_mark` - (Optional) Type: `string`.
* `new_ttl` - (Optional) Type: `string`.
* `nth` - (Optional) Type: `string`.
* `out_bridge_port` - (Optional) Type: `string`.
* `out_bridge_port_list` - (Optional) Type: `string`.
* `out_interface` - (Optional) Type: `string`.
* `out_interface_list` - (Optional) Type: `string`.
* `packet_mark` - (Optional) Type: `string`.
* `packet_size` - (Optional) Type: `string`.
* `passthrough` - (Optional) Type: `string`.
* `per_connection_classifier` - (Optional) Type: `string`.
* `port` - (Optional) Type: `string`.
* `priority` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `string`.
* `psd` - (Optional) Type: `string`.
* `random` - (Optional) Type: `string`.
* `realm` - (Optional) Type: `string`.
* `route_dst` - (Optional) Type: `string`.
* `routing_mark` - (Optional) Type: `string`.
* `sniff_id` - (Optional) Type: `string`.
* `sniff_target` - (Optional) Type: `string`.
* `sniff_target_port` - (Optional) Type: `string`.
* `src_address` - (Optional) Type: `string`.
* `src_address_list` - (Optional) Type: `string`.
* `src_address_type` - (Optional) Type: `string`.
* `src_mac_address` - (Optional) Type: `string`.
* `src_port` - (Optional) Type: `string`.
* `tcp_flags` - (Optional) Type: `string`.
* `tcp_mss` - (Optional) Type: `string`.
* `time` - (Optional) Type: `string`.
* `tls_host` - (Optional) Type: `string`.
* `ttl` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `bytes` - Type: `string`.
* `packets` - Type: `string`.

