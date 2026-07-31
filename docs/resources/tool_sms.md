---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_sms"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_sms

Manages the RouterOS `/tool/sms` menu.

## Example Usage

```terraform
resource "routeros_tool_sms" "sms_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # allowed_number = "replace-me"
  # channel = 0
  # polling = false
  # port = "443"
  # receive_enabled = false
  # remove_sent_sms_after_send = false
  # secret = "REDACTED"
  # sim_pin = "replace-me"
  # sms_storage = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `allowed_number` - (Optional) Type: `string`.
* `channel` - (Optional) Type: `int`.
* `last_ussd` - (Optional) Type: `string`.
* `polling` - (Optional) Type: `bool`.
* `port` - (Optional) Type: `string`.
* `receive_enabled` - (Optional) Type: `bool`.
* `remove_sent_sms_after_send` - (Optional) Type: `bool`.
* `secret` - (Optional) Type: `string`. **Sensitive.**
* `sim_pin` - (Optional) Type: `string`. **Sensitive.**
* `sms_storage` - (Optional) Type: `string`.
* `status` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_tool_sms.this 'home'
```
