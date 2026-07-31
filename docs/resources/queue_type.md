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
  name = "tf-example"

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
  # type_name = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `bfifo_limit` - (Optional) Type: `string`. RouterOS `bfifo-limit`.
* `cake_ack_filter` - (Optional) Type: `string`. RouterOS `cake-ack-filter`.
* `cake_atm` - (Optional) Type: `string`. RouterOS `cake-atm`.
* `cake_autorate_ingress` - (Optional) Type: `string`. RouterOS `cake-autorate-ingress`.
* `cake_bandwidth` - (Optional) Type: `string`. RouterOS `cake-bandwidth`.
* `cake_diffserv` - (Optional) Type: `string`. RouterOS `cake-diffserv`.
* `cake_flowmode` - (Optional) Type: `string`. RouterOS `cake-flowmode`.
* `cake_memlimit` - (Optional) Type: `string`. RouterOS `cake-memlimit`.
* `cake_mpu` - (Optional) Type: `string`. RouterOS `cake-mpu`.
* `cake_nat` - (Optional) Type: `string`. RouterOS `cake-nat`.
* `cake_overhead` - (Optional) Type: `string`. RouterOS `cake-overhead`.
* `cake_overhead_scheme` - (Optional) Type: `string`. RouterOS `cake-overhead-scheme`.
* `cake_rtt` - (Optional) Type: `string`. RouterOS `cake-rtt`.
* `cake_rtt_scheme` - (Optional) Type: `string`. RouterOS `cake-rtt-scheme`.
* `cake_wash` - (Optional) Type: `string`. RouterOS `cake-wash`.
* `codel_ce_threshold` - (Optional) Type: `string`. RouterOS `codel-ce-threshold`.
* `codel_ecn` - (Optional) Type: `string`. RouterOS `codel-ecn`.
* `codel_interval` - (Optional) Type: `string`. RouterOS `codel-interval`.
* `codel_limit` - (Optional) Type: `string`. RouterOS `codel-limit`.
* `codel_target` - (Optional) Type: `string`. RouterOS `codel-target`.
* `default` - (Read-only) Type: `bool`.
* `fq_codel_ce_threshold` - (Optional) Type: `string`. RouterOS `fq-codel-ce-threshold`.
* `fq_codel_ecn` - (Optional) Type: `string`. RouterOS `fq-codel-ecn`.
* `fq_codel_flows` - (Optional) Type: `string`. RouterOS `fq-codel-flows`.
* `fq_codel_interval` - (Optional) Type: `string`. RouterOS `fq-codel-interval`.
* `fq_codel_limit` - (Optional) Type: `string`. RouterOS `fq-codel-limit`.
* `fq_codel_memlimit` - (Optional) Type: `string`. RouterOS `fq-codel-memlimit`.
* `fq_codel_quantum` - (Optional) Type: `string`. RouterOS `fq-codel-quantum`.
* `fq_codel_target` - (Optional) Type: `string`. RouterOS `fq-codel-target`.
* `kind` - (Required) Type: `string`.
* `mq_pfifo_limit` - (Optional) Type: `int`.
* `name` - (Required) Type: `string`.
* `pcq_burst_rate` - (Optional) Type: `int`.
* `pcq_burst_threshold` - (Optional) Type: `int`.
* `pcq_burst_time` - (Optional) Type: `string`.
* `pcq_classifier` - (Optional) Type: `string`.
* `pcq_dst_address6_mask` - (Optional) Type: `int`.
* `pcq_dst_address_mask` - (Optional) Type: `int`.
* `pcq_limit` - (Optional) Type: `int`.
* `pcq_rate` - (Optional) Type: `int`.
* `pcq_src_address6_mask` - (Optional) Type: `int`.
* `pcq_src_address_mask` - (Optional) Type: `int`.
* `pcq_total_limit` - (Optional) Type: `int`.
* `pfifo_limit` - (Optional) Type: `int`.
* `red_avg_packet` - (Optional) Type: `int`.
* `red_burst` - (Optional) Type: `int`.
* `red_limit` - (Optional) Type: `int`.
* `red_max_threshold` - (Optional) Type: `int`.
* `red_min_threshold` - (Optional) Type: `int`.
* `sfq_allot` - (Optional) Type: `int`.
* `sfq_perturb` - (Optional) Type: `int`.
* `type_name` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
