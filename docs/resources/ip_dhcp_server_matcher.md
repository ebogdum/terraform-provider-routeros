---
page_title: "RouterOS: routeros_ip_dhcp_server_matcher"
description: |-
  matching-type field has version-specific accepted values; needs hand-tuning per ROS.
---

# Resource: routeros_ip_dhcp_server_matcher

matching-type field has version-specific accepted values; needs hand-tuning per ROS.

## Example Usage

```terraform
resource "routeros_ip_dhcp_server_matcher" "matcher_example" {
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
terraform import routeros_ip_dhcp_server_matcher.example '*3'

# Named router
terraform import routeros_ip_dhcp_server_matcher.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_dhcp_server_matcher.example 'home/my-resource-name'
```
