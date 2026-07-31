---
subcategory: "Users & RADIUS"
page_title: "RouterOS: routeros_user_manager_router"
description: |-
  RouterOS resource.
---

# Resource: routeros_user_manager_router

Manages the RouterOS `/user-manager/router` menu.

## Example Usage

```terraform
resource "routeros_user_manager_router" "router_example" {
  # router = "my-router"  # which router to target; omit for the default
  disabled = false

  # Optional attributes (uncomment as needed):
  # address = "replace-me"
  # coa_port = "443"
  # name = "tf-example"
  # protocol = "replace-me"
  # shared_secret = "REDACTED"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `address` - (Optional) Type: `string`.
* `coa_port` - (Optional) Type: `string`.
* `disabled` - (Optional) Type: `string`.
* `name` - (Optional) Type: `string`.
* `protocol` - (Optional) Type: `string`.
* `shared_secret` - (Optional) Type: `string`. **Sensitive.**

## Attribute Reference

* `id` - RouterOS internal .id.


## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_user_manager_router.example '*3'

# Named router
terraform import routeros_user_manager_router.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_user_manager_router.example 'home/my-resource-name'
```
