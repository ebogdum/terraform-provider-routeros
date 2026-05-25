---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_netwatch"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_netwatch

Manages the RouterOS `/tool/netwatch` menu.

## Example Usage

```terraform
resource "routeros_tool_netwatch" "netwatch_example" {
  # router = "my-router"  # which router to target; omit for the default
  host = "1.1.1.1"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # certificate = "replace-me"
  # dns_server = "1.1.1.1,8.8.8.8"
  # interval = "1m"
  # name = "tf-example"
  # port = "443"
  # src_address = "10.99.0.0/24"
  # timeout = "replace-me"
  # ttl = "replace-me"
  # type = "icmp"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `certificate` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dns_server` - (Optional) Type: `string`.
* `host` - (Required) Type: `string`. Default: `127.0.0.1`.
* `interval` - (Optional) Type: `duration`. Default: `1m`.
* `name` - (Optional) Type: `string`.
* `port` - (Optional) Type: `string`.
* `src_address` - (Optional) Type: `string`.
* `timeout` - (Optional) Type: `string`.
* `ttl` - (Optional) Type: `string`.
* `type` - (Optional) Type: `string`. Default: `icmp`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_tool_netwatch.example '*3'

# Named router
terraform import routeros_tool_netwatch.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_tool_netwatch.example 'home/my-resource-name'
```
