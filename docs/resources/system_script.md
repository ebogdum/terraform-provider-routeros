---
subcategory: "Scripts & scheduler"
page_title: "RouterOS: routeros_system_script"
description: |-
  RouterOS resource.
---

# Resource: routeros_system_script

Manages the RouterOS `/system/script` menu.

## Example Usage

```terraform
resource "routeros_system_script" "script_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"
  source = ":put \"hello\""

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # don_t_require_permissions = false
  # policy = []
  # run_script = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `don_t_require_permissions` - (Optional) Type: `bool`.
* `name` - (Required) Type: `string`. Default: `tf-acc-script`.
* `policy` - (Optional) Type: `list`.
* `run_script` - (Optional) Type: `string`.
* `source` - (Required) Type: `string`. Default: `:put "hello"`.

## Attribute Reference

In addition to the arguments above, the following attributes are exported:

* `id` - Provider-managed identifier (`<router>:<menu-path>` for singletons, RouterOS `.id` for collection rows).
* `invalid` - Type: `bool`.
* `last_time_started` - Type: `string`.
* `owner` - Type: `string`.
* `run_count` - Type: `int`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_system_script.example '*3'

# Named router
terraform import routeros_system_script.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_system_script.example 'home/my-resource-name'
```
