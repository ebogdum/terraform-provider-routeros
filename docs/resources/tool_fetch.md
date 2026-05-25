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

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `mode` - (Optional) Type: `string`. Default: `http`.
* `url` - (Required) Type: `string`. Default: `http://127.0.0.1/`.

