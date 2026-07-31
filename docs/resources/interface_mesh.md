---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_mesh"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_mesh

Manages the RouterOS `/interface/mesh` menu.

## Example Usage

```terraform
resource "routeros_interface_mesh" "mesh_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # admin_mac_address = "10.99.0.0/24"
  # arp = "enabled"
  # arp_timeout = "1h"
  # default_hoplimit = 32
  # mesh_portal = false
  # mesh_traceroute = "replace-me"
  # mtu = 1500
  # prep_lifetime = "300"
  # preq_destination_only = true
  # preq_reply_and_forward = true
  # preq_retries = 2
  # preq_waiting_time = 4
  # rann_interval = "10"
  # rann_lifetime = "22"
  # rann_propagation_delay = 500
  # reoptimize_paths = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `admin_mac` - (Optional) Type: `string`. RouterOS `admin-mac`.
* `admin_mac_address` - (Read-only) Type: `string`.
* `arp` - (Optional) Type: `string`.
* `arp_timeout` - (Optional) Type: `string`.
* `auto_mac` - (Optional) Type: `string`. RouterOS `auto-mac`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_hoplimit` - (Read-only) Type: `int`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `hwmp_default_hoplimit` - (Optional) Type: `string`. RouterOS `hwmp-default-hoplimit`.
* `hwmp_prep_lifetime` - (Optional) Type: `string`. RouterOS `hwmp-prep-lifetime`.
* `hwmp_preq_destination_only` - (Optional) Type: `string`. RouterOS `hwmp-preq-destination-only`.
* `hwmp_preq_reply_and_forward` - (Optional) Type: `string`. RouterOS `hwmp-preq-reply-and-forward`.
* `hwmp_preq_retries` - (Optional) Type: `string`. RouterOS `hwmp-preq-retries`.
* `hwmp_preq_waiting_time` - (Optional) Type: `string`. RouterOS `hwmp-preq-waiting-time`.
* `hwmp_rann_interval` - (Optional) Type: `string`. RouterOS `hwmp-rann-interval`.
* `hwmp_rann_lifetime` - (Optional) Type: `string`. RouterOS `hwmp-rann-lifetime`.
* `hwmp_rann_propagation_delay` - (Optional) Type: `string`. RouterOS `hwmp-rann-propagation-delay`.
* `mac_address` - (Read-only) Type: `string`.
* `mesh_portal` - (Optional) Type: `bool`.
* `mesh_traceroute` - (Read-only) Type: `string`.
* `mtu` - (Optional) Type: `int`.
* `name` - (Optional) Type: `string`. RouterOS `name`.
* `prep_lifetime` - (Read-only) Type: `string`.
* `preq_destination_only` - (Read-only) Type: `bool`.
* `preq_reply_and_forward` - (Read-only) Type: `bool`.
* `preq_retries` - (Read-only) Type: `int`.
* `preq_waiting_time` - (Read-only) Type: `int`.
* `rann_interval` - (Read-only) Type: `string`.
* `rann_lifetime` - (Read-only) Type: `string`.
* `rann_propagation_delay` - (Read-only) Type: `int`.
* `reoptimize_paths` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_mesh.example '*3'

# Named router
terraform import routeros_interface_mesh.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_mesh.example 'home/my-resource-name'
```
