---
subcategory: "DHCP"
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
  # add_arp_for_leases = false
  # address_list = "replace-me"
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
  # dynamic_lease_identifiers = "mac-address"
  # dynbootp = "replace-me"
  # insert_queue_before = "0"
  # lease_script = "replace-me"
  # lease_time = "1800"
  # parent_queue = "replace-me"
  # relay = "10.99.0.1"
  # server_address = "10.99.0.0/24"
  # support_the_broadband_forum_tr_101 = false
  # use_framed_as_classless = true
  # use_radius = "no"
  # use_reconfigure = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `add_arp` - (Optional) Type: `string`. RouterOS `add-arp`.
* `add_arp_for_leases` - (Read-only) Type: `bool`.
* `add_dns_entries` - (Optional) Type: `string`. RouterOS `add-dns-entries`.
* `add_dns_entries_suffix` - (Optional) Type: `string`. RouterOS `add-dns-entries-suffix`.
* `address_list` - (Read-only) Type: `string`.
* `address_lists` - (Optional) Type: `string`. RouterOS `address-lists`.
* `address_pool` - (Optional) Type: `string`.
* `allow_dual_stack_queue` - (Optional) Type: `bool`.
* `always_broadcast` - (Optional) Type: `bool`.
* `authoritative` - (Optional) Type: `string`.
* `bootp_lease_time` - (Optional) Type: `string`.
* `bootp_support` - (Optional) Type: `string`.
* `client_mac_limit` - (Optional) Type: `string`. A number, or `unlimited`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `conflict_detection` - (Optional) Type: `bool`.
* `delay_threshold` - (Optional) Type: `string`.
* `dhcp_option_set` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dynamic_lease_identifiers` - (Optional) Type: `string`.
* `dynbootp` - (Read-only) Type: `string`.
* `insert_queue_before` - (Optional) Type: `string`.
* `interface` - (Required) Type: `string`.
* `invalid` - (Read-only) Type: `bool`.
* `lease_script` - (Optional) Type: `string`.
* `lease_time` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `parent_queue` - (Optional) Type: `string`.
* `relay` - (Optional) Type: `string`.
* `server_address` - (Optional) Type: `string`.
* `support_broadband_tr101` - (Optional) Type: `string`. RouterOS `support-broadband-tr101`.
* `support_the_broadband_forum_tr_101` - (Read-only) Type: `bool`.
* `use_framed_as_classless` - (Optional) Type: `bool`.
* `use_radius` - (Optional) Type: `string`.
* `use_reconfigure` - (Optional) Type: `bool`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
