---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_interworking"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_interworking

Manages the RouterOS `/interface/wifi/interworking` menu.

## Example Usage

```terraform
resource "routeros_interface_wifi_interworking" "interworking_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # x3gpp_info = "replace-me"
  # x3gpp_info_raw = "replace-me"
  # authentication_types = "replace-me"
  # connection_capabilities = "replace-me"
  # dgaf = "replace-me"
  # domain_names = "replace-me"
  # esr = "replace-me"
  # hessid = "replace-me"
  # hotspot_2_0 = "replace-me"
  # internet = "replace-me"
  # ipv4_availability = "replace-me"
  # ipv6_availability = "replace-me"
  # network_type = "replace-me"
  # operational_classes = "replace-me"
  # operator_names = "replace-me"
  # realms = "replace-me"
  # realms_raw = "replace-me"
  # roaming_ois = "replace-me"
  # uesa = "replace-me"
  # venue = "replace-me"
  # venue_names = "replace-me"
  # wan_at_capacity = "replace-me"
  # wan_downlink = "replace-me"
  # wan_downlink_load = "replace-me"
  # wan_measurement_duration = "replace-me"
  # wan_status = "replace-me"
  # wan_symmetric = "replace-me"
  # wan_uplink = "replace-me"
  # wan_uplink_load = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `x3gpp_info` - (Optional) Type: `string`.
* `x3gpp_info_raw` - (Optional) Type: `string`.
* `authentication_types` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `connection_capabilities` - (Optional) Type: `string`.
* `dgaf` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `domain_names` - (Optional) Type: `string`.
* `esr` - (Optional) Type: `string`.
* `hessid` - (Optional) Type: `string`.
* `hotspot_2_0` - (Optional) Type: `string`.
* `internet` - (Optional) Type: `string`.
* `ipv4_availability` - (Optional) Type: `string`.
* `ipv6_availability` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_wiw`.
* `network_type` - (Optional) Type: `string`.
* `operational_classes` - (Optional) Type: `string`.
* `operator_names` - (Optional) Type: `string`.
* `realms` - (Optional) Type: `string`.
* `realms_raw` - (Optional) Type: `string`.
* `roaming_ois` - (Optional) Type: `string`.
* `uesa` - (Optional) Type: `string`.
* `venue` - (Optional) Type: `string`.
* `venue_names` - (Optional) Type: `string`.
* `wan_at_capacity` - (Optional) Type: `string`.
* `wan_downlink` - (Optional) Type: `string`.
* `wan_downlink_load` - (Optional) Type: `string`.
* `wan_measurement_duration` - (Optional) Type: `string`.
* `wan_status` - (Optional) Type: `string`.
* `wan_symmetric` - (Optional) Type: `string`.
* `wan_uplink` - (Optional) Type: `string`.
* `wan_uplink_load` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wifi_interworking.example '*3'

# Named router
terraform import routeros_interface_wifi_interworking.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wifi_interworking.example 'home/my-resource-name'
```
