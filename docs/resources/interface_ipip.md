---
subcategory: "Interfaces"
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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `allow_fast_path` - (Optional) Type: `string`. RouterOS `allow-fast-path`.
* `clamp_tcp_mss` - (Optional) Type: `bool`. Controls whether to change MSS size for received TCP SYN packets. When enabled, a router will change the MSS size for received TCP SYN packets if the current MSS size exceeds the tunnel interface MTU (taking into account the TCP/IP overhead).The received encapsulated packet will still contain the original MSS, and only after decapsulation the MSS is changed.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dont_fragment` - (Optional) Type: `string`. Whether to include DF bit in related packets: no - fragment if needed, inherit - use Dont Fragment flag of original packet. (Without Dont Fragment: inherit - packet may be fragmented).
* `dscp` - (Optional) Type: `string`. Set dscp value in IPIP header to a fixed value or inherit from dscp value taken from tunnelled traffic.
* `ipsec_secret` - (Optional) Type: `string`. When secret is specified, router adds dynamic ipsec peer to remote-address with pre-shared key and policy with default values (by default phase2 uses sha1/aes128cbc). **Sensitive.**
* `keepalive` - (Optional) Type: `string`. Tunnel keepalive parameter sets the time interval in which the tunnel running flag will remain even if the remote end of tunnel goes down. If configured time,retries fail, interface running flag is removed. Parameters are written in following format: KeepaliveInterval,KeepaliveRetries where KeepaliveInterval is time interval and KeepaliveRetries - number of retry attempts. By default keepalive is set to 10 seconds and 10 retries. To disable set *set ipipv6-tunnel1 !keepalive.
* `local_address` - (Optional) Type: `string`. IP address on a router that will be used by IPIP tunnel.
* `mtu` - (Optional) Type: `string`. Layer3 Maximum transmission unit. A number, or `auto`.
* `name` - (Optional) Type: `string`. Interface name.
* `remote_address` - (Optional) Type: `string`. IP address of remote end of IPIP tunnel.

## Attribute Reference

* `id` - RouterOS internal .id.


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
