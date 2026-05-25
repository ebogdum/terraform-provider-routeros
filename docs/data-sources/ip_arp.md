---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_arp"
description: |-
  RouterOS resource.
---

# Data Source: routeros_ip_arp

Manages the RouterOS `/ip/arp` menu.

## Example Usage

```terraform
data "routeros_ip_arp" "arp_example" {
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
* `address` - (Required) Type: `ip`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `interface` - (Required) Type: `string`.
* `ip_address` - (Optional) Type: `ip`.
* `mac_address` - (Optional) Type: `mac`.
* `mac_ping` - (Optional) Type: `string`.
* `mac_telnet` - (Optional) Type: `string`.
* `make_static` - (Optional) Type: `string`.
* `ping` - (Optional) Type: `string`.
* `published` - (Optional) Type: `bool`.
* `telnet` - (Optional) Type: `string`.
* `torch` - (Optional) Type: `string`.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.
* `bridge_port` - Type: `string`.
* `complete` - Type: `bool`.
* `dhcp` - Type: `bool`.
* `dynamic` - Type: `bool`.
* `host_name` - Type: `string`.
* `invalid` - Type: `bool`.
* `status` - Type: `string`.
* `vrf` - Type: `string`.

