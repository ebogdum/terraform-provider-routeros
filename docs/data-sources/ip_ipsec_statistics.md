---
page_title: "RouterOS: routeros_ip_ipsec_statistics"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_ipsec_statistics

Manages the RouterOS `/ip/ipsec/statistics` menu.

## Example Usage

```terraform
data "routeros_ip_ipsec_statistics" "statistics_example" {
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
* `in_buffer_errors` - (Optional) Type: `int`.
* `in_errors` - (Optional) Type: `int`.
* `in_header_errors` - (Optional) Type: `int`.
* `in_no_policies` - (Optional) Type: `int`.
* `in_no_states` - (Optional) Type: `int`.
* `in_policy_blocked` - (Optional) Type: `int`.
* `in_policy_errors` - (Optional) Type: `int`.
* `in_state_expired` - (Optional) Type: `int`.
* `in_state_invalid` - (Optional) Type: `int`.
* `in_state_mismatches` - (Optional) Type: `int`.
* `in_state_mode_errors` - (Optional) Type: `int`.
* `in_state_protocol_errors` - (Optional) Type: `int`.
* `in_state_sequence_errors` - (Optional) Type: `int`.
* `in_template_mismatches` - (Optional) Type: `int`.
* `out_bundle_check_errors` - (Optional) Type: `int`.
* `out_bundle_errors` - (Optional) Type: `int`.
* `out_errors` - (Optional) Type: `int`.
* `out_no_states` - (Optional) Type: `int`.
* `out_policy_blocked` - (Optional) Type: `int`.
* `out_policy_dead` - (Optional) Type: `int`.
* `out_policy_errors` - (Optional) Type: `int`.
* `out_state_expired` - (Optional) Type: `int`.
* `out_state_mode_errors` - (Optional) Type: `int`.
* `out_state_protocol_errors` - (Optional) Type: `int`.
* `out_state_sequence_errors` - (Optional) Type: `int`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

