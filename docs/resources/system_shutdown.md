---
subcategory: "System"
page_title: "RouterOS: routeros_system_shutdown"
description: |-
  Powers off the router. Verified working against a CHR VM. Skipped from the
---

# Resource: routeros_system_shutdown

Powers off the router. Verified working against a CHR VM. Skipped from the
general suite because it terminates the test target.


## Example Usage

```terraform
resource "routeros_system_shutdown" "shutdown_example" {
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

