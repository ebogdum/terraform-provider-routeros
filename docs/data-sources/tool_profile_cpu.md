---
page_title: "RouterOS: routeros_tool_profile_cpu"
description: |-
  Long-running monitor command, not CRUD
---

# Data Source: routeros_tool_profile_cpu

Long-running monitor command, not CRUD

## Example Usage

```terraform
data "routeros_tool_profile_cpu" "cpu_example" {
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
* `backup` - (Optional) Type: `string`. Backup service.
* `bfd` - (Optional) Type: `string`. BFD service.
* `bgp` - (Optional) Type: `string`. BGP service.
* `bridging` - (Optional) Type: `string`. Bridging service.
* `btest` - (Optional) Type: `string`. Bandwidth test.
* `certificate` - (Optional) Type: `string`. Certificate service.
* `console` - (Optional) Type: `string`. Console.
* `container` - (Optional) Type: `string`. combined container usage.
* `dhcp` - (Optional) Type: `string`. DHCP-Server and DHCP-Client services.
* `disk` - (Optional) Type: `string`. storage-related services.
* `dns` - (Optional) Type: `string`. DNS-related services.
* `dude` - (Optional) Type: `string`. The Dude package services.
* `e_mail` - (Optional) Type: `string`. e-mail tool.
* `encrypting` - (Optional) Type: `string`. encrypting processes.
* `eoip` - (Optional) Type: `string`. EoIP.
* `ethernet` - (Optional) Type: `string`. Ethernet-related properties like link speed, auto-negotiation, duplex mode, monitor a transceiver diagnostic information, etc.
* `fetcher` - (Optional) Type: `string`. Fetch tool.
* `fileman` - (Optional) Type: `string`. File manager.
* `firewall` - (Optional) Type: `string`. Firewall-related processes.
* `firewall_mgmt` - (Optional) Type: `string`. Firewall Management: Filtering, NAT, Mangle.
* `flash` - (Optional) Type: `string`. storage-related services.
* `ftp` - (Optional) Type: `string`. FTP Service.
* `gps` - (Optional) Type: `string`. GPS Service.
* `graphing` - (Optional) Type: `string`. Graphing tool.
* `gre` - (Optional) Type: `string`. GRE.
* `health` - (Optional) Type: `string`. system monitoring, workd health.
* `hotspot` - (Optional) Type: `string`. Hotspot service.
* `idle` - (Optional) Type: `string`. Free CPU resources.
* `igmp_proxy` - (Optional) Type: `string`. IGMP Proxy service.
* `internet_detect` - (Optional) Type: `string`. Detect Internet tool.
* `ip_pool` - (Optional) Type: `string`. IP Pool service.
* `ipsec` - (Optional) Type: `string`. IPsec service: xfrm -  set of statistics showing numbers of packets dropped by the transformation code and why.  drivers/crypto - drivers that provide access to the hardware cryptographic accelerators. ipsec - processes that relate to the Internet Key Exchange (IKE) protocols, Authentication Header (AH), Encapsulating Security Payload (ESP).
* `kvm` - (Optional) Type: `string`. KVM virtual machine functionality.
* `l7_matcher` - (Optional) Type: `string`. L7 matcher.
* `lcd` - (Optional) Type: `string`. LCD Interfaces system.
* `ldp` - (Optional) Type: `string`. Label Distribution Protocol (LDP).
* `logging` - (Optional) Type: `string`. Logging system.
* `management` - (Optional) Type: `string`. different subsystems: scheduler, networking, file management, etc.
* `mpls` - (Optional) Type: `string`. MPLS-related features.
* `neighbour_discovery` - (Optional) Type: `string`. Neighbour discovery service.
* `networking` - (Optional) Type: `string`. common set of services included in the networking.
* `ntp` - (Optional) Type: `string`. NTP service.
* `ospf` - (Optional) Type: `string`. OSPF service.
* `ovpn` - (Optional) Type: `string`. OVPN service.
* `pim` - (Optional) Type: `string`. Protocol Independent Multicast.
* `profiling` - (Optional) Type: `string`. Profiler service.
* `queue_mgmt` - (Optional) Type: `string`. Queues: Simple queues, Queue tree, Queue types.
* `queuing` - (Optional) Type: `string`. Intermediate Queuing.
* `radius` - (Optional) Type: `string`. RADIUS service.
* `radv` - (Optional) Type: `string`. IPv6 radv daemon log messages service.
* `remote_access` - (Optional) Type: `string`. accessing the device directly without logging into RouterOS.
* `rip` - (Optional) Type: `string`. Routing Information Protocol.
* `routing` - (Optional) Type: `string`. Routing-related services.
* `serial` - (Optional) Type: `string`. serial console and terminal tool.
* `sniffing` - (Optional) Type: `string`. packet Sniffer tool.
* `snmp` - (Optional) Type: `string`. SNMP.
* `socks` - (Optional) Type: `string`. Socket Secure.
* `spi` - (Optional) Type: `string`. storage-related services.
* `ssh` - (Optional) Type: `string`. SSH Server.
* `ssl` - (Optional) Type: `string`. SSL.
* `telnet` - (Optional) Type: `string`. Telnet service.
* `tftp` - (Optional) Type: `string`. TFTP service.
* `traffic_accounting` - (Optional) Type: `string`. Traffic-Flow log system.
* `traffic_flow` - (Optional) Type: `string`. Traffic-Flow system.
* `unclassified` - (Optional) Type: `string`. processes or services that are not defined by this classifier.
* `upnp` - (Optional) Type: `string`. UPnP protocol.
* `usb` - (Optional) Type: `string`. USB features.
* `user_manager` - (Optional) Type: `string`. User Manager service.
* `vrrp` - (Optional) Type: `string`. VRRP.
* `web_proxy` - (Optional) Type: `string`. Web Proxy.
* `winbox` - (Optional) Type: `string`. Winbox.
* `wireguard` - (Optional) Type: `string`. Wireguard.
* `wireless` - (Optional) Type: `string`. common set of services using Wireless systems.
* `www` - (Optional) Type: `string`. Webfig HTTP service.
* `zerotier` - (Optional) Type: `string`. ZeroTier.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

