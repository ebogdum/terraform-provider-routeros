---
subcategory: "Queues"
page_title: "RouterOS: routeros_queue_type"
description: |-
  RouterOS resource.
---

# Data Source: routeros_queue_type

Manages the RouterOS `/queue/type` menu.

## Example Usage

```terraform
data "routeros_queue_type" "type_example" {
  # router   = "my-router"  # omit for the default router
  # filter   = { name = "some-name" }
  # proplist = ["name", "address"]
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `proplist` - (Optional) List of property names to project; smaller payload.
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

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.

