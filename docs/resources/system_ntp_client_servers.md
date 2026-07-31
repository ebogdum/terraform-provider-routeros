---
subcategory: "NTP"
page_title: "RouterOS: routeros_system_ntp_client_servers"
description: |-
  NTP server list — accepts add but validator differs per ROS. Skipped from acc tests.
---

# Resource: routeros_system_ntp_client_servers

NTP server list — accepts add but validator differs per ROS. Skipped from acc tests.

## Example Usage

```terraform
resource "routeros_system_ntp_client_servers" "servers_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # auth_key = "REDACTED"
  # iburst = true
  # keys = "replace-me"
  # max_poll = 10
  # min_poll = 6
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Optional) Type: `string`.
* `auth_key` - (Optional) Type: `string`. **Sensitive.**
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`.
* `dynamic` - (Read-only) Type: `bool`.
* `iburst` - (Optional) Type: `bool`.
* `keys` - (Read-only) Type: `string`.
* `max_poll` - (Optional) Type: `int`.
* `min_poll` - (Optional) Type: `int`.
* `resolved_address` - (Read-only) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_system_ntp_client_servers.example '*3'

# Named router
terraform import routeros_system_ntp_client_servers.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_system_ntp_client_servers.example 'home/my-resource-name'
```
