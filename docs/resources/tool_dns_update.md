---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_dns_update"
description: |-
  Sends an RFC 2136 DDNS update; needs a real DDNS server.
---

# Resource: routeros_tool_dns_update

Sends an RFC 2136 DDNS update; needs a real DDNS server.

## Example Usage

```terraform
resource "routeros_tool_dns_update" "dns_update_example" {
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

