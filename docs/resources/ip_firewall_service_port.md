---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_firewall_service_port"
description: |-
  Mirrors RouterOS /ip/firewall/service-port.
---

# Resource: routeros_ip_firewall_service_port

Mirrors RouterOS `/ip/firewall/service-port`.

## Example Usage

```terraform
resource "routeros_ip_firewall_service_port" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # disabled = true
  # name = "replace-me"
  # ports = 0
  # sip_direct_media = true
  # sip_timeout = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `disabled` - (Optional) Type: `bool`. RouterOS `disabled`.
* `name` - (Optional) Type: `string`. RouterOS `name`.
* `ports` - (Optional) Type: `string`. RouterOS `ports`.
* `sip_direct_media` - (Optional) Type: `bool`. RouterOS `sip-direct-media`.
* `sip_timeout` - (Optional) Type: `string`. RouterOS `sip-timeout`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
terraform import routeros_ip_firewall_service_port.example 'home::*3'
```
