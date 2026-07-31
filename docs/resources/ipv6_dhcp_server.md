---
subcategory: "DHCP"
page_title: "RouterOS: routeros_ipv6_dhcp_server"
description: |-
  RouterOS resource.
---

# Resource: routeros_ipv6_dhcp_server

Manages the RouterOS `/ipv6/dhcp-server` menu.

## Example Usage

```terraform
resource "routeros_ipv6_dhcp_server" "dhcp_server_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # dhcp_option = "replace-me"
  # lease_time = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address_lists` - (Optional) Type: `string`. RouterOS `address-lists`.
* `address_pool` - (Optional) Type: `string`. RouterOS `address-pool`.
* `allow_dual_stack_queue` - (Optional) Type: `string`. RouterOS `allow-dual-stack-queue`.
* `binding_script` - (Optional) Type: `string`. RouterOS `binding-script`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dhcp_option` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ignore_ia_na_bindings` - (Optional) Type: `string`. RouterOS `ignore-ia-na-bindings`.
* `insert_queue_before` - (Optional) Type: `string`. RouterOS `insert-queue-before`.
* `interface` - (Required) Type: `string`.
* `lease_time` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`.
* `parent_queue` - (Optional) Type: `string`. RouterOS `parent-queue`.
* `preference` - (Optional) Type: `string`. RouterOS `preference`.
* `prefix_pool` - (Optional) Type: `string`. RouterOS `prefix-pool`.
* `rapid_commit` - (Optional) Type: `string`. RouterOS `rapid-commit`.
* `route_distance` - (Optional) Type: `string`. RouterOS `route-distance`.
* `use_radius` - (Optional) Type: `string`. RouterOS `use-radius`.
* `use_reconfigure` - (Optional) Type: `string`. RouterOS `use-reconfigure`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_dhcp_server.example '*3'

# Named router
terraform import routeros_ipv6_dhcp_server.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_dhcp_server.example 'home/my-resource-name'
```
