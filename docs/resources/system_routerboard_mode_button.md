---
subcategory: "System"
page_title: "RouterOS: routeros_system_routerboard_mode_button"
description: |-
  Binds the mode button on a RouterBOARD to a script or built-in event.
---

# Resource: routeros_system_routerboard_mode_button

Manages RouterOS `/system/routerboard/mode-button` — what the mode button runs
when it is held.

The script named by `on_event` is a separate object managed as
`routeros_system_script`. This resource manages only the *binding*: without it a
device rebuilt from Terraform gets its scripts back but comes up with the button
unbound.

The mode button is commonly bound to `dark-mode` to turn the LEDs off.

## Example Usage

```terraform
resource "routeros_system_routerboard_mode_button" "this" {
  # router = "my-router"  # which router to target; omit for the default
  enabled   = true
  hold_time = "0s..1m"
  on_event  = "dark-mode"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `enabled` - (Optional) Type: `bool`. Whether the button fires `on_event` at all.
* `hold_time` - (Optional) Type: `string`. How long the button must be held for the event to fire, as a RouterOS range -- e.g. `0s..1m`. Not a plain duration.
* `on_event` - (Optional) Type: `string`. Name of the script to run, or a built-in event such as `dark-mode` or `wps-accept`. Empty means nothing is bound.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_system_routerboard_mode_button.this 'home'
```
