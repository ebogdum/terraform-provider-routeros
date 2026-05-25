---
subcategory: "Interfaces"
page_title: "RouterOS: routeros_interface_wifi"
description: |-
  WiFi virtual interface needs exactly one of radio-mac or master-interface. Skipped -- requires WiFi-enabled hardware.
---

# Resource: routeros_interface_wifi

WiFi virtual interface needs exactly one of radio-mac or master-interface. Skipped -- requires WiFi-enabled hardware.

## Example Usage

```terraform
resource "routeros_interface_wifi" "wifi_example" {
  # router = "my-router"  # which router to target; omit for the default
  comment = "managed by terraform"
  disabled = false

  # Optional attributes (uncomment as needed):
  # name = "tf-example"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `name` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_interface_wifi.example '*3'

# Named router
terraform import routeros_interface_wifi.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_interface_wifi.example 'home/my-resource-name'
```
