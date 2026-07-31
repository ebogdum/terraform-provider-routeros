---
subcategory: "IP"
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
  # v9 = "replace-me"
  # v9_ipfix_template_refresh = 20
  # v9_ipfix_template_timeout = 1800
  # version = "9"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `disabled` - (Optional) Type: `bool`.
* `dst_address` - (Optional) Type: `string`.
* `port` - (Optional) Type: `int`.
* `src_address` - (Optional) Type: `string`.
* `v9` - (Read-only) Type: `string`.
* `v9_ipfix_template_refresh` - (Read-only) Type: `int`.
* `v9_ipfix_template_timeout` - (Read-only) Type: `int`.
* `v9_template_refresh` - (Optional) Type: `string`. RouterOS `v9-template-refresh`.
* `v9_template_timeout` - (Optional) Type: `string`. RouterOS `v9-template-timeout`.
* `version` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
