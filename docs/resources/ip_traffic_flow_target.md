---
page_title: "RouterOS: routeros_ip_traffic_flow_target"
description: |-
  Discovered; required dst-address must be valid
---

# Resource: routeros_ip_traffic_flow_target

Discovered; required dst-address must be valid

## Example Usage

```terraform
resource "routeros_ip_traffic_flow_target" "target_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # dst_address = "10.99.0.0/24"
  # port = "443"
  # src_address = "10.99.0.0/24"
  # version = "9"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `disabled` - (Optional) Type: `bool`.
* `dst_address` - (Optional) Type: `string`.
* `port` - (Optional) Type: `int`. Default: `1234`.
* `src_address` - (Optional) Type: `string`.
* `version` - (Optional) Type: `enum(1|5|9|IPFIX)`. Default: `9`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_traffic_flow_target.example '*3'

# Named router
terraform import routeros_ip_traffic_flow_target.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_traffic_flow_target.example 'home/my-resource-name'
```
