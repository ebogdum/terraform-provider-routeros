---
subcategory: "System"
page_title: "RouterOS: routeros_system_logging_action"
description: |-
  RouterOS rejects hyphens AND underscores in action names on some 7.x versions; not portable.
---

# Resource: routeros_system_logging_action

RouterOS rejects hyphens AND underscores in action names on some 7.x versions; not portable.

## Example Usage

```terraform
resource "routeros_system_logging_action" "action_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # disk_file_count = 0
  # disk_file_name = "replace-me"
  # disk_lines_per_file = 0
  # disk_stop_on_full = false
  # memory_lines = 0
  # memory_stop_on_full = false
  # name = "tf-example"
  # remember = false
  # remote = "replace-me"
  # remote_log_format = "replace-me"
  # remote_port = "443"
  # remote_protocol = "replace-me"
  # src_address = "10.99.0.0/24"
  # target = "replace-me"
  # vrf = "main"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `add_topics_string` - (Optional) Type: `string`. RouterOS `add-topics-string`.
* `cef_event_delimiter` - (Optional) Type: `string`. RouterOS `cef-event-delimiter`.
* `check_certificate` - (Optional) Type: `string`. RouterOS `check-certificate`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `default` - (Read-only) Type: `bool`.
* `disk_file_count` - (Optional) Type: `int`.
* `disk_file_name` - (Optional) Type: `string`.
* `disk_lines_per_file` - (Optional) Type: `int`.
* `disk_stop_on_full` - (Optional) Type: `bool`.
* `email_cc` - (Optional) Type: `string`. RouterOS `email-cc`.
* `email_start_tls` - (Optional) Type: `string`. RouterOS `email-start-tls`.
* `email_to` - (Optional) Type: `string`. RouterOS `email-to`.
* `memory_lines` - (Optional) Type: `int`.
* `memory_stop_on_full` - (Optional) Type: `bool`.
* `name` - (Optional) Type: `string`.
* `remember` - (Optional) Type: `bool`.
* `remote` - (Optional) Type: `string`.
* `remote_log_format` - (Optional) Type: `string`.
* `remote_port` - (Optional) Type: `int`.
* `remote_protocol` - (Optional) Type: `string`.
* `script` - (Optional) Type: `string`. RouterOS `script`.
* `src_address` - (Optional) Type: `string`.
* `syslog_facility` - (Optional) Type: `string`. RouterOS `syslog-facility`.
* `syslog_severity` - (Optional) Type: `string`. RouterOS `syslog-severity`.
* `syslog_time_format` - (Optional) Type: `string`. RouterOS `syslog-time-format`.
* `target` - (Optional) Type: `string`.
* `vrf` - (Optional) Type: `string`.

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_system_logging_action.example '*3'

# Named router
terraform import routeros_system_logging_action.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_system_logging_action.example 'home/my-resource-name'
```
