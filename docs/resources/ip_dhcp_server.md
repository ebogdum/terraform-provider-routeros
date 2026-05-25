---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_dhcp_server"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_dhcp_server

Manages the RouterOS `/ip/dhcp-server` menu.

## Example Usage

```terraform
resource "routeros_ip_dhcp_server" "dhcp_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address_pool = "replace-me"
  # allow_dual_stack_queue = true
  # always_broadcast = false
  # authoritative = "yes"
  # bootp_lease_time = "4.294967295e+09"
  # bootp_support = "static"
  # client_mac_limit = 0
  # conflict_detection = true
  # delay_threshold = "1h"
  # dhcp_option_set = "4.294967295e+09"
  # dynamic_lease_identifiers = "MAC Address"
  # insert_queue_before = "0"
  # lease_script = "replace-me"
  # lease_time = "1800"
  # parent_queue = "replace-me"
  # relay = "10.99.0.1"
  # server_address = "10.99.0.0/24"
  # use_framed_as_classless = true
  # use_radius = "no"
  # use_reconfigure = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address_pool` - (Optional) Type: `string`.
* `allow_dual_stack_queue` - (Optional) Type: `bool`. Default: `1`.
* `always_broadcast` - (Optional) Type: `bool`.
* `authoritative` - (Optional) Type: `enum(yes|after 2s delay|after 10s delay|no)`. Default: `0`.
* `bootp_lease_time` - (Optional) Type: `duration`. Default: `4.294967295e+09`.
* `bootp_support` - (Optional) Type: `enum(none|static|dynamic)`. Default: `1`.
* `client_mac_limit` - (Optional) Type: `int`. Default: `4.294967295e+09`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `conflict_detection` - (Optional) Type: `bool`. Default: `1`.
* `delay_threshold` - (Optional) Type: `duration`.
* `dhcp_option_set` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic_lease_identifiers` - (Optional) Type: `string`.
* `insert_queue_before` - (Optional) Type: `string`. Default: `0`.
* `interface` - (Required) Type: `string`.
* `lease_script` - (Optional) Type: `string`.
* `lease_time` - (Optional) Type: `duration`. Default: `1800`.
* `name` - (Required) Type: `string`. Default: `tf_acc_dhcps`.
* `parent_queue` - (Optional) Type: `string`.
* `relay` - (Optional) Type: `ip`.
* `server_address` - (Optional) Type: `ip`.
* `use_framed_as_classless` - (Optional) Type: `bool`. Default: `1`.
* `use_radius` - (Optional) Type: `enum(no|yes|accounting)`.
* `use_reconfigure` - (Optional) Type: `bool`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_dhcp_server.example '*3'

# Named router
terraform import routeros_ip_dhcp_server.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_dhcp_server.example 'home/my-resource-name'
```
