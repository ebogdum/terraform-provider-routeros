---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_eoip"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_eoip

Manages the RouterOS `/interface/eoip` menu.

## Example Usage

```terraform
resource "routeros_interface_eoip" "eoip_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"
  remote_address = "10.99.0.1"
  tunnel_id = 1

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # allow_fast_path = true
  # arp = "enabled"
  # arp_timeout = "1h"
  # clamp_tcp_mss = true
  # disable_time = "300"
  # dont_fragment = "no"
  # dscp = "inherit"
  # ipsec_secret = "REDACTED"
  # keepalive = "1"
  # local_address = "10.99.0.1"
  # loop_protect = "default"
  # loop_protect_disable_time = "replace-me"
  # loop_protect_send_interval = "replace-me"
  # mac_address = "10.99.0.0/24"
  # mtu = 0
  # send_interval = "5"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `actual_mtu` - (Read-only) Type: `int`.
* `allow_fast_path` - (Optional) Type: `bool`. Whether to allow FastPath processing. Must be disabled if IPsec tunneling is used.
* `arp` - (Optional) Type: `string`. Address Resolution Protocol mode. disabled - the interface will not use ARP enabled - the interface will use ARP proxy-arp - the interface will use the ARP proxy feature reply-only - the interface will only reply to requests originated from matching IP address/MAC address combinations which are entered as static entries in the "/ip arp" table. No dynamic entries will be automatically stored in the "/ip arp" table. Therefore for communications to be successful, a valid static entry must already exist.
* `arp_timeout` - (Optional) Type: `string`. Time interval in which ARP entries should time out.
* `clamp_tcp_mss` - (Optional) Type: `bool`. Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP overhead).The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is changed.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disable_time` - (Read-only) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dont_fragment` - (Optional) Type: `string`. Whether to include DF bit in related packets: no - fragment if needed, inherit - use Dont Fragment flag of original packet. (Without Dont Fragment: inherit - packet may be fragmented).
* `dscp` - (Optional) Type: `string`. DSCP value of packet. Inherited option means that dscp value will be inherited from packet which is going to be encapsulated.
* `ipsec_secret` - (Optional) Type: `string`. When secret is specified, router adds dynamic IPsec peer to remote-address with pre-shared key and policy (by default phase2 uses sha1/aes128cbc). **Sensitive.**
* `keepalive` - (Optional) Type: `string`. Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries - number of retry attempts. By default keepalive is set to 10 seconds and 10 retries.
* `local_address` - (Optional) Type: `string`. Source address of the tunnel packets, local on the router.
* `loop_protect` - (Optional) Type: `string`.
* `loop_protect_disable_time` - (Optional) Type: `string`.
* `loop_protect_send_interval` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`. Media Access Control number of an interface. The address numeration authority IANA allows the use of MAC addresses in the range from 00:00:5E:80:00:00 - 00:00:5E:FF:FF:FF freely
* `mtu` - (Optional) Type: `string`. Layer3 Maximum transmission unit A number, or `auto`.
* `name` - (Required) Type: `string`. Interface name
* `remote_address` - (Required) Type: `string`. IP address of remote end of EoIP tunnel
* `send_interval` - (Read-only) Type: `string`.
* `status` - (Read-only) Type: `string`.
* `tunnel_id` - (Required) Type: `int`. Unique tunnel identifier, which must match other side of the tunnel

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_eoip.example '*3'

# Named router
terraform import routeros_interface_eoip.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_eoip.example 'home/my-resource-name'
```
