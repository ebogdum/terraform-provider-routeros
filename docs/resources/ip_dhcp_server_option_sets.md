---
subcategory: "DHCP"
page_title: "RouterOS: routeros_ip_dhcp_server_option_sets"
description: |-
  Sets reference existing /ip/dhcp-server/option entries; skipped from automated acc tests.
---

# Resource: routeros_ip_dhcp_server_option_sets

Sets reference existing /ip/dhcp-server/option entries; skipped from automated acc tests.

## Example Usage

```terraform
resource "routeros_ip_dhcp_server_option_sets" "sets_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `name` - (Optional) Type: `string`. Name of the option set, referenced by `/ip/dhcp-server` `dhcp-option-set`.
* `options` - (Optional) Type: `string`. Comma-separated list of `/ip/dhcp-server/option` names in this set.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_dhcp_server_option_sets.example '*3'

# Named router
terraform import routeros_ip_dhcp_server_option_sets.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_dhcp_server_option_sets.example 'home/my-resource-name'
```
