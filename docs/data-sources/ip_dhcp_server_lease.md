---
subcategory: "DHCP"
page_title: "RouterOS: routeros_ip_dhcp_server_lease"
description: |-
  Lease entries reference an existing dhcp-server; auto-test can't synthesise the precondition reliably.
---

# Data Source: routeros_ip_dhcp_server_lease

Lease entries reference an existing dhcp-server; auto-test can't synthesise the precondition reliably.

## Example Usage

```terraform
data "routeros_ip_dhcp_server_lease" "lease_example" {
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
* `address` - (Optional) Type: `string`.
* `address_list` - (Optional) Type: `string`.
* `agent_circuit_id` - (Optional) Type: `string`.
* `agent_remote_id` - (Optional) Type: `string`.
* `allow_dual_stack_queue` - (Optional) Type: `string`.
* `always_broadcast` - (Optional) Type: `bool`.
* `block_access` - (Optional) Type: `bool`.
* `blocked` - (Optional) Type: `bool`.
* `check_status` - (Optional) Type: `string`.
* `client_id` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dhcp_option_set` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `dhcp_options` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dyn` - (Optional) Type: `string`.
* `insert_queue_before` - (Optional) Type: `string`.
* `lease_time` - (Optional) Type: `duration`.
* `mac_address` - (Optional) Type: `string`.
* `make_static` - (Optional) Type: `string`.
* `parent_queue` - (Optional) Type: `string`.
* `ping` - (Optional) Type: `string`.
* `queue_type` - (Optional) Type: `string`.
* `radius` - (Optional) Type: `bool`.
* `rate_limit` - (Optional) Type: `string`.
* `rostat` - (Optional) Type: `string`.
* `routes` - (Optional) Type: `string`.
* `send_reconfigure` - (Optional) Type: `string`.
* `server` - (Optional) Type: `string`.
* `stat` - (Optional) Type: `string`.
* `use_src_mac_address` - (Optional) Type: `bool`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `active_address` - Type: `ip`.
* `active_agent_circuit_id` - Type: `string`.
* `active_agent_remote_id` - Type: `string`.
* `active_class_id` - Type: `string`.
* `active_client_id` - Type: `string`.
* `active_host_name` - Type: `string`.
* `active_mac_address` - Type: `string`.
* `active_server` - Type: `string`.
* `age` - Type: `string`.
* `bridge_port` - Type: `string`.
* `dynamic` - Type: `bool`.
* `expires_after` - Type: `string`.
* `last_seen` - Type: `string`.
* `last_sent_counter` - Type: `string`.
* `reconfigure_key` - Type: `string`.
* `reconfigure_status` - Type: `string`.
* `src_mac_address` - Type: `string`.

