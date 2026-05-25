---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_ethernet"
description: |-
  /interface/ethernet entries are auto-created from the physical NICs; can't be added via TF.
---

# Resource: routeros_interface_ethernet

/interface/ethernet entries are auto-created from the physical NICs; can't be added via TF.

## Example Usage

```terraform
resource "routeros_interface_ethernet" "ethernet_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # advertise = []
  # arp = "disabled"
  # arp_timeout = "1h"
  # cable_settings = "replace-me"
  # combo_mode = "auto"
  # disable_running_check = false
  # fec_mode = "off"
  # interface = "ether1"
  # l2mtu = "replace-me"
  # loop_protect = "default"
  # loop_protect_disable_time = "1h"
  # loop_protect_send_interval = "1h"
  # mac_address = "10.99.0.0/24"
  # mtu = 1500
  # name = "tf-example"
  # orig_mac_address = "10.99.0.0/24"
  # published = "replace-me"
  # rx_flow_control = "off"
  # sfp_shutdown_temperature = 0
  # speed = "10M baseT half"
  # tx_flow_control = "off"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Optional) Type: `string`. IP address to be mapped.
* `advertise` - (Optional) Type: `list`.
* `arp` - (Optional) Type: `enum(disabled|enabled|proxy-arp|reply-only|local-proxy-arp)`.
* `arp_timeout` - (Optional) Type: `duration`.
* `cable_settings` - (Optional) Type: `string`.
* `combo_mode` - (Optional) Type: `enum(auto|copper|sfp)`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disable_running_check` - (Optional) Type: `bool`.
* `disabled` - (Optional) Type: `bool`.
* `fec_mode` - (Optional) Type: `enum(off|auto|fec74|fec91)`.
* `interface` - (Optional) Type: `string`. Interface name the IP address is assigned to.
* `l2mtu` - (Optional) Type: `string`.
* `loop_protect` - (Optional) Type: `enum(default|off|on)`.
* `loop_protect_disable_time` - (Optional) Type: `duration`.
* `loop_protect_send_interval` - (Optional) Type: `duration`.
* `mac_address` - (Optional) Type: `mac`. MAC address to be mapped to.
* `mtu` - (Optional) Type: `int`. Default: `1500`.
* `name` - (Optional) Type: `string`.
* `orig_mac_address` - (Optional) Type: `mac`.
* `published` - (Optional) Type: `string`. Static proxy-arp entry for individual IP addresses. When an ARP query is received for the specific IP address, the device will respond with its own MAC address. No need to set proxy-arp on the interface itself for all the MAC addresses to be proxied. The interface will respond to an ARP request only when the device has an active route towards the destination.
* `rx_flow_control` - (Optional) Type: `enum(off|on|auto)`.
* `sfp_shutdown_temperature` - (Optional) Type: `int`.
* `speed` - (Optional) Type: `enum(10M baseT half|10M baseT full|100M baseT half|100M baseT full|1G baseT half|1G baseT full, ...)`.
* `tx_flow_control` - (Optional) Type: `enum(off|on|auto)`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `auto_negotiation` - Type: `enum(incomplete|done|no negotiation|failed|restarted|disabled, ...)`.
* `default_name` - Type: `string`.
* `loop_protect_status` - Type: `string`.
* `running` - Type: `bool`.
* `rx_broadcast` - Type: `int`.
* `rx_bytes` - Type: `int`.
* `rx_multicast` - Type: `int`.
* `rx_packet` - Type: `int`.
* `tx_broadcast` - Type: `int`.
* `tx_bytes` - Type: `int`.
* `tx_multicast` - Type: `int`.
* `tx_packet` - Type: `int`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_ethernet.example '*3'

# Named router
terraform import routeros_interface_ethernet.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_ethernet.example 'home/my-resource-name'
```
