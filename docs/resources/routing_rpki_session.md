---
page_title: "RouterOS: routeros_routing_rpki_session"
description: |-
  Discovered; needs rpki backend
---

# Resource: routeros_routing_rpki_session

Discovered; needs rpki backend

## Example Usage

```terraform
resource "routeros_routing_rpki_session" "session_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # expires = "1h"
  # group = "replace-me"
  # port = "443"
  # serial = 0
  # session = 0
  # state = "idle"
  # version = 0
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `address` - (Optional) Type: `string`.
* `expires` - (Optional) Type: `duration`.
* `group` - (Optional) Type: `string`.
* `port` - (Optional) Type: `int`.
* `serial` - (Optional) Type: `int`.
* `session` - (Optional) Type: `int`.
* `state` - (Optional) Type: `enum(idle|connecting|prepare|loading|sync|error)`.
* `version` - (Optional) Type: `int`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_rpki_session.example '*3'

# Named router
terraform import routeros_routing_rpki_session.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_rpki_session.example 'home/my-resource-name'
```
