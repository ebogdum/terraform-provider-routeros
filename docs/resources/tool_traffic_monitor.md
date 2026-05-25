---
page_title: "RouterOS: routeros_tool_traffic_monitor"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_traffic_monitor

Manages the RouterOS `/tool/traffic-monitor` menu.

## Example Usage

```terraform
resource "routeros_tool_traffic_monitor" "traffic_monitor_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # on_event = "replace-me"
  # threshold = "replace-me"
  # traffic = "transmitted"
  # trigger = "v1"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_trmon`.
* `on_event` - (Optional) Type: `string`.
* `threshold` - (Optional) Type: `string`.
* `traffic` - (Optional) Type: `enum(|transmitted|received)`. Default: `1`.
* `trigger` - (Optional) Type: `enum(|above|below|always)`. Default: `1`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_tool_traffic_monitor.example '*3'

# Named router
terraform import routeros_tool_traffic_monitor.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_tool_traffic_monitor.example 'home/my-resource-name'
```
