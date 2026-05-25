---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_dhcp_client"
description: |-
  RouterOS resource.
---

# Resource: routeros_ipv6_dhcp_client

Manages the RouterOS `/ipv6/dhcp-client` menu.

## Example Usage

```terraform
resource "routeros_ipv6_dhcp_client" "dhcp_client_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  request = "address"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # default_route_distance = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default_route_distance` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.
* `request` - (Required) Type: `string`. Default: `address`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_dhcp_client.example '*3'

# Named router
terraform import routeros_ipv6_dhcp_client.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_dhcp_client.example 'home/my-resource-name'
```
