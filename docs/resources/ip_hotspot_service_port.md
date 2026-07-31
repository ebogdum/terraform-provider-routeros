---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_hotspot_service_port"
description: |-
  Mirrors RouterOS /ip/hotspot/service-port.
---

# Resource: routeros_ip_hotspot_service_port

Mirrors RouterOS `/ip/hotspot/service-port`.

## Example Usage

```terraform
resource "routeros_ip_hotspot_service_port" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # disabled = true
  # name = "replace-me"
  # ports = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `disabled` - (Optional) Type: `bool`. RouterOS `disabled`.
* `name` - (Optional) Type: `string`. RouterOS `name`.
* `ports` - (Optional) Type: `int`. RouterOS `ports`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
terraform import routeros_ip_hotspot_service_port.example 'home::*3'
```
