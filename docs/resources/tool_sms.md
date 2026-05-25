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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `allowed_number` - (Optional) Type: `string`.
* `channel` - (Optional) Type: `int`.
* `polling` - (Optional) Type: `bool`.
* `port` - (Optional) Type: `string`.
* `receive_enabled` - (Optional) Type: `bool`.
* `remove_sent_sms_after_send` - (Optional) Type: `bool`.
* `secret` - (Optional) Type: `string`. **Sensitive.**
* `sim_pin` - (Optional) Type: `string`. **Sensitive.**
* `sms_storage` - (Optional) Type: `string`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `last_ussd` - Type: `string`.
* `status` - Type: `string`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_tool_sms.this 'home'
```
