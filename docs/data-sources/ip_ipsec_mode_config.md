---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_ipsec_mode_config"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_ipsec_mode_config

Manages the RouterOS `/ip/ipsec/mode-config` menu.

## Example Usage

```terraform
data "routeros_ip_ipsec_mode_config" "mode_config_example" {
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
* `address` - (Optional) Type: `ip`.
* `address_pool` - (Optional) Type: `string`. Default: `4.294967295e+09`.
* `address_prefix_length` - (Optional) Type: `int`. Default: `24`.
* `connection_mark` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `tf_acc_modecfg`.
* `responder` - (Optional) Type: `bool`.
* `split_dns` - (Optional) Type: `string`.
* `split_include` - (Optional) Type: `string`.
* `src_address_list` - (Optional) Type: `string`.
* `static_dns` - (Optional) Type: `string`.
* `system_dns` - (Optional) Type: `bool`.
* `use_responder_dns` - (Optional) Type: `enum(no|yes|exclusively)`. Default: `2`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `default` - Type: `bool`.

