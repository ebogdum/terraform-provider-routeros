---
subcategory: "IP"
page_title: "RouterOS: routeros_ip_dhcp_server_option"
description: |-
  RouterOS resource.
---

# Resource: routeros_ip_dhcp_server_option

Manages the RouterOS `/ip/dhcp-server/option` menu.

## Example Usage

```terraform
resource "routeros_ip_dhcp_server_option" "option_example" {
  # router = "my-router"  # which router to target; omit for the default
  code = 60
  name = "tf-example"
  value = "'tf-acc'"

  comment = "managed by terraform"

  # Optional attributes (uncomment as needed):
  # force = false
}
```

## Argument Reference

This resource supports the following arguments:

* `router` - (Optional) Name of the router in the provider's `routers` map to target. Omit to use the default router.
* `code` - (Required) Type: `int`. Default: `60`.
* `comment` - (Optional) Type: `string`. Free-form comment.
* `force` - (Optional) Type: `bool`.
* `name` - (Required) Type: `string`. Default: `tf_acc_opt`.
* `value` - (Required) Type: `string`. Default: `'tf-acc'`.

## Import

Rows are imported by RouterOS `.id`, optionally prefixed by the router name:

```sh
# Default router, .id = *3
terraform import routeros_ip_dhcp_server_option.example '*3'

# Named router
terraform import routeros_ip_dhcp_server_option.example 'home/*3'

# By natural key (the resource's `name` attribute, when present)
terraform import routeros_ip_dhcp_server_option.example 'home/my-resource-name'
```
