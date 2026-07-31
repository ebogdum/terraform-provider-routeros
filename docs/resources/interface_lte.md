---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_lte"
description: |-
  LTE interfaces are physical-device backed; skipped on virtual devices.
---

# Resource: routeros_interface_lte

LTE interfaces are physical-device backed; skipped on virtual devices.

## Example Usage

```terraform
resource "routeros_interface_lte" "lte_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `allow_roaming` - (Optional) Type: `string`. RouterOS `allow-roaming`.
* `apn_profiles` - (Optional) Type: `string`. RouterOS `apn-profiles`.
* `band` - (Optional) Type: `string`. RouterOS `band`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `modem_init` - (Optional) Type: `string`. RouterOS `modem-init`.
* `mtu` - (Optional) Type: `string`. RouterOS `mtu`.
* `name` - (Optional) Type: `string`. RouterOS `name`.
* `network_mode` - (Optional) Type: `string`. RouterOS `network-mode`.
* `nr_band` - (Optional) Type: `string`. RouterOS `nr-band`.
* `operator` - (Optional) Type: `string`. RouterOS `operator`.
* `pin` - (Optional) Type: `string`. RouterOS `pin`.
* `sms_protocol` - (Optional) Type: `string`. RouterOS `sms-protocol`.
* `sms_read` - (Optional) Type: `string`. RouterOS `sms-read`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_lte.example '*3'

# Named router
terraform import routeros_interface_lte.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_lte.example 'home/my-resource-name'
```
