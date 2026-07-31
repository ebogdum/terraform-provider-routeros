---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_traffic_flow_ipfix"
description: |-
  Mirrors RouterOS /ip/traffic-flow/ipfix.
---

# Resource: routeros_ip_traffic_flow_ipfix

Mirrors RouterOS `/ip/traffic-flow/ipfix`.

## Example Usage

```terraform
resource "routeros_ip_traffic_flow_ipfix" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # bytes = true
  # dst_address = true
  # dst_address_mask = true
  # dst_mac_address = true
  # dst_port = true
  # first_forwarded = true
  # gateway = true
  # icmp_code = true
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `bytes` - (Optional) Type: `bool`. RouterOS `bytes`.
* `dst_address` - (Optional) Type: `bool`. RouterOS `dst-address`.
* `dst_address_mask` - (Optional) Type: `bool`. RouterOS `dst-address-mask`.
* `dst_mac_address` - (Optional) Type: `bool`. RouterOS `dst-mac-address`.
* `dst_port` - (Optional) Type: `bool`. RouterOS `dst-port`.
* `first_forwarded` - (Optional) Type: `bool`. RouterOS `first-forwarded`.
* `gateway` - (Optional) Type: `bool`. RouterOS `gateway`.
* `icmp_code` - (Optional) Type: `bool`. RouterOS `icmp-code`.
* `icmp_type` - (Optional) Type: `bool`. RouterOS `icmp-type`.
* `igmp_type` - (Optional) Type: `bool`. RouterOS `igmp-type`.
* `in_interface` - (Optional) Type: `bool`. RouterOS `in-interface`.
* `ip_header_length` - (Optional) Type: `bool`. RouterOS `ip-header-length`.
* `ip_total_length` - (Optional) Type: `bool`. RouterOS `ip-total-length`.
* `ipv6_flow_label` - (Optional) Type: `bool`. RouterOS `ipv6-flow-label`.
* `is_multicast` - (Optional) Type: `bool`. RouterOS `is-multicast`.
* `last_forwarded` - (Optional) Type: `bool`. RouterOS `last-forwarded`.
* `nat_dst_address` - (Optional) Type: `bool`. RouterOS `nat-dst-address`.
* `nat_dst_port` - (Optional) Type: `bool`. RouterOS `nat-dst-port`.
* `nat_events` - (Optional) Type: `bool`. RouterOS `nat-events`.
* `nat_src_address` - (Optional) Type: `bool`. RouterOS `nat-src-address`.
* `nat_src_port` - (Optional) Type: `bool`. RouterOS `nat-src-port`.
* `out_interface` - (Optional) Type: `bool`. RouterOS `out-interface`.
* `packets` - (Optional) Type: `bool`. RouterOS `packets`.
* `protocol` - (Optional) Type: `bool`. RouterOS `protocol`.
* `src_address` - (Optional) Type: `bool`. RouterOS `src-address`.
* `src_address_mask` - (Optional) Type: `bool`. RouterOS `src-address-mask`.
* `src_mac_address` - (Optional) Type: `bool`. RouterOS `src-mac-address`.
* `src_port` - (Optional) Type: `bool`. RouterOS `src-port`.
* `sys_init_time` - (Optional) Type: `bool`. RouterOS `sys-init-time`.
* `tcp_ack_num` - (Optional) Type: `bool`. RouterOS `tcp-ack-num`.
* `tcp_flags` - (Optional) Type: `bool`. RouterOS `tcp-flags`.
* `tcp_seq_num` - (Optional) Type: `bool`. RouterOS `tcp-seq-num`.
* `tcp_window_size` - (Optional) Type: `bool`. RouterOS `tcp-window-size`.
* `tos` - (Optional) Type: `bool`. RouterOS `tos`.
* `ttl` - (Optional) Type: `bool`. RouterOS `ttl`.
* `udp_length` - (Optional) Type: `bool`. RouterOS `udp-length`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_traffic_flow_ipfix.this 'home'
```
