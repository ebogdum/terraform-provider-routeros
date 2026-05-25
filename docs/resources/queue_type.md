---
subcategory: "Queues"
page_title: "RouterOS: routeros_queue_type"
description: |-
  RouterOS resource.
---

# Resource: routeros_queue_type

Manages the RouterOS `/queue/type` menu.

## Example Usage

```terraform
resource "routeros_queue_type" "type_example" {
  # router = "my-router"  # which router to target; omit for the default
  kind = "pfifo"
  name = "example"

  # Optional attributes (uncomment as needed):
  # mq_pfifo_limit = 0
  # pcq_burst_rate = 0
  # pcq_burst_threshold = 0
  # pcq_burst_time = "1h"
  # pcq_classifier = "replace-me"
  # pcq_dst_address_mask = 0
  # pcq_dst_address6_mask = 0
  # pcq_limit = 0
  # pcq_rate = 0
  # pcq_src_address_mask = 0
  # pcq_src_address6_mask = 0
  # pcq_total_limit = 0
  # pfifo_limit = 0
  # red_avg_packet = 0
  # red_burst = 0
  # red_limit = 0
  # red_max_threshold = 0
  # red_min_threshold = 0
  # sfq_allot = 0
  # sfq_perturb = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `kind` - (Required) Type: `enum(|bfifo|pfifo|red|sfq|pcq, ...)`. Default: `pfifo`.
* `mq_pfifo_limit` - (Optional) Type: `int`.
* `name` - (Required) Type: `string`. Default: `tf-acc-qtype`.
* `pcq_burst_rate` - (Optional) Type: `int`.
* `pcq_burst_threshold` - (Optional) Type: `int`.
* `pcq_burst_time` - (Optional) Type: `duration`.
* `pcq_classifier` - (Optional) Type: `string`.
* `pcq_dst_address_mask` - (Optional) Type: `int`.
* `pcq_dst_address6_mask` - (Optional) Type: `int`.
* `pcq_limit` - (Optional) Type: `int`.
* `pcq_rate` - (Optional) Type: `int`.
* `pcq_src_address_mask` - (Optional) Type: `int`.
* `pcq_src_address6_mask` - (Optional) Type: `int`.
* `pcq_total_limit` - (Optional) Type: `int`.
* `pfifo_limit` - (Optional) Type: `int`.
* `red_avg_packet` - (Optional) Type: `int`.
* `red_burst` - (Optional) Type: `int`.
* `red_limit` - (Optional) Type: `int`.
* `red_max_threshold` - (Optional) Type: `int`.
* `red_min_threshold` - (Optional) Type: `int`.
* `sfq_allot` - (Optional) Type: `int`.
* `sfq_perturb` - (Optional) Type: `int`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `default` - Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_queue_type.example '*3'

# Named router
terraform import routeros_queue_type.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_queue_type.example 'home/my-resource-name'
```
