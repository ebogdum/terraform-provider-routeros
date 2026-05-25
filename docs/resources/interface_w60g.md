---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_w60g"
description: |-
  Requires 60GHz wAP60G hardware
---

# Resource: routeros_interface_w60g

Requires 60GHz wAP60G hardware

## Example Usage

```terraform
resource "routeros_interface_w60g" "w60g_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # frequency = "replace-me"
  # isolate_stations = "replace-me"
  # l2mtu = "replace-me"
  # mac_address = "10.99.0.0/24"
  # mdmg_fix = "replace-me"
  # mode = "replace-me"
  # mtu = "replace-me"
  # name = "tf-example"
  # password = "REDACTED"
  # put_stations_in_bridge = "replace-me"
  # region = "replace-me"
  # scan_list = "replace-me"
  # ssid = "replace-me"
  # tx_sector = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `arp` - (Optional) Type: `string`. Read more >>.
* `arp_timeout` - (Optional) Type: `string`. ARP timeout is time how long ARP record is kept in ARP table after no packets are received from IP. Value auto equals to the value of arp-timeout in /ip settings , default is 30s.
* `comment` - (Optional) Type: `string`. Short description of the interface.
* `disabled` - (Optional) Type: `string`. Whether interface is disabled.
* `frequency` - (Optional) Type: `string`. Frequency used in communication (Only active on bridge device).
* `isolate_stations` - (Optional) Type: `string`. Don't allow communication between connected clients (from RouterOS 6.41).
* `l2mtu` - (Optional) Type: `string`. Layer2 Maximum transmission unit.
* `mac_address` - (Optional) Type: `string`. MAC address of the radio interface.
* `mdmg_fix` - (Optional) Type: `string`. Experimental feature working only on wAP60Gx3 devices, providing better point to multi point stability in some cases.
* `mode` - (Optional) Type: `string`. Operation mode.
* `mtu` - (Optional) Type: `string`. Layer3 Maximum transmission unit.
* `name` - (Optional) Type: `string`. Name of the interface.
* `password` - (Optional) Type: `string`. Password used for AES encryption. **Sensitive.**
* `put_stations_in_bridge` - (Optional) Type: `string`. Put newly created station device interfaces in this bridge.
* `region` - (Optional) Type: `string`. Parameter to limit frequency use.
* `scan_list` - (Optional) Type: `string`. Scan list to limit connectivity over frequencies in station mode.
* `ssid` - (Optional) Type: `string`. SSID (service set identifier) is a name that identifies wireless network.
* `tx_sector` - (Optional) Type: `string`. Disables beamforming and locks to selected radiation pattern.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_w60g.example '*3'

# Named router
terraform import routeros_interface_w60g.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_w60g.example 'home/my-resource-name'
```
