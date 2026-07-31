---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_socksify"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_socksify

Manages the RouterOS `/ip/socksify` menu.

## Example Usage

```terraform
resource "routeros_ip_socksify" "socksify_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # port = "443"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connection_timeout` - (Optional) Type: `string`. RouterOS `connection-timeout`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `name` - (Required) Type: `string`.
* `port` - (Optional) Type: `int`.
* `socks5_password` - (Optional) Type: `string`. RouterOS `socks5-password`. **Sensitive.**
* `socks5_port` - (Optional) Type: `string`. RouterOS `socks5-port`.
* `socks5_server` - (Optional) Type: `string`. RouterOS `socks5-server`.
* `socks5_user` - (Optional) Type: `string`. RouterOS `socks5-user`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_socksify.example '*3'

# Named router
terraform import routeros_ip_socksify.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_socksify.example 'home/my-resource-name'
```
