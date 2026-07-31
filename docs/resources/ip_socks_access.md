---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_socks_access"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_socks_access

Manages the RouterOS `/ip/socks/access` menu.

## Example Usage

```terraform
resource "routeros_ip_socks_access" "access_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # action = "replace-me"
  # dst_address = "10.99.0.0/24"
  # dst_port = "443"
  # src_address = "10.99.0.0/24"
  # src_port = "443"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `action` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `dst_address` - (Optional) Type: `string`.
* `dst_port` - (Optional) Type: `string`.
* `src_address` - (Optional) Type: `string`.
* `src_port` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_socks_access.example '*3'

# Named router
terraform import routeros_ip_socks_access.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_socks_access.example 'home/my-resource-name'
```
