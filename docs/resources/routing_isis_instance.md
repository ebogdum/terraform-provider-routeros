---
subcategory: "ISIS"
page_title: "RouterOS: routeros_routing_isis_instance"
description: |-
  ISIS instance argument set differs across ROS releases. Skipped.
---

# Resource: routeros_routing_isis_instance

ISIS instance argument set differs across ROS releases. Skipped.

## Example Usage

```terraform
resource "routeros_routing_isis_instance" "instance_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `afi` - (Optional) Type: `string`. RouterOS `afi`.
* `areas` - (Optional) Type: `string`. RouterOS `areas`.
* `areas_max` - (Optional) Type: `string`. RouterOS `areas-max`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `in_filter_chain` - (Optional) Type: `string`. RouterOS `in-filter-chain`.
* `l1_lsp_max_age` - (Optional) Type: `string`. RouterOS `l1.lsp-max-age`.
* `l1_lsp_max_size` - (Optional) Type: `string`. RouterOS `l1.lsp-max-size`.
* `l1_lsp_refresh_interval` - (Optional) Type: `string`. RouterOS `l1.lsp-refresh-interval`.
* `l1_lsp_update_interval` - (Optional) Type: `string`. RouterOS `l1.lsp-update-interval`.
* `l1_originate_default` - (Optional) Type: `string`. RouterOS `l1.originate-default`.
* `l1_out_filter_chain` - (Optional) Type: `string`. RouterOS `l1.out-filter-chain`.
* `l1_out_filter_select` - (Optional) Type: `string`. RouterOS `l1.out-filter-select`.
* `l1_redistribute` - (Optional) Type: `string`. RouterOS `l1.redistribute`.
* `l2_lsp_max_age` - (Optional) Type: `string`. RouterOS `l2.lsp-max-age`.
* `l2_lsp_max_size` - (Optional) Type: `string`. RouterOS `l2.lsp-max-size`.
* `l2_lsp_update_interval` - (Optional) Type: `string`. RouterOS `l2.lsp-update-interval`.
* `l2_originate_default` - (Optional) Type: `string`. RouterOS `l2.originate-default`.
* `l2_out_filter_chain` - (Optional) Type: `string`. RouterOS `l2.out-filter-chain`.
* `l2_out_filter_select` - (Optional) Type: `string`. RouterOS `l2.out-filter-select`.
* `l2_redistribute` - (Optional) Type: `string`. RouterOS `l2.redistribute`.
* `metric_type` - (Optional) Type: `string`. RouterOS `metric-type`.
* `name` - (Optional) Type: `string`. RouterOS `name`.
* `system_id` - (Optional) Type: `string`. RouterOS `system-id`.
* `vrf` - (Optional) Type: `string`. RouterOS `vrf`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_isis_instance.example '*3'

# Named router
terraform import routeros_routing_isis_instance.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_isis_instance.example 'home/my-resource-name'
```
