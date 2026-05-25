---
page_title: "RouterOS: routeros_log"
description: |-
  RouterOS resource.
---

# Data Source: routeros_log

Manages the RouterOS `/log` menu.

## Example Usage

```terraform
data "routeros_log" "log_example" {
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
* `action` - (Optional) Type: `string`. specifies one of the system default actions or user specified action listed in actions menu.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `extra_info` - (Optional) Type: `string`.
* `message` - (Optional) Type: `string`.
* `prefix` - (Optional) Type: `string`. prefix added at the beginning of log messages.
* `regex` - (Optional) Type: `string`. regex which will be used in order to match or not match message. If the regex is not matched, then even if topic is configured to be logged, but log message does not match regex, action will not be performed.
* `time` - (Optional) Type: `string`.
* `topics` - (Optional) Type: `list`. log all messages that falls into specified topic or list of topics. '!' character can be used before topic to exclude messages falling under this topic. For example, we want to log NTP debug info without too much details: /system logging add topics=ntp,debug,!packet.

## Attribute Reference

This data source exports the following attributes in addition to the arguments above:

* `records` - List of matching rows. Each row has the same fields as the resource above (string-typed), plus the device's `.id`.

