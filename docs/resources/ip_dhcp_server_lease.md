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
  # agent_circuit_id = "replace-me"
  # agent_remote_id = "replace-me"
  # allow_dual_stack_queue = "replace-me"
  # always_broadcast = false
  # block_access = false
  # client_id = "replace-me"
  # dhcp_option_set = "4.294967295e+09"
  # insert_queue_before = "replace-me"
  # lease_time = "1h"
  # mac_address = "10.99.0.0/24"
  # parent_queue = "replace-me"
  # queue_type = "replace-me"
  # rate_limit = "replace-me"
  # routes = "replace-me"
  # server = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Optional) Type: `string`.
* `agent_circuit_id` - (Optional) Type: `string`.
* `agent_remote_id` - (Optional) Type: `string`.
* `allow_dual_stack_queue` - (Optional) Type: `string`.
* `always_broadcast` - (Optional) Type: `bool`.
* `block_access` - (Optional) Type: `bool`.
* `client_id` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dhcp_option_set` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `insert_queue_before` - (Optional) Type: `string`.
* `lease_time` - (Optional) Type: `duration`.
* `mac_address` - (Optional) Type: `string`.
* `parent_queue` - (Optional) Type: `string`.
* `queue_type` - (Optional) Type: `string`.
* `rate_limit` - (Optional) Type: `string`.
* `routes` - (Optional) Type: `string`.
* `server` - (Optional) Type: `string`.

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
