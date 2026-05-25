---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_vlan"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_vlan

Manages the RouterOS `/interface/vlan` menu.

## Example Usage

```terraform
data "routeros_interface_vlan" "vlan_example" {
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
* `arp` - (Optional) Type: `string`. Address Resolution Protocol setting disabled   - the interface will not use ARP enabled   - the interface will use ARP local-proxy-arp   -     the router performs proxy ARP on the interface and sends replies to the same interface proxy-arp   -   the router performs proxy ARP on the interface and sends replies to other interfaces reply-only   - the interface will only reply to requests originated from matching IP address/MAC address combinations which are entered as static entries in the   IP/ARP   table. No dynamic entries will be automatically stored in the   IP/ARP   table. Therefore for communications to be successful, a valid static entry must already exist.
* `arp_timeout` - (Optional) Type: `string`. How long the ARP record is kept in the ARP table after no packets are received from IP. Value   auto   equals to the value of   arp-timeout   in   IP/Settings, default is 30s.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`. Name of the interface on top of which VLAN will work. Adding a VLAN interface to a bridge with vlan-filtering enabled will automatically tag the bridge interface as a member port. A dynamic entry with the comment "added by vlan on bridge" will appear under the /interface/bridge/vlan menu.
* `l3_hw_offloading` - (Optional) Type: `string`. Enables or disabled L3HW on a per-VLAN interface. This setting is only applicable to devices that support L3HW offloading and is available starting from RouterOS v7.21. More details - Per-VLAN offloading .
* `mtu` - (Optional) Type: `string`. Layer3 Maximum transmission unit.
* `mvrp` - (Optional) Type: `string`. Specifies whether this VLAN should declare its attributes through Multiple VLAN Registration Protocol (MVRP) as an applicant. Its main use case is for VLANs that is created on Ethernet interface (such as a "router on a stick" setup) that is connected to a bridge supporting MVRP . Enabling this option on a VLAN interface that is already part of an MVRP-enabled bridge has no effect, as the bridge manages MVRP in that case.   This property only has an effect when use-service-tag   is disabled .
* `name` - (Required) Type: `string`. Interface name. Default: `tf_acc_vlan`.
* `use_service_tag` - (Optional) Type: `string`. IEEE 802.1ad compatible Service Tag.
* `vlan_id` - (Required) Type: `string`. Virtual LAN identifier or tag that is used to distinguish VLANs. Must be equal for all computers that belong to the same VLAN. Default: `100`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

