---
subcategory: "IPv6"
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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `dhcp_option` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.
* `lease_time` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_dhcp6`.

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
