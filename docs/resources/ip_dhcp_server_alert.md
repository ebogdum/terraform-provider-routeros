---
subcategory: "DHCP"
page_title: "RouterOS: routeros_ip_dhcp_server_alert"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_dhcp_server_alert

Manages the RouterOS `/ip/dhcp-server/alert` menu.

## Example Usage

```terraform
resource "routeros_ip_dhcp_server_alert" "alert_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"

  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # alert_timeout = "3600"
  # on_alert = "replace-me"
  # reset_alert = "replace-me"
  # valid_servers = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `alert_timeout` - (Optional) Type: `string`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.
* `on_alert` - (Optional) Type: `string`.
* `reset_alert` - (Read-only) Type: `string`.
* `unknown_servers` - (Read-only) Type: `string`.
* `valid_server` - (Optional) Type: `string`. RouterOS `valid-server`.
* `valid_servers` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_dhcp_server_alert.example '*3'

# Named router
terraform import routeros_ip_dhcp_server_alert.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_dhcp_server_alert.example 'home/my-resource-name'
```
