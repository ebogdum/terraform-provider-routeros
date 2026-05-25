---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_eoip"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_eoip

Manages the RouterOS `/interface/eoip` menu.

## Example Usage

```terraform
data "routeros_interface_eoip" "eoip_example" {
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
* `allow_fast_path` - (Optional) Type: `bool`. Whether to allow FastPath processing. Must be disabled if IPsec tunneling is used. Default: `1`.
* `arp` - (Optional) Type: `enum(disabled|enabled|proxy-arp|reply-only|local-proxy-arp)`. Address Resolution Protocol mode. disabled - the interface will not use ARP enabled - the interface will use ARP proxy-arp - the interface will use the ARP proxy feature reply-only - the interface will only reply to requests originated from matching IP address/MAC address combinations which are entered as static entries in the "/ip arp" table. No dynamic entries will be automatically stored in the "/ip arp" table. Therefore for communications to be successful, a valid static entry must already exist. Default: `1`.
* `arp_timeout` - (Optional) Type: `duration`. Time interval in which ARP entries should time out.
* `clamp_tcp_mss` - (Optional) Type: `bool`. Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP overhead).The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is changed. Default: `1`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disable_time` - (Optional) Type: `duration`. Default: `300`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dont_fragment` - (Optional) Type: `enum(no|inherit)`. Whether to include DF bit in related packets: no   - fragment if needed,   inherit   - use Dont Fragment flag of original packet. (Without Dont Fragment: inherit - packet may be fragmented).
* `dscp` - (Optional) Type: `enum(inherit)`. DSCP value of packet. Inherited option means that dscp value will be inherited from packet which is going to be encapsulated. Default: `256`.
* `ipsec_secret` - (Optional) Type: `string`. When secret is specified, router adds dynamic IPsec peer to remote-address with pre-shared key and policy (by default phase2 uses sha1/aes128cbc). **Sensitive.**
* `keepalive` - (Optional) Type: `string`. Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries - number of retry attempts. By default keepalive is set to 10 seconds and 10 retries. Default: `1`.
* `local_address` - (Optional) Type: `ip`. Source address of the tunnel packets, local on the router.
* `loop_protect` - (Optional) Type: `enum(default|off|on)`.
* `loop_protect_disable_time` - (Optional) Type: `string`.
* `loop_protect_send_interval` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`. Media Access Control number of an interface. The address numeration authority IANA allows the use of MAC addresses in the range from 00:00:5E:80:00:00 - 00:00:5E:FF:FF:FF freely.
* `mtu` - (Optional) Type: `int`. Layer3 Maximum transmission unit. Default: `0`.
* `name` - (Required) Type: `string`. Interface name. Default: `tf_acc_eoip`.
* `remote_address` - (Required) Type: `string`. IP address of remote end of EoIP tunnel. Default: `10.255.255.1`.
* `send_interval` - (Optional) Type: `duration`. Default: `5`.
* `tunnel_id` - (Required) Type: `int`. Unique tunnel identifier, which must match other side of the tunnel. Default: `1`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `actual_mtu` - Type: `int`.
* `status` - Type: `enum(|off|on|disabled)`.

