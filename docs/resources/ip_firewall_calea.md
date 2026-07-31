---
subcategory: "Firewall"
page_title: "RouterOS: routeros_ip_firewall_calea"
description: |-
  IP firewall CALEA creates a session that drops the management connection on CHR. Skipped.
---

# Resource: routeros_ip_firewall_calea

IP firewall CALEA creates a session that drops the management connection on CHR. Skipped.

## Example Usage

```terraform
resource "routeros_ip_firewall_calea" "calea_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `action` - (Optional) Type: `string`. RouterOS `action`.
* `address_list` - (Optional) Type: `string`. RouterOS `address-list`.
* `address_list_timeout` - (Optional) Type: `string`. RouterOS `address-list-timeout`.
* `chain` - (Optional) Type: `string`. RouterOS `chain`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connection_bytes` - (Optional) Type: `string`. RouterOS `connection-bytes`.
* `connection_limit` - (Optional) Type: `string`. RouterOS `connection-limit`.
* `connection_mark` - (Optional) Type: `string`. RouterOS `connection-mark`.
* `connection_rate` - (Optional) Type: `string`. RouterOS `connection-rate`.
* `connection_type` - (Optional) Type: `string`. RouterOS `connection-type`.
* `content` - (Optional) Type: `string`. RouterOS `content`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dscp` - (Optional) Type: `string`. RouterOS `dscp`.
* `dst_address` - (Optional) Type: `string`. RouterOS `dst-address`.
* `dst_address_list` - (Optional) Type: `string`. RouterOS `dst-address-list`.
* `dst_address_type` - (Optional) Type: `string`. RouterOS `dst-address-type`.
* `dst_limit` - (Optional) Type: `string`. RouterOS `dst-limit`.
* `dst_port` - (Optional) Type: `string`. RouterOS `dst-port`.
* `fragment` - (Optional) Type: `string`. RouterOS `fragment`.
* `hotspot` - (Optional) Type: `string`. RouterOS `hotspot`.
* `icmp_options` - (Optional) Type: `string`. RouterOS `icmp-options`.
* `in_bridge_port` - (Optional) Type: `string`. RouterOS `in-bridge-port`.
* `in_bridge_port_list` - (Optional) Type: `string`. RouterOS `in-bridge-port-list`.
* `in_interface` - (Optional) Type: `string`. RouterOS `in-interface`.
* `in_interface_list` - (Optional) Type: `string`. RouterOS `in-interface-list`.
* `ingress_priority` - (Optional) Type: `string`. RouterOS `ingress-priority`.
* `ipsec_policy` - (Optional) Type: `string`. RouterOS `ipsec-policy`.
* `ipv4_options` - (Optional) Type: `string`. RouterOS `ipv4-options`.
* `layer7_protocol` - (Optional) Type: `string`. RouterOS `layer7-protocol`.
* `limit` - (Optional) Type: `string`. RouterOS `limit`.
* `log` - (Optional) Type: `string`. RouterOS `log`.
* `log_prefix` - (Optional) Type: `string`. RouterOS `log-prefix`.
* `nth` - (Optional) Type: `string`. RouterOS `nth`.
* `out_bridge_port` - (Optional) Type: `string`. RouterOS `out-bridge-port`.
* `out_bridge_port_list` - (Optional) Type: `string`. RouterOS `out-bridge-port-list`.
* `out_interface` - (Optional) Type: `string`. RouterOS `out-interface`.
* `out_interface_list` - (Optional) Type: `string`. RouterOS `out-interface-list`.
* `packet_mark` - (Optional) Type: `string`. RouterOS `packet-mark`.
* `packet_size` - (Optional) Type: `string`. RouterOS `packet-size`.
* `per_connection_classifier` - (Optional) Type: `string`. RouterOS `per-connection-classifier`.
* `port` - (Optional) Type: `string`. RouterOS `port`.
* `priority` - (Optional) Type: `string`. RouterOS `priority`.
* `protocol` - (Optional) Type: `string`. RouterOS `protocol`.
* `psd` - (Optional) Type: `string`. RouterOS `psd`.
* `random` - (Optional) Type: `string`. RouterOS `random`.
* `realm` - (Optional) Type: `string`. RouterOS `realm`.
* `routing_mark` - (Optional) Type: `string`. RouterOS `routing-mark`.
* `sniff_id` - (Optional) Type: `string`. RouterOS `sniff-id`.
* `sniff_target` - (Optional) Type: `string`. RouterOS `sniff-target`.
* `sniff_target_port` - (Optional) Type: `string`. RouterOS `sniff-target-port`.
* `src_address` - (Optional) Type: `string`. RouterOS `src-address`.
* `src_address_list` - (Optional) Type: `string`. RouterOS `src-address-list`.
* `src_address_type` - (Optional) Type: `string`. RouterOS `src-address-type`.
* `src_mac_address` - (Optional) Type: `string`. RouterOS `src-mac-address`.
* `src_port` - (Optional) Type: `string`. RouterOS `src-port`.
* `tcp_mss` - (Optional) Type: `string`. RouterOS `tcp-mss`.
* `time` - (Optional) Type: `string`. RouterOS `time`.
* `tls_host` - (Optional) Type: `string`. RouterOS `tls-host`.
* `tos` - (Optional) Type: `string`. RouterOS `tos`.
* `ttl` - (Optional) Type: `string`. RouterOS `ttl`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_firewall_calea.example '*3'

# Named router
terraform import routeros_ip_firewall_calea.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_firewall_calea.example 'home/my-resource-name'
```
