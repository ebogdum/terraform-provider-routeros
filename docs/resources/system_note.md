---
page_title: "RouterOS: routeros_system_note"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_note

Manages the RouterOS `/system/note` menu.

## Example Usage

```terraform
resource "routeros_system_note" "note_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # note = "replace-me"
  # show_at_cli_login = false
  # show_at_login = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `note` - (Optional) Type: `string`.
* `show_at_cli_login` - (Optional) Type: `bool`.
* `show_at_login` - (Optional) Type: `bool`.

## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_note.this 'home'
```
