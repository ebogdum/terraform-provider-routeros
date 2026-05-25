---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_hotspot"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_hotspot

Manages the RouterOS `/ip/hotspot` menu.

## Example Usage

```terraform
resource "routeros_ip_hotspot" "hotspot_example" {
  # router = "my-router"  # which router to target; omit for the default
  interface = "ether1"
  name = "tf-example"

  disabled = false

  # Optional attributes (uncomment as needed):
  # keepalive_timeout = "replace-me"
  # profile = "replace-me"
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `disabled` - (Optional) Type: `bool`. Whether the entry is disabled.
* `interface` - (Required) Type: `string`.
* `keepalive_timeout` - (Optional) Type: `string`.
* `name` - (Required) Type: `string`. Descriptive name of the profile. Default: `tf_acc_hotspot`.
* `profile` - (Optional) Type: `string`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_hotspot.example '*3'

# Named router
terraform import routeros_ip_hotspot.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_hotspot.example 'home/my-resource-name'
```
