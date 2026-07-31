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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `active_address` - (Read-only) Type: `string`.
* `active_agent_circuit_id` - (Read-only) Type: `string`.
* `active_agent_remote_id` - (Read-only) Type: `string`.
* `active_class_id` - (Read-only) Type: `string`.
* `active_client_id` - (Read-only) Type: `string`.
* `active_host_name` - (Read-only) Type: `string`.
* `active_mac_address` - (Read-only) Type: `string`.
* `active_server` - (Read-only) Type: `string`.
* `address` - (Optional) Type: `string`.
* `address_list` - (Read-only) Type: `string`.
* `address_lists` - (Optional) Type: `string`. RouterOS `address-lists`.
* `age` - (Read-only) Type: `string`.
* `agent_circuit_id` - (Optional) Type: `string`.
* `agent_remote_id` - (Optional) Type: `string`.
* `allow_dual_stack_queue` - (Optional) Type: `string`.
* `always_broadcast` - (Optional) Type: `bool`.
* `block_access` - (Optional) Type: `bool`.
* `blocked` - (Read-only) Type: `bool`.
* `bridge_port` - (Read-only) Type: `string`.
* `check_status` - (Read-only) Type: `string`.
* `client_id` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dhcp_option` - (Optional) Type: `string`. RouterOS `dhcp-option`.
* `dhcp_option_set` - (Optional) Type: `string`.
* `dhcp_options` - (Read-only) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dyn` - (Read-only) Type: `string`.
* `dynamic` - (Read-only) Type: `bool`.
* `expires_after` - (Read-only) Type: `string`.
* `insert_queue_before` - (Optional) Type: `string`.
* `last_seen` - (Read-only) Type: `string`.
* `last_sent_counter` - (Read-only) Type: `string`.
* `lease_time` - (Optional) Type: `string`.
* `mac_address` - (Optional) Type: `string`.
* `make_static` - (Read-only) Type: `string`.
* `parent_queue` - (Optional) Type: `string`.
* `ping` - (Read-only) Type: `string`.
* `queue_type` - (Optional) Type: `string`.
* `radius` - (Read-only) Type: `bool`.
* `rate_limit` - (Optional) Type: `string`.
* `reconfigure_key` - (Read-only) Type: `string`.
* `reconfigure_status` - (Read-only) Type: `string`.
* `rostat` - (Read-only) Type: `string`.
* `routes` - (Optional) Type: `string`.
* `send_reconfigure` - (Read-only) Type: `string`.
* `server` - (Optional) Type: `string`.
* `src_mac_address` - (Read-only) Type: `string`.
* `stat` - (Read-only) Type: `string`.
* `use_src_mac` - (Optional) Type: `string`. RouterOS `use-src-mac`.
* `use_src_mac_address` - (Read-only) Type: `bool`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
