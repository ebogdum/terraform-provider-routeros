---
subcategory: "Wireless"
page_title: "RouterOS: routeros_interface_wifi_interworking"
description: |-
  RouterOS resource.
---

# Data Source: routeros_interface_wifi_interworking

Manages the RouterOS `/interface/wifi/interworking` menu.

## Example Usage

```terraform
data "routeros_interface_wifi_interworking" "interworking_example" {
  # router   = "my-router"  # omit for the default router
  # filter   = { name = "some-name" }
  # proplist = ["name", "address"]
}
```

## Argument Reference

This data source supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to query.
* `filter` - (Optional) Map of field=value pairs to narrow the result set.
* `proplist` - (Optional) List of property names to project; smaller payload.
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

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

