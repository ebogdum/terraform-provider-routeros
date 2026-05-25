---
page_title: "RouterOS: routeros_ipv6_dhcp_relay"
description: |-
  server address validator rejects literal addresses on this ROS version. Skipped.
---

# Resource: routeros_ipv6_dhcp_relay

server address validator rejects literal addresses on this ROS version. Skipped.

## Example Usage

```terraform
resource "routeros_ipv6_dhcp_relay" "dhcp_relay_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ipv6_dhcp_relay.example '*3'

# Named router
terraform import routeros_ipv6_dhcp_relay.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ipv6_dhcp_relay.example 'home/my-resource-name'
```
