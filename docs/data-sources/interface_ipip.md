---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ipip"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_ipip

Manages the RouterOS `/interface/ipip` menu.

## Example Usage

```terraform
data "routeros_interface_ipip" "ipip_example" {
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

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`. **Marked sensitive**: this menu holds a secret, which RouterOS returns in the row like any other column, so an unprojected read puts it in your state file. Use `proplist` to name the columns you need.

