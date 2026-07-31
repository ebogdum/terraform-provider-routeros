---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_vrrp"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_vrrp

Manages the RouterOS `/interface/vrrp` menu.

## Example Usage

```terraform
resource "routeros_interface_vrrp" "vrrp_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # arp = "replace-me"
  # arp_timeout = "replace-me"
  # authentication = "replace-me"
  # interval = "replace-me"
  # password = "REDACTED"
  # priority = "replace-me"
  # remote_address = "10.99.0.1"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `arp` - (Optional) Type: `string`.
* `arp_timeout` - (Optional) Type: `string`.
* `authentication` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connection_tracking_mode` - (Optional) Type: `string`. RouterOS `connection-tracking-mode`.
* `connection_tracking_port` - (Optional) Type: `string`. RouterOS `connection-tracking-port`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `group_authority` - (Optional) Type: `string`. RouterOS `group-authority`.
* `group_master` - (Optional) Type: `string`. RouterOS `group-master`.
* `interface` - (Required) Type: `string`.
* `interval` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `on_backup` - (Optional) Type: `string`. RouterOS `on-backup`.
* `on_fail` - (Optional) Type: `string`. RouterOS `on-fail`.
* `on_master` - (Optional) Type: `string`. RouterOS `on-master`.
* `password` - (Optional) Type: `string`. **Sensitive.**
* `preemption_mode` - (Optional) Type: `string`. RouterOS `preemption-mode`.
* `priority` - (Optional) Type: `string`.
* `remote_address` - (Optional) Type: `string`.
* `sync_connection_tracking` - (Optional) Type: `string`. RouterOS `sync-connection-tracking`.
* `v3_protocol` - (Optional) Type: `string`. RouterOS `v3-protocol`.
* `version` - (Optional) Type: `string`. RouterOS `version`.
* `vrid` - (Optional) Type: `string`. RouterOS `vrid`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_vrrp.example '*3'

# Named router
terraform import routeros_interface_vrrp.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_vrrp.example 'home/my-resource-name'
```
