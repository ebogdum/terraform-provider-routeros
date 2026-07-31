---
subcategory: "Tools"
page_title: "RouterOS: routeros_tool_romon"
description: |-
  RouterOS resource.
---

# Resource: routeros_tool_romon

Manages the RouterOS `/tool/romon` menu.

## Example Usage

```terraform
resource "routeros_tool_romon" "romon_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # enabled = false
  # secrets = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `enabled` - (Optional) Type: `bool`.
* `secrets` - (Optional) Type: `string`. **Sensitive.**

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_tool_romon.this 'home'
```
