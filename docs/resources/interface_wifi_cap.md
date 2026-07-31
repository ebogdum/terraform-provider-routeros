---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_cap"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_cap

Manages the RouterOS `/interface/wifi/cap` menu.

## Example Usage

```terraform
resource "routeros_interface_wifi_cap" "cap_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # enabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `caps_man_addresses` - (Optional) Type: `string`. Comma-separated list of CAPsMAN controller addresses to connect to. Without this the CAP relies on discovery over `discovery_interfaces`.
* `caps_man_certificate_common_names` - (Optional) Type: `string`. Comma-separated list of controller certificate CommonNames to accept.
* `caps_man_names` - (Optional) Type: `string`. Comma-separated list of CAPsMAN controller names to accept.
* `certificate` - (Optional) Type: `string`. Certificate used towards the controller: a certificate name, `none`, or `request`.
* `discovery_interfaces` - (Optional) Type: `string`. Comma-separated list of interfaces used to discover a CAPsMAN controller.
* `enabled` - (Optional) Type: `bool`. Whether CAP mode is enabled on this device.
* `lock_to_caps_man` - (Optional) Type: `bool`. Lock the CAP to the first controller it successfully connects to.
* `mld_datapath` - (Optional) Type: `string`. RouterOS `mld-datapath`.
* `mld_static` - (Optional) Type: `string`. RouterOS `mld-static`.
* `slaves_datapath` - (Optional) Type: `string`. RouterOS `slaves-datapath`.
* `slaves_static` - (Optional) Type: `bool`. Create static rather than dynamic interfaces for provisioned radios.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Rebuilding a CAP

`enabled = true` on its own is not a working CAP. Unless the controller is
discoverable on the same L2 segment, set `caps_man_addresses` too -- otherwise the
radios come up unmanaged after a rebuild and the provisioned SSIDs never appear.

```terraform
resource "routeros_interface_wifi_cap" "this" {
  enabled            = true
  caps_man_addresses = "192.168.10.1"
  slaves_static      = false
}
```

* `mld_datapath` - (Optional) Type: `string`. RouterOS `mld-datapath`.
* `mld_static` - (Optional) Type: `string`. RouterOS `mld-static`.
* `slaves_datapath` - (Optional) Type: `string`. RouterOS `slaves-datapath`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_interface_wifi_cap.this 'home'
```
