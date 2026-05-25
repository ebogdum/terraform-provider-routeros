---
page_title: "RouterOS: routeros_tool_sniffer_stop"
description: |-
  Action only; not CRUD
---

# Resource: routeros_tool_sniffer_stop

Action only; not CRUD

## Example Usage

```terraform
resource "routeros_tool_sniffer_stop" "stop_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # file_limit = "replace-me"
  # file_name = "replace-me"
  # filter_cpu = "replace-me"
  # filter_direction = "replace-me"
  # filter_dst_ip_address = "10.99.0.0/24"
  # filter_dst_ipv6_address = "10.99.0.0/24"
  # filter_dst_mac_address = "10.99.0.0/24"
  # filter_dst_port = "443"
  # filter_interface = "replace-me"
  # filter_ip_address = "10.99.0.0/24"
  # filter_ip_protocol = "replace-me"
  # filter_ipv6_address = "10.99.0.0/24"
  # filter_mac_address = "10.99.0.0/24"
  # filter_mac_protocol = "replace-me"
  # filter_operator_between_entries = "replace-me"
  # filter_port = "443"
  # filter_size = "replace-me"
  # filter_src_ip_address = "10.99.0.0/24"
  # filter_src_ipv6_address = "10.99.0.0/24"
  # filter_src_mac_address = "10.99.0.0/24"
  # filter_src_port = "443"
  # filter_stream = "replace-me"
  # filter_vlan = "replace-me"
  # memory_limit = "replace-me"
  # memory_scroll = "replace-me"
  # only_headers = "replace-me"
  # show_frame = "replace-me"
  # streaming_enabled = "replace-me"
  # streaming_port = "443"
  # streaming_server = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `file_limit` - (Optional) Type: `string`. File size limit. Sniffer will stop when a limit is reached.
* `file_name` - (Optional) Type: `string`. Name of the file where sniffed packets will be saved.
* `filter_cpu` - (Optional) Type: `string`. CPU core used as a filter.
* `filter_direction` - (Optional) Type: `string`. Specifies which direction filtering will be applied.
* `filter_dst_ip_address` - (Optional) Type: `string`. Up to 16 IP destination addresses used as a filter.
* `filter_dst_ipv6_address` - (Optional) Type: `string`. Up to 16 IPv6 destination addresses used as a filter.
* `filter_dst_mac_address` - (Optional) Type: `string`. Up to 16 MAC destination addresses and MAC address masks used as a filter.
* `filter_dst_port` - (Optional) Type: `string`. Up to 16 comma-separated destination ports used as a filter. A list of predefined port names is also available, like ssh and telnet.
* `filter_interface` - (Optional) Type: `string`. Interface name on which sniffer will be running.   all   indicates that the sniffer will sniff packets on all interfaces.
* `filter_ip_address` - (Optional) Type: `string`. Up to 16 IP addresses used as a filter.
* `filter_ip_protocol` - (Optional) Type: `string`. Up to 16 comma-separated IP/IPv6 protocols used as a filter. IP protocols (instead of protocol names, protocol numbers can be used): ipsec-ah   - IPsec AH protocol ipsec-esp   - IPsec ESP protocol ddp   - datagram delivery protocol egp   - exterior gateway protocol ggp   - gateway-gateway protocol gre   - general routing encapsulation hmp   - host monitoring protocol idpr-cmtp   - idpr control message transport icmp   - internet control message protocol icmpv6   - internet control message protocol v6 igmp   - internet group management protocol ipencap   - ip encapsulated in ip ipip   - ip encapsulation encap   - ip encapsulation iso-tp4   - iso transport protocol class 4 ospf   - open shortest path first pup   - parc universal packet protocol pim   - protocol independent multicast rspf   - radio shortest path first rdp   - reliable datagram protocol st   - st datagram mode tcp   - transmission control protocol udp   - user datagram protocol vmtp   - versatile message transport vrrp   - virtual router redundancy protocol xns-idp   - xerox xns idp xtp   - xpress transfer protocol.
* `filter_ipv6_address` - (Optional) Type: `string`. Up to 16 IPv6 addresses used as a filter.
* `filter_mac_address` - (Optional) Type: `string`. Up to 16 MAC addresses and MAC address masks used as a filter.
* `filter_mac_protocol` - (Optional) Type: `string`. Up to 16 comma separated entries used as a filter. Mac protocols (instead of protocol names, protocol number can be used): 802.2   - 802.2 Frames (0x0004) arp   - Address Resolution Protocol (0x0806) capsman - CAPsMAN to CAP MAC layer connection ( 0x88BB ) dot1x - EAPoL IEEE 802.1X ( 0x888E ) homeplug-av   - HomePlug AV MME (0x88E1) ip   - Internet Protocol version 4 (0x0800) ipv6   - Internet Protocol Version 6 (0x86DD) ipx   - Internetwork Packet Exchange (0x8137) lacp - Link Aggregation Control Protocol ( 0x8809 ) lldp   - Link Layer Discovery Protocol (0x88CC) loop-protect   - Loop Protect Protocol (0x9003) macsec - MAC security IEEE 802.1AE (0x88E5) mpls-multicast   - MPLS multicast (0x8848) mpls-unicast   - MPLS unicast (0x8847) mvrp - Multiple VLAN Registration protocol (0x88F5) packing-compr   - Encapsulated packets with compressed   IP packing   (0x9001) packing-simple   - Encapsulated packets with simple   IP packing   (0x9000) pppoe   - PPPoE Session Stage (0x8864) pppoe-discovery   - PPPoE Discovery Stage (0x8863) rarp   - Reverse Address Resolution Protocol (0x8035) romon - Router Management Overlay Network RoMON ( 0x88BF ) service-vlan   - Provider Bridging (IEEE 802.1ad) & Shortest Path Bridging IEEE 802.1aq (0x88A8) vlan   - VLAN-tagged frame (IEEE 802.1Q) and Shortest Path Bridging IEEE 802.1aq with NNI compatibility (0x8100).
* `filter_operator_between_entries` - (Optional) Type: `string`. Changes the logic for filters with multiple entries.
* `filter_port` - (Optional) Type: `string`. Up to 16 comma-separated ports used as a filter. A list of predefined port names is also available, like ssh and telnet.
* `filter_size` - (Optional) Type: `string`. Filters packets of specified size or size range in bytes.
* `filter_src_ip_address` - (Optional) Type: `string`. Up to 16 IP source addresses used as a filter.
* `filter_src_ipv6_address` - (Optional) Type: `string`. Up to 16 IPv6 source addresses used as a filter.
* `filter_src_mac_address` - (Optional) Type: `string`. Up to 16 MAC source addresses and MAC address masks used as a filter.
* `filter_src_port` - (Optional) Type: `string`. Up to 16 comma-separated source ports used as a filter. A list of predefined port names is also available, like ssh and telnet.
* `filter_stream` - (Optional) Type: `string`. Sniffed packets that are devised for the sniffer server are ignored.
* `filter_vlan` - (Optional) Type: `string`. Up to 16 VLAN IDs used as a filter.
* `memory_limit` - (Optional) Type: `string`. Memory amount used to store sniffed data.
* `memory_scroll` - (Optional) Type: `string`. Whether to rewrite older sniffed data when the memory limit is reached.
* `only_headers` - (Optional) Type: `string`. Save in the memory only the packet's headers, not the whole packet.
* `show_frame` - (Optional) Type: `string`. Whether to see the content of the frame when running quick sniffer in command line.
* `streaming_enabled` - (Optional) Type: `string`. Defines whether to send sniffed packets to the streaming server.
* `streaming_port` - (Optional) Type: `string`. Port to stream the TZSP packets to.
* `streaming_server` - (Optional) Type: `string`. Tazmen Sniffer Protocol (TZSP) stream receiver.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_tool_sniffer_stop.example '*3'

# Named router
terraform import routeros_tool_sniffer_stop.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_tool_sniffer_stop.example 'home/my-resource-name'
```
