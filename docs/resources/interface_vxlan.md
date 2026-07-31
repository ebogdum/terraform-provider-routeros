---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_vxlan"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_vxlan

Manages the RouterOS `/interface/vxlan` menu.

## Example Usage

```terraform
resource "routeros_interface_vxlan" "vxlan_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"
  vni = "100"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # bridge = "bridge1"
  # interface = "ether1"
  # local_address = "10.99.0.1"
  # mac_address = "10.99.0.0/24"
  # mtu = "replace-me"
  # port = "443"
  # ttl = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `allow_fast_path` - (Optional) Type: `string`. RouterOS `allow-fast-path`.
* `arp` - (Optional) Type: `string`.
* `arp_timeout` - (Optional) Type: `string`.
* `bridge` - (Optional) Type: `string`.
* `bridge_pvid` - (Optional) Type: `string`. RouterOS `bridge-pvid`.
* `checksum` - (Optional) Type: `string`. RouterOS `checksum`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dont_fragment` - (Optional) Type: `string`. RouterOS `dont-fragment`.
* `group` - (Optional) Type: `string`. RouterOS `group`.
* `hw` - (Optional) Type: `string`. RouterOS `hw`.
* `interface` - (Optional) Type: `string`.
* `learning` - (Optional) Type: `string`. RouterOS `learning`.
* `local_address` - (Optional) Type: `string`.
* `loop_protect` - (Optional) Type: `string`. RouterOS `loop-protect`.
* `loop_protect_disable_time` - (Optional) Type: `string`. RouterOS `loop-protect-disable-time`.
* `loop_protect_send_interval` - (Optional) Type: `string`. RouterOS `loop-protect-send-interval`.
* `mac_address` - (Optional) Type: `string`.
* `max_fdb_size` - (Optional) Type: `string`. RouterOS `max-fdb-size`.
* `mtu` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `port` - (Optional) Type: `string`.
* `rem_csum` - (Optional) Type: `string`. RouterOS `rem-csum`.
* `ttl` - (Optional) Type: `string`.
* `vni` - (Required) Type: `string`.
* `vtep_vrf` - (Optional) Type: `string`. RouterOS `vtep-vrf`.
* `vteps_ip_version` - (Optional) Type: `string`. RouterOS `vteps-ip-version`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_vxlan.example '*3'

# Named router
terraform import routeros_interface_vxlan.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_vxlan.example 'home/my-resource-name'
```
