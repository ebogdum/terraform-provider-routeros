---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_datapath"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_datapath

Manages the RouterOS `/interface/wifi/datapath` menu.

## Example Usage

```terraform
resource "routeros_interface_wifi_datapath" "datapath_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # bridge = "bridge1"
  # bridge_cost = "replace-me"
  # bridge_horizon = "replace-me"
  # client_isolation = "replace-me"
  # interface_list = "replace-me"
  # open_flow_switch = "replace-me"
  # openflow = "replace-me"
  # traffic_processing = "replace-me"
  # vlan_id = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `bridge` - (Optional) Type: `string`.
* `bridge_cost` - (Optional) Type: `string`.
* `bridge_horizon` - (Optional) Type: `string`.
* `client_isolation` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface_list` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `open_flow_switch` - (Optional) Type: `string`.
* `openflow` - (Optional) Type: `string`.
* `traffic_processing` - (Optional) Type: `string`.
* `vlan_id` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wifi_datapath.example '*3'

# Named router
terraform import routeros_interface_wifi_datapath.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wifi_datapath.example 'home/my-resource-name'
```
