---
subcategory: "System & misc"
page_title: "RouterOS: routeros_caps_man_rates"
description: |-
  RouterOS resource.
---

# Resource: routeros_caps_man_rates

Manages the RouterOS `/caps-man/rates` menu.

## Example Usage

```terraform
resource "routeros_caps_man_rates" "rates_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `name` - (Required) Type: `string`. Default: `ros_audit_20260523213235_9`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_caps_man_rates.example '*3'

# Named router
terraform import routeros_caps_man_rates.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_caps_man_rates.example 'home/my-resource-name'
```
