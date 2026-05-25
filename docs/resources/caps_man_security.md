---
subcategory: "System & misc"
page_title: "RouterOS: routeros_caps_man_security"
description: |-
  RouterOS resource.
---

# Resource: routeros_caps_man_security

Manages the RouterOS `/caps-man/security` menu.

## Example Usage

```terraform
resource "routeros_caps_man_security" "security_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # encryption = "replace-me"
  # passphrase = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `encryption` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Default: `ros_audit_20260523213235_12`.
* `passphrase` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_caps_man_security.example '*3'

# Named router
terraform import routeros_caps_man_security.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_caps_man_security.example 'home/my-resource-name'
```
