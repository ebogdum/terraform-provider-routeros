---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_traffic_generator"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_traffic_generator

Manages the RouterOS `/tool/traffic-generator` menu.

## Example Usage

```terraform
resource "routeros_tool_traffic_generator" "traffic_generator_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # latency_distribution_max = "1h"
  # measure_out_of_order = false
  # stats_samples_to_keep = 0
  # test_id = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `latency_distribution_max` - (Optional) Type: `string`.
* `latency_distribution_measure_interval` - (Optional) Type: `string`.
* `latency_distribution_samples` - (Optional) Type: `int`.
* `measure_out_of_order` - (Optional) Type: `bool`.
* `running` - (Optional) Type: `bool`.
* `stats_samples_to_keep` - (Optional) Type: `int`.
* `test_id` - (Optional) Type: `int`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_tool_traffic_generator.this 'home'
```
