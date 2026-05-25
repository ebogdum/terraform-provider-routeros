---
page_title: "RouterOS: routeros_interface"
description: |-
  /interface is mostly read-only -- interfaces are created by their specific subtypes (/interface/bridge, /interface/wireguard, etc.).
---

# Resource: routeros_interface

/interface is mostly read-only -- interfaces are created by their specific subtypes (/interface/bridge, /interface/wireguard, etc.).

## Example Usage

```terraform
resource "routeros_interface" "interface_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # mtu = 1500
  # name = "tf-example"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `mtu` - (Optional) Type: `int`. Default: `1500`.
* `name` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `actual_mtu` - Type: `int`.
* `default_name` - Type: `string`.
* `fp_rps_drop` - Type: `int`.
* `fp_rx_byte` - Type: `int`.
* `fp_rx_packet` - Type: `int`.
* `fp_tx_byte` - Type: `int`.
* `fp_tx_packet` - Type: `int`.
* `last_link_up_time` - Type: `string`.
* `link_downs` - Type: `int`.
* `mac_address` - Type: `mac`.
* `running` - Type: `bool`.
* `rx_byte` - Type: `int`.
* `rx_drop` - Type: `int`.
* `rx_error` - Type: `int`.
* `rx_packet` - Type: `int`.
* `tx_byte` - Type: `int`.
* `tx_drop` - Type: `int`.
* `tx_error` - Type: `int`.
* `tx_packet` - Type: `int`.
* `tx_queue_drop` - Type: `int`.
* `type` - Type: `int`.
* `vrf` - Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface.example '*3'

# Named router
terraform import routeros_interface.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface.example 'home/my-resource-name'
```
