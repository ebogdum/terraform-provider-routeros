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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `latency_distribution_max` - (Optional) Type: `duration`.
* `measure_out_of_order` - (Optional) Type: `bool`.
* `stats_samples_to_keep` - (Optional) Type: `int`.
* `test_id` - (Optional) Type: `int`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `latency_distribution_measure_interval` - Type: `string`.
* `latency_distribution_samples` - Type: `int`.
* `running` - Type: `bool`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_tool_traffic_generator.this 'home'
```
