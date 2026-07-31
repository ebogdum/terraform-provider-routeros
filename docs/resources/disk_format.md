---
subcategory: "Storage"
page_title: "RouterOS: routeros_disk_format"
description: |-
  Formats a block device. Requires the .id of a real /disk entry, which an
---

# Resource: routeros_disk_format

Formats a block device. Requires the .id of a real /disk entry, which an
automated test cannot create. Skipped.


## Example Usage

```terraform
resource "routeros_disk_format" "format_example" {
  # router = "my-router"  # which router to target; omit for the default
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `output` - (Read-only) Type: `list`. Server response rows.
* `params` - (Optional) Type: `map`. Extra parameters forwarded to RouterOS verbatim. Keys with dots are allowed. Example: { ca = "my-ca", name = "new-cert" }.
* `target_id` - (Optional) Type: `string`. RouterOS .id of the row this action targets. Required by per-row actions (e.g. /certificate/sign, /interface/reset-counters, /disk/format).
* `trigger` - (Optional) Type: `string`. Change to force re-execution.

## Attribute Reference

* `id` - Hash of the inputs that produced this run.

