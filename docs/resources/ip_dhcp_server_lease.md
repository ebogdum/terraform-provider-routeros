---
subcategory: "DHCP"
page_title: "RouterOS: routeros_ip_dhcp_server_lease"
description: |-
  Lease entries reference an existing dhcp-server; auto-test can't synthesise the precondition reliably.
---

# Resource: routeros_ip_dhcp_server_lease

Lease entries reference an existing dhcp-server; auto-test can't synthesise the precondition reliably.

## Example Usage

```terraform
resource "routeros_ip_dhcp_server_lease" "lease_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # address_list = "replace-me"
  # agent_circuit_id = "replace-me"
  # agent_remote_id = "replace-me"
  # allow_dual_stack_queue = "replace-me"
  # always_broadcast = false
  # block_access = false
  # blocked = false
  # check_status = "replace-me"
  # client_id = "replace-me"
  # dhcp_option_set = "4.294967295e+09"
  # dhcp_options = "replace-me"
  # dyn = "replace-me"
  # insert_queue_before = "replace-me"
  # lease_time = "1h"
  # mac_address = "10.99.0.0/24"
  # make_static = "replace-me"
  # parent_queue = "replace-me"
  # ping = "replace-me"
  # queue_type = "replace-me"
  # radius = false
  # rate_limit = "replace-me"
  # rostat = "replace-me"
  # routes = "replace-me"
  # send_reconfigure = "replace-me"
  # server = "replace-me"
  # stat = "replace-me"
  # use_src_mac_address = "10.99.0.0/24"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
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

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
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

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_dhcp_server_lease.example '*3'

# Named router
terraform import routeros_ip_dhcp_server_lease.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_dhcp_server_lease.example 'home/my-resource-name'
```
