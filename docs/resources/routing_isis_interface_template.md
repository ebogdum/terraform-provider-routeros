---
subcategory: "ISIS"
page_title: "RouterOS: routeros_routing_isis_interface_template"
description: |-
  References an existing isis instance; auto-test can't synthesise.
---

# Resource: routeros_routing_isis_interface_template

References an existing isis instance; auto-test can't synthesise.

## Example Usage

```terraform
resource "routeros_routing_isis_interface_template" "interface_template_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `bcast_l1_csnp_interval` - (Optional) Type: `string`. RouterOS `bcast.l1.csnp-interval`.
* `bcast_l1_hello_interval` - (Optional) Type: `string`. RouterOS `bcast.l1.hello-interval`.
* `bcast_l1_hello_interval_dr` - (Optional) Type: `string`. RouterOS `bcast.l1.hello-interval-dr`.
* `bcast_l1_hello_multiplier` - (Optional) Type: `string`. RouterOS `bcast.l1.hello-multiplier`.
* `bcast_l1_metric` - (Optional) Type: `string`. RouterOS `bcast.l1.metric`.
* `bcast_l1_priority` - (Optional) Type: `string`. RouterOS `bcast.l1.priority`.
* `bcast_l1_psnp_interval` - (Optional) Type: `string`. RouterOS `bcast.l1.psnp-interval`.
* `bcast_l2_csnp_interval` - (Optional) Type: `string`. RouterOS `bcast.l2.csnp-interval`.
* `bcast_l2_hello_interval` - (Optional) Type: `string`. RouterOS `bcast.l2.hello-interval`.
* `bcast_l2_hello_interval_dr` - (Optional) Type: `string`. RouterOS `bcast.l2.hello-interval-dr`.
* `bcast_l2_hello_multiplier` - (Optional) Type: `string`. RouterOS `bcast.l2.hello-multiplier`.
* `bcast_l2_metric` - (Optional) Type: `string`. RouterOS `bcast.l2.metric`.
* `bcast_l2_priority` - (Optional) Type: `string`. RouterOS `bcast.l2.priority`.
* `bcast_l2_psnp_interval` - (Optional) Type: `string`. RouterOS `bcast.l2.psnp-interval`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `instance` - (Optional) Type: `string`. RouterOS `instance`.
* `interfaces` - (Optional) Type: `string`. RouterOS `interfaces`.
* `levels` - (Optional) Type: `string`. RouterOS `levels`.
* `passive` - (Optional) Type: `string`. RouterOS `passive`.
* `ptp` - (Optional) Type: `string`. RouterOS `ptp`.
* `ptp_hello_3way` - (Optional) Type: `string`. RouterOS `ptp.hello-3way`.
* `ptp_hello_interval` - (Optional) Type: `string`. RouterOS `ptp.hello-interval`.
* `ptp_hello_multiplier` - (Optional) Type: `string`. RouterOS `ptp.hello-multiplier`.
* `ptp_l1_csnp_interval` - (Optional) Type: `string`. RouterOS `ptp.l1.csnp-interval`.
* `ptp_l1_metric` - (Optional) Type: `string`. RouterOS `ptp.l1.metric`.
* `ptp_l1_psnp_interval` - (Optional) Type: `string`. RouterOS `ptp.l1.psnp-interval`.
* `ptp_l2_csnp_interval` - (Optional) Type: `string`. RouterOS `ptp.l2.csnp-interval`.
* `ptp_l2_metric` - (Optional) Type: `string`. RouterOS `ptp.l2.metric`.
* `ptp_l2_psnp_interval` - (Optional) Type: `string`. RouterOS `ptp.l2.psnp-interval`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_isis_interface_template.example '*3'

# Named router
terraform import routeros_routing_isis_interface_template.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_isis_interface_template.example 'home/my-resource-name'
```
