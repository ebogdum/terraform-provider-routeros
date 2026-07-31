---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_steering"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_steering

Manages the RouterOS `/interface/wifi/steering` menu.

## Example Usage

```terraform
resource "routeros_interface_wifi_steering" "steering_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # x2g_probe_delay = "replace-me"
  # neighbor_group = "replace-me"
  # neighbor_groups = "replace-me"
  # rrm = "replace-me"
  # transition_request_count = "replace-me"
  # transition_threshold = "replace-me"
  # transition_threshold_period = "replace-me"
  # transition_threshold_time = "replace-me"
  # transition_time = "replace-me"
  # wnm = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `name` - (Required) Type: `string`.
* `neighbor_group` - (Optional) Type: `string`.
* `neighbor_groups` - (Optional) Type: `string`.
* `rrm` - (Optional) Type: `string`.
* `transition_request_count` - (Optional) Type: `string`.
* `transition_request_period` - (Optional) Type: `string`. RouterOS `transition-request-period`.
* `transition_threshold` - (Optional) Type: `string`.
* `transition_threshold_period` - (Optional) Type: `string`.
* `transition_threshold_time` - (Optional) Type: `string`.
* `transition_time` - (Optional) Type: `string`.
* `wnm` - (Optional) Type: `string`.
* `x2g_probe_delay` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wifi_steering.example '*3'

# Named router
terraform import routeros_interface_wifi_steering.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wifi_steering.example 'home/my-resource-name'
```
