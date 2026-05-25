---
subcategory: "Port"
page_title: "RouterOS: routeros_port_remote_access"
description: |-
  RouterOS resource.
---

# Resource: routeros_port_remote_access

Manages the RouterOS `/port/remote-access` menu.

## Example Usage

```terraform
resource "routeros_port_remote_access" "remote_access_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # channel = "replace-me"
  # local_address = "10.99.0.1"
  # port = "443"
  # protocol = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `channel` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `local_address` - (Optional) Type: `string`.
* `port` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_port_remote_access.example '*3'

# Named router
terraform import routeros_port_remote_access.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_port_remote_access.example 'home/my-resource-name'
```
