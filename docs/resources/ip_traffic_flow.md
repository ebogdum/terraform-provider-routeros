---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_traffic_flow"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_traffic_flow

Manages the RouterOS `/ip/traffic-flow` menu.

## Example Usage

```terraform
resource "routeros_ip_traffic_flow" "traffic_flow_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # active_flow_timeout = "1h"
  # cache_entries = "replace-me"
  # enabled = false
  # inactive_flow_timeout = "1h"
  # interfaces = "replace-me"
  # packet_sampling = false
  # sampling_interval = 0
  # sampling_space = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `active_flow_timeout` - (Optional) Type: `duration`. Maximum life-time of a flow.
* `cache_entries` - (Optional) Type: `string`. Number of flows which can be in router's memory simultaneously.
* `enabled` - (Optional) Type: `bool`.
* `inactive_flow_timeout` - (Optional) Type: `duration`. How long to keep the flow active, if it is idle. If a connection does not see any packet within this timeout, then traffic-flow will send a packet out as a new flow. If this timeout is too small it can create a significant amount of flows and overflow the buffer.
* `interfaces` - (Optional) Type: `string`. Names of those interfaces will be used to gather statistics for traffic-flow. To specify more than one interface, separate them with a comma.
* `packet_sampling` - (Optional) Type: `bool`. Enable or disable packet sampling feature.
* `sampling_interval` - (Optional) Type: `int`. The number of packets that are consecutively sampled.
* `sampling_space` - (Optional) Type: `int`. The number of packets that are consecutively omitted.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_traffic_flow.this 'home'
```
