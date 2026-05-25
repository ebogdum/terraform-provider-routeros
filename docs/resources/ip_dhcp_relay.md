---
page_title: "RouterOS: routeros_ip_dhcp_relay"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_dhcp_relay

Manages the RouterOS `/ip/dhcp-relay` menu.

## Example Usage

```terraform
resource "routeros_ip_dhcp_relay" "dhcp_relay_example" {
  # router = "my-router"  # which router to target; omit for the default
  dhcp_server = "127.0.0.1"
  interface = "ether1"
  name = "tf-example"

  disabled = false

  # Optional attributes (uncomment as needed):
  # add_relay_info = false
  # delay_threshold = "1h"
  # dhcp_server_vrf = "replace-me"
  # local_address = "10.99.0.1"
  # relay_info_remote_id = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `add_relay_info` - (Optional) Type: `bool`.
* `delay_threshold` - (Optional) Type: `duration`.
* `dhcp_server` - (Required) Type: `string`. Default: `127.0.0.1`.
* `dhcp_server_vrf` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.
* `local_address` - (Optional) Type: `ip`.
* `name` - (Required) Type: `string`. Default: `tf-acc-relay`.
* `relay_info_remote_id` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_dhcp_relay.example '*3'

# Named router
terraform import routeros_ip_dhcp_relay.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_dhcp_relay.example 'home/my-resource-name'
```
