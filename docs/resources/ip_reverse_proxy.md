---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_reverse_proxy"
description: |-
  Reverse proxy listener — properties differ across ROS versions; safe defaults rejected on 7.x. Skipped.
---

# Resource: routeros_ip_reverse_proxy

Reverse proxy listener — properties differ across ROS versions; safe defaults rejected on 7.x. Skipped.

## Example Usage

```terraform
resource "routeros_ip_reverse_proxy" "reverse_proxy_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `certificate` - (Optional) Type: `string`. RouterOS `certificate`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `ip_address` - (Optional) Type: `string`. RouterOS `ip-address`.
* `port` - (Optional) Type: `string`. RouterOS `port`.
* `sni` - (Optional) Type: `string`. RouterOS `sni`.
* `vrf` - (Optional) Type: `string`. RouterOS `vrf`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_reverse_proxy.example '*3'

# Named router
terraform import routeros_ip_reverse_proxy.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_reverse_proxy.example 'home/my-resource-name'
```
