---
subcategory: "System & misc"
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

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `channel` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ip_port` - (Optional) Type: `string`. RouterOS `ip-port`.
* `local_address` - (Optional) Type: `string`.
* `log_file` - (Optional) Type: `string`. RouterOS `log-file`.
* `port` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `string`.
* `remote_addresses` - (Optional) Type: `string`. RouterOS `remote-addresses`.

## Attribute Reference

* `id` - RouterOS internal .id.


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
