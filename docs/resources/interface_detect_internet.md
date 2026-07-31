---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_detect_internet"
description: |-
  RouterOS resource.
---

# Resource: routeros_interface_detect_internet

Manages the RouterOS `/interface/detect-internet` menu.

## Example Usage

```terraform
resource "routeros_interface_detect_internet" "detect_internet_example" {
  # router = "my-router"  # which router to target; omit for the default

  # Optional attributes (uncomment as needed):
  # detect_interface_list = "replace-me"
  # internet_interface_list = "replace-me"
  # lan_interface_list = "replace-me"
  # request_interval = "1h"
  # wan_interface_list = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Type: `string`. Name of the router (key in provider's `routers` map). Omit to use the default.
* `detect_interface_list` - (Optional) Type: `string`.
* `internet_interface_list` - (Optional) Type: `string`.
* `lan_interface_list` - (Optional) Type: `string`.
* `request_interval` - (Optional) Type: `string`.
* `wan_interface_list` - (Optional) Type: `string`.

## Attribute Reference

* `id` - Stable identifier (the singleton's menu path, optionally namespaced by router).


## Import

Singletons are imported by router name:

```sh
terraform import routeros_interface_detect_internet.this 'home'
```
