---
subcategory: "Firewall"
page_title: "RouterOS: routeros_ip_firewall_filter"
description: |-
  IP firewall filter rule. Ordered top-down by `position` (sort key, not
---

# Resource: routeros_ip_firewall_filter

IP firewall filter rule, ordered one of two ways.

`position` is a top-down sort key (not an identity), ordering this rule only against other rules
managed by the same apply. It is persisted on the device via `[tf:pos=N]` in the comment, so a
destroy and re-apply rebuilds the same order.

`place_before` takes a RouterOS `.id` and orders the rule against any rule on the device, including
one owned by a different Terraform state or shipped as a RouterOS default. The two are mutually
exclusive.

Safety: refuses an unconditional `chain=input`/`forward` `action=drop|reject|tarpit` rule unless
`lockout_ack = true`. Note that the guard reads a rule's own content, not its position: a rule that
passes it can still be placed above an accept rule that Terraform does not manage.


## Example Usage

```terraform
resource "routeros_ip_firewall_filter" "filter_example" {
  # router = "my-router"  # which router to target; omit for the default
  action = "accept"
  chain = "forward"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address_list = "replace-me"
  # address_list_timeout = "replace-me"
  # connection_bytes = "replace-me"
  # connection_limit = "replace-me"
  # connection_mark = "replace-me"
  # connection_nat_state = "replace-me"
  # connection_rate = "replace-me"
  # connection_state = "replace-me"
  # connection_type = "replace-me"
  # content = "replace-me"
  # dscp = "replace-me"
  # dst_address = "10.99.0.0/24"
  # dst_address_list = "my-list"
  # dst_address_type = "replace-me"
  # dst_limit = "replace-me"
  # dst_port = "443"
  # fragment = "replace-me"
  # hotspot = "replace-me"
  # icmp_options = "replace-me"
  # in_bridge_port = "443"
  # in_bridge_port_list = "replace-me"
  # in_interface = "ether1"
  # in_interface_list = "LAN"
  # ingress_priority = "replace-me"
  # ipsec_policy = "replace-me"
  # ipv4_options = "replace-me"
  # jump_target = "replace-me"
  # layer7_protocol = "replace-me"
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
  # psd = "replace-me"
  # random = "replace-me"
  # realm = "replace-me"
  # reject_with = "replace-me"
  # routing_mark = "replace-me"
  # src_address = "10.99.0.0/24"
  # src_address_list = "my-list"
  # src_address_type = "replace-me"
  # src_mac_address = "10.99.0.0/24"
  # src_port = "443"
  # tcp_flags = "replace-me"
  # tcp_mss = "replace-me"
  # time = "replace-me"
  # tls_host = "replace-me"
  # ttl = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `action` - (Required) Type: `string`.
* `address_list` - (Optional) Type: `string`.
* `address_list_timeout` - (Optional) Type: `string`.
* `bytes` - (Optional) Type: `string`.
* `chain` - (Required) Type: `string`.
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
* `lockout_ack` - (Optional) Type: `bool`. Acknowledge that this rule may sever management traffic (required for unconditional input/forward drop/reject/tarpit rules with no match).
* `log` - (Optional) Type: `string`.
* `log_prefix` - (Optional) Type: `string`.
* `nth` - (Optional) Type: `string`.
* `out_bridge_port` - (Optional) Type: `string`.
* `out_bridge_port_list` - (Optional) Type: `string`.
* `out_interface` - (Optional) Type: `string`.
* `out_interface_list` - (Optional) Type: `string`.
* `p2p` - (Optional) Type: `string`. RouterOS `p2p`.
* `packet_mark` - (Optional) Type: `string`.
* `packet_size` - (Optional) Type: `string`.
* `packets` - (Optional) Type: `string`.
* `per_connection_classifier` - (Optional) Type: `string`.
* `port` - (Optional) Type: `string`.
* `position` - (Optional) Type: `int`. Sort key for placement in the ordered chain. Lower = higher in the chain. Persisted on the device via a [tf:pos=N] prefix in the comment so destroy+apply rebuilds the same order.
* `place_before` - (Optional) Type: `string`. RouterOS `.id` (e.g. `*3`) of the rule this one is inserted before, typically from a `data "routeros_ip_firewall_filter"` lookup. Unlike `position`, which only orders rules managed by the same apply, this orders against any rule on the device. Mutually exclusive with `position`.

  Two limits worth knowing. The anchor is resolved once, when `place_before` changes: if someone reorders the chain by hand, or a rule is inserted between the anchor and this one, the next plan reports no changes. And RouterOS reuses an `.id` after the rule holding it is deleted, so a stored anchor can come to name a different rule. Pin the anchor with a data-source lookup on a stable field such as `comment` rather than hardcoding an `.id`.
* `priority` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `string`.
* `psd` - (Optional) Type: `string`.
* `random` - (Optional) Type: `string`.
* `realm` - (Optional) Type: `string`.
* `reject_with` - (Optional) Type: `string`.
* `routing_mark` - (Optional) Type: `string`.
* `src_address` - (Optional) Type: `string`.
* `src_address_list` - (Optional) Type: `string`.
* `src_address_type` - (Optional) Type: `string`.
* `src_mac_address` - (Optional) Type: `string`.
* `src_port` - (Optional) Type: `string`.
* `tcp_flags` - (Optional) Type: `string`.
* `tcp_mss` - (Optional) Type: `string`.
* `time` - (Optional) Type: `string`.
* `tls_host` - (Optional) Type: `string`.
* `tos` - (Optional) Type: `string`. RouterOS `tos`.
* `ttl` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_firewall_filter.example '*3'

# Named router
terraform import routeros_ip_firewall_filter.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_firewall_filter.example 'home/my-resource-name'
```
