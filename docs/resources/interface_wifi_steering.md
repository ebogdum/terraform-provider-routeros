---
page_title: "RouterOS: routeros_interface_wifi_steering"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_wifi_steering

Manages the RouterOS `/interface/wifi/steering` menu.

## Example Usage

```terraform
resource "routeros_interface_wifi_steering" "steering_example" {
  # router = "my-router"  # which router to target; omit for the default
  name = "tf-example"

  comment = "managed by terraform"
  disabled = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `name` - (Required) Type: `string`. Default: `tf_acc_wstr`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wifi_steering.example '*3'

# Named router
terraform import routeros_interface_wifi_steering.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wifi_steering.example 'home/my-resource-name'
```
