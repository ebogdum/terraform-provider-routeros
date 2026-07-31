---
subcategory: "IPv6"
page_title: "RouterOS: routeros_ipv6_dhcp_relay_option"
description: |-
  Mirrors RouterOS /ipv6/dhcp-relay/option.
---

# Resource: routeros_ipv6_dhcp_relay_option

Mirrors RouterOS `/ipv6/dhcp-relay/option`.

## Example Usage

```terraform
resource "routeros_ipv6_dhcp_relay_option" "this" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # code = 0
  # name = "replace-me"
  # only_if_mac_available = true
  # value = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `code` - (Optional) Type: `int`. RouterOS `code`.
* `name` - (Optional) Type: `string`. RouterOS `name`.
* `only_if_mac_available` - (Optional) Type: `bool`. RouterOS `only-if-mac-available`.
* `value` - (Optional) Type: `string`. RouterOS `value`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
terraform import routeros_ipv6_dhcp_relay_option.example 'home::*3'
```
