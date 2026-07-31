---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_kid_control"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_kid_control

Manages the RouterOS `/ip/kid-control` menu.

## Example Usage

```terraform
resource "routeros_ip_kid_control" "kid_control_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # name = "tf-example"
  # rate_limit = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `fri` - (Optional) Type: `string`. RouterOS `fri`.
* `mon` - (Optional) Type: `string`. RouterOS `mon`.
* `name` - (Optional) Type: `string`. Name of the Kid's profile
* `rate_limit` - (Optional) Type: `string`. The maximum available data rate for flow
* `sat` - (Optional) Type: `string`. RouterOS `sat`.
* `sun` - (Optional) Type: `string`. RouterOS `sun`.
* `thu` - (Optional) Type: `string`. RouterOS `thu`.
* `tue` - (Optional) Type: `string`. RouterOS `tue`.
* `tur_fri` - (Optional) Type: `string`. RouterOS `tur-fri`.
* `tur_mon` - (Optional) Type: `string`. RouterOS `tur-mon`.
* `tur_sat` - (Optional) Type: `string`. RouterOS `tur-sat`.
* `tur_sun` - (Optional) Type: `string`. RouterOS `tur-sun`.
* `tur_thu` - (Optional) Type: `string`. RouterOS `tur-thu`.
* `tur_tue` - (Optional) Type: `string`. RouterOS `tur-tue`.
* `tur_wed` - (Optional) Type: `string`. RouterOS `tur-wed`.
* `wed` - (Optional) Type: `string`. RouterOS `wed`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_kid_control.example '*3'

# Named router
terraform import routeros_ip_kid_control.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_kid_control.example 'home/my-resource-name'
```
