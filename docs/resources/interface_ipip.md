---
page_title: "RouterOS: routeros_interface_ipip"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_ipip

Manages the RouterOS `/interface/ipip` menu.

## Example Usage

```terraform
resource "routeros_interface_ipip" "ipip_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # clamp_tcp_mss = true
  # dont_fragment = "no"
  # dscp = "inherit"
  # ipsec_secret = "REDACTED"
  # keepalive = "replace-me"
  # local_address = "10.99.0.1"
  # mtu = 0
  # name = "tf-example"
  # remote_address = "10.99.0.1"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `clamp_tcp_mss` - (Optional) Type: `bool`. Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP overhead).The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is changed. Default: `1`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dont_fragment` - (Optional) Type: `enum(no|inherit)`. Whether to include DF bit in related packets: no   - fragment if needed,   inherit   - use Dont Fragment flag of original packet. (Without Dont Fragment: inherit - packet may be fragmented).
* `dscp` - (Optional) Type: `enum(inherit)`. Set dscp value in IPIP header to a fixed value or inherit from dscp value taken from tunnelled traffic. Default: `256`.
* `ipsec_secret` - (Optional) Type: `string`. When secret is specified, router adds dynamic ipsec peer to remote-address with pre-shared key and policy with default values (by default phase2 uses sha1/aes128cbc). **Sensitive.**
* `keepalive` - (Optional) Type: `string`. Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries - number of retry attempts. By default keepalive is set to 10 seconds and 10 retries. To disable set *set ipipv6-tunnel1 !keepalive.
* `local_address` - (Optional) Type: `ip`. IP address on a router that will be used by IPIP tunnel.
* `mtu` - (Optional) Type: `int`. Layer3 Maximum transmission unit. Default: `0`.
* `name` - (Optional) Type: `string`. Interface name.
* `remote_address` - (Optional) Type: `string`. IP address of remote end of IPIP tunnel.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_ipip.example '*3'

# Named router
terraform import routeros_interface_ipip.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_ipip.example 'home/my-resource-name'
```
