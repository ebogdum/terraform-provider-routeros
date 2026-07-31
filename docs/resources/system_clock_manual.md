---
subcategory: "System"
page_title: "RouterOS: routeros_system_clock_manual"
description: |-
  Singleton for manual time configuration. Curl confirms POST /set with an
---

# Resource: routeros_system_clock_manual

Singleton for manual time configuration. Curl confirms POST /set with an
empty body returns 200, but the acc test framework times out — likely
because RouterOS adjusts the clock to a value that breaks TLS validation
on the next request. Skipped from automated acc tests.


## Example Usage

```terraform
resource "routeros_system_clock_manual" "manual_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # dst_delta = "replace-me"
  # dst_end = "replace-me"
  # dst_start = "replace-me"
  # time_zone = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `dst_delta` - (Optional) Type: `string`.
* `dst_end` - (Optional) Type: `string`.
* `dst_start` - (Optional) Type: `string`.
* `time_zone` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_clock_manual.this 'home'
```
