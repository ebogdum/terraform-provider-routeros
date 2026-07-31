---
subcategory: "DHCP"
page_title: "RouterOS: routeros_ip_dhcp_server_config"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_dhcp_server_config

Manages the RouterOS `/ip/dhcp-server/config` menu.

## Example Usage

```terraform
resource "routeros_ip_dhcp_server_config" "config_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # accounting = false
  # interim_update = "1h"
  # radius_password = "REDACTED"
  # store_leases_disk = "1h"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `accounting` - (Optional) Type: `bool`.
* `interim_update` - (Optional) Type: `string`.
* `radius_password` - (Optional) Type: `string`. **Sensitive.**
* `store_leases_disk` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_ip_dhcp_server_config.this 'home'
```
