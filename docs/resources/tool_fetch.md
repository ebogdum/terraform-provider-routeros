---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_fetch"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_fetch

Manages the RouterOS `/tool/fetch` menu.

## Example Usage

```terraform
resource "routeros_tool_fetch" "fetch_example" {
  # router = "my-router"  # which router to target; omit for the default
  url = "https://example.com"

  # Optional attributes (uncomment as needed):
  # mode = "http"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `mode` - (Optional) Type: `string`.
* `output` - (Read-only) Type: `list`. Server response rows.
* `params` - (Optional) Type: `map`. Extra parameters forwarded to RouterOS verbatim. Keys with dots are allowed. Example: { ca = "my-ca", name = "new-cert" }.
* `target_id` - (Optional) Type: `string`. RouterOS .id of the row this action targets. Required by per-row actions (e.g. /certificate/sign, /interface/reset-counters, /disk/format).
* `trigger` - (Optional) Type: `string`. Change to force re-execution.
* `url` - (Required) Type: `string`.

## Attribute Reference

* `id` - Hash of the inputs that produced this run.

