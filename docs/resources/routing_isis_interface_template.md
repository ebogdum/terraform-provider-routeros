---
subcategory: "ISIS"
page_title: "RouterOS: routeros_routing_isis_interface_template"
description: |-
  References an existing isis instance; auto-test can't synthesise.
---

# Resource: routeros_routing_isis_interface_template

References an existing isis instance; auto-test can't synthesise.

## Example Usage

```terraform
resource "routeros_routing_isis_interface_template" "interface_template_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_routing_isis_interface_template.example '*3'

# Named router
terraform import routeros_routing_isis_interface_template.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_routing_isis_interface_template.example 'home/my-resource-name'
```
